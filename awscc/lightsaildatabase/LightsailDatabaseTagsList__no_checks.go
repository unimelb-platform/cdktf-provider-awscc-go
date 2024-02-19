//go:build no_runtime_type_checking

package lightsaildatabase

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LightsailDatabaseTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LightsailDatabaseTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LightsailDatabaseTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LightsailDatabaseTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LightsailDatabaseTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LightsailDatabaseTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLightsailDatabaseTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

