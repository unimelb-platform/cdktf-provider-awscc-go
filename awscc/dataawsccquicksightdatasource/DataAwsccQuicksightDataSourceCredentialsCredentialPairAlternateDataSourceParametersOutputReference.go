package dataawsccquicksightdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccquicksightdatasource/internal"
)

type DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference interface {
	cdktf.ComplexObject
	AmazonElasticsearchParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonElasticsearchParametersOutputReference
	AmazonOpenSearchParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonOpenSearchParametersOutputReference
	AthenaParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAthenaParametersOutputReference
	AuroraParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraParametersOutputReference
	AuroraPostgreSqlParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraPostgreSqlParametersOutputReference
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
	DatabricksParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersDatabricksParametersOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParameters
	SetInternalValue(val *DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParameters)
	MariaDbParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMariaDbParametersOutputReference
	MySqlParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMySqlParametersOutputReference
	OracleParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOracleParametersOutputReference
	PostgreSqlParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPostgreSqlParametersOutputReference
	PrestoParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPrestoParametersOutputReference
	RdsParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRdsParametersOutputReference
	RedshiftParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference
	S3Parameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersS3ParametersOutputReference
	SnowflakeParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSnowflakeParametersOutputReference
	SparkParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSparkParametersOutputReference
	SqlServerParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSqlServerParametersOutputReference
	StarburstParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference
	TeradataParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTeradataParametersOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrinoParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference
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

// The jsii proxy struct for DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference
type jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) AmazonElasticsearchParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonElasticsearchParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonElasticsearchParametersOutputReference
	_jsii_.Get(
		j,
		"amazonElasticsearchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) AmazonOpenSearchParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonOpenSearchParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonOpenSearchParametersOutputReference
	_jsii_.Get(
		j,
		"amazonOpenSearchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) AthenaParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAthenaParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAthenaParametersOutputReference
	_jsii_.Get(
		j,
		"athenaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) AuroraParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraParametersOutputReference
	_jsii_.Get(
		j,
		"auroraParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) AuroraPostgreSqlParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraPostgreSqlParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraPostgreSqlParametersOutputReference
	_jsii_.Get(
		j,
		"auroraPostgreSqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) DatabricksParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersDatabricksParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersDatabricksParametersOutputReference
	_jsii_.Get(
		j,
		"databricksParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) InternalValue() *DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParameters {
	var returns *DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) MariaDbParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMariaDbParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMariaDbParametersOutputReference
	_jsii_.Get(
		j,
		"mariaDbParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) MySqlParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMySqlParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMySqlParametersOutputReference
	_jsii_.Get(
		j,
		"mySqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) OracleParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOracleParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOracleParametersOutputReference
	_jsii_.Get(
		j,
		"oracleParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PostgreSqlParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPostgreSqlParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPostgreSqlParametersOutputReference
	_jsii_.Get(
		j,
		"postgreSqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) PrestoParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPrestoParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPrestoParametersOutputReference
	_jsii_.Get(
		j,
		"prestoParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) RdsParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRdsParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRdsParametersOutputReference
	_jsii_.Get(
		j,
		"rdsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) RedshiftParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference
	_jsii_.Get(
		j,
		"redshiftParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) S3Parameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersS3ParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersS3ParametersOutputReference
	_jsii_.Get(
		j,
		"s3Parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) SnowflakeParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSnowflakeParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSnowflakeParametersOutputReference
	_jsii_.Get(
		j,
		"snowflakeParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) SparkParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSparkParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSparkParametersOutputReference
	_jsii_.Get(
		j,
		"sparkParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) SqlServerParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSqlServerParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSqlServerParametersOutputReference
	_jsii_.Get(
		j,
		"sqlServerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) StarburstParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference
	_jsii_.Get(
		j,
		"starburstParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) TeradataParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTeradataParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTeradataParametersOutputReference
	_jsii_.Get(
		j,
		"teradataParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) TrinoParameters() DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference {
	var returns DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference
	_jsii_.Get(
		j,
		"trinoParameters",
		&returns,
	)
	return returns
}


func NewDataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccQuicksightDataSource.DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference_Override(d DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccQuicksightDataSource.DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference)SetInternalValue(val *DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

