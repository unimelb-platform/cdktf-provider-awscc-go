package opensearchserverlesscollection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpensearchserverlessCollectionConfig struct {
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
	// The name of the collection.
	//
	// The name must meet the following criteria:
	// Unique to your account and AWS Region
	// Starts with a lowercase letter
	// Contains only lowercase letters a-z, the numbers 0-9 and the hyphen (-)
	// Contains between 3 and 32 characters
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/opensearchserverless_collection#name OpensearchserverlessCollection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/opensearchserverless_collection#description OpensearchserverlessCollection#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The possible standby replicas for the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/opensearchserverless_collection#standby_replicas OpensearchserverlessCollection#standby_replicas}
	StandbyReplicas *string `field:"optional" json:"standbyReplicas" yaml:"standbyReplicas"`
	// List of tags to be added to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/opensearchserverless_collection#tags OpensearchserverlessCollection#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The possible types for the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/opensearchserverless_collection#type OpensearchserverlessCollection#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

