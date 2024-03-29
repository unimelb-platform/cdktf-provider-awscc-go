package lexbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lexbot/internal"
)

type LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList interface {
	cdktf.ComplexList
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
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList
type jsiiProxy_LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList {
	_init_.Initialize()

	if err := validateNewLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList{}

	_jsii_.Create(
		"awscc.lexBot.LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList_Override(l LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lexBot.LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		l,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (l *jsiiProxy_LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList) Get(index *float64) LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference {
	if err := l.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference

	_jsii_.Invoke(
		l,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListVariationsImageResponseCardButtonsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

