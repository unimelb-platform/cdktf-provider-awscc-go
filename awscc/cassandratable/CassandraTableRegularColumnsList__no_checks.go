//go:build no_runtime_type_checking

package cassandratable

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CassandraTableRegularColumnsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CassandraTableRegularColumnsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CassandraTableRegularColumnsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CassandraTableRegularColumnsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CassandraTableRegularColumnsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CassandraTableRegularColumnsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CassandraTableRegularColumnsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCassandraTableRegularColumnsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

