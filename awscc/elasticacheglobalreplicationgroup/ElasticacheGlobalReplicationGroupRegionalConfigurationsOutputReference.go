package elasticacheglobalreplicationgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/elasticacheglobalreplicationgroup/internal"
)

type ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ReplicationGroupId() *string
	SetReplicationGroupId(val *string)
	ReplicationGroupIdInput() *string
	ReplicationGroupRegion() *string
	SetReplicationGroupRegion(val *string)
	ReplicationGroupRegionInput() *string
	ReshardingConfigurations() ElasticacheGlobalReplicationGroupRegionalConfigurationsReshardingConfigurationsList
	ReshardingConfigurationsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutReshardingConfigurations(value interface{})
	ResetReplicationGroupId()
	ResetReplicationGroupRegion()
	ResetReshardingConfigurations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference
type jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) ReplicationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) ReplicationGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) ReplicationGroupRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) ReplicationGroupRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) ReshardingConfigurations() ElasticacheGlobalReplicationGroupRegionalConfigurationsReshardingConfigurationsList {
	var returns ElasticacheGlobalReplicationGroupRegionalConfigurationsReshardingConfigurationsList
	_jsii_.Get(
		j,
		"reshardingConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) ReshardingConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reshardingConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference{}

	_jsii_.Create(
		"awscc.elasticacheGlobalReplicationGroup.ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference_Override(e ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.elasticacheGlobalReplicationGroup.ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference)SetReplicationGroupId(val *string) {
	if err := j.validateSetReplicationGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationGroupId",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference)SetReplicationGroupRegion(val *string) {
	if err := j.validateSetReplicationGroupRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationGroupRegion",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) PutReshardingConfigurations(value interface{}) {
	if err := e.validatePutReshardingConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putReshardingConfigurations",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) ResetReplicationGroupId() {
	_jsii_.InvokeVoid(
		e,
		"resetReplicationGroupId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) ResetReplicationGroupRegion() {
	_jsii_.InvokeVoid(
		e,
		"resetReplicationGroupRegion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) ResetReshardingConfigurations() {
	_jsii_.InvokeVoid(
		e,
		"resetReshardingConfigurations",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroupRegionalConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

