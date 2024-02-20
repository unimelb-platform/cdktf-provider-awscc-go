package databrewjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/databrewjob/internal"
)

type DatabrewJobProfileConfigurationOutputReference interface {
	cdktf.ComplexObject
	ColumnStatisticsConfigurations() DatabrewJobProfileConfigurationColumnStatisticsConfigurationsList
	ColumnStatisticsConfigurationsInput() interface{}
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
	DatasetStatisticsConfiguration() DatabrewJobProfileConfigurationDatasetStatisticsConfigurationOutputReference
	DatasetStatisticsConfigurationInput() interface{}
	EntityDetectorConfiguration() DatabrewJobProfileConfigurationEntityDetectorConfigurationOutputReference
	EntityDetectorConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProfileColumns() DatabrewJobProfileConfigurationProfileColumnsList
	ProfileColumnsInput() interface{}
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
	PutColumnStatisticsConfigurations(value interface{})
	PutDatasetStatisticsConfiguration(value *DatabrewJobProfileConfigurationDatasetStatisticsConfiguration)
	PutEntityDetectorConfiguration(value *DatabrewJobProfileConfigurationEntityDetectorConfiguration)
	PutProfileColumns(value interface{})
	ResetColumnStatisticsConfigurations()
	ResetDatasetStatisticsConfiguration()
	ResetEntityDetectorConfiguration()
	ResetProfileColumns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabrewJobProfileConfigurationOutputReference
type jsiiProxy_DatabrewJobProfileConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) ColumnStatisticsConfigurations() DatabrewJobProfileConfigurationColumnStatisticsConfigurationsList {
	var returns DatabrewJobProfileConfigurationColumnStatisticsConfigurationsList
	_jsii_.Get(
		j,
		"columnStatisticsConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) ColumnStatisticsConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnStatisticsConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) DatasetStatisticsConfiguration() DatabrewJobProfileConfigurationDatasetStatisticsConfigurationOutputReference {
	var returns DatabrewJobProfileConfigurationDatasetStatisticsConfigurationOutputReference
	_jsii_.Get(
		j,
		"datasetStatisticsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) DatasetStatisticsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datasetStatisticsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) EntityDetectorConfiguration() DatabrewJobProfileConfigurationEntityDetectorConfigurationOutputReference {
	var returns DatabrewJobProfileConfigurationEntityDetectorConfigurationOutputReference
	_jsii_.Get(
		j,
		"entityDetectorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) EntityDetectorConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entityDetectorConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) ProfileColumns() DatabrewJobProfileConfigurationProfileColumnsList {
	var returns DatabrewJobProfileConfigurationProfileColumnsList
	_jsii_.Get(
		j,
		"profileColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) ProfileColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatabrewJobProfileConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabrewJobProfileConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDatabrewJobProfileConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabrewJobProfileConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.databrewJob.DatabrewJobProfileConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabrewJobProfileConfigurationOutputReference_Override(d DatabrewJobProfileConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.databrewJob.DatabrewJobProfileConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobProfileConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) PutColumnStatisticsConfigurations(value interface{}) {
	if err := d.validatePutColumnStatisticsConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putColumnStatisticsConfigurations",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) PutDatasetStatisticsConfiguration(value *DatabrewJobProfileConfigurationDatasetStatisticsConfiguration) {
	if err := d.validatePutDatasetStatisticsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDatasetStatisticsConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) PutEntityDetectorConfiguration(value *DatabrewJobProfileConfigurationEntityDetectorConfiguration) {
	if err := d.validatePutEntityDetectorConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEntityDetectorConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) PutProfileColumns(value interface{}) {
	if err := d.validatePutProfileColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProfileColumns",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) ResetColumnStatisticsConfigurations() {
	_jsii_.InvokeVoid(
		d,
		"resetColumnStatisticsConfigurations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) ResetDatasetStatisticsConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetDatasetStatisticsConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) ResetEntityDetectorConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetEntityDetectorConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) ResetProfileColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetProfileColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabrewJobProfileConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

