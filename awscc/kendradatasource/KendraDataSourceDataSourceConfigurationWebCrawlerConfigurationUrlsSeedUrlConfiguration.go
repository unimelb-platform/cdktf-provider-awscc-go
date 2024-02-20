package kendradatasource


type KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#seed_urls KendraDataSource#seed_urls}.
	SeedUrls *[]*string `field:"required" json:"seedUrls" yaml:"seedUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#web_crawler_mode KendraDataSource#web_crawler_mode}.
	WebCrawlerMode *string `field:"optional" json:"webCrawlerMode" yaml:"webCrawlerMode"`
}

