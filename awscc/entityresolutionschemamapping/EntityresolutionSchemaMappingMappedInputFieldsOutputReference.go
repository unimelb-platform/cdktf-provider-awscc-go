package entityresolutionschemamapping

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/entityresolutionschemamapping/internal"
)

type EntityresolutionSchemaMappingMappedInputFieldsOutputReference interface {
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
	GroupName() *string
	SetGroupName(val *string)
	GroupNameInput() *string
	Hashed() interface{}
	SetHashed(val interface{})
	HashedInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MatchKey() *string
	SetMatchKey(val *string)
	MatchKeyInput() *string
	SubType() *string
	SetSubType(val *string)
	SubTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetGroupName()
	ResetHashed()
	ResetMatchKey()
	ResetSubType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EntityresolutionSchemaMappingMappedInputFieldsOutputReference
type jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) FieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) FieldNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) GroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) GroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) Hashed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hashed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) HashedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hashedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) MatchKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) MatchKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) SubType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) SubTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewEntityresolutionSchemaMappingMappedInputFieldsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EntityresolutionSchemaMappingMappedInputFieldsOutputReference {
	_init_.Initialize()

	if err := validateNewEntityresolutionSchemaMappingMappedInputFieldsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference{}

	_jsii_.Create(
		"awscc.entityresolutionSchemaMapping.EntityresolutionSchemaMappingMappedInputFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEntityresolutionSchemaMappingMappedInputFieldsOutputReference_Override(e EntityresolutionSchemaMappingMappedInputFieldsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.entityresolutionSchemaMapping.EntityresolutionSchemaMappingMappedInputFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference)SetFieldName(val *string) {
	if err := j.validateSetFieldNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldName",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference)SetGroupName(val *string) {
	if err := j.validateSetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupName",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference)SetHashed(val interface{}) {
	if err := j.validateSetHashedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashed",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference)SetMatchKey(val *string) {
	if err := j.validateSetMatchKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchKey",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference)SetSubType(val *string) {
	if err := j.validateSetSubTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subType",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) ResetGroupName() {
	_jsii_.InvokeVoid(
		e,
		"resetGroupName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) ResetHashed() {
	_jsii_.InvokeVoid(
		e,
		"resetHashed",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) ResetMatchKey() {
	_jsii_.InvokeVoid(
		e,
		"resetMatchKey",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) ResetSubType() {
	_jsii_.InvokeVoid(
		e,
		"resetSubType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionSchemaMappingMappedInputFieldsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

