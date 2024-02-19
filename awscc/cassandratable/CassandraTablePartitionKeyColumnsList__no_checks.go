//go:build no_runtime_type_checking

package cassandratable

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CassandraTablePartitionKeyColumnsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CassandraTablePartitionKeyColumnsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CassandraTablePartitionKeyColumnsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CassandraTablePartitionKeyColumnsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CassandraTablePartitionKeyColumnsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CassandraTablePartitionKeyColumnsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCassandraTablePartitionKeyColumnsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

