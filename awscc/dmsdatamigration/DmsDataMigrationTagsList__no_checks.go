//go:build no_runtime_type_checking

package dmsdatamigration

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DmsDataMigrationTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DmsDataMigrationTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DmsDataMigrationTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DmsDataMigrationTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DmsDataMigrationTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DmsDataMigrationTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DmsDataMigrationTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDmsDataMigrationTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

