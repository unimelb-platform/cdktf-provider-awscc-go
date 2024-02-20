package logsloganomalydetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/logsloganomalydetector/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_log_anomaly_detector awscc_logs_log_anomaly_detector}.
type LogsLogAnomalyDetector interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AnomalyDetectorArn() *string
	AnomalyDetectorStatus() *string
	AnomalyVisibilityTime() *float64
	SetAnomalyVisibilityTime(val *float64)
	AnomalyVisibilityTimeInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CreationTimeStamp() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DetectorName() *string
	SetDetectorName(val *string)
	DetectorNameInput() *string
	EvaluationFrequency() *string
	SetEvaluationFrequency(val *string)
	EvaluationFrequencyInput() *string
	FilterPattern() *string
	SetFilterPattern(val *string)
	FilterPatternInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	LastModifiedTimeStamp() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogGroupArnList() *[]*string
	SetLogGroupArnList(val *[]*string)
	LogGroupArnListInput() *[]*string
	// The tree node.
	Node() constructs.Node
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetAccountId()
	ResetAnomalyVisibilityTime()
	ResetDetectorName()
	ResetEvaluationFrequency()
	ResetFilterPattern()
	ResetKmsKeyId()
	ResetLogGroupArnList()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LogsLogAnomalyDetector
type jsiiProxy_LogsLogAnomalyDetector struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LogsLogAnomalyDetector) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) AnomalyDetectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anomalyDetectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) AnomalyDetectorStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anomalyDetectorStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) AnomalyVisibilityTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anomalyVisibilityTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) AnomalyVisibilityTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anomalyVisibilityTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) CreationTimeStamp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTimeStamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) DetectorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) DetectorNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) EvaluationFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) EvaluationFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) FilterPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) FilterPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) LastModifiedTimeStamp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastModifiedTimeStamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) LogGroupArnList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logGroupArnList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) LogGroupArnListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logGroupArnListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsLogAnomalyDetector) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_log_anomaly_detector awscc_logs_log_anomaly_detector} Resource.
func NewLogsLogAnomalyDetector(scope constructs.Construct, id *string, config *LogsLogAnomalyDetectorConfig) LogsLogAnomalyDetector {
	_init_.Initialize()

	if err := validateNewLogsLogAnomalyDetectorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogsLogAnomalyDetector{}

	_jsii_.Create(
		"awscc.logsLogAnomalyDetector.LogsLogAnomalyDetector",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_log_anomaly_detector awscc_logs_log_anomaly_detector} Resource.
func NewLogsLogAnomalyDetector_Override(l LogsLogAnomalyDetector, scope constructs.Construct, id *string, config *LogsLogAnomalyDetectorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.logsLogAnomalyDetector.LogsLogAnomalyDetector",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LogsLogAnomalyDetector)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_LogsLogAnomalyDetector)SetAnomalyVisibilityTime(val *float64) {
	if err := j.validateSetAnomalyVisibilityTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anomalyVisibilityTime",
		val,
	)
}

func (j *jsiiProxy_LogsLogAnomalyDetector)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LogsLogAnomalyDetector)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LogsLogAnomalyDetector)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LogsLogAnomalyDetector)SetDetectorName(val *string) {
	if err := j.validateSetDetectorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detectorName",
		val,
	)
}

func (j *jsiiProxy_LogsLogAnomalyDetector)SetEvaluationFrequency(val *string) {
	if err := j.validateSetEvaluationFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationFrequency",
		val,
	)
}

func (j *jsiiProxy_LogsLogAnomalyDetector)SetFilterPattern(val *string) {
	if err := j.validateSetFilterPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterPattern",
		val,
	)
}

func (j *jsiiProxy_LogsLogAnomalyDetector)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LogsLogAnomalyDetector)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_LogsLogAnomalyDetector)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LogsLogAnomalyDetector)SetLogGroupArnList(val *[]*string) {
	if err := j.validateSetLogGroupArnListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupArnList",
		val,
	)
}

func (j *jsiiProxy_LogsLogAnomalyDetector)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LogsLogAnomalyDetector)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a LogsLogAnomalyDetector resource upon running "cdktf plan <stack-name>".
func LogsLogAnomalyDetector_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLogsLogAnomalyDetector_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.logsLogAnomalyDetector.LogsLogAnomalyDetector",
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
func LogsLogAnomalyDetector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogsLogAnomalyDetector_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.logsLogAnomalyDetector.LogsLogAnomalyDetector",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogsLogAnomalyDetector_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogsLogAnomalyDetector_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.logsLogAnomalyDetector.LogsLogAnomalyDetector",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogsLogAnomalyDetector_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogsLogAnomalyDetector_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.logsLogAnomalyDetector.LogsLogAnomalyDetector",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LogsLogAnomalyDetector_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.logsLogAnomalyDetector.LogsLogAnomalyDetector",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LogsLogAnomalyDetector) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LogsLogAnomalyDetector) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LogsLogAnomalyDetector) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsLogAnomalyDetector) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsLogAnomalyDetector) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsLogAnomalyDetector) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsLogAnomalyDetector) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsLogAnomalyDetector) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsLogAnomalyDetector) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsLogAnomalyDetector) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsLogAnomalyDetector) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsLogAnomalyDetector) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LogsLogAnomalyDetector) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsLogAnomalyDetector) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LogsLogAnomalyDetector) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LogsLogAnomalyDetector) ResetAccountId() {
	_jsii_.InvokeVoid(
		l,
		"resetAccountId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsLogAnomalyDetector) ResetAnomalyVisibilityTime() {
	_jsii_.InvokeVoid(
		l,
		"resetAnomalyVisibilityTime",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsLogAnomalyDetector) ResetDetectorName() {
	_jsii_.InvokeVoid(
		l,
		"resetDetectorName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsLogAnomalyDetector) ResetEvaluationFrequency() {
	_jsii_.InvokeVoid(
		l,
		"resetEvaluationFrequency",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsLogAnomalyDetector) ResetFilterPattern() {
	_jsii_.InvokeVoid(
		l,
		"resetFilterPattern",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsLogAnomalyDetector) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		l,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsLogAnomalyDetector) ResetLogGroupArnList() {
	_jsii_.InvokeVoid(
		l,
		"resetLogGroupArnList",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsLogAnomalyDetector) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsLogAnomalyDetector) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsLogAnomalyDetector) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsLogAnomalyDetector) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsLogAnomalyDetector) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

