package sagemakermodelpackage


type SagemakerModelPackageMetadataProperties struct {
	// The commit ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#commit_id SagemakerModelPackage#commit_id}
	CommitId *string `field:"optional" json:"commitId" yaml:"commitId"`
	// The entity this entity was generated by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#generated_by SagemakerModelPackage#generated_by}
	GeneratedBy *string `field:"optional" json:"generatedBy" yaml:"generatedBy"`
	// The project ID metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#project_id SagemakerModelPackage#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// The repository metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#repository SagemakerModelPackage#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
}
