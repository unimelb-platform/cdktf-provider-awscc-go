package gameliftmatchmakingconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/gameliftmatchmakingconfiguration/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_matchmaking_configuration awscc_gamelift_matchmaking_configuration}.
type GameliftMatchmakingConfiguration interface {
	cdktf.TerraformResource
	AcceptanceRequired() interface{}
	SetAcceptanceRequired(val interface{})
	AcceptanceRequiredInput() interface{}
	AcceptanceTimeoutSeconds() *float64
	SetAcceptanceTimeoutSeconds(val *float64)
	AcceptanceTimeoutSecondsInput() *float64
	AdditionalPlayerCount() *float64
	SetAdditionalPlayerCount(val *float64)
	AdditionalPlayerCountInput() *float64
	Arn() *string
	BackfillMode() *string
	SetBackfillMode(val *string)
	BackfillModeInput() *string
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
	CreationTime() *string
	SetCreationTime(val *string)
	CreationTimeInput() *string
	CustomEventData() *string
	SetCustomEventData(val *string)
	CustomEventDataInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	FlexMatchMode() *string
	SetFlexMatchMode(val *string)
	FlexMatchModeInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GameProperties() GameliftMatchmakingConfigurationGamePropertiesList
	GamePropertiesInput() interface{}
	GameSessionData() *string
	SetGameSessionData(val *string)
	GameSessionDataInput() *string
	GameSessionQueueArns() *[]*string
	SetGameSessionQueueArns(val *[]*string)
	GameSessionQueueArnsInput() *[]*string
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
	NotificationTarget() *string
	SetNotificationTarget(val *string)
	NotificationTargetInput() *string
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
	RequestTimeoutSeconds() *float64
	SetRequestTimeoutSeconds(val *float64)
	RequestTimeoutSecondsInput() *float64
	RuleSetArn() *string
	SetRuleSetArn(val *string)
	RuleSetArnInput() *string
	RuleSetName() *string
	SetRuleSetName(val *string)
	RuleSetNameInput() *string
	Tags() GameliftMatchmakingConfigurationTagsList
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
	PutGameProperties(value interface{})
	PutTags(value interface{})
	ResetAcceptanceTimeoutSeconds()
	ResetAdditionalPlayerCount()
	ResetBackfillMode()
	ResetCreationTime()
	ResetCustomEventData()
	ResetDescription()
	ResetFlexMatchMode()
	ResetGameProperties()
	ResetGameSessionData()
	ResetGameSessionQueueArns()
	ResetNotificationTarget()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRuleSetArn()
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

// The jsii proxy struct for GameliftMatchmakingConfiguration
type jsiiProxy_GameliftMatchmakingConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) AcceptanceRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptanceRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) AcceptanceRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptanceRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) AcceptanceTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"acceptanceTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) AcceptanceTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"acceptanceTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) AdditionalPlayerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalPlayerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) AdditionalPlayerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalPlayerCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) BackfillMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backfillMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) BackfillModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backfillModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) CreationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) CustomEventData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEventData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) CustomEventDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEventDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) FlexMatchMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flexMatchMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) FlexMatchModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flexMatchModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) GameProperties() GameliftMatchmakingConfigurationGamePropertiesList {
	var returns GameliftMatchmakingConfigurationGamePropertiesList
	_jsii_.Get(
		j,
		"gameProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) GamePropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gamePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) GameSessionData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameSessionData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) GameSessionDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameSessionDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) GameSessionQueueArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gameSessionQueueArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) GameSessionQueueArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gameSessionQueueArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) NotificationTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) NotificationTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) RequestTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) RequestTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) RuleSetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) RuleSetArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSetArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) RuleSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) RuleSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) Tags() GameliftMatchmakingConfigurationTagsList {
	var returns GameliftMatchmakingConfigurationTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_matchmaking_configuration awscc_gamelift_matchmaking_configuration} Resource.
func NewGameliftMatchmakingConfiguration(scope constructs.Construct, id *string, config *GameliftMatchmakingConfigurationConfig) GameliftMatchmakingConfiguration {
	_init_.Initialize()

	if err := validateNewGameliftMatchmakingConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameliftMatchmakingConfiguration{}

	_jsii_.Create(
		"awscc.gameliftMatchmakingConfiguration.GameliftMatchmakingConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_matchmaking_configuration awscc_gamelift_matchmaking_configuration} Resource.
func NewGameliftMatchmakingConfiguration_Override(g GameliftMatchmakingConfiguration, scope constructs.Construct, id *string, config *GameliftMatchmakingConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.gameliftMatchmakingConfiguration.GameliftMatchmakingConfiguration",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetAcceptanceRequired(val interface{}) {
	if err := j.validateSetAcceptanceRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptanceRequired",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetAcceptanceTimeoutSeconds(val *float64) {
	if err := j.validateSetAcceptanceTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptanceTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetAdditionalPlayerCount(val *float64) {
	if err := j.validateSetAdditionalPlayerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalPlayerCount",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetBackfillMode(val *string) {
	if err := j.validateSetBackfillModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backfillMode",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetCreationTime(val *string) {
	if err := j.validateSetCreationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creationTime",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetCustomEventData(val *string) {
	if err := j.validateSetCustomEventDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customEventData",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetFlexMatchMode(val *string) {
	if err := j.validateSetFlexMatchModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flexMatchMode",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetGameSessionData(val *string) {
	if err := j.validateSetGameSessionDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gameSessionData",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetGameSessionQueueArns(val *[]*string) {
	if err := j.validateSetGameSessionQueueArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gameSessionQueueArns",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetNotificationTarget(val *string) {
	if err := j.validateSetNotificationTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationTarget",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetRequestTimeoutSeconds(val *float64) {
	if err := j.validateSetRequestTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetRuleSetArn(val *string) {
	if err := j.validateSetRuleSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleSetArn",
		val,
	)
}

func (j *jsiiProxy_GameliftMatchmakingConfiguration)SetRuleSetName(val *string) {
	if err := j.validateSetRuleSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleSetName",
		val,
	)
}

// Generates CDKTF code for importing a GameliftMatchmakingConfiguration resource upon running "cdktf plan <stack-name>".
func GameliftMatchmakingConfiguration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGameliftMatchmakingConfiguration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.gameliftMatchmakingConfiguration.GameliftMatchmakingConfiguration",
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
func GameliftMatchmakingConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftMatchmakingConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftMatchmakingConfiguration.GameliftMatchmakingConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftMatchmakingConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftMatchmakingConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftMatchmakingConfiguration.GameliftMatchmakingConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftMatchmakingConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftMatchmakingConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftMatchmakingConfiguration.GameliftMatchmakingConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GameliftMatchmakingConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.gameliftMatchmakingConfiguration.GameliftMatchmakingConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) PutGameProperties(value interface{}) {
	if err := g.validatePutGamePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGameProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) PutTags(value interface{}) {
	if err := g.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTags",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ResetAcceptanceTimeoutSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetAcceptanceTimeoutSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ResetAdditionalPlayerCount() {
	_jsii_.InvokeVoid(
		g,
		"resetAdditionalPlayerCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ResetBackfillMode() {
	_jsii_.InvokeVoid(
		g,
		"resetBackfillMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ResetCreationTime() {
	_jsii_.InvokeVoid(
		g,
		"resetCreationTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ResetCustomEventData() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomEventData",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ResetFlexMatchMode() {
	_jsii_.InvokeVoid(
		g,
		"resetFlexMatchMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ResetGameProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetGameProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ResetGameSessionData() {
	_jsii_.InvokeVoid(
		g,
		"resetGameSessionData",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ResetGameSessionQueueArns() {
	_jsii_.InvokeVoid(
		g,
		"resetGameSessionQueueArns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ResetNotificationTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetNotificationTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ResetRuleSetArn() {
	_jsii_.InvokeVoid(
		g,
		"resetRuleSetArn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftMatchmakingConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

