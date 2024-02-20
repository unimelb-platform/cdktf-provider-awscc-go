package devopsgururesourcecollection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DevopsguruResourceCollectionConfig struct {
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
	// Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/devopsguru_resource_collection#resource_collection_filter DevopsguruResourceCollection#resource_collection_filter}
	ResourceCollectionFilter *DevopsguruResourceCollectionResourceCollectionFilter `field:"required" json:"resourceCollectionFilter" yaml:"resourceCollectionFilter"`
}

