package eventsconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/eventsconnection/internal"
)

type EventsConnectionAuthParametersOAuthParametersOutputReference interface {
	cdktf.ComplexObject
	AuthorizationEndpoint() *string
	SetAuthorizationEndpoint(val *string)
	AuthorizationEndpointInput() *string
	ClientParameters() EventsConnectionAuthParametersOAuthParametersClientParametersOutputReference
	ClientParametersInput() interface{}
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
	HttpMethod() *string
	SetHttpMethod(val *string)
	HttpMethodInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OAuthHttpParameters() EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference
	OAuthHttpParametersInput() interface{}
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
	PutClientParameters(value *EventsConnectionAuthParametersOAuthParametersClientParameters)
	PutOAuthHttpParameters(value *EventsConnectionAuthParametersOAuthParametersOAuthHttpParameters)
	ResetAuthorizationEndpoint()
	ResetClientParameters()
	ResetHttpMethod()
	ResetOAuthHttpParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventsConnectionAuthParametersOAuthParametersOutputReference
type jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) AuthorizationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) AuthorizationEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) ClientParameters() EventsConnectionAuthParametersOAuthParametersClientParametersOutputReference {
	var returns EventsConnectionAuthParametersOAuthParametersClientParametersOutputReference
	_jsii_.Get(
		j,
		"clientParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) ClientParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) OAuthHttpParameters() EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference {
	var returns EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference
	_jsii_.Get(
		j,
		"oAuthHttpParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) OAuthHttpParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oAuthHttpParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventsConnectionAuthParametersOAuthParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventsConnectionAuthParametersOAuthParametersOutputReference {
	_init_.Initialize()

	if err := validateNewEventsConnectionAuthParametersOAuthParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference{}

	_jsii_.Create(
		"awscc.eventsConnection.EventsConnectionAuthParametersOAuthParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventsConnectionAuthParametersOAuthParametersOutputReference_Override(e EventsConnectionAuthParametersOAuthParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.eventsConnection.EventsConnectionAuthParametersOAuthParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference)SetAuthorizationEndpoint(val *string) {
	if err := j.validateSetAuthorizationEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationEndpoint",
		val,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) PutClientParameters(value *EventsConnectionAuthParametersOAuthParametersClientParameters) {
	if err := e.validatePutClientParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putClientParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) PutOAuthHttpParameters(value *EventsConnectionAuthParametersOAuthParametersOAuthHttpParameters) {
	if err := e.validatePutOAuthHttpParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOAuthHttpParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) ResetAuthorizationEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetAuthorizationEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) ResetClientParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetClientParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		e,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) ResetOAuthHttpParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetOAuthHttpParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

