package ssmcontactsplan

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsmcontactsPlanConfig struct {
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
	// Contact ID for the AWS SSM Incident Manager Contact to associate the plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_plan#contact_id SsmcontactsPlan#contact_id}
	ContactId *string `field:"optional" json:"contactId" yaml:"contactId"`
	// Rotation Ids to associate with Oncall Contact for engagement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_plan#rotation_ids SsmcontactsPlan#rotation_ids}
	RotationIds *[]*string `field:"optional" json:"rotationIds" yaml:"rotationIds"`
	// The stages that an escalation plan or engagement plan engages contacts and contact methods in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_plan#stages SsmcontactsPlan#stages}
	Stages interface{} `field:"optional" json:"stages" yaml:"stages"`
}

