package eventsconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/eventsconnection/internal"
)

type EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference interface {
	cdktf.ComplexObject
	BodyParameters() EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersBodyParametersList
	BodyParametersInput() interface{}
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
	HeaderParameters() EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersHeaderParametersList
	HeaderParametersInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	QueryStringParameters() EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersQueryStringParametersList
	QueryStringParametersInput() interface{}
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
	PutBodyParameters(value interface{})
	PutHeaderParameters(value interface{})
	PutQueryStringParameters(value interface{})
	ResetBodyParameters()
	ResetHeaderParameters()
	ResetQueryStringParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference
type jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) BodyParameters() EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersBodyParametersList {
	var returns EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersBodyParametersList
	_jsii_.Get(
		j,
		"bodyParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) BodyParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bodyParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) HeaderParameters() EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersHeaderParametersList {
	var returns EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersHeaderParametersList
	_jsii_.Get(
		j,
		"headerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) HeaderParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) QueryStringParameters() EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersQueryStringParametersList {
	var returns EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersQueryStringParametersList
	_jsii_.Get(
		j,
		"queryStringParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) QueryStringParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference {
	_init_.Initialize()

	if err := validateNewEventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference{}

	_jsii_.Create(
		"awscc.eventsConnection.EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference_Override(e EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.eventsConnection.EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) PutBodyParameters(value interface{}) {
	if err := e.validatePutBodyParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBodyParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) PutHeaderParameters(value interface{}) {
	if err := e.validatePutHeaderParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHeaderParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) PutQueryStringParameters(value interface{}) {
	if err := e.validatePutQueryStringParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putQueryStringParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) ResetBodyParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetBodyParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) ResetHeaderParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetHeaderParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) ResetQueryStringParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetQueryStringParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EventsConnectionAuthParametersOAuthParametersOAuthHttpParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

