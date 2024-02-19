//go:build no_runtime_type_checking

package cassandratable

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CassandraTableTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CassandraTableTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CassandraTableTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CassandraTableTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CassandraTableTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CassandraTableTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CassandraTableTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCassandraTableTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

