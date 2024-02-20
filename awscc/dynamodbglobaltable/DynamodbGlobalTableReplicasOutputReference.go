package dynamodbglobaltable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dynamodbglobaltable/internal"
)

type DynamodbGlobalTableReplicasOutputReference interface {
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
	ContributorInsightsSpecification() DynamodbGlobalTableReplicasContributorInsightsSpecificationOutputReference
	ContributorInsightsSpecificationInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeletionProtectionEnabled() interface{}
	SetDeletionProtectionEnabled(val interface{})
	DeletionProtectionEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	GlobalSecondaryIndexes() DynamodbGlobalTableReplicasGlobalSecondaryIndexesList
	GlobalSecondaryIndexesInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KinesisStreamSpecification() DynamodbGlobalTableReplicasKinesisStreamSpecificationOutputReference
	KinesisStreamSpecificationInput() interface{}
	PointInTimeRecoverySpecification() DynamodbGlobalTableReplicasPointInTimeRecoverySpecificationOutputReference
	PointInTimeRecoverySpecificationInput() interface{}
	ReadProvisionedThroughputSettings() DynamodbGlobalTableReplicasReadProvisionedThroughputSettingsOutputReference
	ReadProvisionedThroughputSettingsInput() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SseSpecification() DynamodbGlobalTableReplicasSseSpecificationOutputReference
	SseSpecificationInput() interface{}
	TableClass() *string
	SetTableClass(val *string)
	TableClassInput() *string
	Tags() DynamodbGlobalTableReplicasTagsList
	TagsInput() interface{}
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
	PutContributorInsightsSpecification(value *DynamodbGlobalTableReplicasContributorInsightsSpecification)
	PutGlobalSecondaryIndexes(value interface{})
	PutKinesisStreamSpecification(value *DynamodbGlobalTableReplicasKinesisStreamSpecification)
	PutPointInTimeRecoverySpecification(value *DynamodbGlobalTableReplicasPointInTimeRecoverySpecification)
	PutReadProvisionedThroughputSettings(value *DynamodbGlobalTableReplicasReadProvisionedThroughputSettings)
	PutSseSpecification(value *DynamodbGlobalTableReplicasSseSpecification)
	PutTags(value interface{})
	ResetContributorInsightsSpecification()
	ResetDeletionProtectionEnabled()
	ResetGlobalSecondaryIndexes()
	ResetKinesisStreamSpecification()
	ResetPointInTimeRecoverySpecification()
	ResetReadProvisionedThroughputSettings()
	ResetSseSpecification()
	ResetTableClass()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DynamodbGlobalTableReplicasOutputReference
type jsiiProxy_DynamodbGlobalTableReplicasOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ContributorInsightsSpecification() DynamodbGlobalTableReplicasContributorInsightsSpecificationOutputReference {
	var returns DynamodbGlobalTableReplicasContributorInsightsSpecificationOutputReference
	_jsii_.Get(
		j,
		"contributorInsightsSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ContributorInsightsSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contributorInsightsSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) DeletionProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) GlobalSecondaryIndexes() DynamodbGlobalTableReplicasGlobalSecondaryIndexesList {
	var returns DynamodbGlobalTableReplicasGlobalSecondaryIndexesList
	_jsii_.Get(
		j,
		"globalSecondaryIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) GlobalSecondaryIndexesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalSecondaryIndexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) KinesisStreamSpecification() DynamodbGlobalTableReplicasKinesisStreamSpecificationOutputReference {
	var returns DynamodbGlobalTableReplicasKinesisStreamSpecificationOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) KinesisStreamSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisStreamSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) PointInTimeRecoverySpecification() DynamodbGlobalTableReplicasPointInTimeRecoverySpecificationOutputReference {
	var returns DynamodbGlobalTableReplicasPointInTimeRecoverySpecificationOutputReference
	_jsii_.Get(
		j,
		"pointInTimeRecoverySpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) PointInTimeRecoverySpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pointInTimeRecoverySpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ReadProvisionedThroughputSettings() DynamodbGlobalTableReplicasReadProvisionedThroughputSettingsOutputReference {
	var returns DynamodbGlobalTableReplicasReadProvisionedThroughputSettingsOutputReference
	_jsii_.Get(
		j,
		"readProvisionedThroughputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ReadProvisionedThroughputSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readProvisionedThroughputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) SseSpecification() DynamodbGlobalTableReplicasSseSpecificationOutputReference {
	var returns DynamodbGlobalTableReplicasSseSpecificationOutputReference
	_jsii_.Get(
		j,
		"sseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) SseSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) TableClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) TableClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) Tags() DynamodbGlobalTableReplicasTagsList {
	var returns DynamodbGlobalTableReplicasTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDynamodbGlobalTableReplicasOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DynamodbGlobalTableReplicasOutputReference {
	_init_.Initialize()

	if err := validateNewDynamodbGlobalTableReplicasOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamodbGlobalTableReplicasOutputReference{}

	_jsii_.Create(
		"awscc.dynamodbGlobalTable.DynamodbGlobalTableReplicasOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDynamodbGlobalTableReplicasOutputReference_Override(d DynamodbGlobalTableReplicasOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dynamodbGlobalTable.DynamodbGlobalTableReplicasOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference)SetTableClass(val *string) {
	if err := j.validateSetTableClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableClass",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) PutContributorInsightsSpecification(value *DynamodbGlobalTableReplicasContributorInsightsSpecification) {
	if err := d.validatePutContributorInsightsSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContributorInsightsSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) PutGlobalSecondaryIndexes(value interface{}) {
	if err := d.validatePutGlobalSecondaryIndexesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGlobalSecondaryIndexes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) PutKinesisStreamSpecification(value *DynamodbGlobalTableReplicasKinesisStreamSpecification) {
	if err := d.validatePutKinesisStreamSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKinesisStreamSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) PutPointInTimeRecoverySpecification(value *DynamodbGlobalTableReplicasPointInTimeRecoverySpecification) {
	if err := d.validatePutPointInTimeRecoverySpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPointInTimeRecoverySpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) PutReadProvisionedThroughputSettings(value *DynamodbGlobalTableReplicasReadProvisionedThroughputSettings) {
	if err := d.validatePutReadProvisionedThroughputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReadProvisionedThroughputSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) PutSseSpecification(value *DynamodbGlobalTableReplicasSseSpecification) {
	if err := d.validatePutSseSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSseSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) PutTags(value interface{}) {
	if err := d.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ResetContributorInsightsSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetContributorInsightsSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ResetDeletionProtectionEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDeletionProtectionEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ResetGlobalSecondaryIndexes() {
	_jsii_.InvokeVoid(
		d,
		"resetGlobalSecondaryIndexes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ResetKinesisStreamSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetKinesisStreamSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ResetPointInTimeRecoverySpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetPointInTimeRecoverySpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ResetReadProvisionedThroughputSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetReadProvisionedThroughputSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ResetSseSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetSseSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ResetTableClass() {
	_jsii_.InvokeVoid(
		d,
		"resetTableClass",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DynamodbGlobalTableReplicasOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

