package iotsitewisegateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/iotsitewisegateway/internal"
)

type IotsitewiseGatewayGatewayPlatformOutputReference interface {
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
	Greengrass() IotsitewiseGatewayGatewayPlatformGreengrassOutputReference
	GreengrassInput() interface{}
	GreengrassV2() IotsitewiseGatewayGatewayPlatformGreengrassV2OutputReference
	GreengrassV2Input() interface{}
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
	PutGreengrass(value *IotsitewiseGatewayGatewayPlatformGreengrass)
	PutGreengrassV2(value *IotsitewiseGatewayGatewayPlatformGreengrassV2)
	ResetGreengrass()
	ResetGreengrassV2()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotsitewiseGatewayGatewayPlatformOutputReference
type jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) Greengrass() IotsitewiseGatewayGatewayPlatformGreengrassOutputReference {
	var returns IotsitewiseGatewayGatewayPlatformGreengrassOutputReference
	_jsii_.Get(
		j,
		"greengrass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) GreengrassInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"greengrassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) GreengrassV2() IotsitewiseGatewayGatewayPlatformGreengrassV2OutputReference {
	var returns IotsitewiseGatewayGatewayPlatformGreengrassV2OutputReference
	_jsii_.Get(
		j,
		"greengrassV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) GreengrassV2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"greengrassV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotsitewiseGatewayGatewayPlatformOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IotsitewiseGatewayGatewayPlatformOutputReference {
	_init_.Initialize()

	if err := validateNewIotsitewiseGatewayGatewayPlatformOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference{}

	_jsii_.Create(
		"awscc.iotsitewiseGateway.IotsitewiseGatewayGatewayPlatformOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotsitewiseGatewayGatewayPlatformOutputReference_Override(i IotsitewiseGatewayGatewayPlatformOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.iotsitewiseGateway.IotsitewiseGatewayGatewayPlatformOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) PutGreengrass(value *IotsitewiseGatewayGatewayPlatformGreengrass) {
	if err := i.validatePutGreengrassParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putGreengrass",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) PutGreengrassV2(value *IotsitewiseGatewayGatewayPlatformGreengrassV2) {
	if err := i.validatePutGreengrassV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putGreengrassV2",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) ResetGreengrass() {
	_jsii_.InvokeVoid(
		i,
		"resetGreengrass",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) ResetGreengrassV2() {
	_jsii_.InvokeVoid(
		i,
		"resetGreengrassV2",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseGatewayGatewayPlatformOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

