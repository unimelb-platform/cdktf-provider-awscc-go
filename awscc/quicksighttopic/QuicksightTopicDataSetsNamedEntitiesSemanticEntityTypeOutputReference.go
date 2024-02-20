package quicksighttopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/quicksighttopic/internal"
)

type QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference interface {
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
	SubTypeName() *string
	SetSubTypeName(val *string)
	SubTypeNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TypeName() *string
	SetTypeName(val *string)
	TypeNameInput() *string
	TypeParameters() *map[string]*string
	SetTypeParameters(val *map[string]*string)
	TypeParametersInput() *map[string]*string
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
	ResetSubTypeName()
	ResetTypeName()
	ResetTypeParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference
type jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) SubTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) SubTypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTypeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) TypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) TypeParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"typeParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) TypeParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"typeParametersInput",
		&returns,
	)
	return returns
}


func NewQuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference{}

	_jsii_.Create(
		"awscc.quicksightTopic.QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference_Override(q QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.quicksightTopic.QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference)SetSubTypeName(val *string) {
	if err := j.validateSetSubTypeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subTypeName",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference)SetTypeName(val *string) {
	if err := j.validateSetTypeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference)SetTypeParameters(val *map[string]*string) {
	if err := j.validateSetTypeParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeParameters",
		val,
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) ResetSubTypeName() {
	_jsii_.InvokeVoid(
		q,
		"resetSubTypeName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) ResetTypeName() {
	_jsii_.InvokeVoid(
		q,
		"resetTypeName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) ResetTypeParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetTypeParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

