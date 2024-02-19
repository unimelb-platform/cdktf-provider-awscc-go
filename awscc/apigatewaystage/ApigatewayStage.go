package apigatewaystage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/apigatewaystage/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage awscc_apigateway_stage}.
type ApigatewayStage interface {
	cdktf.TerraformResource
	AccessLogSetting() ApigatewayStageAccessLogSettingOutputReference
	AccessLogSettingInput() interface{}
	CacheClusterEnabled() interface{}
	SetCacheClusterEnabled(val interface{})
	CacheClusterEnabledInput() interface{}
	CacheClusterSize() *string
	SetCacheClusterSize(val *string)
	CacheClusterSizeInput() *string
	CanarySetting() ApigatewayStageCanarySettingOutputReference
	CanarySettingInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCertificateId() *string
	SetClientCertificateId(val *string)
	ClientCertificateIdInput() *string
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
	DeploymentId() *string
	SetDeploymentId(val *string)
	DeploymentIdInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DocumentationVersion() *string
	SetDocumentationVersion(val *string)
	DocumentationVersionInput() *string
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
	MethodSettings() ApigatewayStageMethodSettingsList
	MethodSettingsInput() interface{}
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
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
	StageName() *string
	SetStageName(val *string)
	StageNameInput() *string
	Tags() ApigatewayStageTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TracingEnabled() interface{}
	SetTracingEnabled(val interface{})
	TracingEnabledInput() interface{}
	Variables() *map[string]*string
	SetVariables(val *map[string]*string)
	VariablesInput() *map[string]*string
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
	PutAccessLogSetting(value *ApigatewayStageAccessLogSetting)
	PutCanarySetting(value *ApigatewayStageCanarySetting)
	PutMethodSettings(value interface{})
	PutTags(value interface{})
	ResetAccessLogSetting()
	ResetCacheClusterEnabled()
	ResetCacheClusterSize()
	ResetCanarySetting()
	ResetClientCertificateId()
	ResetDeploymentId()
	ResetDescription()
	ResetDocumentationVersion()
	ResetMethodSettings()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStageName()
	ResetTags()
	ResetTracingEnabled()
	ResetVariables()
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

// The jsii proxy struct for ApigatewayStage
type jsiiProxy_ApigatewayStage struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApigatewayStage) AccessLogSetting() ApigatewayStageAccessLogSettingOutputReference {
	var returns ApigatewayStageAccessLogSettingOutputReference
	_jsii_.Get(
		j,
		"accessLogSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) AccessLogSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessLogSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) CacheClusterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheClusterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) CacheClusterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheClusterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) CacheClusterSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheClusterSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) CacheClusterSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheClusterSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) CanarySetting() ApigatewayStageCanarySettingOutputReference {
	var returns ApigatewayStageCanarySettingOutputReference
	_jsii_.Get(
		j,
		"canarySetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) CanarySettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canarySettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) ClientCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) ClientCertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) DeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) DeploymentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) DocumentationVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentationVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) DocumentationVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentationVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) MethodSettings() ApigatewayStageMethodSettingsList {
	var returns ApigatewayStageMethodSettingsList
	_jsii_.Get(
		j,
		"methodSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) MethodSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"methodSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) StageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) Tags() ApigatewayStageTagsList {
	var returns ApigatewayStageTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) TracingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) TracingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) Variables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayStage) VariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variablesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage awscc_apigateway_stage} Resource.
func NewApigatewayStage(scope constructs.Construct, id *string, config *ApigatewayStageConfig) ApigatewayStage {
	_init_.Initialize()

	if err := validateNewApigatewayStageParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigatewayStage{}

	_jsii_.Create(
		"awscc.apigatewayStage.ApigatewayStage",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage awscc_apigateway_stage} Resource.
func NewApigatewayStage_Override(a ApigatewayStage, scope constructs.Construct, id *string, config *ApigatewayStageConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.apigatewayStage.ApigatewayStage",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetCacheClusterEnabled(val interface{}) {
	if err := j.validateSetCacheClusterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheClusterEnabled",
		val,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetCacheClusterSize(val *string) {
	if err := j.validateSetCacheClusterSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheClusterSize",
		val,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetClientCertificateId(val *string) {
	if err := j.validateSetClientCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateId",
		val,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetDeploymentId(val *string) {
	if err := j.validateSetDeploymentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentId",
		val,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetDocumentationVersion(val *string) {
	if err := j.validateSetDocumentationVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentationVersion",
		val,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetRestApiId(val *string) {
	if err := j.validateSetRestApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetStageName(val *string) {
	if err := j.validateSetStageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stageName",
		val,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetTracingEnabled(val interface{}) {
	if err := j.validateSetTracingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tracingEnabled",
		val,
	)
}

func (j *jsiiProxy_ApigatewayStage)SetVariables(val *map[string]*string) {
	if err := j.validateSetVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variables",
		val,
	)
}

// Generates CDKTF code for importing a ApigatewayStage resource upon running "cdktf plan <stack-name>".
func ApigatewayStage_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApigatewayStage_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.apigatewayStage.ApigatewayStage",
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
func ApigatewayStage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigatewayStage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.apigatewayStage.ApigatewayStage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigatewayStage_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigatewayStage_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.apigatewayStage.ApigatewayStage",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigatewayStage_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigatewayStage_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.apigatewayStage.ApigatewayStage",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApigatewayStage_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.apigatewayStage.ApigatewayStage",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApigatewayStage) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApigatewayStage) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApigatewayStage) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigatewayStage) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigatewayStage) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigatewayStage) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigatewayStage) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigatewayStage) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigatewayStage) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigatewayStage) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigatewayStage) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigatewayStage) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayStage) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApigatewayStage) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigatewayStage) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApigatewayStage) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApigatewayStage) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApigatewayStage) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApigatewayStage) PutAccessLogSetting(value *ApigatewayStageAccessLogSetting) {
	if err := a.validatePutAccessLogSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessLogSetting",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigatewayStage) PutCanarySetting(value *ApigatewayStageCanarySetting) {
	if err := a.validatePutCanarySettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCanarySetting",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigatewayStage) PutMethodSettings(value interface{}) {
	if err := a.validatePutMethodSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMethodSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigatewayStage) PutTags(value interface{}) {
	if err := a.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigatewayStage) ResetAccessLogSetting() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessLogSetting",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayStage) ResetCacheClusterEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheClusterEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayStage) ResetCacheClusterSize() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheClusterSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayStage) ResetCanarySetting() {
	_jsii_.InvokeVoid(
		a,
		"resetCanarySetting",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayStage) ResetClientCertificateId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificateId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayStage) ResetDeploymentId() {
	_jsii_.InvokeVoid(
		a,
		"resetDeploymentId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayStage) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayStage) ResetDocumentationVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetDocumentationVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayStage) ResetMethodSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetMethodSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayStage) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayStage) ResetStageName() {
	_jsii_.InvokeVoid(
		a,
		"resetStageName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayStage) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayStage) ResetTracingEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTracingEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayStage) ResetVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayStage) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayStage) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayStage) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayStage) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayStage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayStage) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

