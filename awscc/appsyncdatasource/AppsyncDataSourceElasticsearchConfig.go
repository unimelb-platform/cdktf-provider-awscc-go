package appsyncdatasource


type AppsyncDataSourceElasticsearchConfig struct {
	// The AWS Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_data_source#aws_region AppsyncDataSource#aws_region}
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_data_source#endpoint AppsyncDataSource#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

