//go:build no_runtime_type_checking

package cassandratype

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CassandraTypeFieldsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CassandraTypeFieldsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CassandraTypeFieldsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CassandraTypeFieldsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CassandraTypeFieldsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CassandraTypeFieldsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CassandraTypeFieldsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCassandraTypeFieldsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

