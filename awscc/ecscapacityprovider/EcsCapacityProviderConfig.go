package ecscapacityprovider

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsCapacityProviderConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_capacity_provider#auto_scaling_group_provider EcsCapacityProvider#auto_scaling_group_provider}.
	AutoScalingGroupProvider *EcsCapacityProviderAutoScalingGroupProvider `field:"required" json:"autoScalingGroupProvider" yaml:"autoScalingGroupProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_capacity_provider#name EcsCapacityProvider#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_capacity_provider#tags EcsCapacityProvider#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

