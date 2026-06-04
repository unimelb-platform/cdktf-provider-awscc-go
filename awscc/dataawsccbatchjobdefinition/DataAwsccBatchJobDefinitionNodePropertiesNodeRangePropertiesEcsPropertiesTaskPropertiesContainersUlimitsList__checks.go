//go:build !no_runtime_type_checking

package dataawsccbatchjobdefinition

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesContainersUlimitsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesContainersUlimitsList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesContainersUlimitsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesContainersUlimitsList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesContainersUlimitsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesContainersUlimitsList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesContainersUlimitsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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

