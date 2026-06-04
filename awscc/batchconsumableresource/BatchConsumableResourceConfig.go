package batchconsumableresource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchConsumableResourceConfig struct {
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
	// Type of Consumable Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_consumable_resource#resource_type BatchConsumableResource#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Total Quantity of ConsumableResource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_consumable_resource#total_quantity BatchConsumableResource#total_quantity}
	TotalQuantity *float64 `field:"required" json:"totalQuantity" yaml:"totalQuantity"`
	// Name of ConsumableResource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_consumable_resource#consumable_resource_name BatchConsumableResource#consumable_resource_name}
	ConsumableResourceName *string `field:"optional" json:"consumableResourceName" yaml:"consumableResourceName"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_consumable_resource#tags BatchConsumableResource#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

