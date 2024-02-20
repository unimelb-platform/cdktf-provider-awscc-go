package healthimagingdatastore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthimagingDatastoreConfig struct {
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
	// User friendly name for Datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/healthimaging_datastore#datastore_name HealthimagingDatastore#datastore_name}
	DatastoreName *string `field:"optional" json:"datastoreName" yaml:"datastoreName"`
	// ARN referencing a KMS key or KMS key alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/healthimaging_datastore#kms_key_arn HealthimagingDatastore#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// A Map of key value pairs for Tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/healthimaging_datastore#tags HealthimagingDatastore#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

