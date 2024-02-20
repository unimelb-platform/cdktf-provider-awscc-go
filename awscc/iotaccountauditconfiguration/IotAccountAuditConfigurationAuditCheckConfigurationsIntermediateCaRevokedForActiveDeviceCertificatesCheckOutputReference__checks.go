//go:build !no_runtime_type_checking

package iotaccountauditconfiguration

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateSetEnabledParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case cdktf.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *bool, cdktf.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheck:
		val := val.(*IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheck)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheck:
		val_ := val.(IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheck)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheck; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewIotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

