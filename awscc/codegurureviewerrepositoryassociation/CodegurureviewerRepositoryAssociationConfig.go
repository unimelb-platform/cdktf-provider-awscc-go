package codegurureviewerrepositoryassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodegurureviewerRepositoryAssociationConfig struct {
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
	// Name of the repository to be associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codegurureviewer_repository_association#name CodegurureviewerRepositoryAssociation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of repository to be associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codegurureviewer_repository_association#type CodegurureviewerRepositoryAssociation#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The name of the S3 bucket associated with an associated S3 repository. It must start with `codeguru-reviewer-`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codegurureviewer_repository_association#bucket_name CodegurureviewerRepositoryAssociation#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The Amazon Resource Name (ARN) of an AWS CodeStar Connections connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codegurureviewer_repository_association#connection_arn CodegurureviewerRepositoryAssociation#connection_arn}
	ConnectionArn *string `field:"optional" json:"connectionArn" yaml:"connectionArn"`
	// The owner of the repository.
	//
	// For a Bitbucket repository, this is the username for the account that owns the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codegurureviewer_repository_association#owner CodegurureviewerRepositoryAssociation#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// The tags associated with a repository association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codegurureviewer_repository_association#tags CodegurureviewerRepositoryAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

