//go:build no_runtime_type_checking

package apsscraper

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApsScraperTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApsScraperTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApsScraperTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApsScraperTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApsScraperTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApsScraperTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApsScraperTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApsScraperTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

