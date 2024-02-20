package forecastdatasetgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ForecastDatasetGroupConfig struct {
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
	// A name for the dataset group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/forecast_dataset_group#dataset_group_name ForecastDatasetGroup#dataset_group_name}
	DatasetGroupName *string `field:"required" json:"datasetGroupName" yaml:"datasetGroupName"`
	// The domain associated with the dataset group.
	//
	// When you add a dataset to a dataset group, this value and the value specified for the Domain parameter of the CreateDataset operation must match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/forecast_dataset_group#domain ForecastDatasetGroup#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// An array of Amazon Resource Names (ARNs) of the datasets that you want to include in the dataset group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/forecast_dataset_group#dataset_arns ForecastDatasetGroup#dataset_arns}
	DatasetArns *[]*string `field:"optional" json:"datasetArns" yaml:"datasetArns"`
	// The tags of Application Insights application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/forecast_dataset_group#tags ForecastDatasetGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

