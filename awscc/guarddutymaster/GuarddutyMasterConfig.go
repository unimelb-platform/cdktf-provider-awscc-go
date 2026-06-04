package guarddutymaster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GuarddutyMasterConfig struct {
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
	// Unique ID of the detector of the GuardDuty member account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/guardduty_master#detector_id GuarddutyMaster#detector_id}
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// ID of the account used as the master account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/guardduty_master#master_id GuarddutyMaster#master_id}
	MasterId *string `field:"required" json:"masterId" yaml:"masterId"`
	// Value used to validate the master account to the member account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/guardduty_master#invitation_id GuarddutyMaster#invitation_id}
	InvitationId *string `field:"optional" json:"invitationId" yaml:"invitationId"`
}

