package pinpointinapptemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PinpointInAppTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pinpoint_in_app_template#template_name PinpointInAppTemplate#template_name}.
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pinpoint_in_app_template#content PinpointInAppTemplate#content}.
	Content interface{} `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pinpoint_in_app_template#custom_config PinpointInAppTemplate#custom_config}.
	CustomConfig *string `field:"optional" json:"customConfig" yaml:"customConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pinpoint_in_app_template#layout PinpointInAppTemplate#layout}.
	Layout *string `field:"optional" json:"layout" yaml:"layout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pinpoint_in_app_template#tags PinpointInAppTemplate#tags}.
	Tags *string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pinpoint_in_app_template#template_description PinpointInAppTemplate#template_description}.
	TemplateDescription *string `field:"optional" json:"templateDescription" yaml:"templateDescription"`
}

