package dmsdataprovider


type DmsDataProviderSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_data_provider#microsoft_sql_server_settings DmsDataProvider#microsoft_sql_server_settings}.
	MicrosoftSqlServerSettings *DmsDataProviderSettingsMicrosoftSqlServerSettings `field:"optional" json:"microsoftSqlServerSettings" yaml:"microsoftSqlServerSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_data_provider#my_sql_settings DmsDataProvider#my_sql_settings}.
	MySqlSettings *DmsDataProviderSettingsMySqlSettings `field:"optional" json:"mySqlSettings" yaml:"mySqlSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_data_provider#oracle_settings DmsDataProvider#oracle_settings}.
	OracleSettings *DmsDataProviderSettingsOracleSettings `field:"optional" json:"oracleSettings" yaml:"oracleSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_data_provider#postgre_sql_settings DmsDataProvider#postgre_sql_settings}.
	PostgreSqlSettings *DmsDataProviderSettingsPostgreSqlSettings `field:"optional" json:"postgreSqlSettings" yaml:"postgreSqlSettings"`
}

