package vpclatticeresourceconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/vpclatticeresourceconfiguration/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_resource_configuration awscc_vpclattice_resource_configuration}.
type VpclatticeResourceConfiguration interface {
	cdktf.TerraformResource
	AllowAssociationToSharableServiceNetwork() interface{}
	SetAllowAssociationToSharableServiceNetwork(val interface{})
	AllowAssociationToSharableServiceNetworkInput() interface{}
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PortRanges() *[]*string
	SetPortRanges(val *[]*string)
	PortRangesInput() *[]*string
	ProtocolType() *string
	SetProtocolType(val *string)
	ProtocolTypeInput() *string
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
	ResourceConfigurationAuthType() *string
	SetResourceConfigurationAuthType(val *string)
	ResourceConfigurationAuthTypeInput() *string
	ResourceConfigurationDefinition() VpclatticeResourceConfigurationResourceConfigurationDefinitionOutputReference
	ResourceConfigurationDefinitionInput() interface{}
	ResourceConfigurationGroupId() *string
	SetResourceConfigurationGroupId(val *string)
	ResourceConfigurationGroupIdInput() *string
	ResourceConfigurationId() *string
	ResourceConfigurationType() *string
	SetResourceConfigurationType(val *string)
	ResourceConfigurationTypeInput() *string
	ResourceGatewayId() *string
	SetResourceGatewayId(val *string)
	ResourceGatewayIdInput() *string
	Tags() VpclatticeResourceConfigurationTagsList
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
	PutResourceConfigurationDefinition(value *VpclatticeResourceConfigurationResourceConfigurationDefinition)
	PutTags(value interface{})
	ResetAllowAssociationToSharableServiceNetwork()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPortRanges()
	ResetProtocolType()
	ResetResourceConfigurationAuthType()
	ResetResourceConfigurationDefinition()
	ResetResourceConfigurationGroupId()
	ResetResourceGatewayId()
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

// The jsii proxy struct for VpclatticeResourceConfiguration
type jsiiProxy_VpclatticeResourceConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) AllowAssociationToSharableServiceNetwork() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAssociationToSharableServiceNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) AllowAssociationToSharableServiceNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAssociationToSharableServiceNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) PortRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"portRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) PortRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"portRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) ProtocolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) ProtocolTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) ResourceConfigurationAuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceConfigurationAuthType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) ResourceConfigurationAuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceConfigurationAuthTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) ResourceConfigurationDefinition() VpclatticeResourceConfigurationResourceConfigurationDefinitionOutputReference {
	var returns VpclatticeResourceConfigurationResourceConfigurationDefinitionOutputReference
	_jsii_.Get(
		j,
		"resourceConfigurationDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) ResourceConfigurationDefinitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceConfigurationDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) ResourceConfigurationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceConfigurationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) ResourceConfigurationGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceConfigurationGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) ResourceConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) ResourceConfigurationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceConfigurationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) ResourceConfigurationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceConfigurationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) ResourceGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) ResourceGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) Tags() VpclatticeResourceConfigurationTagsList {
	var returns VpclatticeResourceConfigurationTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeResourceConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_resource_configuration awscc_vpclattice_resource_configuration} Resource.
func NewVpclatticeResourceConfiguration(scope constructs.Construct, id *string, config *VpclatticeResourceConfigurationConfig) VpclatticeResourceConfiguration {
	_init_.Initialize()

	if err := validateNewVpclatticeResourceConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpclatticeResourceConfiguration{}

	_jsii_.Create(
		"awscc.vpclatticeResourceConfiguration.VpclatticeResourceConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_resource_configuration awscc_vpclattice_resource_configuration} Resource.
func NewVpclatticeResourceConfiguration_Override(v VpclatticeResourceConfiguration, scope constructs.Construct, id *string, config *VpclatticeResourceConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.vpclatticeResourceConfiguration.VpclatticeResourceConfiguration",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpclatticeResourceConfiguration)SetAllowAssociationToSharableServiceNetwork(val interface{}) {
	if err := j.validateSetAllowAssociationToSharableServiceNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAssociationToSharableServiceNetwork",
		val,
	)
}

func (j *jsiiProxy_VpclatticeResourceConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpclatticeResourceConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpclatticeResourceConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpclatticeResourceConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpclatticeResourceConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpclatticeResourceConfiguration)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VpclatticeResourceConfiguration)SetPortRanges(val *[]*string) {
	if err := j.validateSetPortRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRanges",
		val,
	)
}

func (j *jsiiProxy_VpclatticeResourceConfiguration)SetProtocolType(val *string) {
	if err := j.validateSetProtocolTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolType",
		val,
	)
}

func (j *jsiiProxy_VpclatticeResourceConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpclatticeResourceConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VpclatticeResourceConfiguration)SetResourceConfigurationAuthType(val *string) {
	if err := j.validateSetResourceConfigurationAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceConfigurationAuthType",
		val,
	)
}

func (j *jsiiProxy_VpclatticeResourceConfiguration)SetResourceConfigurationGroupId(val *string) {
	if err := j.validateSetResourceConfigurationGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceConfigurationGroupId",
		val,
	)
}

func (j *jsiiProxy_VpclatticeResourceConfiguration)SetResourceConfigurationType(val *string) {
	if err := j.validateSetResourceConfigurationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceConfigurationType",
		val,
	)
}

func (j *jsiiProxy_VpclatticeResourceConfiguration)SetResourceGatewayId(val *string) {
	if err := j.validateSetResourceGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGatewayId",
		val,
	)
}

// Generates CDKTF code for importing a VpclatticeResourceConfiguration resource upon running "cdktf plan <stack-name>".
func VpclatticeResourceConfiguration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVpclatticeResourceConfiguration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.vpclatticeResourceConfiguration.VpclatticeResourceConfiguration",
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
func VpclatticeResourceConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpclatticeResourceConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.vpclatticeResourceConfiguration.VpclatticeResourceConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpclatticeResourceConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpclatticeResourceConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.vpclatticeResourceConfiguration.VpclatticeResourceConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpclatticeResourceConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpclatticeResourceConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.vpclatticeResourceConfiguration.VpclatticeResourceConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpclatticeResourceConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.vpclatticeResourceConfiguration.VpclatticeResourceConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) PutResourceConfigurationDefinition(value *VpclatticeResourceConfigurationResourceConfigurationDefinition) {
	if err := v.validatePutResourceConfigurationDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putResourceConfigurationDefinition",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) PutTags(value interface{}) {
	if err := v.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTags",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) ResetAllowAssociationToSharableServiceNetwork() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowAssociationToSharableServiceNetwork",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) ResetPortRanges() {
	_jsii_.InvokeVoid(
		v,
		"resetPortRanges",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) ResetProtocolType() {
	_jsii_.InvokeVoid(
		v,
		"resetProtocolType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) ResetResourceConfigurationAuthType() {
	_jsii_.InvokeVoid(
		v,
		"resetResourceConfigurationAuthType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) ResetResourceConfigurationDefinition() {
	_jsii_.InvokeVoid(
		v,
		"resetResourceConfigurationDefinition",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) ResetResourceConfigurationGroupId() {
	_jsii_.InvokeVoid(
		v,
		"resetResourceConfigurationGroupId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) ResetResourceGatewayId() {
	_jsii_.InvokeVoid(
		v,
		"resetResourceGatewayId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeResourceConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

