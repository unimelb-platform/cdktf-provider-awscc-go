package dataawscciotwirelessserviceprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscciotwirelessserviceprofile/internal"
)

type DataAwsccIotwirelessServiceProfileLoRaWanOutputReference interface {
	cdktf.ComplexObject
	AddGwMetadata() cdktf.IResolvable
	ChannelMask() *string
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
	DevStatusReqFreq() *float64
	DlBucketSize() *float64
	DlRate() *float64
	DlRatePolicy() *string
	DrMax() *float64
	DrMin() *float64
	// Experimental.
	Fqn() *string
	HrAllowed() cdktf.IResolvable
	InternalValue() *DataAwsccIotwirelessServiceProfileLoRaWan
	SetInternalValue(val *DataAwsccIotwirelessServiceProfileLoRaWan)
	MinGwDiversity() *float64
	NwkGeoLoc() cdktf.IResolvable
	PrAllowed() cdktf.IResolvable
	RaAllowed() cdktf.IResolvable
	ReportDevStatusBattery() cdktf.IResolvable
	ReportDevStatusMargin() cdktf.IResolvable
	TargetPer() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UlBucketSize() *float64
	UlRate() *float64
	UlRatePolicy() *string
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

// The jsii proxy struct for DataAwsccIotwirelessServiceProfileLoRaWanOutputReference
type jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) AddGwMetadata() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"addGwMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) ChannelMask() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelMask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) DevStatusReqFreq() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"devStatusReqFreq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) DlBucketSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dlBucketSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) DlRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dlRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) DlRatePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dlRatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) DrMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) DrMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) HrAllowed() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hrAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) InternalValue() *DataAwsccIotwirelessServiceProfileLoRaWan {
	var returns *DataAwsccIotwirelessServiceProfileLoRaWan
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) MinGwDiversity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minGwDiversity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) NwkGeoLoc() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"nwkGeoLoc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) PrAllowed() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"prAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) RaAllowed() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"raAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) ReportDevStatusBattery() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reportDevStatusBattery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) ReportDevStatusMargin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reportDevStatusMargin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) TargetPer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) UlBucketSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ulBucketSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) UlRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ulRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) UlRatePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ulRatePolicy",
		&returns,
	)
	return returns
}


func NewDataAwsccIotwirelessServiceProfileLoRaWanOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccIotwirelessServiceProfileLoRaWanOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccIotwirelessServiceProfileLoRaWanOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccIotwirelessServiceProfile.DataAwsccIotwirelessServiceProfileLoRaWanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccIotwirelessServiceProfileLoRaWanOutputReference_Override(d DataAwsccIotwirelessServiceProfileLoRaWanOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccIotwirelessServiceProfile.DataAwsccIotwirelessServiceProfileLoRaWanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference)SetInternalValue(val *DataAwsccIotwirelessServiceProfileLoRaWan) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccIotwirelessServiceProfileLoRaWanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

