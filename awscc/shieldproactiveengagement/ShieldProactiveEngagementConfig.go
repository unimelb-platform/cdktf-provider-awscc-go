package shieldproactiveengagement

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ShieldProactiveEngagementConfig struct {
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
	// A list of email addresses and phone numbers that the Shield Response Team (SRT) can use to contact you for escalations to the SRT and to initiate proactive customer support.
	//
	// To enable proactive engagement, the contact list must include at least one phone number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/shield_proactive_engagement#emergency_contact_list ShieldProactiveEngagement#emergency_contact_list}
	EmergencyContactList interface{} `field:"required" json:"emergencyContactList" yaml:"emergencyContactList"`
	// If `ENABLED`, the Shield Response Team (SRT) will use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.
	//
	// If `DISABLED`, the SRT will not proactively notify contacts about escalations or to initiate proactive customer support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/shield_proactive_engagement#proactive_engagement_status ShieldProactiveEngagement#proactive_engagement_status}
	ProactiveEngagementStatus *string `field:"required" json:"proactiveEngagementStatus" yaml:"proactiveEngagementStatus"`
}

