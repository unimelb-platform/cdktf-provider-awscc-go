package quicksightdataset


type QuicksightDataSetColumnGroupsGeoSpatialColumnGroup struct {
	// <p>Columns in this hierarchy.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#columns QuicksightDataSet#columns}
	Columns *[]*string `field:"required" json:"columns" yaml:"columns"`
	// <p>A display name for the hierarchy.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#name QuicksightDataSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#country_code QuicksightDataSet#country_code}.
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
}

