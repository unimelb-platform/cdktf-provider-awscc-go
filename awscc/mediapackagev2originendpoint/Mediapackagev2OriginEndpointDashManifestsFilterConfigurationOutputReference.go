package mediapackagev2originendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediapackagev2originendpoint/internal"
)

type Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference interface {
	cdktf.ComplexObject
	ClipStartTime() *string
	SetClipStartTime(val *string)
	ClipStartTimeInput() *string
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
	End() *string
	SetEnd(val *string)
	EndInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ManifestFilter() *string
	SetManifestFilter(val *string)
	ManifestFilterInput() *string
	Start() *string
	SetStart(val *string)
	StartInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeDelaySeconds() *float64
	SetTimeDelaySeconds(val *float64)
	TimeDelaySecondsInput() *float64
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
	ResetClipStartTime()
	ResetEnd()
	ResetManifestFilter()
	ResetStart()
	ResetTimeDelaySeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference
type jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) ClipStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clipStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) ClipStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clipStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) End() *string {
	var returns *string
	_jsii_.Get(
		j,
		"end",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) EndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) ManifestFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) ManifestFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) Start() *string {
	var returns *string
	_jsii_.Get(
		j,
		"start",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) StartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) TimeDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) TimeDelaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeDelaySecondsInput",
		&returns,
	)
	return returns
}


func NewMediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference_Override(m Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference)SetClipStartTime(val *string) {
	if err := j.validateSetClipStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clipStartTime",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference)SetEnd(val *string) {
	if err := j.validateSetEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"end",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference)SetManifestFilter(val *string) {
	if err := j.validateSetManifestFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestFilter",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference)SetStart(val *string) {
	if err := j.validateSetStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"start",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference)SetTimeDelaySeconds(val *float64) {
	if err := j.validateSetTimeDelaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeDelaySeconds",
		val,
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) ResetClipStartTime() {
	_jsii_.InvokeVoid(
		m,
		"resetClipStartTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) ResetEnd() {
	_jsii_.InvokeVoid(
		m,
		"resetEnd",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) ResetManifestFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetManifestFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) ResetStart() {
	_jsii_.InvokeVoid(
		m,
		"resetStart",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) ResetTimeDelaySeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeDelaySeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_Mediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

