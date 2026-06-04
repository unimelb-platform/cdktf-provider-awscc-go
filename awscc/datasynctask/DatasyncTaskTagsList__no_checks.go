//go:build no_runtime_type_checking

package datasynctask

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatasyncTaskTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatasyncTaskTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatasyncTaskTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatasyncTaskTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatasyncTaskTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatasyncTaskTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatasyncTaskTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatasyncTaskTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

