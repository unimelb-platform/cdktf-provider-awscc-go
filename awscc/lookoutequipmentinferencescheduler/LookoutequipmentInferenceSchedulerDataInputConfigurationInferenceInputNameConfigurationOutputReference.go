package lookoutequipmentinferencescheduler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lookoutequipmentinferencescheduler/internal"
)

type LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference interface {
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
	ComponentTimestampDelimiter() *string
	SetComponentTimestampDelimiter(val *string)
	ComponentTimestampDelimiterInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimestampFormat() *string
	SetTimestampFormat(val *string)
	TimestampFormatInput() *string
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
	ResetComponentTimestampDelimiter()
	ResetTimestampFormat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference
type jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) ComponentTimestampDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentTimestampDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) ComponentTimestampDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentTimestampDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) TimestampFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) TimestampFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampFormatInput",
		&returns,
	)
	return returns
}


func NewLookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewLookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.lookoutequipmentInferenceScheduler.LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference_Override(l LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lookoutequipmentInferenceScheduler.LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference)SetComponentTimestampDelimiter(val *string) {
	if err := j.validateSetComponentTimestampDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"componentTimestampDelimiter",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference)SetTimestampFormat(val *string) {
	if err := j.validateSetTimestampFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampFormat",
		val,
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) ResetComponentTimestampDelimiter() {
	_jsii_.InvokeVoid(
		l,
		"resetComponentTimestampDelimiter",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) ResetTimestampFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetTimestampFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

