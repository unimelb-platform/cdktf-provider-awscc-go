package sagemakermodelpackage


type SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersModelInput struct {
	// The input configuration object for the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#data_input_config SagemakerModelPackage#data_input_config}
	DataInputConfig *string `field:"required" json:"dataInputConfig" yaml:"dataInputConfig"`
}

