package quicksighttopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/quicksighttopic/internal"
)

type QuicksightTopicDataSetsNamedEntitiesOutputReference interface {
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
	Definition() QuicksightTopicDataSetsNamedEntitiesDefinitionList
	DefinitionInput() interface{}
	EntityDescription() *string
	SetEntityDescription(val *string)
	EntityDescriptionInput() *string
	EntityName() *string
	SetEntityName(val *string)
	EntityNameInput() *string
	EntitySynonyms() *[]*string
	SetEntitySynonyms(val *[]*string)
	EntitySynonymsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SemanticEntityType() QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference
	SemanticEntityTypeInput() interface{}
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
	PutDefinition(value interface{})
	PutSemanticEntityType(value *QuicksightTopicDataSetsNamedEntitiesSemanticEntityType)
	ResetDefinition()
	ResetEntityDescription()
	ResetEntitySynonyms()
	ResetSemanticEntityType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightTopicDataSetsNamedEntitiesOutputReference
type jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) Definition() QuicksightTopicDataSetsNamedEntitiesDefinitionList {
	var returns QuicksightTopicDataSetsNamedEntitiesDefinitionList
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) DefinitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) EntityDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) EntityDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) EntityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) EntityNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) EntitySynonyms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entitySynonyms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) EntitySynonymsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entitySynonymsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) SemanticEntityType() QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference {
	var returns QuicksightTopicDataSetsNamedEntitiesSemanticEntityTypeOutputReference
	_jsii_.Get(
		j,
		"semanticEntityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) SemanticEntityTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"semanticEntityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightTopicDataSetsNamedEntitiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) QuicksightTopicDataSetsNamedEntitiesOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightTopicDataSetsNamedEntitiesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference{}

	_jsii_.Create(
		"awscc.quicksightTopic.QuicksightTopicDataSetsNamedEntitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewQuicksightTopicDataSetsNamedEntitiesOutputReference_Override(q QuicksightTopicDataSetsNamedEntitiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.quicksightTopic.QuicksightTopicDataSetsNamedEntitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference)SetEntityDescription(val *string) {
	if err := j.validateSetEntityDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityDescription",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference)SetEntityName(val *string) {
	if err := j.validateSetEntityNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityName",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference)SetEntitySynonyms(val *[]*string) {
	if err := j.validateSetEntitySynonymsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entitySynonyms",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) PutDefinition(value interface{}) {
	if err := q.validatePutDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDefinition",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) PutSemanticEntityType(value *QuicksightTopicDataSetsNamedEntitiesSemanticEntityType) {
	if err := q.validatePutSemanticEntityTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSemanticEntityType",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) ResetDefinition() {
	_jsii_.InvokeVoid(
		q,
		"resetDefinition",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) ResetEntityDescription() {
	_jsii_.InvokeVoid(
		q,
		"resetEntityDescription",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) ResetEntitySynonyms() {
	_jsii_.InvokeVoid(
		q,
		"resetEntitySynonyms",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) ResetSemanticEntityType() {
	_jsii_.InvokeVoid(
		q,
		"resetSemanticEntityType",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightTopicDataSetsNamedEntitiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

