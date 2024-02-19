package mediaconnectflowoutput

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediaconnectflowoutput/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output awscc_mediaconnect_flow_output}.
type MediaconnectFlowOutput interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CidrAllowList() *[]*string
	SetCidrAllowList(val *[]*string)
	CidrAllowListInput() *[]*string
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
	Destination() *string
	SetDestination(val *string)
	DestinationInput() *string
	Encryption() MediaconnectFlowOutputEncryptionOutputReference
	EncryptionInput() interface{}
	FlowArn() *string
	SetFlowArn(val *string)
	FlowArnInput() *string
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
	MaxLatency() *float64
	SetMaxLatency(val *float64)
	MaxLatencyInput() *float64
	MinLatency() *float64
	SetMinLatency(val *float64)
	MinLatencyInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutputArn() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	RemoteId() *string
	SetRemoteId(val *string)
	RemoteIdInput() *string
	SmoothingLatency() *float64
	SetSmoothingLatency(val *float64)
	SmoothingLatencyInput() *float64
	StreamId() *string
	SetStreamId(val *string)
	StreamIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VpcInterfaceAttachment() MediaconnectFlowOutputVpcInterfaceAttachmentOutputReference
	VpcInterfaceAttachmentInput() interface{}
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
	PutEncryption(value *MediaconnectFlowOutputEncryption)
	PutVpcInterfaceAttachment(value *MediaconnectFlowOutputVpcInterfaceAttachment)
	ResetCidrAllowList()
	ResetDescription()
	ResetDestination()
	ResetEncryption()
	ResetMaxLatency()
	ResetMinLatency()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetRemoteId()
	ResetSmoothingLatency()
	ResetStreamId()
	ResetVpcInterfaceAttachment()
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

// The jsii proxy struct for MediaconnectFlowOutput
type jsiiProxy_MediaconnectFlowOutput struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MediaconnectFlowOutput) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) CidrAllowList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrAllowList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) CidrAllowListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrAllowListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) DestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) Encryption() MediaconnectFlowOutputEncryptionOutputReference {
	var returns MediaconnectFlowOutputEncryptionOutputReference
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) EncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) FlowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) FlowArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) MaxLatency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) MaxLatencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) MinLatency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) MinLatencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) OutputArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) RemoteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) RemoteIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) SmoothingLatency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"smoothingLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) SmoothingLatencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"smoothingLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) StreamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) StreamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) VpcInterfaceAttachment() MediaconnectFlowOutputVpcInterfaceAttachmentOutputReference {
	var returns MediaconnectFlowOutputVpcInterfaceAttachmentOutputReference
	_jsii_.Get(
		j,
		"vpcInterfaceAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutput) VpcInterfaceAttachmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcInterfaceAttachmentInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output awscc_mediaconnect_flow_output} Resource.
func NewMediaconnectFlowOutput(scope constructs.Construct, id *string, config *MediaconnectFlowOutputConfig) MediaconnectFlowOutput {
	_init_.Initialize()

	if err := validateNewMediaconnectFlowOutputParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectFlowOutput{}

	_jsii_.Create(
		"awscc.mediaconnectFlowOutput.MediaconnectFlowOutput",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output awscc_mediaconnect_flow_output} Resource.
func NewMediaconnectFlowOutput_Override(m MediaconnectFlowOutput, scope constructs.Construct, id *string, config *MediaconnectFlowOutputConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediaconnectFlowOutput.MediaconnectFlowOutput",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetCidrAllowList(val *[]*string) {
	if err := j.validateSetCidrAllowListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrAllowList",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetDestination(val *string) {
	if err := j.validateSetDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetFlowArn(val *string) {
	if err := j.validateSetFlowArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowArn",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetMaxLatency(val *float64) {
	if err := j.validateSetMaxLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLatency",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetMinLatency(val *float64) {
	if err := j.validateSetMinLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minLatency",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetRemoteId(val *string) {
	if err := j.validateSetRemoteIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteId",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetSmoothingLatency(val *float64) {
	if err := j.validateSetSmoothingLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smoothingLatency",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutput)SetStreamId(val *string) {
	if err := j.validateSetStreamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamId",
		val,
	)
}

// Generates CDKTF code for importing a MediaconnectFlowOutput resource upon running "cdktf plan <stack-name>".
func MediaconnectFlowOutput_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMediaconnectFlowOutput_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.mediaconnectFlowOutput.MediaconnectFlowOutput",
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
func MediaconnectFlowOutput_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaconnectFlowOutput_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mediaconnectFlowOutput.MediaconnectFlowOutput",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MediaconnectFlowOutput_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaconnectFlowOutput_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mediaconnectFlowOutput.MediaconnectFlowOutput",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MediaconnectFlowOutput_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaconnectFlowOutput_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mediaconnectFlowOutput.MediaconnectFlowOutput",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MediaconnectFlowOutput_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.mediaconnectFlowOutput.MediaconnectFlowOutput",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutput) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectFlowOutput) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaconnectFlowOutput) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectFlowOutput) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectFlowOutput) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectFlowOutput) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectFlowOutput) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectFlowOutput) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectFlowOutput) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectFlowOutput) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutput) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaconnectFlowOutput) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) PutEncryption(value *MediaconnectFlowOutputEncryption) {
	if err := m.validatePutEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncryption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) PutVpcInterfaceAttachment(value *MediaconnectFlowOutputVpcInterfaceAttachment) {
	if err := m.validatePutVpcInterfaceAttachmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putVpcInterfaceAttachment",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) ResetCidrAllowList() {
	_jsii_.InvokeVoid(
		m,
		"resetCidrAllowList",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) ResetDestination() {
	_jsii_.InvokeVoid(
		m,
		"resetDestination",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) ResetEncryption() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryption",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) ResetMaxLatency() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxLatency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) ResetMinLatency() {
	_jsii_.InvokeVoid(
		m,
		"resetMinLatency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) ResetName() {
	_jsii_.InvokeVoid(
		m,
		"resetName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) ResetPort() {
	_jsii_.InvokeVoid(
		m,
		"resetPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) ResetRemoteId() {
	_jsii_.InvokeVoid(
		m,
		"resetRemoteId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) ResetSmoothingLatency() {
	_jsii_.InvokeVoid(
		m,
		"resetSmoothingLatency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) ResetStreamId() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) ResetVpcInterfaceAttachment() {
	_jsii_.InvokeVoid(
		m,
		"resetVpcInterfaceAttachment",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowOutput) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutput) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutput) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutput) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutput) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutput) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

