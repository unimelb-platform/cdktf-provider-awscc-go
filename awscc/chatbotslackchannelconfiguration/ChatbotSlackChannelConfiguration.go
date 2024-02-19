package chatbotslackchannelconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/chatbotslackchannelconfiguration/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/chatbot_slack_channel_configuration awscc_chatbot_slack_channel_configuration}.
type ChatbotSlackChannelConfiguration interface {
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
	SlackChannelId() *string
	SetSlackChannelId(val *string)
	SlackChannelIdInput() *string
	SlackWorkspaceId() *string
	SetSlackWorkspaceId(val *string)
	SlackWorkspaceIdInput() *string
	SnsTopicArns() *[]*string
	SetSnsTopicArns(val *[]*string)
	SnsTopicArnsInput() *[]*string
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
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
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ChatbotSlackChannelConfiguration
type jsiiProxy_ChatbotSlackChannelConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) ConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) ConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) GuardrailPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guardrailPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) GuardrailPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guardrailPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) IamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) IamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) LoggingLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) LoggingLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) SlackChannelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackChannelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) SlackChannelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackChannelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) SlackWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) SlackWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) SnsTopicArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snsTopicArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) SnsTopicArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snsTopicArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) UserRoleRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userRoleRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration) UserRoleRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userRoleRequiredInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/chatbot_slack_channel_configuration awscc_chatbot_slack_channel_configuration} Resource.
func NewChatbotSlackChannelConfiguration(scope constructs.Construct, id *string, config *ChatbotSlackChannelConfigurationConfig) ChatbotSlackChannelConfiguration {
	_init_.Initialize()

	if err := validateNewChatbotSlackChannelConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChatbotSlackChannelConfiguration{}

	_jsii_.Create(
		"awscc.chatbotSlackChannelConfiguration.ChatbotSlackChannelConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/chatbot_slack_channel_configuration awscc_chatbot_slack_channel_configuration} Resource.
func NewChatbotSlackChannelConfiguration_Override(c ChatbotSlackChannelConfiguration, scope constructs.Construct, id *string, config *ChatbotSlackChannelConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.chatbotSlackChannelConfiguration.ChatbotSlackChannelConfiguration",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration)SetConfigurationName(val *string) {
	if err := j.validateSetConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationName",
		val,
	)
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration)SetGuardrailPolicies(val *[]*string) {
	if err := j.validateSetGuardrailPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guardrailPolicies",
		val,
	)
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration)SetIamRoleArn(val *string) {
	if err := j.validateSetIamRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRoleArn",
		val,
	)
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration)SetLoggingLevel(val *string) {
	if err := j.validateSetLoggingLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingLevel",
		val,
	)
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration)SetSlackChannelId(val *string) {
	if err := j.validateSetSlackChannelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slackChannelId",
		val,
	)
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration)SetSlackWorkspaceId(val *string) {
	if err := j.validateSetSlackWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slackWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration)SetSnsTopicArns(val *[]*string) {
	if err := j.validateSetSnsTopicArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snsTopicArns",
		val,
	)
}

func (j *jsiiProxy_ChatbotSlackChannelConfiguration)SetUserRoleRequired(val interface{}) {
	if err := j.validateSetUserRoleRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userRoleRequired",
		val,
	)
}

// Generates CDKTF code for importing a ChatbotSlackChannelConfiguration resource upon running "cdktf plan <stack-name>".
func ChatbotSlackChannelConfiguration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateChatbotSlackChannelConfiguration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.chatbotSlackChannelConfiguration.ChatbotSlackChannelConfiguration",
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
func ChatbotSlackChannelConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateChatbotSlackChannelConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.chatbotSlackChannelConfiguration.ChatbotSlackChannelConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ChatbotSlackChannelConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateChatbotSlackChannelConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.chatbotSlackChannelConfiguration.ChatbotSlackChannelConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ChatbotSlackChannelConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateChatbotSlackChannelConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.chatbotSlackChannelConfiguration.ChatbotSlackChannelConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ChatbotSlackChannelConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.chatbotSlackChannelConfiguration.ChatbotSlackChannelConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) ResetGuardrailPolicies() {
	_jsii_.InvokeVoid(
		c,
		"resetGuardrailPolicies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) ResetLoggingLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetLoggingLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) ResetSnsTopicArns() {
	_jsii_.InvokeVoid(
		c,
		"resetSnsTopicArns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) ResetUserRoleRequired() {
	_jsii_.InvokeVoid(
		c,
		"resetUserRoleRequired",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChatbotSlackChannelConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

