package amplifybranch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/amplifybranch/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_branch awscc_amplify_branch}.
type AmplifyBranch interface {
	cdktf.TerraformResource
	AppId() *string
	SetAppId(val *string)
	AppIdInput() *string
	Arn() *string
	Backend() AmplifyBranchBackendOutputReference
	BackendInput() interface{}
	BasicAuthConfig() AmplifyBranchBasicAuthConfigOutputReference
	BasicAuthConfigInput() interface{}
	BranchName() *string
	SetBranchName(val *string)
	BranchNameInput() *string
	BuildSpec() *string
	SetBuildSpec(val *string)
	BuildSpecInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnableAutoBuild() interface{}
	SetEnableAutoBuild(val interface{})
	EnableAutoBuildInput() interface{}
	EnablePerformanceMode() interface{}
	SetEnablePerformanceMode(val interface{})
	EnablePerformanceModeInput() interface{}
	EnablePullRequestPreview() interface{}
	SetEnablePullRequestPreview(val interface{})
	EnablePullRequestPreviewInput() interface{}
	EnvironmentVariables() AmplifyBranchEnvironmentVariablesList
	EnvironmentVariablesInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	Framework() *string
	SetFramework(val *string)
	FrameworkInput() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	PullRequestEnvironmentName() *string
	SetPullRequestEnvironmentName(val *string)
	PullRequestEnvironmentNameInput() *string
	// Experimental.
	RawOverrides() interface{}
	Stage() *string
	SetStage(val *string)
	StageInput() *string
	Tags() AmplifyBranchTagsList
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
	PutBackend(value *AmplifyBranchBackend)
	PutBasicAuthConfig(value *AmplifyBranchBasicAuthConfig)
	PutEnvironmentVariables(value interface{})
	PutTags(value interface{})
	ResetBackend()
	ResetBasicAuthConfig()
	ResetBuildSpec()
	ResetDescription()
	ResetEnableAutoBuild()
	ResetEnablePerformanceMode()
	ResetEnablePullRequestPreview()
	ResetEnvironmentVariables()
	ResetFramework()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPullRequestEnvironmentName()
	ResetStage()
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

// The jsii proxy struct for AmplifyBranch
type jsiiProxy_AmplifyBranch struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AmplifyBranch) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) AppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Backend() AmplifyBranchBackendOutputReference {
	var returns AmplifyBranchBackendOutputReference
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) BackendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) BasicAuthConfig() AmplifyBranchBasicAuthConfigOutputReference {
	var returns AmplifyBranchBasicAuthConfigOutputReference
	_jsii_.Get(
		j,
		"basicAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) BasicAuthConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basicAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) BranchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) BranchNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) BuildSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) BuildSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnableAutoBuild() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnableAutoBuildInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoBuildInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnablePerformanceMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePerformanceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnablePerformanceModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePerformanceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnablePullRequestPreview() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePullRequestPreview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnablePullRequestPreviewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePullRequestPreviewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnvironmentVariables() AmplifyBranchEnvironmentVariablesList {
	var returns AmplifyBranchEnvironmentVariablesList
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnvironmentVariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Framework() *string {
	var returns *string
	_jsii_.Get(
		j,
		"framework",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) FrameworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) PullRequestEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pullRequestEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) PullRequestEnvironmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pullRequestEnvironmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) StageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Tags() AmplifyBranchTagsList {
	var returns AmplifyBranchTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_branch awscc_amplify_branch} Resource.
func NewAmplifyBranch(scope constructs.Construct, id *string, config *AmplifyBranchConfig) AmplifyBranch {
	_init_.Initialize()

	if err := validateNewAmplifyBranchParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AmplifyBranch{}

	_jsii_.Create(
		"awscc.amplifyBranch.AmplifyBranch",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_branch awscc_amplify_branch} Resource.
func NewAmplifyBranch_Override(a AmplifyBranch, scope constructs.Construct, id *string, config *AmplifyBranchConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.amplifyBranch.AmplifyBranch",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetAppId(val *string) {
	if err := j.validateSetAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetBranchName(val *string) {
	if err := j.validateSetBranchNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branchName",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetBuildSpec(val *string) {
	if err := j.validateSetBuildSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildSpec",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetEnableAutoBuild(val interface{}) {
	if err := j.validateSetEnableAutoBuildParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutoBuild",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetEnablePerformanceMode(val interface{}) {
	if err := j.validateSetEnablePerformanceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePerformanceMode",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetEnablePullRequestPreview(val interface{}) {
	if err := j.validateSetEnablePullRequestPreviewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePullRequestPreview",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetFramework(val *string) {
	if err := j.validateSetFrameworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"framework",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetPullRequestEnvironmentName(val *string) {
	if err := j.validateSetPullRequestEnvironmentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pullRequestEnvironmentName",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch)SetStage(val *string) {
	if err := j.validateSetStageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stage",
		val,
	)
}

// Generates CDKTF code for importing a AmplifyBranch resource upon running "cdktf plan <stack-name>".
func AmplifyBranch_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAmplifyBranch_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.amplifyBranch.AmplifyBranch",
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
func AmplifyBranch_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAmplifyBranch_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.amplifyBranch.AmplifyBranch",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AmplifyBranch_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAmplifyBranch_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.amplifyBranch.AmplifyBranch",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AmplifyBranch_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAmplifyBranch_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.amplifyBranch.AmplifyBranch",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AmplifyBranch_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.amplifyBranch.AmplifyBranch",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AmplifyBranch) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AmplifyBranch) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AmplifyBranch) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AmplifyBranch) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AmplifyBranch) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AmplifyBranch) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AmplifyBranch) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AmplifyBranch) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AmplifyBranch) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AmplifyBranch) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AmplifyBranch) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AmplifyBranch) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AmplifyBranch) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AmplifyBranch) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AmplifyBranch) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AmplifyBranch) PutBackend(value *AmplifyBranchBackend) {
	if err := a.validatePutBackendParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBackend",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AmplifyBranch) PutBasicAuthConfig(value *AmplifyBranchBasicAuthConfig) {
	if err := a.validatePutBasicAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBasicAuthConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AmplifyBranch) PutEnvironmentVariables(value interface{}) {
	if err := a.validatePutEnvironmentVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnvironmentVariables",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AmplifyBranch) PutTags(value interface{}) {
	if err := a.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetBackend() {
	_jsii_.InvokeVoid(
		a,
		"resetBackend",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetBasicAuthConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetBasicAuthConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetBuildSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetBuildSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetEnableAutoBuild() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableAutoBuild",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetEnablePerformanceMode() {
	_jsii_.InvokeVoid(
		a,
		"resetEnablePerformanceMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetEnablePullRequestPreview() {
	_jsii_.InvokeVoid(
		a,
		"resetEnablePullRequestPreview",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetFramework() {
	_jsii_.InvokeVoid(
		a,
		"resetFramework",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetPullRequestEnvironmentName() {
	_jsii_.InvokeVoid(
		a,
		"resetPullRequestEnvironmentName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetStage() {
	_jsii_.InvokeVoid(
		a,
		"resetStage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyBranch) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyBranch) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyBranch) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

