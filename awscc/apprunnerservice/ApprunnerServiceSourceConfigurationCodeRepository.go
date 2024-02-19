package apprunnerservice


type ApprunnerServiceSourceConfigurationCodeRepository struct {
	// Repository Url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#repository_url ApprunnerService#repository_url}
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
	// Source Code Version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#source_code_version ApprunnerService#source_code_version}
	SourceCodeVersion *ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersion `field:"required" json:"sourceCodeVersion" yaml:"sourceCodeVersion"`
	// Code Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#code_configuration ApprunnerService#code_configuration}
	CodeConfiguration *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfiguration `field:"optional" json:"codeConfiguration" yaml:"codeConfiguration"`
	// Source Directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#source_directory ApprunnerService#source_directory}
	SourceDirectory *string `field:"optional" json:"sourceDirectory" yaml:"sourceDirectory"`
}

