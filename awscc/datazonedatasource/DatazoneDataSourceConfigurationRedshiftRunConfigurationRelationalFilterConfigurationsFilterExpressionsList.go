package datazonedatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/datazonedatasource/internal"
)

type DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList
type jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList {
	_init_.Initialize()

	if err := validateNewDatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList{}

	_jsii_.Create(
		"awscc.datazoneDataSource.DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList_Override(d DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.datazoneDataSource.DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) Get(index *float64) DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

