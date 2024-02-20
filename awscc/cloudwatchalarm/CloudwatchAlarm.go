package cloudwatchalarm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cloudwatchalarm/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm awscc_cloudwatch_alarm}.
type CloudwatchAlarm interface {
	cdktf.TerraformResource
	ActionsEnabled() interface{}
	SetActionsEnabled(val interface{})
	ActionsEnabledInput() interface{}
	AlarmActions() *[]*string
	SetAlarmActions(val *[]*string)
	AlarmActionsInput() *[]*string
	AlarmDescription() *string
	SetAlarmDescription(val *string)
	AlarmDescriptionInput() *string
	AlarmName() *string
	SetAlarmName(val *string)
	AlarmNameInput() *string
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComparisonOperator() *string
	SetComparisonOperator(val *string)
	ComparisonOperatorInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DatapointsToAlarm() *float64
	SetDatapointsToAlarm(val *float64)
	DatapointsToAlarmInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Dimensions() CloudwatchAlarmDimensionsList
	DimensionsInput() interface{}
	EvaluateLowSampleCountPercentile() *string
	SetEvaluateLowSampleCountPercentile(val *string)
	EvaluateLowSampleCountPercentileInput() *string
	EvaluationPeriods() *float64
	SetEvaluationPeriods(val *float64)
	EvaluationPeriodsInput() *float64
	ExtendedStatistic() *string
	SetExtendedStatistic(val *string)
	ExtendedStatisticInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	InsufficientDataActions() *[]*string
	SetInsufficientDataActions(val *[]*string)
	InsufficientDataActionsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MetricName() *string
	SetMetricName(val *string)
	MetricNameInput() *string
	Metrics() CloudwatchAlarmMetricsList
	MetricsInput() interface{}
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	OkActions() *[]*string
	SetOkActions(val *[]*string)
	OkActionsInput() *[]*string
	Period() *float64
	SetPeriod(val *float64)
	PeriodInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Statistic() *string
	SetStatistic(val *string)
	StatisticInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Threshold() *float64
	SetThreshold(val *float64)
	ThresholdInput() *float64
	ThresholdMetricId() *string
	SetThresholdMetricId(val *string)
	ThresholdMetricIdInput() *string
	TreatMissingData() *string
	SetTreatMissingData(val *string)
	TreatMissingDataInput() *string
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDimensions(value interface{})
	PutMetrics(value interface{})
	ResetActionsEnabled()
	ResetAlarmActions()
	ResetAlarmDescription()
	ResetAlarmName()
	ResetDatapointsToAlarm()
	ResetDimensions()
	ResetEvaluateLowSampleCountPercentile()
	ResetExtendedStatistic()
	ResetInsufficientDataActions()
	ResetMetricName()
	ResetMetrics()
	ResetNamespace()
	ResetOkActions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeriod()
	ResetStatistic()
	ResetThreshold()
	ResetThresholdMetricId()
	ResetTreatMissingData()
	ResetUnit()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudwatchAlarm
type jsiiProxy_CloudwatchAlarm struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudwatchAlarm) ActionsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) ActionsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) AlarmActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarmActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) AlarmActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarmActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) AlarmDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) AlarmDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) AlarmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) AlarmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) ComparisonOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) ComparisonOperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonOperatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) DatapointsToAlarm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"datapointsToAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) DatapointsToAlarmInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"datapointsToAlarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) Dimensions() CloudwatchAlarmDimensionsList {
	var returns CloudwatchAlarmDimensionsList
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) DimensionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) EvaluateLowSampleCountPercentile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluateLowSampleCountPercentile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) EvaluateLowSampleCountPercentileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluateLowSampleCountPercentileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) EvaluationPeriods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) EvaluationPeriodsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) ExtendedStatistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extendedStatistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) ExtendedStatisticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extendedStatisticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) InsufficientDataActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insufficientDataActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) InsufficientDataActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insufficientDataActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) Metrics() CloudwatchAlarmMetricsList {
	var returns CloudwatchAlarmMetricsList
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) MetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) OkActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"okActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) OkActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"okActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) PeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) StatisticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statisticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) ThresholdMetricId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdMetricId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) ThresholdMetricIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdMetricIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) TreatMissingData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treatMissingData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) TreatMissingDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treatMissingDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlarm) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm awscc_cloudwatch_alarm} Resource.
func NewCloudwatchAlarm(scope constructs.Construct, id *string, config *CloudwatchAlarmConfig) CloudwatchAlarm {
	_init_.Initialize()

	if err := validateNewCloudwatchAlarmParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudwatchAlarm{}

	_jsii_.Create(
		"awscc.cloudwatchAlarm.CloudwatchAlarm",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm awscc_cloudwatch_alarm} Resource.
func NewCloudwatchAlarm_Override(c CloudwatchAlarm, scope constructs.Construct, id *string, config *CloudwatchAlarmConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cloudwatchAlarm.CloudwatchAlarm",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetActionsEnabled(val interface{}) {
	if err := j.validateSetActionsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionsEnabled",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetAlarmActions(val *[]*string) {
	if err := j.validateSetAlarmActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmActions",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetAlarmDescription(val *string) {
	if err := j.validateSetAlarmDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmDescription",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetAlarmName(val *string) {
	if err := j.validateSetAlarmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmName",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetComparisonOperator(val *string) {
	if err := j.validateSetComparisonOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparisonOperator",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetDatapointsToAlarm(val *float64) {
	if err := j.validateSetDatapointsToAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datapointsToAlarm",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetEvaluateLowSampleCountPercentile(val *string) {
	if err := j.validateSetEvaluateLowSampleCountPercentileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluateLowSampleCountPercentile",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetEvaluationPeriods(val *float64) {
	if err := j.validateSetEvaluationPeriodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationPeriods",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetExtendedStatistic(val *string) {
	if err := j.validateSetExtendedStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendedStatistic",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetInsufficientDataActions(val *[]*string) {
	if err := j.validateSetInsufficientDataActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insufficientDataActions",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetOkActions(val *[]*string) {
	if err := j.validateSetOkActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"okActions",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetPeriod(val *float64) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetThreshold(val *float64) {
	if err := j.validateSetThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetThresholdMetricId(val *string) {
	if err := j.validateSetThresholdMetricIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdMetricId",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetTreatMissingData(val *string) {
	if err := j.validateSetTreatMissingDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"treatMissingData",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlarm)SetUnit(val *string) {
	if err := j.validateSetUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

// Generates CDKTF code for importing a CloudwatchAlarm resource upon running "cdktf plan <stack-name>".
func CloudwatchAlarm_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCloudwatchAlarm_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.cloudwatchAlarm.CloudwatchAlarm",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CloudwatchAlarm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudwatchAlarm_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cloudwatchAlarm.CloudwatchAlarm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudwatchAlarm_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudwatchAlarm_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cloudwatchAlarm.CloudwatchAlarm",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudwatchAlarm_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudwatchAlarm_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cloudwatchAlarm.CloudwatchAlarm",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudwatchAlarm_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.cloudwatchAlarm.CloudwatchAlarm",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudwatchAlarm) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CloudwatchAlarm) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudwatchAlarm) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlarm) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlarm) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlarm) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlarm) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlarm) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlarm) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlarm) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlarm) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlarm) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CloudwatchAlarm) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlarm) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudwatchAlarm) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudwatchAlarm) PutDimensions(value interface{}) {
	if err := c.validatePutDimensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDimensions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchAlarm) PutMetrics(value interface{}) {
	if err := c.validatePutMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMetrics",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetActionsEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetActionsEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetAlarmActions() {
	_jsii_.InvokeVoid(
		c,
		"resetAlarmActions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetAlarmDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetAlarmDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetAlarmName() {
	_jsii_.InvokeVoid(
		c,
		"resetAlarmName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetDatapointsToAlarm() {
	_jsii_.InvokeVoid(
		c,
		"resetDatapointsToAlarm",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetDimensions() {
	_jsii_.InvokeVoid(
		c,
		"resetDimensions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetEvaluateLowSampleCountPercentile() {
	_jsii_.InvokeVoid(
		c,
		"resetEvaluateLowSampleCountPercentile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetExtendedStatistic() {
	_jsii_.InvokeVoid(
		c,
		"resetExtendedStatistic",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetInsufficientDataActions() {
	_jsii_.InvokeVoid(
		c,
		"resetInsufficientDataActions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetMetricName() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetMetrics() {
	_jsii_.InvokeVoid(
		c,
		"resetMetrics",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetNamespace() {
	_jsii_.InvokeVoid(
		c,
		"resetNamespace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetOkActions() {
	_jsii_.InvokeVoid(
		c,
		"resetOkActions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetPeriod() {
	_jsii_.InvokeVoid(
		c,
		"resetPeriod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetStatistic() {
	_jsii_.InvokeVoid(
		c,
		"resetStatistic",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetThresholdMetricId() {
	_jsii_.InvokeVoid(
		c,
		"resetThresholdMetricId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetTreatMissingData() {
	_jsii_.InvokeVoid(
		c,
		"resetTreatMissingData",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) ResetUnit() {
	_jsii_.InvokeVoid(
		c,
		"resetUnit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlarm) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlarm) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlarm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlarm) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

