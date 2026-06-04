package mediapackagev2channel


type Mediapackagev2ChannelInputSwitchConfiguration struct {
	// <p>When true, AWS Elemental MediaPackage performs input switching based on the MQCS.
	//
	// Default is true. This setting is valid only when <code>InputType</code> is <code>CMAF</code>.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_channel#mqcs_input_switching Mediapackagev2Channel#mqcs_input_switching}
	MqcsInputSwitching interface{} `field:"optional" json:"mqcsInputSwitching" yaml:"mqcsInputSwitching"`
}

