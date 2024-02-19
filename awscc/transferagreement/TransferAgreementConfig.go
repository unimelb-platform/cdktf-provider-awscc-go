package transferagreement

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransferAgreementConfig struct {
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
	// Specifies the access role for the agreement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_agreement#access_role TransferAgreement#access_role}
	AccessRole *string `field:"required" json:"accessRole" yaml:"accessRole"`
	// Specifies the base directory for the agreement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_agreement#base_directory TransferAgreement#base_directory}
	BaseDirectory *string `field:"required" json:"baseDirectory" yaml:"baseDirectory"`
	// A unique identifier for the local profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_agreement#local_profile_id TransferAgreement#local_profile_id}
	LocalProfileId *string `field:"required" json:"localProfileId" yaml:"localProfileId"`
	// A unique identifier for the partner profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_agreement#partner_profile_id TransferAgreement#partner_profile_id}
	PartnerProfileId *string `field:"required" json:"partnerProfileId" yaml:"partnerProfileId"`
	// A unique identifier for the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_agreement#server_id TransferAgreement#server_id}
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// A textual description for the agreement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_agreement#description TransferAgreement#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the status of the agreement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_agreement#status TransferAgreement#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Key-value pairs that can be used to group and search for agreements.
	//
	// Tags are metadata attached to agreements for any purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_agreement#tags TransferAgreement#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

