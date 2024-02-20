package ecsservice


type EcsServiceServiceConnectConfigurationLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#log_driver EcsService#log_driver}.
	LogDriver *string `field:"optional" json:"logDriver" yaml:"logDriver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#options EcsService#options}.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#secret_options EcsService#secret_options}.
	SecretOptions interface{} `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}

