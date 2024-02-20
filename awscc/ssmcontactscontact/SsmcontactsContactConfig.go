package ssmcontactscontact

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsmcontactsContactConfig struct {
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
	// Alias of the contact. String value with 20 to 256 characters. Only alphabetical, numeric characters, dash, or underscore allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_contact#alias SsmcontactsContact#alias}
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// Name of the contact.
	//
	// String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_contact#display_name SsmcontactsContact#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Contact type, which specify type of contact. Currently supported values: ?PERSONAL?, ?SHARED?, ?OTHER?.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_contact#type SsmcontactsContact#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The stages that an escalation plan or engagement plan engages contacts and contact methods in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_contact#plan SsmcontactsContact#plan}
	Plan interface{} `field:"optional" json:"plan" yaml:"plan"`
}

