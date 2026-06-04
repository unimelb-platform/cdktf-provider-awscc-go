package cloudformationlambdahook

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cloudformationlambdahook/internal"
)

type CloudformationLambdaHookStackFiltersOutputReference interface {
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
	FilteringCriteria() *string
	SetFilteringCriteria(val *string)
	FilteringCriteriaInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StackNames() CloudformationLambdaHookStackFiltersStackNamesOutputReference
	StackNamesInput() interface{}
	StackRoles() CloudformationLambdaHookStackFiltersStackRolesOutputReference
	StackRolesInput() interface{}
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
	PutStackNames(value *CloudformationLambdaHookStackFiltersStackNames)
	PutStackRoles(value *CloudformationLambdaHookStackFiltersStackRoles)
	ResetFilteringCriteria()
	ResetStackNames()
	ResetStackRoles()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudformationLambdaHookStackFiltersOutputReference
type jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) FilteringCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filteringCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) FilteringCriteriaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filteringCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) StackNames() CloudformationLambdaHookStackFiltersStackNamesOutputReference {
	var returns CloudformationLambdaHookStackFiltersStackNamesOutputReference
	_jsii_.Get(
		j,
		"stackNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) StackNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stackNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) StackRoles() CloudformationLambdaHookStackFiltersStackRolesOutputReference {
	var returns CloudformationLambdaHookStackFiltersStackRolesOutputReference
	_jsii_.Get(
		j,
		"stackRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) StackRolesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stackRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudformationLambdaHookStackFiltersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudformationLambdaHookStackFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewCloudformationLambdaHookStackFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference{}

	_jsii_.Create(
		"awscc.cloudformationLambdaHook.CloudformationLambdaHookStackFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudformationLambdaHookStackFiltersOutputReference_Override(c CloudformationLambdaHookStackFiltersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cloudformationLambdaHook.CloudformationLambdaHookStackFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference)SetFilteringCriteria(val *string) {
	if err := j.validateSetFilteringCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filteringCriteria",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) PutStackNames(value *CloudformationLambdaHookStackFiltersStackNames) {
	if err := c.validatePutStackNamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStackNames",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) PutStackRoles(value *CloudformationLambdaHookStackFiltersStackRoles) {
	if err := c.validatePutStackRolesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStackRoles",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) ResetFilteringCriteria() {
	_jsii_.InvokeVoid(
		c,
		"resetFilteringCriteria",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) ResetStackNames() {
	_jsii_.InvokeVoid(
		c,
		"resetStackNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) ResetStackRoles() {
	_jsii_.InvokeVoid(
		c,
		"resetStackRoles",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudformationLambdaHookStackFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

