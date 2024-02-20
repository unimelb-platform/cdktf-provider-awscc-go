//go:build no_runtime_type_checking

package datasyncstoragesystem

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatasyncStorageSystemTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatasyncStorageSystemTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatasyncStorageSystemTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatasyncStorageSystemTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatasyncStorageSystemTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatasyncStorageSystemTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatasyncStorageSystemTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

