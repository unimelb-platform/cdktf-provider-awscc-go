package autoscalingautoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/autoscalingautoscalinggroup/internal"
)

type AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference interface {
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
	InstancesDistribution() AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference
	InstancesDistributionInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LaunchTemplate() AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOutputReference
	LaunchTemplateInput() *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplate
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
	PutInstancesDistribution(value *AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistribution)
	PutLaunchTemplate(value *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplate)
	ResetInstancesDistribution()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference
type jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) InstancesDistribution() AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference {
	var returns AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference
	_jsii_.Get(
		j,
		"instancesDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) InstancesDistributionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instancesDistributionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) LaunchTemplate() AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOutputReference {
	var returns AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) LaunchTemplateInput() *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplate {
	var returns *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplate
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingAutoScalingGroupMixedInstancesPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference{}

	_jsii_.Create(
		"awscc.autoscalingAutoScalingGroup.AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference_Override(a AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.autoscalingAutoScalingGroup.AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) PutInstancesDistribution(value *AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistribution) {
	if err := a.validatePutInstancesDistributionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstancesDistribution",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) PutLaunchTemplate(value *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplate) {
	if err := a.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) ResetInstancesDistribution() {
	_jsii_.InvokeVoid(
		a,
		"resetInstancesDistribution",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

