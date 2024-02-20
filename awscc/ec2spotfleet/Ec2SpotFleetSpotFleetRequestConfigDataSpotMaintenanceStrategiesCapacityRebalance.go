package ec2spotfleet


type Ec2SpotFleetSpotFleetRequestConfigDataSpotMaintenanceStrategiesCapacityRebalance struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_spot_fleet#replacement_strategy Ec2SpotFleet#replacement_strategy}.
	ReplacementStrategy *string `field:"optional" json:"replacementStrategy" yaml:"replacementStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_spot_fleet#termination_delay Ec2SpotFleet#termination_delay}.
	TerminationDelay *float64 `field:"optional" json:"terminationDelay" yaml:"terminationDelay"`
}

