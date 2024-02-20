package dataawscciotmitigationaction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscciotmitigationaction/internal"
)

type DataAwsccIotMitigationActionActionParamsOutputReference interface {
	cdktf.ComplexObject
	AddThingsToThingGroupParams() DataAwsccIotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference
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
	EnableIoTLoggingParams() DataAwsccIotMitigationActionActionParamsEnableIoTLoggingParamsOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccIotMitigationActionActionParams
	SetInternalValue(val *DataAwsccIotMitigationActionActionParams)
	PublishFindingToSnsParams() DataAwsccIotMitigationActionActionParamsPublishFindingToSnsParamsOutputReference
	ReplaceDefaultPolicyVersionParams() DataAwsccIotMitigationActionActionParamsReplaceDefaultPolicyVersionParamsOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdateCaCertificateParams() DataAwsccIotMitigationActionActionParamsUpdateCaCertificateParamsOutputReference
	UpdateDeviceCertificateParams() DataAwsccIotMitigationActionActionParamsUpdateDeviceCertificateParamsOutputReference
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

// The jsii proxy struct for DataAwsccIotMitigationActionActionParamsOutputReference
type jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) AddThingsToThingGroupParams() DataAwsccIotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference {
	var returns DataAwsccIotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference
	_jsii_.Get(
		j,
		"addThingsToThingGroupParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) EnableIoTLoggingParams() DataAwsccIotMitigationActionActionParamsEnableIoTLoggingParamsOutputReference {
	var returns DataAwsccIotMitigationActionActionParamsEnableIoTLoggingParamsOutputReference
	_jsii_.Get(
		j,
		"enableIoTLoggingParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) InternalValue() *DataAwsccIotMitigationActionActionParams {
	var returns *DataAwsccIotMitigationActionActionParams
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) PublishFindingToSnsParams() DataAwsccIotMitigationActionActionParamsPublishFindingToSnsParamsOutputReference {
	var returns DataAwsccIotMitigationActionActionParamsPublishFindingToSnsParamsOutputReference
	_jsii_.Get(
		j,
		"publishFindingToSnsParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) ReplaceDefaultPolicyVersionParams() DataAwsccIotMitigationActionActionParamsReplaceDefaultPolicyVersionParamsOutputReference {
	var returns DataAwsccIotMitigationActionActionParamsReplaceDefaultPolicyVersionParamsOutputReference
	_jsii_.Get(
		j,
		"replaceDefaultPolicyVersionParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) UpdateCaCertificateParams() DataAwsccIotMitigationActionActionParamsUpdateCaCertificateParamsOutputReference {
	var returns DataAwsccIotMitigationActionActionParamsUpdateCaCertificateParamsOutputReference
	_jsii_.Get(
		j,
		"updateCaCertificateParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) UpdateDeviceCertificateParams() DataAwsccIotMitigationActionActionParamsUpdateDeviceCertificateParamsOutputReference {
	var returns DataAwsccIotMitigationActionActionParamsUpdateDeviceCertificateParamsOutputReference
	_jsii_.Get(
		j,
		"updateDeviceCertificateParams",
		&returns,
	)
	return returns
}


func NewDataAwsccIotMitigationActionActionParamsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccIotMitigationActionActionParamsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccIotMitigationActionActionParamsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccIotMitigationAction.DataAwsccIotMitigationActionActionParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccIotMitigationActionActionParamsOutputReference_Override(d DataAwsccIotMitigationActionActionParamsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccIotMitigationAction.DataAwsccIotMitigationActionActionParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference)SetInternalValue(val *DataAwsccIotMitigationActionActionParams) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccIotMitigationActionActionParamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

