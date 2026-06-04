package bedrockdatasource


type BedrockDataSourceDataSourceConfiguration struct {
	// The type of the data source location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_source#type BedrockDataSource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The configuration information to connect to Confluence as your data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_source#confluence_configuration BedrockDataSource#confluence_configuration}
	ConfluenceConfiguration *BedrockDataSourceDataSourceConfigurationConfluenceConfiguration `field:"optional" json:"confluenceConfiguration" yaml:"confluenceConfiguration"`
	// The configuration information to connect to Amazon S3 as your data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_source#s3_configuration BedrockDataSource#s3_configuration}
	S3Configuration *BedrockDataSourceDataSourceConfigurationS3Configuration `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
	// The configuration information to connect to Salesforce as your data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_source#salesforce_configuration BedrockDataSource#salesforce_configuration}
	SalesforceConfiguration *BedrockDataSourceDataSourceConfigurationSalesforceConfiguration `field:"optional" json:"salesforceConfiguration" yaml:"salesforceConfiguration"`
	// The configuration information to connect to SharePoint as your data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_source#share_point_configuration BedrockDataSource#share_point_configuration}
	SharePointConfiguration *BedrockDataSourceDataSourceConfigurationSharePointConfiguration `field:"optional" json:"sharePointConfiguration" yaml:"sharePointConfiguration"`
	// Configures a web data source location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_source#web_configuration BedrockDataSource#web_configuration}
	WebConfiguration *BedrockDataSourceDataSourceConfigurationWebConfiguration `field:"optional" json:"webConfiguration" yaml:"webConfiguration"`
}

