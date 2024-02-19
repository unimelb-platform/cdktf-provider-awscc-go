package apigatewayrestapi


type ApigatewayRestApiEndpointConfiguration struct {
	// A list of endpoint types of an API (RestApi) or its custom domain name (DomainName).
	//
	// For an edge-optimized API and its custom domain name, the endpoint type is ``"EDGE"``. For a regional API and its custom domain name, the endpoint type is ``REGIONAL``. For a private API, the endpoint type is ``PRIVATE``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#types ApigatewayRestApi#types}
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
	// A list of VpcEndpointIds of an API (RestApi) against which to create Route53 ALIASes.
	//
	// It is only supported for ``PRIVATE`` endpoint type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#vpc_endpoint_ids ApigatewayRestApi#vpc_endpoint_ids}
	VpcEndpointIds *[]*string `field:"optional" json:"vpcEndpointIds" yaml:"vpcEndpointIds"`
}

