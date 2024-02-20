package apprunnerservice


type ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues struct {
	// Runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#runtime ApprunnerService#runtime}
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// Build Command.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#build_command ApprunnerService#build_command}
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// Port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#port ApprunnerService#port}
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The secrets and parameters that get referenced by your service as environment variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#runtime_environment_secrets ApprunnerService#runtime_environment_secrets}
	RuntimeEnvironmentSecrets interface{} `field:"optional" json:"runtimeEnvironmentSecrets" yaml:"runtimeEnvironmentSecrets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#runtime_environment_variables ApprunnerService#runtime_environment_variables}.
	RuntimeEnvironmentVariables interface{} `field:"optional" json:"runtimeEnvironmentVariables" yaml:"runtimeEnvironmentVariables"`
	// Start Command.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#start_command ApprunnerService#start_command}
	StartCommand *string `field:"optional" json:"startCommand" yaml:"startCommand"`
}

