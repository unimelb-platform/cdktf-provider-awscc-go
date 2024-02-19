package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointSegmentEncryption struct {
	// <p>The encryption type.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#encryption_method Mediapackagev2OriginEndpoint#encryption_method}
	EncryptionMethod *Mediapackagev2OriginEndpointSegmentEncryptionEncryptionMethod `field:"required" json:"encryptionMethod" yaml:"encryptionMethod"`
	// <p>The parameters for the SPEKE key provider.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#speke_key_provider Mediapackagev2OriginEndpoint#speke_key_provider}
	SpekeKeyProvider *Mediapackagev2OriginEndpointSegmentEncryptionSpekeKeyProvider `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
	// <p>A 128-bit, 16-byte hex value represented by a 32-character string, used in conjunction with the key for encrypting content.
	//
	// If you don't specify a value, then MediaPackage creates the constant initialization vector (IV).</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#constant_initialization_vector Mediapackagev2OriginEndpoint#constant_initialization_vector}
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// <p>The frequency (in seconds) of key changes for live workflows, in which content is streamed real time.
	//
	// The service retrieves content keys before the live content begins streaming, and then retrieves them as needed over the lifetime of the workflow. By default, key rotation is set to 300 seconds (5 minutes), the minimum rotation interval, which is equivalent to setting it to 300. If you don't enter an interval, content keys aren't rotated.</p>
	//          <p>The following example setting causes the service to rotate keys every thirty minutes: <code>1800</code>
	//          </p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#key_rotation_interval_seconds Mediapackagev2OriginEndpoint#key_rotation_interval_seconds}
	KeyRotationIntervalSeconds *float64 `field:"optional" json:"keyRotationIntervalSeconds" yaml:"keyRotationIntervalSeconds"`
}

