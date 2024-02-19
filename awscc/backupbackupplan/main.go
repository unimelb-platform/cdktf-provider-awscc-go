package backupbackupplan

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"awscc.backupBackupPlan.BackupBackupPlan",
		reflect.TypeOf((*BackupBackupPlan)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "backupPlan", GoGetter: "BackupPlan"},
			_jsii_.MemberProperty{JsiiProperty: "backupPlanArn", GoGetter: "BackupPlanArn"},
			_jsii_.MemberProperty{JsiiProperty: "backupPlanId", GoGetter: "BackupPlanId"},
			_jsii_.MemberProperty{JsiiProperty: "backupPlanInput", GoGetter: "BackupPlanInput"},
			_jsii_.MemberProperty{JsiiProperty: "backupPlanTags", GoGetter: "BackupPlanTags"},
			_jsii_.MemberProperty{JsiiProperty: "backupPlanTagsInput", GoGetter: "BackupPlanTagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putBackupPlan", GoMethod: "PutBackupPlan"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackupPlanTags", GoMethod: "ResetBackupPlanTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "versionId", GoGetter: "VersionId"},
		},
		func() interface{} {
			j := jsiiProxy_BackupBackupPlan{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlan",
		reflect.TypeOf((*BackupBackupPlanBackupPlan)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanAdvancedBackupSettings",
		reflect.TypeOf((*BackupBackupPlanBackupPlanAdvancedBackupSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanAdvancedBackupSettingsList",
		reflect.TypeOf((*BackupBackupPlanBackupPlanAdvancedBackupSettingsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_BackupBackupPlanBackupPlanAdvancedBackupSettingsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanAdvancedBackupSettingsOutputReference",
		reflect.TypeOf((*BackupBackupPlanBackupPlanAdvancedBackupSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backupOptions", GoGetter: "BackupOptions"},
			_jsii_.MemberProperty{JsiiProperty: "backupOptionsInput", GoGetter: "BackupOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "resourceTypeInput", GoGetter: "ResourceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BackupBackupPlanBackupPlanAdvancedBackupSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanBackupPlanRule",
		reflect.TypeOf((*BackupBackupPlanBackupPlanBackupPlanRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanBackupPlanRuleCopyActions",
		reflect.TypeOf((*BackupBackupPlanBackupPlanBackupPlanRuleCopyActions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanBackupPlanRuleCopyActionsLifecycle",
		reflect.TypeOf((*BackupBackupPlanBackupPlanBackupPlanRuleCopyActionsLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanBackupPlanRuleCopyActionsLifecycleOutputReference",
		reflect.TypeOf((*BackupBackupPlanBackupPlanBackupPlanRuleCopyActionsLifecycleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAfterDays", GoGetter: "DeleteAfterDays"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAfterDaysInput", GoGetter: "DeleteAfterDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "moveToColdStorageAfterDays", GoGetter: "MoveToColdStorageAfterDays"},
			_jsii_.MemberProperty{JsiiProperty: "moveToColdStorageAfterDaysInput", GoGetter: "MoveToColdStorageAfterDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "optInToArchiveForSupportedResources", GoGetter: "OptInToArchiveForSupportedResources"},
			_jsii_.MemberProperty{JsiiProperty: "optInToArchiveForSupportedResourcesInput", GoGetter: "OptInToArchiveForSupportedResourcesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeleteAfterDays", GoMethod: "ResetDeleteAfterDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetMoveToColdStorageAfterDays", GoMethod: "ResetMoveToColdStorageAfterDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptInToArchiveForSupportedResources", GoMethod: "ResetOptInToArchiveForSupportedResources"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleCopyActionsLifecycleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanBackupPlanRuleCopyActionsList",
		reflect.TypeOf((*BackupBackupPlanBackupPlanBackupPlanRuleCopyActionsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleCopyActionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanBackupPlanRuleCopyActionsOutputReference",
		reflect.TypeOf((*BackupBackupPlanBackupPlanBackupPlanRuleCopyActionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destinationBackupVaultArn", GoGetter: "DestinationBackupVaultArn"},
			_jsii_.MemberProperty{JsiiProperty: "destinationBackupVaultArnInput", GoGetter: "DestinationBackupVaultArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleInput", GoGetter: "LifecycleInput"},
			_jsii_.MemberMethod{JsiiMethod: "putLifecycle", GoMethod: "PutLifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycle", GoMethod: "ResetLifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleCopyActionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanBackupPlanRuleLifecycle",
		reflect.TypeOf((*BackupBackupPlanBackupPlanBackupPlanRuleLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference",
		reflect.TypeOf((*BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAfterDays", GoGetter: "DeleteAfterDays"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAfterDaysInput", GoGetter: "DeleteAfterDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "moveToColdStorageAfterDays", GoGetter: "MoveToColdStorageAfterDays"},
			_jsii_.MemberProperty{JsiiProperty: "moveToColdStorageAfterDaysInput", GoGetter: "MoveToColdStorageAfterDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "optInToArchiveForSupportedResources", GoGetter: "OptInToArchiveForSupportedResources"},
			_jsii_.MemberProperty{JsiiProperty: "optInToArchiveForSupportedResourcesInput", GoGetter: "OptInToArchiveForSupportedResourcesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeleteAfterDays", GoMethod: "ResetDeleteAfterDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetMoveToColdStorageAfterDays", GoMethod: "ResetMoveToColdStorageAfterDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptInToArchiveForSupportedResources", GoMethod: "ResetOptInToArchiveForSupportedResources"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanBackupPlanRuleList",
		reflect.TypeOf((*BackupBackupPlanBackupPlanBackupPlanRuleList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanBackupPlanRuleOutputReference",
		reflect.TypeOf((*BackupBackupPlanBackupPlanBackupPlanRuleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "completionWindowMinutes", GoGetter: "CompletionWindowMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "completionWindowMinutesInput", GoGetter: "CompletionWindowMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "copyActions", GoGetter: "CopyActions"},
			_jsii_.MemberProperty{JsiiProperty: "copyActionsInput", GoGetter: "CopyActionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableContinuousBackup", GoGetter: "EnableContinuousBackup"},
			_jsii_.MemberProperty{JsiiProperty: "enableContinuousBackupInput", GoGetter: "EnableContinuousBackupInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleInput", GoGetter: "LifecycleInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCopyActions", GoMethod: "PutCopyActions"},
			_jsii_.MemberMethod{JsiiMethod: "putLifecycle", GoMethod: "PutLifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryPointTags", GoGetter: "RecoveryPointTags"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryPointTagsInput", GoGetter: "RecoveryPointTagsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompletionWindowMinutes", GoMethod: "ResetCompletionWindowMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetCopyActions", GoMethod: "ResetCopyActions"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableContinuousBackup", GoMethod: "ResetEnableContinuousBackup"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycle", GoMethod: "ResetLifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecoveryPointTags", GoMethod: "ResetRecoveryPointTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheduleExpression", GoMethod: "ResetScheduleExpression"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheduleExpressionTimezone", GoMethod: "ResetScheduleExpressionTimezone"},
			_jsii_.MemberMethod{JsiiMethod: "resetStartWindowMinutes", GoMethod: "ResetStartWindowMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "ruleName", GoGetter: "RuleName"},
			_jsii_.MemberProperty{JsiiProperty: "ruleNameInput", GoGetter: "RuleNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleExpression", GoGetter: "ScheduleExpression"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleExpressionInput", GoGetter: "ScheduleExpressionInput"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleExpressionTimezone", GoGetter: "ScheduleExpressionTimezone"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleExpressionTimezoneInput", GoGetter: "ScheduleExpressionTimezoneInput"},
			_jsii_.MemberProperty{JsiiProperty: "startWindowMinutes", GoGetter: "StartWindowMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "startWindowMinutesInput", GoGetter: "StartWindowMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetBackupVault", GoGetter: "TargetBackupVault"},
			_jsii_.MemberProperty{JsiiProperty: "targetBackupVaultInput", GoGetter: "TargetBackupVaultInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanOutputReference",
		reflect.TypeOf((*BackupBackupPlanBackupPlanOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "advancedBackupSettings", GoGetter: "AdvancedBackupSettings"},
			_jsii_.MemberProperty{JsiiProperty: "advancedBackupSettingsInput", GoGetter: "AdvancedBackupSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "backupPlanName", GoGetter: "BackupPlanName"},
			_jsii_.MemberProperty{JsiiProperty: "backupPlanNameInput", GoGetter: "BackupPlanNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "backupPlanRule", GoGetter: "BackupPlanRule"},
			_jsii_.MemberProperty{JsiiProperty: "backupPlanRuleInput", GoGetter: "BackupPlanRuleInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putAdvancedBackupSettings", GoMethod: "PutAdvancedBackupSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putBackupPlanRule", GoMethod: "PutBackupPlanRule"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdvancedBackupSettings", GoMethod: "ResetAdvancedBackupSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BackupBackupPlanBackupPlanOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.backupBackupPlan.BackupBackupPlanConfig",
		reflect.TypeOf((*BackupBackupPlanConfig)(nil)).Elem(),
	)
}
