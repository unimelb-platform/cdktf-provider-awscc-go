package sagemakermodelexplainabilityjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sagemakermodelexplainabilityjobdefinition/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_explainability_job_definition awscc_sagemaker_model_explainability_job_definition}.
type SagemakerModelExplainabilityJobDefinition interface {
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
	JobResources() SagemakerModelExplainabilityJobDefinitionJobResourcesOutputReference
	JobResourcesInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ModelExplainabilityAppSpecification() SagemakerModelExplainabilityJobDefinitionModelExplainabilityAppSpecificationOutputReference
	ModelExplainabilityAppSpecificationInput() interface{}
	ModelExplainabilityBaselineConfig() SagemakerModelExplainabilityJobDefinitionModelExplainabilityBaselineConfigOutputReference
	ModelExplainabilityBaselineConfigInput() interface{}
	ModelExplainabilityJobInput() SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobInputOutputReference
	ModelExplainabilityJobInputInput() interface{}
	ModelExplainabilityJobOutputConfig() SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobOutputConfigOutputReference
	ModelExplainabilityJobOutputConfigInput() interface{}
	NetworkConfig() SagemakerModelExplainabilityJobDefinitionNetworkConfigOutputReference
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
	StoppingCondition() SagemakerModelExplainabilityJobDefinitionStoppingConditionOutputReference
	StoppingConditionInput() interface{}
	Tags() SagemakerModelExplainabilityJobDefinitionTagsList
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutJobResources(value *SagemakerModelExplainabilityJobDefinitionJobResources)
	PutModelExplainabilityAppSpecification(value *SagemakerModelExplainabilityJobDefinitionModelExplainabilityAppSpecification)
	PutModelExplainabilityBaselineConfig(value *SagemakerModelExplainabilityJobDefinitionModelExplainabilityBaselineConfig)
	PutModelExplainabilityJobInput(value *SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobInput)
	PutModelExplainabilityJobOutputConfig(value *SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobOutputConfig)
	PutNetworkConfig(value *SagemakerModelExplainabilityJobDefinitionNetworkConfig)
	PutStoppingCondition(value *SagemakerModelExplainabilityJobDefinitionStoppingCondition)
	PutTags(value interface{})
	ResetEndpointName()
	ResetJobDefinitionName()
	ResetModelExplainabilityBaselineConfig()
	ResetNetworkConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStoppingCondition()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SagemakerModelExplainabilityJobDefinition
type jsiiProxy_SagemakerModelExplainabilityJobDefinition struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) EndpointNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) JobDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) JobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) JobDefinitionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) JobResources() SagemakerModelExplainabilityJobDefinitionJobResourcesOutputReference {
	var returns SagemakerModelExplainabilityJobDefinitionJobResourcesOutputReference
	_jsii_.Get(
		j,
		"jobResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) JobResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ModelExplainabilityAppSpecification() SagemakerModelExplainabilityJobDefinitionModelExplainabilityAppSpecificationOutputReference {
	var returns SagemakerModelExplainabilityJobDefinitionModelExplainabilityAppSpecificationOutputReference
	_jsii_.Get(
		j,
		"modelExplainabilityAppSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ModelExplainabilityAppSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelExplainabilityAppSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ModelExplainabilityBaselineConfig() SagemakerModelExplainabilityJobDefinitionModelExplainabilityBaselineConfigOutputReference {
	var returns SagemakerModelExplainabilityJobDefinitionModelExplainabilityBaselineConfigOutputReference
	_jsii_.Get(
		j,
		"modelExplainabilityBaselineConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ModelExplainabilityBaselineConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelExplainabilityBaselineConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ModelExplainabilityJobInput() SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobInputOutputReference {
	var returns SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobInputOutputReference
	_jsii_.Get(
		j,
		"modelExplainabilityJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ModelExplainabilityJobInputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelExplainabilityJobInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ModelExplainabilityJobOutputConfig() SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobOutputConfigOutputReference {
	var returns SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobOutputConfigOutputReference
	_jsii_.Get(
		j,
		"modelExplainabilityJobOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ModelExplainabilityJobOutputConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelExplainabilityJobOutputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) NetworkConfig() SagemakerModelExplainabilityJobDefinitionNetworkConfigOutputReference {
	var returns SagemakerModelExplainabilityJobDefinitionNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) NetworkConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) StoppingCondition() SagemakerModelExplainabilityJobDefinitionStoppingConditionOutputReference {
	var returns SagemakerModelExplainabilityJobDefinitionStoppingConditionOutputReference
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) StoppingConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) Tags() SagemakerModelExplainabilityJobDefinitionTagsList {
	var returns SagemakerModelExplainabilityJobDefinitionTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_explainability_job_definition awscc_sagemaker_model_explainability_job_definition} Resource.
func NewSagemakerModelExplainabilityJobDefinition(scope constructs.Construct, id *string, config *SagemakerModelExplainabilityJobDefinitionConfig) SagemakerModelExplainabilityJobDefinition {
	_init_.Initialize()

	if err := validateNewSagemakerModelExplainabilityJobDefinitionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerModelExplainabilityJobDefinition{}

	_jsii_.Create(
		"awscc.sagemakerModelExplainabilityJobDefinition.SagemakerModelExplainabilityJobDefinition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_explainability_job_definition awscc_sagemaker_model_explainability_job_definition} Resource.
func NewSagemakerModelExplainabilityJobDefinition_Override(s SagemakerModelExplainabilityJobDefinition, scope constructs.Construct, id *string, config *SagemakerModelExplainabilityJobDefinitionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sagemakerModelExplainabilityJobDefinition.SagemakerModelExplainabilityJobDefinition",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition)SetEndpointName(val *string) {
	if err := j.validateSetEndpointNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointName",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition)SetJobDefinitionName(val *string) {
	if err := j.validateSetJobDefinitionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobDefinitionName",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelExplainabilityJobDefinition)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Generates CDKTF code for importing a SagemakerModelExplainabilityJobDefinition resource upon running "cdktf plan <stack-name>".
func SagemakerModelExplainabilityJobDefinition_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSagemakerModelExplainabilityJobDefinition_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.sagemakerModelExplainabilityJobDefinition.SagemakerModelExplainabilityJobDefinition",
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
func SagemakerModelExplainabilityJobDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerModelExplainabilityJobDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sagemakerModelExplainabilityJobDefinition.SagemakerModelExplainabilityJobDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerModelExplainabilityJobDefinition_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerModelExplainabilityJobDefinition_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sagemakerModelExplainabilityJobDefinition.SagemakerModelExplainabilityJobDefinition",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerModelExplainabilityJobDefinition_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerModelExplainabilityJobDefinition_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sagemakerModelExplainabilityJobDefinition.SagemakerModelExplainabilityJobDefinition",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SagemakerModelExplainabilityJobDefinition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.sagemakerModelExplainabilityJobDefinition.SagemakerModelExplainabilityJobDefinition",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) PutJobResources(value *SagemakerModelExplainabilityJobDefinitionJobResources) {
	if err := s.validatePutJobResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putJobResources",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) PutModelExplainabilityAppSpecification(value *SagemakerModelExplainabilityJobDefinitionModelExplainabilityAppSpecification) {
	if err := s.validatePutModelExplainabilityAppSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelExplainabilityAppSpecification",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) PutModelExplainabilityBaselineConfig(value *SagemakerModelExplainabilityJobDefinitionModelExplainabilityBaselineConfig) {
	if err := s.validatePutModelExplainabilityBaselineConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelExplainabilityBaselineConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) PutModelExplainabilityJobInput(value *SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobInput) {
	if err := s.validatePutModelExplainabilityJobInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelExplainabilityJobInput",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) PutModelExplainabilityJobOutputConfig(value *SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobOutputConfig) {
	if err := s.validatePutModelExplainabilityJobOutputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelExplainabilityJobOutputConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) PutNetworkConfig(value *SagemakerModelExplainabilityJobDefinitionNetworkConfig) {
	if err := s.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) PutStoppingCondition(value *SagemakerModelExplainabilityJobDefinitionStoppingCondition) {
	if err := s.validatePutStoppingConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStoppingCondition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) PutTags(value interface{}) {
	if err := s.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTags",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ResetEndpointName() {
	_jsii_.InvokeVoid(
		s,
		"resetEndpointName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ResetJobDefinitionName() {
	_jsii_.InvokeVoid(
		s,
		"resetJobDefinitionName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ResetModelExplainabilityBaselineConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetModelExplainabilityBaselineConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ResetStoppingCondition() {
	_jsii_.InvokeVoid(
		s,
		"resetStoppingCondition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelExplainabilityJobDefinition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

