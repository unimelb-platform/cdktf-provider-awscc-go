package dataawsccquicksighttopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccquicksighttopic/internal"
)

type DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference interface {
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
	FalseyCellValue() *string
	FalseyCellValueSynonyms() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticType
	SetInternalValue(val *DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticType)
	SubTypeName() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TruthyCellValue() *string
	TruthyCellValueSynonyms() *[]*string
	TypeName() *string
	TypeParameters() cdktf.StringMap
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

// The jsii proxy struct for DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference
type jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) FalseyCellValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"falseyCellValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) FalseyCellValueSynonyms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"falseyCellValueSynonyms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) InternalValue() *DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticType {
	var returns *DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticType
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) SubTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) TruthyCellValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truthyCellValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) TruthyCellValueSynonyms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"truthyCellValueSynonyms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) TypeParameters() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"typeParameters",
		&returns,
	)
	return returns
}


func NewDataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccQuicksightTopic.DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference_Override(d DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccQuicksightTopic.DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetInternalValue(val *DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticType) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

