package entityresolutionmatchingworkflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/entityresolutionmatchingworkflow/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_matching_workflow awscc_entityresolution_matching_workflow}.
type EntityresolutionMatchingWorkflow interface {
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
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IncrementalRunConfig() EntityresolutionMatchingWorkflowIncrementalRunConfigOutputReference
	IncrementalRunConfigInput() interface{}
	InputSourceConfig() EntityresolutionMatchingWorkflowInputSourceConfigList
	InputSourceConfigInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OutputSourceConfig() EntityresolutionMatchingWorkflowOutputSourceConfigList
	OutputSourceConfigInput() interface{}
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
	ResolutionTechniques() EntityresolutionMatchingWorkflowResolutionTechniquesOutputReference
	ResolutionTechniquesInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Tags() EntityresolutionMatchingWorkflowTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatedAt() *string
	WorkflowArn() *string
	WorkflowName() *string
	SetWorkflowName(val *string)
	WorkflowNameInput() *string
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
	PutIncrementalRunConfig(value *EntityresolutionMatchingWorkflowIncrementalRunConfig)
	PutInputSourceConfig(value interface{})
	PutOutputSourceConfig(value interface{})
	PutResolutionTechniques(value *EntityresolutionMatchingWorkflowResolutionTechniques)
	PutTags(value interface{})
	ResetDescription()
	ResetIncrementalRunConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for EntityresolutionMatchingWorkflow
type jsiiProxy_EntityresolutionMatchingWorkflow struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) IncrementalRunConfig() EntityresolutionMatchingWorkflowIncrementalRunConfigOutputReference {
	var returns EntityresolutionMatchingWorkflowIncrementalRunConfigOutputReference
	_jsii_.Get(
		j,
		"incrementalRunConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) IncrementalRunConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incrementalRunConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) InputSourceConfig() EntityresolutionMatchingWorkflowInputSourceConfigList {
	var returns EntityresolutionMatchingWorkflowInputSourceConfigList
	_jsii_.Get(
		j,
		"inputSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) InputSourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) OutputSourceConfig() EntityresolutionMatchingWorkflowOutputSourceConfigList {
	var returns EntityresolutionMatchingWorkflowOutputSourceConfigList
	_jsii_.Get(
		j,
		"outputSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) OutputSourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) ResolutionTechniques() EntityresolutionMatchingWorkflowResolutionTechniquesOutputReference {
	var returns EntityresolutionMatchingWorkflowResolutionTechniquesOutputReference
	_jsii_.Get(
		j,
		"resolutionTechniques",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) ResolutionTechniquesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resolutionTechniquesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) Tags() EntityresolutionMatchingWorkflowTagsList {
	var returns EntityresolutionMatchingWorkflowTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) WorkflowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) WorkflowName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow) WorkflowNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_matching_workflow awscc_entityresolution_matching_workflow} Resource.
func NewEntityresolutionMatchingWorkflow(scope constructs.Construct, id *string, config *EntityresolutionMatchingWorkflowConfig) EntityresolutionMatchingWorkflow {
	_init_.Initialize()

	if err := validateNewEntityresolutionMatchingWorkflowParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EntityresolutionMatchingWorkflow{}

	_jsii_.Create(
		"awscc.entityresolutionMatchingWorkflow.EntityresolutionMatchingWorkflow",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_matching_workflow awscc_entityresolution_matching_workflow} Resource.
func NewEntityresolutionMatchingWorkflow_Override(e EntityresolutionMatchingWorkflow, scope constructs.Construct, id *string, config *EntityresolutionMatchingWorkflowConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.entityresolutionMatchingWorkflow.EntityresolutionMatchingWorkflow",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflow)SetWorkflowName(val *string) {
	if err := j.validateSetWorkflowNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workflowName",
		val,
	)
}

// Generates CDKTF code for importing a EntityresolutionMatchingWorkflow resource upon running "cdktf plan <stack-name>".
func EntityresolutionMatchingWorkflow_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEntityresolutionMatchingWorkflow_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.entityresolutionMatchingWorkflow.EntityresolutionMatchingWorkflow",
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
func EntityresolutionMatchingWorkflow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEntityresolutionMatchingWorkflow_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.entityresolutionMatchingWorkflow.EntityresolutionMatchingWorkflow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EntityresolutionMatchingWorkflow_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEntityresolutionMatchingWorkflow_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.entityresolutionMatchingWorkflow.EntityresolutionMatchingWorkflow",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EntityresolutionMatchingWorkflow_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEntityresolutionMatchingWorkflow_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.entityresolutionMatchingWorkflow.EntityresolutionMatchingWorkflow",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EntityresolutionMatchingWorkflow_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.entityresolutionMatchingWorkflow.EntityresolutionMatchingWorkflow",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) PutIncrementalRunConfig(value *EntityresolutionMatchingWorkflowIncrementalRunConfig) {
	if err := e.validatePutIncrementalRunConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIncrementalRunConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) PutInputSourceConfig(value interface{}) {
	if err := e.validatePutInputSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInputSourceConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) PutOutputSourceConfig(value interface{}) {
	if err := e.validatePutOutputSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOutputSourceConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) PutResolutionTechniques(value *EntityresolutionMatchingWorkflowResolutionTechniques) {
	if err := e.validatePutResolutionTechniquesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putResolutionTechniques",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) ResetIncrementalRunConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetIncrementalRunConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflow) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

