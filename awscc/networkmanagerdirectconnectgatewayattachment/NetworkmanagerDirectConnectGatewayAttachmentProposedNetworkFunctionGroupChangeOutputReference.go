package networkmanagerdirectconnectgatewayattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/networkmanagerdirectconnectgatewayattachment/internal"
)

type NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference interface {
	cdktf.ComplexObject
	AttachmentPolicyRuleNumber() *float64
	SetAttachmentPolicyRuleNumber(val *float64)
	AttachmentPolicyRuleNumberInput() *float64
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NetworkFunctionGroupName() *string
	SetNetworkFunctionGroupName(val *string)
	NetworkFunctionGroupNameInput() *string
	Tags() NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutTags(value interface{})
	ResetAttachmentPolicyRuleNumber()
	ResetNetworkFunctionGroupName()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference
type jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) AttachmentPolicyRuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attachmentPolicyRuleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) AttachmentPolicyRuleNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attachmentPolicyRuleNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) NetworkFunctionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkFunctionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) NetworkFunctionGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkFunctionGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) Tags() NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeTagsList {
	var returns NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference{}

	_jsii_.Create(
		"awscc.networkmanagerDirectConnectGatewayAttachment.NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference_Override(n NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.networkmanagerDirectConnectGatewayAttachment.NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetAttachmentPolicyRuleNumber(val *float64) {
	if err := j.validateSetAttachmentPolicyRuleNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachmentPolicyRuleNumber",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetNetworkFunctionGroupName(val *string) {
	if err := j.validateSetNetworkFunctionGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkFunctionGroupName",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) PutTags(value interface{}) {
	if err := n.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTags",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) ResetAttachmentPolicyRuleNumber() {
	_jsii_.InvokeVoid(
		n,
		"resetAttachmentPolicyRuleNumber",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) ResetNetworkFunctionGroupName() {
	_jsii_.InvokeVoid(
		n,
		"resetNetworkFunctionGroupName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		n,
		"resetTags",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

