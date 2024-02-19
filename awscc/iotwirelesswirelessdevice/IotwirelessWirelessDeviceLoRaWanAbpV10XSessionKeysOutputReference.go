package iotwirelesswirelessdevice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/iotwirelesswirelessdevice/internal"
)

type IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference interface {
	cdktf.ComplexObject
	AppSKey() *string
	SetAppSKey(val *string)
	AppSKeyInput() *string
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
	InternalValue() *IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeys
	SetInternalValue(val *IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeys)
	NwkSKey() *string
	SetNwkSKey(val *string)
	NwkSKeyInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference
type jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) AppSKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) AppSKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) InternalValue() *IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeys {
	var returns *IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeys
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) NwkSKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nwkSKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) NwkSKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nwkSKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference {
	_init_.Initialize()

	if err := validateNewIotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference{}

	_jsii_.Create(
		"awscc.iotwirelessWirelessDevice.IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference_Override(i IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.iotwirelessWirelessDevice.IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference)SetAppSKey(val *string) {
	if err := j.validateSetAppSKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appSKey",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference)SetInternalValue(val *IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeys) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference)SetNwkSKey(val *string) {
	if err := j.validateSetNwkSKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nwkSKey",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeysOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

