package ec2launchtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2launchtemplate/internal"
)

type Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference interface {
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
	DeviceName() *string
	SetDeviceName(val *string)
	DeviceNameInput() *string
	Ebs() Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsEbsOutputReference
	EbsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NoDevice() *string
	SetNoDevice(val *string)
	NoDeviceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualName() *string
	SetVirtualName(val *string)
	VirtualNameInput() *string
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
	PutEbs(value *Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsEbs)
	ResetDeviceName()
	ResetEbs()
	ResetNoDevice()
	ResetVirtualName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference
type jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) DeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) DeviceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) Ebs() Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsEbsOutputReference {
	var returns Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsEbsOutputReference
	_jsii_.Get(
		j,
		"ebs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) EbsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) NoDevice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) NoDeviceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) VirtualName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) VirtualNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNameInput",
		&returns,
	)
	return returns
}


func NewEc2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference{}

	_jsii_.Create(
		"awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference_Override(e Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference)SetDeviceName(val *string) {
	if err := j.validateSetDeviceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceName",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference)SetNoDevice(val *string) {
	if err := j.validateSetNoDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDevice",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference)SetVirtualName(val *string) {
	if err := j.validateSetVirtualNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualName",
		val,
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) PutEbs(value *Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsEbs) {
	if err := e.validatePutEbsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEbs",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) ResetDeviceName() {
	_jsii_.InvokeVoid(
		e,
		"resetDeviceName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) ResetEbs() {
	_jsii_.InvokeVoid(
		e,
		"resetEbs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) ResetNoDevice() {
	_jsii_.InvokeVoid(
		e,
		"resetNoDevice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) ResetVirtualName() {
	_jsii_.InvokeVoid(
		e,
		"resetVirtualName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

