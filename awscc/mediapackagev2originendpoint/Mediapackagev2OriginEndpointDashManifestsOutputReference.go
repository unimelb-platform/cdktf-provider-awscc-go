package mediapackagev2originendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediapackagev2originendpoint/internal"
)

type Mediapackagev2OriginEndpointDashManifestsOutputReference interface {
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
	DrmSignaling() *string
	SetDrmSignaling(val *string)
	DrmSignalingInput() *string
	FilterConfiguration() Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference
	FilterConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ManifestName() *string
	SetManifestName(val *string)
	ManifestNameInput() *string
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
	ScteDash() Mediapackagev2OriginEndpointDashManifestsScteDashOutputReference
	ScteDashInput() interface{}
	SegmentTemplateFormat() *string
	SetSegmentTemplateFormat(val *string)
	SegmentTemplateFormatInput() *string
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
	UtcTiming() Mediapackagev2OriginEndpointDashManifestsUtcTimingOutputReference
	UtcTimingInput() interface{}
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
	PutFilterConfiguration(value *Mediapackagev2OriginEndpointDashManifestsFilterConfiguration)
	PutScteDash(value *Mediapackagev2OriginEndpointDashManifestsScteDash)
	PutUtcTiming(value *Mediapackagev2OriginEndpointDashManifestsUtcTiming)
	ResetDrmSignaling()
	ResetFilterConfiguration()
	ResetManifestName()
	ResetManifestWindowSeconds()
	ResetMinBufferTimeSeconds()
	ResetMinUpdatePeriodSeconds()
	ResetPeriodTriggers()
	ResetScteDash()
	ResetSegmentTemplateFormat()
	ResetSuggestedPresentationDelaySeconds()
	ResetUtcTiming()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Mediapackagev2OriginEndpointDashManifestsOutputReference
type jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) DrmSignaling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drmSignaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) DrmSignalingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drmSignalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) FilterConfiguration() Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference {
	var returns Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference
	_jsii_.Get(
		j,
		"filterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) FilterConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ManifestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ManifestNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ManifestWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"manifestWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ManifestWindowSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"manifestWindowSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) MinBufferTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minBufferTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) MinBufferTimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minBufferTimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) MinUpdatePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minUpdatePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) MinUpdatePeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minUpdatePeriodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) PeriodTriggers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"periodTriggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) PeriodTriggersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"periodTriggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ScteDash() Mediapackagev2OriginEndpointDashManifestsScteDashOutputReference {
	var returns Mediapackagev2OriginEndpointDashManifestsScteDashOutputReference
	_jsii_.Get(
		j,
		"scteDash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ScteDashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scteDashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) SegmentTemplateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentTemplateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) SegmentTemplateFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentTemplateFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) SuggestedPresentationDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suggestedPresentationDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) SuggestedPresentationDelaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suggestedPresentationDelaySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) UtcTiming() Mediapackagev2OriginEndpointDashManifestsUtcTimingOutputReference {
	var returns Mediapackagev2OriginEndpointDashManifestsUtcTimingOutputReference
	_jsii_.Get(
		j,
		"utcTiming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) UtcTimingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utcTimingInput",
		&returns,
	)
	return returns
}


func NewMediapackagev2OriginEndpointDashManifestsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Mediapackagev2OriginEndpointDashManifestsOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackagev2OriginEndpointDashManifestsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference{}

	_jsii_.Create(
		"awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpointDashManifestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediapackagev2OriginEndpointDashManifestsOutputReference_Override(m Mediapackagev2OriginEndpointDashManifestsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpointDashManifestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference)SetDrmSignaling(val *string) {
	if err := j.validateSetDrmSignalingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drmSignaling",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference)SetManifestName(val *string) {
	if err := j.validateSetManifestNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestName",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference)SetManifestWindowSeconds(val *float64) {
	if err := j.validateSetManifestWindowSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestWindowSeconds",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference)SetMinBufferTimeSeconds(val *float64) {
	if err := j.validateSetMinBufferTimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minBufferTimeSeconds",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference)SetMinUpdatePeriodSeconds(val *float64) {
	if err := j.validateSetMinUpdatePeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minUpdatePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference)SetPeriodTriggers(val *[]*string) {
	if err := j.validateSetPeriodTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodTriggers",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference)SetSegmentTemplateFormat(val *string) {
	if err := j.validateSetSegmentTemplateFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentTemplateFormat",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference)SetSuggestedPresentationDelaySeconds(val *float64) {
	if err := j.validateSetSuggestedPresentationDelaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suggestedPresentationDelaySeconds",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) PutFilterConfiguration(value *Mediapackagev2OriginEndpointDashManifestsFilterConfiguration) {
	if err := m.validatePutFilterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFilterConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) PutScteDash(value *Mediapackagev2OriginEndpointDashManifestsScteDash) {
	if err := m.validatePutScteDashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putScteDash",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) PutUtcTiming(value *Mediapackagev2OriginEndpointDashManifestsUtcTiming) {
	if err := m.validatePutUtcTimingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putUtcTiming",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ResetDrmSignaling() {
	_jsii_.InvokeVoid(
		m,
		"resetDrmSignaling",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ResetFilterConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetFilterConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ResetManifestName() {
	_jsii_.InvokeVoid(
		m,
		"resetManifestName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ResetManifestWindowSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetManifestWindowSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ResetMinBufferTimeSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetMinBufferTimeSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ResetMinUpdatePeriodSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetMinUpdatePeriodSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ResetPeriodTriggers() {
	_jsii_.InvokeVoid(
		m,
		"resetPeriodTriggers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ResetScteDash() {
	_jsii_.InvokeVoid(
		m,
		"resetScteDash",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ResetSegmentTemplateFormat() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentTemplateFormat",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ResetSuggestedPresentationDelaySeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetSuggestedPresentationDelaySeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ResetUtcTiming() {
	_jsii_.InvokeVoid(
		m,
		"resetUtcTiming",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

