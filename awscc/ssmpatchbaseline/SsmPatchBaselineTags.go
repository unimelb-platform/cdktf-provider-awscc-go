package ssmpatchbaseline


type SsmPatchBaselineTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_patch_baseline#key SsmPatchBaseline#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_patch_baseline#value SsmPatchBaseline#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

