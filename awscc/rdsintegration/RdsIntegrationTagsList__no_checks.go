//go:build no_runtime_type_checking

package rdsintegration

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsIntegrationTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RdsIntegrationTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RdsIntegrationTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RdsIntegrationTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RdsIntegrationTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RdsIntegrationTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RdsIntegrationTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRdsIntegrationTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

