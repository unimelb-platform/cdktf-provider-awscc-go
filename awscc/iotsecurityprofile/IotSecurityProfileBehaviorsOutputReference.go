package iotsecurityprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/iotsecurityprofile/internal"
)

type IotSecurityProfileBehaviorsOutputReference interface {
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
	Criteria() IotSecurityProfileBehaviorsCriteriaOutputReference
	CriteriaInput() interface{}
	ExportMetric() interface{}
	SetExportMetric(val interface{})
	ExportMetricInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Metric() *string
	SetMetric(val *string)
	MetricDimension() IotSecurityProfileBehaviorsMetricDimensionOutputReference
	MetricDimensionInput() interface{}
	MetricInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	SuppressAlerts() interface{}
	SetSuppressAlerts(val interface{})
	SuppressAlertsInput() interface{}
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
	PutCriteria(value *IotSecurityProfileBehaviorsCriteria)
	PutMetricDimension(value *IotSecurityProfileBehaviorsMetricDimension)
	ResetCriteria()
	ResetExportMetric()
	ResetMetric()
	ResetMetricDimension()
	ResetSuppressAlerts()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotSecurityProfileBehaviorsOutputReference
type jsiiProxy_IotSecurityProfileBehaviorsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) Criteria() IotSecurityProfileBehaviorsCriteriaOutputReference {
	var returns IotSecurityProfileBehaviorsCriteriaOutputReference
	_jsii_.Get(
		j,
		"criteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) CriteriaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"criteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) ExportMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) ExportMetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) Metric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) MetricDimension() IotSecurityProfileBehaviorsMetricDimensionOutputReference {
	var returns IotSecurityProfileBehaviorsMetricDimensionOutputReference
	_jsii_.Get(
		j,
		"metricDimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) MetricDimensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricDimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) MetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) SuppressAlerts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressAlerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) SuppressAlertsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressAlertsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotSecurityProfileBehaviorsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IotSecurityProfileBehaviorsOutputReference {
	_init_.Initialize()

	if err := validateNewIotSecurityProfileBehaviorsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotSecurityProfileBehaviorsOutputReference{}

	_jsii_.Create(
		"awscc.iotSecurityProfile.IotSecurityProfileBehaviorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIotSecurityProfileBehaviorsOutputReference_Override(i IotSecurityProfileBehaviorsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.iotSecurityProfile.IotSecurityProfileBehaviorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference)SetExportMetric(val interface{}) {
	if err := j.validateSetExportMetricParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportMetric",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference)SetMetric(val *string) {
	if err := j.validateSetMetricParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metric",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference)SetSuppressAlerts(val interface{}) {
	if err := j.validateSetSuppressAlertsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suppressAlerts",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) PutCriteria(value *IotSecurityProfileBehaviorsCriteria) {
	if err := i.validatePutCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCriteria",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) PutMetricDimension(value *IotSecurityProfileBehaviorsMetricDimension) {
	if err := i.validatePutMetricDimensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMetricDimension",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) ResetCriteria() {
	_jsii_.InvokeVoid(
		i,
		"resetCriteria",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) ResetExportMetric() {
	_jsii_.InvokeVoid(
		i,
		"resetExportMetric",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) ResetMetric() {
	_jsii_.InvokeVoid(
		i,
		"resetMetric",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) ResetMetricDimension() {
	_jsii_.InvokeVoid(
		i,
		"resetMetricDimension",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) ResetSuppressAlerts() {
	_jsii_.InvokeVoid(
		i,
		"resetSuppressAlerts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

