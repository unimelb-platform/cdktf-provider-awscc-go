package cassandratable


type CassandraTableClusteringKeyColumnsColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#column_name CassandraTable#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#column_type CassandraTable#column_type}.
	ColumnType *string `field:"required" json:"columnType" yaml:"columnType"`
}

