package wafv2loggingconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/wafv2loggingconfiguration/internal"
)

type Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference interface {
	cdktf.ComplexObject
	ActionCondition() Wafv2LoggingConfigurationLoggingFilterFiltersConditionsActionConditionOutputReference
	ActionConditionInput() interface{}
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
	InternalValue() *Wafv2LoggingConfigurationLoggingFilterFiltersConditions
	SetInternalValue(val *Wafv2LoggingConfigurationLoggingFilterFiltersConditions)
	LabelNameCondition() Wafv2LoggingConfigurationLoggingFilterFiltersConditionsLabelNameConditionOutputReference
	LabelNameConditionInput() interface{}
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
	PutActionCondition(value *Wafv2LoggingConfigurationLoggingFilterFiltersConditionsActionCondition)
	PutLabelNameCondition(value *Wafv2LoggingConfigurationLoggingFilterFiltersConditionsLabelNameCondition)
	ResetActionCondition()
	ResetLabelNameCondition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference
type jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) ActionCondition() Wafv2LoggingConfigurationLoggingFilterFiltersConditionsActionConditionOutputReference {
	var returns Wafv2LoggingConfigurationLoggingFilterFiltersConditionsActionConditionOutputReference
	_jsii_.Get(
		j,
		"actionCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) ActionConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) InternalValue() *Wafv2LoggingConfigurationLoggingFilterFiltersConditions {
	var returns *Wafv2LoggingConfigurationLoggingFilterFiltersConditions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) LabelNameCondition() Wafv2LoggingConfigurationLoggingFilterFiltersConditionsLabelNameConditionOutputReference {
	var returns Wafv2LoggingConfigurationLoggingFilterFiltersConditionsLabelNameConditionOutputReference
	_jsii_.Get(
		j,
		"labelNameCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) LabelNameConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelNameConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference{}

	_jsii_.Create(
		"awscc.wafv2LoggingConfiguration.Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference_Override(w Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.wafv2LoggingConfiguration.Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference)SetInternalValue(val *Wafv2LoggingConfigurationLoggingFilterFiltersConditions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) PutActionCondition(value *Wafv2LoggingConfigurationLoggingFilterFiltersConditionsActionCondition) {
	if err := w.validatePutActionConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putActionCondition",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) PutLabelNameCondition(value *Wafv2LoggingConfigurationLoggingFilterFiltersConditionsLabelNameCondition) {
	if err := w.validatePutLabelNameConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLabelNameCondition",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) ResetActionCondition() {
	_jsii_.InvokeVoid(
		w,
		"resetActionCondition",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) ResetLabelNameCondition() {
	_jsii_.InvokeVoid(
		w,
		"resetLabelNameCondition",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2LoggingConfigurationLoggingFilterFiltersConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

