package appsyncchannelnamespace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/appsyncchannelnamespace/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_channel_namespace awscc_appsync_channel_namespace}.
type AppsyncChannelNamespace interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChannelNamespaceArn() *string
	CodeHandlers() *string
	SetCodeHandlers(val *string)
	CodeHandlersInput() *string
	CodeS3Location() *string
	SetCodeS3Location(val *string)
	CodeS3LocationInput() *string
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
	HandlerConfigs() AppsyncChannelNamespaceHandlerConfigsOutputReference
	HandlerConfigsInput() interface{}
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
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublishAuthModes() AppsyncChannelNamespacePublishAuthModesList
	PublishAuthModesInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	SubscribeAuthModes() AppsyncChannelNamespaceSubscribeAuthModesList
	SubscribeAuthModesInput() interface{}
	Tags() AppsyncChannelNamespaceTagsList
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
	PutHandlerConfigs(value *AppsyncChannelNamespaceHandlerConfigs)
	PutPublishAuthModes(value interface{})
	PutSubscribeAuthModes(value interface{})
	PutTags(value interface{})
	ResetCodeHandlers()
	ResetCodeS3Location()
	ResetHandlerConfigs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublishAuthModes()
	ResetSubscribeAuthModes()
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

// The jsii proxy struct for AppsyncChannelNamespace
type jsiiProxy_AppsyncChannelNamespace struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppsyncChannelNamespace) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) ChannelNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelNamespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) CodeHandlers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeHandlers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) CodeHandlersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeHandlersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) CodeS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) CodeS3LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeS3LocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) HandlerConfigs() AppsyncChannelNamespaceHandlerConfigsOutputReference {
	var returns AppsyncChannelNamespaceHandlerConfigsOutputReference
	_jsii_.Get(
		j,
		"handlerConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) HandlerConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"handlerConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) PublishAuthModes() AppsyncChannelNamespacePublishAuthModesList {
	var returns AppsyncChannelNamespacePublishAuthModesList
	_jsii_.Get(
		j,
		"publishAuthModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) PublishAuthModesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishAuthModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) SubscribeAuthModes() AppsyncChannelNamespaceSubscribeAuthModesList {
	var returns AppsyncChannelNamespaceSubscribeAuthModesList
	_jsii_.Get(
		j,
		"subscribeAuthModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) SubscribeAuthModesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subscribeAuthModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) Tags() AppsyncChannelNamespaceTagsList {
	var returns AppsyncChannelNamespaceTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespace) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_channel_namespace awscc_appsync_channel_namespace} Resource.
func NewAppsyncChannelNamespace(scope constructs.Construct, id *string, config *AppsyncChannelNamespaceConfig) AppsyncChannelNamespace {
	_init_.Initialize()

	if err := validateNewAppsyncChannelNamespaceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncChannelNamespace{}

	_jsii_.Create(
		"awscc.appsyncChannelNamespace.AppsyncChannelNamespace",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_channel_namespace awscc_appsync_channel_namespace} Resource.
func NewAppsyncChannelNamespace_Override(a AppsyncChannelNamespace, scope constructs.Construct, id *string, config *AppsyncChannelNamespaceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.appsyncChannelNamespace.AppsyncChannelNamespace",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespace)SetApiId(val *string) {
	if err := j.validateSetApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespace)SetCodeHandlers(val *string) {
	if err := j.validateSetCodeHandlersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeHandlers",
		val,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespace)SetCodeS3Location(val *string) {
	if err := j.validateSetCodeS3LocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeS3Location",
		val,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespace)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespace)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespace)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespace)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespace)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespace)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespace)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespace)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a AppsyncChannelNamespace resource upon running "cdktf plan <stack-name>".
func AppsyncChannelNamespace_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAppsyncChannelNamespace_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.appsyncChannelNamespace.AppsyncChannelNamespace",
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
func AppsyncChannelNamespace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncChannelNamespace_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.appsyncChannelNamespace.AppsyncChannelNamespace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppsyncChannelNamespace_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncChannelNamespace_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.appsyncChannelNamespace.AppsyncChannelNamespace",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppsyncChannelNamespace_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncChannelNamespace_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.appsyncChannelNamespace.AppsyncChannelNamespace",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppsyncChannelNamespace_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.appsyncChannelNamespace.AppsyncChannelNamespace",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppsyncChannelNamespace) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppsyncChannelNamespace) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncChannelNamespace) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppsyncChannelNamespace) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppsyncChannelNamespace) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppsyncChannelNamespace) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppsyncChannelNamespace) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppsyncChannelNamespace) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppsyncChannelNamespace) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppsyncChannelNamespace) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncChannelNamespace) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncChannelNamespace) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) PutHandlerConfigs(value *AppsyncChannelNamespaceHandlerConfigs) {
	if err := a.validatePutHandlerConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHandlerConfigs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) PutPublishAuthModes(value interface{}) {
	if err := a.validatePutPublishAuthModesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPublishAuthModes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) PutSubscribeAuthModes(value interface{}) {
	if err := a.validatePutSubscribeAuthModesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSubscribeAuthModes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) PutTags(value interface{}) {
	if err := a.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) ResetCodeHandlers() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeHandlers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) ResetCodeS3Location() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeS3Location",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) ResetHandlerConfigs() {
	_jsii_.InvokeVoid(
		a,
		"resetHandlerConfigs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) ResetPublishAuthModes() {
	_jsii_.InvokeVoid(
		a,
		"resetPublishAuthModes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) ResetSubscribeAuthModes() {
	_jsii_.InvokeVoid(
		a,
		"resetSubscribeAuthModes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncChannelNamespace) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncChannelNamespace) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncChannelNamespace) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncChannelNamespace) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncChannelNamespace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncChannelNamespace) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

