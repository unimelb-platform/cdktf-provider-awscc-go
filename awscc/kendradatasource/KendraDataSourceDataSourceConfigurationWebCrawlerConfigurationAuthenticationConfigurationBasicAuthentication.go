package kendradatasource


type KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationBasicAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kendra_data_source#credentials KendraDataSource#credentials}.
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kendra_data_source#host KendraDataSource#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kendra_data_source#port KendraDataSource#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

