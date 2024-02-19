package sesdedicatedippool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SesDedicatedIpPoolConfig struct {
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
	// The name of the dedicated IP pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_dedicated_ip_pool#pool_name SesDedicatedIpPool#pool_name}
	PoolName *string `field:"optional" json:"poolName" yaml:"poolName"`
	// Specifies whether the dedicated IP pool is managed or not. The default value is STANDARD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_dedicated_ip_pool#scaling_mode SesDedicatedIpPool#scaling_mode}
	ScalingMode *string `field:"optional" json:"scalingMode" yaml:"scalingMode"`
}

