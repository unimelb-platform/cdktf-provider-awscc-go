package cleanroomsconfiguredtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cleanroomsconfiguredtable/internal"
)

type CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference interface {
	cdktf.ComplexObject
	Aggregation() CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference
	AggregationInput() interface{}
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
	Custom() CleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference
	CustomInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *CleanroomsConfiguredTableAnalysisRulesPolicyV1
	SetInternalValue(val *CleanroomsConfiguredTableAnalysisRulesPolicyV1)
	List() CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference
	ListInput() interface{}
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
	PutAggregation(value *CleanroomsConfiguredTableAnalysisRulesPolicyV1Aggregation)
	PutCustom(value *CleanroomsConfiguredTableAnalysisRulesPolicyV1Custom)
	PutList(value *CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStruct)
	ResetAggregation()
	ResetCustom()
	ResetList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference
type jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) Aggregation() CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference {
	var returns CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference
	_jsii_.Get(
		j,
		"aggregation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) AggregationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aggregationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) Custom() CleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference {
	var returns CleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) CustomInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) InternalValue() *CleanroomsConfiguredTableAnalysisRulesPolicyV1 {
	var returns *CleanroomsConfiguredTableAnalysisRulesPolicyV1
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) List() CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference {
	var returns CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference
	_jsii_.Get(
		j,
		"list",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) ListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"listInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference {
	_init_.Initialize()

	if err := validateNewCleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference{}

	_jsii_.Create(
		"awscc.cleanroomsConfiguredTable.CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference_Override(c CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cleanroomsConfiguredTable.CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference)SetInternalValue(val *CleanroomsConfiguredTableAnalysisRulesPolicyV1) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) PutAggregation(value *CleanroomsConfiguredTableAnalysisRulesPolicyV1Aggregation) {
	if err := c.validatePutAggregationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAggregation",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) PutCustom(value *CleanroomsConfiguredTableAnalysisRulesPolicyV1Custom) {
	if err := c.validatePutCustomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustom",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) PutList(value *CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStruct) {
	if err := c.validatePutListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putList",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) ResetAggregation() {
	_jsii_.InvokeVoid(
		c,
		"resetAggregation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) ResetCustom() {
	_jsii_.InvokeVoid(
		c,
		"resetCustom",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) ResetList() {
	_jsii_.InvokeVoid(
		c,
		"resetList",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

