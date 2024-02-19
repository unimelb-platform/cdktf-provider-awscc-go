package applicationautoscalingscalingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/applicationautoscalingscalingpolicy/internal"
)

type ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference interface {
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
	Metric() ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricOutputReference
	MetricInput() interface{}
	Stat() *string
	SetStat(val *string)
	StatInput() *string
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
	PutMetric(value *ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetric)
	ResetMetric()
	ResetStat()
	ResetUnit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference
type jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) Metric() ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricOutputReference {
	var returns ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricOutputReference
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) MetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) Stat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) StatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}


func NewApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference{}

	_jsii_.Create(
		"awscc.applicationautoscalingScalingPolicy.ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference_Override(a ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.applicationautoscalingScalingPolicy.ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference)SetStat(val *string) {
	if err := j.validateSetStatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stat",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference)SetUnit(val *string) {
	if err := j.validateSetUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) PutMetric(value *ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetric) {
	if err := a.validatePutMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetric",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) ResetMetric() {
	_jsii_.InvokeVoid(
		a,
		"resetMetric",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) ResetStat() {
	_jsii_.InvokeVoid(
		a,
		"resetStat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) ResetUnit() {
	_jsii_.InvokeVoid(
		a,
		"resetUnit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

