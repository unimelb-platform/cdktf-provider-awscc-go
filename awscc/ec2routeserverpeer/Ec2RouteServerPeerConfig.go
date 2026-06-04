package ec2routeserverpeer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2RouteServerPeerConfig struct {
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
	// BGP Options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_route_server_peer#bgp_options Ec2RouteServerPeer#bgp_options}
	BgpOptions *Ec2RouteServerPeerBgpOptions `field:"required" json:"bgpOptions" yaml:"bgpOptions"`
	// IP address of the Route Server Peer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_route_server_peer#peer_address Ec2RouteServerPeer#peer_address}
	PeerAddress *string `field:"required" json:"peerAddress" yaml:"peerAddress"`
	// Route Server Endpoint ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_route_server_peer#route_server_endpoint_id Ec2RouteServerPeer#route_server_endpoint_id}
	RouteServerEndpointId *string `field:"required" json:"routeServerEndpointId" yaml:"routeServerEndpointId"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_route_server_peer#tags Ec2RouteServerPeer#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

