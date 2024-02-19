package gameliftgamesessionqueue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/gameliftgamesessionqueue/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_game_session_queue awscc_gamelift_game_session_queue}.
type GameliftGameSessionQueue interface {
	cdktf.TerraformResource
	Arn() *string
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
	CustomEventData() *string
	SetCustomEventData(val *string)
	CustomEventDataInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Destinations() GameliftGameSessionQueueDestinationsList
	DestinationsInput() interface{}
	FilterConfiguration() GameliftGameSessionQueueFilterConfigurationOutputReference
	FilterConfigurationInput() interface{}
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NotificationTarget() *string
	SetNotificationTarget(val *string)
	NotificationTargetInput() *string
	PlayerLatencyPolicies() GameliftGameSessionQueuePlayerLatencyPoliciesList
	PlayerLatencyPoliciesInput() interface{}
	PriorityConfiguration() GameliftGameSessionQueuePriorityConfigurationOutputReference
	PriorityConfigurationInput() interface{}
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
	Tags() GameliftGameSessionQueueTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeoutInSeconds() *float64
	SetTimeoutInSeconds(val *float64)
	TimeoutInSecondsInput() *float64
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
	PutDestinations(value interface{})
	PutFilterConfiguration(value *GameliftGameSessionQueueFilterConfiguration)
	PutPlayerLatencyPolicies(value interface{})
	PutPriorityConfiguration(value *GameliftGameSessionQueuePriorityConfiguration)
	PutTags(value interface{})
	ResetCustomEventData()
	ResetDestinations()
	ResetFilterConfiguration()
	ResetNotificationTarget()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlayerLatencyPolicies()
	ResetPriorityConfiguration()
	ResetTags()
	ResetTimeoutInSeconds()
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

// The jsii proxy struct for GameliftGameSessionQueue
type jsiiProxy_GameliftGameSessionQueue struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GameliftGameSessionQueue) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) CustomEventData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEventData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) CustomEventDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEventDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) Destinations() GameliftGameSessionQueueDestinationsList {
	var returns GameliftGameSessionQueueDestinationsList
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) DestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) FilterConfiguration() GameliftGameSessionQueueFilterConfigurationOutputReference {
	var returns GameliftGameSessionQueueFilterConfigurationOutputReference
	_jsii_.Get(
		j,
		"filterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) FilterConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) NotificationTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) NotificationTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) PlayerLatencyPolicies() GameliftGameSessionQueuePlayerLatencyPoliciesList {
	var returns GameliftGameSessionQueuePlayerLatencyPoliciesList
	_jsii_.Get(
		j,
		"playerLatencyPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) PlayerLatencyPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"playerLatencyPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) PriorityConfiguration() GameliftGameSessionQueuePriorityConfigurationOutputReference {
	var returns GameliftGameSessionQueuePriorityConfigurationOutputReference
	_jsii_.Get(
		j,
		"priorityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) PriorityConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"priorityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) Tags() GameliftGameSessionQueueTagsList {
	var returns GameliftGameSessionQueueTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameSessionQueue) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_game_session_queue awscc_gamelift_game_session_queue} Resource.
func NewGameliftGameSessionQueue(scope constructs.Construct, id *string, config *GameliftGameSessionQueueConfig) GameliftGameSessionQueue {
	_init_.Initialize()

	if err := validateNewGameliftGameSessionQueueParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameliftGameSessionQueue{}

	_jsii_.Create(
		"awscc.gameliftGameSessionQueue.GameliftGameSessionQueue",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_game_session_queue awscc_gamelift_game_session_queue} Resource.
func NewGameliftGameSessionQueue_Override(g GameliftGameSessionQueue, scope constructs.Construct, id *string, config *GameliftGameSessionQueueConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.gameliftGameSessionQueue.GameliftGameSessionQueue",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GameliftGameSessionQueue)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GameliftGameSessionQueue)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GameliftGameSessionQueue)SetCustomEventData(val *string) {
	if err := j.validateSetCustomEventDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customEventData",
		val,
	)
}

func (j *jsiiProxy_GameliftGameSessionQueue)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GameliftGameSessionQueue)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GameliftGameSessionQueue)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GameliftGameSessionQueue)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GameliftGameSessionQueue)SetNotificationTarget(val *string) {
	if err := j.validateSetNotificationTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationTarget",
		val,
	)
}

func (j *jsiiProxy_GameliftGameSessionQueue)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GameliftGameSessionQueue)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GameliftGameSessionQueue)SetTimeoutInSeconds(val *float64) {
	if err := j.validateSetTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInSeconds",
		val,
	)
}

// Generates CDKTF code for importing a GameliftGameSessionQueue resource upon running "cdktf plan <stack-name>".
func GameliftGameSessionQueue_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGameliftGameSessionQueue_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.gameliftGameSessionQueue.GameliftGameSessionQueue",
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
func GameliftGameSessionQueue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftGameSessionQueue_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftGameSessionQueue.GameliftGameSessionQueue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftGameSessionQueue_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftGameSessionQueue_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftGameSessionQueue.GameliftGameSessionQueue",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftGameSessionQueue_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftGameSessionQueue_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftGameSessionQueue.GameliftGameSessionQueue",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GameliftGameSessionQueue_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.gameliftGameSessionQueue.GameliftGameSessionQueue",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GameliftGameSessionQueue) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GameliftGameSessionQueue) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GameliftGameSessionQueue) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GameliftGameSessionQueue) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GameliftGameSessionQueue) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GameliftGameSessionQueue) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GameliftGameSessionQueue) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GameliftGameSessionQueue) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GameliftGameSessionQueue) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GameliftGameSessionQueue) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftGameSessionQueue) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GameliftGameSessionQueue) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) PutDestinations(value interface{}) {
	if err := g.validatePutDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDestinations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) PutFilterConfiguration(value *GameliftGameSessionQueueFilterConfiguration) {
	if err := g.validatePutFilterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFilterConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) PutPlayerLatencyPolicies(value interface{}) {
	if err := g.validatePutPlayerLatencyPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPlayerLatencyPolicies",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) PutPriorityConfiguration(value *GameliftGameSessionQueuePriorityConfiguration) {
	if err := g.validatePutPriorityConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPriorityConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) PutTags(value interface{}) {
	if err := g.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTags",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) ResetCustomEventData() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomEventData",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) ResetDestinations() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) ResetFilterConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetFilterConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) ResetNotificationTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetNotificationTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) ResetPlayerLatencyPolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetPlayerLatencyPolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) ResetPriorityConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetPriorityConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) ResetTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeoutInSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameSessionQueue) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftGameSessionQueue) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftGameSessionQueue) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftGameSessionQueue) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftGameSessionQueue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftGameSessionQueue) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

