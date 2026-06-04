package dataawsccconnectcampaignsv2campaign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccconnectcampaignsv2campaign/internal"
)

type DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference interface {
	cdktf.ComplexObject
	AnswerMachineDetectionConfig() DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigAnswerMachineDetectionConfigOutputReference
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
	ConnectSourcePhoneNumber() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfig
	SetInternalValue(val *DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfig)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference
type jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) AnswerMachineDetectionConfig() DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigAnswerMachineDetectionConfigOutputReference {
	var returns DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigAnswerMachineDetectionConfigOutputReference
	_jsii_.Get(
		j,
		"answerMachineDetectionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ConnectContactFlowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectContactFlowId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ConnectSourcePhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectSourcePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) InternalValue() *DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfig {
	var returns *DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccConnectcampaignsv2Campaign.DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference_Override(d DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccConnectcampaignsv2Campaign.DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference)SetInternalValue(val *DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

