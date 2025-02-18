/**
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
 *
 * @flow
 * @format
 */

import AccountSettings from '../AccountSettings';
import Admin from '../admin/Admin';
import AppContext from '../../../fbc_js_core/ui/context/AppContext';
import LoadingFiller from '../../../fbc_js_core/ui/components/LoadingFiller';
import React, {useContext} from 'react';
import useSections from './useSections';
import {Redirect, Route, Switch} from 'react-router-dom';
import {useRouter} from '../../../fbc_js_core/ui/hooks';

export default function SectionRoutes() {
  const {relativePath, match} = useRouter();
  const [landingPath, sections] = useSections();
  const {user, ssoEnabled} = useContext(AppContext);

  return (
    <Switch>
      {sections.map(section => (
        <Route
          key={section.path}
          path={relativePath(`/${section.path}`)}
          component={section.component}
        />
      ))}
      {user.isSuperUser && (
        <Route key="admin" path={relativePath(`/admin`)} component={Admin} />
      )}
      {!ssoEnabled && (
        <Route path={relativePath('/settings')} component={AccountSettings} />
      )}
      {landingPath && (
        <Route
          path={relativePath('')}
          render={() => (
            <Redirect to={`/nms/${match.params.networkId}/${landingPath}`} />
          )}
        />
      )}
      <LoadingFiller />
    </Switch>
  );
}
