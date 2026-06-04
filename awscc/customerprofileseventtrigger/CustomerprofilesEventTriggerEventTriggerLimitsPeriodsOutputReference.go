package customerprofileseventtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/customerprofileseventtrigger/internal"
)

type CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxInvocationsPerProfile() *float64
	SetMaxInvocationsPerProfile(val *float64)
	MaxInvocationsPerProfileInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Unlimited() interface{}
	SetUnlimited(val interface{})
	UnlimitedInput() interface{}
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
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
	ResetMaxInvocationsPerProfile()
	ResetUnit()
	ResetUnlimited()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference
type jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) MaxInvocationsPerProfile() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInvocationsPerProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) MaxInvocationsPerProfileInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInvocationsPerProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) Unlimited() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unlimited",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) UnlimitedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unlimitedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewCustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference {
	_init_.Initialize()

	if err := validateNewCustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference{}

	_jsii_.Create(
		"awscc.customerprofilesEventTrigger.CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference_Override(c CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.customerprofilesEventTrigger.CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference)SetMaxInvocationsPerProfile(val *float64) {
	if err := j.validateSetMaxInvocationsPerProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInvocationsPerProfile",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference)SetUnit(val *string) {
	if err := j.validateSetUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference)SetUnlimited(val interface{}) {
	if err := j.validateSetUnlimitedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unlimited",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference)SetValue(val *float64) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) ResetMaxInvocationsPerProfile() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxInvocationsPerProfile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) ResetUnit() {
	_jsii_.InvokeVoid(
		c,
		"resetUnit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) ResetUnlimited() {
	_jsii_.InvokeVoid(
		c,
		"resetUnlimited",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		c,
		"resetValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CustomerprofilesEventTriggerEventTriggerLimitsPeriodsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

