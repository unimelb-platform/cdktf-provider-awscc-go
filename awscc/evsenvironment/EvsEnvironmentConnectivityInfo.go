package evsenvironment


type EvsEnvironmentConnectivityInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#private_route_server_peerings EvsEnvironment#private_route_server_peerings}.
	PrivateRouteServerPeerings *[]*string `field:"required" json:"privateRouteServerPeerings" yaml:"privateRouteServerPeerings"`
}

