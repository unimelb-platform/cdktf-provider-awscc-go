package pcaconnectoradtemplategroupaccesscontrolentry

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PcaconnectoradTemplateGroupAccessControlEntryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template_group_access_control_entry#access_rights PcaconnectoradTemplateGroupAccessControlEntry#access_rights}.
	AccessRights *PcaconnectoradTemplateGroupAccessControlEntryAccessRights `field:"required" json:"accessRights" yaml:"accessRights"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template_group_access_control_entry#group_display_name PcaconnectoradTemplateGroupAccessControlEntry#group_display_name}.
	GroupDisplayName *string `field:"required" json:"groupDisplayName" yaml:"groupDisplayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template_group_access_control_entry#group_security_identifier PcaconnectoradTemplateGroupAccessControlEntry#group_security_identifier}.
	GroupSecurityIdentifier *string `field:"optional" json:"groupSecurityIdentifier" yaml:"groupSecurityIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template_group_access_control_entry#template_arn PcaconnectoradTemplateGroupAccessControlEntry#template_arn}.
	TemplateArn *string `field:"optional" json:"templateArn" yaml:"templateArn"`
}

