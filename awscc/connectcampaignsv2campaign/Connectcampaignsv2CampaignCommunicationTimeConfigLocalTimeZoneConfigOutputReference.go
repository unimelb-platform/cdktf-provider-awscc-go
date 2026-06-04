package connectcampaignsv2campaign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/connectcampaignsv2campaign/internal"
)

type Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference interface {
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
	DefaultTimeZone() *string
	SetDefaultTimeZone(val *string)
	DefaultTimeZoneInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LocalTimeZoneDetection() *[]*string
	SetLocalTimeZoneDetection(val *[]*string)
	LocalTimeZoneDetectionInput() *[]*string
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
	ResetDefaultTimeZone()
	ResetLocalTimeZoneDetection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference
type jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) DefaultTimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTimeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) DefaultTimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTimeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) LocalTimeZoneDetection() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localTimeZoneDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) LocalTimeZoneDetectionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localTimeZoneDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference {
	_init_.Initialize()

	if err := validateNewConnectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference{}

	_jsii_.Create(
		"awscc.connectcampaignsv2Campaign.Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference_Override(c Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.connectcampaignsv2Campaign.Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference)SetDefaultTimeZone(val *string) {
	if err := j.validateSetDefaultTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTimeZone",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference)SetLocalTimeZoneDetection(val *[]*string) {
	if err := j.validateSetLocalTimeZoneDetectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localTimeZoneDetection",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) ResetDefaultTimeZone() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultTimeZone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) ResetLocalTimeZoneDetection() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalTimeZoneDetection",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

