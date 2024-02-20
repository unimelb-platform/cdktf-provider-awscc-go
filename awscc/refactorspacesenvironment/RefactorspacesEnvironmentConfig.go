package refactorspacesenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RefactorspacesEnvironmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/refactorspaces_environment#name RefactorspacesEnvironment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/refactorspaces_environment#network_fabric_type RefactorspacesEnvironment#network_fabric_type}.
	NetworkFabricType *string `field:"required" json:"networkFabricType" yaml:"networkFabricType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/refactorspaces_environment#description RefactorspacesEnvironment#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/refactorspaces_environment#tags RefactorspacesEnvironment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

