package codedeploydeploymentconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/codedeploydeploymentconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config awscc_codedeploy_deployment_config}.
type CodedeployDeploymentConfig interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComputePlatform() *string
	SetComputePlatform(val *string)
	ComputePlatformInput() *string
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
	DeploymentConfigName() *string
	SetDeploymentConfigName(val *string)
	DeploymentConfigNameInput() *string
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
	MinimumHealthyHosts() CodedeployDeploymentConfigMinimumHealthyHostsOutputReference
	MinimumHealthyHostsInput() interface{}
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
	TrafficRoutingConfig() CodedeployDeploymentConfigTrafficRoutingConfigOutputReference
	TrafficRoutingConfigInput() interface{}
	ZonalConfig() CodedeployDeploymentConfigZonalConfigOutputReference
	ZonalConfigInput() interface{}
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
	PutMinimumHealthyHosts(value *CodedeployDeploymentConfigMinimumHealthyHosts)
	PutTrafficRoutingConfig(value *CodedeployDeploymentConfigTrafficRoutingConfig)
	PutZonalConfig(value *CodedeployDeploymentConfigZonalConfig)
	ResetComputePlatform()
	ResetDeploymentConfigName()
	ResetMinimumHealthyHosts()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTrafficRoutingConfig()
	ResetZonalConfig()
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

// The jsii proxy struct for CodedeployDeploymentConfig
type jsiiProxy_CodedeployDeploymentConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CodedeployDeploymentConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) ComputePlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) ComputePlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computePlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) DeploymentConfigNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) MinimumHealthyHosts() CodedeployDeploymentConfigMinimumHealthyHostsOutputReference {
	var returns CodedeployDeploymentConfigMinimumHealthyHostsOutputReference
	_jsii_.Get(
		j,
		"minimumHealthyHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) MinimumHealthyHostsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"minimumHealthyHostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) TrafficRoutingConfig() CodedeployDeploymentConfigTrafficRoutingConfigOutputReference {
	var returns CodedeployDeploymentConfigTrafficRoutingConfigOutputReference
	_jsii_.Get(
		j,
		"trafficRoutingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) TrafficRoutingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficRoutingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) ZonalConfig() CodedeployDeploymentConfigZonalConfigOutputReference {
	var returns CodedeployDeploymentConfigZonalConfigOutputReference
	_jsii_.Get(
		j,
		"zonalConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) ZonalConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zonalConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config awscc_codedeploy_deployment_config} Resource.
func NewCodedeployDeploymentConfig(scope constructs.Construct, id *string, config *CodedeployDeploymentConfigConfig) CodedeployDeploymentConfig {
	_init_.Initialize()

	if err := validateNewCodedeployDeploymentConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodedeployDeploymentConfig{}

	_jsii_.Create(
		"awscc.codedeployDeploymentConfig.CodedeployDeploymentConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config awscc_codedeploy_deployment_config} Resource.
func NewCodedeployDeploymentConfig_Override(c CodedeployDeploymentConfig, scope constructs.Construct, id *string, config *CodedeployDeploymentConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.codedeployDeploymentConfig.CodedeployDeploymentConfig",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfig)SetComputePlatform(val *string) {
	if err := j.validateSetComputePlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computePlatform",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfig)SetDeploymentConfigName(val *string) {
	if err := j.validateSetDeploymentConfigNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentConfigName",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a CodedeployDeploymentConfig resource upon running "cdktf plan <stack-name>".
func CodedeployDeploymentConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCodedeployDeploymentConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.codedeployDeploymentConfig.CodedeployDeploymentConfig",
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
func CodedeployDeploymentConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCodedeployDeploymentConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.codedeployDeploymentConfig.CodedeployDeploymentConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CodedeployDeploymentConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCodedeployDeploymentConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.codedeployDeploymentConfig.CodedeployDeploymentConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CodedeployDeploymentConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCodedeployDeploymentConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.codedeployDeploymentConfig.CodedeployDeploymentConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CodedeployDeploymentConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.codedeployDeploymentConfig.CodedeployDeploymentConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfig) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CodedeployDeploymentConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodedeployDeploymentConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CodedeployDeploymentConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CodedeployDeploymentConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CodedeployDeploymentConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CodedeployDeploymentConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CodedeployDeploymentConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CodedeployDeploymentConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CodedeployDeploymentConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodedeployDeploymentConfig) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) PutMinimumHealthyHosts(value *CodedeployDeploymentConfigMinimumHealthyHosts) {
	if err := c.validatePutMinimumHealthyHostsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMinimumHealthyHosts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) PutTrafficRoutingConfig(value *CodedeployDeploymentConfigTrafficRoutingConfig) {
	if err := c.validatePutTrafficRoutingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTrafficRoutingConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) PutZonalConfig(value *CodedeployDeploymentConfigZonalConfig) {
	if err := c.validatePutZonalConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putZonalConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) ResetComputePlatform() {
	_jsii_.InvokeVoid(
		c,
		"resetComputePlatform",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) ResetDeploymentConfigName() {
	_jsii_.InvokeVoid(
		c,
		"resetDeploymentConfigName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) ResetMinimumHealthyHosts() {
	_jsii_.InvokeVoid(
		c,
		"resetMinimumHealthyHosts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) ResetTrafficRoutingConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetTrafficRoutingConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) ResetZonalConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetZonalConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

