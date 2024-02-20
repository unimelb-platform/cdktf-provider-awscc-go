package iotwirelesswirelessdevice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/iotwirelesswirelessdevice/internal"
)

type IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference interface {
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
	FNwkSIntKey() *string
	SetFNwkSIntKey(val *string)
	FNwkSIntKeyInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeys
	SetInternalValue(val *IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeys)
	NwkSEncKey() *string
	SetNwkSEncKey(val *string)
	NwkSEncKeyInput() *string
	SNwkSIntKey() *string
	SetSNwkSIntKey(val *string)
	SNwkSIntKeyInput() *string
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

// The jsii proxy struct for IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference
type jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) AppSKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) AppSKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) FNwkSIntKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fNwkSIntKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) FNwkSIntKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fNwkSIntKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) InternalValue() *IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeys {
	var returns *IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeys
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) NwkSEncKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nwkSEncKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) NwkSEncKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nwkSEncKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) SNwkSIntKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sNwkSIntKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) SNwkSIntKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sNwkSIntKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference {
	_init_.Initialize()

	if err := validateNewIotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference{}

	_jsii_.Create(
		"awscc.iotwirelessWirelessDevice.IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference_Override(i IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.iotwirelessWirelessDevice.IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference)SetAppSKey(val *string) {
	if err := j.validateSetAppSKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appSKey",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference)SetFNwkSIntKey(val *string) {
	if err := j.validateSetFNwkSIntKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fNwkSIntKey",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference)SetInternalValue(val *IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeys) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference)SetNwkSEncKey(val *string) {
	if err := j.validateSetNwkSEncKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nwkSEncKey",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference)SetSNwkSIntKey(val *string) {
	if err := j.validateSetSNwkSIntKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sNwkSIntKey",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeysOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

