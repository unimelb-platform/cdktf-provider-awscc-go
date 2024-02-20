package cloudfrontdistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cloudfrontdistribution/internal"
)

type CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference interface {
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
	EventType() *string
	SetEventType(val *string)
	EventTypeInput() *string
	// Experimental.
	Fqn() *string
	IncludeBody() interface{}
	SetIncludeBody(val interface{})
	IncludeBodyInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LambdaFunctionArn() *string
	SetLambdaFunctionArn(val *string)
	LambdaFunctionArnInput() *string
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
	ResetEventType()
	ResetIncludeBody()
	ResetLambdaFunctionArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference
type jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) EventType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) EventTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) IncludeBody() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) IncludeBodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) LambdaFunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFunctionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) LambdaFunctionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFunctionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference{}

	_jsii_.Create(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference_Override(c CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference)SetEventType(val *string) {
	if err := j.validateSetEventTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventType",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference)SetIncludeBody(val interface{}) {
	if err := j.validateSetIncludeBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeBody",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference)SetLambdaFunctionArn(val *string) {
	if err := j.validateSetLambdaFunctionArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaFunctionArn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) ResetEventType() {
	_jsii_.InvokeVoid(
		c,
		"resetEventType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) ResetIncludeBody() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeBody",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) ResetLambdaFunctionArn() {
	_jsii_.InvokeVoid(
		c,
		"resetLambdaFunctionArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsLambdaFunctionAssociationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

