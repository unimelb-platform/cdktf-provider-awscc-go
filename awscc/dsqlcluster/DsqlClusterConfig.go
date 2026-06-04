package dsqlcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DsqlClusterConfig struct {
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
	// Whether deletion protection is enabled in this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/dsql_cluster#deletion_protection_enabled DsqlCluster#deletion_protection_enabled}
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// The KMS key that encrypts data on the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/dsql_cluster#kms_encryption_key DsqlCluster#kms_encryption_key}
	KmsEncryptionKey *string `field:"optional" json:"kmsEncryptionKey" yaml:"kmsEncryptionKey"`
	// The Multi-region properties associated to this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/dsql_cluster#multi_region_properties DsqlCluster#multi_region_properties}
	MultiRegionProperties *DsqlClusterMultiRegionProperties `field:"optional" json:"multiRegionProperties" yaml:"multiRegionProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/dsql_cluster#tags DsqlCluster#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

