package greengrassv2deployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/greengrassv2deployment/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_deployment awscc_greengrassv2_deployment}.
type Greengrassv2Deployment interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Components() Greengrassv2DeploymentComponentsMap
	ComponentsInput() interface{}
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
	DeploymentId() *string
	DeploymentName() *string
	SetDeploymentName(val *string)
	DeploymentNameInput() *string
	DeploymentPolicies() Greengrassv2DeploymentDeploymentPoliciesOutputReference
	DeploymentPoliciesInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IotJobConfiguration() Greengrassv2DeploymentIotJobConfigurationOutputReference
	IotJobConfigurationInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	ParentTargetArn() *string
	SetParentTargetArn(val *string)
	ParentTargetArnInput() *string
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TargetArn() *string
	SetTargetArn(val *string)
	TargetArnInput() *string
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
	PutComponents(value interface{})
	PutDeploymentPolicies(value *Greengrassv2DeploymentDeploymentPolicies)
	PutIotJobConfiguration(value *Greengrassv2DeploymentIotJobConfiguration)
	ResetComponents()
	ResetDeploymentName()
	ResetDeploymentPolicies()
	ResetIotJobConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentTargetArn()
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

// The jsii proxy struct for Greengrassv2Deployment
type jsiiProxy_Greengrassv2Deployment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Greengrassv2Deployment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) Components() Greengrassv2DeploymentComponentsMap {
	var returns Greengrassv2DeploymentComponentsMap
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) ComponentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) DeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) DeploymentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) DeploymentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) DeploymentPolicies() Greengrassv2DeploymentDeploymentPoliciesOutputReference {
	var returns Greengrassv2DeploymentDeploymentPoliciesOutputReference
	_jsii_.Get(
		j,
		"deploymentPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) DeploymentPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) IotJobConfiguration() Greengrassv2DeploymentIotJobConfigurationOutputReference {
	var returns Greengrassv2DeploymentIotJobConfigurationOutputReference
	_jsii_.Get(
		j,
		"iotJobConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) IotJobConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotJobConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) ParentTargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentTargetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) ParentTargetArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentTargetArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) TargetArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2Deployment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_deployment awscc_greengrassv2_deployment} Resource.
func NewGreengrassv2Deployment(scope constructs.Construct, id *string, config *Greengrassv2DeploymentConfig) Greengrassv2Deployment {
	_init_.Initialize()

	if err := validateNewGreengrassv2DeploymentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Greengrassv2Deployment{}

	_jsii_.Create(
		"awscc.greengrassv2Deployment.Greengrassv2Deployment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_deployment awscc_greengrassv2_deployment} Resource.
func NewGreengrassv2Deployment_Override(g Greengrassv2Deployment, scope constructs.Construct, id *string, config *Greengrassv2DeploymentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.greengrassv2Deployment.Greengrassv2Deployment",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_Greengrassv2Deployment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2Deployment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2Deployment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2Deployment)SetDeploymentName(val *string) {
	if err := j.validateSetDeploymentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentName",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2Deployment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2Deployment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2Deployment)SetParentTargetArn(val *string) {
	if err := j.validateSetParentTargetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentTargetArn",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2Deployment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2Deployment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2Deployment)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2Deployment)SetTargetArn(val *string) {
	if err := j.validateSetTargetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetArn",
		val,
	)
}

// Generates CDKTF code for importing a Greengrassv2Deployment resource upon running "cdktf plan <stack-name>".
func Greengrassv2Deployment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGreengrassv2Deployment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.greengrassv2Deployment.Greengrassv2Deployment",
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
func Greengrassv2Deployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGreengrassv2Deployment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.greengrassv2Deployment.Greengrassv2Deployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Greengrassv2Deployment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGreengrassv2Deployment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.greengrassv2Deployment.Greengrassv2Deployment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Greengrassv2Deployment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGreengrassv2Deployment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.greengrassv2Deployment.Greengrassv2Deployment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Greengrassv2Deployment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.greengrassv2Deployment.Greengrassv2Deployment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) PutComponents(value interface{}) {
	if err := g.validatePutComponentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putComponents",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) PutDeploymentPolicies(value *Greengrassv2DeploymentDeploymentPolicies) {
	if err := g.validatePutDeploymentPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeploymentPolicies",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) PutIotJobConfiguration(value *Greengrassv2DeploymentIotJobConfiguration) {
	if err := g.validatePutIotJobConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIotJobConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) ResetComponents() {
	_jsii_.InvokeVoid(
		g,
		"resetComponents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) ResetDeploymentName() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) ResetDeploymentPolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentPolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) ResetIotJobConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetIotJobConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) ResetParentTargetArn() {
	_jsii_.InvokeVoid(
		g,
		"resetParentTargetArn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2Deployment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2Deployment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

