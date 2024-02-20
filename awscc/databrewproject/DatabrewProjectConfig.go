package databrewproject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabrewProjectConfig struct {
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
	// Dataset name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_project#dataset_name DatabrewProject#dataset_name}
	DatasetName *string `field:"required" json:"datasetName" yaml:"datasetName"`
	// Project name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_project#name DatabrewProject#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Recipe name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_project#recipe_name DatabrewProject#recipe_name}
	RecipeName *string `field:"required" json:"recipeName" yaml:"recipeName"`
	// Role arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_project#role_arn DatabrewProject#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Sample.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_project#sample DatabrewProject#sample}
	Sample *DatabrewProjectSample `field:"optional" json:"sample" yaml:"sample"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_project#tags DatabrewProject#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

