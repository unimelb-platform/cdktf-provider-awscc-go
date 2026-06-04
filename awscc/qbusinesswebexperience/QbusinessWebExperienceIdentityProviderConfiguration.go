package qbusinesswebexperience


type QbusinessWebExperienceIdentityProviderConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/qbusiness_web_experience#open_id_connect_configuration QbusinessWebExperience#open_id_connect_configuration}.
	OpenIdConnectConfiguration *QbusinessWebExperienceIdentityProviderConfigurationOpenIdConnectConfiguration `field:"optional" json:"openIdConnectConfiguration" yaml:"openIdConnectConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/qbusiness_web_experience#saml_configuration QbusinessWebExperience#saml_configuration}.
	SamlConfiguration *QbusinessWebExperienceIdentityProviderConfigurationSamlConfiguration `field:"optional" json:"samlConfiguration" yaml:"samlConfiguration"`
}

