package cleanroomsmembership

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CleanroomsMembershipConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_membership#collaboration_identifier CleanroomsMembership#collaboration_identifier}.
	CollaborationIdentifier *string `field:"required" json:"collaborationIdentifier" yaml:"collaborationIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_membership#query_log_status CleanroomsMembership#query_log_status}.
	QueryLogStatus *string `field:"required" json:"queryLogStatus" yaml:"queryLogStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_membership#default_result_configuration CleanroomsMembership#default_result_configuration}.
	DefaultResultConfiguration *CleanroomsMembershipDefaultResultConfiguration `field:"optional" json:"defaultResultConfiguration" yaml:"defaultResultConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_membership#payment_configuration CleanroomsMembership#payment_configuration}.
	PaymentConfiguration *CleanroomsMembershipPaymentConfiguration `field:"optional" json:"paymentConfiguration" yaml:"paymentConfiguration"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms membership.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_membership#tags CleanroomsMembership#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

