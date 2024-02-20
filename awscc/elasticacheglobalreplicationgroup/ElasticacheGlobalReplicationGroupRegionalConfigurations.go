package elasticacheglobalreplicationgroup


type ElasticacheGlobalReplicationGroupRegionalConfigurations struct {
	// The replication group id of the Global Datastore member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_global_replication_group#replication_group_id ElasticacheGlobalReplicationGroup#replication_group_id}
	ReplicationGroupId *string `field:"optional" json:"replicationGroupId" yaml:"replicationGroupId"`
	// The AWS region of the Global Datastore member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_global_replication_group#replication_group_region ElasticacheGlobalReplicationGroup#replication_group_region}
	ReplicationGroupRegion *string `field:"optional" json:"replicationGroupRegion" yaml:"replicationGroupRegion"`
	// A list of PreferredAvailabilityZones objects that specifies the configuration of a node group in the resharded cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_global_replication_group#resharding_configurations ElasticacheGlobalReplicationGroup#resharding_configurations}
	ReshardingConfigurations interface{} `field:"optional" json:"reshardingConfigurations" yaml:"reshardingConfigurations"`
}

