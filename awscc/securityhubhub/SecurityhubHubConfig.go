package securityhubhub

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityhubHubConfig struct {
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
	// Whether to automatically enable new controls when they are added to standards that are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/securityhub_hub#auto_enable_controls SecurityhubHub#auto_enable_controls}
	AutoEnableControls interface{} `field:"optional" json:"autoEnableControls" yaml:"autoEnableControls"`
	// This field, used when enabling Security Hub, specifies whether the calling account has consolidated control findings turned on.
	//
	// If the value for this field is set to SECURITY_CONTROL, Security Hub generates a single finding for a control check even when the check applies to multiple enabled standards.  If the value for this field is set to STANDARD_CONTROL, Security Hub generates separate findings for a control check when the check applies to multiple enabled standards.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/securityhub_hub#control_finding_generator SecurityhubHub#control_finding_generator}
	ControlFindingGenerator *string `field:"optional" json:"controlFindingGenerator" yaml:"controlFindingGenerator"`
	// Whether to enable the security standards that Security Hub has designated as automatically enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/securityhub_hub#enable_default_standards SecurityhubHub#enable_default_standards}
	EnableDefaultStandards interface{} `field:"optional" json:"enableDefaultStandards" yaml:"enableDefaultStandards"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/securityhub_hub#tags SecurityhubHub#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

