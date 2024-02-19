package quicksightdatasource


type QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParameters struct {
	// <p>Amazon Elasticsearch Service parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#amazon_elasticsearch_parameters QuicksightDataSource#amazon_elasticsearch_parameters}
	AmazonElasticsearchParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonElasticsearchParameters `field:"optional" json:"amazonElasticsearchParameters" yaml:"amazonElasticsearchParameters"`
	// <p>Amazon OpenSearch Service parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#amazon_open_search_parameters QuicksightDataSource#amazon_open_search_parameters}
	AmazonOpenSearchParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAmazonOpenSearchParameters `field:"optional" json:"amazonOpenSearchParameters" yaml:"amazonOpenSearchParameters"`
	// <p>Amazon Athena parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#athena_parameters QuicksightDataSource#athena_parameters}
	AthenaParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAthenaParameters `field:"optional" json:"athenaParameters" yaml:"athenaParameters"`
	// <p>Amazon Aurora parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#aurora_parameters QuicksightDataSource#aurora_parameters}
	AuroraParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraParameters `field:"optional" json:"auroraParameters" yaml:"auroraParameters"`
	// <p>Amazon Aurora with PostgreSQL compatibility parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#aurora_postgre_sql_parameters QuicksightDataSource#aurora_postgre_sql_parameters}
	AuroraPostgreSqlParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersAuroraPostgreSqlParameters `field:"optional" json:"auroraPostgreSqlParameters" yaml:"auroraPostgreSqlParameters"`
	// <p>Databricks parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#databricks_parameters QuicksightDataSource#databricks_parameters}
	DatabricksParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersDatabricksParameters `field:"optional" json:"databricksParameters" yaml:"databricksParameters"`
	// <p>MariaDB parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#maria_db_parameters QuicksightDataSource#maria_db_parameters}
	MariaDbParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMariaDbParameters `field:"optional" json:"mariaDbParameters" yaml:"mariaDbParameters"`
	// <p>MySQL parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#my_sql_parameters QuicksightDataSource#my_sql_parameters}
	MySqlParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersMySqlParameters `field:"optional" json:"mySqlParameters" yaml:"mySqlParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#oracle_parameters QuicksightDataSource#oracle_parameters}.
	OracleParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersOracleParameters `field:"optional" json:"oracleParameters" yaml:"oracleParameters"`
	// <p>PostgreSQL parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#postgre_sql_parameters QuicksightDataSource#postgre_sql_parameters}
	PostgreSqlParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPostgreSqlParameters `field:"optional" json:"postgreSqlParameters" yaml:"postgreSqlParameters"`
	// <p>Presto parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#presto_parameters QuicksightDataSource#presto_parameters}
	PrestoParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersPrestoParameters `field:"optional" json:"prestoParameters" yaml:"prestoParameters"`
	// <p>Amazon RDS parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#rds_parameters QuicksightDataSource#rds_parameters}
	RdsParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRdsParameters `field:"optional" json:"rdsParameters" yaml:"rdsParameters"`
	// <p>Amazon Redshift parameters.
	//
	// The <code>ClusterId</code> field can be blank if
	//             <code>Host</code> and <code>Port</code> are both set. The <code>Host</code> and
	//             <code>Port</code> fields can be blank if the <code>ClusterId</code> field is set.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#redshift_parameters QuicksightDataSource#redshift_parameters}
	RedshiftParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParameters `field:"optional" json:"redshiftParameters" yaml:"redshiftParameters"`
	// <p>S3 parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#s3_parameters QuicksightDataSource#s3_parameters}
	S3Parameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersS3Parameters `field:"optional" json:"s3Parameters" yaml:"s3Parameters"`
	// <p>Snowflake parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#snowflake_parameters QuicksightDataSource#snowflake_parameters}
	SnowflakeParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSnowflakeParameters `field:"optional" json:"snowflakeParameters" yaml:"snowflakeParameters"`
	// <p>Spark parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#spark_parameters QuicksightDataSource#spark_parameters}
	SparkParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSparkParameters `field:"optional" json:"sparkParameters" yaml:"sparkParameters"`
	// <p>SQL Server parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#sql_server_parameters QuicksightDataSource#sql_server_parameters}
	SqlServerParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSqlServerParameters `field:"optional" json:"sqlServerParameters" yaml:"sqlServerParameters"`
	// <p>Starburst parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#starburst_parameters QuicksightDataSource#starburst_parameters}
	StarburstParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParameters `field:"optional" json:"starburstParameters" yaml:"starburstParameters"`
	// <p>Teradata parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#teradata_parameters QuicksightDataSource#teradata_parameters}
	TeradataParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTeradataParameters `field:"optional" json:"teradataParameters" yaml:"teradataParameters"`
	// <p>Trino parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#trino_parameters QuicksightDataSource#trino_parameters}
	TrinoParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParameters `field:"optional" json:"trinoParameters" yaml:"trinoParameters"`
}

