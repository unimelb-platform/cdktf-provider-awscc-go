package backupframework

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupFrameworkConfig struct {
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
	// Contains detailed information about all of the controls of a framework. Each framework must contain at least one control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_framework#framework_controls BackupFramework#framework_controls}
	FrameworkControls interface{} `field:"required" json:"frameworkControls" yaml:"frameworkControls"`
	// An optional description of the framework with a maximum 1,024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_framework#framework_description BackupFramework#framework_description}
	FrameworkDescription *string `field:"optional" json:"frameworkDescription" yaml:"frameworkDescription"`
	// The unique name of a framework.
	//
	// This name is between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_framework#framework_name BackupFramework#framework_name}
	FrameworkName *string `field:"optional" json:"frameworkName" yaml:"frameworkName"`
	// Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_framework#framework_tags BackupFramework#framework_tags}
	FrameworkTags interface{} `field:"optional" json:"frameworkTags" yaml:"frameworkTags"`
}

