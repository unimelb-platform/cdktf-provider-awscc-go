package transferprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransferProfileConfig struct {
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
	// AS2 identifier agreed with a trading partner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_profile#as_2_id TransferProfile#as_2_id}
	As2Id *string `field:"required" json:"as2Id" yaml:"as2Id"`
	// Enum specifying whether the profile is local or associated with a trading partner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_profile#profile_type TransferProfile#profile_type}
	ProfileType *string `field:"required" json:"profileType" yaml:"profileType"`
	// List of the certificate IDs associated with this profile to be used for encryption and signing of AS2 messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_profile#certificate_ids TransferProfile#certificate_ids}
	CertificateIds *[]*string `field:"optional" json:"certificateIds" yaml:"certificateIds"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_profile#tags TransferProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

