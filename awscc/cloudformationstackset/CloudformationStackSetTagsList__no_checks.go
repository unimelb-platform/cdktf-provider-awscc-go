//go:build no_runtime_type_checking

package cloudformationstackset

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudformationStackSetTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudformationStackSetTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudformationStackSetTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudformationStackSetTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudformationStackSetTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudformationStackSetTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudformationStackSetTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

