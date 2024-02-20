//go:build no_runtime_type_checking

package ec2ipam

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2IpamOperatingRegionsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2IpamOperatingRegionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2IpamOperatingRegionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2IpamOperatingRegionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2IpamOperatingRegionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2IpamOperatingRegionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2IpamOperatingRegionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

