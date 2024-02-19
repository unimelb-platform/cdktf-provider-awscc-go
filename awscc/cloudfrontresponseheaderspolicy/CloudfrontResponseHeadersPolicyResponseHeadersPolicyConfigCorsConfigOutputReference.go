package cloudfrontresponseheaderspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cloudfrontresponseheaderspolicy/internal"
)

type CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference interface {
	cdktf.ComplexObject
	AccessControlAllowCredentials() interface{}
	SetAccessControlAllowCredentials(val interface{})
	AccessControlAllowCredentialsInput() interface{}
	AccessControlAllowHeaders() CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowHeadersOutputReference
	AccessControlAllowHeadersInput() *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowHeaders
	AccessControlAllowMethods() CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowMethodsOutputReference
	AccessControlAllowMethodsInput() *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowMethods
	AccessControlAllowOrigins() CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowOriginsOutputReference
	AccessControlAllowOriginsInput() *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowOrigins
	AccessControlExposeHeaders() CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlExposeHeadersOutputReference
	AccessControlExposeHeadersInput() interface{}
	AccessControlMaxAgeSec() *float64
	SetAccessControlMaxAgeSec(val *float64)
	AccessControlMaxAgeSecInput() *float64
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OriginOverride() interface{}
	SetOriginOverride(val interface{})
	OriginOverrideInput() interface{}
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
	PutAccessControlAllowHeaders(value *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowHeaders)
	PutAccessControlAllowMethods(value *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowMethods)
	PutAccessControlAllowOrigins(value *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowOrigins)
	PutAccessControlExposeHeaders(value *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlExposeHeaders)
	ResetAccessControlExposeHeaders()
	ResetAccessControlMaxAgeSec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) AccessControlAllowCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlAllowCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) AccessControlAllowCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlAllowCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) AccessControlAllowHeaders() CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowHeadersOutputReference {
	var returns CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowHeadersOutputReference
	_jsii_.Get(
		j,
		"accessControlAllowHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) AccessControlAllowHeadersInput() *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowHeaders {
	var returns *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowHeaders
	_jsii_.Get(
		j,
		"accessControlAllowHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) AccessControlAllowMethods() CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowMethodsOutputReference {
	var returns CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowMethodsOutputReference
	_jsii_.Get(
		j,
		"accessControlAllowMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) AccessControlAllowMethodsInput() *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowMethods {
	var returns *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowMethods
	_jsii_.Get(
		j,
		"accessControlAllowMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) AccessControlAllowOrigins() CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowOriginsOutputReference {
	var returns CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowOriginsOutputReference
	_jsii_.Get(
		j,
		"accessControlAllowOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) AccessControlAllowOriginsInput() *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowOrigins {
	var returns *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowOrigins
	_jsii_.Get(
		j,
		"accessControlAllowOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) AccessControlExposeHeaders() CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlExposeHeadersOutputReference {
	var returns CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlExposeHeadersOutputReference
	_jsii_.Get(
		j,
		"accessControlExposeHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) AccessControlExposeHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlExposeHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) AccessControlMaxAgeSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) AccessControlMaxAgeSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) OriginOverride() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) OriginOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference{}

	_jsii_.Create(
		"awscc.cloudfrontResponseHeadersPolicy.CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference_Override(c CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cloudfrontResponseHeadersPolicy.CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference)SetAccessControlAllowCredentials(val interface{}) {
	if err := j.validateSetAccessControlAllowCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessControlAllowCredentials",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference)SetAccessControlMaxAgeSec(val *float64) {
	if err := j.validateSetAccessControlMaxAgeSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessControlMaxAgeSec",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference)SetOriginOverride(val interface{}) {
	if err := j.validateSetOriginOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originOverride",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) PutAccessControlAllowHeaders(value *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowHeaders) {
	if err := c.validatePutAccessControlAllowHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAccessControlAllowHeaders",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) PutAccessControlAllowMethods(value *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowMethods) {
	if err := c.validatePutAccessControlAllowMethodsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAccessControlAllowMethods",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) PutAccessControlAllowOrigins(value *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowOrigins) {
	if err := c.validatePutAccessControlAllowOriginsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAccessControlAllowOrigins",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) PutAccessControlExposeHeaders(value *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlExposeHeaders) {
	if err := c.validatePutAccessControlExposeHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAccessControlExposeHeaders",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) ResetAccessControlExposeHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessControlExposeHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) ResetAccessControlMaxAgeSec() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessControlMaxAgeSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

