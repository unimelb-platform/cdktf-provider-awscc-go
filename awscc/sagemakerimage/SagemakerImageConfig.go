package sagemakerimage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerImageConfig struct {
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
	// The name of the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_image#image_name SagemakerImage#image_name}
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on behalf of the customer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_image#image_role_arn SagemakerImage#image_role_arn}
	ImageRoleArn *string `field:"required" json:"imageRoleArn" yaml:"imageRoleArn"`
	// A description of the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_image#image_description SagemakerImage#image_description}
	ImageDescription *string `field:"optional" json:"imageDescription" yaml:"imageDescription"`
	// The display name of the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_image#image_display_name SagemakerImage#image_display_name}
	ImageDisplayName *string `field:"optional" json:"imageDisplayName" yaml:"imageDisplayName"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_image#tags SagemakerImage#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

