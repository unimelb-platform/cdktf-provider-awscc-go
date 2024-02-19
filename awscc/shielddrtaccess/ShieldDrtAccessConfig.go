package shielddrtaccess

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ShieldDrtAccessConfig struct {
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
	// Authorizes the Shield Response Team (SRT) using the specified role, to access your AWS account to assist with DDoS attack mitigation during potential attacks.
	//
	// This enables the SRT to inspect your AWS WAF configuration and create or update AWS WAF rules and web ACLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/shield_drt_access#role_arn ShieldDrtAccess#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Authorizes the Shield Response Team (SRT) to access the specified Amazon S3 bucket containing log data such as Application Load Balancer access logs, CloudFront logs, or logs from third party sources.
	//
	// You can associate up to 10 Amazon S3 buckets with your subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/shield_drt_access#log_bucket_list ShieldDrtAccess#log_bucket_list}
	LogBucketList *[]*string `field:"optional" json:"logBucketList" yaml:"logBucketList"`
}

