package bedrockapplicationinferenceprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockApplicationInferenceProfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_application_inference_profile#inference_profile_name BedrockApplicationInferenceProfile#inference_profile_name}.
	InferenceProfileName *string `field:"required" json:"inferenceProfileName" yaml:"inferenceProfileName"`
	// Description of the inference profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_application_inference_profile#description BedrockApplicationInferenceProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Various ways to encode a list of models in a CreateInferenceProfile request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_application_inference_profile#model_source BedrockApplicationInferenceProfile#model_source}
	ModelSource *BedrockApplicationInferenceProfileModelSource `field:"optional" json:"modelSource" yaml:"modelSource"`
	// List of Tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_application_inference_profile#tags BedrockApplicationInferenceProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

