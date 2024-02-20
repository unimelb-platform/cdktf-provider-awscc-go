package mediaconnectflowsource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediaconnectflowsource/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_source awscc_mediaconnect_flow_source}.
type MediaconnectFlowSourceA interface {
	cdktf.TerraformResource
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
	Decryption() MediaconnectFlowSourceDecryptionAOutputReference
	DecryptionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EntitlementArn() *string
	SetEntitlementArn(val *string)
	EntitlementArnInput() *string
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
	GatewayBridgeSource() MediaconnectFlowSourceGatewayBridgeSourceAOutputReference
	GatewayBridgeSourceInput() interface{}
	Id() *string
	IngestIp() *string
	IngestPort() *float64
	SetIngestPort(val *float64)
	IngestPortInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxBitrate() *float64
	SetMaxBitrate(val *float64)
	MaxBitrateInput() *float64
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
	SenderControlPort() *float64
	SetSenderControlPort(val *float64)
	SenderControlPortInput() *float64
	SenderIpAddress() *string
	SetSenderIpAddress(val *string)
	SenderIpAddressInput() *string
	SourceArn() *string
	SourceIngestPort() *string
	SourceListenerAddress() *string
	SetSourceListenerAddress(val *string)
	SourceListenerAddressInput() *string
	SourceListenerPort() *float64
	SetSourceListenerPort(val *float64)
	SourceListenerPortInput() *float64
	StreamId() *string
	SetStreamId(val *string)
	StreamIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VpcInterfaceName() *string
	SetVpcInterfaceName(val *string)
	VpcInterfaceNameInput() *string
	WhitelistCidr() *string
	SetWhitelistCidr(val *string)
	WhitelistCidrInput() *string
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
	PutDecryption(value *MediaconnectFlowSourceDecryptionA)
	PutGatewayBridgeSource(value *MediaconnectFlowSourceGatewayBridgeSourceA)
	ResetDecryption()
	ResetEntitlementArn()
	ResetFlowArn()
	ResetGatewayBridgeSource()
	ResetIngestPort()
	ResetMaxBitrate()
	ResetMaxLatency()
	ResetMinLatency()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProtocol()
	ResetSenderControlPort()
	ResetSenderIpAddress()
	ResetSourceListenerAddress()
	ResetSourceListenerPort()
	ResetStreamId()
	ResetVpcInterfaceName()
	ResetWhitelistCidr()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MediaconnectFlowSourceA
type jsiiProxy_MediaconnectFlowSourceA struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MediaconnectFlowSourceA) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) Decryption() MediaconnectFlowSourceDecryptionAOutputReference {
	var returns MediaconnectFlowSourceDecryptionAOutputReference
	_jsii_.Get(
		j,
		"decryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) DecryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) EntitlementArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entitlementArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) EntitlementArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entitlementArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) FlowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) FlowArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) GatewayBridgeSource() MediaconnectFlowSourceGatewayBridgeSourceAOutputReference {
	var returns MediaconnectFlowSourceGatewayBridgeSourceAOutputReference
	_jsii_.Get(
		j,
		"gatewayBridgeSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) GatewayBridgeSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewayBridgeSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) IngestIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingestIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) IngestPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingestPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) IngestPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingestPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) MaxBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) MaxBitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) MaxLatency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) MaxLatencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) MinLatency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) MinLatencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) SenderControlPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"senderControlPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) SenderControlPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"senderControlPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) SenderIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"senderIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) SenderIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"senderIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) SourceIngestPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIngestPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) SourceListenerAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceListenerAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) SourceListenerAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceListenerAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) SourceListenerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceListenerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) SourceListenerPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceListenerPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) StreamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) StreamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) VpcInterfaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcInterfaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) VpcInterfaceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcInterfaceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) WhitelistCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whitelistCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceA) WhitelistCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whitelistCidrInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_source awscc_mediaconnect_flow_source} Resource.
func NewMediaconnectFlowSourceA(scope constructs.Construct, id *string, config *MediaconnectFlowSourceAConfig) MediaconnectFlowSourceA {
	_init_.Initialize()

	if err := validateNewMediaconnectFlowSourceAParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectFlowSourceA{}

	_jsii_.Create(
		"awscc.mediaconnectFlowSource.MediaconnectFlowSourceA",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_source awscc_mediaconnect_flow_source} Resource.
func NewMediaconnectFlowSourceA_Override(m MediaconnectFlowSourceA, scope constructs.Construct, id *string, config *MediaconnectFlowSourceAConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediaconnectFlowSource.MediaconnectFlowSourceA",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetEntitlementArn(val *string) {
	if err := j.validateSetEntitlementArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entitlementArn",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetFlowArn(val *string) {
	if err := j.validateSetFlowArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowArn",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetIngestPort(val *float64) {
	if err := j.validateSetIngestPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingestPort",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetMaxBitrate(val *float64) {
	if err := j.validateSetMaxBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBitrate",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetMaxLatency(val *float64) {
	if err := j.validateSetMaxLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLatency",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetMinLatency(val *float64) {
	if err := j.validateSetMinLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minLatency",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetSenderControlPort(val *float64) {
	if err := j.validateSetSenderControlPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"senderControlPort",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetSenderIpAddress(val *string) {
	if err := j.validateSetSenderIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"senderIpAddress",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetSourceListenerAddress(val *string) {
	if err := j.validateSetSourceListenerAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceListenerAddress",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetSourceListenerPort(val *float64) {
	if err := j.validateSetSourceListenerPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceListenerPort",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetStreamId(val *string) {
	if err := j.validateSetStreamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamId",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetVpcInterfaceName(val *string) {
	if err := j.validateSetVpcInterfaceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcInterfaceName",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceA)SetWhitelistCidr(val *string) {
	if err := j.validateSetWhitelistCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whitelistCidr",
		val,
	)
}

// Generates CDKTF code for importing a MediaconnectFlowSourceA resource upon running "cdktf plan <stack-name>".
func MediaconnectFlowSourceA_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMediaconnectFlowSourceA_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.mediaconnectFlowSource.MediaconnectFlowSourceA",
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
func MediaconnectFlowSourceA_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaconnectFlowSourceA_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mediaconnectFlowSource.MediaconnectFlowSourceA",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MediaconnectFlowSourceA_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaconnectFlowSourceA_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mediaconnectFlowSource.MediaconnectFlowSourceA",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MediaconnectFlowSourceA_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaconnectFlowSourceA_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mediaconnectFlowSource.MediaconnectFlowSourceA",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MediaconnectFlowSourceA_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.mediaconnectFlowSource.MediaconnectFlowSourceA",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceA) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectFlowSourceA) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaconnectFlowSourceA) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectFlowSourceA) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectFlowSourceA) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectFlowSourceA) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectFlowSourceA) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectFlowSourceA) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectFlowSourceA) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectFlowSourceA) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaconnectFlowSourceA) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) PutDecryption(value *MediaconnectFlowSourceDecryptionA) {
	if err := m.validatePutDecryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDecryption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) PutGatewayBridgeSource(value *MediaconnectFlowSourceGatewayBridgeSourceA) {
	if err := m.validatePutGatewayBridgeSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putGatewayBridgeSource",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetDecryption() {
	_jsii_.InvokeVoid(
		m,
		"resetDecryption",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetEntitlementArn() {
	_jsii_.InvokeVoid(
		m,
		"resetEntitlementArn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetFlowArn() {
	_jsii_.InvokeVoid(
		m,
		"resetFlowArn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetGatewayBridgeSource() {
	_jsii_.InvokeVoid(
		m,
		"resetGatewayBridgeSource",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetIngestPort() {
	_jsii_.InvokeVoid(
		m,
		"resetIngestPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetMaxBitrate() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBitrate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetMaxLatency() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxLatency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetMinLatency() {
	_jsii_.InvokeVoid(
		m,
		"resetMinLatency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetProtocol() {
	_jsii_.InvokeVoid(
		m,
		"resetProtocol",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetSenderControlPort() {
	_jsii_.InvokeVoid(
		m,
		"resetSenderControlPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetSenderIpAddress() {
	_jsii_.InvokeVoid(
		m,
		"resetSenderIpAddress",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetSourceListenerAddress() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceListenerAddress",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetSourceListenerPort() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceListenerPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetStreamId() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetVpcInterfaceName() {
	_jsii_.InvokeVoid(
		m,
		"resetVpcInterfaceName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ResetWhitelistCidr() {
	_jsii_.InvokeVoid(
		m,
		"resetWhitelistCidr",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceA) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceA) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

