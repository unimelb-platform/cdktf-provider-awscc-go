package sagemakerdataqualityjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sagemakerdataqualityjobdefinition/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition awscc_sagemaker_data_quality_job_definition}.
type SagemakerDataQualityJobDefinition interface {
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
	DataQualityAppSpecification() SagemakerDataQualityJobDefinitionDataQualityAppSpecificationOutputReference
	DataQualityAppSpecificationInput() interface{}
	DataQualityBaselineConfig() SagemakerDataQualityJobDefinitionDataQualityBaselineConfigOutputReference
	DataQualityBaselineConfigInput() interface{}
	DataQualityJobInput() SagemakerDataQualityJobDefinitionDataQualityJobInputOutputReference
	DataQualityJobInputInput() interface{}
	DataQualityJobOutputConfig() SagemakerDataQualityJobDefinitionDataQualityJobOutputConfigOutputReference
	DataQualityJobOutputConfigInput() interface{}
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
	JobResources() SagemakerDataQualityJobDefinitionJobResourcesOutputReference
	JobResourcesInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NetworkConfig() SagemakerDataQualityJobDefinitionNetworkConfigOutputReference
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
	StoppingCondition() SagemakerDataQualityJobDefinitionStoppingConditionOutputReference
	StoppingConditionInput() interface{}
	Tags() SagemakerDataQualityJobDefinitionTagsList
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
	PutDataQualityAppSpecification(value *SagemakerDataQualityJobDefinitionDataQualityAppSpecification)
	PutDataQualityBaselineConfig(value *SagemakerDataQualityJobDefinitionDataQualityBaselineConfig)
	PutDataQualityJobInput(value *SagemakerDataQualityJobDefinitionDataQualityJobInput)
	PutDataQualityJobOutputConfig(value *SagemakerDataQualityJobDefinitionDataQualityJobOutputConfig)
	PutJobResources(value *SagemakerDataQualityJobDefinitionJobResources)
	PutNetworkConfig(value *SagemakerDataQualityJobDefinitionNetworkConfig)
	PutStoppingCondition(value *SagemakerDataQualityJobDefinitionStoppingCondition)
	PutTags(value interface{})
	ResetDataQualityBaselineConfig()
	ResetEndpointName()
	ResetJobDefinitionName()
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

// The jsii proxy struct for SagemakerDataQualityJobDefinition
type jsiiProxy_SagemakerDataQualityJobDefinition struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) DataQualityAppSpecification() SagemakerDataQualityJobDefinitionDataQualityAppSpecificationOutputReference {
	var returns SagemakerDataQualityJobDefinitionDataQualityAppSpecificationOutputReference
	_jsii_.Get(
		j,
		"dataQualityAppSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) DataQualityAppSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataQualityAppSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) DataQualityBaselineConfig() SagemakerDataQualityJobDefinitionDataQualityBaselineConfigOutputReference {
	var returns SagemakerDataQualityJobDefinitionDataQualityBaselineConfigOutputReference
	_jsii_.Get(
		j,
		"dataQualityBaselineConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) DataQualityBaselineConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataQualityBaselineConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) DataQualityJobInput() SagemakerDataQualityJobDefinitionDataQualityJobInputOutputReference {
	var returns SagemakerDataQualityJobDefinitionDataQualityJobInputOutputReference
	_jsii_.Get(
		j,
		"dataQualityJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) DataQualityJobInputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataQualityJobInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) DataQualityJobOutputConfig() SagemakerDataQualityJobDefinitionDataQualityJobOutputConfigOutputReference {
	var returns SagemakerDataQualityJobDefinitionDataQualityJobOutputConfigOutputReference
	_jsii_.Get(
		j,
		"dataQualityJobOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) DataQualityJobOutputConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataQualityJobOutputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) EndpointNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) JobDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) JobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) JobDefinitionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) JobResources() SagemakerDataQualityJobDefinitionJobResourcesOutputReference {
	var returns SagemakerDataQualityJobDefinitionJobResourcesOutputReference
	_jsii_.Get(
		j,
		"jobResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) JobResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) NetworkConfig() SagemakerDataQualityJobDefinitionNetworkConfigOutputReference {
	var returns SagemakerDataQualityJobDefinitionNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) NetworkConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) StoppingCondition() SagemakerDataQualityJobDefinitionStoppingConditionOutputReference {
	var returns SagemakerDataQualityJobDefinitionStoppingConditionOutputReference
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) StoppingConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) Tags() SagemakerDataQualityJobDefinitionTagsList {
	var returns SagemakerDataQualityJobDefinitionTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition awscc_sagemaker_data_quality_job_definition} Resource.
func NewSagemakerDataQualityJobDefinition(scope constructs.Construct, id *string, config *SagemakerDataQualityJobDefinitionConfig) SagemakerDataQualityJobDefinition {
	_init_.Initialize()

	if err := validateNewSagemakerDataQualityJobDefinitionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerDataQualityJobDefinition{}

	_jsii_.Create(
		"awscc.sagemakerDataQualityJobDefinition.SagemakerDataQualityJobDefinition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition awscc_sagemaker_data_quality_job_definition} Resource.
func NewSagemakerDataQualityJobDefinition_Override(s SagemakerDataQualityJobDefinition, scope constructs.Construct, id *string, config *SagemakerDataQualityJobDefinitionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sagemakerDataQualityJobDefinition.SagemakerDataQualityJobDefinition",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition)SetEndpointName(val *string) {
	if err := j.validateSetEndpointNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointName",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition)SetJobDefinitionName(val *string) {
	if err := j.validateSetJobDefinitionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobDefinitionName",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinition)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Generates CDKTF code for importing a SagemakerDataQualityJobDefinition resource upon running "cdktf plan <stack-name>".
func SagemakerDataQualityJobDefinition_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSagemakerDataQualityJobDefinition_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.sagemakerDataQualityJobDefinition.SagemakerDataQualityJobDefinition",
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
func SagemakerDataQualityJobDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerDataQualityJobDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sagemakerDataQualityJobDefinition.SagemakerDataQualityJobDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerDataQualityJobDefinition_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerDataQualityJobDefinition_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sagemakerDataQualityJobDefinition.SagemakerDataQualityJobDefinition",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerDataQualityJobDefinition_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerDataQualityJobDefinition_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sagemakerDataQualityJobDefinition.SagemakerDataQualityJobDefinition",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SagemakerDataQualityJobDefinition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.sagemakerDataQualityJobDefinition.SagemakerDataQualityJobDefinition",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) PutDataQualityAppSpecification(value *SagemakerDataQualityJobDefinitionDataQualityAppSpecification) {
	if err := s.validatePutDataQualityAppSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDataQualityAppSpecification",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) PutDataQualityBaselineConfig(value *SagemakerDataQualityJobDefinitionDataQualityBaselineConfig) {
	if err := s.validatePutDataQualityBaselineConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDataQualityBaselineConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) PutDataQualityJobInput(value *SagemakerDataQualityJobDefinitionDataQualityJobInput) {
	if err := s.validatePutDataQualityJobInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDataQualityJobInput",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) PutDataQualityJobOutputConfig(value *SagemakerDataQualityJobDefinitionDataQualityJobOutputConfig) {
	if err := s.validatePutDataQualityJobOutputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDataQualityJobOutputConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) PutJobResources(value *SagemakerDataQualityJobDefinitionJobResources) {
	if err := s.validatePutJobResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putJobResources",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) PutNetworkConfig(value *SagemakerDataQualityJobDefinitionNetworkConfig) {
	if err := s.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) PutStoppingCondition(value *SagemakerDataQualityJobDefinitionStoppingCondition) {
	if err := s.validatePutStoppingConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStoppingCondition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) PutTags(value interface{}) {
	if err := s.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTags",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) ResetDataQualityBaselineConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetDataQualityBaselineConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) ResetEndpointName() {
	_jsii_.InvokeVoid(
		s,
		"resetEndpointName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) ResetJobDefinitionName() {
	_jsii_.InvokeVoid(
		s,
		"resetJobDefinitionName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) ResetStoppingCondition() {
	_jsii_.InvokeVoid(
		s,
		"resetStoppingCondition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

