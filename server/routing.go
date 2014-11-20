package server

func RegisterRoutes() {

	conditionBase := Router.Path("/Condition").Subrouter()
	conditionBase.Methods("GET").HandlerFunc(ConditionIndexHandler)
	conditionBase.Methods("POST").HandlerFunc(ConditionCreateHandler)

	condition := Router.Path("/Condition/{id}").Subrouter()
	condition.Methods("GET").HandlerFunc(ConditionShowHandler)
	condition.Methods("PUT").HandlerFunc(ConditionUpdateHandler)
	condition.Methods("DELETE").HandlerFunc(ConditionDeleteHandler)

	supplyBase := Router.Path("/Supply").Subrouter()
	supplyBase.Methods("GET").HandlerFunc(SupplyIndexHandler)
	supplyBase.Methods("POST").HandlerFunc(SupplyCreateHandler)

	supply := Router.Path("/Supply/{id}").Subrouter()
	supply.Methods("GET").HandlerFunc(SupplyShowHandler)
	supply.Methods("PUT").HandlerFunc(SupplyUpdateHandler)
	supply.Methods("DELETE").HandlerFunc(SupplyDeleteHandler)

	organizationBase := Router.Path("/Organization").Subrouter()
	organizationBase.Methods("GET").HandlerFunc(OrganizationIndexHandler)
	organizationBase.Methods("POST").HandlerFunc(OrganizationCreateHandler)

	organization := Router.Path("/Organization/{id}").Subrouter()
	organization.Methods("GET").HandlerFunc(OrganizationShowHandler)
	organization.Methods("PUT").HandlerFunc(OrganizationUpdateHandler)
	organization.Methods("DELETE").HandlerFunc(OrganizationDeleteHandler)

	groupBase := Router.Path("/Group").Subrouter()
	groupBase.Methods("GET").HandlerFunc(GroupIndexHandler)
	groupBase.Methods("POST").HandlerFunc(GroupCreateHandler)

	group := Router.Path("/Group/{id}").Subrouter()
	group.Methods("GET").HandlerFunc(GroupShowHandler)
	group.Methods("PUT").HandlerFunc(GroupUpdateHandler)
	group.Methods("DELETE").HandlerFunc(GroupDeleteHandler)

	valuesetBase := Router.Path("/ValueSet").Subrouter()
	valuesetBase.Methods("GET").HandlerFunc(ValueSetIndexHandler)
	valuesetBase.Methods("POST").HandlerFunc(ValueSetCreateHandler)

	valueset := Router.Path("/ValueSet/{id}").Subrouter()
	valueset.Methods("GET").HandlerFunc(ValueSetShowHandler)
	valueset.Methods("PUT").HandlerFunc(ValueSetUpdateHandler)
	valueset.Methods("DELETE").HandlerFunc(ValueSetDeleteHandler)

	immunizationrecommendationBase := Router.Path("/ImmunizationRecommendation").Subrouter()
	immunizationrecommendationBase.Methods("GET").HandlerFunc(ImmunizationRecommendationIndexHandler)
	immunizationrecommendationBase.Methods("POST").HandlerFunc(ImmunizationRecommendationCreateHandler)

	immunizationrecommendation := Router.Path("/ImmunizationRecommendation/{id}").Subrouter()
	immunizationrecommendation.Methods("GET").HandlerFunc(ImmunizationRecommendationShowHandler)
	immunizationrecommendation.Methods("PUT").HandlerFunc(ImmunizationRecommendationUpdateHandler)
	immunizationrecommendation.Methods("DELETE").HandlerFunc(ImmunizationRecommendationDeleteHandler)

	appointmentBase := Router.Path("/Appointment").Subrouter()
	appointmentBase.Methods("GET").HandlerFunc(AppointmentIndexHandler)
	appointmentBase.Methods("POST").HandlerFunc(AppointmentCreateHandler)

	appointment := Router.Path("/Appointment/{id}").Subrouter()
	appointment.Methods("GET").HandlerFunc(AppointmentShowHandler)
	appointment.Methods("PUT").HandlerFunc(AppointmentUpdateHandler)
	appointment.Methods("DELETE").HandlerFunc(AppointmentDeleteHandler)

	medicationdispenseBase := Router.Path("/MedicationDispense").Subrouter()
	medicationdispenseBase.Methods("GET").HandlerFunc(MedicationDispenseIndexHandler)
	medicationdispenseBase.Methods("POST").HandlerFunc(MedicationDispenseCreateHandler)

	medicationdispense := Router.Path("/MedicationDispense/{id}").Subrouter()
	medicationdispense.Methods("GET").HandlerFunc(MedicationDispenseShowHandler)
	medicationdispense.Methods("PUT").HandlerFunc(MedicationDispenseUpdateHandler)
	medicationdispense.Methods("DELETE").HandlerFunc(MedicationDispenseDeleteHandler)

	medicationprescriptionBase := Router.Path("/MedicationPrescription").Subrouter()
	medicationprescriptionBase.Methods("GET").HandlerFunc(MedicationPrescriptionIndexHandler)
	medicationprescriptionBase.Methods("POST").HandlerFunc(MedicationPrescriptionCreateHandler)

	medicationprescription := Router.Path("/MedicationPrescription/{id}").Subrouter()
	medicationprescription.Methods("GET").HandlerFunc(MedicationPrescriptionShowHandler)
	medicationprescription.Methods("PUT").HandlerFunc(MedicationPrescriptionUpdateHandler)
	medicationprescription.Methods("DELETE").HandlerFunc(MedicationPrescriptionDeleteHandler)

	slotBase := Router.Path("/Slot").Subrouter()
	slotBase.Methods("GET").HandlerFunc(SlotIndexHandler)
	slotBase.Methods("POST").HandlerFunc(SlotCreateHandler)

	slot := Router.Path("/Slot/{id}").Subrouter()
	slot.Methods("GET").HandlerFunc(SlotShowHandler)
	slot.Methods("PUT").HandlerFunc(SlotUpdateHandler)
	slot.Methods("DELETE").HandlerFunc(SlotDeleteHandler)

	contraindicationBase := Router.Path("/Contraindication").Subrouter()
	contraindicationBase.Methods("GET").HandlerFunc(ContraindicationIndexHandler)
	contraindicationBase.Methods("POST").HandlerFunc(ContraindicationCreateHandler)

	contraindication := Router.Path("/Contraindication/{id}").Subrouter()
	contraindication.Methods("GET").HandlerFunc(ContraindicationShowHandler)
	contraindication.Methods("PUT").HandlerFunc(ContraindicationUpdateHandler)
	contraindication.Methods("DELETE").HandlerFunc(ContraindicationDeleteHandler)

	appointmentresponseBase := Router.Path("/AppointmentResponse").Subrouter()
	appointmentresponseBase.Methods("GET").HandlerFunc(AppointmentResponseIndexHandler)
	appointmentresponseBase.Methods("POST").HandlerFunc(AppointmentResponseCreateHandler)

	appointmentresponse := Router.Path("/AppointmentResponse/{id}").Subrouter()
	appointmentresponse.Methods("GET").HandlerFunc(AppointmentResponseShowHandler)
	appointmentresponse.Methods("PUT").HandlerFunc(AppointmentResponseUpdateHandler)
	appointmentresponse.Methods("DELETE").HandlerFunc(AppointmentResponseDeleteHandler)

	medicationstatementBase := Router.Path("/MedicationStatement").Subrouter()
	medicationstatementBase.Methods("GET").HandlerFunc(MedicationStatementIndexHandler)
	medicationstatementBase.Methods("POST").HandlerFunc(MedicationStatementCreateHandler)

	medicationstatement := Router.Path("/MedicationStatement/{id}").Subrouter()
	medicationstatement.Methods("GET").HandlerFunc(MedicationStatementShowHandler)
	medicationstatement.Methods("PUT").HandlerFunc(MedicationStatementUpdateHandler)
	medicationstatement.Methods("DELETE").HandlerFunc(MedicationStatementDeleteHandler)

	compositionBase := Router.Path("/Composition").Subrouter()
	compositionBase.Methods("GET").HandlerFunc(CompositionIndexHandler)
	compositionBase.Methods("POST").HandlerFunc(CompositionCreateHandler)

	composition := Router.Path("/Composition/{id}").Subrouter()
	composition.Methods("GET").HandlerFunc(CompositionShowHandler)
	composition.Methods("PUT").HandlerFunc(CompositionUpdateHandler)
	composition.Methods("DELETE").HandlerFunc(CompositionDeleteHandler)

	questionnaireBase := Router.Path("/Questionnaire").Subrouter()
	questionnaireBase.Methods("GET").HandlerFunc(QuestionnaireIndexHandler)
	questionnaireBase.Methods("POST").HandlerFunc(QuestionnaireCreateHandler)

	questionnaire := Router.Path("/Questionnaire/{id}").Subrouter()
	questionnaire.Methods("GET").HandlerFunc(QuestionnaireShowHandler)
	questionnaire.Methods("PUT").HandlerFunc(QuestionnaireUpdateHandler)
	questionnaire.Methods("DELETE").HandlerFunc(QuestionnaireDeleteHandler)

	operationoutcomeBase := Router.Path("/OperationOutcome").Subrouter()
	operationoutcomeBase.Methods("GET").HandlerFunc(OperationOutcomeIndexHandler)
	operationoutcomeBase.Methods("POST").HandlerFunc(OperationOutcomeCreateHandler)

	operationoutcome := Router.Path("/OperationOutcome/{id}").Subrouter()
	operationoutcome.Methods("GET").HandlerFunc(OperationOutcomeShowHandler)
	operationoutcome.Methods("PUT").HandlerFunc(OperationOutcomeUpdateHandler)
	operationoutcome.Methods("DELETE").HandlerFunc(OperationOutcomeDeleteHandler)

	conformanceBase := Router.Path("/Conformance").Subrouter()
	conformanceBase.Methods("GET").HandlerFunc(ConformanceIndexHandler)
	conformanceBase.Methods("POST").HandlerFunc(ConformanceCreateHandler)

	conformance := Router.Path("/Conformance/{id}").Subrouter()
	conformance.Methods("GET").HandlerFunc(ConformanceShowHandler)
	conformance.Methods("PUT").HandlerFunc(ConformanceUpdateHandler)
	conformance.Methods("DELETE").HandlerFunc(ConformanceDeleteHandler)

	mediaBase := Router.Path("/Media").Subrouter()
	mediaBase.Methods("GET").HandlerFunc(MediaIndexHandler)
	mediaBase.Methods("POST").HandlerFunc(MediaCreateHandler)

	media := Router.Path("/Media/{id}").Subrouter()
	media.Methods("GET").HandlerFunc(MediaShowHandler)
	media.Methods("PUT").HandlerFunc(MediaUpdateHandler)
	media.Methods("DELETE").HandlerFunc(MediaDeleteHandler)

	otherBase := Router.Path("/Other").Subrouter()
	otherBase.Methods("GET").HandlerFunc(OtherIndexHandler)
	otherBase.Methods("POST").HandlerFunc(OtherCreateHandler)

	other := Router.Path("/Other/{id}").Subrouter()
	other.Methods("GET").HandlerFunc(OtherShowHandler)
	other.Methods("PUT").HandlerFunc(OtherUpdateHandler)
	other.Methods("DELETE").HandlerFunc(OtherDeleteHandler)

	profileBase := Router.Path("/Profile").Subrouter()
	profileBase.Methods("GET").HandlerFunc(ProfileIndexHandler)
	profileBase.Methods("POST").HandlerFunc(ProfileCreateHandler)

	profile := Router.Path("/Profile/{id}").Subrouter()
	profile.Methods("GET").HandlerFunc(ProfileShowHandler)
	profile.Methods("PUT").HandlerFunc(ProfileUpdateHandler)
	profile.Methods("DELETE").HandlerFunc(ProfileDeleteHandler)

	documentreferenceBase := Router.Path("/DocumentReference").Subrouter()
	documentreferenceBase.Methods("GET").HandlerFunc(DocumentReferenceIndexHandler)
	documentreferenceBase.Methods("POST").HandlerFunc(DocumentReferenceCreateHandler)

	documentreference := Router.Path("/DocumentReference/{id}").Subrouter()
	documentreference.Methods("GET").HandlerFunc(DocumentReferenceShowHandler)
	documentreference.Methods("PUT").HandlerFunc(DocumentReferenceUpdateHandler)
	documentreference.Methods("DELETE").HandlerFunc(DocumentReferenceDeleteHandler)

	immunizationBase := Router.Path("/Immunization").Subrouter()
	immunizationBase.Methods("GET").HandlerFunc(ImmunizationIndexHandler)
	immunizationBase.Methods("POST").HandlerFunc(ImmunizationCreateHandler)

	immunization := Router.Path("/Immunization/{id}").Subrouter()
	immunization.Methods("GET").HandlerFunc(ImmunizationShowHandler)
	immunization.Methods("PUT").HandlerFunc(ImmunizationUpdateHandler)
	immunization.Methods("DELETE").HandlerFunc(ImmunizationDeleteHandler)

	subscriptionBase := Router.Path("/Subscription").Subrouter()
	subscriptionBase.Methods("GET").HandlerFunc(SubscriptionIndexHandler)
	subscriptionBase.Methods("POST").HandlerFunc(SubscriptionCreateHandler)

	subscription := Router.Path("/Subscription/{id}").Subrouter()
	subscription.Methods("GET").HandlerFunc(SubscriptionShowHandler)
	subscription.Methods("PUT").HandlerFunc(SubscriptionUpdateHandler)
	subscription.Methods("DELETE").HandlerFunc(SubscriptionDeleteHandler)

	orderresponseBase := Router.Path("/OrderResponse").Subrouter()
	orderresponseBase.Methods("GET").HandlerFunc(OrderResponseIndexHandler)
	orderresponseBase.Methods("POST").HandlerFunc(OrderResponseCreateHandler)

	orderresponse := Router.Path("/OrderResponse/{id}").Subrouter()
	orderresponse.Methods("GET").HandlerFunc(OrderResponseShowHandler)
	orderresponse.Methods("PUT").HandlerFunc(OrderResponseUpdateHandler)
	orderresponse.Methods("DELETE").HandlerFunc(OrderResponseDeleteHandler)

	conceptmapBase := Router.Path("/ConceptMap").Subrouter()
	conceptmapBase.Methods("GET").HandlerFunc(ConceptMapIndexHandler)
	conceptmapBase.Methods("POST").HandlerFunc(ConceptMapCreateHandler)

	conceptmap := Router.Path("/ConceptMap/{id}").Subrouter()
	conceptmap.Methods("GET").HandlerFunc(ConceptMapShowHandler)
	conceptmap.Methods("PUT").HandlerFunc(ConceptMapUpdateHandler)
	conceptmap.Methods("DELETE").HandlerFunc(ConceptMapDeleteHandler)

	imagingstudyBase := Router.Path("/ImagingStudy").Subrouter()
	imagingstudyBase.Methods("GET").HandlerFunc(ImagingStudyIndexHandler)
	imagingstudyBase.Methods("POST").HandlerFunc(ImagingStudyCreateHandler)

	imagingstudy := Router.Path("/ImagingStudy/{id}").Subrouter()
	imagingstudy.Methods("GET").HandlerFunc(ImagingStudyShowHandler)
	imagingstudy.Methods("PUT").HandlerFunc(ImagingStudyUpdateHandler)
	imagingstudy.Methods("DELETE").HandlerFunc(ImagingStudyDeleteHandler)

	practitionerBase := Router.Path("/Practitioner").Subrouter()
	practitionerBase.Methods("GET").HandlerFunc(PractitionerIndexHandler)
	practitionerBase.Methods("POST").HandlerFunc(PractitionerCreateHandler)

	practitioner := Router.Path("/Practitioner/{id}").Subrouter()
	practitioner.Methods("GET").HandlerFunc(PractitionerShowHandler)
	practitioner.Methods("PUT").HandlerFunc(PractitionerUpdateHandler)
	practitioner.Methods("DELETE").HandlerFunc(PractitionerDeleteHandler)

	careplanBase := Router.Path("/CarePlan").Subrouter()
	careplanBase.Methods("GET").HandlerFunc(CarePlanIndexHandler)
	careplanBase.Methods("POST").HandlerFunc(CarePlanCreateHandler)

	careplan := Router.Path("/CarePlan/{id}").Subrouter()
	careplan.Methods("GET").HandlerFunc(CarePlanShowHandler)
	careplan.Methods("PUT").HandlerFunc(CarePlanUpdateHandler)
	careplan.Methods("DELETE").HandlerFunc(CarePlanDeleteHandler)

	provenanceBase := Router.Path("/Provenance").Subrouter()
	provenanceBase.Methods("GET").HandlerFunc(ProvenanceIndexHandler)
	provenanceBase.Methods("POST").HandlerFunc(ProvenanceCreateHandler)

	provenance := Router.Path("/Provenance/{id}").Subrouter()
	provenance.Methods("GET").HandlerFunc(ProvenanceShowHandler)
	provenance.Methods("PUT").HandlerFunc(ProvenanceUpdateHandler)
	provenance.Methods("DELETE").HandlerFunc(ProvenanceDeleteHandler)

	deviceBase := Router.Path("/Device").Subrouter()
	deviceBase.Methods("GET").HandlerFunc(DeviceIndexHandler)
	deviceBase.Methods("POST").HandlerFunc(DeviceCreateHandler)

	device := Router.Path("/Device/{id}").Subrouter()
	device.Methods("GET").HandlerFunc(DeviceShowHandler)
	device.Methods("PUT").HandlerFunc(DeviceUpdateHandler)
	device.Methods("DELETE").HandlerFunc(DeviceDeleteHandler)

	queryBase := Router.Path("/Query").Subrouter()
	queryBase.Methods("GET").HandlerFunc(QueryIndexHandler)
	queryBase.Methods("POST").HandlerFunc(QueryCreateHandler)

	query := Router.Path("/Query/{id}").Subrouter()
	query.Methods("GET").HandlerFunc(QueryShowHandler)
	query.Methods("PUT").HandlerFunc(QueryUpdateHandler)
	query.Methods("DELETE").HandlerFunc(QueryDeleteHandler)

	orderBase := Router.Path("/Order").Subrouter()
	orderBase.Methods("GET").HandlerFunc(OrderIndexHandler)
	orderBase.Methods("POST").HandlerFunc(OrderCreateHandler)

	order := Router.Path("/Order/{id}").Subrouter()
	order.Methods("GET").HandlerFunc(OrderShowHandler)
	order.Methods("PUT").HandlerFunc(OrderUpdateHandler)
	order.Methods("DELETE").HandlerFunc(OrderDeleteHandler)

	procedureBase := Router.Path("/Procedure").Subrouter()
	procedureBase.Methods("GET").HandlerFunc(ProcedureIndexHandler)
	procedureBase.Methods("POST").HandlerFunc(ProcedureCreateHandler)

	procedure := Router.Path("/Procedure/{id}").Subrouter()
	procedure.Methods("GET").HandlerFunc(ProcedureShowHandler)
	procedure.Methods("PUT").HandlerFunc(ProcedureUpdateHandler)
	procedure.Methods("DELETE").HandlerFunc(ProcedureDeleteHandler)

	substanceBase := Router.Path("/Substance").Subrouter()
	substanceBase.Methods("GET").HandlerFunc(SubstanceIndexHandler)
	substanceBase.Methods("POST").HandlerFunc(SubstanceCreateHandler)

	substance := Router.Path("/Substance/{id}").Subrouter()
	substance.Methods("GET").HandlerFunc(SubstanceShowHandler)
	substance.Methods("PUT").HandlerFunc(SubstanceUpdateHandler)
	substance.Methods("DELETE").HandlerFunc(SubstanceDeleteHandler)

	diagnosticreportBase := Router.Path("/DiagnosticReport").Subrouter()
	diagnosticreportBase.Methods("GET").HandlerFunc(DiagnosticReportIndexHandler)
	diagnosticreportBase.Methods("POST").HandlerFunc(DiagnosticReportCreateHandler)

	diagnosticreport := Router.Path("/DiagnosticReport/{id}").Subrouter()
	diagnosticreport.Methods("GET").HandlerFunc(DiagnosticReportShowHandler)
	diagnosticreport.Methods("PUT").HandlerFunc(DiagnosticReportUpdateHandler)
	diagnosticreport.Methods("DELETE").HandlerFunc(DiagnosticReportDeleteHandler)

	medicationBase := Router.Path("/Medication").Subrouter()
	medicationBase.Methods("GET").HandlerFunc(MedicationIndexHandler)
	medicationBase.Methods("POST").HandlerFunc(MedicationCreateHandler)

	medication := Router.Path("/Medication/{id}").Subrouter()
	medication.Methods("GET").HandlerFunc(MedicationShowHandler)
	medication.Methods("PUT").HandlerFunc(MedicationUpdateHandler)
	medication.Methods("DELETE").HandlerFunc(MedicationDeleteHandler)

	messageheaderBase := Router.Path("/MessageHeader").Subrouter()
	messageheaderBase.Methods("GET").HandlerFunc(MessageHeaderIndexHandler)
	messageheaderBase.Methods("POST").HandlerFunc(MessageHeaderCreateHandler)

	messageheader := Router.Path("/MessageHeader/{id}").Subrouter()
	messageheader.Methods("GET").HandlerFunc(MessageHeaderShowHandler)
	messageheader.Methods("PUT").HandlerFunc(MessageHeaderUpdateHandler)
	messageheader.Methods("DELETE").HandlerFunc(MessageHeaderDeleteHandler)

	documentmanifestBase := Router.Path("/DocumentManifest").Subrouter()
	documentmanifestBase.Methods("GET").HandlerFunc(DocumentManifestIndexHandler)
	documentmanifestBase.Methods("POST").HandlerFunc(DocumentManifestCreateHandler)

	documentmanifest := Router.Path("/DocumentManifest/{id}").Subrouter()
	documentmanifest.Methods("GET").HandlerFunc(DocumentManifestShowHandler)
	documentmanifest.Methods("PUT").HandlerFunc(DocumentManifestUpdateHandler)
	documentmanifest.Methods("DELETE").HandlerFunc(DocumentManifestDeleteHandler)

	dataelementBase := Router.Path("/DataElement").Subrouter()
	dataelementBase.Methods("GET").HandlerFunc(DataElementIndexHandler)
	dataelementBase.Methods("POST").HandlerFunc(DataElementCreateHandler)

	dataelement := Router.Path("/DataElement/{id}").Subrouter()
	dataelement.Methods("GET").HandlerFunc(DataElementShowHandler)
	dataelement.Methods("PUT").HandlerFunc(DataElementUpdateHandler)
	dataelement.Methods("DELETE").HandlerFunc(DataElementDeleteHandler)

	availabilityBase := Router.Path("/Availability").Subrouter()
	availabilityBase.Methods("GET").HandlerFunc(AvailabilityIndexHandler)
	availabilityBase.Methods("POST").HandlerFunc(AvailabilityCreateHandler)

	availability := Router.Path("/Availability/{id}").Subrouter()
	availability.Methods("GET").HandlerFunc(AvailabilityShowHandler)
	availability.Methods("PUT").HandlerFunc(AvailabilityUpdateHandler)
	availability.Methods("DELETE").HandlerFunc(AvailabilityDeleteHandler)

	medicationadministrationBase := Router.Path("/MedicationAdministration").Subrouter()
	medicationadministrationBase.Methods("GET").HandlerFunc(MedicationAdministrationIndexHandler)
	medicationadministrationBase.Methods("POST").HandlerFunc(MedicationAdministrationCreateHandler)

	medicationadministration := Router.Path("/MedicationAdministration/{id}").Subrouter()
	medicationadministration.Methods("GET").HandlerFunc(MedicationAdministrationShowHandler)
	medicationadministration.Methods("PUT").HandlerFunc(MedicationAdministrationUpdateHandler)
	medicationadministration.Methods("DELETE").HandlerFunc(MedicationAdministrationDeleteHandler)

	questionnaireanswersBase := Router.Path("/QuestionnaireAnswers").Subrouter()
	questionnaireanswersBase.Methods("GET").HandlerFunc(QuestionnaireAnswersIndexHandler)
	questionnaireanswersBase.Methods("POST").HandlerFunc(QuestionnaireAnswersCreateHandler)

	questionnaireanswers := Router.Path("/QuestionnaireAnswers/{id}").Subrouter()
	questionnaireanswers.Methods("GET").HandlerFunc(QuestionnaireAnswersShowHandler)
	questionnaireanswers.Methods("PUT").HandlerFunc(QuestionnaireAnswersUpdateHandler)
	questionnaireanswers.Methods("DELETE").HandlerFunc(QuestionnaireAnswersDeleteHandler)

	encounterBase := Router.Path("/Encounter").Subrouter()
	encounterBase.Methods("GET").HandlerFunc(EncounterIndexHandler)
	encounterBase.Methods("POST").HandlerFunc(EncounterCreateHandler)

	encounter := Router.Path("/Encounter/{id}").Subrouter()
	encounter.Methods("GET").HandlerFunc(EncounterShowHandler)
	encounter.Methods("PUT").HandlerFunc(EncounterUpdateHandler)
	encounter.Methods("DELETE").HandlerFunc(EncounterDeleteHandler)

	securityeventBase := Router.Path("/SecurityEvent").Subrouter()
	securityeventBase.Methods("GET").HandlerFunc(SecurityEventIndexHandler)
	securityeventBase.Methods("POST").HandlerFunc(SecurityEventCreateHandler)

	securityevent := Router.Path("/SecurityEvent/{id}").Subrouter()
	securityevent.Methods("GET").HandlerFunc(SecurityEventShowHandler)
	securityevent.Methods("PUT").HandlerFunc(SecurityEventUpdateHandler)
	securityevent.Methods("DELETE").HandlerFunc(SecurityEventDeleteHandler)

	listBase := Router.Path("/List").Subrouter()
	listBase.Methods("GET").HandlerFunc(ListIndexHandler)
	listBase.Methods("POST").HandlerFunc(ListCreateHandler)

	list := Router.Path("/List/{id}").Subrouter()
	list.Methods("GET").HandlerFunc(ListShowHandler)
	list.Methods("PUT").HandlerFunc(ListUpdateHandler)
	list.Methods("DELETE").HandlerFunc(ListDeleteHandler)

	operationdefinitionBase := Router.Path("/OperationDefinition").Subrouter()
	operationdefinitionBase.Methods("GET").HandlerFunc(OperationDefinitionIndexHandler)
	operationdefinitionBase.Methods("POST").HandlerFunc(OperationDefinitionCreateHandler)

	operationdefinition := Router.Path("/OperationDefinition/{id}").Subrouter()
	operationdefinition.Methods("GET").HandlerFunc(OperationDefinitionShowHandler)
	operationdefinition.Methods("PUT").HandlerFunc(OperationDefinitionUpdateHandler)
	operationdefinition.Methods("DELETE").HandlerFunc(OperationDefinitionDeleteHandler)

	deviceobservationreportBase := Router.Path("/DeviceObservationReport").Subrouter()
	deviceobservationreportBase.Methods("GET").HandlerFunc(DeviceObservationReportIndexHandler)
	deviceobservationreportBase.Methods("POST").HandlerFunc(DeviceObservationReportCreateHandler)

	deviceobservationreport := Router.Path("/DeviceObservationReport/{id}").Subrouter()
	deviceobservationreport.Methods("GET").HandlerFunc(DeviceObservationReportShowHandler)
	deviceobservationreport.Methods("PUT").HandlerFunc(DeviceObservationReportUpdateHandler)
	deviceobservationreport.Methods("DELETE").HandlerFunc(DeviceObservationReportDeleteHandler)

	nutritionorderBase := Router.Path("/NutritionOrder").Subrouter()
	nutritionorderBase.Methods("GET").HandlerFunc(NutritionOrderIndexHandler)
	nutritionorderBase.Methods("POST").HandlerFunc(NutritionOrderCreateHandler)

	nutritionorder := Router.Path("/NutritionOrder/{id}").Subrouter()
	nutritionorder.Methods("GET").HandlerFunc(NutritionOrderShowHandler)
	nutritionorder.Methods("PUT").HandlerFunc(NutritionOrderUpdateHandler)
	nutritionorder.Methods("DELETE").HandlerFunc(NutritionOrderDeleteHandler)

	referralrequestBase := Router.Path("/ReferralRequest").Subrouter()
	referralrequestBase.Methods("GET").HandlerFunc(ReferralRequestIndexHandler)
	referralrequestBase.Methods("POST").HandlerFunc(ReferralRequestCreateHandler)

	referralrequest := Router.Path("/ReferralRequest/{id}").Subrouter()
	referralrequest.Methods("GET").HandlerFunc(ReferralRequestShowHandler)
	referralrequest.Methods("PUT").HandlerFunc(ReferralRequestUpdateHandler)
	referralrequest.Methods("DELETE").HandlerFunc(ReferralRequestDeleteHandler)

	riskassessmentBase := Router.Path("/RiskAssessment").Subrouter()
	riskassessmentBase.Methods("GET").HandlerFunc(RiskAssessmentIndexHandler)
	riskassessmentBase.Methods("POST").HandlerFunc(RiskAssessmentCreateHandler)

	riskassessment := Router.Path("/RiskAssessment/{id}").Subrouter()
	riskassessment.Methods("GET").HandlerFunc(RiskAssessmentShowHandler)
	riskassessment.Methods("PUT").HandlerFunc(RiskAssessmentUpdateHandler)
	riskassessment.Methods("DELETE").HandlerFunc(RiskAssessmentDeleteHandler)

	familyhistoryBase := Router.Path("/FamilyHistory").Subrouter()
	familyhistoryBase.Methods("GET").HandlerFunc(FamilyHistoryIndexHandler)
	familyhistoryBase.Methods("POST").HandlerFunc(FamilyHistoryCreateHandler)

	familyhistory := Router.Path("/FamilyHistory/{id}").Subrouter()
	familyhistory.Methods("GET").HandlerFunc(FamilyHistoryShowHandler)
	familyhistory.Methods("PUT").HandlerFunc(FamilyHistoryUpdateHandler)
	familyhistory.Methods("DELETE").HandlerFunc(FamilyHistoryDeleteHandler)

	locationBase := Router.Path("/Location").Subrouter()
	locationBase.Methods("GET").HandlerFunc(LocationIndexHandler)
	locationBase.Methods("POST").HandlerFunc(LocationCreateHandler)

	location := Router.Path("/Location/{id}").Subrouter()
	location.Methods("GET").HandlerFunc(LocationShowHandler)
	location.Methods("PUT").HandlerFunc(LocationUpdateHandler)
	location.Methods("DELETE").HandlerFunc(LocationDeleteHandler)

	allergyintoleranceBase := Router.Path("/AllergyIntolerance").Subrouter()
	allergyintoleranceBase.Methods("GET").HandlerFunc(AllergyIntoleranceIndexHandler)
	allergyintoleranceBase.Methods("POST").HandlerFunc(AllergyIntoleranceCreateHandler)

	allergyintolerance := Router.Path("/AllergyIntolerance/{id}").Subrouter()
	allergyintolerance.Methods("GET").HandlerFunc(AllergyIntoleranceShowHandler)
	allergyintolerance.Methods("PUT").HandlerFunc(AllergyIntoleranceUpdateHandler)
	allergyintolerance.Methods("DELETE").HandlerFunc(AllergyIntoleranceDeleteHandler)

	observationBase := Router.Path("/Observation").Subrouter()
	observationBase.Methods("GET").HandlerFunc(ObservationIndexHandler)
	observationBase.Methods("POST").HandlerFunc(ObservationCreateHandler)

	observation := Router.Path("/Observation/{id}").Subrouter()
	observation.Methods("GET").HandlerFunc(ObservationShowHandler)
	observation.Methods("PUT").HandlerFunc(ObservationUpdateHandler)
	observation.Methods("DELETE").HandlerFunc(ObservationDeleteHandler)

	relatedpersonBase := Router.Path("/RelatedPerson").Subrouter()
	relatedpersonBase.Methods("GET").HandlerFunc(RelatedPersonIndexHandler)
	relatedpersonBase.Methods("POST").HandlerFunc(RelatedPersonCreateHandler)

	relatedperson := Router.Path("/RelatedPerson/{id}").Subrouter()
	relatedperson.Methods("GET").HandlerFunc(RelatedPersonShowHandler)
	relatedperson.Methods("PUT").HandlerFunc(RelatedPersonUpdateHandler)
	relatedperson.Methods("DELETE").HandlerFunc(RelatedPersonDeleteHandler)

	specimenBase := Router.Path("/Specimen").Subrouter()
	specimenBase.Methods("GET").HandlerFunc(SpecimenIndexHandler)
	specimenBase.Methods("POST").HandlerFunc(SpecimenCreateHandler)

	specimen := Router.Path("/Specimen/{id}").Subrouter()
	specimen.Methods("GET").HandlerFunc(SpecimenShowHandler)
	specimen.Methods("PUT").HandlerFunc(SpecimenUpdateHandler)
	specimen.Methods("DELETE").HandlerFunc(SpecimenDeleteHandler)

	alertBase := Router.Path("/Alert").Subrouter()
	alertBase.Methods("GET").HandlerFunc(AlertIndexHandler)
	alertBase.Methods("POST").HandlerFunc(AlertCreateHandler)

	alert := Router.Path("/Alert/{id}").Subrouter()
	alert.Methods("GET").HandlerFunc(AlertShowHandler)
	alert.Methods("PUT").HandlerFunc(AlertUpdateHandler)
	alert.Methods("DELETE").HandlerFunc(AlertDeleteHandler)

	namespaceBase := Router.Path("/Namespace").Subrouter()
	namespaceBase.Methods("GET").HandlerFunc(NamespaceIndexHandler)
	namespaceBase.Methods("POST").HandlerFunc(NamespaceCreateHandler)

	namespace := Router.Path("/Namespace/{id}").Subrouter()
	namespace.Methods("GET").HandlerFunc(NamespaceShowHandler)
	namespace.Methods("PUT").HandlerFunc(NamespaceUpdateHandler)
	namespace.Methods("DELETE").HandlerFunc(NamespaceDeleteHandler)

	patientBase := Router.Path("/Patient").Subrouter()
	patientBase.Methods("GET").HandlerFunc(PatientIndexHandler)
	patientBase.Methods("POST").HandlerFunc(PatientCreateHandler)

	patient := Router.Path("/Patient/{id}").Subrouter()
	patient.Methods("GET").HandlerFunc(PatientShowHandler)
	patient.Methods("PUT").HandlerFunc(PatientUpdateHandler)
	patient.Methods("DELETE").HandlerFunc(PatientDeleteHandler)

	adversereactionBase := Router.Path("/AdverseReaction").Subrouter()
	adversereactionBase.Methods("GET").HandlerFunc(AdverseReactionIndexHandler)
	adversereactionBase.Methods("POST").HandlerFunc(AdverseReactionCreateHandler)

	adversereaction := Router.Path("/AdverseReaction/{id}").Subrouter()
	adversereaction.Methods("GET").HandlerFunc(AdverseReactionShowHandler)
	adversereaction.Methods("PUT").HandlerFunc(AdverseReactionUpdateHandler)
	adversereaction.Methods("DELETE").HandlerFunc(AdverseReactionDeleteHandler)

	diagnosticorderBase := Router.Path("/DiagnosticOrder").Subrouter()
	diagnosticorderBase.Methods("GET").HandlerFunc(DiagnosticOrderIndexHandler)
	diagnosticorderBase.Methods("POST").HandlerFunc(DiagnosticOrderCreateHandler)

	diagnosticorder := Router.Path("/DiagnosticOrder/{id}").Subrouter()
	diagnosticorder.Methods("GET").HandlerFunc(DiagnosticOrderShowHandler)
	diagnosticorder.Methods("PUT").HandlerFunc(DiagnosticOrderUpdateHandler)
	diagnosticorder.Methods("DELETE").HandlerFunc(DiagnosticOrderDeleteHandler)

}
