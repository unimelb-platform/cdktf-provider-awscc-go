package mediapackagepackagingconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediapackagepackagingconfiguration/internal"
)

type MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference interface {
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
	InternalValue() *MediapackagePackagingConfigurationDashPackageDashManifests
	SetInternalValue(val *MediapackagePackagingConfigurationDashPackageDashManifests)
	ManifestLayout() *string
	SetManifestLayout(val *string)
	ManifestLayoutInput() *string
	ManifestName() *string
	SetManifestName(val *string)
	ManifestNameInput() *string
	MinBufferTimeSeconds() *float64
	SetMinBufferTimeSeconds(val *float64)
	MinBufferTimeSecondsInput() *float64
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
	ScteMarkersSource() *string
	SetScteMarkersSource(val *string)
	ScteMarkersSourceInput() *string
	StreamSelection() MediapackagePackagingConfigurationDashPackageDashManifestsStreamSelectionOutputReference
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
	PutStreamSelection(value *MediapackagePackagingConfigurationDashPackageDashManifestsStreamSelection)
	ResetManifestLayout()
	ResetManifestName()
	ResetMinBufferTimeSeconds()
	ResetProfile()
	ResetScteMarkersSource()
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

// The jsii proxy struct for MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference
type jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) InternalValue() *MediapackagePackagingConfigurationDashPackageDashManifests {
	var returns *MediapackagePackagingConfigurationDashPackageDashManifests
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ManifestLayout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestLayout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ManifestLayoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestLayoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ManifestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ManifestNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) MinBufferTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minBufferTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) MinBufferTimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minBufferTimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ScteMarkersSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scteMarkersSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ScteMarkersSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scteMarkersSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) StreamSelection() MediapackagePackagingConfigurationDashPackageDashManifestsStreamSelectionOutputReference {
	var returns MediapackagePackagingConfigurationDashPackageDashManifestsStreamSelectionOutputReference
	_jsii_.Get(
		j,
		"streamSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) StreamSelectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediapackagePackagingConfigurationDashPackageDashManifestsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackagePackagingConfigurationDashPackageDashManifestsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference{}

	_jsii_.Create(
		"awscc.mediapackagePackagingConfiguration.MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediapackagePackagingConfigurationDashPackageDashManifestsOutputReference_Override(m MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediapackagePackagingConfiguration.MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference)SetInternalValue(val *MediapackagePackagingConfigurationDashPackageDashManifests) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference)SetManifestLayout(val *string) {
	if err := j.validateSetManifestLayoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestLayout",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference)SetManifestName(val *string) {
	if err := j.validateSetManifestNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestName",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference)SetMinBufferTimeSeconds(val *float64) {
	if err := j.validateSetMinBufferTimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minBufferTimeSeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference)SetProfile(val *string) {
	if err := j.validateSetProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference)SetScteMarkersSource(val *string) {
	if err := j.validateSetScteMarkersSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scteMarkersSource",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) PutStreamSelection(value *MediapackagePackagingConfigurationDashPackageDashManifestsStreamSelection) {
	if err := m.validatePutStreamSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStreamSelection",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ResetManifestLayout() {
	_jsii_.InvokeVoid(
		m,
		"resetManifestLayout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ResetManifestName() {
	_jsii_.InvokeVoid(
		m,
		"resetManifestName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ResetMinBufferTimeSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetMinBufferTimeSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ResetProfile() {
	_jsii_.InvokeVoid(
		m,
		"resetProfile",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ResetScteMarkersSource() {
	_jsii_.InvokeVoid(
		m,
		"resetScteMarkersSource",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ResetStreamSelection() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamSelection",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageDashManifestsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

