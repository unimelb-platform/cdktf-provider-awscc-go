package transferconnector


type TransferConnectorAs2Config struct {
	// ARN or name of the secret in AWS Secrets Manager which contains the credentials for Basic authentication.
	//
	// If empty, Basic authentication is disabled for the AS2 connector
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#basic_auth_secret_id TransferConnector#basic_auth_secret_id}
	BasicAuthSecretId *string `field:"optional" json:"basicAuthSecretId" yaml:"basicAuthSecretId"`
	// Compression setting for this AS2 connector configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#compression TransferConnector#compression}
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// Encryption algorithm for this AS2 connector configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#encryption_algorithm TransferConnector#encryption_algorithm}
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// A unique identifier for the local profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#local_profile_id TransferConnector#local_profile_id}
	LocalProfileId *string `field:"optional" json:"localProfileId" yaml:"localProfileId"`
	// MDN Response setting for this AS2 connector configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#mdn_response TransferConnector#mdn_response}
	MdnResponse *string `field:"optional" json:"mdnResponse" yaml:"mdnResponse"`
	// MDN Signing algorithm for this AS2 connector configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#mdn_signing_algorithm TransferConnector#mdn_signing_algorithm}
	MdnSigningAlgorithm *string `field:"optional" json:"mdnSigningAlgorithm" yaml:"mdnSigningAlgorithm"`
	// The message subject for this AS2 connector configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#message_subject TransferConnector#message_subject}
	MessageSubject *string `field:"optional" json:"messageSubject" yaml:"messageSubject"`
	// A unique identifier for the partner profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#partner_profile_id TransferConnector#partner_profile_id}
	PartnerProfileId *string `field:"optional" json:"partnerProfileId" yaml:"partnerProfileId"`
	// Signing algorithm for this AS2 connector configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#signing_algorithm TransferConnector#signing_algorithm}
	SigningAlgorithm *string `field:"optional" json:"signingAlgorithm" yaml:"signingAlgorithm"`
}

