package route53recoverycontrolsafetyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/route53recoverycontrolsafetyrule/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoverycontrol_safety_rule awscc_route53recoverycontrol_safety_rule}.
type Route53RecoverycontrolSafetyRule interface {
	cdktf.TerraformResource
	AssertionRule() Route53RecoverycontrolSafetyRuleAssertionRuleOutputReference
	AssertionRuleInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControlPanelArn() *string
	SetControlPanelArn(val *string)
	ControlPanelArnInput() *string
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
	GatingRule() Route53RecoverycontrolSafetyRuleGatingRuleOutputReference
	GatingRuleInput() interface{}
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	RuleConfig() Route53RecoverycontrolSafetyRuleRuleConfigOutputReference
	RuleConfigInput() interface{}
	SafetyRuleArn() *string
	Status() *string
	Tags() Route53RecoverycontrolSafetyRuleTagsList
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
	PutAssertionRule(value *Route53RecoverycontrolSafetyRuleAssertionRule)
	PutGatingRule(value *Route53RecoverycontrolSafetyRuleGatingRule)
	PutRuleConfig(value *Route53RecoverycontrolSafetyRuleRuleConfig)
	PutTags(value interface{})
	ResetAssertionRule()
	ResetControlPanelArn()
	ResetGatingRule()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRuleConfig()
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

// The jsii proxy struct for Route53RecoverycontrolSafetyRule
type jsiiProxy_Route53RecoverycontrolSafetyRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) AssertionRule() Route53RecoverycontrolSafetyRuleAssertionRuleOutputReference {
	var returns Route53RecoverycontrolSafetyRuleAssertionRuleOutputReference
	_jsii_.Get(
		j,
		"assertionRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) AssertionRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assertionRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) ControlPanelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPanelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) ControlPanelArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPanelArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) GatingRule() Route53RecoverycontrolSafetyRuleGatingRuleOutputReference {
	var returns Route53RecoverycontrolSafetyRuleGatingRuleOutputReference
	_jsii_.Get(
		j,
		"gatingRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) GatingRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatingRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) RuleConfig() Route53RecoverycontrolSafetyRuleRuleConfigOutputReference {
	var returns Route53RecoverycontrolSafetyRuleRuleConfigOutputReference
	_jsii_.Get(
		j,
		"ruleConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) RuleConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) SafetyRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"safetyRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) Tags() Route53RecoverycontrolSafetyRuleTagsList {
	var returns Route53RecoverycontrolSafetyRuleTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoverycontrol_safety_rule awscc_route53recoverycontrol_safety_rule} Resource.
func NewRoute53RecoverycontrolSafetyRule(scope constructs.Construct, id *string, config *Route53RecoverycontrolSafetyRuleConfig) Route53RecoverycontrolSafetyRule {
	_init_.Initialize()

	if err := validateNewRoute53RecoverycontrolSafetyRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53RecoverycontrolSafetyRule{}

	_jsii_.Create(
		"awscc.route53RecoverycontrolSafetyRule.Route53RecoverycontrolSafetyRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoverycontrol_safety_rule awscc_route53recoverycontrol_safety_rule} Resource.
func NewRoute53RecoverycontrolSafetyRule_Override(r Route53RecoverycontrolSafetyRule, scope constructs.Construct, id *string, config *Route53RecoverycontrolSafetyRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.route53RecoverycontrolSafetyRule.Route53RecoverycontrolSafetyRule",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule)SetControlPanelArn(val *string) {
	if err := j.validateSetControlPanelArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlPanelArn",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolSafetyRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a Route53RecoverycontrolSafetyRule resource upon running "cdktf plan <stack-name>".
func Route53RecoverycontrolSafetyRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRoute53RecoverycontrolSafetyRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.route53RecoverycontrolSafetyRule.Route53RecoverycontrolSafetyRule",
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
func Route53RecoverycontrolSafetyRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53RecoverycontrolSafetyRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.route53RecoverycontrolSafetyRule.Route53RecoverycontrolSafetyRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Route53RecoverycontrolSafetyRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53RecoverycontrolSafetyRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.route53RecoverycontrolSafetyRule.Route53RecoverycontrolSafetyRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Route53RecoverycontrolSafetyRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53RecoverycontrolSafetyRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.route53RecoverycontrolSafetyRule.Route53RecoverycontrolSafetyRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Route53RecoverycontrolSafetyRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.route53RecoverycontrolSafetyRule.Route53RecoverycontrolSafetyRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) PutAssertionRule(value *Route53RecoverycontrolSafetyRuleAssertionRule) {
	if err := r.validatePutAssertionRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAssertionRule",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) PutGatingRule(value *Route53RecoverycontrolSafetyRuleGatingRule) {
	if err := r.validatePutGatingRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putGatingRule",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) PutRuleConfig(value *Route53RecoverycontrolSafetyRuleRuleConfig) {
	if err := r.validatePutRuleConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRuleConfig",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) PutTags(value interface{}) {
	if err := r.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTags",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) ResetAssertionRule() {
	_jsii_.InvokeVoid(
		r,
		"resetAssertionRule",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) ResetControlPanelArn() {
	_jsii_.InvokeVoid(
		r,
		"resetControlPanelArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) ResetGatingRule() {
	_jsii_.InvokeVoid(
		r,
		"resetGatingRule",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) ResetName() {
	_jsii_.InvokeVoid(
		r,
		"resetName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) ResetRuleConfig() {
	_jsii_.InvokeVoid(
		r,
		"resetRuleConfig",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolSafetyRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

