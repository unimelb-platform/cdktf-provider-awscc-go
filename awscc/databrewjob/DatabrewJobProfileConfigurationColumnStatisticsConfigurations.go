package databrewjob


type DatabrewJobProfileConfigurationColumnStatisticsConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/databrew_job#selectors DatabrewJob#selectors}.
	Selectors interface{} `field:"optional" json:"selectors" yaml:"selectors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/databrew_job#statistics DatabrewJob#statistics}.
	Statistics *DatabrewJobProfileConfigurationColumnStatisticsConfigurationsStatistics `field:"optional" json:"statistics" yaml:"statistics"`
}

