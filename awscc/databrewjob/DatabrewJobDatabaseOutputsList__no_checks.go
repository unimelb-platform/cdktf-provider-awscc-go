//go:build no_runtime_type_checking

package databrewjob

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabrewJobDatabaseOutputsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabrewJobDatabaseOutputsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabrewJobDatabaseOutputsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabrewJobDatabaseOutputsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabrewJobDatabaseOutputsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabrewJobDatabaseOutputsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabrewJobDatabaseOutputsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

