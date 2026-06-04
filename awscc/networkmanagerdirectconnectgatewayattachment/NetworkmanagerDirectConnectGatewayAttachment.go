package networkmanagerdirectconnectgatewayattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/networkmanagerdirectconnectgatewayattachment/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_direct_connect_gateway_attachment awscc_networkmanager_direct_connect_gateway_attachment}.
type NetworkmanagerDirectConnectGatewayAttachment interface {
	cdktf.TerraformResource
	AttachmentId() *string
	AttachmentPolicyRuleNumber() *float64
	AttachmentType() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CoreNetworkArn() *string
	CoreNetworkId() *string
	SetCoreNetworkId(val *string)
	CoreNetworkIdInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DirectConnectGatewayArn() *string
	SetDirectConnectGatewayArn(val *string)
	DirectConnectGatewayArnInput() *string
	EdgeLocations() *[]*string
	SetEdgeLocations(val *[]*string)
	EdgeLocationsInput() *[]*string
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
	NetworkFunctionGroupName() *string
	// The tree node.
	Node() constructs.Node
	OwnerAccountId() *string
	ProposedNetworkFunctionGroupChange() NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference
	ProposedNetworkFunctionGroupChangeInput() interface{}
	ProposedSegmentChange() NetworkmanagerDirectConnectGatewayAttachmentProposedSegmentChangeOutputReference
	ProposedSegmentChangeInput() interface{}
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
	ResourceArn() *string
	SegmentName() *string
	State() *string
	Tags() NetworkmanagerDirectConnectGatewayAttachmentTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatedAt() *string
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
	PutProposedNetworkFunctionGroupChange(value *NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChange)
	PutProposedSegmentChange(value *NetworkmanagerDirectConnectGatewayAttachmentProposedSegmentChange)
	PutTags(value interface{})
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProposedNetworkFunctionGroupChange()
	ResetProposedSegmentChange()
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

// The jsii proxy struct for NetworkmanagerDirectConnectGatewayAttachment
type jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) AttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) AttachmentPolicyRuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attachmentPolicyRuleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) AttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) CoreNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) CoreNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) CoreNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) DirectConnectGatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directConnectGatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) DirectConnectGatewayArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directConnectGatewayArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) EdgeLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"edgeLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) EdgeLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"edgeLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) NetworkFunctionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkFunctionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) OwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) ProposedNetworkFunctionGroupChange() NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference {
	var returns NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference
	_jsii_.Get(
		j,
		"proposedNetworkFunctionGroupChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) ProposedNetworkFunctionGroupChangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proposedNetworkFunctionGroupChangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) ProposedSegmentChange() NetworkmanagerDirectConnectGatewayAttachmentProposedSegmentChangeOutputReference {
	var returns NetworkmanagerDirectConnectGatewayAttachmentProposedSegmentChangeOutputReference
	_jsii_.Get(
		j,
		"proposedSegmentChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) ProposedSegmentChangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proposedSegmentChangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) SegmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) Tags() NetworkmanagerDirectConnectGatewayAttachmentTagsList {
	var returns NetworkmanagerDirectConnectGatewayAttachmentTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_direct_connect_gateway_attachment awscc_networkmanager_direct_connect_gateway_attachment} Resource.
func NewNetworkmanagerDirectConnectGatewayAttachment(scope constructs.Construct, id *string, config *NetworkmanagerDirectConnectGatewayAttachmentConfig) NetworkmanagerDirectConnectGatewayAttachment {
	_init_.Initialize()

	if err := validateNewNetworkmanagerDirectConnectGatewayAttachmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment{}

	_jsii_.Create(
		"awscc.networkmanagerDirectConnectGatewayAttachment.NetworkmanagerDirectConnectGatewayAttachment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_direct_connect_gateway_attachment awscc_networkmanager_direct_connect_gateway_attachment} Resource.
func NewNetworkmanagerDirectConnectGatewayAttachment_Override(n NetworkmanagerDirectConnectGatewayAttachment, scope constructs.Construct, id *string, config *NetworkmanagerDirectConnectGatewayAttachmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.networkmanagerDirectConnectGatewayAttachment.NetworkmanagerDirectConnectGatewayAttachment",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment)SetCoreNetworkId(val *string) {
	if err := j.validateSetCoreNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreNetworkId",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment)SetDirectConnectGatewayArn(val *string) {
	if err := j.validateSetDirectConnectGatewayArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directConnectGatewayArn",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment)SetEdgeLocations(val *[]*string) {
	if err := j.validateSetEdgeLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeLocations",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a NetworkmanagerDirectConnectGatewayAttachment resource upon running "cdktf plan <stack-name>".
func NetworkmanagerDirectConnectGatewayAttachment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetworkmanagerDirectConnectGatewayAttachment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.networkmanagerDirectConnectGatewayAttachment.NetworkmanagerDirectConnectGatewayAttachment",
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
func NetworkmanagerDirectConnectGatewayAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkmanagerDirectConnectGatewayAttachment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.networkmanagerDirectConnectGatewayAttachment.NetworkmanagerDirectConnectGatewayAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkmanagerDirectConnectGatewayAttachment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkmanagerDirectConnectGatewayAttachment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.networkmanagerDirectConnectGatewayAttachment.NetworkmanagerDirectConnectGatewayAttachment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkmanagerDirectConnectGatewayAttachment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkmanagerDirectConnectGatewayAttachment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.networkmanagerDirectConnectGatewayAttachment.NetworkmanagerDirectConnectGatewayAttachment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetworkmanagerDirectConnectGatewayAttachment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.networkmanagerDirectConnectGatewayAttachment.NetworkmanagerDirectConnectGatewayAttachment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) PutProposedNetworkFunctionGroupChange(value *NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChange) {
	if err := n.validatePutProposedNetworkFunctionGroupChangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putProposedNetworkFunctionGroupChange",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) PutProposedSegmentChange(value *NetworkmanagerDirectConnectGatewayAttachmentProposedSegmentChange) {
	if err := n.validatePutProposedSegmentChangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putProposedSegmentChange",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) PutTags(value interface{}) {
	if err := n.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTags",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) ResetProposedNetworkFunctionGroupChange() {
	_jsii_.InvokeVoid(
		n,
		"resetProposedNetworkFunctionGroupChange",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) ResetProposedSegmentChange() {
	_jsii_.InvokeVoid(
		n,
		"resetProposedSegmentChange",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) ResetTags() {
	_jsii_.InvokeVoid(
		n,
		"resetTags",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

