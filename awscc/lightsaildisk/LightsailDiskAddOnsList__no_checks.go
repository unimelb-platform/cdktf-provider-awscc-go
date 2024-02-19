//go:build no_runtime_type_checking

package lightsaildisk

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LightsailDiskAddOnsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LightsailDiskAddOnsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LightsailDiskAddOnsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LightsailDiskAddOnsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LightsailDiskAddOnsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LightsailDiskAddOnsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLightsailDiskAddOnsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

