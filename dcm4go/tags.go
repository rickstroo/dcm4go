package dcm4go

// auto-generated, do not edit

const (
	// CommandGroupLengthTag is tag for Command Group Length
	CommandGroupLengthTag = 0x00000000
	// CommandLengthToEndTag is tag for Command Length to End
	CommandLengthToEndTag = 0x00000001
	// AffectedSOPClassUIDTag is tag for Affected SOP Class UID
	AffectedSOPClassUIDTag = 0x00000002
	// RequestedSOPClassUIDTag is tag for Requested SOP Class UID
	RequestedSOPClassUIDTag = 0x00000003
	// CommandRecognitionCodeTag is tag for Command Recognition Code
	CommandRecognitionCodeTag = 0x00000010
	// CommandFieldTag is tag for Command Field
	CommandFieldTag = 0x00000100
	// MessageIDTag is tag for Message ID
	MessageIDTag = 0x00000110
	// MessageIDBeingRespondedToTag is tag for Message ID Being Responded To
	MessageIDBeingRespondedToTag = 0x00000120
	// InitiatorTag is tag for Initiator
	InitiatorTag = 0x00000200
	// ReceiverTag is tag for Receiver
	ReceiverTag = 0x00000300
	// FindLocationTag is tag for Find Location
	FindLocationTag = 0x00000400
	// MoveDestinationTag is tag for Move Destination
	MoveDestinationTag = 0x00000600
	// PriorityTag is tag for Priority
	PriorityTag = 0x00000700
	// CommandDataSetTypeTag is tag for Command Data Set Type
	CommandDataSetTypeTag = 0x00000800
	// NumberOfMatchesTag is tag for Number of Matches
	NumberOfMatchesTag = 0x00000850
	// ResponseSequenceNumberTag is tag for Response Sequence Number
	ResponseSequenceNumberTag = 0x00000860
	// StatusTag is tag for Status
	StatusTag = 0x00000900
	// OffendingElementTag is tag for Offending Element
	OffendingElementTag = 0x00000901
	// ErrorCommentTag is tag for Error Comment
	ErrorCommentTag = 0x00000902
	// ErrorIDTag is tag for Error ID
	ErrorIDTag = 0x00000903
	// AffectedSOPInstanceUIDTag is tag for Affected SOP Instance UID
	AffectedSOPInstanceUIDTag = 0x00001000
	// RequestedSOPInstanceUIDTag is tag for Requested SOP Instance UID
	RequestedSOPInstanceUIDTag = 0x00001001
	// EventTypeIDTag is tag for Event Type ID
	EventTypeIDTag = 0x00001002
	// AttributeIdentifierListTag is tag for Attribute Identifier List
	AttributeIdentifierListTag = 0x00001005
	// ActionTypeIDTag is tag for Action Type ID
	ActionTypeIDTag = 0x00001008
	// NumberOfRemainingSuboperationsTag is tag for Number of Remaining Sub-operations
	NumberOfRemainingSuboperationsTag = 0x00001020
	// NumberOfCompletedSuboperationsTag is tag for Number of Completed Sub-operations
	NumberOfCompletedSuboperationsTag = 0x00001021
	// NumberOfFailedSuboperationsTag is tag for Number of Failed Sub-operations
	NumberOfFailedSuboperationsTag = 0x00001022
	// NumberOfWarningSuboperationsTag is tag for Number of Warning Sub-operations
	NumberOfWarningSuboperationsTag = 0x00001023
	// MoveOriginatorApplicationEntityTitleTag is tag for Move Originator Application Entity Title
	MoveOriginatorApplicationEntityTitleTag = 0x00001030
	// MoveOriginatorMessageIDTag is tag for Move Originator Message ID
	MoveOriginatorMessageIDTag = 0x00001031
	// DialogReceiverTag is tag for Dialog Receiver
	DialogReceiverTag = 0x00004000
	// TerminalTypeTag is tag for Terminal Type
	TerminalTypeTag = 0x00004010
	// MessageSetIDTag is tag for Message Set ID
	MessageSetIDTag = 0x00005010
	// EndMessageIDTag is tag for End Message ID
	EndMessageIDTag = 0x00005020
	// DisplayFormatTag is tag for Display Format
	DisplayFormatTag = 0x00005110
	// PagePositionIDTag is tag for Page Position ID
	PagePositionIDTag = 0x00005120
	// TextFormatIDTag is tag for Text Format ID
	TextFormatIDTag = 0x00005130
	// NormalReverseTag is tag for Normal/Reverse
	NormalReverseTag = 0x00005140
	// AddGrayScaleTag is tag for Add Gray Scale
	AddGrayScaleTag = 0x00005150
	// BordersTag is tag for Borders
	BordersTag = 0x00005160
	// CopiesTag is tag for Copies
	CopiesTag = 0x00005170
	// CommandMagnificationTypeTag is tag for Command Magnification Type
	CommandMagnificationTypeTag = 0x00005180
	// EraseTag is tag for Erase
	EraseTag = 0x00005190
	// PrintTag is tag for Print
	PrintTag = 0x000051A0
	// OverlaysTag is tag for Overlays
	OverlaysTag = 0x000051B0
	// FileMetaInformationGroupLengthTag is tag for File Meta Information Group Length
	FileMetaInformationGroupLengthTag = 0x00020000
	// FileMetaInformationVersionTag is tag for File Meta Information Version
	FileMetaInformationVersionTag = 0x00020001
	// MediaStorageSOPClassUIDTag is tag for Media Storage SOP Class UID
	MediaStorageSOPClassUIDTag = 0x00020002
	// MediaStorageSOPInstanceUIDTag is tag for Media Storage SOP Instance UID
	MediaStorageSOPInstanceUIDTag = 0x00020003
	// TransferSyntaxUIDTag is tag for Transfer Syntax UID
	TransferSyntaxUIDTag = 0x00020010
	// ImplementationClassUIDTag is tag for Implementation Class UID
	ImplementationClassUIDTag = 0x00020012
	// ImplementationVersionNameTag is tag for Implementation Version Name
	ImplementationVersionNameTag = 0x00020013
	// SourceApplicationEntityTitleTag is tag for Source Application Entity Title
	SourceApplicationEntityTitleTag = 0x00020016
	// SendingApplicationEntityTitleTag is tag for Sending Application Entity Title
	SendingApplicationEntityTitleTag = 0x00020017
	// ReceivingApplicationEntityTitleTag is tag for Receiving Application Entity Title
	ReceivingApplicationEntityTitleTag = 0x00020018
	// SourcePresentationAddressTag is tag for Source Presentation Address
	SourcePresentationAddressTag = 0x00020026
	// SendingPresentationAddressTag is tag for Sending Presentation Address
	SendingPresentationAddressTag = 0x00020027
	// ReceivingPresentationAddressTag is tag for Receiving Presentation Address
	ReceivingPresentationAddressTag = 0x00020028
	// RTVMetaInformationVersionTag is tag for RTV Meta Information Version
	RTVMetaInformationVersionTag = 0x00020031
	// RTVCommunicationSOPClassUIDTag is tag for RTV Communication SOP Class UID
	RTVCommunicationSOPClassUIDTag = 0x00020032
	// RTVCommunicationSOPInstanceUIDTag is tag for RTV Communication SOP Instance UID
	RTVCommunicationSOPInstanceUIDTag = 0x00020033
	// RTVSourceIdentifierTag is tag for RTV Source Identifier
	RTVSourceIdentifierTag = 0x00020035
	// RTVFlowIdentifierTag is tag for RTV Flow Identifier
	RTVFlowIdentifierTag = 0x00020036
	// RTVFlowRTPSamplingRateTag is tag for RTV Flow RTP Sampling Rate
	RTVFlowRTPSamplingRateTag = 0x00020037
	// RTVFlowActualFrameDurationTag is tag for RTV Flow Actual Frame Duration
	RTVFlowActualFrameDurationTag = 0x00020038
	// PrivateInformationCreatorUIDTag is tag for Private Information Creator UID
	PrivateInformationCreatorUIDTag = 0x00020100
	// PrivateInformationTag is tag for Private Information
	PrivateInformationTag = 0x00020102
	// FileSetIDTag is tag for File-set ID
	FileSetIDTag = 0x00041130
	// FileSetDescriptorFileIDTag is tag for File-set Descriptor File ID
	FileSetDescriptorFileIDTag = 0x00041141
	// SpecificCharacterSetOfFileSetDescriptorFileTag is tag for Specific Character Set of File-set Descriptor File
	SpecificCharacterSetOfFileSetDescriptorFileTag = 0x00041142
	// OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntityTag is tag for Offset of the First Directory Record of the Root Directory Entity
	OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntityTag = 0x00041200
	// OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntityTag is tag for Offset of the Last Directory Record of the Root Directory Entity
	OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntityTag = 0x00041202
	// FileSetConsistencyFlagTag is tag for File-set Consistency Flag
	FileSetConsistencyFlagTag = 0x00041212
	// DirectoryRecordSequenceTag is tag for Directory Record Sequence
	DirectoryRecordSequenceTag = 0x00041220
	// OffsetOfTheNextDirectoryRecordTag is tag for Offset of the Next Directory Record
	OffsetOfTheNextDirectoryRecordTag = 0x00041400
	// RecordInUseFlagTag is tag for Record In-use Flag
	RecordInUseFlagTag = 0x00041410
	// OffsetOfReferencedLowerLevelDirectoryEntityTag is tag for Offset of Referenced Lower-Level Directory Entity
	OffsetOfReferencedLowerLevelDirectoryEntityTag = 0x00041420
	// DirectoryRecordTypeTag is tag for Directory Record Type
	DirectoryRecordTypeTag = 0x00041430
	// PrivateRecordUIDTag is tag for Private Record UID
	PrivateRecordUIDTag = 0x00041432
	// ReferencedFileIDTag is tag for Referenced File ID
	ReferencedFileIDTag = 0x00041500
	// MRDRDirectoryRecordOffsetTag is tag for MRDR Directory Record Offset
	MRDRDirectoryRecordOffsetTag = 0x00041504
	// ReferencedSOPClassUIDInFileTag is tag for Referenced SOP Class UID in File
	ReferencedSOPClassUIDInFileTag = 0x00041510
	// ReferencedSOPInstanceUIDInFileTag is tag for Referenced SOP Instance UID in File
	ReferencedSOPInstanceUIDInFileTag = 0x00041511
	// ReferencedTransferSyntaxUIDInFileTag is tag for Referenced Transfer Syntax UID in File
	ReferencedTransferSyntaxUIDInFileTag = 0x00041512
	// ReferencedRelatedGeneralSOPClassUIDInFileTag is tag for Referenced Related General SOP Class UID in File
	ReferencedRelatedGeneralSOPClassUIDInFileTag = 0x0004151A
	// NumberOfReferencesTag is tag for Number of References
	NumberOfReferencesTag = 0x00041600
	// LengthToEndTag is tag for Length to End
	LengthToEndTag = 0x00080001
	// SpecificCharacterSetTag is tag for Specific Character Set
	SpecificCharacterSetTag = 0x00080005
	// LanguageCodeSequenceTag is tag for Language Code Sequence
	LanguageCodeSequenceTag = 0x00080006
	// ImageTypeTag is tag for Image Type
	ImageTypeTag = 0x00080008
	// RecognitionCodeTag is tag for Recognition Code
	RecognitionCodeTag = 0x00080010
	// InstanceCreationDateTag is tag for Instance Creation Date
	InstanceCreationDateTag = 0x00080012
	// InstanceCreationTimeTag is tag for Instance Creation Time
	InstanceCreationTimeTag = 0x00080013
	// InstanceCreatorUIDTag is tag for Instance Creator UID
	InstanceCreatorUIDTag = 0x00080014
	// InstanceCoercionDateTimeTag is tag for Instance Coercion DateTime
	InstanceCoercionDateTimeTag = 0x00080015
	// SOPClassUIDTag is tag for SOP Class UID
	SOPClassUIDTag = 0x00080016
	// SOPInstanceUIDTag is tag for SOP Instance UID
	SOPInstanceUIDTag = 0x00080018
	// RelatedGeneralSOPClassUIDTag is tag for Related General SOP Class UID
	RelatedGeneralSOPClassUIDTag = 0x0008001A
	// OriginalSpecializedSOPClassUIDTag is tag for Original Specialized SOP Class UID
	OriginalSpecializedSOPClassUIDTag = 0x0008001B
	// StudyDateTag is tag for Study Date
	StudyDateTag = 0x00080020
	// SeriesDateTag is tag for Series Date
	SeriesDateTag = 0x00080021
	// AcquisitionDateTag is tag for Acquisition Date
	AcquisitionDateTag = 0x00080022
	// ContentDateTag is tag for Content Date
	ContentDateTag = 0x00080023
	// OverlayDateTag is tag for Overlay Date
	OverlayDateTag = 0x00080024
	// CurveDateTag is tag for Curve Date
	CurveDateTag = 0x00080025
	// AcquisitionDateTimeTag is tag for Acquisition DateTime
	AcquisitionDateTimeTag = 0x0008002A
	// StudyTimeTag is tag for Study Time
	StudyTimeTag = 0x00080030
	// SeriesTimeTag is tag for Series Time
	SeriesTimeTag = 0x00080031
	// AcquisitionTimeTag is tag for Acquisition Time
	AcquisitionTimeTag = 0x00080032
	// ContentTimeTag is tag for Content Time
	ContentTimeTag = 0x00080033
	// OverlayTimeTag is tag for Overlay Time
	OverlayTimeTag = 0x00080034
	// CurveTimeTag is tag for Curve Time
	CurveTimeTag = 0x00080035
	// DataSetTypeTag is tag for Data Set Type
	DataSetTypeTag = 0x00080040
	// DataSetSubtypeTag is tag for Data Set Subtype
	DataSetSubtypeTag = 0x00080041
	// NuclearMedicineSeriesTypeTag is tag for Nuclear Medicine Series Type
	NuclearMedicineSeriesTypeTag = 0x00080042
	// AccessionNumberTag is tag for Accession Number
	AccessionNumberTag = 0x00080050
	// IssuerOfAccessionNumberSequenceTag is tag for Issuer of Accession Number Sequence
	IssuerOfAccessionNumberSequenceTag = 0x00080051
	// QueryRetrieveLevelTag is tag for Query/Retrieve Level
	QueryRetrieveLevelTag = 0x00080052
	// QueryRetrieveViewTag is tag for Query/Retrieve View
	QueryRetrieveViewTag = 0x00080053
	// RetrieveAETitleTag is tag for Retrieve AE Title
	RetrieveAETitleTag = 0x00080054
	// StationAETitleTag is tag for Station  AE Title
	StationAETitleTag = 0x00080055
	// InstanceAvailabilityTag is tag for Instance Availability
	InstanceAvailabilityTag = 0x00080056
	// FailedSOPInstanceUIDListTag is tag for Failed SOP Instance UID List
	FailedSOPInstanceUIDListTag = 0x00080058
	// ModalityTag is tag for Modality
	ModalityTag = 0x00080060
	// ModalitiesInStudyTag is tag for Modalities in Study
	ModalitiesInStudyTag = 0x00080061
	// SOPClassesInStudyTag is tag for SOP Classes in Study
	SOPClassesInStudyTag = 0x00080062
	// AnatomicRegionsInStudyCodeSequenceTag is tag for Anatomic Regions in Study Code Sequence
	AnatomicRegionsInStudyCodeSequenceTag = 0x00080063
	// ConversionTypeTag is tag for Conversion Type
	ConversionTypeTag = 0x00080064
	// PresentationIntentTypeTag is tag for Presentation Intent Type
	PresentationIntentTypeTag = 0x00080068
	// ManufacturerTag is tag for Manufacturer
	ManufacturerTag = 0x00080070
	// InstitutionNameTag is tag for Institution Name
	InstitutionNameTag = 0x00080080
	// InstitutionAddressTag is tag for Institution Address
	InstitutionAddressTag = 0x00080081
	// InstitutionCodeSequenceTag is tag for Institution Code Sequence
	InstitutionCodeSequenceTag = 0x00080082
	// ReferringPhysicianNameTag is tag for Referring Physician's Name
	ReferringPhysicianNameTag = 0x00080090
	// ReferringPhysicianAddressTag is tag for Referring Physician's Address
	ReferringPhysicianAddressTag = 0x00080092
	// ReferringPhysicianTelephoneNumbersTag is tag for Referring Physician's Telephone Numbers
	ReferringPhysicianTelephoneNumbersTag = 0x00080094
	// ReferringPhysicianIdentificationSequenceTag is tag for Referring Physician Identification Sequence
	ReferringPhysicianIdentificationSequenceTag = 0x00080096
	// ConsultingPhysicianNameTag is tag for Consulting Physician's Name
	ConsultingPhysicianNameTag = 0x0008009C
	// ConsultingPhysicianIdentificationSequenceTag is tag for Consulting Physician Identification Sequence
	ConsultingPhysicianIdentificationSequenceTag = 0x0008009D
	// CodeValueTag is tag for Code Value
	CodeValueTag = 0x00080100
	// ExtendedCodeValueTag is tag for Extended Code Value
	ExtendedCodeValueTag = 0x00080101
	// CodingSchemeDesignatorTag is tag for Coding Scheme Designator
	CodingSchemeDesignatorTag = 0x00080102
	// CodingSchemeVersionTag is tag for Coding Scheme Version
	CodingSchemeVersionTag = 0x00080103
	// CodeMeaningTag is tag for Code Meaning
	CodeMeaningTag = 0x00080104
	// MappingResourceTag is tag for Mapping Resource
	MappingResourceTag = 0x00080105
	// ContextGroupVersionTag is tag for Context Group Version
	ContextGroupVersionTag = 0x00080106
	// ContextGroupLocalVersionTag is tag for Context Group Local Version
	ContextGroupLocalVersionTag = 0x00080107
	// ExtendedCodeMeaningTag is tag for Extended Code Meaning
	ExtendedCodeMeaningTag = 0x00080108
	// CodingSchemeResourcesSequenceTag is tag for Coding Scheme Resources Sequence
	CodingSchemeResourcesSequenceTag = 0x00080109
	// CodingSchemeURLTypeTag is tag for Coding Scheme URL Type
	CodingSchemeURLTypeTag = 0x0008010A
	// ContextGroupExtensionFlagTag is tag for Context Group Extension Flag
	ContextGroupExtensionFlagTag = 0x0008010B
	// CodingSchemeUIDTag is tag for Coding Scheme UID
	CodingSchemeUIDTag = 0x0008010C
	// ContextGroupExtensionCreatorUIDTag is tag for Context Group Extension Creator UID
	ContextGroupExtensionCreatorUIDTag = 0x0008010D
	// CodingSchemeURLTag is tag for Coding Scheme URL
	CodingSchemeURLTag = 0x0008010E
	// ContextIdentifierTag is tag for Context Identifier
	ContextIdentifierTag = 0x0008010F
	// CodingSchemeIdentificationSequenceTag is tag for Coding Scheme Identification Sequence
	CodingSchemeIdentificationSequenceTag = 0x00080110
	// CodingSchemeRegistryTag is tag for Coding Scheme Registry
	CodingSchemeRegistryTag = 0x00080112
	// CodingSchemeExternalIDTag is tag for Coding Scheme External ID
	CodingSchemeExternalIDTag = 0x00080114
	// CodingSchemeNameTag is tag for Coding Scheme Name
	CodingSchemeNameTag = 0x00080115
	// CodingSchemeResponsibleOrganizationTag is tag for Coding Scheme Responsible Organization
	CodingSchemeResponsibleOrganizationTag = 0x00080116
	// ContextUIDTag is tag for Context UID
	ContextUIDTag = 0x00080117
	// MappingResourceUIDTag is tag for Mapping Resource UID
	MappingResourceUIDTag = 0x00080118
	// LongCodeValueTag is tag for Long Code Value
	LongCodeValueTag = 0x00080119
	// URNCodeValueTag is tag for URN Code Value
	URNCodeValueTag = 0x00080120
	// EquivalentCodeSequenceTag is tag for Equivalent Code Sequence
	EquivalentCodeSequenceTag = 0x00080121
	// MappingResourceNameTag is tag for Mapping Resource Name
	MappingResourceNameTag = 0x00080122
	// ContextGroupIdentificationSequenceTag is tag for Context Group Identification Sequence
	ContextGroupIdentificationSequenceTag = 0x00080123
	// MappingResourceIdentificationSequenceTag is tag for Mapping Resource Identification Sequence
	MappingResourceIdentificationSequenceTag = 0x00080124
	// TimezoneOffsetFromUTCTag is tag for Timezone Offset From UTC
	TimezoneOffsetFromUTCTag = 0x00080201
	// ResponsibleGroupCodeSequenceTag is tag for Responsible Group Code Sequence
	ResponsibleGroupCodeSequenceTag = 0x00080220
	// EquipmentModalityTag is tag for Equipment Modality
	EquipmentModalityTag = 0x00080221
	// ManufacturerRelatedModelGroupTag is tag for Manufacturer's Related Model Group
	ManufacturerRelatedModelGroupTag = 0x00080222
	// PrivateDataElementCharacteristicsSequenceTag is tag for Private Data Element Characteristics Sequence
	PrivateDataElementCharacteristicsSequenceTag = 0x00080300
	// PrivateGroupReferenceTag is tag for Private Group Reference
	PrivateGroupReferenceTag = 0x00080301
	// PrivateCreatorReferenceTag is tag for Private Creator Reference
	PrivateCreatorReferenceTag = 0x00080302
	// BlockIdentifyingInformationStatusTag is tag for Block Identifying Information Status
	BlockIdentifyingInformationStatusTag = 0x00080303
	// NonidentifyingPrivateElementsTag is tag for Nonidentifying Private Elements
	NonidentifyingPrivateElementsTag = 0x00080304
	// DeidentificationActionSequenceTag is tag for Deidentification Action Sequence
	DeidentificationActionSequenceTag = 0x00080305
	// IdentifyingPrivateElementsTag is tag for Identifying Private Elements
	IdentifyingPrivateElementsTag = 0x00080306
	// DeidentificationActionTag is tag for Deidentification Action
	DeidentificationActionTag = 0x00080307
	// PrivateDataElementTag is tag for Private Data Element
	PrivateDataElementTag = 0x00080308
	// PrivateDataElementValueMultiplicityTag is tag for Private Data Element Value Multiplicity
	PrivateDataElementValueMultiplicityTag = 0x00080309
	// PrivateDataElementValueRepresentationTag is tag for Private Data Element Value Representation
	PrivateDataElementValueRepresentationTag = 0x0008030A
	// PrivateDataElementNumberOfItemsTag is tag for Private Data Element Number of Items
	PrivateDataElementNumberOfItemsTag = 0x0008030B
	// PrivateDataElementNameTag is tag for Private Data Element Name
	PrivateDataElementNameTag = 0x0008030C
	// PrivateDataElementKeywordTag is tag for Private Data Element Keyword
	PrivateDataElementKeywordTag = 0x0008030D
	// PrivateDataElementDescriptionTag is tag for Private Data Element Description
	PrivateDataElementDescriptionTag = 0x0008030E
	// PrivateDataElementEncodingTag is tag for Private Data Element Encoding
	PrivateDataElementEncodingTag = 0x0008030F
	// PrivateDataElementDefinitionSequenceTag is tag for Private Data Element Definition Sequence
	PrivateDataElementDefinitionSequenceTag = 0x00080310
	// NetworkIDTag is tag for Network ID
	NetworkIDTag = 0x00081000
	// StationNameTag is tag for Station Name
	StationNameTag = 0x00081010
	// StudyDescriptionTag is tag for Study Description
	StudyDescriptionTag = 0x00081030
	// ProcedureCodeSequenceTag is tag for Procedure Code Sequence
	ProcedureCodeSequenceTag = 0x00081032
	// SeriesDescriptionTag is tag for Series Description
	SeriesDescriptionTag = 0x0008103E
	// SeriesDescriptionCodeSequenceTag is tag for Series Description Code Sequence
	SeriesDescriptionCodeSequenceTag = 0x0008103F
	// InstitutionalDepartmentNameTag is tag for Institutional Department Name
	InstitutionalDepartmentNameTag = 0x00081040
	// InstitutionalDepartmentTypeCodeSequenceTag is tag for Institutional Department Type Code Sequence
	InstitutionalDepartmentTypeCodeSequenceTag = 0x00081041
	// PhysiciansOfRecordTag is tag for Physician(s) of Record
	PhysiciansOfRecordTag = 0x00081048
	// PhysiciansOfRecordIdentificationSequenceTag is tag for Physician(s) of Record Identification Sequence
	PhysiciansOfRecordIdentificationSequenceTag = 0x00081049
	// PerformingPhysicianNameTag is tag for Performing Physician's Name
	PerformingPhysicianNameTag = 0x00081050
	// PerformingPhysicianIdentificationSequenceTag is tag for Performing Physician Identification Sequence
	PerformingPhysicianIdentificationSequenceTag = 0x00081052
	// NameOfPhysiciansReadingStudyTag is tag for Name of Physician(s) Reading Study
	NameOfPhysiciansReadingStudyTag = 0x00081060
	// PhysiciansReadingStudyIdentificationSequenceTag is tag for Physician(s) Reading Study Identification Sequence
	PhysiciansReadingStudyIdentificationSequenceTag = 0x00081062
	// OperatorsNameTag is tag for Operators' Name
	OperatorsNameTag = 0x00081070
	// OperatorIdentificationSequenceTag is tag for Operator Identification Sequence
	OperatorIdentificationSequenceTag = 0x00081072
	// AdmittingDiagnosesDescriptionTag is tag for Admitting Diagnoses Description
	AdmittingDiagnosesDescriptionTag = 0x00081080
	// AdmittingDiagnosesCodeSequenceTag is tag for Admitting Diagnoses Code Sequence
	AdmittingDiagnosesCodeSequenceTag = 0x00081084
	// ManufacturerModelNameTag is tag for Manufacturer's Model Name
	ManufacturerModelNameTag = 0x00081090
	// ReferencedResultsSequenceTag is tag for Referenced Results Sequence
	ReferencedResultsSequenceTag = 0x00081100
	// ReferencedStudySequenceTag is tag for Referenced Study Sequence
	ReferencedStudySequenceTag = 0x00081110
	// ReferencedPerformedProcedureStepSequenceTag is tag for Referenced Performed Procedure Step Sequence
	ReferencedPerformedProcedureStepSequenceTag = 0x00081111
	// ReferencedSeriesSequenceTag is tag for Referenced Series Sequence
	ReferencedSeriesSequenceTag = 0x00081115
	// ReferencedPatientSequenceTag is tag for Referenced Patient Sequence
	ReferencedPatientSequenceTag = 0x00081120
	// ReferencedVisitSequenceTag is tag for Referenced Visit Sequence
	ReferencedVisitSequenceTag = 0x00081125
	// ReferencedOverlaySequenceTag is tag for Referenced Overlay Sequence
	ReferencedOverlaySequenceTag = 0x00081130
	// ReferencedStereometricInstanceSequenceTag is tag for Referenced Stereometric Instance Sequence
	ReferencedStereometricInstanceSequenceTag = 0x00081134
	// ReferencedWaveformSequenceTag is tag for Referenced Waveform Sequence
	ReferencedWaveformSequenceTag = 0x0008113A
	// ReferencedImageSequenceTag is tag for Referenced Image Sequence
	ReferencedImageSequenceTag = 0x00081140
	// ReferencedCurveSequenceTag is tag for Referenced Curve Sequence
	ReferencedCurveSequenceTag = 0x00081145
	// ReferencedInstanceSequenceTag is tag for Referenced Instance Sequence
	ReferencedInstanceSequenceTag = 0x0008114A
	// ReferencedRealWorldValueMappingInstanceSequenceTag is tag for Referenced Real World Value Mapping Instance Sequence
	ReferencedRealWorldValueMappingInstanceSequenceTag = 0x0008114B
	// ReferencedSOPClassUIDTag is tag for Referenced SOP Class UID
	ReferencedSOPClassUIDTag = 0x00081150
	// ReferencedSOPInstanceUIDTag is tag for Referenced SOP Instance UID
	ReferencedSOPInstanceUIDTag = 0x00081155
	// DefinitionSourceSequenceTag is tag for Definition Source Sequence
	DefinitionSourceSequenceTag = 0x00081156
	// SOPClassesSupportedTag is tag for SOP Classes Supported
	SOPClassesSupportedTag = 0x0008115A
	// ReferencedFrameNumberTag is tag for Referenced Frame Number
	ReferencedFrameNumberTag = 0x00081160
	// SimpleFrameListTag is tag for Simple Frame List
	SimpleFrameListTag = 0x00081161
	// CalculatedFrameListTag is tag for Calculated Frame List
	CalculatedFrameListTag = 0x00081162
	// TimeRangeTag is tag for Time Range
	TimeRangeTag = 0x00081163
	// FrameExtractionSequenceTag is tag for Frame Extraction Sequence
	FrameExtractionSequenceTag = 0x00081164
	// MultiFrameSourceSOPInstanceUIDTag is tag for Multi-frame Source SOP Instance UID
	MultiFrameSourceSOPInstanceUIDTag = 0x00081167
	// RetrieveURLTag is tag for Retrieve URL
	RetrieveURLTag = 0x00081190
	// TransactionUIDTag is tag for Transaction UID
	TransactionUIDTag = 0x00081195
	// WarningReasonTag is tag for Warning Reason
	WarningReasonTag = 0x00081196
	// FailureReasonTag is tag for Failure Reason
	FailureReasonTag = 0x00081197
	// FailedSOPSequenceTag is tag for Failed SOP Sequence
	FailedSOPSequenceTag = 0x00081198
	// ReferencedSOPSequenceTag is tag for Referenced SOP Sequence
	ReferencedSOPSequenceTag = 0x00081199
	// OtherFailuresSequenceTag is tag for Other Failures Sequence
	OtherFailuresSequenceTag = 0x0008119A
	// StudiesContainingOtherReferencedInstancesSequenceTag is tag for Studies Containing Other Referenced Instances Sequence
	StudiesContainingOtherReferencedInstancesSequenceTag = 0x00081200
	// RelatedSeriesSequenceTag is tag for Related Series Sequence
	RelatedSeriesSequenceTag = 0x00081250
	// LossyImageCompressionRetiredTag is tag for Lossy Image Compression (Retired)
	LossyImageCompressionRetiredTag = 0x00082110
	// DerivationDescriptionTag is tag for Derivation Description
	DerivationDescriptionTag = 0x00082111
	// SourceImageSequenceTag is tag for Source Image Sequence
	SourceImageSequenceTag = 0x00082112
	// StageNameTag is tag for Stage Name
	StageNameTag = 0x00082120
	// StageNumberTag is tag for Stage Number
	StageNumberTag = 0x00082122
	// NumberOfStagesTag is tag for Number of Stages
	NumberOfStagesTag = 0x00082124
	// ViewNameTag is tag for View Name
	ViewNameTag = 0x00082127
	// ViewNumberTag is tag for View Number
	ViewNumberTag = 0x00082128
	// NumberOfEventTimersTag is tag for Number of Event Timers
	NumberOfEventTimersTag = 0x00082129
	// NumberOfViewsInStageTag is tag for Number of Views in Stage
	NumberOfViewsInStageTag = 0x0008212A
	// EventElapsedTimesTag is tag for Event Elapsed Time(s)
	EventElapsedTimesTag = 0x00082130
	// EventTimerNamesTag is tag for Event Timer Name(s)
	EventTimerNamesTag = 0x00082132
	// EventTimerSequenceTag is tag for Event Timer Sequence
	EventTimerSequenceTag = 0x00082133
	// EventTimeOffsetTag is tag for Event Time Offset
	EventTimeOffsetTag = 0x00082134
	// EventCodeSequenceTag is tag for Event Code Sequence
	EventCodeSequenceTag = 0x00082135
	// StartTrimTag is tag for Start Trim
	StartTrimTag = 0x00082142
	// StopTrimTag is tag for Stop Trim
	StopTrimTag = 0x00082143
	// RecommendedDisplayFrameRateTag is tag for Recommended Display Frame Rate
	RecommendedDisplayFrameRateTag = 0x00082144
	// TransducerPositionTag is tag for Transducer Position
	TransducerPositionTag = 0x00082200
	// TransducerOrientationTag is tag for Transducer Orientation
	TransducerOrientationTag = 0x00082204
	// AnatomicStructureTag is tag for Anatomic Structure
	AnatomicStructureTag = 0x00082208
	// AnatomicRegionSequenceTag is tag for Anatomic Region Sequence
	AnatomicRegionSequenceTag = 0x00082218
	// AnatomicRegionModifierSequenceTag is tag for Anatomic Region Modifier Sequence
	AnatomicRegionModifierSequenceTag = 0x00082220
	// PrimaryAnatomicStructureSequenceTag is tag for Primary Anatomic Structure Sequence
	PrimaryAnatomicStructureSequenceTag = 0x00082228
	// AnatomicStructureSpaceOrRegionSequenceTag is tag for Anatomic Structure, Space or Region Sequence
	AnatomicStructureSpaceOrRegionSequenceTag = 0x00082229
	// PrimaryAnatomicStructureModifierSequenceTag is tag for Primary Anatomic Structure Modifier Sequence
	PrimaryAnatomicStructureModifierSequenceTag = 0x00082230
	// TransducerPositionSequenceTag is tag for Transducer Position Sequence
	TransducerPositionSequenceTag = 0x00082240
	// TransducerPositionModifierSequenceTag is tag for Transducer Position Modifier Sequence
	TransducerPositionModifierSequenceTag = 0x00082242
	// TransducerOrientationSequenceTag is tag for Transducer Orientation Sequence
	TransducerOrientationSequenceTag = 0x00082244
	// TransducerOrientationModifierSequenceTag is tag for Transducer Orientation Modifier Sequence
	TransducerOrientationModifierSequenceTag = 0x00082246
	// AnatomicStructureSpaceOrRegionCodeSequenceTrialTag is tag for Anatomic Structure Space Or Region Code Sequence (Trial)
	AnatomicStructureSpaceOrRegionCodeSequenceTrialTag = 0x00082251
	// AnatomicPortalOfEntranceCodeSequenceTrialTag is tag for Anatomic Portal Of Entrance Code Sequence (Trial)
	AnatomicPortalOfEntranceCodeSequenceTrialTag = 0x00082253
	// AnatomicApproachDirectionCodeSequenceTrialTag is tag for Anatomic Approach Direction Code Sequence (Trial)
	AnatomicApproachDirectionCodeSequenceTrialTag = 0x00082255
	// AnatomicPerspectiveDescriptionTrialTag is tag for Anatomic Perspective Description (Trial)
	AnatomicPerspectiveDescriptionTrialTag = 0x00082256
	// AnatomicPerspectiveCodeSequenceTrialTag is tag for Anatomic Perspective Code Sequence (Trial)
	AnatomicPerspectiveCodeSequenceTrialTag = 0x00082257
	// AnatomicLocationOfExaminingInstrumentDescriptionTrialTag is tag for Anatomic Location Of Examining Instrument Description (Trial)
	AnatomicLocationOfExaminingInstrumentDescriptionTrialTag = 0x00082258
	// AnatomicLocationOfExaminingInstrumentCodeSequenceTrialTag is tag for Anatomic Location Of Examining Instrument Code Sequence (Trial)
	AnatomicLocationOfExaminingInstrumentCodeSequenceTrialTag = 0x00082259
	// AnatomicStructureSpaceOrRegionModifierCodeSequenceTrialTag is tag for Anatomic Structure Space Or Region Modifier Code Sequence (Trial)
	AnatomicStructureSpaceOrRegionModifierCodeSequenceTrialTag = 0x0008225A
	// OnAxisBackgroundAnatomicStructureCodeSequenceTrialTag is tag for On Axis Background Anatomic Structure Code Sequence (Trial)
	OnAxisBackgroundAnatomicStructureCodeSequenceTrialTag = 0x0008225C
	// AlternateRepresentationSequenceTag is tag for Alternate Representation Sequence
	AlternateRepresentationSequenceTag = 0x00083001
	// AvailableTransferSyntaxUIDTag is tag for Available Transfer Syntax UID
	AvailableTransferSyntaxUIDTag = 0x00083002
	// IrradiationEventUIDTag is tag for Irradiation Event UID
	IrradiationEventUIDTag = 0x00083010
	// SourceIrradiationEventSequenceTag is tag for Source Irradiation Event Sequence
	SourceIrradiationEventSequenceTag = 0x00083011
	// RadiopharmaceuticalAdministrationEventUIDTag is tag for Radiopharmaceutical Administration Event UID
	RadiopharmaceuticalAdministrationEventUIDTag = 0x00083012
	// IdentifyingCommentsTag is tag for Identifying Comments
	IdentifyingCommentsTag = 0x00084000
	// FrameTypeTag is tag for Frame Type
	FrameTypeTag = 0x00089007
	// ReferencedImageEvidenceSequenceTag is tag for Referenced Image Evidence Sequence
	ReferencedImageEvidenceSequenceTag = 0x00089092
	// ReferencedRawDataSequenceTag is tag for Referenced Raw Data Sequence
	ReferencedRawDataSequenceTag = 0x00089121
	// CreatorVersionUIDTag is tag for Creator-Version UID
	CreatorVersionUIDTag = 0x00089123
	// DerivationImageSequenceTag is tag for Derivation Image Sequence
	DerivationImageSequenceTag = 0x00089124
	// SourceImageEvidenceSequenceTag is tag for Source Image Evidence Sequence
	SourceImageEvidenceSequenceTag = 0x00089154
	// PixelPresentationTag is tag for Pixel Presentation
	PixelPresentationTag = 0x00089205
	// VolumetricPropertiesTag is tag for Volumetric Properties
	VolumetricPropertiesTag = 0x00089206
	// VolumeBasedCalculationTechniqueTag is tag for Volume Based Calculation Technique
	VolumeBasedCalculationTechniqueTag = 0x00089207
	// ComplexImageComponentTag is tag for Complex Image Component
	ComplexImageComponentTag = 0x00089208
	// AcquisitionContrastTag is tag for Acquisition Contrast
	AcquisitionContrastTag = 0x00089209
	// DerivationCodeSequenceTag is tag for Derivation Code Sequence
	DerivationCodeSequenceTag = 0x00089215
	// ReferencedPresentationStateSequenceTag is tag for Referenced Presentation State Sequence
	ReferencedPresentationStateSequenceTag = 0x00089237
	// ReferencedOtherPlaneSequenceTag is tag for Referenced Other Plane Sequence
	ReferencedOtherPlaneSequenceTag = 0x00089410
	// FrameDisplaySequenceTag is tag for Frame Display Sequence
	FrameDisplaySequenceTag = 0x00089458
	// RecommendedDisplayFrameRateInFloatTag is tag for Recommended Display Frame Rate in Float
	RecommendedDisplayFrameRateInFloatTag = 0x00089459
	// SkipFrameRangeFlagTag is tag for Skip Frame Range Flag
	SkipFrameRangeFlagTag = 0x00089460
	// PatientNameTag is tag for Patient's Name
	PatientNameTag = 0x00100010
	// PatientIDTag is tag for Patient ID
	PatientIDTag = 0x00100020
	// IssuerOfPatientIDTag is tag for Issuer of Patient ID
	IssuerOfPatientIDTag = 0x00100021
	// TypeOfPatientIDTag is tag for Type of Patient ID
	TypeOfPatientIDTag = 0x00100022
	// IssuerOfPatientIDQualifiersSequenceTag is tag for Issuer of Patient ID Qualifiers Sequence
	IssuerOfPatientIDQualifiersSequenceTag = 0x00100024
	// SourcePatientGroupIdentificationSequenceTag is tag for Source Patient Group Identification Sequence
	SourcePatientGroupIdentificationSequenceTag = 0x00100026
	// GroupOfPatientsIdentificationSequenceTag is tag for Group of Patients Identification Sequence
	GroupOfPatientsIdentificationSequenceTag = 0x00100027
	// SubjectRelativePositionInImageTag is tag for Subject Relative Position in Image
	SubjectRelativePositionInImageTag = 0x00100028
	// PatientBirthDateTag is tag for Patient's Birth Date
	PatientBirthDateTag = 0x00100030
	// PatientBirthTimeTag is tag for Patient's Birth Time
	PatientBirthTimeTag = 0x00100032
	// PatientBirthDateInAlternativeCalendarTag is tag for Patient's Birth Date in Alternative Calendar
	PatientBirthDateInAlternativeCalendarTag = 0x00100033
	// PatientDeathDateInAlternativeCalendarTag is tag for Patient's Death Date in Alternative Calendar
	PatientDeathDateInAlternativeCalendarTag = 0x00100034
	// PatientAlternativeCalendarTag is tag for Patient's Alternative Calendar
	PatientAlternativeCalendarTag = 0x00100035
	// PatientSexTag is tag for Patient's Sex
	PatientSexTag = 0x00100040
	// PatientInsurancePlanCodeSequenceTag is tag for Patient's Insurance Plan Code Sequence
	PatientInsurancePlanCodeSequenceTag = 0x00100050
	// PatientPrimaryLanguageCodeSequenceTag is tag for Patient's Primary Language Code Sequence
	PatientPrimaryLanguageCodeSequenceTag = 0x00100101
	// PatientPrimaryLanguageModifierCodeSequenceTag is tag for Patient's Primary Language Modifier Code Sequence
	PatientPrimaryLanguageModifierCodeSequenceTag = 0x00100102
	// QualityControlSubjectTag is tag for Quality Control Subject
	QualityControlSubjectTag = 0x00100200
	// QualityControlSubjectTypeCodeSequenceTag is tag for Quality Control Subject Type Code Sequence
	QualityControlSubjectTypeCodeSequenceTag = 0x00100201
	// StrainDescriptionTag is tag for Strain Description
	StrainDescriptionTag = 0x00100212
	// StrainNomenclatureTag is tag for Strain Nomenclature
	StrainNomenclatureTag = 0x00100213
	// StrainStockNumberTag is tag for Strain Stock Number
	StrainStockNumberTag = 0x00100214
	// StrainSourceRegistryCodeSequenceTag is tag for Strain Source Registry Code Sequence
	StrainSourceRegistryCodeSequenceTag = 0x00100215
	// StrainStockSequenceTag is tag for Strain Stock Sequence
	StrainStockSequenceTag = 0x00100216
	// StrainSourceTag is tag for Strain Source
	StrainSourceTag = 0x00100217
	// StrainAdditionalInformationTag is tag for Strain Additional Information
	StrainAdditionalInformationTag = 0x00100218
	// StrainCodeSequenceTag is tag for Strain Code Sequence
	StrainCodeSequenceTag = 0x00100219
	// GeneticModificationsSequenceTag is tag for Genetic Modifications Sequence
	GeneticModificationsSequenceTag = 0x00100221
	// GeneticModificationsDescriptionTag is tag for Genetic Modifications Description
	GeneticModificationsDescriptionTag = 0x00100222
	// GeneticModificationsNomenclatureTag is tag for Genetic Modifications Nomenclature
	GeneticModificationsNomenclatureTag = 0x00100223
	// GeneticModificationsCodeSequenceTag is tag for Genetic Modifications Code Sequence
	GeneticModificationsCodeSequenceTag = 0x00100229
	// OtherPatientIDsTag is tag for Other Patient IDs
	OtherPatientIDsTag = 0x00101000
	// OtherPatientNamesTag is tag for Other Patient Names
	OtherPatientNamesTag = 0x00101001
	// OtherPatientIDsSequenceTag is tag for Other Patient IDs Sequence
	OtherPatientIDsSequenceTag = 0x00101002
	// PatientBirthNameTag is tag for Patient's Birth Name
	PatientBirthNameTag = 0x00101005
	// PatientAgeTag is tag for Patient's Age
	PatientAgeTag = 0x00101010
	// PatientSizeTag is tag for Patient's Size
	PatientSizeTag = 0x00101020
	// PatientSizeCodeSequenceTag is tag for Patient's Size Code Sequence
	PatientSizeCodeSequenceTag = 0x00101021
	// PatientBodyMassIndexTag is tag for Patient's Body Mass Index
	PatientBodyMassIndexTag = 0x00101022
	// MeasuredAPDimensionTag is tag for Measured AP Dimension
	MeasuredAPDimensionTag = 0x00101023
	// MeasuredLateralDimensionTag is tag for Measured Lateral Dimension
	MeasuredLateralDimensionTag = 0x00101024
	// PatientWeightTag is tag for Patient's Weight
	PatientWeightTag = 0x00101030
	// PatientAddressTag is tag for Patient's Address
	PatientAddressTag = 0x00101040
	// InsurancePlanIdentificationTag is tag for Insurance Plan Identification
	InsurancePlanIdentificationTag = 0x00101050
	// PatientMotherBirthNameTag is tag for Patient's Mother's Birth Name
	PatientMotherBirthNameTag = 0x00101060
	// MilitaryRankTag is tag for Military Rank
	MilitaryRankTag = 0x00101080
	// BranchOfServiceTag is tag for Branch of Service
	BranchOfServiceTag = 0x00101081
	// MedicalRecordLocatorTag is tag for Medical Record Locator
	MedicalRecordLocatorTag = 0x00101090
	// ReferencedPatientPhotoSequenceTag is tag for Referenced Patient Photo Sequence
	ReferencedPatientPhotoSequenceTag = 0x00101100
	// MedicalAlertsTag is tag for Medical Alerts
	MedicalAlertsTag = 0x00102000
	// AllergiesTag is tag for Allergies
	AllergiesTag = 0x00102110
	// CountryOfResidenceTag is tag for Country of Residence
	CountryOfResidenceTag = 0x00102150
	// RegionOfResidenceTag is tag for Region of Residence
	RegionOfResidenceTag = 0x00102152
	// PatientTelephoneNumbersTag is tag for Patient's Telephone Numbers
	PatientTelephoneNumbersTag = 0x00102154
	// PatientTelecomInformationTag is tag for Patient's Telecom Information
	PatientTelecomInformationTag = 0x00102155
	// EthnicGroupTag is tag for Ethnic Group
	EthnicGroupTag = 0x00102160
	// OccupationTag is tag for Occupation
	OccupationTag = 0x00102180
	// SmokingStatusTag is tag for Smoking Status
	SmokingStatusTag = 0x001021A0
	// AdditionalPatientHistoryTag is tag for Additional Patient History
	AdditionalPatientHistoryTag = 0x001021B0
	// PregnancyStatusTag is tag for Pregnancy Status
	PregnancyStatusTag = 0x001021C0
	// LastMenstrualDateTag is tag for Last Menstrual Date
	LastMenstrualDateTag = 0x001021D0
	// PatientReligiousPreferenceTag is tag for Patient's Religious Preference
	PatientReligiousPreferenceTag = 0x001021F0
	// PatientSpeciesDescriptionTag is tag for Patient Species Description
	PatientSpeciesDescriptionTag = 0x00102201
	// PatientSpeciesCodeSequenceTag is tag for Patient Species Code Sequence
	PatientSpeciesCodeSequenceTag = 0x00102202
	// PatientSexNeuteredTag is tag for Patient's Sex Neutered
	PatientSexNeuteredTag = 0x00102203
	// AnatomicalOrientationTypeTag is tag for Anatomical Orientation Type
	AnatomicalOrientationTypeTag = 0x00102210
	// PatientBreedDescriptionTag is tag for Patient Breed Description
	PatientBreedDescriptionTag = 0x00102292
	// PatientBreedCodeSequenceTag is tag for Patient Breed Code Sequence
	PatientBreedCodeSequenceTag = 0x00102293
	// BreedRegistrationSequenceTag is tag for Breed Registration Sequence
	BreedRegistrationSequenceTag = 0x00102294
	// BreedRegistrationNumberTag is tag for Breed Registration Number
	BreedRegistrationNumberTag = 0x00102295
	// BreedRegistryCodeSequenceTag is tag for Breed Registry Code Sequence
	BreedRegistryCodeSequenceTag = 0x00102296
	// ResponsiblePersonTag is tag for Responsible Person
	ResponsiblePersonTag = 0x00102297
	// ResponsiblePersonRoleTag is tag for Responsible Person Role
	ResponsiblePersonRoleTag = 0x00102298
	// ResponsibleOrganizationTag is tag for Responsible Organization
	ResponsibleOrganizationTag = 0x00102299
	// PatientCommentsTag is tag for Patient Comments
	PatientCommentsTag = 0x00104000
	// ExaminedBodyThicknessTag is tag for Examined Body Thickness
	ExaminedBodyThicknessTag = 0x00109431
	// ClinicalTrialSponsorNameTag is tag for Clinical Trial Sponsor Name
	ClinicalTrialSponsorNameTag = 0x00120010
	// ClinicalTrialProtocolIDTag is tag for Clinical Trial Protocol ID
	ClinicalTrialProtocolIDTag = 0x00120020
	// ClinicalTrialProtocolNameTag is tag for Clinical Trial Protocol Name
	ClinicalTrialProtocolNameTag = 0x00120021
	// ClinicalTrialSiteIDTag is tag for Clinical Trial Site ID
	ClinicalTrialSiteIDTag = 0x00120030
	// ClinicalTrialSiteNameTag is tag for Clinical Trial Site Name
	ClinicalTrialSiteNameTag = 0x00120031
	// ClinicalTrialSubjectIDTag is tag for Clinical Trial Subject ID
	ClinicalTrialSubjectIDTag = 0x00120040
	// ClinicalTrialSubjectReadingIDTag is tag for Clinical Trial Subject Reading ID
	ClinicalTrialSubjectReadingIDTag = 0x00120042
	// ClinicalTrialTimePointIDTag is tag for Clinical Trial Time Point ID
	ClinicalTrialTimePointIDTag = 0x00120050
	// ClinicalTrialTimePointDescriptionTag is tag for Clinical Trial Time Point Description
	ClinicalTrialTimePointDescriptionTag = 0x00120051
	// LongitudinalTemporalOffsetFromEventTag is tag for Longitudinal Temporal Offset from Event
	LongitudinalTemporalOffsetFromEventTag = 0x00120052
	// LongitudinalTemporalEventTypeTag is tag for Longitudinal Temporal Event Type
	LongitudinalTemporalEventTypeTag = 0x00120053
	// ClinicalTrialCoordinatingCenterNameTag is tag for Clinical Trial Coordinating Center Name
	ClinicalTrialCoordinatingCenterNameTag = 0x00120060
	// PatientIdentityRemovedTag is tag for Patient Identity Removed
	PatientIdentityRemovedTag = 0x00120062
	// DeidentificationMethodTag is tag for De-identification Method
	DeidentificationMethodTag = 0x00120063
	// DeidentificationMethodCodeSequenceTag is tag for De-identification Method Code Sequence
	DeidentificationMethodCodeSequenceTag = 0x00120064
	// ClinicalTrialSeriesIDTag is tag for Clinical Trial Series ID
	ClinicalTrialSeriesIDTag = 0x00120071
	// ClinicalTrialSeriesDescriptionTag is tag for Clinical Trial Series Description
	ClinicalTrialSeriesDescriptionTag = 0x00120072
	// ClinicalTrialProtocolEthicsCommitteeNameTag is tag for Clinical Trial Protocol Ethics Committee Name
	ClinicalTrialProtocolEthicsCommitteeNameTag = 0x00120081
	// ClinicalTrialProtocolEthicsCommitteeApprovalNumberTag is tag for Clinical Trial Protocol Ethics Committee Approval Number
	ClinicalTrialProtocolEthicsCommitteeApprovalNumberTag = 0x00120082
	// ConsentForClinicalTrialUseSequenceTag is tag for Consent for Clinical Trial Use Sequence
	ConsentForClinicalTrialUseSequenceTag = 0x00120083
	// DistributionTypeTag is tag for Distribution Type
	DistributionTypeTag = 0x00120084
	// ConsentForDistributionFlagTag is tag for Consent for Distribution Flag
	ConsentForDistributionFlagTag = 0x00120085
	// EthicsCommitteeApprovalEffectivenessStartDateTag is tag for Ethics Committee Approval Effectiveness Start Date
	EthicsCommitteeApprovalEffectivenessStartDateTag = 0x00120086
	// EthicsCommitteeApprovalEffectivenessEndDateTag is tag for Ethics Committee Approval Effectiveness End Date
	EthicsCommitteeApprovalEffectivenessEndDateTag = 0x00120087
	// CADFileFormatTag is tag for CAD File Format
	CADFileFormatTag = 0x00140023
	// ComponentReferenceSystemTag is tag for Component Reference System
	ComponentReferenceSystemTag = 0x00140024
	// ComponentManufacturingProcedureTag is tag for Component Manufacturing Procedure
	ComponentManufacturingProcedureTag = 0x00140025
	// ComponentManufacturerTag is tag for Component Manufacturer
	ComponentManufacturerTag = 0x00140028
	// MaterialThicknessTag is tag for Material Thickness
	MaterialThicknessTag = 0x00140030
	// MaterialPipeDiameterTag is tag for Material Pipe Diameter
	MaterialPipeDiameterTag = 0x00140032
	// MaterialIsolationDiameterTag is tag for Material Isolation Diameter
	MaterialIsolationDiameterTag = 0x00140034
	// MaterialGradeTag is tag for Material Grade
	MaterialGradeTag = 0x00140042
	// MaterialPropertiesDescriptionTag is tag for Material Properties Description
	MaterialPropertiesDescriptionTag = 0x00140044
	// MaterialPropertiesFileFormatRetiredTag is tag for Material Properties File Format (Retired)
	MaterialPropertiesFileFormatRetiredTag = 0x00140045
	// MaterialNotesTag is tag for Material Notes
	MaterialNotesTag = 0x00140046
	// ComponentShapeTag is tag for Component Shape
	ComponentShapeTag = 0x00140050
	// CurvatureTypeTag is tag for Curvature Type
	CurvatureTypeTag = 0x00140052
	// OuterDiameterTag is tag for Outer Diameter
	OuterDiameterTag = 0x00140054
	// InnerDiameterTag is tag for Inner Diameter
	InnerDiameterTag = 0x00140056
	// ComponentWelderIDsTag is tag for Component Welder IDs
	ComponentWelderIDsTag = 0x00140100
	// SecondaryApprovalStatusTag is tag for Secondary Approval Status
	SecondaryApprovalStatusTag = 0x00140101
	// SecondaryReviewDateTag is tag for Secondary Review Date
	SecondaryReviewDateTag = 0x00140102
	// SecondaryReviewTimeTag is tag for Secondary Review Time
	SecondaryReviewTimeTag = 0x00140103
	// SecondaryReviewerNameTag is tag for Secondary Reviewer Name
	SecondaryReviewerNameTag = 0x00140104
	// RepairIDTag is tag for Repair ID
	RepairIDTag = 0x00140105
	// MultipleComponentApprovalSequenceTag is tag for Multiple Component Approval Sequence
	MultipleComponentApprovalSequenceTag = 0x00140106
	// OtherApprovalStatusTag is tag for Other Approval Status
	OtherApprovalStatusTag = 0x00140107
	// OtherSecondaryApprovalStatusTag is tag for Other Secondary Approval Status
	OtherSecondaryApprovalStatusTag = 0x00140108
	// ActualEnvironmentalConditionsTag is tag for Actual Environmental Conditions
	ActualEnvironmentalConditionsTag = 0x00141010
	// ExpiryDateTag is tag for Expiry Date
	ExpiryDateTag = 0x00141020
	// EnvironmentalConditionsTag is tag for Environmental Conditions
	EnvironmentalConditionsTag = 0x00141040
	// EvaluatorSequenceTag is tag for Evaluator Sequence
	EvaluatorSequenceTag = 0x00142002
	// EvaluatorNumberTag is tag for Evaluator Number
	EvaluatorNumberTag = 0x00142004
	// EvaluatorNameTag is tag for Evaluator Name
	EvaluatorNameTag = 0x00142006
	// EvaluationAttemptTag is tag for Evaluation Attempt
	EvaluationAttemptTag = 0x00142008
	// IndicationSequenceTag is tag for Indication Sequence
	IndicationSequenceTag = 0x00142012
	// IndicationNumberTag is tag for Indication Number
	IndicationNumberTag = 0x00142014
	// IndicationLabelTag is tag for Indication Label
	IndicationLabelTag = 0x00142016
	// IndicationDescriptionTag is tag for Indication Description
	IndicationDescriptionTag = 0x00142018
	// IndicationTypeTag is tag for Indication Type
	IndicationTypeTag = 0x0014201A
	// IndicationDispositionTag is tag for Indication Disposition
	IndicationDispositionTag = 0x0014201C
	// IndicationROISequenceTag is tag for Indication ROI Sequence
	IndicationROISequenceTag = 0x0014201E
	// IndicationPhysicalPropertySequenceTag is tag for Indication Physical Property Sequence
	IndicationPhysicalPropertySequenceTag = 0x00142030
	// PropertyLabelTag is tag for Property Label
	PropertyLabelTag = 0x00142032
	// CoordinateSystemNumberOfAxesTag is tag for Coordinate System Number of Axes
	CoordinateSystemNumberOfAxesTag = 0x00142202
	// CoordinateSystemAxesSequenceTag is tag for Coordinate System Axes Sequence
	CoordinateSystemAxesSequenceTag = 0x00142204
	// CoordinateSystemAxisDescriptionTag is tag for Coordinate System Axis Description
	CoordinateSystemAxisDescriptionTag = 0x00142206
	// CoordinateSystemDataSetMappingTag is tag for Coordinate System Data Set Mapping
	CoordinateSystemDataSetMappingTag = 0x00142208
	// CoordinateSystemAxisNumberTag is tag for Coordinate System Axis Number
	CoordinateSystemAxisNumberTag = 0x0014220A
	// CoordinateSystemAxisTypeTag is tag for Coordinate System Axis Type
	CoordinateSystemAxisTypeTag = 0x0014220C
	// CoordinateSystemAxisUnitsTag is tag for Coordinate System Axis Units
	CoordinateSystemAxisUnitsTag = 0x0014220E
	// CoordinateSystemAxisValuesTag is tag for Coordinate System Axis Values
	CoordinateSystemAxisValuesTag = 0x00142210
	// CoordinateSystemTransformSequenceTag is tag for Coordinate System Transform Sequence
	CoordinateSystemTransformSequenceTag = 0x00142220
	// TransformDescriptionTag is tag for Transform Description
	TransformDescriptionTag = 0x00142222
	// TransformNumberOfAxesTag is tag for Transform Number of Axes
	TransformNumberOfAxesTag = 0x00142224
	// TransformOrderOfAxesTag is tag for Transform Order of Axes
	TransformOrderOfAxesTag = 0x00142226
	// TransformedAxisUnitsTag is tag for Transformed Axis Units
	TransformedAxisUnitsTag = 0x00142228
	// CoordinateSystemTransformRotationAndScaleMatrixTag is tag for Coordinate System Transform Rotation and Scale Matrix
	CoordinateSystemTransformRotationAndScaleMatrixTag = 0x0014222A
	// CoordinateSystemTransformTranslationMatrixTag is tag for Coordinate System Transform Translation Matrix
	CoordinateSystemTransformTranslationMatrixTag = 0x0014222C
	// InternalDetectorFrameTimeTag is tag for Internal Detector Frame Time
	InternalDetectorFrameTimeTag = 0x00143011
	// NumberOfFramesIntegratedTag is tag for Number of Frames Integrated
	NumberOfFramesIntegratedTag = 0x00143012
	// DetectorTemperatureSequenceTag is tag for Detector Temperature Sequence
	DetectorTemperatureSequenceTag = 0x00143020
	// SensorNameTag is tag for Sensor Name
	SensorNameTag = 0x00143022
	// HorizontalOffsetOfSensorTag is tag for Horizontal Offset of Sensor
	HorizontalOffsetOfSensorTag = 0x00143024
	// VerticalOffsetOfSensorTag is tag for Vertical Offset of Sensor
	VerticalOffsetOfSensorTag = 0x00143026
	// SensorTemperatureTag is tag for Sensor Temperature
	SensorTemperatureTag = 0x00143028
	// DarkCurrentSequenceTag is tag for Dark Current Sequence
	DarkCurrentSequenceTag = 0x00143040
	// DarkCurrentCountsTag is tag for Dark Current Counts
	DarkCurrentCountsTag = 0x00143050
	// GainCorrectionReferenceSequenceTag is tag for Gain Correction Reference Sequence
	GainCorrectionReferenceSequenceTag = 0x00143060
	// AirCountsTag is tag for Air Counts
	AirCountsTag = 0x00143070
	// KVUsedInGainCalibrationTag is tag for KV Used in Gain Calibration
	KVUsedInGainCalibrationTag = 0x00143071
	// MAUsedInGainCalibrationTag is tag for MA Used in Gain Calibration
	MAUsedInGainCalibrationTag = 0x00143072
	// NumberOfFramesUsedForIntegrationTag is tag for Number of Frames Used for Integration
	NumberOfFramesUsedForIntegrationTag = 0x00143073
	// FilterMaterialUsedInGainCalibrationTag is tag for Filter Material Used in Gain Calibration
	FilterMaterialUsedInGainCalibrationTag = 0x00143074
	// FilterThicknessUsedInGainCalibrationTag is tag for Filter Thickness Used in Gain Calibration
	FilterThicknessUsedInGainCalibrationTag = 0x00143075
	// DateOfGainCalibrationTag is tag for Date of Gain Calibration
	DateOfGainCalibrationTag = 0x00143076
	// TimeOfGainCalibrationTag is tag for Time of Gain Calibration
	TimeOfGainCalibrationTag = 0x00143077
	// BadPixelImageTag is tag for Bad Pixel Image
	BadPixelImageTag = 0x00143080
	// CalibrationNotesTag is tag for Calibration Notes
	CalibrationNotesTag = 0x00143099
	// PulserEquipmentSequenceTag is tag for Pulser Equipment Sequence
	PulserEquipmentSequenceTag = 0x00144002
	// PulserTypeTag is tag for Pulser Type
	PulserTypeTag = 0x00144004
	// PulserNotesTag is tag for Pulser Notes
	PulserNotesTag = 0x00144006
	// ReceiverEquipmentSequenceTag is tag for Receiver Equipment Sequence
	ReceiverEquipmentSequenceTag = 0x00144008
	// AmplifierTypeTag is tag for Amplifier Type
	AmplifierTypeTag = 0x0014400A
	// ReceiverNotesTag is tag for Receiver Notes
	ReceiverNotesTag = 0x0014400C
	// PreAmplifierEquipmentSequenceTag is tag for Pre-Amplifier Equipment Sequence
	PreAmplifierEquipmentSequenceTag = 0x0014400E
	// PreAmplifierNotesTag is tag for Pre-Amplifier Notes
	PreAmplifierNotesTag = 0x0014400F
	// TransmitTransducerSequenceTag is tag for Transmit Transducer Sequence
	TransmitTransducerSequenceTag = 0x00144010
	// ReceiveTransducerSequenceTag is tag for Receive Transducer Sequence
	ReceiveTransducerSequenceTag = 0x00144011
	// NumberOfElementsTag is tag for Number of Elements
	NumberOfElementsTag = 0x00144012
	// ElementShapeTag is tag for Element Shape
	ElementShapeTag = 0x00144013
	// ElementDimensionATag is tag for Element Dimension A
	ElementDimensionATag = 0x00144014
	// ElementDimensionBTag is tag for Element Dimension B
	ElementDimensionBTag = 0x00144015
	// ElementPitchATag is tag for Element Pitch A
	ElementPitchATag = 0x00144016
	// MeasuredBeamDimensionATag is tag for Measured Beam Dimension A
	MeasuredBeamDimensionATag = 0x00144017
	// MeasuredBeamDimensionBTag is tag for Measured Beam Dimension B
	MeasuredBeamDimensionBTag = 0x00144018
	// LocationOfMeasuredBeamDiameterTag is tag for Location of Measured Beam Diameter
	LocationOfMeasuredBeamDiameterTag = 0x00144019
	// NominalFrequencyTag is tag for Nominal Frequency
	NominalFrequencyTag = 0x0014401A
	// MeasuredCenterFrequencyTag is tag for Measured Center Frequency
	MeasuredCenterFrequencyTag = 0x0014401B
	// MeasuredBandwidthTag is tag for Measured Bandwidth
	MeasuredBandwidthTag = 0x0014401C
	// ElementPitchBTag is tag for Element Pitch B
	ElementPitchBTag = 0x0014401D
	// PulserSettingsSequenceTag is tag for Pulser Settings Sequence
	PulserSettingsSequenceTag = 0x00144020
	// PulseWidthTag is tag for Pulse Width
	PulseWidthTag = 0x00144022
	// ExcitationFrequencyTag is tag for Excitation Frequency
	ExcitationFrequencyTag = 0x00144024
	// ModulationTypeTag is tag for Modulation Type
	ModulationTypeTag = 0x00144026
	// DampingTag is tag for Damping
	DampingTag = 0x00144028
	// ReceiverSettingsSequenceTag is tag for Receiver Settings Sequence
	ReceiverSettingsSequenceTag = 0x00144030
	// AcquiredSoundpathLengthTag is tag for Acquired Soundpath Length
	AcquiredSoundpathLengthTag = 0x00144031
	// AcquisitionCompressionTypeTag is tag for Acquisition Compression Type
	AcquisitionCompressionTypeTag = 0x00144032
	// AcquisitionSampleSizeTag is tag for Acquisition Sample Size
	AcquisitionSampleSizeTag = 0x00144033
	// RectifierSmoothingTag is tag for Rectifier Smoothing
	RectifierSmoothingTag = 0x00144034
	// DACSequenceTag is tag for DAC Sequence
	DACSequenceTag = 0x00144035
	// DACTypeTag is tag for DAC Type
	DACTypeTag = 0x00144036
	// DACGainPointsTag is tag for DAC Gain Points
	DACGainPointsTag = 0x00144038
	// DACTimePointsTag is tag for DAC Time Points
	DACTimePointsTag = 0x0014403A
	// DACAmplitudeTag is tag for DAC Amplitude
	DACAmplitudeTag = 0x0014403C
	// PreAmplifierSettingsSequenceTag is tag for Pre-Amplifier Settings Sequence
	PreAmplifierSettingsSequenceTag = 0x00144040
	// TransmitTransducerSettingsSequenceTag is tag for Transmit Transducer Settings Sequence
	TransmitTransducerSettingsSequenceTag = 0x00144050
	// ReceiveTransducerSettingsSequenceTag is tag for Receive Transducer Settings Sequence
	ReceiveTransducerSettingsSequenceTag = 0x00144051
	// IncidentAngleTag is tag for Incident Angle
	IncidentAngleTag = 0x00144052
	// CouplingTechniqueTag is tag for Coupling Technique
	CouplingTechniqueTag = 0x00144054
	// CouplingMediumTag is tag for Coupling Medium
	CouplingMediumTag = 0x00144056
	// CouplingVelocityTag is tag for Coupling Velocity
	CouplingVelocityTag = 0x00144057
	// ProbeCenterLocationXTag is tag for Probe Center Location X
	ProbeCenterLocationXTag = 0x00144058
	// ProbeCenterLocationZTag is tag for Probe Center Location Z
	ProbeCenterLocationZTag = 0x00144059
	// SoundPathLengthTag is tag for Sound Path Length
	SoundPathLengthTag = 0x0014405A
	// DelayLawIdentifierTag is tag for Delay Law Identifier
	DelayLawIdentifierTag = 0x0014405C
	// GateSettingsSequenceTag is tag for Gate Settings Sequence
	GateSettingsSequenceTag = 0x00144060
	// GateThresholdTag is tag for Gate Threshold
	GateThresholdTag = 0x00144062
	// VelocityOfSoundTag is tag for Velocity of Sound
	VelocityOfSoundTag = 0x00144064
	// CalibrationSettingsSequenceTag is tag for Calibration Settings Sequence
	CalibrationSettingsSequenceTag = 0x00144070
	// CalibrationProcedureTag is tag for Calibration Procedure
	CalibrationProcedureTag = 0x00144072
	// ProcedureVersionTag is tag for Procedure Version
	ProcedureVersionTag = 0x00144074
	// ProcedureCreationDateTag is tag for Procedure Creation Date
	ProcedureCreationDateTag = 0x00144076
	// ProcedureExpirationDateTag is tag for Procedure Expiration Date
	ProcedureExpirationDateTag = 0x00144078
	// ProcedureLastModifiedDateTag is tag for Procedure Last Modified Date
	ProcedureLastModifiedDateTag = 0x0014407A
	// CalibrationTimeTag is tag for Calibration Time
	CalibrationTimeTag = 0x0014407C
	// CalibrationDateTag is tag for Calibration Date
	CalibrationDateTag = 0x0014407E
	// ProbeDriveEquipmentSequenceTag is tag for Probe Drive Equipment Sequence
	ProbeDriveEquipmentSequenceTag = 0x00144080
	// DriveTypeTag is tag for Drive Type
	DriveTypeTag = 0x00144081
	// ProbeDriveNotesTag is tag for Probe Drive Notes
	ProbeDriveNotesTag = 0x00144082
	// DriveProbeSequenceTag is tag for Drive Probe Sequence
	DriveProbeSequenceTag = 0x00144083
	// ProbeInductanceTag is tag for Probe Inductance
	ProbeInductanceTag = 0x00144084
	// ProbeResistanceTag is tag for Probe Resistance
	ProbeResistanceTag = 0x00144085
	// ReceiveProbeSequenceTag is tag for Receive Probe Sequence
	ReceiveProbeSequenceTag = 0x00144086
	// ProbeDriveSettingsSequenceTag is tag for Probe Drive Settings Sequence
	ProbeDriveSettingsSequenceTag = 0x00144087
	// BridgeResistorsTag is tag for Bridge Resistors
	BridgeResistorsTag = 0x00144088
	// ProbeOrientationAngleTag is tag for Probe Orientation Angle
	ProbeOrientationAngleTag = 0x00144089
	// UserSelectedGainYTag is tag for User Selected Gain Y
	UserSelectedGainYTag = 0x0014408B
	// UserSelectedPhaseTag is tag for User Selected Phase
	UserSelectedPhaseTag = 0x0014408C
	// UserSelectedOffsetXTag is tag for User Selected Offset X
	UserSelectedOffsetXTag = 0x0014408D
	// UserSelectedOffsetYTag is tag for User Selected Offset Y
	UserSelectedOffsetYTag = 0x0014408E
	// ChannelSettingsSequenceTag is tag for Channel Settings Sequence
	ChannelSettingsSequenceTag = 0x00144091
	// ChannelThresholdTag is tag for Channel Threshold
	ChannelThresholdTag = 0x00144092
	// ScannerSettingsSequenceTag is tag for Scanner Settings Sequence
	ScannerSettingsSequenceTag = 0x0014409A
	// ScanProcedureTag is tag for Scan Procedure
	ScanProcedureTag = 0x0014409B
	// TranslationRateXTag is tag for Translation Rate X
	TranslationRateXTag = 0x0014409C
	// TranslationRateYTag is tag for Translation Rate Y
	TranslationRateYTag = 0x0014409D
	// ChannelOverlapTag is tag for Channel Overlap
	ChannelOverlapTag = 0x0014409F
	// ImageQualityIndicatorTypeTag is tag for Image Quality Indicator Type
	ImageQualityIndicatorTypeTag = 0x001440A0
	// ImageQualityIndicatorMaterialTag is tag for Image Quality Indicator Material
	ImageQualityIndicatorMaterialTag = 0x001440A1
	// ImageQualityIndicatorSizeTag is tag for Image Quality Indicator Size
	ImageQualityIndicatorSizeTag = 0x001440A2
	// LINACEnergyTag is tag for LINAC Energy
	LINACEnergyTag = 0x00145002
	// LINACOutputTag is tag for LINAC Output
	LINACOutputTag = 0x00145004
	// ActiveApertureTag is tag for Active Aperture
	ActiveApertureTag = 0x00145100
	// TotalApertureTag is tag for Total Aperture
	TotalApertureTag = 0x00145101
	// ApertureElevationTag is tag for Aperture Elevation
	ApertureElevationTag = 0x00145102
	// MainLobeAngleTag is tag for Main Lobe Angle
	MainLobeAngleTag = 0x00145103
	// MainRoofAngleTag is tag for Main Roof Angle
	MainRoofAngleTag = 0x00145104
	// ConnectorTypeTag is tag for Connector Type
	ConnectorTypeTag = 0x00145105
	// WedgeModelNumberTag is tag for Wedge Model Number
	WedgeModelNumberTag = 0x00145106
	// WedgeAngleFloatTag is tag for Wedge Angle Float
	WedgeAngleFloatTag = 0x00145107
	// WedgeRoofAngleTag is tag for Wedge Roof Angle
	WedgeRoofAngleTag = 0x00145108
	// WedgeElement1PositionTag is tag for Wedge Element 1 Position
	WedgeElement1PositionTag = 0x00145109
	// WedgeMaterialVelocityTag is tag for Wedge Material Velocity
	WedgeMaterialVelocityTag = 0x0014510A
	// WedgeMaterialTag is tag for Wedge Material
	WedgeMaterialTag = 0x0014510B
	// WedgeOffsetZTag is tag for Wedge Offset Z
	WedgeOffsetZTag = 0x0014510C
	// WedgeOriginOffsetXTag is tag for Wedge Origin Offset X
	WedgeOriginOffsetXTag = 0x0014510D
	// WedgeTimeDelayTag is tag for Wedge Time Delay
	WedgeTimeDelayTag = 0x0014510E
	// WedgeNameTag is tag for Wedge Name
	WedgeNameTag = 0x0014510F
	// WedgeManufacturerNameTag is tag for Wedge Manufacturer Name
	WedgeManufacturerNameTag = 0x00145110
	// WedgeDescriptionTag is tag for Wedge Description
	WedgeDescriptionTag = 0x00145111
	// NominalBeamAngleTag is tag for Nominal Beam Angle
	NominalBeamAngleTag = 0x00145112
	// WedgeOffsetXTag is tag for Wedge Offset X
	WedgeOffsetXTag = 0x00145113
	// WedgeOffsetYTag is tag for Wedge Offset Y
	WedgeOffsetYTag = 0x00145114
	// WedgeTotalLengthTag is tag for Wedge Total Length
	WedgeTotalLengthTag = 0x00145115
	// WedgeInContactLengthTag is tag for Wedge In Contact Length
	WedgeInContactLengthTag = 0x00145116
	// WedgeFrontGapTag is tag for Wedge Front Gap
	WedgeFrontGapTag = 0x00145117
	// WedgeTotalHeightTag is tag for Wedge Total Height
	WedgeTotalHeightTag = 0x00145118
	// WedgeFrontHeightTag is tag for Wedge Front Height
	WedgeFrontHeightTag = 0x00145119
	// WedgeRearHeightTag is tag for Wedge Rear Height
	WedgeRearHeightTag = 0x0014511A
	// WedgeTotalWidthTag is tag for Wedge Total Width
	WedgeTotalWidthTag = 0x0014511B
	// WedgeInContactWidthTag is tag for Wedge In Contact Width
	WedgeInContactWidthTag = 0x0014511C
	// WedgeChamferHeightTag is tag for Wedge Chamfer Height
	WedgeChamferHeightTag = 0x0014511D
	// WedgeCurveTag is tag for Wedge Curve
	WedgeCurveTag = 0x0014511E
	// RadiusAlongWedgeTag is tag for Radius Along the Wedge
	RadiusAlongWedgeTag = 0x0014511F
	// WhitePointTag is tag for White Point
	WhitePointTag = 0x00160001
	// PrimaryChromaticitiesTag is tag for Primary Chromaticities
	PrimaryChromaticitiesTag = 0x00160002
	// BatteryLevelTag is tag for Battery Level
	BatteryLevelTag = 0x00160003
	// ExposureTimeInSecondsTag is tag for Exposure Time in Seconds
	ExposureTimeInSecondsTag = 0x00160004
	// FNumberTag is tag for F-Number
	FNumberTag = 0x00160005
	// OECFRowsTag is tag for OECF Rows
	OECFRowsTag = 0x00160006
	// OECFColumnsTag is tag for OECF Columns
	OECFColumnsTag = 0x00160007
	// OECFColumnNamesTag is tag for OECF Column Names
	OECFColumnNamesTag = 0x00160008
	// OECFValuesTag is tag for OECF Values
	OECFValuesTag = 0x00160009
	// SpatialFrequencyResponseRowsTag is tag for Spatial Frequency Response Rows
	SpatialFrequencyResponseRowsTag = 0x0016000A
	// SpatialFrequencyResponseColumnsTag is tag for Spatial Frequency Response Columns
	SpatialFrequencyResponseColumnsTag = 0x0016000B
	// SpatialFrequencyResponseColumnNamesTag is tag for Spatial Frequency Response Column Names
	SpatialFrequencyResponseColumnNamesTag = 0x0016000C
	// SpatialFrequencyResponseValuesTag is tag for Spatial Frequency Response Values
	SpatialFrequencyResponseValuesTag = 0x0016000D
	// ColorFilterArrayPatternRowsTag is tag for Color Filter Array Pattern Rows
	ColorFilterArrayPatternRowsTag = 0x0016000E
	// ColorFilterArrayPatternColumnsTag is tag for Color Filter Array Pattern Columns
	ColorFilterArrayPatternColumnsTag = 0x0016000F
	// ColorFilterArrayPatternValuesTag is tag for Color Filter Array Pattern Values
	ColorFilterArrayPatternValuesTag = 0x00160010
	// FlashFiringStatusTag is tag for Flash Firing Status
	FlashFiringStatusTag = 0x00160011
	// FlashReturnStatusTag is tag for Flash Return Status
	FlashReturnStatusTag = 0x00160012
	// FlashModeTag is tag for Flash Mode
	FlashModeTag = 0x00160013
	// FlashFunctionPresentTag is tag for Flash Function Present
	FlashFunctionPresentTag = 0x00160014
	// FlashRedEyeModeTag is tag for Flash Red Eye Mode
	FlashRedEyeModeTag = 0x00160015
	// ExposureProgramTag is tag for Exposure Program
	ExposureProgramTag = 0x00160016
	// SpectralSensitivityTag is tag for Spectral Sensitivity
	SpectralSensitivityTag = 0x00160017
	// PhotographicSensitivityTag is tag for Photographic Sensitivity
	PhotographicSensitivityTag = 0x00160018
	// SelfTimerModeTag is tag for Self Timer Mode
	SelfTimerModeTag = 0x00160019
	// SensitivityTypeTag is tag for Sensitivity Type
	SensitivityTypeTag = 0x0016001A
	// StandardOutputSensitivityTag is tag for Standard Output Sensitivity
	StandardOutputSensitivityTag = 0x0016001B
	// RecommendedExposureIndexTag is tag for Recommended Exposure Index
	RecommendedExposureIndexTag = 0x0016001C
	// ISOSpeedTag is tag for ISO Speed​
	ISOSpeedTag = 0x0016001D
	// ISOSpeedLatitudeyyyTag is tag for ISO Speed​​ Latitude yyy
	ISOSpeedLatitudeyyyTag = 0x0016001E
	// ISOSpeedLatitudezzzTag is tag for ISO Speed​​ Latitude zzz
	ISOSpeedLatitudezzzTag = 0x0016001F
	// EXIFVersionTag is tag for EXIF Version
	EXIFVersionTag = 0x00160020
	// ShutterSpeedValueTag is tag for Shutter Speed Value
	ShutterSpeedValueTag = 0x00160021
	// ApertureValueTag is tag for Aperture Value
	ApertureValueTag = 0x00160022
	// BrightnessValueTag is tag for Brightness Value
	BrightnessValueTag = 0x00160023
	// ExposureBiasValueTag is tag for Exposure Bias Value
	ExposureBiasValueTag = 0x00160024
	// MaxApertureValueTag is tag for Max Aperture Value
	MaxApertureValueTag = 0x00160025
	// SubjectDistanceTag is tag for Subject Distance
	SubjectDistanceTag = 0x00160026
	// MeteringModeTag is tag for Metering Mode
	MeteringModeTag = 0x00160027
	// LightSourceTag is tag for Light Source
	LightSourceTag = 0x00160028
	// FocalLengthTag is tag for Focal Length
	FocalLengthTag = 0x00160029
	// SubjectAreaTag is tag for Subject Area
	SubjectAreaTag = 0x0016002A
	// MakerNoteTag is tag for Maker Note
	MakerNoteTag = 0x0016002B
	// TemperatureTag is tag for Temperature
	TemperatureTag = 0x00160030
	// HumidityTag is tag for Humidity
	HumidityTag = 0x00160031
	// PressureTag is tag for Pressure
	PressureTag = 0x00160032
	// WaterDepthTag is tag for Water Depth
	WaterDepthTag = 0x00160033
	// AccelerationTag is tag for Acceleration
	AccelerationTag = 0x00160034
	// CameraElevationAngleTag is tag for Camera Elevation Angle
	CameraElevationAngleTag = 0x00160035
	// FlashEnergyTag is tag for Flash Energy
	FlashEnergyTag = 0x00160036
	// SubjectLocationTag is tag for Subject Location
	SubjectLocationTag = 0x00160037
	// PhotographicExposureIndexTag is tag for Photographic Exposure Index
	PhotographicExposureIndexTag = 0x00160038
	// SensingMethodTag is tag for Sensing Method
	SensingMethodTag = 0x00160039
	// FileSourceTag is tag for File Source
	FileSourceTag = 0x0016003A
	// SceneTypeTag is tag for Scene Type
	SceneTypeTag = 0x0016003B
	// CustomRenderedTag is tag for Custom Rendered
	CustomRenderedTag = 0x00160041
	// ExposureModeTag is tag for Exposure Mode
	ExposureModeTag = 0x00160042
	// WhiteBalanceTag is tag for White Balance
	WhiteBalanceTag = 0x00160043
	// DigitalZoomRatioTag is tag for Digital Zoom Ratio
	DigitalZoomRatioTag = 0x00160044
	// FocalLengthIn35mmFilmTag is tag for Focal Length In 35mm​ Film
	FocalLengthIn35mmFilmTag = 0x00160045
	// SceneCaptureTypeTag is tag for Scene Capture Type
	SceneCaptureTypeTag = 0x00160046
	// GainControlTag is tag for Gain Control
	GainControlTag = 0x00160047
	// ContrastTag is tag for Contrast
	ContrastTag = 0x00160048
	// SaturationTag is tag for Saturation
	SaturationTag = 0x00160049
	// SharpnessTag is tag for Sharpness
	SharpnessTag = 0x0016004A
	// DeviceSettingDescriptionTag is tag for Device Setting Description
	DeviceSettingDescriptionTag = 0x0016004B
	// SubjectDistanceRangeTag is tag for Subject Distance Range
	SubjectDistanceRangeTag = 0x0016004C
	// CameraOwnerNameTag is tag for Camera Owner Name
	CameraOwnerNameTag = 0x0016004D
	// LensSpecificationTag is tag for Lens Specification
	LensSpecificationTag = 0x0016004E
	// LensMakeTag is tag for Lens Make
	LensMakeTag = 0x0016004F
	// LensModelTag is tag for Lens Model
	LensModelTag = 0x00160050
	// LensSerialNumberTag is tag for Lens Serial Number
	LensSerialNumberTag = 0x00160051
	// InteroperabilityIndexTag is tag for Interoperability Index
	InteroperabilityIndexTag = 0x00160061
	// InteroperabilityVersionTag is tag for Interoperability Version
	InteroperabilityVersionTag = 0x00160062
	// GPSVersionIDTag is tag for GPS Version ID
	GPSVersionIDTag = 0x00160070
	// GPSLatitudeRefTag is tag for GPS Latitude​ Ref
	GPSLatitudeRefTag = 0x00160071
	// GPSLatitudeTag is tag for GPS Latitude​
	GPSLatitudeTag = 0x00160072
	// GPSLongitudeRefTag is tag for GPS Longitude Ref
	GPSLongitudeRefTag = 0x00160073
	// GPSLongitudeTag is tag for GPS Longitude
	GPSLongitudeTag = 0x00160074
	// GPSAltitudeRefTag is tag for GPS Altitude​ Ref
	GPSAltitudeRefTag = 0x00160075
	// GPSAltitudeTag is tag for GPS Altitude​
	GPSAltitudeTag = 0x00160076
	// GPSTimeStampTag is tag for GPS Time​ Stamp
	GPSTimeStampTag = 0x00160077
	// GPSSatellitesTag is tag for GPS Satellites
	GPSSatellitesTag = 0x00160078
	// GPSStatusTag is tag for GPS Status
	GPSStatusTag = 0x00160079
	// GPSMeasureModeTag is tag for GPS Measure ​Mode
	GPSMeasureModeTag = 0x0016007A
	// GPSDOPTag is tag for GPS DOP
	GPSDOPTag = 0x0016007B
	// GPSSpeedRefTag is tag for GPS Speed​ Ref
	GPSSpeedRefTag = 0x0016007C
	// GPSSpeedTag is tag for GPS Speed​
	GPSSpeedTag = 0x0016007D
	// GPSTrackRefTag is tag for GPS Track ​Ref
	GPSTrackRefTag = 0x0016007E
	// GPSTrackTag is tag for GPS Track
	GPSTrackTag = 0x0016007F
	// GPSImgDirectionRefTag is tag for GPS Img​ Direction Ref
	GPSImgDirectionRefTag = 0x00160080
	// GPSImgDirectionTag is tag for GPS Img​ Direction
	GPSImgDirectionTag = 0x00160081
	// GPSMapDatumTag is tag for GPS Map​ Datum
	GPSMapDatumTag = 0x00160082
	// GPSDestLatitudeRefTag is tag for GPS Dest ​Latitude Ref
	GPSDestLatitudeRefTag = 0x00160083
	// GPSDestLatitudeTag is tag for GPS Dest​ Latitude
	GPSDestLatitudeTag = 0x00160084
	// GPSDestLongitudeRefTag is tag for GPS Dest​ Longitude Ref
	GPSDestLongitudeRefTag = 0x00160085
	// GPSDestLongitudeTag is tag for GPS Dest ​Longitude
	GPSDestLongitudeTag = 0x00160086
	// GPSDestBearingRefTag is tag for GPS Dest ​Bearing Ref
	GPSDestBearingRefTag = 0x00160087
	// GPSDestBearingTag is tag for GPS Dest ​Bearing
	GPSDestBearingTag = 0x00160088
	// GPSDestDistanceRefTag is tag for GPS Dest ​Distance Ref
	GPSDestDistanceRefTag = 0x00160089
	// GPSDestDistanceTag is tag for GPS Dest​ Distance
	GPSDestDistanceTag = 0x0016008A
	// GPSProcessingMethodTag is tag for GPS Processing​ Method
	GPSProcessingMethodTag = 0x0016008B
	// GPSAreaInformationTag is tag for GPS Area ​Information
	GPSAreaInformationTag = 0x0016008C
	// GPSDateStampTag is tag for GPS Date​ Stamp
	GPSDateStampTag = 0x0016008D
	// GPSDifferentialTag is tag for GPS Differential
	GPSDifferentialTag = 0x0016008E
	// ContrastBolusAgentTag is tag for Contrast/Bolus Agent
	ContrastBolusAgentTag = 0x00180010
	// ContrastBolusAgentSequenceTag is tag for Contrast/Bolus Agent Sequence
	ContrastBolusAgentSequenceTag = 0x00180012
	// ContrastBolusT1RelaxivityTag is tag for Contrast/Bolus T1 Relaxivity
	ContrastBolusT1RelaxivityTag = 0x00180013
	// ContrastBolusAdministrationRouteSequenceTag is tag for Contrast/Bolus Administration Route Sequence
	ContrastBolusAdministrationRouteSequenceTag = 0x00180014
	// BodyPartExaminedTag is tag for Body Part Examined
	BodyPartExaminedTag = 0x00180015
	// ScanningSequenceTag is tag for Scanning Sequence
	ScanningSequenceTag = 0x00180020
	// SequenceVariantTag is tag for Sequence Variant
	SequenceVariantTag = 0x00180021
	// ScanOptionsTag is tag for Scan Options
	ScanOptionsTag = 0x00180022
	// MRAcquisitionTypeTag is tag for MR Acquisition Type
	MRAcquisitionTypeTag = 0x00180023
	// SequenceNameTag is tag for Sequence Name
	SequenceNameTag = 0x00180024
	// AngioFlagTag is tag for Angio Flag
	AngioFlagTag = 0x00180025
	// InterventionDrugInformationSequenceTag is tag for Intervention Drug Information Sequence
	InterventionDrugInformationSequenceTag = 0x00180026
	// InterventionDrugStopTimeTag is tag for Intervention Drug Stop Time
	InterventionDrugStopTimeTag = 0x00180027
	// InterventionDrugDoseTag is tag for Intervention Drug Dose
	InterventionDrugDoseTag = 0x00180028
	// InterventionDrugCodeSequenceTag is tag for Intervention Drug Code Sequence
	InterventionDrugCodeSequenceTag = 0x00180029
	// AdditionalDrugSequenceTag is tag for Additional Drug Sequence
	AdditionalDrugSequenceTag = 0x0018002A
	// RadionuclideTag is tag for Radionuclide
	RadionuclideTag = 0x00180030
	// RadiopharmaceuticalTag is tag for Radiopharmaceutical
	RadiopharmaceuticalTag = 0x00180031
	// EnergyWindowCenterlineTag is tag for Energy Window Centerline
	EnergyWindowCenterlineTag = 0x00180032
	// EnergyWindowTotalWidthTag is tag for Energy Window Total Width
	EnergyWindowTotalWidthTag = 0x00180033
	// InterventionDrugNameTag is tag for Intervention Drug Name
	InterventionDrugNameTag = 0x00180034
	// InterventionDrugStartTimeTag is tag for Intervention Drug Start Time
	InterventionDrugStartTimeTag = 0x00180035
	// InterventionSequenceTag is tag for Intervention Sequence
	InterventionSequenceTag = 0x00180036
	// TherapyTypeTag is tag for Therapy Type
	TherapyTypeTag = 0x00180037
	// InterventionStatusTag is tag for Intervention Status
	InterventionStatusTag = 0x00180038
	// TherapyDescriptionTag is tag for Therapy Description
	TherapyDescriptionTag = 0x00180039
	// InterventionDescriptionTag is tag for Intervention Description
	InterventionDescriptionTag = 0x0018003A
	// CineRateTag is tag for Cine Rate
	CineRateTag = 0x00180040
	// InitialCineRunStateTag is tag for Initial Cine Run State
	InitialCineRunStateTag = 0x00180042
	// SliceThicknessTag is tag for Slice Thickness
	SliceThicknessTag = 0x00180050
	// KVPTag is tag for KVP
	KVPTag = 0x00180060
	// CountsAccumulatedTag is tag for Counts Accumulated
	CountsAccumulatedTag = 0x00180070
	// AcquisitionTerminationConditionTag is tag for Acquisition Termination Condition
	AcquisitionTerminationConditionTag = 0x00180071
	// EffectiveDurationTag is tag for Effective Duration
	EffectiveDurationTag = 0x00180072
	// AcquisitionStartConditionTag is tag for Acquisition Start Condition
	AcquisitionStartConditionTag = 0x00180073
	// AcquisitionStartConditionDataTag is tag for Acquisition Start Condition Data
	AcquisitionStartConditionDataTag = 0x00180074
	// AcquisitionTerminationConditionDataTag is tag for Acquisition Termination Condition Data
	AcquisitionTerminationConditionDataTag = 0x00180075
	// RepetitionTimeTag is tag for Repetition Time
	RepetitionTimeTag = 0x00180080
	// EchoTimeTag is tag for Echo Time
	EchoTimeTag = 0x00180081
	// InversionTimeTag is tag for Inversion Time
	InversionTimeTag = 0x00180082
	// NumberOfAveragesTag is tag for Number of Averages
	NumberOfAveragesTag = 0x00180083
	// ImagingFrequencyTag is tag for Imaging Frequency
	ImagingFrequencyTag = 0x00180084
	// ImagedNucleusTag is tag for Imaged Nucleus
	ImagedNucleusTag = 0x00180085
	// EchoNumbersTag is tag for Echo Number(s)
	EchoNumbersTag = 0x00180086
	// MagneticFieldStrengthTag is tag for Magnetic Field Strength
	MagneticFieldStrengthTag = 0x00180087
	// SpacingBetweenSlicesTag is tag for Spacing Between Slices
	SpacingBetweenSlicesTag = 0x00180088
	// NumberOfPhaseEncodingStepsTag is tag for Number of Phase Encoding Steps
	NumberOfPhaseEncodingStepsTag = 0x00180089
	// DataCollectionDiameterTag is tag for Data Collection Diameter
	DataCollectionDiameterTag = 0x00180090
	// EchoTrainLengthTag is tag for Echo Train Length
	EchoTrainLengthTag = 0x00180091
	// PercentSamplingTag is tag for Percent Sampling
	PercentSamplingTag = 0x00180093
	// PercentPhaseFieldOfViewTag is tag for Percent Phase Field of View
	PercentPhaseFieldOfViewTag = 0x00180094
	// PixelBandwidthTag is tag for Pixel Bandwidth
	PixelBandwidthTag = 0x00180095
	// DeviceSerialNumberTag is tag for Device Serial Number
	DeviceSerialNumberTag = 0x00181000
	// DeviceUIDTag is tag for Device UID
	DeviceUIDTag = 0x00181002
	// DeviceIDTag is tag for Device ID
	DeviceIDTag = 0x00181003
	// PlateIDTag is tag for Plate ID
	PlateIDTag = 0x00181004
	// GeneratorIDTag is tag for Generator ID
	GeneratorIDTag = 0x00181005
	// GridIDTag is tag for Grid ID
	GridIDTag = 0x00181006
	// CassetteIDTag is tag for Cassette ID
	CassetteIDTag = 0x00181007
	// GantryIDTag is tag for Gantry ID
	GantryIDTag = 0x00181008
	// UniqueDeviceIdentifierTag is tag for Unique Device Identifier
	UniqueDeviceIdentifierTag = 0x00181009
	// UDISequenceTag is tag for UDI Sequence
	UDISequenceTag = 0x0018100A
	// ManufacturerDeviceClassUIDTag is tag for Manufacturer's Device Class UID
	ManufacturerDeviceClassUIDTag = 0x0018100B
	// SecondaryCaptureDeviceIDTag is tag for Secondary Capture Device ID
	SecondaryCaptureDeviceIDTag = 0x00181010
	// HardcopyCreationDeviceIDTag is tag for Hardcopy Creation Device ID
	HardcopyCreationDeviceIDTag = 0x00181011
	// DateOfSecondaryCaptureTag is tag for Date of Secondary Capture
	DateOfSecondaryCaptureTag = 0x00181012
	// TimeOfSecondaryCaptureTag is tag for Time of Secondary Capture
	TimeOfSecondaryCaptureTag = 0x00181014
	// SecondaryCaptureDeviceManufacturerTag is tag for Secondary Capture Device Manufacturer
	SecondaryCaptureDeviceManufacturerTag = 0x00181016
	// HardcopyDeviceManufacturerTag is tag for Hardcopy Device Manufacturer
	HardcopyDeviceManufacturerTag = 0x00181017
	// SecondaryCaptureDeviceManufacturerModelNameTag is tag for Secondary Capture Device Manufacturer's Model Name
	SecondaryCaptureDeviceManufacturerModelNameTag = 0x00181018
	// SecondaryCaptureDeviceSoftwareVersionsTag is tag for Secondary Capture Device Software Versions
	SecondaryCaptureDeviceSoftwareVersionsTag = 0x00181019
	// HardcopyDeviceSoftwareVersionTag is tag for Hardcopy Device Software Version
	HardcopyDeviceSoftwareVersionTag = 0x0018101A
	// HardcopyDeviceManufacturerModelNameTag is tag for Hardcopy Device Manufacturer's Model Name
	HardcopyDeviceManufacturerModelNameTag = 0x0018101B
	// SoftwareVersionsTag is tag for Software Versions
	SoftwareVersionsTag = 0x00181020
	// VideoImageFormatAcquiredTag is tag for Video Image Format Acquired
	VideoImageFormatAcquiredTag = 0x00181022
	// DigitalImageFormatAcquiredTag is tag for Digital Image Format Acquired
	DigitalImageFormatAcquiredTag = 0x00181023
	// ProtocolNameTag is tag for Protocol Name
	ProtocolNameTag = 0x00181030
	// ContrastBolusRouteTag is tag for Contrast/Bolus Route
	ContrastBolusRouteTag = 0x00181040
	// ContrastBolusVolumeTag is tag for Contrast/Bolus Volume
	ContrastBolusVolumeTag = 0x00181041
	// ContrastBolusStartTimeTag is tag for Contrast/Bolus Start Time
	ContrastBolusStartTimeTag = 0x00181042
	// ContrastBolusStopTimeTag is tag for Contrast/Bolus Stop Time
	ContrastBolusStopTimeTag = 0x00181043
	// ContrastBolusTotalDoseTag is tag for Contrast/Bolus Total Dose
	ContrastBolusTotalDoseTag = 0x00181044
	// SyringeCountsTag is tag for Syringe Counts
	SyringeCountsTag = 0x00181045
	// ContrastFlowRateTag is tag for Contrast Flow Rate
	ContrastFlowRateTag = 0x00181046
	// ContrastFlowDurationTag is tag for Contrast Flow Duration
	ContrastFlowDurationTag = 0x00181047
	// ContrastBolusIngredientTag is tag for Contrast/Bolus Ingredient
	ContrastBolusIngredientTag = 0x00181048
	// ContrastBolusIngredientConcentrationTag is tag for Contrast/Bolus Ingredient Concentration
	ContrastBolusIngredientConcentrationTag = 0x00181049
	// SpatialResolutionTag is tag for Spatial Resolution
	SpatialResolutionTag = 0x00181050
	// TriggerTimeTag is tag for Trigger Time
	TriggerTimeTag = 0x00181060
	// TriggerSourceOrTypeTag is tag for Trigger Source or Type
	TriggerSourceOrTypeTag = 0x00181061
	// NominalIntervalTag is tag for Nominal Interval
	NominalIntervalTag = 0x00181062
	// FrameTimeTag is tag for Frame Time
	FrameTimeTag = 0x00181063
	// CardiacFramingTypeTag is tag for Cardiac Framing Type
	CardiacFramingTypeTag = 0x00181064
	// FrameTimeVectorTag is tag for Frame Time Vector
	FrameTimeVectorTag = 0x00181065
	// FrameDelayTag is tag for Frame Delay
	FrameDelayTag = 0x00181066
	// ImageTriggerDelayTag is tag for Image Trigger Delay
	ImageTriggerDelayTag = 0x00181067
	// MultiplexGroupTimeOffsetTag is tag for Multiplex Group Time Offset
	MultiplexGroupTimeOffsetTag = 0x00181068
	// TriggerTimeOffsetTag is tag for Trigger Time Offset
	TriggerTimeOffsetTag = 0x00181069
	// SynchronizationTriggerTag is tag for Synchronization Trigger
	SynchronizationTriggerTag = 0x0018106A
	// SynchronizationChannelTag is tag for Synchronization Channel
	SynchronizationChannelTag = 0x0018106C
	// TriggerSamplePositionTag is tag for Trigger Sample Position
	TriggerSamplePositionTag = 0x0018106E
	// RadiopharmaceuticalRouteTag is tag for Radiopharmaceutical Route
	RadiopharmaceuticalRouteTag = 0x00181070
	// RadiopharmaceuticalVolumeTag is tag for Radiopharmaceutical Volume
	RadiopharmaceuticalVolumeTag = 0x00181071
	// RadiopharmaceuticalStartTimeTag is tag for Radiopharmaceutical Start Time
	RadiopharmaceuticalStartTimeTag = 0x00181072
	// RadiopharmaceuticalStopTimeTag is tag for Radiopharmaceutical Stop Time
	RadiopharmaceuticalStopTimeTag = 0x00181073
	// RadionuclideTotalDoseTag is tag for Radionuclide Total Dose
	RadionuclideTotalDoseTag = 0x00181074
	// RadionuclideHalfLifeTag is tag for Radionuclide Half Life
	RadionuclideHalfLifeTag = 0x00181075
	// RadionuclidePositronFractionTag is tag for Radionuclide Positron Fraction
	RadionuclidePositronFractionTag = 0x00181076
	// RadiopharmaceuticalSpecificActivityTag is tag for Radiopharmaceutical Specific Activity
	RadiopharmaceuticalSpecificActivityTag = 0x00181077
	// RadiopharmaceuticalStartDateTimeTag is tag for Radiopharmaceutical Start DateTime
	RadiopharmaceuticalStartDateTimeTag = 0x00181078
	// RadiopharmaceuticalStopDateTimeTag is tag for Radiopharmaceutical Stop DateTime
	RadiopharmaceuticalStopDateTimeTag = 0x00181079
	// BeatRejectionFlagTag is tag for Beat Rejection Flag
	BeatRejectionFlagTag = 0x00181080
	// LowRRValueTag is tag for Low R-R Value
	LowRRValueTag = 0x00181081
	// HighRRValueTag is tag for High R-R Value
	HighRRValueTag = 0x00181082
	// IntervalsAcquiredTag is tag for Intervals Acquired
	IntervalsAcquiredTag = 0x00181083
	// IntervalsRejectedTag is tag for Intervals Rejected
	IntervalsRejectedTag = 0x00181084
	// PVCRejectionTag is tag for PVC Rejection
	PVCRejectionTag = 0x00181085
	// SkipBeatsTag is tag for Skip Beats
	SkipBeatsTag = 0x00181086
	// HeartRateTag is tag for Heart Rate
	HeartRateTag = 0x00181088
	// CardiacNumberOfImagesTag is tag for Cardiac Number of Images
	CardiacNumberOfImagesTag = 0x00181090
	// TriggerWindowTag is tag for Trigger Window
	TriggerWindowTag = 0x00181094
	// ReconstructionDiameterTag is tag for Reconstruction Diameter
	ReconstructionDiameterTag = 0x00181100
	// DistanceSourceToDetectorTag is tag for Distance Source to Detector
	DistanceSourceToDetectorTag = 0x00181110
	// DistanceSourceToPatientTag is tag for Distance Source to Patient
	DistanceSourceToPatientTag = 0x00181111
	// EstimatedRadiographicMagnificationFactorTag is tag for Estimated Radiographic Magnification Factor
	EstimatedRadiographicMagnificationFactorTag = 0x00181114
	// GantryDetectorTiltTag is tag for Gantry/Detector Tilt
	GantryDetectorTiltTag = 0x00181120
	// GantryDetectorSlewTag is tag for Gantry/Detector Slew
	GantryDetectorSlewTag = 0x00181121
	// TableHeightTag is tag for Table Height
	TableHeightTag = 0x00181130
	// TableTraverseTag is tag for Table Traverse
	TableTraverseTag = 0x00181131
	// TableMotionTag is tag for Table Motion
	TableMotionTag = 0x00181134
	// TableVerticalIncrementTag is tag for Table Vertical Increment
	TableVerticalIncrementTag = 0x00181135
	// TableLateralIncrementTag is tag for Table Lateral Increment
	TableLateralIncrementTag = 0x00181136
	// TableLongitudinalIncrementTag is tag for Table Longitudinal Increment
	TableLongitudinalIncrementTag = 0x00181137
	// TableAngleTag is tag for Table Angle
	TableAngleTag = 0x00181138
	// TableTypeTag is tag for Table Type
	TableTypeTag = 0x0018113A
	// RotationDirectionTag is tag for Rotation Direction
	RotationDirectionTag = 0x00181140
	// AngularPositionTag is tag for Angular Position
	AngularPositionTag = 0x00181141
	// RadialPositionTag is tag for Radial Position
	RadialPositionTag = 0x00181142
	// ScanArcTag is tag for Scan Arc
	ScanArcTag = 0x00181143
	// AngularStepTag is tag for Angular Step
	AngularStepTag = 0x00181144
	// CenterOfRotationOffsetTag is tag for Center of Rotation Offset
	CenterOfRotationOffsetTag = 0x00181145
	// RotationOffsetTag is tag for Rotation Offset
	RotationOffsetTag = 0x00181146
	// FieldOfViewShapeTag is tag for Field of View Shape
	FieldOfViewShapeTag = 0x00181147
	// FieldOfViewDimensionsTag is tag for Field of View Dimension(s)
	FieldOfViewDimensionsTag = 0x00181149
	// ExposureTimeTag is tag for Exposure Time
	ExposureTimeTag = 0x00181150
	// XRayTubeCurrentTag is tag for X-Ray Tube Current
	XRayTubeCurrentTag = 0x00181151
	// ExposureTag is tag for Exposure
	ExposureTag = 0x00181152
	// ExposureInuAsTag is tag for Exposure in µAs
	ExposureInuAsTag = 0x00181153
	// AveragePulseWidthTag is tag for Average Pulse Width
	AveragePulseWidthTag = 0x00181154
	// RadiationSettingTag is tag for Radiation Setting
	RadiationSettingTag = 0x00181155
	// RectificationTypeTag is tag for Rectification Type
	RectificationTypeTag = 0x00181156
	// RadiationModeTag is tag for Radiation Mode
	RadiationModeTag = 0x0018115A
	// ImageAndFluoroscopyAreaDoseProductTag is tag for Image and Fluoroscopy Area Dose Product
	ImageAndFluoroscopyAreaDoseProductTag = 0x0018115E
	// FilterTypeTag is tag for Filter Type
	FilterTypeTag = 0x00181160
	// TypeOfFiltersTag is tag for Type of Filters
	TypeOfFiltersTag = 0x00181161
	// IntensifierSizeTag is tag for Intensifier Size
	IntensifierSizeTag = 0x00181162
	// ImagerPixelSpacingTag is tag for Imager Pixel Spacing
	ImagerPixelSpacingTag = 0x00181164
	// GridTag is tag for Grid
	GridTag = 0x00181166
	// GeneratorPowerTag is tag for Generator Power
	GeneratorPowerTag = 0x00181170
	// CollimatorGridNameTag is tag for Collimator/grid Name
	CollimatorGridNameTag = 0x00181180
	// CollimatorTypeTag is tag for Collimator Type
	CollimatorTypeTag = 0x00181181
	// FocalDistanceTag is tag for Focal Distance
	FocalDistanceTag = 0x00181182
	// XFocusCenterTag is tag for X Focus Center
	XFocusCenterTag = 0x00181183
	// YFocusCenterTag is tag for Y Focus Center
	YFocusCenterTag = 0x00181184
	// FocalSpotsTag is tag for Focal Spot(s)
	FocalSpotsTag = 0x00181190
	// AnodeTargetMaterialTag is tag for Anode Target Material
	AnodeTargetMaterialTag = 0x00181191
	// BodyPartThicknessTag is tag for Body Part Thickness
	BodyPartThicknessTag = 0x001811A0
	// CompressionForceTag is tag for Compression Force
	CompressionForceTag = 0x001811A2
	// CompressionPressureTag is tag for Compression Pressure
	CompressionPressureTag = 0x001811A3
	// PaddleDescriptionTag is tag for Paddle Description
	PaddleDescriptionTag = 0x001811A4
	// CompressionContactAreaTag is tag for Compression Contact Area
	CompressionContactAreaTag = 0x001811A5
	// DateOfLastCalibrationTag is tag for Date of Last Calibration
	DateOfLastCalibrationTag = 0x00181200
	// TimeOfLastCalibrationTag is tag for Time of Last Calibration
	TimeOfLastCalibrationTag = 0x00181201
	// DateTimeOfLastCalibrationTag is tag for DateTime of Last Calibration
	DateTimeOfLastCalibrationTag = 0x00181202
	// ConvolutionKernelTag is tag for Convolution Kernel
	ConvolutionKernelTag = 0x00181210
	// UpperLowerPixelValuesTag is tag for Upper/Lower Pixel Values
	UpperLowerPixelValuesTag = 0x00181240
	// ActualFrameDurationTag is tag for Actual Frame Duration
	ActualFrameDurationTag = 0x00181242
	// CountRateTag is tag for Count Rate
	CountRateTag = 0x00181243
	// PreferredPlaybackSequencingTag is tag for Preferred Playback Sequencing
	PreferredPlaybackSequencingTag = 0x00181244
	// ReceiveCoilNameTag is tag for Receive Coil Name
	ReceiveCoilNameTag = 0x00181250
	// TransmitCoilNameTag is tag for Transmit Coil Name
	TransmitCoilNameTag = 0x00181251
	// PlateTypeTag is tag for Plate Type
	PlateTypeTag = 0x00181260
	// PhosphorTypeTag is tag for Phosphor Type
	PhosphorTypeTag = 0x00181261
	// WaterEquivalentDiameterTag is tag for Water Equivalent Diameter
	WaterEquivalentDiameterTag = 0x00181271
	// WaterEquivalentDiameterCalculationMethodCodeSequenceTag is tag for Water Equivalent Diameter Calculation Method Code Sequence
	WaterEquivalentDiameterCalculationMethodCodeSequenceTag = 0x00181272
	// ScanVelocityTag is tag for Scan Velocity
	ScanVelocityTag = 0x00181300
	// WholeBodyTechniqueTag is tag for Whole Body Technique
	WholeBodyTechniqueTag = 0x00181301
	// ScanLengthTag is tag for Scan Length
	ScanLengthTag = 0x00181302
	// AcquisitionMatrixTag is tag for Acquisition Matrix
	AcquisitionMatrixTag = 0x00181310
	// InPlanePhaseEncodingDirectionTag is tag for In-plane Phase Encoding Direction
	InPlanePhaseEncodingDirectionTag = 0x00181312
	// FlipAngleTag is tag for Flip Angle
	FlipAngleTag = 0x00181314
	// VariableFlipAngleFlagTag is tag for Variable Flip Angle Flag
	VariableFlipAngleFlagTag = 0x00181315
	// SARTag is tag for SAR
	SARTag = 0x00181316
	// dBdtTag is tag for dB/dt
	dBdtTag = 0x00181318
	// B1rmsTag is tag for B1rms
	B1rmsTag = 0x00181320
	// AcquisitionDeviceProcessingDescriptionTag is tag for Acquisition Device Processing Description
	AcquisitionDeviceProcessingDescriptionTag = 0x00181400
	// AcquisitionDeviceProcessingCodeTag is tag for Acquisition Device Processing Code
	AcquisitionDeviceProcessingCodeTag = 0x00181401
	// CassetteOrientationTag is tag for Cassette Orientation
	CassetteOrientationTag = 0x00181402
	// CassetteSizeTag is tag for Cassette Size
	CassetteSizeTag = 0x00181403
	// ExposuresOnPlateTag is tag for Exposures on Plate
	ExposuresOnPlateTag = 0x00181404
	// RelativeXRayExposureTag is tag for Relative X-Ray Exposure
	RelativeXRayExposureTag = 0x00181405
	// ExposureIndexTag is tag for Exposure Index
	ExposureIndexTag = 0x00181411
	// TargetExposureIndexTag is tag for Target Exposure Index
	TargetExposureIndexTag = 0x00181412
	// DeviationIndexTag is tag for Deviation Index
	DeviationIndexTag = 0x00181413
	// ColumnAngulationTag is tag for Column Angulation
	ColumnAngulationTag = 0x00181450
	// TomoLayerHeightTag is tag for Tomo Layer Height
	TomoLayerHeightTag = 0x00181460
	// TomoAngleTag is tag for Tomo Angle
	TomoAngleTag = 0x00181470
	// TomoTimeTag is tag for Tomo Time
	TomoTimeTag = 0x00181480
	// TomoTypeTag is tag for Tomo Type
	TomoTypeTag = 0x00181490
	// TomoClassTag is tag for Tomo Class
	TomoClassTag = 0x00181491
	// NumberOfTomosynthesisSourceImagesTag is tag for Number of Tomosynthesis Source Images
	NumberOfTomosynthesisSourceImagesTag = 0x00181495
	// PositionerMotionTag is tag for Positioner Motion
	PositionerMotionTag = 0x00181500
	// PositionerTypeTag is tag for Positioner Type
	PositionerTypeTag = 0x00181508
	// PositionerPrimaryAngleTag is tag for Positioner Primary Angle
	PositionerPrimaryAngleTag = 0x00181510
	// PositionerSecondaryAngleTag is tag for Positioner Secondary Angle
	PositionerSecondaryAngleTag = 0x00181511
	// PositionerPrimaryAngleIncrementTag is tag for Positioner Primary Angle Increment
	PositionerPrimaryAngleIncrementTag = 0x00181520
	// PositionerSecondaryAngleIncrementTag is tag for Positioner Secondary Angle Increment
	PositionerSecondaryAngleIncrementTag = 0x00181521
	// DetectorPrimaryAngleTag is tag for Detector Primary Angle
	DetectorPrimaryAngleTag = 0x00181530
	// DetectorSecondaryAngleTag is tag for Detector Secondary Angle
	DetectorSecondaryAngleTag = 0x00181531
	// ShutterShapeTag is tag for Shutter Shape
	ShutterShapeTag = 0x00181600
	// ShutterLeftVerticalEdgeTag is tag for Shutter Left Vertical Edge
	ShutterLeftVerticalEdgeTag = 0x00181602
	// ShutterRightVerticalEdgeTag is tag for Shutter Right Vertical Edge
	ShutterRightVerticalEdgeTag = 0x00181604
	// ShutterUpperHorizontalEdgeTag is tag for Shutter Upper Horizontal Edge
	ShutterUpperHorizontalEdgeTag = 0x00181606
	// ShutterLowerHorizontalEdgeTag is tag for Shutter Lower Horizontal Edge
	ShutterLowerHorizontalEdgeTag = 0x00181608
	// CenterOfCircularShutterTag is tag for Center of Circular Shutter
	CenterOfCircularShutterTag = 0x00181610
	// RadiusOfCircularShutterTag is tag for Radius of Circular Shutter
	RadiusOfCircularShutterTag = 0x00181612
	// VerticesOfThePolygonalShutterTag is tag for Vertices of the Polygonal Shutter
	VerticesOfThePolygonalShutterTag = 0x00181620
	// ShutterPresentationValueTag is tag for Shutter Presentation Value
	ShutterPresentationValueTag = 0x00181622
	// ShutterOverlayGroupTag is tag for Shutter Overlay Group
	ShutterOverlayGroupTag = 0x00181623
	// ShutterPresentationColorCIELabValueTag is tag for Shutter Presentation Color CIELab Value
	ShutterPresentationColorCIELabValueTag = 0x00181624
	// OutlineShapeTypeTag is tag for Outline Shape Type
	OutlineShapeTypeTag = 0x00181630
	// OutlineLeftVerticalEdgeTag is tag for Outline Left Vertical Edge
	OutlineLeftVerticalEdgeTag = 0x00181631
	// OutlineRightVerticalEdgeTag is tag for Outline Right Vertical Edge
	OutlineRightVerticalEdgeTag = 0x00181632
	// OutlineUpperHorizontalEdgeTag is tag for Outline Upper Horizontal Edge
	OutlineUpperHorizontalEdgeTag = 0x00181633
	// OutlineLowerHorizontalEdgeTag is tag for Outline Lower Horizontal Edge
	OutlineLowerHorizontalEdgeTag = 0x00181634
	// CenterOfCircularOutlineTag is tag for Center of Circular Outline
	CenterOfCircularOutlineTag = 0x00181635
	// DiameterOfCircularOutlineTag is tag for Diameter of Circular Outline
	DiameterOfCircularOutlineTag = 0x00181636
	// NumberOfPolygonalVerticesTag is tag for Number of Polygonal Vertices
	NumberOfPolygonalVerticesTag = 0x00181637
	// VerticesOfThePolygonalOutlineTag is tag for Vertices of the Polygonal Outline
	VerticesOfThePolygonalOutlineTag = 0x00181638
	// CollimatorShapeTag is tag for Collimator Shape
	CollimatorShapeTag = 0x00181700
	// CollimatorLeftVerticalEdgeTag is tag for Collimator Left Vertical Edge
	CollimatorLeftVerticalEdgeTag = 0x00181702
	// CollimatorRightVerticalEdgeTag is tag for Collimator Right Vertical Edge
	CollimatorRightVerticalEdgeTag = 0x00181704
	// CollimatorUpperHorizontalEdgeTag is tag for Collimator Upper Horizontal Edge
	CollimatorUpperHorizontalEdgeTag = 0x00181706
	// CollimatorLowerHorizontalEdgeTag is tag for Collimator Lower Horizontal Edge
	CollimatorLowerHorizontalEdgeTag = 0x00181708
	// CenterOfCircularCollimatorTag is tag for Center of Circular Collimator
	CenterOfCircularCollimatorTag = 0x00181710
	// RadiusOfCircularCollimatorTag is tag for Radius of Circular Collimator
	RadiusOfCircularCollimatorTag = 0x00181712
	// VerticesOfThePolygonalCollimatorTag is tag for Vertices of the Polygonal Collimator
	VerticesOfThePolygonalCollimatorTag = 0x00181720
	// AcquisitionTimeSynchronizedTag is tag for Acquisition Time Synchronized
	AcquisitionTimeSynchronizedTag = 0x00181800
	// TimeSourceTag is tag for Time Source
	TimeSourceTag = 0x00181801
	// TimeDistributionProtocolTag is tag for Time Distribution Protocol
	TimeDistributionProtocolTag = 0x00181802
	// NTPSourceAddressTag is tag for NTP Source Address
	NTPSourceAddressTag = 0x00181803
	// PageNumberVectorTag is tag for Page Number Vector
	PageNumberVectorTag = 0x00182001
	// FrameLabelVectorTag is tag for Frame Label Vector
	FrameLabelVectorTag = 0x00182002
	// FramePrimaryAngleVectorTag is tag for Frame Primary Angle Vector
	FramePrimaryAngleVectorTag = 0x00182003
	// FrameSecondaryAngleVectorTag is tag for Frame Secondary Angle Vector
	FrameSecondaryAngleVectorTag = 0x00182004
	// SliceLocationVectorTag is tag for Slice Location Vector
	SliceLocationVectorTag = 0x00182005
	// DisplayWindowLabelVectorTag is tag for Display Window Label Vector
	DisplayWindowLabelVectorTag = 0x00182006
	// NominalScannedPixelSpacingTag is tag for Nominal Scanned Pixel Spacing
	NominalScannedPixelSpacingTag = 0x00182010
	// DigitizingDeviceTransportDirectionTag is tag for Digitizing Device Transport Direction
	DigitizingDeviceTransportDirectionTag = 0x00182020
	// RotationOfScannedFilmTag is tag for Rotation of Scanned Film
	RotationOfScannedFilmTag = 0x00182030
	// BiopsyTargetSequenceTag is tag for Biopsy Target Sequence
	BiopsyTargetSequenceTag = 0x00182041
	// TargetUIDTag is tag for Target UID
	TargetUIDTag = 0x00182042
	// LocalizingCursorPositionTag is tag for Localizing Cursor Position
	LocalizingCursorPositionTag = 0x00182043
	// CalculatedTargetPositionTag is tag for Calculated Target Position
	CalculatedTargetPositionTag = 0x00182044
	// TargetLabelTag is tag for Target Label
	TargetLabelTag = 0x00182045
	// DisplayedZValueTag is tag for Displayed Z Value
	DisplayedZValueTag = 0x00182046
	// IVUSAcquisitionTag is tag for IVUS Acquisition
	IVUSAcquisitionTag = 0x00183100
	// IVUSPullbackRateTag is tag for IVUS Pullback Rate
	IVUSPullbackRateTag = 0x00183101
	// IVUSGatedRateTag is tag for IVUS Gated Rate
	IVUSGatedRateTag = 0x00183102
	// IVUSPullbackStartFrameNumberTag is tag for IVUS Pullback Start Frame Number
	IVUSPullbackStartFrameNumberTag = 0x00183103
	// IVUSPullbackStopFrameNumberTag is tag for IVUS Pullback Stop Frame Number
	IVUSPullbackStopFrameNumberTag = 0x00183104
	// LesionNumberTag is tag for Lesion Number
	LesionNumberTag = 0x00183105
	// AcquisitionCommentsTag is tag for Acquisition Comments
	AcquisitionCommentsTag = 0x00184000
	// OutputPowerTag is tag for Output Power
	OutputPowerTag = 0x00185000
	// TransducerDataTag is tag for Transducer Data
	TransducerDataTag = 0x00185010
	// FocusDepthTag is tag for Focus Depth
	FocusDepthTag = 0x00185012
	// ProcessingFunctionTag is tag for Processing Function
	ProcessingFunctionTag = 0x00185020
	// PostprocessingFunctionTag is tag for Postprocessing Function
	PostprocessingFunctionTag = 0x00185021
	// MechanicalIndexTag is tag for Mechanical Index
	MechanicalIndexTag = 0x00185022
	// BoneThermalIndexTag is tag for Bone Thermal Index
	BoneThermalIndexTag = 0x00185024
	// CranialThermalIndexTag is tag for Cranial Thermal Index
	CranialThermalIndexTag = 0x00185026
	// SoftTissueThermalIndexTag is tag for Soft Tissue Thermal Index
	SoftTissueThermalIndexTag = 0x00185027
	// SoftTissueFocusThermalIndexTag is tag for Soft Tissue-focus Thermal Index
	SoftTissueFocusThermalIndexTag = 0x00185028
	// SoftTissueSurfaceThermalIndexTag is tag for Soft Tissue-surface Thermal Index
	SoftTissueSurfaceThermalIndexTag = 0x00185029
	// DynamicRangeTag is tag for Dynamic Range
	DynamicRangeTag = 0x00185030
	// TotalGainTag is tag for Total Gain
	TotalGainTag = 0x00185040
	// DepthOfScanFieldTag is tag for Depth of Scan Field
	DepthOfScanFieldTag = 0x00185050
	// PatientPositionTag is tag for Patient Position
	PatientPositionTag = 0x00185100
	// ViewPositionTag is tag for View Position
	ViewPositionTag = 0x00185101
	// ProjectionEponymousNameCodeSequenceTag is tag for Projection Eponymous Name Code Sequence
	ProjectionEponymousNameCodeSequenceTag = 0x00185104
	// ImageTransformationMatrixTag is tag for Image Transformation Matrix
	ImageTransformationMatrixTag = 0x00185210
	// ImageTranslationVectorTag is tag for Image Translation Vector
	ImageTranslationVectorTag = 0x00185212
	// SensitivityTag is tag for Sensitivity
	SensitivityTag = 0x00186000
	// SequenceOfUltrasoundRegionsTag is tag for Sequence of Ultrasound Regions
	SequenceOfUltrasoundRegionsTag = 0x00186011
	// RegionSpatialFormatTag is tag for Region Spatial Format
	RegionSpatialFormatTag = 0x00186012
	// RegionDataTypeTag is tag for Region Data Type
	RegionDataTypeTag = 0x00186014
	// RegionFlagsTag is tag for Region Flags
	RegionFlagsTag = 0x00186016
	// RegionLocationMinX0Tag is tag for Region Location Min X0
	RegionLocationMinX0Tag = 0x00186018
	// RegionLocationMinY0Tag is tag for Region Location Min Y0
	RegionLocationMinY0Tag = 0x0018601A
	// RegionLocationMaxX1Tag is tag for Region Location Max X1
	RegionLocationMaxX1Tag = 0x0018601C
	// RegionLocationMaxY1Tag is tag for Region Location Max Y1
	RegionLocationMaxY1Tag = 0x0018601E
	// ReferencePixelX0Tag is tag for Reference Pixel X0
	ReferencePixelX0Tag = 0x00186020
	// ReferencePixelY0Tag is tag for Reference Pixel Y0
	ReferencePixelY0Tag = 0x00186022
	// PhysicalUnitsXDirectionTag is tag for Physical Units X Direction
	PhysicalUnitsXDirectionTag = 0x00186024
	// PhysicalUnitsYDirectionTag is tag for Physical Units Y Direction
	PhysicalUnitsYDirectionTag = 0x00186026
	// ReferencePixelPhysicalValueXTag is tag for Reference Pixel Physical Value X
	ReferencePixelPhysicalValueXTag = 0x00186028
	// ReferencePixelPhysicalValueYTag is tag for Reference Pixel Physical Value Y
	ReferencePixelPhysicalValueYTag = 0x0018602A
	// PhysicalDeltaXTag is tag for Physical Delta X
	PhysicalDeltaXTag = 0x0018602C
	// PhysicalDeltaYTag is tag for Physical Delta Y
	PhysicalDeltaYTag = 0x0018602E
	// TransducerFrequencyTag is tag for Transducer Frequency
	TransducerFrequencyTag = 0x00186030
	// TransducerTypeTag is tag for Transducer Type
	TransducerTypeTag = 0x00186031
	// PulseRepetitionFrequencyTag is tag for Pulse Repetition Frequency
	PulseRepetitionFrequencyTag = 0x00186032
	// DopplerCorrectionAngleTag is tag for Doppler Correction Angle
	DopplerCorrectionAngleTag = 0x00186034
	// SteeringAngleTag is tag for Steering Angle
	SteeringAngleTag = 0x00186036
	// DopplerSampleVolumeXPositionRetiredTag is tag for Doppler Sample Volume X Position (Retired)
	DopplerSampleVolumeXPositionRetiredTag = 0x00186038
	// DopplerSampleVolumeXPositionTag is tag for Doppler Sample Volume X Position
	DopplerSampleVolumeXPositionTag = 0x00186039
	// DopplerSampleVolumeYPositionRetiredTag is tag for Doppler Sample Volume Y Position (Retired)
	DopplerSampleVolumeYPositionRetiredTag = 0x0018603A
	// DopplerSampleVolumeYPositionTag is tag for Doppler Sample Volume Y Position
	DopplerSampleVolumeYPositionTag = 0x0018603B
	// TMLinePositionX0RetiredTag is tag for TM-Line Position X0 (Retired)
	TMLinePositionX0RetiredTag = 0x0018603C
	// TMLinePositionX0Tag is tag for TM-Line Position X0
	TMLinePositionX0Tag = 0x0018603D
	// TMLinePositionY0RetiredTag is tag for TM-Line Position Y0 (Retired)
	TMLinePositionY0RetiredTag = 0x0018603E
	// TMLinePositionY0Tag is tag for TM-Line Position Y0
	TMLinePositionY0Tag = 0x0018603F
	// TMLinePositionX1RetiredTag is tag for TM-Line Position X1 (Retired)
	TMLinePositionX1RetiredTag = 0x00186040
	// TMLinePositionX1Tag is tag for TM-Line Position X1
	TMLinePositionX1Tag = 0x00186041
	// TMLinePositionY1RetiredTag is tag for TM-Line Position Y1 (Retired)
	TMLinePositionY1RetiredTag = 0x00186042
	// TMLinePositionY1Tag is tag for TM-Line Position Y1
	TMLinePositionY1Tag = 0x00186043
	// PixelComponentOrganizationTag is tag for Pixel Component Organization
	PixelComponentOrganizationTag = 0x00186044
	// PixelComponentMaskTag is tag for Pixel Component Mask
	PixelComponentMaskTag = 0x00186046
	// PixelComponentRangeStartTag is tag for Pixel Component Range Start
	PixelComponentRangeStartTag = 0x00186048
	// PixelComponentRangeStopTag is tag for Pixel Component Range Stop
	PixelComponentRangeStopTag = 0x0018604A
	// PixelComponentPhysicalUnitsTag is tag for Pixel Component Physical Units
	PixelComponentPhysicalUnitsTag = 0x0018604C
	// PixelComponentDataTypeTag is tag for Pixel Component Data Type
	PixelComponentDataTypeTag = 0x0018604E
	// NumberOfTableBreakPointsTag is tag for Number of Table Break Points
	NumberOfTableBreakPointsTag = 0x00186050
	// TableOfXBreakPointsTag is tag for Table of X Break Points
	TableOfXBreakPointsTag = 0x00186052
	// TableOfYBreakPointsTag is tag for Table of Y Break Points
	TableOfYBreakPointsTag = 0x00186054
	// NumberOfTableEntriesTag is tag for Number of Table Entries
	NumberOfTableEntriesTag = 0x00186056
	// TableOfPixelValuesTag is tag for Table of Pixel Values
	TableOfPixelValuesTag = 0x00186058
	// TableOfParameterValuesTag is tag for Table of Parameter Values
	TableOfParameterValuesTag = 0x0018605A
	// RWaveTimeVectorTag is tag for R Wave Time Vector
	RWaveTimeVectorTag = 0x00186060
	// DetectorConditionsNominalFlagTag is tag for Detector Conditions Nominal Flag
	DetectorConditionsNominalFlagTag = 0x00187000
	// DetectorTemperatureTag is tag for Detector Temperature
	DetectorTemperatureTag = 0x00187001
	// DetectorTypeTag is tag for Detector Type
	DetectorTypeTag = 0x00187004
	// DetectorConfigurationTag is tag for Detector Configuration
	DetectorConfigurationTag = 0x00187005
	// DetectorDescriptionTag is tag for Detector Description
	DetectorDescriptionTag = 0x00187006
	// DetectorModeTag is tag for Detector Mode
	DetectorModeTag = 0x00187008
	// DetectorIDTag is tag for Detector ID
	DetectorIDTag = 0x0018700A
	// DateOfLastDetectorCalibrationTag is tag for Date of Last Detector Calibration
	DateOfLastDetectorCalibrationTag = 0x0018700C
	// TimeOfLastDetectorCalibrationTag is tag for Time of Last Detector Calibration
	TimeOfLastDetectorCalibrationTag = 0x0018700E
	// ExposuresOnDetectorSinceLastCalibrationTag is tag for Exposures on Detector Since Last Calibration
	ExposuresOnDetectorSinceLastCalibrationTag = 0x00187010
	// ExposuresOnDetectorSinceManufacturedTag is tag for Exposures on Detector Since Manufactured
	ExposuresOnDetectorSinceManufacturedTag = 0x00187011
	// DetectorTimeSinceLastExposureTag is tag for Detector Time Since Last Exposure
	DetectorTimeSinceLastExposureTag = 0x00187012
	// DetectorActiveTimeTag is tag for Detector Active Time
	DetectorActiveTimeTag = 0x00187014
	// DetectorActivationOffsetFromExposureTag is tag for Detector Activation Offset From Exposure
	DetectorActivationOffsetFromExposureTag = 0x00187016
	// DetectorBinningTag is tag for Detector Binning
	DetectorBinningTag = 0x0018701A
	// DetectorElementPhysicalSizeTag is tag for Detector Element Physical Size
	DetectorElementPhysicalSizeTag = 0x00187020
	// DetectorElementSpacingTag is tag for Detector Element Spacing
	DetectorElementSpacingTag = 0x00187022
	// DetectorActiveShapeTag is tag for Detector Active Shape
	DetectorActiveShapeTag = 0x00187024
	// DetectorActiveDimensionsTag is tag for Detector Active Dimension(s)
	DetectorActiveDimensionsTag = 0x00187026
	// DetectorActiveOriginTag is tag for Detector Active Origin
	DetectorActiveOriginTag = 0x00187028
	// DetectorManufacturerNameTag is tag for Detector Manufacturer Name
	DetectorManufacturerNameTag = 0x0018702A
	// DetectorManufacturerModelNameTag is tag for Detector Manufacturer's Model Name
	DetectorManufacturerModelNameTag = 0x0018702B
	// FieldOfViewOriginTag is tag for Field of View Origin
	FieldOfViewOriginTag = 0x00187030
	// FieldOfViewRotationTag is tag for Field of View Rotation
	FieldOfViewRotationTag = 0x00187032
	// FieldOfViewHorizontalFlipTag is tag for Field of View Horizontal Flip
	FieldOfViewHorizontalFlipTag = 0x00187034
	// PixelDataAreaOriginRelativeToFOVTag is tag for Pixel Data Area Origin Relative To FOV
	PixelDataAreaOriginRelativeToFOVTag = 0x00187036
	// PixelDataAreaRotationAngleRelativeToFOVTag is tag for Pixel Data Area Rotation Angle Relative To FOV
	PixelDataAreaRotationAngleRelativeToFOVTag = 0x00187038
	// GridAbsorbingMaterialTag is tag for Grid Absorbing Material
	GridAbsorbingMaterialTag = 0x00187040
	// GridSpacingMaterialTag is tag for Grid Spacing Material
	GridSpacingMaterialTag = 0x00187041
	// GridThicknessTag is tag for Grid Thickness
	GridThicknessTag = 0x00187042
	// GridPitchTag is tag for Grid Pitch
	GridPitchTag = 0x00187044
	// GridAspectRatioTag is tag for Grid Aspect Ratio
	GridAspectRatioTag = 0x00187046
	// GridPeriodTag is tag for Grid Period
	GridPeriodTag = 0x00187048
	// GridFocalDistanceTag is tag for Grid Focal Distance
	GridFocalDistanceTag = 0x0018704C
	// FilterMaterialTag is tag for Filter Material
	FilterMaterialTag = 0x00187050
	// FilterThicknessMinimumTag is tag for Filter Thickness Minimum
	FilterThicknessMinimumTag = 0x00187052
	// FilterThicknessMaximumTag is tag for Filter Thickness Maximum
	FilterThicknessMaximumTag = 0x00187054
	// FilterBeamPathLengthMinimumTag is tag for Filter Beam Path Length Minimum
	FilterBeamPathLengthMinimumTag = 0x00187056
	// FilterBeamPathLengthMaximumTag is tag for Filter Beam Path Length Maximum
	FilterBeamPathLengthMaximumTag = 0x00187058
	// ExposureControlModeTag is tag for Exposure Control Mode
	ExposureControlModeTag = 0x00187060
	// ExposureControlModeDescriptionTag is tag for Exposure Control Mode Description
	ExposureControlModeDescriptionTag = 0x00187062
	// ExposureStatusTag is tag for Exposure Status
	ExposureStatusTag = 0x00187064
	// PhototimerSettingTag is tag for Phototimer Setting
	PhototimerSettingTag = 0x00187065
	// ExposureTimeInuSTag is tag for Exposure Time in µS
	ExposureTimeInuSTag = 0x00188150
	// XRayTubeCurrentInuATag is tag for X-Ray Tube Current in µA
	XRayTubeCurrentInuATag = 0x00188151
	// ContentQualificationTag is tag for Content Qualification
	ContentQualificationTag = 0x00189004
	// PulseSequenceNameTag is tag for Pulse Sequence Name
	PulseSequenceNameTag = 0x00189005
	// MRImagingModifierSequenceTag is tag for MR Imaging Modifier Sequence
	MRImagingModifierSequenceTag = 0x00189006
	// EchoPulseSequenceTag is tag for Echo Pulse Sequence
	EchoPulseSequenceTag = 0x00189008
	// InversionRecoveryTag is tag for Inversion Recovery
	InversionRecoveryTag = 0x00189009
	// FlowCompensationTag is tag for Flow Compensation
	FlowCompensationTag = 0x00189010
	// MultipleSpinEchoTag is tag for Multiple Spin Echo
	MultipleSpinEchoTag = 0x00189011
	// MultiPlanarExcitationTag is tag for Multi-planar Excitation
	MultiPlanarExcitationTag = 0x00189012
	// PhaseContrastTag is tag for Phase Contrast
	PhaseContrastTag = 0x00189014
	// TimeOfFlightContrastTag is tag for Time of Flight Contrast
	TimeOfFlightContrastTag = 0x00189015
	// SpoilingTag is tag for Spoiling
	SpoilingTag = 0x00189016
	// SteadyStatePulseSequenceTag is tag for Steady State Pulse Sequence
	SteadyStatePulseSequenceTag = 0x00189017
	// EchoPlanarPulseSequenceTag is tag for Echo Planar Pulse Sequence
	EchoPlanarPulseSequenceTag = 0x00189018
	// TagAngleFirstAxisTag is tag for Tag Angle First Axis
	TagAngleFirstAxisTag = 0x00189019
	// MagnetizationTransferTag is tag for Magnetization Transfer
	MagnetizationTransferTag = 0x00189020
	// T2PreparationTag is tag for T2 Preparation
	T2PreparationTag = 0x00189021
	// BloodSignalNullingTag is tag for Blood Signal Nulling
	BloodSignalNullingTag = 0x00189022
	// SaturationRecoveryTag is tag for Saturation Recovery
	SaturationRecoveryTag = 0x00189024
	// SpectrallySelectedSuppressionTag is tag for Spectrally Selected Suppression
	SpectrallySelectedSuppressionTag = 0x00189025
	// SpectrallySelectedExcitationTag is tag for Spectrally Selected Excitation
	SpectrallySelectedExcitationTag = 0x00189026
	// SpatialPresaturationTag is tag for Spatial Pre-saturation
	SpatialPresaturationTag = 0x00189027
	// TaggingTag is tag for Tagging
	TaggingTag = 0x00189028
	// OversamplingPhaseTag is tag for Oversampling Phase
	OversamplingPhaseTag = 0x00189029
	// TagSpacingFirstDimensionTag is tag for Tag Spacing First Dimension
	TagSpacingFirstDimensionTag = 0x00189030
	// GeometryOfKSpaceTraversalTag is tag for Geometry of k-Space Traversal
	GeometryOfKSpaceTraversalTag = 0x00189032
	// SegmentedKSpaceTraversalTag is tag for Segmented k-Space Traversal
	SegmentedKSpaceTraversalTag = 0x00189033
	// RectilinearPhaseEncodeReorderingTag is tag for Rectilinear Phase Encode Reordering
	RectilinearPhaseEncodeReorderingTag = 0x00189034
	// TagThicknessTag is tag for Tag Thickness
	TagThicknessTag = 0x00189035
	// PartialFourierDirectionTag is tag for Partial Fourier Direction
	PartialFourierDirectionTag = 0x00189036
	// CardiacSynchronizationTechniqueTag is tag for Cardiac Synchronization Technique
	CardiacSynchronizationTechniqueTag = 0x00189037
	// ReceiveCoilManufacturerNameTag is tag for Receive Coil Manufacturer Name
	ReceiveCoilManufacturerNameTag = 0x00189041
	// MRReceiveCoilSequenceTag is tag for MR Receive Coil Sequence
	MRReceiveCoilSequenceTag = 0x00189042
	// ReceiveCoilTypeTag is tag for Receive Coil Type
	ReceiveCoilTypeTag = 0x00189043
	// QuadratureReceiveCoilTag is tag for Quadrature Receive Coil
	QuadratureReceiveCoilTag = 0x00189044
	// MultiCoilDefinitionSequenceTag is tag for Multi-Coil Definition Sequence
	MultiCoilDefinitionSequenceTag = 0x00189045
	// MultiCoilConfigurationTag is tag for Multi-Coil Configuration
	MultiCoilConfigurationTag = 0x00189046
	// MultiCoilElementNameTag is tag for Multi-Coil Element Name
	MultiCoilElementNameTag = 0x00189047
	// MultiCoilElementUsedTag is tag for Multi-Coil Element Used
	MultiCoilElementUsedTag = 0x00189048
	// MRTransmitCoilSequenceTag is tag for MR Transmit Coil Sequence
	MRTransmitCoilSequenceTag = 0x00189049
	// TransmitCoilManufacturerNameTag is tag for Transmit Coil Manufacturer Name
	TransmitCoilManufacturerNameTag = 0x00189050
	// TransmitCoilTypeTag is tag for Transmit Coil Type
	TransmitCoilTypeTag = 0x00189051
	// SpectralWidthTag is tag for Spectral Width
	SpectralWidthTag = 0x00189052
	// ChemicalShiftReferenceTag is tag for Chemical Shift Reference
	ChemicalShiftReferenceTag = 0x00189053
	// VolumeLocalizationTechniqueTag is tag for Volume Localization Technique
	VolumeLocalizationTechniqueTag = 0x00189054
	// MRAcquisitionFrequencyEncodingStepsTag is tag for MR Acquisition Frequency Encoding Steps
	MRAcquisitionFrequencyEncodingStepsTag = 0x00189058
	// DecouplingTag is tag for De-coupling
	DecouplingTag = 0x00189059
	// DecoupledNucleusTag is tag for De-coupled Nucleus
	DecoupledNucleusTag = 0x00189060
	// DecouplingFrequencyTag is tag for De-coupling Frequency
	DecouplingFrequencyTag = 0x00189061
	// DecouplingMethodTag is tag for De-coupling Method
	DecouplingMethodTag = 0x00189062
	// DecouplingChemicalShiftReferenceTag is tag for De-coupling Chemical Shift Reference
	DecouplingChemicalShiftReferenceTag = 0x00189063
	// KSpaceFilteringTag is tag for k-space Filtering
	KSpaceFilteringTag = 0x00189064
	// TimeDomainFilteringTag is tag for Time Domain Filtering
	TimeDomainFilteringTag = 0x00189065
	// NumberOfZeroFillsTag is tag for Number of Zero Fills
	NumberOfZeroFillsTag = 0x00189066
	// BaselineCorrectionTag is tag for Baseline Correction
	BaselineCorrectionTag = 0x00189067
	// ParallelReductionFactorInPlaneTag is tag for Parallel Reduction Factor In-plane
	ParallelReductionFactorInPlaneTag = 0x00189069
	// CardiacRRIntervalSpecifiedTag is tag for Cardiac R-R Interval Specified
	CardiacRRIntervalSpecifiedTag = 0x00189070
	// AcquisitionDurationTag is tag for Acquisition Duration
	AcquisitionDurationTag = 0x00189073
	// FrameAcquisitionDateTimeTag is tag for Frame Acquisition DateTime
	FrameAcquisitionDateTimeTag = 0x00189074
	// DiffusionDirectionalityTag is tag for Diffusion Directionality
	DiffusionDirectionalityTag = 0x00189075
	// DiffusionGradientDirectionSequenceTag is tag for Diffusion Gradient Direction Sequence
	DiffusionGradientDirectionSequenceTag = 0x00189076
	// ParallelAcquisitionTag is tag for Parallel Acquisition
	ParallelAcquisitionTag = 0x00189077
	// ParallelAcquisitionTechniqueTag is tag for Parallel Acquisition Technique
	ParallelAcquisitionTechniqueTag = 0x00189078
	// InversionTimesTag is tag for Inversion Times
	InversionTimesTag = 0x00189079
	// MetaboliteMapDescriptionTag is tag for Metabolite Map Description
	MetaboliteMapDescriptionTag = 0x00189080
	// PartialFourierTag is tag for Partial Fourier
	PartialFourierTag = 0x00189081
	// EffectiveEchoTimeTag is tag for Effective Echo Time
	EffectiveEchoTimeTag = 0x00189082
	// MetaboliteMapCodeSequenceTag is tag for Metabolite Map Code Sequence
	MetaboliteMapCodeSequenceTag = 0x00189083
	// ChemicalShiftSequenceTag is tag for Chemical Shift Sequence
	ChemicalShiftSequenceTag = 0x00189084
	// CardiacSignalSourceTag is tag for Cardiac Signal Source
	CardiacSignalSourceTag = 0x00189085
	// DiffusionBValueTag is tag for Diffusion b-value
	DiffusionBValueTag = 0x00189087
	// DiffusionGradientOrientationTag is tag for Diffusion Gradient Orientation
	DiffusionGradientOrientationTag = 0x00189089
	// VelocityEncodingDirectionTag is tag for Velocity Encoding Direction
	VelocityEncodingDirectionTag = 0x00189090
	// VelocityEncodingMinimumValueTag is tag for Velocity Encoding Minimum Value
	VelocityEncodingMinimumValueTag = 0x00189091
	// VelocityEncodingAcquisitionSequenceTag is tag for Velocity Encoding Acquisition Sequence
	VelocityEncodingAcquisitionSequenceTag = 0x00189092
	// NumberOfKSpaceTrajectoriesTag is tag for Number of k-Space Trajectories
	NumberOfKSpaceTrajectoriesTag = 0x00189093
	// CoverageOfKSpaceTag is tag for Coverage of k-Space
	CoverageOfKSpaceTag = 0x00189094
	// SpectroscopyAcquisitionPhaseRowsTag is tag for Spectroscopy Acquisition Phase Rows
	SpectroscopyAcquisitionPhaseRowsTag = 0x00189095
	// ParallelReductionFactorInPlaneRetiredTag is tag for Parallel Reduction Factor In-plane (Retired)
	ParallelReductionFactorInPlaneRetiredTag = 0x00189096
	// TransmitterFrequencyTag is tag for Transmitter Frequency
	TransmitterFrequencyTag = 0x00189098
	// ResonantNucleusTag is tag for Resonant Nucleus
	ResonantNucleusTag = 0x00189100
	// FrequencyCorrectionTag is tag for Frequency Correction
	FrequencyCorrectionTag = 0x00189101
	// MRSpectroscopyFOVGeometrySequenceTag is tag for MR Spectroscopy FOV/Geometry Sequence
	MRSpectroscopyFOVGeometrySequenceTag = 0x00189103
	// SlabThicknessTag is tag for Slab Thickness
	SlabThicknessTag = 0x00189104
	// SlabOrientationTag is tag for Slab Orientation
	SlabOrientationTag = 0x00189105
	// MidSlabPositionTag is tag for Mid Slab Position
	MidSlabPositionTag = 0x00189106
	// MRSpatialSaturationSequenceTag is tag for MR Spatial Saturation Sequence
	MRSpatialSaturationSequenceTag = 0x00189107
	// MRTimingAndRelatedParametersSequenceTag is tag for MR Timing and Related Parameters Sequence
	MRTimingAndRelatedParametersSequenceTag = 0x00189112
	// MREchoSequenceTag is tag for MR Echo Sequence
	MREchoSequenceTag = 0x00189114
	// MRModifierSequenceTag is tag for MR Modifier Sequence
	MRModifierSequenceTag = 0x00189115
	// MRDiffusionSequenceTag is tag for MR Diffusion Sequence
	MRDiffusionSequenceTag = 0x00189117
	// CardiacSynchronizationSequenceTag is tag for Cardiac Synchronization Sequence
	CardiacSynchronizationSequenceTag = 0x00189118
	// MRAveragesSequenceTag is tag for MR Averages Sequence
	MRAveragesSequenceTag = 0x00189119
	// MRFOVGeometrySequenceTag is tag for MR FOV/Geometry Sequence
	MRFOVGeometrySequenceTag = 0x00189125
	// VolumeLocalizationSequenceTag is tag for Volume Localization Sequence
	VolumeLocalizationSequenceTag = 0x00189126
	// SpectroscopyAcquisitionDataColumnsTag is tag for Spectroscopy Acquisition Data Columns
	SpectroscopyAcquisitionDataColumnsTag = 0x00189127
	// DiffusionAnisotropyTypeTag is tag for Diffusion Anisotropy Type
	DiffusionAnisotropyTypeTag = 0x00189147
	// FrameReferenceDateTimeTag is tag for Frame Reference DateTime
	FrameReferenceDateTimeTag = 0x00189151
	// MRMetaboliteMapSequenceTag is tag for MR Metabolite Map Sequence
	MRMetaboliteMapSequenceTag = 0x00189152
	// ParallelReductionFactorOutOfPlaneTag is tag for Parallel Reduction Factor out-of-plane
	ParallelReductionFactorOutOfPlaneTag = 0x00189155
	// SpectroscopyAcquisitionOutOfPlanePhaseStepsTag is tag for Spectroscopy Acquisition Out-of-plane Phase Steps
	SpectroscopyAcquisitionOutOfPlanePhaseStepsTag = 0x00189159
	// BulkMotionStatusTag is tag for Bulk Motion Status
	BulkMotionStatusTag = 0x00189166
	// ParallelReductionFactorSecondInPlaneTag is tag for Parallel Reduction Factor Second In-plane
	ParallelReductionFactorSecondInPlaneTag = 0x00189168
	// CardiacBeatRejectionTechniqueTag is tag for Cardiac Beat Rejection Technique
	CardiacBeatRejectionTechniqueTag = 0x00189169
	// RespiratoryMotionCompensationTechniqueTag is tag for Respiratory Motion Compensation Technique
	RespiratoryMotionCompensationTechniqueTag = 0x00189170
	// RespiratorySignalSourceTag is tag for Respiratory Signal Source
	RespiratorySignalSourceTag = 0x00189171
	// BulkMotionCompensationTechniqueTag is tag for Bulk Motion Compensation Technique
	BulkMotionCompensationTechniqueTag = 0x00189172
	// BulkMotionSignalSourceTag is tag for Bulk Motion Signal Source
	BulkMotionSignalSourceTag = 0x00189173
	// ApplicableSafetyStandardAgencyTag is tag for Applicable Safety Standard Agency
	ApplicableSafetyStandardAgencyTag = 0x00189174
	// ApplicableSafetyStandardDescriptionTag is tag for Applicable Safety Standard Description
	ApplicableSafetyStandardDescriptionTag = 0x00189175
	// OperatingModeSequenceTag is tag for Operating Mode Sequence
	OperatingModeSequenceTag = 0x00189176
	// OperatingModeTypeTag is tag for Operating Mode Type
	OperatingModeTypeTag = 0x00189177
	// OperatingModeTag is tag for Operating Mode
	OperatingModeTag = 0x00189178
	// SpecificAbsorptionRateDefinitionTag is tag for Specific Absorption Rate Definition
	SpecificAbsorptionRateDefinitionTag = 0x00189179
	// GradientOutputTypeTag is tag for Gradient Output Type
	GradientOutputTypeTag = 0x00189180
	// SpecificAbsorptionRateValueTag is tag for Specific Absorption Rate Value
	SpecificAbsorptionRateValueTag = 0x00189181
	// GradientOutputTag is tag for Gradient Output
	GradientOutputTag = 0x00189182
	// FlowCompensationDirectionTag is tag for Flow Compensation Direction
	FlowCompensationDirectionTag = 0x00189183
	// TaggingDelayTag is tag for Tagging Delay
	TaggingDelayTag = 0x00189184
	// RespiratoryMotionCompensationTechniqueDescriptionTag is tag for Respiratory Motion Compensation Technique Description
	RespiratoryMotionCompensationTechniqueDescriptionTag = 0x00189185
	// RespiratorySignalSourceIDTag is tag for Respiratory Signal Source ID
	RespiratorySignalSourceIDTag = 0x00189186
	// ChemicalShiftMinimumIntegrationLimitInHzTag is tag for Chemical Shift Minimum Integration Limit in Hz
	ChemicalShiftMinimumIntegrationLimitInHzTag = 0x00189195
	// ChemicalShiftMaximumIntegrationLimitInHzTag is tag for Chemical Shift Maximum Integration Limit in Hz
	ChemicalShiftMaximumIntegrationLimitInHzTag = 0x00189196
	// MRVelocityEncodingSequenceTag is tag for MR Velocity Encoding Sequence
	MRVelocityEncodingSequenceTag = 0x00189197
	// FirstOrderPhaseCorrectionTag is tag for First Order Phase Correction
	FirstOrderPhaseCorrectionTag = 0x00189198
	// WaterReferencedPhaseCorrectionTag is tag for Water Referenced Phase Correction
	WaterReferencedPhaseCorrectionTag = 0x00189199
	// MRSpectroscopyAcquisitionTypeTag is tag for MR Spectroscopy Acquisition Type
	MRSpectroscopyAcquisitionTypeTag = 0x00189200
	// RespiratoryCyclePositionTag is tag for Respiratory Cycle Position
	RespiratoryCyclePositionTag = 0x00189214
	// VelocityEncodingMaximumValueTag is tag for Velocity Encoding Maximum Value
	VelocityEncodingMaximumValueTag = 0x00189217
	// TagSpacingSecondDimensionTag is tag for Tag Spacing Second Dimension
	TagSpacingSecondDimensionTag = 0x00189218
	// TagAngleSecondAxisTag is tag for Tag Angle Second Axis
	TagAngleSecondAxisTag = 0x00189219
	// FrameAcquisitionDurationTag is tag for Frame Acquisition Duration
	FrameAcquisitionDurationTag = 0x00189220
	// MRImageFrameTypeSequenceTag is tag for MR Image Frame Type Sequence
	MRImageFrameTypeSequenceTag = 0x00189226
	// MRSpectroscopyFrameTypeSequenceTag is tag for MR Spectroscopy Frame Type Sequence
	MRSpectroscopyFrameTypeSequenceTag = 0x00189227
	// MRAcquisitionPhaseEncodingStepsInPlaneTag is tag for MR Acquisition Phase Encoding Steps in-plane
	MRAcquisitionPhaseEncodingStepsInPlaneTag = 0x00189231
	// MRAcquisitionPhaseEncodingStepsOutOfPlaneTag is tag for MR Acquisition Phase Encoding Steps out-of-plane
	MRAcquisitionPhaseEncodingStepsOutOfPlaneTag = 0x00189232
	// SpectroscopyAcquisitionPhaseColumnsTag is tag for Spectroscopy Acquisition Phase Columns
	SpectroscopyAcquisitionPhaseColumnsTag = 0x00189234
	// CardiacCyclePositionTag is tag for Cardiac Cycle Position
	CardiacCyclePositionTag = 0x00189236
	// SpecificAbsorptionRateSequenceTag is tag for Specific Absorption Rate Sequence
	SpecificAbsorptionRateSequenceTag = 0x00189239
	// RFEchoTrainLengthTag is tag for RF Echo Train Length
	RFEchoTrainLengthTag = 0x00189240
	// GradientEchoTrainLengthTag is tag for Gradient Echo Train Length
	GradientEchoTrainLengthTag = 0x00189241
	// ArterialSpinLabelingContrastTag is tag for Arterial Spin Labeling Contrast
	ArterialSpinLabelingContrastTag = 0x00189250
	// MRArterialSpinLabelingSequenceTag is tag for MR Arterial Spin Labeling Sequence
	MRArterialSpinLabelingSequenceTag = 0x00189251
	// ASLTechniqueDescriptionTag is tag for ASL Technique Description
	ASLTechniqueDescriptionTag = 0x00189252
	// ASLSlabNumberTag is tag for ASL Slab Number
	ASLSlabNumberTag = 0x00189253
	// ASLSlabThicknessTag is tag for ASL Slab Thickness
	ASLSlabThicknessTag = 0x00189254
	// ASLSlabOrientationTag is tag for ASL Slab Orientation
	ASLSlabOrientationTag = 0x00189255
	// ASLMidSlabPositionTag is tag for ASL Mid Slab Position
	ASLMidSlabPositionTag = 0x00189256
	// ASLContextTag is tag for ASL Context
	ASLContextTag = 0x00189257
	// ASLPulseTrainDurationTag is tag for ASL Pulse Train Duration
	ASLPulseTrainDurationTag = 0x00189258
	// ASLCrusherFlagTag is tag for ASL Crusher Flag
	ASLCrusherFlagTag = 0x00189259
	// ASLCrusherFlowLimitTag is tag for ASL Crusher Flow Limit
	ASLCrusherFlowLimitTag = 0x0018925A
	// ASLCrusherDescriptionTag is tag for ASL Crusher Description
	ASLCrusherDescriptionTag = 0x0018925B
	// ASLBolusCutoffFlagTag is tag for ASL Bolus Cut-off Flag
	ASLBolusCutoffFlagTag = 0x0018925C
	// ASLBolusCutoffTimingSequenceTag is tag for ASL Bolus Cut-off Timing Sequence
	ASLBolusCutoffTimingSequenceTag = 0x0018925D
	// ASLBolusCutoffTechniqueTag is tag for ASL Bolus Cut-off Technique
	ASLBolusCutoffTechniqueTag = 0x0018925E
	// ASLBolusCutoffDelayTimeTag is tag for ASL Bolus Cut-off Delay Time
	ASLBolusCutoffDelayTimeTag = 0x0018925F
	// ASLSlabSequenceTag is tag for ASL Slab Sequence
	ASLSlabSequenceTag = 0x00189260
	// ChemicalShiftMinimumIntegrationLimitInppmTag is tag for Chemical Shift Minimum Integration Limit in ppm
	ChemicalShiftMinimumIntegrationLimitInppmTag = 0x00189295
	// ChemicalShiftMaximumIntegrationLimitInppmTag is tag for Chemical Shift Maximum Integration Limit in ppm
	ChemicalShiftMaximumIntegrationLimitInppmTag = 0x00189296
	// WaterReferenceAcquisitionTag is tag for Water Reference Acquisition
	WaterReferenceAcquisitionTag = 0x00189297
	// EchoPeakPositionTag is tag for Echo Peak Position
	EchoPeakPositionTag = 0x00189298
	// CTAcquisitionTypeSequenceTag is tag for CT Acquisition Type Sequence
	CTAcquisitionTypeSequenceTag = 0x00189301
	// AcquisitionTypeTag is tag for Acquisition Type
	AcquisitionTypeTag = 0x00189302
	// TubeAngleTag is tag for Tube Angle
	TubeAngleTag = 0x00189303
	// CTAcquisitionDetailsSequenceTag is tag for CT Acquisition Details Sequence
	CTAcquisitionDetailsSequenceTag = 0x00189304
	// RevolutionTimeTag is tag for Revolution Time
	RevolutionTimeTag = 0x00189305
	// SingleCollimationWidthTag is tag for Single Collimation Width
	SingleCollimationWidthTag = 0x00189306
	// TotalCollimationWidthTag is tag for Total Collimation Width
	TotalCollimationWidthTag = 0x00189307
	// CTTableDynamicsSequenceTag is tag for CT Table Dynamics Sequence
	CTTableDynamicsSequenceTag = 0x00189308
	// TableSpeedTag is tag for Table Speed
	TableSpeedTag = 0x00189309
	// TableFeedPerRotationTag is tag for Table Feed per Rotation
	TableFeedPerRotationTag = 0x00189310
	// SpiralPitchFactorTag is tag for Spiral Pitch Factor
	SpiralPitchFactorTag = 0x00189311
	// CTGeometrySequenceTag is tag for CT Geometry Sequence
	CTGeometrySequenceTag = 0x00189312
	// DataCollectionCenterPatientTag is tag for Data Collection Center (Patient)
	DataCollectionCenterPatientTag = 0x00189313
	// CTReconstructionSequenceTag is tag for CT Reconstruction Sequence
	CTReconstructionSequenceTag = 0x00189314
	// ReconstructionAlgorithmTag is tag for Reconstruction Algorithm
	ReconstructionAlgorithmTag = 0x00189315
	// ConvolutionKernelGroupTag is tag for Convolution Kernel Group
	ConvolutionKernelGroupTag = 0x00189316
	// ReconstructionFieldOfViewTag is tag for Reconstruction Field of View
	ReconstructionFieldOfViewTag = 0x00189317
	// ReconstructionTargetCenterPatientTag is tag for Reconstruction Target Center (Patient)
	ReconstructionTargetCenterPatientTag = 0x00189318
	// ReconstructionAngleTag is tag for Reconstruction Angle
	ReconstructionAngleTag = 0x00189319
	// ImageFilterTag is tag for Image Filter
	ImageFilterTag = 0x00189320
	// CTExposureSequenceTag is tag for CT Exposure Sequence
	CTExposureSequenceTag = 0x00189321
	// ReconstructionPixelSpacingTag is tag for Reconstruction Pixel Spacing
	ReconstructionPixelSpacingTag = 0x00189322
	// ExposureModulationTypeTag is tag for Exposure Modulation Type
	ExposureModulationTypeTag = 0x00189323
	// EstimatedDoseSavingTag is tag for Estimated Dose Saving
	EstimatedDoseSavingTag = 0x00189324
	// CTXRayDetailsSequenceTag is tag for CT X-Ray Details Sequence
	CTXRayDetailsSequenceTag = 0x00189325
	// CTPositionSequenceTag is tag for CT Position Sequence
	CTPositionSequenceTag = 0x00189326
	// TablePositionTag is tag for Table Position
	TablePositionTag = 0x00189327
	// ExposureTimeInmsTag is tag for Exposure Time in ms
	ExposureTimeInmsTag = 0x00189328
	// CTImageFrameTypeSequenceTag is tag for CT Image Frame Type Sequence
	CTImageFrameTypeSequenceTag = 0x00189329
	// XRayTubeCurrentInmATag is tag for X-Ray Tube Current in mA
	XRayTubeCurrentInmATag = 0x00189330
	// ExposureInmAsTag is tag for Exposure in mAs
	ExposureInmAsTag = 0x00189332
	// ConstantVolumeFlagTag is tag for Constant Volume Flag
	ConstantVolumeFlagTag = 0x00189333
	// FluoroscopyFlagTag is tag for Fluoroscopy Flag
	FluoroscopyFlagTag = 0x00189334
	// DistanceSourceToDataCollectionCenterTag is tag for Distance Source to Data Collection Center
	DistanceSourceToDataCollectionCenterTag = 0x00189335
	// ContrastBolusAgentNumberTag is tag for Contrast/Bolus Agent Number
	ContrastBolusAgentNumberTag = 0x00189337
	// ContrastBolusIngredientCodeSequenceTag is tag for Contrast/Bolus Ingredient Code Sequence
	ContrastBolusIngredientCodeSequenceTag = 0x00189338
	// ContrastAdministrationProfileSequenceTag is tag for Contrast Administration Profile Sequence
	ContrastAdministrationProfileSequenceTag = 0x00189340
	// ContrastBolusUsageSequenceTag is tag for Contrast/Bolus Usage Sequence
	ContrastBolusUsageSequenceTag = 0x00189341
	// ContrastBolusAgentAdministeredTag is tag for Contrast/Bolus Agent Administered
	ContrastBolusAgentAdministeredTag = 0x00189342
	// ContrastBolusAgentDetectedTag is tag for Contrast/Bolus Agent Detected
	ContrastBolusAgentDetectedTag = 0x00189343
	// ContrastBolusAgentPhaseTag is tag for Contrast/Bolus Agent Phase
	ContrastBolusAgentPhaseTag = 0x00189344
	// CTDIvolTag is tag for CTDIvol
	CTDIvolTag = 0x00189345
	// CTDIPhantomTypeCodeSequenceTag is tag for CTDI Phantom Type Code Sequence
	CTDIPhantomTypeCodeSequenceTag = 0x00189346
	// CalciumScoringMassFactorPatientTag is tag for Calcium Scoring Mass Factor Patient
	CalciumScoringMassFactorPatientTag = 0x00189351
	// CalciumScoringMassFactorDeviceTag is tag for Calcium Scoring Mass Factor Device
	CalciumScoringMassFactorDeviceTag = 0x00189352
	// EnergyWeightingFactorTag is tag for Energy Weighting Factor
	EnergyWeightingFactorTag = 0x00189353
	// CTAdditionalXRaySourceSequenceTag is tag for CT Additional X-Ray Source Sequence
	CTAdditionalXRaySourceSequenceTag = 0x00189360
	// MultienergyCTAcquisitionTag is tag for Multi-energy CT Acquisition
	MultienergyCTAcquisitionTag = 0x00189361
	// MultienergyCTAcquisitionSequenceTag is tag for Multi-energy CT Acquisition Sequence
	MultienergyCTAcquisitionSequenceTag = 0x00189362
	// MultienergyCTProcessingSequenceTag is tag for Multi-energy CT Processing Sequence
	MultienergyCTProcessingSequenceTag = 0x00189363
	// MultienergyCTCharacteristicsSequenceTag is tag for Multi-energy CT Characteristics Sequence
	MultienergyCTCharacteristicsSequenceTag = 0x00189364
	// MultienergyCTXRaySourceSequenceTag is tag for Multi-energy CT X-Ray Source Sequence
	MultienergyCTXRaySourceSequenceTag = 0x00189365
	// XRaySourceIndexTag is tag for X-Ray Source Index
	XRaySourceIndexTag = 0x00189366
	// XRaySourceIDTag is tag for X-Ray Source ID
	XRaySourceIDTag = 0x00189367
	// MultienergySourceTechniqueTag is tag for Multi-energy Source Technique
	MultienergySourceTechniqueTag = 0x00189368
	// SourceStartDateTimeTag is tag for Source Start DateTime
	SourceStartDateTimeTag = 0x00189369
	// SourceEndDateTimeTag is tag for Source End DateTime
	SourceEndDateTimeTag = 0x0018936A
	// SwitchingPhaseNumberTag is tag for Switching Phase Number
	SwitchingPhaseNumberTag = 0x0018936B
	// SwitchingPhaseNominalDurationTag is tag for Switching Phase Nominal Duration
	SwitchingPhaseNominalDurationTag = 0x0018936C
	// SwitchingPhaseTransitionDurationTag is tag for Switching Phase Transition Duration
	SwitchingPhaseTransitionDurationTag = 0x0018936D
	// EffectiveBinEnergyTag is tag for Effective Bin Energy
	EffectiveBinEnergyTag = 0x0018936E
	// MultienergyCTXRayDetectorSequenceTag is tag for Multi-energy CT X-Ray Detector Sequence
	MultienergyCTXRayDetectorSequenceTag = 0x0018936F
	// XRayDetectorIndexTag is tag for X-Ray Detector Index
	XRayDetectorIndexTag = 0x00189370
	// XRayDetectorIDTag is tag for X-Ray Detector ID
	XRayDetectorIDTag = 0x00189371
	// MultienergyDetectorTypeTag is tag for Multi-energy Detector Type
	MultienergyDetectorTypeTag = 0x00189372
	// XRayDetectorLabelTag is tag for X-Ray Detector Label
	XRayDetectorLabelTag = 0x00189373
	// NominalMaxEnergyTag is tag for Nominal Max Energy
	NominalMaxEnergyTag = 0x00189374
	// NominalMinEnergyTag is tag for Nominal Min Energy
	NominalMinEnergyTag = 0x00189375
	// ReferencedXRayDetectorIndexTag is tag for Referenced X-Ray Detector Index
	ReferencedXRayDetectorIndexTag = 0x00189376
	// ReferencedXRaySourceIndexTag is tag for Referenced X-Ray Source Index
	ReferencedXRaySourceIndexTag = 0x00189377
	// ReferencedPathIndexTag is tag for Referenced Path Index
	ReferencedPathIndexTag = 0x00189378
	// MultienergyCTPathSequenceTag is tag for Multi-energy CT Path Sequence
	MultienergyCTPathSequenceTag = 0x00189379
	// MultienergyCTPathIndexTag is tag for Multi-energy CT Path Index
	MultienergyCTPathIndexTag = 0x0018937A
	// MultienergyAcquisitionDescriptionTag is tag for Multi-energy Acquisition Description
	MultienergyAcquisitionDescriptionTag = 0x0018937B
	// MonoenergeticEnergyEquivalentTag is tag for Monoenergetic Energy Equivalent
	MonoenergeticEnergyEquivalentTag = 0x0018937C
	// MaterialCodeSequenceTag is tag for Material Code Sequence
	MaterialCodeSequenceTag = 0x0018937D
	// DecompositionMethodTag is tag for Decomposition Method
	DecompositionMethodTag = 0x0018937E
	// DecompositionDescriptionTag is tag for Decomposition Description
	DecompositionDescriptionTag = 0x0018937F
	// DecompositionAlgorithmIdentificationSequenceTag is tag for Decomposition Algorithm Identification Sequence
	DecompositionAlgorithmIdentificationSequenceTag = 0x00189380
	// DecompositionMaterialSequenceTag is tag for Decomposition Material Sequence
	DecompositionMaterialSequenceTag = 0x00189381
	// MaterialAttenuationSequenceTag is tag for Material Attenuation Sequence
	MaterialAttenuationSequenceTag = 0x00189382
	// PhotonEnergyTag is tag for Photon Energy
	PhotonEnergyTag = 0x00189383
	// XRayMassAttenuationCoefficientTag is tag for X-Ray Mass Attenuation Coefficient
	XRayMassAttenuationCoefficientTag = 0x00189384
	// ProjectionPixelCalibrationSequenceTag is tag for Projection Pixel Calibration Sequence
	ProjectionPixelCalibrationSequenceTag = 0x00189401
	// DistanceSourceToIsocenterTag is tag for Distance Source to Isocenter
	DistanceSourceToIsocenterTag = 0x00189402
	// DistanceObjectToTableTopTag is tag for Distance Object to Table Top
	DistanceObjectToTableTopTag = 0x00189403
	// ObjectPixelSpacingInCenterOfBeamTag is tag for Object Pixel Spacing in Center of Beam
	ObjectPixelSpacingInCenterOfBeamTag = 0x00189404
	// PositionerPositionSequenceTag is tag for Positioner Position Sequence
	PositionerPositionSequenceTag = 0x00189405
	// TablePositionSequenceTag is tag for Table Position Sequence
	TablePositionSequenceTag = 0x00189406
	// CollimatorShapeSequenceTag is tag for Collimator Shape Sequence
	CollimatorShapeSequenceTag = 0x00189407
	// PlanesInAcquisitionTag is tag for Planes in Acquisition
	PlanesInAcquisitionTag = 0x00189410
	// XAXRFFrameCharacteristicsSequenceTag is tag for XA/XRF Frame Characteristics Sequence
	XAXRFFrameCharacteristicsSequenceTag = 0x00189412
	// FrameAcquisitionSequenceTag is tag for Frame Acquisition Sequence
	FrameAcquisitionSequenceTag = 0x00189417
	// XRayReceptorTypeTag is tag for X-Ray Receptor Type
	XRayReceptorTypeTag = 0x00189420
	// AcquisitionProtocolNameTag is tag for Acquisition Protocol Name
	AcquisitionProtocolNameTag = 0x00189423
	// AcquisitionProtocolDescriptionTag is tag for Acquisition Protocol Description
	AcquisitionProtocolDescriptionTag = 0x00189424
	// ContrastBolusIngredientOpaqueTag is tag for Contrast/Bolus Ingredient Opaque
	ContrastBolusIngredientOpaqueTag = 0x00189425
	// DistanceReceptorPlaneToDetectorHousingTag is tag for Distance Receptor Plane to Detector Housing
	DistanceReceptorPlaneToDetectorHousingTag = 0x00189426
	// IntensifierActiveShapeTag is tag for Intensifier Active Shape
	IntensifierActiveShapeTag = 0x00189427
	// IntensifierActiveDimensionsTag is tag for Intensifier Active Dimension(s)
	IntensifierActiveDimensionsTag = 0x00189428
	// PhysicalDetectorSizeTag is tag for Physical Detector Size
	PhysicalDetectorSizeTag = 0x00189429
	// PositionOfIsocenterProjectionTag is tag for Position of Isocenter Projection
	PositionOfIsocenterProjectionTag = 0x00189430
	// FieldOfViewSequenceTag is tag for Field of View Sequence
	FieldOfViewSequenceTag = 0x00189432
	// FieldOfViewDescriptionTag is tag for Field of View Description
	FieldOfViewDescriptionTag = 0x00189433
	// ExposureControlSensingRegionsSequenceTag is tag for Exposure Control Sensing Regions Sequence
	ExposureControlSensingRegionsSequenceTag = 0x00189434
	// ExposureControlSensingRegionShapeTag is tag for Exposure Control Sensing Region Shape
	ExposureControlSensingRegionShapeTag = 0x00189435
	// ExposureControlSensingRegionLeftVerticalEdgeTag is tag for Exposure Control Sensing Region Left Vertical Edge
	ExposureControlSensingRegionLeftVerticalEdgeTag = 0x00189436
	// ExposureControlSensingRegionRightVerticalEdgeTag is tag for Exposure Control Sensing Region Right Vertical Edge
	ExposureControlSensingRegionRightVerticalEdgeTag = 0x00189437
	// ExposureControlSensingRegionUpperHorizontalEdgeTag is tag for Exposure Control Sensing Region Upper Horizontal Edge
	ExposureControlSensingRegionUpperHorizontalEdgeTag = 0x00189438
	// ExposureControlSensingRegionLowerHorizontalEdgeTag is tag for Exposure Control Sensing Region Lower Horizontal Edge
	ExposureControlSensingRegionLowerHorizontalEdgeTag = 0x00189439
	// CenterOfCircularExposureControlSensingRegionTag is tag for Center of Circular Exposure Control Sensing Region
	CenterOfCircularExposureControlSensingRegionTag = 0x00189440
	// RadiusOfCircularExposureControlSensingRegionTag is tag for Radius of Circular Exposure Control Sensing Region
	RadiusOfCircularExposureControlSensingRegionTag = 0x00189441
	// VerticesOfThePolygonalExposureControlSensingRegionTag is tag for Vertices of the Polygonal Exposure Control Sensing Region
	VerticesOfThePolygonalExposureControlSensingRegionTag = 0x00189442
	// ColumnAngulationPatientTag is tag for Column Angulation (Patient)
	ColumnAngulationPatientTag = 0x00189447
	// BeamAngleTag is tag for Beam Angle
	BeamAngleTag = 0x00189449
	// FrameDetectorParametersSequenceTag is tag for Frame Detector Parameters Sequence
	FrameDetectorParametersSequenceTag = 0x00189451
	// CalculatedAnatomyThicknessTag is tag for Calculated Anatomy Thickness
	CalculatedAnatomyThicknessTag = 0x00189452
	// CalibrationSequenceTag is tag for Calibration Sequence
	CalibrationSequenceTag = 0x00189455
	// ObjectThicknessSequenceTag is tag for Object Thickness Sequence
	ObjectThicknessSequenceTag = 0x00189456
	// PlaneIdentificationTag is tag for Plane Identification
	PlaneIdentificationTag = 0x00189457
	// FieldOfViewDimensionsInFloatTag is tag for Field of View Dimension(s) in Float
	FieldOfViewDimensionsInFloatTag = 0x00189461
	// IsocenterReferenceSystemSequenceTag is tag for Isocenter Reference System Sequence
	IsocenterReferenceSystemSequenceTag = 0x00189462
	// PositionerIsocenterPrimaryAngleTag is tag for Positioner Isocenter Primary Angle
	PositionerIsocenterPrimaryAngleTag = 0x00189463
	// PositionerIsocenterSecondaryAngleTag is tag for Positioner Isocenter Secondary Angle
	PositionerIsocenterSecondaryAngleTag = 0x00189464
	// PositionerIsocenterDetectorRotationAngleTag is tag for Positioner Isocenter Detector Rotation Angle
	PositionerIsocenterDetectorRotationAngleTag = 0x00189465
	// TableXPositionToIsocenterTag is tag for Table X Position to Isocenter
	TableXPositionToIsocenterTag = 0x00189466
	// TableYPositionToIsocenterTag is tag for Table Y Position to Isocenter
	TableYPositionToIsocenterTag = 0x00189467
	// TableZPositionToIsocenterTag is tag for Table Z Position to Isocenter
	TableZPositionToIsocenterTag = 0x00189468
	// TableHorizontalRotationAngleTag is tag for Table Horizontal Rotation Angle
	TableHorizontalRotationAngleTag = 0x00189469
	// TableHeadTiltAngleTag is tag for Table Head Tilt Angle
	TableHeadTiltAngleTag = 0x00189470
	// TableCradleTiltAngleTag is tag for Table Cradle Tilt Angle
	TableCradleTiltAngleTag = 0x00189471
	// FrameDisplayShutterSequenceTag is tag for Frame Display Shutter Sequence
	FrameDisplayShutterSequenceTag = 0x00189472
	// AcquiredImageAreaDoseProductTag is tag for Acquired Image Area Dose Product
	AcquiredImageAreaDoseProductTag = 0x00189473
	// CArmPositionerTabletopRelationshipTag is tag for C-arm Positioner Tabletop Relationship
	CArmPositionerTabletopRelationshipTag = 0x00189474
	// XRayGeometrySequenceTag is tag for X-Ray Geometry Sequence
	XRayGeometrySequenceTag = 0x00189476
	// IrradiationEventIdentificationSequenceTag is tag for Irradiation Event Identification Sequence
	IrradiationEventIdentificationSequenceTag = 0x00189477
	// XRay3DFrameTypeSequenceTag is tag for X-Ray 3D Frame Type Sequence
	XRay3DFrameTypeSequenceTag = 0x00189504
	// ContributingSourcesSequenceTag is tag for Contributing Sources Sequence
	ContributingSourcesSequenceTag = 0x00189506
	// XRay3DAcquisitionSequenceTag is tag for X-Ray 3D Acquisition Sequence
	XRay3DAcquisitionSequenceTag = 0x00189507
	// PrimaryPositionerScanArcTag is tag for Primary Positioner Scan Arc
	PrimaryPositionerScanArcTag = 0x00189508
	// SecondaryPositionerScanArcTag is tag for Secondary Positioner Scan Arc
	SecondaryPositionerScanArcTag = 0x00189509
	// PrimaryPositionerScanStartAngleTag is tag for Primary Positioner Scan Start Angle
	PrimaryPositionerScanStartAngleTag = 0x00189510
	// SecondaryPositionerScanStartAngleTag is tag for Secondary Positioner Scan Start Angle
	SecondaryPositionerScanStartAngleTag = 0x00189511
	// PrimaryPositionerIncrementTag is tag for Primary Positioner Increment
	PrimaryPositionerIncrementTag = 0x00189514
	// SecondaryPositionerIncrementTag is tag for Secondary Positioner Increment
	SecondaryPositionerIncrementTag = 0x00189515
	// StartAcquisitionDateTimeTag is tag for Start Acquisition DateTime
	StartAcquisitionDateTimeTag = 0x00189516
	// EndAcquisitionDateTimeTag is tag for End Acquisition DateTime
	EndAcquisitionDateTimeTag = 0x00189517
	// PrimaryPositionerIncrementSignTag is tag for Primary Positioner Increment Sign
	PrimaryPositionerIncrementSignTag = 0x00189518
	// SecondaryPositionerIncrementSignTag is tag for Secondary Positioner Increment Sign
	SecondaryPositionerIncrementSignTag = 0x00189519
	// ApplicationNameTag is tag for Application Name
	ApplicationNameTag = 0x00189524
	// ApplicationVersionTag is tag for Application Version
	ApplicationVersionTag = 0x00189525
	// ApplicationManufacturerTag is tag for Application Manufacturer
	ApplicationManufacturerTag = 0x00189526
	// AlgorithmTypeTag is tag for Algorithm Type
	AlgorithmTypeTag = 0x00189527
	// AlgorithmDescriptionTag is tag for Algorithm Description
	AlgorithmDescriptionTag = 0x00189528
	// XRay3DReconstructionSequenceTag is tag for X-Ray 3D Reconstruction Sequence
	XRay3DReconstructionSequenceTag = 0x00189530
	// ReconstructionDescriptionTag is tag for Reconstruction Description
	ReconstructionDescriptionTag = 0x00189531
	// PerProjectionAcquisitionSequenceTag is tag for Per Projection Acquisition Sequence
	PerProjectionAcquisitionSequenceTag = 0x00189538
	// DetectorPositionSequenceTag is tag for Detector Position Sequence
	DetectorPositionSequenceTag = 0x00189541
	// XRayAcquisitionDoseSequenceTag is tag for X-Ray Acquisition Dose Sequence
	XRayAcquisitionDoseSequenceTag = 0x00189542
	// XRaySourceIsocenterPrimaryAngleTag is tag for X-Ray Source Isocenter Primary Angle
	XRaySourceIsocenterPrimaryAngleTag = 0x00189543
	// XRaySourceIsocenterSecondaryAngleTag is tag for X-Ray Source Isocenter Secondary Angle
	XRaySourceIsocenterSecondaryAngleTag = 0x00189544
	// BreastSupportIsocenterPrimaryAngleTag is tag for Breast Support Isocenter Primary Angle
	BreastSupportIsocenterPrimaryAngleTag = 0x00189545
	// BreastSupportIsocenterSecondaryAngleTag is tag for Breast Support Isocenter Secondary Angle
	BreastSupportIsocenterSecondaryAngleTag = 0x00189546
	// BreastSupportXPositionToIsocenterTag is tag for Breast Support X Position to Isocenter
	BreastSupportXPositionToIsocenterTag = 0x00189547
	// BreastSupportYPositionToIsocenterTag is tag for Breast Support Y Position to Isocenter
	BreastSupportYPositionToIsocenterTag = 0x00189548
	// BreastSupportZPositionToIsocenterTag is tag for Breast Support Z Position to Isocenter
	BreastSupportZPositionToIsocenterTag = 0x00189549
	// DetectorIsocenterPrimaryAngleTag is tag for Detector Isocenter Primary Angle
	DetectorIsocenterPrimaryAngleTag = 0x00189550
	// DetectorIsocenterSecondaryAngleTag is tag for Detector Isocenter Secondary Angle
	DetectorIsocenterSecondaryAngleTag = 0x00189551
	// DetectorXPositionToIsocenterTag is tag for Detector X Position to Isocenter
	DetectorXPositionToIsocenterTag = 0x00189552
	// DetectorYPositionToIsocenterTag is tag for Detector Y Position to Isocenter
	DetectorYPositionToIsocenterTag = 0x00189553
	// DetectorZPositionToIsocenterTag is tag for Detector Z Position to Isocenter
	DetectorZPositionToIsocenterTag = 0x00189554
	// XRayGridSequenceTag is tag for X-Ray Grid Sequence
	XRayGridSequenceTag = 0x00189555
	// XRayFilterSequenceTag is tag for X-Ray Filter Sequence
	XRayFilterSequenceTag = 0x00189556
	// DetectorActiveAreaTLHCPositionTag is tag for Detector Active Area TLHC Position
	DetectorActiveAreaTLHCPositionTag = 0x00189557
	// DetectorActiveAreaOrientationTag is tag for Detector Active Area Orientation
	DetectorActiveAreaOrientationTag = 0x00189558
	// PositionerPrimaryAngleDirectionTag is tag for Positioner Primary Angle Direction
	PositionerPrimaryAngleDirectionTag = 0x00189559
	// DiffusionBMatrixSequenceTag is tag for Diffusion b-matrix Sequence
	DiffusionBMatrixSequenceTag = 0x00189601
	// DiffusionBValueXXTag is tag for Diffusion b-value XX
	DiffusionBValueXXTag = 0x00189602
	// DiffusionBValueXYTag is tag for Diffusion b-value XY
	DiffusionBValueXYTag = 0x00189603
	// DiffusionBValueXZTag is tag for Diffusion b-value XZ
	DiffusionBValueXZTag = 0x00189604
	// DiffusionBValueYYTag is tag for Diffusion b-value YY
	DiffusionBValueYYTag = 0x00189605
	// DiffusionBValueYZTag is tag for Diffusion b-value YZ
	DiffusionBValueYZTag = 0x00189606
	// DiffusionBValueZZTag is tag for Diffusion b-value ZZ
	DiffusionBValueZZTag = 0x00189607
	// FunctionalMRSequenceTag is tag for Functional MR Sequence
	FunctionalMRSequenceTag = 0x00189621
	// FunctionalSettlingPhaseFramesPresentTag is tag for Functional Settling Phase Frames Present
	FunctionalSettlingPhaseFramesPresentTag = 0x00189622
	// FunctionalSyncPulseTag is tag for Functional Sync Pulse
	FunctionalSyncPulseTag = 0x00189623
	// SettlingPhaseFrameTag is tag for Settling Phase Frame
	SettlingPhaseFrameTag = 0x00189624
	// DecayCorrectionDateTimeTag is tag for Decay Correction DateTime
	DecayCorrectionDateTimeTag = 0x00189701
	// StartDensityThresholdTag is tag for Start Density Threshold
	StartDensityThresholdTag = 0x00189715
	// StartRelativeDensityDifferenceThresholdTag is tag for Start Relative Density Difference Threshold
	StartRelativeDensityDifferenceThresholdTag = 0x00189716
	// StartCardiacTriggerCountThresholdTag is tag for Start Cardiac Trigger Count Threshold
	StartCardiacTriggerCountThresholdTag = 0x00189717
	// StartRespiratoryTriggerCountThresholdTag is tag for Start Respiratory Trigger Count Threshold
	StartRespiratoryTriggerCountThresholdTag = 0x00189718
	// TerminationCountsThresholdTag is tag for Termination Counts Threshold
	TerminationCountsThresholdTag = 0x00189719
	// TerminationDensityThresholdTag is tag for Termination Density Threshold
	TerminationDensityThresholdTag = 0x00189720
	// TerminationRelativeDensityThresholdTag is tag for Termination Relative Density Threshold
	TerminationRelativeDensityThresholdTag = 0x00189721
	// TerminationTimeThresholdTag is tag for Termination Time Threshold
	TerminationTimeThresholdTag = 0x00189722
	// TerminationCardiacTriggerCountThresholdTag is tag for Termination Cardiac Trigger Count Threshold
	TerminationCardiacTriggerCountThresholdTag = 0x00189723
	// TerminationRespiratoryTriggerCountThresholdTag is tag for Termination Respiratory Trigger Count Threshold
	TerminationRespiratoryTriggerCountThresholdTag = 0x00189724
	// DetectorGeometryTag is tag for Detector Geometry
	DetectorGeometryTag = 0x00189725
	// TransverseDetectorSeparationTag is tag for Transverse Detector Separation
	TransverseDetectorSeparationTag = 0x00189726
	// AxialDetectorDimensionTag is tag for Axial Detector Dimension
	AxialDetectorDimensionTag = 0x00189727
	// RadiopharmaceuticalAgentNumberTag is tag for Radiopharmaceutical Agent Number
	RadiopharmaceuticalAgentNumberTag = 0x00189729
	// PETFrameAcquisitionSequenceTag is tag for PET Frame Acquisition Sequence
	PETFrameAcquisitionSequenceTag = 0x00189732
	// PETDetectorMotionDetailsSequenceTag is tag for PET Detector Motion Details Sequence
	PETDetectorMotionDetailsSequenceTag = 0x00189733
	// PETTableDynamicsSequenceTag is tag for PET Table Dynamics Sequence
	PETTableDynamicsSequenceTag = 0x00189734
	// PETPositionSequenceTag is tag for PET Position Sequence
	PETPositionSequenceTag = 0x00189735
	// PETFrameCorrectionFactorsSequenceTag is tag for PET Frame Correction Factors Sequence
	PETFrameCorrectionFactorsSequenceTag = 0x00189736
	// RadiopharmaceuticalUsageSequenceTag is tag for Radiopharmaceutical Usage Sequence
	RadiopharmaceuticalUsageSequenceTag = 0x00189737
	// AttenuationCorrectionSourceTag is tag for Attenuation Correction Source
	AttenuationCorrectionSourceTag = 0x00189738
	// NumberOfIterationsTag is tag for Number of Iterations
	NumberOfIterationsTag = 0x00189739
	// NumberOfSubsetsTag is tag for Number of Subsets
	NumberOfSubsetsTag = 0x00189740
	// PETReconstructionSequenceTag is tag for PET Reconstruction Sequence
	PETReconstructionSequenceTag = 0x00189749
	// PETFrameTypeSequenceTag is tag for PET Frame Type Sequence
	PETFrameTypeSequenceTag = 0x00189751
	// TimeOfFlightInformationUsedTag is tag for Time of Flight Information Used
	TimeOfFlightInformationUsedTag = 0x00189755
	// ReconstructionTypeTag is tag for Reconstruction Type
	ReconstructionTypeTag = 0x00189756
	// DecayCorrectedTag is tag for Decay Corrected
	DecayCorrectedTag = 0x00189758
	// AttenuationCorrectedTag is tag for Attenuation Corrected
	AttenuationCorrectedTag = 0x00189759
	// ScatterCorrectedTag is tag for Scatter Corrected
	ScatterCorrectedTag = 0x00189760
	// DeadTimeCorrectedTag is tag for Dead Time Corrected
	DeadTimeCorrectedTag = 0x00189761
	// GantryMotionCorrectedTag is tag for Gantry Motion Corrected
	GantryMotionCorrectedTag = 0x00189762
	// PatientMotionCorrectedTag is tag for Patient Motion Corrected
	PatientMotionCorrectedTag = 0x00189763
	// CountLossNormalizationCorrectedTag is tag for Count Loss Normalization Corrected
	CountLossNormalizationCorrectedTag = 0x00189764
	// RandomsCorrectedTag is tag for Randoms Corrected
	RandomsCorrectedTag = 0x00189765
	// NonUniformRadialSamplingCorrectedTag is tag for Non-uniform Radial Sampling Corrected
	NonUniformRadialSamplingCorrectedTag = 0x00189766
	// SensitivityCalibratedTag is tag for Sensitivity Calibrated
	SensitivityCalibratedTag = 0x00189767
	// DetectorNormalizationCorrectionTag is tag for Detector Normalization Correction
	DetectorNormalizationCorrectionTag = 0x00189768
	// IterativeReconstructionMethodTag is tag for Iterative Reconstruction Method
	IterativeReconstructionMethodTag = 0x00189769
	// AttenuationCorrectionTemporalRelationshipTag is tag for Attenuation Correction Temporal Relationship
	AttenuationCorrectionTemporalRelationshipTag = 0x00189770
	// PatientPhysiologicalStateSequenceTag is tag for Patient Physiological State Sequence
	PatientPhysiologicalStateSequenceTag = 0x00189771
	// PatientPhysiologicalStateCodeSequenceTag is tag for Patient Physiological State Code Sequence
	PatientPhysiologicalStateCodeSequenceTag = 0x00189772
	// DepthsOfFocusTag is tag for Depth(s) of Focus
	DepthsOfFocusTag = 0x00189801
	// ExcludedIntervalsSequenceTag is tag for Excluded Intervals Sequence
	ExcludedIntervalsSequenceTag = 0x00189803
	// ExclusionStartDateTimeTag is tag for Exclusion Start DateTime
	ExclusionStartDateTimeTag = 0x00189804
	// ExclusionDurationTag is tag for Exclusion Duration
	ExclusionDurationTag = 0x00189805
	// USImageDescriptionSequenceTag is tag for US Image Description Sequence
	USImageDescriptionSequenceTag = 0x00189806
	// ImageDataTypeSequenceTag is tag for Image Data Type Sequence
	ImageDataTypeSequenceTag = 0x00189807
	// DataTypeTag is tag for Data Type
	DataTypeTag = 0x00189808
	// TransducerScanPatternCodeSequenceTag is tag for Transducer Scan Pattern Code Sequence
	TransducerScanPatternCodeSequenceTag = 0x00189809
	// AliasedDataTypeTag is tag for Aliased Data Type
	AliasedDataTypeTag = 0x0018980B
	// PositionMeasuringDeviceUsedTag is tag for Position Measuring Device Used
	PositionMeasuringDeviceUsedTag = 0x0018980C
	// TransducerGeometryCodeSequenceTag is tag for Transducer Geometry Code Sequence
	TransducerGeometryCodeSequenceTag = 0x0018980D
	// TransducerBeamSteeringCodeSequenceTag is tag for Transducer Beam Steering Code Sequence
	TransducerBeamSteeringCodeSequenceTag = 0x0018980E
	// TransducerApplicationCodeSequenceTag is tag for Transducer Application Code Sequence
	TransducerApplicationCodeSequenceTag = 0x0018980F
	// ZeroVelocityPixelValueTag is tag for Zero Velocity Pixel Value
	ZeroVelocityPixelValueTag = 0x00189810
	// ReferenceLocationLabelTag is tag for Reference Location Label
	ReferenceLocationLabelTag = 0x00189900
	// ReferenceLocationDescriptionTag is tag for Reference Location Description
	ReferenceLocationDescriptionTag = 0x00189901
	// ReferenceBasisCodeSequenceTag is tag for Reference Basis Code Sequence
	ReferenceBasisCodeSequenceTag = 0x00189902
	// ReferenceGeometryCodeSequenceTag is tag for Reference Geometry Code Sequence
	ReferenceGeometryCodeSequenceTag = 0x00189903
	// OffsetDistanceTag is tag for Offset Distance
	OffsetDistanceTag = 0x00189904
	// OffsetDirectionTag is tag for Offset Direction
	OffsetDirectionTag = 0x00189905
	// PotentialScheduledProtocolCodeSequenceTag is tag for Potential Scheduled Protocol Code Sequence
	PotentialScheduledProtocolCodeSequenceTag = 0x00189906
	// PotentialRequestedProcedureCodeSequenceTag is tag for Potential Requested Procedure Code Sequence
	PotentialRequestedProcedureCodeSequenceTag = 0x00189907
	// PotentialReasonsForProcedureTag is tag for Potential Reasons for Procedure
	PotentialReasonsForProcedureTag = 0x00189908
	// PotentialReasonsForProcedureCodeSequenceTag is tag for Potential Reasons for Procedure Code Sequence
	PotentialReasonsForProcedureCodeSequenceTag = 0x00189909
	// PotentialDiagnosticTasksTag is tag for Potential Diagnostic Tasks
	PotentialDiagnosticTasksTag = 0x0018990A
	// ContraindicationsCodeSequenceTag is tag for Contraindications Code Sequence
	ContraindicationsCodeSequenceTag = 0x0018990B
	// ReferencedDefinedProtocolSequenceTag is tag for Referenced Defined Protocol Sequence
	ReferencedDefinedProtocolSequenceTag = 0x0018990C
	// ReferencedPerformedProtocolSequenceTag is tag for Referenced Performed Protocol Sequence
	ReferencedPerformedProtocolSequenceTag = 0x0018990D
	// PredecessorProtocolSequenceTag is tag for Predecessor Protocol Sequence
	PredecessorProtocolSequenceTag = 0x0018990E
	// ProtocolPlanningInformationTag is tag for Protocol Planning Information
	ProtocolPlanningInformationTag = 0x0018990F
	// ProtocolDesignRationaleTag is tag for Protocol Design Rationale
	ProtocolDesignRationaleTag = 0x00189910
	// PatientSpecificationSequenceTag is tag for Patient Specification Sequence
	PatientSpecificationSequenceTag = 0x00189911
	// ModelSpecificationSequenceTag is tag for Model Specification Sequence
	ModelSpecificationSequenceTag = 0x00189912
	// ParametersSpecificationSequenceTag is tag for Parameters Specification Sequence
	ParametersSpecificationSequenceTag = 0x00189913
	// InstructionSequenceTag is tag for Instruction Sequence
	InstructionSequenceTag = 0x00189914
	// InstructionIndexTag is tag for Instruction Index
	InstructionIndexTag = 0x00189915
	// InstructionTextTag is tag for Instruction Text
	InstructionTextTag = 0x00189916
	// InstructionDescriptionTag is tag for Instruction Description
	InstructionDescriptionTag = 0x00189917
	// InstructionPerformedFlagTag is tag for Instruction Performed Flag
	InstructionPerformedFlagTag = 0x00189918
	// InstructionPerformedDateTimeTag is tag for Instruction Performed DateTime
	InstructionPerformedDateTimeTag = 0x00189919
	// InstructionPerformanceCommentTag is tag for Instruction Performance Comment
	InstructionPerformanceCommentTag = 0x0018991A
	// PatientPositioningInstructionSequenceTag is tag for Patient Positioning Instruction Sequence
	PatientPositioningInstructionSequenceTag = 0x0018991B
	// PositioningMethodCodeSequenceTag is tag for Positioning Method Code Sequence
	PositioningMethodCodeSequenceTag = 0x0018991C
	// PositioningLandmarkSequenceTag is tag for Positioning Landmark Sequence
	PositioningLandmarkSequenceTag = 0x0018991D
	// TargetFrameOfReferenceUIDTag is tag for Target Frame of Reference UID
	TargetFrameOfReferenceUIDTag = 0x0018991E
	// AcquisitionProtocolElementSpecificationSequenceTag is tag for Acquisition Protocol Element Specification Sequence
	AcquisitionProtocolElementSpecificationSequenceTag = 0x0018991F
	// AcquisitionProtocolElementSequenceTag is tag for Acquisition Protocol Element Sequence
	AcquisitionProtocolElementSequenceTag = 0x00189920
	// ProtocolElementNumberTag is tag for Protocol Element Number
	ProtocolElementNumberTag = 0x00189921
	// ProtocolElementNameTag is tag for Protocol Element Name
	ProtocolElementNameTag = 0x00189922
	// ProtocolElementCharacteristicsSummaryTag is tag for Protocol Element Characteristics Summary
	ProtocolElementCharacteristicsSummaryTag = 0x00189923
	// ProtocolElementPurposeTag is tag for Protocol Element Purpose
	ProtocolElementPurposeTag = 0x00189924
	// AcquisitionMotionTag is tag for Acquisition Motion
	AcquisitionMotionTag = 0x00189930
	// AcquisitionStartLocationSequenceTag is tag for Acquisition Start Location Sequence
	AcquisitionStartLocationSequenceTag = 0x00189931
	// AcquisitionEndLocationSequenceTag is tag for Acquisition End Location Sequence
	AcquisitionEndLocationSequenceTag = 0x00189932
	// ReconstructionProtocolElementSpecificationSequenceTag is tag for Reconstruction Protocol Element Specification Sequence
	ReconstructionProtocolElementSpecificationSequenceTag = 0x00189933
	// ReconstructionProtocolElementSequenceTag is tag for Reconstruction Protocol Element Sequence
	ReconstructionProtocolElementSequenceTag = 0x00189934
	// StorageProtocolElementSpecificationSequenceTag is tag for Storage Protocol Element Specification Sequence
	StorageProtocolElementSpecificationSequenceTag = 0x00189935
	// StorageProtocolElementSequenceTag is tag for Storage Protocol Element Sequence
	StorageProtocolElementSequenceTag = 0x00189936
	// RequestedSeriesDescriptionTag is tag for Requested Series Description
	RequestedSeriesDescriptionTag = 0x00189937
	// SourceAcquisitionProtocolElementNumberTag is tag for Source Acquisition Protocol Element Number
	SourceAcquisitionProtocolElementNumberTag = 0x00189938
	// SourceAcquisitionBeamNumberTag is tag for Source Acquisition Beam Number
	SourceAcquisitionBeamNumberTag = 0x00189939
	// SourceReconstructionProtocolElementNumberTag is tag for Source Reconstruction Protocol Element Number
	SourceReconstructionProtocolElementNumberTag = 0x0018993A
	// ReconstructionStartLocationSequenceTag is tag for Reconstruction Start Location Sequence
	ReconstructionStartLocationSequenceTag = 0x0018993B
	// ReconstructionEndLocationSequenceTag is tag for Reconstruction End Location Sequence
	ReconstructionEndLocationSequenceTag = 0x0018993C
	// ReconstructionAlgorithmSequenceTag is tag for Reconstruction Algorithm Sequence
	ReconstructionAlgorithmSequenceTag = 0x0018993D
	// ReconstructionTargetCenterLocationSequenceTag is tag for Reconstruction Target Center Location Sequence
	ReconstructionTargetCenterLocationSequenceTag = 0x0018993E
	// ImageFilterDescriptionTag is tag for Image Filter Description
	ImageFilterDescriptionTag = 0x00189941
	// CTDIvolNotificationTriggerTag is tag for CTDIvol Notification Trigger
	CTDIvolNotificationTriggerTag = 0x00189942
	// DLPNotificationTriggerTag is tag for DLP Notification Trigger
	DLPNotificationTriggerTag = 0x00189943
	// AutoKVPSelectionTypeTag is tag for Auto KVP Selection Type
	AutoKVPSelectionTypeTag = 0x00189944
	// AutoKVPUpperBoundTag is tag for Auto KVP Upper Bound
	AutoKVPUpperBoundTag = 0x00189945
	// AutoKVPLowerBoundTag is tag for Auto KVP Lower Bound
	AutoKVPLowerBoundTag = 0x00189946
	// ProtocolDefinedPatientPositionTag is tag for Protocol Defined Patient Position
	ProtocolDefinedPatientPositionTag = 0x00189947
	// ContributingEquipmentSequenceTag is tag for Contributing Equipment Sequence
	ContributingEquipmentSequenceTag = 0x0018A001
	// ContributionDateTimeTag is tag for Contribution DateTime
	ContributionDateTimeTag = 0x0018A002
	// ContributionDescriptionTag is tag for Contribution Description
	ContributionDescriptionTag = 0x0018A003
	// StudyInstanceUIDTag is tag for Study Instance UID
	StudyInstanceUIDTag = 0x0020000D
	// SeriesInstanceUIDTag is tag for Series Instance UID
	SeriesInstanceUIDTag = 0x0020000E
	// StudyIDTag is tag for Study ID
	StudyIDTag = 0x00200010
	// SeriesNumberTag is tag for Series Number
	SeriesNumberTag = 0x00200011
	// AcquisitionNumberTag is tag for Acquisition Number
	AcquisitionNumberTag = 0x00200012
	// InstanceNumberTag is tag for Instance Number
	InstanceNumberTag = 0x00200013
	// IsotopeNumberTag is tag for Isotope Number
	IsotopeNumberTag = 0x00200014
	// PhaseNumberTag is tag for Phase Number
	PhaseNumberTag = 0x00200015
	// IntervalNumberTag is tag for Interval Number
	IntervalNumberTag = 0x00200016
	// TimeSlotNumberTag is tag for Time Slot Number
	TimeSlotNumberTag = 0x00200017
	// AngleNumberTag is tag for Angle Number
	AngleNumberTag = 0x00200018
	// ItemNumberTag is tag for Item Number
	ItemNumberTag = 0x00200019
	// PatientOrientationTag is tag for Patient Orientation
	PatientOrientationTag = 0x00200020
	// OverlayNumberTag is tag for Overlay Number
	OverlayNumberTag = 0x00200022
	// CurveNumberTag is tag for Curve Number
	CurveNumberTag = 0x00200024
	// LUTNumberTag is tag for LUT Number
	LUTNumberTag = 0x00200026
	// ImagePositionTag is tag for Image Position
	ImagePositionTag = 0x00200030
	// ImagePositionPatientTag is tag for Image Position (Patient)
	ImagePositionPatientTag = 0x00200032
	// ImageOrientationTag is tag for Image Orientation
	ImageOrientationTag = 0x00200035
	// ImageOrientationPatientTag is tag for Image Orientation (Patient)
	ImageOrientationPatientTag = 0x00200037
	// LocationTag is tag for Location
	LocationTag = 0x00200050
	// FrameOfReferenceUIDTag is tag for Frame of Reference UID
	FrameOfReferenceUIDTag = 0x00200052
	// LateralityTag is tag for Laterality
	LateralityTag = 0x00200060
	// ImageLateralityTag is tag for Image Laterality
	ImageLateralityTag = 0x00200062
	// ImageGeometryTypeTag is tag for Image Geometry Type
	ImageGeometryTypeTag = 0x00200070
	// MaskingImageTag is tag for Masking Image
	MaskingImageTag = 0x00200080
	// ReportNumberTag is tag for Report Number
	ReportNumberTag = 0x002000AA
	// TemporalPositionIdentifierTag is tag for Temporal Position Identifier
	TemporalPositionIdentifierTag = 0x00200100
	// NumberOfTemporalPositionsTag is tag for Number of Temporal Positions
	NumberOfTemporalPositionsTag = 0x00200105
	// TemporalResolutionTag is tag for Temporal Resolution
	TemporalResolutionTag = 0x00200110
	// SynchronizationFrameOfReferenceUIDTag is tag for Synchronization Frame of Reference UID
	SynchronizationFrameOfReferenceUIDTag = 0x00200200
	// SOPInstanceUIDOfConcatenationSourceTag is tag for SOP Instance UID of Concatenation Source
	SOPInstanceUIDOfConcatenationSourceTag = 0x00200242
	// SeriesInStudyTag is tag for Series in Study
	SeriesInStudyTag = 0x00201000
	// AcquisitionsInSeriesTag is tag for Acquisitions in Series
	AcquisitionsInSeriesTag = 0x00201001
	// ImagesInAcquisitionTag is tag for Images in Acquisition
	ImagesInAcquisitionTag = 0x00201002
	// ImagesInSeriesTag is tag for Images in Series
	ImagesInSeriesTag = 0x00201003
	// AcquisitionsInStudyTag is tag for Acquisitions in Study
	AcquisitionsInStudyTag = 0x00201004
	// ImagesInStudyTag is tag for Images in Study
	ImagesInStudyTag = 0x00201005
	// ReferenceTag is tag for Reference
	ReferenceTag = 0x00201020
	// TargetPositionReferenceIndicatorTag is tag for Target Position Reference Indicator
	TargetPositionReferenceIndicatorTag = 0x0020103F
	// PositionReferenceIndicatorTag is tag for Position Reference Indicator
	PositionReferenceIndicatorTag = 0x00201040
	// SliceLocationTag is tag for Slice Location
	SliceLocationTag = 0x00201041
	// OtherStudyNumbersTag is tag for Other Study Numbers
	OtherStudyNumbersTag = 0x00201070
	// NumberOfPatientRelatedStudiesTag is tag for Number of Patient Related Studies
	NumberOfPatientRelatedStudiesTag = 0x00201200
	// NumberOfPatientRelatedSeriesTag is tag for Number of Patient Related Series
	NumberOfPatientRelatedSeriesTag = 0x00201202
	// NumberOfPatientRelatedInstancesTag is tag for Number of Patient Related Instances
	NumberOfPatientRelatedInstancesTag = 0x00201204
	// NumberOfStudyRelatedSeriesTag is tag for Number of Study Related Series
	NumberOfStudyRelatedSeriesTag = 0x00201206
	// NumberOfStudyRelatedInstancesTag is tag for Number of Study Related Instances
	NumberOfStudyRelatedInstancesTag = 0x00201208
	// NumberOfSeriesRelatedInstancesTag is tag for Number of Series Related Instances
	NumberOfSeriesRelatedInstancesTag = 0x00201209
	// SourceImageIDsTag is tag for Source Image IDs
	SourceImageIDsTag = 0x00203100
	// ModifyingDeviceIDTag is tag for Modifying Device ID
	ModifyingDeviceIDTag = 0x00203401
	// ModifiedImageIDTag is tag for Modified Image ID
	ModifiedImageIDTag = 0x00203402
	// ModifiedImageDateTag is tag for Modified Image Date
	ModifiedImageDateTag = 0x00203403
	// ModifyingDeviceManufacturerTag is tag for Modifying Device Manufacturer
	ModifyingDeviceManufacturerTag = 0x00203404
	// ModifiedImageTimeTag is tag for Modified Image Time
	ModifiedImageTimeTag = 0x00203405
	// ModifiedImageDescriptionTag is tag for Modified Image Description
	ModifiedImageDescriptionTag = 0x00203406
	// ImageCommentsTag is tag for Image Comments
	ImageCommentsTag = 0x00204000
	// OriginalImageIdentificationTag is tag for Original Image Identification
	OriginalImageIdentificationTag = 0x00205000
	// OriginalImageIdentificationNomenclatureTag is tag for Original Image Identification Nomenclature
	OriginalImageIdentificationNomenclatureTag = 0x00205002
	// StackIDTag is tag for Stack ID
	StackIDTag = 0x00209056
	// InStackPositionNumberTag is tag for In-Stack Position Number
	InStackPositionNumberTag = 0x00209057
	// FrameAnatomySequenceTag is tag for Frame Anatomy Sequence
	FrameAnatomySequenceTag = 0x00209071
	// FrameLateralityTag is tag for Frame Laterality
	FrameLateralityTag = 0x00209072
	// FrameContentSequenceTag is tag for Frame Content Sequence
	FrameContentSequenceTag = 0x00209111
	// PlanePositionSequenceTag is tag for Plane Position Sequence
	PlanePositionSequenceTag = 0x00209113
	// PlaneOrientationSequenceTag is tag for Plane Orientation Sequence
	PlaneOrientationSequenceTag = 0x00209116
	// TemporalPositionIndexTag is tag for Temporal Position Index
	TemporalPositionIndexTag = 0x00209128
	// NominalCardiacTriggerDelayTimeTag is tag for Nominal Cardiac Trigger Delay Time
	NominalCardiacTriggerDelayTimeTag = 0x00209153
	// NominalCardiacTriggerTimePriorToRPeakTag is tag for Nominal Cardiac Trigger Time Prior To R-Peak
	NominalCardiacTriggerTimePriorToRPeakTag = 0x00209154
	// ActualCardiacTriggerTimePriorToRPeakTag is tag for Actual Cardiac Trigger Time Prior To R-Peak
	ActualCardiacTriggerTimePriorToRPeakTag = 0x00209155
	// FrameAcquisitionNumberTag is tag for Frame Acquisition Number
	FrameAcquisitionNumberTag = 0x00209156
	// DimensionIndexValuesTag is tag for Dimension Index Values
	DimensionIndexValuesTag = 0x00209157
	// FrameCommentsTag is tag for Frame Comments
	FrameCommentsTag = 0x00209158
	// ConcatenationUIDTag is tag for Concatenation UID
	ConcatenationUIDTag = 0x00209161
	// InConcatenationNumberTag is tag for In-concatenation Number
	InConcatenationNumberTag = 0x00209162
	// InConcatenationTotalNumberTag is tag for In-concatenation Total Number
	InConcatenationTotalNumberTag = 0x00209163
	// DimensionOrganizationUIDTag is tag for Dimension Organization UID
	DimensionOrganizationUIDTag = 0x00209164
	// DimensionIndexPointerTag is tag for Dimension Index Pointer
	DimensionIndexPointerTag = 0x00209165
	// FunctionalGroupPointerTag is tag for Functional Group Pointer
	FunctionalGroupPointerTag = 0x00209167
	// UnassignedSharedConvertedAttributesSequenceTag is tag for Unassigned Shared Converted Attributes Sequence
	UnassignedSharedConvertedAttributesSequenceTag = 0x00209170
	// UnassignedPerFrameConvertedAttributesSequenceTag is tag for Unassigned Per-Frame Converted Attributes Sequence
	UnassignedPerFrameConvertedAttributesSequenceTag = 0x00209171
	// ConversionSourceAttributesSequenceTag is tag for Conversion Source Attributes Sequence
	ConversionSourceAttributesSequenceTag = 0x00209172
	// DimensionIndexPrivateCreatorTag is tag for Dimension Index Private Creator
	DimensionIndexPrivateCreatorTag = 0x00209213
	// DimensionOrganizationSequenceTag is tag for Dimension Organization Sequence
	DimensionOrganizationSequenceTag = 0x00209221
	// DimensionIndexSequenceTag is tag for Dimension Index Sequence
	DimensionIndexSequenceTag = 0x00209222
	// ConcatenationFrameOffsetNumberTag is tag for Concatenation Frame Offset Number
	ConcatenationFrameOffsetNumberTag = 0x00209228
	// FunctionalGroupPrivateCreatorTag is tag for Functional Group Private Creator
	FunctionalGroupPrivateCreatorTag = 0x00209238
	// NominalPercentageOfCardiacPhaseTag is tag for Nominal Percentage of Cardiac Phase
	NominalPercentageOfCardiacPhaseTag = 0x00209241
	// NominalPercentageOfRespiratoryPhaseTag is tag for Nominal Percentage of Respiratory Phase
	NominalPercentageOfRespiratoryPhaseTag = 0x00209245
	// StartingRespiratoryAmplitudeTag is tag for Starting Respiratory Amplitude
	StartingRespiratoryAmplitudeTag = 0x00209246
	// StartingRespiratoryPhaseTag is tag for Starting Respiratory Phase
	StartingRespiratoryPhaseTag = 0x00209247
	// EndingRespiratoryAmplitudeTag is tag for Ending Respiratory Amplitude
	EndingRespiratoryAmplitudeTag = 0x00209248
	// EndingRespiratoryPhaseTag is tag for Ending Respiratory Phase
	EndingRespiratoryPhaseTag = 0x00209249
	// RespiratoryTriggerTypeTag is tag for Respiratory Trigger Type
	RespiratoryTriggerTypeTag = 0x00209250
	// RRIntervalTimeNominalTag is tag for R-R Interval Time Nominal
	RRIntervalTimeNominalTag = 0x00209251
	// ActualCardiacTriggerDelayTimeTag is tag for Actual Cardiac Trigger Delay Time
	ActualCardiacTriggerDelayTimeTag = 0x00209252
	// RespiratorySynchronizationSequenceTag is tag for Respiratory Synchronization Sequence
	RespiratorySynchronizationSequenceTag = 0x00209253
	// RespiratoryIntervalTimeTag is tag for Respiratory Interval Time
	RespiratoryIntervalTimeTag = 0x00209254
	// NominalRespiratoryTriggerDelayTimeTag is tag for Nominal Respiratory Trigger Delay Time
	NominalRespiratoryTriggerDelayTimeTag = 0x00209255
	// RespiratoryTriggerDelayThresholdTag is tag for Respiratory Trigger Delay Threshold
	RespiratoryTriggerDelayThresholdTag = 0x00209256
	// ActualRespiratoryTriggerDelayTimeTag is tag for Actual Respiratory Trigger Delay Time
	ActualRespiratoryTriggerDelayTimeTag = 0x00209257
	// ImagePositionVolumeTag is tag for Image Position (Volume)
	ImagePositionVolumeTag = 0x00209301
	// ImageOrientationVolumeTag is tag for Image Orientation (Volume)
	ImageOrientationVolumeTag = 0x00209302
	// UltrasoundAcquisitionGeometryTag is tag for Ultrasound Acquisition Geometry
	UltrasoundAcquisitionGeometryTag = 0x00209307
	// ApexPositionTag is tag for Apex Position
	ApexPositionTag = 0x00209308
	// VolumeToTransducerMappingMatrixTag is tag for Volume to Transducer Mapping Matrix
	VolumeToTransducerMappingMatrixTag = 0x00209309
	// VolumeToTableMappingMatrixTag is tag for Volume to Table Mapping Matrix
	VolumeToTableMappingMatrixTag = 0x0020930A
	// VolumeToTransducerRelationshipTag is tag for Volume to Transducer Relationship
	VolumeToTransducerRelationshipTag = 0x0020930B
	// PatientFrameOfReferenceSourceTag is tag for Patient Frame of Reference Source
	PatientFrameOfReferenceSourceTag = 0x0020930C
	// TemporalPositionTimeOffsetTag is tag for Temporal Position Time Offset
	TemporalPositionTimeOffsetTag = 0x0020930D
	// PlanePositionVolumeSequenceTag is tag for Plane Position (Volume) Sequence
	PlanePositionVolumeSequenceTag = 0x0020930E
	// PlaneOrientationVolumeSequenceTag is tag for Plane Orientation (Volume) Sequence
	PlaneOrientationVolumeSequenceTag = 0x0020930F
	// TemporalPositionSequenceTag is tag for Temporal Position Sequence
	TemporalPositionSequenceTag = 0x00209310
	// DimensionOrganizationTypeTag is tag for Dimension Organization Type
	DimensionOrganizationTypeTag = 0x00209311
	// VolumeFrameOfReferenceUIDTag is tag for Volume Frame of Reference UID
	VolumeFrameOfReferenceUIDTag = 0x00209312
	// TableFrameOfReferenceUIDTag is tag for Table Frame of Reference UID
	TableFrameOfReferenceUIDTag = 0x00209313
	// DimensionDescriptionLabelTag is tag for Dimension Description Label
	DimensionDescriptionLabelTag = 0x00209421
	// PatientOrientationInFrameSequenceTag is tag for Patient Orientation in Frame Sequence
	PatientOrientationInFrameSequenceTag = 0x00209450
	// FrameLabelTag is tag for Frame Label
	FrameLabelTag = 0x00209453
	// AcquisitionIndexTag is tag for Acquisition Index
	AcquisitionIndexTag = 0x00209518
	// ContributingSOPInstancesReferenceSequenceTag is tag for Contributing SOP Instances Reference Sequence
	ContributingSOPInstancesReferenceSequenceTag = 0x00209529
	// ReconstructionIndexTag is tag for Reconstruction Index
	ReconstructionIndexTag = 0x00209536
	// LightPathFilterPassThroughWavelengthTag is tag for Light Path Filter Pass-Through Wavelength
	LightPathFilterPassThroughWavelengthTag = 0x00220001
	// LightPathFilterPassBandTag is tag for Light Path Filter Pass Band
	LightPathFilterPassBandTag = 0x00220002
	// ImagePathFilterPassThroughWavelengthTag is tag for Image Path Filter Pass-Through Wavelength
	ImagePathFilterPassThroughWavelengthTag = 0x00220003
	// ImagePathFilterPassBandTag is tag for Image Path Filter Pass Band
	ImagePathFilterPassBandTag = 0x00220004
	// PatientEyeMovementCommandedTag is tag for Patient Eye Movement Commanded
	PatientEyeMovementCommandedTag = 0x00220005
	// PatientEyeMovementCommandCodeSequenceTag is tag for Patient Eye Movement Command Code Sequence
	PatientEyeMovementCommandCodeSequenceTag = 0x00220006
	// SphericalLensPowerTag is tag for Spherical Lens Power
	SphericalLensPowerTag = 0x00220007
	// CylinderLensPowerTag is tag for Cylinder Lens Power
	CylinderLensPowerTag = 0x00220008
	// CylinderAxisTag is tag for Cylinder Axis
	CylinderAxisTag = 0x00220009
	// EmmetropicMagnificationTag is tag for Emmetropic Magnification
	EmmetropicMagnificationTag = 0x0022000A
	// IntraOcularPressureTag is tag for Intra Ocular Pressure
	IntraOcularPressureTag = 0x0022000B
	// HorizontalFieldOfViewTag is tag for Horizontal Field of View
	HorizontalFieldOfViewTag = 0x0022000C
	// PupilDilatedTag is tag for Pupil Dilated
	PupilDilatedTag = 0x0022000D
	// DegreeOfDilationTag is tag for Degree of Dilation
	DegreeOfDilationTag = 0x0022000E
	// StereoBaselineAngleTag is tag for Stereo Baseline Angle
	StereoBaselineAngleTag = 0x00220010
	// StereoBaselineDisplacementTag is tag for Stereo Baseline Displacement
	StereoBaselineDisplacementTag = 0x00220011
	// StereoHorizontalPixelOffsetTag is tag for Stereo Horizontal Pixel Offset
	StereoHorizontalPixelOffsetTag = 0x00220012
	// StereoVerticalPixelOffsetTag is tag for Stereo Vertical Pixel Offset
	StereoVerticalPixelOffsetTag = 0x00220013
	// StereoRotationTag is tag for Stereo Rotation
	StereoRotationTag = 0x00220014
	// AcquisitionDeviceTypeCodeSequenceTag is tag for Acquisition Device Type Code Sequence
	AcquisitionDeviceTypeCodeSequenceTag = 0x00220015
	// IlluminationTypeCodeSequenceTag is tag for Illumination Type Code Sequence
	IlluminationTypeCodeSequenceTag = 0x00220016
	// LightPathFilterTypeStackCodeSequenceTag is tag for Light Path Filter Type Stack Code Sequence
	LightPathFilterTypeStackCodeSequenceTag = 0x00220017
	// ImagePathFilterTypeStackCodeSequenceTag is tag for Image Path Filter Type Stack Code Sequence
	ImagePathFilterTypeStackCodeSequenceTag = 0x00220018
	// LensesCodeSequenceTag is tag for Lenses Code Sequence
	LensesCodeSequenceTag = 0x00220019
	// ChannelDescriptionCodeSequenceTag is tag for Channel Description Code Sequence
	ChannelDescriptionCodeSequenceTag = 0x0022001A
	// RefractiveStateSequenceTag is tag for Refractive State Sequence
	RefractiveStateSequenceTag = 0x0022001B
	// MydriaticAgentCodeSequenceTag is tag for Mydriatic Agent Code Sequence
	MydriaticAgentCodeSequenceTag = 0x0022001C
	// RelativeImagePositionCodeSequenceTag is tag for Relative Image Position Code Sequence
	RelativeImagePositionCodeSequenceTag = 0x0022001D
	// CameraAngleOfViewTag is tag for Camera Angle of View
	CameraAngleOfViewTag = 0x0022001E
	// StereoPairsSequenceTag is tag for Stereo Pairs Sequence
	StereoPairsSequenceTag = 0x00220020
	// LeftImageSequenceTag is tag for Left Image Sequence
	LeftImageSequenceTag = 0x00220021
	// RightImageSequenceTag is tag for Right Image Sequence
	RightImageSequenceTag = 0x00220022
	// StereoPairsPresentTag is tag for Stereo Pairs Present
	StereoPairsPresentTag = 0x00220028
	// AxialLengthOfTheEyeTag is tag for Axial Length of the Eye
	AxialLengthOfTheEyeTag = 0x00220030
	// OphthalmicFrameLocationSequenceTag is tag for Ophthalmic Frame Location Sequence
	OphthalmicFrameLocationSequenceTag = 0x00220031
	// ReferenceCoordinatesTag is tag for Reference Coordinates
	ReferenceCoordinatesTag = 0x00220032
	// DepthSpatialResolutionTag is tag for Depth Spatial Resolution
	DepthSpatialResolutionTag = 0x00220035
	// MaximumDepthDistortionTag is tag for Maximum Depth Distortion
	MaximumDepthDistortionTag = 0x00220036
	// AlongScanSpatialResolutionTag is tag for Along-scan Spatial Resolution
	AlongScanSpatialResolutionTag = 0x00220037
	// MaximumAlongScanDistortionTag is tag for Maximum Along-scan Distortion
	MaximumAlongScanDistortionTag = 0x00220038
	// OphthalmicImageOrientationTag is tag for Ophthalmic Image Orientation
	OphthalmicImageOrientationTag = 0x00220039
	// DepthOfTransverseImageTag is tag for Depth of Transverse Image
	DepthOfTransverseImageTag = 0x00220041
	// MydriaticAgentConcentrationUnitsSequenceTag is tag for Mydriatic Agent Concentration Units Sequence
	MydriaticAgentConcentrationUnitsSequenceTag = 0x00220042
	// AcrossScanSpatialResolutionTag is tag for Across-scan Spatial Resolution
	AcrossScanSpatialResolutionTag = 0x00220048
	// MaximumAcrossScanDistortionTag is tag for Maximum Across-scan Distortion
	MaximumAcrossScanDistortionTag = 0x00220049
	// MydriaticAgentConcentrationTag is tag for Mydriatic Agent Concentration
	MydriaticAgentConcentrationTag = 0x0022004E
	// IlluminationWaveLengthTag is tag for Illumination Wave Length
	IlluminationWaveLengthTag = 0x00220055
	// IlluminationPowerTag is tag for Illumination Power
	IlluminationPowerTag = 0x00220056
	// IlluminationBandwidthTag is tag for Illumination Bandwidth
	IlluminationBandwidthTag = 0x00220057
	// MydriaticAgentSequenceTag is tag for Mydriatic Agent Sequence
	MydriaticAgentSequenceTag = 0x00220058
	// OphthalmicAxialMeasurementsRightEyeSequenceTag is tag for Ophthalmic Axial Measurements Right Eye Sequence
	OphthalmicAxialMeasurementsRightEyeSequenceTag = 0x00221007
	// OphthalmicAxialMeasurementsLeftEyeSequenceTag is tag for Ophthalmic Axial Measurements Left Eye Sequence
	OphthalmicAxialMeasurementsLeftEyeSequenceTag = 0x00221008
	// OphthalmicAxialMeasurementsDeviceTypeTag is tag for Ophthalmic Axial Measurements Device Type
	OphthalmicAxialMeasurementsDeviceTypeTag = 0x00221009
	// OphthalmicAxialLengthMeasurementsTypeTag is tag for Ophthalmic Axial Length Measurements Type
	OphthalmicAxialLengthMeasurementsTypeTag = 0x00221010
	// OphthalmicAxialLengthSequenceTag is tag for Ophthalmic Axial Length Sequence
	OphthalmicAxialLengthSequenceTag = 0x00221012
	// OphthalmicAxialLengthTag is tag for Ophthalmic Axial Length
	OphthalmicAxialLengthTag = 0x00221019
	// LensStatusCodeSequenceTag is tag for Lens Status Code Sequence
	LensStatusCodeSequenceTag = 0x00221024
	// VitreousStatusCodeSequenceTag is tag for Vitreous Status Code Sequence
	VitreousStatusCodeSequenceTag = 0x00221025
	// IOLFormulaCodeSequenceTag is tag for IOL Formula Code Sequence
	IOLFormulaCodeSequenceTag = 0x00221028
	// IOLFormulaDetailTag is tag for IOL Formula Detail
	IOLFormulaDetailTag = 0x00221029
	// KeratometerIndexTag is tag for Keratometer Index
	KeratometerIndexTag = 0x00221033
	// SourceOfOphthalmicAxialLengthCodeSequenceTag is tag for Source of Ophthalmic Axial Length Code Sequence
	SourceOfOphthalmicAxialLengthCodeSequenceTag = 0x00221035
	// SourceOfCornealSizeDataCodeSequenceTag is tag for Source of Corneal Size Data Code Sequence
	SourceOfCornealSizeDataCodeSequenceTag = 0x00221036
	// TargetRefractionTag is tag for Target Refraction
	TargetRefractionTag = 0x00221037
	// RefractiveProcedureOccurredTag is tag for Refractive Procedure Occurred
	RefractiveProcedureOccurredTag = 0x00221039
	// RefractiveSurgeryTypeCodeSequenceTag is tag for Refractive Surgery Type Code Sequence
	RefractiveSurgeryTypeCodeSequenceTag = 0x00221040
	// OphthalmicUltrasoundMethodCodeSequenceTag is tag for Ophthalmic Ultrasound Method Code Sequence
	OphthalmicUltrasoundMethodCodeSequenceTag = 0x00221044
	// SurgicallyInducedAstigmatismSequenceTag is tag for Surgically Induced Astigmatism Sequence
	SurgicallyInducedAstigmatismSequenceTag = 0x00221045
	// TypeOfOpticalCorrectionTag is tag for Type of Optical Correction
	TypeOfOpticalCorrectionTag = 0x00221046
	// ToricIOLPowerSequenceTag is tag for Toric IOL Power Sequence
	ToricIOLPowerSequenceTag = 0x00221047
	// PredictedToricErrorSequenceTag is tag for Predicted Toric Error Sequence
	PredictedToricErrorSequenceTag = 0x00221048
	// PreSelectedForImplantationTag is tag for Pre-Selected for Implantation
	PreSelectedForImplantationTag = 0x00221049
	// ToricIOLPowerForExactEmmetropiaSequenceTag is tag for Toric IOL Power for Exact Emmetropia Sequence
	ToricIOLPowerForExactEmmetropiaSequenceTag = 0x0022104A
	// ToricIOLPowerForExactTargetRefractionSequenceTag is tag for Toric IOL Power for Exact Target Refraction Sequence
	ToricIOLPowerForExactTargetRefractionSequenceTag = 0x0022104B
	// OphthalmicAxialLengthMeasurementsSequenceTag is tag for Ophthalmic Axial Length Measurements Sequence
	OphthalmicAxialLengthMeasurementsSequenceTag = 0x00221050
	// IOLPowerTag is tag for IOL Power
	IOLPowerTag = 0x00221053
	// PredictedRefractiveErrorTag is tag for Predicted Refractive Error
	PredictedRefractiveErrorTag = 0x00221054
	// OphthalmicAxialLengthVelocityTag is tag for Ophthalmic Axial Length Velocity
	OphthalmicAxialLengthVelocityTag = 0x00221059
	// LensStatusDescriptionTag is tag for Lens Status Description
	LensStatusDescriptionTag = 0x00221065
	// VitreousStatusDescriptionTag is tag for Vitreous Status Description
	VitreousStatusDescriptionTag = 0x00221066
	// IOLPowerSequenceTag is tag for IOL Power Sequence
	IOLPowerSequenceTag = 0x00221090
	// LensConstantSequenceTag is tag for Lens Constant Sequence
	LensConstantSequenceTag = 0x00221092
	// IOLManufacturerTag is tag for IOL Manufacturer
	IOLManufacturerTag = 0x00221093
	// LensConstantDescriptionTag is tag for Lens Constant Description
	LensConstantDescriptionTag = 0x00221094
	// ImplantNameTag is tag for Implant Name
	ImplantNameTag = 0x00221095
	// KeratometryMeasurementTypeCodeSequenceTag is tag for Keratometry Measurement Type Code Sequence
	KeratometryMeasurementTypeCodeSequenceTag = 0x00221096
	// ImplantPartNumberTag is tag for Implant Part Number
	ImplantPartNumberTag = 0x00221097
	// ReferencedOphthalmicAxialMeasurementsSequenceTag is tag for Referenced Ophthalmic Axial Measurements Sequence
	ReferencedOphthalmicAxialMeasurementsSequenceTag = 0x00221100
	// OphthalmicAxialLengthMeasurementsSegmentNameCodeSequenceTag is tag for Ophthalmic Axial Length Measurements Segment Name Code Sequence
	OphthalmicAxialLengthMeasurementsSegmentNameCodeSequenceTag = 0x00221101
	// RefractiveErrorBeforeRefractiveSurgeryCodeSequenceTag is tag for Refractive Error Before Refractive Surgery Code Sequence
	RefractiveErrorBeforeRefractiveSurgeryCodeSequenceTag = 0x00221103
	// IOLPowerForExactEmmetropiaTag is tag for IOL Power For Exact Emmetropia
	IOLPowerForExactEmmetropiaTag = 0x00221121
	// IOLPowerForExactTargetRefractionTag is tag for IOL Power For Exact Target Refraction
	IOLPowerForExactTargetRefractionTag = 0x00221122
	// AnteriorChamberDepthDefinitionCodeSequenceTag is tag for Anterior Chamber Depth Definition Code Sequence
	AnteriorChamberDepthDefinitionCodeSequenceTag = 0x00221125
	// LensThicknessSequenceTag is tag for Lens Thickness Sequence
	LensThicknessSequenceTag = 0x00221127
	// AnteriorChamberDepthSequenceTag is tag for Anterior Chamber Depth Sequence
	AnteriorChamberDepthSequenceTag = 0x00221128
	// CalculationCommentSequenceTag is tag for Calculation Comment Sequence
	CalculationCommentSequenceTag = 0x0022112A
	// CalculationCommentTypeTag is tag for Calculation Comment Type
	CalculationCommentTypeTag = 0x0022112B
	// CalculationCommentTag is tag for Calculation Comment
	CalculationCommentTag = 0x0022112C
	// LensThicknessTag is tag for Lens Thickness
	LensThicknessTag = 0x00221130
	// AnteriorChamberDepthTag is tag for Anterior Chamber Depth
	AnteriorChamberDepthTag = 0x00221131
	// SourceOfLensThicknessDataCodeSequenceTag is tag for Source of Lens Thickness Data Code Sequence
	SourceOfLensThicknessDataCodeSequenceTag = 0x00221132
	// SourceOfAnteriorChamberDepthDataCodeSequenceTag is tag for Source of Anterior Chamber Depth Data Code Sequence
	SourceOfAnteriorChamberDepthDataCodeSequenceTag = 0x00221133
	// SourceOfRefractiveMeasurementsSequenceTag is tag for Source of Refractive Measurements Sequence
	SourceOfRefractiveMeasurementsSequenceTag = 0x00221134
	// SourceOfRefractiveMeasurementsCodeSequenceTag is tag for Source of Refractive Measurements Code Sequence
	SourceOfRefractiveMeasurementsCodeSequenceTag = 0x00221135
	// OphthalmicAxialLengthMeasurementModifiedTag is tag for Ophthalmic Axial Length Measurement Modified
	OphthalmicAxialLengthMeasurementModifiedTag = 0x00221140
	// OphthalmicAxialLengthDataSourceCodeSequenceTag is tag for Ophthalmic Axial Length Data Source Code Sequence
	OphthalmicAxialLengthDataSourceCodeSequenceTag = 0x00221150
	// OphthalmicAxialLengthAcquisitionMethodCodeSequenceTag is tag for Ophthalmic Axial Length Acquisition Method Code Sequence
	OphthalmicAxialLengthAcquisitionMethodCodeSequenceTag = 0x00221153
	// SignalToNoiseRatioTag is tag for Signal to Noise Ratio
	SignalToNoiseRatioTag = 0x00221155
	// OphthalmicAxialLengthDataSourceDescriptionTag is tag for Ophthalmic Axial Length Data Source Description
	OphthalmicAxialLengthDataSourceDescriptionTag = 0x00221159
	// OphthalmicAxialLengthMeasurementsTotalLengthSequenceTag is tag for Ophthalmic Axial Length Measurements Total Length Sequence
	OphthalmicAxialLengthMeasurementsTotalLengthSequenceTag = 0x00221210
	// OphthalmicAxialLengthMeasurementsSegmentalLengthSequenceTag is tag for Ophthalmic Axial Length Measurements Segmental Length Sequence
	OphthalmicAxialLengthMeasurementsSegmentalLengthSequenceTag = 0x00221211
	// OphthalmicAxialLengthMeasurementsLengthSummationSequenceTag is tag for Ophthalmic Axial Length Measurements Length Summation Sequence
	OphthalmicAxialLengthMeasurementsLengthSummationSequenceTag = 0x00221212
	// UltrasoundOphthalmicAxialLengthMeasurementsSequenceTag is tag for Ultrasound Ophthalmic Axial Length Measurements Sequence
	UltrasoundOphthalmicAxialLengthMeasurementsSequenceTag = 0x00221220
	// OpticalOphthalmicAxialLengthMeasurementsSequenceTag is tag for Optical Ophthalmic Axial Length Measurements Sequence
	OpticalOphthalmicAxialLengthMeasurementsSequenceTag = 0x00221225
	// UltrasoundSelectedOphthalmicAxialLengthSequenceTag is tag for Ultrasound Selected Ophthalmic Axial Length Sequence
	UltrasoundSelectedOphthalmicAxialLengthSequenceTag = 0x00221230
	// OphthalmicAxialLengthSelectionMethodCodeSequenceTag is tag for Ophthalmic Axial Length Selection Method Code Sequence
	OphthalmicAxialLengthSelectionMethodCodeSequenceTag = 0x00221250
	// OpticalSelectedOphthalmicAxialLengthSequenceTag is tag for Optical Selected Ophthalmic Axial Length Sequence
	OpticalSelectedOphthalmicAxialLengthSequenceTag = 0x00221255
	// SelectedSegmentalOphthalmicAxialLengthSequenceTag is tag for Selected Segmental Ophthalmic Axial Length Sequence
	SelectedSegmentalOphthalmicAxialLengthSequenceTag = 0x00221257
	// SelectedTotalOphthalmicAxialLengthSequenceTag is tag for Selected Total Ophthalmic Axial Length Sequence
	SelectedTotalOphthalmicAxialLengthSequenceTag = 0x00221260
	// OphthalmicAxialLengthQualityMetricSequenceTag is tag for Ophthalmic Axial Length Quality Metric Sequence
	OphthalmicAxialLengthQualityMetricSequenceTag = 0x00221262
	// OphthalmicAxialLengthQualityMetricTypeCodeSequenceTag is tag for Ophthalmic Axial Length Quality Metric Type Code Sequence
	OphthalmicAxialLengthQualityMetricTypeCodeSequenceTag = 0x00221265
	// OphthalmicAxialLengthQualityMetricTypeDescriptionTag is tag for Ophthalmic Axial Length Quality Metric Type Description
	OphthalmicAxialLengthQualityMetricTypeDescriptionTag = 0x00221273
	// IntraocularLensCalculationsRightEyeSequenceTag is tag for Intraocular Lens Calculations Right Eye Sequence
	IntraocularLensCalculationsRightEyeSequenceTag = 0x00221300
	// IntraocularLensCalculationsLeftEyeSequenceTag is tag for Intraocular Lens Calculations Left Eye Sequence
	IntraocularLensCalculationsLeftEyeSequenceTag = 0x00221310
	// ReferencedOphthalmicAxialLengthMeasurementQCImageSequenceTag is tag for Referenced Ophthalmic Axial Length Measurement QC Image Sequence
	ReferencedOphthalmicAxialLengthMeasurementQCImageSequenceTag = 0x00221330
	// OphthalmicMappingDeviceTypeTag is tag for Ophthalmic Mapping Device Type
	OphthalmicMappingDeviceTypeTag = 0x00221415
	// AcquisitionMethodCodeSequenceTag is tag for Acquisition Method Code Sequence
	AcquisitionMethodCodeSequenceTag = 0x00221420
	// AcquisitionMethodAlgorithmSequenceTag is tag for Acquisition Method Algorithm Sequence
	AcquisitionMethodAlgorithmSequenceTag = 0x00221423
	// OphthalmicThicknessMapTypeCodeSequenceTag is tag for Ophthalmic Thickness Map Type Code Sequence
	OphthalmicThicknessMapTypeCodeSequenceTag = 0x00221436
	// OphthalmicThicknessMappingNormalsSequenceTag is tag for Ophthalmic Thickness Mapping Normals Sequence
	OphthalmicThicknessMappingNormalsSequenceTag = 0x00221443
	// RetinalThicknessDefinitionCodeSequenceTag is tag for Retinal Thickness Definition Code Sequence
	RetinalThicknessDefinitionCodeSequenceTag = 0x00221445
	// PixelValueMappingToCodedConceptSequenceTag is tag for Pixel Value Mapping to Coded Concept Sequence
	PixelValueMappingToCodedConceptSequenceTag = 0x00221450
	// MappedPixelValueTag is tag for Mapped Pixel Value
	MappedPixelValueTag = 0x00221452
	// PixelValueMappingExplanationTag is tag for Pixel Value Mapping Explanation
	PixelValueMappingExplanationTag = 0x00221454
	// OphthalmicThicknessMapQualityThresholdSequenceTag is tag for Ophthalmic Thickness Map Quality Threshold Sequence
	OphthalmicThicknessMapQualityThresholdSequenceTag = 0x00221458
	// OphthalmicThicknessMapThresholdQualityRatingTag is tag for Ophthalmic Thickness Map Threshold Quality Rating
	OphthalmicThicknessMapThresholdQualityRatingTag = 0x00221460
	// AnatomicStructureReferencePointTag is tag for Anatomic Structure Reference Point
	AnatomicStructureReferencePointTag = 0x00221463
	// RegistrationToLocalizerSequenceTag is tag for Registration to Localizer Sequence
	RegistrationToLocalizerSequenceTag = 0x00221465
	// RegisteredLocalizerUnitsTag is tag for Registered Localizer Units
	RegisteredLocalizerUnitsTag = 0x00221466
	// RegisteredLocalizerTopLeftHandCornerTag is tag for Registered Localizer Top Left Hand Corner
	RegisteredLocalizerTopLeftHandCornerTag = 0x00221467
	// RegisteredLocalizerBottomRightHandCornerTag is tag for Registered Localizer Bottom Right Hand Corner
	RegisteredLocalizerBottomRightHandCornerTag = 0x00221468
	// OphthalmicThicknessMapQualityRatingSequenceTag is tag for Ophthalmic Thickness Map Quality Rating Sequence
	OphthalmicThicknessMapQualityRatingSequenceTag = 0x00221470
	// RelevantOPTAttributesSequenceTag is tag for Relevant OPT Attributes Sequence
	RelevantOPTAttributesSequenceTag = 0x00221472
	// TransformationMethodCodeSequenceTag is tag for Transformation Method Code Sequence
	TransformationMethodCodeSequenceTag = 0x00221512
	// TransformationAlgorithmSequenceTag is tag for Transformation Algorithm Sequence
	TransformationAlgorithmSequenceTag = 0x00221513
	// OphthalmicAxialLengthMethodTag is tag for Ophthalmic Axial Length Method
	OphthalmicAxialLengthMethodTag = 0x00221515
	// OphthalmicFOVTag is tag for Ophthalmic FOV
	OphthalmicFOVTag = 0x00221517
	// TwoDimensionalToThreeDimensionalMapSequenceTag is tag for Two Dimensional to Three Dimensional Map Sequence
	TwoDimensionalToThreeDimensionalMapSequenceTag = 0x00221518
	// WideFieldOphthalmicPhotographyQualityRatingSequenceTag is tag for Wide Field Ophthalmic Photography Quality Rating Sequence
	WideFieldOphthalmicPhotographyQualityRatingSequenceTag = 0x00221525
	// WideFieldOphthalmicPhotographyQualityThresholdSequenceTag is tag for Wide Field Ophthalmic Photography Quality Threshold Sequence
	WideFieldOphthalmicPhotographyQualityThresholdSequenceTag = 0x00221526
	// WideFieldOphthalmicPhotographyThresholdQualityRatingTag is tag for Wide Field Ophthalmic Photography Threshold Quality Rating
	WideFieldOphthalmicPhotographyThresholdQualityRatingTag = 0x00221527
	// XCoordinatesCenterPixelViewAngleTag is tag for X Coordinates Center Pixel View Angle
	XCoordinatesCenterPixelViewAngleTag = 0x00221528
	// YCoordinatesCenterPixelViewAngleTag is tag for Y Coordinates Center Pixel View Angle
	YCoordinatesCenterPixelViewAngleTag = 0x00221529
	// NumberOfMapPointsTag is tag for Number of Map Points
	NumberOfMapPointsTag = 0x00221530
	// TwoDimensionalToThreeDimensionalMapDataTag is tag for Two Dimensional to Three Dimensional Map Data
	TwoDimensionalToThreeDimensionalMapDataTag = 0x00221531
	// DerivationAlgorithmSequenceTag is tag for Derivation Algorithm Sequence
	DerivationAlgorithmSequenceTag = 0x00221612
	// OphthalmicImageTypeCodeSequenceTag is tag for Ophthalmic Image Type Code Sequence
	OphthalmicImageTypeCodeSequenceTag = 0x00221615
	// OphthalmicImageTypeDescriptionTag is tag for Ophthalmic Image Type Description
	OphthalmicImageTypeDescriptionTag = 0x00221616
	// ScanPatternTypeCodeSequenceTag is tag for Scan Pattern Type Code Sequence
	ScanPatternTypeCodeSequenceTag = 0x00221618
	// ReferencedSurfaceMeshIdentificationSequenceTag is tag for Referenced Surface Mesh Identification Sequence
	ReferencedSurfaceMeshIdentificationSequenceTag = 0x00221620
	// OphthalmicVolumetricPropertiesFlagTag is tag for Ophthalmic Volumetric Properties Flag
	OphthalmicVolumetricPropertiesFlagTag = 0x00221622
	// OphthalmicAnatomicReferencePointXCoordinateTag is tag for Ophthalmic Anatomic Reference Point X-Coordinate
	OphthalmicAnatomicReferencePointXCoordinateTag = 0x00221624
	// OphthalmicAnatomicReferencePointYCoordinateTag is tag for Ophthalmic Anatomic Reference Point Y-Coordinate
	OphthalmicAnatomicReferencePointYCoordinateTag = 0x00221626
	// OphthalmicEnFaceImageQualityRatingSequenceTag is tag for Ophthalmic En Face Image Quality Rating Sequence
	OphthalmicEnFaceImageQualityRatingSequenceTag = 0x00221628
	// QualityThresholdTag is tag for Quality Threshold
	QualityThresholdTag = 0x00221630
	// OCTBscanAnalysisAcquisitionParametersSequenceTag is tag for OCT B-scan Analysis Acquisition Parameters Sequence
	OCTBscanAnalysisAcquisitionParametersSequenceTag = 0x00221640
	// NumberofBscansPerFrameTag is tag for Number of B-scans Per Frame
	NumberofBscansPerFrameTag = 0x00221642
	// BscanSlabThicknessTag is tag for B-scan Slab Thickness
	BscanSlabThicknessTag = 0x00221643
	// DistanceBetweenBscanSlabsTag is tag for Distance Between B-scan Slabs
	DistanceBetweenBscanSlabsTag = 0x00221644
	// BscanCycleTimeTag is tag for B-scan Cycle Time
	BscanCycleTimeTag = 0x00221645
	// BscanCycleTimeVectorTag is tag for B-scan Cycle Time Vector
	BscanCycleTimeVectorTag = 0x00221646
	// AscanRateTag is tag for A-scan Rate
	AscanRateTag = 0x00221649
	// BscanRateTag is tag for B-scan Rate
	BscanRateTag = 0x00221650
	// SurfaceMeshZPixelOffsetTag is tag for Surface Mesh Z-Pixel Offset
	SurfaceMeshZPixelOffsetTag = 0x00221658
	// VisualFieldHorizontalExtentTag is tag for Visual Field Horizontal Extent
	VisualFieldHorizontalExtentTag = 0x00240010
	// VisualFieldVerticalExtentTag is tag for Visual Field Vertical Extent
	VisualFieldVerticalExtentTag = 0x00240011
	// VisualFieldShapeTag is tag for Visual Field Shape
	VisualFieldShapeTag = 0x00240012
	// ScreeningTestModeCodeSequenceTag is tag for Screening Test Mode Code Sequence
	ScreeningTestModeCodeSequenceTag = 0x00240016
	// MaximumStimulusLuminanceTag is tag for Maximum Stimulus Luminance
	MaximumStimulusLuminanceTag = 0x00240018
	// BackgroundLuminanceTag is tag for Background Luminance
	BackgroundLuminanceTag = 0x00240020
	// StimulusColorCodeSequenceTag is tag for Stimulus Color Code Sequence
	StimulusColorCodeSequenceTag = 0x00240021
	// BackgroundIlluminationColorCodeSequenceTag is tag for Background Illumination Color Code Sequence
	BackgroundIlluminationColorCodeSequenceTag = 0x00240024
	// StimulusAreaTag is tag for Stimulus Area
	StimulusAreaTag = 0x00240025
	// StimulusPresentationTimeTag is tag for Stimulus Presentation Time
	StimulusPresentationTimeTag = 0x00240028
	// FixationSequenceTag is tag for Fixation Sequence
	FixationSequenceTag = 0x00240032
	// FixationMonitoringCodeSequenceTag is tag for Fixation Monitoring Code Sequence
	FixationMonitoringCodeSequenceTag = 0x00240033
	// VisualFieldCatchTrialSequenceTag is tag for Visual Field Catch Trial Sequence
	VisualFieldCatchTrialSequenceTag = 0x00240034
	// FixationCheckedQuantityTag is tag for Fixation Checked Quantity
	FixationCheckedQuantityTag = 0x00240035
	// PatientNotProperlyFixatedQuantityTag is tag for Patient Not Properly Fixated Quantity
	PatientNotProperlyFixatedQuantityTag = 0x00240036
	// PresentedVisualStimuliDataFlagTag is tag for Presented Visual Stimuli Data Flag
	PresentedVisualStimuliDataFlagTag = 0x00240037
	// NumberOfVisualStimuliTag is tag for Number of Visual Stimuli
	NumberOfVisualStimuliTag = 0x00240038
	// ExcessiveFixationLossesDataFlagTag is tag for Excessive Fixation Losses Data Flag
	ExcessiveFixationLossesDataFlagTag = 0x00240039
	// ExcessiveFixationLossesTag is tag for Excessive Fixation Losses
	ExcessiveFixationLossesTag = 0x00240040
	// StimuliRetestingQuantityTag is tag for Stimuli Retesting Quantity
	StimuliRetestingQuantityTag = 0x00240042
	// CommentsOnPatientPerformanceOfVisualFieldTag is tag for Comments on Patient's Performance of Visual Field
	CommentsOnPatientPerformanceOfVisualFieldTag = 0x00240044
	// FalseNegativesEstimateFlagTag is tag for False Negatives Estimate Flag
	FalseNegativesEstimateFlagTag = 0x00240045
	// FalseNegativesEstimateTag is tag for False Negatives Estimate
	FalseNegativesEstimateTag = 0x00240046
	// NegativeCatchTrialsQuantityTag is tag for Negative Catch Trials Quantity
	NegativeCatchTrialsQuantityTag = 0x00240048
	// FalseNegativesQuantityTag is tag for False Negatives Quantity
	FalseNegativesQuantityTag = 0x00240050
	// ExcessiveFalseNegativesDataFlagTag is tag for Excessive False Negatives Data Flag
	ExcessiveFalseNegativesDataFlagTag = 0x00240051
	// ExcessiveFalseNegativesTag is tag for Excessive False Negatives
	ExcessiveFalseNegativesTag = 0x00240052
	// FalsePositivesEstimateFlagTag is tag for False Positives Estimate Flag
	FalsePositivesEstimateFlagTag = 0x00240053
	// FalsePositivesEstimateTag is tag for False Positives Estimate
	FalsePositivesEstimateTag = 0x00240054
	// CatchTrialsDataFlagTag is tag for Catch Trials Data Flag
	CatchTrialsDataFlagTag = 0x00240055
	// PositiveCatchTrialsQuantityTag is tag for Positive Catch Trials Quantity
	PositiveCatchTrialsQuantityTag = 0x00240056
	// TestPointNormalsDataFlagTag is tag for Test Point Normals Data Flag
	TestPointNormalsDataFlagTag = 0x00240057
	// TestPointNormalsSequenceTag is tag for Test Point Normals Sequence
	TestPointNormalsSequenceTag = 0x00240058
	// GlobalDeviationProbabilityNormalsFlagTag is tag for Global Deviation Probability Normals Flag
	GlobalDeviationProbabilityNormalsFlagTag = 0x00240059
	// FalsePositivesQuantityTag is tag for False Positives Quantity
	FalsePositivesQuantityTag = 0x00240060
	// ExcessiveFalsePositivesDataFlagTag is tag for Excessive False Positives Data Flag
	ExcessiveFalsePositivesDataFlagTag = 0x00240061
	// ExcessiveFalsePositivesTag is tag for Excessive False Positives
	ExcessiveFalsePositivesTag = 0x00240062
	// VisualFieldTestNormalsFlagTag is tag for Visual Field Test Normals Flag
	VisualFieldTestNormalsFlagTag = 0x00240063
	// ResultsNormalsSequenceTag is tag for Results Normals Sequence
	ResultsNormalsSequenceTag = 0x00240064
	// AgeCorrectedSensitivityDeviationAlgorithmSequenceTag is tag for Age Corrected Sensitivity Deviation Algorithm Sequence
	AgeCorrectedSensitivityDeviationAlgorithmSequenceTag = 0x00240065
	// GlobalDeviationFromNormalTag is tag for Global Deviation From Normal
	GlobalDeviationFromNormalTag = 0x00240066
	// GeneralizedDefectSensitivityDeviationAlgorithmSequenceTag is tag for Generalized Defect Sensitivity Deviation Algorithm Sequence
	GeneralizedDefectSensitivityDeviationAlgorithmSequenceTag = 0x00240067
	// LocalizedDeviationFromNormalTag is tag for Localized Deviation From Normal
	LocalizedDeviationFromNormalTag = 0x00240068
	// PatientReliabilityIndicatorTag is tag for Patient Reliability Indicator
	PatientReliabilityIndicatorTag = 0x00240069
	// VisualFieldMeanSensitivityTag is tag for Visual Field Mean Sensitivity
	VisualFieldMeanSensitivityTag = 0x00240070
	// GlobalDeviationProbabilityTag is tag for Global Deviation Probability
	GlobalDeviationProbabilityTag = 0x00240071
	// LocalDeviationProbabilityNormalsFlagTag is tag for Local Deviation Probability Normals Flag
	LocalDeviationProbabilityNormalsFlagTag = 0x00240072
	// LocalizedDeviationProbabilityTag is tag for Localized Deviation Probability
	LocalizedDeviationProbabilityTag = 0x00240073
	// ShortTermFluctuationCalculatedTag is tag for Short Term Fluctuation Calculated
	ShortTermFluctuationCalculatedTag = 0x00240074
	// ShortTermFluctuationTag is tag for Short Term Fluctuation
	ShortTermFluctuationTag = 0x00240075
	// ShortTermFluctuationProbabilityCalculatedTag is tag for Short Term Fluctuation Probability Calculated
	ShortTermFluctuationProbabilityCalculatedTag = 0x00240076
	// ShortTermFluctuationProbabilityTag is tag for Short Term Fluctuation Probability
	ShortTermFluctuationProbabilityTag = 0x00240077
	// CorrectedLocalizedDeviationFromNormalCalculatedTag is tag for Corrected Localized Deviation From Normal Calculated
	CorrectedLocalizedDeviationFromNormalCalculatedTag = 0x00240078
	// CorrectedLocalizedDeviationFromNormalTag is tag for Corrected Localized Deviation From Normal
	CorrectedLocalizedDeviationFromNormalTag = 0x00240079
	// CorrectedLocalizedDeviationFromNormalProbabilityCalculatedTag is tag for Corrected Localized Deviation From Normal Probability Calculated
	CorrectedLocalizedDeviationFromNormalProbabilityCalculatedTag = 0x00240080
	// CorrectedLocalizedDeviationFromNormalProbabilityTag is tag for Corrected Localized Deviation From Normal Probability
	CorrectedLocalizedDeviationFromNormalProbabilityTag = 0x00240081
	// GlobalDeviationProbabilitySequenceTag is tag for Global Deviation Probability Sequence
	GlobalDeviationProbabilitySequenceTag = 0x00240083
	// LocalizedDeviationProbabilitySequenceTag is tag for Localized Deviation Probability Sequence
	LocalizedDeviationProbabilitySequenceTag = 0x00240085
	// FovealSensitivityMeasuredTag is tag for Foveal Sensitivity Measured
	FovealSensitivityMeasuredTag = 0x00240086
	// FovealSensitivityTag is tag for Foveal Sensitivity
	FovealSensitivityTag = 0x00240087
	// VisualFieldTestDurationTag is tag for Visual Field Test Duration
	VisualFieldTestDurationTag = 0x00240088
	// VisualFieldTestPointSequenceTag is tag for Visual Field Test Point Sequence
	VisualFieldTestPointSequenceTag = 0x00240089
	// VisualFieldTestPointXCoordinateTag is tag for Visual Field Test Point X-Coordinate
	VisualFieldTestPointXCoordinateTag = 0x00240090
	// VisualFieldTestPointYCoordinateTag is tag for Visual Field Test Point Y-Coordinate
	VisualFieldTestPointYCoordinateTag = 0x00240091
	// AgeCorrectedSensitivityDeviationValueTag is tag for Age Corrected Sensitivity Deviation Value
	AgeCorrectedSensitivityDeviationValueTag = 0x00240092
	// StimulusResultsTag is tag for Stimulus Results
	StimulusResultsTag = 0x00240093
	// SensitivityValueTag is tag for Sensitivity Value
	SensitivityValueTag = 0x00240094
	// RetestStimulusSeenTag is tag for Retest Stimulus Seen
	RetestStimulusSeenTag = 0x00240095
	// RetestSensitivityValueTag is tag for Retest Sensitivity Value
	RetestSensitivityValueTag = 0x00240096
	// VisualFieldTestPointNormalsSequenceTag is tag for Visual Field Test Point Normals Sequence
	VisualFieldTestPointNormalsSequenceTag = 0x00240097
	// QuantifiedDefectTag is tag for Quantified Defect
	QuantifiedDefectTag = 0x00240098
	// AgeCorrectedSensitivityDeviationProbabilityValueTag is tag for Age Corrected Sensitivity Deviation Probability Value
	AgeCorrectedSensitivityDeviationProbabilityValueTag = 0x00240100
	// GeneralizedDefectCorrectedSensitivityDeviationFlagTag is tag for Generalized Defect Corrected Sensitivity Deviation Flag
	GeneralizedDefectCorrectedSensitivityDeviationFlagTag = 0x00240102
	// GeneralizedDefectCorrectedSensitivityDeviationValueTag is tag for Generalized Defect Corrected Sensitivity Deviation Value
	GeneralizedDefectCorrectedSensitivityDeviationValueTag = 0x00240103
	// GeneralizedDefectCorrectedSensitivityDeviationProbabilityValueTag is tag for Generalized Defect Corrected Sensitivity Deviation Probability Value
	GeneralizedDefectCorrectedSensitivityDeviationProbabilityValueTag = 0x00240104
	// MinimumSensitivityValueTag is tag for Minimum Sensitivity Value
	MinimumSensitivityValueTag = 0x00240105
	// BlindSpotLocalizedTag is tag for Blind Spot Localized
	BlindSpotLocalizedTag = 0x00240106
	// BlindSpotXCoordinateTag is tag for Blind Spot X-Coordinate
	BlindSpotXCoordinateTag = 0x00240107
	// BlindSpotYCoordinateTag is tag for Blind Spot Y-Coordinate
	BlindSpotYCoordinateTag = 0x00240108
	// VisualAcuityMeasurementSequenceTag is tag for Visual Acuity Measurement Sequence
	VisualAcuityMeasurementSequenceTag = 0x00240110
	// RefractiveParametersUsedOnPatientSequenceTag is tag for Refractive Parameters Used on Patient Sequence
	RefractiveParametersUsedOnPatientSequenceTag = 0x00240112
	// MeasurementLateralityTag is tag for Measurement Laterality
	MeasurementLateralityTag = 0x00240113
	// OphthalmicPatientClinicalInformationLeftEyeSequenceTag is tag for Ophthalmic Patient Clinical Information Left Eye Sequence
	OphthalmicPatientClinicalInformationLeftEyeSequenceTag = 0x00240114
	// OphthalmicPatientClinicalInformationRightEyeSequenceTag is tag for Ophthalmic Patient Clinical Information Right Eye Sequence
	OphthalmicPatientClinicalInformationRightEyeSequenceTag = 0x00240115
	// FovealPointNormativeDataFlagTag is tag for Foveal Point Normative Data Flag
	FovealPointNormativeDataFlagTag = 0x00240117
	// FovealPointProbabilityValueTag is tag for Foveal Point Probability Value
	FovealPointProbabilityValueTag = 0x00240118
	// ScreeningBaselineMeasuredTag is tag for Screening Baseline Measured
	ScreeningBaselineMeasuredTag = 0x00240120
	// ScreeningBaselineMeasuredSequenceTag is tag for Screening Baseline Measured Sequence
	ScreeningBaselineMeasuredSequenceTag = 0x00240122
	// ScreeningBaselineTypeTag is tag for Screening Baseline Type
	ScreeningBaselineTypeTag = 0x00240124
	// ScreeningBaselineValueTag is tag for Screening Baseline Value
	ScreeningBaselineValueTag = 0x00240126
	// AlgorithmSourceTag is tag for Algorithm Source
	AlgorithmSourceTag = 0x00240202
	// DataSetNameTag is tag for Data Set Name
	DataSetNameTag = 0x00240306
	// DataSetVersionTag is tag for Data Set Version
	DataSetVersionTag = 0x00240307
	// DataSetSourceTag is tag for Data Set Source
	DataSetSourceTag = 0x00240308
	// DataSetDescriptionTag is tag for Data Set Description
	DataSetDescriptionTag = 0x00240309
	// VisualFieldTestReliabilityGlobalIndexSequenceTag is tag for Visual Field Test Reliability Global Index Sequence
	VisualFieldTestReliabilityGlobalIndexSequenceTag = 0x00240317
	// VisualFieldGlobalResultsIndexSequenceTag is tag for Visual Field Global Results Index Sequence
	VisualFieldGlobalResultsIndexSequenceTag = 0x00240320
	// DataObservationSequenceTag is tag for Data Observation Sequence
	DataObservationSequenceTag = 0x00240325
	// IndexNormalsFlagTag is tag for Index Normals Flag
	IndexNormalsFlagTag = 0x00240338
	// IndexProbabilityTag is tag for Index Probability
	IndexProbabilityTag = 0x00240341
	// IndexProbabilitySequenceTag is tag for Index Probability Sequence
	IndexProbabilitySequenceTag = 0x00240344
	// SamplesPerPixelTag is tag for Samples per Pixel
	SamplesPerPixelTag = 0x00280002
	// SamplesPerPixelUsedTag is tag for Samples per Pixel Used
	SamplesPerPixelUsedTag = 0x00280003
	// PhotometricInterpretationTag is tag for Photometric Interpretation
	PhotometricInterpretationTag = 0x00280004
	// ImageDimensionsTag is tag for Image Dimensions
	ImageDimensionsTag = 0x00280005
	// PlanarConfigurationTag is tag for Planar Configuration
	PlanarConfigurationTag = 0x00280006
	// NumberOfFramesTag is tag for Number of Frames
	NumberOfFramesTag = 0x00280008
	// FrameIncrementPointerTag is tag for Frame Increment Pointer
	FrameIncrementPointerTag = 0x00280009
	// FrameDimensionPointerTag is tag for Frame Dimension Pointer
	FrameDimensionPointerTag = 0x0028000A
	// RowsTag is tag for Rows
	RowsTag = 0x00280010
	// ColumnsTag is tag for Columns
	ColumnsTag = 0x00280011
	// PlanesTag is tag for Planes
	PlanesTag = 0x00280012
	// UltrasoundColorDataPresentTag is tag for Ultrasound Color Data Present
	UltrasoundColorDataPresentTag = 0x00280014
	// PixelSpacingTag is tag for Pixel Spacing
	PixelSpacingTag = 0x00280030
	// ZoomFactorTag is tag for Zoom Factor
	ZoomFactorTag = 0x00280031
	// ZoomCenterTag is tag for Zoom Center
	ZoomCenterTag = 0x00280032
	// PixelAspectRatioTag is tag for Pixel Aspect Ratio
	PixelAspectRatioTag = 0x00280034
	// ImageFormatTag is tag for Image Format
	ImageFormatTag = 0x00280040
	// ManipulatedImageTag is tag for Manipulated Image
	ManipulatedImageTag = 0x00280050
	// CorrectedImageTag is tag for Corrected Image
	CorrectedImageTag = 0x00280051
	// CompressionRecognitionCodeTag is tag for Compression Recognition Code
	CompressionRecognitionCodeTag = 0x0028005F
	// CompressionCodeTag is tag for Compression Code
	CompressionCodeTag = 0x00280060
	// CompressionOriginatorTag is tag for Compression Originator
	CompressionOriginatorTag = 0x00280061
	// CompressionLabelTag is tag for Compression Label
	CompressionLabelTag = 0x00280062
	// CompressionDescriptionTag is tag for Compression Description
	CompressionDescriptionTag = 0x00280063
	// CompressionSequenceTag is tag for Compression Sequence
	CompressionSequenceTag = 0x00280065
	// CompressionStepPointersTag is tag for Compression Step Pointers
	CompressionStepPointersTag = 0x00280066
	// RepeatIntervalTag is tag for Repeat Interval
	RepeatIntervalTag = 0x00280068
	// BitsGroupedTag is tag for Bits Grouped
	BitsGroupedTag = 0x00280069
	// PerimeterTableTag is tag for Perimeter Table
	PerimeterTableTag = 0x00280070
	// PerimeterValueTag is tag for Perimeter Value
	PerimeterValueTag = 0x00280071
	// PredictorRowsTag is tag for Predictor Rows
	PredictorRowsTag = 0x00280080
	// PredictorColumnsTag is tag for Predictor Columns
	PredictorColumnsTag = 0x00280081
	// PredictorConstantsTag is tag for Predictor Constants
	PredictorConstantsTag = 0x00280082
	// BlockedPixelsTag is tag for Blocked Pixels
	BlockedPixelsTag = 0x00280090
	// BlockRowsTag is tag for Block Rows
	BlockRowsTag = 0x00280091
	// BlockColumnsTag is tag for Block Columns
	BlockColumnsTag = 0x00280092
	// RowOverlapTag is tag for Row Overlap
	RowOverlapTag = 0x00280093
	// ColumnOverlapTag is tag for Column Overlap
	ColumnOverlapTag = 0x00280094
	// BitsAllocatedTag is tag for Bits Allocated
	BitsAllocatedTag = 0x00280100
	// BitsStoredTag is tag for Bits Stored
	BitsStoredTag = 0x00280101
	// HighBitTag is tag for High Bit
	HighBitTag = 0x00280102
	// PixelRepresentationTag is tag for Pixel Representation
	PixelRepresentationTag = 0x00280103
	// SmallestValidPixelValueTag is tag for Smallest Valid Pixel Value
	SmallestValidPixelValueTag = 0x00280104
	// LargestValidPixelValueTag is tag for Largest Valid Pixel Value
	LargestValidPixelValueTag = 0x00280105
	// SmallestImagePixelValueTag is tag for Smallest Image Pixel Value
	SmallestImagePixelValueTag = 0x00280106
	// LargestImagePixelValueTag is tag for Largest Image Pixel Value
	LargestImagePixelValueTag = 0x00280107
	// SmallestPixelValueInSeriesTag is tag for Smallest Pixel Value in Series
	SmallestPixelValueInSeriesTag = 0x00280108
	// LargestPixelValueInSeriesTag is tag for Largest Pixel Value in Series
	LargestPixelValueInSeriesTag = 0x00280109
	// SmallestImagePixelValueInPlaneTag is tag for Smallest Image Pixel Value in Plane
	SmallestImagePixelValueInPlaneTag = 0x00280110
	// LargestImagePixelValueInPlaneTag is tag for Largest Image Pixel Value in Plane
	LargestImagePixelValueInPlaneTag = 0x00280111
	// PixelPaddingValueTag is tag for Pixel Padding Value
	PixelPaddingValueTag = 0x00280120
	// PixelPaddingRangeLimitTag is tag for Pixel Padding Range Limit
	PixelPaddingRangeLimitTag = 0x00280121
	// FloatPixelPaddingValueTag is tag for Float Pixel Padding Value
	FloatPixelPaddingValueTag = 0x00280122
	// DoubleFloatPixelPaddingValueTag is tag for Double Float Pixel Padding Value
	DoubleFloatPixelPaddingValueTag = 0x00280123
	// FloatPixelPaddingRangeLimitTag is tag for Float Pixel Padding Range Limit
	FloatPixelPaddingRangeLimitTag = 0x00280124
	// DoubleFloatPixelPaddingRangeLimitTag is tag for Double Float Pixel Padding Range Limit
	DoubleFloatPixelPaddingRangeLimitTag = 0x00280125
	// ImageLocationTag is tag for Image Location
	ImageLocationTag = 0x00280200
	// QualityControlImageTag is tag for Quality Control Image
	QualityControlImageTag = 0x00280300
	// BurnedInAnnotationTag is tag for Burned In Annotation
	BurnedInAnnotationTag = 0x00280301
	// RecognizableVisualFeaturesTag is tag for Recognizable Visual Features
	RecognizableVisualFeaturesTag = 0x00280302
	// LongitudinalTemporalInformationModifiedTag is tag for Longitudinal Temporal Information Modified
	LongitudinalTemporalInformationModifiedTag = 0x00280303
	// ReferencedColorPaletteInstanceUIDTag is tag for Referenced Color Palette Instance UID
	ReferencedColorPaletteInstanceUIDTag = 0x00280304
	// RowsForNthOrderCoefficientsTag is tag for Rows For Nth Order Coefficients
	RowsForNthOrderCoefficientsTag = 0x00280400
	// ColumnsForNthOrderCoefficientsTag is tag for Columns For Nth Order Coefficients
	ColumnsForNthOrderCoefficientsTag = 0x00280401
	// CoefficientCodingTag is tag for Coefficient Coding
	CoefficientCodingTag = 0x00280402
	// CoefficientCodingPointersTag is tag for Coefficient Coding Pointers
	CoefficientCodingPointersTag = 0x00280403
	// DetailsOfCoefficientsTag is tag for Details of Coefficients
	DetailsOfCoefficientsTag = 0x00280404
	// DCTLabelTag is tag for DCT Label
	DCTLabelTag = 0x00280700
	// DataBlockDescriptionTag is tag for Data Block Description
	DataBlockDescriptionTag = 0x00280701
	// DataBlockTag is tag for Data Block
	DataBlockTag = 0x00280702
	// NormalizationFactorFormatTag is tag for Normalization Factor Format
	NormalizationFactorFormatTag = 0x00280710
	// ZonalMapNumberFormatTag is tag for Zonal Map Number Format
	ZonalMapNumberFormatTag = 0x00280720
	// ZonalMapLocationTag is tag for Zonal Map Location
	ZonalMapLocationTag = 0x00280721
	// ZonalMapFormatTag is tag for Zonal Map Format
	ZonalMapFormatTag = 0x00280722
	// AdaptiveMapFormatTag is tag for Adaptive Map Format
	AdaptiveMapFormatTag = 0x00280730
	// CodeNumberFormatTag is tag for Code Number Format
	CodeNumberFormatTag = 0x00280740
	// CodeLabelTag is tag for Code Label
	CodeLabelTag = 0x00280800
	// NumberOfTablesTag is tag for Number of Tables
	NumberOfTablesTag = 0x00280802
	// CodeTableLocationTag is tag for Code Table Location
	CodeTableLocationTag = 0x00280803
	// BitsForCodeWordTag is tag for Bits For Code Word
	BitsForCodeWordTag = 0x00280804
	// ImageDataLocationTag is tag for Image Data Location
	ImageDataLocationTag = 0x00280808
	// PixelSpacingCalibrationTypeTag is tag for Pixel Spacing Calibration Type
	PixelSpacingCalibrationTypeTag = 0x00280A02
	// PixelSpacingCalibrationDescriptionTag is tag for Pixel Spacing Calibration Description
	PixelSpacingCalibrationDescriptionTag = 0x00280A04
	// PixelIntensityRelationshipTag is tag for Pixel Intensity Relationship
	PixelIntensityRelationshipTag = 0x00281040
	// PixelIntensityRelationshipSignTag is tag for Pixel Intensity Relationship Sign
	PixelIntensityRelationshipSignTag = 0x00281041
	// WindowCenterTag is tag for Window Center
	WindowCenterTag = 0x00281050
	// WindowWidthTag is tag for Window Width
	WindowWidthTag = 0x00281051
	// RescaleInterceptTag is tag for Rescale Intercept
	RescaleInterceptTag = 0x00281052
	// RescaleSlopeTag is tag for Rescale Slope
	RescaleSlopeTag = 0x00281053
	// RescaleTypeTag is tag for Rescale Type
	RescaleTypeTag = 0x00281054
	// WindowCenterWidthExplanationTag is tag for Window Center & Width Explanation
	WindowCenterWidthExplanationTag = 0x00281055
	// VOILUTFunctionTag is tag for VOI LUT Function
	VOILUTFunctionTag = 0x00281056
	// GrayScaleTag is tag for Gray Scale
	GrayScaleTag = 0x00281080
	// RecommendedViewingModeTag is tag for Recommended Viewing Mode
	RecommendedViewingModeTag = 0x00281090
	// GrayLookupTableDescriptorTag is tag for Gray Lookup Table Descriptor
	GrayLookupTableDescriptorTag = 0x00281100
	// RedPaletteColorLookupTableDescriptorTag is tag for Red Palette Color Lookup Table Descriptor
	RedPaletteColorLookupTableDescriptorTag = 0x00281101
	// GreenPaletteColorLookupTableDescriptorTag is tag for Green Palette Color Lookup Table Descriptor
	GreenPaletteColorLookupTableDescriptorTag = 0x00281102
	// BluePaletteColorLookupTableDescriptorTag is tag for Blue Palette Color Lookup Table Descriptor
	BluePaletteColorLookupTableDescriptorTag = 0x00281103
	// AlphaPaletteColorLookupTableDescriptorTag is tag for Alpha Palette Color Lookup Table Descriptor
	AlphaPaletteColorLookupTableDescriptorTag = 0x00281104
	// LargeRedPaletteColorLookupTableDescriptorTag is tag for Large Red Palette Color Lookup Table Descriptor
	LargeRedPaletteColorLookupTableDescriptorTag = 0x00281111
	// LargeGreenPaletteColorLookupTableDescriptorTag is tag for Large Green Palette Color Lookup Table Descriptor
	LargeGreenPaletteColorLookupTableDescriptorTag = 0x00281112
	// LargeBluePaletteColorLookupTableDescriptorTag is tag for Large Blue Palette Color Lookup Table Descriptor
	LargeBluePaletteColorLookupTableDescriptorTag = 0x00281113
	// PaletteColorLookupTableUIDTag is tag for Palette Color Lookup Table UID
	PaletteColorLookupTableUIDTag = 0x00281199
	// GrayLookupTableDataTag is tag for Gray Lookup Table Data
	GrayLookupTableDataTag = 0x00281200
	// RedPaletteColorLookupTableDataTag is tag for Red Palette Color Lookup Table Data
	RedPaletteColorLookupTableDataTag = 0x00281201
	// GreenPaletteColorLookupTableDataTag is tag for Green Palette Color Lookup Table Data
	GreenPaletteColorLookupTableDataTag = 0x00281202
	// BluePaletteColorLookupTableDataTag is tag for Blue Palette Color Lookup Table Data
	BluePaletteColorLookupTableDataTag = 0x00281203
	// AlphaPaletteColorLookupTableDataTag is tag for Alpha Palette Color Lookup Table Data
	AlphaPaletteColorLookupTableDataTag = 0x00281204
	// LargeRedPaletteColorLookupTableDataTag is tag for Large Red Palette Color Lookup Table Data
	LargeRedPaletteColorLookupTableDataTag = 0x00281211
	// LargeGreenPaletteColorLookupTableDataTag is tag for Large Green Palette Color Lookup Table Data
	LargeGreenPaletteColorLookupTableDataTag = 0x00281212
	// LargeBluePaletteColorLookupTableDataTag is tag for Large Blue Palette Color Lookup Table Data
	LargeBluePaletteColorLookupTableDataTag = 0x00281213
	// LargePaletteColorLookupTableUIDTag is tag for Large Palette Color Lookup Table UID
	LargePaletteColorLookupTableUIDTag = 0x00281214
	// SegmentedRedPaletteColorLookupTableDataTag is tag for Segmented Red Palette Color Lookup Table Data
	SegmentedRedPaletteColorLookupTableDataTag = 0x00281221
	// SegmentedGreenPaletteColorLookupTableDataTag is tag for Segmented Green Palette Color Lookup Table Data
	SegmentedGreenPaletteColorLookupTableDataTag = 0x00281222
	// SegmentedBluePaletteColorLookupTableDataTag is tag for Segmented Blue Palette Color Lookup Table Data
	SegmentedBluePaletteColorLookupTableDataTag = 0x00281223
	// SegmentedAlphaPaletteColorLookupTableDataTag is tag for Segmented Alpha Palette Color Lookup Table Data
	SegmentedAlphaPaletteColorLookupTableDataTag = 0x00281224
	// StoredValueColorRangeSequenceTag is tag for Stored Value Color Range Sequence
	StoredValueColorRangeSequenceTag = 0x00281230
	// MinimumStoredValueMappedTag is tag for Minimum Stored Value Mapped
	MinimumStoredValueMappedTag = 0x00281231
	// MaximumStoredValueMappedTag is tag for Maximum Stored Value Mapped
	MaximumStoredValueMappedTag = 0x00281232
	// BreastImplantPresentTag is tag for Breast Implant Present
	BreastImplantPresentTag = 0x00281300
	// PartialViewTag is tag for Partial View
	PartialViewTag = 0x00281350
	// PartialViewDescriptionTag is tag for Partial View Description
	PartialViewDescriptionTag = 0x00281351
	// PartialViewCodeSequenceTag is tag for Partial View Code Sequence
	PartialViewCodeSequenceTag = 0x00281352
	// SpatialLocationsPreservedTag is tag for Spatial Locations Preserved
	SpatialLocationsPreservedTag = 0x0028135A
	// DataFrameAssignmentSequenceTag is tag for Data Frame Assignment Sequence
	DataFrameAssignmentSequenceTag = 0x00281401
	// DataPathAssignmentTag is tag for Data Path Assignment
	DataPathAssignmentTag = 0x00281402
	// BitsMappedToColorLookupTableTag is tag for Bits Mapped to Color Lookup Table
	BitsMappedToColorLookupTableTag = 0x00281403
	// BlendingLUT1SequenceTag is tag for Blending LUT 1 Sequence
	BlendingLUT1SequenceTag = 0x00281404
	// BlendingLUT1TransferFunctionTag is tag for Blending LUT 1 Transfer Function
	BlendingLUT1TransferFunctionTag = 0x00281405
	// BlendingWeightConstantTag is tag for Blending Weight Constant
	BlendingWeightConstantTag = 0x00281406
	// BlendingLookupTableDescriptorTag is tag for Blending Lookup Table Descriptor
	BlendingLookupTableDescriptorTag = 0x00281407
	// BlendingLookupTableDataTag is tag for Blending Lookup Table Data
	BlendingLookupTableDataTag = 0x00281408
	// EnhancedPaletteColorLookupTableSequenceTag is tag for Enhanced Palette Color Lookup Table Sequence
	EnhancedPaletteColorLookupTableSequenceTag = 0x0028140B
	// BlendingLUT2SequenceTag is tag for Blending LUT 2 Sequence
	BlendingLUT2SequenceTag = 0x0028140C
	// BlendingLUT2TransferFunctionTag is tag for Blending LUT 2 Transfer Function
	BlendingLUT2TransferFunctionTag = 0x0028140D
	// DataPathIDTag is tag for Data Path ID
	DataPathIDTag = 0x0028140E
	// RGBLUTTransferFunctionTag is tag for RGB LUT Transfer Function
	RGBLUTTransferFunctionTag = 0x0028140F
	// AlphaLUTTransferFunctionTag is tag for Alpha LUT Transfer Function
	AlphaLUTTransferFunctionTag = 0x00281410
	// ICCProfileTag is tag for ICC Profile
	ICCProfileTag = 0x00282000
	// ColorSpaceTag is tag for Color Space
	ColorSpaceTag = 0x00282002
	// LossyImageCompressionTag is tag for Lossy Image Compression
	LossyImageCompressionTag = 0x00282110
	// LossyImageCompressionRatioTag is tag for Lossy Image Compression Ratio
	LossyImageCompressionRatioTag = 0x00282112
	// LossyImageCompressionMethodTag is tag for Lossy Image Compression Method
	LossyImageCompressionMethodTag = 0x00282114
	// ModalityLUTSequenceTag is tag for Modality LUT Sequence
	ModalityLUTSequenceTag = 0x00283000
	// LUTDescriptorTag is tag for LUT Descriptor
	LUTDescriptorTag = 0x00283002
	// LUTExplanationTag is tag for LUT Explanation
	LUTExplanationTag = 0x00283003
	// ModalityLUTTypeTag is tag for Modality LUT Type
	ModalityLUTTypeTag = 0x00283004
	// LUTDataTag is tag for LUT Data
	LUTDataTag = 0x00283006
	// VOILUTSequenceTag is tag for VOI LUT Sequence
	VOILUTSequenceTag = 0x00283010
	// SoftcopyVOILUTSequenceTag is tag for Softcopy VOI LUT Sequence
	SoftcopyVOILUTSequenceTag = 0x00283110
	// ImagePresentationCommentsTag is tag for Image Presentation Comments
	ImagePresentationCommentsTag = 0x00284000
	// BiPlaneAcquisitionSequenceTag is tag for Bi-Plane Acquisition Sequence
	BiPlaneAcquisitionSequenceTag = 0x00285000
	// RepresentativeFrameNumberTag is tag for Representative Frame Number
	RepresentativeFrameNumberTag = 0x00286010
	// FrameNumbersOfInterestTag is tag for Frame Numbers of Interest (FOI)
	FrameNumbersOfInterestTag = 0x00286020
	// FrameOfInterestDescriptionTag is tag for Frame of Interest Description
	FrameOfInterestDescriptionTag = 0x00286022
	// FrameOfInterestTypeTag is tag for Frame of Interest Type
	FrameOfInterestTypeTag = 0x00286023
	// MaskPointersTag is tag for Mask Pointer(s)
	MaskPointersTag = 0x00286030
	// RWavePointerTag is tag for R Wave Pointer
	RWavePointerTag = 0x00286040
	// MaskSubtractionSequenceTag is tag for Mask Subtraction Sequence
	MaskSubtractionSequenceTag = 0x00286100
	// MaskOperationTag is tag for Mask Operation
	MaskOperationTag = 0x00286101
	// ApplicableFrameRangeTag is tag for Applicable Frame Range
	ApplicableFrameRangeTag = 0x00286102
	// MaskFrameNumbersTag is tag for Mask Frame Numbers
	MaskFrameNumbersTag = 0x00286110
	// ContrastFrameAveragingTag is tag for Contrast Frame Averaging
	ContrastFrameAveragingTag = 0x00286112
	// MaskSubPixelShiftTag is tag for Mask Sub-pixel Shift
	MaskSubPixelShiftTag = 0x00286114
	// TIDOffsetTag is tag for TID Offset
	TIDOffsetTag = 0x00286120
	// MaskOperationExplanationTag is tag for Mask Operation Explanation
	MaskOperationExplanationTag = 0x00286190
	// EquipmentAdministratorSequenceTag is tag for Equipment Administrator Sequence
	EquipmentAdministratorSequenceTag = 0x00287000
	// NumberOfDisplaySubsystemsTag is tag for Number of Display Subsystems
	NumberOfDisplaySubsystemsTag = 0x00287001
	// CurrentConfigurationIDTag is tag for Current Configuration ID
	CurrentConfigurationIDTag = 0x00287002
	// DisplaySubsystemIDTag is tag for Display Subsystem ID
	DisplaySubsystemIDTag = 0x00287003
	// DisplaySubsystemNameTag is tag for Display Subsystem Name
	DisplaySubsystemNameTag = 0x00287004
	// DisplaySubsystemDescriptionTag is tag for Display Subsystem Description
	DisplaySubsystemDescriptionTag = 0x00287005
	// SystemStatusTag is tag for System Status
	SystemStatusTag = 0x00287006
	// SystemStatusCommentTag is tag for System Status Comment
	SystemStatusCommentTag = 0x00287007
	// TargetLuminanceCharacteristicsSequenceTag is tag for Target Luminance Characteristics Sequence
	TargetLuminanceCharacteristicsSequenceTag = 0x00287008
	// LuminanceCharacteristicsIDTag is tag for Luminance Characteristics ID
	LuminanceCharacteristicsIDTag = 0x00287009
	// DisplaySubsystemConfigurationSequenceTag is tag for Display Subsystem Configuration Sequence
	DisplaySubsystemConfigurationSequenceTag = 0x0028700A
	// ConfigurationIDTag is tag for Configuration ID
	ConfigurationIDTag = 0x0028700B
	// ConfigurationNameTag is tag for Configuration Name
	ConfigurationNameTag = 0x0028700C
	// ConfigurationDescriptionTag is tag for Configuration Description
	ConfigurationDescriptionTag = 0x0028700D
	// ReferencedTargetLuminanceCharacteristicsIDTag is tag for Referenced Target Luminance Characteristics ID
	ReferencedTargetLuminanceCharacteristicsIDTag = 0x0028700E
	// QAResultsSequenceTag is tag for QA Results Sequence
	QAResultsSequenceTag = 0x0028700F
	// DisplaySubsystemQAResultsSequenceTag is tag for Display Subsystem QA Results Sequence
	DisplaySubsystemQAResultsSequenceTag = 0x00287010
	// ConfigurationQAResultsSequenceTag is tag for Configuration QA Results Sequence
	ConfigurationQAResultsSequenceTag = 0x00287011
	// MeasurementEquipmentSequenceTag is tag for Measurement Equipment Sequence
	MeasurementEquipmentSequenceTag = 0x00287012
	// MeasurementFunctionsTag is tag for Measurement Functions
	MeasurementFunctionsTag = 0x00287013
	// MeasurementEquipmentTypeTag is tag for Measurement Equipment Type
	MeasurementEquipmentTypeTag = 0x00287014
	// VisualEvaluationResultSequenceTag is tag for Visual Evaluation Result Sequence
	VisualEvaluationResultSequenceTag = 0x00287015
	// DisplayCalibrationResultSequenceTag is tag for Display Calibration Result Sequence
	DisplayCalibrationResultSequenceTag = 0x00287016
	// DDLValueTag is tag for DDL Value
	DDLValueTag = 0x00287017
	// CIExyWhitePointTag is tag for CIExy White Point
	CIExyWhitePointTag = 0x00287018
	// DisplayFunctionTypeTag is tag for Display Function Type
	DisplayFunctionTypeTag = 0x00287019
	// GammaValueTag is tag for Gamma Value
	GammaValueTag = 0x0028701A
	// NumberOfLuminancePointsTag is tag for Number of Luminance Points
	NumberOfLuminancePointsTag = 0x0028701B
	// LuminanceResponseSequenceTag is tag for Luminance Response Sequence
	LuminanceResponseSequenceTag = 0x0028701C
	// TargetMinimumLuminanceTag is tag for Target Minimum Luminance
	TargetMinimumLuminanceTag = 0x0028701D
	// TargetMaximumLuminanceTag is tag for Target Maximum Luminance
	TargetMaximumLuminanceTag = 0x0028701E
	// LuminanceValueTag is tag for Luminance Value
	LuminanceValueTag = 0x0028701F
	// LuminanceResponseDescriptionTag is tag for Luminance Response Description
	LuminanceResponseDescriptionTag = 0x00287020
	// WhitePointFlagTag is tag for White Point Flag
	WhitePointFlagTag = 0x00287021
	// DisplayDeviceTypeCodeSequenceTag is tag for Display Device Type Code Sequence
	DisplayDeviceTypeCodeSequenceTag = 0x00287022
	// DisplaySubsystemSequenceTag is tag for Display Subsystem Sequence
	DisplaySubsystemSequenceTag = 0x00287023
	// LuminanceResultSequenceTag is tag for Luminance Result Sequence
	LuminanceResultSequenceTag = 0x00287024
	// AmbientLightValueSourceTag is tag for Ambient Light Value Source
	AmbientLightValueSourceTag = 0x00287025
	// MeasuredCharacteristicsTag is tag for Measured Characteristics
	MeasuredCharacteristicsTag = 0x00287026
	// LuminanceUniformityResultSequenceTag is tag for Luminance Uniformity Result Sequence
	LuminanceUniformityResultSequenceTag = 0x00287027
	// VisualEvaluationTestSequenceTag is tag for Visual Evaluation Test Sequence
	VisualEvaluationTestSequenceTag = 0x00287028
	// TestResultTag is tag for Test Result
	TestResultTag = 0x00287029
	// TestResultCommentTag is tag for Test Result Comment
	TestResultCommentTag = 0x0028702A
	// TestImageValidationTag is tag for Test Image Validation
	TestImageValidationTag = 0x0028702B
	// TestPatternCodeSequenceTag is tag for Test Pattern Code Sequence
	TestPatternCodeSequenceTag = 0x0028702C
	// MeasurementPatternCodeSequenceTag is tag for Measurement Pattern Code Sequence
	MeasurementPatternCodeSequenceTag = 0x0028702D
	// VisualEvaluationMethodCodeSequenceTag is tag for Visual Evaluation Method Code Sequence
	VisualEvaluationMethodCodeSequenceTag = 0x0028702E
	// PixelDataProviderURLTag is tag for Pixel Data Provider URL
	PixelDataProviderURLTag = 0x00287FE0
	// DataPointRowsTag is tag for Data Point Rows
	DataPointRowsTag = 0x00289001
	// DataPointColumnsTag is tag for Data Point Columns
	DataPointColumnsTag = 0x00289002
	// SignalDomainColumnsTag is tag for Signal Domain Columns
	SignalDomainColumnsTag = 0x00289003
	// LargestMonochromePixelValueTag is tag for Largest Monochrome Pixel Value
	LargestMonochromePixelValueTag = 0x00289099
	// DataRepresentationTag is tag for Data Representation
	DataRepresentationTag = 0x00289108
	// PixelMeasuresSequenceTag is tag for Pixel Measures Sequence
	PixelMeasuresSequenceTag = 0x00289110
	// FrameVOILUTSequenceTag is tag for Frame VOI LUT Sequence
	FrameVOILUTSequenceTag = 0x00289132
	// PixelValueTransformationSequenceTag is tag for Pixel Value Transformation Sequence
	PixelValueTransformationSequenceTag = 0x00289145
	// SignalDomainRowsTag is tag for Signal Domain Rows
	SignalDomainRowsTag = 0x00289235
	// DisplayFilterPercentageTag is tag for Display Filter Percentage
	DisplayFilterPercentageTag = 0x00289411
	// FramePixelShiftSequenceTag is tag for Frame Pixel Shift Sequence
	FramePixelShiftSequenceTag = 0x00289415
	// SubtractionItemIDTag is tag for Subtraction Item ID
	SubtractionItemIDTag = 0x00289416
	// PixelIntensityRelationshipLUTSequenceTag is tag for Pixel Intensity Relationship LUT Sequence
	PixelIntensityRelationshipLUTSequenceTag = 0x00289422
	// FramePixelDataPropertiesSequenceTag is tag for Frame Pixel Data Properties Sequence
	FramePixelDataPropertiesSequenceTag = 0x00289443
	// GeometricalPropertiesTag is tag for Geometrical Properties
	GeometricalPropertiesTag = 0x00289444
	// GeometricMaximumDistortionTag is tag for Geometric Maximum Distortion
	GeometricMaximumDistortionTag = 0x00289445
	// ImageProcessingAppliedTag is tag for Image Processing Applied
	ImageProcessingAppliedTag = 0x00289446
	// MaskSelectionModeTag is tag for Mask Selection Mode
	MaskSelectionModeTag = 0x00289454
	// LUTFunctionTag is tag for LUT Function
	LUTFunctionTag = 0x00289474
	// MaskVisibilityPercentageTag is tag for Mask Visibility Percentage
	MaskVisibilityPercentageTag = 0x00289478
	// PixelShiftSequenceTag is tag for Pixel Shift Sequence
	PixelShiftSequenceTag = 0x00289501
	// RegionPixelShiftSequenceTag is tag for Region Pixel Shift Sequence
	RegionPixelShiftSequenceTag = 0x00289502
	// VerticesOfTheRegionTag is tag for Vertices of the Region
	VerticesOfTheRegionTag = 0x00289503
	// MultiFramePresentationSequenceTag is tag for Multi-frame Presentation Sequence
	MultiFramePresentationSequenceTag = 0x00289505
	// PixelShiftFrameRangeTag is tag for Pixel Shift Frame Range
	PixelShiftFrameRangeTag = 0x00289506
	// LUTFrameRangeTag is tag for LUT Frame Range
	LUTFrameRangeTag = 0x00289507
	// ImageToEquipmentMappingMatrixTag is tag for Image to Equipment Mapping Matrix
	ImageToEquipmentMappingMatrixTag = 0x00289520
	// EquipmentCoordinateSystemIdentificationTag is tag for Equipment Coordinate System Identification
	EquipmentCoordinateSystemIdentificationTag = 0x00289537
	// StudyStatusIDTag is tag for Study Status ID
	StudyStatusIDTag = 0x0032000A
	// StudyPriorityIDTag is tag for Study Priority ID
	StudyPriorityIDTag = 0x0032000C
	// StudyIDIssuerTag is tag for Study ID Issuer
	StudyIDIssuerTag = 0x00320012
	// StudyVerifiedDateTag is tag for Study Verified Date
	StudyVerifiedDateTag = 0x00320032
	// StudyVerifiedTimeTag is tag for Study Verified Time
	StudyVerifiedTimeTag = 0x00320033
	// StudyReadDateTag is tag for Study Read Date
	StudyReadDateTag = 0x00320034
	// StudyReadTimeTag is tag for Study Read Time
	StudyReadTimeTag = 0x00320035
	// ScheduledStudyStartDateTag is tag for Scheduled Study Start Date
	ScheduledStudyStartDateTag = 0x00321000
	// ScheduledStudyStartTimeTag is tag for Scheduled Study Start Time
	ScheduledStudyStartTimeTag = 0x00321001
	// ScheduledStudyStopDateTag is tag for Scheduled Study Stop Date
	ScheduledStudyStopDateTag = 0x00321010
	// ScheduledStudyStopTimeTag is tag for Scheduled Study Stop Time
	ScheduledStudyStopTimeTag = 0x00321011
	// ScheduledStudyLocationTag is tag for Scheduled Study Location
	ScheduledStudyLocationTag = 0x00321020
	// ScheduledStudyLocationAETitleTag is tag for Scheduled Study Location AE Title
	ScheduledStudyLocationAETitleTag = 0x00321021
	// ReasonForStudyTag is tag for Reason for Study
	ReasonForStudyTag = 0x00321030
	// RequestingPhysicianIdentificationSequenceTag is tag for Requesting Physician Identification Sequence
	RequestingPhysicianIdentificationSequenceTag = 0x00321031
	// RequestingPhysicianTag is tag for Requesting Physician
	RequestingPhysicianTag = 0x00321032
	// RequestingServiceTag is tag for Requesting Service
	RequestingServiceTag = 0x00321033
	// RequestingServiceCodeSequenceTag is tag for Requesting Service Code Sequence
	RequestingServiceCodeSequenceTag = 0x00321034
	// StudyArrivalDateTag is tag for Study Arrival Date
	StudyArrivalDateTag = 0x00321040
	// StudyArrivalTimeTag is tag for Study Arrival Time
	StudyArrivalTimeTag = 0x00321041
	// StudyCompletionDateTag is tag for Study Completion Date
	StudyCompletionDateTag = 0x00321050
	// StudyCompletionTimeTag is tag for Study Completion Time
	StudyCompletionTimeTag = 0x00321051
	// StudyComponentStatusIDTag is tag for Study Component Status ID
	StudyComponentStatusIDTag = 0x00321055
	// RequestedProcedureDescriptionTag is tag for Requested Procedure Description
	RequestedProcedureDescriptionTag = 0x00321060
	// RequestedProcedureCodeSequenceTag is tag for Requested Procedure Code Sequence
	RequestedProcedureCodeSequenceTag = 0x00321064
	// ReasonForVisitTag is tag for Reason for Visit
	ReasonForVisitTag = 0x00321066
	// ReasonForVisitCodeSequenceTag is tag for Reason for Visit Code Sequence
	ReasonForVisitCodeSequenceTag = 0x00321067
	// RequestedContrastAgentTag is tag for Requested Contrast Agent
	RequestedContrastAgentTag = 0x00321070
	// StudyCommentsTag is tag for Study Comments
	StudyCommentsTag = 0x00324000
	// FlowIdentifierSequenceTag is tag for Flow Identifier Sequence
	FlowIdentifierSequenceTag = 0x00340001
	// FlowIdentifierTag is tag for Flow Identifier
	FlowIdentifierTag = 0x00340002
	// FlowTransferSyntaxUIDTag is tag for Flow Transfer Syntax UID
	FlowTransferSyntaxUIDTag = 0x00340003
	// FlowRTPSamplingRateTag is tag for Flow RTP Sampling Rate
	FlowRTPSamplingRateTag = 0x00340004
	// SourceIdentifierTag is tag for Source Identifier
	SourceIdentifierTag = 0x00340005
	// FrameOriginTimestampTag is tag for Frame Origin Timestamp
	FrameOriginTimestampTag = 0x00340007
	// IncludesImagingSubjectTag is tag for Includes Imaging Subject
	IncludesImagingSubjectTag = 0x00340008
	// FrameUsefulnessGroupSequenceTag is tag for Frame Usefulness Group Sequence
	FrameUsefulnessGroupSequenceTag = 0x00340009
	// RealTimeBulkDataFlowSequenceTag is tag for Real-Time Bulk Data Flow Sequence
	RealTimeBulkDataFlowSequenceTag = 0x0034000A
	// CameraPositionGroupSequenceTag is tag for Camera Position Group Sequence
	CameraPositionGroupSequenceTag = 0x0034000B
	// IncludesInformationTag is tag for Includes Information
	IncludesInformationTag = 0x0034000C
	// TimeOfFrameGroupSequenceTag is tag for Time of Frame Group Sequence
	TimeOfFrameGroupSequenceTag = 0x0034000D
	// ReferencedPatientAliasSequenceTag is tag for Referenced Patient Alias Sequence
	ReferencedPatientAliasSequenceTag = 0x00380004
	// VisitStatusIDTag is tag for Visit Status ID
	VisitStatusIDTag = 0x00380008
	// AdmissionIDTag is tag for Admission ID
	AdmissionIDTag = 0x00380010
	// IssuerOfAdmissionIDTag is tag for Issuer of Admission ID
	IssuerOfAdmissionIDTag = 0x00380011
	// IssuerOfAdmissionIDSequenceTag is tag for Issuer of Admission ID Sequence
	IssuerOfAdmissionIDSequenceTag = 0x00380014
	// RouteOfAdmissionsTag is tag for Route of Admissions
	RouteOfAdmissionsTag = 0x00380016
	// ScheduledAdmissionDateTag is tag for Scheduled Admission Date
	ScheduledAdmissionDateTag = 0x0038001A
	// ScheduledAdmissionTimeTag is tag for Scheduled Admission Time
	ScheduledAdmissionTimeTag = 0x0038001B
	// ScheduledDischargeDateTag is tag for Scheduled Discharge Date
	ScheduledDischargeDateTag = 0x0038001C
	// ScheduledDischargeTimeTag is tag for Scheduled Discharge Time
	ScheduledDischargeTimeTag = 0x0038001D
	// ScheduledPatientInstitutionResidenceTag is tag for Scheduled Patient Institution Residence
	ScheduledPatientInstitutionResidenceTag = 0x0038001E
	// AdmittingDateTag is tag for Admitting Date
	AdmittingDateTag = 0x00380020
	// AdmittingTimeTag is tag for Admitting Time
	AdmittingTimeTag = 0x00380021
	// DischargeDateTag is tag for Discharge Date
	DischargeDateTag = 0x00380030
	// DischargeTimeTag is tag for Discharge Time
	DischargeTimeTag = 0x00380032
	// DischargeDiagnosisDescriptionTag is tag for Discharge Diagnosis Description
	DischargeDiagnosisDescriptionTag = 0x00380040
	// DischargeDiagnosisCodeSequenceTag is tag for Discharge Diagnosis Code Sequence
	DischargeDiagnosisCodeSequenceTag = 0x00380044
	// SpecialNeedsTag is tag for Special Needs
	SpecialNeedsTag = 0x00380050
	// ServiceEpisodeIDTag is tag for Service Episode ID
	ServiceEpisodeIDTag = 0x00380060
	// IssuerOfServiceEpisodeIDTag is tag for Issuer of Service Episode ID
	IssuerOfServiceEpisodeIDTag = 0x00380061
	// ServiceEpisodeDescriptionTag is tag for Service Episode Description
	ServiceEpisodeDescriptionTag = 0x00380062
	// IssuerOfServiceEpisodeIDSequenceTag is tag for Issuer of Service Episode ID Sequence
	IssuerOfServiceEpisodeIDSequenceTag = 0x00380064
	// PertinentDocumentsSequenceTag is tag for Pertinent Documents Sequence
	PertinentDocumentsSequenceTag = 0x00380100
	// PertinentResourcesSequenceTag is tag for Pertinent Resources Sequence
	PertinentResourcesSequenceTag = 0x00380101
	// ResourceDescriptionTag is tag for Resource Description
	ResourceDescriptionTag = 0x00380102
	// CurrentPatientLocationTag is tag for Current Patient Location
	CurrentPatientLocationTag = 0x00380300
	// PatientInstitutionResidenceTag is tag for Patient's Institution Residence
	PatientInstitutionResidenceTag = 0x00380400
	// PatientStateTag is tag for Patient State
	PatientStateTag = 0x00380500
	// PatientClinicalTrialParticipationSequenceTag is tag for Patient Clinical Trial Participation Sequence
	PatientClinicalTrialParticipationSequenceTag = 0x00380502
	// VisitCommentsTag is tag for Visit Comments
	VisitCommentsTag = 0x00384000
	// WaveformOriginalityTag is tag for Waveform Originality
	WaveformOriginalityTag = 0x003A0004
	// NumberOfWaveformChannelsTag is tag for Number of Waveform Channels
	NumberOfWaveformChannelsTag = 0x003A0005
	// NumberOfWaveformSamplesTag is tag for Number of Waveform Samples
	NumberOfWaveformSamplesTag = 0x003A0010
	// SamplingFrequencyTag is tag for Sampling Frequency
	SamplingFrequencyTag = 0x003A001A
	// MultiplexGroupLabelTag is tag for Multiplex Group Label
	MultiplexGroupLabelTag = 0x003A0020
	// ChannelDefinitionSequenceTag is tag for Channel Definition Sequence
	ChannelDefinitionSequenceTag = 0x003A0200
	// WaveformChannelNumberTag is tag for Waveform Channel Number
	WaveformChannelNumberTag = 0x003A0202
	// ChannelLabelTag is tag for Channel Label
	ChannelLabelTag = 0x003A0203
	// ChannelStatusTag is tag for Channel Status
	ChannelStatusTag = 0x003A0205
	// ChannelSourceSequenceTag is tag for Channel Source Sequence
	ChannelSourceSequenceTag = 0x003A0208
	// ChannelSourceModifiersSequenceTag is tag for Channel Source Modifiers Sequence
	ChannelSourceModifiersSequenceTag = 0x003A0209
	// SourceWaveformSequenceTag is tag for Source Waveform Sequence
	SourceWaveformSequenceTag = 0x003A020A
	// ChannelDerivationDescriptionTag is tag for Channel Derivation Description
	ChannelDerivationDescriptionTag = 0x003A020C
	// ChannelSensitivityTag is tag for Channel Sensitivity
	ChannelSensitivityTag = 0x003A0210
	// ChannelSensitivityUnitsSequenceTag is tag for Channel Sensitivity Units Sequence
	ChannelSensitivityUnitsSequenceTag = 0x003A0211
	// ChannelSensitivityCorrectionFactorTag is tag for Channel Sensitivity Correction Factor
	ChannelSensitivityCorrectionFactorTag = 0x003A0212
	// ChannelBaselineTag is tag for Channel Baseline
	ChannelBaselineTag = 0x003A0213
	// ChannelTimeSkewTag is tag for Channel Time Skew
	ChannelTimeSkewTag = 0x003A0214
	// ChannelSampleSkewTag is tag for Channel Sample Skew
	ChannelSampleSkewTag = 0x003A0215
	// ChannelOffsetTag is tag for Channel Offset
	ChannelOffsetTag = 0x003A0218
	// WaveformBitsStoredTag is tag for Waveform Bits Stored
	WaveformBitsStoredTag = 0x003A021A
	// FilterLowFrequencyTag is tag for Filter Low Frequency
	FilterLowFrequencyTag = 0x003A0220
	// FilterHighFrequencyTag is tag for Filter High Frequency
	FilterHighFrequencyTag = 0x003A0221
	// NotchFilterFrequencyTag is tag for Notch Filter Frequency
	NotchFilterFrequencyTag = 0x003A0222
	// NotchFilterBandwidthTag is tag for Notch Filter Bandwidth
	NotchFilterBandwidthTag = 0x003A0223
	// WaveformDataDisplayScaleTag is tag for Waveform Data Display Scale
	WaveformDataDisplayScaleTag = 0x003A0230
	// WaveformDisplayBackgroundCIELabValueTag is tag for Waveform Display Background CIELab Value
	WaveformDisplayBackgroundCIELabValueTag = 0x003A0231
	// WaveformPresentationGroupSequenceTag is tag for Waveform Presentation Group Sequence
	WaveformPresentationGroupSequenceTag = 0x003A0240
	// PresentationGroupNumberTag is tag for Presentation Group Number
	PresentationGroupNumberTag = 0x003A0241
	// ChannelDisplaySequenceTag is tag for Channel Display Sequence
	ChannelDisplaySequenceTag = 0x003A0242
	// ChannelRecommendedDisplayCIELabValueTag is tag for Channel Recommended Display CIELab Value
	ChannelRecommendedDisplayCIELabValueTag = 0x003A0244
	// ChannelPositionTag is tag for Channel Position
	ChannelPositionTag = 0x003A0245
	// DisplayShadingFlagTag is tag for Display Shading Flag
	DisplayShadingFlagTag = 0x003A0246
	// FractionalChannelDisplayScaleTag is tag for Fractional Channel Display Scale
	FractionalChannelDisplayScaleTag = 0x003A0247
	// AbsoluteChannelDisplayScaleTag is tag for Absolute Channel Display Scale
	AbsoluteChannelDisplayScaleTag = 0x003A0248
	// MultiplexedAudioChannelsDescriptionCodeSequenceTag is tag for Multiplexed Audio Channels Description Code Sequence
	MultiplexedAudioChannelsDescriptionCodeSequenceTag = 0x003A0300
	// ChannelIdentificationCodeTag is tag for Channel Identification Code
	ChannelIdentificationCodeTag = 0x003A0301
	// ChannelModeTag is tag for Channel Mode
	ChannelModeTag = 0x003A0302
	// ScheduledStationAETitleTag is tag for Scheduled Station AE Title
	ScheduledStationAETitleTag = 0x00400001
	// ScheduledProcedureStepStartDateTag is tag for Scheduled Procedure Step Start Date
	ScheduledProcedureStepStartDateTag = 0x00400002
	// ScheduledProcedureStepStartTimeTag is tag for Scheduled Procedure Step Start Time
	ScheduledProcedureStepStartTimeTag = 0x00400003
	// ScheduledProcedureStepEndDateTag is tag for Scheduled Procedure Step End Date
	ScheduledProcedureStepEndDateTag = 0x00400004
	// ScheduledProcedureStepEndTimeTag is tag for Scheduled Procedure Step End Time
	ScheduledProcedureStepEndTimeTag = 0x00400005
	// ScheduledPerformingPhysicianNameTag is tag for Scheduled Performing Physician's Name
	ScheduledPerformingPhysicianNameTag = 0x00400006
	// ScheduledProcedureStepDescriptionTag is tag for Scheduled Procedure Step Description
	ScheduledProcedureStepDescriptionTag = 0x00400007
	// ScheduledProtocolCodeSequenceTag is tag for Scheduled Protocol Code Sequence
	ScheduledProtocolCodeSequenceTag = 0x00400008
	// ScheduledProcedureStepIDTag is tag for Scheduled Procedure Step ID
	ScheduledProcedureStepIDTag = 0x00400009
	// StageCodeSequenceTag is tag for Stage Code Sequence
	StageCodeSequenceTag = 0x0040000A
	// ScheduledPerformingPhysicianIdentificationSequenceTag is tag for Scheduled Performing Physician Identification Sequence
	ScheduledPerformingPhysicianIdentificationSequenceTag = 0x0040000B
	// ScheduledStationNameTag is tag for Scheduled Station Name
	ScheduledStationNameTag = 0x00400010
	// ScheduledProcedureStepLocationTag is tag for Scheduled Procedure Step Location
	ScheduledProcedureStepLocationTag = 0x00400011
	// PreMedicationTag is tag for Pre-Medication
	PreMedicationTag = 0x00400012
	// ScheduledProcedureStepStatusTag is tag for Scheduled Procedure Step Status
	ScheduledProcedureStepStatusTag = 0x00400020
	// OrderPlacerIdentifierSequenceTag is tag for Order Placer Identifier Sequence
	OrderPlacerIdentifierSequenceTag = 0x00400026
	// OrderFillerIdentifierSequenceTag is tag for Order Filler Identifier Sequence
	OrderFillerIdentifierSequenceTag = 0x00400027
	// LocalNamespaceEntityIDTag is tag for Local Namespace Entity ID
	LocalNamespaceEntityIDTag = 0x00400031
	// UniversalEntityIDTag is tag for Universal Entity ID
	UniversalEntityIDTag = 0x00400032
	// UniversalEntityIDTypeTag is tag for Universal Entity ID Type
	UniversalEntityIDTypeTag = 0x00400033
	// IdentifierTypeCodeTag is tag for Identifier Type Code
	IdentifierTypeCodeTag = 0x00400035
	// AssigningFacilitySequenceTag is tag for Assigning Facility Sequence
	AssigningFacilitySequenceTag = 0x00400036
	// AssigningJurisdictionCodeSequenceTag is tag for Assigning Jurisdiction Code Sequence
	AssigningJurisdictionCodeSequenceTag = 0x00400039
	// AssigningAgencyOrDepartmentCodeSequenceTag is tag for Assigning Agency or Department Code Sequence
	AssigningAgencyOrDepartmentCodeSequenceTag = 0x0040003A
	// ScheduledProcedureStepSequenceTag is tag for Scheduled Procedure Step Sequence
	ScheduledProcedureStepSequenceTag = 0x00400100
	// ReferencedNonImageCompositeSOPInstanceSequenceTag is tag for Referenced Non-Image Composite SOP Instance Sequence
	ReferencedNonImageCompositeSOPInstanceSequenceTag = 0x00400220
	// PerformedStationAETitleTag is tag for Performed Station AE Title
	PerformedStationAETitleTag = 0x00400241
	// PerformedStationNameTag is tag for Performed Station Name
	PerformedStationNameTag = 0x00400242
	// PerformedLocationTag is tag for Performed Location
	PerformedLocationTag = 0x00400243
	// PerformedProcedureStepStartDateTag is tag for Performed Procedure Step Start Date
	PerformedProcedureStepStartDateTag = 0x00400244
	// PerformedProcedureStepStartTimeTag is tag for Performed Procedure Step Start Time
	PerformedProcedureStepStartTimeTag = 0x00400245
	// PerformedProcedureStepEndDateTag is tag for Performed Procedure Step End Date
	PerformedProcedureStepEndDateTag = 0x00400250
	// PerformedProcedureStepEndTimeTag is tag for Performed Procedure Step End Time
	PerformedProcedureStepEndTimeTag = 0x00400251
	// PerformedProcedureStepStatusTag is tag for Performed Procedure Step Status
	PerformedProcedureStepStatusTag = 0x00400252
	// PerformedProcedureStepIDTag is tag for Performed Procedure Step ID
	PerformedProcedureStepIDTag = 0x00400253
	// PerformedProcedureStepDescriptionTag is tag for Performed Procedure Step Description
	PerformedProcedureStepDescriptionTag = 0x00400254
	// PerformedProcedureTypeDescriptionTag is tag for Performed Procedure Type Description
	PerformedProcedureTypeDescriptionTag = 0x00400255
	// PerformedProtocolCodeSequenceTag is tag for Performed Protocol Code Sequence
	PerformedProtocolCodeSequenceTag = 0x00400260
	// PerformedProtocolTypeTag is tag for Performed Protocol Type
	PerformedProtocolTypeTag = 0x00400261
	// ScheduledStepAttributesSequenceTag is tag for Scheduled Step Attributes Sequence
	ScheduledStepAttributesSequenceTag = 0x00400270
	// RequestAttributesSequenceTag is tag for Request Attributes Sequence
	RequestAttributesSequenceTag = 0x00400275
	// CommentsOnThePerformedProcedureStepTag is tag for Comments on the Performed Procedure Step
	CommentsOnThePerformedProcedureStepTag = 0x00400280
	// PerformedProcedureStepDiscontinuationReasonCodeSequenceTag is tag for Performed Procedure Step Discontinuation Reason Code Sequence
	PerformedProcedureStepDiscontinuationReasonCodeSequenceTag = 0x00400281
	// QuantitySequenceTag is tag for Quantity Sequence
	QuantitySequenceTag = 0x00400293
	// QuantityTag is tag for Quantity
	QuantityTag = 0x00400294
	// MeasuringUnitsSequenceTag is tag for Measuring Units Sequence
	MeasuringUnitsSequenceTag = 0x00400295
	// BillingItemSequenceTag is tag for Billing Item Sequence
	BillingItemSequenceTag = 0x00400296
	// TotalTimeOfFluoroscopyTag is tag for Total Time of Fluoroscopy
	TotalTimeOfFluoroscopyTag = 0x00400300
	// TotalNumberOfExposuresTag is tag for Total Number of Exposures
	TotalNumberOfExposuresTag = 0x00400301
	// EntranceDoseTag is tag for Entrance Dose
	EntranceDoseTag = 0x00400302
	// ExposedAreaTag is tag for Exposed Area
	ExposedAreaTag = 0x00400303
	// DistanceSourceToEntranceTag is tag for Distance Source to Entrance
	DistanceSourceToEntranceTag = 0x00400306
	// DistanceSourceToSupportTag is tag for Distance Source to Support
	DistanceSourceToSupportTag = 0x00400307
	// ExposureDoseSequenceTag is tag for Exposure Dose Sequence
	ExposureDoseSequenceTag = 0x0040030E
	// CommentsOnRadiationDoseTag is tag for Comments on Radiation Dose
	CommentsOnRadiationDoseTag = 0x00400310
	// XRayOutputTag is tag for X-Ray Output
	XRayOutputTag = 0x00400312
	// HalfValueLayerTag is tag for Half Value Layer
	HalfValueLayerTag = 0x00400314
	// OrganDoseTag is tag for Organ Dose
	OrganDoseTag = 0x00400316
	// OrganExposedTag is tag for Organ Exposed
	OrganExposedTag = 0x00400318
	// BillingProcedureStepSequenceTag is tag for Billing Procedure Step Sequence
	BillingProcedureStepSequenceTag = 0x00400320
	// FilmConsumptionSequenceTag is tag for Film Consumption Sequence
	FilmConsumptionSequenceTag = 0x00400321
	// BillingSuppliesAndDevicesSequenceTag is tag for Billing Supplies and Devices Sequence
	BillingSuppliesAndDevicesSequenceTag = 0x00400324
	// ReferencedProcedureStepSequenceTag is tag for Referenced Procedure Step Sequence
	ReferencedProcedureStepSequenceTag = 0x00400330
	// PerformedSeriesSequenceTag is tag for Performed Series Sequence
	PerformedSeriesSequenceTag = 0x00400340
	// CommentsOnTheScheduledProcedureStepTag is tag for Comments on the Scheduled Procedure Step
	CommentsOnTheScheduledProcedureStepTag = 0x00400400
	// ProtocolContextSequenceTag is tag for Protocol Context Sequence
	ProtocolContextSequenceTag = 0x00400440
	// ContentItemModifierSequenceTag is tag for Content Item Modifier Sequence
	ContentItemModifierSequenceTag = 0x00400441
	// ScheduledSpecimenSequenceTag is tag for Scheduled Specimen Sequence
	ScheduledSpecimenSequenceTag = 0x00400500
	// SpecimenAccessionNumberTag is tag for Specimen Accession Number
	SpecimenAccessionNumberTag = 0x0040050A
	// ContainerIdentifierTag is tag for Container Identifier
	ContainerIdentifierTag = 0x00400512
	// IssuerOfTheContainerIdentifierSequenceTag is tag for Issuer of the Container Identifier Sequence
	IssuerOfTheContainerIdentifierSequenceTag = 0x00400513
	// AlternateContainerIdentifierSequenceTag is tag for Alternate Container Identifier Sequence
	AlternateContainerIdentifierSequenceTag = 0x00400515
	// ContainerTypeCodeSequenceTag is tag for Container Type Code Sequence
	ContainerTypeCodeSequenceTag = 0x00400518
	// ContainerDescriptionTag is tag for Container Description
	ContainerDescriptionTag = 0x0040051A
	// ContainerComponentSequenceTag is tag for Container Component Sequence
	ContainerComponentSequenceTag = 0x00400520
	// SpecimenSequenceTag is tag for Specimen Sequence
	SpecimenSequenceTag = 0x00400550
	// SpecimenIdentifierTag is tag for Specimen Identifier
	SpecimenIdentifierTag = 0x00400551
	// SpecimenDescriptionSequenceTrialTag is tag for Specimen Description Sequence (Trial)
	SpecimenDescriptionSequenceTrialTag = 0x00400552
	// SpecimenDescriptionTrialTag is tag for Specimen Description (Trial)
	SpecimenDescriptionTrialTag = 0x00400553
	// SpecimenUIDTag is tag for Specimen UID
	SpecimenUIDTag = 0x00400554
	// AcquisitionContextSequenceTag is tag for Acquisition Context Sequence
	AcquisitionContextSequenceTag = 0x00400555
	// AcquisitionContextDescriptionTag is tag for Acquisition Context Description
	AcquisitionContextDescriptionTag = 0x00400556
	// SpecimenDescriptionSequenceTag is tag for Specimen Description Sequence
	SpecimenDescriptionSequenceTag = 0x00400560
	// IssuerOfTheSpecimenIdentifierSequenceTag is tag for Issuer of the Specimen Identifier Sequence
	IssuerOfTheSpecimenIdentifierSequenceTag = 0x00400562
	// SpecimenTypeCodeSequenceTag is tag for Specimen Type Code Sequence
	SpecimenTypeCodeSequenceTag = 0x0040059A
	// SpecimenShortDescriptionTag is tag for Specimen Short Description
	SpecimenShortDescriptionTag = 0x00400600
	// SpecimenDetailedDescriptionTag is tag for Specimen Detailed Description
	SpecimenDetailedDescriptionTag = 0x00400602
	// SpecimenPreparationSequenceTag is tag for Specimen Preparation Sequence
	SpecimenPreparationSequenceTag = 0x00400610
	// SpecimenPreparationStepContentItemSequenceTag is tag for Specimen Preparation Step Content Item Sequence
	SpecimenPreparationStepContentItemSequenceTag = 0x00400612
	// SpecimenLocalizationContentItemSequenceTag is tag for Specimen Localization Content Item Sequence
	SpecimenLocalizationContentItemSequenceTag = 0x00400620
	// SlideIdentifierTag is tag for Slide Identifier
	SlideIdentifierTag = 0x004006FA
	// WholeSlideMicroscopyImageFrameTypeSequenceTag is tag for Whole Slide Microscopy Image Frame Type Sequence
	WholeSlideMicroscopyImageFrameTypeSequenceTag = 0x00400710
	// ImageCenterPointCoordinatesSequenceTag is tag for Image Center Point Coordinates Sequence
	ImageCenterPointCoordinatesSequenceTag = 0x0040071A
	// XOffsetInSlideCoordinateSystemTag is tag for X Offset in Slide Coordinate System
	XOffsetInSlideCoordinateSystemTag = 0x0040072A
	// YOffsetInSlideCoordinateSystemTag is tag for Y Offset in Slide Coordinate System
	YOffsetInSlideCoordinateSystemTag = 0x0040073A
	// ZOffsetInSlideCoordinateSystemTag is tag for Z Offset in Slide Coordinate System
	ZOffsetInSlideCoordinateSystemTag = 0x0040074A
	// PixelSpacingSequenceTag is tag for Pixel Spacing Sequence
	PixelSpacingSequenceTag = 0x004008D8
	// CoordinateSystemAxisCodeSequenceTag is tag for Coordinate System Axis Code Sequence
	CoordinateSystemAxisCodeSequenceTag = 0x004008DA
	// MeasurementUnitsCodeSequenceTag is tag for Measurement Units Code Sequence
	MeasurementUnitsCodeSequenceTag = 0x004008EA
	// VitalStainCodeSequenceTrialTag is tag for Vital Stain Code Sequence (Trial)
	VitalStainCodeSequenceTrialTag = 0x004009F8
	// RequestedProcedureIDTag is tag for Requested Procedure ID
	RequestedProcedureIDTag = 0x00401001
	// ReasonForTheRequestedProcedureTag is tag for Reason for the Requested Procedure
	ReasonForTheRequestedProcedureTag = 0x00401002
	// RequestedProcedurePriorityTag is tag for Requested Procedure Priority
	RequestedProcedurePriorityTag = 0x00401003
	// PatientTransportArrangementsTag is tag for Patient Transport Arrangements
	PatientTransportArrangementsTag = 0x00401004
	// RequestedProcedureLocationTag is tag for Requested Procedure Location
	RequestedProcedureLocationTag = 0x00401005
	// PlacerOrderNumberProcedureTag is tag for Placer Order Number / Procedure
	PlacerOrderNumberProcedureTag = 0x00401006
	// FillerOrderNumberProcedureTag is tag for Filler Order Number / Procedure
	FillerOrderNumberProcedureTag = 0x00401007
	// ConfidentialityCodeTag is tag for Confidentiality Code
	ConfidentialityCodeTag = 0x00401008
	// ReportingPriorityTag is tag for Reporting Priority
	ReportingPriorityTag = 0x00401009
	// ReasonForRequestedProcedureCodeSequenceTag is tag for Reason for Requested Procedure Code Sequence
	ReasonForRequestedProcedureCodeSequenceTag = 0x0040100A
	// NamesOfIntendedRecipientsOfResultsTag is tag for Names of Intended Recipients of Results
	NamesOfIntendedRecipientsOfResultsTag = 0x00401010
	// IntendedRecipientsOfResultsIdentificationSequenceTag is tag for Intended Recipients of Results Identification Sequence
	IntendedRecipientsOfResultsIdentificationSequenceTag = 0x00401011
	// ReasonForPerformedProcedureCodeSequenceTag is tag for Reason For Performed Procedure Code Sequence
	ReasonForPerformedProcedureCodeSequenceTag = 0x00401012
	// RequestedProcedureDescriptionTrialTag is tag for Requested Procedure Description (Trial)
	RequestedProcedureDescriptionTrialTag = 0x00401060
	// PersonIdentificationCodeSequenceTag is tag for Person Identification Code Sequence
	PersonIdentificationCodeSequenceTag = 0x00401101
	// PersonAddressTag is tag for Person's Address
	PersonAddressTag = 0x00401102
	// PersonTelephoneNumbersTag is tag for Person's Telephone Numbers
	PersonTelephoneNumbersTag = 0x00401103
	// PersonTelecomInformationTag is tag for Person's Telecom Information
	PersonTelecomInformationTag = 0x00401104
	// RequestedProcedureCommentsTag is tag for Requested Procedure Comments
	RequestedProcedureCommentsTag = 0x00401400
	// ReasonForTheImagingServiceRequestTag is tag for Reason for the Imaging Service Request
	ReasonForTheImagingServiceRequestTag = 0x00402001
	// IssueDateOfImagingServiceRequestTag is tag for Issue Date of Imaging Service Request
	IssueDateOfImagingServiceRequestTag = 0x00402004
	// IssueTimeOfImagingServiceRequestTag is tag for Issue Time of Imaging Service Request
	IssueTimeOfImagingServiceRequestTag = 0x00402005
	// PlacerOrderNumberImagingServiceRequestRetiredTag is tag for Placer Order Number / Imaging Service Request (Retired)
	PlacerOrderNumberImagingServiceRequestRetiredTag = 0x00402006
	// FillerOrderNumberImagingServiceRequestRetiredTag is tag for Filler Order Number / Imaging Service Request (Retired)
	FillerOrderNumberImagingServiceRequestRetiredTag = 0x00402007
	// OrderEnteredByTag is tag for Order Entered By
	OrderEnteredByTag = 0x00402008
	// OrderEntererLocationTag is tag for Order Enterer's Location
	OrderEntererLocationTag = 0x00402009
	// OrderCallbackPhoneNumberTag is tag for Order Callback Phone Number
	OrderCallbackPhoneNumberTag = 0x00402010
	// OrderCallbackTelecomInformationTag is tag for Order Callback Telecom Information
	OrderCallbackTelecomInformationTag = 0x00402011
	// PlacerOrderNumberImagingServiceRequestTag is tag for Placer Order Number / Imaging Service Request
	PlacerOrderNumberImagingServiceRequestTag = 0x00402016
	// FillerOrderNumberImagingServiceRequestTag is tag for Filler Order Number / Imaging Service Request
	FillerOrderNumberImagingServiceRequestTag = 0x00402017
	// ImagingServiceRequestCommentsTag is tag for Imaging Service Request Comments
	ImagingServiceRequestCommentsTag = 0x00402400
	// ConfidentialityConstraintOnPatientDataDescriptionTag is tag for Confidentiality Constraint on Patient Data Description
	ConfidentialityConstraintOnPatientDataDescriptionTag = 0x00403001
	// GeneralPurposeScheduledProcedureStepStatusTag is tag for General Purpose Scheduled Procedure Step Status
	GeneralPurposeScheduledProcedureStepStatusTag = 0x00404001
	// GeneralPurposePerformedProcedureStepStatusTag is tag for General Purpose Performed Procedure Step Status
	GeneralPurposePerformedProcedureStepStatusTag = 0x00404002
	// GeneralPurposeScheduledProcedureStepPriorityTag is tag for General Purpose Scheduled Procedure Step Priority
	GeneralPurposeScheduledProcedureStepPriorityTag = 0x00404003
	// ScheduledProcessingApplicationsCodeSequenceTag is tag for Scheduled Processing Applications Code Sequence
	ScheduledProcessingApplicationsCodeSequenceTag = 0x00404004
	// ScheduledProcedureStepStartDateTimeTag is tag for Scheduled Procedure Step Start DateTime
	ScheduledProcedureStepStartDateTimeTag = 0x00404005
	// MultipleCopiesFlagTag is tag for Multiple Copies Flag
	MultipleCopiesFlagTag = 0x00404006
	// PerformedProcessingApplicationsCodeSequenceTag is tag for Performed Processing Applications Code Sequence
	PerformedProcessingApplicationsCodeSequenceTag = 0x00404007
	// ScheduledProcedureStepExpirationDateTimeTag is tag for Scheduled Procedure Step Expiration DateTime
	ScheduledProcedureStepExpirationDateTimeTag = 0x00404008
	// HumanPerformerCodeSequenceTag is tag for Human Performer Code Sequence
	HumanPerformerCodeSequenceTag = 0x00404009
	// ScheduledProcedureStepModificationDateTimeTag is tag for Scheduled Procedure Step Modification DateTime
	ScheduledProcedureStepModificationDateTimeTag = 0x00404010
	// ExpectedCompletionDateTimeTag is tag for Expected Completion DateTime
	ExpectedCompletionDateTimeTag = 0x00404011
	// ResultingGeneralPurposePerformedProcedureStepsSequenceTag is tag for Resulting General Purpose Performed Procedure Steps Sequence
	ResultingGeneralPurposePerformedProcedureStepsSequenceTag = 0x00404015
	// ReferencedGeneralPurposeScheduledProcedureStepSequenceTag is tag for Referenced General Purpose Scheduled Procedure Step Sequence
	ReferencedGeneralPurposeScheduledProcedureStepSequenceTag = 0x00404016
	// ScheduledWorkitemCodeSequenceTag is tag for Scheduled Workitem Code Sequence
	ScheduledWorkitemCodeSequenceTag = 0x00404018
	// PerformedWorkitemCodeSequenceTag is tag for Performed Workitem Code Sequence
	PerformedWorkitemCodeSequenceTag = 0x00404019
	// InputAvailabilityFlagTag is tag for Input Availability Flag
	InputAvailabilityFlagTag = 0x00404020
	// InputInformationSequenceTag is tag for Input Information Sequence
	InputInformationSequenceTag = 0x00404021
	// RelevantInformationSequenceTag is tag for Relevant Information Sequence
	RelevantInformationSequenceTag = 0x00404022
	// ReferencedGeneralPurposeScheduledProcedureStepTransactionUIDTag is tag for Referenced General Purpose Scheduled Procedure Step Transaction UID
	ReferencedGeneralPurposeScheduledProcedureStepTransactionUIDTag = 0x00404023
	// ScheduledStationNameCodeSequenceTag is tag for Scheduled Station Name Code Sequence
	ScheduledStationNameCodeSequenceTag = 0x00404025
	// ScheduledStationClassCodeSequenceTag is tag for Scheduled Station Class Code Sequence
	ScheduledStationClassCodeSequenceTag = 0x00404026
	// ScheduledStationGeographicLocationCodeSequenceTag is tag for Scheduled Station Geographic Location Code Sequence
	ScheduledStationGeographicLocationCodeSequenceTag = 0x00404027
	// PerformedStationNameCodeSequenceTag is tag for Performed Station Name Code Sequence
	PerformedStationNameCodeSequenceTag = 0x00404028
	// PerformedStationClassCodeSequenceTag is tag for Performed Station Class Code Sequence
	PerformedStationClassCodeSequenceTag = 0x00404029
	// PerformedStationGeographicLocationCodeSequenceTag is tag for Performed Station Geographic Location Code Sequence
	PerformedStationGeographicLocationCodeSequenceTag = 0x00404030
	// RequestedSubsequentWorkitemCodeSequenceTag is tag for Requested Subsequent Workitem Code Sequence
	RequestedSubsequentWorkitemCodeSequenceTag = 0x00404031
	// NonDICOMOutputCodeSequenceTag is tag for Non-DICOM Output Code Sequence
	NonDICOMOutputCodeSequenceTag = 0x00404032
	// OutputInformationSequenceTag is tag for Output Information Sequence
	OutputInformationSequenceTag = 0x00404033
	// ScheduledHumanPerformersSequenceTag is tag for Scheduled Human Performers Sequence
	ScheduledHumanPerformersSequenceTag = 0x00404034
	// ActualHumanPerformersSequenceTag is tag for Actual Human Performers Sequence
	ActualHumanPerformersSequenceTag = 0x00404035
	// HumanPerformerOrganizationTag is tag for Human Performer's Organization
	HumanPerformerOrganizationTag = 0x00404036
	// HumanPerformerNameTag is tag for Human Performer's Name
	HumanPerformerNameTag = 0x00404037
	// RawDataHandlingTag is tag for Raw Data Handling
	RawDataHandlingTag = 0x00404040
	// InputReadinessStateTag is tag for Input Readiness State
	InputReadinessStateTag = 0x00404041
	// PerformedProcedureStepStartDateTimeTag is tag for Performed Procedure Step Start DateTime
	PerformedProcedureStepStartDateTimeTag = 0x00404050
	// PerformedProcedureStepEndDateTimeTag is tag for Performed Procedure Step End DateTime
	PerformedProcedureStepEndDateTimeTag = 0x00404051
	// ProcedureStepCancellationDateTimeTag is tag for Procedure Step Cancellation DateTime
	ProcedureStepCancellationDateTimeTag = 0x00404052
	// OutputDestinationSequenceTag is tag for Output Destination Sequence
	OutputDestinationSequenceTag = 0x00404070
	// DICOMStorageSequenceTag is tag for DICOM Storage Sequence
	DICOMStorageSequenceTag = 0x00404071
	// STOWRSStorageSequenceTag is tag for STOW-RS Storage Sequence
	STOWRSStorageSequenceTag = 0x00404072
	// StorageURLTag is tag for Storage URL
	StorageURLTag = 0x00404073
	// XDSStorageSequenceTag is tag for XDS Storage Sequence
	XDSStorageSequenceTag = 0x00404074
	// EntranceDoseInmGyTag is tag for Entrance Dose in mGy
	EntranceDoseInmGyTag = 0x00408302
	// EntranceDoseDerivationTag is tag for Entrance Dose Derivation
	EntranceDoseDerivationTag = 0x00408303
	// ParametricMapFrameTypeSequenceTag is tag for Parametric Map Frame Type Sequence
	ParametricMapFrameTypeSequenceTag = 0x00409092
	// ReferencedImageRealWorldValueMappingSequenceTag is tag for Referenced Image Real World Value Mapping Sequence
	ReferencedImageRealWorldValueMappingSequenceTag = 0x00409094
	// RealWorldValueMappingSequenceTag is tag for Real World Value Mapping Sequence
	RealWorldValueMappingSequenceTag = 0x00409096
	// PixelValueMappingCodeSequenceTag is tag for Pixel Value Mapping Code Sequence
	PixelValueMappingCodeSequenceTag = 0x00409098
	// LUTLabelTag is tag for LUT Label
	LUTLabelTag = 0x00409210
	// RealWorldValueLastValueMappedTag is tag for Real World Value Last Value Mapped
	RealWorldValueLastValueMappedTag = 0x00409211
	// RealWorldValueLUTDataTag is tag for Real World Value LUT Data
	RealWorldValueLUTDataTag = 0x00409212
	// DoubleFloatRealWorldValueLastValueMappedTag is tag for Double Float Real World Value Last Value Mapped
	DoubleFloatRealWorldValueLastValueMappedTag = 0x00409213
	// DoubleFloatRealWorldValueFirstValueMappedTag is tag for Double Float Real World Value First Value Mapped
	DoubleFloatRealWorldValueFirstValueMappedTag = 0x00409214
	// RealWorldValueFirstValueMappedTag is tag for Real World Value First Value Mapped
	RealWorldValueFirstValueMappedTag = 0x00409216
	// QuantityDefinitionSequenceTag is tag for Quantity Definition Sequence
	QuantityDefinitionSequenceTag = 0x00409220
	// RealWorldValueInterceptTag is tag for Real World Value Intercept
	RealWorldValueInterceptTag = 0x00409224
	// RealWorldValueSlopeTag is tag for Real World Value Slope
	RealWorldValueSlopeTag = 0x00409225
	// FindingsFlagTrialTag is tag for Findings Flag (Trial)
	FindingsFlagTrialTag = 0x0040A007
	// RelationshipTypeTag is tag for Relationship Type
	RelationshipTypeTag = 0x0040A010
	// FindingsSequenceTrialTag is tag for Findings Sequence (Trial)
	FindingsSequenceTrialTag = 0x0040A020
	// FindingsGroupUIDTrialTag is tag for Findings Group UID (Trial)
	FindingsGroupUIDTrialTag = 0x0040A021
	// ReferencedFindingsGroupUIDTrialTag is tag for Referenced Findings Group UID (Trial)
	ReferencedFindingsGroupUIDTrialTag = 0x0040A022
	// FindingsGroupRecordingDateTrialTag is tag for Findings Group Recording Date (Trial)
	FindingsGroupRecordingDateTrialTag = 0x0040A023
	// FindingsGroupRecordingTimeTrialTag is tag for Findings Group Recording Time (Trial)
	FindingsGroupRecordingTimeTrialTag = 0x0040A024
	// FindingsSourceCategoryCodeSequenceTrialTag is tag for Findings Source Category Code Sequence (Trial)
	FindingsSourceCategoryCodeSequenceTrialTag = 0x0040A026
	// VerifyingOrganizationTag is tag for Verifying Organization
	VerifyingOrganizationTag = 0x0040A027
	// DocumentingOrganizationIdentifierCodeSequenceTrialTag is tag for Documenting Organization Identifier Code Sequence (Trial)
	DocumentingOrganizationIdentifierCodeSequenceTrialTag = 0x0040A028
	// VerificationDateTimeTag is tag for Verification DateTime
	VerificationDateTimeTag = 0x0040A030
	// ObservationDateTimeTag is tag for Observation DateTime
	ObservationDateTimeTag = 0x0040A032
	// ValueTypeTag is tag for Value Type
	ValueTypeTag = 0x0040A040
	// ConceptNameCodeSequenceTag is tag for Concept Name Code Sequence
	ConceptNameCodeSequenceTag = 0x0040A043
	// MeasurementPrecisionDescriptionTrialTag is tag for Measurement Precision Description (Trial)
	MeasurementPrecisionDescriptionTrialTag = 0x0040A047
	// ContinuityOfContentTag is tag for Continuity Of Content
	ContinuityOfContentTag = 0x0040A050
	// UrgencyOrPriorityAlertsTrialTag is tag for Urgency or Priority Alerts (Trial)
	UrgencyOrPriorityAlertsTrialTag = 0x0040A057
	// SequencingIndicatorTrialTag is tag for Sequencing Indicator (Trial)
	SequencingIndicatorTrialTag = 0x0040A060
	// DocumentIdentifierCodeSequenceTrialTag is tag for Document Identifier Code Sequence (Trial)
	DocumentIdentifierCodeSequenceTrialTag = 0x0040A066
	// DocumentAuthorTrialTag is tag for Document Author (Trial)
	DocumentAuthorTrialTag = 0x0040A067
	// DocumentAuthorIdentifierCodeSequenceTrialTag is tag for Document Author Identifier Code Sequence (Trial)
	DocumentAuthorIdentifierCodeSequenceTrialTag = 0x0040A068
	// IdentifierCodeSequenceTrialTag is tag for Identifier Code Sequence (Trial)
	IdentifierCodeSequenceTrialTag = 0x0040A070
	// VerifyingObserverSequenceTag is tag for Verifying Observer Sequence
	VerifyingObserverSequenceTag = 0x0040A073
	// ObjectBinaryIdentifierTrialTag is tag for Object Binary Identifier (Trial)
	ObjectBinaryIdentifierTrialTag = 0x0040A074
	// VerifyingObserverNameTag is tag for Verifying Observer Name
	VerifyingObserverNameTag = 0x0040A075
	// DocumentingObserverIdentifierCodeSequenceTrialTag is tag for Documenting Observer Identifier Code Sequence (Trial)
	DocumentingObserverIdentifierCodeSequenceTrialTag = 0x0040A076
	// AuthorObserverSequenceTag is tag for Author Observer Sequence
	AuthorObserverSequenceTag = 0x0040A078
	// ParticipantSequenceTag is tag for Participant Sequence
	ParticipantSequenceTag = 0x0040A07A
	// CustodialOrganizationSequenceTag is tag for Custodial Organization Sequence
	CustodialOrganizationSequenceTag = 0x0040A07C
	// ParticipationTypeTag is tag for Participation Type
	ParticipationTypeTag = 0x0040A080
	// ParticipationDateTimeTag is tag for Participation DateTime
	ParticipationDateTimeTag = 0x0040A082
	// ObserverTypeTag is tag for Observer Type
	ObserverTypeTag = 0x0040A084
	// ProcedureIdentifierCodeSequenceTrialTag is tag for Procedure Identifier Code Sequence (Trial)
	ProcedureIdentifierCodeSequenceTrialTag = 0x0040A085
	// VerifyingObserverIdentificationCodeSequenceTag is tag for Verifying Observer Identification Code Sequence
	VerifyingObserverIdentificationCodeSequenceTag = 0x0040A088
	// ObjectDirectoryBinaryIdentifierTrialTag is tag for Object Directory Binary Identifier (Trial)
	ObjectDirectoryBinaryIdentifierTrialTag = 0x0040A089
	// EquivalentCDADocumentSequenceTag is tag for Equivalent CDA Document Sequence
	EquivalentCDADocumentSequenceTag = 0x0040A090
	// ReferencedWaveformChannelsTag is tag for Referenced Waveform Channels
	ReferencedWaveformChannelsTag = 0x0040A0B0
	// DateOfDocumentOrVerbalTransactionTrialTag is tag for Date of Document or Verbal Transaction (Trial)
	DateOfDocumentOrVerbalTransactionTrialTag = 0x0040A110
	// TimeOfDocumentCreationOrVerbalTransactionTrialTag is tag for Time of Document Creation or Verbal Transaction (Trial)
	TimeOfDocumentCreationOrVerbalTransactionTrialTag = 0x0040A112
	// DateTimeTag is tag for DateTime
	DateTimeTag = 0x0040A120
	// DateTag is tag for Date
	DateTag = 0x0040A121
	// TimeTag is tag for Time
	TimeTag = 0x0040A122
	// PersonNameTag is tag for Person Name
	PersonNameTag = 0x0040A123
	// UIDTag is tag for UID
	UIDTag = 0x0040A124
	// ReportStatusIDTrialTag is tag for Report Status ID (Trial)
	ReportStatusIDTrialTag = 0x0040A125
	// TemporalRangeTypeTag is tag for Temporal Range Type
	TemporalRangeTypeTag = 0x0040A130
	// ReferencedSamplePositionsTag is tag for Referenced Sample Positions
	ReferencedSamplePositionsTag = 0x0040A132
	// ReferencedFrameNumbersTag is tag for Referenced Frame Numbers
	ReferencedFrameNumbersTag = 0x0040A136
	// ReferencedTimeOffsetsTag is tag for Referenced Time Offsets
	ReferencedTimeOffsetsTag = 0x0040A138
	// ReferencedDateTimeTag is tag for Referenced DateTime
	ReferencedDateTimeTag = 0x0040A13A
	// TextValueTag is tag for Text Value
	TextValueTag = 0x0040A160
	// FloatingPointValueTag is tag for Floating Point Value
	FloatingPointValueTag = 0x0040A161
	// RationalNumeratorValueTag is tag for Rational Numerator Value
	RationalNumeratorValueTag = 0x0040A162
	// RationalDenominatorValueTag is tag for Rational Denominator Value
	RationalDenominatorValueTag = 0x0040A163
	// ObservationCategoryCodeSequenceTrialTag is tag for Observation Category Code Sequence (Trial)
	ObservationCategoryCodeSequenceTrialTag = 0x0040A167
	// ConceptCodeSequenceTag is tag for Concept Code Sequence
	ConceptCodeSequenceTag = 0x0040A168
	// BibliographicCitationTrialTag is tag for Bibliographic Citation (Trial)
	BibliographicCitationTrialTag = 0x0040A16A
	// PurposeOfReferenceCodeSequenceTag is tag for Purpose of Reference Code Sequence
	PurposeOfReferenceCodeSequenceTag = 0x0040A170
	// ObservationUIDTag is tag for Observation UID
	ObservationUIDTag = 0x0040A171
	// ReferencedObservationUIDTrialTag is tag for Referenced Observation UID (Trial)
	ReferencedObservationUIDTrialTag = 0x0040A172
	// ReferencedObservationClassTrialTag is tag for Referenced Observation Class (Trial)
	ReferencedObservationClassTrialTag = 0x0040A173
	// ReferencedObjectObservationClassTrialTag is tag for Referenced Object Observation Class (Trial)
	ReferencedObjectObservationClassTrialTag = 0x0040A174
	// AnnotationGroupNumberTag is tag for Annotation Group Number
	AnnotationGroupNumberTag = 0x0040A180
	// ObservationDateTrialTag is tag for Observation Date (Trial)
	ObservationDateTrialTag = 0x0040A192
	// ObservationTimeTrialTag is tag for Observation Time (Trial)
	ObservationTimeTrialTag = 0x0040A193
	// MeasurementAutomationTrialTag is tag for Measurement Automation (Trial)
	MeasurementAutomationTrialTag = 0x0040A194
	// ModifierCodeSequenceTag is tag for Modifier Code Sequence
	ModifierCodeSequenceTag = 0x0040A195
	// IdentificationDescriptionTrialTag is tag for Identification Description (Trial)
	IdentificationDescriptionTrialTag = 0x0040A224
	// CoordinatesSetGeometricTypeTrialTag is tag for Coordinates Set Geometric Type (Trial)
	CoordinatesSetGeometricTypeTrialTag = 0x0040A290
	// AlgorithmCodeSequenceTrialTag is tag for Algorithm Code Sequence (Trial)
	AlgorithmCodeSequenceTrialTag = 0x0040A296
	// AlgorithmDescriptionTrialTag is tag for Algorithm Description (Trial)
	AlgorithmDescriptionTrialTag = 0x0040A297
	// PixelCoordinatesSetTrialTag is tag for Pixel Coordinates Set (Trial)
	PixelCoordinatesSetTrialTag = 0x0040A29A
	// MeasuredValueSequenceTag is tag for Measured Value Sequence
	MeasuredValueSequenceTag = 0x0040A300
	// NumericValueQualifierCodeSequenceTag is tag for Numeric Value Qualifier Code Sequence
	NumericValueQualifierCodeSequenceTag = 0x0040A301
	// CurrentObserverTrialTag is tag for Current Observer (Trial)
	CurrentObserverTrialTag = 0x0040A307
	// NumericValueTag is tag for Numeric Value
	NumericValueTag = 0x0040A30A
	// ReferencedAccessionSequenceTrialTag is tag for Referenced Accession Sequence (Trial)
	ReferencedAccessionSequenceTrialTag = 0x0040A313
	// ReportStatusCommentTrialTag is tag for Report Status Comment (Trial)
	ReportStatusCommentTrialTag = 0x0040A33A
	// ProcedureContextSequenceTrialTag is tag for Procedure Context Sequence (Trial)
	ProcedureContextSequenceTrialTag = 0x0040A340
	// VerbalSourceTrialTag is tag for Verbal Source (Trial)
	VerbalSourceTrialTag = 0x0040A352
	// AddressTrialTag is tag for Address (Trial)
	AddressTrialTag = 0x0040A353
	// TelephoneNumberTrialTag is tag for Telephone Number (Trial)
	TelephoneNumberTrialTag = 0x0040A354
	// VerbalSourceIdentifierCodeSequenceTrialTag is tag for Verbal Source Identifier Code Sequence (Trial)
	VerbalSourceIdentifierCodeSequenceTrialTag = 0x0040A358
	// PredecessorDocumentsSequenceTag is tag for Predecessor Documents Sequence
	PredecessorDocumentsSequenceTag = 0x0040A360
	// ReferencedRequestSequenceTag is tag for Referenced Request Sequence
	ReferencedRequestSequenceTag = 0x0040A370
	// PerformedProcedureCodeSequenceTag is tag for Performed Procedure Code Sequence
	PerformedProcedureCodeSequenceTag = 0x0040A372
	// CurrentRequestedProcedureEvidenceSequenceTag is tag for Current Requested Procedure Evidence Sequence
	CurrentRequestedProcedureEvidenceSequenceTag = 0x0040A375
	// ReportDetailSequenceTrialTag is tag for Report Detail Sequence (Trial)
	ReportDetailSequenceTrialTag = 0x0040A380
	// PertinentOtherEvidenceSequenceTag is tag for Pertinent Other Evidence Sequence
	PertinentOtherEvidenceSequenceTag = 0x0040A385
	// HL7StructuredDocumentReferenceSequenceTag is tag for HL7 Structured Document Reference Sequence
	HL7StructuredDocumentReferenceSequenceTag = 0x0040A390
	// ObservationSubjectUIDTrialTag is tag for Observation Subject UID (Trial)
	ObservationSubjectUIDTrialTag = 0x0040A402
	// ObservationSubjectClassTrialTag is tag for Observation Subject Class (Trial)
	ObservationSubjectClassTrialTag = 0x0040A403
	// ObservationSubjectTypeCodeSequenceTrialTag is tag for Observation Subject Type Code Sequence (Trial)
	ObservationSubjectTypeCodeSequenceTrialTag = 0x0040A404
	// CompletionFlagTag is tag for Completion Flag
	CompletionFlagTag = 0x0040A491
	// CompletionFlagDescriptionTag is tag for Completion Flag Description
	CompletionFlagDescriptionTag = 0x0040A492
	// VerificationFlagTag is tag for Verification Flag
	VerificationFlagTag = 0x0040A493
	// ArchiveRequestedTag is tag for Archive Requested
	ArchiveRequestedTag = 0x0040A494
	// PreliminaryFlagTag is tag for Preliminary Flag
	PreliminaryFlagTag = 0x0040A496
	// ContentTemplateSequenceTag is tag for Content Template Sequence
	ContentTemplateSequenceTag = 0x0040A504
	// IdenticalDocumentsSequenceTag is tag for Identical Documents Sequence
	IdenticalDocumentsSequenceTag = 0x0040A525
	// ObservationSubjectContextFlagTrialTag is tag for Observation Subject Context Flag (Trial)
	ObservationSubjectContextFlagTrialTag = 0x0040A600
	// ObserverContextFlagTrialTag is tag for Observer Context Flag (Trial)
	ObserverContextFlagTrialTag = 0x0040A601
	// ProcedureContextFlagTrialTag is tag for Procedure Context Flag (Trial)
	ProcedureContextFlagTrialTag = 0x0040A603
	// ContentSequenceTag is tag for Content Sequence
	ContentSequenceTag = 0x0040A730
	// RelationshipSequenceTrialTag is tag for Relationship Sequence (Trial)
	RelationshipSequenceTrialTag = 0x0040A731
	// RelationshipTypeCodeSequenceTrialTag is tag for Relationship Type Code Sequence (Trial)
	RelationshipTypeCodeSequenceTrialTag = 0x0040A732
	// LanguageCodeSequenceTrialTag is tag for Language Code Sequence (Trial)
	LanguageCodeSequenceTrialTag = 0x0040A744
	// UniformResourceLocatorTrialTag is tag for Uniform Resource Locator (Trial)
	UniformResourceLocatorTrialTag = 0x0040A992
	// WaveformAnnotationSequenceTag is tag for Waveform Annotation Sequence
	WaveformAnnotationSequenceTag = 0x0040B020
	// TemplateIdentifierTag is tag for Template Identifier
	TemplateIdentifierTag = 0x0040DB00
	// TemplateVersionTag is tag for Template Version
	TemplateVersionTag = 0x0040DB06
	// TemplateLocalVersionTag is tag for Template Local Version
	TemplateLocalVersionTag = 0x0040DB07
	// TemplateExtensionFlagTag is tag for Template Extension Flag
	TemplateExtensionFlagTag = 0x0040DB0B
	// TemplateExtensionOrganizationUIDTag is tag for Template Extension Organization UID
	TemplateExtensionOrganizationUIDTag = 0x0040DB0C
	// TemplateExtensionCreatorUIDTag is tag for Template Extension Creator UID
	TemplateExtensionCreatorUIDTag = 0x0040DB0D
	// ReferencedContentItemIdentifierTag is tag for Referenced Content Item Identifier
	ReferencedContentItemIdentifierTag = 0x0040DB73
	// HL7InstanceIdentifierTag is tag for HL7 Instance Identifier
	HL7InstanceIdentifierTag = 0x0040E001
	// HL7DocumentEffectiveTimeTag is tag for HL7 Document Effective Time
	HL7DocumentEffectiveTimeTag = 0x0040E004
	// HL7DocumentTypeCodeSequenceTag is tag for HL7 Document Type Code Sequence
	HL7DocumentTypeCodeSequenceTag = 0x0040E006
	// DocumentClassCodeSequenceTag is tag for Document Class Code Sequence
	DocumentClassCodeSequenceTag = 0x0040E008
	// RetrieveURITag is tag for Retrieve URI
	RetrieveURITag = 0x0040E010
	// RetrieveLocationUIDTag is tag for Retrieve Location UID
	RetrieveLocationUIDTag = 0x0040E011
	// TypeOfInstancesTag is tag for Type of Instances
	TypeOfInstancesTag = 0x0040E020
	// DICOMRetrievalSequenceTag is tag for DICOM Retrieval Sequence
	DICOMRetrievalSequenceTag = 0x0040E021
	// DICOMMediaRetrievalSequenceTag is tag for DICOM Media Retrieval Sequence
	DICOMMediaRetrievalSequenceTag = 0x0040E022
	// WADORetrievalSequenceTag is tag for WADO Retrieval Sequence
	WADORetrievalSequenceTag = 0x0040E023
	// XDSRetrievalSequenceTag is tag for XDS Retrieval Sequence
	XDSRetrievalSequenceTag = 0x0040E024
	// WADORSRetrievalSequenceTag is tag for WADO-RS Retrieval Sequence
	WADORSRetrievalSequenceTag = 0x0040E025
	// RepositoryUniqueIDTag is tag for Repository Unique ID
	RepositoryUniqueIDTag = 0x0040E030
	// HomeCommunityIDTag is tag for Home Community ID
	HomeCommunityIDTag = 0x0040E031
	// DocumentTitleTag is tag for Document Title
	DocumentTitleTag = 0x00420010
	// EncapsulatedDocumentTag is tag for Encapsulated Document
	EncapsulatedDocumentTag = 0x00420011
	// MIMETypeOfEncapsulatedDocumentTag is tag for MIME Type of Encapsulated Document
	MIMETypeOfEncapsulatedDocumentTag = 0x00420012
	// SourceInstanceSequenceTag is tag for Source Instance Sequence
	SourceInstanceSequenceTag = 0x00420013
	// ListOfMIMETypesTag is tag for List of MIME Types
	ListOfMIMETypesTag = 0x00420014
	// EncapsulatedDocumentLengthTag is tag for Encapsulated Document Length
	EncapsulatedDocumentLengthTag = 0x00420015
	// ProductPackageIdentifierTag is tag for Product Package Identifier
	ProductPackageIdentifierTag = 0x00440001
	// SubstanceAdministrationApprovalTag is tag for Substance Administration Approval
	SubstanceAdministrationApprovalTag = 0x00440002
	// ApprovalStatusFurtherDescriptionTag is tag for Approval Status Further Description
	ApprovalStatusFurtherDescriptionTag = 0x00440003
	// ApprovalStatusDateTimeTag is tag for Approval Status DateTime
	ApprovalStatusDateTimeTag = 0x00440004
	// ProductTypeCodeSequenceTag is tag for Product Type Code Sequence
	ProductTypeCodeSequenceTag = 0x00440007
	// ProductNameTag is tag for Product Name
	ProductNameTag = 0x00440008
	// ProductDescriptionTag is tag for Product Description
	ProductDescriptionTag = 0x00440009
	// ProductLotIdentifierTag is tag for Product Lot Identifier
	ProductLotIdentifierTag = 0x0044000A
	// ProductExpirationDateTimeTag is tag for Product Expiration DateTime
	ProductExpirationDateTimeTag = 0x0044000B
	// SubstanceAdministrationDateTimeTag is tag for Substance Administration DateTime
	SubstanceAdministrationDateTimeTag = 0x00440010
	// SubstanceAdministrationNotesTag is tag for Substance Administration Notes
	SubstanceAdministrationNotesTag = 0x00440011
	// SubstanceAdministrationDeviceIDTag is tag for Substance Administration Device ID
	SubstanceAdministrationDeviceIDTag = 0x00440012
	// ProductParameterSequenceTag is tag for Product Parameter Sequence
	ProductParameterSequenceTag = 0x00440013
	// SubstanceAdministrationParameterSequenceTag is tag for Substance Administration Parameter Sequence
	SubstanceAdministrationParameterSequenceTag = 0x00440019
	// ApprovalSequenceTag is tag for Approval Sequence
	ApprovalSequenceTag = 0x00440100
	// AssertionCodeSequenceTag is tag for Assertion Code Sequence
	AssertionCodeSequenceTag = 0x00440101
	// AssertionUIDTag is tag for Assertion UID
	AssertionUIDTag = 0x00440102
	// AsserterIdentificationSequenceTag is tag for Asserter Identification Sequence
	AsserterIdentificationSequenceTag = 0x00440103
	// AssertionDateTimeTag is tag for Assertion DateTime
	AssertionDateTimeTag = 0x00440104
	// AssertionExpirationDateTimeTag is tag for Assertion Expiration DateTime
	AssertionExpirationDateTimeTag = 0x00440105
	// AssertionCommentsTag is tag for Assertion Comments
	AssertionCommentsTag = 0x00440106
	// RelatedAssertionSequenceTag is tag for Related Assertion Sequence
	RelatedAssertionSequenceTag = 0x00440107
	// ReferencedAssertionUIDTag is tag for Referenced Assertion UID
	ReferencedAssertionUIDTag = 0x00440108
	// ApprovalSubjectSequenceTag is tag for Approval Subject Sequence
	ApprovalSubjectSequenceTag = 0x00440109
	// OrganizationalRoleCodeSequenceTag is tag for Organizational Role Code Sequence
	OrganizationalRoleCodeSequenceTag = 0x0044010A
	// LensDescriptionTag is tag for Lens Description
	LensDescriptionTag = 0x00460012
	// RightLensSequenceTag is tag for Right Lens Sequence
	RightLensSequenceTag = 0x00460014
	// LeftLensSequenceTag is tag for Left Lens Sequence
	LeftLensSequenceTag = 0x00460015
	// UnspecifiedLateralityLensSequenceTag is tag for Unspecified Laterality Lens Sequence
	UnspecifiedLateralityLensSequenceTag = 0x00460016
	// CylinderSequenceTag is tag for Cylinder Sequence
	CylinderSequenceTag = 0x00460018
	// PrismSequenceTag is tag for Prism Sequence
	PrismSequenceTag = 0x00460028
	// HorizontalPrismPowerTag is tag for Horizontal Prism Power
	HorizontalPrismPowerTag = 0x00460030
	// HorizontalPrismBaseTag is tag for Horizontal Prism Base
	HorizontalPrismBaseTag = 0x00460032
	// VerticalPrismPowerTag is tag for Vertical Prism Power
	VerticalPrismPowerTag = 0x00460034
	// VerticalPrismBaseTag is tag for Vertical Prism Base
	VerticalPrismBaseTag = 0x00460036
	// LensSegmentTypeTag is tag for Lens Segment Type
	LensSegmentTypeTag = 0x00460038
	// OpticalTransmittanceTag is tag for Optical Transmittance
	OpticalTransmittanceTag = 0x00460040
	// ChannelWidthTag is tag for Channel Width
	ChannelWidthTag = 0x00460042
	// PupilSizeTag is tag for Pupil Size
	PupilSizeTag = 0x00460044
	// CornealSizeTag is tag for Corneal Size
	CornealSizeTag = 0x00460046
	// CornealSizeSequenceTag is tag for Corneal Size Sequence
	CornealSizeSequenceTag = 0x00460047
	// AutorefractionRightEyeSequenceTag is tag for Autorefraction Right Eye Sequence
	AutorefractionRightEyeSequenceTag = 0x00460050
	// AutorefractionLeftEyeSequenceTag is tag for Autorefraction Left Eye Sequence
	AutorefractionLeftEyeSequenceTag = 0x00460052
	// DistancePupillaryDistanceTag is tag for Distance Pupillary Distance
	DistancePupillaryDistanceTag = 0x00460060
	// NearPupillaryDistanceTag is tag for Near Pupillary Distance
	NearPupillaryDistanceTag = 0x00460062
	// IntermediatePupillaryDistanceTag is tag for Intermediate Pupillary Distance
	IntermediatePupillaryDistanceTag = 0x00460063
	// OtherPupillaryDistanceTag is tag for Other Pupillary Distance
	OtherPupillaryDistanceTag = 0x00460064
	// KeratometryRightEyeSequenceTag is tag for Keratometry Right Eye Sequence
	KeratometryRightEyeSequenceTag = 0x00460070
	// KeratometryLeftEyeSequenceTag is tag for Keratometry Left Eye Sequence
	KeratometryLeftEyeSequenceTag = 0x00460071
	// SteepKeratometricAxisSequenceTag is tag for Steep Keratometric Axis Sequence
	SteepKeratometricAxisSequenceTag = 0x00460074
	// RadiusOfCurvatureTag is tag for Radius of Curvature
	RadiusOfCurvatureTag = 0x00460075
	// KeratometricPowerTag is tag for Keratometric Power
	KeratometricPowerTag = 0x00460076
	// KeratometricAxisTag is tag for Keratometric Axis
	KeratometricAxisTag = 0x00460077
	// FlatKeratometricAxisSequenceTag is tag for Flat Keratometric Axis Sequence
	FlatKeratometricAxisSequenceTag = 0x00460080
	// BackgroundColorTag is tag for Background Color
	BackgroundColorTag = 0x00460092
	// OptotypeTag is tag for Optotype
	OptotypeTag = 0x00460094
	// OptotypePresentationTag is tag for Optotype Presentation
	OptotypePresentationTag = 0x00460095
	// SubjectiveRefractionRightEyeSequenceTag is tag for Subjective Refraction Right Eye Sequence
	SubjectiveRefractionRightEyeSequenceTag = 0x00460097
	// SubjectiveRefractionLeftEyeSequenceTag is tag for Subjective Refraction Left Eye Sequence
	SubjectiveRefractionLeftEyeSequenceTag = 0x00460098
	// AddNearSequenceTag is tag for Add Near Sequence
	AddNearSequenceTag = 0x00460100
	// AddIntermediateSequenceTag is tag for Add Intermediate Sequence
	AddIntermediateSequenceTag = 0x00460101
	// AddOtherSequenceTag is tag for Add Other Sequence
	AddOtherSequenceTag = 0x00460102
	// AddPowerTag is tag for Add Power
	AddPowerTag = 0x00460104
	// ViewingDistanceTag is tag for Viewing Distance
	ViewingDistanceTag = 0x00460106
	// CorneaMeasurementsSequenceTag is tag for Cornea Measurements Sequence
	CorneaMeasurementsSequenceTag = 0x00460110
	// SourceOfCorneaMeasurementDataCodeSequenceTag is tag for Source of Cornea Measurement Data Code Sequence
	SourceOfCorneaMeasurementDataCodeSequenceTag = 0x00460111
	// SteepCornealAxisSequenceTag is tag for Steep Corneal Axis Sequence
	SteepCornealAxisSequenceTag = 0x00460112
	// FlatCornealAxisSequenceTag is tag for Flat Corneal Axis Sequence
	FlatCornealAxisSequenceTag = 0x00460113
	// CornealPowerTag is tag for Corneal Power
	CornealPowerTag = 0x00460114
	// CornealAxisTag is tag for Corneal Axis
	CornealAxisTag = 0x00460115
	// CorneaMeasurementMethodCodeSequenceTag is tag for Cornea Measurement Method Code Sequence
	CorneaMeasurementMethodCodeSequenceTag = 0x00460116
	// RefractiveIndexOfCorneaTag is tag for Refractive Index of Cornea
	RefractiveIndexOfCorneaTag = 0x00460117
	// RefractiveIndexOfAqueousHumorTag is tag for Refractive Index of Aqueous Humor
	RefractiveIndexOfAqueousHumorTag = 0x00460118
	// VisualAcuityTypeCodeSequenceTag is tag for Visual Acuity Type Code Sequence
	VisualAcuityTypeCodeSequenceTag = 0x00460121
	// VisualAcuityRightEyeSequenceTag is tag for Visual Acuity Right Eye Sequence
	VisualAcuityRightEyeSequenceTag = 0x00460122
	// VisualAcuityLeftEyeSequenceTag is tag for Visual Acuity Left Eye Sequence
	VisualAcuityLeftEyeSequenceTag = 0x00460123
	// VisualAcuityBothEyesOpenSequenceTag is tag for Visual Acuity Both Eyes Open Sequence
	VisualAcuityBothEyesOpenSequenceTag = 0x00460124
	// ViewingDistanceTypeTag is tag for Viewing Distance Type
	ViewingDistanceTypeTag = 0x00460125
	// VisualAcuityModifiersTag is tag for Visual Acuity Modifiers
	VisualAcuityModifiersTag = 0x00460135
	// DecimalVisualAcuityTag is tag for Decimal Visual Acuity
	DecimalVisualAcuityTag = 0x00460137
	// OptotypeDetailedDefinitionTag is tag for Optotype Detailed Definition
	OptotypeDetailedDefinitionTag = 0x00460139
	// ReferencedRefractiveMeasurementsSequenceTag is tag for Referenced Refractive Measurements Sequence
	ReferencedRefractiveMeasurementsSequenceTag = 0x00460145
	// SpherePowerTag is tag for Sphere Power
	SpherePowerTag = 0x00460146
	// CylinderPowerTag is tag for Cylinder Power
	CylinderPowerTag = 0x00460147
	// CornealTopographySurfaceTag is tag for Corneal Topography Surface
	CornealTopographySurfaceTag = 0x00460201
	// CornealVertexLocationTag is tag for Corneal Vertex Location
	CornealVertexLocationTag = 0x00460202
	// PupilCentroidXCoordinateTag is tag for Pupil Centroid X-Coordinate
	PupilCentroidXCoordinateTag = 0x00460203
	// PupilCentroidYCoordinateTag is tag for Pupil Centroid Y-Coordinate
	PupilCentroidYCoordinateTag = 0x00460204
	// EquivalentPupilRadiusTag is tag for Equivalent Pupil Radius
	EquivalentPupilRadiusTag = 0x00460205
	// CornealTopographyMapTypeCodeSequenceTag is tag for Corneal Topography Map Type Code Sequence
	CornealTopographyMapTypeCodeSequenceTag = 0x00460207
	// VerticesOfTheOutlineOfPupilTag is tag for Vertices of the Outline of Pupil
	VerticesOfTheOutlineOfPupilTag = 0x00460208
	// CornealTopographyMappingNormalsSequenceTag is tag for Corneal Topography Mapping Normals Sequence
	CornealTopographyMappingNormalsSequenceTag = 0x00460210
	// MaximumCornealCurvatureSequenceTag is tag for Maximum Corneal Curvature Sequence
	MaximumCornealCurvatureSequenceTag = 0x00460211
	// MaximumCornealCurvatureTag is tag for Maximum Corneal Curvature
	MaximumCornealCurvatureTag = 0x00460212
	// MaximumCornealCurvatureLocationTag is tag for Maximum Corneal Curvature Location
	MaximumCornealCurvatureLocationTag = 0x00460213
	// MinimumKeratometricSequenceTag is tag for Minimum Keratometric Sequence
	MinimumKeratometricSequenceTag = 0x00460215
	// SimulatedKeratometricCylinderSequenceTag is tag for Simulated Keratometric Cylinder Sequence
	SimulatedKeratometricCylinderSequenceTag = 0x00460218
	// AverageCornealPowerTag is tag for Average Corneal Power
	AverageCornealPowerTag = 0x00460220
	// CornealISValueTag is tag for Corneal I-S Value
	CornealISValueTag = 0x00460224
	// AnalyzedAreaTag is tag for Analyzed Area
	AnalyzedAreaTag = 0x00460227
	// SurfaceRegularityIndexTag is tag for Surface Regularity Index
	SurfaceRegularityIndexTag = 0x00460230
	// SurfaceAsymmetryIndexTag is tag for Surface Asymmetry Index
	SurfaceAsymmetryIndexTag = 0x00460232
	// CornealEccentricityIndexTag is tag for Corneal Eccentricity Index
	CornealEccentricityIndexTag = 0x00460234
	// KeratoconusPredictionIndexTag is tag for Keratoconus Prediction Index
	KeratoconusPredictionIndexTag = 0x00460236
	// DecimalPotentialVisualAcuityTag is tag for Decimal Potential Visual Acuity
	DecimalPotentialVisualAcuityTag = 0x00460238
	// CornealTopographyMapQualityEvaluationTag is tag for Corneal Topography Map Quality Evaluation
	CornealTopographyMapQualityEvaluationTag = 0x00460242
	// SourceImageCornealProcessedDataSequenceTag is tag for Source Image Corneal Processed Data Sequence
	SourceImageCornealProcessedDataSequenceTag = 0x00460244
	// CornealPointLocationTag is tag for Corneal Point Location
	CornealPointLocationTag = 0x00460247
	// CornealPointEstimatedTag is tag for Corneal Point Estimated
	CornealPointEstimatedTag = 0x00460248
	// AxialPowerTag is tag for Axial Power
	AxialPowerTag = 0x00460249
	// TangentialPowerTag is tag for Tangential Power
	TangentialPowerTag = 0x00460250
	// RefractivePowerTag is tag for Refractive Power
	RefractivePowerTag = 0x00460251
	// RelativeElevationTag is tag for Relative Elevation
	RelativeElevationTag = 0x00460252
	// CornealWavefrontTag is tag for Corneal Wavefront
	CornealWavefrontTag = 0x00460253
	// ImagedVolumeWidthTag is tag for Imaged Volume Width
	ImagedVolumeWidthTag = 0x00480001
	// ImagedVolumeHeightTag is tag for Imaged Volume Height
	ImagedVolumeHeightTag = 0x00480002
	// ImagedVolumeDepthTag is tag for Imaged Volume Depth
	ImagedVolumeDepthTag = 0x00480003
	// TotalPixelMatrixColumnsTag is tag for Total Pixel Matrix Columns
	TotalPixelMatrixColumnsTag = 0x00480006
	// TotalPixelMatrixRowsTag is tag for Total Pixel Matrix Rows
	TotalPixelMatrixRowsTag = 0x00480007
	// TotalPixelMatrixOriginSequenceTag is tag for Total Pixel Matrix Origin Sequence
	TotalPixelMatrixOriginSequenceTag = 0x00480008
	// SpecimenLabelInImageTag is tag for Specimen Label in Image
	SpecimenLabelInImageTag = 0x00480010
	// FocusMethodTag is tag for Focus Method
	FocusMethodTag = 0x00480011
	// ExtendedDepthOfFieldTag is tag for Extended Depth of Field
	ExtendedDepthOfFieldTag = 0x00480012
	// NumberOfFocalPlanesTag is tag for Number of Focal Planes
	NumberOfFocalPlanesTag = 0x00480013
	// DistanceBetweenFocalPlanesTag is tag for Distance Between Focal Planes
	DistanceBetweenFocalPlanesTag = 0x00480014
	// RecommendedAbsentPixelCIELabValueTag is tag for Recommended Absent Pixel CIELab Value
	RecommendedAbsentPixelCIELabValueTag = 0x00480015
	// IlluminatorTypeCodeSequenceTag is tag for Illuminator Type Code Sequence
	IlluminatorTypeCodeSequenceTag = 0x00480100
	// ImageOrientationSlideTag is tag for Image Orientation (Slide)
	ImageOrientationSlideTag = 0x00480102
	// OpticalPathSequenceTag is tag for Optical Path Sequence
	OpticalPathSequenceTag = 0x00480105
	// OpticalPathIdentifierTag is tag for Optical Path Identifier
	OpticalPathIdentifierTag = 0x00480106
	// OpticalPathDescriptionTag is tag for Optical Path Description
	OpticalPathDescriptionTag = 0x00480107
	// IlluminationColorCodeSequenceTag is tag for Illumination Color Code Sequence
	IlluminationColorCodeSequenceTag = 0x00480108
	// SpecimenReferenceSequenceTag is tag for Specimen Reference Sequence
	SpecimenReferenceSequenceTag = 0x00480110
	// CondenserLensPowerTag is tag for Condenser Lens Power
	CondenserLensPowerTag = 0x00480111
	// ObjectiveLensPowerTag is tag for Objective Lens Power
	ObjectiveLensPowerTag = 0x00480112
	// ObjectiveLensNumericalApertureTag is tag for Objective Lens Numerical Aperture
	ObjectiveLensNumericalApertureTag = 0x00480113
	// PaletteColorLookupTableSequenceTag is tag for Palette Color Lookup Table Sequence
	PaletteColorLookupTableSequenceTag = 0x00480120
	// ReferencedImageNavigationSequenceTag is tag for Referenced Image Navigation Sequence
	ReferencedImageNavigationSequenceTag = 0x00480200
	// TopLeftHandCornerOfLocalizerAreaTag is tag for Top Left Hand Corner of Localizer Area
	TopLeftHandCornerOfLocalizerAreaTag = 0x00480201
	// BottomRightHandCornerOfLocalizerAreaTag is tag for Bottom Right Hand Corner of Localizer Area
	BottomRightHandCornerOfLocalizerAreaTag = 0x00480202
	// OpticalPathIdentificationSequenceTag is tag for Optical Path Identification Sequence
	OpticalPathIdentificationSequenceTag = 0x00480207
	// PlanePositionSlideSequenceTag is tag for Plane Position (Slide) Sequence
	PlanePositionSlideSequenceTag = 0x0048021A
	// ColumnPositionInTotalImagePixelMatrixTag is tag for Column Position In Total Image Pixel Matrix
	ColumnPositionInTotalImagePixelMatrixTag = 0x0048021E
	// RowPositionInTotalImagePixelMatrixTag is tag for Row Position In Total Image Pixel Matrix
	RowPositionInTotalImagePixelMatrixTag = 0x0048021F
	// PixelOriginInterpretationTag is tag for Pixel Origin Interpretation
	PixelOriginInterpretationTag = 0x00480301
	// NumberOfOpticalPathsTag is tag for Number of Optical Paths
	NumberOfOpticalPathsTag = 0x00480302
	// TotalPixelMatrixFocalPlanesTag is tag for Total Pixel Matrix Focal Planes
	TotalPixelMatrixFocalPlanesTag = 0x00480303
	// CalibrationImageTag is tag for Calibration Image
	CalibrationImageTag = 0x00500004
	// DeviceSequenceTag is tag for Device Sequence
	DeviceSequenceTag = 0x00500010
	// ContainerComponentTypeCodeSequenceTag is tag for Container Component Type Code Sequence
	ContainerComponentTypeCodeSequenceTag = 0x00500012
	// ContainerComponentThicknessTag is tag for Container Component Thickness
	ContainerComponentThicknessTag = 0x00500013
	// DeviceLengthTag is tag for Device Length
	DeviceLengthTag = 0x00500014
	// ContainerComponentWidthTag is tag for Container Component Width
	ContainerComponentWidthTag = 0x00500015
	// DeviceDiameterTag is tag for Device Diameter
	DeviceDiameterTag = 0x00500016
	// DeviceDiameterUnitsTag is tag for Device Diameter Units
	DeviceDiameterUnitsTag = 0x00500017
	// DeviceVolumeTag is tag for Device Volume
	DeviceVolumeTag = 0x00500018
	// InterMarkerDistanceTag is tag for Inter-Marker Distance
	InterMarkerDistanceTag = 0x00500019
	// ContainerComponentMaterialTag is tag for Container Component Material
	ContainerComponentMaterialTag = 0x0050001A
	// ContainerComponentIDTag is tag for Container Component ID
	ContainerComponentIDTag = 0x0050001B
	// ContainerComponentLengthTag is tag for Container Component Length
	ContainerComponentLengthTag = 0x0050001C
	// ContainerComponentDiameterTag is tag for Container Component Diameter
	ContainerComponentDiameterTag = 0x0050001D
	// ContainerComponentDescriptionTag is tag for Container Component Description
	ContainerComponentDescriptionTag = 0x0050001E
	// DeviceDescriptionTag is tag for Device Description
	DeviceDescriptionTag = 0x00500020
	// LongDeviceDescriptionTag is tag for Long Device Description
	LongDeviceDescriptionTag = 0x00500021
	// ContrastBolusIngredientPercentByVolumeTag is tag for Contrast/Bolus Ingredient Percent by Volume
	ContrastBolusIngredientPercentByVolumeTag = 0x00520001
	// OCTFocalDistanceTag is tag for OCT Focal Distance
	OCTFocalDistanceTag = 0x00520002
	// BeamSpotSizeTag is tag for Beam Spot Size
	BeamSpotSizeTag = 0x00520003
	// EffectiveRefractiveIndexTag is tag for Effective Refractive Index
	EffectiveRefractiveIndexTag = 0x00520004
	// OCTAcquisitionDomainTag is tag for OCT Acquisition Domain
	OCTAcquisitionDomainTag = 0x00520006
	// OCTOpticalCenterWavelengthTag is tag for OCT Optical Center Wavelength
	OCTOpticalCenterWavelengthTag = 0x00520007
	// AxialResolutionTag is tag for Axial Resolution
	AxialResolutionTag = 0x00520008
	// RangingDepthTag is tag for Ranging Depth
	RangingDepthTag = 0x00520009
	// ALineRateTag is tag for A-line Rate
	ALineRateTag = 0x00520011
	// ALinesPerFrameTag is tag for A-lines Per Frame
	ALinesPerFrameTag = 0x00520012
	// CatheterRotationalRateTag is tag for Catheter Rotational Rate
	CatheterRotationalRateTag = 0x00520013
	// ALinePixelSpacingTag is tag for A-line Pixel Spacing
	ALinePixelSpacingTag = 0x00520014
	// ModeOfPercutaneousAccessSequenceTag is tag for Mode of Percutaneous Access Sequence
	ModeOfPercutaneousAccessSequenceTag = 0x00520016
	// IntravascularOCTFrameTypeSequenceTag is tag for Intravascular OCT Frame Type Sequence
	IntravascularOCTFrameTypeSequenceTag = 0x00520025
	// OCTZOffsetAppliedTag is tag for OCT Z Offset Applied
	OCTZOffsetAppliedTag = 0x00520026
	// IntravascularFrameContentSequenceTag is tag for Intravascular Frame Content Sequence
	IntravascularFrameContentSequenceTag = 0x00520027
	// IntravascularLongitudinalDistanceTag is tag for Intravascular Longitudinal Distance
	IntravascularLongitudinalDistanceTag = 0x00520028
	// IntravascularOCTFrameContentSequenceTag is tag for Intravascular OCT Frame Content Sequence
	IntravascularOCTFrameContentSequenceTag = 0x00520029
	// OCTZOffsetCorrectionTag is tag for OCT Z Offset Correction
	OCTZOffsetCorrectionTag = 0x00520030
	// CatheterDirectionOfRotationTag is tag for Catheter Direction of Rotation
	CatheterDirectionOfRotationTag = 0x00520031
	// SeamLineLocationTag is tag for Seam Line Location
	SeamLineLocationTag = 0x00520033
	// FirstALineLocationTag is tag for First A-line Location
	FirstALineLocationTag = 0x00520034
	// SeamLineIndexTag is tag for Seam Line Index
	SeamLineIndexTag = 0x00520036
	// NumberOfPaddedALinesTag is tag for Number of Padded A-lines
	NumberOfPaddedALinesTag = 0x00520038
	// InterpolationTypeTag is tag for Interpolation Type
	InterpolationTypeTag = 0x00520039
	// RefractiveIndexAppliedTag is tag for Refractive Index Applied
	RefractiveIndexAppliedTag = 0x0052003A
	// EnergyWindowVectorTag is tag for Energy Window Vector
	EnergyWindowVectorTag = 0x00540010
	// NumberOfEnergyWindowsTag is tag for Number of Energy Windows
	NumberOfEnergyWindowsTag = 0x00540011
	// EnergyWindowInformationSequenceTag is tag for Energy Window Information Sequence
	EnergyWindowInformationSequenceTag = 0x00540012
	// EnergyWindowRangeSequenceTag is tag for Energy Window Range Sequence
	EnergyWindowRangeSequenceTag = 0x00540013
	// EnergyWindowLowerLimitTag is tag for Energy Window Lower Limit
	EnergyWindowLowerLimitTag = 0x00540014
	// EnergyWindowUpperLimitTag is tag for Energy Window Upper Limit
	EnergyWindowUpperLimitTag = 0x00540015
	// RadiopharmaceuticalInformationSequenceTag is tag for Radiopharmaceutical Information Sequence
	RadiopharmaceuticalInformationSequenceTag = 0x00540016
	// ResidualSyringeCountsTag is tag for Residual Syringe Counts
	ResidualSyringeCountsTag = 0x00540017
	// EnergyWindowNameTag is tag for Energy Window Name
	EnergyWindowNameTag = 0x00540018
	// DetectorVectorTag is tag for Detector Vector
	DetectorVectorTag = 0x00540020
	// NumberOfDetectorsTag is tag for Number of Detectors
	NumberOfDetectorsTag = 0x00540021
	// DetectorInformationSequenceTag is tag for Detector Information Sequence
	DetectorInformationSequenceTag = 0x00540022
	// PhaseVectorTag is tag for Phase Vector
	PhaseVectorTag = 0x00540030
	// NumberOfPhasesTag is tag for Number of Phases
	NumberOfPhasesTag = 0x00540031
	// PhaseInformationSequenceTag is tag for Phase Information Sequence
	PhaseInformationSequenceTag = 0x00540032
	// NumberOfFramesInPhaseTag is tag for Number of Frames in Phase
	NumberOfFramesInPhaseTag = 0x00540033
	// PhaseDelayTag is tag for Phase Delay
	PhaseDelayTag = 0x00540036
	// PauseBetweenFramesTag is tag for Pause Between Frames
	PauseBetweenFramesTag = 0x00540038
	// PhaseDescriptionTag is tag for Phase Description
	PhaseDescriptionTag = 0x00540039
	// RotationVectorTag is tag for Rotation Vector
	RotationVectorTag = 0x00540050
	// NumberOfRotationsTag is tag for Number of Rotations
	NumberOfRotationsTag = 0x00540051
	// RotationInformationSequenceTag is tag for Rotation Information Sequence
	RotationInformationSequenceTag = 0x00540052
	// NumberOfFramesInRotationTag is tag for Number of Frames in Rotation
	NumberOfFramesInRotationTag = 0x00540053
	// RRIntervalVectorTag is tag for R-R Interval Vector
	RRIntervalVectorTag = 0x00540060
	// NumberOfRRIntervalsTag is tag for Number of R-R Intervals
	NumberOfRRIntervalsTag = 0x00540061
	// GatedInformationSequenceTag is tag for Gated Information Sequence
	GatedInformationSequenceTag = 0x00540062
	// DataInformationSequenceTag is tag for Data Information Sequence
	DataInformationSequenceTag = 0x00540063
	// TimeSlotVectorTag is tag for Time Slot Vector
	TimeSlotVectorTag = 0x00540070
	// NumberOfTimeSlotsTag is tag for Number of Time Slots
	NumberOfTimeSlotsTag = 0x00540071
	// TimeSlotInformationSequenceTag is tag for Time Slot Information Sequence
	TimeSlotInformationSequenceTag = 0x00540072
	// TimeSlotTimeTag is tag for Time Slot Time
	TimeSlotTimeTag = 0x00540073
	// SliceVectorTag is tag for Slice Vector
	SliceVectorTag = 0x00540080
	// NumberOfSlicesTag is tag for Number of Slices
	NumberOfSlicesTag = 0x00540081
	// AngularViewVectorTag is tag for Angular View Vector
	AngularViewVectorTag = 0x00540090
	// TimeSliceVectorTag is tag for Time Slice Vector
	TimeSliceVectorTag = 0x00540100
	// NumberOfTimeSlicesTag is tag for Number of Time Slices
	NumberOfTimeSlicesTag = 0x00540101
	// StartAngleTag is tag for Start Angle
	StartAngleTag = 0x00540200
	// TypeOfDetectorMotionTag is tag for Type of Detector Motion
	TypeOfDetectorMotionTag = 0x00540202
	// TriggerVectorTag is tag for Trigger Vector
	TriggerVectorTag = 0x00540210
	// NumberOfTriggersInPhaseTag is tag for Number of Triggers in Phase
	NumberOfTriggersInPhaseTag = 0x00540211
	// ViewCodeSequenceTag is tag for View Code Sequence
	ViewCodeSequenceTag = 0x00540220
	// ViewModifierCodeSequenceTag is tag for View Modifier Code Sequence
	ViewModifierCodeSequenceTag = 0x00540222
	// RadionuclideCodeSequenceTag is tag for Radionuclide Code Sequence
	RadionuclideCodeSequenceTag = 0x00540300
	// AdministrationRouteCodeSequenceTag is tag for Administration Route Code Sequence
	AdministrationRouteCodeSequenceTag = 0x00540302
	// RadiopharmaceuticalCodeSequenceTag is tag for Radiopharmaceutical Code Sequence
	RadiopharmaceuticalCodeSequenceTag = 0x00540304
	// CalibrationDataSequenceTag is tag for Calibration Data Sequence
	CalibrationDataSequenceTag = 0x00540306
	// EnergyWindowNumberTag is tag for Energy Window Number
	EnergyWindowNumberTag = 0x00540308
	// ImageIDTag is tag for Image ID
	ImageIDTag = 0x00540400
	// PatientOrientationCodeSequenceTag is tag for Patient Orientation Code Sequence
	PatientOrientationCodeSequenceTag = 0x00540410
	// PatientOrientationModifierCodeSequenceTag is tag for Patient Orientation Modifier Code Sequence
	PatientOrientationModifierCodeSequenceTag = 0x00540412
	// PatientGantryRelationshipCodeSequenceTag is tag for Patient Gantry Relationship Code Sequence
	PatientGantryRelationshipCodeSequenceTag = 0x00540414
	// SliceProgressionDirectionTag is tag for Slice Progression Direction
	SliceProgressionDirectionTag = 0x00540500
	// ScanProgressionDirectionTag is tag for Scan Progression Direction
	ScanProgressionDirectionTag = 0x00540501
	// SeriesTypeTag is tag for Series Type
	SeriesTypeTag = 0x00541000
	// UnitsTag is tag for Units
	UnitsTag = 0x00541001
	// CountsSourceTag is tag for Counts Source
	CountsSourceTag = 0x00541002
	// ReprojectionMethodTag is tag for Reprojection Method
	ReprojectionMethodTag = 0x00541004
	// SUVTypeTag is tag for SUV Type
	SUVTypeTag = 0x00541006
	// RandomsCorrectionMethodTag is tag for Randoms Correction Method
	RandomsCorrectionMethodTag = 0x00541100
	// AttenuationCorrectionMethodTag is tag for Attenuation Correction Method
	AttenuationCorrectionMethodTag = 0x00541101
	// DecayCorrectionTag is tag for Decay Correction
	DecayCorrectionTag = 0x00541102
	// ReconstructionMethodTag is tag for Reconstruction Method
	ReconstructionMethodTag = 0x00541103
	// DetectorLinesOfResponseUsedTag is tag for Detector Lines of Response Used
	DetectorLinesOfResponseUsedTag = 0x00541104
	// ScatterCorrectionMethodTag is tag for Scatter Correction Method
	ScatterCorrectionMethodTag = 0x00541105
	// AxialAcceptanceTag is tag for Axial Acceptance
	AxialAcceptanceTag = 0x00541200
	// AxialMashTag is tag for Axial Mash
	AxialMashTag = 0x00541201
	// TransverseMashTag is tag for Transverse Mash
	TransverseMashTag = 0x00541202
	// DetectorElementSizeTag is tag for Detector Element Size
	DetectorElementSizeTag = 0x00541203
	// CoincidenceWindowWidthTag is tag for Coincidence Window Width
	CoincidenceWindowWidthTag = 0x00541210
	// SecondaryCountsTypeTag is tag for Secondary Counts Type
	SecondaryCountsTypeTag = 0x00541220
	// FrameReferenceTimeTag is tag for Frame Reference Time
	FrameReferenceTimeTag = 0x00541300
	// PrimaryPromptsCountsAccumulatedTag is tag for Primary (Prompts) Counts Accumulated
	PrimaryPromptsCountsAccumulatedTag = 0x00541310
	// SecondaryCountsAccumulatedTag is tag for Secondary Counts Accumulated
	SecondaryCountsAccumulatedTag = 0x00541311
	// SliceSensitivityFactorTag is tag for Slice Sensitivity Factor
	SliceSensitivityFactorTag = 0x00541320
	// DecayFactorTag is tag for Decay Factor
	DecayFactorTag = 0x00541321
	// DoseCalibrationFactorTag is tag for Dose Calibration Factor
	DoseCalibrationFactorTag = 0x00541322
	// ScatterFractionFactorTag is tag for Scatter Fraction Factor
	ScatterFractionFactorTag = 0x00541323
	// DeadTimeFactorTag is tag for Dead Time Factor
	DeadTimeFactorTag = 0x00541324
	// ImageIndexTag is tag for Image Index
	ImageIndexTag = 0x00541330
	// CountsIncludedTag is tag for Counts Included
	CountsIncludedTag = 0x00541400
	// DeadTimeCorrectionFlagTag is tag for Dead Time Correction Flag
	DeadTimeCorrectionFlagTag = 0x00541401
	// HistogramSequenceTag is tag for Histogram Sequence
	HistogramSequenceTag = 0x00603000
	// HistogramNumberOfBinsTag is tag for Histogram Number of Bins
	HistogramNumberOfBinsTag = 0x00603002
	// HistogramFirstBinValueTag is tag for Histogram First Bin Value
	HistogramFirstBinValueTag = 0x00603004
	// HistogramLastBinValueTag is tag for Histogram Last Bin Value
	HistogramLastBinValueTag = 0x00603006
	// HistogramBinWidthTag is tag for Histogram Bin Width
	HistogramBinWidthTag = 0x00603008
	// HistogramExplanationTag is tag for Histogram Explanation
	HistogramExplanationTag = 0x00603010
	// HistogramDataTag is tag for Histogram Data
	HistogramDataTag = 0x00603020
	// SegmentationTypeTag is tag for Segmentation Type
	SegmentationTypeTag = 0x00620001
	// SegmentSequenceTag is tag for Segment Sequence
	SegmentSequenceTag = 0x00620002
	// SegmentedPropertyCategoryCodeSequenceTag is tag for Segmented Property Category Code Sequence
	SegmentedPropertyCategoryCodeSequenceTag = 0x00620003
	// SegmentNumberTag is tag for Segment Number
	SegmentNumberTag = 0x00620004
	// SegmentLabelTag is tag for Segment Label
	SegmentLabelTag = 0x00620005
	// SegmentDescriptionTag is tag for Segment Description
	SegmentDescriptionTag = 0x00620006
	// SegmentationAlgorithmIdentificationSequenceTag is tag for Segmentation Algorithm Identification Sequence
	SegmentationAlgorithmIdentificationSequenceTag = 0x00620007
	// SegmentAlgorithmTypeTag is tag for Segment Algorithm Type
	SegmentAlgorithmTypeTag = 0x00620008
	// SegmentAlgorithmNameTag is tag for Segment Algorithm Name
	SegmentAlgorithmNameTag = 0x00620009
	// SegmentIdentificationSequenceTag is tag for Segment Identification Sequence
	SegmentIdentificationSequenceTag = 0x0062000A
	// ReferencedSegmentNumberTag is tag for Referenced Segment Number
	ReferencedSegmentNumberTag = 0x0062000B
	// RecommendedDisplayGrayscaleValueTag is tag for Recommended Display Grayscale Value
	RecommendedDisplayGrayscaleValueTag = 0x0062000C
	// RecommendedDisplayCIELabValueTag is tag for Recommended Display CIELab Value
	RecommendedDisplayCIELabValueTag = 0x0062000D
	// MaximumFractionalValueTag is tag for Maximum Fractional Value
	MaximumFractionalValueTag = 0x0062000E
	// SegmentedPropertyTypeCodeSequenceTag is tag for Segmented Property Type Code Sequence
	SegmentedPropertyTypeCodeSequenceTag = 0x0062000F
	// SegmentationFractionalTypeTag is tag for Segmentation Fractional Type
	SegmentationFractionalTypeTag = 0x00620010
	// SegmentedPropertyTypeModifierCodeSequenceTag is tag for Segmented Property Type Modifier Code Sequence
	SegmentedPropertyTypeModifierCodeSequenceTag = 0x00620011
	// UsedSegmentsSequenceTag is tag for Used Segments Sequence
	UsedSegmentsSequenceTag = 0x00620012
	// SegmentsOverlapTag is tag for Segments Overlap
	SegmentsOverlapTag = 0x00620013
	// TrackingIDTag is tag for Tracking ID
	TrackingIDTag = 0x00620020
	// TrackingUIDTag is tag for Tracking UID
	TrackingUIDTag = 0x00620021
	// DeformableRegistrationSequenceTag is tag for Deformable Registration Sequence
	DeformableRegistrationSequenceTag = 0x00640002
	// SourceFrameOfReferenceUIDTag is tag for Source Frame of Reference UID
	SourceFrameOfReferenceUIDTag = 0x00640003
	// DeformableRegistrationGridSequenceTag is tag for Deformable Registration Grid Sequence
	DeformableRegistrationGridSequenceTag = 0x00640005
	// GridDimensionsTag is tag for Grid Dimensions
	GridDimensionsTag = 0x00640007
	// GridResolutionTag is tag for Grid Resolution
	GridResolutionTag = 0x00640008
	// VectorGridDataTag is tag for Vector Grid Data
	VectorGridDataTag = 0x00640009
	// PreDeformationMatrixRegistrationSequenceTag is tag for Pre Deformation Matrix Registration Sequence
	PreDeformationMatrixRegistrationSequenceTag = 0x0064000F
	// PostDeformationMatrixRegistrationSequenceTag is tag for Post Deformation Matrix Registration Sequence
	PostDeformationMatrixRegistrationSequenceTag = 0x00640010
	// NumberOfSurfacesTag is tag for Number of Surfaces
	NumberOfSurfacesTag = 0x00660001
	// SurfaceSequenceTag is tag for Surface Sequence
	SurfaceSequenceTag = 0x00660002
	// SurfaceNumberTag is tag for Surface Number
	SurfaceNumberTag = 0x00660003
	// SurfaceCommentsTag is tag for Surface Comments
	SurfaceCommentsTag = 0x00660004
	// SurfaceProcessingTag is tag for Surface Processing
	SurfaceProcessingTag = 0x00660009
	// SurfaceProcessingRatioTag is tag for Surface Processing Ratio
	SurfaceProcessingRatioTag = 0x0066000A
	// SurfaceProcessingDescriptionTag is tag for Surface Processing Description
	SurfaceProcessingDescriptionTag = 0x0066000B
	// RecommendedPresentationOpacityTag is tag for Recommended Presentation Opacity
	RecommendedPresentationOpacityTag = 0x0066000C
	// RecommendedPresentationTypeTag is tag for Recommended Presentation Type
	RecommendedPresentationTypeTag = 0x0066000D
	// FiniteVolumeTag is tag for Finite Volume
	FiniteVolumeTag = 0x0066000E
	// ManifoldTag is tag for Manifold
	ManifoldTag = 0x00660010
	// SurfacePointsSequenceTag is tag for Surface Points Sequence
	SurfacePointsSequenceTag = 0x00660011
	// SurfacePointsNormalsSequenceTag is tag for Surface Points Normals Sequence
	SurfacePointsNormalsSequenceTag = 0x00660012
	// SurfaceMeshPrimitivesSequenceTag is tag for Surface Mesh Primitives Sequence
	SurfaceMeshPrimitivesSequenceTag = 0x00660013
	// NumberOfSurfacePointsTag is tag for Number of Surface Points
	NumberOfSurfacePointsTag = 0x00660015
	// PointCoordinatesDataTag is tag for Point Coordinates Data
	PointCoordinatesDataTag = 0x00660016
	// PointPositionAccuracyTag is tag for Point Position Accuracy
	PointPositionAccuracyTag = 0x00660017
	// MeanPointDistanceTag is tag for Mean Point Distance
	MeanPointDistanceTag = 0x00660018
	// MaximumPointDistanceTag is tag for Maximum Point Distance
	MaximumPointDistanceTag = 0x00660019
	// PointsBoundingBoxCoordinatesTag is tag for Points Bounding Box Coordinates
	PointsBoundingBoxCoordinatesTag = 0x0066001A
	// AxisOfRotationTag is tag for Axis of Rotation
	AxisOfRotationTag = 0x0066001B
	// CenterOfRotationTag is tag for Center of Rotation
	CenterOfRotationTag = 0x0066001C
	// NumberOfVectorsTag is tag for Number of Vectors
	NumberOfVectorsTag = 0x0066001E
	// VectorDimensionalityTag is tag for Vector Dimensionality
	VectorDimensionalityTag = 0x0066001F
	// VectorAccuracyTag is tag for Vector Accuracy
	VectorAccuracyTag = 0x00660020
	// VectorCoordinateDataTag is tag for Vector Coordinate Data
	VectorCoordinateDataTag = 0x00660021
	// TrianglePointIndexListTag is tag for Triangle Point Index List
	TrianglePointIndexListTag = 0x00660023
	// EdgePointIndexListTag is tag for Edge Point Index List
	EdgePointIndexListTag = 0x00660024
	// VertexPointIndexListTag is tag for Vertex Point Index List
	VertexPointIndexListTag = 0x00660025
	// TriangleStripSequenceTag is tag for Triangle Strip Sequence
	TriangleStripSequenceTag = 0x00660026
	// TriangleFanSequenceTag is tag for Triangle Fan Sequence
	TriangleFanSequenceTag = 0x00660027
	// LineSequenceTag is tag for Line Sequence
	LineSequenceTag = 0x00660028
	// PrimitivePointIndexListTag is tag for Primitive Point Index List
	PrimitivePointIndexListTag = 0x00660029
	// SurfaceCountTag is tag for Surface Count
	SurfaceCountTag = 0x0066002A
	// ReferencedSurfaceSequenceTag is tag for Referenced Surface Sequence
	ReferencedSurfaceSequenceTag = 0x0066002B
	// ReferencedSurfaceNumberTag is tag for Referenced Surface Number
	ReferencedSurfaceNumberTag = 0x0066002C
	// SegmentSurfaceGenerationAlgorithmIdentificationSequenceTag is tag for Segment Surface Generation Algorithm Identification Sequence
	SegmentSurfaceGenerationAlgorithmIdentificationSequenceTag = 0x0066002D
	// SegmentSurfaceSourceInstanceSequenceTag is tag for Segment Surface Source Instance Sequence
	SegmentSurfaceSourceInstanceSequenceTag = 0x0066002E
	// AlgorithmFamilyCodeSequenceTag is tag for Algorithm Family Code Sequence
	AlgorithmFamilyCodeSequenceTag = 0x0066002F
	// AlgorithmNameCodeSequenceTag is tag for Algorithm Name Code Sequence
	AlgorithmNameCodeSequenceTag = 0x00660030
	// AlgorithmVersionTag is tag for Algorithm Version
	AlgorithmVersionTag = 0x00660031
	// AlgorithmParametersTag is tag for Algorithm Parameters
	AlgorithmParametersTag = 0x00660032
	// FacetSequenceTag is tag for Facet Sequence
	FacetSequenceTag = 0x00660034
	// SurfaceProcessingAlgorithmIdentificationSequenceTag is tag for Surface Processing Algorithm Identification Sequence
	SurfaceProcessingAlgorithmIdentificationSequenceTag = 0x00660035
	// AlgorithmNameTag is tag for Algorithm Name
	AlgorithmNameTag = 0x00660036
	// RecommendedPointRadiusTag is tag for Recommended Point Radius
	RecommendedPointRadiusTag = 0x00660037
	// RecommendedLineThicknessTag is tag for Recommended Line Thickness
	RecommendedLineThicknessTag = 0x00660038
	// LongPrimitivePointIndexListTag is tag for Long Primitive Point Index List
	LongPrimitivePointIndexListTag = 0x00660040
	// LongTrianglePointIndexListTag is tag for Long Triangle Point Index List
	LongTrianglePointIndexListTag = 0x00660041
	// LongEdgePointIndexListTag is tag for Long Edge Point Index List
	LongEdgePointIndexListTag = 0x00660042
	// LongVertexPointIndexListTag is tag for Long Vertex Point Index List
	LongVertexPointIndexListTag = 0x00660043
	// TrackSetSequenceTag is tag for Track Set Sequence
	TrackSetSequenceTag = 0x00660101
	// TrackSequenceTag is tag for Track Sequence
	TrackSequenceTag = 0x00660102
	// RecommendedDisplayCIELabValueListTag is tag for Recommended Display CIELab Value List
	RecommendedDisplayCIELabValueListTag = 0x00660103
	// TrackingAlgorithmIdentificationSequenceTag is tag for Tracking Algorithm Identification Sequence
	TrackingAlgorithmIdentificationSequenceTag = 0x00660104
	// TrackSetNumberTag is tag for Track Set Number
	TrackSetNumberTag = 0x00660105
	// TrackSetLabelTag is tag for Track Set Label
	TrackSetLabelTag = 0x00660106
	// TrackSetDescriptionTag is tag for Track Set Description
	TrackSetDescriptionTag = 0x00660107
	// TrackSetAnatomicalTypeCodeSequenceTag is tag for Track Set Anatomical Type Code Sequence
	TrackSetAnatomicalTypeCodeSequenceTag = 0x00660108
	// MeasurementsSequenceTag is tag for Measurements Sequence
	MeasurementsSequenceTag = 0x00660121
	// TrackSetStatisticsSequenceTag is tag for Track Set Statistics Sequence
	TrackSetStatisticsSequenceTag = 0x00660124
	// FloatingPointValuesTag is tag for Floating Point Values
	FloatingPointValuesTag = 0x00660125
	// TrackPointIndexListTag is tag for Track Point Index List
	TrackPointIndexListTag = 0x00660129
	// TrackStatisticsSequenceTag is tag for Track Statistics Sequence
	TrackStatisticsSequenceTag = 0x00660130
	// MeasurementValuesSequenceTag is tag for Measurement Values Sequence
	MeasurementValuesSequenceTag = 0x00660132
	// DiffusionAcquisitionCodeSequenceTag is tag for Diffusion Acquisition Code Sequence
	DiffusionAcquisitionCodeSequenceTag = 0x00660133
	// DiffusionModelCodeSequenceTag is tag for Diffusion Model Code Sequence
	DiffusionModelCodeSequenceTag = 0x00660134
	// ImplantSizeTag is tag for Implant Size
	ImplantSizeTag = 0x00686210
	// ImplantTemplateVersionTag is tag for Implant Template Version
	ImplantTemplateVersionTag = 0x00686221
	// ReplacedImplantTemplateSequenceTag is tag for Replaced Implant Template Sequence
	ReplacedImplantTemplateSequenceTag = 0x00686222
	// ImplantTypeTag is tag for Implant Type
	ImplantTypeTag = 0x00686223
	// DerivationImplantTemplateSequenceTag is tag for Derivation Implant Template Sequence
	DerivationImplantTemplateSequenceTag = 0x00686224
	// OriginalImplantTemplateSequenceTag is tag for Original Implant Template Sequence
	OriginalImplantTemplateSequenceTag = 0x00686225
	// EffectiveDateTimeTag is tag for Effective DateTime
	EffectiveDateTimeTag = 0x00686226
	// ImplantTargetAnatomySequenceTag is tag for Implant Target Anatomy Sequence
	ImplantTargetAnatomySequenceTag = 0x00686230
	// InformationFromManufacturerSequenceTag is tag for Information From Manufacturer Sequence
	InformationFromManufacturerSequenceTag = 0x00686260
	// NotificationFromManufacturerSequenceTag is tag for Notification From Manufacturer Sequence
	NotificationFromManufacturerSequenceTag = 0x00686265
	// InformationIssueDateTimeTag is tag for Information Issue DateTime
	InformationIssueDateTimeTag = 0x00686270
	// InformationSummaryTag is tag for Information Summary
	InformationSummaryTag = 0x00686280
	// ImplantRegulatoryDisapprovalCodeSequenceTag is tag for Implant Regulatory Disapproval Code Sequence
	ImplantRegulatoryDisapprovalCodeSequenceTag = 0x006862A0
	// OverallTemplateSpatialToleranceTag is tag for Overall Template Spatial Tolerance
	OverallTemplateSpatialToleranceTag = 0x006862A5
	// HPGLDocumentSequenceTag is tag for HPGL Document Sequence
	HPGLDocumentSequenceTag = 0x006862C0
	// HPGLDocumentIDTag is tag for HPGL Document ID
	HPGLDocumentIDTag = 0x006862D0
	// HPGLDocumentLabelTag is tag for HPGL Document Label
	HPGLDocumentLabelTag = 0x006862D5
	// ViewOrientationCodeSequenceTag is tag for View Orientation Code Sequence
	ViewOrientationCodeSequenceTag = 0x006862E0
	// ViewOrientationModifierCodeSequenceTag is tag for View Orientation Modifier Code Sequence
	ViewOrientationModifierCodeSequenceTag = 0x006862F0
	// HPGLDocumentScalingTag is tag for HPGL Document Scaling
	HPGLDocumentScalingTag = 0x006862F2
	// HPGLDocumentTag is tag for HPGL Document
	HPGLDocumentTag = 0x00686300
	// HPGLContourPenNumberTag is tag for HPGL Contour Pen Number
	HPGLContourPenNumberTag = 0x00686310
	// HPGLPenSequenceTag is tag for HPGL Pen Sequence
	HPGLPenSequenceTag = 0x00686320
	// HPGLPenNumberTag is tag for HPGL Pen Number
	HPGLPenNumberTag = 0x00686330
	// HPGLPenLabelTag is tag for HPGL Pen Label
	HPGLPenLabelTag = 0x00686340
	// HPGLPenDescriptionTag is tag for HPGL Pen Description
	HPGLPenDescriptionTag = 0x00686345
	// RecommendedRotationPointTag is tag for Recommended Rotation Point
	RecommendedRotationPointTag = 0x00686346
	// BoundingRectangleTag is tag for Bounding Rectangle
	BoundingRectangleTag = 0x00686347
	// ImplantTemplate3DModelSurfaceNumberTag is tag for Implant Template 3D Model Surface Number
	ImplantTemplate3DModelSurfaceNumberTag = 0x00686350
	// SurfaceModelDescriptionSequenceTag is tag for Surface Model Description Sequence
	SurfaceModelDescriptionSequenceTag = 0x00686360
	// SurfaceModelLabelTag is tag for Surface Model Label
	SurfaceModelLabelTag = 0x00686380
	// SurfaceModelScalingFactorTag is tag for Surface Model Scaling Factor
	SurfaceModelScalingFactorTag = 0x00686390
	// MaterialsCodeSequenceTag is tag for Materials Code Sequence
	MaterialsCodeSequenceTag = 0x006863A0
	// CoatingMaterialsCodeSequenceTag is tag for Coating Materials Code Sequence
	CoatingMaterialsCodeSequenceTag = 0x006863A4
	// ImplantTypeCodeSequenceTag is tag for Implant Type Code Sequence
	ImplantTypeCodeSequenceTag = 0x006863A8
	// FixationMethodCodeSequenceTag is tag for Fixation Method Code Sequence
	FixationMethodCodeSequenceTag = 0x006863AC
	// MatingFeatureSetsSequenceTag is tag for Mating Feature Sets Sequence
	MatingFeatureSetsSequenceTag = 0x006863B0
	// MatingFeatureSetIDTag is tag for Mating Feature Set ID
	MatingFeatureSetIDTag = 0x006863C0
	// MatingFeatureSetLabelTag is tag for Mating Feature Set Label
	MatingFeatureSetLabelTag = 0x006863D0
	// MatingFeatureSequenceTag is tag for Mating Feature Sequence
	MatingFeatureSequenceTag = 0x006863E0
	// MatingFeatureIDTag is tag for Mating Feature ID
	MatingFeatureIDTag = 0x006863F0
	// MatingFeatureDegreeOfFreedomSequenceTag is tag for Mating Feature Degree of Freedom Sequence
	MatingFeatureDegreeOfFreedomSequenceTag = 0x00686400
	// DegreeOfFreedomIDTag is tag for Degree of Freedom ID
	DegreeOfFreedomIDTag = 0x00686410
	// DegreeOfFreedomTypeTag is tag for Degree of Freedom Type
	DegreeOfFreedomTypeTag = 0x00686420
	// TwoDMatingFeatureCoordinatesSequenceTag is tag for 2D Mating Feature Coordinates Sequence
	TwoDMatingFeatureCoordinatesSequenceTag = 0x00686430
	// ReferencedHPGLDocumentIDTag is tag for Referenced HPGL Document ID
	ReferencedHPGLDocumentIDTag = 0x00686440
	// TwoDMatingPointTag is tag for 2D Mating Point
	TwoDMatingPointTag = 0x00686450
	// TwoDMatingAxesTag is tag for 2D Mating Axes
	TwoDMatingAxesTag = 0x00686460
	// TwoDDegreeOfFreedomSequenceTag is tag for 2D Degree of Freedom Sequence
	TwoDDegreeOfFreedomSequenceTag = 0x00686470
	// ThreeDDegreeOfFreedomAxisTag is tag for 3D Degree of Freedom Axis
	ThreeDDegreeOfFreedomAxisTag = 0x00686490
	// RangeOfFreedomTag is tag for Range of Freedom
	RangeOfFreedomTag = 0x006864A0
	// ThreeDMatingPointTag is tag for 3D Mating Point
	ThreeDMatingPointTag = 0x006864C0
	// ThreeDMatingAxesTag is tag for 3D Mating Axes
	ThreeDMatingAxesTag = 0x006864D0
	// TwoDDegreeOfFreedomAxisTag is tag for 2D Degree of Freedom Axis
	TwoDDegreeOfFreedomAxisTag = 0x006864F0
	// PlanningLandmarkPointSequenceTag is tag for Planning Landmark Point Sequence
	PlanningLandmarkPointSequenceTag = 0x00686500
	// PlanningLandmarkLineSequenceTag is tag for Planning Landmark Line Sequence
	PlanningLandmarkLineSequenceTag = 0x00686510
	// PlanningLandmarkPlaneSequenceTag is tag for Planning Landmark Plane Sequence
	PlanningLandmarkPlaneSequenceTag = 0x00686520
	// PlanningLandmarkIDTag is tag for Planning Landmark ID
	PlanningLandmarkIDTag = 0x00686530
	// PlanningLandmarkDescriptionTag is tag for Planning Landmark Description
	PlanningLandmarkDescriptionTag = 0x00686540
	// PlanningLandmarkIdentificationCodeSequenceTag is tag for Planning Landmark Identification Code Sequence
	PlanningLandmarkIdentificationCodeSequenceTag = 0x00686545
	// TwoDPointCoordinatesSequenceTag is tag for 2D Point Coordinates Sequence
	TwoDPointCoordinatesSequenceTag = 0x00686550
	// TwoDPointCoordinatesTag is tag for 2D Point Coordinates
	TwoDPointCoordinatesTag = 0x00686560
	// ThreeDPointCoordinatesTag is tag for 3D Point Coordinates
	ThreeDPointCoordinatesTag = 0x00686590
	// TwoDLineCoordinatesSequenceTag is tag for 2D Line Coordinates Sequence
	TwoDLineCoordinatesSequenceTag = 0x006865A0
	// TwoDLineCoordinatesTag is tag for 2D Line Coordinates
	TwoDLineCoordinatesTag = 0x006865B0
	// ThreeDLineCoordinatesTag is tag for 3D Line Coordinates
	ThreeDLineCoordinatesTag = 0x006865D0
	// TwoDPlaneCoordinatesSequenceTag is tag for 2D Plane Coordinates Sequence
	TwoDPlaneCoordinatesSequenceTag = 0x006865E0
	// TwoDPlaneIntersectionTag is tag for 2D Plane Intersection
	TwoDPlaneIntersectionTag = 0x006865F0
	// ThreeDPlaneOriginTag is tag for 3D Plane Origin
	ThreeDPlaneOriginTag = 0x00686610
	// ThreeDPlaneNormalTag is tag for 3D Plane Normal
	ThreeDPlaneNormalTag = 0x00686620
	// ModelModificationTag is tag for Model Modification
	ModelModificationTag = 0x00687001
	// ModelMirroringTag is tag for Model Mirroring
	ModelMirroringTag = 0x00687002
	// ModelUsageCodeSequenceTag is tag for Model Usage Code Sequence
	ModelUsageCodeSequenceTag = 0x00687003
	// GraphicAnnotationSequenceTag is tag for Graphic Annotation Sequence
	GraphicAnnotationSequenceTag = 0x00700001
	// GraphicLayerTag is tag for Graphic Layer
	GraphicLayerTag = 0x00700002
	// BoundingBoxAnnotationUnitsTag is tag for Bounding Box Annotation Units
	BoundingBoxAnnotationUnitsTag = 0x00700003
	// AnchorPointAnnotationUnitsTag is tag for Anchor Point Annotation Units
	AnchorPointAnnotationUnitsTag = 0x00700004
	// GraphicAnnotationUnitsTag is tag for Graphic Annotation Units
	GraphicAnnotationUnitsTag = 0x00700005
	// UnformattedTextValueTag is tag for Unformatted Text Value
	UnformattedTextValueTag = 0x00700006
	// TextObjectSequenceTag is tag for Text Object Sequence
	TextObjectSequenceTag = 0x00700008
	// GraphicObjectSequenceTag is tag for Graphic Object Sequence
	GraphicObjectSequenceTag = 0x00700009
	// BoundingBoxTopLeftHandCornerTag is tag for Bounding Box Top Left Hand Corner
	BoundingBoxTopLeftHandCornerTag = 0x00700010
	// BoundingBoxBottomRightHandCornerTag is tag for Bounding Box Bottom Right Hand Corner
	BoundingBoxBottomRightHandCornerTag = 0x00700011
	// BoundingBoxTextHorizontalJustificationTag is tag for Bounding Box Text Horizontal Justification
	BoundingBoxTextHorizontalJustificationTag = 0x00700012
	// AnchorPointTag is tag for Anchor Point
	AnchorPointTag = 0x00700014
	// AnchorPointVisibilityTag is tag for Anchor Point Visibility
	AnchorPointVisibilityTag = 0x00700015
	// GraphicDimensionsTag is tag for Graphic Dimensions
	GraphicDimensionsTag = 0x00700020
	// NumberOfGraphicPointsTag is tag for Number of Graphic Points
	NumberOfGraphicPointsTag = 0x00700021
	// GraphicDataTag is tag for Graphic Data
	GraphicDataTag = 0x00700022
	// GraphicTypeTag is tag for Graphic Type
	GraphicTypeTag = 0x00700023
	// GraphicFilledTag is tag for Graphic Filled
	GraphicFilledTag = 0x00700024
	// ImageRotationRetiredTag is tag for Image Rotation (Retired)
	ImageRotationRetiredTag = 0x00700040
	// ImageHorizontalFlipTag is tag for Image Horizontal Flip
	ImageHorizontalFlipTag = 0x00700041
	// ImageRotationTag is tag for Image Rotation
	ImageRotationTag = 0x00700042
	// DisplayedAreaTopLeftHandCornerTrialTag is tag for Displayed Area Top Left Hand Corner (Trial)
	DisplayedAreaTopLeftHandCornerTrialTag = 0x00700050
	// DisplayedAreaBottomRightHandCornerTrialTag is tag for Displayed Area Bottom Right Hand Corner (Trial)
	DisplayedAreaBottomRightHandCornerTrialTag = 0x00700051
	// DisplayedAreaTopLeftHandCornerTag is tag for Displayed Area Top Left Hand Corner
	DisplayedAreaTopLeftHandCornerTag = 0x00700052
	// DisplayedAreaBottomRightHandCornerTag is tag for Displayed Area Bottom Right Hand Corner
	DisplayedAreaBottomRightHandCornerTag = 0x00700053
	// DisplayedAreaSelectionSequenceTag is tag for Displayed Area Selection Sequence
	DisplayedAreaSelectionSequenceTag = 0x0070005A
	// GraphicLayerSequenceTag is tag for Graphic Layer Sequence
	GraphicLayerSequenceTag = 0x00700060
	// GraphicLayerOrderTag is tag for Graphic Layer Order
	GraphicLayerOrderTag = 0x00700062
	// GraphicLayerRecommendedDisplayGrayscaleValueTag is tag for Graphic Layer Recommended Display Grayscale Value
	GraphicLayerRecommendedDisplayGrayscaleValueTag = 0x00700066
	// GraphicLayerRecommendedDisplayRGBValueTag is tag for Graphic Layer Recommended Display RGB Value
	GraphicLayerRecommendedDisplayRGBValueTag = 0x00700067
	// GraphicLayerDescriptionTag is tag for Graphic Layer Description
	GraphicLayerDescriptionTag = 0x00700068
	// ContentLabelTag is tag for Content Label
	ContentLabelTag = 0x00700080
	// ContentDescriptionTag is tag for Content Description
	ContentDescriptionTag = 0x00700081
	// PresentationCreationDateTag is tag for Presentation Creation Date
	PresentationCreationDateTag = 0x00700082
	// PresentationCreationTimeTag is tag for Presentation Creation Time
	PresentationCreationTimeTag = 0x00700083
	// ContentCreatorNameTag is tag for Content Creator's Name
	ContentCreatorNameTag = 0x00700084
	// ContentCreatorIdentificationCodeSequenceTag is tag for Content Creator's Identification Code Sequence
	ContentCreatorIdentificationCodeSequenceTag = 0x00700086
	// AlternateContentDescriptionSequenceTag is tag for Alternate Content Description Sequence
	AlternateContentDescriptionSequenceTag = 0x00700087
	// PresentationSizeModeTag is tag for Presentation Size Mode
	PresentationSizeModeTag = 0x00700100
	// PresentationPixelSpacingTag is tag for Presentation Pixel Spacing
	PresentationPixelSpacingTag = 0x00700101
	// PresentationPixelAspectRatioTag is tag for Presentation Pixel Aspect Ratio
	PresentationPixelAspectRatioTag = 0x00700102
	// PresentationPixelMagnificationRatioTag is tag for Presentation Pixel Magnification Ratio
	PresentationPixelMagnificationRatioTag = 0x00700103
	// GraphicGroupLabelTag is tag for Graphic Group Label
	GraphicGroupLabelTag = 0x00700207
	// GraphicGroupDescriptionTag is tag for Graphic Group Description
	GraphicGroupDescriptionTag = 0x00700208
	// CompoundGraphicSequenceTag is tag for Compound Graphic Sequence
	CompoundGraphicSequenceTag = 0x00700209
	// CompoundGraphicInstanceIDTag is tag for Compound Graphic Instance ID
	CompoundGraphicInstanceIDTag = 0x00700226
	// FontNameTag is tag for Font Name
	FontNameTag = 0x00700227
	// FontNameTypeTag is tag for Font Name Type
	FontNameTypeTag = 0x00700228
	// CSSFontNameTag is tag for CSS Font Name
	CSSFontNameTag = 0x00700229
	// RotationAngleTag is tag for Rotation Angle
	RotationAngleTag = 0x00700230
	// TextStyleSequenceTag is tag for Text Style Sequence
	TextStyleSequenceTag = 0x00700231
	// LineStyleSequenceTag is tag for Line Style Sequence
	LineStyleSequenceTag = 0x00700232
	// FillStyleSequenceTag is tag for Fill Style Sequence
	FillStyleSequenceTag = 0x00700233
	// GraphicGroupSequenceTag is tag for Graphic Group Sequence
	GraphicGroupSequenceTag = 0x00700234
	// TextColorCIELabValueTag is tag for Text Color CIELab Value
	TextColorCIELabValueTag = 0x00700241
	// HorizontalAlignmentTag is tag for Horizontal Alignment
	HorizontalAlignmentTag = 0x00700242
	// VerticalAlignmentTag is tag for Vertical Alignment
	VerticalAlignmentTag = 0x00700243
	// ShadowStyleTag is tag for Shadow Style
	ShadowStyleTag = 0x00700244
	// ShadowOffsetXTag is tag for Shadow Offset X
	ShadowOffsetXTag = 0x00700245
	// ShadowOffsetYTag is tag for Shadow Offset Y
	ShadowOffsetYTag = 0x00700246
	// ShadowColorCIELabValueTag is tag for Shadow Color CIELab Value
	ShadowColorCIELabValueTag = 0x00700247
	// UnderlinedTag is tag for Underlined
	UnderlinedTag = 0x00700248
	// BoldTag is tag for Bold
	BoldTag = 0x00700249
	// ItalicTag is tag for Italic
	ItalicTag = 0x00700250
	// PatternOnColorCIELabValueTag is tag for Pattern On Color CIELab Value
	PatternOnColorCIELabValueTag = 0x00700251
	// PatternOffColorCIELabValueTag is tag for Pattern Off Color CIELab Value
	PatternOffColorCIELabValueTag = 0x00700252
	// LineThicknessTag is tag for Line Thickness
	LineThicknessTag = 0x00700253
	// LineDashingStyleTag is tag for Line Dashing Style
	LineDashingStyleTag = 0x00700254
	// LinePatternTag is tag for Line Pattern
	LinePatternTag = 0x00700255
	// FillPatternTag is tag for Fill Pattern
	FillPatternTag = 0x00700256
	// FillModeTag is tag for Fill Mode
	FillModeTag = 0x00700257
	// ShadowOpacityTag is tag for Shadow Opacity
	ShadowOpacityTag = 0x00700258
	// GapLengthTag is tag for Gap Length
	GapLengthTag = 0x00700261
	// DiameterOfVisibilityTag is tag for Diameter of Visibility
	DiameterOfVisibilityTag = 0x00700262
	// RotationPointTag is tag for Rotation Point
	RotationPointTag = 0x00700273
	// TickAlignmentTag is tag for Tick Alignment
	TickAlignmentTag = 0x00700274
	// ShowTickLabelTag is tag for Show Tick Label
	ShowTickLabelTag = 0x00700278
	// TickLabelAlignmentTag is tag for Tick Label Alignment
	TickLabelAlignmentTag = 0x00700279
	// CompoundGraphicUnitsTag is tag for Compound Graphic Units
	CompoundGraphicUnitsTag = 0x00700282
	// PatternOnOpacityTag is tag for Pattern On Opacity
	PatternOnOpacityTag = 0x00700284
	// PatternOffOpacityTag is tag for Pattern Off Opacity
	PatternOffOpacityTag = 0x00700285
	// MajorTicksSequenceTag is tag for Major Ticks Sequence
	MajorTicksSequenceTag = 0x00700287
	// TickPositionTag is tag for Tick Position
	TickPositionTag = 0x00700288
	// TickLabelTag is tag for Tick Label
	TickLabelTag = 0x00700289
	// CompoundGraphicTypeTag is tag for Compound Graphic Type
	CompoundGraphicTypeTag = 0x00700294
	// GraphicGroupIDTag is tag for Graphic Group ID
	GraphicGroupIDTag = 0x00700295
	// ShapeTypeTag is tag for Shape Type
	ShapeTypeTag = 0x00700306
	// RegistrationSequenceTag is tag for Registration Sequence
	RegistrationSequenceTag = 0x00700308
	// MatrixRegistrationSequenceTag is tag for Matrix Registration Sequence
	MatrixRegistrationSequenceTag = 0x00700309
	// MatrixSequenceTag is tag for Matrix Sequence
	MatrixSequenceTag = 0x0070030A
	// FrameOfReferenceToDisplayedCoordinateSystemTransformationMatrixTag is tag for Frame of Reference to Displayed Coordinate System Transformation Matrix
	FrameOfReferenceToDisplayedCoordinateSystemTransformationMatrixTag = 0x0070030B
	// FrameOfReferenceTransformationMatrixTypeTag is tag for Frame of Reference Transformation Matrix Type
	FrameOfReferenceTransformationMatrixTypeTag = 0x0070030C
	// RegistrationTypeCodeSequenceTag is tag for Registration Type Code Sequence
	RegistrationTypeCodeSequenceTag = 0x0070030D
	// FiducialDescriptionTag is tag for Fiducial Description
	FiducialDescriptionTag = 0x0070030F
	// FiducialIdentifierTag is tag for Fiducial Identifier
	FiducialIdentifierTag = 0x00700310
	// FiducialIdentifierCodeSequenceTag is tag for Fiducial Identifier Code Sequence
	FiducialIdentifierCodeSequenceTag = 0x00700311
	// ContourUncertaintyRadiusTag is tag for Contour Uncertainty Radius
	ContourUncertaintyRadiusTag = 0x00700312
	// UsedFiducialsSequenceTag is tag for Used Fiducials Sequence
	UsedFiducialsSequenceTag = 0x00700314
	// GraphicCoordinatesDataSequenceTag is tag for Graphic Coordinates Data Sequence
	GraphicCoordinatesDataSequenceTag = 0x00700318
	// FiducialUIDTag is tag for Fiducial UID
	FiducialUIDTag = 0x0070031A
	// ReferencedFiducialUIDTag is tag for Referenced Fiducial UID
	ReferencedFiducialUIDTag = 0x0070031B
	// FiducialSetSequenceTag is tag for Fiducial Set Sequence
	FiducialSetSequenceTag = 0x0070031C
	// FiducialSequenceTag is tag for Fiducial Sequence
	FiducialSequenceTag = 0x0070031E
	// FiducialsPropertyCategoryCodeSequenceTag is tag for Fiducials Property Category Code Sequence
	FiducialsPropertyCategoryCodeSequenceTag = 0x0070031F
	// GraphicLayerRecommendedDisplayCIELabValueTag is tag for Graphic Layer Recommended Display CIELab Value
	GraphicLayerRecommendedDisplayCIELabValueTag = 0x00700401
	// BlendingSequenceTag is tag for Blending Sequence
	BlendingSequenceTag = 0x00700402
	// RelativeOpacityTag is tag for Relative Opacity
	RelativeOpacityTag = 0x00700403
	// ReferencedSpatialRegistrationSequenceTag is tag for Referenced Spatial Registration Sequence
	ReferencedSpatialRegistrationSequenceTag = 0x00700404
	// BlendingPositionTag is tag for Blending Position
	BlendingPositionTag = 0x00700405
	// PresentationDisplayCollectionUIDTag is tag for Presentation Display Collection UID
	PresentationDisplayCollectionUIDTag = 0x00701101
	// PresentationSequenceCollectionUIDTag is tag for Presentation Sequence Collection UID
	PresentationSequenceCollectionUIDTag = 0x00701102
	// PresentationSequencePositionIndexTag is tag for Presentation Sequence Position Index
	PresentationSequencePositionIndexTag = 0x00701103
	// RenderedImageReferenceSequenceTag is tag for Rendered Image Reference Sequence
	RenderedImageReferenceSequenceTag = 0x00701104
	// VolumetricPresentationStateInputSequenceTag is tag for Volumetric Presentation State Input Sequence
	VolumetricPresentationStateInputSequenceTag = 0x00701201
	// PresentationInputTypeTag is tag for Presentation Input Type
	PresentationInputTypeTag = 0x00701202
	// InputSequencePositionIndexTag is tag for Input Sequence Position Index
	InputSequencePositionIndexTag = 0x00701203
	// CropTag is tag for Crop
	CropTag = 0x00701204
	// CroppingSpecificationIndexTag is tag for Cropping Specification Index
	CroppingSpecificationIndexTag = 0x00701205
	// CompositingMethodTag is tag for Compositing Method
	CompositingMethodTag = 0x00701206
	// VolumetricPresentationInputNumberTag is tag for Volumetric Presentation Input Number
	VolumetricPresentationInputNumberTag = 0x00701207
	// ImageVolumeGeometryTag is tag for Image Volume Geometry
	ImageVolumeGeometryTag = 0x00701208
	// VolumetricPresentationInputSetUIDTag is tag for Volumetric Presentation Input Set UID
	VolumetricPresentationInputSetUIDTag = 0x00701209
	// VolumetricPresentationInputSetSequenceTag is tag for Volumetric Presentation Input Set Sequence
	VolumetricPresentationInputSetSequenceTag = 0x0070120A
	// GlobalCropTag is tag for Global Crop
	GlobalCropTag = 0x0070120B
	// GlobalCroppingSpecificationIndexTag is tag for Global Cropping Specification Index
	GlobalCroppingSpecificationIndexTag = 0x0070120C
	// RenderingMethodTag is tag for Rendering Method
	RenderingMethodTag = 0x0070120D
	// VolumeCroppingSequenceTag is tag for Volume Cropping Sequence
	VolumeCroppingSequenceTag = 0x00701301
	// VolumeCroppingMethodTag is tag for Volume Cropping Method
	VolumeCroppingMethodTag = 0x00701302
	// BoundingBoxCropTag is tag for Bounding Box Crop
	BoundingBoxCropTag = 0x00701303
	// ObliqueCroppingPlaneSequenceTag is tag for Oblique Cropping Plane Sequence
	ObliqueCroppingPlaneSequenceTag = 0x00701304
	// PlaneTag is tag for Plane
	PlaneTag = 0x00701305
	// PlaneNormalTag is tag for Plane Normal
	PlaneNormalTag = 0x00701306
	// CroppingSpecificationNumberTag is tag for Cropping Specification Number
	CroppingSpecificationNumberTag = 0x00701309
	// MultiPlanarReconstructionStyleTag is tag for Multi-Planar Reconstruction Style
	MultiPlanarReconstructionStyleTag = 0x00701501
	// MPRThicknessTypeTag is tag for MPR Thickness Type
	MPRThicknessTypeTag = 0x00701502
	// MPRSlabThicknessTag is tag for MPR Slab Thickness
	MPRSlabThicknessTag = 0x00701503
	// MPRTopLeftHandCornerTag is tag for MPR Top Left Hand Corner
	MPRTopLeftHandCornerTag = 0x00701505
	// MPRViewWidthDirectionTag is tag for MPR View Width Direction
	MPRViewWidthDirectionTag = 0x00701507
	// MPRViewWidthTag is tag for MPR View Width
	MPRViewWidthTag = 0x00701508
	// NumberOfVolumetricCurvePointsTag is tag for Number of Volumetric Curve Points
	NumberOfVolumetricCurvePointsTag = 0x0070150C
	// VolumetricCurvePointsTag is tag for Volumetric Curve Points
	VolumetricCurvePointsTag = 0x0070150D
	// MPRViewHeightDirectionTag is tag for MPR View Height Direction
	MPRViewHeightDirectionTag = 0x00701511
	// MPRViewHeightTag is tag for MPR View Height
	MPRViewHeightTag = 0x00701512
	// RenderProjectionTag is tag for Render Projection
	RenderProjectionTag = 0x00701602
	// ViewpointPositionTag is tag for Viewpoint Position
	ViewpointPositionTag = 0x00701603
	// ViewpointLookAtPointTag is tag for Viewpoint LookAt Point
	ViewpointLookAtPointTag = 0x00701604
	// ViewpointUpDirectionTag is tag for Viewpoint Up Direction
	ViewpointUpDirectionTag = 0x00701605
	// RenderFieldOfViewTag is tag for Render Field of View
	RenderFieldOfViewTag = 0x00701606
	// SamplingStepSizeTag is tag for Sampling Step Size
	SamplingStepSizeTag = 0x00701607
	// ShadingStyleTag is tag for Shading Style
	ShadingStyleTag = 0x00701701
	// AmbientReflectionIntensityTag is tag for Ambient Reflection Intensity
	AmbientReflectionIntensityTag = 0x00701702
	// LightDirectionTag is tag for Light Direction
	LightDirectionTag = 0x00701703
	// DiffuseReflectionIntensityTag is tag for Diffuse Reflection Intensity
	DiffuseReflectionIntensityTag = 0x00701704
	// SpecularReflectionIntensityTag is tag for Specular Reflection Intensity
	SpecularReflectionIntensityTag = 0x00701705
	// ShininessTag is tag for Shininess
	ShininessTag = 0x00701706
	// PresentationStateClassificationComponentSequenceTag is tag for Presentation State Classification Component Sequence
	PresentationStateClassificationComponentSequenceTag = 0x00701801
	// ComponentTypeTag is tag for Component Type
	ComponentTypeTag = 0x00701802
	// ComponentInputSequenceTag is tag for Component Input Sequence
	ComponentInputSequenceTag = 0x00701803
	// VolumetricPresentationInputIndexTag is tag for Volumetric Presentation Input Index
	VolumetricPresentationInputIndexTag = 0x00701804
	// PresentationStateCompositorComponentSequenceTag is tag for Presentation State Compositor Component Sequence
	PresentationStateCompositorComponentSequenceTag = 0x00701805
	// WeightingTransferFunctionSequenceTag is tag for Weighting Transfer Function Sequence
	WeightingTransferFunctionSequenceTag = 0x00701806
	// WeightingLookupTableDescriptorTag is tag for Weighting Lookup Table Descriptor
	WeightingLookupTableDescriptorTag = 0x00701807
	// WeightingLookupTableDataTag is tag for Weighting Lookup Table Data
	WeightingLookupTableDataTag = 0x00701808
	// VolumetricAnnotationSequenceTag is tag for Volumetric Annotation Sequence
	VolumetricAnnotationSequenceTag = 0x00701901
	// ReferencedStructuredContextSequenceTag is tag for Referenced Structured Context Sequence
	ReferencedStructuredContextSequenceTag = 0x00701903
	// ReferencedContentItemTag is tag for Referenced Content Item
	ReferencedContentItemTag = 0x00701904
	// VolumetricPresentationInputAnnotationSequenceTag is tag for Volumetric Presentation Input Annotation Sequence
	VolumetricPresentationInputAnnotationSequenceTag = 0x00701905
	// AnnotationClippingTag is tag for Annotation Clipping
	AnnotationClippingTag = 0x00701907
	// PresentationAnimationStyleTag is tag for Presentation Animation Style
	PresentationAnimationStyleTag = 0x00701A01
	// RecommendedAnimationRateTag is tag for Recommended Animation Rate
	RecommendedAnimationRateTag = 0x00701A03
	// AnimationCurveSequenceTag is tag for Animation Curve Sequence
	AnimationCurveSequenceTag = 0x00701A04
	// AnimationStepSizeTag is tag for Animation Step Size
	AnimationStepSizeTag = 0x00701A05
	// SwivelRangeTag is tag for Swivel Range
	SwivelRangeTag = 0x00701A06
	// VolumetricCurveUpDirectionsTag is tag for Volumetric Curve Up Directions
	VolumetricCurveUpDirectionsTag = 0x00701A07
	// VolumeStreamSequenceTag is tag for Volume Stream Sequence
	VolumeStreamSequenceTag = 0x00701A08
	// RGBATransferFunctionDescriptionTag is tag for RGBA Transfer Function Description
	RGBATransferFunctionDescriptionTag = 0x00701A09
	// AdvancedBlendingSequenceTag is tag for Advanced Blending Sequence
	AdvancedBlendingSequenceTag = 0x00701B01
	// BlendingInputNumberTag is tag for Blending Input Number
	BlendingInputNumberTag = 0x00701B02
	// BlendingDisplayInputSequenceTag is tag for Blending Display Input Sequence
	BlendingDisplayInputSequenceTag = 0x00701B03
	// BlendingDisplaySequenceTag is tag for Blending Display Sequence
	BlendingDisplaySequenceTag = 0x00701B04
	// BlendingModeTag is tag for Blending Mode
	BlendingModeTag = 0x00701B06
	// TimeSeriesBlendingTag is tag for Time Series Blending
	TimeSeriesBlendingTag = 0x00701B07
	// GeometryForDisplayTag is tag for Geometry for Display
	GeometryForDisplayTag = 0x00701B08
	// ThresholdSequenceTag is tag for Threshold Sequence
	ThresholdSequenceTag = 0x00701B11
	// ThresholdValueSequenceTag is tag for Threshold Value Sequence
	ThresholdValueSequenceTag = 0x00701B12
	// ThresholdTypeTag is tag for Threshold Type
	ThresholdTypeTag = 0x00701B13
	// ThresholdValueTag is tag for Threshold Value
	ThresholdValueTag = 0x00701B14
	// HangingProtocolNameTag is tag for Hanging Protocol Name
	HangingProtocolNameTag = 0x00720002
	// HangingProtocolDescriptionTag is tag for Hanging Protocol Description
	HangingProtocolDescriptionTag = 0x00720004
	// HangingProtocolLevelTag is tag for Hanging Protocol Level
	HangingProtocolLevelTag = 0x00720006
	// HangingProtocolCreatorTag is tag for Hanging Protocol Creator
	HangingProtocolCreatorTag = 0x00720008
	// HangingProtocolCreationDateTimeTag is tag for Hanging Protocol Creation Date​Time
	HangingProtocolCreationDateTimeTag = 0x0072000A
	// HangingProtocolDefinitionSequenceTag is tag for Hanging Protocol Definition Sequence
	HangingProtocolDefinitionSequenceTag = 0x0072000C
	// HangingProtocolUserIdentificationCodeSequenceTag is tag for Hanging Protocol User Identification Code Sequence
	HangingProtocolUserIdentificationCodeSequenceTag = 0x0072000E
	// HangingProtocolUserGroupNameTag is tag for Hanging Protocol User Group Name
	HangingProtocolUserGroupNameTag = 0x00720010
	// SourceHangingProtocolSequenceTag is tag for Source Hanging Protocol Sequence
	SourceHangingProtocolSequenceTag = 0x00720012
	// NumberOfPriorsReferencedTag is tag for Number of Priors Referenced
	NumberOfPriorsReferencedTag = 0x00720014
	// ImageSetsSequenceTag is tag for Image Sets Sequence
	ImageSetsSequenceTag = 0x00720020
	// ImageSetSelectorSequenceTag is tag for Image Set Selector Sequence
	ImageSetSelectorSequenceTag = 0x00720022
	// ImageSetSelectorUsageFlagTag is tag for Image Set Selector Usage Flag
	ImageSetSelectorUsageFlagTag = 0x00720024
	// SelectorAttributeTag is tag for Selector Attribute
	SelectorAttributeTag = 0x00720026
	// SelectorValueNumberTag is tag for Selector Value Number
	SelectorValueNumberTag = 0x00720028
	// TimeBasedImageSetsSequenceTag is tag for Time Based Image Sets Sequence
	TimeBasedImageSetsSequenceTag = 0x00720030
	// ImageSetNumberTag is tag for Image Set Number
	ImageSetNumberTag = 0x00720032
	// ImageSetSelectorCategoryTag is tag for Image Set Selector Category
	ImageSetSelectorCategoryTag = 0x00720034
	// RelativeTimeTag is tag for Relative Time
	RelativeTimeTag = 0x00720038
	// RelativeTimeUnitsTag is tag for Relative Time Units
	RelativeTimeUnitsTag = 0x0072003A
	// AbstractPriorValueTag is tag for Abstract Prior Value
	AbstractPriorValueTag = 0x0072003C
	// AbstractPriorCodeSequenceTag is tag for Abstract Prior Code Sequence
	AbstractPriorCodeSequenceTag = 0x0072003E
	// ImageSetLabelTag is tag for Image Set Label
	ImageSetLabelTag = 0x00720040
	// SelectorAttributeVRTag is tag for Selector Attribute VR
	SelectorAttributeVRTag = 0x00720050
	// SelectorSequencePointerTag is tag for Selector Sequence Pointer
	SelectorSequencePointerTag = 0x00720052
	// SelectorSequencePointerPrivateCreatorTag is tag for Selector Sequence Pointer Private Creator
	SelectorSequencePointerPrivateCreatorTag = 0x00720054
	// SelectorAttributePrivateCreatorTag is tag for Selector Attribute Private Creator
	SelectorAttributePrivateCreatorTag = 0x00720056
	// SelectorAEValueTag is tag for Selector AE Value
	SelectorAEValueTag = 0x0072005E
	// SelectorASValueTag is tag for Selector AS Value
	SelectorASValueTag = 0x0072005F
	// SelectorATValueTag is tag for Selector AT Value
	SelectorATValueTag = 0x00720060
	// SelectorDAValueTag is tag for Selector DA Value
	SelectorDAValueTag = 0x00720061
	// SelectorCSValueTag is tag for Selector CS Value
	SelectorCSValueTag = 0x00720062
	// SelectorDTValueTag is tag for Selector DT Value
	SelectorDTValueTag = 0x00720063
	// SelectorISValueTag is tag for Selector IS Value
	SelectorISValueTag = 0x00720064
	// SelectorOBValueTag is tag for Selector OB Value
	SelectorOBValueTag = 0x00720065
	// SelectorLOValueTag is tag for Selector LO Value
	SelectorLOValueTag = 0x00720066
	// SelectorOFValueTag is tag for Selector OF Value
	SelectorOFValueTag = 0x00720067
	// SelectorLTValueTag is tag for Selector LT Value
	SelectorLTValueTag = 0x00720068
	// SelectorOWValueTag is tag for Selector OW Value
	SelectorOWValueTag = 0x00720069
	// SelectorPNValueTag is tag for Selector PN Value
	SelectorPNValueTag = 0x0072006A
	// SelectorTMValueTag is tag for Selector TM Value
	SelectorTMValueTag = 0x0072006B
	// SelectorSHValueTag is tag for Selector SH Value
	SelectorSHValueTag = 0x0072006C
	// SelectorUNValueTag is tag for Selector UN Value
	SelectorUNValueTag = 0x0072006D
	// SelectorSTValueTag is tag for Selector ST Value
	SelectorSTValueTag = 0x0072006E
	// SelectorUCValueTag is tag for Selector UC Value
	SelectorUCValueTag = 0x0072006F
	// SelectorUTValueTag is tag for Selector UT Value
	SelectorUTValueTag = 0x00720070
	// SelectorURValueTag is tag for Selector UR Value
	SelectorURValueTag = 0x00720071
	// SelectorDSValueTag is tag for Selector DS Value
	SelectorDSValueTag = 0x00720072
	// SelectorODValueTag is tag for Selector OD Value
	SelectorODValueTag = 0x00720073
	// SelectorFDValueTag is tag for Selector FD Value
	SelectorFDValueTag = 0x00720074
	// SelectorOLValueTag is tag for Selector OL Value
	SelectorOLValueTag = 0x00720075
	// SelectorFLValueTag is tag for Selector FL Value
	SelectorFLValueTag = 0x00720076
	// SelectorULValueTag is tag for Selector UL Value
	SelectorULValueTag = 0x00720078
	// SelectorUSValueTag is tag for Selector US Value
	SelectorUSValueTag = 0x0072007A
	// SelectorSLValueTag is tag for Selector SL Value
	SelectorSLValueTag = 0x0072007C
	// SelectorSSValueTag is tag for Selector SS Value
	SelectorSSValueTag = 0x0072007E
	// SelectorUIValueTag is tag for Selector UI Value
	SelectorUIValueTag = 0x0072007F
	// SelectorCodeSequenceValueTag is tag for Selector Code Sequence Value
	SelectorCodeSequenceValueTag = 0x00720080
	// NumberOfScreensTag is tag for Number of Screens
	NumberOfScreensTag = 0x00720100
	// NominalScreenDefinitionSequenceTag is tag for Nominal Screen Definition Sequence
	NominalScreenDefinitionSequenceTag = 0x00720102
	// NumberOfVerticalPixelsTag is tag for Number of Vertical Pixels
	NumberOfVerticalPixelsTag = 0x00720104
	// NumberOfHorizontalPixelsTag is tag for Number of Horizontal Pixels
	NumberOfHorizontalPixelsTag = 0x00720106
	// DisplayEnvironmentSpatialPositionTag is tag for Display Environment Spatial Position
	DisplayEnvironmentSpatialPositionTag = 0x00720108
	// ScreenMinimumGrayscaleBitDepthTag is tag for Screen Minimum Grayscale Bit Depth
	ScreenMinimumGrayscaleBitDepthTag = 0x0072010A
	// ScreenMinimumColorBitDepthTag is tag for Screen Minimum Color Bit Depth
	ScreenMinimumColorBitDepthTag = 0x0072010C
	// ApplicationMaximumRepaintTimeTag is tag for Application Maximum Repaint Time
	ApplicationMaximumRepaintTimeTag = 0x0072010E
	// DisplaySetsSequenceTag is tag for Display Sets Sequence
	DisplaySetsSequenceTag = 0x00720200
	// DisplaySetNumberTag is tag for Display Set Number
	DisplaySetNumberTag = 0x00720202
	// DisplaySetLabelTag is tag for Display Set Label
	DisplaySetLabelTag = 0x00720203
	// DisplaySetPresentationGroupTag is tag for Display Set Presentation Group
	DisplaySetPresentationGroupTag = 0x00720204
	// DisplaySetPresentationGroupDescriptionTag is tag for Display Set Presentation Group Description
	DisplaySetPresentationGroupDescriptionTag = 0x00720206
	// PartialDataDisplayHandlingTag is tag for Partial Data Display Handling
	PartialDataDisplayHandlingTag = 0x00720208
	// SynchronizedScrollingSequenceTag is tag for Synchronized Scrolling Sequence
	SynchronizedScrollingSequenceTag = 0x00720210
	// DisplaySetScrollingGroupTag is tag for Display Set Scrolling Group
	DisplaySetScrollingGroupTag = 0x00720212
	// NavigationIndicatorSequenceTag is tag for Navigation Indicator Sequence
	NavigationIndicatorSequenceTag = 0x00720214
	// NavigationDisplaySetTag is tag for Navigation Display Set
	NavigationDisplaySetTag = 0x00720216
	// ReferenceDisplaySetsTag is tag for Reference Display Sets
	ReferenceDisplaySetsTag = 0x00720218
	// ImageBoxesSequenceTag is tag for Image Boxes Sequence
	ImageBoxesSequenceTag = 0x00720300
	// ImageBoxNumberTag is tag for Image Box Number
	ImageBoxNumberTag = 0x00720302
	// ImageBoxLayoutTypeTag is tag for Image Box Layout Type
	ImageBoxLayoutTypeTag = 0x00720304
	// ImageBoxTileHorizontalDimensionTag is tag for Image Box Tile Horizontal Dimension
	ImageBoxTileHorizontalDimensionTag = 0x00720306
	// ImageBoxTileVerticalDimensionTag is tag for Image Box Tile Vertical Dimension
	ImageBoxTileVerticalDimensionTag = 0x00720308
	// ImageBoxScrollDirectionTag is tag for Image Box Scroll Direction
	ImageBoxScrollDirectionTag = 0x00720310
	// ImageBoxSmallScrollTypeTag is tag for Image Box Small Scroll Type
	ImageBoxSmallScrollTypeTag = 0x00720312
	// ImageBoxSmallScrollAmountTag is tag for Image Box Small Scroll Amount
	ImageBoxSmallScrollAmountTag = 0x00720314
	// ImageBoxLargeScrollTypeTag is tag for Image Box Large Scroll Type
	ImageBoxLargeScrollTypeTag = 0x00720316
	// ImageBoxLargeScrollAmountTag is tag for Image Box Large Scroll Amount
	ImageBoxLargeScrollAmountTag = 0x00720318
	// ImageBoxOverlapPriorityTag is tag for Image Box Overlap Priority
	ImageBoxOverlapPriorityTag = 0x00720320
	// CineRelativeToRealTimeTag is tag for Cine Relative to Real-Time
	CineRelativeToRealTimeTag = 0x00720330
	// FilterOperationsSequenceTag is tag for Filter Operations Sequence
	FilterOperationsSequenceTag = 0x00720400
	// FilterByCategoryTag is tag for Filter-by Category
	FilterByCategoryTag = 0x00720402
	// FilterByAttributePresenceTag is tag for Filter-by Attribute Presence
	FilterByAttributePresenceTag = 0x00720404
	// FilterByOperatorTag is tag for Filter-by Operator
	FilterByOperatorTag = 0x00720406
	// StructuredDisplayBackgroundCIELabValueTag is tag for Structured Display Background CIELab Value
	StructuredDisplayBackgroundCIELabValueTag = 0x00720420
	// EmptyImageBoxCIELabValueTag is tag for Empty Image Box CIELab Value
	EmptyImageBoxCIELabValueTag = 0x00720421
	// StructuredDisplayImageBoxSequenceTag is tag for Structured Display Image Box Sequence
	StructuredDisplayImageBoxSequenceTag = 0x00720422
	// StructuredDisplayTextBoxSequenceTag is tag for Structured Display Text Box Sequence
	StructuredDisplayTextBoxSequenceTag = 0x00720424
	// ReferencedFirstFrameSequenceTag is tag for Referenced First Frame Sequence
	ReferencedFirstFrameSequenceTag = 0x00720427
	// ImageBoxSynchronizationSequenceTag is tag for Image Box Synchronization Sequence
	ImageBoxSynchronizationSequenceTag = 0x00720430
	// SynchronizedImageBoxListTag is tag for Synchronized Image Box List
	SynchronizedImageBoxListTag = 0x00720432
	// TypeOfSynchronizationTag is tag for Type of Synchronization
	TypeOfSynchronizationTag = 0x00720434
	// BlendingOperationTypeTag is tag for Blending Operation Type
	BlendingOperationTypeTag = 0x00720500
	// ReformattingOperationTypeTag is tag for Reformatting Operation Type
	ReformattingOperationTypeTag = 0x00720510
	// ReformattingThicknessTag is tag for Reformatting Thickness
	ReformattingThicknessTag = 0x00720512
	// ReformattingIntervalTag is tag for Reformatting Interval
	ReformattingIntervalTag = 0x00720514
	// ReformattingOperationInitialViewDirectionTag is tag for Reformatting Operation Initial View Direction
	ReformattingOperationInitialViewDirectionTag = 0x00720516
	// ThreeDRenderingTypeTag is tag for 3D Rendering Type
	ThreeDRenderingTypeTag = 0x00720520
	// SortingOperationsSequenceTag is tag for Sorting Operations Sequence
	SortingOperationsSequenceTag = 0x00720600
	// SortByCategoryTag is tag for Sort-by Category
	SortByCategoryTag = 0x00720602
	// SortingDirectionTag is tag for Sorting Direction
	SortingDirectionTag = 0x00720604
	// DisplaySetPatientOrientationTag is tag for Display Set Patient Orientation
	DisplaySetPatientOrientationTag = 0x00720700
	// VOITypeTag is tag for VOI Type
	VOITypeTag = 0x00720702
	// PseudoColorTypeTag is tag for Pseudo-Color Type
	PseudoColorTypeTag = 0x00720704
	// PseudoColorPaletteInstanceReferenceSequenceTag is tag for Pseudo-Color Palette Instance Reference Sequence
	PseudoColorPaletteInstanceReferenceSequenceTag = 0x00720705
	// ShowGrayscaleInvertedTag is tag for Show Grayscale Inverted
	ShowGrayscaleInvertedTag = 0x00720706
	// ShowImageTrueSizeFlagTag is tag for Show Image True Size Flag
	ShowImageTrueSizeFlagTag = 0x00720710
	// ShowGraphicAnnotationFlagTag is tag for Show Graphic Annotation Flag
	ShowGraphicAnnotationFlagTag = 0x00720712
	// ShowPatientDemographicsFlagTag is tag for Show Patient Demographics Flag
	ShowPatientDemographicsFlagTag = 0x00720714
	// ShowAcquisitionTechniquesFlagTag is tag for Show Acquisition Techniques Flag
	ShowAcquisitionTechniquesFlagTag = 0x00720716
	// DisplaySetHorizontalJustificationTag is tag for Display Set Horizontal Justification
	DisplaySetHorizontalJustificationTag = 0x00720717
	// DisplaySetVerticalJustificationTag is tag for Display Set Vertical Justification
	DisplaySetVerticalJustificationTag = 0x00720718
	// ContinuationStartMetersetTag is tag for Continuation Start Meterset
	ContinuationStartMetersetTag = 0x00740120
	// ContinuationEndMetersetTag is tag for Continuation End Meterset
	ContinuationEndMetersetTag = 0x00740121
	// ProcedureStepStateTag is tag for Procedure Step State
	ProcedureStepStateTag = 0x00741000
	// ProcedureStepProgressInformationSequenceTag is tag for Procedure Step Progress Information Sequence
	ProcedureStepProgressInformationSequenceTag = 0x00741002
	// ProcedureStepProgressTag is tag for Procedure Step Progress
	ProcedureStepProgressTag = 0x00741004
	// ProcedureStepProgressDescriptionTag is tag for Procedure Step Progress Description
	ProcedureStepProgressDescriptionTag = 0x00741006
	// ProcedureStepProgressParametersSequenceTag is tag for Procedure Step Progress Parameters Sequence
	ProcedureStepProgressParametersSequenceTag = 0x00741007
	// ProcedureStepCommunicationsURISequenceTag is tag for Procedure Step Communications URI Sequence
	ProcedureStepCommunicationsURISequenceTag = 0x00741008
	// ContactURITag is tag for Contact URI
	ContactURITag = 0x0074100A
	// ContactDisplayNameTag is tag for Contact Display Name
	ContactDisplayNameTag = 0x0074100C
	// ProcedureStepDiscontinuationReasonCodeSequenceTag is tag for Procedure Step Discontinuation Reason Code Sequence
	ProcedureStepDiscontinuationReasonCodeSequenceTag = 0x0074100E
	// BeamTaskSequenceTag is tag for Beam Task Sequence
	BeamTaskSequenceTag = 0x00741020
	// BeamTaskTypeTag is tag for Beam Task Type
	BeamTaskTypeTag = 0x00741022
	// BeamOrderIndexTrialTag is tag for Beam Order Index (Trial)
	BeamOrderIndexTrialTag = 0x00741024
	// AutosequenceFlagTag is tag for Autosequence Flag
	AutosequenceFlagTag = 0x00741025
	// TableTopVerticalAdjustedPositionTag is tag for Table Top Vertical Adjusted Position
	TableTopVerticalAdjustedPositionTag = 0x00741026
	// TableTopLongitudinalAdjustedPositionTag is tag for Table Top Longitudinal Adjusted Position
	TableTopLongitudinalAdjustedPositionTag = 0x00741027
	// TableTopLateralAdjustedPositionTag is tag for Table Top Lateral Adjusted Position
	TableTopLateralAdjustedPositionTag = 0x00741028
	// PatientSupportAdjustedAngleTag is tag for Patient Support Adjusted Angle
	PatientSupportAdjustedAngleTag = 0x0074102A
	// TableTopEccentricAdjustedAngleTag is tag for Table Top Eccentric Adjusted Angle
	TableTopEccentricAdjustedAngleTag = 0x0074102B
	// TableTopPitchAdjustedAngleTag is tag for Table Top Pitch Adjusted Angle
	TableTopPitchAdjustedAngleTag = 0x0074102C
	// TableTopRollAdjustedAngleTag is tag for Table Top Roll Adjusted Angle
	TableTopRollAdjustedAngleTag = 0x0074102D
	// DeliveryVerificationImageSequenceTag is tag for Delivery Verification Image Sequence
	DeliveryVerificationImageSequenceTag = 0x00741030
	// VerificationImageTimingTag is tag for Verification Image Timing
	VerificationImageTimingTag = 0x00741032
	// DoubleExposureFlagTag is tag for Double Exposure Flag
	DoubleExposureFlagTag = 0x00741034
	// DoubleExposureOrderingTag is tag for Double Exposure Ordering
	DoubleExposureOrderingTag = 0x00741036
	// DoubleExposureMetersetTrialTag is tag for Double Exposure Meterset (Trial)
	DoubleExposureMetersetTrialTag = 0x00741038
	// DoubleExposureFieldDeltaTrialTag is tag for Double Exposure Field Delta (Trial)
	DoubleExposureFieldDeltaTrialTag = 0x0074103A
	// RelatedReferenceRTImageSequenceTag is tag for Related Reference RT Image Sequence
	RelatedReferenceRTImageSequenceTag = 0x00741040
	// GeneralMachineVerificationSequenceTag is tag for General Machine Verification Sequence
	GeneralMachineVerificationSequenceTag = 0x00741042
	// ConventionalMachineVerificationSequenceTag is tag for Conventional Machine Verification Sequence
	ConventionalMachineVerificationSequenceTag = 0x00741044
	// IonMachineVerificationSequenceTag is tag for Ion Machine Verification Sequence
	IonMachineVerificationSequenceTag = 0x00741046
	// FailedAttributesSequenceTag is tag for Failed Attributes Sequence
	FailedAttributesSequenceTag = 0x00741048
	// OverriddenAttributesSequenceTag is tag for Overridden Attributes Sequence
	OverriddenAttributesSequenceTag = 0x0074104A
	// ConventionalControlPointVerificationSequenceTag is tag for Conventional Control Point Verification Sequence
	ConventionalControlPointVerificationSequenceTag = 0x0074104C
	// IonControlPointVerificationSequenceTag is tag for Ion Control Point Verification Sequence
	IonControlPointVerificationSequenceTag = 0x0074104E
	// AttributeOccurrenceSequenceTag is tag for Attribute Occurrence Sequence
	AttributeOccurrenceSequenceTag = 0x00741050
	// AttributeOccurrencePointerTag is tag for Attribute Occurrence Pointer
	AttributeOccurrencePointerTag = 0x00741052
	// AttributeItemSelectorTag is tag for Attribute Item Selector
	AttributeItemSelectorTag = 0x00741054
	// AttributeOccurrencePrivateCreatorTag is tag for Attribute Occurrence Private Creator
	AttributeOccurrencePrivateCreatorTag = 0x00741056
	// SelectorSequencePointerItemsTag is tag for Selector Sequence Pointer Items
	SelectorSequencePointerItemsTag = 0x00741057
	// ScheduledProcedureStepPriorityTag is tag for Scheduled Procedure Step Priority
	ScheduledProcedureStepPriorityTag = 0x00741200
	// WorklistLabelTag is tag for Worklist Label
	WorklistLabelTag = 0x00741202
	// ProcedureStepLabelTag is tag for Procedure Step Label
	ProcedureStepLabelTag = 0x00741204
	// ScheduledProcessingParametersSequenceTag is tag for Scheduled Processing Parameters Sequence
	ScheduledProcessingParametersSequenceTag = 0x00741210
	// PerformedProcessingParametersSequenceTag is tag for Performed Processing Parameters Sequence
	PerformedProcessingParametersSequenceTag = 0x00741212
	// UnifiedProcedureStepPerformedProcedureSequenceTag is tag for Unified Procedure Step Performed Procedure Sequence
	UnifiedProcedureStepPerformedProcedureSequenceTag = 0x00741216
	// RelatedProcedureStepSequenceTag is tag for Related Procedure Step Sequence
	RelatedProcedureStepSequenceTag = 0x00741220
	// ProcedureStepRelationshipTypeTag is tag for Procedure Step Relationship Type
	ProcedureStepRelationshipTypeTag = 0x00741222
	// ReplacedProcedureStepSequenceTag is tag for Replaced Procedure Step Sequence
	ReplacedProcedureStepSequenceTag = 0x00741224
	// DeletionLockTag is tag for Deletion Lock
	DeletionLockTag = 0x00741230
	// ReceivingAETag is tag for Receiving AE
	ReceivingAETag = 0x00741234
	// RequestingAETag is tag for Requesting AE
	RequestingAETag = 0x00741236
	// ReasonForCancellationTag is tag for Reason for Cancellation
	ReasonForCancellationTag = 0x00741238
	// SCPStatusTag is tag for SCP Status
	SCPStatusTag = 0x00741242
	// SubscriptionListStatusTag is tag for Subscription List Status
	SubscriptionListStatusTag = 0x00741244
	// UnifiedProcedureStepListStatusTag is tag for Unified Procedure Step List Status
	UnifiedProcedureStepListStatusTag = 0x00741246
	// BeamOrderIndexTag is tag for Beam Order Index
	BeamOrderIndexTag = 0x00741324
	// DoubleExposureMetersetTag is tag for Double Exposure Meterset
	DoubleExposureMetersetTag = 0x00741338
	// DoubleExposureFieldDeltaTag is tag for Double Exposure Field Delta
	DoubleExposureFieldDeltaTag = 0x0074133A
	// BrachyTaskSequenceTag is tag for Brachy Task Sequence
	BrachyTaskSequenceTag = 0x00741401
	// ContinuationStartTotalReferenceAirKermaTag is tag for Continuation Start Total Reference Air Kerma
	ContinuationStartTotalReferenceAirKermaTag = 0x00741402
	// ContinuationEndTotalReferenceAirKermaTag is tag for Continuation End Total Reference Air Kerma
	ContinuationEndTotalReferenceAirKermaTag = 0x00741403
	// ContinuationPulseNumberTag is tag for Continuation Pulse Number
	ContinuationPulseNumberTag = 0x00741404
	// ChannelDeliveryOrderSequenceTag is tag for Channel Delivery Order Sequence
	ChannelDeliveryOrderSequenceTag = 0x00741405
	// ReferencedChannelNumberTag is tag for Referenced Channel Number
	ReferencedChannelNumberTag = 0x00741406
	// StartCumulativeTimeWeightTag is tag for Start Cumulative Time Weight
	StartCumulativeTimeWeightTag = 0x00741407
	// EndCumulativeTimeWeightTag is tag for End Cumulative Time Weight
	EndCumulativeTimeWeightTag = 0x00741408
	// OmittedChannelSequenceTag is tag for Omitted Channel Sequence
	OmittedChannelSequenceTag = 0x00741409
	// ReasonForChannelOmissionTag is tag for Reason for Channel Omission
	ReasonForChannelOmissionTag = 0x0074140A
	// ReasonForChannelOmissionDescriptionTag is tag for Reason for Channel Omission Description
	ReasonForChannelOmissionDescriptionTag = 0x0074140B
	// ChannelDeliveryOrderIndexTag is tag for Channel Delivery Order Index
	ChannelDeliveryOrderIndexTag = 0x0074140C
	// ChannelDeliveryContinuationSequenceTag is tag for Channel Delivery Continuation Sequence
	ChannelDeliveryContinuationSequenceTag = 0x0074140D
	// OmittedApplicationSetupSequenceTag is tag for Omitted Application Setup Sequence
	OmittedApplicationSetupSequenceTag = 0x0074140E
	// ImplantAssemblyTemplateNameTag is tag for Implant Assembly Template Name
	ImplantAssemblyTemplateNameTag = 0x00760001
	// ImplantAssemblyTemplateIssuerTag is tag for Implant Assembly Template Issuer
	ImplantAssemblyTemplateIssuerTag = 0x00760003
	// ImplantAssemblyTemplateVersionTag is tag for Implant Assembly Template Version
	ImplantAssemblyTemplateVersionTag = 0x00760006
	// ReplacedImplantAssemblyTemplateSequenceTag is tag for Replaced Implant Assembly Template Sequence
	ReplacedImplantAssemblyTemplateSequenceTag = 0x00760008
	// ImplantAssemblyTemplateTypeTag is tag for Implant Assembly Template Type
	ImplantAssemblyTemplateTypeTag = 0x0076000A
	// OriginalImplantAssemblyTemplateSequenceTag is tag for Original Implant Assembly Template Sequence
	OriginalImplantAssemblyTemplateSequenceTag = 0x0076000C
	// DerivationImplantAssemblyTemplateSequenceTag is tag for Derivation Implant Assembly Template Sequence
	DerivationImplantAssemblyTemplateSequenceTag = 0x0076000E
	// ImplantAssemblyTemplateTargetAnatomySequenceTag is tag for Implant Assembly Template Target Anatomy Sequence
	ImplantAssemblyTemplateTargetAnatomySequenceTag = 0x00760010
	// ProcedureTypeCodeSequenceTag is tag for Procedure Type Code Sequence
	ProcedureTypeCodeSequenceTag = 0x00760020
	// SurgicalTechniqueTag is tag for Surgical Technique
	SurgicalTechniqueTag = 0x00760030
	// ComponentTypesSequenceTag is tag for Component Types Sequence
	ComponentTypesSequenceTag = 0x00760032
	// ComponentTypeCodeSequenceTag is tag for Component Type Code Sequence
	ComponentTypeCodeSequenceTag = 0x00760034
	// ExclusiveComponentTypeTag is tag for Exclusive Component Type
	ExclusiveComponentTypeTag = 0x00760036
	// MandatoryComponentTypeTag is tag for Mandatory Component Type
	MandatoryComponentTypeTag = 0x00760038
	// ComponentSequenceTag is tag for Component Sequence
	ComponentSequenceTag = 0x00760040
	// ComponentIDTag is tag for Component ID
	ComponentIDTag = 0x00760055
	// ComponentAssemblySequenceTag is tag for Component Assembly Sequence
	ComponentAssemblySequenceTag = 0x00760060
	// Component1ReferencedIDTag is tag for Component 1 Referenced ID
	Component1ReferencedIDTag = 0x00760070
	// Component1ReferencedMatingFeatureSetIDTag is tag for Component 1 Referenced Mating Feature Set ID
	Component1ReferencedMatingFeatureSetIDTag = 0x00760080
	// Component1ReferencedMatingFeatureIDTag is tag for Component 1 Referenced Mating Feature ID
	Component1ReferencedMatingFeatureIDTag = 0x00760090
	// Component2ReferencedIDTag is tag for Component 2 Referenced ID
	Component2ReferencedIDTag = 0x007600A0
	// Component2ReferencedMatingFeatureSetIDTag is tag for Component 2 Referenced Mating Feature Set ID
	Component2ReferencedMatingFeatureSetIDTag = 0x007600B0
	// Component2ReferencedMatingFeatureIDTag is tag for Component 2 Referenced Mating Feature ID
	Component2ReferencedMatingFeatureIDTag = 0x007600C0
	// ImplantTemplateGroupNameTag is tag for Implant Template Group Name
	ImplantTemplateGroupNameTag = 0x00780001
	// ImplantTemplateGroupDescriptionTag is tag for Implant Template Group Description
	ImplantTemplateGroupDescriptionTag = 0x00780010
	// ImplantTemplateGroupIssuerTag is tag for Implant Template Group Issuer
	ImplantTemplateGroupIssuerTag = 0x00780020
	// ImplantTemplateGroupVersionTag is tag for Implant Template Group Version
	ImplantTemplateGroupVersionTag = 0x00780024
	// ReplacedImplantTemplateGroupSequenceTag is tag for Replaced Implant Template Group Sequence
	ReplacedImplantTemplateGroupSequenceTag = 0x00780026
	// ImplantTemplateGroupTargetAnatomySequenceTag is tag for Implant Template Group Target Anatomy Sequence
	ImplantTemplateGroupTargetAnatomySequenceTag = 0x00780028
	// ImplantTemplateGroupMembersSequenceTag is tag for Implant Template Group Members Sequence
	ImplantTemplateGroupMembersSequenceTag = 0x0078002A
	// ImplantTemplateGroupMemberIDTag is tag for Implant Template Group Member ID
	ImplantTemplateGroupMemberIDTag = 0x0078002E
	// ThreeDImplantTemplateGroupMemberMatchingPointTag is tag for 3D Implant Template Group Member Matching Point
	ThreeDImplantTemplateGroupMemberMatchingPointTag = 0x00780050
	// ThreeDImplantTemplateGroupMemberMatchingAxesTag is tag for 3D Implant Template Group Member Matching Axes
	ThreeDImplantTemplateGroupMemberMatchingAxesTag = 0x00780060
	// ImplantTemplateGroupMemberMatching2DCoordinatesSequenceTag is tag for Implant Template Group Member Matching 2D Coordinates Sequence
	ImplantTemplateGroupMemberMatching2DCoordinatesSequenceTag = 0x00780070
	// TwoDImplantTemplateGroupMemberMatchingPointTag is tag for 2D Implant Template Group Member Matching Point
	TwoDImplantTemplateGroupMemberMatchingPointTag = 0x00780090
	// TwoDImplantTemplateGroupMemberMatchingAxesTag is tag for 2D Implant Template Group Member Matching Axes
	TwoDImplantTemplateGroupMemberMatchingAxesTag = 0x007800A0
	// ImplantTemplateGroupVariationDimensionSequenceTag is tag for Implant Template Group Variation Dimension Sequence
	ImplantTemplateGroupVariationDimensionSequenceTag = 0x007800B0
	// ImplantTemplateGroupVariationDimensionNameTag is tag for Implant Template Group Variation Dimension Name
	ImplantTemplateGroupVariationDimensionNameTag = 0x007800B2
	// ImplantTemplateGroupVariationDimensionRankSequenceTag is tag for Implant Template Group Variation Dimension Rank Sequence
	ImplantTemplateGroupVariationDimensionRankSequenceTag = 0x007800B4
	// ReferencedImplantTemplateGroupMemberIDTag is tag for Referenced Implant Template Group Member ID
	ReferencedImplantTemplateGroupMemberIDTag = 0x007800B6
	// ImplantTemplateGroupVariationDimensionRankTag is tag for Implant Template Group Variation Dimension Rank
	ImplantTemplateGroupVariationDimensionRankTag = 0x007800B8
	// SurfaceScanAcquisitionTypeCodeSequenceTag is tag for Surface Scan Acquisition Type Code Sequence
	SurfaceScanAcquisitionTypeCodeSequenceTag = 0x00800001
	// SurfaceScanModeCodeSequenceTag is tag for Surface Scan Mode Code Sequence
	SurfaceScanModeCodeSequenceTag = 0x00800002
	// RegistrationMethodCodeSequenceTag is tag for Registration Method Code Sequence
	RegistrationMethodCodeSequenceTag = 0x00800003
	// ShotDurationTimeTag is tag for Shot Duration Time
	ShotDurationTimeTag = 0x00800004
	// ShotOffsetTimeTag is tag for Shot Offset Time
	ShotOffsetTimeTag = 0x00800005
	// SurfacePointPresentationValueDataTag is tag for Surface Point Presentation Value Data
	SurfacePointPresentationValueDataTag = 0x00800006
	// SurfacePointColorCIELabValueDataTag is tag for Surface Point Color CIELab Value Data
	SurfacePointColorCIELabValueDataTag = 0x00800007
	// UVMappingSequenceTag is tag for UV Mapping Sequence
	UVMappingSequenceTag = 0x00800008
	// TextureLabelTag is tag for Texture Label
	TextureLabelTag = 0x00800009
	// UValueDataTag is tag for U Value Data
	UValueDataTag = 0x00800010
	// VValueDataTag is tag for V Value Data
	VValueDataTag = 0x00800011
	// ReferencedTextureSequenceTag is tag for Referenced Texture Sequence
	ReferencedTextureSequenceTag = 0x00800012
	// ReferencedSurfaceDataSequenceTag is tag for Referenced Surface Data Sequence
	ReferencedSurfaceDataSequenceTag = 0x00800013
	// AssessmentSummaryTag is tag for Assessment Summary
	AssessmentSummaryTag = 0x00820001
	// AssessmentSummaryDescriptionTag is tag for Assessment Summary Description
	AssessmentSummaryDescriptionTag = 0x00820003
	// AssessedSOPInstanceSequenceTag is tag for Assessed SOP Instance Sequence
	AssessedSOPInstanceSequenceTag = 0x00820004
	// ReferencedComparisonSOPInstanceSequenceTag is tag for Referenced Comparison SOP Instance Sequence
	ReferencedComparisonSOPInstanceSequenceTag = 0x00820005
	// NumberOfAssessmentObservationsTag is tag for Number of Assessment Observations
	NumberOfAssessmentObservationsTag = 0x00820006
	// AssessmentObservationsSequenceTag is tag for Assessment Observations Sequence
	AssessmentObservationsSequenceTag = 0x00820007
	// ObservationSignificanceTag is tag for Observation Significance
	ObservationSignificanceTag = 0x00820008
	// ObservationDescriptionTag is tag for Observation Description
	ObservationDescriptionTag = 0x0082000A
	// StructuredConstraintObservationSequenceTag is tag for Structured Constraint Observation Sequence
	StructuredConstraintObservationSequenceTag = 0x0082000C
	// AssessedAttributeValueSequenceTag is tag for Assessed Attribute Value Sequence
	AssessedAttributeValueSequenceTag = 0x00820010
	// AssessmentSetIDTag is tag for Assessment Set ID
	AssessmentSetIDTag = 0x00820016
	// AssessmentRequesterSequenceTag is tag for Assessment Requester Sequence
	AssessmentRequesterSequenceTag = 0x00820017
	// SelectorAttributeNameTag is tag for Selector Attribute Name
	SelectorAttributeNameTag = 0x00820018
	// SelectorAttributeKeywordTag is tag for Selector Attribute Keyword
	SelectorAttributeKeywordTag = 0x00820019
	// AssessmentTypeCodeSequenceTag is tag for Assessment Type Code Sequence
	AssessmentTypeCodeSequenceTag = 0x00820021
	// ObservationBasisCodeSequenceTag is tag for Observation Basis Code Sequence
	ObservationBasisCodeSequenceTag = 0x00820022
	// AssessmentLabelTag is tag for Assessment Label
	AssessmentLabelTag = 0x00820023
	// ConstraintTypeTag is tag for Constraint Type
	ConstraintTypeTag = 0x00820032
	// SpecificationSelectionGuidanceTag is tag for Specification Selection Guidance
	SpecificationSelectionGuidanceTag = 0x00820033
	// ConstraintValueSequenceTag is tag for Constraint Value Sequence
	ConstraintValueSequenceTag = 0x00820034
	// RecommendedDefaultValueSequenceTag is tag for Recommended Default Value Sequence
	RecommendedDefaultValueSequenceTag = 0x00820035
	// ConstraintViolationSignificanceTag is tag for Constraint Violation Significance
	ConstraintViolationSignificanceTag = 0x00820036
	// ConstraintViolationConditionTag is tag for Constraint Violation Condition
	ConstraintViolationConditionTag = 0x00820037
	// ModifiableConstraintFlagTag is tag for Modifiable Constraint Flag
	ModifiableConstraintFlagTag = 0x00820038
	// StorageMediaFileSetIDTag is tag for Storage Media File-set ID
	StorageMediaFileSetIDTag = 0x00880130
	// StorageMediaFileSetUIDTag is tag for Storage Media File-set UID
	StorageMediaFileSetUIDTag = 0x00880140
	// IconImageSequenceTag is tag for Icon Image Sequence
	IconImageSequenceTag = 0x00880200
	// TopicTitleTag is tag for Topic Title
	TopicTitleTag = 0x00880904
	// TopicSubjectTag is tag for Topic Subject
	TopicSubjectTag = 0x00880906
	// TopicAuthorTag is tag for Topic Author
	TopicAuthorTag = 0x00880910
	// TopicKeywordsTag is tag for Topic Keywords
	TopicKeywordsTag = 0x00880912
	// SOPInstanceStatusTag is tag for SOP Instance Status
	SOPInstanceStatusTag = 0x01000410
	// SOPAuthorizationDateTimeTag is tag for SOP Authorization DateTime
	SOPAuthorizationDateTimeTag = 0x01000420
	// SOPAuthorizationCommentTag is tag for SOP Authorization Comment
	SOPAuthorizationCommentTag = 0x01000424
	// AuthorizationEquipmentCertificationNumberTag is tag for Authorization Equipment Certification Number
	AuthorizationEquipmentCertificationNumberTag = 0x01000426
	// MACIDNumberTag is tag for MAC ID Number
	MACIDNumberTag = 0x04000005
	// MACCalculationTransferSyntaxUIDTag is tag for MAC Calculation Transfer Syntax UID
	MACCalculationTransferSyntaxUIDTag = 0x04000010
	// MACAlgorithmTag is tag for MAC Algorithm
	MACAlgorithmTag = 0x04000015
	// DataElementsSignedTag is tag for Data Elements Signed
	DataElementsSignedTag = 0x04000020
	// DigitalSignatureUIDTag is tag for Digital Signature UID
	DigitalSignatureUIDTag = 0x04000100
	// DigitalSignatureDateTimeTag is tag for Digital Signature DateTime
	DigitalSignatureDateTimeTag = 0x04000105
	// CertificateTypeTag is tag for Certificate Type
	CertificateTypeTag = 0x04000110
	// CertificateOfSignerTag is tag for Certificate of Signer
	CertificateOfSignerTag = 0x04000115
	// SignatureTag is tag for Signature
	SignatureTag = 0x04000120
	// CertifiedTimestampTypeTag is tag for Certified Timestamp Type
	CertifiedTimestampTypeTag = 0x04000305
	// CertifiedTimestampTag is tag for Certified Timestamp
	CertifiedTimestampTag = 0x04000310
	// DigitalSignaturePurposeCodeSequenceTag is tag for Digital Signature Purpose Code Sequence
	DigitalSignaturePurposeCodeSequenceTag = 0x04000401
	// ReferencedDigitalSignatureSequenceTag is tag for Referenced Digital Signature Sequence
	ReferencedDigitalSignatureSequenceTag = 0x04000402
	// ReferencedSOPInstanceMACSequenceTag is tag for Referenced SOP Instance MAC Sequence
	ReferencedSOPInstanceMACSequenceTag = 0x04000403
	// MACTag is tag for MAC
	MACTag = 0x04000404
	// EncryptedAttributesSequenceTag is tag for Encrypted Attributes Sequence
	EncryptedAttributesSequenceTag = 0x04000500
	// EncryptedContentTransferSyntaxUIDTag is tag for Encrypted Content Transfer Syntax UID
	EncryptedContentTransferSyntaxUIDTag = 0x04000510
	// EncryptedContentTag is tag for Encrypted Content
	EncryptedContentTag = 0x04000520
	// ModifiedAttributesSequenceTag is tag for Modified Attributes Sequence
	ModifiedAttributesSequenceTag = 0x04000550
	// NonconformingModifiedAttributesSequenceTag is tag for Nonconforming Modified Attributes Sequence
	NonconformingModifiedAttributesSequenceTag = 0x04000551
	// NonconformingDataElementValueTag is tag for Nonconforming Data Element Value
	NonconformingDataElementValueTag = 0x04000552
	// OriginalAttributesSequenceTag is tag for Original Attributes Sequence
	OriginalAttributesSequenceTag = 0x04000561
	// AttributeModificationDateTimeTag is tag for Attribute Modification DateTime
	AttributeModificationDateTimeTag = 0x04000562
	// ModifyingSystemTag is tag for Modifying System
	ModifyingSystemTag = 0x04000563
	// SourceOfPreviousValuesTag is tag for Source of Previous Values
	SourceOfPreviousValuesTag = 0x04000564
	// ReasonForTheAttributeModificationTag is tag for Reason for the Attribute Modification
	ReasonForTheAttributeModificationTag = 0x04000565
	// InstanceOriginStatusTag is tag for Instance Origin Status
	InstanceOriginStatusTag = 0x04000600
	// EscapeTripletTag is tag for Escape Triplet
	EscapeTripletTag = 0x10000000
	// RunLengthTripletTag is tag for Run Length Triplet
	RunLengthTripletTag = 0x10000001
	// HuffmanTableSizeTag is tag for Huffman Table Size
	HuffmanTableSizeTag = 0x10000002
	// HuffmanTableTripletTag is tag for Huffman Table Triplet
	HuffmanTableTripletTag = 0x10000003
	// ShiftTableSizeTag is tag for Shift Table Size
	ShiftTableSizeTag = 0x10000004
	// ShiftTableTripletTag is tag for Shift Table Triplet
	ShiftTableTripletTag = 0x10000005
	// ZonalMapTag is tag for Zonal Map
	ZonalMapTag = 0x10100000
	// NumberOfCopiesTag is tag for Number of Copies
	NumberOfCopiesTag = 0x20000010
	// PrinterConfigurationSequenceTag is tag for Printer Configuration Sequence
	PrinterConfigurationSequenceTag = 0x2000001E
	// PrintPriorityTag is tag for Print Priority
	PrintPriorityTag = 0x20000020
	// MediumTypeTag is tag for Medium Type
	MediumTypeTag = 0x20000030
	// FilmDestinationTag is tag for Film Destination
	FilmDestinationTag = 0x20000040
	// FilmSessionLabelTag is tag for Film Session Label
	FilmSessionLabelTag = 0x20000050
	// MemoryAllocationTag is tag for Memory Allocation
	MemoryAllocationTag = 0x20000060
	// MaximumMemoryAllocationTag is tag for Maximum Memory Allocation
	MaximumMemoryAllocationTag = 0x20000061
	// ColorImagePrintingFlagTag is tag for Color Image Printing Flag
	ColorImagePrintingFlagTag = 0x20000062
	// CollationFlagTag is tag for Collation Flag
	CollationFlagTag = 0x20000063
	// AnnotationFlagTag is tag for Annotation Flag
	AnnotationFlagTag = 0x20000065
	// ImageOverlayFlagTag is tag for Image Overlay Flag
	ImageOverlayFlagTag = 0x20000067
	// PresentationLUTFlagTag is tag for Presentation LUT Flag
	PresentationLUTFlagTag = 0x20000069
	// ImageBoxPresentationLUTFlagTag is tag for Image Box Presentation LUT Flag
	ImageBoxPresentationLUTFlagTag = 0x2000006A
	// MemoryBitDepthTag is tag for Memory Bit Depth
	MemoryBitDepthTag = 0x200000A0
	// PrintingBitDepthTag is tag for Printing Bit Depth
	PrintingBitDepthTag = 0x200000A1
	// MediaInstalledSequenceTag is tag for Media Installed Sequence
	MediaInstalledSequenceTag = 0x200000A2
	// OtherMediaAvailableSequenceTag is tag for Other Media Available Sequence
	OtherMediaAvailableSequenceTag = 0x200000A4
	// SupportedImageDisplayFormatsSequenceTag is tag for Supported Image Display Formats Sequence
	SupportedImageDisplayFormatsSequenceTag = 0x200000A8
	// ReferencedFilmBoxSequenceTag is tag for Referenced Film Box Sequence
	ReferencedFilmBoxSequenceTag = 0x20000500
	// ReferencedStoredPrintSequenceTag is tag for Referenced Stored Print Sequence
	ReferencedStoredPrintSequenceTag = 0x20000510
	// ImageDisplayFormatTag is tag for Image Display Format
	ImageDisplayFormatTag = 0x20100010
	// AnnotationDisplayFormatIDTag is tag for Annotation Display Format ID
	AnnotationDisplayFormatIDTag = 0x20100030
	// FilmOrientationTag is tag for Film Orientation
	FilmOrientationTag = 0x20100040
	// FilmSizeIDTag is tag for Film Size ID
	FilmSizeIDTag = 0x20100050
	// PrinterResolutionIDTag is tag for Printer Resolution ID
	PrinterResolutionIDTag = 0x20100052
	// DefaultPrinterResolutionIDTag is tag for Default Printer Resolution ID
	DefaultPrinterResolutionIDTag = 0x20100054
	// MagnificationTypeTag is tag for Magnification Type
	MagnificationTypeTag = 0x20100060
	// SmoothingTypeTag is tag for Smoothing Type
	SmoothingTypeTag = 0x20100080
	// DefaultMagnificationTypeTag is tag for Default Magnification Type
	DefaultMagnificationTypeTag = 0x201000A6
	// OtherMagnificationTypesAvailableTag is tag for Other Magnification Types Available
	OtherMagnificationTypesAvailableTag = 0x201000A7
	// DefaultSmoothingTypeTag is tag for Default Smoothing Type
	DefaultSmoothingTypeTag = 0x201000A8
	// OtherSmoothingTypesAvailableTag is tag for Other Smoothing Types Available
	OtherSmoothingTypesAvailableTag = 0x201000A9
	// BorderDensityTag is tag for Border Density
	BorderDensityTag = 0x20100100
	// EmptyImageDensityTag is tag for Empty Image Density
	EmptyImageDensityTag = 0x20100110
	// MinDensityTag is tag for Min Density
	MinDensityTag = 0x20100120
	// MaxDensityTag is tag for Max Density
	MaxDensityTag = 0x20100130
	// TrimTag is tag for Trim
	TrimTag = 0x20100140
	// ConfigurationInformationTag is tag for Configuration Information
	ConfigurationInformationTag = 0x20100150
	// ConfigurationInformationDescriptionTag is tag for Configuration Information Description
	ConfigurationInformationDescriptionTag = 0x20100152
	// MaximumCollatedFilmsTag is tag for Maximum Collated Films
	MaximumCollatedFilmsTag = 0x20100154
	// IlluminationTag is tag for Illumination
	IlluminationTag = 0x2010015E
	// ReflectedAmbientLightTag is tag for Reflected Ambient Light
	ReflectedAmbientLightTag = 0x20100160
	// PrinterPixelSpacingTag is tag for Printer Pixel Spacing
	PrinterPixelSpacingTag = 0x20100376
	// ReferencedFilmSessionSequenceTag is tag for Referenced Film Session Sequence
	ReferencedFilmSessionSequenceTag = 0x20100500
	// ReferencedImageBoxSequenceTag is tag for Referenced Image Box Sequence
	ReferencedImageBoxSequenceTag = 0x20100510
	// ReferencedBasicAnnotationBoxSequenceTag is tag for Referenced Basic Annotation Box Sequence
	ReferencedBasicAnnotationBoxSequenceTag = 0x20100520
	// ImageBoxPositionTag is tag for Image Box Position
	ImageBoxPositionTag = 0x20200010
	// PolarityTag is tag for Polarity
	PolarityTag = 0x20200020
	// RequestedImageSizeTag is tag for Requested Image Size
	RequestedImageSizeTag = 0x20200030
	// RequestedDecimateCropBehaviorTag is tag for Requested Decimate/Crop Behavior
	RequestedDecimateCropBehaviorTag = 0x20200040
	// RequestedResolutionIDTag is tag for Requested Resolution ID
	RequestedResolutionIDTag = 0x20200050
	// RequestedImageSizeFlagTag is tag for Requested Image Size Flag
	RequestedImageSizeFlagTag = 0x202000A0
	// DecimateCropResultTag is tag for Decimate/Crop Result
	DecimateCropResultTag = 0x202000A2
	// BasicGrayscaleImageSequenceTag is tag for Basic Grayscale Image Sequence
	BasicGrayscaleImageSequenceTag = 0x20200110
	// BasicColorImageSequenceTag is tag for Basic Color Image Sequence
	BasicColorImageSequenceTag = 0x20200111
	// ReferencedImageOverlayBoxSequenceTag is tag for Referenced Image Overlay Box Sequence
	ReferencedImageOverlayBoxSequenceTag = 0x20200130
	// ReferencedVOILUTBoxSequenceTag is tag for Referenced VOI LUT Box Sequence
	ReferencedVOILUTBoxSequenceTag = 0x20200140
	// AnnotationPositionTag is tag for Annotation Position
	AnnotationPositionTag = 0x20300010
	// TextStringTag is tag for Text String
	TextStringTag = 0x20300020
	// ReferencedOverlayPlaneSequenceTag is tag for Referenced Overlay Plane Sequence
	ReferencedOverlayPlaneSequenceTag = 0x20400010
	// ReferencedOverlayPlaneGroupsTag is tag for Referenced Overlay Plane Groups
	ReferencedOverlayPlaneGroupsTag = 0x20400011
	// OverlayPixelDataSequenceTag is tag for Overlay Pixel Data Sequence
	OverlayPixelDataSequenceTag = 0x20400020
	// OverlayMagnificationTypeTag is tag for Overlay Magnification Type
	OverlayMagnificationTypeTag = 0x20400060
	// OverlaySmoothingTypeTag is tag for Overlay Smoothing Type
	OverlaySmoothingTypeTag = 0x20400070
	// OverlayOrImageMagnificationTag is tag for Overlay or Image Magnification
	OverlayOrImageMagnificationTag = 0x20400072
	// MagnifyToNumberOfColumnsTag is tag for Magnify to Number of Columns
	MagnifyToNumberOfColumnsTag = 0x20400074
	// OverlayForegroundDensityTag is tag for Overlay Foreground Density
	OverlayForegroundDensityTag = 0x20400080
	// OverlayBackgroundDensityTag is tag for Overlay Background Density
	OverlayBackgroundDensityTag = 0x20400082
	// OverlayModeTag is tag for Overlay Mode
	OverlayModeTag = 0x20400090
	// ThresholdDensityTag is tag for Threshold Density
	ThresholdDensityTag = 0x20400100
	// ReferencedImageBoxSequenceRetiredTag is tag for Referenced Image Box Sequence (Retired)
	ReferencedImageBoxSequenceRetiredTag = 0x20400500
	// PresentationLUTSequenceTag is tag for Presentation LUT Sequence
	PresentationLUTSequenceTag = 0x20500010
	// PresentationLUTShapeTag is tag for Presentation LUT Shape
	PresentationLUTShapeTag = 0x20500020
	// ReferencedPresentationLUTSequenceTag is tag for Referenced Presentation LUT Sequence
	ReferencedPresentationLUTSequenceTag = 0x20500500
	// PrintJobIDTag is tag for Print Job ID
	PrintJobIDTag = 0x21000010
	// ExecutionStatusTag is tag for Execution Status
	ExecutionStatusTag = 0x21000020
	// ExecutionStatusInfoTag is tag for Execution Status Info
	ExecutionStatusInfoTag = 0x21000030
	// CreationDateTag is tag for Creation Date
	CreationDateTag = 0x21000040
	// CreationTimeTag is tag for Creation Time
	CreationTimeTag = 0x21000050
	// OriginatorTag is tag for Originator
	OriginatorTag = 0x21000070
	// DestinationAETag is tag for Destination AE
	DestinationAETag = 0x21000140
	// OwnerIDTag is tag for Owner ID
	OwnerIDTag = 0x21000160
	// NumberOfFilmsTag is tag for Number of Films
	NumberOfFilmsTag = 0x21000170
	// ReferencedPrintJobSequencePullStoredPrintTag is tag for Referenced Print Job Sequence (Pull Stored Print)
	ReferencedPrintJobSequencePullStoredPrintTag = 0x21000500
	// PrinterStatusTag is tag for Printer Status
	PrinterStatusTag = 0x21100010
	// PrinterStatusInfoTag is tag for Printer Status Info
	PrinterStatusInfoTag = 0x21100020
	// PrinterNameTag is tag for Printer Name
	PrinterNameTag = 0x21100030
	// PrintQueueIDTag is tag for Print Queue ID
	PrintQueueIDTag = 0x21100099
	// QueueStatusTag is tag for Queue Status
	QueueStatusTag = 0x21200010
	// PrintJobDescriptionSequenceTag is tag for Print Job Description Sequence
	PrintJobDescriptionSequenceTag = 0x21200050
	// ReferencedPrintJobSequenceTag is tag for Referenced Print Job Sequence
	ReferencedPrintJobSequenceTag = 0x21200070
	// PrintManagementCapabilitiesSequenceTag is tag for Print Management Capabilities Sequence
	PrintManagementCapabilitiesSequenceTag = 0x21300010
	// PrinterCharacteristicsSequenceTag is tag for Printer Characteristics Sequence
	PrinterCharacteristicsSequenceTag = 0x21300015
	// FilmBoxContentSequenceTag is tag for Film Box Content Sequence
	FilmBoxContentSequenceTag = 0x21300030
	// ImageBoxContentSequenceTag is tag for Image Box Content Sequence
	ImageBoxContentSequenceTag = 0x21300040
	// AnnotationContentSequenceTag is tag for Annotation Content Sequence
	AnnotationContentSequenceTag = 0x21300050
	// ImageOverlayBoxContentSequenceTag is tag for Image Overlay Box Content Sequence
	ImageOverlayBoxContentSequenceTag = 0x21300060
	// PresentationLUTContentSequenceTag is tag for Presentation LUT Content Sequence
	PresentationLUTContentSequenceTag = 0x21300080
	// ProposedStudySequenceTag is tag for Proposed Study Sequence
	ProposedStudySequenceTag = 0x213000A0
	// OriginalImageSequenceTag is tag for Original Image Sequence
	OriginalImageSequenceTag = 0x213000C0
	// LabelUsingInformationExtractedFromInstancesTag is tag for Label Using Information Extracted From Instances
	LabelUsingInformationExtractedFromInstancesTag = 0x22000001
	// LabelTextTag is tag for Label Text
	LabelTextTag = 0x22000002
	// LabelStyleSelectionTag is tag for Label Style Selection
	LabelStyleSelectionTag = 0x22000003
	// MediaDispositionTag is tag for Media Disposition
	MediaDispositionTag = 0x22000004
	// BarcodeValueTag is tag for Barcode Value
	BarcodeValueTag = 0x22000005
	// BarcodeSymbologyTag is tag for Barcode Symbology
	BarcodeSymbologyTag = 0x22000006
	// AllowMediaSplittingTag is tag for Allow Media Splitting
	AllowMediaSplittingTag = 0x22000007
	// IncludeNonDICOMObjectsTag is tag for Include Non-DICOM Objects
	IncludeNonDICOMObjectsTag = 0x22000008
	// IncludeDisplayApplicationTag is tag for Include Display Application
	IncludeDisplayApplicationTag = 0x22000009
	// PreserveCompositeInstancesAfterMediaCreationTag is tag for Preserve Composite Instances After Media Creation
	PreserveCompositeInstancesAfterMediaCreationTag = 0x2200000A
	// TotalNumberOfPiecesOfMediaCreatedTag is tag for Total Number of Pieces of Media Created
	TotalNumberOfPiecesOfMediaCreatedTag = 0x2200000B
	// RequestedMediaApplicationProfileTag is tag for Requested Media Application Profile
	RequestedMediaApplicationProfileTag = 0x2200000C
	// ReferencedStorageMediaSequenceTag is tag for Referenced Storage Media Sequence
	ReferencedStorageMediaSequenceTag = 0x2200000D
	// FailureAttributesTag is tag for Failure Attributes
	FailureAttributesTag = 0x2200000E
	// AllowLossyCompressionTag is tag for Allow Lossy Compression
	AllowLossyCompressionTag = 0x2200000F
	// RequestPriorityTag is tag for Request Priority
	RequestPriorityTag = 0x22000020
	// RTImageLabelTag is tag for RT Image Label
	RTImageLabelTag = 0x30020002
	// RTImageNameTag is tag for RT Image Name
	RTImageNameTag = 0x30020003
	// RTImageDescriptionTag is tag for RT Image Description
	RTImageDescriptionTag = 0x30020004
	// ReportedValuesOriginTag is tag for Reported Values Origin
	ReportedValuesOriginTag = 0x3002000A
	// RTImagePlaneTag is tag for RT Image Plane
	RTImagePlaneTag = 0x3002000C
	// XRayImageReceptorTranslationTag is tag for X-Ray Image Receptor Translation
	XRayImageReceptorTranslationTag = 0x3002000D
	// XRayImageReceptorAngleTag is tag for X-Ray Image Receptor Angle
	XRayImageReceptorAngleTag = 0x3002000E
	// RTImageOrientationTag is tag for RT Image Orientation
	RTImageOrientationTag = 0x30020010
	// ImagePlanePixelSpacingTag is tag for Image Plane Pixel Spacing
	ImagePlanePixelSpacingTag = 0x30020011
	// RTImagePositionTag is tag for RT Image Position
	RTImagePositionTag = 0x30020012
	// RadiationMachineNameTag is tag for Radiation Machine Name
	RadiationMachineNameTag = 0x30020020
	// RadiationMachineSADTag is tag for Radiation Machine SAD
	RadiationMachineSADTag = 0x30020022
	// RadiationMachineSSDTag is tag for Radiation Machine SSD
	RadiationMachineSSDTag = 0x30020024
	// RTImageSIDTag is tag for RT Image SID
	RTImageSIDTag = 0x30020026
	// SourceToReferenceObjectDistanceTag is tag for Source to Reference Object Distance
	SourceToReferenceObjectDistanceTag = 0x30020028
	// FractionNumberTag is tag for Fraction Number
	FractionNumberTag = 0x30020029
	// ExposureSequenceTag is tag for Exposure Sequence
	ExposureSequenceTag = 0x30020030
	// MetersetExposureTag is tag for Meterset Exposure
	MetersetExposureTag = 0x30020032
	// DiaphragmPositionTag is tag for Diaphragm Position
	DiaphragmPositionTag = 0x30020034
	// FluenceMapSequenceTag is tag for Fluence Map Sequence
	FluenceMapSequenceTag = 0x30020040
	// FluenceDataSourceTag is tag for Fluence Data Source
	FluenceDataSourceTag = 0x30020041
	// FluenceDataScaleTag is tag for Fluence Data Scale
	FluenceDataScaleTag = 0x30020042
	// PrimaryFluenceModeSequenceTag is tag for Primary Fluence Mode Sequence
	PrimaryFluenceModeSequenceTag = 0x30020050
	// FluenceModeTag is tag for Fluence Mode
	FluenceModeTag = 0x30020051
	// FluenceModeIDTag is tag for Fluence Mode ID
	FluenceModeIDTag = 0x30020052
	// DVHTypeTag is tag for DVH Type
	DVHTypeTag = 0x30040001
	// DoseUnitsTag is tag for Dose Units
	DoseUnitsTag = 0x30040002
	// DoseTypeTag is tag for Dose Type
	DoseTypeTag = 0x30040004
	// SpatialTransformOfDoseTag is tag for Spatial Transform of Dose
	SpatialTransformOfDoseTag = 0x30040005
	// DoseCommentTag is tag for Dose Comment
	DoseCommentTag = 0x30040006
	// NormalizationPointTag is tag for Normalization Point
	NormalizationPointTag = 0x30040008
	// DoseSummationTypeTag is tag for Dose Summation Type
	DoseSummationTypeTag = 0x3004000A
	// GridFrameOffsetVectorTag is tag for Grid Frame Offset Vector
	GridFrameOffsetVectorTag = 0x3004000C
	// DoseGridScalingTag is tag for Dose Grid Scaling
	DoseGridScalingTag = 0x3004000E
	// RTDoseROISequenceTag is tag for RT Dose ROI Sequence
	RTDoseROISequenceTag = 0x30040010
	// DoseValueTag is tag for Dose Value
	DoseValueTag = 0x30040012
	// TissueHeterogeneityCorrectionTag is tag for Tissue Heterogeneity Correction
	TissueHeterogeneityCorrectionTag = 0x30040014
	// DVHNormalizationPointTag is tag for DVH Normalization Point
	DVHNormalizationPointTag = 0x30040040
	// DVHNormalizationDoseValueTag is tag for DVH Normalization Dose Value
	DVHNormalizationDoseValueTag = 0x30040042
	// DVHSequenceTag is tag for DVH Sequence
	DVHSequenceTag = 0x30040050
	// DVHDoseScalingTag is tag for DVH Dose Scaling
	DVHDoseScalingTag = 0x30040052
	// DVHVolumeUnitsTag is tag for DVH Volume Units
	DVHVolumeUnitsTag = 0x30040054
	// DVHNumberOfBinsTag is tag for DVH Number of Bins
	DVHNumberOfBinsTag = 0x30040056
	// DVHDataTag is tag for DVH Data
	DVHDataTag = 0x30040058
	// DVHReferencedROISequenceTag is tag for DVH Referenced ROI Sequence
	DVHReferencedROISequenceTag = 0x30040060
	// DVHROIContributionTypeTag is tag for DVH ROI Contribution Type
	DVHROIContributionTypeTag = 0x30040062
	// DVHMinimumDoseTag is tag for DVH Minimum Dose
	DVHMinimumDoseTag = 0x30040070
	// DVHMaximumDoseTag is tag for DVH Maximum Dose
	DVHMaximumDoseTag = 0x30040072
	// DVHMeanDoseTag is tag for DVH Mean Dose
	DVHMeanDoseTag = 0x30040074
	// StructureSetLabelTag is tag for Structure Set Label
	StructureSetLabelTag = 0x30060002
	// StructureSetNameTag is tag for Structure Set Name
	StructureSetNameTag = 0x30060004
	// StructureSetDescriptionTag is tag for Structure Set Description
	StructureSetDescriptionTag = 0x30060006
	// StructureSetDateTag is tag for Structure Set Date
	StructureSetDateTag = 0x30060008
	// StructureSetTimeTag is tag for Structure Set Time
	StructureSetTimeTag = 0x30060009
	// ReferencedFrameOfReferenceSequenceTag is tag for Referenced Frame of Reference Sequence
	ReferencedFrameOfReferenceSequenceTag = 0x30060010
	// RTReferencedStudySequenceTag is tag for RT Referenced Study Sequence
	RTReferencedStudySequenceTag = 0x30060012
	// RTReferencedSeriesSequenceTag is tag for RT Referenced Series Sequence
	RTReferencedSeriesSequenceTag = 0x30060014
	// ContourImageSequenceTag is tag for Contour Image Sequence
	ContourImageSequenceTag = 0x30060016
	// PredecessorStructureSetSequenceTag is tag for Predecessor Structure Set Sequence
	PredecessorStructureSetSequenceTag = 0x30060018
	// StructureSetROISequenceTag is tag for Structure Set ROI Sequence
	StructureSetROISequenceTag = 0x30060020
	// ROINumberTag is tag for ROI Number
	ROINumberTag = 0x30060022
	// ReferencedFrameOfReferenceUIDTag is tag for Referenced Frame of Reference UID
	ReferencedFrameOfReferenceUIDTag = 0x30060024
	// ROINameTag is tag for ROI Name
	ROINameTag = 0x30060026
	// ROIDescriptionTag is tag for ROI Description
	ROIDescriptionTag = 0x30060028
	// ROIDisplayColorTag is tag for ROI Display Color
	ROIDisplayColorTag = 0x3006002A
	// ROIVolumeTag is tag for ROI Volume
	ROIVolumeTag = 0x3006002C
	// RTRelatedROISequenceTag is tag for RT Related ROI Sequence
	RTRelatedROISequenceTag = 0x30060030
	// RTROIRelationshipTag is tag for RT ROI Relationship
	RTROIRelationshipTag = 0x30060033
	// ROIGenerationAlgorithmTag is tag for ROI Generation Algorithm
	ROIGenerationAlgorithmTag = 0x30060036
	// ROIDerivationAlgorithmIdentificationSequenceTag is tag for ROI Derivation Algorithm Identification Sequence
	ROIDerivationAlgorithmIdentificationSequenceTag = 0x30060037
	// ROIGenerationDescriptionTag is tag for ROI Generation Description
	ROIGenerationDescriptionTag = 0x30060038
	// ROIContourSequenceTag is tag for ROI Contour Sequence
	ROIContourSequenceTag = 0x30060039
	// ContourSequenceTag is tag for Contour Sequence
	ContourSequenceTag = 0x30060040
	// ContourGeometricTypeTag is tag for Contour Geometric Type
	ContourGeometricTypeTag = 0x30060042
	// ContourSlabThicknessTag is tag for Contour Slab Thickness
	ContourSlabThicknessTag = 0x30060044
	// ContourOffsetVectorTag is tag for Contour Offset Vector
	ContourOffsetVectorTag = 0x30060045
	// NumberOfContourPointsTag is tag for Number of Contour Points
	NumberOfContourPointsTag = 0x30060046
	// ContourNumberTag is tag for Contour Number
	ContourNumberTag = 0x30060048
	// AttachedContoursTag is tag for Attached Contours
	AttachedContoursTag = 0x30060049
	// ContourDataTag is tag for Contour Data
	ContourDataTag = 0x30060050
	// RTROIObservationsSequenceTag is tag for RT ROI Observations Sequence
	RTROIObservationsSequenceTag = 0x30060080
	// ObservationNumberTag is tag for Observation Number
	ObservationNumberTag = 0x30060082
	// ReferencedROINumberTag is tag for Referenced ROI Number
	ReferencedROINumberTag = 0x30060084
	// ROIObservationLabelTag is tag for ROI Observation Label
	ROIObservationLabelTag = 0x30060085
	// RTROIIdentificationCodeSequenceTag is tag for RT ROI Identification Code Sequence
	RTROIIdentificationCodeSequenceTag = 0x30060086
	// ROIObservationDescriptionTag is tag for ROI Observation Description
	ROIObservationDescriptionTag = 0x30060088
	// RelatedRTROIObservationsSequenceTag is tag for Related RT ROI Observations Sequence
	RelatedRTROIObservationsSequenceTag = 0x300600A0
	// RTROIInterpretedTypeTag is tag for RT ROI Interpreted Type
	RTROIInterpretedTypeTag = 0x300600A4
	// ROIInterpreterTag is tag for ROI Interpreter
	ROIInterpreterTag = 0x300600A6
	// ROIPhysicalPropertiesSequenceTag is tag for ROI Physical Properties Sequence
	ROIPhysicalPropertiesSequenceTag = 0x300600B0
	// ROIPhysicalPropertyTag is tag for ROI Physical Property
	ROIPhysicalPropertyTag = 0x300600B2
	// ROIPhysicalPropertyValueTag is tag for ROI Physical Property Value
	ROIPhysicalPropertyValueTag = 0x300600B4
	// ROIElementalCompositionSequenceTag is tag for ROI Elemental Composition Sequence
	ROIElementalCompositionSequenceTag = 0x300600B6
	// ROIElementalCompositionAtomicNumberTag is tag for ROI Elemental Composition Atomic Number
	ROIElementalCompositionAtomicNumberTag = 0x300600B7
	// ROIElementalCompositionAtomicMassFractionTag is tag for ROI Elemental Composition Atomic Mass Fraction
	ROIElementalCompositionAtomicMassFractionTag = 0x300600B8
	// AdditionalRTROIIdentificationCodeSequenceTag is tag for Additional RT ROI Identification Code Sequence
	AdditionalRTROIIdentificationCodeSequenceTag = 0x300600B9
	// FrameOfReferenceRelationshipSequenceTag is tag for Frame of Reference Relationship Sequence
	FrameOfReferenceRelationshipSequenceTag = 0x300600C0
	// RelatedFrameOfReferenceUIDTag is tag for Related Frame of Reference UID
	RelatedFrameOfReferenceUIDTag = 0x300600C2
	// FrameOfReferenceTransformationTypeTag is tag for Frame of Reference Transformation Type
	FrameOfReferenceTransformationTypeTag = 0x300600C4
	// FrameOfReferenceTransformationMatrixTag is tag for Frame of Reference Transformation Matrix
	FrameOfReferenceTransformationMatrixTag = 0x300600C6
	// FrameOfReferenceTransformationCommentTag is tag for Frame of Reference Transformation Comment
	FrameOfReferenceTransformationCommentTag = 0x300600C8
	// PatientLocationCoordinatesSequenceTag is tag for Patient Location Coordinates Sequence
	PatientLocationCoordinatesSequenceTag = 0x300600C9
	// PatientLocationCoordinatesCodeSequenceTag is tag for Patient Location Coordinates Code Sequence
	PatientLocationCoordinatesCodeSequenceTag = 0x300600CA
	// PatientSupportPositionSequenceTag is tag for Patient Support Position Sequence
	PatientSupportPositionSequenceTag = 0x300600CB
	// MeasuredDoseReferenceSequenceTag is tag for Measured Dose Reference Sequence
	MeasuredDoseReferenceSequenceTag = 0x30080010
	// MeasuredDoseDescriptionTag is tag for Measured Dose Description
	MeasuredDoseDescriptionTag = 0x30080012
	// MeasuredDoseTypeTag is tag for Measured Dose Type
	MeasuredDoseTypeTag = 0x30080014
	// MeasuredDoseValueTag is tag for Measured Dose Value
	MeasuredDoseValueTag = 0x30080016
	// TreatmentSessionBeamSequenceTag is tag for Treatment Session Beam Sequence
	TreatmentSessionBeamSequenceTag = 0x30080020
	// TreatmentSessionIonBeamSequenceTag is tag for Treatment Session Ion Beam Sequence
	TreatmentSessionIonBeamSequenceTag = 0x30080021
	// CurrentFractionNumberTag is tag for Current Fraction Number
	CurrentFractionNumberTag = 0x30080022
	// TreatmentControlPointDateTag is tag for Treatment Control Point Date
	TreatmentControlPointDateTag = 0x30080024
	// TreatmentControlPointTimeTag is tag for Treatment Control Point Time
	TreatmentControlPointTimeTag = 0x30080025
	// TreatmentTerminationStatusTag is tag for Treatment Termination Status
	TreatmentTerminationStatusTag = 0x3008002A
	// TreatmentTerminationCodeTag is tag for Treatment Termination Code
	TreatmentTerminationCodeTag = 0x3008002B
	// TreatmentVerificationStatusTag is tag for Treatment Verification Status
	TreatmentVerificationStatusTag = 0x3008002C
	// ReferencedTreatmentRecordSequenceTag is tag for Referenced Treatment Record Sequence
	ReferencedTreatmentRecordSequenceTag = 0x30080030
	// SpecifiedPrimaryMetersetTag is tag for Specified Primary Meterset
	SpecifiedPrimaryMetersetTag = 0x30080032
	// SpecifiedSecondaryMetersetTag is tag for Specified Secondary Meterset
	SpecifiedSecondaryMetersetTag = 0x30080033
	// DeliveredPrimaryMetersetTag is tag for Delivered Primary Meterset
	DeliveredPrimaryMetersetTag = 0x30080036
	// DeliveredSecondaryMetersetTag is tag for Delivered Secondary Meterset
	DeliveredSecondaryMetersetTag = 0x30080037
	// SpecifiedTreatmentTimeTag is tag for Specified Treatment Time
	SpecifiedTreatmentTimeTag = 0x3008003A
	// DeliveredTreatmentTimeTag is tag for Delivered Treatment Time
	DeliveredTreatmentTimeTag = 0x3008003B
	// ControlPointDeliverySequenceTag is tag for Control Point Delivery Sequence
	ControlPointDeliverySequenceTag = 0x30080040
	// IonControlPointDeliverySequenceTag is tag for Ion Control Point Delivery Sequence
	IonControlPointDeliverySequenceTag = 0x30080041
	// SpecifiedMetersetTag is tag for Specified Meterset
	SpecifiedMetersetTag = 0x30080042
	// DeliveredMetersetTag is tag for Delivered Meterset
	DeliveredMetersetTag = 0x30080044
	// MetersetRateSetTag is tag for Meterset Rate Set
	MetersetRateSetTag = 0x30080045
	// MetersetRateDeliveredTag is tag for Meterset Rate Delivered
	MetersetRateDeliveredTag = 0x30080046
	// ScanSpotMetersetsDeliveredTag is tag for Scan Spot Metersets Delivered
	ScanSpotMetersetsDeliveredTag = 0x30080047
	// DoseRateDeliveredTag is tag for Dose Rate Delivered
	DoseRateDeliveredTag = 0x30080048
	// TreatmentSummaryCalculatedDoseReferenceSequenceTag is tag for Treatment Summary Calculated Dose Reference Sequence
	TreatmentSummaryCalculatedDoseReferenceSequenceTag = 0x30080050
	// CumulativeDoseToDoseReferenceTag is tag for Cumulative Dose to Dose Reference
	CumulativeDoseToDoseReferenceTag = 0x30080052
	// FirstTreatmentDateTag is tag for First Treatment Date
	FirstTreatmentDateTag = 0x30080054
	// MostRecentTreatmentDateTag is tag for Most Recent Treatment Date
	MostRecentTreatmentDateTag = 0x30080056
	// NumberOfFractionsDeliveredTag is tag for Number of Fractions Delivered
	NumberOfFractionsDeliveredTag = 0x3008005A
	// OverrideSequenceTag is tag for Override Sequence
	OverrideSequenceTag = 0x30080060
	// ParameterSequencePointerTag is tag for Parameter Sequence Pointer
	ParameterSequencePointerTag = 0x30080061
	// OverrideParameterPointerTag is tag for Override Parameter Pointer
	OverrideParameterPointerTag = 0x30080062
	// ParameterItemIndexTag is tag for Parameter Item Index
	ParameterItemIndexTag = 0x30080063
	// MeasuredDoseReferenceNumberTag is tag for Measured Dose Reference Number
	MeasuredDoseReferenceNumberTag = 0x30080064
	// ParameterPointerTag is tag for Parameter Pointer
	ParameterPointerTag = 0x30080065
	// OverrideReasonTag is tag for Override Reason
	OverrideReasonTag = 0x30080066
	// ParameterValueNumberTag is tag for Parameter Value Number
	ParameterValueNumberTag = 0x30080067
	// CorrectedParameterSequenceTag is tag for Corrected Parameter Sequence
	CorrectedParameterSequenceTag = 0x30080068
	// CorrectionValueTag is tag for Correction Value
	CorrectionValueTag = 0x3008006A
	// CalculatedDoseReferenceSequenceTag is tag for Calculated Dose Reference Sequence
	CalculatedDoseReferenceSequenceTag = 0x30080070
	// CalculatedDoseReferenceNumberTag is tag for Calculated Dose Reference Number
	CalculatedDoseReferenceNumberTag = 0x30080072
	// CalculatedDoseReferenceDescriptionTag is tag for Calculated Dose Reference Description
	CalculatedDoseReferenceDescriptionTag = 0x30080074
	// CalculatedDoseReferenceDoseValueTag is tag for Calculated Dose Reference Dose Value
	CalculatedDoseReferenceDoseValueTag = 0x30080076
	// StartMetersetTag is tag for Start Meterset
	StartMetersetTag = 0x30080078
	// EndMetersetTag is tag for End Meterset
	EndMetersetTag = 0x3008007A
	// ReferencedMeasuredDoseReferenceSequenceTag is tag for Referenced Measured Dose Reference Sequence
	ReferencedMeasuredDoseReferenceSequenceTag = 0x30080080
	// ReferencedMeasuredDoseReferenceNumberTag is tag for Referenced Measured Dose Reference Number
	ReferencedMeasuredDoseReferenceNumberTag = 0x30080082
	// ReferencedCalculatedDoseReferenceSequenceTag is tag for Referenced Calculated Dose Reference Sequence
	ReferencedCalculatedDoseReferenceSequenceTag = 0x30080090
	// ReferencedCalculatedDoseReferenceNumberTag is tag for Referenced Calculated Dose Reference Number
	ReferencedCalculatedDoseReferenceNumberTag = 0x30080092
	// BeamLimitingDeviceLeafPairsSequenceTag is tag for Beam Limiting Device Leaf Pairs Sequence
	BeamLimitingDeviceLeafPairsSequenceTag = 0x300800A0
	// RecordedWedgeSequenceTag is tag for Recorded Wedge Sequence
	RecordedWedgeSequenceTag = 0x300800B0
	// RecordedCompensatorSequenceTag is tag for Recorded Compensator Sequence
	RecordedCompensatorSequenceTag = 0x300800C0
	// RecordedBlockSequenceTag is tag for Recorded Block Sequence
	RecordedBlockSequenceTag = 0x300800D0
	// TreatmentSummaryMeasuredDoseReferenceSequenceTag is tag for Treatment Summary Measured Dose Reference Sequence
	TreatmentSummaryMeasuredDoseReferenceSequenceTag = 0x300800E0
	// RecordedSnoutSequenceTag is tag for Recorded Snout Sequence
	RecordedSnoutSequenceTag = 0x300800F0
	// RecordedRangeShifterSequenceTag is tag for Recorded Range Shifter Sequence
	RecordedRangeShifterSequenceTag = 0x300800F2
	// RecordedLateralSpreadingDeviceSequenceTag is tag for Recorded Lateral Spreading Device Sequence
	RecordedLateralSpreadingDeviceSequenceTag = 0x300800F4
	// RecordedRangeModulatorSequenceTag is tag for Recorded Range Modulator Sequence
	RecordedRangeModulatorSequenceTag = 0x300800F6
	// RecordedSourceSequenceTag is tag for Recorded Source Sequence
	RecordedSourceSequenceTag = 0x30080100
	// SourceSerialNumberTag is tag for Source Serial Number
	SourceSerialNumberTag = 0x30080105
	// TreatmentSessionApplicationSetupSequenceTag is tag for Treatment Session Application Setup Sequence
	TreatmentSessionApplicationSetupSequenceTag = 0x30080110
	// ApplicationSetupCheckTag is tag for Application Setup Check
	ApplicationSetupCheckTag = 0x30080116
	// RecordedBrachyAccessoryDeviceSequenceTag is tag for Recorded Brachy Accessory Device Sequence
	RecordedBrachyAccessoryDeviceSequenceTag = 0x30080120
	// ReferencedBrachyAccessoryDeviceNumberTag is tag for Referenced Brachy Accessory Device Number
	ReferencedBrachyAccessoryDeviceNumberTag = 0x30080122
	// RecordedChannelSequenceTag is tag for Recorded Channel Sequence
	RecordedChannelSequenceTag = 0x30080130
	// SpecifiedChannelTotalTimeTag is tag for Specified Channel Total Time
	SpecifiedChannelTotalTimeTag = 0x30080132
	// DeliveredChannelTotalTimeTag is tag for Delivered Channel Total Time
	DeliveredChannelTotalTimeTag = 0x30080134
	// SpecifiedNumberOfPulsesTag is tag for Specified Number of Pulses
	SpecifiedNumberOfPulsesTag = 0x30080136
	// DeliveredNumberOfPulsesTag is tag for Delivered Number of Pulses
	DeliveredNumberOfPulsesTag = 0x30080138
	// SpecifiedPulseRepetitionIntervalTag is tag for Specified Pulse Repetition Interval
	SpecifiedPulseRepetitionIntervalTag = 0x3008013A
	// DeliveredPulseRepetitionIntervalTag is tag for Delivered Pulse Repetition Interval
	DeliveredPulseRepetitionIntervalTag = 0x3008013C
	// RecordedSourceApplicatorSequenceTag is tag for Recorded Source Applicator Sequence
	RecordedSourceApplicatorSequenceTag = 0x30080140
	// ReferencedSourceApplicatorNumberTag is tag for Referenced Source Applicator Number
	ReferencedSourceApplicatorNumberTag = 0x30080142
	// RecordedChannelShieldSequenceTag is tag for Recorded Channel Shield Sequence
	RecordedChannelShieldSequenceTag = 0x30080150
	// ReferencedChannelShieldNumberTag is tag for Referenced Channel Shield Number
	ReferencedChannelShieldNumberTag = 0x30080152
	// BrachyControlPointDeliveredSequenceTag is tag for Brachy Control Point Delivered Sequence
	BrachyControlPointDeliveredSequenceTag = 0x30080160
	// SafePositionExitDateTag is tag for Safe Position Exit Date
	SafePositionExitDateTag = 0x30080162
	// SafePositionExitTimeTag is tag for Safe Position Exit Time
	SafePositionExitTimeTag = 0x30080164
	// SafePositionReturnDateTag is tag for Safe Position Return Date
	SafePositionReturnDateTag = 0x30080166
	// SafePositionReturnTimeTag is tag for Safe Position Return Time
	SafePositionReturnTimeTag = 0x30080168
	// PulseSpecificBrachyControlPointDeliveredSequenceTag is tag for Pulse Specific Brachy Control Point Delivered Sequence
	PulseSpecificBrachyControlPointDeliveredSequenceTag = 0x30080171
	// PulseNumberTag is tag for Pulse Number
	PulseNumberTag = 0x30080172
	// BrachyPulseControlPointDeliveredSequenceTag is tag for Brachy Pulse Control Point Delivered Sequence
	BrachyPulseControlPointDeliveredSequenceTag = 0x30080173
	// CurrentTreatmentStatusTag is tag for Current Treatment Status
	CurrentTreatmentStatusTag = 0x30080200
	// TreatmentStatusCommentTag is tag for Treatment Status Comment
	TreatmentStatusCommentTag = 0x30080202
	// FractionGroupSummarySequenceTag is tag for Fraction Group Summary Sequence
	FractionGroupSummarySequenceTag = 0x30080220
	// ReferencedFractionNumberTag is tag for Referenced Fraction Number
	ReferencedFractionNumberTag = 0x30080223
	// FractionGroupTypeTag is tag for Fraction Group Type
	FractionGroupTypeTag = 0x30080224
	// BeamStopperPositionTag is tag for Beam Stopper Position
	BeamStopperPositionTag = 0x30080230
	// FractionStatusSummarySequenceTag is tag for Fraction Status Summary Sequence
	FractionStatusSummarySequenceTag = 0x30080240
	// TreatmentDateTag is tag for Treatment Date
	TreatmentDateTag = 0x30080250
	// TreatmentTimeTag is tag for Treatment Time
	TreatmentTimeTag = 0x30080251
	// RTPlanLabelTag is tag for RT Plan Label
	RTPlanLabelTag = 0x300A0002
	// RTPlanNameTag is tag for RT Plan Name
	RTPlanNameTag = 0x300A0003
	// RTPlanDescriptionTag is tag for RT Plan Description
	RTPlanDescriptionTag = 0x300A0004
	// RTPlanDateTag is tag for RT Plan Date
	RTPlanDateTag = 0x300A0006
	// RTPlanTimeTag is tag for RT Plan Time
	RTPlanTimeTag = 0x300A0007
	// TreatmentProtocolsTag is tag for Treatment Protocols
	TreatmentProtocolsTag = 0x300A0009
	// PlanIntentTag is tag for Plan Intent
	PlanIntentTag = 0x300A000A
	// TreatmentSitesTag is tag for Treatment Sites
	TreatmentSitesTag = 0x300A000B
	// RTPlanGeometryTag is tag for RT Plan Geometry
	RTPlanGeometryTag = 0x300A000C
	// PrescriptionDescriptionTag is tag for Prescription Description
	PrescriptionDescriptionTag = 0x300A000E
	// DoseReferenceSequenceTag is tag for Dose Reference Sequence
	DoseReferenceSequenceTag = 0x300A0010
	// DoseReferenceNumberTag is tag for Dose Reference Number
	DoseReferenceNumberTag = 0x300A0012
	// DoseReferenceUIDTag is tag for Dose Reference UID
	DoseReferenceUIDTag = 0x300A0013
	// DoseReferenceStructureTypeTag is tag for Dose Reference Structure Type
	DoseReferenceStructureTypeTag = 0x300A0014
	// NominalBeamEnergyUnitTag is tag for Nominal Beam Energy Unit
	NominalBeamEnergyUnitTag = 0x300A0015
	// DoseReferenceDescriptionTag is tag for Dose Reference Description
	DoseReferenceDescriptionTag = 0x300A0016
	// DoseReferencePointCoordinatesTag is tag for Dose Reference Point Coordinates
	DoseReferencePointCoordinatesTag = 0x300A0018
	// NominalPriorDoseTag is tag for Nominal Prior Dose
	NominalPriorDoseTag = 0x300A001A
	// DoseReferenceTypeTag is tag for Dose Reference Type
	DoseReferenceTypeTag = 0x300A0020
	// ConstraintWeightTag is tag for Constraint Weight
	ConstraintWeightTag = 0x300A0021
	// DeliveryWarningDoseTag is tag for Delivery Warning Dose
	DeliveryWarningDoseTag = 0x300A0022
	// DeliveryMaximumDoseTag is tag for Delivery Maximum Dose
	DeliveryMaximumDoseTag = 0x300A0023
	// TargetMinimumDoseTag is tag for Target Minimum Dose
	TargetMinimumDoseTag = 0x300A0025
	// TargetPrescriptionDoseTag is tag for Target Prescription Dose
	TargetPrescriptionDoseTag = 0x300A0026
	// TargetMaximumDoseTag is tag for Target Maximum Dose
	TargetMaximumDoseTag = 0x300A0027
	// TargetUnderdoseVolumeFractionTag is tag for Target Underdose Volume Fraction
	TargetUnderdoseVolumeFractionTag = 0x300A0028
	// OrganAtRiskFullVolumeDoseTag is tag for Organ at Risk Full-volume Dose
	OrganAtRiskFullVolumeDoseTag = 0x300A002A
	// OrganAtRiskLimitDoseTag is tag for Organ at Risk Limit Dose
	OrganAtRiskLimitDoseTag = 0x300A002B
	// OrganAtRiskMaximumDoseTag is tag for Organ at Risk Maximum Dose
	OrganAtRiskMaximumDoseTag = 0x300A002C
	// OrganAtRiskOverdoseVolumeFractionTag is tag for Organ at Risk Overdose Volume Fraction
	OrganAtRiskOverdoseVolumeFractionTag = 0x300A002D
	// ToleranceTableSequenceTag is tag for Tolerance Table Sequence
	ToleranceTableSequenceTag = 0x300A0040
	// ToleranceTableNumberTag is tag for Tolerance Table Number
	ToleranceTableNumberTag = 0x300A0042
	// ToleranceTableLabelTag is tag for Tolerance Table Label
	ToleranceTableLabelTag = 0x300A0043
	// GantryAngleToleranceTag is tag for Gantry Angle Tolerance
	GantryAngleToleranceTag = 0x300A0044
	// BeamLimitingDeviceAngleToleranceTag is tag for Beam Limiting Device Angle Tolerance
	BeamLimitingDeviceAngleToleranceTag = 0x300A0046
	// BeamLimitingDeviceToleranceSequenceTag is tag for Beam Limiting Device Tolerance Sequence
	BeamLimitingDeviceToleranceSequenceTag = 0x300A0048
	// BeamLimitingDevicePositionToleranceTag is tag for Beam Limiting Device Position Tolerance
	BeamLimitingDevicePositionToleranceTag = 0x300A004A
	// SnoutPositionToleranceTag is tag for Snout Position Tolerance
	SnoutPositionToleranceTag = 0x300A004B
	// PatientSupportAngleToleranceTag is tag for Patient Support Angle Tolerance
	PatientSupportAngleToleranceTag = 0x300A004C
	// TableTopEccentricAngleToleranceTag is tag for Table Top Eccentric Angle Tolerance
	TableTopEccentricAngleToleranceTag = 0x300A004E
	// TableTopPitchAngleToleranceTag is tag for Table Top Pitch Angle Tolerance
	TableTopPitchAngleToleranceTag = 0x300A004F
	// TableTopRollAngleToleranceTag is tag for Table Top Roll Angle Tolerance
	TableTopRollAngleToleranceTag = 0x300A0050
	// TableTopVerticalPositionToleranceTag is tag for Table Top Vertical Position Tolerance
	TableTopVerticalPositionToleranceTag = 0x300A0051
	// TableTopLongitudinalPositionToleranceTag is tag for Table Top Longitudinal Position Tolerance
	TableTopLongitudinalPositionToleranceTag = 0x300A0052
	// TableTopLateralPositionToleranceTag is tag for Table Top Lateral Position Tolerance
	TableTopLateralPositionToleranceTag = 0x300A0053
	// RTPlanRelationshipTag is tag for RT Plan Relationship
	RTPlanRelationshipTag = 0x300A0055
	// FractionGroupSequenceTag is tag for Fraction Group Sequence
	FractionGroupSequenceTag = 0x300A0070
	// FractionGroupNumberTag is tag for Fraction Group Number
	FractionGroupNumberTag = 0x300A0071
	// FractionGroupDescriptionTag is tag for Fraction Group Description
	FractionGroupDescriptionTag = 0x300A0072
	// NumberOfFractionsPlannedTag is tag for Number of Fractions Planned
	NumberOfFractionsPlannedTag = 0x300A0078
	// NumberOfFractionPatternDigitsPerDayTag is tag for Number of Fraction Pattern Digits Per Day
	NumberOfFractionPatternDigitsPerDayTag = 0x300A0079
	// RepeatFractionCycleLengthTag is tag for Repeat Fraction Cycle Length
	RepeatFractionCycleLengthTag = 0x300A007A
	// FractionPatternTag is tag for Fraction Pattern
	FractionPatternTag = 0x300A007B
	// NumberOfBeamsTag is tag for Number of Beams
	NumberOfBeamsTag = 0x300A0080
	// BeamDoseSpecificationPointTag is tag for Beam Dose Specification Point
	BeamDoseSpecificationPointTag = 0x300A0082
	// ReferencedDoseReferenceUIDTag is tag for Referenced Dose Reference UID
	ReferencedDoseReferenceUIDTag = 0x300A0083
	// BeamDoseTag is tag for Beam Dose
	BeamDoseTag = 0x300A0084
	// BeamMetersetTag is tag for Beam Meterset
	BeamMetersetTag = 0x300A0086
	// BeamDosePointDepthTag is tag for Beam Dose Point Depth
	BeamDosePointDepthTag = 0x300A0088
	// BeamDosePointEquivalentDepthTag is tag for Beam Dose Point Equivalent Depth
	BeamDosePointEquivalentDepthTag = 0x300A0089
	// BeamDosePointSSDTag is tag for Beam Dose Point SSD
	BeamDosePointSSDTag = 0x300A008A
	// BeamDoseMeaningTag is tag for Beam Dose Meaning
	BeamDoseMeaningTag = 0x300A008B
	// BeamDoseVerificationControlPointSequenceTag is tag for Beam Dose Verification Control Point Sequence
	BeamDoseVerificationControlPointSequenceTag = 0x300A008C
	// AverageBeamDosePointDepthTag is tag for Average Beam Dose Point Depth
	AverageBeamDosePointDepthTag = 0x300A008D
	// AverageBeamDosePointEquivalentDepthTag is tag for Average Beam Dose Point Equivalent Depth
	AverageBeamDosePointEquivalentDepthTag = 0x300A008E
	// AverageBeamDosePointSSDTag is tag for Average Beam Dose Point SSD
	AverageBeamDosePointSSDTag = 0x300A008F
	// BeamDoseTypeTag is tag for Beam Dose Type
	BeamDoseTypeTag = 0x300A0090
	// AlternateBeamDoseTag is tag for Alternate Beam Dose
	AlternateBeamDoseTag = 0x300A0091
	// AlternateBeamDoseTypeTag is tag for Alternate Beam Dose Type
	AlternateBeamDoseTypeTag = 0x300A0092
	// DepthValueAveragingFlagTag is tag for Depth Value Averaging Flag
	DepthValueAveragingFlagTag = 0x300A0093
	// BeamDosePointSourceToExternalContourDistanceTag is tag for Beam Dose Point Source to External Contour Distance
	BeamDosePointSourceToExternalContourDistanceTag = 0x300A0094
	// NumberOfBrachyApplicationSetupsTag is tag for Number of Brachy Application Setups
	NumberOfBrachyApplicationSetupsTag = 0x300A00A0
	// BrachyApplicationSetupDoseSpecificationPointTag is tag for Brachy Application Setup Dose Specification Point
	BrachyApplicationSetupDoseSpecificationPointTag = 0x300A00A2
	// BrachyApplicationSetupDoseTag is tag for Brachy Application Setup Dose
	BrachyApplicationSetupDoseTag = 0x300A00A4
	// BeamSequenceTag is tag for Beam Sequence
	BeamSequenceTag = 0x300A00B0
	// TreatmentMachineNameTag is tag for Treatment Machine Name
	TreatmentMachineNameTag = 0x300A00B2
	// PrimaryDosimeterUnitTag is tag for Primary Dosimeter Unit
	PrimaryDosimeterUnitTag = 0x300A00B3
	// SourceAxisDistanceTag is tag for Source-Axis Distance
	SourceAxisDistanceTag = 0x300A00B4
	// BeamLimitingDeviceSequenceTag is tag for Beam Limiting Device Sequence
	BeamLimitingDeviceSequenceTag = 0x300A00B6
	// RTBeamLimitingDeviceTypeTag is tag for RT Beam Limiting Device Type
	RTBeamLimitingDeviceTypeTag = 0x300A00B8
	// SourceToBeamLimitingDeviceDistanceTag is tag for Source to Beam Limiting Device Distance
	SourceToBeamLimitingDeviceDistanceTag = 0x300A00BA
	// IsocenterToBeamLimitingDeviceDistanceTag is tag for Isocenter to Beam Limiting Device Distance
	IsocenterToBeamLimitingDeviceDistanceTag = 0x300A00BB
	// NumberOfLeafJawPairsTag is tag for Number of Leaf/Jaw Pairs
	NumberOfLeafJawPairsTag = 0x300A00BC
	// LeafPositionBoundariesTag is tag for Leaf Position Boundaries
	LeafPositionBoundariesTag = 0x300A00BE
	// BeamNumberTag is tag for Beam Number
	BeamNumberTag = 0x300A00C0
	// BeamNameTag is tag for Beam Name
	BeamNameTag = 0x300A00C2
	// BeamDescriptionTag is tag for Beam Description
	BeamDescriptionTag = 0x300A00C3
	// BeamTypeTag is tag for Beam Type
	BeamTypeTag = 0x300A00C4
	// BeamDeliveryDurationLimitTag is tag for Beam Delivery Duration Limit
	BeamDeliveryDurationLimitTag = 0x300A00C5
	// RadiationTypeTag is tag for Radiation Type
	RadiationTypeTag = 0x300A00C6
	// HighDoseTechniqueTypeTag is tag for High-Dose Technique Type
	HighDoseTechniqueTypeTag = 0x300A00C7
	// ReferenceImageNumberTag is tag for Reference Image Number
	ReferenceImageNumberTag = 0x300A00C8
	// PlannedVerificationImageSequenceTag is tag for Planned Verification Image Sequence
	PlannedVerificationImageSequenceTag = 0x300A00CA
	// ImagingDeviceSpecificAcquisitionParametersTag is tag for Imaging Device-Specific Acquisition Parameters
	ImagingDeviceSpecificAcquisitionParametersTag = 0x300A00CC
	// TreatmentDeliveryTypeTag is tag for Treatment Delivery Type
	TreatmentDeliveryTypeTag = 0x300A00CE
	// NumberOfWedgesTag is tag for Number of Wedges
	NumberOfWedgesTag = 0x300A00D0
	// WedgeSequenceTag is tag for Wedge Sequence
	WedgeSequenceTag = 0x300A00D1
	// WedgeNumberTag is tag for Wedge Number
	WedgeNumberTag = 0x300A00D2
	// WedgeTypeTag is tag for Wedge Type
	WedgeTypeTag = 0x300A00D3
	// WedgeIDTag is tag for Wedge ID
	WedgeIDTag = 0x300A00D4
	// WedgeAngleTag is tag for Wedge Angle
	WedgeAngleTag = 0x300A00D5
	// WedgeFactorTag is tag for Wedge Factor
	WedgeFactorTag = 0x300A00D6
	// TotalWedgeTrayWaterEquivalentThicknessTag is tag for Total Wedge Tray Water-Equivalent Thickness
	TotalWedgeTrayWaterEquivalentThicknessTag = 0x300A00D7
	// WedgeOrientationTag is tag for Wedge Orientation
	WedgeOrientationTag = 0x300A00D8
	// IsocenterToWedgeTrayDistanceTag is tag for Isocenter to Wedge Tray Distance
	IsocenterToWedgeTrayDistanceTag = 0x300A00D9
	// SourceToWedgeTrayDistanceTag is tag for Source to Wedge Tray Distance
	SourceToWedgeTrayDistanceTag = 0x300A00DA
	// WedgeThinEdgePositionTag is tag for Wedge Thin Edge Position
	WedgeThinEdgePositionTag = 0x300A00DB
	// BolusIDTag is tag for Bolus ID
	BolusIDTag = 0x300A00DC
	// BolusDescriptionTag is tag for Bolus Description
	BolusDescriptionTag = 0x300A00DD
	// EffectiveWedgeAngleTag is tag for Effective Wedge Angle
	EffectiveWedgeAngleTag = 0x300A00DE
	// NumberOfCompensatorsTag is tag for Number of Compensators
	NumberOfCompensatorsTag = 0x300A00E0
	// MaterialIDTag is tag for Material ID
	MaterialIDTag = 0x300A00E1
	// TotalCompensatorTrayFactorTag is tag for Total Compensator Tray Factor
	TotalCompensatorTrayFactorTag = 0x300A00E2
	// CompensatorSequenceTag is tag for Compensator Sequence
	CompensatorSequenceTag = 0x300A00E3
	// CompensatorNumberTag is tag for Compensator Number
	CompensatorNumberTag = 0x300A00E4
	// CompensatorIDTag is tag for Compensator ID
	CompensatorIDTag = 0x300A00E5
	// SourceToCompensatorTrayDistanceTag is tag for Source to Compensator Tray Distance
	SourceToCompensatorTrayDistanceTag = 0x300A00E6
	// CompensatorRowsTag is tag for Compensator Rows
	CompensatorRowsTag = 0x300A00E7
	// CompensatorColumnsTag is tag for Compensator Columns
	CompensatorColumnsTag = 0x300A00E8
	// CompensatorPixelSpacingTag is tag for Compensator Pixel Spacing
	CompensatorPixelSpacingTag = 0x300A00E9
	// CompensatorPositionTag is tag for Compensator Position
	CompensatorPositionTag = 0x300A00EA
	// CompensatorTransmissionDataTag is tag for Compensator Transmission Data
	CompensatorTransmissionDataTag = 0x300A00EB
	// CompensatorThicknessDataTag is tag for Compensator Thickness Data
	CompensatorThicknessDataTag = 0x300A00EC
	// NumberOfBoliTag is tag for Number of Boli
	NumberOfBoliTag = 0x300A00ED
	// CompensatorTypeTag is tag for Compensator Type
	CompensatorTypeTag = 0x300A00EE
	// CompensatorTrayIDTag is tag for Compensator Tray ID
	CompensatorTrayIDTag = 0x300A00EF
	// NumberOfBlocksTag is tag for Number of Blocks
	NumberOfBlocksTag = 0x300A00F0
	// TotalBlockTrayFactorTag is tag for Total Block Tray Factor
	TotalBlockTrayFactorTag = 0x300A00F2
	// TotalBlockTrayWaterEquivalentThicknessTag is tag for Total Block Tray Water-Equivalent Thickness
	TotalBlockTrayWaterEquivalentThicknessTag = 0x300A00F3
	// BlockSequenceTag is tag for Block Sequence
	BlockSequenceTag = 0x300A00F4
	// BlockTrayIDTag is tag for Block Tray ID
	BlockTrayIDTag = 0x300A00F5
	// SourceToBlockTrayDistanceTag is tag for Source to Block Tray Distance
	SourceToBlockTrayDistanceTag = 0x300A00F6
	// IsocenterToBlockTrayDistanceTag is tag for Isocenter to Block Tray Distance
	IsocenterToBlockTrayDistanceTag = 0x300A00F7
	// BlockTypeTag is tag for Block Type
	BlockTypeTag = 0x300A00F8
	// AccessoryCodeTag is tag for Accessory Code
	AccessoryCodeTag = 0x300A00F9
	// BlockDivergenceTag is tag for Block Divergence
	BlockDivergenceTag = 0x300A00FA
	// BlockMountingPositionTag is tag for Block Mounting Position
	BlockMountingPositionTag = 0x300A00FB
	// BlockNumberTag is tag for Block Number
	BlockNumberTag = 0x300A00FC
	// BlockNameTag is tag for Block Name
	BlockNameTag = 0x300A00FE
	// BlockThicknessTag is tag for Block Thickness
	BlockThicknessTag = 0x300A0100
	// BlockTransmissionTag is tag for Block Transmission
	BlockTransmissionTag = 0x300A0102
	// BlockNumberOfPointsTag is tag for Block Number of Points
	BlockNumberOfPointsTag = 0x300A0104
	// BlockDataTag is tag for Block Data
	BlockDataTag = 0x300A0106
	// ApplicatorSequenceTag is tag for Applicator Sequence
	ApplicatorSequenceTag = 0x300A0107
	// ApplicatorIDTag is tag for Applicator ID
	ApplicatorIDTag = 0x300A0108
	// ApplicatorTypeTag is tag for Applicator Type
	ApplicatorTypeTag = 0x300A0109
	// ApplicatorDescriptionTag is tag for Applicator Description
	ApplicatorDescriptionTag = 0x300A010A
	// CumulativeDoseReferenceCoefficientTag is tag for Cumulative Dose Reference Coefficient
	CumulativeDoseReferenceCoefficientTag = 0x300A010C
	// FinalCumulativeMetersetWeightTag is tag for Final Cumulative Meterset Weight
	FinalCumulativeMetersetWeightTag = 0x300A010E
	// NumberOfControlPointsTag is tag for Number of Control Points
	NumberOfControlPointsTag = 0x300A0110
	// ControlPointSequenceTag is tag for Control Point Sequence
	ControlPointSequenceTag = 0x300A0111
	// ControlPointIndexTag is tag for Control Point Index
	ControlPointIndexTag = 0x300A0112
	// NominalBeamEnergyTag is tag for Nominal Beam Energy
	NominalBeamEnergyTag = 0x300A0114
	// DoseRateSetTag is tag for Dose Rate Set
	DoseRateSetTag = 0x300A0115
	// WedgePositionSequenceTag is tag for Wedge Position Sequence
	WedgePositionSequenceTag = 0x300A0116
	// WedgePositionTag is tag for Wedge Position
	WedgePositionTag = 0x300A0118
	// BeamLimitingDevicePositionSequenceTag is tag for Beam Limiting Device Position Sequence
	BeamLimitingDevicePositionSequenceTag = 0x300A011A
	// LeafJawPositionsTag is tag for Leaf/Jaw Positions
	LeafJawPositionsTag = 0x300A011C
	// GantryAngleTag is tag for Gantry Angle
	GantryAngleTag = 0x300A011E
	// GantryRotationDirectionTag is tag for Gantry Rotation Direction
	GantryRotationDirectionTag = 0x300A011F
	// BeamLimitingDeviceAngleTag is tag for Beam Limiting Device Angle
	BeamLimitingDeviceAngleTag = 0x300A0120
	// BeamLimitingDeviceRotationDirectionTag is tag for Beam Limiting Device Rotation Direction
	BeamLimitingDeviceRotationDirectionTag = 0x300A0121
	// PatientSupportAngleTag is tag for Patient Support Angle
	PatientSupportAngleTag = 0x300A0122
	// PatientSupportRotationDirectionTag is tag for Patient Support Rotation Direction
	PatientSupportRotationDirectionTag = 0x300A0123
	// TableTopEccentricAxisDistanceTag is tag for Table Top Eccentric Axis Distance
	TableTopEccentricAxisDistanceTag = 0x300A0124
	// TableTopEccentricAngleTag is tag for Table Top Eccentric Angle
	TableTopEccentricAngleTag = 0x300A0125
	// TableTopEccentricRotationDirectionTag is tag for Table Top Eccentric Rotation Direction
	TableTopEccentricRotationDirectionTag = 0x300A0126
	// TableTopVerticalPositionTag is tag for Table Top Vertical Position
	TableTopVerticalPositionTag = 0x300A0128
	// TableTopLongitudinalPositionTag is tag for Table Top Longitudinal Position
	TableTopLongitudinalPositionTag = 0x300A0129
	// TableTopLateralPositionTag is tag for Table Top Lateral Position
	TableTopLateralPositionTag = 0x300A012A
	// IsocenterPositionTag is tag for Isocenter Position
	IsocenterPositionTag = 0x300A012C
	// SurfaceEntryPointTag is tag for Surface Entry Point
	SurfaceEntryPointTag = 0x300A012E
	// SourceToSurfaceDistanceTag is tag for Source to Surface Distance
	SourceToSurfaceDistanceTag = 0x300A0130
	// AverageBeamDosePointSourceToExternalContourDistanceTag is tag for Average Beam Dose Point Source to External Contour Distance
	AverageBeamDosePointSourceToExternalContourDistanceTag = 0x300A0131
	// SourceToExternalContourDistanceTag is tag for Source to External Contour Distance
	SourceToExternalContourDistanceTag = 0x300A0132
	// ExternalContourEntryPointTag is tag for External Contour Entry Point
	ExternalContourEntryPointTag = 0x300A0133
	// CumulativeMetersetWeightTag is tag for Cumulative Meterset Weight
	CumulativeMetersetWeightTag = 0x300A0134
	// TableTopPitchAngleTag is tag for Table Top Pitch Angle
	TableTopPitchAngleTag = 0x300A0140
	// TableTopPitchRotationDirectionTag is tag for Table Top Pitch Rotation Direction
	TableTopPitchRotationDirectionTag = 0x300A0142
	// TableTopRollAngleTag is tag for Table Top Roll Angle
	TableTopRollAngleTag = 0x300A0144
	// TableTopRollRotationDirectionTag is tag for Table Top Roll Rotation Direction
	TableTopRollRotationDirectionTag = 0x300A0146
	// HeadFixationAngleTag is tag for Head Fixation Angle
	HeadFixationAngleTag = 0x300A0148
	// GantryPitchAngleTag is tag for Gantry Pitch Angle
	GantryPitchAngleTag = 0x300A014A
	// GantryPitchRotationDirectionTag is tag for Gantry Pitch Rotation Direction
	GantryPitchRotationDirectionTag = 0x300A014C
	// GantryPitchAngleToleranceTag is tag for Gantry Pitch Angle Tolerance
	GantryPitchAngleToleranceTag = 0x300A014E
	// FixationEyeTag is tag for Fixation Eye
	FixationEyeTag = 0x300A0150
	// ChairHeadFramePositionTag is tag for Chair Head Frame Position
	ChairHeadFramePositionTag = 0x300A0151
	// HeadFixationAngleToleranceTag is tag for Head Fixation Angle Tolerance
	HeadFixationAngleToleranceTag = 0x300A0152
	// ChairHeadFramePositionToleranceTag is tag for Chair Head Frame Position Tolerance
	ChairHeadFramePositionToleranceTag = 0x300A0153
	// FixationLightAzimuthalAngleToleranceTag is tag for Fixation Light Azimuthal Angle Tolerance
	FixationLightAzimuthalAngleToleranceTag = 0x300A0154
	// FixationLightPolarAngleToleranceTag is tag for Fixation Light Polar Angle Tolerance
	FixationLightPolarAngleToleranceTag = 0x300A0155
	// PatientSetupSequenceTag is tag for Patient Setup Sequence
	PatientSetupSequenceTag = 0x300A0180
	// PatientSetupNumberTag is tag for Patient Setup Number
	PatientSetupNumberTag = 0x300A0182
	// PatientSetupLabelTag is tag for Patient Setup Label
	PatientSetupLabelTag = 0x300A0183
	// PatientAdditionalPositionTag is tag for Patient Additional Position
	PatientAdditionalPositionTag = 0x300A0184
	// FixationDeviceSequenceTag is tag for Fixation Device Sequence
	FixationDeviceSequenceTag = 0x300A0190
	// FixationDeviceTypeTag is tag for Fixation Device Type
	FixationDeviceTypeTag = 0x300A0192
	// FixationDeviceLabelTag is tag for Fixation Device Label
	FixationDeviceLabelTag = 0x300A0194
	// FixationDeviceDescriptionTag is tag for Fixation Device Description
	FixationDeviceDescriptionTag = 0x300A0196
	// FixationDevicePositionTag is tag for Fixation Device Position
	FixationDevicePositionTag = 0x300A0198
	// FixationDevicePitchAngleTag is tag for Fixation Device Pitch Angle
	FixationDevicePitchAngleTag = 0x300A0199
	// FixationDeviceRollAngleTag is tag for Fixation Device Roll Angle
	FixationDeviceRollAngleTag = 0x300A019A
	// ShieldingDeviceSequenceTag is tag for Shielding Device Sequence
	ShieldingDeviceSequenceTag = 0x300A01A0
	// ShieldingDeviceTypeTag is tag for Shielding Device Type
	ShieldingDeviceTypeTag = 0x300A01A2
	// ShieldingDeviceLabelTag is tag for Shielding Device Label
	ShieldingDeviceLabelTag = 0x300A01A4
	// ShieldingDeviceDescriptionTag is tag for Shielding Device Description
	ShieldingDeviceDescriptionTag = 0x300A01A6
	// ShieldingDevicePositionTag is tag for Shielding Device Position
	ShieldingDevicePositionTag = 0x300A01A8
	// SetupTechniqueTag is tag for Setup Technique
	SetupTechniqueTag = 0x300A01B0
	// SetupTechniqueDescriptionTag is tag for Setup Technique Description
	SetupTechniqueDescriptionTag = 0x300A01B2
	// SetupDeviceSequenceTag is tag for Setup Device Sequence
	SetupDeviceSequenceTag = 0x300A01B4
	// SetupDeviceTypeTag is tag for Setup Device Type
	SetupDeviceTypeTag = 0x300A01B6
	// SetupDeviceLabelTag is tag for Setup Device Label
	SetupDeviceLabelTag = 0x300A01B8
	// SetupDeviceDescriptionTag is tag for Setup Device Description
	SetupDeviceDescriptionTag = 0x300A01BA
	// SetupDeviceParameterTag is tag for Setup Device Parameter
	SetupDeviceParameterTag = 0x300A01BC
	// SetupReferenceDescriptionTag is tag for Setup Reference Description
	SetupReferenceDescriptionTag = 0x300A01D0
	// TableTopVerticalSetupDisplacementTag is tag for Table Top Vertical Setup Displacement
	TableTopVerticalSetupDisplacementTag = 0x300A01D2
	// TableTopLongitudinalSetupDisplacementTag is tag for Table Top Longitudinal Setup Displacement
	TableTopLongitudinalSetupDisplacementTag = 0x300A01D4
	// TableTopLateralSetupDisplacementTag is tag for Table Top Lateral Setup Displacement
	TableTopLateralSetupDisplacementTag = 0x300A01D6
	// BrachyTreatmentTechniqueTag is tag for Brachy Treatment Technique
	BrachyTreatmentTechniqueTag = 0x300A0200
	// BrachyTreatmentTypeTag is tag for Brachy Treatment Type
	BrachyTreatmentTypeTag = 0x300A0202
	// TreatmentMachineSequenceTag is tag for Treatment Machine Sequence
	TreatmentMachineSequenceTag = 0x300A0206
	// SourceSequenceTag is tag for Source Sequence
	SourceSequenceTag = 0x300A0210
	// SourceNumberTag is tag for Source Number
	SourceNumberTag = 0x300A0212
	// SourceTypeTag is tag for Source Type
	SourceTypeTag = 0x300A0214
	// SourceManufacturerTag is tag for Source Manufacturer
	SourceManufacturerTag = 0x300A0216
	// ActiveSourceDiameterTag is tag for Active Source Diameter
	ActiveSourceDiameterTag = 0x300A0218
	// ActiveSourceLengthTag is tag for Active Source Length
	ActiveSourceLengthTag = 0x300A021A
	// SourceModelIDTag is tag for Source Model ID
	SourceModelIDTag = 0x300A021B
	// SourceDescriptionTag is tag for Source Description
	SourceDescriptionTag = 0x300A021C
	// SourceEncapsulationNominalThicknessTag is tag for Source Encapsulation Nominal Thickness
	SourceEncapsulationNominalThicknessTag = 0x300A0222
	// SourceEncapsulationNominalTransmissionTag is tag for Source Encapsulation Nominal Transmission
	SourceEncapsulationNominalTransmissionTag = 0x300A0224
	// SourceIsotopeNameTag is tag for Source Isotope Name
	SourceIsotopeNameTag = 0x300A0226
	// SourceIsotopeHalfLifeTag is tag for Source Isotope Half Life
	SourceIsotopeHalfLifeTag = 0x300A0228
	// SourceStrengthUnitsTag is tag for Source Strength Units
	SourceStrengthUnitsTag = 0x300A0229
	// ReferenceAirKermaRateTag is tag for Reference Air Kerma Rate
	ReferenceAirKermaRateTag = 0x300A022A
	// SourceStrengthTag is tag for Source Strength
	SourceStrengthTag = 0x300A022B
	// SourceStrengthReferenceDateTag is tag for Source Strength Reference Date
	SourceStrengthReferenceDateTag = 0x300A022C
	// SourceStrengthReferenceTimeTag is tag for Source Strength Reference Time
	SourceStrengthReferenceTimeTag = 0x300A022E
	// ApplicationSetupSequenceTag is tag for Application Setup Sequence
	ApplicationSetupSequenceTag = 0x300A0230
	// ApplicationSetupTypeTag is tag for Application Setup Type
	ApplicationSetupTypeTag = 0x300A0232
	// ApplicationSetupNumberTag is tag for Application Setup Number
	ApplicationSetupNumberTag = 0x300A0234
	// ApplicationSetupNameTag is tag for Application Setup Name
	ApplicationSetupNameTag = 0x300A0236
	// ApplicationSetupManufacturerTag is tag for Application Setup Manufacturer
	ApplicationSetupManufacturerTag = 0x300A0238
	// TemplateNumberTag is tag for Template Number
	TemplateNumberTag = 0x300A0240
	// TemplateTypeTag is tag for Template Type
	TemplateTypeTag = 0x300A0242
	// TemplateNameTag is tag for Template Name
	TemplateNameTag = 0x300A0244
	// TotalReferenceAirKermaTag is tag for Total Reference Air Kerma
	TotalReferenceAirKermaTag = 0x300A0250
	// BrachyAccessoryDeviceSequenceTag is tag for Brachy Accessory Device Sequence
	BrachyAccessoryDeviceSequenceTag = 0x300A0260
	// BrachyAccessoryDeviceNumberTag is tag for Brachy Accessory Device Number
	BrachyAccessoryDeviceNumberTag = 0x300A0262
	// BrachyAccessoryDeviceIDTag is tag for Brachy Accessory Device ID
	BrachyAccessoryDeviceIDTag = 0x300A0263
	// BrachyAccessoryDeviceTypeTag is tag for Brachy Accessory Device Type
	BrachyAccessoryDeviceTypeTag = 0x300A0264
	// BrachyAccessoryDeviceNameTag is tag for Brachy Accessory Device Name
	BrachyAccessoryDeviceNameTag = 0x300A0266
	// BrachyAccessoryDeviceNominalThicknessTag is tag for Brachy Accessory Device Nominal Thickness
	BrachyAccessoryDeviceNominalThicknessTag = 0x300A026A
	// BrachyAccessoryDeviceNominalTransmissionTag is tag for Brachy Accessory Device Nominal Transmission
	BrachyAccessoryDeviceNominalTransmissionTag = 0x300A026C
	// ChannelEffectiveLengthTag is tag for Channel Effective Length
	ChannelEffectiveLengthTag = 0x300A0271
	// ChannelInnerLengthTag is tag for Channel Inner Length
	ChannelInnerLengthTag = 0x300A0272
	// AfterloaderChannelIDTag is tag for Afterloader Channel ID
	AfterloaderChannelIDTag = 0x300A0273
	// SourceApplicatorTipLengthTag is tag for Source Applicator Tip Length
	SourceApplicatorTipLengthTag = 0x300A0274
	// ChannelSequenceTag is tag for Channel Sequence
	ChannelSequenceTag = 0x300A0280
	// ChannelNumberTag is tag for Channel Number
	ChannelNumberTag = 0x300A0282
	// ChannelLengthTag is tag for Channel Length
	ChannelLengthTag = 0x300A0284
	// ChannelTotalTimeTag is tag for Channel Total Time
	ChannelTotalTimeTag = 0x300A0286
	// SourceMovementTypeTag is tag for Source Movement Type
	SourceMovementTypeTag = 0x300A0288
	// NumberOfPulsesTag is tag for Number of Pulses
	NumberOfPulsesTag = 0x300A028A
	// PulseRepetitionIntervalTag is tag for Pulse Repetition Interval
	PulseRepetitionIntervalTag = 0x300A028C
	// SourceApplicatorNumberTag is tag for Source Applicator Number
	SourceApplicatorNumberTag = 0x300A0290
	// SourceApplicatorIDTag is tag for Source Applicator ID
	SourceApplicatorIDTag = 0x300A0291
	// SourceApplicatorTypeTag is tag for Source Applicator Type
	SourceApplicatorTypeTag = 0x300A0292
	// SourceApplicatorNameTag is tag for Source Applicator Name
	SourceApplicatorNameTag = 0x300A0294
	// SourceApplicatorLengthTag is tag for Source Applicator Length
	SourceApplicatorLengthTag = 0x300A0296
	// SourceApplicatorManufacturerTag is tag for Source Applicator Manufacturer
	SourceApplicatorManufacturerTag = 0x300A0298
	// SourceApplicatorWallNominalThicknessTag is tag for Source Applicator Wall Nominal Thickness
	SourceApplicatorWallNominalThicknessTag = 0x300A029C
	// SourceApplicatorWallNominalTransmissionTag is tag for Source Applicator Wall Nominal Transmission
	SourceApplicatorWallNominalTransmissionTag = 0x300A029E
	// SourceApplicatorStepSizeTag is tag for Source Applicator Step Size
	SourceApplicatorStepSizeTag = 0x300A02A0
	// TransferTubeNumberTag is tag for Transfer Tube Number
	TransferTubeNumberTag = 0x300A02A2
	// TransferTubeLengthTag is tag for Transfer Tube Length
	TransferTubeLengthTag = 0x300A02A4
	// ChannelShieldSequenceTag is tag for Channel Shield Sequence
	ChannelShieldSequenceTag = 0x300A02B0
	// ChannelShieldNumberTag is tag for Channel Shield Number
	ChannelShieldNumberTag = 0x300A02B2
	// ChannelShieldIDTag is tag for Channel Shield ID
	ChannelShieldIDTag = 0x300A02B3
	// ChannelShieldNameTag is tag for Channel Shield Name
	ChannelShieldNameTag = 0x300A02B4
	// ChannelShieldNominalThicknessTag is tag for Channel Shield Nominal Thickness
	ChannelShieldNominalThicknessTag = 0x300A02B8
	// ChannelShieldNominalTransmissionTag is tag for Channel Shield Nominal Transmission
	ChannelShieldNominalTransmissionTag = 0x300A02BA
	// FinalCumulativeTimeWeightTag is tag for Final Cumulative Time Weight
	FinalCumulativeTimeWeightTag = 0x300A02C8
	// BrachyControlPointSequenceTag is tag for Brachy Control Point Sequence
	BrachyControlPointSequenceTag = 0x300A02D0
	// ControlPointRelativePositionTag is tag for Control Point Relative Position
	ControlPointRelativePositionTag = 0x300A02D2
	// ControlPoint3DPositionTag is tag for Control Point 3D Position
	ControlPoint3DPositionTag = 0x300A02D4
	// CumulativeTimeWeightTag is tag for Cumulative Time Weight
	CumulativeTimeWeightTag = 0x300A02D6
	// CompensatorDivergenceTag is tag for Compensator Divergence
	CompensatorDivergenceTag = 0x300A02E0
	// CompensatorMountingPositionTag is tag for Compensator Mounting Position
	CompensatorMountingPositionTag = 0x300A02E1
	// SourceToCompensatorDistanceTag is tag for Source to Compensator Distance
	SourceToCompensatorDistanceTag = 0x300A02E2
	// TotalCompensatorTrayWaterEquivalentThicknessTag is tag for Total Compensator Tray Water-Equivalent Thickness
	TotalCompensatorTrayWaterEquivalentThicknessTag = 0x300A02E3
	// IsocenterToCompensatorTrayDistanceTag is tag for Isocenter to Compensator Tray Distance
	IsocenterToCompensatorTrayDistanceTag = 0x300A02E4
	// CompensatorColumnOffsetTag is tag for Compensator Column Offset
	CompensatorColumnOffsetTag = 0x300A02E5
	// IsocenterToCompensatorDistancesTag is tag for Isocenter to Compensator Distances
	IsocenterToCompensatorDistancesTag = 0x300A02E6
	// CompensatorRelativeStoppingPowerRatioTag is tag for Compensator Relative Stopping Power Ratio
	CompensatorRelativeStoppingPowerRatioTag = 0x300A02E7
	// CompensatorMillingToolDiameterTag is tag for Compensator Milling Tool Diameter
	CompensatorMillingToolDiameterTag = 0x300A02E8
	// IonRangeCompensatorSequenceTag is tag for Ion Range Compensator Sequence
	IonRangeCompensatorSequenceTag = 0x300A02EA
	// CompensatorDescriptionTag is tag for Compensator Description
	CompensatorDescriptionTag = 0x300A02EB
	// RadiationMassNumberTag is tag for Radiation Mass Number
	RadiationMassNumberTag = 0x300A0302
	// RadiationAtomicNumberTag is tag for Radiation Atomic Number
	RadiationAtomicNumberTag = 0x300A0304
	// RadiationChargeStateTag is tag for Radiation Charge State
	RadiationChargeStateTag = 0x300A0306
	// ScanModeTag is tag for Scan Mode
	ScanModeTag = 0x300A0308
	// ModulatedScanModeTypeTag is tag for Modulated Scan Mode Type
	ModulatedScanModeTypeTag = 0x300A0309
	// VirtualSourceAxisDistancesTag is tag for Virtual Source-Axis Distances
	VirtualSourceAxisDistancesTag = 0x300A030A
	// SnoutSequenceTag is tag for Snout Sequence
	SnoutSequenceTag = 0x300A030C
	// SnoutPositionTag is tag for Snout Position
	SnoutPositionTag = 0x300A030D
	// SnoutIDTag is tag for Snout ID
	SnoutIDTag = 0x300A030F
	// NumberOfRangeShiftersTag is tag for Number of Range Shifters
	NumberOfRangeShiftersTag = 0x300A0312
	// RangeShifterSequenceTag is tag for Range Shifter Sequence
	RangeShifterSequenceTag = 0x300A0314
	// RangeShifterNumberTag is tag for Range Shifter Number
	RangeShifterNumberTag = 0x300A0316
	// RangeShifterIDTag is tag for Range Shifter ID
	RangeShifterIDTag = 0x300A0318
	// RangeShifterTypeTag is tag for Range Shifter Type
	RangeShifterTypeTag = 0x300A0320
	// RangeShifterDescriptionTag is tag for Range Shifter Description
	RangeShifterDescriptionTag = 0x300A0322
	// NumberOfLateralSpreadingDevicesTag is tag for Number of Lateral Spreading Devices
	NumberOfLateralSpreadingDevicesTag = 0x300A0330
	// LateralSpreadingDeviceSequenceTag is tag for Lateral Spreading Device Sequence
	LateralSpreadingDeviceSequenceTag = 0x300A0332
	// LateralSpreadingDeviceNumberTag is tag for Lateral Spreading Device Number
	LateralSpreadingDeviceNumberTag = 0x300A0334
	// LateralSpreadingDeviceIDTag is tag for Lateral Spreading Device ID
	LateralSpreadingDeviceIDTag = 0x300A0336
	// LateralSpreadingDeviceTypeTag is tag for Lateral Spreading Device Type
	LateralSpreadingDeviceTypeTag = 0x300A0338
	// LateralSpreadingDeviceDescriptionTag is tag for Lateral Spreading Device Description
	LateralSpreadingDeviceDescriptionTag = 0x300A033A
	// LateralSpreadingDeviceWaterEquivalentThicknessTag is tag for Lateral Spreading Device Water Equivalent Thickness
	LateralSpreadingDeviceWaterEquivalentThicknessTag = 0x300A033C
	// NumberOfRangeModulatorsTag is tag for Number of Range Modulators
	NumberOfRangeModulatorsTag = 0x300A0340
	// RangeModulatorSequenceTag is tag for Range Modulator Sequence
	RangeModulatorSequenceTag = 0x300A0342
	// RangeModulatorNumberTag is tag for Range Modulator Number
	RangeModulatorNumberTag = 0x300A0344
	// RangeModulatorIDTag is tag for Range Modulator ID
	RangeModulatorIDTag = 0x300A0346
	// RangeModulatorTypeTag is tag for Range Modulator Type
	RangeModulatorTypeTag = 0x300A0348
	// RangeModulatorDescriptionTag is tag for Range Modulator Description
	RangeModulatorDescriptionTag = 0x300A034A
	// BeamCurrentModulationIDTag is tag for Beam Current Modulation ID
	BeamCurrentModulationIDTag = 0x300A034C
	// PatientSupportTypeTag is tag for Patient Support Type
	PatientSupportTypeTag = 0x300A0350
	// PatientSupportIDTag is tag for Patient Support ID
	PatientSupportIDTag = 0x300A0352
	// PatientSupportAccessoryCodeTag is tag for Patient Support Accessory Code
	PatientSupportAccessoryCodeTag = 0x300A0354
	// TrayAccessoryCodeTag is tag for Tray Accessory Code
	TrayAccessoryCodeTag = 0x300A0355
	// FixationLightAzimuthalAngleTag is tag for Fixation Light Azimuthal Angle
	FixationLightAzimuthalAngleTag = 0x300A0356
	// FixationLightPolarAngleTag is tag for Fixation Light Polar Angle
	FixationLightPolarAngleTag = 0x300A0358
	// MetersetRateTag is tag for Meterset Rate
	MetersetRateTag = 0x300A035A
	// RangeShifterSettingsSequenceTag is tag for Range Shifter Settings Sequence
	RangeShifterSettingsSequenceTag = 0x300A0360
	// RangeShifterSettingTag is tag for Range Shifter Setting
	RangeShifterSettingTag = 0x300A0362
	// IsocenterToRangeShifterDistanceTag is tag for Isocenter to Range Shifter Distance
	IsocenterToRangeShifterDistanceTag = 0x300A0364
	// RangeShifterWaterEquivalentThicknessTag is tag for Range Shifter Water Equivalent Thickness
	RangeShifterWaterEquivalentThicknessTag = 0x300A0366
	// LateralSpreadingDeviceSettingsSequenceTag is tag for Lateral Spreading Device Settings Sequence
	LateralSpreadingDeviceSettingsSequenceTag = 0x300A0370
	// LateralSpreadingDeviceSettingTag is tag for Lateral Spreading Device Setting
	LateralSpreadingDeviceSettingTag = 0x300A0372
	// IsocenterToLateralSpreadingDeviceDistanceTag is tag for Isocenter to Lateral Spreading Device Distance
	IsocenterToLateralSpreadingDeviceDistanceTag = 0x300A0374
	// RangeModulatorSettingsSequenceTag is tag for Range Modulator Settings Sequence
	RangeModulatorSettingsSequenceTag = 0x300A0380
	// RangeModulatorGatingStartValueTag is tag for Range Modulator Gating Start Value
	RangeModulatorGatingStartValueTag = 0x300A0382
	// RangeModulatorGatingStopValueTag is tag for Range Modulator Gating Stop Value
	RangeModulatorGatingStopValueTag = 0x300A0384
	// RangeModulatorGatingStartWaterEquivalentThicknessTag is tag for Range Modulator Gating Start Water Equivalent Thickness
	RangeModulatorGatingStartWaterEquivalentThicknessTag = 0x300A0386
	// RangeModulatorGatingStopWaterEquivalentThicknessTag is tag for Range Modulator Gating Stop Water Equivalent Thickness
	RangeModulatorGatingStopWaterEquivalentThicknessTag = 0x300A0388
	// IsocenterToRangeModulatorDistanceTag is tag for Isocenter to Range Modulator Distance
	IsocenterToRangeModulatorDistanceTag = 0x300A038A
	// ScanSpotTimeOffsetTag is tag for Scan Spot Time Offset
	ScanSpotTimeOffsetTag = 0x300A038F
	// ScanSpotTuneIDTag is tag for Scan Spot Tune ID
	ScanSpotTuneIDTag = 0x300A0390
	// ScanSpotPrescribedIndicesTag is tag for Scan Spot Prescribed Indices
	ScanSpotPrescribedIndicesTag = 0x300A0391
	// NumberOfScanSpotPositionsTag is tag for Number of Scan Spot Positions
	NumberOfScanSpotPositionsTag = 0x300A0392
	// ScanSpotReorderedTag is tag for Scan Spot Reordered
	ScanSpotReorderedTag = 0x300A0393
	// ScanSpotPositionMapTag is tag for Scan Spot Position Map
	ScanSpotPositionMapTag = 0x300A0394
	// ScanSpotReorderingAllowedTag is tag for Scan Spot Reordering Allowed
	ScanSpotReorderingAllowedTag = 0x300A0395
	// ScanSpotMetersetWeightsTag is tag for Scan Spot Meterset Weights
	ScanSpotMetersetWeightsTag = 0x300A0396
	// ScanningSpotSizeTag is tag for Scanning Spot Size
	ScanningSpotSizeTag = 0x300A0398
	// NumberOfPaintingsTag is tag for Number of Paintings
	NumberOfPaintingsTag = 0x300A039A
	// IonToleranceTableSequenceTag is tag for Ion Tolerance Table Sequence
	IonToleranceTableSequenceTag = 0x300A03A0
	// IonBeamSequenceTag is tag for Ion Beam Sequence
	IonBeamSequenceTag = 0x300A03A2
	// IonBeamLimitingDeviceSequenceTag is tag for Ion Beam Limiting Device Sequence
	IonBeamLimitingDeviceSequenceTag = 0x300A03A4
	// IonBlockSequenceTag is tag for Ion Block Sequence
	IonBlockSequenceTag = 0x300A03A6
	// IonControlPointSequenceTag is tag for Ion Control Point Sequence
	IonControlPointSequenceTag = 0x300A03A8
	// IonWedgeSequenceTag is tag for Ion Wedge Sequence
	IonWedgeSequenceTag = 0x300A03AA
	// IonWedgePositionSequenceTag is tag for Ion Wedge Position Sequence
	IonWedgePositionSequenceTag = 0x300A03AC
	// ReferencedSetupImageSequenceTag is tag for Referenced Setup Image Sequence
	ReferencedSetupImageSequenceTag = 0x300A0401
	// SetupImageCommentTag is tag for Setup Image Comment
	SetupImageCommentTag = 0x300A0402
	// MotionSynchronizationSequenceTag is tag for Motion Synchronization Sequence
	MotionSynchronizationSequenceTag = 0x300A0410
	// ControlPointOrientationTag is tag for Control Point Orientation
	ControlPointOrientationTag = 0x300A0412
	// GeneralAccessorySequenceTag is tag for General Accessory Sequence
	GeneralAccessorySequenceTag = 0x300A0420
	// GeneralAccessoryIDTag is tag for General Accessory ID
	GeneralAccessoryIDTag = 0x300A0421
	// GeneralAccessoryDescriptionTag is tag for General Accessory Description
	GeneralAccessoryDescriptionTag = 0x300A0422
	// GeneralAccessoryTypeTag is tag for General Accessory Type
	GeneralAccessoryTypeTag = 0x300A0423
	// GeneralAccessoryNumberTag is tag for General Accessory Number
	GeneralAccessoryNumberTag = 0x300A0424
	// SourceToGeneralAccessoryDistanceTag is tag for Source to General Accessory Distance
	SourceToGeneralAccessoryDistanceTag = 0x300A0425
	// IsocenterToGeneralAccessoryDistanceTag is tag for Isocenter to General Accessory Distance
	IsocenterToGeneralAccessoryDistanceTag = 0x300A0426
	// ApplicatorGeometrySequenceTag is tag for Applicator Geometry Sequence
	ApplicatorGeometrySequenceTag = 0x300A0431
	// ApplicatorApertureShapeTag is tag for Applicator Aperture Shape
	ApplicatorApertureShapeTag = 0x300A0432
	// ApplicatorOpeningTag is tag for Applicator Opening
	ApplicatorOpeningTag = 0x300A0433
	// ApplicatorOpeningXTag is tag for Applicator Opening X
	ApplicatorOpeningXTag = 0x300A0434
	// ApplicatorOpeningYTag is tag for Applicator Opening Y
	ApplicatorOpeningYTag = 0x300A0435
	// SourceToApplicatorMountingPositionDistanceTag is tag for Source to Applicator Mounting Position Distance
	SourceToApplicatorMountingPositionDistanceTag = 0x300A0436
	// NumberOfBlockSlabItemsTag is tag for Number of Block Slab Items
	NumberOfBlockSlabItemsTag = 0x300A0440
	// BlockSlabSequenceTag is tag for Block Slab Sequence
	BlockSlabSequenceTag = 0x300A0441
	// BlockSlabThicknessTag is tag for Block Slab Thickness
	BlockSlabThicknessTag = 0x300A0442
	// BlockSlabNumberTag is tag for Block Slab Number
	BlockSlabNumberTag = 0x300A0443
	// DeviceMotionControlSequenceTag is tag for Device Motion Control Sequence
	DeviceMotionControlSequenceTag = 0x300A0450
	// DeviceMotionExecutionModeTag is tag for Device Motion Execution Mode
	DeviceMotionExecutionModeTag = 0x300A0451
	// DeviceMotionObservationModeTag is tag for Device Motion Observation Mode
	DeviceMotionObservationModeTag = 0x300A0452
	// DeviceMotionParameterCodeSequenceTag is tag for Device Motion Parameter Code Sequence
	DeviceMotionParameterCodeSequenceTag = 0x300A0453
	// DistalDepthFractionTag is tag for Distal Depth Fraction
	DistalDepthFractionTag = 0x300A0501
	// DistalDepthTag is tag for Distal Depth
	DistalDepthTag = 0x300A0502
	// NominalRangeModulationFractionsTag is tag for Nominal Range Modulation Fractions
	NominalRangeModulationFractionsTag = 0x300A0503
	// NominalRangeModulatedRegionDepthsTag is tag for Nominal Range Modulated Region Depths
	NominalRangeModulatedRegionDepthsTag = 0x300A0504
	// DepthDoseParametersSequenceTag is tag for Depth Dose Parameters Sequence
	DepthDoseParametersSequenceTag = 0x300A0505
	// DeliveredDepthDoseParametersSequenceTag is tag for Delivered Depth Dose Parameters Sequence
	DeliveredDepthDoseParametersSequenceTag = 0x300A0506
	// DeliveredDistalDepthFractionTag is tag for Delivered Distal Depth Fraction
	DeliveredDistalDepthFractionTag = 0x300A0507
	// DeliveredDistalDepthTag is tag for Delivered Distal Depth
	DeliveredDistalDepthTag = 0x300A0508
	// DeliveredNominalRangeModulationFractionsTag is tag for Delivered Nominal Range Modulation Fractions
	DeliveredNominalRangeModulationFractionsTag = 0x300A0509
	// DeliveredNominalRangeModulatedRegionDepthsTag is tag for Delivered Nominal Range Modulated Region Depths
	DeliveredNominalRangeModulatedRegionDepthsTag = 0x300A0510
	// DeliveredReferenceDoseDefinitionTag is tag for Delivered Reference Dose Definition
	DeliveredReferenceDoseDefinitionTag = 0x300A0511
	// ReferenceDoseDefinitionTag is tag for Reference Dose Definition
	ReferenceDoseDefinitionTag = 0x300A0512
	// RTControlPointIndexTag is tag for RT Control Point Index
	RTControlPointIndexTag = 0x300A0600
	// RadiationGenerationModeIndexTag is tag for Radiation Generation Mode Index
	RadiationGenerationModeIndexTag = 0x300A0601
	// ReferencedDefinedDeviceIndexTag is tag for Referenced Defined Device Index
	ReferencedDefinedDeviceIndexTag = 0x300A0602
	// RadiationDoseIdentificationIndexTag is tag for Radiation Dose Identification Index
	RadiationDoseIdentificationIndexTag = 0x300A0603
	// NumberOfRTControlPointsTag is tag for Number of RT Control Points
	NumberOfRTControlPointsTag = 0x300A0604
	// ReferencedRadiationGenerationModeIndexTag is tag for Referenced Radiation Generation Mode Index
	ReferencedRadiationGenerationModeIndexTag = 0x300A0605
	// TreatmentPositionIndexTag is tag for Treatment Position Index
	TreatmentPositionIndexTag = 0x300A0606
	// ReferencedDeviceIndexTag is tag for Referenced Device Index
	ReferencedDeviceIndexTag = 0x300A0607
	// TreatmentPositionGroupLabelTag is tag for Treatment Position Group Label
	TreatmentPositionGroupLabelTag = 0x300A0608
	// TreatmentPositionGroupUIDTag is tag for Treatment Position Group UID
	TreatmentPositionGroupUIDTag = 0x300A0609
	// TreatmentPositionGroupSequenceTag is tag for Treatment Position Group Sequence
	TreatmentPositionGroupSequenceTag = 0x300A060A
	// ReferencedTreatmentPositionIndexTag is tag for Referenced Treatment Position Index
	ReferencedTreatmentPositionIndexTag = 0x300A060B
	// ReferencedRadiationDoseIdentificationIndexTag is tag for Referenced Radiation Dose Identification Index
	ReferencedRadiationDoseIdentificationIndexTag = 0x300A060C
	// RTAccessoryHolderWaterEquivalentThicknessTag is tag for RT Accessory Holder Water-Equivalent Thickness
	RTAccessoryHolderWaterEquivalentThicknessTag = 0x300A060D
	// ReferencedRTAccessoryHolderDeviceIndexTag is tag for Referenced RT Accessory Holder Device Index
	ReferencedRTAccessoryHolderDeviceIndexTag = 0x300A060E
	// RTAccessoryHolderSlotExistenceFlagTag is tag for RT Accessory Holder Slot Existence Flag
	RTAccessoryHolderSlotExistenceFlagTag = 0x300A060F
	// RTAccessoryHolderSlotSequenceTag is tag for RT Accessory Holder Slot Sequence
	RTAccessoryHolderSlotSequenceTag = 0x300A0610
	// RTAccessoryHolderSlotIDTag is tag for RT Accessory Holder Slot ID
	RTAccessoryHolderSlotIDTag = 0x300A0611
	// RTAccessoryHolderSlotDistanceTag is tag for RT Accessory Holder Slot Distance
	RTAccessoryHolderSlotDistanceTag = 0x300A0612
	// RTAccessorySlotDistanceTag is tag for RT Accessory Slot Distance
	RTAccessorySlotDistanceTag = 0x300A0613
	// RTAccessoryHolderDefinitionSequenceTag is tag for RT Accessory Holder Definition Sequence
	RTAccessoryHolderDefinitionSequenceTag = 0x300A0614
	// RTAccessoryDeviceSlotIDTag is tag for RT Accessory Device Slot ID
	RTAccessoryDeviceSlotIDTag = 0x300A0615
	// RTRadiationSequenceTag is tag for RT Radiation Sequence
	RTRadiationSequenceTag = 0x300A0616
	// RadiationDoseSequenceTag is tag for Radiation Dose Sequence
	RadiationDoseSequenceTag = 0x300A0617
	// RadiationDoseIdentificationSequenceTag is tag for Radiation Dose Identification Sequence
	RadiationDoseIdentificationSequenceTag = 0x300A0618
	// RadiationDoseIdentificationLabelTag is tag for Radiation Dose Identification Label
	RadiationDoseIdentificationLabelTag = 0x300A0619
	// ReferenceDoseTypeTag is tag for Reference Dose Type
	ReferenceDoseTypeTag = 0x300A061A
	// PrimaryDoseValueIndicatorTag is tag for Primary Dose Value Indicator
	PrimaryDoseValueIndicatorTag = 0x300A061B
	// DoseValuesSequenceTag is tag for Dose Values Sequence
	DoseValuesSequenceTag = 0x300A061C
	// DoseValuePurposeTag is tag for Dose Value Purpose
	DoseValuePurposeTag = 0x300A061D
	// ReferenceDosePointCoordinatesTag is tag for Reference Dose Point Coordinates
	ReferenceDosePointCoordinatesTag = 0x300A061E
	// RadiationDoseValuesParametersSequenceTag is tag for Radiation Dose Values Parameters Sequence
	RadiationDoseValuesParametersSequenceTag = 0x300A061F
	// MetersetToDoseMappingSequenceTag is tag for Meterset to Dose Mapping Sequence
	MetersetToDoseMappingSequenceTag = 0x300A0620
	// ExpectedInVivoMeasurementValuesSequenceTag is tag for Expected In-Vivo Measurement Values Sequence
	ExpectedInVivoMeasurementValuesSequenceTag = 0x300A0621
	// ExpectedInVivoMeasurementValueIndexTag is tag for Expected In-Vivo Measurement Value Index
	ExpectedInVivoMeasurementValueIndexTag = 0x300A0622
	// RadiationDoseInVivoMeasurementLabelTag is tag for Radiation Dose In-Vivo Measurement Label
	RadiationDoseInVivoMeasurementLabelTag = 0x300A0623
	// RadiationDoseCentralAxisDisplacementTag is tag for Radiation Dose Central Axis Displacement
	RadiationDoseCentralAxisDisplacementTag = 0x300A0624
	// RadiationDoseValueTag is tag for Radiation Dose Value
	RadiationDoseValueTag = 0x300A0625
	// RadiationDoseSourceToSkinDistanceTag is tag for Radiation Dose Source to Skin Distance
	RadiationDoseSourceToSkinDistanceTag = 0x300A0626
	// RadiationDoseMeasurementPointCoordinatesTag is tag for Radiation Dose Measurement Point Coordinates
	RadiationDoseMeasurementPointCoordinatesTag = 0x300A0627
	// RadiationDoseSourceToExternalContourDistanceTag is tag for Radiation Dose Source to External Contour Distance
	RadiationDoseSourceToExternalContourDistanceTag = 0x300A0628
	// RTToleranceSetSequenceTag is tag for RT Tolerance Set Sequence
	RTToleranceSetSequenceTag = 0x300A0629
	// RTToleranceSetLabelTag is tag for RT Tolerance Set Label
	RTToleranceSetLabelTag = 0x300A062A
	// AttributeToleranceValuesSequenceTag is tag for Attribute Tolerance Values Sequence
	AttributeToleranceValuesSequenceTag = 0x300A062B
	// ToleranceValueTag is tag for Tolerance Value
	ToleranceValueTag = 0x300A062C
	// PatientSupportPositionToleranceSequenceTag is tag for Patient Support Position Tolerance Sequence
	PatientSupportPositionToleranceSequenceTag = 0x300A062D
	// TreatmentTimeLimitTag is tag for Treatment Time Limit
	TreatmentTimeLimitTag = 0x300A062E
	// CArmPhotonElectronControlPointSequenceTag is tag for C-Arm Photon-Electron Control Point Sequence
	CArmPhotonElectronControlPointSequenceTag = 0x300A062F
	// ReferencedRTRadiationSequenceTag is tag for Referenced RT Radiation Sequence
	ReferencedRTRadiationSequenceTag = 0x300A0630
	// ReferencedRTInstanceSequenceTag is tag for Referenced RT Instance Sequence
	ReferencedRTInstanceSequenceTag = 0x300A0631
	// ReferencedRTPatientSetupSequenceTag is tag for Referenced RT Patient Setup Sequence
	ReferencedRTPatientSetupSequenceTag = 0x300A0632
	// SourceToPatientSurfaceDistanceTag is tag for Source to Patient Surface Distance
	SourceToPatientSurfaceDistanceTag = 0x300A0634
	// TreatmentMachineSpecialModeCodeSequenceTag is tag for Treatment Machine Special Mode Code Sequence
	TreatmentMachineSpecialModeCodeSequenceTag = 0x300A0635
	// IntendedNumberOfFractionsTag is tag for Intended Number of Fractions
	IntendedNumberOfFractionsTag = 0x300A0636
	// RTRadiationSetIntentTag is tag for RT Radiation Set Intent
	RTRadiationSetIntentTag = 0x300A0637
	// RTRadiationPhysicalAndGeometricContentDetailFlagTag is tag for RT Radiation Physical and Geometric Content Detail Flag
	RTRadiationPhysicalAndGeometricContentDetailFlagTag = 0x300A0638
	// RTRecordFlagTag is tag for RT Record Flag
	RTRecordFlagTag = 0x300A0639
	// TreatmentDeviceIdentificationSequenceTag is tag for Treatment Device Identification Sequence
	TreatmentDeviceIdentificationSequenceTag = 0x300A063A
	// ReferencedRTPhysicianIntentSequenceTag is tag for Referenced RT Physician Intent Sequence
	ReferencedRTPhysicianIntentSequenceTag = 0x300A063B
	// CumulativeMetersetTag is tag for Cumulative Meterset
	CumulativeMetersetTag = 0x300A063C
	// DeliveryRateTag is tag for Delivery Rate
	DeliveryRateTag = 0x300A063D
	// DeliveryRateUnitSequenceTag is tag for Delivery Rate Unit Sequence
	DeliveryRateUnitSequenceTag = 0x300A063E
	// TreatmentPositionSequenceTag is tag for Treatment Position Sequence
	TreatmentPositionSequenceTag = 0x300A063F
	// RadiationSourceAxisDistanceTag is tag for Radiation Source-Axis Distance
	RadiationSourceAxisDistanceTag = 0x300A0640
	// NumberOfRTBeamLimitingDevicesTag is tag for Number of RT Beam Limiting Devices
	NumberOfRTBeamLimitingDevicesTag = 0x300A0641
	// RTBeamLimitingDeviceProximalDistanceTag is tag for RT Beam Limiting Device Proximal Distance
	RTBeamLimitingDeviceProximalDistanceTag = 0x300A0642
	// RTBeamLimitingDeviceDistalDistanceTag is tag for RT Beam Limiting Device Distal Distance
	RTBeamLimitingDeviceDistalDistanceTag = 0x300A0643
	// ParallelRTBeamDelimiterDeviceOrientationLabelCodeSequenceTag is tag for Parallel RT Beam Delimiter Device Orientation Label Code Sequence
	ParallelRTBeamDelimiterDeviceOrientationLabelCodeSequenceTag = 0x300A0644
	// BeamModifierOrientationAngleTag is tag for Beam Modifier Orientation Angle
	BeamModifierOrientationAngleTag = 0x300A0645
	// FixedRTBeamDelimiterDeviceSequenceTag is tag for Fixed RT Beam Delimiter Device Sequence
	FixedRTBeamDelimiterDeviceSequenceTag = 0x300A0646
	// ParallelRTBeamDelimiterDeviceSequenceTag is tag for Parallel RT Beam Delimiter Device Sequence
	ParallelRTBeamDelimiterDeviceSequenceTag = 0x300A0647
	// NumberOfParallelRTBeamDelimitersTag is tag for Number of Parallel RT Beam Delimiters
	NumberOfParallelRTBeamDelimitersTag = 0x300A0648
	// ParallelRTBeamDelimiterBoundariesTag is tag for Parallel RT Beam Delimiter Boundaries
	ParallelRTBeamDelimiterBoundariesTag = 0x300A0649
	// ParallelRTBeamDelimiterPositionsTag is tag for Parallel RT Beam Delimiter Positions
	ParallelRTBeamDelimiterPositionsTag = 0x300A064A
	// RTBeamLimitingDeviceOffsetTag is tag for RT Beam Limiting Device Offset
	RTBeamLimitingDeviceOffsetTag = 0x300A064B
	// RTBeamDelimiterGeometrySequenceTag is tag for RT Beam Delimiter Geometry Sequence
	RTBeamDelimiterGeometrySequenceTag = 0x300A064C
	// RTBeamLimitingDeviceDefinitionSequenceTag is tag for RT Beam Limiting Device Definition Sequence
	RTBeamLimitingDeviceDefinitionSequenceTag = 0x300A064D
	// ParallelRTBeamDelimiterOpeningModeTag is tag for Parallel RT Beam Delimiter Opening Mode
	ParallelRTBeamDelimiterOpeningModeTag = 0x300A064E
	// ParallelRTBeamDelimiterLeafMountingSideTag is tag for Parallel RT Beam Delimiter Leaf Mounting Side
	ParallelRTBeamDelimiterLeafMountingSideTag = 0x300A064F
	// PatientSetupUIDTag is tag for Patient Setup UID
	PatientSetupUIDTag = 0x300A0650
	// WedgeDefinitionSequenceTag is tag for Wedge Definition Sequence
	WedgeDefinitionSequenceTag = 0x300A0651
	// RadiationBeamWedgeAngleTag is tag for Radiation Beam Wedge Angle
	RadiationBeamWedgeAngleTag = 0x300A0652
	// RadiationBeamWedgeThinEdgeDistanceTag is tag for Radiation Beam Wedge Thin Edge Distance
	RadiationBeamWedgeThinEdgeDistanceTag = 0x300A0653
	// RadiationBeamEffectiveWedgeAngleTag is tag for Radiation Beam Effective Wedge Angle
	RadiationBeamEffectiveWedgeAngleTag = 0x300A0654
	// NumberOfWedgePositionsTag is tag for Number of Wedge Positions
	NumberOfWedgePositionsTag = 0x300A0655
	// RTBeamLimitingDeviceOpeningSequenceTag is tag for RT Beam Limiting Device Opening Sequence
	RTBeamLimitingDeviceOpeningSequenceTag = 0x300A0656
	// NumberOfRTBeamLimitingDeviceOpeningsTag is tag for Number of RT Beam Limiting Device Openings
	NumberOfRTBeamLimitingDeviceOpeningsTag = 0x300A0657
	// RadiationDosimeterUnitSequenceTag is tag for Radiation Dosimeter Unit Sequence
	RadiationDosimeterUnitSequenceTag = 0x300A0658
	// RTDeviceDistanceReferenceLocationCodeSequenceTag is tag for RT Device Distance Reference Location Code Sequence
	RTDeviceDistanceReferenceLocationCodeSequenceTag = 0x300A0659
	// RadiationDeviceConfigurationAndCommissioningKeySequenceTag is tag for Radiation Device Configuration and Commissioning Key Sequence
	RadiationDeviceConfigurationAndCommissioningKeySequenceTag = 0x300A065A
	// PatientSupportPositionParameterSequenceTag is tag for Patient Support Position Parameter Sequence
	PatientSupportPositionParameterSequenceTag = 0x300A065B
	// PatientSupportPositionSpecificationMethodTag is tag for Patient Support Position Specification Method
	PatientSupportPositionSpecificationMethodTag = 0x300A065C
	// PatientSupportPositionDeviceParameterSequenceTag is tag for Patient Support Position Device Parameter Sequence
	PatientSupportPositionDeviceParameterSequenceTag = 0x300A065D
	// DeviceOrderIndexTag is tag for Device Order Index
	DeviceOrderIndexTag = 0x300A065E
	// PatientSupportPositionParameterOrderIndexTag is tag for Patient Support Position Parameter Order Index
	PatientSupportPositionParameterOrderIndexTag = 0x300A065F
	// PatientSupportPositionDeviceToleranceSequenceTag is tag for Patient Support Position Device Tolerance Sequence
	PatientSupportPositionDeviceToleranceSequenceTag = 0x300A0660
	// PatientSupportPositionToleranceOrderIndexTag is tag for Patient Support Position Tolerance Order Index
	PatientSupportPositionToleranceOrderIndexTag = 0x300A0661
	// CompensatorDefinitionSequenceTag is tag for Compensator Definition Sequence
	CompensatorDefinitionSequenceTag = 0x300A0662
	// CompensatorMapOrientationTag is tag for Compensator Map Orientation
	CompensatorMapOrientationTag = 0x300A0663
	// CompensatorProximalThicknessMapTag is tag for Compensator Proximal Thickness Map
	CompensatorProximalThicknessMapTag = 0x300A0664
	// CompensatorDistalThicknessMapTag is tag for Compensator Distal Thickness Map
	CompensatorDistalThicknessMapTag = 0x300A0665
	// CompensatorBasePlaneOffsetTag is tag for Compensator Base Plane Offset
	CompensatorBasePlaneOffsetTag = 0x300A0666
	// CompensatorShapeFabricationCodeSequenceTag is tag for Compensator Shape Fabrication Code Sequence
	CompensatorShapeFabricationCodeSequenceTag = 0x300A0667
	// CompensatorShapeSequenceTag is tag for Compensator Shape Sequence
	CompensatorShapeSequenceTag = 0x300A0668
	// RadiationBeamCompensatorMillingToolDiameterTag is tag for Radiation Beam Compensator Milling Tool Diameter
	RadiationBeamCompensatorMillingToolDiameterTag = 0x300A0669
	// BlockDefinitionSequenceTag is tag for Block Definition Sequence
	BlockDefinitionSequenceTag = 0x300A066A
	// BlockEdgeDataTag is tag for Block Edge Data
	BlockEdgeDataTag = 0x300A066B
	// BlockOrientationTag is tag for Block Orientation
	BlockOrientationTag = 0x300A066C
	// RadiationBeamBlockThicknessTag is tag for Radiation Beam Block Thickness
	RadiationBeamBlockThicknessTag = 0x300A066D
	// RadiationBeamBlockSlabThicknessTag is tag for Radiation Beam Block Slab Thickness
	RadiationBeamBlockSlabThicknessTag = 0x300A066E
	// BlockEdgeDataSequenceTag is tag for Block Edge Data Sequence
	BlockEdgeDataSequenceTag = 0x300A066F
	// NumberOfRTAccessoryHoldersTag is tag for Number of RT Accessory Holders
	NumberOfRTAccessoryHoldersTag = 0x300A0670
	// GeneralAccessoryDefinitionSequenceTag is tag for General Accessory Definition Sequence
	GeneralAccessoryDefinitionSequenceTag = 0x300A0671
	// NumberOfGeneralAccessoriesTag is tag for Number of General Accessories
	NumberOfGeneralAccessoriesTag = 0x300A0672
	// BolusDefinitionSequenceTag is tag for Bolus Definition Sequence
	BolusDefinitionSequenceTag = 0x300A0673
	// NumberOfBolusesTag is tag for Number of Boluses
	NumberOfBolusesTag = 0x300A0674
	// EquipmentFrameOfReferenceUIDTag is tag for Equipment Frame of Reference UID
	EquipmentFrameOfReferenceUIDTag = 0x300A0675
	// EquipmentFrameOfReferenceDescriptionTag is tag for Equipment Frame of Reference Description
	EquipmentFrameOfReferenceDescriptionTag = 0x300A0676
	// EquipmentReferencePointCoordinatesSequenceTag is tag for Equipment Reference Point Coordinates Sequence
	EquipmentReferencePointCoordinatesSequenceTag = 0x300A0677
	// EquipmentReferencePointCodeSequenceTag is tag for Equipment Reference Point Code Sequence
	EquipmentReferencePointCodeSequenceTag = 0x300A0678
	// RTBeamLimitingDeviceAngleTag is tag for RT Beam Limiting Device Angle
	RTBeamLimitingDeviceAngleTag = 0x300A0679
	// SourceRollAngleTag is tag for Source Roll Angle
	SourceRollAngleTag = 0x300A067A
	// RadiationGenerationModeSequenceTag is tag for Radiation Generation​Mode Sequence
	RadiationGenerationModeSequenceTag = 0x300A067B
	// RadiationGenerationModeLabelTag is tag for Radiation Generation​Mode Label
	RadiationGenerationModeLabelTag = 0x300A067C
	// RadiationGenerationModeDescriptionTag is tag for Radiation Generation​Mode Description
	RadiationGenerationModeDescriptionTag = 0x300A067D
	// RadiationGenerationModeMachineCodeSequenceTag is tag for Radiation Generation​Mode Machine Code Sequence
	RadiationGenerationModeMachineCodeSequenceTag = 0x300A067E
	// RadiationTypeCodeSequenceTag is tag for Radiation Type Code Sequence
	RadiationTypeCodeSequenceTag = 0x300A067F
	// NominalEnergyTag is tag for Nominal Energy
	NominalEnergyTag = 0x300A0680
	// MinimumNominalEnergyTag is tag for Minimum Nominal Energy
	MinimumNominalEnergyTag = 0x300A0681
	// MaximumNominalEnergyTag is tag for Maximum Nominal Energy
	MaximumNominalEnergyTag = 0x300A0682
	// RadiationFluenceModifierCodeSequenceTag is tag for Radiation Fluence Modifier Code Sequence
	RadiationFluenceModifierCodeSequenceTag = 0x300A0683
	// EnergyUnitCodeSequenceTag is tag for Energy Unit Code Sequence
	EnergyUnitCodeSequenceTag = 0x300A0684
	// NumberOfRadiationGenerationModesTag is tag for Number of Radiation Generation​Modes
	NumberOfRadiationGenerationModesTag = 0x300A0685
	// PatientSupportDevicesSequenceTag is tag for Patient Support Devices Sequence
	PatientSupportDevicesSequenceTag = 0x300A0686
	// NumberOfPatientSupportDevicesTag is tag for Number of Patient Support Devices
	NumberOfPatientSupportDevicesTag = 0x300A0687
	// RTBeamModifierDefinitionDistanceTag is tag for RT Beam Modifier Definition Distance
	RTBeamModifierDefinitionDistanceTag = 0x300A0688
	// BeamAreaLimitSequenceTag is tag for Beam Area Limit Sequence
	BeamAreaLimitSequenceTag = 0x300A0689
	// ReferencedRTPrescriptionSequenceTag is tag for Referenced RT Prescription Sequence
	ReferencedRTPrescriptionSequenceTag = 0x300A068A
	// ReferencedRTPlanSequenceTag is tag for Referenced RT Plan Sequence
	ReferencedRTPlanSequenceTag = 0x300C0002
	// ReferencedBeamSequenceTag is tag for Referenced Beam Sequence
	ReferencedBeamSequenceTag = 0x300C0004
	// ReferencedBeamNumberTag is tag for Referenced Beam Number
	ReferencedBeamNumberTag = 0x300C0006
	// ReferencedReferenceImageNumberTag is tag for Referenced Reference Image Number
	ReferencedReferenceImageNumberTag = 0x300C0007
	// StartCumulativeMetersetWeightTag is tag for Start Cumulative Meterset Weight
	StartCumulativeMetersetWeightTag = 0x300C0008
	// EndCumulativeMetersetWeightTag is tag for End Cumulative Meterset Weight
	EndCumulativeMetersetWeightTag = 0x300C0009
	// ReferencedBrachyApplicationSetupSequenceTag is tag for Referenced Brachy Application Setup Sequence
	ReferencedBrachyApplicationSetupSequenceTag = 0x300C000A
	// ReferencedBrachyApplicationSetupNumberTag is tag for Referenced Brachy Application Setup Number
	ReferencedBrachyApplicationSetupNumberTag = 0x300C000C
	// ReferencedSourceNumberTag is tag for Referenced Source Number
	ReferencedSourceNumberTag = 0x300C000E
	// ReferencedFractionGroupSequenceTag is tag for Referenced Fraction Group Sequence
	ReferencedFractionGroupSequenceTag = 0x300C0020
	// ReferencedFractionGroupNumberTag is tag for Referenced Fraction Group Number
	ReferencedFractionGroupNumberTag = 0x300C0022
	// ReferencedVerificationImageSequenceTag is tag for Referenced Verification Image Sequence
	ReferencedVerificationImageSequenceTag = 0x300C0040
	// ReferencedReferenceImageSequenceTag is tag for Referenced Reference Image Sequence
	ReferencedReferenceImageSequenceTag = 0x300C0042
	// ReferencedDoseReferenceSequenceTag is tag for Referenced Dose Reference Sequence
	ReferencedDoseReferenceSequenceTag = 0x300C0050
	// ReferencedDoseReferenceNumberTag is tag for Referenced Dose Reference Number
	ReferencedDoseReferenceNumberTag = 0x300C0051
	// BrachyReferencedDoseReferenceSequenceTag is tag for Brachy Referenced Dose Reference Sequence
	BrachyReferencedDoseReferenceSequenceTag = 0x300C0055
	// ReferencedStructureSetSequenceTag is tag for Referenced Structure Set Sequence
	ReferencedStructureSetSequenceTag = 0x300C0060
	// ReferencedPatientSetupNumberTag is tag for Referenced Patient Setup Number
	ReferencedPatientSetupNumberTag = 0x300C006A
	// ReferencedDoseSequenceTag is tag for Referenced Dose Sequence
	ReferencedDoseSequenceTag = 0x300C0080
	// ReferencedToleranceTableNumberTag is tag for Referenced Tolerance Table Number
	ReferencedToleranceTableNumberTag = 0x300C00A0
	// ReferencedBolusSequenceTag is tag for Referenced Bolus Sequence
	ReferencedBolusSequenceTag = 0x300C00B0
	// ReferencedWedgeNumberTag is tag for Referenced Wedge Number
	ReferencedWedgeNumberTag = 0x300C00C0
	// ReferencedCompensatorNumberTag is tag for Referenced Compensator Number
	ReferencedCompensatorNumberTag = 0x300C00D0
	// ReferencedBlockNumberTag is tag for Referenced Block Number
	ReferencedBlockNumberTag = 0x300C00E0
	// ReferencedControlPointIndexTag is tag for Referenced Control Point Index
	ReferencedControlPointIndexTag = 0x300C00F0
	// ReferencedControlPointSequenceTag is tag for Referenced Control Point Sequence
	ReferencedControlPointSequenceTag = 0x300C00F2
	// ReferencedStartControlPointIndexTag is tag for Referenced Start Control Point Index
	ReferencedStartControlPointIndexTag = 0x300C00F4
	// ReferencedStopControlPointIndexTag is tag for Referenced Stop Control Point Index
	ReferencedStopControlPointIndexTag = 0x300C00F6
	// ReferencedRangeShifterNumberTag is tag for Referenced Range Shifter Number
	ReferencedRangeShifterNumberTag = 0x300C0100
	// ReferencedLateralSpreadingDeviceNumberTag is tag for Referenced Lateral Spreading Device Number
	ReferencedLateralSpreadingDeviceNumberTag = 0x300C0102
	// ReferencedRangeModulatorNumberTag is tag for Referenced Range Modulator Number
	ReferencedRangeModulatorNumberTag = 0x300C0104
	// OmittedBeamTaskSequenceTag is tag for Omitted Beam Task Sequence
	OmittedBeamTaskSequenceTag = 0x300C0111
	// ReasonForOmissionTag is tag for Reason for Omission
	ReasonForOmissionTag = 0x300C0112
	// ReasonForOmissionDescriptionTag is tag for Reason for Omission Description
	ReasonForOmissionDescriptionTag = 0x300C0113
	// ApprovalStatusTag is tag for Approval Status
	ApprovalStatusTag = 0x300E0002
	// ReviewDateTag is tag for Review Date
	ReviewDateTag = 0x300E0004
	// ReviewTimeTag is tag for Review Time
	ReviewTimeTag = 0x300E0005
	// ReviewerNameTag is tag for Reviewer Name
	ReviewerNameTag = 0x300E0008
	// RadiobiologicalDoseEffectSequenceTag is tag for Radiobiological Dose Effect Sequence
	RadiobiologicalDoseEffectSequenceTag = 0x30100001
	// RadiobiologicalDoseEffectFlagTag is tag for Radiobiological Dose Effect Flag
	RadiobiologicalDoseEffectFlagTag = 0x30100002
	// EffectiveDoseCalculationMethodCategoryCodeSequenceTag is tag for Effective Dose Calculation Method Category Code Sequence
	EffectiveDoseCalculationMethodCategoryCodeSequenceTag = 0x30100003
	// EffectiveDoseCalculationMethodCodeSequenceTag is tag for Effective Dose Calculation Method Code Sequence
	EffectiveDoseCalculationMethodCodeSequenceTag = 0x30100004
	// EffectiveDoseCalculationMethodDescriptionTag is tag for Effective Dose Calculation Method Description
	EffectiveDoseCalculationMethodDescriptionTag = 0x30100005
	// ConceptualVolumeUIDTag is tag for Conceptual Volume UID
	ConceptualVolumeUIDTag = 0x30100006
	// OriginatingSOPInstanceReferenceSequenceTag is tag for Originating SOP Instance Reference Sequence
	OriginatingSOPInstanceReferenceSequenceTag = 0x30100007
	// ConceptualVolumeConstituentSequenceTag is tag for Conceptual Volume Constituent Sequence
	ConceptualVolumeConstituentSequenceTag = 0x30100008
	// EquivalentConceptualVolumeInstanceReferenceSequenceTag is tag for Equivalent Conceptual Volume Instance Reference Sequence
	EquivalentConceptualVolumeInstanceReferenceSequenceTag = 0x30100009
	// EquivalentConceptualVolumesSequenceTag is tag for Equivalent Conceptual Volumes Sequence
	EquivalentConceptualVolumesSequenceTag = 0x3010000A
	// ReferencedConceptualVolumeUIDTag is tag for Referenced Conceptual Volume UID
	ReferencedConceptualVolumeUIDTag = 0x3010000B
	// ConceptualVolumeCombinationExpressionTag is tag for Conceptual Volume Combination Expression
	ConceptualVolumeCombinationExpressionTag = 0x3010000C
	// ConceptualVolumeConstituentIndexTag is tag for Conceptual Volume Constituent Index
	ConceptualVolumeConstituentIndexTag = 0x3010000D
	// ConceptualVolumeCombinationFlagTag is tag for Conceptual Volume Combination Flag
	ConceptualVolumeCombinationFlagTag = 0x3010000E
	// ConceptualVolumeCombinationDescriptionTag is tag for Conceptual Volume Combination Description
	ConceptualVolumeCombinationDescriptionTag = 0x3010000F
	// ConceptualVolumeSegmentationDefinedFlagTag is tag for Conceptual Volume Segmentation Defined Flag
	ConceptualVolumeSegmentationDefinedFlagTag = 0x30100010
	// ConceptualVolumeSegmentationReferenceSequenceTag is tag for Conceptual Volume Segmentation Reference Sequence
	ConceptualVolumeSegmentationReferenceSequenceTag = 0x30100011
	// ConceptualVolumeConstituentSegmentationReferenceSequenceTag is tag for Conceptual Volume Constituent Segmentation Reference Sequence
	ConceptualVolumeConstituentSegmentationReferenceSequenceTag = 0x30100012
	// ConstituentConceptualVolumeUIDTag is tag for Constituent Conceptual Volume UID
	ConstituentConceptualVolumeUIDTag = 0x30100013
	// DerivationConceptualVolumeSequenceTag is tag for Derivation Conceptual Volume Sequence
	DerivationConceptualVolumeSequenceTag = 0x30100014
	// SourceConceptualVolumeUIDTag is tag for Source Conceptual Volume UID
	SourceConceptualVolumeUIDTag = 0x30100015
	// ConceptualVolumeDerivationAlgorithmSequenceTag is tag for Conceptual Volume Derivation Algorithm Sequence
	ConceptualVolumeDerivationAlgorithmSequenceTag = 0x30100016
	// ConceptualVolumeDescriptionTag is tag for Conceptual Volume Description
	ConceptualVolumeDescriptionTag = 0x30100017
	// SourceConceptualVolumeSequenceTag is tag for Source Conceptual Volume Sequence
	SourceConceptualVolumeSequenceTag = 0x30100018
	// AuthorIdentificationSequenceTag is tag for Author Identification Sequence
	AuthorIdentificationSequenceTag = 0x30100019
	// ManufacturerModelVersionTag is tag for Manufacturer's Model Version
	ManufacturerModelVersionTag = 0x3010001A
	// DeviceAlternateIdentifierTag is tag for Device Alternate Identifier
	DeviceAlternateIdentifierTag = 0x3010001B
	// DeviceAlternateIdentifierTypeTag is tag for Device Alternate Identifier Type
	DeviceAlternateIdentifierTypeTag = 0x3010001C
	// DeviceAlternateIdentifierFormatTag is tag for Device Alternate Identifier Format
	DeviceAlternateIdentifierFormatTag = 0x3010001D
	// SegmentationCreationTemplateLabelTag is tag for Segmentation Creation Template Label
	SegmentationCreationTemplateLabelTag = 0x3010001E
	// SegmentationTemplateUIDTag is tag for Segmentation Template UID
	SegmentationTemplateUIDTag = 0x3010001F
	// ReferencedSegmentReferenceIndexTag is tag for Referenced Segment Reference Index
	ReferencedSegmentReferenceIndexTag = 0x30100020
	// SegmentReferenceSequenceTag is tag for Segment Reference Sequence
	SegmentReferenceSequenceTag = 0x30100021
	// SegmentReferenceIndexTag is tag for Segment Reference Index
	SegmentReferenceIndexTag = 0x30100022
	// DirectSegmentReferenceSequenceTag is tag for Direct Segment Reference Sequence
	DirectSegmentReferenceSequenceTag = 0x30100023
	// CombinationSegmentReferenceSequenceTag is tag for Combination Segment Reference Sequence
	CombinationSegmentReferenceSequenceTag = 0x30100024
	// ConceptualVolumeSequenceTag is tag for Conceptual Volume Sequence
	ConceptualVolumeSequenceTag = 0x30100025
	// SegmentedRTAccessoryDeviceSequenceTag is tag for Segmented RT Accessory Device Sequence
	SegmentedRTAccessoryDeviceSequenceTag = 0x30100026
	// SegmentCharacteristicsSequenceTag is tag for Segment Characteristics Sequence
	SegmentCharacteristicsSequenceTag = 0x30100027
	// RelatedSegmentCharacteristicsSequenceTag is tag for Related Segment Characteristics Sequence
	RelatedSegmentCharacteristicsSequenceTag = 0x30100028
	// SegmentCharacteristicsPrecedenceTag is tag for Segment Characteristics Precedence
	SegmentCharacteristicsPrecedenceTag = 0x30100029
	// RTSegmentAnnotationSequenceTag is tag for RT Segment Annotation Sequence
	RTSegmentAnnotationSequenceTag = 0x3010002A
	// SegmentAnnotationCategoryCodeSequenceTag is tag for Segment Annotation Category Code Sequence
	SegmentAnnotationCategoryCodeSequenceTag = 0x3010002B
	// SegmentAnnotationTypeCodeSequenceTag is tag for Segment Annotation Type Code Sequence
	SegmentAnnotationTypeCodeSequenceTag = 0x3010002C
	// DeviceLabelTag is tag for Device Label
	DeviceLabelTag = 0x3010002D
	// DeviceTypeCodeSequenceTag is tag for Device Type Code Sequence
	DeviceTypeCodeSequenceTag = 0x3010002E
	// SegmentAnnotationTypeModifierCodeSequenceTag is tag for Segment Annotation Type Modifier Code Sequence
	SegmentAnnotationTypeModifierCodeSequenceTag = 0x3010002F
	// PatientEquipmentRelationshipCodeSequenceTag is tag for Patient Equipment Relationship Code Sequence
	PatientEquipmentRelationshipCodeSequenceTag = 0x30100030
	// ReferencedFiducialsUIDTag is tag for Referenced Fiducials UID
	ReferencedFiducialsUIDTag = 0x30100031
	// PatientTreatmentOrientationSequenceTag is tag for Patient Treatment Orientation Sequence
	PatientTreatmentOrientationSequenceTag = 0x30100032
	// UserContentLabelTag is tag for User Content Label
	UserContentLabelTag = 0x30100033
	// UserContentLongLabelTag is tag for User Content Long Label
	UserContentLongLabelTag = 0x30100034
	// EntityLabelTag is tag for Entity Label
	EntityLabelTag = 0x30100035
	// EntityNameTag is tag for Entity Name
	EntityNameTag = 0x30100036
	// EntityDescriptionTag is tag for Entity Description
	EntityDescriptionTag = 0x30100037
	// EntityLongLabelTag is tag for Entity Long Label
	EntityLongLabelTag = 0x30100038
	// DeviceIndexTag is tag for Device Index
	DeviceIndexTag = 0x30100039
	// RTTreatmentPhaseIndexTag is tag for RT Treatment Phase Index
	RTTreatmentPhaseIndexTag = 0x3010003A
	// RTTreatmentPhaseUIDTag is tag for RT Treatment Phase UID
	RTTreatmentPhaseUIDTag = 0x3010003B
	// RTPrescriptionIndexTag is tag for RT Prescription Index
	RTPrescriptionIndexTag = 0x3010003C
	// RTSegmentAnnotationIndexTag is tag for RT Segment Annotation Index
	RTSegmentAnnotationIndexTag = 0x3010003D
	// BasisRTTreatmentPhaseIndexTag is tag for Basis RT Treatment Phase Index
	BasisRTTreatmentPhaseIndexTag = 0x3010003E
	// RelatedRTTreatmentPhaseIndexTag is tag for Related RT Treatment Phase Index
	RelatedRTTreatmentPhaseIndexTag = 0x3010003F
	// ReferencedRTTreatmentPhaseIndexTag is tag for Referenced RT Treatment Phase Index
	ReferencedRTTreatmentPhaseIndexTag = 0x30100040
	// ReferencedRTPrescriptionIndexTag is tag for Referenced RT Prescription Index
	ReferencedRTPrescriptionIndexTag = 0x30100041
	// ReferencedParentRTPrescriptionIndexTag is tag for Referenced Parent RT Prescription Index
	ReferencedParentRTPrescriptionIndexTag = 0x30100042
	// ManufacturerDeviceIdentifierTag is tag for Manufacturer's Device Identifier
	ManufacturerDeviceIdentifierTag = 0x30100043
	// InstanceLevelReferencedPerformedProcedureStepSequenceTag is tag for Instance-Level Referenced Performed Procedure Step Sequence
	InstanceLevelReferencedPerformedProcedureStepSequenceTag = 0x30100044
	// RTTreatmentPhaseIntentPresenceFlagTag is tag for RT Treatment Phase Intent Presence Flag
	RTTreatmentPhaseIntentPresenceFlagTag = 0x30100045
	// RadiotherapyTreatmentTypeTag is tag for Radiotherapy Treatment Type
	RadiotherapyTreatmentTypeTag = 0x30100046
	// TeletherapyRadiationTypeTag is tag for Teletherapy Radiation Type
	TeletherapyRadiationTypeTag = 0x30100047
	// BrachytherapySourceTypeTag is tag for Brachytherapy Source Type
	BrachytherapySourceTypeTag = 0x30100048
	// ReferencedRTTreatmentPhaseSequenceTag is tag for Referenced RT Treatment Phase Sequence
	ReferencedRTTreatmentPhaseSequenceTag = 0x30100049
	// ReferencedDirectSegmentInstanceSequenceTag is tag for Referenced Direct Segment Instance Sequence
	ReferencedDirectSegmentInstanceSequenceTag = 0x3010004A
	// IntendedRTTreatmentPhaseSequenceTag is tag for Intended RT Treatment Phase Sequence
	IntendedRTTreatmentPhaseSequenceTag = 0x3010004B
	// IntendedPhaseStartDateTag is tag for Intended Phase Start Date
	IntendedPhaseStartDateTag = 0x3010004C
	// IntendedPhaseEndDateTag is tag for Intended Phase End Date
	IntendedPhaseEndDateTag = 0x3010004D
	// RTTreatmentPhaseIntervalSequenceTag is tag for RT Treatment Phase Interval Sequence
	RTTreatmentPhaseIntervalSequenceTag = 0x3010004E
	// TemporalRelationshipIntervalAnchorTag is tag for Temporal Relationship Interval Anchor
	TemporalRelationshipIntervalAnchorTag = 0x3010004F
	// MinimumNumberOfIntervalDaysTag is tag for Minimum Number of Interval Days
	MinimumNumberOfIntervalDaysTag = 0x30100050
	// MaximumNumberOfIntervalDaysTag is tag for Maximum Number of Interval Days
	MaximumNumberOfIntervalDaysTag = 0x30100051
	// PertinentSOPClassesInStudyTag is tag for Pertinent SOP Classes in Study
	PertinentSOPClassesInStudyTag = 0x30100052
	// PertinentSOPClassesInSeriesTag is tag for Pertinent SOP Classes in Series
	PertinentSOPClassesInSeriesTag = 0x30100053
	// RTPrescriptionLabelTag is tag for RT Prescription Label
	RTPrescriptionLabelTag = 0x30100054
	// RTPhysicianIntentPredecessorSequenceTag is tag for RT Physician Intent Predecessor Sequence
	RTPhysicianIntentPredecessorSequenceTag = 0x30100055
	// RTTreatmentApproachLabelTag is tag for RT Treatment Approach Label
	RTTreatmentApproachLabelTag = 0x30100056
	// RTPhysicianIntentSequenceTag is tag for RT Physician Intent Sequence
	RTPhysicianIntentSequenceTag = 0x30100057
	// RTPhysicianIntentIndexTag is tag for RT Physician Intent Index
	RTPhysicianIntentIndexTag = 0x30100058
	// RTTreatmentIntentTypeTag is tag for RT Treatment Intent Type
	RTTreatmentIntentTypeTag = 0x30100059
	// RTPhysicianIntentNarrativeTag is tag for RT Physician Intent Narrative
	RTPhysicianIntentNarrativeTag = 0x3010005A
	// RTProtocolCodeSequenceTag is tag for RT Protocol Code Sequence
	RTProtocolCodeSequenceTag = 0x3010005B
	// ReasonForSupersedingTag is tag for Reason for Superseding
	ReasonForSupersedingTag = 0x3010005C
	// RTDiagnosisCodeSequenceTag is tag for RT Diagnosis Code Sequence
	RTDiagnosisCodeSequenceTag = 0x3010005D
	// ReferencedRTPhysicianIntentIndexTag is tag for Referenced RT Physician Intent Index
	ReferencedRTPhysicianIntentIndexTag = 0x3010005E
	// RTPhysicianIntentInputInstanceSequenceTag is tag for RT Physician Intent Input Instance Sequence
	RTPhysicianIntentInputInstanceSequenceTag = 0x3010005F
	// RTAnatomicPrescriptionSequenceTag is tag for RT Anatomic Prescription Sequence
	RTAnatomicPrescriptionSequenceTag = 0x30100060
	// PriorTreatmentDoseDescriptionTag is tag for Prior Treatment Dose Description
	PriorTreatmentDoseDescriptionTag = 0x30100061
	// PriorTreatmentReferenceSequenceTag is tag for Prior Treatment Reference Sequence
	PriorTreatmentReferenceSequenceTag = 0x30100062
	// DosimetricObjectiveEvaluationScopeTag is tag for Dosimetric Objective Evaluation Scope
	DosimetricObjectiveEvaluationScopeTag = 0x30100063
	// TherapeuticRoleCategoryCodeSequenceTag is tag for Therapeutic Role Category Code Sequence
	TherapeuticRoleCategoryCodeSequenceTag = 0x30100064
	// TherapeuticRoleTypeCodeSequenceTag is tag for Therapeutic Role Type Code Sequence
	TherapeuticRoleTypeCodeSequenceTag = 0x30100065
	// ConceptualVolumeOptimizationPrecedenceTag is tag for Conceptual Volume Optimization Precedence
	ConceptualVolumeOptimizationPrecedenceTag = 0x30100066
	// ConceptualVolumeCategoryCodeSequenceTag is tag for Conceptual Volume Category Code Sequence
	ConceptualVolumeCategoryCodeSequenceTag = 0x30100067
	// ConceptualVolumeBlockingConstraintTag is tag for Conceptual Volume Blocking Constraint
	ConceptualVolumeBlockingConstraintTag = 0x30100068
	// ConceptualVolumeTypeCodeSequenceTag is tag for Conceptual Volume Type Code Sequence
	ConceptualVolumeTypeCodeSequenceTag = 0x30100069
	// ConceptualVolumeTypeModifierCodeSequenceTag is tag for Conceptual Volume Type Modifier Code Sequence
	ConceptualVolumeTypeModifierCodeSequenceTag = 0x3010006A
	// RTPrescriptionSequenceTag is tag for RT Prescription Sequence
	RTPrescriptionSequenceTag = 0x3010006B
	// DosimetricObjectiveSequenceTag is tag for Dosimetric Objective Sequence
	DosimetricObjectiveSequenceTag = 0x3010006C
	// DosimetricObjectiveTypeCodeSequenceTag is tag for Dosimetric Objective Type Code Sequence
	DosimetricObjectiveTypeCodeSequenceTag = 0x3010006D
	// DosimetricObjectiveUIDTag is tag for Dosimetric Objective UID
	DosimetricObjectiveUIDTag = 0x3010006E
	// ReferencedDosimetricObjectiveUIDTag is tag for Referenced Dosimetric Objective UID
	ReferencedDosimetricObjectiveUIDTag = 0x3010006F
	// DosimetricObjectiveParameterSequenceTag is tag for Dosimetric Objective Parameter Sequence
	DosimetricObjectiveParameterSequenceTag = 0x30100070
	// ReferencedDosimetricObjectivesSequenceTag is tag for Referenced Dosimetric Objectives Sequence
	ReferencedDosimetricObjectivesSequenceTag = 0x30100071
	// AbsoluteDosimetricObjectiveFlagTag is tag for Absolute Dosimetric Objective Flag
	AbsoluteDosimetricObjectiveFlagTag = 0x30100073
	// DosimetricObjectiveWeightTag is tag for Dosimetric Objective Weight
	DosimetricObjectiveWeightTag = 0x30100074
	// DosimetricObjectivePurposeTag is tag for Dosimetric Objective Purpose
	DosimetricObjectivePurposeTag = 0x30100075
	// PlanningInputInformationSequenceTag is tag for Planning Input Information Sequence
	PlanningInputInformationSequenceTag = 0x30100076
	// TreatmentSiteTag is tag for Treatment Site
	TreatmentSiteTag = 0x30100077
	// TreatmentSiteCodeSequenceTag is tag for Treatment Site Code Sequence
	TreatmentSiteCodeSequenceTag = 0x30100078
	// FractionPatternSequenceTag is tag for Fraction Pattern Sequence
	FractionPatternSequenceTag = 0x30100079
	// TreatmentTechniqueNotesTag is tag for Treatment Technique Notes
	TreatmentTechniqueNotesTag = 0x3010007A
	// PrescriptionNotesTag is tag for Prescription Notes
	PrescriptionNotesTag = 0x3010007B
	// NumberOfIntervalFractionsTag is tag for Number of Interval Fractions
	NumberOfIntervalFractionsTag = 0x3010007C
	// NumberOfFractionsTag is tag for Number of Fractions
	NumberOfFractionsTag = 0x3010007D
	// IntendedDeliveryDurationTag is tag for Intended Delivery Duration
	IntendedDeliveryDurationTag = 0x3010007E
	// FractionationNotesTag is tag for Fractionation Notes
	FractionationNotesTag = 0x3010007F
	// RTTreatmentTechniqueCodeSequenceTag is tag for RT Treatment Technique Code Sequence
	RTTreatmentTechniqueCodeSequenceTag = 0x30100080
	// PrescriptionNotesSequenceTag is tag for Prescription Notes Sequence
	PrescriptionNotesSequenceTag = 0x30100081
	// FractionBasedRelationshipSequenceTag is tag for Fraction-Based Relationship Sequence
	FractionBasedRelationshipSequenceTag = 0x30100082
	// FractionBasedRelationshipIntervalAnchorTag is tag for Fraction-Based Relationship Interval Anchor
	FractionBasedRelationshipIntervalAnchorTag = 0x30100083
	// MinimumHoursBetweenFractionsTag is tag for Minimum Hours between Fractions
	MinimumHoursBetweenFractionsTag = 0x30100084
	// IntendedFractionStartTimeTag is tag for Intended Fraction Start Time
	IntendedFractionStartTimeTag = 0x30100085
	// IntendedStartDayOfWeekTag is tag for Intended Start Day of Week
	IntendedStartDayOfWeekTag = 0x30100086
	// WeekdayFractionPatternSequenceTag is tag for Weekday Fraction Pattern Sequence
	WeekdayFractionPatternSequenceTag = 0x30100087
	// DeliveryTimeStructureCodeSequenceTag is tag for Delivery Time Structure Code Sequence
	DeliveryTimeStructureCodeSequenceTag = 0x30100088
	// TreatmentSiteModifierCodeSequenceTag is tag for Treatment Site Modifier Code Sequence
	TreatmentSiteModifierCodeSequenceTag = 0x30100089
	// ArbitraryTag is tag for Arbitrary
	ArbitraryTag = 0x40000010
	// TextCommentsTag is tag for Text Comments
	TextCommentsTag = 0x40004000
	// ResultsIDTag is tag for Results ID
	ResultsIDTag = 0x40080040
	// ResultsIDIssuerTag is tag for Results ID Issuer
	ResultsIDIssuerTag = 0x40080042
	// ReferencedInterpretationSequenceTag is tag for Referenced Interpretation Sequence
	ReferencedInterpretationSequenceTag = 0x40080050
	// ReportProductionStatusTrialTag is tag for Report Production Status (Trial)
	ReportProductionStatusTrialTag = 0x400800FF
	// InterpretationRecordedDateTag is tag for Interpretation Recorded Date
	InterpretationRecordedDateTag = 0x40080100
	// InterpretationRecordedTimeTag is tag for Interpretation Recorded Time
	InterpretationRecordedTimeTag = 0x40080101
	// InterpretationRecorderTag is tag for Interpretation Recorder
	InterpretationRecorderTag = 0x40080102
	// ReferenceToRecordedSoundTag is tag for Reference to Recorded Sound
	ReferenceToRecordedSoundTag = 0x40080103
	// InterpretationTranscriptionDateTag is tag for Interpretation Transcription Date
	InterpretationTranscriptionDateTag = 0x40080108
	// InterpretationTranscriptionTimeTag is tag for Interpretation Transcription Time
	InterpretationTranscriptionTimeTag = 0x40080109
	// InterpretationTranscriberTag is tag for Interpretation Transcriber
	InterpretationTranscriberTag = 0x4008010A
	// InterpretationTextTag is tag for Interpretation Text
	InterpretationTextTag = 0x4008010B
	// InterpretationAuthorTag is tag for Interpretation Author
	InterpretationAuthorTag = 0x4008010C
	// InterpretationApproverSequenceTag is tag for Interpretation Approver Sequence
	InterpretationApproverSequenceTag = 0x40080111
	// InterpretationApprovalDateTag is tag for Interpretation Approval Date
	InterpretationApprovalDateTag = 0x40080112
	// InterpretationApprovalTimeTag is tag for Interpretation Approval Time
	InterpretationApprovalTimeTag = 0x40080113
	// PhysicianApprovingInterpretationTag is tag for Physician Approving Interpretation
	PhysicianApprovingInterpretationTag = 0x40080114
	// InterpretationDiagnosisDescriptionTag is tag for Interpretation Diagnosis Description
	InterpretationDiagnosisDescriptionTag = 0x40080115
	// InterpretationDiagnosisCodeSequenceTag is tag for Interpretation Diagnosis Code Sequence
	InterpretationDiagnosisCodeSequenceTag = 0x40080117
	// ResultsDistributionListSequenceTag is tag for Results Distribution List Sequence
	ResultsDistributionListSequenceTag = 0x40080118
	// DistributionNameTag is tag for Distribution Name
	DistributionNameTag = 0x40080119
	// DistributionAddressTag is tag for Distribution Address
	DistributionAddressTag = 0x4008011A
	// InterpretationIDTag is tag for Interpretation ID
	InterpretationIDTag = 0x40080200
	// InterpretationIDIssuerTag is tag for Interpretation ID Issuer
	InterpretationIDIssuerTag = 0x40080202
	// InterpretationTypeIDTag is tag for Interpretation Type ID
	InterpretationTypeIDTag = 0x40080210
	// InterpretationStatusIDTag is tag for Interpretation Status ID
	InterpretationStatusIDTag = 0x40080212
	// ImpressionsTag is tag for Impressions
	ImpressionsTag = 0x40080300
	// ResultsCommentsTag is tag for Results Comments
	ResultsCommentsTag = 0x40084000
	// LowEnergyDetectorsTag is tag for Low Energy Detectors
	LowEnergyDetectorsTag = 0x40100001
	// HighEnergyDetectorsTag is tag for High Energy Detectors
	HighEnergyDetectorsTag = 0x40100002
	// DetectorGeometrySequenceTag is tag for Detector Geometry Sequence
	DetectorGeometrySequenceTag = 0x40100004
	// ThreatROIVoxelSequenceTag is tag for Threat ROI Voxel Sequence
	ThreatROIVoxelSequenceTag = 0x40101001
	// ThreatROIBaseTag is tag for Threat ROI Base
	ThreatROIBaseTag = 0x40101004
	// ThreatROIExtentsTag is tag for Threat ROI Extents
	ThreatROIExtentsTag = 0x40101005
	// ThreatROIBitmapTag is tag for Threat ROI Bitmap
	ThreatROIBitmapTag = 0x40101006
	// RouteSegmentIDTag is tag for Route Segment ID
	RouteSegmentIDTag = 0x40101007
	// GantryTypeTag is tag for Gantry Type
	GantryTypeTag = 0x40101008
	// OOIOwnerTypeTag is tag for OOI Owner Type
	OOIOwnerTypeTag = 0x40101009
	// RouteSegmentSequenceTag is tag for Route Segment Sequence
	RouteSegmentSequenceTag = 0x4010100A
	// PotentialThreatObjectIDTag is tag for Potential Threat Object ID
	PotentialThreatObjectIDTag = 0x40101010
	// ThreatSequenceTag is tag for Threat Sequence
	ThreatSequenceTag = 0x40101011
	// ThreatCategoryTag is tag for Threat Category
	ThreatCategoryTag = 0x40101012
	// ThreatCategoryDescriptionTag is tag for Threat Category Description
	ThreatCategoryDescriptionTag = 0x40101013
	// ATDAbilityAssessmentTag is tag for ATD Ability Assessment
	ATDAbilityAssessmentTag = 0x40101014
	// ATDAssessmentFlagTag is tag for ATD Assessment Flag
	ATDAssessmentFlagTag = 0x40101015
	// ATDAssessmentProbabilityTag is tag for ATD Assessment Probability
	ATDAssessmentProbabilityTag = 0x40101016
	// MassTag is tag for Mass
	MassTag = 0x40101017
	// DensityTag is tag for Density
	DensityTag = 0x40101018
	// ZEffectiveTag is tag for Z Effective
	ZEffectiveTag = 0x40101019
	// BoardingPassIDTag is tag for Boarding Pass ID
	BoardingPassIDTag = 0x4010101A
	// CenterOfMassTag is tag for Center of Mass
	CenterOfMassTag = 0x4010101B
	// CenterOfPTOTag is tag for Center of PTO
	CenterOfPTOTag = 0x4010101C
	// BoundingPolygonTag is tag for Bounding Polygon
	BoundingPolygonTag = 0x4010101D
	// RouteSegmentStartLocationIDTag is tag for Route Segment Start Location ID
	RouteSegmentStartLocationIDTag = 0x4010101E
	// RouteSegmentEndLocationIDTag is tag for Route Segment End Location ID
	RouteSegmentEndLocationIDTag = 0x4010101F
	// RouteSegmentLocationIDTypeTag is tag for Route Segment Location ID Type
	RouteSegmentLocationIDTypeTag = 0x40101020
	// AbortReasonTag is tag for Abort Reason
	AbortReasonTag = 0x40101021
	// VolumeOfPTOTag is tag for Volume of PTO
	VolumeOfPTOTag = 0x40101023
	// AbortFlagTag is tag for Abort Flag
	AbortFlagTag = 0x40101024
	// RouteSegmentStartTimeTag is tag for Route Segment Start Time
	RouteSegmentStartTimeTag = 0x40101025
	// RouteSegmentEndTimeTag is tag for Route Segment End Time
	RouteSegmentEndTimeTag = 0x40101026
	// TDRTypeTag is tag for TDR Type
	TDRTypeTag = 0x40101027
	// InternationalRouteSegmentTag is tag for International Route Segment
	InternationalRouteSegmentTag = 0x40101028
	// ThreatDetectionAlgorithmandVersionTag is tag for Threat Detection Algorithm and Version
	ThreatDetectionAlgorithmandVersionTag = 0x40101029
	// AssignedLocationTag is tag for Assigned Location
	AssignedLocationTag = 0x4010102A
	// AlarmDecisionTimeTag is tag for Alarm Decision Time
	AlarmDecisionTimeTag = 0x4010102B
	// AlarmDecisionTag is tag for Alarm Decision
	AlarmDecisionTag = 0x40101031
	// NumberOfTotalObjectsTag is tag for Number of Total Objects
	NumberOfTotalObjectsTag = 0x40101033
	// NumberOfAlarmObjectsTag is tag for Number of Alarm Objects
	NumberOfAlarmObjectsTag = 0x40101034
	// PTORepresentationSequenceTag is tag for PTO Representation Sequence
	PTORepresentationSequenceTag = 0x40101037
	// ATDAssessmentSequenceTag is tag for ATD Assessment Sequence
	ATDAssessmentSequenceTag = 0x40101038
	// TIPTypeTag is tag for TIP Type
	TIPTypeTag = 0x40101039
	// DICOSVersionTag is tag for DICOS Version
	DICOSVersionTag = 0x4010103A
	// OOIOwnerCreationTimeTag is tag for OOI Owner Creation Time
	OOIOwnerCreationTimeTag = 0x40101041
	// OOITypeTag is tag for OOI Type
	OOITypeTag = 0x40101042
	// OOISizeTag is tag for OOI Size
	OOISizeTag = 0x40101043
	// AcquisitionStatusTag is tag for Acquisition Status
	AcquisitionStatusTag = 0x40101044
	// BasisMaterialsCodeSequenceTag is tag for Basis Materials Code Sequence
	BasisMaterialsCodeSequenceTag = 0x40101045
	// PhantomTypeTag is tag for Phantom Type
	PhantomTypeTag = 0x40101046
	// OOIOwnerSequenceTag is tag for OOI Owner Sequence
	OOIOwnerSequenceTag = 0x40101047
	// ScanTypeTag is tag for Scan Type
	ScanTypeTag = 0x40101048
	// ItineraryIDTag is tag for Itinerary ID
	ItineraryIDTag = 0x40101051
	// ItineraryIDTypeTag is tag for Itinerary ID Type
	ItineraryIDTypeTag = 0x40101052
	// ItineraryIDAssigningAuthorityTag is tag for Itinerary ID Assigning Authority
	ItineraryIDAssigningAuthorityTag = 0x40101053
	// RouteIDTag is tag for Route ID
	RouteIDTag = 0x40101054
	// RouteIDAssigningAuthorityTag is tag for Route ID Assigning Authority
	RouteIDAssigningAuthorityTag = 0x40101055
	// InboundArrivalTypeTag is tag for Inbound Arrival Type
	InboundArrivalTypeTag = 0x40101056
	// CarrierIDTag is tag for Carrier ID
	CarrierIDTag = 0x40101058
	// CarrierIDAssigningAuthorityTag is tag for Carrier ID Assigning Authority
	CarrierIDAssigningAuthorityTag = 0x40101059
	// SourceOrientationTag is tag for Source Orientation
	SourceOrientationTag = 0x40101060
	// SourcePositionTag is tag for Source Position
	SourcePositionTag = 0x40101061
	// BeltHeightTag is tag for Belt Height
	BeltHeightTag = 0x40101062
	// AlgorithmRoutingCodeSequenceTag is tag for Algorithm Routing Code Sequence
	AlgorithmRoutingCodeSequenceTag = 0x40101064
	// TransportClassificationTag is tag for Transport Classification
	TransportClassificationTag = 0x40101067
	// OOITypeDescriptorTag is tag for OOI Type Descriptor
	OOITypeDescriptorTag = 0x40101068
	// TotalProcessingTimeTag is tag for Total Processing Time
	TotalProcessingTimeTag = 0x40101069
	// DetectorCalibrationDataTag is tag for Detector Calibration Data
	DetectorCalibrationDataTag = 0x4010106C
	// AdditionalScreeningPerformedTag is tag for Additional Screening Performed
	AdditionalScreeningPerformedTag = 0x4010106D
	// AdditionalInspectionSelectionCriteriaTag is tag for Additional Inspection Selection Criteria
	AdditionalInspectionSelectionCriteriaTag = 0x4010106E
	// AdditionalInspectionMethodSequenceTag is tag for Additional Inspection Method Sequence
	AdditionalInspectionMethodSequenceTag = 0x4010106F
	// AITDeviceTypeTag is tag for AIT Device Type
	AITDeviceTypeTag = 0x40101070
	// QRMeasurementsSequenceTag is tag for QR Measurements Sequence
	QRMeasurementsSequenceTag = 0x40101071
	// TargetMaterialSequenceTag is tag for Target Material Sequence
	TargetMaterialSequenceTag = 0x40101072
	// SNRThresholdTag is tag for SNR Threshold
	SNRThresholdTag = 0x40101073
	// ImageScaleRepresentationTag is tag for Image Scale Representation
	ImageScaleRepresentationTag = 0x40101075
	// ReferencedPTOSequenceTag is tag for Referenced PTO Sequence
	ReferencedPTOSequenceTag = 0x40101076
	// ReferencedTDRInstanceSequenceTag is tag for Referenced TDR Instance Sequence
	ReferencedTDRInstanceSequenceTag = 0x40101077
	// PTOLocationDescriptionTag is tag for PTO Location Description
	PTOLocationDescriptionTag = 0x40101078
	// AnomalyLocatorIndicatorSequenceTag is tag for Anomaly Locator Indicator Sequence
	AnomalyLocatorIndicatorSequenceTag = 0x40101079
	// AnomalyLocatorIndicatorTag is tag for Anomaly Locator Indicator
	AnomalyLocatorIndicatorTag = 0x4010107A
	// PTORegionSequenceTag is tag for PTO Region Sequence
	PTORegionSequenceTag = 0x4010107B
	// InspectionSelectionCriteriaTag is tag for Inspection Selection Criteria
	InspectionSelectionCriteriaTag = 0x4010107C
	// SecondaryInspectionMethodSequenceTag is tag for Secondary Inspection Method Sequence
	SecondaryInspectionMethodSequenceTag = 0x4010107D
	// PRCSToRCSOrientationTag is tag for PRCS to RCS Orientation
	PRCSToRCSOrientationTag = 0x4010107E
	// MACParametersSequenceTag is tag for MAC Parameters Sequence
	MACParametersSequenceTag = 0x4FFE0001
	// CurveDimensionsTag is tag for Curve Dimensions
	CurveDimensionsTag = 0x50000005
	// NumberOfPointsTag is tag for Number of Points
	NumberOfPointsTag = 0x50000010
	// TypeOfDataTag is tag for Type of Data
	TypeOfDataTag = 0x50000020
	// CurveDescriptionTag is tag for Curve Description
	CurveDescriptionTag = 0x50000022
	// AxisUnitsTag is tag for Axis Units
	AxisUnitsTag = 0x50000030
	// AxisLabelsTag is tag for Axis Labels
	AxisLabelsTag = 0x50000040
	// DataValueRepresentationTag is tag for Data Value Representation
	DataValueRepresentationTag = 0x50000103
	// MinimumCoordinateValueTag is tag for Minimum Coordinate Value
	MinimumCoordinateValueTag = 0x50000104
	// MaximumCoordinateValueTag is tag for Maximum Coordinate Value
	MaximumCoordinateValueTag = 0x50000105
	// CurveRangeTag is tag for Curve Range
	CurveRangeTag = 0x50000106
	// CurveDataDescriptorTag is tag for Curve Data Descriptor
	CurveDataDescriptorTag = 0x50000110
	// CoordinateStartValueTag is tag for Coordinate Start Value
	CoordinateStartValueTag = 0x50000112
	// CoordinateStepValueTag is tag for Coordinate Step Value
	CoordinateStepValueTag = 0x50000114
	// CurveActivationLayerTag is tag for Curve Activation Layer
	CurveActivationLayerTag = 0x50001001
	// AudioTypeTag is tag for Audio Type
	AudioTypeTag = 0x50002000
	// AudioSampleFormatTag is tag for Audio Sample Format
	AudioSampleFormatTag = 0x50002002
	// NumberOfChannelsTag is tag for Number of Channels
	NumberOfChannelsTag = 0x50002004
	// NumberOfSamplesTag is tag for Number of Samples
	NumberOfSamplesTag = 0x50002006
	// SampleRateTag is tag for Sample Rate
	SampleRateTag = 0x50002008
	// TotalTimeTag is tag for Total Time
	TotalTimeTag = 0x5000200A
	// AudioSampleDataTag is tag for Audio Sample Data
	AudioSampleDataTag = 0x5000200C
	// AudioCommentsTag is tag for Audio Comments
	AudioCommentsTag = 0x5000200E
	// CurveLabelTag is tag for Curve Label
	CurveLabelTag = 0x50002500
	// CurveReferencedOverlaySequenceTag is tag for Curve Referenced Overlay Sequence
	CurveReferencedOverlaySequenceTag = 0x50002600
	// CurveReferencedOverlayGroupTag is tag for Curve Referenced Overlay Group
	CurveReferencedOverlayGroupTag = 0x50002610
	// CurveDataTag is tag for Curve Data
	CurveDataTag = 0x50003000
	// SharedFunctionalGroupsSequenceTag is tag for Shared Functional Groups Sequence
	SharedFunctionalGroupsSequenceTag = 0x52009229
	// PerFrameFunctionalGroupsSequenceTag is tag for Per-frame Functional Groups Sequence
	PerFrameFunctionalGroupsSequenceTag = 0x52009230
	// WaveformSequenceTag is tag for Waveform Sequence
	WaveformSequenceTag = 0x54000100
	// ChannelMinimumValueTag is tag for Channel Minimum Value
	ChannelMinimumValueTag = 0x54000110
	// ChannelMaximumValueTag is tag for Channel Maximum Value
	ChannelMaximumValueTag = 0x54000112
	// WaveformBitsAllocatedTag is tag for Waveform Bits Allocated
	WaveformBitsAllocatedTag = 0x54001004
	// WaveformSampleInterpretationTag is tag for Waveform Sample Interpretation
	WaveformSampleInterpretationTag = 0x54001006
	// WaveformPaddingValueTag is tag for Waveform Padding Value
	WaveformPaddingValueTag = 0x5400100A
	// WaveformDataTag is tag for Waveform Data
	WaveformDataTag = 0x54001010
	// FirstOrderPhaseCorrectionAngleTag is tag for First Order Phase Correction Angle
	FirstOrderPhaseCorrectionAngleTag = 0x56000010
	// SpectroscopyDataTag is tag for Spectroscopy Data
	SpectroscopyDataTag = 0x56000020
	// OverlayRowsTag is tag for Overlay Rows
	OverlayRowsTag = 0x60000010
	// OverlayColumnsTag is tag for Overlay Columns
	OverlayColumnsTag = 0x60000011
	// OverlayPlanesTag is tag for Overlay Planes
	OverlayPlanesTag = 0x60000012
	// NumberOfFramesInOverlayTag is tag for Number of Frames in Overlay
	NumberOfFramesInOverlayTag = 0x60000015
	// OverlayDescriptionTag is tag for Overlay Description
	OverlayDescriptionTag = 0x60000022
	// OverlayTypeTag is tag for Overlay Type
	OverlayTypeTag = 0x60000040
	// OverlaySubtypeTag is tag for Overlay Subtype
	OverlaySubtypeTag = 0x60000045
	// OverlayOriginTag is tag for Overlay Origin
	OverlayOriginTag = 0x60000050
	// ImageFrameOriginTag is tag for Image Frame Origin
	ImageFrameOriginTag = 0x60000051
	// OverlayPlaneOriginTag is tag for Overlay Plane Origin
	OverlayPlaneOriginTag = 0x60000052
	// OverlayCompressionCodeTag is tag for Overlay Compression Code
	OverlayCompressionCodeTag = 0x60000060
	// OverlayCompressionOriginatorTag is tag for Overlay Compression Originator
	OverlayCompressionOriginatorTag = 0x60000061
	// OverlayCompressionLabelTag is tag for Overlay Compression Label
	OverlayCompressionLabelTag = 0x60000062
	// OverlayCompressionDescriptionTag is tag for Overlay Compression Description
	OverlayCompressionDescriptionTag = 0x60000063
	// OverlayCompressionStepPointersTag is tag for Overlay Compression Step Pointers
	OverlayCompressionStepPointersTag = 0x60000066
	// OverlayRepeatIntervalTag is tag for Overlay Repeat Interval
	OverlayRepeatIntervalTag = 0x60000068
	// OverlayBitsGroupedTag is tag for Overlay Bits Grouped
	OverlayBitsGroupedTag = 0x60000069
	// OverlayBitsAllocatedTag is tag for Overlay Bits Allocated
	OverlayBitsAllocatedTag = 0x60000100
	// OverlayBitPositionTag is tag for Overlay Bit Position
	OverlayBitPositionTag = 0x60000102
	// OverlayFormatTag is tag for Overlay Format
	OverlayFormatTag = 0x60000110
	// OverlayLocationTag is tag for Overlay Location
	OverlayLocationTag = 0x60000200
	// OverlayCodeLabelTag is tag for Overlay Code Label
	OverlayCodeLabelTag = 0x60000800
	// OverlayNumberOfTablesTag is tag for Overlay Number of Tables
	OverlayNumberOfTablesTag = 0x60000802
	// OverlayCodeTableLocationTag is tag for Overlay Code Table Location
	OverlayCodeTableLocationTag = 0x60000803
	// OverlayBitsForCodeWordTag is tag for Overlay Bits For Code Word
	OverlayBitsForCodeWordTag = 0x60000804
	// OverlayActivationLayerTag is tag for Overlay Activation Layer
	OverlayActivationLayerTag = 0x60001001
	// OverlayDescriptorGrayTag is tag for Overlay Descriptor - Gray
	OverlayDescriptorGrayTag = 0x60001100
	// OverlayDescriptorRedTag is tag for Overlay Descriptor - Red
	OverlayDescriptorRedTag = 0x60001101
	// OverlayDescriptorGreenTag is tag for Overlay Descriptor - Green
	OverlayDescriptorGreenTag = 0x60001102
	// OverlayDescriptorBlueTag is tag for Overlay Descriptor - Blue
	OverlayDescriptorBlueTag = 0x60001103
	// OverlaysGrayTag is tag for Overlays - Gray
	OverlaysGrayTag = 0x60001200
	// OverlaysRedTag is tag for Overlays - Red
	OverlaysRedTag = 0x60001201
	// OverlaysGreenTag is tag for Overlays - Green
	OverlaysGreenTag = 0x60001202
	// OverlaysBlueTag is tag for Overlays - Blue
	OverlaysBlueTag = 0x60001203
	// ROIAreaTag is tag for ROI Area
	ROIAreaTag = 0x60001301
	// ROIMeanTag is tag for ROI Mean
	ROIMeanTag = 0x60001302
	// ROIStandardDeviationTag is tag for ROI Standard Deviation
	ROIStandardDeviationTag = 0x60001303
	// OverlayLabelTag is tag for Overlay Label
	OverlayLabelTag = 0x60001500
	// OverlayDataTag is tag for Overlay Data
	OverlayDataTag = 0x60003000
	// OverlayCommentsTag is tag for Overlay Comments
	OverlayCommentsTag = 0x60004000
	// VariablePixelDataTag is tag for Variable Pixel Data
	VariablePixelDataTag = 0x7F000010
	// VariableNextDataGroupTag is tag for Variable Next Data Group
	VariableNextDataGroupTag = 0x7F000011
	// VariableCoefficientsSDVNTag is tag for Variable Coefficients SDVN
	VariableCoefficientsSDVNTag = 0x7F000020
	// VariableCoefficientsSDHNTag is tag for Variable Coefficients SDHN
	VariableCoefficientsSDHNTag = 0x7F000030
	// VariableCoefficientsSDDNTag is tag for Variable Coefficients SDDN
	VariableCoefficientsSDDNTag = 0x7F000040
	// ExtendedOffsetTableTag is tag for Extended Offset Table
	ExtendedOffsetTableTag = 0x7FE00001
	// ExtendedOffsetTableLengthsTag is tag for Extended Offset Table Lengths
	ExtendedOffsetTableLengthsTag = 0x7FE00002
	// FloatPixelDataTag is tag for Float Pixel Data
	FloatPixelDataTag = 0x7FE00008
	// DoubleFloatPixelDataTag is tag for Double Float Pixel Data
	DoubleFloatPixelDataTag = 0x7FE00009
	// PixelDataTag is tag for Pixel Data
	PixelDataTag = 0x7FE00010
	// CoefficientsSDVNTag is tag for Coefficients SDVN
	CoefficientsSDVNTag = 0x7FE00020
	// CoefficientsSDHNTag is tag for Coefficients SDHN
	CoefficientsSDHNTag = 0x7FE00030
	// CoefficientsSDDNTag is tag for Coefficients SDDN
	CoefficientsSDDNTag = 0x7FE00040
	// DigitalSignaturesSequenceTag is tag for Digital Signatures Sequence
	DigitalSignaturesSequenceTag = 0xFFFAFFFA
	// DataSetTrailingPaddingTag is tag for Data Set Trailing Padding
	DataSetTrailingPaddingTag = 0xFFFCFFFC
	// ItemTag is tag for Item
	ItemTag = 0xFFFEE000
	// ItemDelimitationItemTag is tag for Item Delimitation Item
	ItemDelimitationItemTag = 0xFFFEE00D
	// SequenceDelimitationItemTag is tag for Sequence Delimitation Item
	SequenceDelimitationItemTag = 0xFFFEE0DD
)
