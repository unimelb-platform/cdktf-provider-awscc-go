package mediaconnectflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediaconnectflow/internal"
)

type MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference interface {
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
	State() *string
	SetState(val *string)
	StateInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThresholdSeconds() *float64
	SetThresholdSeconds(val *float64)
	ThresholdSecondsInput() *float64
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
	ResetState()
	ResetThresholdSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference
type jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) ThresholdSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) ThresholdSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdSecondsInput",
		&returns,
	)
	return returns
}


func NewMediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference{}

	_jsii_.Create(
		"awscc.mediaconnectFlow.MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference_Override(m MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediaconnectFlow.MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference)SetThresholdSeconds(val *float64) {
	if err := j.validateSetThresholdSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdSeconds",
		val,
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		m,
		"resetState",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) ResetThresholdSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetThresholdSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudioOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

