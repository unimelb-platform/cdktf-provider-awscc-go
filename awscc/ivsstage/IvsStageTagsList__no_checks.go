//go:build no_runtime_type_checking

package ivsstage

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IvsStageTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IvsStageTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IvsStageTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IvsStageTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IvsStageTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IvsStageTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IvsStageTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIvsStageTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

