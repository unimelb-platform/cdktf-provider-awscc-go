//go:build no_runtime_type_checking

package datazonedatasource

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatazoneDataSourceAssetFormsInputList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatazoneDataSourceAssetFormsInputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatazoneDataSourceAssetFormsInputList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatazoneDataSourceAssetFormsInputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatazoneDataSourceAssetFormsInputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatazoneDataSourceAssetFormsInputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatazoneDataSourceAssetFormsInputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

