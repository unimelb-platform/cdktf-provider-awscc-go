package sagemakerdomain


type SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCodeRepositories struct {
	// A CodeRepository (valid URL) to be used within Jupyter's Git extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#repository_url SagemakerDomain#repository_url}
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
}

