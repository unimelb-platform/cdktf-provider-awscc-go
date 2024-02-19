package ssmpatchbaseline


type SsmPatchBaselineSources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_patch_baseline#configuration SsmPatchBaseline#configuration}.
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_patch_baseline#name SsmPatchBaseline#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_patch_baseline#products SsmPatchBaseline#products}.
	Products *[]*string `field:"optional" json:"products" yaml:"products"`
}

