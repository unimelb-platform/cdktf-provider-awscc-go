//go:build no_runtime_type_checking

package datasynctask

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatasyncTaskIncludesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatasyncTaskIncludesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatasyncTaskIncludesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatasyncTaskIncludesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatasyncTaskIncludesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatasyncTaskIncludesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatasyncTaskIncludesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatasyncTaskIncludesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

