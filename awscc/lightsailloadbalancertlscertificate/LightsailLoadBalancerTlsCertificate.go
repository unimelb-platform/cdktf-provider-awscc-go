package lightsailloadbalancertlscertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lightsailloadbalancertlscertificate/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_load_balancer_tls_certificate awscc_lightsail_load_balancer_tls_certificate}.
type LightsailLoadBalancerTlsCertificate interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateAlternativeNames() *[]*string
	SetCertificateAlternativeNames(val *[]*string)
	CertificateAlternativeNamesInput() *[]*string
	CertificateDomainName() *string
	SetCertificateDomainName(val *string)
	CertificateDomainNameInput() *string
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
	HttpsRedirectionEnabled() interface{}
	SetHttpsRedirectionEnabled(val interface{})
	HttpsRedirectionEnabledInput() interface{}
	Id() *string
	IsAttached() interface{}
	SetIsAttached(val interface{})
	IsAttachedInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancerName() *string
	SetLoadBalancerName(val *string)
	LoadBalancerNameInput() *string
	LoadBalancerTlsCertificateArn() *string
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
	Status() *string
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
	ResetCertificateAlternativeNames()
	ResetHttpsRedirectionEnabled()
	ResetIsAttached()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LightsailLoadBalancerTlsCertificate
type jsiiProxy_LightsailLoadBalancerTlsCertificate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) CertificateAlternativeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificateAlternativeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) CertificateAlternativeNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificateAlternativeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) CertificateDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) CertificateDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) CertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) CertificateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) HttpsRedirectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsRedirectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) HttpsRedirectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsRedirectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) IsAttached() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAttached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) IsAttachedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAttachedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) LoadBalancerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) LoadBalancerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) LoadBalancerTlsCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerTlsCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_load_balancer_tls_certificate awscc_lightsail_load_balancer_tls_certificate} Resource.
func NewLightsailLoadBalancerTlsCertificate(scope constructs.Construct, id *string, config *LightsailLoadBalancerTlsCertificateConfig) LightsailLoadBalancerTlsCertificate {
	_init_.Initialize()

	if err := validateNewLightsailLoadBalancerTlsCertificateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LightsailLoadBalancerTlsCertificate{}

	_jsii_.Create(
		"awscc.lightsailLoadBalancerTlsCertificate.LightsailLoadBalancerTlsCertificate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_load_balancer_tls_certificate awscc_lightsail_load_balancer_tls_certificate} Resource.
func NewLightsailLoadBalancerTlsCertificate_Override(l LightsailLoadBalancerTlsCertificate, scope constructs.Construct, id *string, config *LightsailLoadBalancerTlsCertificateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lightsailLoadBalancerTlsCertificate.LightsailLoadBalancerTlsCertificate",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate)SetCertificateAlternativeNames(val *[]*string) {
	if err := j.validateSetCertificateAlternativeNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateAlternativeNames",
		val,
	)
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate)SetCertificateDomainName(val *string) {
	if err := j.validateSetCertificateDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateDomainName",
		val,
	)
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate)SetCertificateName(val *string) {
	if err := j.validateSetCertificateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateName",
		val,
	)
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate)SetHttpsRedirectionEnabled(val interface{}) {
	if err := j.validateSetHttpsRedirectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsRedirectionEnabled",
		val,
	)
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate)SetIsAttached(val interface{}) {
	if err := j.validateSetIsAttachedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isAttached",
		val,
	)
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate)SetLoadBalancerName(val *string) {
	if err := j.validateSetLoadBalancerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerName",
		val,
	)
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LightsailLoadBalancerTlsCertificate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a LightsailLoadBalancerTlsCertificate resource upon running "cdktf plan <stack-name>".
func LightsailLoadBalancerTlsCertificate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLightsailLoadBalancerTlsCertificate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.lightsailLoadBalancerTlsCertificate.LightsailLoadBalancerTlsCertificate",
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
func LightsailLoadBalancerTlsCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLightsailLoadBalancerTlsCertificate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lightsailLoadBalancerTlsCertificate.LightsailLoadBalancerTlsCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LightsailLoadBalancerTlsCertificate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLightsailLoadBalancerTlsCertificate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lightsailLoadBalancerTlsCertificate.LightsailLoadBalancerTlsCertificate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LightsailLoadBalancerTlsCertificate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLightsailLoadBalancerTlsCertificate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lightsailLoadBalancerTlsCertificate.LightsailLoadBalancerTlsCertificate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LightsailLoadBalancerTlsCertificate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.lightsailLoadBalancerTlsCertificate.LightsailLoadBalancerTlsCertificate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) ResetCertificateAlternativeNames() {
	_jsii_.InvokeVoid(
		l,
		"resetCertificateAlternativeNames",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) ResetHttpsRedirectionEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpsRedirectionEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) ResetIsAttached() {
	_jsii_.InvokeVoid(
		l,
		"resetIsAttached",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailLoadBalancerTlsCertificate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

