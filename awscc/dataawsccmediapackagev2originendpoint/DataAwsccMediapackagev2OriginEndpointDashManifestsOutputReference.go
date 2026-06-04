package dataawsccmediapackagev2originendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccmediapackagev2originendpoint/internal"
)

type DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference interface {
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
	FilterConfiguration() DataAwsccMediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccMediapackagev2OriginEndpointDashManifests
	SetInternalValue(val *DataAwsccMediapackagev2OriginEndpointDashManifests)
	ManifestName() *string
	ManifestWindowSeconds() *float64
	MinBufferTimeSeconds() *float64
	MinUpdatePeriodSeconds() *float64
	PeriodTriggers() *[]*string
	ScteDash() DataAwsccMediapackagev2OriginEndpointDashManifestsScteDashOutputReference
	SegmentTemplateFormat() *string
	SuggestedPresentationDelaySeconds() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UtcTiming() DataAwsccMediapackagev2OriginEndpointDashManifestsUtcTimingOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference
type jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) DrmSignaling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drmSignaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) FilterConfiguration() DataAwsccMediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference {
	var returns DataAwsccMediapackagev2OriginEndpointDashManifestsFilterConfigurationOutputReference
	_jsii_.Get(
		j,
		"filterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) InternalValue() *DataAwsccMediapackagev2OriginEndpointDashManifests {
	var returns *DataAwsccMediapackagev2OriginEndpointDashManifests
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) ManifestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) ManifestWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"manifestWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) MinBufferTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minBufferTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) MinUpdatePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minUpdatePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) PeriodTriggers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"periodTriggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) ScteDash() DataAwsccMediapackagev2OriginEndpointDashManifestsScteDashOutputReference {
	var returns DataAwsccMediapackagev2OriginEndpointDashManifestsScteDashOutputReference
	_jsii_.Get(
		j,
		"scteDash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) SegmentTemplateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentTemplateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) SuggestedPresentationDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suggestedPresentationDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) UtcTiming() DataAwsccMediapackagev2OriginEndpointDashManifestsUtcTimingOutputReference {
	var returns DataAwsccMediapackagev2OriginEndpointDashManifestsUtcTimingOutputReference
	_jsii_.Get(
		j,
		"utcTiming",
		&returns,
	)
	return returns
}


func NewDataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccMediapackagev2OriginEndpointDashManifestsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccMediapackagev2OriginEndpoint.DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference_Override(d DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccMediapackagev2OriginEndpoint.DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference)SetInternalValue(val *DataAwsccMediapackagev2OriginEndpointDashManifests) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediapackagev2OriginEndpointDashManifestsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

