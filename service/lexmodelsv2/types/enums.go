// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AggregatedUtterancesFilterName string

// Enum values for AggregatedUtterancesFilterName
const (
	AggregatedUtterancesFilterNameUtterance AggregatedUtterancesFilterName = "Utterance"
)

// Values returns all known values for AggregatedUtterancesFilterName. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (AggregatedUtterancesFilterName) Values() []AggregatedUtterancesFilterName {
	return []AggregatedUtterancesFilterName{
		"Utterance",
	}
}

type AggregatedUtterancesFilterOperator string

// Enum values for AggregatedUtterancesFilterOperator
const (
	AggregatedUtterancesFilterOperatorContains AggregatedUtterancesFilterOperator = "CO"
	AggregatedUtterancesFilterOperatorEquals   AggregatedUtterancesFilterOperator = "EQ"
)

// Values returns all known values for AggregatedUtterancesFilterOperator. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (AggregatedUtterancesFilterOperator) Values() []AggregatedUtterancesFilterOperator {
	return []AggregatedUtterancesFilterOperator{
		"CO",
		"EQ",
	}
}

type AggregatedUtterancesSortAttribute string

// Enum values for AggregatedUtterancesSortAttribute
const (
	AggregatedUtterancesSortAttributeHitCount    AggregatedUtterancesSortAttribute = "HitCount"
	AggregatedUtterancesSortAttributeMissedCount AggregatedUtterancesSortAttribute = "MissedCount"
)

// Values returns all known values for AggregatedUtterancesSortAttribute. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (AggregatedUtterancesSortAttribute) Values() []AggregatedUtterancesSortAttribute {
	return []AggregatedUtterancesSortAttribute{
		"HitCount",
		"MissedCount",
	}
}

type AssociatedTranscriptFilterName string

// Enum values for AssociatedTranscriptFilterName
const (
	AssociatedTranscriptFilterNameIntentId   AssociatedTranscriptFilterName = "IntentId"
	AssociatedTranscriptFilterNameSlotTypeId AssociatedTranscriptFilterName = "SlotTypeId"
)

// Values returns all known values for AssociatedTranscriptFilterName. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (AssociatedTranscriptFilterName) Values() []AssociatedTranscriptFilterName {
	return []AssociatedTranscriptFilterName{
		"IntentId",
		"SlotTypeId",
	}
}

type AudioRecognitionStrategy string

// Enum values for AudioRecognitionStrategy
const (
	AudioRecognitionStrategyUseSlotValuesAsCustomVocabulary AudioRecognitionStrategy = "UseSlotValuesAsCustomVocabulary"
)

// Values returns all known values for AudioRecognitionStrategy. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AudioRecognitionStrategy) Values() []AudioRecognitionStrategy {
	return []AudioRecognitionStrategy{
		"UseSlotValuesAsCustomVocabulary",
	}
}

type BotAliasStatus string

// Enum values for BotAliasStatus
const (
	BotAliasStatusCreating  BotAliasStatus = "Creating"
	BotAliasStatusAvailable BotAliasStatus = "Available"
	BotAliasStatusDeleting  BotAliasStatus = "Deleting"
	BotAliasStatusFailed    BotAliasStatus = "Failed"
)

// Values returns all known values for BotAliasStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotAliasStatus) Values() []BotAliasStatus {
	return []BotAliasStatus{
		"Creating",
		"Available",
		"Deleting",
		"Failed",
	}
}

type BotFilterName string

// Enum values for BotFilterName
const (
	BotFilterNameBotName BotFilterName = "BotName"
)

// Values returns all known values for BotFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotFilterName) Values() []BotFilterName {
	return []BotFilterName{
		"BotName",
	}
}

type BotFilterOperator string

// Enum values for BotFilterOperator
const (
	BotFilterOperatorContains BotFilterOperator = "CO"
	BotFilterOperatorEquals   BotFilterOperator = "EQ"
)

// Values returns all known values for BotFilterOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotFilterOperator) Values() []BotFilterOperator {
	return []BotFilterOperator{
		"CO",
		"EQ",
	}
}

type BotLocaleFilterName string

// Enum values for BotLocaleFilterName
const (
	BotLocaleFilterNameBotLocaleName BotLocaleFilterName = "BotLocaleName"
)

// Values returns all known values for BotLocaleFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotLocaleFilterName) Values() []BotLocaleFilterName {
	return []BotLocaleFilterName{
		"BotLocaleName",
	}
}

type BotLocaleFilterOperator string

// Enum values for BotLocaleFilterOperator
const (
	BotLocaleFilterOperatorContains BotLocaleFilterOperator = "CO"
	BotLocaleFilterOperatorEquals   BotLocaleFilterOperator = "EQ"
)

// Values returns all known values for BotLocaleFilterOperator. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotLocaleFilterOperator) Values() []BotLocaleFilterOperator {
	return []BotLocaleFilterOperator{
		"CO",
		"EQ",
	}
}

type BotLocaleSortAttribute string

// Enum values for BotLocaleSortAttribute
const (
	BotLocaleSortAttributeBotLocaleName BotLocaleSortAttribute = "BotLocaleName"
)

// Values returns all known values for BotLocaleSortAttribute. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotLocaleSortAttribute) Values() []BotLocaleSortAttribute {
	return []BotLocaleSortAttribute{
		"BotLocaleName",
	}
}

type BotLocaleStatus string

// Enum values for BotLocaleStatus
const (
	BotLocaleStatusCreating            BotLocaleStatus = "Creating"
	BotLocaleStatusBuilding            BotLocaleStatus = "Building"
	BotLocaleStatusBuilt               BotLocaleStatus = "Built"
	BotLocaleStatusReadyExpressTesting BotLocaleStatus = "ReadyExpressTesting"
	BotLocaleStatusFailed              BotLocaleStatus = "Failed"
	BotLocaleStatusDeleting            BotLocaleStatus = "Deleting"
	BotLocaleStatusNotBuilt            BotLocaleStatus = "NotBuilt"
	BotLocaleStatusImporting           BotLocaleStatus = "Importing"
	BotLocaleStatusProcessing          BotLocaleStatus = "Processing"
)

// Values returns all known values for BotLocaleStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotLocaleStatus) Values() []BotLocaleStatus {
	return []BotLocaleStatus{
		"Creating",
		"Building",
		"Built",
		"ReadyExpressTesting",
		"Failed",
		"Deleting",
		"NotBuilt",
		"Importing",
		"Processing",
	}
}

type BotRecommendationStatus string

// Enum values for BotRecommendationStatus
const (
	BotRecommendationStatusProcessing  BotRecommendationStatus = "Processing"
	BotRecommendationStatusDeleting    BotRecommendationStatus = "Deleting"
	BotRecommendationStatusDeleted     BotRecommendationStatus = "Deleted"
	BotRecommendationStatusDownloading BotRecommendationStatus = "Downloading"
	BotRecommendationStatusUpdating    BotRecommendationStatus = "Updating"
	BotRecommendationStatusAvailable   BotRecommendationStatus = "Available"
	BotRecommendationStatusFailed      BotRecommendationStatus = "Failed"
)

// Values returns all known values for BotRecommendationStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotRecommendationStatus) Values() []BotRecommendationStatus {
	return []BotRecommendationStatus{
		"Processing",
		"Deleting",
		"Deleted",
		"Downloading",
		"Updating",
		"Available",
		"Failed",
	}
}

type BotSortAttribute string

// Enum values for BotSortAttribute
const (
	BotSortAttributeBotName BotSortAttribute = "BotName"
)

// Values returns all known values for BotSortAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotSortAttribute) Values() []BotSortAttribute {
	return []BotSortAttribute{
		"BotName",
	}
}

type BotStatus string

// Enum values for BotStatus
const (
	BotStatusCreating   BotStatus = "Creating"
	BotStatusAvailable  BotStatus = "Available"
	BotStatusInactive   BotStatus = "Inactive"
	BotStatusDeleting   BotStatus = "Deleting"
	BotStatusFailed     BotStatus = "Failed"
	BotStatusVersioning BotStatus = "Versioning"
	BotStatusImporting  BotStatus = "Importing"
)

// Values returns all known values for BotStatus. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (BotStatus) Values() []BotStatus {
	return []BotStatus{
		"Creating",
		"Available",
		"Inactive",
		"Deleting",
		"Failed",
		"Versioning",
		"Importing",
	}
}

type BotVersionSortAttribute string

// Enum values for BotVersionSortAttribute
const (
	BotVersionSortAttributeBotVersion BotVersionSortAttribute = "BotVersion"
)

// Values returns all known values for BotVersionSortAttribute. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotVersionSortAttribute) Values() []BotVersionSortAttribute {
	return []BotVersionSortAttribute{
		"BotVersion",
	}
}

type BuiltInIntentSortAttribute string

// Enum values for BuiltInIntentSortAttribute
const (
	BuiltInIntentSortAttributeIntentSignature BuiltInIntentSortAttribute = "IntentSignature"
)

// Values returns all known values for BuiltInIntentSortAttribute. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (BuiltInIntentSortAttribute) Values() []BuiltInIntentSortAttribute {
	return []BuiltInIntentSortAttribute{
		"IntentSignature",
	}
}

type BuiltInSlotTypeSortAttribute string

// Enum values for BuiltInSlotTypeSortAttribute
const (
	BuiltInSlotTypeSortAttributeSlotTypeSignature BuiltInSlotTypeSortAttribute = "SlotTypeSignature"
)

// Values returns all known values for BuiltInSlotTypeSortAttribute. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (BuiltInSlotTypeSortAttribute) Values() []BuiltInSlotTypeSortAttribute {
	return []BuiltInSlotTypeSortAttribute{
		"SlotTypeSignature",
	}
}

type CustomVocabularyStatus string

// Enum values for CustomVocabularyStatus
const (
	CustomVocabularyStatusReady     CustomVocabularyStatus = "Ready"
	CustomVocabularyStatusDeleting  CustomVocabularyStatus = "Deleting"
	CustomVocabularyStatusExporting CustomVocabularyStatus = "Exporting"
	CustomVocabularyStatusImporting CustomVocabularyStatus = "Importing"
	CustomVocabularyStatusCreating  CustomVocabularyStatus = "Creating"
)

// Values returns all known values for CustomVocabularyStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CustomVocabularyStatus) Values() []CustomVocabularyStatus {
	return []CustomVocabularyStatus{
		"Ready",
		"Deleting",
		"Exporting",
		"Importing",
		"Creating",
	}
}

type Effect string

// Enum values for Effect
const (
	EffectAllow Effect = "Allow"
	EffectDeny  Effect = "Deny"
)

// Values returns all known values for Effect. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Effect) Values() []Effect {
	return []Effect{
		"Allow",
		"Deny",
	}
}

type ExportFilterName string

// Enum values for ExportFilterName
const (
	ExportFilterNameExportResourceType ExportFilterName = "ExportResourceType"
)

// Values returns all known values for ExportFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExportFilterName) Values() []ExportFilterName {
	return []ExportFilterName{
		"ExportResourceType",
	}
}

type ExportFilterOperator string

// Enum values for ExportFilterOperator
const (
	ExportFilterOperatorContains ExportFilterOperator = "CO"
	ExportFilterOperatorEquals   ExportFilterOperator = "EQ"
)

// Values returns all known values for ExportFilterOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExportFilterOperator) Values() []ExportFilterOperator {
	return []ExportFilterOperator{
		"CO",
		"EQ",
	}
}

type ExportSortAttribute string

// Enum values for ExportSortAttribute
const (
	ExportSortAttributeLastUpdatedDateTime ExportSortAttribute = "LastUpdatedDateTime"
)

// Values returns all known values for ExportSortAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExportSortAttribute) Values() []ExportSortAttribute {
	return []ExportSortAttribute{
		"LastUpdatedDateTime",
	}
}

type ExportStatus string

// Enum values for ExportStatus
const (
	ExportStatusInProgress ExportStatus = "InProgress"
	ExportStatusCompleted  ExportStatus = "Completed"
	ExportStatusFailed     ExportStatus = "Failed"
	ExportStatusDeleting   ExportStatus = "Deleting"
)

// Values returns all known values for ExportStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ExportStatus) Values() []ExportStatus {
	return []ExportStatus{
		"InProgress",
		"Completed",
		"Failed",
		"Deleting",
	}
}

type ImportExportFileFormat string

// Enum values for ImportExportFileFormat
const (
	ImportExportFileFormatLexJson ImportExportFileFormat = "LexJson"
	ImportExportFileFormatTsv     ImportExportFileFormat = "TSV"
)

// Values returns all known values for ImportExportFileFormat. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImportExportFileFormat) Values() []ImportExportFileFormat {
	return []ImportExportFileFormat{
		"LexJson",
		"TSV",
	}
}

type ImportFilterName string

// Enum values for ImportFilterName
const (
	ImportFilterNameImportResourceType ImportFilterName = "ImportResourceType"
)

// Values returns all known values for ImportFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImportFilterName) Values() []ImportFilterName {
	return []ImportFilterName{
		"ImportResourceType",
	}
}

type ImportFilterOperator string

// Enum values for ImportFilterOperator
const (
	ImportFilterOperatorContains ImportFilterOperator = "CO"
	ImportFilterOperatorEquals   ImportFilterOperator = "EQ"
)

// Values returns all known values for ImportFilterOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImportFilterOperator) Values() []ImportFilterOperator {
	return []ImportFilterOperator{
		"CO",
		"EQ",
	}
}

type ImportResourceType string

// Enum values for ImportResourceType
const (
	ImportResourceTypeBot              ImportResourceType = "Bot"
	ImportResourceTypeBotLocale        ImportResourceType = "BotLocale"
	ImportResourceTypeCustomVocabulary ImportResourceType = "CustomVocabulary"
)

// Values returns all known values for ImportResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImportResourceType) Values() []ImportResourceType {
	return []ImportResourceType{
		"Bot",
		"BotLocale",
		"CustomVocabulary",
	}
}

type ImportSortAttribute string

// Enum values for ImportSortAttribute
const (
	ImportSortAttributeLastUpdatedDateTime ImportSortAttribute = "LastUpdatedDateTime"
)

// Values returns all known values for ImportSortAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImportSortAttribute) Values() []ImportSortAttribute {
	return []ImportSortAttribute{
		"LastUpdatedDateTime",
	}
}

type ImportStatus string

// Enum values for ImportStatus
const (
	ImportStatusInProgress ImportStatus = "InProgress"
	ImportStatusCompleted  ImportStatus = "Completed"
	ImportStatusFailed     ImportStatus = "Failed"
	ImportStatusDeleting   ImportStatus = "Deleting"
)

// Values returns all known values for ImportStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ImportStatus) Values() []ImportStatus {
	return []ImportStatus{
		"InProgress",
		"Completed",
		"Failed",
		"Deleting",
	}
}

type IntentFilterName string

// Enum values for IntentFilterName
const (
	IntentFilterNameIntentName IntentFilterName = "IntentName"
)

// Values returns all known values for IntentFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IntentFilterName) Values() []IntentFilterName {
	return []IntentFilterName{
		"IntentName",
	}
}

type IntentFilterOperator string

// Enum values for IntentFilterOperator
const (
	IntentFilterOperatorContains IntentFilterOperator = "CO"
	IntentFilterOperatorEquals   IntentFilterOperator = "EQ"
)

// Values returns all known values for IntentFilterOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IntentFilterOperator) Values() []IntentFilterOperator {
	return []IntentFilterOperator{
		"CO",
		"EQ",
	}
}

type IntentSortAttribute string

// Enum values for IntentSortAttribute
const (
	IntentSortAttributeIntentName          IntentSortAttribute = "IntentName"
	IntentSortAttributeLastUpdatedDateTime IntentSortAttribute = "LastUpdatedDateTime"
)

// Values returns all known values for IntentSortAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IntentSortAttribute) Values() []IntentSortAttribute {
	return []IntentSortAttribute{
		"IntentName",
		"LastUpdatedDateTime",
	}
}

type MergeStrategy string

// Enum values for MergeStrategy
const (
	MergeStrategyOverwrite      MergeStrategy = "Overwrite"
	MergeStrategyFailOnConflict MergeStrategy = "FailOnConflict"
	MergeStrategyAppend         MergeStrategy = "Append"
)

// Values returns all known values for MergeStrategy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MergeStrategy) Values() []MergeStrategy {
	return []MergeStrategy{
		"Overwrite",
		"FailOnConflict",
		"Append",
	}
}

type MessageSelectionStrategy string

// Enum values for MessageSelectionStrategy
const (
	MessageSelectionStrategyRandom  MessageSelectionStrategy = "Random"
	MessageSelectionStrategyOrdered MessageSelectionStrategy = "Ordered"
)

// Values returns all known values for MessageSelectionStrategy. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MessageSelectionStrategy) Values() []MessageSelectionStrategy {
	return []MessageSelectionStrategy{
		"Random",
		"Ordered",
	}
}

type ObfuscationSettingType string

// Enum values for ObfuscationSettingType
const (
	ObfuscationSettingTypeNone               ObfuscationSettingType = "None"
	ObfuscationSettingTypeDefaultObfuscation ObfuscationSettingType = "DefaultObfuscation"
)

// Values returns all known values for ObfuscationSettingType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ObfuscationSettingType) Values() []ObfuscationSettingType {
	return []ObfuscationSettingType{
		"None",
		"DefaultObfuscation",
	}
}

type SearchOrder string

// Enum values for SearchOrder
const (
	SearchOrderAscending  SearchOrder = "Ascending"
	SearchOrderDescending SearchOrder = "Descending"
)

// Values returns all known values for SearchOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SearchOrder) Values() []SearchOrder {
	return []SearchOrder{
		"Ascending",
		"Descending",
	}
}

type SlotConstraint string

// Enum values for SlotConstraint
const (
	SlotConstraintRequired SlotConstraint = "Required"
	SlotConstraintOptional SlotConstraint = "Optional"
)

// Values returns all known values for SlotConstraint. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SlotConstraint) Values() []SlotConstraint {
	return []SlotConstraint{
		"Required",
		"Optional",
	}
}

type SlotFilterName string

// Enum values for SlotFilterName
const (
	SlotFilterNameSlotName SlotFilterName = "SlotName"
)

// Values returns all known values for SlotFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SlotFilterName) Values() []SlotFilterName {
	return []SlotFilterName{
		"SlotName",
	}
}

type SlotFilterOperator string

// Enum values for SlotFilterOperator
const (
	SlotFilterOperatorContains SlotFilterOperator = "CO"
	SlotFilterOperatorEquals   SlotFilterOperator = "EQ"
)

// Values returns all known values for SlotFilterOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SlotFilterOperator) Values() []SlotFilterOperator {
	return []SlotFilterOperator{
		"CO",
		"EQ",
	}
}

type SlotSortAttribute string

// Enum values for SlotSortAttribute
const (
	SlotSortAttributeSlotName            SlotSortAttribute = "SlotName"
	SlotSortAttributeLastUpdatedDateTime SlotSortAttribute = "LastUpdatedDateTime"
)

// Values returns all known values for SlotSortAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SlotSortAttribute) Values() []SlotSortAttribute {
	return []SlotSortAttribute{
		"SlotName",
		"LastUpdatedDateTime",
	}
}

type SlotTypeCategory string

// Enum values for SlotTypeCategory
const (
	SlotTypeCategoryCustom          SlotTypeCategory = "Custom"
	SlotTypeCategoryExtended        SlotTypeCategory = "Extended"
	SlotTypeCategoryExternalGrammar SlotTypeCategory = "ExternalGrammar"
)

// Values returns all known values for SlotTypeCategory. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SlotTypeCategory) Values() []SlotTypeCategory {
	return []SlotTypeCategory{
		"Custom",
		"Extended",
		"ExternalGrammar",
	}
}

type SlotTypeFilterName string

// Enum values for SlotTypeFilterName
const (
	SlotTypeFilterNameSlotTypeName       SlotTypeFilterName = "SlotTypeName"
	SlotTypeFilterNameExternalSourceType SlotTypeFilterName = "ExternalSourceType"
)

// Values returns all known values for SlotTypeFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SlotTypeFilterName) Values() []SlotTypeFilterName {
	return []SlotTypeFilterName{
		"SlotTypeName",
		"ExternalSourceType",
	}
}

type SlotTypeFilterOperator string

// Enum values for SlotTypeFilterOperator
const (
	SlotTypeFilterOperatorContains SlotTypeFilterOperator = "CO"
	SlotTypeFilterOperatorEquals   SlotTypeFilterOperator = "EQ"
)

// Values returns all known values for SlotTypeFilterOperator. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SlotTypeFilterOperator) Values() []SlotTypeFilterOperator {
	return []SlotTypeFilterOperator{
		"CO",
		"EQ",
	}
}

type SlotTypeSortAttribute string

// Enum values for SlotTypeSortAttribute
const (
	SlotTypeSortAttributeSlotTypeName        SlotTypeSortAttribute = "SlotTypeName"
	SlotTypeSortAttributeLastUpdatedDateTime SlotTypeSortAttribute = "LastUpdatedDateTime"
)

// Values returns all known values for SlotTypeSortAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SlotTypeSortAttribute) Values() []SlotTypeSortAttribute {
	return []SlotTypeSortAttribute{
		"SlotTypeName",
		"LastUpdatedDateTime",
	}
}

type SlotValueResolutionStrategy string

// Enum values for SlotValueResolutionStrategy
const (
	SlotValueResolutionStrategyOriginalValue SlotValueResolutionStrategy = "OriginalValue"
	SlotValueResolutionStrategyTopResolution SlotValueResolutionStrategy = "TopResolution"
)

// Values returns all known values for SlotValueResolutionStrategy. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SlotValueResolutionStrategy) Values() []SlotValueResolutionStrategy {
	return []SlotValueResolutionStrategy{
		"OriginalValue",
		"TopResolution",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "Ascending"
	SortOrderDescending SortOrder = "Descending"
)

// Values returns all known values for SortOrder. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"Ascending",
		"Descending",
	}
}

type TimeDimension string

// Enum values for TimeDimension
const (
	TimeDimensionHours TimeDimension = "Hours"
	TimeDimensionDays  TimeDimension = "Days"
	TimeDimensionWeeks TimeDimension = "Weeks"
)

// Values returns all known values for TimeDimension. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TimeDimension) Values() []TimeDimension {
	return []TimeDimension{
		"Hours",
		"Days",
		"Weeks",
	}
}

type TranscriptFormat string

// Enum values for TranscriptFormat
const (
	TranscriptFormatLex TranscriptFormat = "Lex"
)

// Values returns all known values for TranscriptFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TranscriptFormat) Values() []TranscriptFormat {
	return []TranscriptFormat{
		"Lex",
	}
}

type VoiceEngine string

// Enum values for VoiceEngine
const (
	VoiceEngineStandard VoiceEngine = "standard"
	VoiceEngineNeural   VoiceEngine = "neural"
)

// Values returns all known values for VoiceEngine. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (VoiceEngine) Values() []VoiceEngine {
	return []VoiceEngine{
		"standard",
		"neural",
	}
}
