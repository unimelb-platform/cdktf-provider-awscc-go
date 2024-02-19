package forecastdataset


type ForecastDatasetSchemaAttributes struct {
	// Name of the dataset field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/forecast_dataset#attribute_name ForecastDataset#attribute_name}
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
	// Data type of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/forecast_dataset#attribute_type ForecastDataset#attribute_type}
	AttributeType *string `field:"optional" json:"attributeType" yaml:"attributeType"`
}

