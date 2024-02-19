package connectuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConnectUserConfig struct {
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
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#instance_arn ConnectUser#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The phone settings for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#phone_config ConnectUser#phone_config}
	PhoneConfig *ConnectUserPhoneConfig `field:"required" json:"phoneConfig" yaml:"phoneConfig"`
	// The identifier of the routing profile for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#routing_profile_arn ConnectUser#routing_profile_arn}
	RoutingProfileArn *string `field:"required" json:"routingProfileArn" yaml:"routingProfileArn"`
	// One or more security profile arns for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#security_profile_arns ConnectUser#security_profile_arns}
	SecurityProfileArns *[]*string `field:"required" json:"securityProfileArns" yaml:"securityProfileArns"`
	// The user name for the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#username ConnectUser#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The identifier of the user account in the directory used for identity management.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#directory_user_id ConnectUser#directory_user_id}
	DirectoryUserId *string `field:"optional" json:"directoryUserId" yaml:"directoryUserId"`
	// The identifier of the hierarchy group for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#hierarchy_group_arn ConnectUser#hierarchy_group_arn}
	HierarchyGroupArn *string `field:"optional" json:"hierarchyGroupArn" yaml:"hierarchyGroupArn"`
	// The information about the identity of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#identity_info ConnectUser#identity_info}
	IdentityInfo *ConnectUserIdentityInfo `field:"optional" json:"identityInfo" yaml:"identityInfo"`
	// The password for the user account.
	//
	// A password is required if you are using Amazon Connect for identity management. Otherwise, it is an error to include a password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#password ConnectUser#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#tags ConnectUser#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// One or more predefined attributes assigned to a user, with a level that indicates how skilled they are.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#user_proficiencies ConnectUser#user_proficiencies}
	UserProficiencies interface{} `field:"optional" json:"userProficiencies" yaml:"userProficiencies"`
}

