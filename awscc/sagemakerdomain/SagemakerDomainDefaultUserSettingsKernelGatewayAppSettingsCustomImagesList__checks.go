//go:build !no_runtime_type_checking

package sagemakerdomain

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImagesList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImagesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImagesList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImages:
		val := val.(*[]*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImages)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImages:
		val_ := val.([]*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImages)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImages; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImagesList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImagesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImagesList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewSagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImagesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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

