package eventsconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/eventsconnection/internal"
)

type EventsConnectionAuthParametersOutputReference interface {
	cdktf.ComplexObject
	ApiKeyAuthParameters() EventsConnectionAuthParametersApiKeyAuthParametersOutputReference
	ApiKeyAuthParametersInput() interface{}
	BasicAuthParameters() EventsConnectionAuthParametersBasicAuthParametersOutputReference
	BasicAuthParametersInput() interface{}
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
	ConnectivityParameters() EventsConnectionAuthParametersConnectivityParametersOutputReference
	ConnectivityParametersInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	InvocationHttpParameters() EventsConnectionAuthParametersInvocationHttpParametersOutputReference
	InvocationHttpParametersInput() interface{}
	OAuthParameters() EventsConnectionAuthParametersOAuthParametersOutputReference
	OAuthParametersInput() interface{}
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
	PutApiKeyAuthParameters(value *EventsConnectionAuthParametersApiKeyAuthParameters)
	PutBasicAuthParameters(value *EventsConnectionAuthParametersBasicAuthParameters)
	PutConnectivityParameters(value *EventsConnectionAuthParametersConnectivityParameters)
	PutInvocationHttpParameters(value *EventsConnectionAuthParametersInvocationHttpParameters)
	PutOAuthParameters(value *EventsConnectionAuthParametersOAuthParameters)
	ResetApiKeyAuthParameters()
	ResetBasicAuthParameters()
	ResetConnectivityParameters()
	ResetInvocationHttpParameters()
	ResetOAuthParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventsConnectionAuthParametersOutputReference
type jsiiProxy_EventsConnectionAuthParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) ApiKeyAuthParameters() EventsConnectionAuthParametersApiKeyAuthParametersOutputReference {
	var returns EventsConnectionAuthParametersApiKeyAuthParametersOutputReference
	_jsii_.Get(
		j,
		"apiKeyAuthParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) ApiKeyAuthParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyAuthParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) BasicAuthParameters() EventsConnectionAuthParametersBasicAuthParametersOutputReference {
	var returns EventsConnectionAuthParametersBasicAuthParametersOutputReference
	_jsii_.Get(
		j,
		"basicAuthParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) BasicAuthParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basicAuthParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) ConnectivityParameters() EventsConnectionAuthParametersConnectivityParametersOutputReference {
	var returns EventsConnectionAuthParametersConnectivityParametersOutputReference
	_jsii_.Get(
		j,
		"connectivityParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) ConnectivityParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectivityParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) InvocationHttpParameters() EventsConnectionAuthParametersInvocationHttpParametersOutputReference {
	var returns EventsConnectionAuthParametersInvocationHttpParametersOutputReference
	_jsii_.Get(
		j,
		"invocationHttpParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) InvocationHttpParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invocationHttpParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) OAuthParameters() EventsConnectionAuthParametersOAuthParametersOutputReference {
	var returns EventsConnectionAuthParametersOAuthParametersOutputReference
	_jsii_.Get(
		j,
		"oAuthParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) OAuthParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oAuthParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventsConnectionAuthParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventsConnectionAuthParametersOutputReference {
	_init_.Initialize()

	if err := validateNewEventsConnectionAuthParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventsConnectionAuthParametersOutputReference{}

	_jsii_.Create(
		"awscc.eventsConnection.EventsConnectionAuthParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventsConnectionAuthParametersOutputReference_Override(e EventsConnectionAuthParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.eventsConnection.EventsConnectionAuthParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) PutApiKeyAuthParameters(value *EventsConnectionAuthParametersApiKeyAuthParameters) {
	if err := e.validatePutApiKeyAuthParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putApiKeyAuthParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) PutBasicAuthParameters(value *EventsConnectionAuthParametersBasicAuthParameters) {
	if err := e.validatePutBasicAuthParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBasicAuthParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) PutConnectivityParameters(value *EventsConnectionAuthParametersConnectivityParameters) {
	if err := e.validatePutConnectivityParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putConnectivityParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) PutInvocationHttpParameters(value *EventsConnectionAuthParametersInvocationHttpParameters) {
	if err := e.validatePutInvocationHttpParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInvocationHttpParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) PutOAuthParameters(value *EventsConnectionAuthParametersOAuthParameters) {
	if err := e.validatePutOAuthParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOAuthParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) ResetApiKeyAuthParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetApiKeyAuthParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) ResetBasicAuthParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetBasicAuthParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) ResetConnectivityParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetConnectivityParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) ResetInvocationHttpParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetInvocationHttpParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) ResetOAuthParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetOAuthParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

