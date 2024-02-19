package mediatailorchannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediatailorchannel/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediatailor_channel awscc_mediatailor_channel}.
type MediatailorChannel interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChannelName() *string
	SetChannelName(val *string)
	ChannelNameInput() *string
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
	FillerSlate() MediatailorChannelFillerSlateOutputReference
	FillerSlateInput() interface{}
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
	LogConfiguration() MediatailorChannelLogConfigurationOutputReference
	LogConfigurationInput() interface{}
	// The tree node.
	Node() constructs.Node
	Outputs() MediatailorChannelOutputsList
	OutputsInput() interface{}
	PlaybackMode() *string
	SetPlaybackMode(val *string)
	PlaybackModeInput() *string
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
	Tags() MediatailorChannelTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tier() *string
	SetTier(val *string)
	TierInput() *string
	TimeShiftConfiguration() MediatailorChannelTimeShiftConfigurationOutputReference
	TimeShiftConfigurationInput() interface{}
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
	PutFillerSlate(value *MediatailorChannelFillerSlate)
	PutLogConfiguration(value *MediatailorChannelLogConfiguration)
	PutOutputs(value interface{})
	PutTags(value interface{})
	PutTimeShiftConfiguration(value *MediatailorChannelTimeShiftConfiguration)
	ResetFillerSlate()
	ResetLogConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTier()
	ResetTimeShiftConfiguration()
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

// The jsii proxy struct for MediatailorChannel
type jsiiProxy_MediatailorChannel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MediatailorChannel) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) ChannelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) ChannelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) FillerSlate() MediatailorChannelFillerSlateOutputReference {
	var returns MediatailorChannelFillerSlateOutputReference
	_jsii_.Get(
		j,
		"fillerSlate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) FillerSlateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fillerSlateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) LogConfiguration() MediatailorChannelLogConfigurationOutputReference {
	var returns MediatailorChannelLogConfigurationOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) LogConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) Outputs() MediatailorChannelOutputsList {
	var returns MediatailorChannelOutputsList
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) OutputsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) PlaybackMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"playbackMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) PlaybackModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"playbackModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) Tags() MediatailorChannelTagsList {
	var returns MediatailorChannelTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) TimeShiftConfiguration() MediatailorChannelTimeShiftConfigurationOutputReference {
	var returns MediatailorChannelTimeShiftConfigurationOutputReference
	_jsii_.Get(
		j,
		"timeShiftConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorChannel) TimeShiftConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeShiftConfigurationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediatailor_channel awscc_mediatailor_channel} Resource.
func NewMediatailorChannel(scope constructs.Construct, id *string, config *MediatailorChannelConfig) MediatailorChannel {
	_init_.Initialize()

	if err := validateNewMediatailorChannelParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediatailorChannel{}

	_jsii_.Create(
		"awscc.mediatailorChannel.MediatailorChannel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediatailor_channel awscc_mediatailor_channel} Resource.
func NewMediatailorChannel_Override(m MediatailorChannel, scope constructs.Construct, id *string, config *MediatailorChannelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediatailorChannel.MediatailorChannel",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MediatailorChannel)SetChannelName(val *string) {
	if err := j.validateSetChannelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelName",
		val,
	)
}

func (j *jsiiProxy_MediatailorChannel)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MediatailorChannel)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MediatailorChannel)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MediatailorChannel)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MediatailorChannel)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MediatailorChannel)SetPlaybackMode(val *string) {
	if err := j.validateSetPlaybackModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"playbackMode",
		val,
	)
}

func (j *jsiiProxy_MediatailorChannel)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MediatailorChannel)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MediatailorChannel)SetTier(val *string) {
	if err := j.validateSetTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

// Generates CDKTF code for importing a MediatailorChannel resource upon running "cdktf plan <stack-name>".
func MediatailorChannel_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMediatailorChannel_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.mediatailorChannel.MediatailorChannel",
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
func MediatailorChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediatailorChannel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mediatailorChannel.MediatailorChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MediatailorChannel_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediatailorChannel_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mediatailorChannel.MediatailorChannel",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MediatailorChannel_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediatailorChannel_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mediatailorChannel.MediatailorChannel",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MediatailorChannel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.mediatailorChannel.MediatailorChannel",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MediatailorChannel) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MediatailorChannel) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MediatailorChannel) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorChannel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorChannel) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorChannel) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorChannel) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorChannel) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorChannel) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorChannel) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorChannel) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorChannel) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorChannel) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MediatailorChannel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorChannel) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MediatailorChannel) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MediatailorChannel) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MediatailorChannel) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MediatailorChannel) PutFillerSlate(value *MediatailorChannelFillerSlate) {
	if err := m.validatePutFillerSlateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFillerSlate",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorChannel) PutLogConfiguration(value *MediatailorChannelLogConfiguration) {
	if err := m.validatePutLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLogConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorChannel) PutOutputs(value interface{}) {
	if err := m.validatePutOutputsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putOutputs",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorChannel) PutTags(value interface{}) {
	if err := m.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTags",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorChannel) PutTimeShiftConfiguration(value *MediatailorChannelTimeShiftConfiguration) {
	if err := m.validatePutTimeShiftConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeShiftConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorChannel) ResetFillerSlate() {
	_jsii_.InvokeVoid(
		m,
		"resetFillerSlate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorChannel) ResetLogConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetLogConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorChannel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorChannel) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorChannel) ResetTier() {
	_jsii_.InvokeVoid(
		m,
		"resetTier",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorChannel) ResetTimeShiftConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeShiftConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorChannel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorChannel) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorChannel) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorChannel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorChannel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

