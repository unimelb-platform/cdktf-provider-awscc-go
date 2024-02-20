package databrewjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/databrewjob/internal"
)

type DatabrewJobDataCatalogOutputsOutputReference interface {
	cdktf.ComplexObject
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
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
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	DatabaseOptions() DatabrewJobDataCatalogOutputsDatabaseOptionsOutputReference
	DatabaseOptionsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Overwrite() interface{}
	SetOverwrite(val interface{})
	OverwriteInput() interface{}
	S3Options() DatabrewJobDataCatalogOutputsS3OptionsOutputReference
	S3OptionsInput() interface{}
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
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
	PutDatabaseOptions(value *DatabrewJobDataCatalogOutputsDatabaseOptions)
	PutS3Options(value *DatabrewJobDataCatalogOutputsS3Options)
	ResetCatalogId()
	ResetDatabaseOptions()
	ResetOverwrite()
	ResetS3Options()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabrewJobDataCatalogOutputsOutputReference
type jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) DatabaseOptions() DatabrewJobDataCatalogOutputsDatabaseOptionsOutputReference {
	var returns DatabrewJobDataCatalogOutputsDatabaseOptionsOutputReference
	_jsii_.Get(
		j,
		"databaseOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) DatabaseOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) Overwrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overwrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) OverwriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overwriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) S3Options() DatabrewJobDataCatalogOutputsS3OptionsOutputReference {
	var returns DatabrewJobDataCatalogOutputsS3OptionsOutputReference
	_jsii_.Get(
		j,
		"s3Options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) S3OptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3OptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatabrewJobDataCatalogOutputsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DatabrewJobDataCatalogOutputsOutputReference {
	_init_.Initialize()

	if err := validateNewDatabrewJobDataCatalogOutputsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference{}

	_jsii_.Create(
		"awscc.databrewJob.DatabrewJobDataCatalogOutputsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDatabrewJobDataCatalogOutputsOutputReference_Override(d DatabrewJobDataCatalogOutputsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.databrewJob.DatabrewJobDataCatalogOutputsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference)SetCatalogId(val *string) {
	if err := j.validateSetCatalogIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference)SetOverwrite(val interface{}) {
	if err := j.validateSetOverwriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overwrite",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) PutDatabaseOptions(value *DatabrewJobDataCatalogOutputsDatabaseOptions) {
	if err := d.validatePutDatabaseOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDatabaseOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) PutS3Options(value *DatabrewJobDataCatalogOutputsS3Options) {
	if err := d.validatePutS3OptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putS3Options",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) ResetCatalogId() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) ResetDatabaseOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) ResetOverwrite() {
	_jsii_.InvokeVoid(
		d,
		"resetOverwrite",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) ResetS3Options() {
	_jsii_.InvokeVoid(
		d,
		"resetS3Options",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabrewJobDataCatalogOutputsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

