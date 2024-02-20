package dynamodbglobaltable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dynamodbglobaltable/internal"
)

type DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference interface {
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
	ContributorInsightsSpecification() DynamodbGlobalTableReplicasGlobalSecondaryIndexesContributorInsightsSpecificationOutputReference
	ContributorInsightsSpecificationInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IndexName() *string
	SetIndexName(val *string)
	IndexNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ReadProvisionedThroughputSettings() DynamodbGlobalTableReplicasGlobalSecondaryIndexesReadProvisionedThroughputSettingsOutputReference
	ReadProvisionedThroughputSettingsInput() interface{}
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
	PutContributorInsightsSpecification(value *DynamodbGlobalTableReplicasGlobalSecondaryIndexesContributorInsightsSpecification)
	PutReadProvisionedThroughputSettings(value *DynamodbGlobalTableReplicasGlobalSecondaryIndexesReadProvisionedThroughputSettings)
	ResetContributorInsightsSpecification()
	ResetReadProvisionedThroughputSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference
type jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) ContributorInsightsSpecification() DynamodbGlobalTableReplicasGlobalSecondaryIndexesContributorInsightsSpecificationOutputReference {
	var returns DynamodbGlobalTableReplicasGlobalSecondaryIndexesContributorInsightsSpecificationOutputReference
	_jsii_.Get(
		j,
		"contributorInsightsSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) ContributorInsightsSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contributorInsightsSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) IndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) IndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) ReadProvisionedThroughputSettings() DynamodbGlobalTableReplicasGlobalSecondaryIndexesReadProvisionedThroughputSettingsOutputReference {
	var returns DynamodbGlobalTableReplicasGlobalSecondaryIndexesReadProvisionedThroughputSettingsOutputReference
	_jsii_.Get(
		j,
		"readProvisionedThroughputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) ReadProvisionedThroughputSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readProvisionedThroughputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference {
	_init_.Initialize()

	if err := validateNewDynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference{}

	_jsii_.Create(
		"awscc.dynamodbGlobalTable.DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference_Override(d DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dynamodbGlobalTable.DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference)SetIndexName(val *string) {
	if err := j.validateSetIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexName",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) PutContributorInsightsSpecification(value *DynamodbGlobalTableReplicasGlobalSecondaryIndexesContributorInsightsSpecification) {
	if err := d.validatePutContributorInsightsSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContributorInsightsSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) PutReadProvisionedThroughputSettings(value *DynamodbGlobalTableReplicasGlobalSecondaryIndexesReadProvisionedThroughputSettings) {
	if err := d.validatePutReadProvisionedThroughputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReadProvisionedThroughputSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) ResetContributorInsightsSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetContributorInsightsSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) ResetReadProvisionedThroughputSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetReadProvisionedThroughputSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasGlobalSecondaryIndexesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

