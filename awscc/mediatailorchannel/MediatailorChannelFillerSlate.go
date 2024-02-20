package mediatailorchannel


type MediatailorChannelFillerSlate struct {
	// <p>The name of the source location where the slate VOD source is stored.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediatailor_channel#source_location_name MediatailorChannel#source_location_name}
	SourceLocationName *string `field:"optional" json:"sourceLocationName" yaml:"sourceLocationName"`
	// <p>The slate VOD source name.
	//
	// The VOD source must already exist in a source location before it can be used for slate.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediatailor_channel#vod_source_name MediatailorChannel#vod_source_name}
	VodSourceName *string `field:"optional" json:"vodSourceName" yaml:"vodSourceName"`
}

