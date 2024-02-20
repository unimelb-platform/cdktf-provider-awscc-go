package mediaconnectflow


type MediaconnectFlowSourceDecryption struct {
	// The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow#role_arn MediaconnectFlow#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow#algorithm MediaconnectFlow#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content.
	//
	// This parameter is not valid for static key encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow#constant_initialization_vector MediaconnectFlow#constant_initialization_vector}
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// The value of one of the devices that you configured with your digital rights management (DRM) platform key provider.
	//
	// This parameter is required for SPEKE encryption and is not valid for static key encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow#device_id MediaconnectFlow#device_id}
	DeviceId *string `field:"optional" json:"deviceId" yaml:"deviceId"`
	// The type of key that is used for the encryption.
	//
	// If no keyType is provided, the service will use the default setting (static-key).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow#key_type MediaconnectFlow#key_type}
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// The AWS Region that the API Gateway proxy endpoint was created in.
	//
	// This parameter is required for SPEKE encryption and is not valid for static key encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow#region MediaconnectFlow#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// An identifier for the content.
	//
	// The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow#resource_id MediaconnectFlow#resource_id}
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// The ARN of the secret that you created in AWS Secrets Manager to store the encryption key.
	//
	// This parameter is required for static key encryption and is not valid for SPEKE encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow#secret_arn MediaconnectFlow#secret_arn}
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// The URL from the API Gateway proxy that you set up to talk to your key server.
	//
	// This parameter is required for SPEKE encryption and is not valid for static key encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow#url MediaconnectFlow#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

