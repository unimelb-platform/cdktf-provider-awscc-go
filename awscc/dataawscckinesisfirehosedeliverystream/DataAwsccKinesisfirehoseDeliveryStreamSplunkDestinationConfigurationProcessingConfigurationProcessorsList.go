package dataawscckinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscckinesisfirehosedeliverystream/internal"
)

type DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList
type jsiiProxy_DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList {
	_init_.Initialize()

	if err := validateNewDataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList{}

	_jsii_.Create(
		"awscc.dataAwsccKinesisfirehoseDeliveryStream.DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList_Override(d DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccKinesisfirehoseDeliveryStream.DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList) Get(index *float64) DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationProcessorsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

