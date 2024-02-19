package fmspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/fmspolicy/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fms_policy awscc_fms_policy}.
type FmsPolicy interface {
	cdktf.TerraformResource
	Arn() *string
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
	DeleteAllPolicyResources() interface{}
	SetDeleteAllPolicyResources(val interface{})
	DeleteAllPolicyResourcesInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExcludeMap() FmsPolicyExcludeMapOutputReference
	ExcludeMapInput() interface{}
	ExcludeResourceTags() interface{}
	SetExcludeResourceTags(val interface{})
	ExcludeResourceTagsInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IncludeMap() FmsPolicyIncludeMapOutputReference
	IncludeMapInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PolicyDescription() *string
	SetPolicyDescription(val *string)
	PolicyDescriptionInput() *string
	PolicyName() *string
	SetPolicyName(val *string)
	PolicyNameInput() *string
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
	RemediationEnabled() interface{}
	SetRemediationEnabled(val interface{})
	RemediationEnabledInput() interface{}
	ResourcesCleanUp() interface{}
	SetResourcesCleanUp(val interface{})
	ResourcesCleanUpInput() interface{}
	ResourceSetIds() *[]*string
	SetResourceSetIds(val *[]*string)
	ResourceSetIdsInput() *[]*string
	ResourceTags() FmsPolicyResourceTagsList
	ResourceTagsInput() interface{}
	ResourceType() *string
	SetResourceType(val *string)
	ResourceTypeInput() *string
	ResourceTypeList() *[]*string
	SetResourceTypeList(val *[]*string)
	ResourceTypeListInput() *[]*string
	SecurityServicePolicyData() FmsPolicySecurityServicePolicyDataOutputReference
	SecurityServicePolicyDataInput() interface{}
	Tags() FmsPolicyTagsList
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
	PutExcludeMap(value *FmsPolicyExcludeMap)
	PutIncludeMap(value *FmsPolicyIncludeMap)
	PutResourceTags(value interface{})
	PutSecurityServicePolicyData(value *FmsPolicySecurityServicePolicyData)
	PutTags(value interface{})
	ResetDeleteAllPolicyResources()
	ResetExcludeMap()
	ResetIncludeMap()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyDescription()
	ResetResourcesCleanUp()
	ResetResourceSetIds()
	ResetResourceTags()
	ResetResourceType()
	ResetResourceTypeList()
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

// The jsii proxy struct for FmsPolicy
type jsiiProxy_FmsPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FmsPolicy) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) DeleteAllPolicyResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAllPolicyResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) DeleteAllPolicyResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAllPolicyResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ExcludeMap() FmsPolicyExcludeMapOutputReference {
	var returns FmsPolicyExcludeMapOutputReference
	_jsii_.Get(
		j,
		"excludeMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ExcludeMapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ExcludeResourceTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeResourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ExcludeResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeResourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) IncludeMap() FmsPolicyIncludeMapOutputReference {
	var returns FmsPolicyIncludeMapOutputReference
	_jsii_.Get(
		j,
		"includeMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) IncludeMapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) PolicyDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) PolicyDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) PolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) RemediationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remediationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) RemediationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remediationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ResourcesCleanUp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcesCleanUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ResourcesCleanUpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcesCleanUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ResourceSetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceSetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ResourceSetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceSetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ResourceTags() FmsPolicyResourceTagsList {
	var returns FmsPolicyResourceTagsList
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ResourceTypeList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypeList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ResourceTypeListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypeListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) SecurityServicePolicyData() FmsPolicySecurityServicePolicyDataOutputReference {
	var returns FmsPolicySecurityServicePolicyDataOutputReference
	_jsii_.Get(
		j,
		"securityServicePolicyData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) SecurityServicePolicyDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityServicePolicyDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) Tags() FmsPolicyTagsList {
	var returns FmsPolicyTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fms_policy awscc_fms_policy} Resource.
func NewFmsPolicy(scope constructs.Construct, id *string, config *FmsPolicyConfig) FmsPolicy {
	_init_.Initialize()

	if err := validateNewFmsPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FmsPolicy{}

	_jsii_.Create(
		"awscc.fmsPolicy.FmsPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fms_policy awscc_fms_policy} Resource.
func NewFmsPolicy_Override(f FmsPolicy, scope constructs.Construct, id *string, config *FmsPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.fmsPolicy.FmsPolicy",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FmsPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy)SetDeleteAllPolicyResources(val interface{}) {
	if err := j.validateSetDeleteAllPolicyResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAllPolicyResources",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy)SetExcludeResourceTags(val interface{}) {
	if err := j.validateSetExcludeResourceTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeResourceTags",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy)SetPolicyDescription(val *string) {
	if err := j.validateSetPolicyDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDescription",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy)SetPolicyName(val *string) {
	if err := j.validateSetPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyName",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy)SetRemediationEnabled(val interface{}) {
	if err := j.validateSetRemediationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remediationEnabled",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy)SetResourcesCleanUp(val interface{}) {
	if err := j.validateSetResourcesCleanUpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcesCleanUp",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy)SetResourceSetIds(val *[]*string) {
	if err := j.validateSetResourceSetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceSetIds",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy)SetResourceTypeList(val *[]*string) {
	if err := j.validateSetResourceTypeListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypeList",
		val,
	)
}

// Generates CDKTF code for importing a FmsPolicy resource upon running "cdktf plan <stack-name>".
func FmsPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFmsPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.fmsPolicy.FmsPolicy",
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
func FmsPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFmsPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.fmsPolicy.FmsPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FmsPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFmsPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.fmsPolicy.FmsPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FmsPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFmsPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.fmsPolicy.FmsPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FmsPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.fmsPolicy.FmsPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FmsPolicy) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FmsPolicy) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FmsPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FmsPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicy) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FmsPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FmsPolicy) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FmsPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FmsPolicy) PutExcludeMap(value *FmsPolicyExcludeMap) {
	if err := f.validatePutExcludeMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putExcludeMap",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FmsPolicy) PutIncludeMap(value *FmsPolicyIncludeMap) {
	if err := f.validatePutIncludeMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putIncludeMap",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FmsPolicy) PutResourceTags(value interface{}) {
	if err := f.validatePutResourceTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putResourceTags",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FmsPolicy) PutSecurityServicePolicyData(value *FmsPolicySecurityServicePolicyData) {
	if err := f.validatePutSecurityServicePolicyDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putSecurityServicePolicyData",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FmsPolicy) PutTags(value interface{}) {
	if err := f.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTags",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FmsPolicy) ResetDeleteAllPolicyResources() {
	_jsii_.InvokeVoid(
		f,
		"resetDeleteAllPolicyResources",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) ResetExcludeMap() {
	_jsii_.InvokeVoid(
		f,
		"resetExcludeMap",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) ResetIncludeMap() {
	_jsii_.InvokeVoid(
		f,
		"resetIncludeMap",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) ResetPolicyDescription() {
	_jsii_.InvokeVoid(
		f,
		"resetPolicyDescription",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) ResetResourcesCleanUp() {
	_jsii_.InvokeVoid(
		f,
		"resetResourcesCleanUp",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) ResetResourceSetIds() {
	_jsii_.InvokeVoid(
		f,
		"resetResourceSetIds",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) ResetResourceTags() {
	_jsii_.InvokeVoid(
		f,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) ResetResourceType() {
	_jsii_.InvokeVoid(
		f,
		"resetResourceType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) ResetResourceTypeList() {
	_jsii_.InvokeVoid(
		f,
		"resetResourceTypeList",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

