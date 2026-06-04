//go:build !no_runtime_type_checking

package ec2ec2fleet

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validatePutCpuParameters(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpu) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactors:
		val := val.(*Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactors)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactors:
		val_ := val.(Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactors)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactors; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

