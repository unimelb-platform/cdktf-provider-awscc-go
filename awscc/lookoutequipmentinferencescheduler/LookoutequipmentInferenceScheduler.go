package lookoutequipmentinferencescheduler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lookoutequipmentinferencescheduler/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutequipment_inference_scheduler awscc_lookoutequipment_inference_scheduler}.
type LookoutequipmentInferenceScheduler interface {
	cdktf.TerraformResource
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
	DataDelayOffsetInMinutes() *float64
	SetDataDelayOffsetInMinutes(val *float64)
	DataDelayOffsetInMinutesInput() *float64
	DataInputConfiguration() LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference
	DataInputConfigurationInput() interface{}
	DataOutputConfiguration() LookoutequipmentInferenceSchedulerDataOutputConfigurationOutputReference
	DataOutputConfigurationInput() interface{}
	DataUploadFrequency() *string
	SetDataUploadFrequency(val *string)
	DataUploadFrequencyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	InferenceSchedulerArn() *string
	InferenceSchedulerName() *string
	SetInferenceSchedulerName(val *string)
	InferenceSchedulerNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ModelName() *string
	SetModelName(val *string)
	ModelNameInput() *string
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
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	ServerSideKmsKeyId() *string
	SetServerSideKmsKeyId(val *string)
	ServerSideKmsKeyIdInput() *string
	Tags() LookoutequipmentInferenceSchedulerTagsList
	TagsInput() interface{}
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
	PutDataInputConfiguration(value *LookoutequipmentInferenceSchedulerDataInputConfiguration)
	PutDataOutputConfiguration(value *LookoutequipmentInferenceSchedulerDataOutputConfiguration)
	PutTags(value interface{})
	ResetDataDelayOffsetInMinutes()
	ResetInferenceSchedulerName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServerSideKmsKeyId()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LookoutequipmentInferenceScheduler
type jsiiProxy_LookoutequipmentInferenceScheduler struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) DataDelayOffsetInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataDelayOffsetInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) DataDelayOffsetInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataDelayOffsetInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) DataInputConfiguration() LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference {
	var returns LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference
	_jsii_.Get(
		j,
		"dataInputConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) DataInputConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataInputConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) DataOutputConfiguration() LookoutequipmentInferenceSchedulerDataOutputConfigurationOutputReference {
	var returns LookoutequipmentInferenceSchedulerDataOutputConfigurationOutputReference
	_jsii_.Get(
		j,
		"dataOutputConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) DataOutputConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataOutputConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) DataUploadFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataUploadFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) DataUploadFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataUploadFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) InferenceSchedulerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceSchedulerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) InferenceSchedulerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceSchedulerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) InferenceSchedulerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceSchedulerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) ModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) ModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) ServerSideKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) ServerSideKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) Tags() LookoutequipmentInferenceSchedulerTagsList {
	var returns LookoutequipmentInferenceSchedulerTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutequipment_inference_scheduler awscc_lookoutequipment_inference_scheduler} Resource.
func NewLookoutequipmentInferenceScheduler(scope constructs.Construct, id *string, config *LookoutequipmentInferenceSchedulerConfig) LookoutequipmentInferenceScheduler {
	_init_.Initialize()

	if err := validateNewLookoutequipmentInferenceSchedulerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LookoutequipmentInferenceScheduler{}

	_jsii_.Create(
		"awscc.lookoutequipmentInferenceScheduler.LookoutequipmentInferenceScheduler",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutequipment_inference_scheduler awscc_lookoutequipment_inference_scheduler} Resource.
func NewLookoutequipmentInferenceScheduler_Override(l LookoutequipmentInferenceScheduler, scope constructs.Construct, id *string, config *LookoutequipmentInferenceSchedulerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lookoutequipmentInferenceScheduler.LookoutequipmentInferenceScheduler",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler)SetDataDelayOffsetInMinutes(val *float64) {
	if err := j.validateSetDataDelayOffsetInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataDelayOffsetInMinutes",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler)SetDataUploadFrequency(val *string) {
	if err := j.validateSetDataUploadFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataUploadFrequency",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler)SetInferenceSchedulerName(val *string) {
	if err := j.validateSetInferenceSchedulerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceSchedulerName",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler)SetModelName(val *string) {
	if err := j.validateSetModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelName",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceScheduler)SetServerSideKmsKeyId(val *string) {
	if err := j.validateSetServerSideKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverSideKmsKeyId",
		val,
	)
}

// Generates CDKTF code for importing a LookoutequipmentInferenceScheduler resource upon running "cdktf plan <stack-name>".
func LookoutequipmentInferenceScheduler_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLookoutequipmentInferenceScheduler_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.lookoutequipmentInferenceScheduler.LookoutequipmentInferenceScheduler",
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
func LookoutequipmentInferenceScheduler_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLookoutequipmentInferenceScheduler_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lookoutequipmentInferenceScheduler.LookoutequipmentInferenceScheduler",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LookoutequipmentInferenceScheduler_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLookoutequipmentInferenceScheduler_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lookoutequipmentInferenceScheduler.LookoutequipmentInferenceScheduler",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LookoutequipmentInferenceScheduler_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLookoutequipmentInferenceScheduler_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lookoutequipmentInferenceScheduler.LookoutequipmentInferenceScheduler",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LookoutequipmentInferenceScheduler_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.lookoutequipmentInferenceScheduler.LookoutequipmentInferenceScheduler",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) PutDataInputConfiguration(value *LookoutequipmentInferenceSchedulerDataInputConfiguration) {
	if err := l.validatePutDataInputConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDataInputConfiguration",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) PutDataOutputConfiguration(value *LookoutequipmentInferenceSchedulerDataOutputConfiguration) {
	if err := l.validatePutDataOutputConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDataOutputConfiguration",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) PutTags(value interface{}) {
	if err := l.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTags",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) ResetDataDelayOffsetInMinutes() {
	_jsii_.InvokeVoid(
		l,
		"resetDataDelayOffsetInMinutes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) ResetInferenceSchedulerName() {
	_jsii_.InvokeVoid(
		l,
		"resetInferenceSchedulerName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) ResetServerSideKmsKeyId() {
	_jsii_.InvokeVoid(
		l,
		"resetServerSideKmsKeyId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceScheduler) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

