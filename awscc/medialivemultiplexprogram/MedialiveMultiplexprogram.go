package medialivemultiplexprogram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/medialivemultiplexprogram/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplexprogram awscc_medialive_multiplexprogram}.
type MedialiveMultiplexprogram interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChannelId() *string
	SetChannelId(val *string)
	ChannelIdInput() *string
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
	MultiplexId() *string
	SetMultiplexId(val *string)
	MultiplexIdInput() *string
	MultiplexProgramSettings() MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference
	MultiplexProgramSettingsInput() interface{}
	// The tree node.
	Node() constructs.Node
	PacketIdentifiersMap() MedialiveMultiplexprogramPacketIdentifiersMapOutputReference
	PacketIdentifiersMapInput() interface{}
	PipelineDetails() MedialiveMultiplexprogramPipelineDetailsList
	PipelineDetailsInput() interface{}
	PreferredChannelPipeline() *string
	SetPreferredChannelPipeline(val *string)
	PreferredChannelPipelineInput() *string
	ProgramName() *string
	SetProgramName(val *string)
	ProgramNameInput() *string
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
	PutMultiplexProgramSettings(value *MedialiveMultiplexprogramMultiplexProgramSettings)
	PutPacketIdentifiersMap(value *MedialiveMultiplexprogramPacketIdentifiersMap)
	PutPipelineDetails(value interface{})
	ResetChannelId()
	ResetMultiplexId()
	ResetMultiplexProgramSettings()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPacketIdentifiersMap()
	ResetPipelineDetails()
	ResetPreferredChannelPipeline()
	ResetProgramName()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MedialiveMultiplexprogram
type jsiiProxy_MedialiveMultiplexprogram struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MedialiveMultiplexprogram) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) ChannelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) ChannelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) MultiplexId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiplexId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) MultiplexIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiplexIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) MultiplexProgramSettings() MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference {
	var returns MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference
	_jsii_.Get(
		j,
		"multiplexProgramSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) MultiplexProgramSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiplexProgramSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) PacketIdentifiersMap() MedialiveMultiplexprogramPacketIdentifiersMapOutputReference {
	var returns MedialiveMultiplexprogramPacketIdentifiersMapOutputReference
	_jsii_.Get(
		j,
		"packetIdentifiersMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) PacketIdentifiersMapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"packetIdentifiersMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) PipelineDetails() MedialiveMultiplexprogramPipelineDetailsList {
	var returns MedialiveMultiplexprogramPipelineDetailsList
	_jsii_.Get(
		j,
		"pipelineDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) PipelineDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) PreferredChannelPipeline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredChannelPipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) PreferredChannelPipelineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredChannelPipelineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) ProgramName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) ProgramNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogram) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplexprogram awscc_medialive_multiplexprogram} Resource.
func NewMedialiveMultiplexprogram(scope constructs.Construct, id *string, config *MedialiveMultiplexprogramConfig) MedialiveMultiplexprogram {
	_init_.Initialize()

	if err := validateNewMedialiveMultiplexprogramParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveMultiplexprogram{}

	_jsii_.Create(
		"awscc.medialiveMultiplexprogram.MedialiveMultiplexprogram",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplexprogram awscc_medialive_multiplexprogram} Resource.
func NewMedialiveMultiplexprogram_Override(m MedialiveMultiplexprogram, scope constructs.Construct, id *string, config *MedialiveMultiplexprogramConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.medialiveMultiplexprogram.MedialiveMultiplexprogram",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogram)SetChannelId(val *string) {
	if err := j.validateSetChannelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelId",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogram)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogram)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogram)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogram)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogram)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogram)SetMultiplexId(val *string) {
	if err := j.validateSetMultiplexIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiplexId",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogram)SetPreferredChannelPipeline(val *string) {
	if err := j.validateSetPreferredChannelPipelineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredChannelPipeline",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogram)SetProgramName(val *string) {
	if err := j.validateSetProgramNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programName",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogram)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogram)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a MedialiveMultiplexprogram resource upon running "cdktf plan <stack-name>".
func MedialiveMultiplexprogram_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMedialiveMultiplexprogram_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.medialiveMultiplexprogram.MedialiveMultiplexprogram",
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
func MedialiveMultiplexprogram_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMedialiveMultiplexprogram_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.medialiveMultiplexprogram.MedialiveMultiplexprogram",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MedialiveMultiplexprogram_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMedialiveMultiplexprogram_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.medialiveMultiplexprogram.MedialiveMultiplexprogram",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MedialiveMultiplexprogram_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMedialiveMultiplexprogram_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.medialiveMultiplexprogram.MedialiveMultiplexprogram",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MedialiveMultiplexprogram_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.medialiveMultiplexprogram.MedialiveMultiplexprogram",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogram) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogram) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogram) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveMultiplexprogram) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveMultiplexprogram) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveMultiplexprogram) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveMultiplexprogram) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveMultiplexprogram) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveMultiplexprogram) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveMultiplexprogram) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveMultiplexprogram) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveMultiplexprogram) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogram) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveMultiplexprogram) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogram) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogram) PutMultiplexProgramSettings(value *MedialiveMultiplexprogramMultiplexProgramSettings) {
	if err := m.validatePutMultiplexProgramSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMultiplexProgramSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogram) PutPacketIdentifiersMap(value *MedialiveMultiplexprogramPacketIdentifiersMap) {
	if err := m.validatePutPacketIdentifiersMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPacketIdentifiersMap",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogram) PutPipelineDetails(value interface{}) {
	if err := m.validatePutPipelineDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPipelineDetails",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogram) ResetChannelId() {
	_jsii_.InvokeVoid(
		m,
		"resetChannelId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogram) ResetMultiplexId() {
	_jsii_.InvokeVoid(
		m,
		"resetMultiplexId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogram) ResetMultiplexProgramSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetMultiplexProgramSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogram) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogram) ResetPacketIdentifiersMap() {
	_jsii_.InvokeVoid(
		m,
		"resetPacketIdentifiersMap",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogram) ResetPipelineDetails() {
	_jsii_.InvokeVoid(
		m,
		"resetPipelineDetails",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogram) ResetPreferredChannelPipeline() {
	_jsii_.InvokeVoid(
		m,
		"resetPreferredChannelPipeline",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogram) ResetProgramName() {
	_jsii_.InvokeVoid(
		m,
		"resetProgramName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogram) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogram) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogram) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogram) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

