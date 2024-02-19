package mediapackageoriginendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediapackageOriginEndpointConfig struct {
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
	// The ID of the Channel the OriginEndpoint is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#channel_id MediapackageOriginEndpoint#channel_id}
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// The ID of the OriginEndpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#id MediapackageOriginEndpoint#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// CDN Authorization credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#authorization MediapackageOriginEndpoint#authorization}
	Authorization *MediapackageOriginEndpointAuthorization `field:"optional" json:"authorization" yaml:"authorization"`
	// A Common Media Application Format (CMAF) packaging configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#cmaf_package MediapackageOriginEndpoint#cmaf_package}
	CmafPackage *MediapackageOriginEndpointCmafPackage `field:"optional" json:"cmafPackage" yaml:"cmafPackage"`
	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#dash_package MediapackageOriginEndpoint#dash_package}
	DashPackage *MediapackageOriginEndpointDashPackage `field:"optional" json:"dashPackage" yaml:"dashPackage"`
	// A short text description of the OriginEndpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#description MediapackageOriginEndpoint#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An HTTP Live Streaming (HLS) packaging configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#hls_package MediapackageOriginEndpoint#hls_package}
	HlsPackage *MediapackageOriginEndpointHlsPackage `field:"optional" json:"hlsPackage" yaml:"hlsPackage"`
	// A short string appended to the end of the OriginEndpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#manifest_name MediapackageOriginEndpoint#manifest_name}
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// A Microsoft Smooth Streaming (MSS) packaging configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#mss_package MediapackageOriginEndpoint#mss_package}
	MssPackage *MediapackageOriginEndpointMssPackage `field:"optional" json:"mssPackage" yaml:"mssPackage"`
	// Control whether origination of video is allowed for this OriginEndpoint.
	//
	// If set to ALLOW, the OriginEndpoint may by requested, pursuant to any other form of access control. If set to DENY, the OriginEndpoint may not be requested. This can be helpful for Live to VOD harvesting, or for temporarily disabling origination
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#origination MediapackageOriginEndpoint#origination}
	Origination *string `field:"optional" json:"origination" yaml:"origination"`
	// Maximum duration (seconds) of content to retain for startover playback.
	//
	// If not specified, startover playback will be disabled for the OriginEndpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#startover_window_seconds MediapackageOriginEndpoint#startover_window_seconds}
	StartoverWindowSeconds *float64 `field:"optional" json:"startoverWindowSeconds" yaml:"startoverWindowSeconds"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#tags MediapackageOriginEndpoint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Amount of delay (seconds) to enforce on the playback of live content.
	//
	// If not specified, there will be no time delay in effect for the OriginEndpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#time_delay_seconds MediapackageOriginEndpoint#time_delay_seconds}
	TimeDelaySeconds *float64 `field:"optional" json:"timeDelaySeconds" yaml:"timeDelaySeconds"`
	// A list of source IP CIDR blocks that will be allowed to access the OriginEndpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#whitelist MediapackageOriginEndpoint#whitelist}
	Whitelist *[]*string `field:"optional" json:"whitelist" yaml:"whitelist"`
}

