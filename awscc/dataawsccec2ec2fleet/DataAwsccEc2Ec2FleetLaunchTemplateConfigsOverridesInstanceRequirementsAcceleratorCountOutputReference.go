package dataawsccec2ec2fleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccec2ec2fleet/internal"
)

type DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference interface {
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
	InternalValue() *DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCount
	SetInternalValue(val *DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCount)
	Max() *float64
	Min() *float64
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

// The jsii proxy struct for DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference
type jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) InternalValue() *DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCount {
	var returns *DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCount
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) Max() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"max",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) Min() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"min",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccEc2Ec2Fleet.DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference_Override(d DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccEc2Ec2Fleet.DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference)SetInternalValue(val *DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCount) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsAcceleratorCountOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

