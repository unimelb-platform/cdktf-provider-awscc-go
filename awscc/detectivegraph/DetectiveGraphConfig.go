package detectivegraph

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DetectiveGraphConfig struct {
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
	// Indicates whether to automatically enable new organization accounts as member accounts in the organization behavior graph.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/detective_graph#auto_enable_members DetectiveGraph#auto_enable_members}
	AutoEnableMembers interface{} `field:"optional" json:"autoEnableMembers" yaml:"autoEnableMembers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/detective_graph#tags DetectiveGraph#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

