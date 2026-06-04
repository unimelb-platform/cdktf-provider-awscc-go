package databrewjob


type DatabrewJobOutputsFormatOptions struct {
	// Output Csv options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/databrew_job#csv DatabrewJob#csv}
	Csv *DatabrewJobOutputsFormatOptionsCsv `field:"optional" json:"csv" yaml:"csv"`
}

