//go:build no_runtime_type_checking

package cassandratable

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CassandraTableReplicaSpecificationsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCassandraTableReplicaSpecificationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

