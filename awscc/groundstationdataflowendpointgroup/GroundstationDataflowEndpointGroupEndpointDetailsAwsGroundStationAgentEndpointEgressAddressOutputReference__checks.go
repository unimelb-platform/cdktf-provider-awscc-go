//go:build !no_runtime_type_checking

package groundstationdataflowendpointgroup

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validatePutSocketAddressParameters(value *GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressSocketAddress) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddress:
		val := val.(*GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddress)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddress:
		val_ := val.(GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddress)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddress; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateSetMtuParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewGroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

