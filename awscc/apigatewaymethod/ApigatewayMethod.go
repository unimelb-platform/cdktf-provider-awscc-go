package apigatewaymethod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/apigatewaymethod/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method awscc_apigateway_method}.
type ApigatewayMethod interface {
	cdktf.TerraformResource
	ApiKeyRequired() interface{}
	SetApiKeyRequired(val interface{})
	ApiKeyRequiredInput() interface{}
	AuthorizationScopes() *[]*string
	SetAuthorizationScopes(val *[]*string)
	AuthorizationScopesInput() *[]*string
	AuthorizationType() *string
	SetAuthorizationType(val *string)
	AuthorizationTypeInput() *string
	AuthorizerId() *string
	SetAuthorizerId(val *string)
	AuthorizerIdInput() *string
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
	HttpMethod() *string
	SetHttpMethod(val *string)
	HttpMethodInput() *string
	Id() *string
	Integration() ApigatewayMethodIntegrationOutputReference
	IntegrationInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MethodResponses() ApigatewayMethodMethodResponsesList
	MethodResponsesInput() interface{}
	// The tree node.
	Node() constructs.Node
	OperationName() *string
	SetOperationName(val *string)
	OperationNameInput() *string
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
	RequestModels() *map[string]*string
	SetRequestModels(val *map[string]*string)
	RequestModelsInput() *map[string]*string
	RequestParameters() *map[string]interface{}
	SetRequestParameters(val *map[string]interface{})
	RequestParametersInput() *map[string]interface{}
	RequestValidatorId() *string
	SetRequestValidatorId(val *string)
	RequestValidatorIdInput() *string
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
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
	PutIntegration(value *ApigatewayMethodIntegration)
	PutMethodResponses(value interface{})
	ResetApiKeyRequired()
	ResetAuthorizationScopes()
	ResetAuthorizationType()
	ResetAuthorizerId()
	ResetIntegration()
	ResetMethodResponses()
	ResetOperationName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequestModels()
	ResetRequestParameters()
	ResetRequestValidatorId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ApigatewayMethod
type jsiiProxy_ApigatewayMethod struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApigatewayMethod) ApiKeyRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) ApiKeyRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) AuthorizationScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizationScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) AuthorizationScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizationScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) AuthorizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) AuthorizationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) AuthorizerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) Integration() ApigatewayMethodIntegrationOutputReference {
	var returns ApigatewayMethodIntegrationOutputReference
	_jsii_.Get(
		j,
		"integration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) IntegrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"integrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) MethodResponses() ApigatewayMethodMethodResponsesList {
	var returns ApigatewayMethodMethodResponsesList
	_jsii_.Get(
		j,
		"methodResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) MethodResponsesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"methodResponsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) OperationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) RequestModels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) RequestModelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestModelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) RequestParameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"requestParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) RequestParametersInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"requestParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) RequestValidatorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestValidatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) RequestValidatorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestValidatorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethod) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method awscc_apigateway_method} Resource.
func NewApigatewayMethod(scope constructs.Construct, id *string, config *ApigatewayMethodConfig) ApigatewayMethod {
	_init_.Initialize()

	if err := validateNewApigatewayMethodParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigatewayMethod{}

	_jsii_.Create(
		"awscc.apigatewayMethod.ApigatewayMethod",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method awscc_apigateway_method} Resource.
func NewApigatewayMethod_Override(a ApigatewayMethod, scope constructs.Construct, id *string, config *ApigatewayMethodConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.apigatewayMethod.ApigatewayMethod",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetApiKeyRequired(val interface{}) {
	if err := j.validateSetApiKeyRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiKeyRequired",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetAuthorizationScopes(val *[]*string) {
	if err := j.validateSetAuthorizationScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationScopes",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetAuthorizationType(val *string) {
	if err := j.validateSetAuthorizationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationType",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetAuthorizerId(val *string) {
	if err := j.validateSetAuthorizerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizerId",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetOperationName(val *string) {
	if err := j.validateSetOperationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationName",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetRequestModels(val *map[string]*string) {
	if err := j.validateSetRequestModelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestModels",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetRequestParameters(val *map[string]interface{}) {
	if err := j.validateSetRequestParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestParameters",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetRequestValidatorId(val *string) {
	if err := j.validateSetRequestValidatorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestValidatorId",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethod)SetRestApiId(val *string) {
	if err := j.validateSetRestApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

// Generates CDKTF code for importing a ApigatewayMethod resource upon running "cdktf plan <stack-name>".
func ApigatewayMethod_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApigatewayMethod_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.apigatewayMethod.ApigatewayMethod",
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
func ApigatewayMethod_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigatewayMethod_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.apigatewayMethod.ApigatewayMethod",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigatewayMethod_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigatewayMethod_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.apigatewayMethod.ApigatewayMethod",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigatewayMethod_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigatewayMethod_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.apigatewayMethod.ApigatewayMethod",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApigatewayMethod_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.apigatewayMethod.ApigatewayMethod",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApigatewayMethod) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApigatewayMethod) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApigatewayMethod) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigatewayMethod) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigatewayMethod) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigatewayMethod) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigatewayMethod) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigatewayMethod) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigatewayMethod) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigatewayMethod) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigatewayMethod) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigatewayMethod) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApigatewayMethod) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigatewayMethod) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApigatewayMethod) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApigatewayMethod) PutIntegration(value *ApigatewayMethodIntegration) {
	if err := a.validatePutIntegrationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIntegration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigatewayMethod) PutMethodResponses(value interface{}) {
	if err := a.validatePutMethodResponsesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMethodResponses",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigatewayMethod) ResetApiKeyRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetApiKeyRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethod) ResetAuthorizationScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizationScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethod) ResetAuthorizationType() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethod) ResetAuthorizerId() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethod) ResetIntegration() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethod) ResetMethodResponses() {
	_jsii_.InvokeVoid(
		a,
		"resetMethodResponses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethod) ResetOperationName() {
	_jsii_.InvokeVoid(
		a,
		"resetOperationName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethod) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethod) ResetRequestModels() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestModels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethod) ResetRequestParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethod) ResetRequestValidatorId() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestValidatorId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethod) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayMethod) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayMethod) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayMethod) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

