package dataawsccmediapackageoriginendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccmediapackageoriginendpoint/internal"
)

type DataAwsccMediapackageOriginEndpointDashPackageOutputReference interface {
	cdktf.ComplexObject
	AdsOnDeliveryRestrictions() *string
	AdTriggers() *[]*string
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
	Encryption() DataAwsccMediapackageOriginEndpointDashPackageEncryptionOutputReference
	// Experimental.
	Fqn() *string
	IncludeIframeOnlyStream() cdktf.IResolvable
	InternalValue() *DataAwsccMediapackageOriginEndpointDashPackage
	SetInternalValue(val *DataAwsccMediapackageOriginEndpointDashPackage)
	ManifestLayout() *string
	ManifestWindowSeconds() *float64
	MinBufferTimeSeconds() *float64
	MinUpdatePeriodSeconds() *float64
	PeriodTriggers() *[]*string
	Profile() *string
	SegmentDurationSeconds() *float64
	SegmentTemplateFormat() *string
	StreamSelection() DataAwsccMediapackageOriginEndpointDashPackageStreamSelectionOutputReference
	SuggestedPresentationDelaySeconds() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UtcTiming() *string
	UtcTimingUri() *string
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

// The jsii proxy struct for DataAwsccMediapackageOriginEndpointDashPackageOutputReference
type jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) AdsOnDeliveryRestrictions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adsOnDeliveryRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) AdTriggers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adTriggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) Encryption() DataAwsccMediapackageOriginEndpointDashPackageEncryptionOutputReference {
	var returns DataAwsccMediapackageOriginEndpointDashPackageEncryptionOutputReference
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) IncludeIframeOnlyStream() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"includeIframeOnlyStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) InternalValue() *DataAwsccMediapackageOriginEndpointDashPackage {
	var returns *DataAwsccMediapackageOriginEndpointDashPackage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) ManifestLayout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestLayout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) ManifestWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"manifestWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) MinBufferTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minBufferTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) MinUpdatePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minUpdatePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) PeriodTriggers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"periodTriggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) SegmentDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) SegmentTemplateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentTemplateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) StreamSelection() DataAwsccMediapackageOriginEndpointDashPackageStreamSelectionOutputReference {
	var returns DataAwsccMediapackageOriginEndpointDashPackageStreamSelectionOutputReference
	_jsii_.Get(
		j,
		"streamSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) SuggestedPresentationDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suggestedPresentationDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) UtcTiming() *string {
	var returns *string
	_jsii_.Get(
		j,
		"utcTiming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) UtcTimingUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"utcTimingUri",
		&returns,
	)
	return returns
}


func NewDataAwsccMediapackageOriginEndpointDashPackageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccMediapackageOriginEndpointDashPackageOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccMediapackageOriginEndpointDashPackageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccMediapackageOriginEndpoint.DataAwsccMediapackageOriginEndpointDashPackageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccMediapackageOriginEndpointDashPackageOutputReference_Override(d DataAwsccMediapackageOriginEndpointDashPackageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccMediapackageOriginEndpoint.DataAwsccMediapackageOriginEndpointDashPackageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference)SetInternalValue(val *DataAwsccMediapackageOriginEndpointDashPackage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccMediapackageOriginEndpointDashPackageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

