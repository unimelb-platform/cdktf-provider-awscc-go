package m2environment


type M2EnvironmentHighAvailabilityConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/m2_environment#desired_capacity M2Environment#desired_capacity}.
	DesiredCapacity *float64 `field:"required" json:"desiredCapacity" yaml:"desiredCapacity"`
}

