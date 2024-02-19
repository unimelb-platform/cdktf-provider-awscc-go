package lightsailcontainer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lightsailcontainer/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_container awscc_lightsail_container}.
type LightsailContainer interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerArn() *string
	ContainerServiceDeployment() LightsailContainerContainerServiceDeploymentOutputReference
	ContainerServiceDeploymentInput() interface{}
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
	IsDisabled() interface{}
	SetIsDisabled(val interface{})
	IsDisabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Power() *string
	SetPower(val *string)
	PowerInput() *string
	PrincipalArn() *string
	PrivateRegistryAccess() LightsailContainerPrivateRegistryAccessOutputReference
	PrivateRegistryAccessInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicDomainNames() LightsailContainerPublicDomainNamesList
	PublicDomainNamesInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	Scale() *float64
	SetScale(val *float64)
	ScaleInput() *float64
	ServiceName() *string
	SetServiceName(val *string)
	ServiceNameInput() *string
	Tags() LightsailContainerTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Url() *string
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
	PutContainerServiceDeployment(value *LightsailContainerContainerServiceDeployment)
	PutPrivateRegistryAccess(value *LightsailContainerPrivateRegistryAccess)
	PutPublicDomainNames(value interface{})
	PutTags(value interface{})
	ResetContainerServiceDeployment()
	ResetIsDisabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateRegistryAccess()
	ResetPublicDomainNames()
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

// The jsii proxy struct for LightsailContainer
type jsiiProxy_LightsailContainer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LightsailContainer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) ContainerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) ContainerServiceDeployment() LightsailContainerContainerServiceDeploymentOutputReference {
	var returns LightsailContainerContainerServiceDeploymentOutputReference
	_jsii_.Get(
		j,
		"containerServiceDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) ContainerServiceDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerServiceDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) IsDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) IsDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) Power() *string {
	var returns *string
	_jsii_.Get(
		j,
		"power",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) PowerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) PrincipalArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) PrivateRegistryAccess() LightsailContainerPrivateRegistryAccessOutputReference {
	var returns LightsailContainerPrivateRegistryAccessOutputReference
	_jsii_.Get(
		j,
		"privateRegistryAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) PrivateRegistryAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateRegistryAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) PublicDomainNames() LightsailContainerPublicDomainNamesList {
	var returns LightsailContainerPublicDomainNamesList
	_jsii_.Get(
		j,
		"publicDomainNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) PublicDomainNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicDomainNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) Scale() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) ScaleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) ServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) Tags() LightsailContainerTagsList {
	var returns LightsailContainerTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailContainer) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_container awscc_lightsail_container} Resource.
func NewLightsailContainer(scope constructs.Construct, id *string, config *LightsailContainerConfig) LightsailContainer {
	_init_.Initialize()

	if err := validateNewLightsailContainerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LightsailContainer{}

	_jsii_.Create(
		"awscc.lightsailContainer.LightsailContainer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_container awscc_lightsail_container} Resource.
func NewLightsailContainer_Override(l LightsailContainer, scope constructs.Construct, id *string, config *LightsailContainerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lightsailContainer.LightsailContainer",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LightsailContainer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LightsailContainer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LightsailContainer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LightsailContainer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LightsailContainer)SetIsDisabled(val interface{}) {
	if err := j.validateSetIsDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isDisabled",
		val,
	)
}

func (j *jsiiProxy_LightsailContainer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LightsailContainer)SetPower(val *string) {
	if err := j.validateSetPowerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"power",
		val,
	)
}

func (j *jsiiProxy_LightsailContainer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LightsailContainer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LightsailContainer)SetScale(val *float64) {
	if err := j.validateSetScaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scale",
		val,
	)
}

func (j *jsiiProxy_LightsailContainer)SetServiceName(val *string) {
	if err := j.validateSetServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

// Generates CDKTF code for importing a LightsailContainer resource upon running "cdktf plan <stack-name>".
func LightsailContainer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLightsailContainer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.lightsailContainer.LightsailContainer",
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
func LightsailContainer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLightsailContainer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lightsailContainer.LightsailContainer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LightsailContainer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLightsailContainer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lightsailContainer.LightsailContainer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LightsailContainer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLightsailContainer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lightsailContainer.LightsailContainer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LightsailContainer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.lightsailContainer.LightsailContainer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LightsailContainer) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LightsailContainer) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LightsailContainer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailContainer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailContainer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailContainer) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailContainer) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailContainer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailContainer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailContainer) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailContainer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailContainer) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LightsailContainer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailContainer) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LightsailContainer) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LightsailContainer) PutContainerServiceDeployment(value *LightsailContainerContainerServiceDeployment) {
	if err := l.validatePutContainerServiceDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putContainerServiceDeployment",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LightsailContainer) PutPrivateRegistryAccess(value *LightsailContainerPrivateRegistryAccess) {
	if err := l.validatePutPrivateRegistryAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPrivateRegistryAccess",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LightsailContainer) PutPublicDomainNames(value interface{}) {
	if err := l.validatePutPublicDomainNamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPublicDomainNames",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LightsailContainer) PutTags(value interface{}) {
	if err := l.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTags",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LightsailContainer) ResetContainerServiceDeployment() {
	_jsii_.InvokeVoid(
		l,
		"resetContainerServiceDeployment",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailContainer) ResetIsDisabled() {
	_jsii_.InvokeVoid(
		l,
		"resetIsDisabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailContainer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailContainer) ResetPrivateRegistryAccess() {
	_jsii_.InvokeVoid(
		l,
		"resetPrivateRegistryAccess",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailContainer) ResetPublicDomainNames() {
	_jsii_.InvokeVoid(
		l,
		"resetPublicDomainNames",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailContainer) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailContainer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailContainer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailContainer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailContainer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

