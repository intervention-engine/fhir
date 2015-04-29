package server

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, config map[string][]negroni.Handler) {

	appointmentBase := router.Path("/Appointment").Subrouter()
	appointmentBase.Methods("GET").Handler(negroni.New(append(config["AppointmentIndex"], negroni.HandlerFunc(AppointmentIndexHandler))...))
	appointmentBase.Methods("POST").Handler(negroni.New(append(config["AppointmentCreate"], negroni.HandlerFunc(AppointmentCreateHandler))...))

	appointment := router.Path("/Appointment/{id}").Subrouter()
	appointment.Methods("GET").Handler(negroni.New(append(config["AppointmentShow"], negroni.HandlerFunc(AppointmentShowHandler))...))
	appointment.Methods("PUT").Handler(negroni.New(append(config["AppointmentUpdate"], negroni.HandlerFunc(AppointmentUpdateHandler))...))
	appointment.Methods("DELETE").Handler(negroni.New(append(config["AppointmentDelete"], negroni.HandlerFunc(AppointmentDeleteHandler))...))

	referralrequestBase := router.Path("/ReferralRequest").Subrouter()
	referralrequestBase.Methods("GET").Handler(negroni.New(append(config["ReferralRequestIndex"], negroni.HandlerFunc(ReferralRequestIndexHandler))...))
	referralrequestBase.Methods("POST").Handler(negroni.New(append(config["ReferralRequestCreate"], negroni.HandlerFunc(ReferralRequestCreateHandler))...))

	referralrequest := router.Path("/ReferralRequest/{id}").Subrouter()
	referralrequest.Methods("GET").Handler(negroni.New(append(config["ReferralRequestShow"], negroni.HandlerFunc(ReferralRequestShowHandler))...))
	referralrequest.Methods("PUT").Handler(negroni.New(append(config["ReferralRequestUpdate"], negroni.HandlerFunc(ReferralRequestUpdateHandler))...))
	referralrequest.Methods("DELETE").Handler(negroni.New(append(config["ReferralRequestDelete"], negroni.HandlerFunc(ReferralRequestDeleteHandler))...))

	provenanceBase := router.Path("/Provenance").Subrouter()
	provenanceBase.Methods("GET").Handler(negroni.New(append(config["ProvenanceIndex"], negroni.HandlerFunc(ProvenanceIndexHandler))...))
	provenanceBase.Methods("POST").Handler(negroni.New(append(config["ProvenanceCreate"], negroni.HandlerFunc(ProvenanceCreateHandler))...))

	provenance := router.Path("/Provenance/{id}").Subrouter()
	provenance.Methods("GET").Handler(negroni.New(append(config["ProvenanceShow"], negroni.HandlerFunc(ProvenanceShowHandler))...))
	provenance.Methods("PUT").Handler(negroni.New(append(config["ProvenanceUpdate"], negroni.HandlerFunc(ProvenanceUpdateHandler))...))
	provenance.Methods("DELETE").Handler(negroni.New(append(config["ProvenanceDelete"], negroni.HandlerFunc(ProvenanceDeleteHandler))...))

	questionnaireBase := router.Path("/Questionnaire").Subrouter()
	questionnaireBase.Methods("GET").Handler(negroni.New(append(config["QuestionnaireIndex"], negroni.HandlerFunc(QuestionnaireIndexHandler))...))
	questionnaireBase.Methods("POST").Handler(negroni.New(append(config["QuestionnaireCreate"], negroni.HandlerFunc(QuestionnaireCreateHandler))...))

	questionnaire := router.Path("/Questionnaire/{id}").Subrouter()
	questionnaire.Methods("GET").Handler(negroni.New(append(config["QuestionnaireShow"], negroni.HandlerFunc(QuestionnaireShowHandler))...))
	questionnaire.Methods("PUT").Handler(negroni.New(append(config["QuestionnaireUpdate"], negroni.HandlerFunc(QuestionnaireUpdateHandler))...))
	questionnaire.Methods("DELETE").Handler(negroni.New(append(config["QuestionnaireDelete"], negroni.HandlerFunc(QuestionnaireDeleteHandler))...))

	queryBase := router.Path("/Query").Subrouter()
	queryBase.Methods("GET").Handler(negroni.New(append(config["QueryIndex"], negroni.HandlerFunc(QueryIndexHandler))...))
	queryBase.Methods("POST").Handler(negroni.New(append(config["QueryCreate"], negroni.HandlerFunc(QueryCreateHandler))...))

	query := router.Path("/Query/{id}").Subrouter()
	query.Methods("GET").Handler(negroni.New(append(config["QueryShow"], negroni.HandlerFunc(QueryShowHandler))...))
	query.Methods("PUT").Handler(negroni.New(append(config["QueryUpdate"], negroni.HandlerFunc(QueryUpdateHandler))...))
	query.Methods("DELETE").Handler(negroni.New(append(config["QueryDelete"], negroni.HandlerFunc(QueryDeleteHandler))...))

	documentmanifestBase := router.Path("/DocumentManifest").Subrouter()
	documentmanifestBase.Methods("GET").Handler(negroni.New(append(config["DocumentManifestIndex"], negroni.HandlerFunc(DocumentManifestIndexHandler))...))
	documentmanifestBase.Methods("POST").Handler(negroni.New(append(config["DocumentManifestCreate"], negroni.HandlerFunc(DocumentManifestCreateHandler))...))

	documentmanifest := router.Path("/DocumentManifest/{id}").Subrouter()
	documentmanifest.Methods("GET").Handler(negroni.New(append(config["DocumentManifestShow"], negroni.HandlerFunc(DocumentManifestShowHandler))...))
	documentmanifest.Methods("PUT").Handler(negroni.New(append(config["DocumentManifestUpdate"], negroni.HandlerFunc(DocumentManifestUpdateHandler))...))
	documentmanifest.Methods("DELETE").Handler(negroni.New(append(config["DocumentManifestDelete"], negroni.HandlerFunc(DocumentManifestDeleteHandler))...))

	specimenBase := router.Path("/Specimen").Subrouter()
	specimenBase.Methods("GET").Handler(negroni.New(append(config["SpecimenIndex"], negroni.HandlerFunc(SpecimenIndexHandler))...))
	specimenBase.Methods("POST").Handler(negroni.New(append(config["SpecimenCreate"], negroni.HandlerFunc(SpecimenCreateHandler))...))

	specimen := router.Path("/Specimen/{id}").Subrouter()
	specimen.Methods("GET").Handler(negroni.New(append(config["SpecimenShow"], negroni.HandlerFunc(SpecimenShowHandler))...))
	specimen.Methods("PUT").Handler(negroni.New(append(config["SpecimenUpdate"], negroni.HandlerFunc(SpecimenUpdateHandler))...))
	specimen.Methods("DELETE").Handler(negroni.New(append(config["SpecimenDelete"], negroni.HandlerFunc(SpecimenDeleteHandler))...))

	allergyintoleranceBase := router.Path("/AllergyIntolerance").Subrouter()
	allergyintoleranceBase.Methods("GET").Handler(negroni.New(append(config["AllergyIntoleranceIndex"], negroni.HandlerFunc(AllergyIntoleranceIndexHandler))...))
	allergyintoleranceBase.Methods("POST").Handler(negroni.New(append(config["AllergyIntoleranceCreate"], negroni.HandlerFunc(AllergyIntoleranceCreateHandler))...))

	allergyintolerance := router.Path("/AllergyIntolerance/{id}").Subrouter()
	allergyintolerance.Methods("GET").Handler(negroni.New(append(config["AllergyIntoleranceShow"], negroni.HandlerFunc(AllergyIntoleranceShowHandler))...))
	allergyintolerance.Methods("PUT").Handler(negroni.New(append(config["AllergyIntoleranceUpdate"], negroni.HandlerFunc(AllergyIntoleranceUpdateHandler))...))
	allergyintolerance.Methods("DELETE").Handler(negroni.New(append(config["AllergyIntoleranceDelete"], negroni.HandlerFunc(AllergyIntoleranceDeleteHandler))...))

	careplanBase := router.Path("/CarePlan").Subrouter()
	careplanBase.Methods("GET").Handler(negroni.New(append(config["CarePlanIndex"], negroni.HandlerFunc(CarePlanIndexHandler))...))
	careplanBase.Methods("POST").Handler(negroni.New(append(config["CarePlanCreate"], negroni.HandlerFunc(CarePlanCreateHandler))...))

	careplan := router.Path("/CarePlan/{id}").Subrouter()
	careplan.Methods("GET").Handler(negroni.New(append(config["CarePlanShow"], negroni.HandlerFunc(CarePlanShowHandler))...))
	careplan.Methods("PUT").Handler(negroni.New(append(config["CarePlanUpdate"], negroni.HandlerFunc(CarePlanUpdateHandler))...))
	careplan.Methods("DELETE").Handler(negroni.New(append(config["CarePlanDelete"], negroni.HandlerFunc(CarePlanDeleteHandler))...))

	medicationprescriptionBase := router.Path("/MedicationPrescription").Subrouter()
	medicationprescriptionBase.Methods("GET").Handler(negroni.New(append(config["MedicationPrescriptionIndex"], negroni.HandlerFunc(MedicationPrescriptionIndexHandler))...))
	medicationprescriptionBase.Methods("POST").Handler(negroni.New(append(config["MedicationPrescriptionCreate"], negroni.HandlerFunc(MedicationPrescriptionCreateHandler))...))

	medicationprescription := router.Path("/MedicationPrescription/{id}").Subrouter()
	medicationprescription.Methods("GET").Handler(negroni.New(append(config["MedicationPrescriptionShow"], negroni.HandlerFunc(MedicationPrescriptionShowHandler))...))
	medicationprescription.Methods("PUT").Handler(negroni.New(append(config["MedicationPrescriptionUpdate"], negroni.HandlerFunc(MedicationPrescriptionUpdateHandler))...))
	medicationprescription.Methods("DELETE").Handler(negroni.New(append(config["MedicationPrescriptionDelete"], negroni.HandlerFunc(MedicationPrescriptionDeleteHandler))...))

	operationoutcomeBase := router.Path("/OperationOutcome").Subrouter()
	operationoutcomeBase.Methods("GET").Handler(negroni.New(append(config["OperationOutcomeIndex"], negroni.HandlerFunc(OperationOutcomeIndexHandler))...))
	operationoutcomeBase.Methods("POST").Handler(negroni.New(append(config["OperationOutcomeCreate"], negroni.HandlerFunc(OperationOutcomeCreateHandler))...))

	operationoutcome := router.Path("/OperationOutcome/{id}").Subrouter()
	operationoutcome.Methods("GET").Handler(negroni.New(append(config["OperationOutcomeShow"], negroni.HandlerFunc(OperationOutcomeShowHandler))...))
	operationoutcome.Methods("PUT").Handler(negroni.New(append(config["OperationOutcomeUpdate"], negroni.HandlerFunc(OperationOutcomeUpdateHandler))...))
	operationoutcome.Methods("DELETE").Handler(negroni.New(append(config["OperationOutcomeDelete"], negroni.HandlerFunc(OperationOutcomeDeleteHandler))...))

	familyhistoryBase := router.Path("/FamilyHistory").Subrouter()
	familyhistoryBase.Methods("GET").Handler(negroni.New(append(config["FamilyHistoryIndex"], negroni.HandlerFunc(FamilyHistoryIndexHandler))...))
	familyhistoryBase.Methods("POST").Handler(negroni.New(append(config["FamilyHistoryCreate"], negroni.HandlerFunc(FamilyHistoryCreateHandler))...))

	familyhistory := router.Path("/FamilyHistory/{id}").Subrouter()
	familyhistory.Methods("GET").Handler(negroni.New(append(config["FamilyHistoryShow"], negroni.HandlerFunc(FamilyHistoryShowHandler))...))
	familyhistory.Methods("PUT").Handler(negroni.New(append(config["FamilyHistoryUpdate"], negroni.HandlerFunc(FamilyHistoryUpdateHandler))...))
	familyhistory.Methods("DELETE").Handler(negroni.New(append(config["FamilyHistoryDelete"], negroni.HandlerFunc(FamilyHistoryDeleteHandler))...))

	medicationBase := router.Path("/Medication").Subrouter()
	medicationBase.Methods("GET").Handler(negroni.New(append(config["MedicationIndex"], negroni.HandlerFunc(MedicationIndexHandler))...))
	medicationBase.Methods("POST").Handler(negroni.New(append(config["MedicationCreate"], negroni.HandlerFunc(MedicationCreateHandler))...))

	medication := router.Path("/Medication/{id}").Subrouter()
	medication.Methods("GET").Handler(negroni.New(append(config["MedicationShow"], negroni.HandlerFunc(MedicationShowHandler))...))
	medication.Methods("PUT").Handler(negroni.New(append(config["MedicationUpdate"], negroni.HandlerFunc(MedicationUpdateHandler))...))
	medication.Methods("DELETE").Handler(negroni.New(append(config["MedicationDelete"], negroni.HandlerFunc(MedicationDeleteHandler))...))

	procedureBase := router.Path("/Procedure").Subrouter()
	procedureBase.Methods("GET").Handler(negroni.New(append(config["ProcedureIndex"], negroni.HandlerFunc(ProcedureIndexHandler))...))
	procedureBase.Methods("POST").Handler(negroni.New(append(config["ProcedureCreate"], negroni.HandlerFunc(ProcedureCreateHandler))...))

	procedure := router.Path("/Procedure/{id}").Subrouter()
	procedure.Methods("GET").Handler(negroni.New(append(config["ProcedureShow"], negroni.HandlerFunc(ProcedureShowHandler))...))
	procedure.Methods("PUT").Handler(negroni.New(append(config["ProcedureUpdate"], negroni.HandlerFunc(ProcedureUpdateHandler))...))
	procedure.Methods("DELETE").Handler(negroni.New(append(config["ProcedureDelete"], negroni.HandlerFunc(ProcedureDeleteHandler))...))

	listBase := router.Path("/List").Subrouter()
	listBase.Methods("GET").Handler(negroni.New(append(config["ListIndex"], negroni.HandlerFunc(ListIndexHandler))...))
	listBase.Methods("POST").Handler(negroni.New(append(config["ListCreate"], negroni.HandlerFunc(ListCreateHandler))...))

	list := router.Path("/List/{id}").Subrouter()
	list.Methods("GET").Handler(negroni.New(append(config["ListShow"], negroni.HandlerFunc(ListShowHandler))...))
	list.Methods("PUT").Handler(negroni.New(append(config["ListUpdate"], negroni.HandlerFunc(ListUpdateHandler))...))
	list.Methods("DELETE").Handler(negroni.New(append(config["ListDelete"], negroni.HandlerFunc(ListDeleteHandler))...))

	conceptmapBase := router.Path("/ConceptMap").Subrouter()
	conceptmapBase.Methods("GET").Handler(negroni.New(append(config["ConceptMapIndex"], negroni.HandlerFunc(ConceptMapIndexHandler))...))
	conceptmapBase.Methods("POST").Handler(negroni.New(append(config["ConceptMapCreate"], negroni.HandlerFunc(ConceptMapCreateHandler))...))

	conceptmap := router.Path("/ConceptMap/{id}").Subrouter()
	conceptmap.Methods("GET").Handler(negroni.New(append(config["ConceptMapShow"], negroni.HandlerFunc(ConceptMapShowHandler))...))
	conceptmap.Methods("PUT").Handler(negroni.New(append(config["ConceptMapUpdate"], negroni.HandlerFunc(ConceptMapUpdateHandler))...))
	conceptmap.Methods("DELETE").Handler(negroni.New(append(config["ConceptMapDelete"], negroni.HandlerFunc(ConceptMapDeleteHandler))...))

	subscriptionBase := router.Path("/Subscription").Subrouter()
	subscriptionBase.Methods("GET").Handler(negroni.New(append(config["SubscriptionIndex"], negroni.HandlerFunc(SubscriptionIndexHandler))...))
	subscriptionBase.Methods("POST").Handler(negroni.New(append(config["SubscriptionCreate"], negroni.HandlerFunc(SubscriptionCreateHandler))...))

	subscription := router.Path("/Subscription/{id}").Subrouter()
	subscription.Methods("GET").Handler(negroni.New(append(config["SubscriptionShow"], negroni.HandlerFunc(SubscriptionShowHandler))...))
	subscription.Methods("PUT").Handler(negroni.New(append(config["SubscriptionUpdate"], negroni.HandlerFunc(SubscriptionUpdateHandler))...))
	subscription.Methods("DELETE").Handler(negroni.New(append(config["SubscriptionDelete"], negroni.HandlerFunc(SubscriptionDeleteHandler))...))

	valuesetBase := router.Path("/ValueSet").Subrouter()
	valuesetBase.Methods("GET").Handler(negroni.New(append(config["ValueSetIndex"], negroni.HandlerFunc(ValueSetIndexHandler))...))
	valuesetBase.Methods("POST").Handler(negroni.New(append(config["ValueSetCreate"], negroni.HandlerFunc(ValueSetCreateHandler))...))

	valueset := router.Path("/ValueSet/{id}").Subrouter()
	valueset.Methods("GET").Handler(negroni.New(append(config["ValueSetShow"], negroni.HandlerFunc(ValueSetShowHandler))...))
	valueset.Methods("PUT").Handler(negroni.New(append(config["ValueSetUpdate"], negroni.HandlerFunc(ValueSetUpdateHandler))...))
	valueset.Methods("DELETE").Handler(negroni.New(append(config["ValueSetDelete"], negroni.HandlerFunc(ValueSetDeleteHandler))...))

	operationdefinitionBase := router.Path("/OperationDefinition").Subrouter()
	operationdefinitionBase.Methods("GET").Handler(negroni.New(append(config["OperationDefinitionIndex"], negroni.HandlerFunc(OperationDefinitionIndexHandler))...))
	operationdefinitionBase.Methods("POST").Handler(negroni.New(append(config["OperationDefinitionCreate"], negroni.HandlerFunc(OperationDefinitionCreateHandler))...))

	operationdefinition := router.Path("/OperationDefinition/{id}").Subrouter()
	operationdefinition.Methods("GET").Handler(negroni.New(append(config["OperationDefinitionShow"], negroni.HandlerFunc(OperationDefinitionShowHandler))...))
	operationdefinition.Methods("PUT").Handler(negroni.New(append(config["OperationDefinitionUpdate"], negroni.HandlerFunc(OperationDefinitionUpdateHandler))...))
	operationdefinition.Methods("DELETE").Handler(negroni.New(append(config["OperationDefinitionDelete"], negroni.HandlerFunc(OperationDefinitionDeleteHandler))...))

	documentreferenceBase := router.Path("/DocumentReference").Subrouter()
	documentreferenceBase.Methods("GET").Handler(negroni.New(append(config["DocumentReferenceIndex"], negroni.HandlerFunc(DocumentReferenceIndexHandler))...))
	documentreferenceBase.Methods("POST").Handler(negroni.New(append(config["DocumentReferenceCreate"], negroni.HandlerFunc(DocumentReferenceCreateHandler))...))

	documentreference := router.Path("/DocumentReference/{id}").Subrouter()
	documentreference.Methods("GET").Handler(negroni.New(append(config["DocumentReferenceShow"], negroni.HandlerFunc(DocumentReferenceShowHandler))...))
	documentreference.Methods("PUT").Handler(negroni.New(append(config["DocumentReferenceUpdate"], negroni.HandlerFunc(DocumentReferenceUpdateHandler))...))
	documentreference.Methods("DELETE").Handler(negroni.New(append(config["DocumentReferenceDelete"], negroni.HandlerFunc(DocumentReferenceDeleteHandler))...))

	orderBase := router.Path("/Order").Subrouter()
	orderBase.Methods("GET").Handler(negroni.New(append(config["OrderIndex"], negroni.HandlerFunc(OrderIndexHandler))...))
	orderBase.Methods("POST").Handler(negroni.New(append(config["OrderCreate"], negroni.HandlerFunc(OrderCreateHandler))...))

	order := router.Path("/Order/{id}").Subrouter()
	order.Methods("GET").Handler(negroni.New(append(config["OrderShow"], negroni.HandlerFunc(OrderShowHandler))...))
	order.Methods("PUT").Handler(negroni.New(append(config["OrderUpdate"], negroni.HandlerFunc(OrderUpdateHandler))...))
	order.Methods("DELETE").Handler(negroni.New(append(config["OrderDelete"], negroni.HandlerFunc(OrderDeleteHandler))...))

	availabilityBase := router.Path("/Availability").Subrouter()
	availabilityBase.Methods("GET").Handler(negroni.New(append(config["AvailabilityIndex"], negroni.HandlerFunc(AvailabilityIndexHandler))...))
	availabilityBase.Methods("POST").Handler(negroni.New(append(config["AvailabilityCreate"], negroni.HandlerFunc(AvailabilityCreateHandler))...))

	availability := router.Path("/Availability/{id}").Subrouter()
	availability.Methods("GET").Handler(negroni.New(append(config["AvailabilityShow"], negroni.HandlerFunc(AvailabilityShowHandler))...))
	availability.Methods("PUT").Handler(negroni.New(append(config["AvailabilityUpdate"], negroni.HandlerFunc(AvailabilityUpdateHandler))...))
	availability.Methods("DELETE").Handler(negroni.New(append(config["AvailabilityDelete"], negroni.HandlerFunc(AvailabilityDeleteHandler))...))

	immunizationBase := router.Path("/Immunization").Subrouter()
	immunizationBase.Methods("GET").Handler(negroni.New(append(config["ImmunizationIndex"], negroni.HandlerFunc(ImmunizationIndexHandler))...))
	immunizationBase.Methods("POST").Handler(negroni.New(append(config["ImmunizationCreate"], negroni.HandlerFunc(ImmunizationCreateHandler))...))

	immunization := router.Path("/Immunization/{id}").Subrouter()
	immunization.Methods("GET").Handler(negroni.New(append(config["ImmunizationShow"], negroni.HandlerFunc(ImmunizationShowHandler))...))
	immunization.Methods("PUT").Handler(negroni.New(append(config["ImmunizationUpdate"], negroni.HandlerFunc(ImmunizationUpdateHandler))...))
	immunization.Methods("DELETE").Handler(negroni.New(append(config["ImmunizationDelete"], negroni.HandlerFunc(ImmunizationDeleteHandler))...))

	securityeventBase := router.Path("/SecurityEvent").Subrouter()
	securityeventBase.Methods("GET").Handler(negroni.New(append(config["SecurityEventIndex"], negroni.HandlerFunc(SecurityEventIndexHandler))...))
	securityeventBase.Methods("POST").Handler(negroni.New(append(config["SecurityEventCreate"], negroni.HandlerFunc(SecurityEventCreateHandler))...))

	securityevent := router.Path("/SecurityEvent/{id}").Subrouter()
	securityevent.Methods("GET").Handler(negroni.New(append(config["SecurityEventShow"], negroni.HandlerFunc(SecurityEventShowHandler))...))
	securityevent.Methods("PUT").Handler(negroni.New(append(config["SecurityEventUpdate"], negroni.HandlerFunc(SecurityEventUpdateHandler))...))
	securityevent.Methods("DELETE").Handler(negroni.New(append(config["SecurityEventDelete"], negroni.HandlerFunc(SecurityEventDeleteHandler))...))

	deviceBase := router.Path("/Device").Subrouter()
	deviceBase.Methods("GET").Handler(negroni.New(append(config["DeviceIndex"], negroni.HandlerFunc(DeviceIndexHandler))...))
	deviceBase.Methods("POST").Handler(negroni.New(append(config["DeviceCreate"], negroni.HandlerFunc(DeviceCreateHandler))...))

	device := router.Path("/Device/{id}").Subrouter()
	device.Methods("GET").Handler(negroni.New(append(config["DeviceShow"], negroni.HandlerFunc(DeviceShowHandler))...))
	device.Methods("PUT").Handler(negroni.New(append(config["DeviceUpdate"], negroni.HandlerFunc(DeviceUpdateHandler))...))
	device.Methods("DELETE").Handler(negroni.New(append(config["DeviceDelete"], negroni.HandlerFunc(DeviceDeleteHandler))...))

	mediaBase := router.Path("/Media").Subrouter()
	mediaBase.Methods("GET").Handler(negroni.New(append(config["MediaIndex"], negroni.HandlerFunc(MediaIndexHandler))...))
	mediaBase.Methods("POST").Handler(negroni.New(append(config["MediaCreate"], negroni.HandlerFunc(MediaCreateHandler))...))

	media := router.Path("/Media/{id}").Subrouter()
	media.Methods("GET").Handler(negroni.New(append(config["MediaShow"], negroni.HandlerFunc(MediaShowHandler))...))
	media.Methods("PUT").Handler(negroni.New(append(config["MediaUpdate"], negroni.HandlerFunc(MediaUpdateHandler))...))
	media.Methods("DELETE").Handler(negroni.New(append(config["MediaDelete"], negroni.HandlerFunc(MediaDeleteHandler))...))

	conformanceBase := router.Path("/Conformance").Subrouter()
	conformanceBase.Methods("GET").Handler(negroni.New(append(config["ConformanceIndex"], negroni.HandlerFunc(ConformanceIndexHandler))...))
	conformanceBase.Methods("POST").Handler(negroni.New(append(config["ConformanceCreate"], negroni.HandlerFunc(ConformanceCreateHandler))...))

	conformance := router.Path("/Conformance/{id}").Subrouter()
	conformance.Methods("GET").Handler(negroni.New(append(config["ConformanceShow"], negroni.HandlerFunc(ConformanceShowHandler))...))
	conformance.Methods("PUT").Handler(negroni.New(append(config["ConformanceUpdate"], negroni.HandlerFunc(ConformanceUpdateHandler))...))
	conformance.Methods("DELETE").Handler(negroni.New(append(config["ConformanceDelete"], negroni.HandlerFunc(ConformanceDeleteHandler))...))

	relatedpersonBase := router.Path("/RelatedPerson").Subrouter()
	relatedpersonBase.Methods("GET").Handler(negroni.New(append(config["RelatedPersonIndex"], negroni.HandlerFunc(RelatedPersonIndexHandler))...))
	relatedpersonBase.Methods("POST").Handler(negroni.New(append(config["RelatedPersonCreate"], negroni.HandlerFunc(RelatedPersonCreateHandler))...))

	relatedperson := router.Path("/RelatedPerson/{id}").Subrouter()
	relatedperson.Methods("GET").Handler(negroni.New(append(config["RelatedPersonShow"], negroni.HandlerFunc(RelatedPersonShowHandler))...))
	relatedperson.Methods("PUT").Handler(negroni.New(append(config["RelatedPersonUpdate"], negroni.HandlerFunc(RelatedPersonUpdateHandler))...))
	relatedperson.Methods("DELETE").Handler(negroni.New(append(config["RelatedPersonDelete"], negroni.HandlerFunc(RelatedPersonDeleteHandler))...))

	namespaceBase := router.Path("/Namespace").Subrouter()
	namespaceBase.Methods("GET").Handler(negroni.New(append(config["NamespaceIndex"], negroni.HandlerFunc(NamespaceIndexHandler))...))
	namespaceBase.Methods("POST").Handler(negroni.New(append(config["NamespaceCreate"], negroni.HandlerFunc(NamespaceCreateHandler))...))

	namespace := router.Path("/Namespace/{id}").Subrouter()
	namespace.Methods("GET").Handler(negroni.New(append(config["NamespaceShow"], negroni.HandlerFunc(NamespaceShowHandler))...))
	namespace.Methods("PUT").Handler(negroni.New(append(config["NamespaceUpdate"], negroni.HandlerFunc(NamespaceUpdateHandler))...))
	namespace.Methods("DELETE").Handler(negroni.New(append(config["NamespaceDelete"], negroni.HandlerFunc(NamespaceDeleteHandler))...))

	practitionerBase := router.Path("/Practitioner").Subrouter()
	practitionerBase.Methods("GET").Handler(negroni.New(append(config["PractitionerIndex"], negroni.HandlerFunc(PractitionerIndexHandler))...))
	practitionerBase.Methods("POST").Handler(negroni.New(append(config["PractitionerCreate"], negroni.HandlerFunc(PractitionerCreateHandler))...))

	practitioner := router.Path("/Practitioner/{id}").Subrouter()
	practitioner.Methods("GET").Handler(negroni.New(append(config["PractitionerShow"], negroni.HandlerFunc(PractitionerShowHandler))...))
	practitioner.Methods("PUT").Handler(negroni.New(append(config["PractitionerUpdate"], negroni.HandlerFunc(PractitionerUpdateHandler))...))
	practitioner.Methods("DELETE").Handler(negroni.New(append(config["PractitionerDelete"], negroni.HandlerFunc(PractitionerDeleteHandler))...))

	appointmentresponseBase := router.Path("/AppointmentResponse").Subrouter()
	appointmentresponseBase.Methods("GET").Handler(negroni.New(append(config["AppointmentResponseIndex"], negroni.HandlerFunc(AppointmentResponseIndexHandler))...))
	appointmentresponseBase.Methods("POST").Handler(negroni.New(append(config["AppointmentResponseCreate"], negroni.HandlerFunc(AppointmentResponseCreateHandler))...))

	appointmentresponse := router.Path("/AppointmentResponse/{id}").Subrouter()
	appointmentresponse.Methods("GET").Handler(negroni.New(append(config["AppointmentResponseShow"], negroni.HandlerFunc(AppointmentResponseShowHandler))...))
	appointmentresponse.Methods("PUT").Handler(negroni.New(append(config["AppointmentResponseUpdate"], negroni.HandlerFunc(AppointmentResponseUpdateHandler))...))
	appointmentresponse.Methods("DELETE").Handler(negroni.New(append(config["AppointmentResponseDelete"], negroni.HandlerFunc(AppointmentResponseDeleteHandler))...))

	observationBase := router.Path("/Observation").Subrouter()
	observationBase.Methods("GET").Handler(negroni.New(append(config["ObservationIndex"], negroni.HandlerFunc(ObservationIndexHandler))...))
	observationBase.Methods("POST").Handler(negroni.New(append(config["ObservationCreate"], negroni.HandlerFunc(ObservationCreateHandler))...))

	observation := router.Path("/Observation/{id}").Subrouter()
	observation.Methods("GET").Handler(negroni.New(append(config["ObservationShow"], negroni.HandlerFunc(ObservationShowHandler))...))
	observation.Methods("PUT").Handler(negroni.New(append(config["ObservationUpdate"], negroni.HandlerFunc(ObservationUpdateHandler))...))
	observation.Methods("DELETE").Handler(negroni.New(append(config["ObservationDelete"], negroni.HandlerFunc(ObservationDeleteHandler))...))

	medicationadministrationBase := router.Path("/MedicationAdministration").Subrouter()
	medicationadministrationBase.Methods("GET").Handler(negroni.New(append(config["MedicationAdministrationIndex"], negroni.HandlerFunc(MedicationAdministrationIndexHandler))...))
	medicationadministrationBase.Methods("POST").Handler(negroni.New(append(config["MedicationAdministrationCreate"], negroni.HandlerFunc(MedicationAdministrationCreateHandler))...))

	medicationadministration := router.Path("/MedicationAdministration/{id}").Subrouter()
	medicationadministration.Methods("GET").Handler(negroni.New(append(config["MedicationAdministrationShow"], negroni.HandlerFunc(MedicationAdministrationShowHandler))...))
	medicationadministration.Methods("PUT").Handler(negroni.New(append(config["MedicationAdministrationUpdate"], negroni.HandlerFunc(MedicationAdministrationUpdateHandler))...))
	medicationadministration.Methods("DELETE").Handler(negroni.New(append(config["MedicationAdministrationDelete"], negroni.HandlerFunc(MedicationAdministrationDeleteHandler))...))

	deviceobservationreportBase := router.Path("/DeviceObservationReport").Subrouter()
	deviceobservationreportBase.Methods("GET").Handler(negroni.New(append(config["DeviceObservationReportIndex"], negroni.HandlerFunc(DeviceObservationReportIndexHandler))...))
	deviceobservationreportBase.Methods("POST").Handler(negroni.New(append(config["DeviceObservationReportCreate"], negroni.HandlerFunc(DeviceObservationReportCreateHandler))...))

	deviceobservationreport := router.Path("/DeviceObservationReport/{id}").Subrouter()
	deviceobservationreport.Methods("GET").Handler(negroni.New(append(config["DeviceObservationReportShow"], negroni.HandlerFunc(DeviceObservationReportShowHandler))...))
	deviceobservationreport.Methods("PUT").Handler(negroni.New(append(config["DeviceObservationReportUpdate"], negroni.HandlerFunc(DeviceObservationReportUpdateHandler))...))
	deviceobservationreport.Methods("DELETE").Handler(negroni.New(append(config["DeviceObservationReportDelete"], negroni.HandlerFunc(DeviceObservationReportDeleteHandler))...))

	slotBase := router.Path("/Slot").Subrouter()
	slotBase.Methods("GET").Handler(negroni.New(append(config["SlotIndex"], negroni.HandlerFunc(SlotIndexHandler))...))
	slotBase.Methods("POST").Handler(negroni.New(append(config["SlotCreate"], negroni.HandlerFunc(SlotCreateHandler))...))

	slot := router.Path("/Slot/{id}").Subrouter()
	slot.Methods("GET").Handler(negroni.New(append(config["SlotShow"], negroni.HandlerFunc(SlotShowHandler))...))
	slot.Methods("PUT").Handler(negroni.New(append(config["SlotUpdate"], negroni.HandlerFunc(SlotUpdateHandler))...))
	slot.Methods("DELETE").Handler(negroni.New(append(config["SlotDelete"], negroni.HandlerFunc(SlotDeleteHandler))...))

	contraindicationBase := router.Path("/Contraindication").Subrouter()
	contraindicationBase.Methods("GET").Handler(negroni.New(append(config["ContraindicationIndex"], negroni.HandlerFunc(ContraindicationIndexHandler))...))
	contraindicationBase.Methods("POST").Handler(negroni.New(append(config["ContraindicationCreate"], negroni.HandlerFunc(ContraindicationCreateHandler))...))

	contraindication := router.Path("/Contraindication/{id}").Subrouter()
	contraindication.Methods("GET").Handler(negroni.New(append(config["ContraindicationShow"], negroni.HandlerFunc(ContraindicationShowHandler))...))
	contraindication.Methods("PUT").Handler(negroni.New(append(config["ContraindicationUpdate"], negroni.HandlerFunc(ContraindicationUpdateHandler))...))
	contraindication.Methods("DELETE").Handler(negroni.New(append(config["ContraindicationDelete"], negroni.HandlerFunc(ContraindicationDeleteHandler))...))

	medicationstatementBase := router.Path("/MedicationStatement").Subrouter()
	medicationstatementBase.Methods("GET").Handler(negroni.New(append(config["MedicationStatementIndex"], negroni.HandlerFunc(MedicationStatementIndexHandler))...))
	medicationstatementBase.Methods("POST").Handler(negroni.New(append(config["MedicationStatementCreate"], negroni.HandlerFunc(MedicationStatementCreateHandler))...))

	medicationstatement := router.Path("/MedicationStatement/{id}").Subrouter()
	medicationstatement.Methods("GET").Handler(negroni.New(append(config["MedicationStatementShow"], negroni.HandlerFunc(MedicationStatementShowHandler))...))
	medicationstatement.Methods("PUT").Handler(negroni.New(append(config["MedicationStatementUpdate"], negroni.HandlerFunc(MedicationStatementUpdateHandler))...))
	medicationstatement.Methods("DELETE").Handler(negroni.New(append(config["MedicationStatementDelete"], negroni.HandlerFunc(MedicationStatementDeleteHandler))...))

	riskassessmentBase := router.Path("/RiskAssessment").Subrouter()
	riskassessmentBase.Methods("GET").Handler(negroni.New(append(config["RiskAssessmentIndex"], negroni.HandlerFunc(RiskAssessmentIndexHandler))...))
	riskassessmentBase.Methods("POST").Handler(negroni.New(append(config["RiskAssessmentCreate"], negroni.HandlerFunc(RiskAssessmentCreateHandler))...))

	riskassessment := router.Path("/RiskAssessment/{id}").Subrouter()
	riskassessment.Methods("GET").Handler(negroni.New(append(config["RiskAssessmentShow"], negroni.HandlerFunc(RiskAssessmentShowHandler))...))
	riskassessment.Methods("PUT").Handler(negroni.New(append(config["RiskAssessmentUpdate"], negroni.HandlerFunc(RiskAssessmentUpdateHandler))...))
	riskassessment.Methods("DELETE").Handler(negroni.New(append(config["RiskAssessmentDelete"], negroni.HandlerFunc(RiskAssessmentDeleteHandler))...))

	groupBase := router.Path("/Group").Subrouter()
	groupBase.Methods("GET").Handler(negroni.New(append(config["GroupIndex"], negroni.HandlerFunc(GroupIndexHandler))...))
	groupBase.Methods("POST").Handler(negroni.New(append(config["GroupCreate"], negroni.HandlerFunc(GroupCreateHandler))...))

	group := router.Path("/Group/{id}").Subrouter()
	group.Methods("GET").Handler(negroni.New(append(config["GroupShow"], negroni.HandlerFunc(GroupShowHandler))...))
	group.Methods("PUT").Handler(negroni.New(append(config["GroupUpdate"], negroni.HandlerFunc(GroupUpdateHandler))...))
	group.Methods("DELETE").Handler(negroni.New(append(config["GroupDelete"], negroni.HandlerFunc(GroupDeleteHandler))...))

	organizationBase := router.Path("/Organization").Subrouter()
	organizationBase.Methods("GET").Handler(negroni.New(append(config["OrganizationIndex"], negroni.HandlerFunc(OrganizationIndexHandler))...))
	organizationBase.Methods("POST").Handler(negroni.New(append(config["OrganizationCreate"], negroni.HandlerFunc(OrganizationCreateHandler))...))

	organization := router.Path("/Organization/{id}").Subrouter()
	organization.Methods("GET").Handler(negroni.New(append(config["OrganizationShow"], negroni.HandlerFunc(OrganizationShowHandler))...))
	organization.Methods("PUT").Handler(negroni.New(append(config["OrganizationUpdate"], negroni.HandlerFunc(OrganizationUpdateHandler))...))
	organization.Methods("DELETE").Handler(negroni.New(append(config["OrganizationDelete"], negroni.HandlerFunc(OrganizationDeleteHandler))...))

	medicationdispenseBase := router.Path("/MedicationDispense").Subrouter()
	medicationdispenseBase.Methods("GET").Handler(negroni.New(append(config["MedicationDispenseIndex"], negroni.HandlerFunc(MedicationDispenseIndexHandler))...))
	medicationdispenseBase.Methods("POST").Handler(negroni.New(append(config["MedicationDispenseCreate"], negroni.HandlerFunc(MedicationDispenseCreateHandler))...))

	medicationdispense := router.Path("/MedicationDispense/{id}").Subrouter()
	medicationdispense.Methods("GET").Handler(negroni.New(append(config["MedicationDispenseShow"], negroni.HandlerFunc(MedicationDispenseShowHandler))...))
	medicationdispense.Methods("PUT").Handler(negroni.New(append(config["MedicationDispenseUpdate"], negroni.HandlerFunc(MedicationDispenseUpdateHandler))...))
	medicationdispense.Methods("DELETE").Handler(negroni.New(append(config["MedicationDispenseDelete"], negroni.HandlerFunc(MedicationDispenseDeleteHandler))...))

	supplyBase := router.Path("/Supply").Subrouter()
	supplyBase.Methods("GET").Handler(negroni.New(append(config["SupplyIndex"], negroni.HandlerFunc(SupplyIndexHandler))...))
	supplyBase.Methods("POST").Handler(negroni.New(append(config["SupplyCreate"], negroni.HandlerFunc(SupplyCreateHandler))...))

	supply := router.Path("/Supply/{id}").Subrouter()
	supply.Methods("GET").Handler(negroni.New(append(config["SupplyShow"], negroni.HandlerFunc(SupplyShowHandler))...))
	supply.Methods("PUT").Handler(negroni.New(append(config["SupplyUpdate"], negroni.HandlerFunc(SupplyUpdateHandler))...))
	supply.Methods("DELETE").Handler(negroni.New(append(config["SupplyDelete"], negroni.HandlerFunc(SupplyDeleteHandler))...))

	diagnosticreportBase := router.Path("/DiagnosticReport").Subrouter()
	diagnosticreportBase.Methods("GET").Handler(negroni.New(append(config["DiagnosticReportIndex"], negroni.HandlerFunc(DiagnosticReportIndexHandler))...))
	diagnosticreportBase.Methods("POST").Handler(negroni.New(append(config["DiagnosticReportCreate"], negroni.HandlerFunc(DiagnosticReportCreateHandler))...))

	diagnosticreport := router.Path("/DiagnosticReport/{id}").Subrouter()
	diagnosticreport.Methods("GET").Handler(negroni.New(append(config["DiagnosticReportShow"], negroni.HandlerFunc(DiagnosticReportShowHandler))...))
	diagnosticreport.Methods("PUT").Handler(negroni.New(append(config["DiagnosticReportUpdate"], negroni.HandlerFunc(DiagnosticReportUpdateHandler))...))
	diagnosticreport.Methods("DELETE").Handler(negroni.New(append(config["DiagnosticReportDelete"], negroni.HandlerFunc(DiagnosticReportDeleteHandler))...))

	imagingstudyBase := router.Path("/ImagingStudy").Subrouter()
	imagingstudyBase.Methods("GET").Handler(negroni.New(append(config["ImagingStudyIndex"], negroni.HandlerFunc(ImagingStudyIndexHandler))...))
	imagingstudyBase.Methods("POST").Handler(negroni.New(append(config["ImagingStudyCreate"], negroni.HandlerFunc(ImagingStudyCreateHandler))...))

	imagingstudy := router.Path("/ImagingStudy/{id}").Subrouter()
	imagingstudy.Methods("GET").Handler(negroni.New(append(config["ImagingStudyShow"], negroni.HandlerFunc(ImagingStudyShowHandler))...))
	imagingstudy.Methods("PUT").Handler(negroni.New(append(config["ImagingStudyUpdate"], negroni.HandlerFunc(ImagingStudyUpdateHandler))...))
	imagingstudy.Methods("DELETE").Handler(negroni.New(append(config["ImagingStudyDelete"], negroni.HandlerFunc(ImagingStudyDeleteHandler))...))

	profileBase := router.Path("/Profile").Subrouter()
	profileBase.Methods("GET").Handler(negroni.New(append(config["ProfileIndex"], negroni.HandlerFunc(ProfileIndexHandler))...))
	profileBase.Methods("POST").Handler(negroni.New(append(config["ProfileCreate"], negroni.HandlerFunc(ProfileCreateHandler))...))

	profile := router.Path("/Profile/{id}").Subrouter()
	profile.Methods("GET").Handler(negroni.New(append(config["ProfileShow"], negroni.HandlerFunc(ProfileShowHandler))...))
	profile.Methods("PUT").Handler(negroni.New(append(config["ProfileUpdate"], negroni.HandlerFunc(ProfileUpdateHandler))...))
	profile.Methods("DELETE").Handler(negroni.New(append(config["ProfileDelete"], negroni.HandlerFunc(ProfileDeleteHandler))...))

	dataelementBase := router.Path("/DataElement").Subrouter()
	dataelementBase.Methods("GET").Handler(negroni.New(append(config["DataElementIndex"], negroni.HandlerFunc(DataElementIndexHandler))...))
	dataelementBase.Methods("POST").Handler(negroni.New(append(config["DataElementCreate"], negroni.HandlerFunc(DataElementCreateHandler))...))

	dataelement := router.Path("/DataElement/{id}").Subrouter()
	dataelement.Methods("GET").Handler(negroni.New(append(config["DataElementShow"], negroni.HandlerFunc(DataElementShowHandler))...))
	dataelement.Methods("PUT").Handler(negroni.New(append(config["DataElementUpdate"], negroni.HandlerFunc(DataElementUpdateHandler))...))
	dataelement.Methods("DELETE").Handler(negroni.New(append(config["DataElementDelete"], negroni.HandlerFunc(DataElementDeleteHandler))...))

	questionnaireanswersBase := router.Path("/QuestionnaireAnswers").Subrouter()
	questionnaireanswersBase.Methods("GET").Handler(negroni.New(append(config["QuestionnaireAnswersIndex"], negroni.HandlerFunc(QuestionnaireAnswersIndexHandler))...))
	questionnaireanswersBase.Methods("POST").Handler(negroni.New(append(config["QuestionnaireAnswersCreate"], negroni.HandlerFunc(QuestionnaireAnswersCreateHandler))...))

	questionnaireanswers := router.Path("/QuestionnaireAnswers/{id}").Subrouter()
	questionnaireanswers.Methods("GET").Handler(negroni.New(append(config["QuestionnaireAnswersShow"], negroni.HandlerFunc(QuestionnaireAnswersShowHandler))...))
	questionnaireanswers.Methods("PUT").Handler(negroni.New(append(config["QuestionnaireAnswersUpdate"], negroni.HandlerFunc(QuestionnaireAnswersUpdateHandler))...))
	questionnaireanswers.Methods("DELETE").Handler(negroni.New(append(config["QuestionnaireAnswersDelete"], negroni.HandlerFunc(QuestionnaireAnswersDeleteHandler))...))

	adversereactionBase := router.Path("/AdverseReaction").Subrouter()
	adversereactionBase.Methods("GET").Handler(negroni.New(append(config["AdverseReactionIndex"], negroni.HandlerFunc(AdverseReactionIndexHandler))...))
	adversereactionBase.Methods("POST").Handler(negroni.New(append(config["AdverseReactionCreate"], negroni.HandlerFunc(AdverseReactionCreateHandler))...))

	adversereaction := router.Path("/AdverseReaction/{id}").Subrouter()
	adversereaction.Methods("GET").Handler(negroni.New(append(config["AdverseReactionShow"], negroni.HandlerFunc(AdverseReactionShowHandler))...))
	adversereaction.Methods("PUT").Handler(negroni.New(append(config["AdverseReactionUpdate"], negroni.HandlerFunc(AdverseReactionUpdateHandler))...))
	adversereaction.Methods("DELETE").Handler(negroni.New(append(config["AdverseReactionDelete"], negroni.HandlerFunc(AdverseReactionDeleteHandler))...))

	nutritionorderBase := router.Path("/NutritionOrder").Subrouter()
	nutritionorderBase.Methods("GET").Handler(negroni.New(append(config["NutritionOrderIndex"], negroni.HandlerFunc(NutritionOrderIndexHandler))...))
	nutritionorderBase.Methods("POST").Handler(negroni.New(append(config["NutritionOrderCreate"], negroni.HandlerFunc(NutritionOrderCreateHandler))...))

	nutritionorder := router.Path("/NutritionOrder/{id}").Subrouter()
	nutritionorder.Methods("GET").Handler(negroni.New(append(config["NutritionOrderShow"], negroni.HandlerFunc(NutritionOrderShowHandler))...))
	nutritionorder.Methods("PUT").Handler(negroni.New(append(config["NutritionOrderUpdate"], negroni.HandlerFunc(NutritionOrderUpdateHandler))...))
	nutritionorder.Methods("DELETE").Handler(negroni.New(append(config["NutritionOrderDelete"], negroni.HandlerFunc(NutritionOrderDeleteHandler))...))

	encounterBase := router.Path("/Encounter").Subrouter()
	encounterBase.Methods("GET").Handler(negroni.New(append(config["EncounterIndex"], negroni.HandlerFunc(EncounterIndexHandler))...))
	encounterBase.Methods("POST").Handler(negroni.New(append(config["EncounterCreate"], negroni.HandlerFunc(EncounterCreateHandler))...))

	encounter := router.Path("/Encounter/{id}").Subrouter()
	encounter.Methods("GET").Handler(negroni.New(append(config["EncounterShow"], negroni.HandlerFunc(EncounterShowHandler))...))
	encounter.Methods("PUT").Handler(negroni.New(append(config["EncounterUpdate"], negroni.HandlerFunc(EncounterUpdateHandler))...))
	encounter.Methods("DELETE").Handler(negroni.New(append(config["EncounterDelete"], negroni.HandlerFunc(EncounterDeleteHandler))...))

	substanceBase := router.Path("/Substance").Subrouter()
	substanceBase.Methods("GET").Handler(negroni.New(append(config["SubstanceIndex"], negroni.HandlerFunc(SubstanceIndexHandler))...))
	substanceBase.Methods("POST").Handler(negroni.New(append(config["SubstanceCreate"], negroni.HandlerFunc(SubstanceCreateHandler))...))

	substance := router.Path("/Substance/{id}").Subrouter()
	substance.Methods("GET").Handler(negroni.New(append(config["SubstanceShow"], negroni.HandlerFunc(SubstanceShowHandler))...))
	substance.Methods("PUT").Handler(negroni.New(append(config["SubstanceUpdate"], negroni.HandlerFunc(SubstanceUpdateHandler))...))
	substance.Methods("DELETE").Handler(negroni.New(append(config["SubstanceDelete"], negroni.HandlerFunc(SubstanceDeleteHandler))...))

	conditionBase := router.Path("/Condition").Subrouter()
	conditionBase.Methods("GET").Handler(negroni.New(append(config["ConditionIndex"], negroni.HandlerFunc(ConditionIndexHandler))...))
	conditionBase.Methods("POST").Handler(negroni.New(append(config["ConditionCreate"], negroni.HandlerFunc(ConditionCreateHandler))...))

	condition := router.Path("/Condition/{id}").Subrouter()
	condition.Methods("GET").Handler(negroni.New(append(config["ConditionShow"], negroni.HandlerFunc(ConditionShowHandler))...))
	condition.Methods("PUT").Handler(negroni.New(append(config["ConditionUpdate"], negroni.HandlerFunc(ConditionUpdateHandler))...))
	condition.Methods("DELETE").Handler(negroni.New(append(config["ConditionDelete"], negroni.HandlerFunc(ConditionDeleteHandler))...))

	compositionBase := router.Path("/Composition").Subrouter()
	compositionBase.Methods("GET").Handler(negroni.New(append(config["CompositionIndex"], negroni.HandlerFunc(CompositionIndexHandler))...))
	compositionBase.Methods("POST").Handler(negroni.New(append(config["CompositionCreate"], negroni.HandlerFunc(CompositionCreateHandler))...))

	composition := router.Path("/Composition/{id}").Subrouter()
	composition.Methods("GET").Handler(negroni.New(append(config["CompositionShow"], negroni.HandlerFunc(CompositionShowHandler))...))
	composition.Methods("PUT").Handler(negroni.New(append(config["CompositionUpdate"], negroni.HandlerFunc(CompositionUpdateHandler))...))
	composition.Methods("DELETE").Handler(negroni.New(append(config["CompositionDelete"], negroni.HandlerFunc(CompositionDeleteHandler))...))

	diagnosticorderBase := router.Path("/DiagnosticOrder").Subrouter()
	diagnosticorderBase.Methods("GET").Handler(negroni.New(append(config["DiagnosticOrderIndex"], negroni.HandlerFunc(DiagnosticOrderIndexHandler))...))
	diagnosticorderBase.Methods("POST").Handler(negroni.New(append(config["DiagnosticOrderCreate"], negroni.HandlerFunc(DiagnosticOrderCreateHandler))...))

	diagnosticorder := router.Path("/DiagnosticOrder/{id}").Subrouter()
	diagnosticorder.Methods("GET").Handler(negroni.New(append(config["DiagnosticOrderShow"], negroni.HandlerFunc(DiagnosticOrderShowHandler))...))
	diagnosticorder.Methods("PUT").Handler(negroni.New(append(config["DiagnosticOrderUpdate"], negroni.HandlerFunc(DiagnosticOrderUpdateHandler))...))
	diagnosticorder.Methods("DELETE").Handler(negroni.New(append(config["DiagnosticOrderDelete"], negroni.HandlerFunc(DiagnosticOrderDeleteHandler))...))

	patientBase := router.Path("/Patient").Subrouter()
	patientBase.Methods("GET").Handler(negroni.New(append(config["PatientIndex"], negroni.HandlerFunc(PatientIndexHandler))...))
	patientBase.Methods("POST").Handler(negroni.New(append(config["PatientCreate"], negroni.HandlerFunc(PatientCreateHandler))...))

	patient := router.Path("/Patient/{id}").Subrouter()
	patient.Methods("GET").Handler(negroni.New(append(config["PatientShow"], negroni.HandlerFunc(PatientShowHandler))...))
	patient.Methods("PUT").Handler(negroni.New(append(config["PatientUpdate"], negroni.HandlerFunc(PatientUpdateHandler))...))
	patient.Methods("DELETE").Handler(negroni.New(append(config["PatientDelete"], negroni.HandlerFunc(PatientDeleteHandler))...))

	orderresponseBase := router.Path("/OrderResponse").Subrouter()
	orderresponseBase.Methods("GET").Handler(negroni.New(append(config["OrderResponseIndex"], negroni.HandlerFunc(OrderResponseIndexHandler))...))
	orderresponseBase.Methods("POST").Handler(negroni.New(append(config["OrderResponseCreate"], negroni.HandlerFunc(OrderResponseCreateHandler))...))

	orderresponse := router.Path("/OrderResponse/{id}").Subrouter()
	orderresponse.Methods("GET").Handler(negroni.New(append(config["OrderResponseShow"], negroni.HandlerFunc(OrderResponseShowHandler))...))
	orderresponse.Methods("PUT").Handler(negroni.New(append(config["OrderResponseUpdate"], negroni.HandlerFunc(OrderResponseUpdateHandler))...))
	orderresponse.Methods("DELETE").Handler(negroni.New(append(config["OrderResponseDelete"], negroni.HandlerFunc(OrderResponseDeleteHandler))...))

	alertBase := router.Path("/Alert").Subrouter()
	alertBase.Methods("GET").Handler(negroni.New(append(config["AlertIndex"], negroni.HandlerFunc(AlertIndexHandler))...))
	alertBase.Methods("POST").Handler(negroni.New(append(config["AlertCreate"], negroni.HandlerFunc(AlertCreateHandler))...))

	alert := router.Path("/Alert/{id}").Subrouter()
	alert.Methods("GET").Handler(negroni.New(append(config["AlertShow"], negroni.HandlerFunc(AlertShowHandler))...))
	alert.Methods("PUT").Handler(negroni.New(append(config["AlertUpdate"], negroni.HandlerFunc(AlertUpdateHandler))...))
	alert.Methods("DELETE").Handler(negroni.New(append(config["AlertDelete"], negroni.HandlerFunc(AlertDeleteHandler))...))

	messageheaderBase := router.Path("/MessageHeader").Subrouter()
	messageheaderBase.Methods("GET").Handler(negroni.New(append(config["MessageHeaderIndex"], negroni.HandlerFunc(MessageHeaderIndexHandler))...))
	messageheaderBase.Methods("POST").Handler(negroni.New(append(config["MessageHeaderCreate"], negroni.HandlerFunc(MessageHeaderCreateHandler))...))

	messageheader := router.Path("/MessageHeader/{id}").Subrouter()
	messageheader.Methods("GET").Handler(negroni.New(append(config["MessageHeaderShow"], negroni.HandlerFunc(MessageHeaderShowHandler))...))
	messageheader.Methods("PUT").Handler(negroni.New(append(config["MessageHeaderUpdate"], negroni.HandlerFunc(MessageHeaderUpdateHandler))...))
	messageheader.Methods("DELETE").Handler(negroni.New(append(config["MessageHeaderDelete"], negroni.HandlerFunc(MessageHeaderDeleteHandler))...))

	immunizationrecommendationBase := router.Path("/ImmunizationRecommendation").Subrouter()
	immunizationrecommendationBase.Methods("GET").Handler(negroni.New(append(config["ImmunizationRecommendationIndex"], negroni.HandlerFunc(ImmunizationRecommendationIndexHandler))...))
	immunizationrecommendationBase.Methods("POST").Handler(negroni.New(append(config["ImmunizationRecommendationCreate"], negroni.HandlerFunc(ImmunizationRecommendationCreateHandler))...))

	immunizationrecommendation := router.Path("/ImmunizationRecommendation/{id}").Subrouter()
	immunizationrecommendation.Methods("GET").Handler(negroni.New(append(config["ImmunizationRecommendationShow"], negroni.HandlerFunc(ImmunizationRecommendationShowHandler))...))
	immunizationrecommendation.Methods("PUT").Handler(negroni.New(append(config["ImmunizationRecommendationUpdate"], negroni.HandlerFunc(ImmunizationRecommendationUpdateHandler))...))
	immunizationrecommendation.Methods("DELETE").Handler(negroni.New(append(config["ImmunizationRecommendationDelete"], negroni.HandlerFunc(ImmunizationRecommendationDeleteHandler))...))

	locationBase := router.Path("/Location").Subrouter()
	locationBase.Methods("GET").Handler(negroni.New(append(config["LocationIndex"], negroni.HandlerFunc(LocationIndexHandler))...))
	locationBase.Methods("POST").Handler(negroni.New(append(config["LocationCreate"], negroni.HandlerFunc(LocationCreateHandler))...))

	location := router.Path("/Location/{id}").Subrouter()
	location.Methods("GET").Handler(negroni.New(append(config["LocationShow"], negroni.HandlerFunc(LocationShowHandler))...))
	location.Methods("PUT").Handler(negroni.New(append(config["LocationUpdate"], negroni.HandlerFunc(LocationUpdateHandler))...))
	location.Methods("DELETE").Handler(negroni.New(append(config["LocationDelete"], negroni.HandlerFunc(LocationDeleteHandler))...))

	otherBase := router.Path("/Other").Subrouter()
	otherBase.Methods("GET").Handler(negroni.New(append(config["OtherIndex"], negroni.HandlerFunc(OtherIndexHandler))...))
	otherBase.Methods("POST").Handler(negroni.New(append(config["OtherCreate"], negroni.HandlerFunc(OtherCreateHandler))...))

	other := router.Path("/Other/{id}").Subrouter()
	other.Methods("GET").Handler(negroni.New(append(config["OtherShow"], negroni.HandlerFunc(OtherShowHandler))...))
	other.Methods("PUT").Handler(negroni.New(append(config["OtherUpdate"], negroni.HandlerFunc(OtherUpdateHandler))...))
	other.Methods("DELETE").Handler(negroni.New(append(config["OtherDelete"], negroni.HandlerFunc(OtherDeleteHandler))...))

}
