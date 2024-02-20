package dataawsccquicksightdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccquicksightdatasource/internal"
)

type DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference interface {
	cdktf.ComplexObject
	AmazonElasticsearchParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersAmazonElasticsearchParametersOutputReference
	AmazonOpenSearchParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersAmazonOpenSearchParametersOutputReference
	AthenaParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersAthenaParametersOutputReference
	AuroraParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersAuroraParametersOutputReference
	AuroraPostgreSqlParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersAuroraPostgreSqlParametersOutputReference
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
	DatabricksParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersDatabricksParametersOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccQuicksightDataSourceAlternateDataSourceParameters
	SetInternalValue(val *DataAwsccQuicksightDataSourceAlternateDataSourceParameters)
	MariaDbParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersMariaDbParametersOutputReference
	MySqlParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersMySqlParametersOutputReference
	OracleParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersOracleParametersOutputReference
	PostgreSqlParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersPostgreSqlParametersOutputReference
	PrestoParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersPrestoParametersOutputReference
	RdsParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersRdsParametersOutputReference
	RedshiftParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersRedshiftParametersOutputReference
	S3Parameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersS3ParametersOutputReference
	SnowflakeParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference
	SparkParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersSparkParametersOutputReference
	SqlServerParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersSqlServerParametersOutputReference
	StarburstParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersStarburstParametersOutputReference
	TeradataParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersTeradataParametersOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrinoParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersTrinoParametersOutputReference
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

// The jsii proxy struct for DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference
type jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) AmazonElasticsearchParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersAmazonElasticsearchParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersAmazonElasticsearchParametersOutputReference
	_jsii_.Get(
		j,
		"amazonElasticsearchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) AmazonOpenSearchParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersAmazonOpenSearchParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersAmazonOpenSearchParametersOutputReference
	_jsii_.Get(
		j,
		"amazonOpenSearchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) AthenaParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersAthenaParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersAthenaParametersOutputReference
	_jsii_.Get(
		j,
		"athenaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) AuroraParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersAuroraParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersAuroraParametersOutputReference
	_jsii_.Get(
		j,
		"auroraParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) AuroraPostgreSqlParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersAuroraPostgreSqlParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersAuroraPostgreSqlParametersOutputReference
	_jsii_.Get(
		j,
		"auroraPostgreSqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) DatabricksParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersDatabricksParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersDatabricksParametersOutputReference
	_jsii_.Get(
		j,
		"databricksParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) InternalValue() *DataAwsccQuicksightDataSourceAlternateDataSourceParameters {
	var returns *DataAwsccQuicksightDataSourceAlternateDataSourceParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) MariaDbParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersMariaDbParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersMariaDbParametersOutputReference
	_jsii_.Get(
		j,
		"mariaDbParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) MySqlParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersMySqlParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersMySqlParametersOutputReference
	_jsii_.Get(
		j,
		"mySqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) OracleParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersOracleParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersOracleParametersOutputReference
	_jsii_.Get(
		j,
		"oracleParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) PostgreSqlParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersPostgreSqlParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersPostgreSqlParametersOutputReference
	_jsii_.Get(
		j,
		"postgreSqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) PrestoParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersPrestoParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersPrestoParametersOutputReference
	_jsii_.Get(
		j,
		"prestoParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) RdsParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersRdsParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersRdsParametersOutputReference
	_jsii_.Get(
		j,
		"rdsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) RedshiftParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersRedshiftParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersRedshiftParametersOutputReference
	_jsii_.Get(
		j,
		"redshiftParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) S3Parameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersS3ParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersS3ParametersOutputReference
	_jsii_.Get(
		j,
		"s3Parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) SnowflakeParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference
	_jsii_.Get(
		j,
		"snowflakeParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) SparkParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersSparkParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersSparkParametersOutputReference
	_jsii_.Get(
		j,
		"sparkParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) SqlServerParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersSqlServerParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersSqlServerParametersOutputReference
	_jsii_.Get(
		j,
		"sqlServerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) StarburstParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersStarburstParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersStarburstParametersOutputReference
	_jsii_.Get(
		j,
		"starburstParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) TeradataParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersTeradataParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersTeradataParametersOutputReference
	_jsii_.Get(
		j,
		"teradataParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) TrinoParameters() DataAwsccQuicksightDataSourceAlternateDataSourceParametersTrinoParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceAlternateDataSourceParametersTrinoParametersOutputReference
	_jsii_.Get(
		j,
		"trinoParameters",
		&returns,
	)
	return returns
}


func NewDataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccQuicksightDataSource.DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference_Override(d DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccQuicksightDataSource.DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference)SetInternalValue(val *DataAwsccQuicksightDataSourceAlternateDataSourceParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceAlternateDataSourceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

