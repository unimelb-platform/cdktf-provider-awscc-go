package grafanaworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/grafanaworkspace/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/grafana_workspace awscc_grafana_workspace}.
type GrafanaWorkspace interface {
	cdktf.TerraformResource
	AccountAccessType() *string
	SetAccountAccessType(val *string)
	AccountAccessTypeInput() *string
	AuthenticationProviders() *[]*string
	SetAuthenticationProviders(val *[]*string)
	AuthenticationProvidersInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientToken() *string
	SetClientToken(val *string)
	ClientTokenInput() *string
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
	CreationTimestamp() *string
	DataSources() *[]*string
	SetDataSources(val *[]*string)
	DataSourcesInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Endpoint() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GrafanaVersion() *string
	SetGrafanaVersion(val *string)
	GrafanaVersionInput() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ModificationTimestamp() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkAccessControl() GrafanaWorkspaceNetworkAccessControlOutputReference
	NetworkAccessControlInput() interface{}
	// The tree node.
	Node() constructs.Node
	NotificationDestinations() *[]*string
	SetNotificationDestinations(val *[]*string)
	NotificationDestinationsInput() *[]*string
	OrganizationalUnits() *[]*string
	SetOrganizationalUnits(val *[]*string)
	OrganizationalUnitsInput() *[]*string
	OrganizationRoleName() *string
	SetOrganizationRoleName(val *string)
	OrganizationRoleNameInput() *string
	PermissionType() *string
	SetPermissionType(val *string)
	PermissionTypeInput() *string
	PluginAdminEnabled() interface{}
	SetPluginAdminEnabled(val interface{})
	PluginAdminEnabledInput() interface{}
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
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	SamlConfiguration() GrafanaWorkspaceSamlConfigurationOutputReference
	SamlConfigurationInput() interface{}
	SamlConfigurationStatus() *string
	SsoClientId() *string
	StackSetName() *string
	SetStackSetName(val *string)
	StackSetNameInput() *string
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VpcConfiguration() GrafanaWorkspaceVpcConfigurationOutputReference
	VpcConfigurationInput() interface{}
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
	PutNetworkAccessControl(value *GrafanaWorkspaceNetworkAccessControl)
	PutSamlConfiguration(value *GrafanaWorkspaceSamlConfiguration)
	PutVpcConfiguration(value *GrafanaWorkspaceVpcConfiguration)
	ResetClientToken()
	ResetDataSources()
	ResetDescription()
	ResetGrafanaVersion()
	ResetName()
	ResetNetworkAccessControl()
	ResetNotificationDestinations()
	ResetOrganizationalUnits()
	ResetOrganizationRoleName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPluginAdminEnabled()
	ResetRoleArn()
	ResetSamlConfiguration()
	ResetStackSetName()
	ResetVpcConfiguration()
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

// The jsii proxy struct for GrafanaWorkspace
type jsiiProxy_GrafanaWorkspace struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GrafanaWorkspace) AccountAccessType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountAccessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) AccountAccessTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountAccessTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) AuthenticationProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authenticationProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) AuthenticationProvidersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authenticationProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) ClientTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) DataSources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) DataSourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) GrafanaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) GrafanaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) ModificationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modificationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) NetworkAccessControl() GrafanaWorkspaceNetworkAccessControlOutputReference {
	var returns GrafanaWorkspaceNetworkAccessControlOutputReference
	_jsii_.Get(
		j,
		"networkAccessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) NetworkAccessControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkAccessControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) NotificationDestinations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) NotificationDestinationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) OrganizationalUnits() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationalUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) OrganizationalUnitsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationalUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) OrganizationRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) OrganizationRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) PermissionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) PermissionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) PluginAdminEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pluginAdminEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) PluginAdminEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pluginAdminEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) SamlConfiguration() GrafanaWorkspaceSamlConfigurationOutputReference {
	var returns GrafanaWorkspaceSamlConfigurationOutputReference
	_jsii_.Get(
		j,
		"samlConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) SamlConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samlConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) SamlConfigurationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlConfigurationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) SsoClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) StackSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) StackSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) VpcConfiguration() GrafanaWorkspaceVpcConfigurationOutputReference {
	var returns GrafanaWorkspaceVpcConfigurationOutputReference
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspace) VpcConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigurationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/grafana_workspace awscc_grafana_workspace} Resource.
func NewGrafanaWorkspace(scope constructs.Construct, id *string, config *GrafanaWorkspaceConfig) GrafanaWorkspace {
	_init_.Initialize()

	if err := validateNewGrafanaWorkspaceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GrafanaWorkspace{}

	_jsii_.Create(
		"awscc.grafanaWorkspace.GrafanaWorkspace",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/grafana_workspace awscc_grafana_workspace} Resource.
func NewGrafanaWorkspace_Override(g GrafanaWorkspace, scope constructs.Construct, id *string, config *GrafanaWorkspaceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.grafanaWorkspace.GrafanaWorkspace",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetAccountAccessType(val *string) {
	if err := j.validateSetAccountAccessTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountAccessType",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetAuthenticationProviders(val *[]*string) {
	if err := j.validateSetAuthenticationProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationProviders",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetClientToken(val *string) {
	if err := j.validateSetClientTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientToken",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetDataSources(val *[]*string) {
	if err := j.validateSetDataSourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSources",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetGrafanaVersion(val *string) {
	if err := j.validateSetGrafanaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grafanaVersion",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetNotificationDestinations(val *[]*string) {
	if err := j.validateSetNotificationDestinationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationDestinations",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetOrganizationalUnits(val *[]*string) {
	if err := j.validateSetOrganizationalUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationalUnits",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetOrganizationRoleName(val *string) {
	if err := j.validateSetOrganizationRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationRoleName",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetPermissionType(val *string) {
	if err := j.validateSetPermissionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissionType",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetPluginAdminEnabled(val interface{}) {
	if err := j.validateSetPluginAdminEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginAdminEnabled",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspace)SetStackSetName(val *string) {
	if err := j.validateSetStackSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackSetName",
		val,
	)
}

// Generates CDKTF code for importing a GrafanaWorkspace resource upon running "cdktf plan <stack-name>".
func GrafanaWorkspace_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGrafanaWorkspace_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.grafanaWorkspace.GrafanaWorkspace",
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
func GrafanaWorkspace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrafanaWorkspace_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.grafanaWorkspace.GrafanaWorkspace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GrafanaWorkspace_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrafanaWorkspace_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.grafanaWorkspace.GrafanaWorkspace",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GrafanaWorkspace_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrafanaWorkspace_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.grafanaWorkspace.GrafanaWorkspace",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GrafanaWorkspace_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.grafanaWorkspace.GrafanaWorkspace",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GrafanaWorkspace) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GrafanaWorkspace) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GrafanaWorkspace) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GrafanaWorkspace) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GrafanaWorkspace) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GrafanaWorkspace) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GrafanaWorkspace) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GrafanaWorkspace) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GrafanaWorkspace) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GrafanaWorkspace) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GrafanaWorkspace) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GrafanaWorkspace) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspace) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GrafanaWorkspace) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GrafanaWorkspace) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GrafanaWorkspace) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GrafanaWorkspace) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GrafanaWorkspace) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GrafanaWorkspace) PutNetworkAccessControl(value *GrafanaWorkspaceNetworkAccessControl) {
	if err := g.validatePutNetworkAccessControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkAccessControl",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrafanaWorkspace) PutSamlConfiguration(value *GrafanaWorkspaceSamlConfiguration) {
	if err := g.validatePutSamlConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSamlConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrafanaWorkspace) PutVpcConfiguration(value *GrafanaWorkspaceVpcConfiguration) {
	if err := g.validatePutVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVpcConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrafanaWorkspace) ResetClientToken() {
	_jsii_.InvokeVoid(
		g,
		"resetClientToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspace) ResetDataSources() {
	_jsii_.InvokeVoid(
		g,
		"resetDataSources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspace) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspace) ResetGrafanaVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetGrafanaVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspace) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspace) ResetNetworkAccessControl() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkAccessControl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspace) ResetNotificationDestinations() {
	_jsii_.InvokeVoid(
		g,
		"resetNotificationDestinations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspace) ResetOrganizationalUnits() {
	_jsii_.InvokeVoid(
		g,
		"resetOrganizationalUnits",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspace) ResetOrganizationRoleName() {
	_jsii_.InvokeVoid(
		g,
		"resetOrganizationRoleName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspace) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspace) ResetPluginAdminEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetPluginAdminEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspace) ResetRoleArn() {
	_jsii_.InvokeVoid(
		g,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspace) ResetSamlConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetSamlConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspace) ResetStackSetName() {
	_jsii_.InvokeVoid(
		g,
		"resetStackSetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspace) ResetVpcConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspace) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspace) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspace) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspace) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspace) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

