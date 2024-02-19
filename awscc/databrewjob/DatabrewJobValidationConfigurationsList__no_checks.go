//go:build no_runtime_type_checking

package databrewjob

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabrewJobValidationConfigurationsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabrewJobValidationConfigurationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabrewJobValidationConfigurationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabrewJobValidationConfigurationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabrewJobValidationConfigurationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabrewJobValidationConfigurationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabrewJobValidationConfigurationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

