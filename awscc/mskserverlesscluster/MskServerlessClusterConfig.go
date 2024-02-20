package mskserverlesscluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MskServerlessClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_serverless_cluster#client_authentication MskServerlessCluster#client_authentication}.
	ClientAuthentication *MskServerlessClusterClientAuthentication `field:"required" json:"clientAuthentication" yaml:"clientAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_serverless_cluster#cluster_name MskServerlessCluster#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_serverless_cluster#vpc_configs MskServerlessCluster#vpc_configs}.
	VpcConfigs interface{} `field:"required" json:"vpcConfigs" yaml:"vpcConfigs"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_serverless_cluster#tags MskServerlessCluster#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

