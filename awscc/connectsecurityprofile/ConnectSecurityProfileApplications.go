package connectsecurityprofile


type ConnectSecurityProfileApplications struct {
	// The permissions that the agent is granted on the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_security_profile#application_permissions ConnectSecurityProfile#application_permissions}
	ApplicationPermissions *[]*string `field:"optional" json:"applicationPermissions" yaml:"applicationPermissions"`
	// Namespace of the application that you want to give access to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_security_profile#namespace ConnectSecurityProfile#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

