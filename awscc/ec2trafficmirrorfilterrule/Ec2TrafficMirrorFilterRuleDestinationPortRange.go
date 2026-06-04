package ec2trafficmirrorfilterrule


type Ec2TrafficMirrorFilterRuleDestinationPortRange struct {
	// The first port in the Traffic Mirror port range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_traffic_mirror_filter_rule#from_port Ec2TrafficMirrorFilterRule#from_port}
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// The last port in the Traffic Mirror port range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_traffic_mirror_filter_rule#to_port Ec2TrafficMirrorFilterRule#to_port}
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

