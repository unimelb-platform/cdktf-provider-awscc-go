package iotsitewiseassetmodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/iotsitewiseassetmodel/internal"
)

type IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference interface {
	cdktf.ComplexObject
	Attribute() IotsitewiseAssetModelAssetModelPropertiesTypeAttributeOutputReference
	AttributeInput() interface{}
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
	InternalValue() *IotsitewiseAssetModelAssetModelPropertiesType
	SetInternalValue(val *IotsitewiseAssetModelAssetModelPropertiesType)
	Metric() IotsitewiseAssetModelAssetModelPropertiesTypeMetricOutputReference
	MetricInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Transform() IotsitewiseAssetModelAssetModelPropertiesTypeTransformOutputReference
	TransformInput() interface{}
	TypeName() *string
	SetTypeName(val *string)
	TypeNameInput() *string
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
	PutAttribute(value *IotsitewiseAssetModelAssetModelPropertiesTypeAttribute)
	PutMetric(value *IotsitewiseAssetModelAssetModelPropertiesTypeMetric)
	PutTransform(value *IotsitewiseAssetModelAssetModelPropertiesTypeTransform)
	ResetAttribute()
	ResetMetric()
	ResetTransform()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference
type jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) Attribute() IotsitewiseAssetModelAssetModelPropertiesTypeAttributeOutputReference {
	var returns IotsitewiseAssetModelAssetModelPropertiesTypeAttributeOutputReference
	_jsii_.Get(
		j,
		"attribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) AttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) InternalValue() *IotsitewiseAssetModelAssetModelPropertiesType {
	var returns *IotsitewiseAssetModelAssetModelPropertiesType
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) Metric() IotsitewiseAssetModelAssetModelPropertiesTypeMetricOutputReference {
	var returns IotsitewiseAssetModelAssetModelPropertiesTypeMetricOutputReference
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) MetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) Transform() IotsitewiseAssetModelAssetModelPropertiesTypeTransformOutputReference {
	var returns IotsitewiseAssetModelAssetModelPropertiesTypeTransformOutputReference
	_jsii_.Get(
		j,
		"transform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) TransformInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) TypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeNameInput",
		&returns,
	)
	return returns
}


func NewIotsitewiseAssetModelAssetModelPropertiesTypeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference {
	_init_.Initialize()

	if err := validateNewIotsitewiseAssetModelAssetModelPropertiesTypeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference{}

	_jsii_.Create(
		"awscc.iotsitewiseAssetModel.IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotsitewiseAssetModelAssetModelPropertiesTypeOutputReference_Override(i IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.iotsitewiseAssetModel.IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference)SetInternalValue(val *IotsitewiseAssetModelAssetModelPropertiesType) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference)SetTypeName(val *string) {
	if err := j.validateSetTypeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) PutAttribute(value *IotsitewiseAssetModelAssetModelPropertiesTypeAttribute) {
	if err := i.validatePutAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAttribute",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) PutMetric(value *IotsitewiseAssetModelAssetModelPropertiesTypeMetric) {
	if err := i.validatePutMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMetric",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) PutTransform(value *IotsitewiseAssetModelAssetModelPropertiesTypeTransform) {
	if err := i.validatePutTransformParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTransform",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) ResetAttribute() {
	_jsii_.InvokeVoid(
		i,
		"resetAttribute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) ResetMetric() {
	_jsii_.InvokeVoid(
		i,
		"resetMetric",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) ResetTransform() {
	_jsii_.InvokeVoid(
		i,
		"resetTransform",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

