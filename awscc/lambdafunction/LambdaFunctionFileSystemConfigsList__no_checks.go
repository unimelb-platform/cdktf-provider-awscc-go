//go:build no_runtime_type_checking

package lambdafunction

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaFunctionFileSystemConfigsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LambdaFunctionFileSystemConfigsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLambdaFunctionFileSystemConfigsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

