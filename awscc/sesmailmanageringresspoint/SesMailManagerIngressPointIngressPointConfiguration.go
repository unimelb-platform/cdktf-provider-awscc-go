package sesmailmanageringresspoint


type SesMailManagerIngressPointIngressPointConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_ingress_point#secret_arn SesMailManagerIngressPoint#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_ingress_point#smtp_password SesMailManagerIngressPoint#smtp_password}.
	SmtpPassword *string `field:"optional" json:"smtpPassword" yaml:"smtpPassword"`
}

