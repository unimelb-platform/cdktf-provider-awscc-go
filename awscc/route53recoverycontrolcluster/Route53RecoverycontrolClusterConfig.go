package route53recoverycontrolcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53RecoverycontrolClusterConfig struct {
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
	// Name of a Cluster. You can use any non-white space character in the name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53recoverycontrol_cluster#name Route53RecoverycontrolCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Cluster supports IPv4 endpoints and Dual-stack IPv4 and IPv6 endpoints. NetworkType can be IPV4 or DUALSTACK.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53recoverycontrol_cluster#network_type Route53RecoverycontrolCluster#network_type}
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53recoverycontrol_cluster#tags Route53RecoverycontrolCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

