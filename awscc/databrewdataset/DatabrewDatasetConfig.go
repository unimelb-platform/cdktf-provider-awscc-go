package databrewdataset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabrewDatasetConfig struct {
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
	// Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#input DatabrewDataset#input}
	Input *DatabrewDatasetInput `field:"required" json:"input" yaml:"input"`
	// Dataset name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#name DatabrewDataset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Dataset format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#format DatabrewDataset#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Format options for dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#format_options DatabrewDataset#format_options}
	FormatOptions *DatabrewDatasetFormatOptions `field:"optional" json:"formatOptions" yaml:"formatOptions"`
	// PathOptions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#path_options DatabrewDataset#path_options}
	PathOptions *DatabrewDatasetPathOptions `field:"optional" json:"pathOptions" yaml:"pathOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#tags DatabrewDataset#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

