package lexbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lexbot/internal"
)

type LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap interface {
	cdktf.ComplexMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	Get(key *string) LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationOutputReference
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap
type jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap struct {
	internal.Type__cdktfComplexMap
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap {
	_init_.Initialize()

	if err := validateNewLexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap{}

	_jsii_.Create(
		"awscc.lexBot.LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap_Override(l LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lexBot.LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap) Get(key *string) LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationOutputReference {
	if err := l.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationOutputReference

	_jsii_.Invoke(
		l,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

