package athenadatacatalog

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AthenaDataCatalogConfig struct {
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
	// The name of the data catalog to create.
	//
	// The catalog name must be unique for the AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or hyphen characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_data_catalog#name AthenaDataCatalog#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of data catalog to create: LAMBDA for a federated catalog, GLUE for AWS Glue Catalog, or HIVE for an external hive metastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_data_catalog#type AthenaDataCatalog#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A description of the data catalog to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_data_catalog#description AthenaDataCatalog#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the Lambda function or functions to use for creating the data catalog.
	//
	// This is a mapping whose values depend on the catalog type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_data_catalog#parameters AthenaDataCatalog#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// A list of comma separated tags to add to the data catalog that is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_data_catalog#tags AthenaDataCatalog#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

