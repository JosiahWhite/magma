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

import * as PromQL from '../../prometheus/PromQL';
import * as React from 'react';
import Button from '@material-ui/core/Button';
import Divider from '@material-ui/core/Divider';
import FormHelperText from '@material-ui/core/FormHelperText';
import Grid from '@material-ui/core/Grid';
import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import RuleEditorBase from '../RuleEditorBase';
import TextField from '@material-ui/core/TextField';
import ToggleableExpressionEditor, {
  AdvancedExpressionEditor,
  thresholdToPromQL,
} from './ToggleableExpressionEditor';
import useForm from '../../../hooks/useForm';
import useRouter from '../../../../../fbc_js_core/ui/hooks/useRouter';
import {Labels} from '../../prometheus/PromQL';
import {Parse} from '../../prometheus/PromQLParser';
import {SEVERITY} from '../../severity/Severity';
import {makeStyles} from '@material-ui/styles';
import {useAlarmContext} from '../../AlarmContext';
import {useSnackbars} from '../../../../../fbc_js_core/ui/hooks/useSnackbar';

import type {AlertConfig, Labels as LabelsMap} from '../../AlarmAPIType';
import type {GenericRule, RuleEditorProps} from '../RuleInterface';
import type {RuleEditorBaseFields} from '../RuleEditorBase';
import type {ThresholdExpression} from './ToggleableExpressionEditor';

type MenuItemProps = {key: string, value: string, children: string};

type TimeUnit = {value: string, label: string};

const useStyles = makeStyles(theme => ({
  button: {
    marginLeft: -theme.spacing(0.5),
    margin: theme.spacing(1.5),
  },
  divider: {
    margin: `${theme.spacing(2)}px 0`,
  },
}));

const timeUnits: Array<TimeUnit> = [
  {
    value: '',
    label: '',
  },
  {
    value: 's',
    label: 'seconds',
  },
  {
    value: 'm',
    label: 'minutes',
  },
  {
    value: 'h',
    label: 'hours',
  },
];

/**
 * An easier to edit representation of the form's state, then convert
 * to and from the AlertConfig type for posting to the api.
 */
type FormState = {
  ruleName: string,
  expression: string,
  severity: string,
  timeNumber: number,
  timeUnit: string,
  description: string,
  labels: LabelsMap,
};

export type InputChangeFunc = (
  formUpdate: (val: string) => $Shape<FormState>,
) => (event: SyntheticInputEvent<HTMLElement>) => void;

type PrometheusEditorProps = {
  ...RuleEditorProps<AlertConfig>,
};
export default function PrometheusEditor(props: PrometheusEditorProps) {
  const {apiUtil, thresholdEditorEnabled} = useAlarmContext();
  const {isNew, onRuleUpdated, onExit, rule} = props;
  const {match} = useRouter();
  const classes = useStyles();
  const snackbars = useSnackbars();

  /**
   * after the user types into the form, map back from FormState and
   * notify the parent component
   */
  const handleFormUpdated = React.useCallback(
    (state: FormState) => {
      const updatedConfig = toAlertConfig(state);
      onRuleUpdated({
        ...rule,
        ...({rawRule: updatedConfig}: $Shape<GenericRule<AlertConfig>>),
      });
    },
    [onRuleUpdated, rule],
  );

  const {formState, handleInputChange, updateFormState} = useForm({
    initialState: fromAlertConfig(rule ? rule.rawRule : null),
    onFormUpdated: handleFormUpdated,
  });

  const {
    advancedEditorMode,
    setAdvancedEditorMode,
    thresholdExpression,
    setThresholdExpression,
  } = useThresholdExpressionEditorState({
    expression: rule?.expression,
    thresholdEditorEnabled,
  });

  /**
   * Handles when the RuleEditorBase form changes, map this from
   * RuleEditorForm -> AlertConfig
   */
  const handleEditorBaseChange = React.useCallback(
    editorBaseState => {
      updateFormState({
        ruleName: editorBaseState.name,
        description: editorBaseState.description,
        labels: editorBaseState.labels,
      });
    },
    [updateFormState],
  );

  const editorBaseInitialState = React.useMemo(() => toBaseFields(rule), [
    rule,
  ]);

  const saveAlert = async () => {
    try {
      if (!formState) {
        throw new Error('Alert config empty');
      }
      const request = {
        networkId: match.params.networkId,
        rule: toAlertConfig(formState),
      };
      if (isNew) {
        await apiUtil.createAlertRule(request);
      } else {
        await apiUtil.editAlertRule(request);
      }
      snackbars.success(`Successfully ${isNew ? 'added' : 'saved'} alert rule`);
      onExit();
    } catch (error) {
      snackbars.error(
        `Unable to create rule: ${
          error.response ? error.response.data.message : error.message
        }.`,
      );
    }
  };

  const updateThresholdExpression = React.useCallback(
    (expression: ThresholdExpression) => {
      setThresholdExpression(expression);
      const stringExpression = thresholdToPromQL(expression);
      updateFormState({expression: stringExpression});
    },
    [setThresholdExpression, updateFormState],
  );

  const severityOptions = React.useMemo<Array<MenuItemProps>>(
    () =>
      Object.keys(SEVERITY).map(key => ({
        key: key,
        value: key,
        children: key.toLowerCase(),
      })),
    [],
  );

  const toggleMode = () => setAdvancedEditorMode(!advancedEditorMode);
  return (
    <RuleEditorBase
      initialState={editorBaseInitialState}
      onChange={handleEditorBaseChange}
      onSave={saveAlert}
      onExit={onExit}
      isNew={isNew}>
      {advancedEditorMode ? (
        <Grid item>
          <AdvancedExpressionEditor
            expression={formState.expression}
            onChange={handleInputChange}
          />
          <Button
            className={classes.button}
            color="primary"
            size="small"
            target="_blank"
            href="https://prometheus.io/docs/prometheus/latest/querying/basics/">
            PromQL FAQ
          </Button>
          <Button
            className={classes.button}
            color="primary"
            size="small"
            onClick={toggleMode}>
            Switch to template
          </Button>
        </Grid>
      ) : (
        <ToggleableExpressionEditor
          onChange={handleInputChange}
          onThresholdExpressionChange={updateThresholdExpression}
          expression={thresholdExpression}
          stringExpression={formState.expression}
          toggleOn={advancedEditorMode}
          onToggleChange={toggleMode}
        />
      )}
      <Divider className={classes.divider} />
      <Grid
        item
        container
        alignItems="flex-start"
        justifyContent="space-between"
        spacing={2}>
        <TimeEditor
          onChange={handleInputChange}
          timeNumber={formState.timeNumber}
          timeUnit={formState.timeUnit}
        />
        <SeverityEditor
          onChange={handleInputChange}
          options={severityOptions}
          severity={formState.severity}
        />
      </Grid>
    </RuleEditorBase>
  );
}

const useSeverityMenuItemStyles = makeStyles(_theme => ({
  root: {
    textTransform: 'capitalize',
  },
}));
const useSeveritySelectStyles = makeStyles(_theme => ({
  root: {
    textTransform: 'capitalize',
  },
}));
function SeverityEditor(props: {
  onChange: InputChangeFunc,
  severity: string,
  options: Array<MenuItemProps>,
}) {
  const severitySelectClasses = useSeveritySelectStyles();
  const severityMenuItemClasses = useSeverityMenuItemStyles();
  return (
    <Grid item xs={3}>
      <InputLabel htmlFor="severity-input">Severity</InputLabel>
      <TextField
        id="severity-input"
        fullWidth
        required
        select
        value={props.severity}
        onChange={props.onChange(value => ({severity: value}))}
        classes={severitySelectClasses}>
        {props.options.map(opt => (
          <MenuItem {...opt} ListItemClasses={severityMenuItemClasses} />
        ))}
      </TextField>
    </Grid>
  );
}

function TimeEditor(props: {
  onChange: InputChangeFunc,
  timeNumber: number,
  timeUnit: string,
}) {
  return (
    <>
      <TimeNumberEditor
        onChange={props.onChange}
        timeNumber={props.timeNumber}
      />
      <TimeUnitEditor
        onChange={props.onChange}
        timeUnit={props.timeUnit}
        timeUnits={timeUnits}
      />
    </>
  );
}

function TimeNumberEditor(props: {
  onChange: InputChangeFunc,
  timeNumber: number,
}) {
  return (
    <Grid item xs={6}>
      <InputLabel htmlFor="duration-input">Duration</InputLabel>
      <Input
        id="duration-input"
        fullWidth
        type="number"
        value={isNaN(props.timeNumber) ? '' : props.timeNumber}
        onChange={props.onChange(val => ({timeNumber: parseInt(val, 10)}))}
      />
      <FormHelperText>
        Amount of time that conditions are true before an alert is triggered
      </FormHelperText>
    </Grid>
  );
}

function TimeUnitEditor(props: {
  onChange: InputChangeFunc,
  timeUnit: string,
  timeUnits: Array<TimeUnit>,
}) {
  const severitySelectClasses = useSeveritySelectStyles();
  const severityMenuItemClasses = useSeverityMenuItemStyles();
  return (
    <Grid item xs={3}>
      <InputLabel htmlFor="unit-input">Unit</InputLabel>
      <TextField
        id="unit-input"
        fullWidth
        select
        value={props.timeUnit}
        onChange={props.onChange(val => ({timeUnit: val}))}
        classes={severitySelectClasses}>
        {props.timeUnits.map(option => (
          <MenuItem
            key={option.value}
            value={option.value}
            classes={severityMenuItemClasses}>
            {option.label}
          </MenuItem>
        ))}
      </TextField>
    </Grid>
  );
}

function fromAlertConfig(rule: ?AlertConfig): FormState {
  if (!rule) {
    return {
      ruleName: '',
      expression: '',
      severity: SEVERITY.WARNING.name,
      description: '',
      timeNumber: 0,
      timeUnit: 's',
      labels: {},
    };
  }
  const timeString = rule.for ?? '';
  const {timeNumber, timeUnit} = getMostSignificantTime(
    parseTimeString(timeString),
  );
  return {
    ruleName: rule.alert || '',
    expression: rule.expr || '',
    severity: rule.labels?.severity || '',
    description: rule.annotations?.description || '',
    timeNumber: timeNumber,
    timeUnit,
    labels: rule.labels || {},
  };
}

function toAlertConfig(form: FormState): AlertConfig {
  return {
    alert: form.ruleName,
    expr: form.expression,
    labels: {
      ...form.labels,
      severity: form.severity,
    },
    for: `${form.timeNumber}${form.timeUnit}`,
    annotations: {
      description: form.description,
    },
  };
}

/**
 * Map from rule-specific type to the generic RuleEditorBaseFields
 */
export function toBaseFields(
  rule: ?GenericRule<AlertConfig>,
): RuleEditorBaseFields {
  return {
    name: rule?.name || '',
    description: rule?.description || '',
    labels: rule?.rawRule?.labels || {},
  };
}

export type Duration = {
  hours: number,
  minutes: number,
  seconds: number,
};

export function parseTimeString(timeStamp: string): Duration {
  if (timeStamp === '') {
    return {hours: 0, minutes: 0, seconds: 0};
  }
  const durationRegex = /^((\d+)h)*((\d+)m)*((\d+)s)*$/;
  const duration = timeStamp.match(durationRegex);
  if (!duration) {
    return {hours: 0, minutes: 0, seconds: 0};
  }
  // Index is corresponding capture group from regex
  const hours = parseInt(duration[2], 10) || 0;
  const minutes = parseInt(duration[4], 10) || 0;
  const seconds = parseInt(duration[6], 10) || 0;
  return {hours, minutes, seconds};
}

function getMostSignificantTime(
  duration: Duration,
): {timeNumber: number, timeUnit: string} {
  if (duration.hours) {
    return {timeNumber: duration.hours, timeUnit: 'h'};
  } else if (duration.minutes) {
    return {timeNumber: duration.minutes, timeUnit: 'm'};
  }
  return {timeNumber: duration.seconds, timeUnit: 's'};
}

function getThresholdExpression(
  exp: PromQL.Expression<string | number>,
): ?ThresholdExpression {
  if (
    !(
      exp instanceof PromQL.BinaryOperation &&
      exp.lh instanceof PromQL.InstantSelector &&
      exp.rh instanceof PromQL.Scalar
    )
  ) {
    return null;
  }

  const metricName = exp.lh.selectorName || '';
  const threshold = exp.rh.value;
  const filters = exp.lh.labels || new PromQL.Labels();
  filters.removeByName('networkID');
  const comparator = asBinaryComparator(exp.operator);
  if (!comparator) {
    return null;
  }
  return {
    metricName,
    filters,
    comparator,
    value: threshold,
  };
}

function asBinaryComparator(
  operator: PromQL.BinaryOperator,
): ?PromQL.BinaryComparator {
  if (operator instanceof Object) {
    return operator;
  }
  return null;
}

function useThresholdExpressionEditorState({
  expression,
  thresholdEditorEnabled,
}: {
  expression: ?string,
  thresholdEditorEnabled: ?boolean,
}): {
  thresholdExpression: ThresholdExpression,
  setThresholdExpression: ThresholdExpression => void,
  advancedEditorMode: boolean,
  setAdvancedEditorMode: boolean => void,
} {
  const [
    thresholdExpression,
    setThresholdExpression,
  ] = React.useState<ThresholdExpression>({
    metricName: '',
    comparator: new PromQL.BinaryComparator('=='),
    value: 0,
    filters: new Labels(),
  });
  const snackbars = useSnackbars();

  const [advancedEditorMode, setAdvancedEditorMode] = React.useState<boolean>(
    !thresholdEditorEnabled,
  );

  // Parse the expression string once when the component mounts
  const parsedExpression = React.useMemo(() => {
    try {
      return Parse(expression);
    } catch {
      return null;
    }
    // We only want to parse the expression on the first
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  /**
   * After parsing the expression, caches the threshold expression in state. If
   * the expression cannot be parsed, swaps to the advanced editor mode.
   */
  React.useEffect(() => {
    if (!thresholdEditorEnabled) {
      return;
    }
    if (!expression) {
      setAdvancedEditorMode(false);
    } else if (parsedExpression) {
      const newThresholdExpression = getThresholdExpression(parsedExpression);
      if (newThresholdExpression) {
        setAdvancedEditorMode(false);
        setThresholdExpression(newThresholdExpression);
      } else {
        setAdvancedEditorMode(true);
      }
    } else {
      snackbars.error(
        "Error parsing alert expression. You can still edit this using the advanced editor, but you won't be able to use the UI expression editor.",
      );
    }
    // we only want this to run when the parsedExpression changes
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [parsedExpression, thresholdEditorEnabled]);

  return {
    thresholdExpression,
    setThresholdExpression,
    advancedEditorMode,
    setAdvancedEditorMode,
  };
}
