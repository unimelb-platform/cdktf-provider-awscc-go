//go:build no_runtime_type_checking

package pcsqueue

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PcsQueueErrorInfoList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PcsQueueErrorInfoList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PcsQueueErrorInfoList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PcsQueueErrorInfoList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PcsQueueErrorInfoList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PcsQueueErrorInfoList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPcsQueueErrorInfoListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

