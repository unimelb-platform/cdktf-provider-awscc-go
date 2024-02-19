package configconformancepack


type ConfigConformancePackTemplateSsmDocumentDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_conformance_pack#document_name ConfigConformancePack#document_name}.
	DocumentName *string `field:"optional" json:"documentName" yaml:"documentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_conformance_pack#document_version ConfigConformancePack#document_version}.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
}

