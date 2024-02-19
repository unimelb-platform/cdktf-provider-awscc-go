package sagemakerinferenceexperiment


type SagemakerInferenceExperimentShadowModeConfig struct {
	// List of shadow variant configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_experiment#shadow_model_variants SagemakerInferenceExperiment#shadow_model_variants}
	ShadowModelVariants interface{} `field:"required" json:"shadowModelVariants" yaml:"shadowModelVariants"`
	// The name of the production variant, which takes all the inference requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_experiment#source_model_variant_name SagemakerInferenceExperiment#source_model_variant_name}
	SourceModelVariantName *string `field:"required" json:"sourceModelVariantName" yaml:"sourceModelVariantName"`
}

