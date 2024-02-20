package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV2 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#certificate_validity PcaconnectoradTemplate#certificate_validity}.
	CertificateValidity *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidity `field:"required" json:"certificateValidity" yaml:"certificateValidity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#enrollment_flags PcaconnectoradTemplate#enrollment_flags}.
	EnrollmentFlags *PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlags `field:"required" json:"enrollmentFlags" yaml:"enrollmentFlags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#extensions PcaconnectoradTemplate#extensions}.
	Extensions *PcaconnectoradTemplateDefinitionTemplateV2Extensions `field:"required" json:"extensions" yaml:"extensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#general_flags PcaconnectoradTemplate#general_flags}.
	GeneralFlags *PcaconnectoradTemplateDefinitionTemplateV2GeneralFlags `field:"required" json:"generalFlags" yaml:"generalFlags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#private_key_attributes PcaconnectoradTemplate#private_key_attributes}.
	PrivateKeyAttributes *PcaconnectoradTemplateDefinitionTemplateV2PrivateKeyAttributes `field:"required" json:"privateKeyAttributes" yaml:"privateKeyAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#private_key_flags PcaconnectoradTemplate#private_key_flags}.
	PrivateKeyFlags *PcaconnectoradTemplateDefinitionTemplateV2PrivateKeyFlags `field:"required" json:"privateKeyFlags" yaml:"privateKeyFlags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#subject_name_flags PcaconnectoradTemplate#subject_name_flags}.
	SubjectNameFlags *PcaconnectoradTemplateDefinitionTemplateV2SubjectNameFlags `field:"required" json:"subjectNameFlags" yaml:"subjectNameFlags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#superseded_templates PcaconnectoradTemplate#superseded_templates}.
	SupersededTemplates *[]*string `field:"optional" json:"supersededTemplates" yaml:"supersededTemplates"`
}

