package databrewproject


type DatabrewProjectSample struct {
	// Sample size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/databrew_project#size DatabrewProject#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// Sample type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/databrew_project#type DatabrewProject#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

