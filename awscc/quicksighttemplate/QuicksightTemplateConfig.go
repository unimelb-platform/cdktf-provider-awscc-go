package quicksighttemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_template#aws_account_id QuicksightTemplate#aws_account_id}.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// <p>The source entity of the template.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_template#source_entity QuicksightTemplate#source_entity}
	SourceEntity *QuicksightTemplateSourceEntity `field:"required" json:"sourceEntity" yaml:"sourceEntity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_template#template_id QuicksightTemplate#template_id}.
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
	// <p>A display name for the template.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_template#name QuicksightTemplate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// <p>A list of resource permissions to be set on the template. </p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_template#permissions QuicksightTemplate#permissions}
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_template#tags QuicksightTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// <p>A description of the current template version being created.
	//
	// This API operation creates the
	// 			first version of the template. Every time <code>UpdateTemplate</code> is called, a new
	// 			version is created. Each version of the template maintains a description of the version
	// 			in the <code>VersionDescription</code> field.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_template#version_description QuicksightTemplate#version_description}
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

