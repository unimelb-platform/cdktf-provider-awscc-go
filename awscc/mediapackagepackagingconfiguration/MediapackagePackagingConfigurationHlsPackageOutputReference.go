package mediapackagepackagingconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediapackagepackagingconfiguration/internal"
)

type MediapackagePackagingConfigurationHlsPackageOutputReference interface {
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
	Encryption() MediapackagePackagingConfigurationHlsPackageEncryptionOutputReference
	EncryptionInput() interface{}
	// Experimental.
	Fqn() *string
	HlsManifests() MediapackagePackagingConfigurationHlsPackageHlsManifestsList
	HlsManifestsInput() interface{}
	IncludeDvbSubtitles() interface{}
	SetIncludeDvbSubtitles(val interface{})
	IncludeDvbSubtitlesInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SegmentDurationSeconds() *float64
	SetSegmentDurationSeconds(val *float64)
	SegmentDurationSecondsInput() *float64
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
	PutEncryption(value *MediapackagePackagingConfigurationHlsPackageEncryption)
	PutHlsManifests(value interface{})
	ResetEncryption()
	ResetIncludeDvbSubtitles()
	ResetSegmentDurationSeconds()
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

// The jsii proxy struct for MediapackagePackagingConfigurationHlsPackageOutputReference
type jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) Encryption() MediapackagePackagingConfigurationHlsPackageEncryptionOutputReference {
	var returns MediapackagePackagingConfigurationHlsPackageEncryptionOutputReference
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) EncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) HlsManifests() MediapackagePackagingConfigurationHlsPackageHlsManifestsList {
	var returns MediapackagePackagingConfigurationHlsPackageHlsManifestsList
	_jsii_.Get(
		j,
		"hlsManifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) HlsManifestsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsManifestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) IncludeDvbSubtitles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeDvbSubtitles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) IncludeDvbSubtitlesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeDvbSubtitlesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) SegmentDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) SegmentDurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentDurationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) UseAudioRenditionGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAudioRenditionGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) UseAudioRenditionGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAudioRenditionGroupInput",
		&returns,
	)
	return returns
}


func NewMediapackagePackagingConfigurationHlsPackageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediapackagePackagingConfigurationHlsPackageOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackagePackagingConfigurationHlsPackageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference{}

	_jsii_.Create(
		"awscc.mediapackagePackagingConfiguration.MediapackagePackagingConfigurationHlsPackageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediapackagePackagingConfigurationHlsPackageOutputReference_Override(m MediapackagePackagingConfigurationHlsPackageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediapackagePackagingConfiguration.MediapackagePackagingConfigurationHlsPackageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference)SetIncludeDvbSubtitles(val interface{}) {
	if err := j.validateSetIncludeDvbSubtitlesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeDvbSubtitles",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference)SetSegmentDurationSeconds(val *float64) {
	if err := j.validateSetSegmentDurationSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference)SetUseAudioRenditionGroup(val interface{}) {
	if err := j.validateSetUseAudioRenditionGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAudioRenditionGroup",
		val,
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) PutEncryption(value *MediapackagePackagingConfigurationHlsPackageEncryption) {
	if err := m.validatePutEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncryption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) PutHlsManifests(value interface{}) {
	if err := m.validatePutHlsManifestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putHlsManifests",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) ResetEncryption() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryption",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) ResetIncludeDvbSubtitles() {
	_jsii_.InvokeVoid(
		m,
		"resetIncludeDvbSubtitles",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) ResetSegmentDurationSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentDurationSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) ResetUseAudioRenditionGroup() {
	_jsii_.InvokeVoid(
		m,
		"resetUseAudioRenditionGroup",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

