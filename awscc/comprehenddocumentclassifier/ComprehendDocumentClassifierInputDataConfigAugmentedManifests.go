package comprehenddocumentclassifier


type ComprehendDocumentClassifierInputDataConfigAugmentedManifests struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#attribute_names ComprehendDocumentClassifier#attribute_names}.
	AttributeNames *[]*string `field:"required" json:"attributeNames" yaml:"attributeNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#s3_uri ComprehendDocumentClassifier#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#split ComprehendDocumentClassifier#split}.
	Split *string `field:"optional" json:"split" yaml:"split"`
}

