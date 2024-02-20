package applicationautoscalingscalingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/applicationautoscalingscalingpolicy/internal"
)

type ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList
type jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList {
	_init_.Initialize()

	if err := validateNewApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList{}

	_jsii_.Create(
		"awscc.applicationautoscalingScalingPolicy.ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList_Override(a ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.applicationautoscalingScalingPolicy.ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList) Get(index *float64) ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

