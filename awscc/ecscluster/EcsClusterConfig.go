package ecscluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#capacity_providers EcsCluster#capacity_providers}.
	CapacityProviders *[]*string `field:"optional" json:"capacityProviders" yaml:"capacityProviders"`
	// A user-generated string that you use to identify your cluster.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID for the name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#cluster_name EcsCluster#cluster_name}
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#cluster_settings EcsCluster#cluster_settings}.
	ClusterSettings interface{} `field:"optional" json:"clusterSettings" yaml:"clusterSettings"`
	// The configurations to be set at cluster level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#configuration EcsCluster#configuration}
	Configuration *EcsClusterConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#default_capacity_provider_strategy EcsCluster#default_capacity_provider_strategy}.
	DefaultCapacityProviderStrategy interface{} `field:"optional" json:"defaultCapacityProviderStrategy" yaml:"defaultCapacityProviderStrategy"`
	// Service Connect Configuration default for all services or tasks within this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#service_connect_defaults EcsCluster#service_connect_defaults}
	ServiceConnectDefaults *EcsClusterServiceConnectDefaults `field:"optional" json:"serviceConnectDefaults" yaml:"serviceConnectDefaults"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#tags EcsCluster#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

