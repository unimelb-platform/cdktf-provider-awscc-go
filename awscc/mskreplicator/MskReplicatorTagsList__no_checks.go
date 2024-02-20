//go:build no_runtime_type_checking

package mskreplicator

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MskReplicatorTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MskReplicatorTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MskReplicatorTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MskReplicatorTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MskReplicatorTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MskReplicatorTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMskReplicatorTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

