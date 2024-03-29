package auditmanagerassessment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/auditmanagerassessment/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/auditmanager_assessment awscc_auditmanager_assessment}.
type AuditmanagerAssessment interface {
	cdktf.TerraformResource
	Arn() *string
	AssessmentId() *string
	AssessmentReportsDestination() AuditmanagerAssessmentAssessmentReportsDestinationOutputReference
	AssessmentReportsDestinationInput() interface{}
	AwsAccount() AuditmanagerAssessmentAwsAccountOutputReference
	AwsAccountInput() interface{}
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
	Delegations() AuditmanagerAssessmentDelegationsList
	DelegationsInput() interface{}
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
	FrameworkId() *string
	SetFrameworkId(val *string)
	FrameworkIdInput() *string
	// Experimental.
	FriendlyUniqueId() *string
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
	Roles() AuditmanagerAssessmentRolesList
	RolesInput() interface{}
	Scope() AuditmanagerAssessmentScopeOutputReference
	ScopeInput() interface{}
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	Tags() AuditmanagerAssessmentTagsList
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
	PutAssessmentReportsDestination(value *AuditmanagerAssessmentAssessmentReportsDestination)
	PutAwsAccount(value *AuditmanagerAssessmentAwsAccount)
	PutDelegations(value interface{})
	PutRoles(value interface{})
	PutScope(value *AuditmanagerAssessmentScope)
	PutTags(value interface{})
	ResetAssessmentReportsDestination()
	ResetAwsAccount()
	ResetDelegations()
	ResetDescription()
	ResetFrameworkId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRoles()
	ResetScope()
	ResetStatus()
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

// The jsii proxy struct for AuditmanagerAssessment
type jsiiProxy_AuditmanagerAssessment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AuditmanagerAssessment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) AssessmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assessmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) AssessmentReportsDestination() AuditmanagerAssessmentAssessmentReportsDestinationOutputReference {
	var returns AuditmanagerAssessmentAssessmentReportsDestinationOutputReference
	_jsii_.Get(
		j,
		"assessmentReportsDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) AssessmentReportsDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assessmentReportsDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) AwsAccount() AuditmanagerAssessmentAwsAccountOutputReference {
	var returns AuditmanagerAssessmentAwsAccountOutputReference
	_jsii_.Get(
		j,
		"awsAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) AwsAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) CreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) Delegations() AuditmanagerAssessmentDelegationsList {
	var returns AuditmanagerAssessmentDelegationsList
	_jsii_.Get(
		j,
		"delegations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) DelegationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"delegationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) FrameworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) FrameworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) Roles() AuditmanagerAssessmentRolesList {
	var returns AuditmanagerAssessmentRolesList
	_jsii_.Get(
		j,
		"roles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) RolesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) Scope() AuditmanagerAssessmentScopeOutputReference {
	var returns AuditmanagerAssessmentScopeOutputReference
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) ScopeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) Tags() AuditmanagerAssessmentTagsList {
	var returns AuditmanagerAssessmentTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuditmanagerAssessment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/auditmanager_assessment awscc_auditmanager_assessment} Resource.
func NewAuditmanagerAssessment(scope constructs.Construct, id *string, config *AuditmanagerAssessmentConfig) AuditmanagerAssessment {
	_init_.Initialize()

	if err := validateNewAuditmanagerAssessmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AuditmanagerAssessment{}

	_jsii_.Create(
		"awscc.auditmanagerAssessment.AuditmanagerAssessment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/auditmanager_assessment awscc_auditmanager_assessment} Resource.
func NewAuditmanagerAssessment_Override(a AuditmanagerAssessment, scope constructs.Construct, id *string, config *AuditmanagerAssessmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.auditmanagerAssessment.AuditmanagerAssessment",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AuditmanagerAssessment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerAssessment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerAssessment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerAssessment)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerAssessment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerAssessment)SetFrameworkId(val *string) {
	if err := j.validateSetFrameworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frameworkId",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerAssessment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerAssessment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerAssessment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerAssessment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AuditmanagerAssessment)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

// Generates CDKTF code for importing a AuditmanagerAssessment resource upon running "cdktf plan <stack-name>".
func AuditmanagerAssessment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAuditmanagerAssessment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.auditmanagerAssessment.AuditmanagerAssessment",
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
func AuditmanagerAssessment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuditmanagerAssessment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.auditmanagerAssessment.AuditmanagerAssessment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AuditmanagerAssessment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuditmanagerAssessment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.auditmanagerAssessment.AuditmanagerAssessment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AuditmanagerAssessment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuditmanagerAssessment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.auditmanagerAssessment.AuditmanagerAssessment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AuditmanagerAssessment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.auditmanagerAssessment.AuditmanagerAssessment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AuditmanagerAssessment) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerAssessment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerAssessment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerAssessment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerAssessment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerAssessment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerAssessment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerAssessment) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerAssessment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerAssessment) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerAssessment) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) PutAssessmentReportsDestination(value *AuditmanagerAssessmentAssessmentReportsDestination) {
	if err := a.validatePutAssessmentReportsDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAssessmentReportsDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) PutAwsAccount(value *AuditmanagerAssessmentAwsAccount) {
	if err := a.validatePutAwsAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsAccount",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) PutDelegations(value interface{}) {
	if err := a.validatePutDelegationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDelegations",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) PutRoles(value interface{}) {
	if err := a.validatePutRolesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRoles",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) PutScope(value *AuditmanagerAssessmentScope) {
	if err := a.validatePutScopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScope",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) PutTags(value interface{}) {
	if err := a.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) ResetAssessmentReportsDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetAssessmentReportsDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) ResetAwsAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) ResetDelegations() {
	_jsii_.InvokeVoid(
		a,
		"resetDelegations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) ResetFrameworkId() {
	_jsii_.InvokeVoid(
		a,
		"resetFrameworkId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) ResetRoles() {
	_jsii_.InvokeVoid(
		a,
		"resetRoles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) ResetScope() {
	_jsii_.InvokeVoid(
		a,
		"resetScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) ResetStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuditmanagerAssessment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerAssessment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerAssessment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuditmanagerAssessment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

