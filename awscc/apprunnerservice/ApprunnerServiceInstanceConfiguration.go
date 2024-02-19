package apprunnerservice


type ApprunnerServiceInstanceConfiguration struct {
	// CPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#cpu ApprunnerService#cpu}
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Instance Role Arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#instance_role_arn ApprunnerService#instance_role_arn}
	InstanceRoleArn *string `field:"optional" json:"instanceRoleArn" yaml:"instanceRoleArn"`
	// Memory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#memory ApprunnerService#memory}
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

