//go:build no_runtime_type_checking

package cloudformationstackset

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudformationStackSetStackInstancesGroupListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

