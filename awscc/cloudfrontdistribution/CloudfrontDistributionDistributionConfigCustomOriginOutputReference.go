package cloudfrontdistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cloudfrontdistribution/internal"
)

type CloudfrontDistributionDistributionConfigCustomOriginOutputReference interface {
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
	DnsName() *string
	SetDnsName(val *string)
	DnsNameInput() *string
	// Experimental.
	Fqn() *string
	HttpPort() *float64
	SetHttpPort(val *float64)
	HttpPortInput() *float64
	HttpsPort() *float64
	SetHttpsPort(val *float64)
	HttpsPortInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OriginProtocolPolicy() *string
	SetOriginProtocolPolicy(val *string)
	OriginProtocolPolicyInput() *string
	OriginSslProtocols() *[]*string
	SetOriginSslProtocols(val *[]*string)
	OriginSslProtocolsInput() *[]*string
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
	ResetHttpPort()
	ResetHttpsPort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontDistributionDistributionConfigCustomOriginOutputReference
type jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) DnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) HttpPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) HttpPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) HttpsPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpsPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) HttpsPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpsPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) OriginProtocolPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originProtocolPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) OriginProtocolPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originProtocolPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) OriginSslProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originSslProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) OriginSslProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originSslProtocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudfrontDistributionDistributionConfigCustomOriginOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudfrontDistributionDistributionConfigCustomOriginOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontDistributionDistributionConfigCustomOriginOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference{}

	_jsii_.Create(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCustomOriginOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionDistributionConfigCustomOriginOutputReference_Override(c CloudfrontDistributionDistributionConfigCustomOriginOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCustomOriginOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference)SetDnsName(val *string) {
	if err := j.validateSetDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsName",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference)SetHttpPort(val *float64) {
	if err := j.validateSetHttpPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpPort",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference)SetHttpsPort(val *float64) {
	if err := j.validateSetHttpsPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsPort",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference)SetOriginProtocolPolicy(val *string) {
	if err := j.validateSetOriginProtocolPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originProtocolPolicy",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference)SetOriginSslProtocols(val *[]*string) {
	if err := j.validateSetOriginSslProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originSslProtocols",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) ResetHttpPort() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) ResetHttpsPort() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpsPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCustomOriginOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

