package sagemakerinferenceexperiment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sagemakerinferenceexperiment/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_experiment awscc_sagemaker_inference_experiment}.
type SagemakerInferenceExperiment interface {
	cdktf.TerraformResource
	Arn() *string
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
	CreationTime() *string
	DataStorageConfig() SagemakerInferenceExperimentDataStorageConfigOutputReference
	DataStorageConfigInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DesiredState() *string
	SetDesiredState(val *string)
	DesiredStateInput() *string
	EndpointMetadata() SagemakerInferenceExperimentEndpointMetadataOutputReference
	EndpointName() *string
	SetEndpointName(val *string)
	EndpointNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	LastModifiedTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ModelVariants() SagemakerInferenceExperimentModelVariantsList
	ModelVariantsInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	Schedule() SagemakerInferenceExperimentScheduleOutputReference
	ScheduleInput() interface{}
	ShadowModeConfig() SagemakerInferenceExperimentShadowModeConfigOutputReference
	ShadowModeConfigInput() interface{}
	Status() *string
	StatusReason() *string
	SetStatusReason(val *string)
	StatusReasonInput() *string
	Tags() SagemakerInferenceExperimentTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutDataStorageConfig(value *SagemakerInferenceExperimentDataStorageConfig)
	PutModelVariants(value interface{})
	PutSchedule(value *SagemakerInferenceExperimentSchedule)
	PutShadowModeConfig(value *SagemakerInferenceExperimentShadowModeConfig)
	PutTags(value interface{})
	ResetDataStorageConfig()
	ResetDescription()
	ResetDesiredState()
	ResetKmsKey()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSchedule()
	ResetShadowModeConfig()
	ResetStatusReason()
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

// The jsii proxy struct for SagemakerInferenceExperiment
type jsiiProxy_SagemakerInferenceExperiment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SagemakerInferenceExperiment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) DataStorageConfig() SagemakerInferenceExperimentDataStorageConfigOutputReference {
	var returns SagemakerInferenceExperimentDataStorageConfigOutputReference
	_jsii_.Get(
		j,
		"dataStorageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) DataStorageConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataStorageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) DesiredState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) DesiredStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) EndpointMetadata() SagemakerInferenceExperimentEndpointMetadataOutputReference {
	var returns SagemakerInferenceExperimentEndpointMetadataOutputReference
	_jsii_.Get(
		j,
		"endpointMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) EndpointNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) ModelVariants() SagemakerInferenceExperimentModelVariantsList {
	var returns SagemakerInferenceExperimentModelVariantsList
	_jsii_.Get(
		j,
		"modelVariants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) ModelVariantsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelVariantsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) Schedule() SagemakerInferenceExperimentScheduleOutputReference {
	var returns SagemakerInferenceExperimentScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) ShadowModeConfig() SagemakerInferenceExperimentShadowModeConfigOutputReference {
	var returns SagemakerInferenceExperimentShadowModeConfigOutputReference
	_jsii_.Get(
		j,
		"shadowModeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) ShadowModeConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shadowModeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) StatusReasonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) Tags() SagemakerInferenceExperimentTagsList {
	var returns SagemakerInferenceExperimentTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceExperiment) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_experiment awscc_sagemaker_inference_experiment} Resource.
func NewSagemakerInferenceExperiment(scope constructs.Construct, id *string, config *SagemakerInferenceExperimentConfig) SagemakerInferenceExperiment {
	_init_.Initialize()

	if err := validateNewSagemakerInferenceExperimentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerInferenceExperiment{}

	_jsii_.Create(
		"awscc.sagemakerInferenceExperiment.SagemakerInferenceExperiment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_experiment awscc_sagemaker_inference_experiment} Resource.
func NewSagemakerInferenceExperiment_Override(s SagemakerInferenceExperiment, scope constructs.Construct, id *string, config *SagemakerInferenceExperimentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sagemakerInferenceExperiment.SagemakerInferenceExperiment",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SagemakerInferenceExperiment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceExperiment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceExperiment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceExperiment)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceExperiment)SetDesiredState(val *string) {
	if err := j.validateSetDesiredStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredState",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceExperiment)SetEndpointName(val *string) {
	if err := j.validateSetEndpointNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointName",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceExperiment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceExperiment)SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceExperiment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceExperiment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceExperiment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceExperiment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceExperiment)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceExperiment)SetStatusReason(val *string) {
	if err := j.validateSetStatusReasonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statusReason",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceExperiment)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a SagemakerInferenceExperiment resource upon running "cdktf plan <stack-name>".
func SagemakerInferenceExperiment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSagemakerInferenceExperiment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.sagemakerInferenceExperiment.SagemakerInferenceExperiment",
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
func SagemakerInferenceExperiment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerInferenceExperiment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sagemakerInferenceExperiment.SagemakerInferenceExperiment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerInferenceExperiment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerInferenceExperiment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sagemakerInferenceExperiment.SagemakerInferenceExperiment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerInferenceExperiment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerInferenceExperiment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sagemakerInferenceExperiment.SagemakerInferenceExperiment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SagemakerInferenceExperiment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.sagemakerInferenceExperiment.SagemakerInferenceExperiment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SagemakerInferenceExperiment) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceExperiment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceExperiment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceExperiment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceExperiment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceExperiment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceExperiment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceExperiment) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceExperiment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceExperiment) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceExperiment) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) PutDataStorageConfig(value *SagemakerInferenceExperimentDataStorageConfig) {
	if err := s.validatePutDataStorageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDataStorageConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) PutModelVariants(value interface{}) {
	if err := s.validatePutModelVariantsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelVariants",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) PutSchedule(value *SagemakerInferenceExperimentSchedule) {
	if err := s.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSchedule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) PutShadowModeConfig(value *SagemakerInferenceExperimentShadowModeConfig) {
	if err := s.validatePutShadowModeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putShadowModeConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) PutTags(value interface{}) {
	if err := s.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTags",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) ResetDataStorageConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetDataStorageConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) ResetDesiredState() {
	_jsii_.InvokeVoid(
		s,
		"resetDesiredState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) ResetKmsKey() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) ResetSchedule() {
	_jsii_.InvokeVoid(
		s,
		"resetSchedule",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) ResetShadowModeConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetShadowModeConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) ResetStatusReason() {
	_jsii_.InvokeVoid(
		s,
		"resetStatusReason",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceExperiment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceExperiment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceExperiment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceExperiment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

