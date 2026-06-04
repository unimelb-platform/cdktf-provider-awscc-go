package cloudformationlambdahook

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cloudformationlambdahook/internal"
)

type CloudformationLambdaHookTargetFiltersOutputReference interface {
	cdktf.ComplexObject
	Actions() *[]*string
	SetActions(val *[]*string)
	ActionsInput() *[]*string
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
	InvocationPoints() *[]*string
	SetInvocationPoints(val *[]*string)
	InvocationPointsInput() *[]*string
	TargetNames() *[]*string
	SetTargetNames(val *[]*string)
	TargetNamesInput() *[]*string
	Targets() CloudformationLambdaHookTargetFiltersTargetsList
	TargetsInput() interface{}
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
	PutTargets(value interface{})
	ResetActions()
	ResetInvocationPoints()
	ResetTargetNames()
	ResetTargets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudformationLambdaHookTargetFiltersOutputReference
type jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) Actions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) ActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) InvocationPoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"invocationPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) InvocationPointsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"invocationPointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) TargetNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) TargetNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) Targets() CloudformationLambdaHookTargetFiltersTargetsList {
	var returns CloudformationLambdaHookTargetFiltersTargetsList
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudformationLambdaHookTargetFiltersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudformationLambdaHookTargetFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewCloudformationLambdaHookTargetFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference{}

	_jsii_.Create(
		"awscc.cloudformationLambdaHook.CloudformationLambdaHookTargetFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudformationLambdaHookTargetFiltersOutputReference_Override(c CloudformationLambdaHookTargetFiltersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cloudformationLambdaHook.CloudformationLambdaHookTargetFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference)SetActions(val *[]*string) {
	if err := j.validateSetActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference)SetInvocationPoints(val *[]*string) {
	if err := j.validateSetInvocationPointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invocationPoints",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference)SetTargetNames(val *[]*string) {
	if err := j.validateSetTargetNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetNames",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) PutTargets(value interface{}) {
	if err := c.validatePutTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTargets",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) ResetActions() {
	_jsii_.InvokeVoid(
		c,
		"resetActions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) ResetInvocationPoints() {
	_jsii_.InvokeVoid(
		c,
		"resetInvocationPoints",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) ResetTargetNames() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) ResetTargets() {
	_jsii_.InvokeVoid(
		c,
		"resetTargets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudformationLambdaHookTargetFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

