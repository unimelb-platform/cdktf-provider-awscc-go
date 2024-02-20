package databrewjob


type DatabrewJobProfileConfigurationColumnStatisticsConfigurationsStatistics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job#included_statistics DatabrewJob#included_statistics}.
	IncludedStatistics *[]*string `field:"optional" json:"includedStatistics" yaml:"includedStatistics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job#overrides DatabrewJob#overrides}.
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}

