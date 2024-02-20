package lexbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lexbot/internal"
)

type LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference interface {
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
	CustomPayload() LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsCustomPayloadOutputReference
	CustomPayloadInput() interface{}
	// Experimental.
	Fqn() *string
	ImageResponseCard() LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardOutputReference
	ImageResponseCardInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PlainTextMessage() LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsPlainTextMessageOutputReference
	PlainTextMessageInput() interface{}
	SsmlMessage() LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsSsmlMessageOutputReference
	SsmlMessageInput() interface{}
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
	PutCustomPayload(value *LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsCustomPayload)
	PutImageResponseCard(value *LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCard)
	PutPlainTextMessage(value *LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsPlainTextMessage)
	PutSsmlMessage(value *LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsSsmlMessage)
	ResetCustomPayload()
	ResetImageResponseCard()
	ResetPlainTextMessage()
	ResetSsmlMessage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference
type jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) CustomPayload() LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsCustomPayloadOutputReference {
	var returns LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsCustomPayloadOutputReference
	_jsii_.Get(
		j,
		"customPayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) CustomPayloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customPayloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) ImageResponseCard() LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardOutputReference {
	var returns LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardOutputReference
	_jsii_.Get(
		j,
		"imageResponseCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) ImageResponseCardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageResponseCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) PlainTextMessage() LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsPlainTextMessageOutputReference {
	var returns LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsPlainTextMessageOutputReference
	_jsii_.Get(
		j,
		"plainTextMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) PlainTextMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"plainTextMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) SsmlMessage() LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsSsmlMessageOutputReference {
	var returns LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsSsmlMessageOutputReference
	_jsii_.Get(
		j,
		"ssmlMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) SsmlMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssmlMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference {
	_init_.Initialize()

	if err := validateNewLexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference{}

	_jsii_.Create(
		"awscc.lexBot.LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference_Override(l LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lexBot.LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) PutCustomPayload(value *LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsCustomPayload) {
	if err := l.validatePutCustomPayloadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCustomPayload",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) PutImageResponseCard(value *LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCard) {
	if err := l.validatePutImageResponseCardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putImageResponseCard",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) PutPlainTextMessage(value *LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsPlainTextMessage) {
	if err := l.validatePutPlainTextMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPlainTextMessage",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) PutSsmlMessage(value *LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsSsmlMessage) {
	if err := l.validatePutSsmlMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSsmlMessage",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) ResetCustomPayload() {
	_jsii_.InvokeVoid(
		l,
		"resetCustomPayload",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) ResetImageResponseCard() {
	_jsii_.InvokeVoid(
		l,
		"resetImageResponseCard",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) ResetPlainTextMessage() {
	_jsii_.InvokeVoid(
		l,
		"resetPlainTextMessage",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) ResetSsmlMessage() {
	_jsii_.InvokeVoid(
		l,
		"resetSsmlMessage",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

