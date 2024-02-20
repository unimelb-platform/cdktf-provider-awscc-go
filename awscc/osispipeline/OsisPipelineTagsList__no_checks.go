//go:build no_runtime_type_checking

package osispipeline

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OsisPipelineTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OsisPipelineTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OsisPipelineTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OsisPipelineTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OsisPipelineTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OsisPipelineTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOsisPipelineTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

