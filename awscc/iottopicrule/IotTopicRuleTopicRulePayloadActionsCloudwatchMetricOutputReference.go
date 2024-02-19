package iottopicrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/iottopicrule/internal"
)

type IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference interface {
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
	MetricName() *string
	SetMetricName(val *string)
	MetricNameInput() *string
	MetricNamespace() *string
	SetMetricNamespace(val *string)
	MetricNamespaceInput() *string
	MetricTimestamp() *string
	SetMetricTimestamp(val *string)
	MetricTimestampInput() *string
	MetricUnit() *string
	SetMetricUnit(val *string)
	MetricUnitInput() *string
	MetricValue() *string
	SetMetricValue(val *string)
	MetricValueInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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
	ResetMetricTimestamp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference
type jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) MetricNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) MetricNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) MetricTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) MetricTimestampInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) MetricUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) MetricUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) MetricValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) MetricValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference {
	_init_.Initialize()

	if err := validateNewIotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference{}

	_jsii_.Create(
		"awscc.iotTopicRule.IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference_Override(i IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.iotTopicRule.IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference)SetMetricNamespace(val *string) {
	if err := j.validateSetMetricNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricNamespace",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference)SetMetricTimestamp(val *string) {
	if err := j.validateSetMetricTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricTimestamp",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference)SetMetricUnit(val *string) {
	if err := j.validateSetMetricUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricUnit",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference)SetMetricValue(val *string) {
	if err := j.validateSetMetricValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricValue",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) ResetMetricTimestamp() {
	_jsii_.InvokeVoid(
		i,
		"resetMetricTimestamp",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

