//go:build no_runtime_type_checking

package ec2keypair

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2KeyPairTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_Ec2KeyPairTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2KeyPairTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2KeyPairTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2KeyPairTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2KeyPairTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2KeyPairTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2KeyPairTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

