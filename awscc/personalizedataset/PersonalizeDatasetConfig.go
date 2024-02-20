package personalizedataset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PersonalizeDatasetConfig struct {
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
	// The Amazon Resource Name (ARN) of the dataset group to add the dataset to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/personalize_dataset#dataset_group_arn PersonalizeDataset#dataset_group_arn}
	DatasetGroupArn *string `field:"required" json:"datasetGroupArn" yaml:"datasetGroupArn"`
	// The type of dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/personalize_dataset#dataset_type PersonalizeDataset#dataset_type}
	DatasetType *string `field:"required" json:"datasetType" yaml:"datasetType"`
	// The name for the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/personalize_dataset#name PersonalizeDataset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the schema to associate with the dataset. The schema defines the dataset fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/personalize_dataset#schema_arn PersonalizeDataset#schema_arn}
	SchemaArn *string `field:"required" json:"schemaArn" yaml:"schemaArn"`
	// Initial DatasetImportJob for the created dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/personalize_dataset#dataset_import_job PersonalizeDataset#dataset_import_job}
	DatasetImportJob *PersonalizeDatasetDatasetImportJob `field:"optional" json:"datasetImportJob" yaml:"datasetImportJob"`
}

