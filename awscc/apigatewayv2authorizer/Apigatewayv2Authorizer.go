package apigatewayv2authorizer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/apigatewayv2authorizer/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_authorizer awscc_apigatewayv2_authorizer}.
type Apigatewayv2Authorizer interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	AuthorizerCredentialsArn() *string
	SetAuthorizerCredentialsArn(val *string)
	AuthorizerCredentialsArnInput() *string
	AuthorizerId() *string
	AuthorizerPayloadFormatVersion() *string
	SetAuthorizerPayloadFormatVersion(val *string)
	AuthorizerPayloadFormatVersionInput() *string
	AuthorizerResultTtlInSeconds() *float64
	SetAuthorizerResultTtlInSeconds(val *float64)
	AuthorizerResultTtlInSecondsInput() *float64
	AuthorizerType() *string
	SetAuthorizerType(val *string)
	AuthorizerTypeInput() *string
	AuthorizerUri() *string
	SetAuthorizerUri(val *string)
	AuthorizerUriInput() *string
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
	EnableSimpleResponses() interface{}
	SetEnableSimpleResponses(val interface{})
	EnableSimpleResponsesInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IdentitySource() *[]*string
	SetIdentitySource(val *[]*string)
	IdentitySourceInput() *[]*string
	IdentityValidationExpression() *string
	SetIdentityValidationExpression(val *string)
	IdentityValidationExpressionInput() *string
	JwtConfiguration() Apigatewayv2AuthorizerJwtConfigurationOutputReference
	JwtConfigurationInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	PutJwtConfiguration(value *Apigatewayv2AuthorizerJwtConfiguration)
	ResetAuthorizerCredentialsArn()
	ResetAuthorizerPayloadFormatVersion()
	ResetAuthorizerResultTtlInSeconds()
	ResetAuthorizerUri()
	ResetEnableSimpleResponses()
	ResetIdentitySource()
	ResetIdentityValidationExpression()
	ResetJwtConfiguration()
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

// The jsii proxy struct for Apigatewayv2Authorizer
type jsiiProxy_Apigatewayv2Authorizer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Apigatewayv2Authorizer) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerCredentialsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerCredentialsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerCredentialsArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerCredentialsArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerPayloadFormatVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerPayloadFormatVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerPayloadFormatVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerPayloadFormatVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerResultTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorizerResultTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerResultTtlInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorizerResultTtlInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) EnableSimpleResponses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSimpleResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) EnableSimpleResponsesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSimpleResponsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) IdentitySource() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identitySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) IdentitySourceInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identitySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) IdentityValidationExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityValidationExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) IdentityValidationExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityValidationExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) JwtConfiguration() Apigatewayv2AuthorizerJwtConfigurationOutputReference {
	var returns Apigatewayv2AuthorizerJwtConfigurationOutputReference
	_jsii_.Get(
		j,
		"jwtConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) JwtConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jwtConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_authorizer awscc_apigatewayv2_authorizer} Resource.
func NewApigatewayv2Authorizer(scope constructs.Construct, id *string, config *Apigatewayv2AuthorizerConfig) Apigatewayv2Authorizer {
	_init_.Initialize()

	if err := validateNewApigatewayv2AuthorizerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Apigatewayv2Authorizer{}

	_jsii_.Create(
		"awscc.apigatewayv2Authorizer.Apigatewayv2Authorizer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_authorizer awscc_apigatewayv2_authorizer} Resource.
func NewApigatewayv2Authorizer_Override(a Apigatewayv2Authorizer, scope constructs.Construct, id *string, config *Apigatewayv2AuthorizerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.apigatewayv2Authorizer.Apigatewayv2Authorizer",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetApiId(val *string) {
	if err := j.validateSetApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetAuthorizerCredentialsArn(val *string) {
	if err := j.validateSetAuthorizerCredentialsArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizerCredentialsArn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetAuthorizerPayloadFormatVersion(val *string) {
	if err := j.validateSetAuthorizerPayloadFormatVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizerPayloadFormatVersion",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetAuthorizerResultTtlInSeconds(val *float64) {
	if err := j.validateSetAuthorizerResultTtlInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizerResultTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetAuthorizerType(val *string) {
	if err := j.validateSetAuthorizerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizerType",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetAuthorizerUri(val *string) {
	if err := j.validateSetAuthorizerUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizerUri",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetEnableSimpleResponses(val interface{}) {
	if err := j.validateSetEnableSimpleResponsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSimpleResponses",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetIdentitySource(val *[]*string) {
	if err := j.validateSetIdentitySourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identitySource",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetIdentityValidationExpression(val *string) {
	if err := j.validateSetIdentityValidationExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityValidationExpression",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a Apigatewayv2Authorizer resource upon running "cdktf plan <stack-name>".
func Apigatewayv2Authorizer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApigatewayv2Authorizer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.apigatewayv2Authorizer.Apigatewayv2Authorizer",
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
func Apigatewayv2Authorizer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigatewayv2Authorizer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.apigatewayv2Authorizer.Apigatewayv2Authorizer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Apigatewayv2Authorizer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigatewayv2Authorizer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.apigatewayv2Authorizer.Apigatewayv2Authorizer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Apigatewayv2Authorizer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigatewayv2Authorizer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.apigatewayv2Authorizer.Apigatewayv2Authorizer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Apigatewayv2Authorizer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.apigatewayv2Authorizer.Apigatewayv2Authorizer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_Apigatewayv2Authorizer) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Authorizer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Authorizer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Authorizer) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Authorizer) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Authorizer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Authorizer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Authorizer) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Authorizer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Authorizer) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) PutJwtConfiguration(value *Apigatewayv2AuthorizerJwtConfiguration) {
	if err := a.validatePutJwtConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJwtConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ResetAuthorizerCredentialsArn() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerCredentialsArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ResetAuthorizerPayloadFormatVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerPayloadFormatVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ResetAuthorizerResultTtlInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerResultTtlInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ResetAuthorizerUri() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ResetEnableSimpleResponses() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableSimpleResponses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ResetIdentitySource() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentitySource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ResetIdentityValidationExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityValidationExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ResetJwtConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetJwtConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

