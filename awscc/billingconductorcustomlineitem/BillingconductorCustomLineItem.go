package billingconductorcustomlineitem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/billingconductorcustomlineitem/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item awscc_billingconductor_custom_line_item}.
type BillingconductorCustomLineItem interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	Arn() *string
	AssociationSize() *float64
	BillingGroupArn() *string
	SetBillingGroupArn(val *string)
	BillingGroupArnInput() *string
	BillingPeriodRange() BillingconductorCustomLineItemBillingPeriodRangeOutputReference
	BillingPeriodRangeInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTime() *float64
	CurrencyCode() *string
	CustomLineItemChargeDetails() BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference
	CustomLineItemChargeDetailsInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	LastModifiedTime() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProductCode() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Tags() BillingconductorCustomLineItemTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutBillingPeriodRange(value *BillingconductorCustomLineItemBillingPeriodRange)
	PutCustomLineItemChargeDetails(value *BillingconductorCustomLineItemCustomLineItemChargeDetails)
	PutTags(value interface{})
	ResetAccountId()
	ResetBillingPeriodRange()
	ResetCustomLineItemChargeDetails()
	ResetDescription()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for BillingconductorCustomLineItem
type jsiiProxy_BillingconductorCustomLineItem struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BillingconductorCustomLineItem) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) AssociationSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"associationSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) BillingGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) BillingGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) BillingPeriodRange() BillingconductorCustomLineItemBillingPeriodRangeOutputReference {
	var returns BillingconductorCustomLineItemBillingPeriodRangeOutputReference
	_jsii_.Get(
		j,
		"billingPeriodRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) BillingPeriodRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"billingPeriodRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) CreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) CurrencyCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currencyCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) CustomLineItemChargeDetails() BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference {
	var returns BillingconductorCustomLineItemCustomLineItemChargeDetailsOutputReference
	_jsii_.Get(
		j,
		"customLineItemChargeDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) CustomLineItemChargeDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customLineItemChargeDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) LastModifiedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) ProductCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) Tags() BillingconductorCustomLineItemTagsList {
	var returns BillingconductorCustomLineItemTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorCustomLineItem) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item awscc_billingconductor_custom_line_item} Resource.
func NewBillingconductorCustomLineItem(scope constructs.Construct, id *string, config *BillingconductorCustomLineItemConfig) BillingconductorCustomLineItem {
	_init_.Initialize()

	if err := validateNewBillingconductorCustomLineItemParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BillingconductorCustomLineItem{}

	_jsii_.Create(
		"awscc.billingconductorCustomLineItem.BillingconductorCustomLineItem",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item awscc_billingconductor_custom_line_item} Resource.
func NewBillingconductorCustomLineItem_Override(b BillingconductorCustomLineItem, scope constructs.Construct, id *string, config *BillingconductorCustomLineItemConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.billingconductorCustomLineItem.BillingconductorCustomLineItem",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItem)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItem)SetBillingGroupArn(val *string) {
	if err := j.validateSetBillingGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingGroupArn",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItem)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItem)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItem)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItem)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItem)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItem)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItem)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItem)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BillingconductorCustomLineItem)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a BillingconductorCustomLineItem resource upon running "cdktf plan <stack-name>".
func BillingconductorCustomLineItem_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBillingconductorCustomLineItem_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.billingconductorCustomLineItem.BillingconductorCustomLineItem",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func BillingconductorCustomLineItem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBillingconductorCustomLineItem_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.billingconductorCustomLineItem.BillingconductorCustomLineItem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BillingconductorCustomLineItem_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBillingconductorCustomLineItem_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.billingconductorCustomLineItem.BillingconductorCustomLineItem",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BillingconductorCustomLineItem_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBillingconductorCustomLineItem_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.billingconductorCustomLineItem.BillingconductorCustomLineItem",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BillingconductorCustomLineItem_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.billingconductorCustomLineItem.BillingconductorCustomLineItem",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItem) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItem) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItem) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItem) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItem) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItem) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItem) PutBillingPeriodRange(value *BillingconductorCustomLineItemBillingPeriodRange) {
	if err := b.validatePutBillingPeriodRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putBillingPeriodRange",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItem) PutCustomLineItemChargeDetails(value *BillingconductorCustomLineItemCustomLineItemChargeDetails) {
	if err := b.validatePutCustomLineItemChargeDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCustomLineItemChargeDetails",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItem) PutTags(value interface{}) {
	if err := b.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTags",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItem) ResetAccountId() {
	_jsii_.InvokeVoid(
		b,
		"resetAccountId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItem) ResetBillingPeriodRange() {
	_jsii_.InvokeVoid(
		b,
		"resetBillingPeriodRange",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItem) ResetCustomLineItemChargeDetails() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomLineItemChargeDetails",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItem) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItem) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItem) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorCustomLineItem) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorCustomLineItem) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

