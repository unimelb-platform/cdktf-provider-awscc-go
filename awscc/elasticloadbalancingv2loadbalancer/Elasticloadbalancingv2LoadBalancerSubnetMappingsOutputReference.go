package elasticloadbalancingv2loadbalancer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/elasticloadbalancingv2loadbalancer/internal"
)

type Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference interface {
	cdktf.ComplexObject
	AllocationId() *string
	SetAllocationId(val *string)
	AllocationIdInput() *string
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
	IPv6Address() *string
	SetIPv6Address(val *string)
	IPv6AddressInput() *string
	PrivateIPv4Address() *string
	SetPrivateIPv4Address(val *string)
	PrivateIPv4AddressInput() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
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
	ResetAllocationId()
	ResetIPv6Address()
	ResetPrivateIPv4Address()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference
type jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) AllocationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) AllocationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) IPv6Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iPv6Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) IPv6AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iPv6AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) PrivateIPv4Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIPv4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) PrivateIPv4AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIPv4AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference {
	_init_.Initialize()

	if err := validateNewElasticloadbalancingv2LoadBalancerSubnetMappingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference{}

	_jsii_.Create(
		"awscc.elasticloadbalancingv2LoadBalancer.Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewElasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference_Override(e Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.elasticloadbalancingv2LoadBalancer.Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference)SetAllocationId(val *string) {
	if err := j.validateSetAllocationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationId",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference)SetIPv6Address(val *string) {
	if err := j.validateSetIPv6AddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iPv6Address",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference)SetPrivateIPv4Address(val *string) {
	if err := j.validateSetPrivateIPv4AddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIPv4Address",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) ResetAllocationId() {
	_jsii_.InvokeVoid(
		e,
		"resetAllocationId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) ResetIPv6Address() {
	_jsii_.InvokeVoid(
		e,
		"resetIPv6Address",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) ResetPrivateIPv4Address() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivateIPv4Address",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Elasticloadbalancingv2LoadBalancerSubnetMappingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

