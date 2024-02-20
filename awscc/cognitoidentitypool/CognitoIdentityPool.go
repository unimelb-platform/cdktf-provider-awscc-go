package cognitoidentitypool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cognitoidentitypool/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cognito_identity_pool awscc_cognito_identity_pool}.
type CognitoIdentityPool interface {
	cdktf.TerraformResource
	AllowClassicFlow() interface{}
	SetAllowClassicFlow(val interface{})
	AllowClassicFlowInput() interface{}
	AllowUnauthenticatedIdentities() interface{}
	SetAllowUnauthenticatedIdentities(val interface{})
	AllowUnauthenticatedIdentitiesInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CognitoEvents() *string
	SetCognitoEvents(val *string)
	CognitoEventsInput() *string
	CognitoIdentityProviders() CognitoIdentityPoolCognitoIdentityProvidersList
	CognitoIdentityProvidersInput() interface{}
	CognitoStreams() CognitoIdentityPoolCognitoStreamsOutputReference
	CognitoStreamsInput() interface{}
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
	DeveloperProviderName() *string
	SetDeveloperProviderName(val *string)
	DeveloperProviderNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IdentityPoolName() *string
	SetIdentityPoolName(val *string)
	IdentityPoolNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	OpenIdConnectProviderArNs() *[]*string
	SetOpenIdConnectProviderArNs(val *[]*string)
	OpenIdConnectProviderArNsInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PushSync() CognitoIdentityPoolPushSyncOutputReference
	PushSyncInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	SamlProviderArNs() *[]*string
	SetSamlProviderArNs(val *[]*string)
	SamlProviderArNsInput() *[]*string
	SupportedLoginProviders() *string
	SetSupportedLoginProviders(val *string)
	SupportedLoginProvidersInput() *string
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
	PutCognitoIdentityProviders(value interface{})
	PutCognitoStreams(value *CognitoIdentityPoolCognitoStreams)
	PutPushSync(value *CognitoIdentityPoolPushSync)
	ResetAllowClassicFlow()
	ResetCognitoEvents()
	ResetCognitoIdentityProviders()
	ResetCognitoStreams()
	ResetDeveloperProviderName()
	ResetIdentityPoolName()
	ResetOpenIdConnectProviderArNs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPushSync()
	ResetSamlProviderArNs()
	ResetSupportedLoginProviders()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CognitoIdentityPool
type jsiiProxy_CognitoIdentityPool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CognitoIdentityPool) AllowClassicFlow() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClassicFlow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) AllowClassicFlowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClassicFlowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) AllowUnauthenticatedIdentities() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnauthenticatedIdentities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) AllowUnauthenticatedIdentitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnauthenticatedIdentitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CognitoEvents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cognitoEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CognitoEventsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cognitoEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CognitoIdentityProviders() CognitoIdentityPoolCognitoIdentityProvidersList {
	var returns CognitoIdentityPoolCognitoIdentityProvidersList
	_jsii_.Get(
		j,
		"cognitoIdentityProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CognitoIdentityProvidersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cognitoIdentityProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CognitoStreams() CognitoIdentityPoolCognitoStreamsOutputReference {
	var returns CognitoIdentityPoolCognitoStreamsOutputReference
	_jsii_.Get(
		j,
		"cognitoStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CognitoStreamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cognitoStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) DeveloperProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) DeveloperProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerProviderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) IdentityPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) IdentityPoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) OpenIdConnectProviderArNs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"openIdConnectProviderArNs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) OpenIdConnectProviderArNsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"openIdConnectProviderArNsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) PushSync() CognitoIdentityPoolPushSyncOutputReference {
	var returns CognitoIdentityPoolPushSyncOutputReference
	_jsii_.Get(
		j,
		"pushSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) PushSyncInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushSyncInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) SamlProviderArNs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"samlProviderArNs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) SamlProviderArNsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"samlProviderArNsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) SupportedLoginProviders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportedLoginProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) SupportedLoginProvidersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportedLoginProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cognito_identity_pool awscc_cognito_identity_pool} Resource.
func NewCognitoIdentityPool(scope constructs.Construct, id *string, config *CognitoIdentityPoolConfig) CognitoIdentityPool {
	_init_.Initialize()

	if err := validateNewCognitoIdentityPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoIdentityPool{}

	_jsii_.Create(
		"awscc.cognitoIdentityPool.CognitoIdentityPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cognito_identity_pool awscc_cognito_identity_pool} Resource.
func NewCognitoIdentityPool_Override(c CognitoIdentityPool, scope constructs.Construct, id *string, config *CognitoIdentityPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cognitoIdentityPool.CognitoIdentityPool",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetAllowClassicFlow(val interface{}) {
	if err := j.validateSetAllowClassicFlowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowClassicFlow",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetAllowUnauthenticatedIdentities(val interface{}) {
	if err := j.validateSetAllowUnauthenticatedIdentitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowUnauthenticatedIdentities",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetCognitoEvents(val *string) {
	if err := j.validateSetCognitoEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cognitoEvents",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetDeveloperProviderName(val *string) {
	if err := j.validateSetDeveloperProviderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developerProviderName",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetIdentityPoolName(val *string) {
	if err := j.validateSetIdentityPoolNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityPoolName",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetOpenIdConnectProviderArNs(val *[]*string) {
	if err := j.validateSetOpenIdConnectProviderArNsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openIdConnectProviderArNs",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetSamlProviderArNs(val *[]*string) {
	if err := j.validateSetSamlProviderArNsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samlProviderArNs",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetSupportedLoginProviders(val *string) {
	if err := j.validateSetSupportedLoginProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedLoginProviders",
		val,
	)
}

// Generates CDKTF code for importing a CognitoIdentityPool resource upon running "cdktf plan <stack-name>".
func CognitoIdentityPool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCognitoIdentityPool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.cognitoIdentityPool.CognitoIdentityPool",
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
func CognitoIdentityPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoIdentityPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cognitoIdentityPool.CognitoIdentityPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CognitoIdentityPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoIdentityPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cognitoIdentityPool.CognitoIdentityPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CognitoIdentityPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoIdentityPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cognitoIdentityPool.CognitoIdentityPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CognitoIdentityPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.cognitoIdentityPool.CognitoIdentityPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CognitoIdentityPool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoIdentityPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CognitoIdentityPool) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CognitoIdentityPool) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CognitoIdentityPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CognitoIdentityPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CognitoIdentityPool) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CognitoIdentityPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CognitoIdentityPool) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoIdentityPool) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) PutCognitoIdentityProviders(value interface{}) {
	if err := c.validatePutCognitoIdentityProvidersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCognitoIdentityProviders",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) PutCognitoStreams(value *CognitoIdentityPoolCognitoStreams) {
	if err := c.validatePutCognitoStreamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCognitoStreams",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) PutPushSync(value *CognitoIdentityPoolPushSync) {
	if err := c.validatePutPushSyncParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPushSync",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetAllowClassicFlow() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowClassicFlow",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetCognitoEvents() {
	_jsii_.InvokeVoid(
		c,
		"resetCognitoEvents",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetCognitoIdentityProviders() {
	_jsii_.InvokeVoid(
		c,
		"resetCognitoIdentityProviders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetCognitoStreams() {
	_jsii_.InvokeVoid(
		c,
		"resetCognitoStreams",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetDeveloperProviderName() {
	_jsii_.InvokeVoid(
		c,
		"resetDeveloperProviderName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetIdentityPoolName() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentityPoolName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetOpenIdConnectProviderArNs() {
	_jsii_.InvokeVoid(
		c,
		"resetOpenIdConnectProviderArNs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetPushSync() {
	_jsii_.InvokeVoid(
		c,
		"resetPushSync",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetSamlProviderArNs() {
	_jsii_.InvokeVoid(
		c,
		"resetSamlProviderArNs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetSupportedLoginProviders() {
	_jsii_.InvokeVoid(
		c,
		"resetSupportedLoginProviders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

