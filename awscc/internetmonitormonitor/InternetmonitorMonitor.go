package internetmonitormonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/internetmonitormonitor/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/internetmonitor_monitor awscc_internetmonitor_monitor}.
type InternetmonitorMonitor interface {
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
	CreatedAt() *string
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
	HealthEventsConfig() InternetmonitorMonitorHealthEventsConfigOutputReference
	HealthEventsConfigInput() interface{}
	Id() *string
	InternetMeasurementsLogDelivery() InternetmonitorMonitorInternetMeasurementsLogDeliveryOutputReference
	InternetMeasurementsLogDeliveryInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxCityNetworksToMonitor() *float64
	SetMaxCityNetworksToMonitor(val *float64)
	MaxCityNetworksToMonitorInput() *float64
	ModifiedAt() *string
	MonitorArn() *string
	MonitorName() *string
	SetMonitorName(val *string)
	MonitorNameInput() *string
	// The tree node.
	Node() constructs.Node
	ProcessingStatus() *string
	ProcessingStatusInfo() *string
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
	Resources() *[]*string
	SetResources(val *[]*string)
	ResourcesInput() *[]*string
	ResourcesToAdd() *[]*string
	SetResourcesToAdd(val *[]*string)
	ResourcesToAddInput() *[]*string
	ResourcesToRemove() *[]*string
	SetResourcesToRemove(val *[]*string)
	ResourcesToRemoveInput() *[]*string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	Tags() InternetmonitorMonitorTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TrafficPercentageToMonitor() *float64
	SetTrafficPercentageToMonitor(val *float64)
	TrafficPercentageToMonitorInput() *float64
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
	PutHealthEventsConfig(value *InternetmonitorMonitorHealthEventsConfig)
	PutInternetMeasurementsLogDelivery(value *InternetmonitorMonitorInternetMeasurementsLogDelivery)
	PutTags(value interface{})
	ResetHealthEventsConfig()
	ResetInternetMeasurementsLogDelivery()
	ResetMaxCityNetworksToMonitor()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResources()
	ResetResourcesToAdd()
	ResetResourcesToRemove()
	ResetStatus()
	ResetTags()
	ResetTrafficPercentageToMonitor()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for InternetmonitorMonitor
type jsiiProxy_InternetmonitorMonitor struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_InternetmonitorMonitor) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) HealthEventsConfig() InternetmonitorMonitorHealthEventsConfigOutputReference {
	var returns InternetmonitorMonitorHealthEventsConfigOutputReference
	_jsii_.Get(
		j,
		"healthEventsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) HealthEventsConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthEventsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) InternetMeasurementsLogDelivery() InternetmonitorMonitorInternetMeasurementsLogDeliveryOutputReference {
	var returns InternetmonitorMonitorInternetMeasurementsLogDeliveryOutputReference
	_jsii_.Get(
		j,
		"internetMeasurementsLogDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) InternetMeasurementsLogDeliveryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetMeasurementsLogDeliveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) MaxCityNetworksToMonitor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCityNetworksToMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) MaxCityNetworksToMonitorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCityNetworksToMonitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) MonitorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) MonitorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) MonitorNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) ProcessingStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) ProcessingStatusInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processingStatusInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) Resources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) ResourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) ResourcesToAdd() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesToAdd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) ResourcesToAddInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesToAddInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) ResourcesToRemove() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesToRemove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) ResourcesToRemoveInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesToRemoveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) Tags() InternetmonitorMonitorTagsList {
	var returns InternetmonitorMonitorTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) TrafficPercentageToMonitor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficPercentageToMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitor) TrafficPercentageToMonitorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficPercentageToMonitorInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/internetmonitor_monitor awscc_internetmonitor_monitor} Resource.
func NewInternetmonitorMonitor(scope constructs.Construct, id *string, config *InternetmonitorMonitorConfig) InternetmonitorMonitor {
	_init_.Initialize()

	if err := validateNewInternetmonitorMonitorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_InternetmonitorMonitor{}

	_jsii_.Create(
		"awscc.internetmonitorMonitor.InternetmonitorMonitor",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/internetmonitor_monitor awscc_internetmonitor_monitor} Resource.
func NewInternetmonitorMonitor_Override(i InternetmonitorMonitor, scope constructs.Construct, id *string, config *InternetmonitorMonitorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.internetmonitorMonitor.InternetmonitorMonitor",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_InternetmonitorMonitor)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitor)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitor)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitor)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitor)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitor)SetMaxCityNetworksToMonitor(val *float64) {
	if err := j.validateSetMaxCityNetworksToMonitorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCityNetworksToMonitor",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitor)SetMonitorName(val *string) {
	if err := j.validateSetMonitorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitorName",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitor)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitor)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitor)SetResources(val *[]*string) {
	if err := j.validateSetResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resources",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitor)SetResourcesToAdd(val *[]*string) {
	if err := j.validateSetResourcesToAddParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcesToAdd",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitor)SetResourcesToRemove(val *[]*string) {
	if err := j.validateSetResourcesToRemoveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcesToRemove",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitor)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitor)SetTrafficPercentageToMonitor(val *float64) {
	if err := j.validateSetTrafficPercentageToMonitorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficPercentageToMonitor",
		val,
	)
}

// Generates CDKTF code for importing a InternetmonitorMonitor resource upon running "cdktf plan <stack-name>".
func InternetmonitorMonitor_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateInternetmonitorMonitor_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.internetmonitorMonitor.InternetmonitorMonitor",
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
func InternetmonitorMonitor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInternetmonitorMonitor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.internetmonitorMonitor.InternetmonitorMonitor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func InternetmonitorMonitor_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInternetmonitorMonitor_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.internetmonitorMonitor.InternetmonitorMonitor",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func InternetmonitorMonitor_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInternetmonitorMonitor_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.internetmonitorMonitor.InternetmonitorMonitor",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func InternetmonitorMonitor_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.internetmonitorMonitor.InternetmonitorMonitor",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_InternetmonitorMonitor) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_InternetmonitorMonitor) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_InternetmonitorMonitor) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_InternetmonitorMonitor) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_InternetmonitorMonitor) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_InternetmonitorMonitor) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_InternetmonitorMonitor) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_InternetmonitorMonitor) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_InternetmonitorMonitor) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_InternetmonitorMonitor) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_InternetmonitorMonitor) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) PutHealthEventsConfig(value *InternetmonitorMonitorHealthEventsConfig) {
	if err := i.validatePutHealthEventsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putHealthEventsConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) PutInternetMeasurementsLogDelivery(value *InternetmonitorMonitorInternetMeasurementsLogDelivery) {
	if err := i.validatePutInternetMeasurementsLogDeliveryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putInternetMeasurementsLogDelivery",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) PutTags(value interface{}) {
	if err := i.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) ResetHealthEventsConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetHealthEventsConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) ResetInternetMeasurementsLogDelivery() {
	_jsii_.InvokeVoid(
		i,
		"resetInternetMeasurementsLogDelivery",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) ResetMaxCityNetworksToMonitor() {
	_jsii_.InvokeVoid(
		i,
		"resetMaxCityNetworksToMonitor",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) ResetResources() {
	_jsii_.InvokeVoid(
		i,
		"resetResources",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) ResetResourcesToAdd() {
	_jsii_.InvokeVoid(
		i,
		"resetResourcesToAdd",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) ResetResourcesToRemove() {
	_jsii_.InvokeVoid(
		i,
		"resetResourcesToRemove",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) ResetStatus() {
	_jsii_.InvokeVoid(
		i,
		"resetStatus",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) ResetTrafficPercentageToMonitor() {
	_jsii_.InvokeVoid(
		i,
		"resetTrafficPercentageToMonitor",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitor) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitor) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitor) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

