package sagemakerinferencecomponent


type SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyMaximumBatchSize struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_inference_component#type SagemakerInferenceComponent#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The number of copies for the inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_inference_component#value SagemakerInferenceComponent#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

