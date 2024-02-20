package sagemakerspace


type SagemakerSpaceSpaceSettingsJupyterLabAppSettingsCodeRepositories struct {
	// A CodeRepository (valid URL) to be used within Jupyter's Git extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#repository_url SagemakerSpace#repository_url}
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
}

