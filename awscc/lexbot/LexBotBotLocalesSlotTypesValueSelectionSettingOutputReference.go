package lexbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lexbot/internal"
)

type LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference interface {
	cdktf.ComplexObject
	AdvancedRecognitionSetting() LexBotBotLocalesSlotTypesValueSelectionSettingAdvancedRecognitionSettingOutputReference
	AdvancedRecognitionSettingInput() interface{}
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
	RegexFilter() LexBotBotLocalesSlotTypesValueSelectionSettingRegexFilterOutputReference
	RegexFilterInput() interface{}
	ResolutionStrategy() *string
	SetResolutionStrategy(val *string)
	ResolutionStrategyInput() *string
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
	PutAdvancedRecognitionSetting(value *LexBotBotLocalesSlotTypesValueSelectionSettingAdvancedRecognitionSetting)
	PutRegexFilter(value *LexBotBotLocalesSlotTypesValueSelectionSettingRegexFilter)
	ResetAdvancedRecognitionSetting()
	ResetRegexFilter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference
type jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) AdvancedRecognitionSetting() LexBotBotLocalesSlotTypesValueSelectionSettingAdvancedRecognitionSettingOutputReference {
	var returns LexBotBotLocalesSlotTypesValueSelectionSettingAdvancedRecognitionSettingOutputReference
	_jsii_.Get(
		j,
		"advancedRecognitionSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) AdvancedRecognitionSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedRecognitionSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) RegexFilter() LexBotBotLocalesSlotTypesValueSelectionSettingRegexFilterOutputReference {
	var returns LexBotBotLocalesSlotTypesValueSelectionSettingRegexFilterOutputReference
	_jsii_.Get(
		j,
		"regexFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) RegexFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regexFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) ResolutionStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolutionStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) ResolutionStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolutionStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLexBotBotLocalesSlotTypesValueSelectionSettingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference {
	_init_.Initialize()

	if err := validateNewLexBotBotLocalesSlotTypesValueSelectionSettingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference{}

	_jsii_.Create(
		"awscc.lexBot.LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLexBotBotLocalesSlotTypesValueSelectionSettingOutputReference_Override(l LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lexBot.LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference)SetResolutionStrategy(val *string) {
	if err := j.validateSetResolutionStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolutionStrategy",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) PutAdvancedRecognitionSetting(value *LexBotBotLocalesSlotTypesValueSelectionSettingAdvancedRecognitionSetting) {
	if err := l.validatePutAdvancedRecognitionSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAdvancedRecognitionSetting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) PutRegexFilter(value *LexBotBotLocalesSlotTypesValueSelectionSettingRegexFilter) {
	if err := l.validatePutRegexFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRegexFilter",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) ResetAdvancedRecognitionSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetAdvancedRecognitionSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) ResetRegexFilter() {
	_jsii_.InvokeVoid(
		l,
		"resetRegexFilter",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LexBotBotLocalesSlotTypesValueSelectionSettingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

