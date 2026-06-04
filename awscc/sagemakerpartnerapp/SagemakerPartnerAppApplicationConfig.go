package sagemakerpartnerapp


type SagemakerPartnerAppApplicationConfig struct {
	// A list of users with administrator privileges for the PartnerApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_partner_app#admin_users SagemakerPartnerApp#admin_users}
	AdminUsers *[]*string `field:"optional" json:"adminUsers" yaml:"adminUsers"`
	// A list of arguments to pass to the PartnerApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_partner_app#arguments SagemakerPartnerApp#arguments}
	Arguments *map[string]*string `field:"optional" json:"arguments" yaml:"arguments"`
}

