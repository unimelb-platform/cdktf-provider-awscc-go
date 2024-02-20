package autoscalingautoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/autoscalingautoscalinggroup/internal"
)

type AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference interface {
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
	InstanceRequirements() AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference
	InstanceRequirementsInput() interface{}
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LaunchTemplateSpecification() AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesLaunchTemplateSpecificationOutputReference
	LaunchTemplateSpecificationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeightedCapacity() *string
	SetWeightedCapacity(val *string)
	WeightedCapacityInput() *string
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
	PutInstanceRequirements(value *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirements)
	PutLaunchTemplateSpecification(value *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesLaunchTemplateSpecification)
	ResetInstanceRequirements()
	ResetInstanceType()
	ResetLaunchTemplateSpecification()
	ResetWeightedCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference
type jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) InstanceRequirements() AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference {
	var returns AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference
	_jsii_.Get(
		j,
		"instanceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) InstanceRequirementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) LaunchTemplateSpecification() AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesLaunchTemplateSpecificationOutputReference {
	var returns AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesLaunchTemplateSpecificationOutputReference
	_jsii_.Get(
		j,
		"launchTemplateSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) LaunchTemplateSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) WeightedCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weightedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) WeightedCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weightedCapacityInput",
		&returns,
	)
	return returns
}


func NewAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference{}

	_jsii_.Create(
		"awscc.autoscalingAutoScalingGroup.AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference_Override(a AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.autoscalingAutoScalingGroup.AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference)SetWeightedCapacity(val *string) {
	if err := j.validateSetWeightedCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weightedCapacity",
		val,
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) PutInstanceRequirements(value *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirements) {
	if err := a.validatePutInstanceRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstanceRequirements",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) PutLaunchTemplateSpecification(value *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesLaunchTemplateSpecification) {
	if err := a.validatePutLaunchTemplateSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLaunchTemplateSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) ResetInstanceRequirements() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceRequirements",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) ResetLaunchTemplateSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchTemplateSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) ResetWeightedCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetWeightedCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

