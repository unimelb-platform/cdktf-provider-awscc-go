package pipespipe


type PipesPipeSourceParametersSelfManagedKafkaParametersCredentials struct {
	// Optional SecretManager ARN which stores the database credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#basic_auth PipesPipe#basic_auth}
	BasicAuth *string `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// Optional SecretManager ARN which stores the database credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#client_certificate_tls_auth PipesPipe#client_certificate_tls_auth}
	ClientCertificateTlsAuth *string `field:"optional" json:"clientCertificateTlsAuth" yaml:"clientCertificateTlsAuth"`
	// Optional SecretManager ARN which stores the database credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#sasl_scram_256_auth PipesPipe#sasl_scram_256_auth}
	SaslScram256Auth *string `field:"optional" json:"saslScram256Auth" yaml:"saslScram256Auth"`
	// Optional SecretManager ARN which stores the database credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#sasl_scram_512_auth PipesPipe#sasl_scram_512_auth}
	SaslScram512Auth *string `field:"optional" json:"saslScram512Auth" yaml:"saslScram512Auth"`
}

