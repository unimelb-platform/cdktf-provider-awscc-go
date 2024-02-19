package lexbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lexbot/internal"
)

type LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference interface {
	cdktf.ComplexObject
	AllowInterrupt() interface{}
	SetAllowInterrupt(val interface{})
	AllowInterruptInput() interface{}
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
	InternalValue() *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecification
	SetInternalValue(val *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecification)
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	MessageGroupsList() LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListStructList
	MessageGroupsListInput() interface{}
	MessageSelectionStrategy() *string
	SetMessageSelectionStrategy(val *string)
	MessageSelectionStrategyInput() *string
	PromptAttemptsSpecification() LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap
	PromptAttemptsSpecificationInput() interface{}
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
	PutMessageGroupsList(value interface{})
	PutPromptAttemptsSpecification(value interface{})
	ResetAllowInterrupt()
	ResetMessageSelectionStrategy()
	ResetPromptAttemptsSpecification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference
type jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) AllowInterrupt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInterrupt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) AllowInterruptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInterruptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) InternalValue() *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecification {
	var returns *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecification
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) MessageGroupsList() LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListStructList {
	var returns LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationMessageGroupsListStructList
	_jsii_.Get(
		j,
		"messageGroupsList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) MessageGroupsListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageGroupsListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) MessageSelectionStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageSelectionStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) MessageSelectionStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageSelectionStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) PromptAttemptsSpecification() LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap {
	var returns LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap
	_jsii_.Get(
		j,
		"promptAttemptsSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) PromptAttemptsSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptAttemptsSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewLexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference{}

	_jsii_.Create(
		"awscc.lexBot.LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference_Override(l LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lexBot.LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference)SetAllowInterrupt(val interface{}) {
	if err := j.validateSetAllowInterruptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowInterrupt",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference)SetInternalValue(val *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecification) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference)SetMessageSelectionStrategy(val *string) {
	if err := j.validateSetMessageSelectionStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageSelectionStrategy",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) PutMessageGroupsList(value interface{}) {
	if err := l.validatePutMessageGroupsListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMessageGroupsList",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) PutPromptAttemptsSpecification(value interface{}) {
	if err := l.validatePutPromptAttemptsSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPromptAttemptsSpecification",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) ResetAllowInterrupt() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowInterrupt",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) ResetMessageSelectionStrategy() {
	_jsii_.InvokeVoid(
		l,
		"resetMessageSelectionStrategy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) ResetPromptAttemptsSpecification() {
	_jsii_.InvokeVoid(
		l,
		"resetPromptAttemptsSpecification",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

