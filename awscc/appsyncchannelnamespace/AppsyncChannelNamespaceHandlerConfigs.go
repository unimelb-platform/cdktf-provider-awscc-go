package appsyncchannelnamespace


type AppsyncChannelNamespaceHandlerConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_channel_namespace#on_publish AppsyncChannelNamespace#on_publish}.
	OnPublish *AppsyncChannelNamespaceHandlerConfigsOnPublish `field:"optional" json:"onPublish" yaml:"onPublish"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_channel_namespace#on_subscribe AppsyncChannelNamespace#on_subscribe}.
	OnSubscribe *AppsyncChannelNamespaceHandlerConfigsOnSubscribe `field:"optional" json:"onSubscribe" yaml:"onSubscribe"`
}

