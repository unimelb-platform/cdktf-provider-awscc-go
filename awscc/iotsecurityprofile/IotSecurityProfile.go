package iotsecurityprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/iotsecurityprofile/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_security_profile awscc_iot_security_profile}.
type IotSecurityProfile interface {
	cdktf.TerraformResource
	AdditionalMetricsToRetainV2() IotSecurityProfileAdditionalMetricsToRetainV2List
	AdditionalMetricsToRetainV2Input() interface{}
	AlertTargets() IotSecurityProfileAlertTargetsMap
	AlertTargetsInput() interface{}
	Behaviors() IotSecurityProfileBehaviorsList
	BehaviorsInput() interface{}
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
	MetricsExportConfig() IotSecurityProfileMetricsExportConfigOutputReference
	MetricsExportConfigInput() interface{}
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
	SecurityProfileArn() *string
	SecurityProfileDescription() *string
	SetSecurityProfileDescription(val *string)
	SecurityProfileDescriptionInput() *string
	SecurityProfileName() *string
	SetSecurityProfileName(val *string)
	SecurityProfileNameInput() *string
	Tags() IotSecurityProfileTagsList
	TagsInput() interface{}
	TargetArns() *[]*string
	SetTargetArns(val *[]*string)
	TargetArnsInput() *[]*string
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
	PutAdditionalMetricsToRetainV2(value interface{})
	PutAlertTargets(value interface{})
	PutBehaviors(value interface{})
	PutMetricsExportConfig(value *IotSecurityProfileMetricsExportConfig)
	PutTags(value interface{})
	ResetAdditionalMetricsToRetainV2()
	ResetAlertTargets()
	ResetBehaviors()
	ResetMetricsExportConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSecurityProfileDescription()
	ResetSecurityProfileName()
	ResetTags()
	ResetTargetArns()
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

// The jsii proxy struct for IotSecurityProfile
type jsiiProxy_IotSecurityProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotSecurityProfile) AdditionalMetricsToRetainV2() IotSecurityProfileAdditionalMetricsToRetainV2List {
	var returns IotSecurityProfileAdditionalMetricsToRetainV2List
	_jsii_.Get(
		j,
		"additionalMetricsToRetainV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) AdditionalMetricsToRetainV2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalMetricsToRetainV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) AlertTargets() IotSecurityProfileAlertTargetsMap {
	var returns IotSecurityProfileAlertTargetsMap
	_jsii_.Get(
		j,
		"alertTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) AlertTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) Behaviors() IotSecurityProfileBehaviorsList {
	var returns IotSecurityProfileBehaviorsList
	_jsii_.Get(
		j,
		"behaviors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) BehaviorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"behaviorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) MetricsExportConfig() IotSecurityProfileMetricsExportConfigOutputReference {
	var returns IotSecurityProfileMetricsExportConfigOutputReference
	_jsii_.Get(
		j,
		"metricsExportConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) MetricsExportConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsExportConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) SecurityProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) SecurityProfileDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) SecurityProfileDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) SecurityProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) SecurityProfileNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) Tags() IotSecurityProfileTagsList {
	var returns IotSecurityProfileTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) TargetArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) TargetArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_security_profile awscc_iot_security_profile} Resource.
func NewIotSecurityProfile(scope constructs.Construct, id *string, config *IotSecurityProfileConfig) IotSecurityProfile {
	_init_.Initialize()

	if err := validateNewIotSecurityProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotSecurityProfile{}

	_jsii_.Create(
		"awscc.iotSecurityProfile.IotSecurityProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_security_profile awscc_iot_security_profile} Resource.
func NewIotSecurityProfile_Override(i IotSecurityProfile, scope constructs.Construct, id *string, config *IotSecurityProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.iotSecurityProfile.IotSecurityProfile",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotSecurityProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfile)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfile)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfile)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfile)SetSecurityProfileDescription(val *string) {
	if err := j.validateSetSecurityProfileDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProfileDescription",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfile)SetSecurityProfileName(val *string) {
	if err := j.validateSetSecurityProfileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProfileName",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfile)SetTargetArns(val *[]*string) {
	if err := j.validateSetTargetArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetArns",
		val,
	)
}

// Generates CDKTF code for importing a IotSecurityProfile resource upon running "cdktf plan <stack-name>".
func IotSecurityProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIotSecurityProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.iotSecurityProfile.IotSecurityProfile",
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
func IotSecurityProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotSecurityProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotSecurityProfile.IotSecurityProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotSecurityProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotSecurityProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotSecurityProfile.IotSecurityProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotSecurityProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotSecurityProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotSecurityProfile.IotSecurityProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotSecurityProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.iotSecurityProfile.IotSecurityProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IotSecurityProfile) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IotSecurityProfile) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IotSecurityProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfile) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfile) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfile) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IotSecurityProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfile) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IotSecurityProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IotSecurityProfile) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IotSecurityProfile) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotSecurityProfile) PutAdditionalMetricsToRetainV2(value interface{}) {
	if err := i.validatePutAdditionalMetricsToRetainV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAdditionalMetricsToRetainV2",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotSecurityProfile) PutAlertTargets(value interface{}) {
	if err := i.validatePutAlertTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAlertTargets",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotSecurityProfile) PutBehaviors(value interface{}) {
	if err := i.validatePutBehaviorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putBehaviors",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotSecurityProfile) PutMetricsExportConfig(value *IotSecurityProfileMetricsExportConfig) {
	if err := i.validatePutMetricsExportConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMetricsExportConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotSecurityProfile) PutTags(value interface{}) {
	if err := i.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotSecurityProfile) ResetAdditionalMetricsToRetainV2() {
	_jsii_.InvokeVoid(
		i,
		"resetAdditionalMetricsToRetainV2",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfile) ResetAlertTargets() {
	_jsii_.InvokeVoid(
		i,
		"resetAlertTargets",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfile) ResetBehaviors() {
	_jsii_.InvokeVoid(
		i,
		"resetBehaviors",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfile) ResetMetricsExportConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetMetricsExportConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfile) ResetSecurityProfileDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetSecurityProfileDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfile) ResetSecurityProfileName() {
	_jsii_.InvokeVoid(
		i,
		"resetSecurityProfileName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfile) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfile) ResetTargetArns() {
	_jsii_.InvokeVoid(
		i,
		"resetTargetArns",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

