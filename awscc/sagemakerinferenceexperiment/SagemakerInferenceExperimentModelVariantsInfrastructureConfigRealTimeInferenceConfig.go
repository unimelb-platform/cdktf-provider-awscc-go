package sagemakerinferenceexperiment


type SagemakerInferenceExperimentModelVariantsInfrastructureConfigRealTimeInferenceConfig struct {
	// The number of instances of the type specified by InstanceType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_experiment#instance_count SagemakerInferenceExperiment#instance_count}
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// The instance type the model is deployed to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_experiment#instance_type SagemakerInferenceExperiment#instance_type}
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
}

