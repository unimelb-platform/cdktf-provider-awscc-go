package networkmanagersitetositevpnattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/networkmanagersitetositevpnattachment/internal"
)

type NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference interface {
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
	SegmentName() *string
	SetSegmentName(val *string)
	SegmentNameInput() *string
	Tags() NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeTagsList
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
	ResetSegmentName()
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

// The jsii proxy struct for NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference
type jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) AttachmentPolicyRuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attachmentPolicyRuleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) AttachmentPolicyRuleNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attachmentPolicyRuleNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) SegmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) SegmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) Tags() NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeTagsList {
	var returns NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference{}

	_jsii_.Create(
		"awscc.networkmanagerSiteToSiteVpnAttachment.NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference_Override(n NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.networkmanagerSiteToSiteVpnAttachment.NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference)SetAttachmentPolicyRuleNumber(val *float64) {
	if err := j.validateSetAttachmentPolicyRuleNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachmentPolicyRuleNumber",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference)SetSegmentName(val *string) {
	if err := j.validateSetSegmentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentName",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) PutTags(value interface{}) {
	if err := n.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTags",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) ResetAttachmentPolicyRuleNumber() {
	_jsii_.InvokeVoid(
		n,
		"resetAttachmentPolicyRuleNumber",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) ResetSegmentName() {
	_jsii_.InvokeVoid(
		n,
		"resetSegmentName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		n,
		"resetTags",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChangeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

