//go:build no_runtime_type_checking

package locationtracker

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LocationTrackerTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LocationTrackerTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LocationTrackerTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LocationTrackerTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LocationTrackerTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LocationTrackerTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LocationTrackerTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLocationTrackerTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

