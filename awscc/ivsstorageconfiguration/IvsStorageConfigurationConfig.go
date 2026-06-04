package ivsstorageconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IvsStorageConfigurationConfig struct {
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
	// A complex type that describes an S3 location where recorded videos will be stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivs_storage_configuration#s3 IvsStorageConfiguration#s3}
	S3 *IvsStorageConfigurationS3 `field:"required" json:"s3" yaml:"s3"`
	// Storage Configuration Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivs_storage_configuration#name IvsStorageConfiguration#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of key-value pairs that contain metadata for the asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivs_storage_configuration#tags IvsStorageConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

