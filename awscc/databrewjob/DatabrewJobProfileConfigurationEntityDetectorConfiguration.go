package databrewjob


type DatabrewJobProfileConfigurationEntityDetectorConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job#entity_types DatabrewJob#entity_types}.
	EntityTypes *[]*string `field:"required" json:"entityTypes" yaml:"entityTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job#allowed_statistics DatabrewJob#allowed_statistics}.
	AllowedStatistics *DatabrewJobProfileConfigurationEntityDetectorConfigurationAllowedStatistics `field:"optional" json:"allowedStatistics" yaml:"allowedStatistics"`
}

