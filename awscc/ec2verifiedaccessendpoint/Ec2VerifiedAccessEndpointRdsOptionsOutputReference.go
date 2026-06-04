package ec2verifiedaccessendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2verifiedaccessendpoint/internal"
)

type Ec2VerifiedAccessEndpointRdsOptionsOutputReference interface {
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
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	RdsDbClusterArn() *string
	SetRdsDbClusterArn(val *string)
	RdsDbClusterArnInput() *string
	RdsDbInstanceArn() *string
	SetRdsDbInstanceArn(val *string)
	RdsDbInstanceArnInput() *string
	RdsDbProxyArn() *string
	SetRdsDbProxyArn(val *string)
	RdsDbProxyArnInput() *string
	RdsEndpoint() *string
	SetRdsEndpoint(val *string)
	RdsEndpointInput() *string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
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
	ResetPort()
	ResetProtocol()
	ResetRdsDbClusterArn()
	ResetRdsDbInstanceArn()
	ResetRdsDbProxyArn()
	ResetRdsEndpoint()
	ResetSubnetIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2VerifiedAccessEndpointRdsOptionsOutputReference
type jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) RdsDbClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) RdsDbClusterArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbClusterArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) RdsDbInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) RdsDbInstanceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbInstanceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) RdsDbProxyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbProxyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) RdsDbProxyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbProxyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) RdsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) RdsEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2VerifiedAccessEndpointRdsOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Ec2VerifiedAccessEndpointRdsOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2VerifiedAccessEndpointRdsOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference{}

	_jsii_.Create(
		"awscc.ec2VerifiedAccessEndpoint.Ec2VerifiedAccessEndpointRdsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2VerifiedAccessEndpointRdsOptionsOutputReference_Override(e Ec2VerifiedAccessEndpointRdsOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2VerifiedAccessEndpoint.Ec2VerifiedAccessEndpointRdsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference)SetRdsDbClusterArn(val *string) {
	if err := j.validateSetRdsDbClusterArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsDbClusterArn",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference)SetRdsDbInstanceArn(val *string) {
	if err := j.validateSetRdsDbInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsDbInstanceArn",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference)SetRdsDbProxyArn(val *string) {
	if err := j.validateSetRdsDbProxyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsDbProxyArn",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference)SetRdsEndpoint(val *string) {
	if err := j.validateSetRdsEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsEndpoint",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		e,
		"resetPort",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		e,
		"resetProtocol",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) ResetRdsDbClusterArn() {
	_jsii_.InvokeVoid(
		e,
		"resetRdsDbClusterArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) ResetRdsDbInstanceArn() {
	_jsii_.InvokeVoid(
		e,
		"resetRdsDbInstanceArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) ResetRdsDbProxyArn() {
	_jsii_.InvokeVoid(
		e,
		"resetRdsDbProxyArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) ResetRdsEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetRdsEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		e,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpointRdsOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

