package mediapackageoriginendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediapackageoriginendpoint/internal"
)

type MediapackageOriginEndpointHlsPackageOutputReference interface {
	cdktf.ComplexObject
	AdMarkers() *string
	SetAdMarkers(val *string)
	AdMarkersInput() *string
	AdsOnDeliveryRestrictions() *string
	SetAdsOnDeliveryRestrictions(val *string)
	AdsOnDeliveryRestrictionsInput() *string
	AdTriggers() *[]*string
	SetAdTriggers(val *[]*string)
	AdTriggersInput() *[]*string
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
	Encryption() MediapackageOriginEndpointHlsPackageEncryptionOutputReference
	EncryptionInput() interface{}
	// Experimental.
	Fqn() *string
	IncludeDvbSubtitles() interface{}
	SetIncludeDvbSubtitles(val interface{})
	IncludeDvbSubtitlesInput() interface{}
	IncludeIframeOnlyStream() interface{}
	SetIncludeIframeOnlyStream(val interface{})
	IncludeIframeOnlyStreamInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PlaylistType() *string
	SetPlaylistType(val *string)
	PlaylistTypeInput() *string
	PlaylistWindowSeconds() *float64
	SetPlaylistWindowSeconds(val *float64)
	PlaylistWindowSecondsInput() *float64
	ProgramDateTimeIntervalSeconds() *float64
	SetProgramDateTimeIntervalSeconds(val *float64)
	ProgramDateTimeIntervalSecondsInput() *float64
	SegmentDurationSeconds() *float64
	SetSegmentDurationSeconds(val *float64)
	SegmentDurationSecondsInput() *float64
	StreamSelection() MediapackageOriginEndpointHlsPackageStreamSelectionOutputReference
	StreamSelectionInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseAudioRenditionGroup() interface{}
	SetUseAudioRenditionGroup(val interface{})
	UseAudioRenditionGroupInput() interface{}
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
	PutEncryption(value *MediapackageOriginEndpointHlsPackageEncryption)
	PutStreamSelection(value *MediapackageOriginEndpointHlsPackageStreamSelection)
	ResetAdMarkers()
	ResetAdsOnDeliveryRestrictions()
	ResetAdTriggers()
	ResetEncryption()
	ResetIncludeDvbSubtitles()
	ResetIncludeIframeOnlyStream()
	ResetPlaylistType()
	ResetPlaylistWindowSeconds()
	ResetProgramDateTimeIntervalSeconds()
	ResetSegmentDurationSeconds()
	ResetStreamSelection()
	ResetUseAudioRenditionGroup()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediapackageOriginEndpointHlsPackageOutputReference
type jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) AdMarkers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adMarkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) AdMarkersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adMarkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) AdsOnDeliveryRestrictions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adsOnDeliveryRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) AdsOnDeliveryRestrictionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adsOnDeliveryRestrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) AdTriggers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adTriggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) AdTriggersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adTriggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) Encryption() MediapackageOriginEndpointHlsPackageEncryptionOutputReference {
	var returns MediapackageOriginEndpointHlsPackageEncryptionOutputReference
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) EncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) IncludeDvbSubtitles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeDvbSubtitles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) IncludeDvbSubtitlesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeDvbSubtitlesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) IncludeIframeOnlyStream() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeIframeOnlyStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) IncludeIframeOnlyStreamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeIframeOnlyStreamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) PlaylistType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"playlistType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) PlaylistTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"playlistTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) PlaylistWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"playlistWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) PlaylistWindowSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"playlistWindowSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ProgramDateTimeIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programDateTimeIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ProgramDateTimeIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programDateTimeIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) SegmentDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) SegmentDurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentDurationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) StreamSelection() MediapackageOriginEndpointHlsPackageStreamSelectionOutputReference {
	var returns MediapackageOriginEndpointHlsPackageStreamSelectionOutputReference
	_jsii_.Get(
		j,
		"streamSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) StreamSelectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) UseAudioRenditionGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAudioRenditionGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) UseAudioRenditionGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAudioRenditionGroupInput",
		&returns,
	)
	return returns
}


func NewMediapackageOriginEndpointHlsPackageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediapackageOriginEndpointHlsPackageOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackageOriginEndpointHlsPackageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference{}

	_jsii_.Create(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpointHlsPackageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediapackageOriginEndpointHlsPackageOutputReference_Override(m MediapackageOriginEndpointHlsPackageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpointHlsPackageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference)SetAdMarkers(val *string) {
	if err := j.validateSetAdMarkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adMarkers",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference)SetAdsOnDeliveryRestrictions(val *string) {
	if err := j.validateSetAdsOnDeliveryRestrictionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adsOnDeliveryRestrictions",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference)SetAdTriggers(val *[]*string) {
	if err := j.validateSetAdTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adTriggers",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference)SetIncludeDvbSubtitles(val interface{}) {
	if err := j.validateSetIncludeDvbSubtitlesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeDvbSubtitles",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference)SetIncludeIframeOnlyStream(val interface{}) {
	if err := j.validateSetIncludeIframeOnlyStreamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeIframeOnlyStream",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference)SetPlaylistType(val *string) {
	if err := j.validateSetPlaylistTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"playlistType",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference)SetPlaylistWindowSeconds(val *float64) {
	if err := j.validateSetPlaylistWindowSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"playlistWindowSeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference)SetProgramDateTimeIntervalSeconds(val *float64) {
	if err := j.validateSetProgramDateTimeIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programDateTimeIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference)SetSegmentDurationSeconds(val *float64) {
	if err := j.validateSetSegmentDurationSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference)SetUseAudioRenditionGroup(val interface{}) {
	if err := j.validateSetUseAudioRenditionGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAudioRenditionGroup",
		val,
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) PutEncryption(value *MediapackageOriginEndpointHlsPackageEncryption) {
	if err := m.validatePutEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncryption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) PutStreamSelection(value *MediapackageOriginEndpointHlsPackageStreamSelection) {
	if err := m.validatePutStreamSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStreamSelection",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ResetAdMarkers() {
	_jsii_.InvokeVoid(
		m,
		"resetAdMarkers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ResetAdsOnDeliveryRestrictions() {
	_jsii_.InvokeVoid(
		m,
		"resetAdsOnDeliveryRestrictions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ResetAdTriggers() {
	_jsii_.InvokeVoid(
		m,
		"resetAdTriggers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ResetEncryption() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryption",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ResetIncludeDvbSubtitles() {
	_jsii_.InvokeVoid(
		m,
		"resetIncludeDvbSubtitles",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ResetIncludeIframeOnlyStream() {
	_jsii_.InvokeVoid(
		m,
		"resetIncludeIframeOnlyStream",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ResetPlaylistType() {
	_jsii_.InvokeVoid(
		m,
		"resetPlaylistType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ResetPlaylistWindowSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetPlaylistWindowSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ResetProgramDateTimeIntervalSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetProgramDateTimeIntervalSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ResetSegmentDurationSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentDurationSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ResetStreamSelection() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamSelection",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ResetUseAudioRenditionGroup() {
	_jsii_.InvokeVoid(
		m,
		"resetUseAudioRenditionGroup",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediapackageOriginEndpointHlsPackageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

