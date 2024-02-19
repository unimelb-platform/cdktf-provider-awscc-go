package datazonedatasource


type DatazoneDataSourceRecommendation struct {
	// Specifies whether automatic business name generation is to be enabled or not as part of the recommendation configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_data_source#enable_business_name_generation DatazoneDataSource#enable_business_name_generation}
	EnableBusinessNameGeneration interface{} `field:"optional" json:"enableBusinessNameGeneration" yaml:"enableBusinessNameGeneration"`
}

