package sesmailmanagertrafficpolicy


type SesMailManagerTrafficPolicyPolicyStatements struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_traffic_policy#action SesMailManagerTrafficPolicy#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_traffic_policy#conditions SesMailManagerTrafficPolicy#conditions}.
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
}

