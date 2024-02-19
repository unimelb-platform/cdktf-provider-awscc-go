package databrewproject


type DatabrewProjectSample struct {
	// Sample type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_project#type DatabrewProject#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Sample size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_project#size DatabrewProject#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

