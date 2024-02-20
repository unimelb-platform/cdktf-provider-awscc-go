package mediapackagev2originendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediapackagev2originendpoint/internal"
)

type Mediapackagev2OriginEndpointSegmentOutputReference interface {
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
	Encryption() Mediapackagev2OriginEndpointSegmentEncryptionOutputReference
	EncryptionInput() interface{}
	// Experimental.
	Fqn() *string
	IncludeIframeOnlyStreams() interface{}
	SetIncludeIframeOnlyStreams(val interface{})
	IncludeIframeOnlyStreamsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Scte() Mediapackagev2OriginEndpointSegmentScteOutputReference
	ScteInput() interface{}
	SegmentDurationSeconds() *float64
	SetSegmentDurationSeconds(val *float64)
	SegmentDurationSecondsInput() *float64
	SegmentName() *string
	SetSegmentName(val *string)
	SegmentNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TsIncludeDvbSubtitles() interface{}
	SetTsIncludeDvbSubtitles(val interface{})
	TsIncludeDvbSubtitlesInput() interface{}
	TsUseAudioRenditionGroup() interface{}
	SetTsUseAudioRenditionGroup(val interface{})
	TsUseAudioRenditionGroupInput() interface{}
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
	PutEncryption(value *Mediapackagev2OriginEndpointSegmentEncryption)
	PutScte(value *Mediapackagev2OriginEndpointSegmentScte)
	ResetEncryption()
	ResetIncludeIframeOnlyStreams()
	ResetScte()
	ResetSegmentDurationSeconds()
	ResetSegmentName()
	ResetTsIncludeDvbSubtitles()
	ResetTsUseAudioRenditionGroup()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Mediapackagev2OriginEndpointSegmentOutputReference
type jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) Encryption() Mediapackagev2OriginEndpointSegmentEncryptionOutputReference {
	var returns Mediapackagev2OriginEndpointSegmentEncryptionOutputReference
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) EncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) IncludeIframeOnlyStreams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeIframeOnlyStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) IncludeIframeOnlyStreamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeIframeOnlyStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) Scte() Mediapackagev2OriginEndpointSegmentScteOutputReference {
	var returns Mediapackagev2OriginEndpointSegmentScteOutputReference
	_jsii_.Get(
		j,
		"scte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) ScteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) SegmentDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) SegmentDurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentDurationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) SegmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) SegmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) TsIncludeDvbSubtitles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tsIncludeDvbSubtitles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) TsIncludeDvbSubtitlesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tsIncludeDvbSubtitlesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) TsUseAudioRenditionGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tsUseAudioRenditionGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) TsUseAudioRenditionGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tsUseAudioRenditionGroupInput",
		&returns,
	)
	return returns
}


func NewMediapackagev2OriginEndpointSegmentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Mediapackagev2OriginEndpointSegmentOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackagev2OriginEndpointSegmentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference{}

	_jsii_.Create(
		"awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpointSegmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediapackagev2OriginEndpointSegmentOutputReference_Override(m Mediapackagev2OriginEndpointSegmentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpointSegmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference)SetIncludeIframeOnlyStreams(val interface{}) {
	if err := j.validateSetIncludeIframeOnlyStreamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeIframeOnlyStreams",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference)SetSegmentDurationSeconds(val *float64) {
	if err := j.validateSetSegmentDurationSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference)SetSegmentName(val *string) {
	if err := j.validateSetSegmentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentName",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference)SetTsIncludeDvbSubtitles(val interface{}) {
	if err := j.validateSetTsIncludeDvbSubtitlesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tsIncludeDvbSubtitles",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference)SetTsUseAudioRenditionGroup(val interface{}) {
	if err := j.validateSetTsUseAudioRenditionGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tsUseAudioRenditionGroup",
		val,
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) PutEncryption(value *Mediapackagev2OriginEndpointSegmentEncryption) {
	if err := m.validatePutEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncryption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) PutScte(value *Mediapackagev2OriginEndpointSegmentScte) {
	if err := m.validatePutScteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putScte",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) ResetEncryption() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryption",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) ResetIncludeIframeOnlyStreams() {
	_jsii_.InvokeVoid(
		m,
		"resetIncludeIframeOnlyStreams",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) ResetScte() {
	_jsii_.InvokeVoid(
		m,
		"resetScte",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) ResetSegmentDurationSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentDurationSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) ResetSegmentName() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) ResetTsIncludeDvbSubtitles() {
	_jsii_.InvokeVoid(
		m,
		"resetTsIncludeDvbSubtitles",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) ResetTsUseAudioRenditionGroup() {
	_jsii_.InvokeVoid(
		m,
		"resetTsUseAudioRenditionGroup",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointSegmentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

