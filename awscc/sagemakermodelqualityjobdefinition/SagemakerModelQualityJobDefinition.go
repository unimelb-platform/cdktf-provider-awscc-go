package sagemakermodelqualityjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sagemakermodelqualityjobdefinition/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_quality_job_definition awscc_sagemaker_model_quality_job_definition}.
type SagemakerModelQualityJobDefinition interface {
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
	CreationTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	JobDefinitionArn() *string
	JobDefinitionName() *string
	SetJobDefinitionName(val *string)
	JobDefinitionNameInput() *string
	JobResources() SagemakerModelQualityJobDefinitionJobResourcesOutputReference
	JobResourcesInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ModelQualityAppSpecification() SagemakerModelQualityJobDefinitionModelQualityAppSpecificationOutputReference
	ModelQualityAppSpecificationInput() interface{}
	ModelQualityBaselineConfig() SagemakerModelQualityJobDefinitionModelQualityBaselineConfigOutputReference
	ModelQualityBaselineConfigInput() interface{}
	ModelQualityJobInput() SagemakerModelQualityJobDefinitionModelQualityJobInputOutputReference
	ModelQualityJobInputInput() interface{}
	ModelQualityJobOutputConfig() SagemakerModelQualityJobDefinitionModelQualityJobOutputConfigOutputReference
	ModelQualityJobOutputConfigInput() interface{}
	NetworkConfig() SagemakerModelQualityJobDefinitionNetworkConfigOutputReference
	NetworkConfigInput() interface{}
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
	StoppingCondition() SagemakerModelQualityJobDefinitionStoppingConditionOutputReference
	StoppingConditionInput() interface{}
	Tags() SagemakerModelQualityJobDefinitionTagsList
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
	PutJobResources(value *SagemakerModelQualityJobDefinitionJobResources)
	PutModelQualityAppSpecification(value *SagemakerModelQualityJobDefinitionModelQualityAppSpecification)
	PutModelQualityBaselineConfig(value *SagemakerModelQualityJobDefinitionModelQualityBaselineConfig)
	PutModelQualityJobInput(value *SagemakerModelQualityJobDefinitionModelQualityJobInput)
	PutModelQualityJobOutputConfig(value *SagemakerModelQualityJobDefinitionModelQualityJobOutputConfig)
	PutNetworkConfig(value *SagemakerModelQualityJobDefinitionNetworkConfig)
	PutStoppingCondition(value *SagemakerModelQualityJobDefinitionStoppingCondition)
	PutTags(value interface{})
	ResetEndpointName()
	ResetJobDefinitionName()
	ResetModelQualityBaselineConfig()
	ResetNetworkConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStoppingCondition()
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

// The jsii proxy struct for SagemakerModelQualityJobDefinition
type jsiiProxy_SagemakerModelQualityJobDefinition struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) EndpointNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) JobDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) JobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) JobDefinitionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) JobResources() SagemakerModelQualityJobDefinitionJobResourcesOutputReference {
	var returns SagemakerModelQualityJobDefinitionJobResourcesOutputReference
	_jsii_.Get(
		j,
		"jobResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) JobResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) ModelQualityAppSpecification() SagemakerModelQualityJobDefinitionModelQualityAppSpecificationOutputReference {
	var returns SagemakerModelQualityJobDefinitionModelQualityAppSpecificationOutputReference
	_jsii_.Get(
		j,
		"modelQualityAppSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) ModelQualityAppSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelQualityAppSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) ModelQualityBaselineConfig() SagemakerModelQualityJobDefinitionModelQualityBaselineConfigOutputReference {
	var returns SagemakerModelQualityJobDefinitionModelQualityBaselineConfigOutputReference
	_jsii_.Get(
		j,
		"modelQualityBaselineConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) ModelQualityBaselineConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelQualityBaselineConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) ModelQualityJobInput() SagemakerModelQualityJobDefinitionModelQualityJobInputOutputReference {
	var returns SagemakerModelQualityJobDefinitionModelQualityJobInputOutputReference
	_jsii_.Get(
		j,
		"modelQualityJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) ModelQualityJobInputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelQualityJobInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) ModelQualityJobOutputConfig() SagemakerModelQualityJobDefinitionModelQualityJobOutputConfigOutputReference {
	var returns SagemakerModelQualityJobDefinitionModelQualityJobOutputConfigOutputReference
	_jsii_.Get(
		j,
		"modelQualityJobOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) ModelQualityJobOutputConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelQualityJobOutputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) NetworkConfig() SagemakerModelQualityJobDefinitionNetworkConfigOutputReference {
	var returns SagemakerModelQualityJobDefinitionNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) NetworkConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) StoppingCondition() SagemakerModelQualityJobDefinitionStoppingConditionOutputReference {
	var returns SagemakerModelQualityJobDefinitionStoppingConditionOutputReference
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) StoppingConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) Tags() SagemakerModelQualityJobDefinitionTagsList {
	var returns SagemakerModelQualityJobDefinitionTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_quality_job_definition awscc_sagemaker_model_quality_job_definition} Resource.
func NewSagemakerModelQualityJobDefinition(scope constructs.Construct, id *string, config *SagemakerModelQualityJobDefinitionConfig) SagemakerModelQualityJobDefinition {
	_init_.Initialize()

	if err := validateNewSagemakerModelQualityJobDefinitionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerModelQualityJobDefinition{}

	_jsii_.Create(
		"awscc.sagemakerModelQualityJobDefinition.SagemakerModelQualityJobDefinition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_quality_job_definition awscc_sagemaker_model_quality_job_definition} Resource.
func NewSagemakerModelQualityJobDefinition_Override(s SagemakerModelQualityJobDefinition, scope constructs.Construct, id *string, config *SagemakerModelQualityJobDefinitionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sagemakerModelQualityJobDefinition.SagemakerModelQualityJobDefinition",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition)SetEndpointName(val *string) {
	if err := j.validateSetEndpointNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointName",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition)SetJobDefinitionName(val *string) {
	if err := j.validateSetJobDefinitionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobDefinitionName",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelQualityJobDefinition)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Generates CDKTF code for importing a SagemakerModelQualityJobDefinition resource upon running "cdktf plan <stack-name>".
func SagemakerModelQualityJobDefinition_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSagemakerModelQualityJobDefinition_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.sagemakerModelQualityJobDefinition.SagemakerModelQualityJobDefinition",
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
func SagemakerModelQualityJobDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerModelQualityJobDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sagemakerModelQualityJobDefinition.SagemakerModelQualityJobDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerModelQualityJobDefinition_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerModelQualityJobDefinition_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sagemakerModelQualityJobDefinition.SagemakerModelQualityJobDefinition",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerModelQualityJobDefinition_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerModelQualityJobDefinition_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sagemakerModelQualityJobDefinition.SagemakerModelQualityJobDefinition",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SagemakerModelQualityJobDefinition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.sagemakerModelQualityJobDefinition.SagemakerModelQualityJobDefinition",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) PutJobResources(value *SagemakerModelQualityJobDefinitionJobResources) {
	if err := s.validatePutJobResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putJobResources",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) PutModelQualityAppSpecification(value *SagemakerModelQualityJobDefinitionModelQualityAppSpecification) {
	if err := s.validatePutModelQualityAppSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelQualityAppSpecification",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) PutModelQualityBaselineConfig(value *SagemakerModelQualityJobDefinitionModelQualityBaselineConfig) {
	if err := s.validatePutModelQualityBaselineConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelQualityBaselineConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) PutModelQualityJobInput(value *SagemakerModelQualityJobDefinitionModelQualityJobInput) {
	if err := s.validatePutModelQualityJobInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelQualityJobInput",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) PutModelQualityJobOutputConfig(value *SagemakerModelQualityJobDefinitionModelQualityJobOutputConfig) {
	if err := s.validatePutModelQualityJobOutputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelQualityJobOutputConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) PutNetworkConfig(value *SagemakerModelQualityJobDefinitionNetworkConfig) {
	if err := s.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) PutStoppingCondition(value *SagemakerModelQualityJobDefinitionStoppingCondition) {
	if err := s.validatePutStoppingConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStoppingCondition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) PutTags(value interface{}) {
	if err := s.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTags",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) ResetEndpointName() {
	_jsii_.InvokeVoid(
		s,
		"resetEndpointName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) ResetJobDefinitionName() {
	_jsii_.InvokeVoid(
		s,
		"resetJobDefinitionName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) ResetModelQualityBaselineConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetModelQualityBaselineConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) ResetStoppingCondition() {
	_jsii_.InvokeVoid(
		s,
		"resetStoppingCondition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelQualityJobDefinition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

