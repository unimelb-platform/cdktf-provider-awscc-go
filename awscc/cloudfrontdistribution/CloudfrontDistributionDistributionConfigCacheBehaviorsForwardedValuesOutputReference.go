package cloudfrontdistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cloudfrontdistribution/internal"
)

type CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference interface {
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
	Cookies() CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference
	CookiesInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() *[]*string
	SetHeaders(val *[]*string)
	HeadersInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	QueryString() interface{}
	SetQueryString(val interface{})
	QueryStringCacheKeys() *[]*string
	SetQueryStringCacheKeys(val *[]*string)
	QueryStringCacheKeysInput() *[]*string
	QueryStringInput() interface{}
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
	PutCookies(value *CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookies)
	ResetCookies()
	ResetHeaders()
	ResetQueryStringCacheKeys()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference
type jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) Cookies() CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference {
	var returns CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) CookiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) Headers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) HeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) QueryString() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) QueryStringCacheKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringCacheKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) QueryStringCacheKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringCacheKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference{}

	_jsii_.Create(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference_Override(c CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference)SetHeaders(val *[]*string) {
	if err := j.validateSetHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference)SetQueryString(val interface{}) {
	if err := j.validateSetQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference)SetQueryStringCacheKeys(val *[]*string) {
	if err := j.validateSetQueryStringCacheKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringCacheKeys",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) PutCookies(value *CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookies) {
	if err := c.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCookies",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		c,
		"resetCookies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) ResetQueryStringCacheKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStringCacheKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

