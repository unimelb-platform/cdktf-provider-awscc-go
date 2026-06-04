package sagemakerdomain


type SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCodeRepositories struct {
	// A CodeRepository (valid URL) to be used within Jupyter's Git extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_domain#repository_url SagemakerDomain#repository_url}
	RepositoryUrl *string `field:"optional" json:"repositoryUrl" yaml:"repositoryUrl"`
}

