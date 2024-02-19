package iotfleetwisesignalcatalog

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotfleetwiseSignalCatalogConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#description IotfleetwiseSignalCatalog#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#name IotfleetwiseSignalCatalog#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#node_counts IotfleetwiseSignalCatalog#node_counts}.
	NodeCounts *IotfleetwiseSignalCatalogNodeCounts `field:"optional" json:"nodeCounts" yaml:"nodeCounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#nodes IotfleetwiseSignalCatalog#nodes}.
	Nodes interface{} `field:"optional" json:"nodes" yaml:"nodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#tags IotfleetwiseSignalCatalog#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

