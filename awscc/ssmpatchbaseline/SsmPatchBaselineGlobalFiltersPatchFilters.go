package ssmpatchbaseline


type SsmPatchBaselineGlobalFiltersPatchFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_patch_baseline#key SsmPatchBaseline#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_patch_baseline#values SsmPatchBaseline#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

