package cloudtrailtrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cloudtrailtrail/internal"
)

type CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference interface {
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
	EndsWith() *[]*string
	SetEndsWith(val *[]*string)
	EndsWithInput() *[]*string
	EqualTo() *[]*string
	SetEqualTo(val *[]*string)
	EqualToInput() *[]*string
	Field() *string
	SetField(val *string)
	FieldInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CloudtrailTrailAdvancedEventSelectorsFieldSelectors
	SetInternalValue(val *CloudtrailTrailAdvancedEventSelectorsFieldSelectors)
	NotEndsWith() *[]*string
	SetNotEndsWith(val *[]*string)
	NotEndsWithInput() *[]*string
	NotEquals() *[]*string
	SetNotEquals(val *[]*string)
	NotEqualsInput() *[]*string
	NotStartsWith() *[]*string
	SetNotStartsWith(val *[]*string)
	NotStartsWithInput() *[]*string
	StartsWith() *[]*string
	SetStartsWith(val *[]*string)
	StartsWithInput() *[]*string
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
	ResetEndsWith()
	ResetEqualTo()
	ResetNotEndsWith()
	ResetNotEquals()
	ResetNotStartsWith()
	ResetStartsWith()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference
type jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) EndsWith() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"endsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) EndsWithInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"endsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) EqualTo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) EqualToInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) Field() *string {
	var returns *string
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) FieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) InternalValue() *CloudtrailTrailAdvancedEventSelectorsFieldSelectors {
	var returns *CloudtrailTrailAdvancedEventSelectorsFieldSelectors
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) NotEndsWith() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEndsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) NotEndsWithInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEndsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) NotEquals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) NotEqualsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) NotStartsWith() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notStartsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) NotStartsWithInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notStartsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) StartsWith() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"startsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) StartsWithInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"startsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference {
	_init_.Initialize()

	if err := validateNewCloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference{}

	_jsii_.Create(
		"awscc.cloudtrailTrail.CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference_Override(c CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cloudtrailTrail.CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference)SetEndsWith(val *[]*string) {
	if err := j.validateSetEndsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endsWith",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference)SetEqualTo(val *[]*string) {
	if err := j.validateSetEqualToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"equalTo",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference)SetField(val *string) {
	if err := j.validateSetFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"field",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference)SetInternalValue(val *CloudtrailTrailAdvancedEventSelectorsFieldSelectors) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference)SetNotEndsWith(val *[]*string) {
	if err := j.validateSetNotEndsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notEndsWith",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference)SetNotEquals(val *[]*string) {
	if err := j.validateSetNotEqualsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notEquals",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference)SetNotStartsWith(val *[]*string) {
	if err := j.validateSetNotStartsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notStartsWith",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference)SetStartsWith(val *[]*string) {
	if err := j.validateSetStartsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startsWith",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) ResetEndsWith() {
	_jsii_.InvokeVoid(
		c,
		"resetEndsWith",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) ResetEqualTo() {
	_jsii_.InvokeVoid(
		c,
		"resetEqualTo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) ResetNotEndsWith() {
	_jsii_.InvokeVoid(
		c,
		"resetNotEndsWith",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) ResetNotEquals() {
	_jsii_.InvokeVoid(
		c,
		"resetNotEquals",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) ResetNotStartsWith() {
	_jsii_.InvokeVoid(
		c,
		"resetNotStartsWith",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) ResetStartsWith() {
	_jsii_.InvokeVoid(
		c,
		"resetStartsWith",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudtrailTrailAdvancedEventSelectorsFieldSelectorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

