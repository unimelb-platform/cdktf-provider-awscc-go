package dataawsccathenaworkgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccathenaworkgroup/internal"
)

type DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference interface {
	cdktf.ComplexObject
	AdditionalConfiguration() *string
	BytesScannedCutoffPerQuery() *float64
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
	CustomerContentEncryptionConfiguration() DataAwsccAthenaWorkGroupWorkGroupConfigurationCustomerContentEncryptionConfigurationOutputReference
	EnforceWorkGroupConfiguration() cdktf.IResolvable
	EngineVersion() DataAwsccAthenaWorkGroupWorkGroupConfigurationEngineVersionOutputReference
	ExecutionRole() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccAthenaWorkGroupWorkGroupConfiguration
	SetInternalValue(val *DataAwsccAthenaWorkGroupWorkGroupConfiguration)
	PublishCloudwatchMetricsEnabled() cdktf.IResolvable
	RequesterPaysEnabled() cdktf.IResolvable
	ResultConfiguration() DataAwsccAthenaWorkGroupWorkGroupConfigurationResultConfigurationOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference
type jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) AdditionalConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) BytesScannedCutoffPerQuery() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesScannedCutoffPerQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) CustomerContentEncryptionConfiguration() DataAwsccAthenaWorkGroupWorkGroupConfigurationCustomerContentEncryptionConfigurationOutputReference {
	var returns DataAwsccAthenaWorkGroupWorkGroupConfigurationCustomerContentEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"customerContentEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) EnforceWorkGroupConfiguration() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enforceWorkGroupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) EngineVersion() DataAwsccAthenaWorkGroupWorkGroupConfigurationEngineVersionOutputReference {
	var returns DataAwsccAthenaWorkGroupWorkGroupConfigurationEngineVersionOutputReference
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) InternalValue() *DataAwsccAthenaWorkGroupWorkGroupConfiguration {
	var returns *DataAwsccAthenaWorkGroupWorkGroupConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) PublishCloudwatchMetricsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"publishCloudwatchMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) RequesterPaysEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"requesterPaysEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) ResultConfiguration() DataAwsccAthenaWorkGroupWorkGroupConfigurationResultConfigurationOutputReference {
	var returns DataAwsccAthenaWorkGroupWorkGroupConfigurationResultConfigurationOutputReference
	_jsii_.Get(
		j,
		"resultConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccAthenaWorkGroup.DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference_Override(d DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccAthenaWorkGroup.DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference)SetInternalValue(val *DataAwsccAthenaWorkGroupWorkGroupConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccAthenaWorkGroupWorkGroupConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

