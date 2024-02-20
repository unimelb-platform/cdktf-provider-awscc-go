package signersigningprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SignerSigningProfileConfig struct {
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
	// The ID of the target signing platform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/signer_signing_profile#platform_id SignerSigningProfile#platform_id}
	PlatformId *string `field:"required" json:"platformId" yaml:"platformId"`
	// Signature validity period of the profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/signer_signing_profile#signature_validity_period SignerSigningProfile#signature_validity_period}
	SignatureValidityPeriod *SignerSigningProfileSignatureValidityPeriod `field:"optional" json:"signatureValidityPeriod" yaml:"signatureValidityPeriod"`
	// A list of tags associated with the signing profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/signer_signing_profile#tags SignerSigningProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

