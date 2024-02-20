package dataawsccappflowflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccappflowflow/internal"
)

type DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomConnector() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesCustomConnectorOutputReference
	EventBridge() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesEventBridgeOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorProperties
	SetInternalValue(val *DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorProperties)
	LookoutMetrics() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesLookoutMetricsOutputReference
	Marketo() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesMarketoOutputReference
	Redshift() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesRedshiftOutputReference
	S3() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3OutputReference
	Salesforce() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSalesforceOutputReference
	SapoData() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataOutputReference
	Snowflake() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Upsolver() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesUpsolverOutputReference
	Zendesk() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskOutputReference
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference
type jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) CustomConnector() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesCustomConnectorOutputReference {
	var returns DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesCustomConnectorOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) EventBridge() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesEventBridgeOutputReference {
	var returns DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesEventBridgeOutputReference
	_jsii_.Get(
		j,
		"eventBridge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) InternalValue() *DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorProperties {
	var returns *DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) LookoutMetrics() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesLookoutMetricsOutputReference {
	var returns DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesLookoutMetricsOutputReference
	_jsii_.Get(
		j,
		"lookoutMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) Marketo() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesMarketoOutputReference {
	var returns DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesMarketoOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) Redshift() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesRedshiftOutputReference {
	var returns DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesRedshiftOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) S3() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3OutputReference {
	var returns DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) Salesforce() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSalesforceOutputReference {
	var returns DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSalesforceOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) SapoData() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataOutputReference {
	var returns DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) Snowflake() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeOutputReference {
	var returns DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) Upsolver() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesUpsolverOutputReference {
	var returns DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesUpsolverOutputReference
	_jsii_.Get(
		j,
		"upsolver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) Zendesk() DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskOutputReference {
	var returns DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}


func NewDataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccAppflowFlow.DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference_Override(d DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccAppflowFlow.DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference)SetInternalValue(val *DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

