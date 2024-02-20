package cleanroomscollaboration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CleanroomsCollaborationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#creator_display_name CleanroomsCollaboration#creator_display_name}.
	CreatorDisplayName *string `field:"required" json:"creatorDisplayName" yaml:"creatorDisplayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#creator_member_abilities CleanroomsCollaboration#creator_member_abilities}.
	CreatorMemberAbilities *[]*string `field:"required" json:"creatorMemberAbilities" yaml:"creatorMemberAbilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#description CleanroomsCollaboration#description}.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#members CleanroomsCollaboration#members}.
	Members interface{} `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#name CleanroomsCollaboration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#query_log_status CleanroomsCollaboration#query_log_status}.
	QueryLogStatus *string `field:"required" json:"queryLogStatus" yaml:"queryLogStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#creator_payment_configuration CleanroomsCollaboration#creator_payment_configuration}.
	CreatorPaymentConfiguration *CleanroomsCollaborationCreatorPaymentConfiguration `field:"optional" json:"creatorPaymentConfiguration" yaml:"creatorPaymentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#data_encryption_metadata CleanroomsCollaboration#data_encryption_metadata}.
	DataEncryptionMetadata *CleanroomsCollaborationDataEncryptionMetadata `field:"optional" json:"dataEncryptionMetadata" yaml:"dataEncryptionMetadata"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#tags CleanroomsCollaboration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

