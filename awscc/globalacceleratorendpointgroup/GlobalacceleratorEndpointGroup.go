package globalacceleratorendpointgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/globalacceleratorendpointgroup/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_endpoint_group awscc_globalaccelerator_endpoint_group}.
type GlobalacceleratorEndpointGroup interface {
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EndpointConfigurations() GlobalacceleratorEndpointGroupEndpointConfigurationsList
	EndpointConfigurationsInput() interface{}
	EndpointGroupArn() *string
	EndpointGroupRegion() *string
	SetEndpointGroupRegion(val *string)
	EndpointGroupRegionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthCheckIntervalSeconds() *float64
	SetHealthCheckIntervalSeconds(val *float64)
	HealthCheckIntervalSecondsInput() *float64
	HealthCheckPath() *string
	SetHealthCheckPath(val *string)
	HealthCheckPathInput() *string
	HealthCheckPort() *float64
	SetHealthCheckPort(val *float64)
	HealthCheckPortInput() *float64
	HealthCheckProtocol() *string
	SetHealthCheckProtocol(val *string)
	HealthCheckProtocolInput() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ListenerArn() *string
	SetListenerArn(val *string)
	ListenerArnInput() *string
	// The tree node.
	Node() constructs.Node
	PortOverrides() GlobalacceleratorEndpointGroupPortOverridesList
	PortOverridesInput() interface{}
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
	ThresholdCount() *float64
	SetThresholdCount(val *float64)
	ThresholdCountInput() *float64
	TrafficDialPercentage() *float64
	SetTrafficDialPercentage(val *float64)
	TrafficDialPercentageInput() *float64
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
	PutEndpointConfigurations(value interface{})
	PutPortOverrides(value interface{})
	ResetEndpointConfigurations()
	ResetHealthCheckIntervalSeconds()
	ResetHealthCheckPath()
	ResetHealthCheckPort()
	ResetHealthCheckProtocol()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPortOverrides()
	ResetThresholdCount()
	ResetTrafficDialPercentage()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GlobalacceleratorEndpointGroup
type jsiiProxy_GlobalacceleratorEndpointGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) EndpointConfigurations() GlobalacceleratorEndpointGroupEndpointConfigurationsList {
	var returns GlobalacceleratorEndpointGroupEndpointConfigurationsList
	_jsii_.Get(
		j,
		"endpointConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) EndpointConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) EndpointGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) EndpointGroupRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) EndpointGroupRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ListenerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) PortOverrides() GlobalacceleratorEndpointGroupPortOverridesList {
	var returns GlobalacceleratorEndpointGroupPortOverridesList
	_jsii_.Get(
		j,
		"portOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) PortOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ThresholdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ThresholdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TrafficDialPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficDialPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TrafficDialPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficDialPercentageInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_endpoint_group awscc_globalaccelerator_endpoint_group} Resource.
func NewGlobalacceleratorEndpointGroup(scope constructs.Construct, id *string, config *GlobalacceleratorEndpointGroupConfig) GlobalacceleratorEndpointGroup {
	_init_.Initialize()

	if err := validateNewGlobalacceleratorEndpointGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlobalacceleratorEndpointGroup{}

	_jsii_.Create(
		"awscc.globalacceleratorEndpointGroup.GlobalacceleratorEndpointGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_endpoint_group awscc_globalaccelerator_endpoint_group} Resource.
func NewGlobalacceleratorEndpointGroup_Override(g GlobalacceleratorEndpointGroup, scope constructs.Construct, id *string, config *GlobalacceleratorEndpointGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.globalacceleratorEndpointGroup.GlobalacceleratorEndpointGroup",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetEndpointGroupRegion(val *string) {
	if err := j.validateSetEndpointGroupRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointGroupRegion",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetHealthCheckIntervalSeconds(val *float64) {
	if err := j.validateSetHealthCheckIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetHealthCheckPath(val *string) {
	if err := j.validateSetHealthCheckPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetHealthCheckPort(val *float64) {
	if err := j.validateSetHealthCheckPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckPort",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetHealthCheckProtocol(val *string) {
	if err := j.validateSetHealthCheckProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckProtocol",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetListenerArn(val *string) {
	if err := j.validateSetListenerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listenerArn",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetThresholdCount(val *float64) {
	if err := j.validateSetThresholdCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdCount",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetTrafficDialPercentage(val *float64) {
	if err := j.validateSetTrafficDialPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficDialPercentage",
		val,
	)
}

// Generates CDKTF code for importing a GlobalacceleratorEndpointGroup resource upon running "cdktf plan <stack-name>".
func GlobalacceleratorEndpointGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGlobalacceleratorEndpointGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.globalacceleratorEndpointGroup.GlobalacceleratorEndpointGroup",
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
func GlobalacceleratorEndpointGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlobalacceleratorEndpointGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.globalacceleratorEndpointGroup.GlobalacceleratorEndpointGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GlobalacceleratorEndpointGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlobalacceleratorEndpointGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.globalacceleratorEndpointGroup.GlobalacceleratorEndpointGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GlobalacceleratorEndpointGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlobalacceleratorEndpointGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.globalacceleratorEndpointGroup.GlobalacceleratorEndpointGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlobalacceleratorEndpointGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.globalacceleratorEndpointGroup.GlobalacceleratorEndpointGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) PutEndpointConfigurations(value interface{}) {
	if err := g.validatePutEndpointConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEndpointConfigurations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) PutPortOverrides(value interface{}) {
	if err := g.validatePutPortOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPortOverrides",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetEndpointConfigurations() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpointConfigurations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetHealthCheckIntervalSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheckIntervalSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetHealthCheckPath() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheckPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetHealthCheckPort() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheckPort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetHealthCheckProtocol() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheckProtocol",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetPortOverrides() {
	_jsii_.InvokeVoid(
		g,
		"resetPortOverrides",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetThresholdCount() {
	_jsii_.InvokeVoid(
		g,
		"resetThresholdCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetTrafficDialPercentage() {
	_jsii_.InvokeVoid(
		g,
		"resetTrafficDialPercentage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

