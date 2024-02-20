//go:build !no_runtime_type_checking

package dataawsccdynamodbglobaltable

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (d *jsiiProxy_DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReference) validateSetInternalValueParameters(val *DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfiguration) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataAwsccDynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfigurationOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

