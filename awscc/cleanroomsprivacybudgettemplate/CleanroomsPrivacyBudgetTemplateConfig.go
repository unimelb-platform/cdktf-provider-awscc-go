package cleanroomsprivacybudgettemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CleanroomsPrivacyBudgetTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_privacy_budget_template#auto_refresh CleanroomsPrivacyBudgetTemplate#auto_refresh}.
	AutoRefresh *string `field:"required" json:"autoRefresh" yaml:"autoRefresh"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_privacy_budget_template#membership_identifier CleanroomsPrivacyBudgetTemplate#membership_identifier}.
	MembershipIdentifier *string `field:"required" json:"membershipIdentifier" yaml:"membershipIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_privacy_budget_template#parameters CleanroomsPrivacyBudgetTemplate#parameters}.
	Parameters *CleanroomsPrivacyBudgetTemplateParameters `field:"required" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_privacy_budget_template#privacy_budget_type CleanroomsPrivacyBudgetTemplate#privacy_budget_type}.
	PrivacyBudgetType *string `field:"required" json:"privacyBudgetType" yaml:"privacyBudgetType"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms privacy budget template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_privacy_budget_template#tags CleanroomsPrivacyBudgetTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

