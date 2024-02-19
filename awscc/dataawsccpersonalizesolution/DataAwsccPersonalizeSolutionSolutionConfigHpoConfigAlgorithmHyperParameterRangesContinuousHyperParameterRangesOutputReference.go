package dataawsccpersonalizesolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccpersonalizesolution/internal"
)

type DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference interface {
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
	InternalValue() *DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRanges
	SetInternalValue(val *DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRanges)
	MaxValue() *float64
	MinValue() *float64
	Name() *string
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

// The jsii proxy struct for DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference
type jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) InternalValue() *DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRanges {
	var returns *DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRanges
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) MaxValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) MinValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccPersonalizeSolution.DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference_Override(d DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccPersonalizeSolution.DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference)SetInternalValue(val *DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRanges) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

