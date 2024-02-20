package sagemakerinferencecomponent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerInferenceComponentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the endpoint the inference component is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_component#endpoint_name SagemakerInferenceComponent#endpoint_name}
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// The runtime config for the inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_component#runtime_config SagemakerInferenceComponent#runtime_config}
	RuntimeConfig *SagemakerInferenceComponentRuntimeConfig `field:"required" json:"runtimeConfig" yaml:"runtimeConfig"`
	// The specification for the inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_component#specification SagemakerInferenceComponent#specification}
	Specification *SagemakerInferenceComponentSpecification `field:"required" json:"specification" yaml:"specification"`
	// The name of the endpoint variant the inference component is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_component#variant_name SagemakerInferenceComponent#variant_name}
	VariantName *string `field:"required" json:"variantName" yaml:"variantName"`
	// The Amazon Resource Name (ARN) of the endpoint the inference component is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_component#endpoint_arn SagemakerInferenceComponent#endpoint_arn}
	EndpointArn *string `field:"optional" json:"endpointArn" yaml:"endpointArn"`
	// The name of the inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_component#inference_component_name SagemakerInferenceComponent#inference_component_name}
	InferenceComponentName *string `field:"optional" json:"inferenceComponentName" yaml:"inferenceComponentName"`
	// An array of tags to apply to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_component#tags SagemakerInferenceComponent#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

