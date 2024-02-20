package appstreamappblockbuilder


type AppstreamAppBlockBuilderAccessEndpoints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block_builder#endpoint_type AppstreamAppBlockBuilder#endpoint_type}.
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block_builder#vpce_id AppstreamAppBlockBuilder#vpce_id}.
	VpceId *string `field:"required" json:"vpceId" yaml:"vpceId"`
}

