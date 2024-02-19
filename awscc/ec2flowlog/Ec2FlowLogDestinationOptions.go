package ec2flowlog


type Ec2FlowLogDestinationOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_flow_log#file_format Ec2FlowLog#file_format}.
	FileFormat *string `field:"required" json:"fileFormat" yaml:"fileFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_flow_log#hive_compatible_partitions Ec2FlowLog#hive_compatible_partitions}.
	HiveCompatiblePartitions interface{} `field:"required" json:"hiveCompatiblePartitions" yaml:"hiveCompatiblePartitions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_flow_log#per_hour_partition Ec2FlowLog#per_hour_partition}.
	PerHourPartition interface{} `field:"required" json:"perHourPartition" yaml:"perHourPartition"`
}

