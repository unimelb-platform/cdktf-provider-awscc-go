package ssmdocument


type SsmDocumentRequires struct {
	// The name of the required SSM document. The name can be an Amazon Resource Name (ARN).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_document#name SsmDocument#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The document version required by the current document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_document#version SsmDocument#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

