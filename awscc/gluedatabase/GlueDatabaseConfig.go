package gluedatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlueDatabaseConfig struct {
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
	// The AWS account ID for the account in which to create the catalog object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/glue_database#catalog_id GlueDatabase#catalog_id}
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The metadata for the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/glue_database#database_input GlueDatabase#database_input}
	DatabaseInput *GlueDatabaseDatabaseInput `field:"required" json:"databaseInput" yaml:"databaseInput"`
	// The name of the database. For hive compatibility, this is folded to lowercase when it is store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/glue_database#database_name GlueDatabase#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
}

