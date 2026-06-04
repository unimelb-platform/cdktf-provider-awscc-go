package iotsitewisedataset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotsitewiseDatasetConfig struct {
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
	// The name of the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotsitewise_dataset#dataset_name IotsitewiseDataset#dataset_name}
	DatasetName *string `field:"required" json:"datasetName" yaml:"datasetName"`
	// The data source for the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotsitewise_dataset#dataset_source IotsitewiseDataset#dataset_source}
	DatasetSource *IotsitewiseDatasetDatasetSource `field:"required" json:"datasetSource" yaml:"datasetSource"`
	// A description about the dataset, and its functionality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotsitewise_dataset#dataset_description IotsitewiseDataset#dataset_description}
	DatasetDescription *string `field:"optional" json:"datasetDescription" yaml:"datasetDescription"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotsitewise_dataset#tags IotsitewiseDataset#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

