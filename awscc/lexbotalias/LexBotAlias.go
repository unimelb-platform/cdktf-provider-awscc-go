package lexbotalias

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lexbotalias/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias awscc_lex_bot_alias}.
type LexBotAlias interface {
	cdktf.TerraformResource
	Arn() *string
	BotAliasId() *string
	BotAliasLocaleSettings() LexBotAliasBotAliasLocaleSettingsList
	BotAliasLocaleSettingsInput() interface{}
	BotAliasName() *string
	SetBotAliasName(val *string)
	BotAliasNameInput() *string
	BotAliasStatus() *string
	BotAliasTags() LexBotAliasBotAliasTagsList
	BotAliasTagsInput() interface{}
	BotId() *string
	SetBotId(val *string)
	BotIdInput() *string
	BotVersion() *string
	SetBotVersion(val *string)
	BotVersionInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ConversationLogSettings() LexBotAliasConversationLogSettingsOutputReference
	ConversationLogSettingsInput() interface{}
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
	SentimentAnalysisSettings() LexBotAliasSentimentAnalysisSettingsOutputReference
	SentimentAnalysisSettingsInput() interface{}
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
	PutBotAliasLocaleSettings(value interface{})
	PutBotAliasTags(value interface{})
	PutConversationLogSettings(value *LexBotAliasConversationLogSettings)
	PutSentimentAnalysisSettings(value *LexBotAliasSentimentAnalysisSettings)
	ResetBotAliasLocaleSettings()
	ResetBotAliasTags()
	ResetBotVersion()
	ResetConversationLogSettings()
	ResetDescription()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSentimentAnalysisSettings()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LexBotAlias
type jsiiProxy_LexBotAlias struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LexBotAlias) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) BotAliasId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botAliasId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) BotAliasLocaleSettings() LexBotAliasBotAliasLocaleSettingsList {
	var returns LexBotAliasBotAliasLocaleSettingsList
	_jsii_.Get(
		j,
		"botAliasLocaleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) BotAliasLocaleSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"botAliasLocaleSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) BotAliasName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botAliasName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) BotAliasNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botAliasNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) BotAliasStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botAliasStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) BotAliasTags() LexBotAliasBotAliasTagsList {
	var returns LexBotAliasBotAliasTagsList
	_jsii_.Get(
		j,
		"botAliasTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) BotAliasTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"botAliasTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) BotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) BotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) BotVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) BotVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) ConversationLogSettings() LexBotAliasConversationLogSettingsOutputReference {
	var returns LexBotAliasConversationLogSettingsOutputReference
	_jsii_.Get(
		j,
		"conversationLogSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) ConversationLogSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conversationLogSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) SentimentAnalysisSettings() LexBotAliasSentimentAnalysisSettingsOutputReference {
	var returns LexBotAliasSentimentAnalysisSettingsOutputReference
	_jsii_.Get(
		j,
		"sentimentAnalysisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) SentimentAnalysisSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sentimentAnalysisSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias awscc_lex_bot_alias} Resource.
func NewLexBotAlias(scope constructs.Construct, id *string, config *LexBotAliasConfig) LexBotAlias {
	_init_.Initialize()

	if err := validateNewLexBotAliasParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LexBotAlias{}

	_jsii_.Create(
		"awscc.lexBotAlias.LexBotAlias",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias awscc_lex_bot_alias} Resource.
func NewLexBotAlias_Override(l LexBotAlias, scope constructs.Construct, id *string, config *LexBotAliasConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lexBotAlias.LexBotAlias",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LexBotAlias)SetBotAliasName(val *string) {
	if err := j.validateSetBotAliasNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"botAliasName",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias)SetBotId(val *string) {
	if err := j.validateSetBotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"botId",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias)SetBotVersion(val *string) {
	if err := j.validateSetBotVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"botVersion",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a LexBotAlias resource upon running "cdktf plan <stack-name>".
func LexBotAlias_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLexBotAlias_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.lexBotAlias.LexBotAlias",
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
func LexBotAlias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLexBotAlias_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lexBotAlias.LexBotAlias",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LexBotAlias_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLexBotAlias_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lexBotAlias.LexBotAlias",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LexBotAlias_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLexBotAlias_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lexBotAlias.LexBotAlias",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LexBotAlias_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.lexBotAlias.LexBotAlias",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LexBotAlias) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LexBotAlias) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LexBotAlias) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotAlias) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotAlias) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotAlias) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotAlias) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotAlias) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotAlias) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotAlias) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotAlias) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotAlias) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LexBotAlias) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotAlias) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LexBotAlias) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LexBotAlias) PutBotAliasLocaleSettings(value interface{}) {
	if err := l.validatePutBotAliasLocaleSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putBotAliasLocaleSettings",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotAlias) PutBotAliasTags(value interface{}) {
	if err := l.validatePutBotAliasTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putBotAliasTags",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotAlias) PutConversationLogSettings(value *LexBotAliasConversationLogSettings) {
	if err := l.validatePutConversationLogSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putConversationLogSettings",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotAlias) PutSentimentAnalysisSettings(value *LexBotAliasSentimentAnalysisSettings) {
	if err := l.validatePutSentimentAnalysisSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSentimentAnalysisSettings",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotAlias) ResetBotAliasLocaleSettings() {
	_jsii_.InvokeVoid(
		l,
		"resetBotAliasLocaleSettings",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotAlias) ResetBotAliasTags() {
	_jsii_.InvokeVoid(
		l,
		"resetBotAliasTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotAlias) ResetBotVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetBotVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotAlias) ResetConversationLogSettings() {
	_jsii_.InvokeVoid(
		l,
		"resetConversationLogSettings",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotAlias) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotAlias) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotAlias) ResetSentimentAnalysisSettings() {
	_jsii_.InvokeVoid(
		l,
		"resetSentimentAnalysisSettings",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotAlias) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotAlias) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotAlias) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotAlias) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

