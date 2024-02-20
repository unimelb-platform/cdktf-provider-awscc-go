package cloudfrontdistribution


type CloudfrontDistributionDistributionConfigOriginsCustomOriginConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#origin_protocol_policy CloudfrontDistribution#origin_protocol_policy}.
	OriginProtocolPolicy *string `field:"required" json:"originProtocolPolicy" yaml:"originProtocolPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#http_port CloudfrontDistribution#http_port}.
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#https_port CloudfrontDistribution#https_port}.
	HttpsPort *float64 `field:"optional" json:"httpsPort" yaml:"httpsPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#origin_keepalive_timeout CloudfrontDistribution#origin_keepalive_timeout}.
	OriginKeepaliveTimeout *float64 `field:"optional" json:"originKeepaliveTimeout" yaml:"originKeepaliveTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#origin_read_timeout CloudfrontDistribution#origin_read_timeout}.
	OriginReadTimeout *float64 `field:"optional" json:"originReadTimeout" yaml:"originReadTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#origin_ssl_protocols CloudfrontDistribution#origin_ssl_protocols}.
	OriginSslProtocols *[]*string `field:"optional" json:"originSslProtocols" yaml:"originSslProtocols"`
}

