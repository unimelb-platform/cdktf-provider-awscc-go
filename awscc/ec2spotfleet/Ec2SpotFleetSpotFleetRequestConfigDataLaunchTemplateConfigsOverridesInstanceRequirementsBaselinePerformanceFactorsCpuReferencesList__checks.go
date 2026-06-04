//go:build !no_runtime_type_checking

package ec2spotfleet

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferencesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferencesList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferencesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferencesList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferences:
		val := val.(*[]*Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferences)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferences:
		val_ := val.([]*Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferences)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferences; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferencesList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferencesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferencesList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewEc2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferencesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

