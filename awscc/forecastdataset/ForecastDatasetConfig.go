package forecastdataset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ForecastDatasetConfig struct {
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
	// A name for the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/forecast_dataset#dataset_name ForecastDataset#dataset_name}
	DatasetName *string `field:"required" json:"datasetName" yaml:"datasetName"`
	// The dataset type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/forecast_dataset#dataset_type ForecastDataset#dataset_type}
	DatasetType *string `field:"required" json:"datasetType" yaml:"datasetType"`
	// The domain associated with the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/forecast_dataset#domain ForecastDataset#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/forecast_dataset#schema ForecastDataset#schema}.
	Schema *ForecastDatasetSchema `field:"required" json:"schema" yaml:"schema"`
	// Frequency of data collection. This parameter is required for RELATED_TIME_SERIES.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/forecast_dataset#data_frequency ForecastDataset#data_frequency}
	DataFrequency *string `field:"optional" json:"dataFrequency" yaml:"dataFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/forecast_dataset#encryption_config ForecastDataset#encryption_config}.
	EncryptionConfig *ForecastDatasetEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/forecast_dataset#tags ForecastDataset#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

