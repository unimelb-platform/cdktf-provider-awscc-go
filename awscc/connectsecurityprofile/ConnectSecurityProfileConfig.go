package connectsecurityprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConnectSecurityProfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_security_profile#instance_arn ConnectSecurityProfile#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the security profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_security_profile#security_profile_name ConnectSecurityProfile#security_profile_name}
	SecurityProfileName *string `field:"required" json:"securityProfileName" yaml:"securityProfileName"`
	// The list of tags that a security profile uses to restrict access to resources in Amazon Connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_security_profile#allowed_access_control_tags ConnectSecurityProfile#allowed_access_control_tags}
	AllowedAccessControlTags interface{} `field:"optional" json:"allowedAccessControlTags" yaml:"allowedAccessControlTags"`
	// The description of the security profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_security_profile#description ConnectSecurityProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Permissions assigned to the security profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_security_profile#permissions ConnectSecurityProfile#permissions}
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// The list of resources that a security profile applies tag restrictions to in Amazon Connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_security_profile#tag_restricted_resources ConnectSecurityProfile#tag_restricted_resources}
	TagRestrictedResources *[]*string `field:"optional" json:"tagRestrictedResources" yaml:"tagRestrictedResources"`
	// The tags used to organize, track, or control access for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_security_profile#tags ConnectSecurityProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

