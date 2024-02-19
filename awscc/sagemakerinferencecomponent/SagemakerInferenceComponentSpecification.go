package sagemakerinferencecomponent


type SagemakerInferenceComponentSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_component#compute_resource_requirements SagemakerInferenceComponent#compute_resource_requirements}.
	ComputeResourceRequirements *SagemakerInferenceComponentSpecificationComputeResourceRequirements `field:"required" json:"computeResourceRequirements" yaml:"computeResourceRequirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_component#container SagemakerInferenceComponent#container}.
	Container *SagemakerInferenceComponentSpecificationContainer `field:"optional" json:"container" yaml:"container"`
	// The name of the model to use with the inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_component#model_name SagemakerInferenceComponent#model_name}
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_component#startup_parameters SagemakerInferenceComponent#startup_parameters}.
	StartupParameters *SagemakerInferenceComponentSpecificationStartupParameters `field:"optional" json:"startupParameters" yaml:"startupParameters"`
}

