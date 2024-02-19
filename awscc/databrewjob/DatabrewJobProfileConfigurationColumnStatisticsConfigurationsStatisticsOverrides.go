package databrewjob


type DatabrewJobProfileConfigurationColumnStatisticsConfigurationsStatisticsOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job#parameters DatabrewJob#parameters}.
	Parameters *map[string]*string `field:"required" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job#statistic DatabrewJob#statistic}.
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
}

