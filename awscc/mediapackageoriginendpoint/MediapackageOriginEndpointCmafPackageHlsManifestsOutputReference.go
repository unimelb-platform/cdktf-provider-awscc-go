package mediapackageoriginendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediapackageoriginendpoint/internal"
)

type MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IncludeIframeOnlyStream() interface{}
	SetIncludeIframeOnlyStream(val interface{})
	IncludeIframeOnlyStreamInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ManifestName() *string
	SetManifestName(val *string)
	ManifestNameInput() *string
	PlaylistType() *string
	SetPlaylistType(val *string)
	PlaylistTypeInput() *string
	PlaylistWindowSeconds() *float64
	SetPlaylistWindowSeconds(val *float64)
	PlaylistWindowSecondsInput() *float64
	ProgramDateTimeIntervalSeconds() *float64
	SetProgramDateTimeIntervalSeconds(val *float64)
	ProgramDateTimeIntervalSecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	ResetAdMarkers()
	ResetAdsOnDeliveryRestrictions()
	ResetAdTriggers()
	ResetIncludeIframeOnlyStream()
	ResetManifestName()
	ResetPlaylistType()
	ResetPlaylistWindowSeconds()
	ResetProgramDateTimeIntervalSeconds()
	ResetUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference
type jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) AdMarkers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adMarkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) AdMarkersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adMarkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) AdsOnDeliveryRestrictions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adsOnDeliveryRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) AdsOnDeliveryRestrictionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adsOnDeliveryRestrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) AdTriggers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adTriggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) AdTriggersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adTriggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) IncludeIframeOnlyStream() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeIframeOnlyStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) IncludeIframeOnlyStreamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeIframeOnlyStreamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ManifestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ManifestNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) PlaylistType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"playlistType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) PlaylistTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"playlistTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) PlaylistWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"playlistWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) PlaylistWindowSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"playlistWindowSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ProgramDateTimeIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programDateTimeIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ProgramDateTimeIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programDateTimeIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewMediapackageOriginEndpointCmafPackageHlsManifestsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackageOriginEndpointCmafPackageHlsManifestsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference{}

	_jsii_.Create(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediapackageOriginEndpointCmafPackageHlsManifestsOutputReference_Override(m MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference)SetAdMarkers(val *string) {
	if err := j.validateSetAdMarkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adMarkers",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference)SetAdsOnDeliveryRestrictions(val *string) {
	if err := j.validateSetAdsOnDeliveryRestrictionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adsOnDeliveryRestrictions",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference)SetAdTriggers(val *[]*string) {
	if err := j.validateSetAdTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adTriggers",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference)SetIncludeIframeOnlyStream(val interface{}) {
	if err := j.validateSetIncludeIframeOnlyStreamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeIframeOnlyStream",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference)SetManifestName(val *string) {
	if err := j.validateSetManifestNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestName",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference)SetPlaylistType(val *string) {
	if err := j.validateSetPlaylistTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"playlistType",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference)SetPlaylistWindowSeconds(val *float64) {
	if err := j.validateSetPlaylistWindowSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"playlistWindowSeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference)SetProgramDateTimeIntervalSeconds(val *float64) {
	if err := j.validateSetProgramDateTimeIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programDateTimeIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ResetAdMarkers() {
	_jsii_.InvokeVoid(
		m,
		"resetAdMarkers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ResetAdsOnDeliveryRestrictions() {
	_jsii_.InvokeVoid(
		m,
		"resetAdsOnDeliveryRestrictions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ResetAdTriggers() {
	_jsii_.InvokeVoid(
		m,
		"resetAdTriggers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ResetIncludeIframeOnlyStream() {
	_jsii_.InvokeVoid(
		m,
		"resetIncludeIframeOnlyStream",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ResetManifestName() {
	_jsii_.InvokeVoid(
		m,
		"resetManifestName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ResetPlaylistType() {
	_jsii_.InvokeVoid(
		m,
		"resetPlaylistType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ResetPlaylistWindowSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetPlaylistWindowSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ResetProgramDateTimeIntervalSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetProgramDateTimeIntervalSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		m,
		"resetUrl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageHlsManifestsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

