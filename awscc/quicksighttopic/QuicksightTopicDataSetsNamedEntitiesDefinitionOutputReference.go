package quicksighttopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/quicksighttopic/internal"
)

type QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference interface {
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
	FieldName() *string
	SetFieldName(val *string)
	FieldNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Metric() QuicksightTopicDataSetsNamedEntitiesDefinitionMetricOutputReference
	MetricInput() interface{}
	PropertyName() *string
	SetPropertyName(val *string)
	PropertyNameInput() *string
	PropertyRole() *string
	SetPropertyRole(val *string)
	PropertyRoleInput() *string
	PropertyUsage() *string
	SetPropertyUsage(val *string)
	PropertyUsageInput() *string
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
	PutMetric(value *QuicksightTopicDataSetsNamedEntitiesDefinitionMetric)
	ResetFieldName()
	ResetMetric()
	ResetPropertyName()
	ResetPropertyRole()
	ResetPropertyUsage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference
type jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) FieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) FieldNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) Metric() QuicksightTopicDataSetsNamedEntitiesDefinitionMetricOutputReference {
	var returns QuicksightTopicDataSetsNamedEntitiesDefinitionMetricOutputReference
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) MetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) PropertyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) PropertyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) PropertyRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) PropertyRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) PropertyUsage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) PropertyUsageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightTopicDataSetsNamedEntitiesDefinitionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference{}

	_jsii_.Create(
		"awscc.quicksightTopic.QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewQuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference_Override(q QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.quicksightTopic.QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference)SetFieldName(val *string) {
	if err := j.validateSetFieldNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldName",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference)SetPropertyName(val *string) {
	if err := j.validateSetPropertyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyName",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference)SetPropertyRole(val *string) {
	if err := j.validateSetPropertyRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyRole",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference)SetPropertyUsage(val *string) {
	if err := j.validateSetPropertyUsageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyUsage",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) PutMetric(value *QuicksightTopicDataSetsNamedEntitiesDefinitionMetric) {
	if err := q.validatePutMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putMetric",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) ResetFieldName() {
	_jsii_.InvokeVoid(
		q,
		"resetFieldName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) ResetMetric() {
	_jsii_.InvokeVoid(
		q,
		"resetMetric",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) ResetPropertyName() {
	_jsii_.InvokeVoid(
		q,
		"resetPropertyName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) ResetPropertyRole() {
	_jsii_.InvokeVoid(
		q,
		"resetPropertyRole",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) ResetPropertyUsage() {
	_jsii_.InvokeVoid(
		q,
		"resetPropertyUsage",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

