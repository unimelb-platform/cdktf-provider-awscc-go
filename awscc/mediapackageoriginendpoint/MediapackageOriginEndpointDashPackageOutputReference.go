package mediapackageoriginendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediapackageoriginendpoint/internal"
)

type MediapackageOriginEndpointDashPackageOutputReference interface {
	cdktf.ComplexObject
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
	Encryption() MediapackageOriginEndpointDashPackageEncryptionOutputReference
	EncryptionInput() interface{}
	// Experimental.
	Fqn() *string
	IncludeIframeOnlyStream() interface{}
	SetIncludeIframeOnlyStream(val interface{})
	IncludeIframeOnlyStreamInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ManifestLayout() *string
	SetManifestLayout(val *string)
	ManifestLayoutInput() *string
	ManifestWindowSeconds() *float64
	SetManifestWindowSeconds(val *float64)
	ManifestWindowSecondsInput() *float64
	MinBufferTimeSeconds() *float64
	SetMinBufferTimeSeconds(val *float64)
	MinBufferTimeSecondsInput() *float64
	MinUpdatePeriodSeconds() *float64
	SetMinUpdatePeriodSeconds(val *float64)
	MinUpdatePeriodSecondsInput() *float64
	PeriodTriggers() *[]*string
	SetPeriodTriggers(val *[]*string)
	PeriodTriggersInput() *[]*string
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
	SegmentDurationSeconds() *float64
	SetSegmentDurationSeconds(val *float64)
	SegmentDurationSecondsInput() *float64
	SegmentTemplateFormat() *string
	SetSegmentTemplateFormat(val *string)
	SegmentTemplateFormatInput() *string
	StreamSelection() MediapackageOriginEndpointDashPackageStreamSelectionOutputReference
	StreamSelectionInput() interface{}
	SuggestedPresentationDelaySeconds() *float64
	SetSuggestedPresentationDelaySeconds(val *float64)
	SuggestedPresentationDelaySecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UtcTiming() *string
	SetUtcTiming(val *string)
	UtcTimingInput() *string
	UtcTimingUri() *string
	SetUtcTimingUri(val *string)
	UtcTimingUriInput() *string
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
	PutEncryption(value *MediapackageOriginEndpointDashPackageEncryption)
	PutStreamSelection(value *MediapackageOriginEndpointDashPackageStreamSelection)
	ResetAdsOnDeliveryRestrictions()
	ResetAdTriggers()
	ResetEncryption()
	ResetIncludeIframeOnlyStream()
	ResetManifestLayout()
	ResetManifestWindowSeconds()
	ResetMinBufferTimeSeconds()
	ResetMinUpdatePeriodSeconds()
	ResetPeriodTriggers()
	ResetProfile()
	ResetSegmentDurationSeconds()
	ResetSegmentTemplateFormat()
	ResetStreamSelection()
	ResetSuggestedPresentationDelaySeconds()
	ResetUtcTiming()
	ResetUtcTimingUri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediapackageOriginEndpointDashPackageOutputReference
type jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) AdsOnDeliveryRestrictions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adsOnDeliveryRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) AdsOnDeliveryRestrictionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adsOnDeliveryRestrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) AdTriggers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adTriggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) AdTriggersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adTriggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) Encryption() MediapackageOriginEndpointDashPackageEncryptionOutputReference {
	var returns MediapackageOriginEndpointDashPackageEncryptionOutputReference
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) EncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) IncludeIframeOnlyStream() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeIframeOnlyStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) IncludeIframeOnlyStreamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeIframeOnlyStreamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ManifestLayout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestLayout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ManifestLayoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestLayoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ManifestWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"manifestWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ManifestWindowSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"manifestWindowSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) MinBufferTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minBufferTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) MinBufferTimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minBufferTimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) MinUpdatePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minUpdatePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) MinUpdatePeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minUpdatePeriodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) PeriodTriggers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"periodTriggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) PeriodTriggersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"periodTriggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) SegmentDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) SegmentDurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentDurationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) SegmentTemplateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentTemplateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) SegmentTemplateFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentTemplateFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) StreamSelection() MediapackageOriginEndpointDashPackageStreamSelectionOutputReference {
	var returns MediapackageOriginEndpointDashPackageStreamSelectionOutputReference
	_jsii_.Get(
		j,
		"streamSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) StreamSelectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) SuggestedPresentationDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suggestedPresentationDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) SuggestedPresentationDelaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suggestedPresentationDelaySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) UtcTiming() *string {
	var returns *string
	_jsii_.Get(
		j,
		"utcTiming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) UtcTimingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"utcTimingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) UtcTimingUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"utcTimingUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) UtcTimingUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"utcTimingUriInput",
		&returns,
	)
	return returns
}


func NewMediapackageOriginEndpointDashPackageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediapackageOriginEndpointDashPackageOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackageOriginEndpointDashPackageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference{}

	_jsii_.Create(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpointDashPackageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediapackageOriginEndpointDashPackageOutputReference_Override(m MediapackageOriginEndpointDashPackageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpointDashPackageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetAdsOnDeliveryRestrictions(val *string) {
	if err := j.validateSetAdsOnDeliveryRestrictionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adsOnDeliveryRestrictions",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetAdTriggers(val *[]*string) {
	if err := j.validateSetAdTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adTriggers",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetIncludeIframeOnlyStream(val interface{}) {
	if err := j.validateSetIncludeIframeOnlyStreamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeIframeOnlyStream",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetManifestLayout(val *string) {
	if err := j.validateSetManifestLayoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestLayout",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetManifestWindowSeconds(val *float64) {
	if err := j.validateSetManifestWindowSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestWindowSeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetMinBufferTimeSeconds(val *float64) {
	if err := j.validateSetMinBufferTimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minBufferTimeSeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetMinUpdatePeriodSeconds(val *float64) {
	if err := j.validateSetMinUpdatePeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minUpdatePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetPeriodTriggers(val *[]*string) {
	if err := j.validateSetPeriodTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodTriggers",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetProfile(val *string) {
	if err := j.validateSetProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetSegmentDurationSeconds(val *float64) {
	if err := j.validateSetSegmentDurationSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetSegmentTemplateFormat(val *string) {
	if err := j.validateSetSegmentTemplateFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentTemplateFormat",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetSuggestedPresentationDelaySeconds(val *float64) {
	if err := j.validateSetSuggestedPresentationDelaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suggestedPresentationDelaySeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetUtcTiming(val *string) {
	if err := j.validateSetUtcTimingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"utcTiming",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference)SetUtcTimingUri(val *string) {
	if err := j.validateSetUtcTimingUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"utcTimingUri",
		val,
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) PutEncryption(value *MediapackageOriginEndpointDashPackageEncryption) {
	if err := m.validatePutEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncryption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) PutStreamSelection(value *MediapackageOriginEndpointDashPackageStreamSelection) {
	if err := m.validatePutStreamSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStreamSelection",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ResetAdsOnDeliveryRestrictions() {
	_jsii_.InvokeVoid(
		m,
		"resetAdsOnDeliveryRestrictions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ResetAdTriggers() {
	_jsii_.InvokeVoid(
		m,
		"resetAdTriggers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ResetEncryption() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryption",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ResetIncludeIframeOnlyStream() {
	_jsii_.InvokeVoid(
		m,
		"resetIncludeIframeOnlyStream",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ResetManifestLayout() {
	_jsii_.InvokeVoid(
		m,
		"resetManifestLayout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ResetManifestWindowSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetManifestWindowSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ResetMinBufferTimeSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetMinBufferTimeSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ResetMinUpdatePeriodSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetMinUpdatePeriodSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ResetPeriodTriggers() {
	_jsii_.InvokeVoid(
		m,
		"resetPeriodTriggers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ResetProfile() {
	_jsii_.InvokeVoid(
		m,
		"resetProfile",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ResetSegmentDurationSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentDurationSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ResetSegmentTemplateFormat() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentTemplateFormat",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ResetStreamSelection() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamSelection",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ResetSuggestedPresentationDelaySeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetSuggestedPresentationDelaySeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ResetUtcTiming() {
	_jsii_.InvokeVoid(
		m,
		"resetUtcTiming",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ResetUtcTimingUri() {
	_jsii_.InvokeVoid(
		m,
		"resetUtcTimingUri",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

