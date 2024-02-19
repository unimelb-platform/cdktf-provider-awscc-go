package medialivemultiplexprogram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/medialivemultiplexprogram/internal"
)

type MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference interface {
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
	PreferredChannelPipeline() *string
	SetPreferredChannelPipeline(val *string)
	PreferredChannelPipelineInput() *string
	ProgramNumber() *float64
	SetProgramNumber(val *float64)
	ProgramNumberInput() *float64
	ServiceDescriptor() MedialiveMultiplexprogramMultiplexProgramSettingsServiceDescriptorOutputReference
	ServiceDescriptorInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VideoSettings() MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettingsOutputReference
	VideoSettingsInput() interface{}
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
	PutServiceDescriptor(value *MedialiveMultiplexprogramMultiplexProgramSettingsServiceDescriptor)
	PutVideoSettings(value *MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettings)
	ResetPreferredChannelPipeline()
	ResetServiceDescriptor()
	ResetVideoSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference
type jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) PreferredChannelPipeline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredChannelPipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) PreferredChannelPipelineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredChannelPipelineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) ProgramNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) ProgramNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) ServiceDescriptor() MedialiveMultiplexprogramMultiplexProgramSettingsServiceDescriptorOutputReference {
	var returns MedialiveMultiplexprogramMultiplexProgramSettingsServiceDescriptorOutputReference
	_jsii_.Get(
		j,
		"serviceDescriptor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) ServiceDescriptorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceDescriptorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) VideoSettings() MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettingsOutputReference {
	var returns MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettingsOutputReference
	_jsii_.Get(
		j,
		"videoSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) VideoSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"videoSettingsInput",
		&returns,
	)
	return returns
}


func NewMedialiveMultiplexprogramMultiplexProgramSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveMultiplexprogramMultiplexProgramSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference{}

	_jsii_.Create(
		"awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveMultiplexprogramMultiplexProgramSettingsOutputReference_Override(m MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference)SetPreferredChannelPipeline(val *string) {
	if err := j.validateSetPreferredChannelPipelineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredChannelPipeline",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference)SetProgramNumber(val *float64) {
	if err := j.validateSetProgramNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programNumber",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) PutServiceDescriptor(value *MedialiveMultiplexprogramMultiplexProgramSettingsServiceDescriptor) {
	if err := m.validatePutServiceDescriptorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putServiceDescriptor",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) PutVideoSettings(value *MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettings) {
	if err := m.validatePutVideoSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putVideoSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) ResetPreferredChannelPipeline() {
	_jsii_.InvokeVoid(
		m,
		"resetPreferredChannelPipeline",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) ResetServiceDescriptor() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceDescriptor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) ResetVideoSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetVideoSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

