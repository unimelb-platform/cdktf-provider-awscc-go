package quicksighttopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/quicksighttopic/internal"
)

type QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference interface {
	cdktf.ComplexObject
	Aggregation() *string
	SetAggregation(val *string)
	AggregationInput() *string
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
	Constant() QuicksightTopicDataSetsFiltersNumericRangeFilterConstantOutputReference
	ConstantInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Inclusive() interface{}
	SetInclusive(val interface{})
	InclusiveInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutConstant(value *QuicksightTopicDataSetsFiltersNumericRangeFilterConstant)
	ResetAggregation()
	ResetConstant()
	ResetInclusive()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference
type jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) Aggregation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) AggregationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) Constant() QuicksightTopicDataSetsFiltersNumericRangeFilterConstantOutputReference {
	var returns QuicksightTopicDataSetsFiltersNumericRangeFilterConstantOutputReference
	_jsii_.Get(
		j,
		"constant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) ConstantInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"constantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) Inclusive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inclusive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) InclusiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inclusiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightTopicDataSetsFiltersNumericRangeFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference{}

	_jsii_.Create(
		"awscc.quicksightTopic.QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference_Override(q QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.quicksightTopic.QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference)SetAggregation(val *string) {
	if err := j.validateSetAggregationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregation",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference)SetInclusive(val interface{}) {
	if err := j.validateSetInclusiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inclusive",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) PutConstant(value *QuicksightTopicDataSetsFiltersNumericRangeFilterConstant) {
	if err := q.validatePutConstantParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putConstant",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) ResetAggregation() {
	_jsii_.InvokeVoid(
		q,
		"resetAggregation",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) ResetConstant() {
	_jsii_.InvokeVoid(
		q,
		"resetConstant",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) ResetInclusive() {
	_jsii_.InvokeVoid(
		q,
		"resetInclusive",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := q.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

