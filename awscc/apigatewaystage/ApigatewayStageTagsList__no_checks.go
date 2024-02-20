//go:build no_runtime_type_checking

package apigatewaystage

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApigatewayStageTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApigatewayStageTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApigatewayStageTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApigatewayStageTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApigatewayStageTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApigatewayStageTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApigatewayStageTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

