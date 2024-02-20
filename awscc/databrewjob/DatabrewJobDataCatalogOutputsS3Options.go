package databrewjob


type DatabrewJobDataCatalogOutputsS3Options struct {
	// S3 Output location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job#location DatabrewJob#location}
	Location *DatabrewJobDataCatalogOutputsS3OptionsLocation `field:"required" json:"location" yaml:"location"`
}

