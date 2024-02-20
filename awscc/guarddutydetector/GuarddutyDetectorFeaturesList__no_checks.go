//go:build no_runtime_type_checking

package guarddutydetector

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GuarddutyDetectorFeaturesList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GuarddutyDetectorFeaturesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GuarddutyDetectorFeaturesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GuarddutyDetectorFeaturesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GuarddutyDetectorFeaturesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GuarddutyDetectorFeaturesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGuarddutyDetectorFeaturesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

