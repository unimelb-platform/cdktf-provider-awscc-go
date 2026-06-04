package cloudformationlambdahook

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cloudformationlambdahook/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudformation_lambda_hook awscc_cloudformation_lambda_hook}.
type CloudformationLambdaHook interface {
	cdktf.TerraformResource
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
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
	ExecutionRole() *string
	SetExecutionRole(val *string)
	ExecutionRoleInput() *string
	FailureMode() *string
	SetFailureMode(val *string)
	FailureModeInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HookArn() *string
	HookStatus() *string
	SetHookStatus(val *string)
	HookStatusInput() *string
	Id() *string
	LambdaFunction() *string
	SetLambdaFunction(val *string)
	LambdaFunctionInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	StackFilters() CloudformationLambdaHookStackFiltersOutputReference
	StackFiltersInput() interface{}
	TargetFilters() CloudformationLambdaHookTargetFiltersOutputReference
	TargetFiltersInput() interface{}
	TargetOperations() *[]*string
	SetTargetOperations(val *[]*string)
	TargetOperationsInput() *[]*string
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
	PutStackFilters(value *CloudformationLambdaHookStackFilters)
	PutTargetFilters(value *CloudformationLambdaHookTargetFilters)
	ResetHookStatus()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStackFilters()
	ResetTargetFilters()
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

// The jsii proxy struct for CloudformationLambdaHook
type jsiiProxy_CloudformationLambdaHook struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudformationLambdaHook) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) FailureMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) FailureModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) HookArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hookArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) HookStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hookStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) HookStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hookStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) LambdaFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) LambdaFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) StackFilters() CloudformationLambdaHookStackFiltersOutputReference {
	var returns CloudformationLambdaHookStackFiltersOutputReference
	_jsii_.Get(
		j,
		"stackFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) StackFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stackFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) TargetFilters() CloudformationLambdaHookTargetFiltersOutputReference {
	var returns CloudformationLambdaHookTargetFiltersOutputReference
	_jsii_.Get(
		j,
		"targetFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) TargetFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) TargetOperations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) TargetOperationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHook) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudformation_lambda_hook awscc_cloudformation_lambda_hook} Resource.
func NewCloudformationLambdaHook(scope constructs.Construct, id *string, config *CloudformationLambdaHookConfig) CloudformationLambdaHook {
	_init_.Initialize()

	if err := validateNewCloudformationLambdaHookParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudformationLambdaHook{}

	_jsii_.Create(
		"awscc.cloudformationLambdaHook.CloudformationLambdaHook",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudformation_lambda_hook awscc_cloudformation_lambda_hook} Resource.
func NewCloudformationLambdaHook_Override(c CloudformationLambdaHook, scope constructs.Construct, id *string, config *CloudformationLambdaHookConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cloudformationLambdaHook.CloudformationLambdaHook",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudformationLambdaHook)SetAlias(val *string) {
	if err := j.validateSetAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHook)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHook)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHook)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHook)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHook)SetFailureMode(val *string) {
	if err := j.validateSetFailureModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureMode",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHook)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHook)SetHookStatus(val *string) {
	if err := j.validateSetHookStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hookStatus",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHook)SetLambdaFunction(val *string) {
	if err := j.validateSetLambdaFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaFunction",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHook)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHook)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHook)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHook)SetTargetOperations(val *[]*string) {
	if err := j.validateSetTargetOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetOperations",
		val,
	)
}

// Generates CDKTF code for importing a CloudformationLambdaHook resource upon running "cdktf plan <stack-name>".
func CloudformationLambdaHook_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCloudformationLambdaHook_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.cloudformationLambdaHook.CloudformationLambdaHook",
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
func CloudformationLambdaHook_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudformationLambdaHook_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cloudformationLambdaHook.CloudformationLambdaHook",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudformationLambdaHook_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudformationLambdaHook_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cloudformationLambdaHook.CloudformationLambdaHook",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudformationLambdaHook_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudformationLambdaHook_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cloudformationLambdaHook.CloudformationLambdaHook",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudformationLambdaHook_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.cloudformationLambdaHook.CloudformationLambdaHook",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudformationLambdaHook) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CloudformationLambdaHook) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudformationLambdaHook) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudformationLambdaHook) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudformationLambdaHook) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudformationLambdaHook) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudformationLambdaHook) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudformationLambdaHook) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudformationLambdaHook) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudformationLambdaHook) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudformationLambdaHook) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudformationLambdaHook) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationLambdaHook) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CloudformationLambdaHook) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudformationLambdaHook) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudformationLambdaHook) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudformationLambdaHook) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudformationLambdaHook) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudformationLambdaHook) PutStackFilters(value *CloudformationLambdaHookStackFilters) {
	if err := c.validatePutStackFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStackFilters",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudformationLambdaHook) PutTargetFilters(value *CloudformationLambdaHookTargetFilters) {
	if err := c.validatePutTargetFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTargetFilters",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudformationLambdaHook) ResetHookStatus() {
	_jsii_.InvokeVoid(
		c,
		"resetHookStatus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationLambdaHook) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationLambdaHook) ResetStackFilters() {
	_jsii_.InvokeVoid(
		c,
		"resetStackFilters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationLambdaHook) ResetTargetFilters() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetFilters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationLambdaHook) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationLambdaHook) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationLambdaHook) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationLambdaHook) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationLambdaHook) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationLambdaHook) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

