package comprehenddocumentclassifier


type ComprehendDocumentClassifierInputDataConfigDocumentReaderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#document_read_action ComprehendDocumentClassifier#document_read_action}.
	DocumentReadAction *string `field:"required" json:"documentReadAction" yaml:"documentReadAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#document_read_mode ComprehendDocumentClassifier#document_read_mode}.
	DocumentReadMode *string `field:"optional" json:"documentReadMode" yaml:"documentReadMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#feature_types ComprehendDocumentClassifier#feature_types}.
	FeatureTypes *[]*string `field:"optional" json:"featureTypes" yaml:"featureTypes"`
}

