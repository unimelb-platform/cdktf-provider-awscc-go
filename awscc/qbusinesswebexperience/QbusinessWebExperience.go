package qbusinesswebexperience

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/qbusinesswebexperience/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/qbusiness_web_experience awscc_qbusiness_web_experience}.
type QbusinessWebExperience interface {
	cdktf.TerraformResource
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	BrowserExtensionConfiguration() QbusinessWebExperienceBrowserExtensionConfigurationOutputReference
	BrowserExtensionConfigurationInput() interface{}
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
	CreatedAt() *string
	CustomizationConfiguration() QbusinessWebExperienceCustomizationConfigurationOutputReference
	CustomizationConfigurationInput() interface{}
	DefaultEndpoint() *string
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
	IdentityProviderConfiguration() QbusinessWebExperienceIdentityProviderConfigurationOutputReference
	IdentityProviderConfigurationInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Origins() *[]*string
	SetOrigins(val *[]*string)
	OriginsInput() *[]*string
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
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	SamplePromptsControlMode() *string
	SetSamplePromptsControlMode(val *string)
	SamplePromptsControlModeInput() *string
	Status() *string
	Subtitle() *string
	SetSubtitle(val *string)
	SubtitleInput() *string
	Tags() QbusinessWebExperienceTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	UpdatedAt() *string
	WebExperienceArn() *string
	WebExperienceId() *string
	WelcomeMessage() *string
	SetWelcomeMessage(val *string)
	WelcomeMessageInput() *string
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
	PutBrowserExtensionConfiguration(value *QbusinessWebExperienceBrowserExtensionConfiguration)
	PutCustomizationConfiguration(value *QbusinessWebExperienceCustomizationConfiguration)
	PutIdentityProviderConfiguration(value *QbusinessWebExperienceIdentityProviderConfiguration)
	PutTags(value interface{})
	ResetBrowserExtensionConfiguration()
	ResetCustomizationConfiguration()
	ResetIdentityProviderConfiguration()
	ResetOrigins()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRoleArn()
	ResetSamplePromptsControlMode()
	ResetSubtitle()
	ResetTags()
	ResetTitle()
	ResetWelcomeMessage()
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

// The jsii proxy struct for QbusinessWebExperience
type jsiiProxy_QbusinessWebExperience struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_QbusinessWebExperience) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) BrowserExtensionConfiguration() QbusinessWebExperienceBrowserExtensionConfigurationOutputReference {
	var returns QbusinessWebExperienceBrowserExtensionConfigurationOutputReference
	_jsii_.Get(
		j,
		"browserExtensionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) BrowserExtensionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browserExtensionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) CustomizationConfiguration() QbusinessWebExperienceCustomizationConfigurationOutputReference {
	var returns QbusinessWebExperienceCustomizationConfigurationOutputReference
	_jsii_.Get(
		j,
		"customizationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) CustomizationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customizationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) DefaultEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) IdentityProviderConfiguration() QbusinessWebExperienceIdentityProviderConfigurationOutputReference {
	var returns QbusinessWebExperienceIdentityProviderConfigurationOutputReference
	_jsii_.Get(
		j,
		"identityProviderConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) IdentityProviderConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityProviderConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) Origins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"origins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) OriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) SamplePromptsControlMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samplePromptsControlMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) SamplePromptsControlModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samplePromptsControlModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) Subtitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) SubtitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) Tags() QbusinessWebExperienceTagsList {
	var returns QbusinessWebExperienceTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) WebExperienceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webExperienceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) WebExperienceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webExperienceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) WelcomeMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"welcomeMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperience) WelcomeMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"welcomeMessageInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/qbusiness_web_experience awscc_qbusiness_web_experience} Resource.
func NewQbusinessWebExperience(scope constructs.Construct, id *string, config *QbusinessWebExperienceConfig) QbusinessWebExperience {
	_init_.Initialize()

	if err := validateNewQbusinessWebExperienceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_QbusinessWebExperience{}

	_jsii_.Create(
		"awscc.qbusinessWebExperience.QbusinessWebExperience",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/qbusiness_web_experience awscc_qbusiness_web_experience} Resource.
func NewQbusinessWebExperience_Override(q QbusinessWebExperience, scope constructs.Construct, id *string, config *QbusinessWebExperienceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.qbusinessWebExperience.QbusinessWebExperience",
		[]interface{}{scope, id, config},
		q,
	)
}

func (j *jsiiProxy_QbusinessWebExperience)SetApplicationId(val *string) {
	if err := j.validateSetApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperience)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperience)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperience)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperience)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperience)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperience)SetOrigins(val *[]*string) {
	if err := j.validateSetOriginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"origins",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperience)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperience)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperience)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperience)SetSamplePromptsControlMode(val *string) {
	if err := j.validateSetSamplePromptsControlModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samplePromptsControlMode",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperience)SetSubtitle(val *string) {
	if err := j.validateSetSubtitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subtitle",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperience)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperience)SetWelcomeMessage(val *string) {
	if err := j.validateSetWelcomeMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"welcomeMessage",
		val,
	)
}

// Generates CDKTF code for importing a QbusinessWebExperience resource upon running "cdktf plan <stack-name>".
func QbusinessWebExperience_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateQbusinessWebExperience_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.qbusinessWebExperience.QbusinessWebExperience",
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
func QbusinessWebExperience_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQbusinessWebExperience_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.qbusinessWebExperience.QbusinessWebExperience",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QbusinessWebExperience_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQbusinessWebExperience_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.qbusinessWebExperience.QbusinessWebExperience",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QbusinessWebExperience_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQbusinessWebExperience_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.qbusinessWebExperience.QbusinessWebExperience",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func QbusinessWebExperience_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.qbusinessWebExperience.QbusinessWebExperience",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) AddMoveTarget(moveTarget *string) {
	if err := q.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (q *jsiiProxy_QbusinessWebExperience) AddOverride(path *string, value interface{}) {
	if err := q.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (q *jsiiProxy_QbusinessWebExperience) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := q.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (q *jsiiProxy_QbusinessWebExperience) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) MoveFromId(id *string) {
	if err := q.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveFromId",
		[]interface{}{id},
	)
}

func (q *jsiiProxy_QbusinessWebExperience) MoveTo(moveTarget *string, index interface{}) {
	if err := q.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (q *jsiiProxy_QbusinessWebExperience) MoveToId(id *string) {
	if err := q.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveToId",
		[]interface{}{id},
	)
}

func (q *jsiiProxy_QbusinessWebExperience) OverrideLogicalId(newLogicalId *string) {
	if err := q.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (q *jsiiProxy_QbusinessWebExperience) PutBrowserExtensionConfiguration(value *QbusinessWebExperienceBrowserExtensionConfiguration) {
	if err := q.validatePutBrowserExtensionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putBrowserExtensionConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessWebExperience) PutCustomizationConfiguration(value *QbusinessWebExperienceCustomizationConfiguration) {
	if err := q.validatePutCustomizationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putCustomizationConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessWebExperience) PutIdentityProviderConfiguration(value *QbusinessWebExperienceIdentityProviderConfiguration) {
	if err := q.validatePutIdentityProviderConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putIdentityProviderConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessWebExperience) PutTags(value interface{}) {
	if err := q.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTags",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessWebExperience) ResetBrowserExtensionConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetBrowserExtensionConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessWebExperience) ResetCustomizationConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetCustomizationConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessWebExperience) ResetIdentityProviderConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetIdentityProviderConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessWebExperience) ResetOrigins() {
	_jsii_.InvokeVoid(
		q,
		"resetOrigins",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessWebExperience) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		q,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessWebExperience) ResetRoleArn() {
	_jsii_.InvokeVoid(
		q,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessWebExperience) ResetSamplePromptsControlMode() {
	_jsii_.InvokeVoid(
		q,
		"resetSamplePromptsControlMode",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessWebExperience) ResetSubtitle() {
	_jsii_.InvokeVoid(
		q,
		"resetSubtitle",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessWebExperience) ResetTags() {
	_jsii_.InvokeVoid(
		q,
		"resetTags",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessWebExperience) ResetTitle() {
	_jsii_.InvokeVoid(
		q,
		"resetTitle",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessWebExperience) ResetWelcomeMessage() {
	_jsii_.InvokeVoid(
		q,
		"resetWelcomeMessage",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessWebExperience) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperience) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

