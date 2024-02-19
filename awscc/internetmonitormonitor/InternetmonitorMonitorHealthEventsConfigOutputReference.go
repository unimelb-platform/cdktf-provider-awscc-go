package internetmonitormonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/internetmonitormonitor/internal"
)

type InternetmonitorMonitorHealthEventsConfigOutputReference interface {
	cdktf.ComplexObject
	AvailabilityLocalHealthEventsConfig() InternetmonitorMonitorHealthEventsConfigAvailabilityLocalHealthEventsConfigOutputReference
	AvailabilityLocalHealthEventsConfigInput() interface{}
	AvailabilityScoreThreshold() *float64
	SetAvailabilityScoreThreshold(val *float64)
	AvailabilityScoreThresholdInput() *float64
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
	PerformanceLocalHealthEventsConfig() InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference
	PerformanceLocalHealthEventsConfigInput() interface{}
	PerformanceScoreThreshold() *float64
	SetPerformanceScoreThreshold(val *float64)
	PerformanceScoreThresholdInput() *float64
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
	PutAvailabilityLocalHealthEventsConfig(value *InternetmonitorMonitorHealthEventsConfigAvailabilityLocalHealthEventsConfig)
	PutPerformanceLocalHealthEventsConfig(value *InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfig)
	ResetAvailabilityLocalHealthEventsConfig()
	ResetAvailabilityScoreThreshold()
	ResetPerformanceLocalHealthEventsConfig()
	ResetPerformanceScoreThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for InternetmonitorMonitorHealthEventsConfigOutputReference
type jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) AvailabilityLocalHealthEventsConfig() InternetmonitorMonitorHealthEventsConfigAvailabilityLocalHealthEventsConfigOutputReference {
	var returns InternetmonitorMonitorHealthEventsConfigAvailabilityLocalHealthEventsConfigOutputReference
	_jsii_.Get(
		j,
		"availabilityLocalHealthEventsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) AvailabilityLocalHealthEventsConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availabilityLocalHealthEventsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) AvailabilityScoreThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityScoreThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) AvailabilityScoreThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityScoreThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) PerformanceLocalHealthEventsConfig() InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference {
	var returns InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference
	_jsii_.Get(
		j,
		"performanceLocalHealthEventsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) PerformanceLocalHealthEventsConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceLocalHealthEventsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) PerformanceScoreThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceScoreThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) PerformanceScoreThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceScoreThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewInternetmonitorMonitorHealthEventsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) InternetmonitorMonitorHealthEventsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewInternetmonitorMonitorHealthEventsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference{}

	_jsii_.Create(
		"awscc.internetmonitorMonitor.InternetmonitorMonitorHealthEventsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewInternetmonitorMonitorHealthEventsConfigOutputReference_Override(i InternetmonitorMonitorHealthEventsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.internetmonitorMonitor.InternetmonitorMonitorHealthEventsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference)SetAvailabilityScoreThreshold(val *float64) {
	if err := j.validateSetAvailabilityScoreThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityScoreThreshold",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference)SetPerformanceScoreThreshold(val *float64) {
	if err := j.validateSetPerformanceScoreThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceScoreThreshold",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) PutAvailabilityLocalHealthEventsConfig(value *InternetmonitorMonitorHealthEventsConfigAvailabilityLocalHealthEventsConfig) {
	if err := i.validatePutAvailabilityLocalHealthEventsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAvailabilityLocalHealthEventsConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) PutPerformanceLocalHealthEventsConfig(value *InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfig) {
	if err := i.validatePutPerformanceLocalHealthEventsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPerformanceLocalHealthEventsConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) ResetAvailabilityLocalHealthEventsConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetAvailabilityLocalHealthEventsConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) ResetAvailabilityScoreThreshold() {
	_jsii_.InvokeVoid(
		i,
		"resetAvailabilityScoreThreshold",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) ResetPerformanceLocalHealthEventsConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetPerformanceLocalHealthEventsConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) ResetPerformanceScoreThreshold() {
	_jsii_.InvokeVoid(
		i,
		"resetPerformanceScoreThreshold",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

