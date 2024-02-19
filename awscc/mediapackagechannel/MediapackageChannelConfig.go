package mediapackagechannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediapackageChannelConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the Channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_channel#id MediapackageChannel#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// A short text description of the Channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_channel#description MediapackageChannel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration parameters for egress access logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_channel#egress_access_logs MediapackageChannel#egress_access_logs}
	EgressAccessLogs *MediapackageChannelEgressAccessLogs `field:"optional" json:"egressAccessLogs" yaml:"egressAccessLogs"`
	// An HTTP Live Streaming (HLS) ingest resource configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_channel#hls_ingest MediapackageChannel#hls_ingest}
	HlsIngest *MediapackageChannelHlsIngest `field:"optional" json:"hlsIngest" yaml:"hlsIngest"`
	// The configuration parameters for egress access logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_channel#ingress_access_logs MediapackageChannel#ingress_access_logs}
	IngressAccessLogs *MediapackageChannelIngressAccessLogs `field:"optional" json:"ingressAccessLogs" yaml:"ingressAccessLogs"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_channel#tags MediapackageChannel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

