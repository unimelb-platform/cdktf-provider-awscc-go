package lexbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lexbot/internal"
)

type LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference interface {
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
	InternalValue() *LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStruct
	SetInternalValue(val *LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStruct)
	Message() LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListMessageOutputReference
	MessageInput() *LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListMessage
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Variations() LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListVariationsList
	VariationsInput() interface{}
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
	PutMessage(value *LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListMessage)
	PutVariations(value interface{})
	ResetVariations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference
type jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) InternalValue() *LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStruct {
	var returns *LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStruct
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) Message() LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListMessageOutputReference {
	var returns LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListMessageOutputReference
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) MessageInput() *LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListMessage {
	var returns *LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListMessage
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) Variations() LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListVariationsList {
	var returns LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListVariationsList
	_jsii_.Get(
		j,
		"variations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) VariationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variationsInput",
		&returns,
	)
	return returns
}


func NewLexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference {
	_init_.Initialize()

	if err := validateNewLexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference{}

	_jsii_.Create(
		"awscc.lexBot.LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference_Override(l LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lexBot.LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference)SetInternalValue(val *LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStruct) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) PutMessage(value *LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListMessage) {
	if err := l.validatePutMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMessage",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) PutVariations(value interface{}) {
	if err := l.validatePutVariationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putVariations",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) ResetVariations() {
	_jsii_.InvokeVoid(
		l,
		"resetVariations",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupsListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

