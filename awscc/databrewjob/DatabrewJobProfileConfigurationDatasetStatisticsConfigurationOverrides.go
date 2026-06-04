package databrewjob


type DatabrewJobProfileConfigurationDatasetStatisticsConfigurationOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/databrew_job#parameters DatabrewJob#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/databrew_job#statistic DatabrewJob#statistic}.
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
}

