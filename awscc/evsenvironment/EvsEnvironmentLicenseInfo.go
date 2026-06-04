package evsenvironment


type EvsEnvironmentLicenseInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#solution_key EvsEnvironment#solution_key}.
	SolutionKey *string `field:"required" json:"solutionKey" yaml:"solutionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#vsan_key EvsEnvironment#vsan_key}.
	VsanKey *string `field:"required" json:"vsanKey" yaml:"vsanKey"`
}

