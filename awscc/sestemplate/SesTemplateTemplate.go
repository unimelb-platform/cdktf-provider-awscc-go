package sestemplate


type SesTemplateTemplate struct {
	// The subject line of the email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_template#subject_part SesTemplate#subject_part}
	SubjectPart *string `field:"required" json:"subjectPart" yaml:"subjectPart"`
	// The HTML body of the email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_template#html_part SesTemplate#html_part}
	HtmlPart *string `field:"optional" json:"htmlPart" yaml:"htmlPart"`
	// The name of the template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_template#template_name SesTemplate#template_name}
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
	// The email body that is visible to recipients whose email clients do not display HTML content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_template#text_part SesTemplate#text_part}
	TextPart *string `field:"optional" json:"textPart" yaml:"textPart"`
}

