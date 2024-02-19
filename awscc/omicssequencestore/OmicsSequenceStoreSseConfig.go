package omicssequencestore


type OmicsSequenceStoreSseConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_sequence_store#type OmicsSequenceStore#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// An encryption key ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_sequence_store#key_arn OmicsSequenceStore#key_arn}
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}

