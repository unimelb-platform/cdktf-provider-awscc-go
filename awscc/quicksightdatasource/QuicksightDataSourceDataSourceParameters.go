package quicksightdatasource


type QuicksightDataSourceDataSourceParameters struct {
	// <p>Amazon Elasticsearch Service parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#amazon_elasticsearch_parameters QuicksightDataSource#amazon_elasticsearch_parameters}
	AmazonElasticsearchParameters *QuicksightDataSourceDataSourceParametersAmazonElasticsearchParameters `field:"optional" json:"amazonElasticsearchParameters" yaml:"amazonElasticsearchParameters"`
	// <p>Amazon OpenSearch Service parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#amazon_open_search_parameters QuicksightDataSource#amazon_open_search_parameters}
	AmazonOpenSearchParameters *QuicksightDataSourceDataSourceParametersAmazonOpenSearchParameters `field:"optional" json:"amazonOpenSearchParameters" yaml:"amazonOpenSearchParameters"`
	// <p>Amazon Athena parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#athena_parameters QuicksightDataSource#athena_parameters}
	AthenaParameters *QuicksightDataSourceDataSourceParametersAthenaParameters `field:"optional" json:"athenaParameters" yaml:"athenaParameters"`
	// <p>Amazon Aurora parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#aurora_parameters QuicksightDataSource#aurora_parameters}
	AuroraParameters *QuicksightDataSourceDataSourceParametersAuroraParameters `field:"optional" json:"auroraParameters" yaml:"auroraParameters"`
	// <p>Amazon Aurora with PostgreSQL compatibility parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#aurora_postgre_sql_parameters QuicksightDataSource#aurora_postgre_sql_parameters}
	AuroraPostgreSqlParameters *QuicksightDataSourceDataSourceParametersAuroraPostgreSqlParameters `field:"optional" json:"auroraPostgreSqlParameters" yaml:"auroraPostgreSqlParameters"`
	// <p>Databricks parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#databricks_parameters QuicksightDataSource#databricks_parameters}
	DatabricksParameters *QuicksightDataSourceDataSourceParametersDatabricksParameters `field:"optional" json:"databricksParameters" yaml:"databricksParameters"`
	// <p>MariaDB parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#maria_db_parameters QuicksightDataSource#maria_db_parameters}
	MariaDbParameters *QuicksightDataSourceDataSourceParametersMariaDbParameters `field:"optional" json:"mariaDbParameters" yaml:"mariaDbParameters"`
	// <p>MySQL parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#my_sql_parameters QuicksightDataSource#my_sql_parameters}
	MySqlParameters *QuicksightDataSourceDataSourceParametersMySqlParameters `field:"optional" json:"mySqlParameters" yaml:"mySqlParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#oracle_parameters QuicksightDataSource#oracle_parameters}.
	OracleParameters *QuicksightDataSourceDataSourceParametersOracleParameters `field:"optional" json:"oracleParameters" yaml:"oracleParameters"`
	// <p>PostgreSQL parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#postgre_sql_parameters QuicksightDataSource#postgre_sql_parameters}
	PostgreSqlParameters *QuicksightDataSourceDataSourceParametersPostgreSqlParameters `field:"optional" json:"postgreSqlParameters" yaml:"postgreSqlParameters"`
	// <p>Presto parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#presto_parameters QuicksightDataSource#presto_parameters}
	PrestoParameters *QuicksightDataSourceDataSourceParametersPrestoParameters `field:"optional" json:"prestoParameters" yaml:"prestoParameters"`
	// <p>Amazon RDS parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#rds_parameters QuicksightDataSource#rds_parameters}
	RdsParameters *QuicksightDataSourceDataSourceParametersRdsParameters `field:"optional" json:"rdsParameters" yaml:"rdsParameters"`
	// <p>Amazon Redshift parameters.
	//
	// The <code>ClusterId</code> field can be blank if
	//             <code>Host</code> and <code>Port</code> are both set. The <code>Host</code> and
	//             <code>Port</code> fields can be blank if the <code>ClusterId</code> field is set.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#redshift_parameters QuicksightDataSource#redshift_parameters}
	RedshiftParameters *QuicksightDataSourceDataSourceParametersRedshiftParameters `field:"optional" json:"redshiftParameters" yaml:"redshiftParameters"`
	// <p>S3 parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#s3_parameters QuicksightDataSource#s3_parameters}
	S3Parameters *QuicksightDataSourceDataSourceParametersS3Parameters `field:"optional" json:"s3Parameters" yaml:"s3Parameters"`
	// <p>Snowflake parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#snowflake_parameters QuicksightDataSource#snowflake_parameters}
	SnowflakeParameters *QuicksightDataSourceDataSourceParametersSnowflakeParameters `field:"optional" json:"snowflakeParameters" yaml:"snowflakeParameters"`
	// <p>Spark parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#spark_parameters QuicksightDataSource#spark_parameters}
	SparkParameters *QuicksightDataSourceDataSourceParametersSparkParameters `field:"optional" json:"sparkParameters" yaml:"sparkParameters"`
	// <p>SQL Server parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#sql_server_parameters QuicksightDataSource#sql_server_parameters}
	SqlServerParameters *QuicksightDataSourceDataSourceParametersSqlServerParameters `field:"optional" json:"sqlServerParameters" yaml:"sqlServerParameters"`
	// <p>Starburst parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#starburst_parameters QuicksightDataSource#starburst_parameters}
	StarburstParameters *QuicksightDataSourceDataSourceParametersStarburstParameters `field:"optional" json:"starburstParameters" yaml:"starburstParameters"`
	// <p>Teradata parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#teradata_parameters QuicksightDataSource#teradata_parameters}
	TeradataParameters *QuicksightDataSourceDataSourceParametersTeradataParameters `field:"optional" json:"teradataParameters" yaml:"teradataParameters"`
	// <p>Trino parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#trino_parameters QuicksightDataSource#trino_parameters}
	TrinoParameters *QuicksightDataSourceDataSourceParametersTrinoParameters `field:"optional" json:"trinoParameters" yaml:"trinoParameters"`
}

