package comprehendflywheel


type ComprehendFlywheelTaskConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel#language_code ComprehendFlywheel#language_code}.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel#document_classification_config ComprehendFlywheel#document_classification_config}.
	DocumentClassificationConfig *ComprehendFlywheelTaskConfigDocumentClassificationConfig `field:"optional" json:"documentClassificationConfig" yaml:"documentClassificationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel#entity_recognition_config ComprehendFlywheel#entity_recognition_config}.
	EntityRecognitionConfig *ComprehendFlywheelTaskConfigEntityRecognitionConfig `field:"optional" json:"entityRecognitionConfig" yaml:"entityRecognitionConfig"`
}

