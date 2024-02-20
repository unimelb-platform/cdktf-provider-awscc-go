package dataawsccgreengrassv2componentversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccgreengrassv2componentversion/internal"
)

type DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference interface {
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
	Devices() DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsDevicesList
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParams
	SetInternalValue(val *DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParams)
	MemorySizeInKb() *float64
	MountRoSysfs() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Volumes() DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsVolumesList
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

// The jsii proxy struct for DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference
type jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) Devices() DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsDevicesList {
	var returns DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsDevicesList
	_jsii_.Get(
		j,
		"devices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) InternalValue() *DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParams {
	var returns *DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParams
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) MemorySizeInKb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInKb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) MountRoSysfs() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mountRoSysfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) Volumes() DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsVolumesList {
	var returns DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsVolumesList
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


func NewDataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccGreengrassv2ComponentVersion.DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference_Override(d DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccGreengrassv2ComponentVersion.DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference)SetInternalValue(val *DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParams) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGreengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

