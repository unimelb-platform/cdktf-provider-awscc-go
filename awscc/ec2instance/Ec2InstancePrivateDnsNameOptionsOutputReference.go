package ec2instance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2instance/internal"
)

type Ec2InstancePrivateDnsNameOptionsOutputReference interface {
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
	EnableResourceNameDnsAaaaRecord() interface{}
	SetEnableResourceNameDnsAaaaRecord(val interface{})
	EnableResourceNameDnsAaaaRecordInput() interface{}
	EnableResourceNameDnsARecord() interface{}
	SetEnableResourceNameDnsARecord(val interface{})
	EnableResourceNameDnsARecordInput() interface{}
	// Experimental.
	Fqn() *string
	HostnameType() *string
	SetHostnameType(val *string)
	HostnameTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetEnableResourceNameDnsAaaaRecord()
	ResetEnableResourceNameDnsARecord()
	ResetHostnameType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2InstancePrivateDnsNameOptionsOutputReference
type jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) EnableResourceNameDnsAaaaRecord() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsAaaaRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) EnableResourceNameDnsAaaaRecordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsAaaaRecordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) EnableResourceNameDnsARecord() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsARecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) EnableResourceNameDnsARecordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsARecordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) HostnameType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) HostnameTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2InstancePrivateDnsNameOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Ec2InstancePrivateDnsNameOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2InstancePrivateDnsNameOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference{}

	_jsii_.Create(
		"awscc.ec2Instance.Ec2InstancePrivateDnsNameOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2InstancePrivateDnsNameOptionsOutputReference_Override(e Ec2InstancePrivateDnsNameOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2Instance.Ec2InstancePrivateDnsNameOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference)SetEnableResourceNameDnsAaaaRecord(val interface{}) {
	if err := j.validateSetEnableResourceNameDnsAaaaRecordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourceNameDnsAaaaRecord",
		val,
	)
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference)SetEnableResourceNameDnsARecord(val interface{}) {
	if err := j.validateSetEnableResourceNameDnsARecordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourceNameDnsARecord",
		val,
	)
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference)SetHostnameType(val *string) {
	if err := j.validateSetHostnameTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnameType",
		val,
	)
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) ResetEnableResourceNameDnsAaaaRecord() {
	_jsii_.InvokeVoid(
		e,
		"resetEnableResourceNameDnsAaaaRecord",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) ResetEnableResourceNameDnsARecord() {
	_jsii_.InvokeVoid(
		e,
		"resetEnableResourceNameDnsARecord",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) ResetHostnameType() {
	_jsii_.InvokeVoid(
		e,
		"resetHostnameType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2InstancePrivateDnsNameOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

