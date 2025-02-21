// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AdTriggersElement string

// Enum values for AdTriggersElement
const (
	AdTriggersElementSpliceInsert                           AdTriggersElement = "SPLICE_INSERT"
	AdTriggersElementBreak                                  AdTriggersElement = "BREAK"
	AdTriggersElementProviderAdvertisement                  AdTriggersElement = "PROVIDER_ADVERTISEMENT"
	AdTriggersElementDistributorAdvertisement               AdTriggersElement = "DISTRIBUTOR_ADVERTISEMENT"
	AdTriggersElementProviderPlacementOpportunity           AdTriggersElement = "PROVIDER_PLACEMENT_OPPORTUNITY"
	AdTriggersElementDistributorPlacementOpportunity        AdTriggersElement = "DISTRIBUTOR_PLACEMENT_OPPORTUNITY"
	AdTriggersElementProviderOverlayPlacementOpportunity    AdTriggersElement = "PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY"
	AdTriggersElementDistributorOverlayPlacementOpportunity AdTriggersElement = "DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY"
)

// Values returns all known values for AdTriggersElement. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AdTriggersElement) Values() []AdTriggersElement {
	return []AdTriggersElement{
		"SPLICE_INSERT",
		"BREAK",
		"PROVIDER_ADVERTISEMENT",
		"DISTRIBUTOR_ADVERTISEMENT",
		"PROVIDER_PLACEMENT_OPPORTUNITY",
		"DISTRIBUTOR_PLACEMENT_OPPORTUNITY",
		"PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY",
		"DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY",
	}
}

type PeriodTriggersElement string

// Enum values for PeriodTriggersElement
const (
	PeriodTriggersElementAds PeriodTriggersElement = "ADS"
)

// Values returns all known values for PeriodTriggersElement. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PeriodTriggersElement) Values() []PeriodTriggersElement {
	return []PeriodTriggersElement{
		"ADS",
	}
}

type AdMarkers string

// Enum values for AdMarkers
const (
	AdMarkersNone           AdMarkers = "NONE"
	AdMarkersScte35Enhanced AdMarkers = "SCTE35_ENHANCED"
	AdMarkersPassthrough    AdMarkers = "PASSTHROUGH"
	AdMarkersDaterange      AdMarkers = "DATERANGE"
)

// Values returns all known values for AdMarkers. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (AdMarkers) Values() []AdMarkers {
	return []AdMarkers{
		"NONE",
		"SCTE35_ENHANCED",
		"PASSTHROUGH",
		"DATERANGE",
	}
}

type AdsOnDeliveryRestrictions string

// Enum values for AdsOnDeliveryRestrictions
const (
	AdsOnDeliveryRestrictionsNone         AdsOnDeliveryRestrictions = "NONE"
	AdsOnDeliveryRestrictionsRestricted   AdsOnDeliveryRestrictions = "RESTRICTED"
	AdsOnDeliveryRestrictionsUnrestricted AdsOnDeliveryRestrictions = "UNRESTRICTED"
	AdsOnDeliveryRestrictionsBoth         AdsOnDeliveryRestrictions = "BOTH"
)

// Values returns all known values for AdsOnDeliveryRestrictions. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (AdsOnDeliveryRestrictions) Values() []AdsOnDeliveryRestrictions {
	return []AdsOnDeliveryRestrictions{
		"NONE",
		"RESTRICTED",
		"UNRESTRICTED",
		"BOTH",
	}
}

type EncryptionMethod string

// Enum values for EncryptionMethod
const (
	EncryptionMethodAes128    EncryptionMethod = "AES_128"
	EncryptionMethodSampleAes EncryptionMethod = "SAMPLE_AES"
)

// Values returns all known values for EncryptionMethod. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionMethod) Values() []EncryptionMethod {
	return []EncryptionMethod{
		"AES_128",
		"SAMPLE_AES",
	}
}

type ManifestLayout string

// Enum values for ManifestLayout
const (
	ManifestLayoutFull    ManifestLayout = "FULL"
	ManifestLayoutCompact ManifestLayout = "COMPACT"
)

// Values returns all known values for ManifestLayout. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ManifestLayout) Values() []ManifestLayout {
	return []ManifestLayout{
		"FULL",
		"COMPACT",
	}
}

type Origination string

// Enum values for Origination
const (
	OriginationAllow Origination = "ALLOW"
	OriginationDeny  Origination = "DENY"
)

// Values returns all known values for Origination. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Origination) Values() []Origination {
	return []Origination{
		"ALLOW",
		"DENY",
	}
}

type PlaylistType string

// Enum values for PlaylistType
const (
	PlaylistTypeNone  PlaylistType = "NONE"
	PlaylistTypeEvent PlaylistType = "EVENT"
	PlaylistTypeVod   PlaylistType = "VOD"
)

// Values returns all known values for PlaylistType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PlaylistType) Values() []PlaylistType {
	return []PlaylistType{
		"NONE",
		"EVENT",
		"VOD",
	}
}

type PresetSpeke20Audio string

// Enum values for PresetSpeke20Audio
const (
	PresetSpeke20AudioPresetAudio1 PresetSpeke20Audio = "PRESET-AUDIO-1"
	PresetSpeke20AudioPresetAudio2 PresetSpeke20Audio = "PRESET-AUDIO-2"
	PresetSpeke20AudioPresetAudio3 PresetSpeke20Audio = "PRESET-AUDIO-3"
	PresetSpeke20AudioShared       PresetSpeke20Audio = "SHARED"
	PresetSpeke20AudioUnencrypted  PresetSpeke20Audio = "UNENCRYPTED"
)

// Values returns all known values for PresetSpeke20Audio. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PresetSpeke20Audio) Values() []PresetSpeke20Audio {
	return []PresetSpeke20Audio{
		"PRESET-AUDIO-1",
		"PRESET-AUDIO-2",
		"PRESET-AUDIO-3",
		"SHARED",
		"UNENCRYPTED",
	}
}

type PresetSpeke20Video string

// Enum values for PresetSpeke20Video
const (
	PresetSpeke20VideoPresetVideo1 PresetSpeke20Video = "PRESET-VIDEO-1"
	PresetSpeke20VideoPresetVideo2 PresetSpeke20Video = "PRESET-VIDEO-2"
	PresetSpeke20VideoPresetVideo3 PresetSpeke20Video = "PRESET-VIDEO-3"
	PresetSpeke20VideoPresetVideo4 PresetSpeke20Video = "PRESET-VIDEO-4"
	PresetSpeke20VideoPresetVideo5 PresetSpeke20Video = "PRESET-VIDEO-5"
	PresetSpeke20VideoPresetVideo6 PresetSpeke20Video = "PRESET-VIDEO-6"
	PresetSpeke20VideoPresetVideo7 PresetSpeke20Video = "PRESET-VIDEO-7"
	PresetSpeke20VideoPresetVideo8 PresetSpeke20Video = "PRESET-VIDEO-8"
	PresetSpeke20VideoShared       PresetSpeke20Video = "SHARED"
	PresetSpeke20VideoUnencrypted  PresetSpeke20Video = "UNENCRYPTED"
)

// Values returns all known values for PresetSpeke20Video. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PresetSpeke20Video) Values() []PresetSpeke20Video {
	return []PresetSpeke20Video{
		"PRESET-VIDEO-1",
		"PRESET-VIDEO-2",
		"PRESET-VIDEO-3",
		"PRESET-VIDEO-4",
		"PRESET-VIDEO-5",
		"PRESET-VIDEO-6",
		"PRESET-VIDEO-7",
		"PRESET-VIDEO-8",
		"SHARED",
		"UNENCRYPTED",
	}
}

type Profile string

// Enum values for Profile
const (
	ProfileNone        Profile = "NONE"
	ProfileHbbtv15     Profile = "HBBTV_1_5"
	ProfileHybridcast  Profile = "HYBRIDCAST"
	ProfileDvbDash2014 Profile = "DVB_DASH_2014"
)

// Values returns all known values for Profile. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Profile) Values() []Profile {
	return []Profile{
		"NONE",
		"HBBTV_1_5",
		"HYBRIDCAST",
		"DVB_DASH_2014",
	}
}

type SegmentTemplateFormat string

// Enum values for SegmentTemplateFormat
const (
	SegmentTemplateFormatNumberWithTimeline SegmentTemplateFormat = "NUMBER_WITH_TIMELINE"
	SegmentTemplateFormatTimeWithTimeline   SegmentTemplateFormat = "TIME_WITH_TIMELINE"
	SegmentTemplateFormatNumberWithDuration SegmentTemplateFormat = "NUMBER_WITH_DURATION"
)

// Values returns all known values for SegmentTemplateFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SegmentTemplateFormat) Values() []SegmentTemplateFormat {
	return []SegmentTemplateFormat{
		"NUMBER_WITH_TIMELINE",
		"TIME_WITH_TIMELINE",
		"NUMBER_WITH_DURATION",
	}
}

type Status string

// Enum values for Status
const (
	StatusInProgress Status = "IN_PROGRESS"
	StatusSucceeded  Status = "SUCCEEDED"
	StatusFailed     Status = "FAILED"
)

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

type StreamOrder string

// Enum values for StreamOrder
const (
	StreamOrderOriginal               StreamOrder = "ORIGINAL"
	StreamOrderVideoBitrateAscending  StreamOrder = "VIDEO_BITRATE_ASCENDING"
	StreamOrderVideoBitrateDescending StreamOrder = "VIDEO_BITRATE_DESCENDING"
)

// Values returns all known values for StreamOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StreamOrder) Values() []StreamOrder {
	return []StreamOrder{
		"ORIGINAL",
		"VIDEO_BITRATE_ASCENDING",
		"VIDEO_BITRATE_DESCENDING",
	}
}

type UtcTiming string

// Enum values for UtcTiming
const (
	UtcTimingNone       UtcTiming = "NONE"
	UtcTimingHttpHead   UtcTiming = "HTTP-HEAD"
	UtcTimingHttpIso    UtcTiming = "HTTP-ISO"
	UtcTimingHttpXsdate UtcTiming = "HTTP-XSDATE"
)

// Values returns all known values for UtcTiming. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (UtcTiming) Values() []UtcTiming {
	return []UtcTiming{
		"NONE",
		"HTTP-HEAD",
		"HTTP-ISO",
		"HTTP-XSDATE",
	}
}
