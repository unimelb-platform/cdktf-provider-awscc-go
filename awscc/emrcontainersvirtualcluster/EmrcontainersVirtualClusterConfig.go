package emrcontainersvirtualcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrcontainersVirtualClusterConfig struct {
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
	// Container provider of the virtual cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrcontainers_virtual_cluster#container_provider EmrcontainersVirtualCluster#container_provider}
	ContainerProvider *EmrcontainersVirtualClusterContainerProvider `field:"required" json:"containerProvider" yaml:"containerProvider"`
	// Name of the virtual cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrcontainers_virtual_cluster#name EmrcontainersVirtualCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this virtual cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrcontainers_virtual_cluster#tags EmrcontainersVirtualCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

