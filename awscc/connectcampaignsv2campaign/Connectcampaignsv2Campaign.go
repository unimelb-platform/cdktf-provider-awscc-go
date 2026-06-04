package connectcampaignsv2campaign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/connectcampaignsv2campaign/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign awscc_connectcampaignsv2_campaign}.
type Connectcampaignsv2Campaign interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChannelSubtypeConfig() Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference
	ChannelSubtypeConfigInput() interface{}
	CommunicationLimitsOverride() Connectcampaignsv2CampaignCommunicationLimitsOverrideOutputReference
	CommunicationLimitsOverrideInput() interface{}
	CommunicationTimeConfig() Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference
	CommunicationTimeConfigInput() interface{}
	ConnectCampaignFlowArn() *string
	SetConnectCampaignFlowArn(val *string)
	ConnectCampaignFlowArnInput() *string
	ConnectInstanceId() *string
	SetConnectInstanceId(val *string)
	ConnectInstanceIdInput() *string
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
	// Experimental.
	RawOverrides() interface{}
	Schedule() Connectcampaignsv2CampaignScheduleOutputReference
	ScheduleInput() interface{}
	Source() Connectcampaignsv2CampaignSourceOutputReference
	SourceInput() interface{}
	Tags() Connectcampaignsv2CampaignTagsList
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
	PutChannelSubtypeConfig(value *Connectcampaignsv2CampaignChannelSubtypeConfig)
	PutCommunicationLimitsOverride(value *Connectcampaignsv2CampaignCommunicationLimitsOverride)
	PutCommunicationTimeConfig(value *Connectcampaignsv2CampaignCommunicationTimeConfig)
	PutSchedule(value *Connectcampaignsv2CampaignSchedule)
	PutSource(value *Connectcampaignsv2CampaignSource)
	PutTags(value interface{})
	ResetCommunicationLimitsOverride()
	ResetCommunicationTimeConfig()
	ResetConnectCampaignFlowArn()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSchedule()
	ResetSource()
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

// The jsii proxy struct for Connectcampaignsv2Campaign
type jsiiProxy_Connectcampaignsv2Campaign struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) ChannelSubtypeConfig() Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference {
	var returns Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference
	_jsii_.Get(
		j,
		"channelSubtypeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) ChannelSubtypeConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"channelSubtypeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) CommunicationLimitsOverride() Connectcampaignsv2CampaignCommunicationLimitsOverrideOutputReference {
	var returns Connectcampaignsv2CampaignCommunicationLimitsOverrideOutputReference
	_jsii_.Get(
		j,
		"communicationLimitsOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) CommunicationLimitsOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"communicationLimitsOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) CommunicationTimeConfig() Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference {
	var returns Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference
	_jsii_.Get(
		j,
		"communicationTimeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) CommunicationTimeConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"communicationTimeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) ConnectCampaignFlowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectCampaignFlowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) ConnectCampaignFlowArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectCampaignFlowArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) ConnectInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) ConnectInstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectInstanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) Schedule() Connectcampaignsv2CampaignScheduleOutputReference {
	var returns Connectcampaignsv2CampaignScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) Source() Connectcampaignsv2CampaignSourceOutputReference {
	var returns Connectcampaignsv2CampaignSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) SourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) Tags() Connectcampaignsv2CampaignTagsList {
	var returns Connectcampaignsv2CampaignTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2Campaign) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign awscc_connectcampaignsv2_campaign} Resource.
func NewConnectcampaignsv2Campaign(scope constructs.Construct, id *string, config *Connectcampaignsv2CampaignConfig) Connectcampaignsv2Campaign {
	_init_.Initialize()

	if err := validateNewConnectcampaignsv2CampaignParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Connectcampaignsv2Campaign{}

	_jsii_.Create(
		"awscc.connectcampaignsv2Campaign.Connectcampaignsv2Campaign",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign awscc_connectcampaignsv2_campaign} Resource.
func NewConnectcampaignsv2Campaign_Override(c Connectcampaignsv2Campaign, scope constructs.Construct, id *string, config *Connectcampaignsv2CampaignConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.connectcampaignsv2Campaign.Connectcampaignsv2Campaign",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_Connectcampaignsv2Campaign)SetConnectCampaignFlowArn(val *string) {
	if err := j.validateSetConnectCampaignFlowArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectCampaignFlowArn",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2Campaign)SetConnectInstanceId(val *string) {
	if err := j.validateSetConnectInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectInstanceId",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2Campaign)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2Campaign)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2Campaign)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2Campaign)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2Campaign)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2Campaign)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2Campaign)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2Campaign)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a Connectcampaignsv2Campaign resource upon running "cdktf plan <stack-name>".
func Connectcampaignsv2Campaign_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateConnectcampaignsv2Campaign_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.connectcampaignsv2Campaign.Connectcampaignsv2Campaign",
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
func Connectcampaignsv2Campaign_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConnectcampaignsv2Campaign_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.connectcampaignsv2Campaign.Connectcampaignsv2Campaign",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Connectcampaignsv2Campaign_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConnectcampaignsv2Campaign_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.connectcampaignsv2Campaign.Connectcampaignsv2Campaign",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Connectcampaignsv2Campaign_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConnectcampaignsv2Campaign_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.connectcampaignsv2Campaign.Connectcampaignsv2Campaign",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Connectcampaignsv2Campaign_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.connectcampaignsv2Campaign.Connectcampaignsv2Campaign",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_Connectcampaignsv2Campaign) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_Connectcampaignsv2Campaign) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_Connectcampaignsv2Campaign) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_Connectcampaignsv2Campaign) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_Connectcampaignsv2Campaign) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_Connectcampaignsv2Campaign) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_Connectcampaignsv2Campaign) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_Connectcampaignsv2Campaign) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_Connectcampaignsv2Campaign) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_Connectcampaignsv2Campaign) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) PutChannelSubtypeConfig(value *Connectcampaignsv2CampaignChannelSubtypeConfig) {
	if err := c.validatePutChannelSubtypeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putChannelSubtypeConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) PutCommunicationLimitsOverride(value *Connectcampaignsv2CampaignCommunicationLimitsOverride) {
	if err := c.validatePutCommunicationLimitsOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCommunicationLimitsOverride",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) PutCommunicationTimeConfig(value *Connectcampaignsv2CampaignCommunicationTimeConfig) {
	if err := c.validatePutCommunicationTimeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCommunicationTimeConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) PutSchedule(value *Connectcampaignsv2CampaignSchedule) {
	if err := c.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSchedule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) PutSource(value *Connectcampaignsv2CampaignSource) {
	if err := c.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSource",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) PutTags(value interface{}) {
	if err := c.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTags",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) ResetCommunicationLimitsOverride() {
	_jsii_.InvokeVoid(
		c,
		"resetCommunicationLimitsOverride",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) ResetCommunicationTimeConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetCommunicationTimeConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) ResetConnectCampaignFlowArn() {
	_jsii_.InvokeVoid(
		c,
		"resetConnectCampaignFlowArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) ResetSchedule() {
	_jsii_.InvokeVoid(
		c,
		"resetSchedule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) ResetSource() {
	_jsii_.InvokeVoid(
		c,
		"resetSource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2Campaign) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

