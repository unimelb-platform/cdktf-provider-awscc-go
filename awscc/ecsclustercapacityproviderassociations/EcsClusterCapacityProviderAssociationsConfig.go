package ecsclustercapacityproviderassociations

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsClusterCapacityProviderAssociationsConfig struct {
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
	// List of capacity providers to associate with the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster_capacity_provider_associations#capacity_providers EcsClusterCapacityProviderAssociations#capacity_providers}
	CapacityProviders *[]*string `field:"required" json:"capacityProviders" yaml:"capacityProviders"`
	// The name of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster_capacity_provider_associations#cluster EcsClusterCapacityProviderAssociations#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// List of capacity providers to associate with the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster_capacity_provider_associations#default_capacity_provider_strategy EcsClusterCapacityProviderAssociations#default_capacity_provider_strategy}
	DefaultCapacityProviderStrategy interface{} `field:"required" json:"defaultCapacityProviderStrategy" yaml:"defaultCapacityProviderStrategy"`
}

