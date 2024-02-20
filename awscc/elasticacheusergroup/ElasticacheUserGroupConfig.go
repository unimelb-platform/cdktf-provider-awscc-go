package elasticacheusergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElasticacheUserGroupConfig struct {
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
	// Must be redis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_user_group#engine ElasticacheUserGroup#engine}
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// The ID of the user group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_user_group#user_group_id ElasticacheUserGroup#user_group_id}
	UserGroupId *string `field:"required" json:"userGroupId" yaml:"userGroupId"`
	// List of users associated to this user group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_user_group#user_ids ElasticacheUserGroup#user_ids}
	UserIds *[]*string `field:"required" json:"userIds" yaml:"userIds"`
	// An array of key-value pairs to apply to this user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_user_group#tags ElasticacheUserGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

