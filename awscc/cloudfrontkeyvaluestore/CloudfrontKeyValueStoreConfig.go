package cloudfrontkeyvaluestore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontKeyValueStoreConfig struct {
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
	// The name of the key value store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudfront_key_value_store#name CloudfrontKeyValueStore#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A comment for the key value store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudfront_key_value_store#comment CloudfrontKeyValueStore#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The import source for the key value store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudfront_key_value_store#import_source CloudfrontKeyValueStore#import_source}
	ImportSource *CloudfrontKeyValueStoreImportSource `field:"optional" json:"importSource" yaml:"importSource"`
}

