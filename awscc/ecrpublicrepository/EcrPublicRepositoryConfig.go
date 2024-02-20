package ecrpublicrepository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcrPublicRepositoryConfig struct {
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
	// The CatalogData property type specifies Catalog data for ECR Public Repository. For information about Catalog Data, see <link>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_public_repository#repository_catalog_data EcrPublicRepository#repository_catalog_data}
	RepositoryCatalogData *EcrPublicRepositoryRepositoryCatalogData `field:"optional" json:"repositoryCatalogData" yaml:"repositoryCatalogData"`
	// The name to use for the repository.
	//
	// The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_public_repository#repository_name EcrPublicRepository#repository_name}
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// The JSON repository policy text to apply to the repository.
	//
	// For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_public_repository#repository_policy_text EcrPublicRepository#repository_policy_text}
	RepositoryPolicyText *string `field:"optional" json:"repositoryPolicyText" yaml:"repositoryPolicyText"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_public_repository#tags EcrPublicRepository#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

