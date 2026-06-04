package elasticloadbalancingv2listener

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/elasticloadbalancingv2listener/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticloadbalancingv2_listener awscc_elasticloadbalancingv2_listener}.
type Elasticloadbalancingv2Listener interface {
	cdktf.TerraformResource
	AlpnPolicy() *[]*string
	SetAlpnPolicy(val *[]*string)
	AlpnPolicyInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Certificates() Elasticloadbalancingv2ListenerCertificatesList
	CertificatesInput() interface{}
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
	DefaultActions() Elasticloadbalancingv2ListenerDefaultActionsList
	DefaultActionsInput() interface{}
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
	ListenerArn() *string
	ListenerAttributes() Elasticloadbalancingv2ListenerListenerAttributesList
	ListenerAttributesInput() interface{}
	LoadBalancerArn() *string
	SetLoadBalancerArn(val *string)
	LoadBalancerArnInput() *string
	MutualAuthentication() Elasticloadbalancingv2ListenerMutualAuthenticationOutputReference
	MutualAuthenticationInput() interface{}
	// The tree node.
	Node() constructs.Node
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	SslPolicy() *string
	SetSslPolicy(val *string)
	SslPolicyInput() *string
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
	PutCertificates(value interface{})
	PutDefaultActions(value interface{})
	PutListenerAttributes(value interface{})
	PutMutualAuthentication(value *Elasticloadbalancingv2ListenerMutualAuthentication)
	ResetAlpnPolicy()
	ResetCertificates()
	ResetListenerAttributes()
	ResetMutualAuthentication()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetProtocol()
	ResetSslPolicy()
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

// The jsii proxy struct for Elasticloadbalancingv2Listener
type jsiiProxy_Elasticloadbalancingv2Listener struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) AlpnPolicy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alpnPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) AlpnPolicyInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alpnPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) Certificates() Elasticloadbalancingv2ListenerCertificatesList {
	var returns Elasticloadbalancingv2ListenerCertificatesList
	_jsii_.Get(
		j,
		"certificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) CertificatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) DefaultActions() Elasticloadbalancingv2ListenerDefaultActionsList {
	var returns Elasticloadbalancingv2ListenerDefaultActionsList
	_jsii_.Get(
		j,
		"defaultActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) DefaultActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) ListenerAttributes() Elasticloadbalancingv2ListenerListenerAttributesList {
	var returns Elasticloadbalancingv2ListenerListenerAttributesList
	_jsii_.Get(
		j,
		"listenerAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) ListenerAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"listenerAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) LoadBalancerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) MutualAuthentication() Elasticloadbalancingv2ListenerMutualAuthenticationOutputReference {
	var returns Elasticloadbalancingv2ListenerMutualAuthenticationOutputReference
	_jsii_.Get(
		j,
		"mutualAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) MutualAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutualAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) SslPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) SslPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticloadbalancingv2_listener awscc_elasticloadbalancingv2_listener} Resource.
func NewElasticloadbalancingv2Listener(scope constructs.Construct, id *string, config *Elasticloadbalancingv2ListenerConfig) Elasticloadbalancingv2Listener {
	_init_.Initialize()

	if err := validateNewElasticloadbalancingv2ListenerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Elasticloadbalancingv2Listener{}

	_jsii_.Create(
		"awscc.elasticloadbalancingv2Listener.Elasticloadbalancingv2Listener",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticloadbalancingv2_listener awscc_elasticloadbalancingv2_listener} Resource.
func NewElasticloadbalancingv2Listener_Override(e Elasticloadbalancingv2Listener, scope constructs.Construct, id *string, config *Elasticloadbalancingv2ListenerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.elasticloadbalancingv2Listener.Elasticloadbalancingv2Listener",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener)SetAlpnPolicy(val *[]*string) {
	if err := j.validateSetAlpnPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alpnPolicy",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener)SetLoadBalancerArn(val *string) {
	if err := j.validateSetLoadBalancerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerArn",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2Listener)SetSslPolicy(val *string) {
	if err := j.validateSetSslPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslPolicy",
		val,
	)
}

// Generates CDKTF code for importing a Elasticloadbalancingv2Listener resource upon running "cdktf plan <stack-name>".
func Elasticloadbalancingv2Listener_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2Listener_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.elasticloadbalancingv2Listener.Elasticloadbalancingv2Listener",
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
func Elasticloadbalancingv2Listener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2Listener_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.elasticloadbalancingv2Listener.Elasticloadbalancingv2Listener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Elasticloadbalancingv2Listener_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2Listener_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.elasticloadbalancingv2Listener.Elasticloadbalancingv2Listener",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Elasticloadbalancingv2Listener_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2Listener_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.elasticloadbalancingv2Listener.Elasticloadbalancingv2Listener",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Elasticloadbalancingv2Listener_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.elasticloadbalancingv2Listener.Elasticloadbalancingv2Listener",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Elasticloadbalancingv2Listener) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Elasticloadbalancingv2Listener) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Elasticloadbalancingv2Listener) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Elasticloadbalancingv2Listener) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2Listener) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2Listener) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2Listener) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Elasticloadbalancingv2Listener) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Elasticloadbalancingv2Listener) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Elasticloadbalancingv2Listener) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) PutCertificates(value interface{}) {
	if err := e.validatePutCertificatesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCertificates",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) PutDefaultActions(value interface{}) {
	if err := e.validatePutDefaultActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDefaultActions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) PutListenerAttributes(value interface{}) {
	if err := e.validatePutListenerAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putListenerAttributes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) PutMutualAuthentication(value *Elasticloadbalancingv2ListenerMutualAuthentication) {
	if err := e.validatePutMutualAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMutualAuthentication",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) ResetAlpnPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetAlpnPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) ResetCertificates() {
	_jsii_.InvokeVoid(
		e,
		"resetCertificates",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) ResetListenerAttributes() {
	_jsii_.InvokeVoid(
		e,
		"resetListenerAttributes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) ResetMutualAuthentication() {
	_jsii_.InvokeVoid(
		e,
		"resetMutualAuthentication",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) ResetPort() {
	_jsii_.InvokeVoid(
		e,
		"resetPort",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) ResetProtocol() {
	_jsii_.InvokeVoid(
		e,
		"resetProtocol",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) ResetSslPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetSslPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2Listener) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

