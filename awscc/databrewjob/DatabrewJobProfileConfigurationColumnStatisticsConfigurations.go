package databrewjob


type DatabrewJobProfileConfigurationColumnStatisticsConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job#statistics DatabrewJob#statistics}.
	Statistics *DatabrewJobProfileConfigurationColumnStatisticsConfigurationsStatistics `field:"required" json:"statistics" yaml:"statistics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job#selectors DatabrewJob#selectors}.
	Selectors interface{} `field:"optional" json:"selectors" yaml:"selectors"`
}

