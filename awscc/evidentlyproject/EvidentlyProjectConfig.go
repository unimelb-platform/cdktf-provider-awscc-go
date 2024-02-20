package evidentlyproject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EvidentlyProjectConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_project#name EvidentlyProject#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_project#app_config_resource EvidentlyProject#app_config_resource}.
	AppConfigResource *EvidentlyProjectAppConfigResource `field:"optional" json:"appConfigResource" yaml:"appConfigResource"`
	// Destinations for data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_project#data_delivery EvidentlyProject#data_delivery}
	DataDelivery *EvidentlyProjectDataDelivery `field:"optional" json:"dataDelivery" yaml:"dataDelivery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_project#description EvidentlyProject#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_project#tags EvidentlyProject#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

