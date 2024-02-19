//go:build no_runtime_type_checking

package ec2ipampool

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2IpamPoolProvisionedCidrsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2IpamPoolProvisionedCidrsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2IpamPoolProvisionedCidrsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2IpamPoolProvisionedCidrsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2IpamPoolProvisionedCidrsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2IpamPoolProvisionedCidrsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2IpamPoolProvisionedCidrsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

