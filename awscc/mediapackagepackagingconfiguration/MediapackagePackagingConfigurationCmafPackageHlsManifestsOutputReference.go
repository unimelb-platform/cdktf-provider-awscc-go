package mediapackagepackagingconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediapackagepackagingconfiguration/internal"
)

type MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference interface {
	cdktf.ComplexObject
	AdMarkers() *string
	SetAdMarkers(val *string)
	AdMarkersInput() *string
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
	IncludeIframeOnlyStream() interface{}
	SetIncludeIframeOnlyStream(val interface{})
	IncludeIframeOnlyStreamInput() interface{}
	InternalValue() *MediapackagePackagingConfigurationCmafPackageHlsManifests
	SetInternalValue(val *MediapackagePackagingConfigurationCmafPackageHlsManifests)
	ManifestName() *string
	SetManifestName(val *string)
	ManifestNameInput() *string
	ProgramDateTimeIntervalSeconds() *float64
	SetProgramDateTimeIntervalSeconds(val *float64)
	ProgramDateTimeIntervalSecondsInput() *float64
	RepeatExtXKey() interface{}
	SetRepeatExtXKey(val interface{})
	RepeatExtXKeyInput() interface{}
	StreamSelection() MediapackagePackagingConfigurationCmafPackageHlsManifestsStreamSelectionOutputReference
	StreamSelectionInput() interface{}
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
	PutStreamSelection(value *MediapackagePackagingConfigurationCmafPackageHlsManifestsStreamSelection)
	ResetAdMarkers()
	ResetIncludeIframeOnlyStream()
	ResetManifestName()
	ResetProgramDateTimeIntervalSeconds()
	ResetRepeatExtXKey()
	ResetStreamSelection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference
type jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) AdMarkers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adMarkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) AdMarkersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adMarkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) IncludeIframeOnlyStream() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeIframeOnlyStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) IncludeIframeOnlyStreamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeIframeOnlyStreamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) InternalValue() *MediapackagePackagingConfigurationCmafPackageHlsManifests {
	var returns *MediapackagePackagingConfigurationCmafPackageHlsManifests
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) ManifestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) ManifestNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) ProgramDateTimeIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programDateTimeIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) ProgramDateTimeIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programDateTimeIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) RepeatExtXKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repeatExtXKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) RepeatExtXKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repeatExtXKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) StreamSelection() MediapackagePackagingConfigurationCmafPackageHlsManifestsStreamSelectionOutputReference {
	var returns MediapackagePackagingConfigurationCmafPackageHlsManifestsStreamSelectionOutputReference
	_jsii_.Get(
		j,
		"streamSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) StreamSelectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference{}

	_jsii_.Create(
		"awscc.mediapackagePackagingConfiguration.MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference_Override(m MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediapackagePackagingConfiguration.MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference)SetAdMarkers(val *string) {
	if err := j.validateSetAdMarkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adMarkers",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference)SetIncludeIframeOnlyStream(val interface{}) {
	if err := j.validateSetIncludeIframeOnlyStreamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeIframeOnlyStream",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference)SetInternalValue(val *MediapackagePackagingConfigurationCmafPackageHlsManifests) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference)SetManifestName(val *string) {
	if err := j.validateSetManifestNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestName",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference)SetProgramDateTimeIntervalSeconds(val *float64) {
	if err := j.validateSetProgramDateTimeIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programDateTimeIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference)SetRepeatExtXKey(val interface{}) {
	if err := j.validateSetRepeatExtXKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repeatExtXKey",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) PutStreamSelection(value *MediapackagePackagingConfigurationCmafPackageHlsManifestsStreamSelection) {
	if err := m.validatePutStreamSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStreamSelection",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) ResetAdMarkers() {
	_jsii_.InvokeVoid(
		m,
		"resetAdMarkers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) ResetIncludeIframeOnlyStream() {
	_jsii_.InvokeVoid(
		m,
		"resetIncludeIframeOnlyStream",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) ResetManifestName() {
	_jsii_.InvokeVoid(
		m,
		"resetManifestName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) ResetProgramDateTimeIntervalSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetProgramDateTimeIntervalSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) ResetRepeatExtXKey() {
	_jsii_.InvokeVoid(
		m,
		"resetRepeatExtXKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) ResetStreamSelection() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamSelection",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationCmafPackageHlsManifestsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

