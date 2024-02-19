package dataawsccguarddutyfilter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccguarddutyfilter/internal"
)

type DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference interface {
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
	Eq() *[]*string
	EqualTo() *[]*string
	// Experimental.
	Fqn() *string
	GreaterThan() *float64
	GreaterThanOrEqual() *float64
	Gt() *float64
	Gte() *float64
	InternalValue() *DataAwsccGuarddutyFilterFindingCriteriaCriterion
	SetInternalValue(val *DataAwsccGuarddutyFilterFindingCriteriaCriterion)
	LessThan() *float64
	LessThanOrEqual() *float64
	Lt() *float64
	Lte() *float64
	Neq() *[]*string
	NotEquals() *[]*string
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

// The jsii proxy struct for DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference
type jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) Eq() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) EqualTo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) GreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"greaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) GreaterThanOrEqual() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"greaterThanOrEqual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) Gt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) Gte() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) InternalValue() *DataAwsccGuarddutyFilterFindingCriteriaCriterion {
	var returns *DataAwsccGuarddutyFilterFindingCriteriaCriterion
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) LessThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lessThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) LessThanOrEqual() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lessThanOrEqual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) Lt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) Lte() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) Neq() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"neq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) NotEquals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectKey *string) DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectKey); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccGuarddutyFilter.DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectKey},
		&j,
	)

	return &j
}

func NewDataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference_Override(d DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectKey *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccGuarddutyFilter.DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectKey},
		d,
	)
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference)SetInternalValue(val *DataAwsccGuarddutyFilterFindingCriteriaCriterion) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccGuarddutyFilterFindingCriteriaCriterionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

