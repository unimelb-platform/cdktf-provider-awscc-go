package chatbotmicrosoftteamschannelconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/chatbotmicrosoftteamschannelconfiguration/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/chatbot_microsoft_teams_channel_configuration awscc_chatbot_microsoft_teams_channel_configuration}.
type ChatbotMicrosoftTeamsChannelConfiguration interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfigurationName() *string
	SetConfigurationName(val *string)
	ConfigurationNameInput() *string
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
	GuardrailPolicies() *[]*string
	SetGuardrailPolicies(val *[]*string)
	GuardrailPoliciesInput() *[]*string
	IamRoleArn() *string
	SetIamRoleArn(val *string)
	IamRoleArnInput() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoggingLevel() *string
	SetLoggingLevel(val *string)
	LoggingLevelInput() *string
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
	SnsTopicArns() *[]*string
	SetSnsTopicArns(val *[]*string)
	SnsTopicArnsInput() *[]*string
	TeamId() *string
	SetTeamId(val *string)
	TeamIdInput() *string
	TeamsChannelId() *string
	SetTeamsChannelId(val *string)
	TeamsChannelIdInput() *string
	TeamsTenantId() *string
	SetTeamsTenantId(val *string)
	TeamsTenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserRoleRequired() interface{}
	SetUserRoleRequired(val interface{})
	UserRoleRequiredInput() interface{}
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
	ResetGuardrailPolicies()
	ResetLoggingLevel()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSnsTopicArns()
	ResetUserRoleRequired()
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

// The jsii proxy struct for ChatbotMicrosoftTeamsChannelConfiguration
type jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) ConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) ConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) GuardrailPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guardrailPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) GuardrailPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guardrailPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) IamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) IamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) LoggingLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) LoggingLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) SnsTopicArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snsTopicArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) SnsTopicArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snsTopicArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) TeamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) TeamsChannelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamsChannelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) TeamsChannelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamsChannelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) TeamsTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamsTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) TeamsTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamsTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) UserRoleRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userRoleRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) UserRoleRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userRoleRequiredInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/chatbot_microsoft_teams_channel_configuration awscc_chatbot_microsoft_teams_channel_configuration} Resource.
func NewChatbotMicrosoftTeamsChannelConfiguration(scope constructs.Construct, id *string, config *ChatbotMicrosoftTeamsChannelConfigurationConfig) ChatbotMicrosoftTeamsChannelConfiguration {
	_init_.Initialize()

	if err := validateNewChatbotMicrosoftTeamsChannelConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration{}

	_jsii_.Create(
		"awscc.chatbotMicrosoftTeamsChannelConfiguration.ChatbotMicrosoftTeamsChannelConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/chatbot_microsoft_teams_channel_configuration awscc_chatbot_microsoft_teams_channel_configuration} Resource.
func NewChatbotMicrosoftTeamsChannelConfiguration_Override(c ChatbotMicrosoftTeamsChannelConfiguration, scope constructs.Construct, id *string, config *ChatbotMicrosoftTeamsChannelConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.chatbotMicrosoftTeamsChannelConfiguration.ChatbotMicrosoftTeamsChannelConfiguration",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration)SetConfigurationName(val *string) {
	if err := j.validateSetConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationName",
		val,
	)
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration)SetGuardrailPolicies(val *[]*string) {
	if err := j.validateSetGuardrailPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guardrailPolicies",
		val,
	)
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration)SetIamRoleArn(val *string) {
	if err := j.validateSetIamRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRoleArn",
		val,
	)
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration)SetLoggingLevel(val *string) {
	if err := j.validateSetLoggingLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingLevel",
		val,
	)
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration)SetSnsTopicArns(val *[]*string) {
	if err := j.validateSetSnsTopicArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snsTopicArns",
		val,
	)
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration)SetTeamId(val *string) {
	if err := j.validateSetTeamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration)SetTeamsChannelId(val *string) {
	if err := j.validateSetTeamsChannelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamsChannelId",
		val,
	)
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration)SetTeamsTenantId(val *string) {
	if err := j.validateSetTeamsTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamsTenantId",
		val,
	)
}

func (j *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration)SetUserRoleRequired(val interface{}) {
	if err := j.validateSetUserRoleRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userRoleRequired",
		val,
	)
}

// Generates CDKTF code for importing a ChatbotMicrosoftTeamsChannelConfiguration resource upon running "cdktf plan <stack-name>".
func ChatbotMicrosoftTeamsChannelConfiguration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateChatbotMicrosoftTeamsChannelConfiguration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.chatbotMicrosoftTeamsChannelConfiguration.ChatbotMicrosoftTeamsChannelConfiguration",
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
func ChatbotMicrosoftTeamsChannelConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateChatbotMicrosoftTeamsChannelConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.chatbotMicrosoftTeamsChannelConfiguration.ChatbotMicrosoftTeamsChannelConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ChatbotMicrosoftTeamsChannelConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateChatbotMicrosoftTeamsChannelConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.chatbotMicrosoftTeamsChannelConfiguration.ChatbotMicrosoftTeamsChannelConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ChatbotMicrosoftTeamsChannelConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateChatbotMicrosoftTeamsChannelConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.chatbotMicrosoftTeamsChannelConfiguration.ChatbotMicrosoftTeamsChannelConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ChatbotMicrosoftTeamsChannelConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.chatbotMicrosoftTeamsChannelConfiguration.ChatbotMicrosoftTeamsChannelConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) ResetGuardrailPolicies() {
	_jsii_.InvokeVoid(
		c,
		"resetGuardrailPolicies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) ResetLoggingLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetLoggingLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) ResetSnsTopicArns() {
	_jsii_.InvokeVoid(
		c,
		"resetSnsTopicArns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) ResetUserRoleRequired() {
	_jsii_.InvokeVoid(
		c,
		"resetUserRoleRequired",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChatbotMicrosoftTeamsChannelConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

