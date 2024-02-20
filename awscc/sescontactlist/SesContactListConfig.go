package sescontactlist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SesContactListConfig struct {
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
	// The name of the contact list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_contact_list#contact_list_name SesContactList#contact_list_name}
	ContactListName *string `field:"optional" json:"contactListName" yaml:"contactListName"`
	// The description of the contact list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_contact_list#description SesContactList#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags (keys and values) associated with the contact list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_contact_list#tags SesContactList#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The topics associated with the contact list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_contact_list#topics SesContactList#topics}
	Topics interface{} `field:"optional" json:"topics" yaml:"topics"`
}

