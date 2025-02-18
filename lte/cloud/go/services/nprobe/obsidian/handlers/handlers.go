/*
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package handlers

import (
	"math/rand"
	"net/http"
	"time"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"magma/lte/cloud/go/lte"
	"magma/lte/cloud/go/serdes"
	"magma/lte/cloud/go/services/lte/obsidian/handlers"
	"magma/lte/cloud/go/services/nprobe/obsidian/models"
	"magma/lte/cloud/go/services/nprobe/storage"
	"magma/orc8r/cloud/go/obsidian"
	"magma/orc8r/cloud/go/services/configurator"
	"magma/orc8r/lib/go/merrors"
)

const (
	NetworkProbePath = handlers.ManageNetworkPath + obsidian.UrlSep + "network_probe"

	NetworkProbeTasksPath        = NetworkProbePath + obsidian.UrlSep + "tasks"
	NetworkProbeDestinationsPath = NetworkProbePath + obsidian.UrlSep + "destinations"

	NetworkProbeTaskDetailsPath        = NetworkProbeTasksPath + obsidian.UrlSep + ":task_id"
	NetworkProbeDestinationDetailsPath = NetworkProbeDestinationsPath + obsidian.UrlSep + ":destination_id"
)

func GetHandlers(storage storage.NProbeStorage) []obsidian.Handler {
	ret := []obsidian.Handler{
		{Path: NetworkProbeTasksPath, Methods: obsidian.GET, HandlerFunc: listNetworkProbeTasks},
		{Path: NetworkProbeTasksPath, Methods: obsidian.POST, HandlerFunc: getCreateNetworkProbeTaskHandlerFunc(storage)},
		{Path: NetworkProbeTaskDetailsPath, Methods: obsidian.GET, HandlerFunc: getNetworkProbeTask},
		{Path: NetworkProbeTaskDetailsPath, Methods: obsidian.PUT, HandlerFunc: updateNetworkProbeTask},
		{Path: NetworkProbeTaskDetailsPath, Methods: obsidian.DELETE, HandlerFunc: getDeleteNetworkProbeTaskHandlerFunc(storage)},

		{Path: NetworkProbeDestinationsPath, Methods: obsidian.GET, HandlerFunc: listNetworkProbeDestinations},
		{Path: NetworkProbeDestinationsPath, Methods: obsidian.POST, HandlerFunc: createNetworkProbeDestination},
		{Path: NetworkProbeDestinationDetailsPath, Methods: obsidian.GET, HandlerFunc: getNetworkProbeDestination},
		{Path: NetworkProbeDestinationDetailsPath, Methods: obsidian.PUT, HandlerFunc: updateNetworkProbeDestination},
		{Path: NetworkProbeDestinationDetailsPath, Methods: obsidian.DELETE, HandlerFunc: deleteNetworkProbeDestination},
	}
	return ret
}

func listNetworkProbeTasks(c echo.Context) error {
	networkID, nerr := obsidian.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}

	ents, _, err := configurator.LoadAllEntitiesOfType(
		c.Request().Context(),
		networkID, lte.NetworkProbeTaskEntityType,
		configurator.EntityLoadCriteria{LoadConfig: true},
		serdes.Entity,
	)
	if err == merrors.ErrNotFound {
		return echo.ErrNotFound
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "failed to load existing NetworkProbeTasks"))
	}

	ret := make(map[string]*models.NetworkProbeTask, len(ents))
	for _, ent := range ents {
		ret[ent.Key] = (&models.NetworkProbeTask{}).FromBackendModels(ent)
	}
	return c.JSON(http.StatusOK, ret)
}

func getCreateNetworkProbeTaskHandlerFunc(storage storage.NProbeStorage) echo.HandlerFunc {
	return func(c echo.Context) error {
		networkID, nerr := obsidian.GetNetworkId(c)
		if nerr != nil {
			return nerr
		}
		reqCtx := c.Request().Context()

		payload := &models.NetworkProbeTask{}
		if err := c.Bind(payload); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		if err := payload.ValidateModel(reqCtx); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		// generate random correlation ID if not provided
		if payload.TaskDetails.CorrelationID == 0 {
			payload.TaskDetails.CorrelationID = rand.Uint64()
		}

		payload.TaskDetails.Timestamp = strfmt.DateTime(time.Now().UTC())
		data := models.NetworkProbeData{
			LastExported:   payload.TaskDetails.Timestamp,
			TargetID:       payload.TaskDetails.TargetID,
			SequenceNumber: 0,
		}

		taskID := string(payload.TaskID)
		if err := storage.StoreNProbeData(networkID, taskID, data); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "failed to store NetworkProbeData"))
		}

		_, err := configurator.CreateEntity(
			reqCtx,
			networkID,
			configurator.NetworkEntity{
				Type:   lte.NetworkProbeTaskEntityType,
				Key:    taskID,
				Config: payload.TaskDetails,
			},
			serdes.Entity,
		)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return c.NoContent(http.StatusCreated)
	}
}

func getNetworkProbeTask(c echo.Context) error {
	paramNames := []string{"network_id", "task_id"}
	values, nerr := obsidian.GetParamValues(c, paramNames...)
	if nerr != nil {
		return nerr
	}

	networkID, taskID := values[0], values[1]
	ent, err := configurator.LoadEntity(
		c.Request().Context(),
		networkID,
		lte.NetworkProbeTaskEntityType,
		taskID,
		configurator.EntityLoadCriteria{LoadConfig: true},
		serdes.Entity)
	if err == merrors.ErrNotFound {
		return echo.ErrNotFound
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	ret := (&models.NetworkProbeTask{}).FromBackendModels(ent)
	return c.JSON(http.StatusOK, ret)
}

func updateNetworkProbeTask(c echo.Context) error {
	networkID, nerr := obsidian.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}
	reqCtx := c.Request().Context()

	payload := &models.NetworkProbeTask{}
	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := payload.ValidateModel(reqCtx); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	_, err := configurator.UpdateEntity(reqCtx, networkID, payload.ToEntityUpdateCriteria(), serdes.Entity)
	if err != nil {
		return obsidian.MakeHTTPError(err, http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}

func getDeleteNetworkProbeTaskHandlerFunc(storage storage.NProbeStorage) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramNames := []string{"network_id", "task_id"}
		values, nerr := obsidian.GetParamValues(c, paramNames...)
		if nerr != nil {
			return nerr
		}

		networkID, taskID := values[0], values[1]
		_ = storage.DeleteNProbeData(networkID, taskID)
		err := configurator.DeleteEntity(c.Request().Context(), networkID, lte.NetworkProbeTaskEntityType, taskID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return c.NoContent(http.StatusNoContent)
	}
}

func listNetworkProbeDestinations(c echo.Context) error {
	networkID, nerr := obsidian.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}

	ents, _, err := configurator.LoadAllEntitiesOfType(
		c.Request().Context(),
		networkID, lte.NetworkProbeDestinationEntityType,
		configurator.EntityLoadCriteria{LoadConfig: true},
		serdes.Entity,
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	ret := make(map[string]*models.NetworkProbeDestination, len(ents))
	for _, ent := range ents {
		ret[ent.Key] = (&models.NetworkProbeDestination{}).FromBackendModels(ent)
	}
	return c.JSON(http.StatusOK, ret)
}

func createNetworkProbeDestination(c echo.Context) error {
	networkID, nerr := obsidian.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}
	reqCtx := c.Request().Context()

	payload := &models.NetworkProbeDestination{}
	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := payload.ValidateModel(reqCtx); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	_, err := configurator.CreateEntity(
		reqCtx,
		networkID,
		configurator.NetworkEntity{
			Type:   lte.NetworkProbeDestinationEntityType,
			Key:    string(payload.DestinationID),
			Config: payload.DestinationDetails,
		},
		serdes.Entity,
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusCreated)
}

func getNetworkProbeDestination(c echo.Context) error {
	paramNames := []string{"network_id", "destination_id"}
	values, nerr := obsidian.GetParamValues(c, paramNames...)
	if nerr != nil {
		return nerr
	}

	networkID, destinationID := values[0], values[1]
	ent, err := configurator.LoadEntity(
		c.Request().Context(),
		networkID,
		lte.NetworkProbeDestinationEntityType,
		destinationID,
		configurator.EntityLoadCriteria{LoadConfig: true},
		serdes.Entity)
	if err == merrors.ErrNotFound {
		return echo.ErrNotFound
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	ret := (&models.NetworkProbeDestination{}).FromBackendModels(ent)
	return c.JSON(http.StatusOK, ret)
}

func updateNetworkProbeDestination(c echo.Context) error {
	networkID, nerr := obsidian.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}
	reqCtx := c.Request().Context()

	payload := &models.NetworkProbeDestination{}
	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := payload.ValidateModel(reqCtx); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	_, err := configurator.UpdateEntity(reqCtx, networkID, payload.ToEntityUpdateCriteria(), serdes.Entity)
	if err != nil {
		return obsidian.MakeHTTPError(err, http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}

func deleteNetworkProbeDestination(c echo.Context) error {
	paramNames := []string{"network_id", "destination_id"}
	values, nerr := obsidian.GetParamValues(c, paramNames...)
	if nerr != nil {
		return nerr
	}

	networkID, destinationID := values[0], values[1]
	err := configurator.DeleteEntity(c.Request().Context(), networkID, lte.NetworkProbeDestinationEntityType, destinationID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
