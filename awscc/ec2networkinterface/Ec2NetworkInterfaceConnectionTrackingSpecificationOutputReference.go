package ec2networkinterface

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2networkinterface/internal"
)

type Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference interface {
	cdktf.ComplexObject
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
	TcpEstablishedTimeout() *float64
	SetTcpEstablishedTimeout(val *float64)
	TcpEstablishedTimeoutInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UdpStreamTimeout() *float64
	SetUdpStreamTimeout(val *float64)
	UdpStreamTimeoutInput() *float64
	UdpTimeout() *float64
	SetUdpTimeout(val *float64)
	UdpTimeoutInput() *float64
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
	ResetTcpEstablishedTimeout()
	ResetUdpStreamTimeout()
	ResetUdpTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference
type jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) TcpEstablishedTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpEstablishedTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) TcpEstablishedTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpEstablishedTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) UdpStreamTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpStreamTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) UdpStreamTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpStreamTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) UdpTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) UdpTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpTimeoutInput",
		&returns,
	)
	return returns
}


func NewEc2NetworkInterfaceConnectionTrackingSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewEc2NetworkInterfaceConnectionTrackingSpecificationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference{}

	_jsii_.Create(
		"awscc.ec2NetworkInterface.Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2NetworkInterfaceConnectionTrackingSpecificationOutputReference_Override(e Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2NetworkInterface.Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference)SetTcpEstablishedTimeout(val *float64) {
	if err := j.validateSetTcpEstablishedTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpEstablishedTimeout",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference)SetUdpStreamTimeout(val *float64) {
	if err := j.validateSetUdpStreamTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"udpStreamTimeout",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference)SetUdpTimeout(val *float64) {
	if err := j.validateSetUdpTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"udpTimeout",
		val,
	)
}

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) ResetTcpEstablishedTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetTcpEstablishedTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) ResetUdpStreamTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetUdpStreamTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) ResetUdpTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetUdpTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

