package connectinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/connectinstance/internal"
)

type ConnectInstanceAttributesOutputReference interface {
	cdktf.ComplexObject
	AutoResolveBestVoices() interface{}
	SetAutoResolveBestVoices(val interface{})
	AutoResolveBestVoicesInput() interface{}
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
	ContactflowLogs() interface{}
	SetContactflowLogs(val interface{})
	ContactflowLogsInput() interface{}
	ContactLens() interface{}
	SetContactLens(val interface{})
	ContactLensInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EarlyMedia() interface{}
	SetEarlyMedia(val interface{})
	EarlyMediaInput() interface{}
	// Experimental.
	Fqn() *string
	InboundCalls() interface{}
	SetInboundCalls(val interface{})
	InboundCallsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OutboundCalls() interface{}
	SetOutboundCalls(val interface{})
	OutboundCallsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseCustomTtsVoices() interface{}
	SetUseCustomTtsVoices(val interface{})
	UseCustomTtsVoicesInput() interface{}
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
	ResetAutoResolveBestVoices()
	ResetContactflowLogs()
	ResetContactLens()
	ResetEarlyMedia()
	ResetUseCustomTtsVoices()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectInstanceAttributesOutputReference
type jsiiProxy_ConnectInstanceAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) AutoResolveBestVoices() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoResolveBestVoices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) AutoResolveBestVoicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoResolveBestVoicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) ContactflowLogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contactflowLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) ContactflowLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contactflowLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) ContactLens() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contactLens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) ContactLensInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contactLensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) EarlyMedia() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"earlyMedia",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) EarlyMediaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"earlyMediaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) InboundCalls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inboundCalls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) InboundCallsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inboundCallsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) OutboundCalls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outboundCalls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) OutboundCallsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outboundCallsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) UseCustomTtsVoices() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCustomTtsVoices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference) UseCustomTtsVoicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCustomTtsVoicesInput",
		&returns,
	)
	return returns
}


func NewConnectInstanceAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConnectInstanceAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewConnectInstanceAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectInstanceAttributesOutputReference{}

	_jsii_.Create(
		"awscc.connectInstance.ConnectInstanceAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectInstanceAttributesOutputReference_Override(c ConnectInstanceAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.connectInstance.ConnectInstanceAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference)SetAutoResolveBestVoices(val interface{}) {
	if err := j.validateSetAutoResolveBestVoicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoResolveBestVoices",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference)SetContactflowLogs(val interface{}) {
	if err := j.validateSetContactflowLogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactflowLogs",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference)SetContactLens(val interface{}) {
	if err := j.validateSetContactLensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactLens",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference)SetEarlyMedia(val interface{}) {
	if err := j.validateSetEarlyMediaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"earlyMedia",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference)SetInboundCalls(val interface{}) {
	if err := j.validateSetInboundCallsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inboundCalls",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference)SetOutboundCalls(val interface{}) {
	if err := j.validateSetOutboundCallsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundCalls",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceAttributesOutputReference)SetUseCustomTtsVoices(val interface{}) {
	if err := j.validateSetUseCustomTtsVoicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCustomTtsVoices",
		val,
	)
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) ResetAutoResolveBestVoices() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoResolveBestVoices",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) ResetContactflowLogs() {
	_jsii_.InvokeVoid(
		c,
		"resetContactflowLogs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) ResetContactLens() {
	_jsii_.InvokeVoid(
		c,
		"resetContactLens",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) ResetEarlyMedia() {
	_jsii_.InvokeVoid(
		c,
		"resetEarlyMedia",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) ResetUseCustomTtsVoices() {
	_jsii_.InvokeVoid(
		c,
		"resetUseCustomTtsVoices",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

