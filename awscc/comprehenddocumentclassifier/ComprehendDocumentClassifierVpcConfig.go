package comprehenddocumentclassifier


type ComprehendDocumentClassifierVpcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/comprehend_document_classifier#security_group_ids ComprehendDocumentClassifier#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/comprehend_document_classifier#subnets ComprehendDocumentClassifier#subnets}.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

