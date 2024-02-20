package cloudformationstackset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cloudformationstackset/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_stack_set awscc_cloudformation_stack_set}.
type CloudformationStackSet interface {
	cdktf.TerraformResource
	AdministrationRoleArn() *string
	SetAdministrationRoleArn(val *string)
	AdministrationRoleArnInput() *string
	AutoDeployment() CloudformationStackSetAutoDeploymentOutputReference
	AutoDeploymentInput() interface{}
	CallAs() *string
	SetCallAs(val *string)
	CallAsInput() *string
	Capabilities() *[]*string
	SetCapabilities(val *[]*string)
	CapabilitiesInput() *[]*string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExecutionRoleName() *string
	SetExecutionRoleName(val *string)
	ExecutionRoleNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManagedExecution() CloudformationStackSetManagedExecutionOutputReference
	ManagedExecutionInput() interface{}
	// The tree node.
	Node() constructs.Node
	OperationPreferences() CloudformationStackSetOperationPreferencesOutputReference
	OperationPreferencesInput() interface{}
	Parameters() CloudformationStackSetParametersList
	ParametersInput() interface{}
	PermissionModel() *string
	SetPermissionModel(val *string)
	PermissionModelInput() *string
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
	StackInstancesGroup() CloudformationStackSetStackInstancesGroupList
	StackInstancesGroupInput() interface{}
	StackSetId() *string
	StackSetName() *string
	SetStackSetName(val *string)
	StackSetNameInput() *string
	Tags() CloudformationStackSetTagsList
	TagsInput() interface{}
	TemplateBody() *string
	SetTemplateBody(val *string)
	TemplateBodyInput() *string
	TemplateUrl() *string
	SetTemplateUrl(val *string)
	TemplateUrlInput() *string
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
	PutAutoDeployment(value *CloudformationStackSetAutoDeployment)
	PutManagedExecution(value *CloudformationStackSetManagedExecution)
	PutOperationPreferences(value *CloudformationStackSetOperationPreferences)
	PutParameters(value interface{})
	PutStackInstancesGroup(value interface{})
	PutTags(value interface{})
	ResetAdministrationRoleArn()
	ResetAutoDeployment()
	ResetCallAs()
	ResetCapabilities()
	ResetDescription()
	ResetExecutionRoleName()
	ResetManagedExecution()
	ResetOperationPreferences()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetStackInstancesGroup()
	ResetTags()
	ResetTemplateBody()
	ResetTemplateUrl()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudformationStackSet
type jsiiProxy_CloudformationStackSet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudformationStackSet) AdministrationRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administrationRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) AdministrationRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administrationRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) AutoDeployment() CloudformationStackSetAutoDeploymentOutputReference {
	var returns CloudformationStackSetAutoDeploymentOutputReference
	_jsii_.Get(
		j,
		"autoDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) AutoDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) CallAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) CallAsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) Capabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) CapabilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) ExecutionRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) ExecutionRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) ManagedExecution() CloudformationStackSetManagedExecutionOutputReference {
	var returns CloudformationStackSetManagedExecutionOutputReference
	_jsii_.Get(
		j,
		"managedExecution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) ManagedExecutionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedExecutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) OperationPreferences() CloudformationStackSetOperationPreferencesOutputReference {
	var returns CloudformationStackSetOperationPreferencesOutputReference
	_jsii_.Get(
		j,
		"operationPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) OperationPreferencesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationPreferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) Parameters() CloudformationStackSetParametersList {
	var returns CloudformationStackSetParametersList
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) ParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) PermissionModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) PermissionModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) StackInstancesGroup() CloudformationStackSetStackInstancesGroupList {
	var returns CloudformationStackSetStackInstancesGroupList
	_jsii_.Get(
		j,
		"stackInstancesGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) StackInstancesGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stackInstancesGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) StackSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) StackSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) StackSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) Tags() CloudformationStackSetTagsList {
	var returns CloudformationStackSetTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) TemplateBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) TemplateBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) TemplateUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) TemplateUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_stack_set awscc_cloudformation_stack_set} Resource.
func NewCloudformationStackSet(scope constructs.Construct, id *string, config *CloudformationStackSetConfig) CloudformationStackSet {
	_init_.Initialize()

	if err := validateNewCloudformationStackSetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudformationStackSet{}

	_jsii_.Create(
		"awscc.cloudformationStackSet.CloudformationStackSet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_stack_set awscc_cloudformation_stack_set} Resource.
func NewCloudformationStackSet_Override(c CloudformationStackSet, scope constructs.Construct, id *string, config *CloudformationStackSetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cloudformationStackSet.CloudformationStackSet",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudformationStackSet)SetAdministrationRoleArn(val *string) {
	if err := j.validateSetAdministrationRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administrationRoleArn",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSet)SetCallAs(val *string) {
	if err := j.validateSetCallAsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"callAs",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSet)SetCapabilities(val *[]*string) {
	if err := j.validateSetCapabilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capabilities",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSet)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSet)SetExecutionRoleName(val *string) {
	if err := j.validateSetExecutionRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleName",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSet)SetPermissionModel(val *string) {
	if err := j.validateSetPermissionModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissionModel",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSet)SetStackSetName(val *string) {
	if err := j.validateSetStackSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackSetName",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSet)SetTemplateBody(val *string) {
	if err := j.validateSetTemplateBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateBody",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSet)SetTemplateUrl(val *string) {
	if err := j.validateSetTemplateUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateUrl",
		val,
	)
}

// Generates CDKTF code for importing a CloudformationStackSet resource upon running "cdktf plan <stack-name>".
func CloudformationStackSet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCloudformationStackSet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.cloudformationStackSet.CloudformationStackSet",
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
func CloudformationStackSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudformationStackSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cloudformationStackSet.CloudformationStackSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudformationStackSet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudformationStackSet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cloudformationStackSet.CloudformationStackSet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudformationStackSet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudformationStackSet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cloudformationStackSet.CloudformationStackSet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudformationStackSet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.cloudformationStackSet.CloudformationStackSet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudformationStackSet) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CloudformationStackSet) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudformationStackSet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudformationStackSet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudformationStackSet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudformationStackSet) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudformationStackSet) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudformationStackSet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudformationStackSet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudformationStackSet) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudformationStackSet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudformationStackSet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CloudformationStackSet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudformationStackSet) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudformationStackSet) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudformationStackSet) PutAutoDeployment(value *CloudformationStackSetAutoDeployment) {
	if err := c.validatePutAutoDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAutoDeployment",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudformationStackSet) PutManagedExecution(value *CloudformationStackSetManagedExecution) {
	if err := c.validatePutManagedExecutionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putManagedExecution",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudformationStackSet) PutOperationPreferences(value *CloudformationStackSetOperationPreferences) {
	if err := c.validatePutOperationPreferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOperationPreferences",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudformationStackSet) PutParameters(value interface{}) {
	if err := c.validatePutParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putParameters",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudformationStackSet) PutStackInstancesGroup(value interface{}) {
	if err := c.validatePutStackInstancesGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStackInstancesGroup",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudformationStackSet) PutTags(value interface{}) {
	if err := c.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTags",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudformationStackSet) ResetAdministrationRoleArn() {
	_jsii_.InvokeVoid(
		c,
		"resetAdministrationRoleArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSet) ResetAutoDeployment() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoDeployment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSet) ResetCallAs() {
	_jsii_.InvokeVoid(
		c,
		"resetCallAs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSet) ResetCapabilities() {
	_jsii_.InvokeVoid(
		c,
		"resetCapabilities",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSet) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSet) ResetExecutionRoleName() {
	_jsii_.InvokeVoid(
		c,
		"resetExecutionRoleName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSet) ResetManagedExecution() {
	_jsii_.InvokeVoid(
		c,
		"resetManagedExecution",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSet) ResetOperationPreferences() {
	_jsii_.InvokeVoid(
		c,
		"resetOperationPreferences",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSet) ResetParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSet) ResetStackInstancesGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetStackInstancesGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSet) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSet) ResetTemplateBody() {
	_jsii_.InvokeVoid(
		c,
		"resetTemplateBody",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSet) ResetTemplateUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetTemplateUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

