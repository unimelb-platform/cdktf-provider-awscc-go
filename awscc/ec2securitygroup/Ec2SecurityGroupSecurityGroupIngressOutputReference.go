package ec2securitygroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2securitygroup/internal"
)

type Ec2SecurityGroupSecurityGroupIngressOutputReference interface {
	cdktf.ComplexObject
	CidrIp() *string
	SetCidrIp(val *string)
	CidrIpInput() *string
	CidrIpv6() *string
	SetCidrIpv6(val *string)
	CidrIpv6Input() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	FromPort() *float64
	SetFromPort(val *float64)
	FromPortInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpProtocol() *string
	SetIpProtocol(val *string)
	IpProtocolInput() *string
	SourcePrefixListId() *string
	SetSourcePrefixListId(val *string)
	SourcePrefixListIdInput() *string
	SourceSecurityGroupId() *string
	SetSourceSecurityGroupId(val *string)
	SourceSecurityGroupIdInput() *string
	SourceSecurityGroupName() *string
	SetSourceSecurityGroupName(val *string)
	SourceSecurityGroupNameInput() *string
	SourceSecurityGroupOwnerId() *string
	SetSourceSecurityGroupOwnerId(val *string)
	SourceSecurityGroupOwnerIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ToPort() *float64
	SetToPort(val *float64)
	ToPortInput() *float64
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
	ResetCidrIp()
	ResetCidrIpv6()
	ResetDescription()
	ResetFromPort()
	ResetSourcePrefixListId()
	ResetSourceSecurityGroupId()
	ResetSourceSecurityGroupName()
	ResetSourceSecurityGroupOwnerId()
	ResetToPort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2SecurityGroupSecurityGroupIngressOutputReference
type jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) CidrIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) CidrIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) CidrIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) CidrIpv6Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) FromPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) FromPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) IpProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) IpProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) SourcePrefixListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePrefixListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) SourcePrefixListIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePrefixListIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) SourceSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) SourceSecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSecurityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) SourceSecurityGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSecurityGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) SourceSecurityGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSecurityGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) SourceSecurityGroupOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSecurityGroupOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) SourceSecurityGroupOwnerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSecurityGroupOwnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) ToPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) ToPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toPortInput",
		&returns,
	)
	return returns
}


func NewEc2SecurityGroupSecurityGroupIngressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2SecurityGroupSecurityGroupIngressOutputReference {
	_init_.Initialize()

	if err := validateNewEc2SecurityGroupSecurityGroupIngressOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference{}

	_jsii_.Create(
		"awscc.ec2SecurityGroup.Ec2SecurityGroupSecurityGroupIngressOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2SecurityGroupSecurityGroupIngressOutputReference_Override(e Ec2SecurityGroupSecurityGroupIngressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2SecurityGroup.Ec2SecurityGroupSecurityGroupIngressOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference)SetCidrIp(val *string) {
	if err := j.validateSetCidrIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrIp",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference)SetCidrIpv6(val *string) {
	if err := j.validateSetCidrIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrIpv6",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference)SetFromPort(val *float64) {
	if err := j.validateSetFromPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fromPort",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference)SetIpProtocol(val *string) {
	if err := j.validateSetIpProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipProtocol",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference)SetSourcePrefixListId(val *string) {
	if err := j.validateSetSourcePrefixListIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePrefixListId",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference)SetSourceSecurityGroupId(val *string) {
	if err := j.validateSetSourceSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceSecurityGroupId",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference)SetSourceSecurityGroupName(val *string) {
	if err := j.validateSetSourceSecurityGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceSecurityGroupName",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference)SetSourceSecurityGroupOwnerId(val *string) {
	if err := j.validateSetSourceSecurityGroupOwnerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceSecurityGroupOwnerId",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference)SetToPort(val *float64) {
	if err := j.validateSetToPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toPort",
		val,
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) ResetCidrIp() {
	_jsii_.InvokeVoid(
		e,
		"resetCidrIp",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) ResetCidrIpv6() {
	_jsii_.InvokeVoid(
		e,
		"resetCidrIpv6",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) ResetFromPort() {
	_jsii_.InvokeVoid(
		e,
		"resetFromPort",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) ResetSourcePrefixListId() {
	_jsii_.InvokeVoid(
		e,
		"resetSourcePrefixListId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) ResetSourceSecurityGroupId() {
	_jsii_.InvokeVoid(
		e,
		"resetSourceSecurityGroupId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) ResetSourceSecurityGroupName() {
	_jsii_.InvokeVoid(
		e,
		"resetSourceSecurityGroupName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) ResetSourceSecurityGroupOwnerId() {
	_jsii_.InvokeVoid(
		e,
		"resetSourceSecurityGroupOwnerId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) ResetToPort() {
	_jsii_.InvokeVoid(
		e,
		"resetToPort",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupIngressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

