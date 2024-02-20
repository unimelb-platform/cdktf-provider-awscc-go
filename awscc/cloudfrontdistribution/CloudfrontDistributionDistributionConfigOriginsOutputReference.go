package cloudfrontdistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cloudfrontdistribution/internal"
)

type CloudfrontDistributionDistributionConfigOriginsOutputReference interface {
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
	ConnectionAttempts() *float64
	SetConnectionAttempts(val *float64)
	ConnectionAttemptsInput() *float64
	ConnectionTimeout() *float64
	SetConnectionTimeout(val *float64)
	ConnectionTimeoutInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomOriginConfig() CloudfrontDistributionDistributionConfigOriginsCustomOriginConfigOutputReference
	CustomOriginConfigInput() interface{}
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OriginAccessControlId() *string
	SetOriginAccessControlId(val *string)
	OriginAccessControlIdInput() *string
	OriginCustomHeaders() CloudfrontDistributionDistributionConfigOriginsOriginCustomHeadersList
	OriginCustomHeadersInput() interface{}
	OriginPath() *string
	SetOriginPath(val *string)
	OriginPathInput() *string
	OriginShield() CloudfrontDistributionDistributionConfigOriginsOriginShieldOutputReference
	OriginShieldInput() interface{}
	S3OriginConfig() CloudfrontDistributionDistributionConfigOriginsS3OriginConfigOutputReference
	S3OriginConfigInput() interface{}
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
	PutCustomOriginConfig(value *CloudfrontDistributionDistributionConfigOriginsCustomOriginConfig)
	PutOriginCustomHeaders(value interface{})
	PutOriginShield(value *CloudfrontDistributionDistributionConfigOriginsOriginShield)
	PutS3OriginConfig(value *CloudfrontDistributionDistributionConfigOriginsS3OriginConfig)
	ResetConnectionAttempts()
	ResetConnectionTimeout()
	ResetCustomOriginConfig()
	ResetOriginAccessControlId()
	ResetOriginCustomHeaders()
	ResetOriginPath()
	ResetOriginShield()
	ResetS3OriginConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontDistributionDistributionConfigOriginsOutputReference
type jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) ConnectionAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) ConnectionAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) ConnectionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) ConnectionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) CustomOriginConfig() CloudfrontDistributionDistributionConfigOriginsCustomOriginConfigOutputReference {
	var returns CloudfrontDistributionDistributionConfigOriginsCustomOriginConfigOutputReference
	_jsii_.Get(
		j,
		"customOriginConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) CustomOriginConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customOriginConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) OriginAccessControlId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessControlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) OriginAccessControlIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessControlIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) OriginCustomHeaders() CloudfrontDistributionDistributionConfigOriginsOriginCustomHeadersList {
	var returns CloudfrontDistributionDistributionConfigOriginsOriginCustomHeadersList
	_jsii_.Get(
		j,
		"originCustomHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) OriginCustomHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originCustomHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) OriginPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) OriginPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) OriginShield() CloudfrontDistributionDistributionConfigOriginsOriginShieldOutputReference {
	var returns CloudfrontDistributionDistributionConfigOriginsOriginShieldOutputReference
	_jsii_.Get(
		j,
		"originShield",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) OriginShieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originShieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) S3OriginConfig() CloudfrontDistributionDistributionConfigOriginsS3OriginConfigOutputReference {
	var returns CloudfrontDistributionDistributionConfigOriginsS3OriginConfigOutputReference
	_jsii_.Get(
		j,
		"s3OriginConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) S3OriginConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3OriginConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudfrontDistributionDistributionConfigOriginsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudfrontDistributionDistributionConfigOriginsOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontDistributionDistributionConfigOriginsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference{}

	_jsii_.Create(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionDistributionConfigOriginsOutputReference_Override(c CloudfrontDistributionDistributionConfigOriginsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOriginsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference)SetConnectionAttempts(val *float64) {
	if err := j.validateSetConnectionAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionAttempts",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference)SetConnectionTimeout(val *float64) {
	if err := j.validateSetConnectionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionTimeout",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference)SetOriginAccessControlId(val *string) {
	if err := j.validateSetOriginAccessControlIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originAccessControlId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference)SetOriginPath(val *string) {
	if err := j.validateSetOriginPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originPath",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) PutCustomOriginConfig(value *CloudfrontDistributionDistributionConfigOriginsCustomOriginConfig) {
	if err := c.validatePutCustomOriginConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomOriginConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) PutOriginCustomHeaders(value interface{}) {
	if err := c.validatePutOriginCustomHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOriginCustomHeaders",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) PutOriginShield(value *CloudfrontDistributionDistributionConfigOriginsOriginShield) {
	if err := c.validatePutOriginShieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOriginShield",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) PutS3OriginConfig(value *CloudfrontDistributionDistributionConfigOriginsS3OriginConfig) {
	if err := c.validatePutS3OriginConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putS3OriginConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) ResetConnectionAttempts() {
	_jsii_.InvokeVoid(
		c,
		"resetConnectionAttempts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) ResetConnectionTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetConnectionTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) ResetCustomOriginConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomOriginConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) ResetOriginAccessControlId() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginAccessControlId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) ResetOriginCustomHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginCustomHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) ResetOriginPath() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) ResetOriginShield() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginShield",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) ResetS3OriginConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetS3OriginConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOriginsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

