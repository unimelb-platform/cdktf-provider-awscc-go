package ec2routeserverpeer


type Ec2RouteServerPeerBgpOptions struct {
	// BGP ASN of the Route Server Peer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_route_server_peer#peer_asn Ec2RouteServerPeer#peer_asn}
	PeerAsn *float64 `field:"optional" json:"peerAsn" yaml:"peerAsn"`
	// BGP Liveness Detection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_route_server_peer#peer_liveness_detection Ec2RouteServerPeer#peer_liveness_detection}
	PeerLivenessDetection *string `field:"optional" json:"peerLivenessDetection" yaml:"peerLivenessDetection"`
}

