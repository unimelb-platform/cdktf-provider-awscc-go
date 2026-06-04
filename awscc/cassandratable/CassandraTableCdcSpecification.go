package cassandratable


type CassandraTableCdcSpecification struct {
	// Indicates whether CDC is enabled or disabled for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cassandra_table#status CassandraTable#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Specifies what data should be captured in the change data stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cassandra_table#view_type CassandraTable#view_type}
	ViewType *string `field:"optional" json:"viewType" yaml:"viewType"`
}

