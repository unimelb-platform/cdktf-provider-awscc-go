package sagemakermodelpackage


type SagemakerModelPackageModelCard struct {
	// The content of the model card.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_model_package#model_card_content SagemakerModelPackage#model_card_content}
	ModelCardContent *string `field:"optional" json:"modelCardContent" yaml:"modelCardContent"`
	// The approval status of the model card within your organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_model_package#model_card_status SagemakerModelPackage#model_card_status}
	ModelCardStatus *string `field:"optional" json:"modelCardStatus" yaml:"modelCardStatus"`
}

