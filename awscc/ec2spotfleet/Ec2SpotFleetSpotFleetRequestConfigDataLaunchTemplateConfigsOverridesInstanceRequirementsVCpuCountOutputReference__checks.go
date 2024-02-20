//go:build !no_runtime_type_checking

package ec2spotfleet

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCount:
		val := val.(*Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCount)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCount:
		val_ := val.(Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCount)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCount; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateSetMaxParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateSetMinParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewEc2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsVCpuCountOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

