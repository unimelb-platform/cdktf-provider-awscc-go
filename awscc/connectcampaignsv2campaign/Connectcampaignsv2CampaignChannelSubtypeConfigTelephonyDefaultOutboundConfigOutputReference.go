package connectcampaignsv2campaign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/connectcampaignsv2campaign/internal"
)

type Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference interface {
	cdktf.ComplexObject
	AnswerMachineDetectionConfig() Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigAnswerMachineDetectionConfigOutputReference
	AnswerMachineDetectionConfigInput() interface{}
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
	ConnectContactFlowId() *string
	SetConnectContactFlowId(val *string)
	ConnectContactFlowIdInput() *string
	ConnectSourcePhoneNumber() *string
	SetConnectSourcePhoneNumber(val *string)
	ConnectSourcePhoneNumberInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutAnswerMachineDetectionConfig(value *Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigAnswerMachineDetectionConfig)
	ResetAnswerMachineDetectionConfig()
	ResetConnectContactFlowId()
	ResetConnectSourcePhoneNumber()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference
type jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) AnswerMachineDetectionConfig() Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigAnswerMachineDetectionConfigOutputReference {
	var returns Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigAnswerMachineDetectionConfigOutputReference
	_jsii_.Get(
		j,
		"answerMachineDetectionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) AnswerMachineDetectionConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"answerMachineDetectionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ConnectContactFlowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectContactFlowId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ConnectContactFlowIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectContactFlowIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ConnectSourcePhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectSourcePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ConnectSourcePhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectSourcePhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference {
	_init_.Initialize()

	if err := validateNewConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference{}

	_jsii_.Create(
		"awscc.connectcampaignsv2Campaign.Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference_Override(c Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.connectcampaignsv2Campaign.Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference)SetConnectContactFlowId(val *string) {
	if err := j.validateSetConnectContactFlowIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectContactFlowId",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference)SetConnectSourcePhoneNumber(val *string) {
	if err := j.validateSetConnectSourcePhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectSourcePhoneNumber",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) PutAnswerMachineDetectionConfig(value *Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigAnswerMachineDetectionConfig) {
	if err := c.validatePutAnswerMachineDetectionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAnswerMachineDetectionConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ResetAnswerMachineDetectionConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAnswerMachineDetectionConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ResetConnectContactFlowId() {
	_jsii_.InvokeVoid(
		c,
		"resetConnectContactFlowId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ResetConnectSourcePhoneNumber() {
	_jsii_.InvokeVoid(
		c,
		"resetConnectSourcePhoneNumber",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

