package iotfleetwisestatetemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotfleetwiseStateTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotfleetwise_state_template#name IotfleetwiseStateTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotfleetwise_state_template#signal_catalog_arn IotfleetwiseStateTemplate#signal_catalog_arn}.
	SignalCatalogArn *string `field:"required" json:"signalCatalogArn" yaml:"signalCatalogArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotfleetwise_state_template#state_template_properties IotfleetwiseStateTemplate#state_template_properties}.
	StateTemplateProperties *[]*string `field:"required" json:"stateTemplateProperties" yaml:"stateTemplateProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotfleetwise_state_template#data_extra_dimensions IotfleetwiseStateTemplate#data_extra_dimensions}.
	DataExtraDimensions *[]*string `field:"optional" json:"dataExtraDimensions" yaml:"dataExtraDimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotfleetwise_state_template#description IotfleetwiseStateTemplate#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotfleetwise_state_template#metadata_extra_dimensions IotfleetwiseStateTemplate#metadata_extra_dimensions}.
	MetadataExtraDimensions *[]*string `field:"optional" json:"metadataExtraDimensions" yaml:"metadataExtraDimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotfleetwise_state_template#tags IotfleetwiseStateTemplate#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

