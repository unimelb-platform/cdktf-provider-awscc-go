package appintegrationsapplication


type AppintegrationsApplicationTags struct {
	// A key to identify the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appintegrations_application#key AppintegrationsApplication#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Corresponding tag value for the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appintegrations_application#value AppintegrationsApplication#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

