package pipespipe


type PipesPipeSourceParametersRabbitMqBrokerParametersCredentials struct {
	// Optional SecretManager ARN which stores the database credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#basic_auth PipesPipe#basic_auth}
	BasicAuth *string `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

