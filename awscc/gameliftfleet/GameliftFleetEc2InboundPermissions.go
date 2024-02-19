package gameliftfleet


type GameliftFleetEc2InboundPermissions struct {
	// A starting value for a range of allowed port numbers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_fleet#from_port GameliftFleet#from_port}
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// A range of allowed IP addresses.
	//
	// This value must be expressed in CIDR notation. Example: "000.000.000.000/[subnet mask]" or optionally the shortened version "0.0.0.0/[subnet mask]".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_fleet#ip_range GameliftFleet#ip_range}
	IpRange *string `field:"required" json:"ipRange" yaml:"ipRange"`
	// The network communication protocol used by the fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_fleet#protocol GameliftFleet#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// An ending value for a range of allowed port numbers.
	//
	// Port numbers are end-inclusive. This value must be higher than FromPort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_fleet#to_port GameliftFleet#to_port}
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

