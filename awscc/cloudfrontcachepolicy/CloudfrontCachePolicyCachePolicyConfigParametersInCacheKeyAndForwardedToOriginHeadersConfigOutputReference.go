package cloudfrontcachepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cloudfrontcachepolicy/internal"
)

type CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	HeaderBehavior() *string
	SetHeaderBehavior(val *string)
	HeaderBehaviorInput() *string
	Headers() *[]*string
	SetHeaders(val *[]*string)
	HeadersInput() *[]*string
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
	ResetHeaders()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference
type jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) HeaderBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) HeaderBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) Headers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) HeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference{}

	_jsii_.Create(
		"awscc.cloudfrontCachePolicy.CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference_Override(c CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cloudfrontCachePolicy.CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference)SetHeaderBehavior(val *string) {
	if err := j.validateSetHeaderBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerBehavior",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference)SetHeaders(val *[]*string) {
	if err := j.validateSetHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

