package lexbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lexbot/internal"
)

type LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference interface {
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
	LambdaCodeHook() LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationLambdaCodeHookOutputReference
	LambdaCodeHookInput() *LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationLambdaCodeHook
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
	PutLambdaCodeHook(value *LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationLambdaCodeHook)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference
type jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) LambdaCodeHook() LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationLambdaCodeHookOutputReference {
	var returns LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationLambdaCodeHookOutputReference
	_jsii_.Get(
		j,
		"lambdaCodeHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) LambdaCodeHookInput() *LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationLambdaCodeHook {
	var returns *LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationLambdaCodeHook
	_jsii_.Get(
		j,
		"lambdaCodeHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewLexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference{}

	_jsii_.Create(
		"awscc.lexBot.LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference_Override(l LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lexBot.LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) PutLambdaCodeHook(value *LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationLambdaCodeHook) {
	if err := l.validatePutLambdaCodeHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLambdaCodeHook",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

