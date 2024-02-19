package xraysamplingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/xraysamplingrule/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule awscc_xray_sampling_rule}.
type XraySamplingRule interface {
	cdktf.TerraformResource
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
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
	RuleArn() *string
	RuleName() *string
	SetRuleName(val *string)
	RuleNameInput() *string
	SamplingRule() XraySamplingRuleSamplingRuleOutputReference
	SamplingRuleInput() interface{}
	SamplingRuleRecord() XraySamplingRuleSamplingRuleRecordOutputReference
	SamplingRuleRecordInput() interface{}
	SamplingRuleUpdate() XraySamplingRuleSamplingRuleUpdateOutputReference
	SamplingRuleUpdateInput() interface{}
	Tags() XraySamplingRuleTagsList
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutSamplingRule(value *XraySamplingRuleSamplingRule)
	PutSamplingRuleRecord(value *XraySamplingRuleSamplingRuleRecord)
	PutSamplingRuleUpdate(value *XraySamplingRuleSamplingRuleUpdate)
	PutTags(value interface{})
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRuleName()
	ResetSamplingRule()
	ResetSamplingRuleRecord()
	ResetSamplingRuleUpdate()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for XraySamplingRule
type jsiiProxy_XraySamplingRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_XraySamplingRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) RuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) RuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) SamplingRule() XraySamplingRuleSamplingRuleOutputReference {
	var returns XraySamplingRuleSamplingRuleOutputReference
	_jsii_.Get(
		j,
		"samplingRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) SamplingRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samplingRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) SamplingRuleRecord() XraySamplingRuleSamplingRuleRecordOutputReference {
	var returns XraySamplingRuleSamplingRuleRecordOutputReference
	_jsii_.Get(
		j,
		"samplingRuleRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) SamplingRuleRecordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samplingRuleRecordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) SamplingRuleUpdate() XraySamplingRuleSamplingRuleUpdateOutputReference {
	var returns XraySamplingRuleSamplingRuleUpdateOutputReference
	_jsii_.Get(
		j,
		"samplingRuleUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) SamplingRuleUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samplingRuleUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) Tags() XraySamplingRuleTagsList {
	var returns XraySamplingRuleTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XraySamplingRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule awscc_xray_sampling_rule} Resource.
func NewXraySamplingRule(scope constructs.Construct, id *string, config *XraySamplingRuleConfig) XraySamplingRule {
	_init_.Initialize()

	if err := validateNewXraySamplingRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_XraySamplingRule{}

	_jsii_.Create(
		"awscc.xraySamplingRule.XraySamplingRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule awscc_xray_sampling_rule} Resource.
func NewXraySamplingRule_Override(x XraySamplingRule, scope constructs.Construct, id *string, config *XraySamplingRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.xraySamplingRule.XraySamplingRule",
		[]interface{}{scope, id, config},
		x,
	)
}

func (j *jsiiProxy_XraySamplingRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_XraySamplingRule)SetRuleName(val *string) {
	if err := j.validateSetRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

// Generates CDKTF code for importing a XraySamplingRule resource upon running "cdktf plan <stack-name>".
func XraySamplingRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateXraySamplingRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.xraySamplingRule.XraySamplingRule",
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
func XraySamplingRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateXraySamplingRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.xraySamplingRule.XraySamplingRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func XraySamplingRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateXraySamplingRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.xraySamplingRule.XraySamplingRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func XraySamplingRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateXraySamplingRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.xraySamplingRule.XraySamplingRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func XraySamplingRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.xraySamplingRule.XraySamplingRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (x *jsiiProxy_XraySamplingRule) AddMoveTarget(moveTarget *string) {
	if err := x.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (x *jsiiProxy_XraySamplingRule) AddOverride(path *string, value interface{}) {
	if err := x.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (x *jsiiProxy_XraySamplingRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := x.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		x,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := x.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		x,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := x.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		x,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := x.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		x,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := x.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		x,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := x.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		x,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := x.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		x,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRule) GetStringAttribute(terraformAttribute *string) *string {
	if err := x.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		x,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := x.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		x,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := x.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (x *jsiiProxy_XraySamplingRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := x.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		x,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRule) MoveTo(moveTarget *string, index interface{}) {
	if err := x.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (x *jsiiProxy_XraySamplingRule) OverrideLogicalId(newLogicalId *string) {
	if err := x.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (x *jsiiProxy_XraySamplingRule) PutSamplingRule(value *XraySamplingRuleSamplingRule) {
	if err := x.validatePutSamplingRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"putSamplingRule",
		[]interface{}{value},
	)
}

func (x *jsiiProxy_XraySamplingRule) PutSamplingRuleRecord(value *XraySamplingRuleSamplingRuleRecord) {
	if err := x.validatePutSamplingRuleRecordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"putSamplingRuleRecord",
		[]interface{}{value},
	)
}

func (x *jsiiProxy_XraySamplingRule) PutSamplingRuleUpdate(value *XraySamplingRuleSamplingRuleUpdate) {
	if err := x.validatePutSamplingRuleUpdateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"putSamplingRuleUpdate",
		[]interface{}{value},
	)
}

func (x *jsiiProxy_XraySamplingRule) PutTags(value interface{}) {
	if err := x.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"putTags",
		[]interface{}{value},
	)
}

func (x *jsiiProxy_XraySamplingRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		x,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRule) ResetRuleName() {
	_jsii_.InvokeVoid(
		x,
		"resetRuleName",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRule) ResetSamplingRule() {
	_jsii_.InvokeVoid(
		x,
		"resetSamplingRule",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRule) ResetSamplingRuleRecord() {
	_jsii_.InvokeVoid(
		x,
		"resetSamplingRuleRecord",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRule) ResetSamplingRuleUpdate() {
	_jsii_.InvokeVoid(
		x,
		"resetSamplingRuleUpdate",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRule) ResetTags() {
	_jsii_.InvokeVoid(
		x,
		"resetTags",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XraySamplingRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		x,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		x,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		x,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XraySamplingRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		x,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

