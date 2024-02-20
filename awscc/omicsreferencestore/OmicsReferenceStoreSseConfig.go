package omicsreferencestore


type OmicsReferenceStoreSseConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_reference_store#type OmicsReferenceStore#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// An encryption key ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_reference_store#key_arn OmicsReferenceStore#key_arn}
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}

