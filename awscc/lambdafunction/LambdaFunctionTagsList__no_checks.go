//go:build no_runtime_type_checking

package lambdafunction

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaFunctionTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LambdaFunctionTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LambdaFunctionTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LambdaFunctionTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunctionTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LambdaFunctionTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLambdaFunctionTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

