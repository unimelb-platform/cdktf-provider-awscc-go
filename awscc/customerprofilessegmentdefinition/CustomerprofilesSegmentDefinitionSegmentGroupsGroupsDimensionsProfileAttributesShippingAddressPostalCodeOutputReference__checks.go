//go:build !no_runtime_type_checking

package customerprofilessegmentdefinition

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateSetDimensionTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCode:
		val := val.(*CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCode)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCode:
		val_ := val.(CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCode)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCode; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference) validateSetValuesParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

