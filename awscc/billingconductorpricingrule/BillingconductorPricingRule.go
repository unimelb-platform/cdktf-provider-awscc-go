package billingconductorpricingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/billingconductorpricingrule/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_pricing_rule awscc_billingconductor_pricing_rule}.
type BillingconductorPricingRule interface {
	cdktf.TerraformResource
	Arn() *string
	AssociatedPricingPlanCount() *float64
	BillingEntity() *string
	SetBillingEntity(val *string)
	BillingEntityInput() *string
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
	ModifierPercentage() *float64
	SetModifierPercentage(val *float64)
	ModifierPercentageInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Operation() *string
	SetOperation(val *string)
	OperationInput() *string
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
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	Tags() BillingconductorPricingRuleTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tiering() BillingconductorPricingRuleTieringOutputReference
	TieringInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UsageType() *string
	SetUsageType(val *string)
	UsageTypeInput() *string
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
	PutTags(value interface{})
	PutTiering(value *BillingconductorPricingRuleTiering)
	ResetBillingEntity()
	ResetDescription()
	ResetModifierPercentage()
	ResetOperation()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetService()
	ResetTags()
	ResetTiering()
	ResetUsageType()
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

// The jsii proxy struct for BillingconductorPricingRule
type jsiiProxy_BillingconductorPricingRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BillingconductorPricingRule) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) AssociatedPricingPlanCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"associatedPricingPlanCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) BillingEntity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingEntity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) BillingEntityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingEntityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) CreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) LastModifiedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) ModifierPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modifierPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) ModifierPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modifierPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) Operation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) OperationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) Tags() BillingconductorPricingRuleTagsList {
	var returns BillingconductorPricingRuleTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) Tiering() BillingconductorPricingRuleTieringOutputReference {
	var returns BillingconductorPricingRuleTieringOutputReference
	_jsii_.Get(
		j,
		"tiering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) TieringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tieringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) UsageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorPricingRule) UsageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_pricing_rule awscc_billingconductor_pricing_rule} Resource.
func NewBillingconductorPricingRule(scope constructs.Construct, id *string, config *BillingconductorPricingRuleConfig) BillingconductorPricingRule {
	_init_.Initialize()

	if err := validateNewBillingconductorPricingRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BillingconductorPricingRule{}

	_jsii_.Create(
		"awscc.billingconductorPricingRule.BillingconductorPricingRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_pricing_rule awscc_billingconductor_pricing_rule} Resource.
func NewBillingconductorPricingRule_Override(b BillingconductorPricingRule, scope constructs.Construct, id *string, config *BillingconductorPricingRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.billingconductorPricingRule.BillingconductorPricingRule",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BillingconductorPricingRule)SetBillingEntity(val *string) {
	if err := j.validateSetBillingEntityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingEntity",
		val,
	)
}

func (j *jsiiProxy_BillingconductorPricingRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BillingconductorPricingRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BillingconductorPricingRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BillingconductorPricingRule)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BillingconductorPricingRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BillingconductorPricingRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BillingconductorPricingRule)SetModifierPercentage(val *float64) {
	if err := j.validateSetModifierPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modifierPercentage",
		val,
	)
}

func (j *jsiiProxy_BillingconductorPricingRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BillingconductorPricingRule)SetOperation(val *string) {
	if err := j.validateSetOperationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operation",
		val,
	)
}

func (j *jsiiProxy_BillingconductorPricingRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BillingconductorPricingRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BillingconductorPricingRule)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_BillingconductorPricingRule)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_BillingconductorPricingRule)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_BillingconductorPricingRule)SetUsageType(val *string) {
	if err := j.validateSetUsageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usageType",
		val,
	)
}

// Generates CDKTF code for importing a BillingconductorPricingRule resource upon running "cdktf plan <stack-name>".
func BillingconductorPricingRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBillingconductorPricingRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.billingconductorPricingRule.BillingconductorPricingRule",
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
func BillingconductorPricingRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBillingconductorPricingRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.billingconductorPricingRule.BillingconductorPricingRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BillingconductorPricingRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBillingconductorPricingRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.billingconductorPricingRule.BillingconductorPricingRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BillingconductorPricingRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBillingconductorPricingRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.billingconductorPricingRule.BillingconductorPricingRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BillingconductorPricingRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.billingconductorPricingRule.BillingconductorPricingRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BillingconductorPricingRule) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BillingconductorPricingRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BillingconductorPricingRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BillingconductorPricingRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BillingconductorPricingRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BillingconductorPricingRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BillingconductorPricingRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BillingconductorPricingRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BillingconductorPricingRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BillingconductorPricingRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorPricingRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BillingconductorPricingRule) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) PutTags(value interface{}) {
	if err := b.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTags",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) PutTiering(value *BillingconductorPricingRuleTiering) {
	if err := b.validatePutTieringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTiering",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) ResetBillingEntity() {
	_jsii_.InvokeVoid(
		b,
		"resetBillingEntity",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) ResetModifierPercentage() {
	_jsii_.InvokeVoid(
		b,
		"resetModifierPercentage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) ResetOperation() {
	_jsii_.InvokeVoid(
		b,
		"resetOperation",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) ResetService() {
	_jsii_.InvokeVoid(
		b,
		"resetService",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) ResetTiering() {
	_jsii_.InvokeVoid(
		b,
		"resetTiering",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) ResetUsageType() {
	_jsii_.InvokeVoid(
		b,
		"resetUsageType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorPricingRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorPricingRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorPricingRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorPricingRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorPricingRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorPricingRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

