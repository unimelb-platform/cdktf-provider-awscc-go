package applicationautoscalingscalingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/applicationautoscalingscalingpolicy/internal"
)

type ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference interface {
	cdktf.ComplexObject
	AdjustmentType() *string
	SetAdjustmentType(val *string)
	AdjustmentTypeInput() *string
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
	Cooldown() *float64
	SetCooldown(val *float64)
	CooldownInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricAggregationType() *string
	SetMetricAggregationType(val *string)
	MetricAggregationTypeInput() *string
	MinAdjustmentMagnitude() *float64
	SetMinAdjustmentMagnitude(val *float64)
	MinAdjustmentMagnitudeInput() *float64
	StepAdjustments() ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationStepAdjustmentsList
	StepAdjustmentsInput() interface{}
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
	PutStepAdjustments(value interface{})
	ResetAdjustmentType()
	ResetCooldown()
	ResetMetricAggregationType()
	ResetMinAdjustmentMagnitude()
	ResetStepAdjustments()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference
type jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) AdjustmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) AdjustmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) Cooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) CooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) MetricAggregationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricAggregationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) MetricAggregationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricAggregationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) MinAdjustmentMagnitude() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAdjustmentMagnitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) MinAdjustmentMagnitudeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAdjustmentMagnitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) StepAdjustments() ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationStepAdjustmentsList {
	var returns ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationStepAdjustmentsList
	_jsii_.Get(
		j,
		"stepAdjustments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) StepAdjustmentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepAdjustmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.applicationautoscalingScalingPolicy.ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference_Override(a ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.applicationautoscalingScalingPolicy.ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference)SetAdjustmentType(val *string) {
	if err := j.validateSetAdjustmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adjustmentType",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference)SetCooldown(val *float64) {
	if err := j.validateSetCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cooldown",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference)SetMetricAggregationType(val *string) {
	if err := j.validateSetMetricAggregationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricAggregationType",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference)SetMinAdjustmentMagnitude(val *float64) {
	if err := j.validateSetMinAdjustmentMagnitudeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minAdjustmentMagnitude",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) PutStepAdjustments(value interface{}) {
	if err := a.validatePutStepAdjustmentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStepAdjustments",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) ResetAdjustmentType() {
	_jsii_.InvokeVoid(
		a,
		"resetAdjustmentType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) ResetCooldown() {
	_jsii_.InvokeVoid(
		a,
		"resetCooldown",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) ResetMetricAggregationType() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricAggregationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) ResetMinAdjustmentMagnitude() {
	_jsii_.InvokeVoid(
		a,
		"resetMinAdjustmentMagnitude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) ResetStepAdjustments() {
	_jsii_.InvokeVoid(
		a,
		"resetStepAdjustments",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

