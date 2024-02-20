package ecrrepository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcrRepositoryConfig struct {
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
	// If true, deleting the repository force deletes the contents of the repository.
	//
	// Without a force delete, you can only delete empty repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_repository#empty_on_delete EcrRepository#empty_on_delete}
	EmptyOnDelete interface{} `field:"optional" json:"emptyOnDelete" yaml:"emptyOnDelete"`
	// The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.
	//
	// By default, when no encryption configuration is set or the AES256 encryption type is used, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts your data at rest using an AES-256 encryption algorithm. This does not require any action on your part.
	//
	// For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_repository#encryption_configuration EcrRepository#encryption_configuration}
	EncryptionConfiguration *EcrRepositoryEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The image scanning configuration for the repository.
	//
	// This setting determines whether images are scanned for known vulnerabilities after being pushed to the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_repository#image_scanning_configuration EcrRepository#image_scanning_configuration}
	ImageScanningConfiguration *EcrRepositoryImageScanningConfiguration `field:"optional" json:"imageScanningConfiguration" yaml:"imageScanningConfiguration"`
	// The image tag mutability setting for the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_repository#image_tag_mutability EcrRepository#image_tag_mutability}
	ImageTagMutability *string `field:"optional" json:"imageTagMutability" yaml:"imageTagMutability"`
	// The LifecyclePolicy property type specifies a lifecycle policy. For information about lifecycle policy syntax, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_repository#lifecycle_policy EcrRepository#lifecycle_policy}
	LifecyclePolicy *EcrRepositoryLifecyclePolicy `field:"optional" json:"lifecyclePolicy" yaml:"lifecyclePolicy"`
	// The name to use for the repository.
	//
	// The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_repository#repository_name EcrRepository#repository_name}
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// The JSON repository policy text to apply to the repository.
	//
	// For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_repository#repository_policy_text EcrRepository#repository_policy_text}
	RepositoryPolicyText *string `field:"optional" json:"repositoryPolicyText" yaml:"repositoryPolicyText"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_repository#tags EcrRepository#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

