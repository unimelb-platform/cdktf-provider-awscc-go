package protonenvironmenttemplate


type ProtonEnvironmentTemplateTags struct {
	// <p>The key of the resource tag.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/proton_environment_template#key ProtonEnvironmentTemplate#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// <p>The value of the resource tag.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/proton_environment_template#value ProtonEnvironmentTemplate#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

