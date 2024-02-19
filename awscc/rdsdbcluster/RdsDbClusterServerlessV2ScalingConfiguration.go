package rdsdbcluster


type RdsDbClusterServerlessV2ScalingConfiguration struct {
	// The maximum number of Aurora capacity units (ACUs) for a DB instance in an Aurora Serverless v2 cluster.
	//
	// You can specify ACU values in half-step increments, such as 40, 40.5, 41, and so on. The largest value that you can use is 128.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#max_capacity RdsDbCluster#max_capacity}
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum number of Aurora capacity units (ACUs) for a DB instance in an Aurora Serverless v2 cluster.
	//
	// You can specify ACU values in half-step increments, such as 8, 8.5, 9, and so on. The smallest value that you can use is 0.5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#min_capacity RdsDbCluster#min_capacity}
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}

