package datazonedatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/datazonedatasource/internal"
)

type DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference interface {
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
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	FilterExpressions() DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList
	FilterExpressionsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurations
	SetInternalValue(val *DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurations)
	SchemaName() *string
	SetSchemaName(val *string)
	SchemaNameInput() *string
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
	PutFilterExpressions(value interface{})
	ResetFilterExpressions()
	ResetSchemaName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference
type jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) FilterExpressions() DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList {
	var returns DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList
	_jsii_.Get(
		j,
		"filterExpressions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) FilterExpressionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterExpressionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) InternalValue() *DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurations {
	var returns *DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) SchemaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewDatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference{}

	_jsii_.Create(
		"awscc.datazoneDataSource.DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference_Override(d DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.datazoneDataSource.DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference)SetInternalValue(val *DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference)SetSchemaName(val *string) {
	if err := j.validateSetSchemaNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaName",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) PutFilterExpressions(value interface{}) {
	if err := d.validatePutFilterExpressionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilterExpressions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) ResetFilterExpressions() {
	_jsii_.InvokeVoid(
		d,
		"resetFilterExpressions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) ResetSchemaName() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemaName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

