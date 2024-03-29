package lightsaildistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lightsaildistribution/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_distribution awscc_lightsail_distribution}.
type LightsailDistribution interface {
	cdktf.TerraformResource
	AbleToUpdateBundle() cdktf.IResolvable
	BundleId() *string
	SetBundleId(val *string)
	BundleIdInput() *string
	CacheBehaviors() LightsailDistributionCacheBehaviorsList
	CacheBehaviorSettings() LightsailDistributionCacheBehaviorSettingsOutputReference
	CacheBehaviorSettingsInput() interface{}
	CacheBehaviorsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateName() *string
	SetCertificateName(val *string)
	CertificateNameInput() *string
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
	DefaultCacheBehavior() LightsailDistributionDefaultCacheBehaviorOutputReference
	DefaultCacheBehaviorInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DistributionArn() *string
	DistributionName() *string
	SetDistributionName(val *string)
	DistributionNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IpAddressType() *string
	SetIpAddressType(val *string)
	IpAddressTypeInput() *string
	IsEnabled() interface{}
	SetIsEnabled(val interface{})
	IsEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Origin() LightsailDistributionOriginOutputReference
	OriginInput() interface{}
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
	Status() *string
	Tags() LightsailDistributionTagsList
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
	PutCacheBehaviors(value interface{})
	PutCacheBehaviorSettings(value *LightsailDistributionCacheBehaviorSettings)
	PutDefaultCacheBehavior(value *LightsailDistributionDefaultCacheBehavior)
	PutOrigin(value *LightsailDistributionOrigin)
	PutTags(value interface{})
	ResetCacheBehaviors()
	ResetCacheBehaviorSettings()
	ResetCertificateName()
	ResetIpAddressType()
	ResetIsEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for LightsailDistribution
type jsiiProxy_LightsailDistribution struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LightsailDistribution) AbleToUpdateBundle() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ableToUpdateBundle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) BundleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) CacheBehaviors() LightsailDistributionCacheBehaviorsList {
	var returns LightsailDistributionCacheBehaviorsList
	_jsii_.Get(
		j,
		"cacheBehaviors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) CacheBehaviorSettings() LightsailDistributionCacheBehaviorSettingsOutputReference {
	var returns LightsailDistributionCacheBehaviorSettingsOutputReference
	_jsii_.Get(
		j,
		"cacheBehaviorSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) CacheBehaviorSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheBehaviorSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) CacheBehaviorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheBehaviorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) CertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) CertificateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) DefaultCacheBehavior() LightsailDistributionDefaultCacheBehaviorOutputReference {
	var returns LightsailDistributionDefaultCacheBehaviorOutputReference
	_jsii_.Get(
		j,
		"defaultCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) DefaultCacheBehaviorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultCacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) DistributionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) DistributionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) DistributionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) IsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) IsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) Origin() LightsailDistributionOriginOutputReference {
	var returns LightsailDistributionOriginOutputReference
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) Tags() LightsailDistributionTagsList {
	var returns LightsailDistributionTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistribution) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_distribution awscc_lightsail_distribution} Resource.
func NewLightsailDistribution(scope constructs.Construct, id *string, config *LightsailDistributionConfig) LightsailDistribution {
	_init_.Initialize()

	if err := validateNewLightsailDistributionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LightsailDistribution{}

	_jsii_.Create(
		"awscc.lightsailDistribution.LightsailDistribution",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_distribution awscc_lightsail_distribution} Resource.
func NewLightsailDistribution_Override(l LightsailDistribution, scope constructs.Construct, id *string, config *LightsailDistributionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lightsailDistribution.LightsailDistribution",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LightsailDistribution)SetBundleId(val *string) {
	if err := j.validateSetBundleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_LightsailDistribution)SetCertificateName(val *string) {
	if err := j.validateSetCertificateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateName",
		val,
	)
}

func (j *jsiiProxy_LightsailDistribution)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LightsailDistribution)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LightsailDistribution)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LightsailDistribution)SetDistributionName(val *string) {
	if err := j.validateSetDistributionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distributionName",
		val,
	)
}

func (j *jsiiProxy_LightsailDistribution)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LightsailDistribution)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_LightsailDistribution)SetIsEnabled(val interface{}) {
	if err := j.validateSetIsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isEnabled",
		val,
	)
}

func (j *jsiiProxy_LightsailDistribution)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LightsailDistribution)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LightsailDistribution)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a LightsailDistribution resource upon running "cdktf plan <stack-name>".
func LightsailDistribution_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLightsailDistribution_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.lightsailDistribution.LightsailDistribution",
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
func LightsailDistribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLightsailDistribution_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lightsailDistribution.LightsailDistribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LightsailDistribution_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLightsailDistribution_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lightsailDistribution.LightsailDistribution",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LightsailDistribution_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLightsailDistribution_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lightsailDistribution.LightsailDistribution",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LightsailDistribution_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.lightsailDistribution.LightsailDistribution",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LightsailDistribution) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LightsailDistribution) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LightsailDistribution) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LightsailDistribution) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LightsailDistribution) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LightsailDistribution) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LightsailDistribution) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LightsailDistribution) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LightsailDistribution) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LightsailDistribution) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LightsailDistribution) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LightsailDistribution) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LightsailDistribution) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LightsailDistribution) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LightsailDistribution) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LightsailDistribution) PutCacheBehaviors(value interface{}) {
	if err := l.validatePutCacheBehaviorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCacheBehaviors",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LightsailDistribution) PutCacheBehaviorSettings(value *LightsailDistributionCacheBehaviorSettings) {
	if err := l.validatePutCacheBehaviorSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCacheBehaviorSettings",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LightsailDistribution) PutDefaultCacheBehavior(value *LightsailDistributionDefaultCacheBehavior) {
	if err := l.validatePutDefaultCacheBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDefaultCacheBehavior",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LightsailDistribution) PutOrigin(value *LightsailDistributionOrigin) {
	if err := l.validatePutOriginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putOrigin",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LightsailDistribution) PutTags(value interface{}) {
	if err := l.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTags",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LightsailDistribution) ResetCacheBehaviors() {
	_jsii_.InvokeVoid(
		l,
		"resetCacheBehaviors",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDistribution) ResetCacheBehaviorSettings() {
	_jsii_.InvokeVoid(
		l,
		"resetCacheBehaviorSettings",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDistribution) ResetCertificateName() {
	_jsii_.InvokeVoid(
		l,
		"resetCertificateName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDistribution) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		l,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDistribution) ResetIsEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetIsEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDistribution) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDistribution) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDistribution) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDistribution) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDistribution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDistribution) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

