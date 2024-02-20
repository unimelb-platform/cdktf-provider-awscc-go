package timestreamscheduledquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/timestreamscheduledquery/internal"
)

type TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MultiMeasureAttributeMappings() TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingsList
	MultiMeasureAttributeMappingsInput() interface{}
	TargetMultiMeasureName() *string
	SetTargetMultiMeasureName(val *string)
	TargetMultiMeasureNameInput() *string
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
	PutMultiMeasureAttributeMappings(value interface{})
	ResetTargetMultiMeasureName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference
type jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) MultiMeasureAttributeMappings() TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingsList {
	var returns TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingsList
	_jsii_.Get(
		j,
		"multiMeasureAttributeMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) MultiMeasureAttributeMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiMeasureAttributeMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) TargetMultiMeasureName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetMultiMeasureName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) TargetMultiMeasureNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetMultiMeasureNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference {
	_init_.Initialize()

	if err := validateNewTimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference{}

	_jsii_.Create(
		"awscc.timestreamScheduledQuery.TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference_Override(t TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.timestreamScheduledQuery.TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference)SetTargetMultiMeasureName(val *string) {
	if err := j.validateSetTargetMultiMeasureNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetMultiMeasureName",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) PutMultiMeasureAttributeMappings(value interface{}) {
	if err := t.validatePutMultiMeasureAttributeMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMultiMeasureAttributeMappings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) ResetTargetMultiMeasureName() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetMultiMeasureName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

