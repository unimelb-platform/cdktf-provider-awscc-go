package evsenvironment


type EvsEnvironmentServiceAccessSecurityGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#security_groups EvsEnvironment#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

