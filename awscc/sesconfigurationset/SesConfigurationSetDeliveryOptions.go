package sesconfigurationset


type SesConfigurationSetDeliveryOptions struct {
	// The name of the dedicated IP pool to associate with the configuration set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_configuration_set#sending_pool_name SesConfigurationSet#sending_pool_name}
	SendingPoolName *string `field:"optional" json:"sendingPoolName" yaml:"sendingPoolName"`
	// Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS).
	//
	// If the value is Require , messages are only delivered if a TLS connection can be established. If the value is Optional , messages can be delivered in plain text if a TLS connection can't be established.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_configuration_set#tls_policy SesConfigurationSet#tls_policy}
	TlsPolicy *string `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
}

