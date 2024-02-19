package sagemakeruserprofile


type SagemakerUserProfileUserSettingsJupyterLabAppSettingsCodeRepositories struct {
	// A CodeRepository (valid URL) to be used within Jupyter's Git extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#repository_url SagemakerUserProfile#repository_url}
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
}

