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

	explanationofbenefitBase := router.Path("/ExplanationOfBenefit").Subrouter()
	explanationofbenefitBase.Methods("GET").Handler(negroni.New(append(config["ExplanationOfBenefitIndex"], negroni.HandlerFunc(ExplanationOfBenefitIndexHandler))...))
	explanationofbenefitBase.Methods("POST").Handler(negroni.New(append(config["ExplanationOfBenefitCreate"], negroni.HandlerFunc(ExplanationOfBenefitCreateHandler))...))

	explanationofbenefit := router.Path("/ExplanationOfBenefit/{id}").Subrouter()
	explanationofbenefit.Methods("GET").Handler(negroni.New(append(config["ExplanationOfBenefitShow"], negroni.HandlerFunc(ExplanationOfBenefitShowHandler))...))
	explanationofbenefit.Methods("PUT").Handler(negroni.New(append(config["ExplanationOfBenefitUpdate"], negroni.HandlerFunc(ExplanationOfBenefitUpdateHandler))...))
	explanationofbenefit.Methods("DELETE").Handler(negroni.New(append(config["ExplanationOfBenefitDelete"], negroni.HandlerFunc(ExplanationOfBenefitDeleteHandler))...))

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

	goalBase := router.Path("/Goal").Subrouter()
	goalBase.Methods("GET").Handler(negroni.New(append(config["GoalIndex"], negroni.HandlerFunc(GoalIndexHandler))...))
	goalBase.Methods("POST").Handler(negroni.New(append(config["GoalCreate"], negroni.HandlerFunc(GoalCreateHandler))...))

	goal := router.Path("/Goal/{id}").Subrouter()
	goal.Methods("GET").Handler(negroni.New(append(config["GoalShow"], negroni.HandlerFunc(GoalShowHandler))...))
	goal.Methods("PUT").Handler(negroni.New(append(config["GoalUpdate"], negroni.HandlerFunc(GoalUpdateHandler))...))
	goal.Methods("DELETE").Handler(negroni.New(append(config["GoalDelete"], negroni.HandlerFunc(GoalDeleteHandler))...))

	structuredefinitionBase := router.Path("/StructureDefinition").Subrouter()
	structuredefinitionBase.Methods("GET").Handler(negroni.New(append(config["StructureDefinitionIndex"], negroni.HandlerFunc(StructureDefinitionIndexHandler))...))
	structuredefinitionBase.Methods("POST").Handler(negroni.New(append(config["StructureDefinitionCreate"], negroni.HandlerFunc(StructureDefinitionCreateHandler))...))

	structuredefinition := router.Path("/StructureDefinition/{id}").Subrouter()
	structuredefinition.Methods("GET").Handler(negroni.New(append(config["StructureDefinitionShow"], negroni.HandlerFunc(StructureDefinitionShowHandler))...))
	structuredefinition.Methods("PUT").Handler(negroni.New(append(config["StructureDefinitionUpdate"], negroni.HandlerFunc(StructureDefinitionUpdateHandler))...))
	structuredefinition.Methods("DELETE").Handler(negroni.New(append(config["StructureDefinitionDelete"], negroni.HandlerFunc(StructureDefinitionDeleteHandler))...))

	enrollmentrequestBase := router.Path("/EnrollmentRequest").Subrouter()
	enrollmentrequestBase.Methods("GET").Handler(negroni.New(append(config["EnrollmentRequestIndex"], negroni.HandlerFunc(EnrollmentRequestIndexHandler))...))
	enrollmentrequestBase.Methods("POST").Handler(negroni.New(append(config["EnrollmentRequestCreate"], negroni.HandlerFunc(EnrollmentRequestCreateHandler))...))

	enrollmentrequest := router.Path("/EnrollmentRequest/{id}").Subrouter()
	enrollmentrequest.Methods("GET").Handler(negroni.New(append(config["EnrollmentRequestShow"], negroni.HandlerFunc(EnrollmentRequestShowHandler))...))
	enrollmentrequest.Methods("PUT").Handler(negroni.New(append(config["EnrollmentRequestUpdate"], negroni.HandlerFunc(EnrollmentRequestUpdateHandler))...))
	enrollmentrequest.Methods("DELETE").Handler(negroni.New(append(config["EnrollmentRequestDelete"], negroni.HandlerFunc(EnrollmentRequestDeleteHandler))...))

	episodeofcareBase := router.Path("/EpisodeOfCare").Subrouter()
	episodeofcareBase.Methods("GET").Handler(negroni.New(append(config["EpisodeOfCareIndex"], negroni.HandlerFunc(EpisodeOfCareIndexHandler))...))
	episodeofcareBase.Methods("POST").Handler(negroni.New(append(config["EpisodeOfCareCreate"], negroni.HandlerFunc(EpisodeOfCareCreateHandler))...))

	episodeofcare := router.Path("/EpisodeOfCare/{id}").Subrouter()
	episodeofcare.Methods("GET").Handler(negroni.New(append(config["EpisodeOfCareShow"], negroni.HandlerFunc(EpisodeOfCareShowHandler))...))
	episodeofcare.Methods("PUT").Handler(negroni.New(append(config["EpisodeOfCareUpdate"], negroni.HandlerFunc(EpisodeOfCareUpdateHandler))...))
	episodeofcare.Methods("DELETE").Handler(negroni.New(append(config["EpisodeOfCareDelete"], negroni.HandlerFunc(EpisodeOfCareDeleteHandler))...))

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

	immunizationBase := router.Path("/Immunization").Subrouter()
	immunizationBase.Methods("GET").Handler(negroni.New(append(config["ImmunizationIndex"], negroni.HandlerFunc(ImmunizationIndexHandler))...))
	immunizationBase.Methods("POST").Handler(negroni.New(append(config["ImmunizationCreate"], negroni.HandlerFunc(ImmunizationCreateHandler))...))

	immunization := router.Path("/Immunization/{id}").Subrouter()
	immunization.Methods("GET").Handler(negroni.New(append(config["ImmunizationShow"], negroni.HandlerFunc(ImmunizationShowHandler))...))
	immunization.Methods("PUT").Handler(negroni.New(append(config["ImmunizationUpdate"], negroni.HandlerFunc(ImmunizationUpdateHandler))...))
	immunization.Methods("DELETE").Handler(negroni.New(append(config["ImmunizationDelete"], negroni.HandlerFunc(ImmunizationDeleteHandler))...))

	deviceBase := router.Path("/Device").Subrouter()
	deviceBase.Methods("GET").Handler(negroni.New(append(config["DeviceIndex"], negroni.HandlerFunc(DeviceIndexHandler))...))
	deviceBase.Methods("POST").Handler(negroni.New(append(config["DeviceCreate"], negroni.HandlerFunc(DeviceCreateHandler))...))

	device := router.Path("/Device/{id}").Subrouter()
	device.Methods("GET").Handler(negroni.New(append(config["DeviceShow"], negroni.HandlerFunc(DeviceShowHandler))...))
	device.Methods("PUT").Handler(negroni.New(append(config["DeviceUpdate"], negroni.HandlerFunc(DeviceUpdateHandler))...))
	device.Methods("DELETE").Handler(negroni.New(append(config["DeviceDelete"], negroni.HandlerFunc(DeviceDeleteHandler))...))

	visionprescriptionBase := router.Path("/VisionPrescription").Subrouter()
	visionprescriptionBase.Methods("GET").Handler(negroni.New(append(config["VisionPrescriptionIndex"], negroni.HandlerFunc(VisionPrescriptionIndexHandler))...))
	visionprescriptionBase.Methods("POST").Handler(negroni.New(append(config["VisionPrescriptionCreate"], negroni.HandlerFunc(VisionPrescriptionCreateHandler))...))

	visionprescription := router.Path("/VisionPrescription/{id}").Subrouter()
	visionprescription.Methods("GET").Handler(negroni.New(append(config["VisionPrescriptionShow"], negroni.HandlerFunc(VisionPrescriptionShowHandler))...))
	visionprescription.Methods("PUT").Handler(negroni.New(append(config["VisionPrescriptionUpdate"], negroni.HandlerFunc(VisionPrescriptionUpdateHandler))...))
	visionprescription.Methods("DELETE").Handler(negroni.New(append(config["VisionPrescriptionDelete"], negroni.HandlerFunc(VisionPrescriptionDeleteHandler))...))

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

	procedurerequestBase := router.Path("/ProcedureRequest").Subrouter()
	procedurerequestBase.Methods("GET").Handler(negroni.New(append(config["ProcedureRequestIndex"], negroni.HandlerFunc(ProcedureRequestIndexHandler))...))
	procedurerequestBase.Methods("POST").Handler(negroni.New(append(config["ProcedureRequestCreate"], negroni.HandlerFunc(ProcedureRequestCreateHandler))...))

	procedurerequest := router.Path("/ProcedureRequest/{id}").Subrouter()
	procedurerequest.Methods("GET").Handler(negroni.New(append(config["ProcedureRequestShow"], negroni.HandlerFunc(ProcedureRequestShowHandler))...))
	procedurerequest.Methods("PUT").Handler(negroni.New(append(config["ProcedureRequestUpdate"], negroni.HandlerFunc(ProcedureRequestUpdateHandler))...))
	procedurerequest.Methods("DELETE").Handler(negroni.New(append(config["ProcedureRequestDelete"], negroni.HandlerFunc(ProcedureRequestDeleteHandler))...))

	eligibilityresponseBase := router.Path("/EligibilityResponse").Subrouter()
	eligibilityresponseBase.Methods("GET").Handler(negroni.New(append(config["EligibilityResponseIndex"], negroni.HandlerFunc(EligibilityResponseIndexHandler))...))
	eligibilityresponseBase.Methods("POST").Handler(negroni.New(append(config["EligibilityResponseCreate"], negroni.HandlerFunc(EligibilityResponseCreateHandler))...))

	eligibilityresponse := router.Path("/EligibilityResponse/{id}").Subrouter()
	eligibilityresponse.Methods("GET").Handler(negroni.New(append(config["EligibilityResponseShow"], negroni.HandlerFunc(EligibilityResponseShowHandler))...))
	eligibilityresponse.Methods("PUT").Handler(negroni.New(append(config["EligibilityResponseUpdate"], negroni.HandlerFunc(EligibilityResponseUpdateHandler))...))
	eligibilityresponse.Methods("DELETE").Handler(negroni.New(append(config["EligibilityResponseDelete"], negroni.HandlerFunc(EligibilityResponseDeleteHandler))...))

	deviceuserequestBase := router.Path("/DeviceUseRequest").Subrouter()
	deviceuserequestBase.Methods("GET").Handler(negroni.New(append(config["DeviceUseRequestIndex"], negroni.HandlerFunc(DeviceUseRequestIndexHandler))...))
	deviceuserequestBase.Methods("POST").Handler(negroni.New(append(config["DeviceUseRequestCreate"], negroni.HandlerFunc(DeviceUseRequestCreateHandler))...))

	deviceuserequest := router.Path("/DeviceUseRequest/{id}").Subrouter()
	deviceuserequest.Methods("GET").Handler(negroni.New(append(config["DeviceUseRequestShow"], negroni.HandlerFunc(DeviceUseRequestShowHandler))...))
	deviceuserequest.Methods("PUT").Handler(negroni.New(append(config["DeviceUseRequestUpdate"], negroni.HandlerFunc(DeviceUseRequestUpdateHandler))...))
	deviceuserequest.Methods("DELETE").Handler(negroni.New(append(config["DeviceUseRequestDelete"], negroni.HandlerFunc(DeviceUseRequestDeleteHandler))...))

	devicemetricBase := router.Path("/DeviceMetric").Subrouter()
	devicemetricBase.Methods("GET").Handler(negroni.New(append(config["DeviceMetricIndex"], negroni.HandlerFunc(DeviceMetricIndexHandler))...))
	devicemetricBase.Methods("POST").Handler(negroni.New(append(config["DeviceMetricCreate"], negroni.HandlerFunc(DeviceMetricCreateHandler))...))

	devicemetric := router.Path("/DeviceMetric/{id}").Subrouter()
	devicemetric.Methods("GET").Handler(negroni.New(append(config["DeviceMetricShow"], negroni.HandlerFunc(DeviceMetricShowHandler))...))
	devicemetric.Methods("PUT").Handler(negroni.New(append(config["DeviceMetricUpdate"], negroni.HandlerFunc(DeviceMetricUpdateHandler))...))
	devicemetric.Methods("DELETE").Handler(negroni.New(append(config["DeviceMetricDelete"], negroni.HandlerFunc(DeviceMetricDeleteHandler))...))

	flagBase := router.Path("/Flag").Subrouter()
	flagBase.Methods("GET").Handler(negroni.New(append(config["FlagIndex"], negroni.HandlerFunc(FlagIndexHandler))...))
	flagBase.Methods("POST").Handler(negroni.New(append(config["FlagCreate"], negroni.HandlerFunc(FlagCreateHandler))...))

	flag := router.Path("/Flag/{id}").Subrouter()
	flag.Methods("GET").Handler(negroni.New(append(config["FlagShow"], negroni.HandlerFunc(FlagShowHandler))...))
	flag.Methods("PUT").Handler(negroni.New(append(config["FlagUpdate"], negroni.HandlerFunc(FlagUpdateHandler))...))
	flag.Methods("DELETE").Handler(negroni.New(append(config["FlagDelete"], negroni.HandlerFunc(FlagDeleteHandler))...))

	relatedpersonBase := router.Path("/RelatedPerson").Subrouter()
	relatedpersonBase.Methods("GET").Handler(negroni.New(append(config["RelatedPersonIndex"], negroni.HandlerFunc(RelatedPersonIndexHandler))...))
	relatedpersonBase.Methods("POST").Handler(negroni.New(append(config["RelatedPersonCreate"], negroni.HandlerFunc(RelatedPersonCreateHandler))...))

	relatedperson := router.Path("/RelatedPerson/{id}").Subrouter()
	relatedperson.Methods("GET").Handler(negroni.New(append(config["RelatedPersonShow"], negroni.HandlerFunc(RelatedPersonShowHandler))...))
	relatedperson.Methods("PUT").Handler(negroni.New(append(config["RelatedPersonUpdate"], negroni.HandlerFunc(RelatedPersonUpdateHandler))...))
	relatedperson.Methods("DELETE").Handler(negroni.New(append(config["RelatedPersonDelete"], negroni.HandlerFunc(RelatedPersonDeleteHandler))...))

	supplyrequestBase := router.Path("/SupplyRequest").Subrouter()
	supplyrequestBase.Methods("GET").Handler(negroni.New(append(config["SupplyRequestIndex"], negroni.HandlerFunc(SupplyRequestIndexHandler))...))
	supplyrequestBase.Methods("POST").Handler(negroni.New(append(config["SupplyRequestCreate"], negroni.HandlerFunc(SupplyRequestCreateHandler))...))

	supplyrequest := router.Path("/SupplyRequest/{id}").Subrouter()
	supplyrequest.Methods("GET").Handler(negroni.New(append(config["SupplyRequestShow"], negroni.HandlerFunc(SupplyRequestShowHandler))...))
	supplyrequest.Methods("PUT").Handler(negroni.New(append(config["SupplyRequestUpdate"], negroni.HandlerFunc(SupplyRequestUpdateHandler))...))
	supplyrequest.Methods("DELETE").Handler(negroni.New(append(config["SupplyRequestDelete"], negroni.HandlerFunc(SupplyRequestDeleteHandler))...))

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

	enrollmentresponseBase := router.Path("/EnrollmentResponse").Subrouter()
	enrollmentresponseBase.Methods("GET").Handler(negroni.New(append(config["EnrollmentResponseIndex"], negroni.HandlerFunc(EnrollmentResponseIndexHandler))...))
	enrollmentresponseBase.Methods("POST").Handler(negroni.New(append(config["EnrollmentResponseCreate"], negroni.HandlerFunc(EnrollmentResponseCreateHandler))...))

	enrollmentresponse := router.Path("/EnrollmentResponse/{id}").Subrouter()
	enrollmentresponse.Methods("GET").Handler(negroni.New(append(config["EnrollmentResponseShow"], negroni.HandlerFunc(EnrollmentResponseShowHandler))...))
	enrollmentresponse.Methods("PUT").Handler(negroni.New(append(config["EnrollmentResponseUpdate"], negroni.HandlerFunc(EnrollmentResponseUpdateHandler))...))
	enrollmentresponse.Methods("DELETE").Handler(negroni.New(append(config["EnrollmentResponseDelete"], negroni.HandlerFunc(EnrollmentResponseDeleteHandler))...))

	binaryBase := router.Path("/Binary").Subrouter()
	binaryBase.Methods("GET").Handler(negroni.New(append(config["BinaryIndex"], negroni.HandlerFunc(BinaryIndexHandler))...))
	binaryBase.Methods("POST").Handler(negroni.New(append(config["BinaryCreate"], negroni.HandlerFunc(BinaryCreateHandler))...))

	binary := router.Path("/Binary/{id}").Subrouter()
	binary.Methods("GET").Handler(negroni.New(append(config["BinaryShow"], negroni.HandlerFunc(BinaryShowHandler))...))
	binary.Methods("PUT").Handler(negroni.New(append(config["BinaryUpdate"], negroni.HandlerFunc(BinaryUpdateHandler))...))
	binary.Methods("DELETE").Handler(negroni.New(append(config["BinaryDelete"], negroni.HandlerFunc(BinaryDeleteHandler))...))

	medicationstatementBase := router.Path("/MedicationStatement").Subrouter()
	medicationstatementBase.Methods("GET").Handler(negroni.New(append(config["MedicationStatementIndex"], negroni.HandlerFunc(MedicationStatementIndexHandler))...))
	medicationstatementBase.Methods("POST").Handler(negroni.New(append(config["MedicationStatementCreate"], negroni.HandlerFunc(MedicationStatementCreateHandler))...))

	medicationstatement := router.Path("/MedicationStatement/{id}").Subrouter()
	medicationstatement.Methods("GET").Handler(negroni.New(append(config["MedicationStatementShow"], negroni.HandlerFunc(MedicationStatementShowHandler))...))
	medicationstatement.Methods("PUT").Handler(negroni.New(append(config["MedicationStatementUpdate"], negroni.HandlerFunc(MedicationStatementUpdateHandler))...))
	medicationstatement.Methods("DELETE").Handler(negroni.New(append(config["MedicationStatementDelete"], negroni.HandlerFunc(MedicationStatementDeleteHandler))...))

	contractBase := router.Path("/Contract").Subrouter()
	contractBase.Methods("GET").Handler(negroni.New(append(config["ContractIndex"], negroni.HandlerFunc(ContractIndexHandler))...))
	contractBase.Methods("POST").Handler(negroni.New(append(config["ContractCreate"], negroni.HandlerFunc(ContractCreateHandler))...))

	contract := router.Path("/Contract/{id}").Subrouter()
	contract.Methods("GET").Handler(negroni.New(append(config["ContractShow"], negroni.HandlerFunc(ContractShowHandler))...))
	contract.Methods("PUT").Handler(negroni.New(append(config["ContractUpdate"], negroni.HandlerFunc(ContractUpdateHandler))...))
	contract.Methods("DELETE").Handler(negroni.New(append(config["ContractDelete"], negroni.HandlerFunc(ContractDeleteHandler))...))

	personBase := router.Path("/Person").Subrouter()
	personBase.Methods("GET").Handler(negroni.New(append(config["PersonIndex"], negroni.HandlerFunc(PersonIndexHandler))...))
	personBase.Methods("POST").Handler(negroni.New(append(config["PersonCreate"], negroni.HandlerFunc(PersonCreateHandler))...))

	person := router.Path("/Person/{id}").Subrouter()
	person.Methods("GET").Handler(negroni.New(append(config["PersonShow"], negroni.HandlerFunc(PersonShowHandler))...))
	person.Methods("PUT").Handler(negroni.New(append(config["PersonUpdate"], negroni.HandlerFunc(PersonUpdateHandler))...))
	person.Methods("DELETE").Handler(negroni.New(append(config["PersonDelete"], negroni.HandlerFunc(PersonDeleteHandler))...))

	communicationrequestBase := router.Path("/CommunicationRequest").Subrouter()
	communicationrequestBase.Methods("GET").Handler(negroni.New(append(config["CommunicationRequestIndex"], negroni.HandlerFunc(CommunicationRequestIndexHandler))...))
	communicationrequestBase.Methods("POST").Handler(negroni.New(append(config["CommunicationRequestCreate"], negroni.HandlerFunc(CommunicationRequestCreateHandler))...))

	communicationrequest := router.Path("/CommunicationRequest/{id}").Subrouter()
	communicationrequest.Methods("GET").Handler(negroni.New(append(config["CommunicationRequestShow"], negroni.HandlerFunc(CommunicationRequestShowHandler))...))
	communicationrequest.Methods("PUT").Handler(negroni.New(append(config["CommunicationRequestUpdate"], negroni.HandlerFunc(CommunicationRequestUpdateHandler))...))
	communicationrequest.Methods("DELETE").Handler(negroni.New(append(config["CommunicationRequestDelete"], negroni.HandlerFunc(CommunicationRequestDeleteHandler))...))

	riskassessmentBase := router.Path("/RiskAssessment").Subrouter()
	riskassessmentBase.Methods("GET").Handler(negroni.New(append(config["RiskAssessmentIndex"], negroni.HandlerFunc(RiskAssessmentIndexHandler))...))
	riskassessmentBase.Methods("POST").Handler(negroni.New(append(config["RiskAssessmentCreate"], negroni.HandlerFunc(RiskAssessmentCreateHandler))...))

	riskassessment := router.Path("/RiskAssessment/{id}").Subrouter()
	riskassessment.Methods("GET").Handler(negroni.New(append(config["RiskAssessmentShow"], negroni.HandlerFunc(RiskAssessmentShowHandler))...))
	riskassessment.Methods("PUT").Handler(negroni.New(append(config["RiskAssessmentUpdate"], negroni.HandlerFunc(RiskAssessmentUpdateHandler))...))
	riskassessment.Methods("DELETE").Handler(negroni.New(append(config["RiskAssessmentDelete"], negroni.HandlerFunc(RiskAssessmentDeleteHandler))...))

	testscriptBase := router.Path("/TestScript").Subrouter()
	testscriptBase.Methods("GET").Handler(negroni.New(append(config["TestScriptIndex"], negroni.HandlerFunc(TestScriptIndexHandler))...))
	testscriptBase.Methods("POST").Handler(negroni.New(append(config["TestScriptCreate"], negroni.HandlerFunc(TestScriptCreateHandler))...))

	testscript := router.Path("/TestScript/{id}").Subrouter()
	testscript.Methods("GET").Handler(negroni.New(append(config["TestScriptShow"], negroni.HandlerFunc(TestScriptShowHandler))...))
	testscript.Methods("PUT").Handler(negroni.New(append(config["TestScriptUpdate"], negroni.HandlerFunc(TestScriptUpdateHandler))...))
	testscript.Methods("DELETE").Handler(negroni.New(append(config["TestScriptDelete"], negroni.HandlerFunc(TestScriptDeleteHandler))...))

	basicBase := router.Path("/Basic").Subrouter()
	basicBase.Methods("GET").Handler(negroni.New(append(config["BasicIndex"], negroni.HandlerFunc(BasicIndexHandler))...))
	basicBase.Methods("POST").Handler(negroni.New(append(config["BasicCreate"], negroni.HandlerFunc(BasicCreateHandler))...))

	basic := router.Path("/Basic/{id}").Subrouter()
	basic.Methods("GET").Handler(negroni.New(append(config["BasicShow"], negroni.HandlerFunc(BasicShowHandler))...))
	basic.Methods("PUT").Handler(negroni.New(append(config["BasicUpdate"], negroni.HandlerFunc(BasicUpdateHandler))...))
	basic.Methods("DELETE").Handler(negroni.New(append(config["BasicDelete"], negroni.HandlerFunc(BasicDeleteHandler))...))

	groupBase := router.Path("/Group").Subrouter()
	groupBase.Methods("GET").Handler(negroni.New(append(config["GroupIndex"], negroni.HandlerFunc(GroupIndexHandler))...))
	groupBase.Methods("POST").Handler(negroni.New(append(config["GroupCreate"], negroni.HandlerFunc(GroupCreateHandler))...))

	group := router.Path("/Group/{id}").Subrouter()
	group.Methods("GET").Handler(negroni.New(append(config["GroupShow"], negroni.HandlerFunc(GroupShowHandler))...))
	group.Methods("PUT").Handler(negroni.New(append(config["GroupUpdate"], negroni.HandlerFunc(GroupUpdateHandler))...))
	group.Methods("DELETE").Handler(negroni.New(append(config["GroupDelete"], negroni.HandlerFunc(GroupDeleteHandler))...))

	paymentnoticeBase := router.Path("/PaymentNotice").Subrouter()
	paymentnoticeBase.Methods("GET").Handler(negroni.New(append(config["PaymentNoticeIndex"], negroni.HandlerFunc(PaymentNoticeIndexHandler))...))
	paymentnoticeBase.Methods("POST").Handler(negroni.New(append(config["PaymentNoticeCreate"], negroni.HandlerFunc(PaymentNoticeCreateHandler))...))

	paymentnotice := router.Path("/PaymentNotice/{id}").Subrouter()
	paymentnotice.Methods("GET").Handler(negroni.New(append(config["PaymentNoticeShow"], negroni.HandlerFunc(PaymentNoticeShowHandler))...))
	paymentnotice.Methods("PUT").Handler(negroni.New(append(config["PaymentNoticeUpdate"], negroni.HandlerFunc(PaymentNoticeUpdateHandler))...))
	paymentnotice.Methods("DELETE").Handler(negroni.New(append(config["PaymentNoticeDelete"], negroni.HandlerFunc(PaymentNoticeDeleteHandler))...))

	organizationBase := router.Path("/Organization").Subrouter()
	organizationBase.Methods("GET").Handler(negroni.New(append(config["OrganizationIndex"], negroni.HandlerFunc(OrganizationIndexHandler))...))
	organizationBase.Methods("POST").Handler(negroni.New(append(config["OrganizationCreate"], negroni.HandlerFunc(OrganizationCreateHandler))...))

	organization := router.Path("/Organization/{id}").Subrouter()
	organization.Methods("GET").Handler(negroni.New(append(config["OrganizationShow"], negroni.HandlerFunc(OrganizationShowHandler))...))
	organization.Methods("PUT").Handler(negroni.New(append(config["OrganizationUpdate"], negroni.HandlerFunc(OrganizationUpdateHandler))...))
	organization.Methods("DELETE").Handler(negroni.New(append(config["OrganizationDelete"], negroni.HandlerFunc(OrganizationDeleteHandler))...))

	claimresponseBase := router.Path("/ClaimResponse").Subrouter()
	claimresponseBase.Methods("GET").Handler(negroni.New(append(config["ClaimResponseIndex"], negroni.HandlerFunc(ClaimResponseIndexHandler))...))
	claimresponseBase.Methods("POST").Handler(negroni.New(append(config["ClaimResponseCreate"], negroni.HandlerFunc(ClaimResponseCreateHandler))...))

	claimresponse := router.Path("/ClaimResponse/{id}").Subrouter()
	claimresponse.Methods("GET").Handler(negroni.New(append(config["ClaimResponseShow"], negroni.HandlerFunc(ClaimResponseShowHandler))...))
	claimresponse.Methods("PUT").Handler(negroni.New(append(config["ClaimResponseUpdate"], negroni.HandlerFunc(ClaimResponseUpdateHandler))...))
	claimresponse.Methods("DELETE").Handler(negroni.New(append(config["ClaimResponseDelete"], negroni.HandlerFunc(ClaimResponseDeleteHandler))...))

	eligibilityrequestBase := router.Path("/EligibilityRequest").Subrouter()
	eligibilityrequestBase.Methods("GET").Handler(negroni.New(append(config["EligibilityRequestIndex"], negroni.HandlerFunc(EligibilityRequestIndexHandler))...))
	eligibilityrequestBase.Methods("POST").Handler(negroni.New(append(config["EligibilityRequestCreate"], negroni.HandlerFunc(EligibilityRequestCreateHandler))...))

	eligibilityrequest := router.Path("/EligibilityRequest/{id}").Subrouter()
	eligibilityrequest.Methods("GET").Handler(negroni.New(append(config["EligibilityRequestShow"], negroni.HandlerFunc(EligibilityRequestShowHandler))...))
	eligibilityrequest.Methods("PUT").Handler(negroni.New(append(config["EligibilityRequestUpdate"], negroni.HandlerFunc(EligibilityRequestUpdateHandler))...))
	eligibilityrequest.Methods("DELETE").Handler(negroni.New(append(config["EligibilityRequestDelete"], negroni.HandlerFunc(EligibilityRequestDeleteHandler))...))

	processrequestBase := router.Path("/ProcessRequest").Subrouter()
	processrequestBase.Methods("GET").Handler(negroni.New(append(config["ProcessRequestIndex"], negroni.HandlerFunc(ProcessRequestIndexHandler))...))
	processrequestBase.Methods("POST").Handler(negroni.New(append(config["ProcessRequestCreate"], negroni.HandlerFunc(ProcessRequestCreateHandler))...))

	processrequest := router.Path("/ProcessRequest/{id}").Subrouter()
	processrequest.Methods("GET").Handler(negroni.New(append(config["ProcessRequestShow"], negroni.HandlerFunc(ProcessRequestShowHandler))...))
	processrequest.Methods("PUT").Handler(negroni.New(append(config["ProcessRequestUpdate"], negroni.HandlerFunc(ProcessRequestUpdateHandler))...))
	processrequest.Methods("DELETE").Handler(negroni.New(append(config["ProcessRequestDelete"], negroni.HandlerFunc(ProcessRequestDeleteHandler))...))

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

	imagingobjectselectionBase := router.Path("/ImagingObjectSelection").Subrouter()
	imagingobjectselectionBase.Methods("GET").Handler(negroni.New(append(config["ImagingObjectSelectionIndex"], negroni.HandlerFunc(ImagingObjectSelectionIndexHandler))...))
	imagingobjectselectionBase.Methods("POST").Handler(negroni.New(append(config["ImagingObjectSelectionCreate"], negroni.HandlerFunc(ImagingObjectSelectionCreateHandler))...))

	imagingobjectselection := router.Path("/ImagingObjectSelection/{id}").Subrouter()
	imagingobjectselection.Methods("GET").Handler(negroni.New(append(config["ImagingObjectSelectionShow"], negroni.HandlerFunc(ImagingObjectSelectionShowHandler))...))
	imagingobjectselection.Methods("PUT").Handler(negroni.New(append(config["ImagingObjectSelectionUpdate"], negroni.HandlerFunc(ImagingObjectSelectionUpdateHandler))...))
	imagingobjectselection.Methods("DELETE").Handler(negroni.New(append(config["ImagingObjectSelectionDelete"], negroni.HandlerFunc(ImagingObjectSelectionDeleteHandler))...))

	healthcareserviceBase := router.Path("/HealthcareService").Subrouter()
	healthcareserviceBase.Methods("GET").Handler(negroni.New(append(config["HealthcareServiceIndex"], negroni.HandlerFunc(HealthcareServiceIndexHandler))...))
	healthcareserviceBase.Methods("POST").Handler(negroni.New(append(config["HealthcareServiceCreate"], negroni.HandlerFunc(HealthcareServiceCreateHandler))...))

	healthcareservice := router.Path("/HealthcareService/{id}").Subrouter()
	healthcareservice.Methods("GET").Handler(negroni.New(append(config["HealthcareServiceShow"], negroni.HandlerFunc(HealthcareServiceShowHandler))...))
	healthcareservice.Methods("PUT").Handler(negroni.New(append(config["HealthcareServiceUpdate"], negroni.HandlerFunc(HealthcareServiceUpdateHandler))...))
	healthcareservice.Methods("DELETE").Handler(negroni.New(append(config["HealthcareServiceDelete"], negroni.HandlerFunc(HealthcareServiceDeleteHandler))...))

	dataelementBase := router.Path("/DataElement").Subrouter()
	dataelementBase.Methods("GET").Handler(negroni.New(append(config["DataElementIndex"], negroni.HandlerFunc(DataElementIndexHandler))...))
	dataelementBase.Methods("POST").Handler(negroni.New(append(config["DataElementCreate"], negroni.HandlerFunc(DataElementCreateHandler))...))

	dataelement := router.Path("/DataElement/{id}").Subrouter()
	dataelement.Methods("GET").Handler(negroni.New(append(config["DataElementShow"], negroni.HandlerFunc(DataElementShowHandler))...))
	dataelement.Methods("PUT").Handler(negroni.New(append(config["DataElementUpdate"], negroni.HandlerFunc(DataElementUpdateHandler))...))
	dataelement.Methods("DELETE").Handler(negroni.New(append(config["DataElementDelete"], negroni.HandlerFunc(DataElementDeleteHandler))...))

	devicecomponentBase := router.Path("/DeviceComponent").Subrouter()
	devicecomponentBase.Methods("GET").Handler(negroni.New(append(config["DeviceComponentIndex"], negroni.HandlerFunc(DeviceComponentIndexHandler))...))
	devicecomponentBase.Methods("POST").Handler(negroni.New(append(config["DeviceComponentCreate"], negroni.HandlerFunc(DeviceComponentCreateHandler))...))

	devicecomponent := router.Path("/DeviceComponent/{id}").Subrouter()
	devicecomponent.Methods("GET").Handler(negroni.New(append(config["DeviceComponentShow"], negroni.HandlerFunc(DeviceComponentShowHandler))...))
	devicecomponent.Methods("PUT").Handler(negroni.New(append(config["DeviceComponentUpdate"], negroni.HandlerFunc(DeviceComponentUpdateHandler))...))
	devicecomponent.Methods("DELETE").Handler(negroni.New(append(config["DeviceComponentDelete"], negroni.HandlerFunc(DeviceComponentDeleteHandler))...))

	familymemberhistoryBase := router.Path("/FamilyMemberHistory").Subrouter()
	familymemberhistoryBase.Methods("GET").Handler(negroni.New(append(config["FamilyMemberHistoryIndex"], negroni.HandlerFunc(FamilyMemberHistoryIndexHandler))...))
	familymemberhistoryBase.Methods("POST").Handler(negroni.New(append(config["FamilyMemberHistoryCreate"], negroni.HandlerFunc(FamilyMemberHistoryCreateHandler))...))

	familymemberhistory := router.Path("/FamilyMemberHistory/{id}").Subrouter()
	familymemberhistory.Methods("GET").Handler(negroni.New(append(config["FamilyMemberHistoryShow"], negroni.HandlerFunc(FamilyMemberHistoryShowHandler))...))
	familymemberhistory.Methods("PUT").Handler(negroni.New(append(config["FamilyMemberHistoryUpdate"], negroni.HandlerFunc(FamilyMemberHistoryUpdateHandler))...))
	familymemberhistory.Methods("DELETE").Handler(negroni.New(append(config["FamilyMemberHistoryDelete"], negroni.HandlerFunc(FamilyMemberHistoryDeleteHandler))...))

	questionnaireanswersBase := router.Path("/QuestionnaireAnswers").Subrouter()
	questionnaireanswersBase.Methods("GET").Handler(negroni.New(append(config["QuestionnaireAnswersIndex"], negroni.HandlerFunc(QuestionnaireAnswersIndexHandler))...))
	questionnaireanswersBase.Methods("POST").Handler(negroni.New(append(config["QuestionnaireAnswersCreate"], negroni.HandlerFunc(QuestionnaireAnswersCreateHandler))...))

	questionnaireanswers := router.Path("/QuestionnaireAnswers/{id}").Subrouter()
	questionnaireanswers.Methods("GET").Handler(negroni.New(append(config["QuestionnaireAnswersShow"], negroni.HandlerFunc(QuestionnaireAnswersShowHandler))...))
	questionnaireanswers.Methods("PUT").Handler(negroni.New(append(config["QuestionnaireAnswersUpdate"], negroni.HandlerFunc(QuestionnaireAnswersUpdateHandler))...))
	questionnaireanswers.Methods("DELETE").Handler(negroni.New(append(config["QuestionnaireAnswersDelete"], negroni.HandlerFunc(QuestionnaireAnswersDeleteHandler))...))

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

	auditeventBase := router.Path("/AuditEvent").Subrouter()
	auditeventBase.Methods("GET").Handler(negroni.New(append(config["AuditEventIndex"], negroni.HandlerFunc(AuditEventIndexHandler))...))
	auditeventBase.Methods("POST").Handler(negroni.New(append(config["AuditEventCreate"], negroni.HandlerFunc(AuditEventCreateHandler))...))

	auditevent := router.Path("/AuditEvent/{id}").Subrouter()
	auditevent.Methods("GET").Handler(negroni.New(append(config["AuditEventShow"], negroni.HandlerFunc(AuditEventShowHandler))...))
	auditevent.Methods("PUT").Handler(negroni.New(append(config["AuditEventUpdate"], negroni.HandlerFunc(AuditEventUpdateHandler))...))
	auditevent.Methods("DELETE").Handler(negroni.New(append(config["AuditEventDelete"], negroni.HandlerFunc(AuditEventDeleteHandler))...))

	searchparameterBase := router.Path("/SearchParameter").Subrouter()
	searchparameterBase.Methods("GET").Handler(negroni.New(append(config["SearchParameterIndex"], negroni.HandlerFunc(SearchParameterIndexHandler))...))
	searchparameterBase.Methods("POST").Handler(negroni.New(append(config["SearchParameterCreate"], negroni.HandlerFunc(SearchParameterCreateHandler))...))

	searchparameter := router.Path("/SearchParameter/{id}").Subrouter()
	searchparameter.Methods("GET").Handler(negroni.New(append(config["SearchParameterShow"], negroni.HandlerFunc(SearchParameterShowHandler))...))
	searchparameter.Methods("PUT").Handler(negroni.New(append(config["SearchParameterUpdate"], negroni.HandlerFunc(SearchParameterUpdateHandler))...))
	searchparameter.Methods("DELETE").Handler(negroni.New(append(config["SearchParameterDelete"], negroni.HandlerFunc(SearchParameterDeleteHandler))...))

	paymentreconciliationBase := router.Path("/PaymentReconciliation").Subrouter()
	paymentreconciliationBase.Methods("GET").Handler(negroni.New(append(config["PaymentReconciliationIndex"], negroni.HandlerFunc(PaymentReconciliationIndexHandler))...))
	paymentreconciliationBase.Methods("POST").Handler(negroni.New(append(config["PaymentReconciliationCreate"], negroni.HandlerFunc(PaymentReconciliationCreateHandler))...))

	paymentreconciliation := router.Path("/PaymentReconciliation/{id}").Subrouter()
	paymentreconciliation.Methods("GET").Handler(negroni.New(append(config["PaymentReconciliationShow"], negroni.HandlerFunc(PaymentReconciliationShowHandler))...))
	paymentreconciliation.Methods("PUT").Handler(negroni.New(append(config["PaymentReconciliationUpdate"], negroni.HandlerFunc(PaymentReconciliationUpdateHandler))...))
	paymentreconciliation.Methods("DELETE").Handler(negroni.New(append(config["PaymentReconciliationDelete"], negroni.HandlerFunc(PaymentReconciliationDeleteHandler))...))

	communicationBase := router.Path("/Communication").Subrouter()
	communicationBase.Methods("GET").Handler(negroni.New(append(config["CommunicationIndex"], negroni.HandlerFunc(CommunicationIndexHandler))...))
	communicationBase.Methods("POST").Handler(negroni.New(append(config["CommunicationCreate"], negroni.HandlerFunc(CommunicationCreateHandler))...))

	communication := router.Path("/Communication/{id}").Subrouter()
	communication.Methods("GET").Handler(negroni.New(append(config["CommunicationShow"], negroni.HandlerFunc(CommunicationShowHandler))...))
	communication.Methods("PUT").Handler(negroni.New(append(config["CommunicationUpdate"], negroni.HandlerFunc(CommunicationUpdateHandler))...))
	communication.Methods("DELETE").Handler(negroni.New(append(config["CommunicationDelete"], negroni.HandlerFunc(CommunicationDeleteHandler))...))

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

	bundleBase := router.Path("/Bundle").Subrouter()
	bundleBase.Methods("GET").Handler(negroni.New(append(config["BundleIndex"], negroni.HandlerFunc(BundleIndexHandler))...))
	bundleBase.Methods("POST").Handler(negroni.New(append(config["BundleCreate"], negroni.HandlerFunc(BundleCreateHandler))...))

	bundle := router.Path("/Bundle/{id}").Subrouter()
	bundle.Methods("GET").Handler(negroni.New(append(config["BundleShow"], negroni.HandlerFunc(BundleShowHandler))...))
	bundle.Methods("PUT").Handler(negroni.New(append(config["BundleUpdate"], negroni.HandlerFunc(BundleUpdateHandler))...))
	bundle.Methods("DELETE").Handler(negroni.New(append(config["BundleDelete"], negroni.HandlerFunc(BundleDeleteHandler))...))

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

	coverageBase := router.Path("/Coverage").Subrouter()
	coverageBase.Methods("GET").Handler(negroni.New(append(config["CoverageIndex"], negroni.HandlerFunc(CoverageIndexHandler))...))
	coverageBase.Methods("POST").Handler(negroni.New(append(config["CoverageCreate"], negroni.HandlerFunc(CoverageCreateHandler))...))

	coverage := router.Path("/Coverage/{id}").Subrouter()
	coverage.Methods("GET").Handler(negroni.New(append(config["CoverageShow"], negroni.HandlerFunc(CoverageShowHandler))...))
	coverage.Methods("PUT").Handler(negroni.New(append(config["CoverageUpdate"], negroni.HandlerFunc(CoverageUpdateHandler))...))
	coverage.Methods("DELETE").Handler(negroni.New(append(config["CoverageDelete"], negroni.HandlerFunc(CoverageDeleteHandler))...))

	deviceusestatementBase := router.Path("/DeviceUseStatement").Subrouter()
	deviceusestatementBase.Methods("GET").Handler(negroni.New(append(config["DeviceUseStatementIndex"], negroni.HandlerFunc(DeviceUseStatementIndexHandler))...))
	deviceusestatementBase.Methods("POST").Handler(negroni.New(append(config["DeviceUseStatementCreate"], negroni.HandlerFunc(DeviceUseStatementCreateHandler))...))

	deviceusestatement := router.Path("/DeviceUseStatement/{id}").Subrouter()
	deviceusestatement.Methods("GET").Handler(negroni.New(append(config["DeviceUseStatementShow"], negroni.HandlerFunc(DeviceUseStatementShowHandler))...))
	deviceusestatement.Methods("PUT").Handler(negroni.New(append(config["DeviceUseStatementUpdate"], negroni.HandlerFunc(DeviceUseStatementUpdateHandler))...))
	deviceusestatement.Methods("DELETE").Handler(negroni.New(append(config["DeviceUseStatementDelete"], negroni.HandlerFunc(DeviceUseStatementDeleteHandler))...))

	processresponseBase := router.Path("/ProcessResponse").Subrouter()
	processresponseBase.Methods("GET").Handler(negroni.New(append(config["ProcessResponseIndex"], negroni.HandlerFunc(ProcessResponseIndexHandler))...))
	processresponseBase.Methods("POST").Handler(negroni.New(append(config["ProcessResponseCreate"], negroni.HandlerFunc(ProcessResponseCreateHandler))...))

	processresponse := router.Path("/ProcessResponse/{id}").Subrouter()
	processresponse.Methods("GET").Handler(negroni.New(append(config["ProcessResponseShow"], negroni.HandlerFunc(ProcessResponseShowHandler))...))
	processresponse.Methods("PUT").Handler(negroni.New(append(config["ProcessResponseUpdate"], negroni.HandlerFunc(ProcessResponseUpdateHandler))...))
	processresponse.Methods("DELETE").Handler(negroni.New(append(config["ProcessResponseDelete"], negroni.HandlerFunc(ProcessResponseDeleteHandler))...))

	namingsystemBase := router.Path("/NamingSystem").Subrouter()
	namingsystemBase.Methods("GET").Handler(negroni.New(append(config["NamingSystemIndex"], negroni.HandlerFunc(NamingSystemIndexHandler))...))
	namingsystemBase.Methods("POST").Handler(negroni.New(append(config["NamingSystemCreate"], negroni.HandlerFunc(NamingSystemCreateHandler))...))

	namingsystem := router.Path("/NamingSystem/{id}").Subrouter()
	namingsystem.Methods("GET").Handler(negroni.New(append(config["NamingSystemShow"], negroni.HandlerFunc(NamingSystemShowHandler))...))
	namingsystem.Methods("PUT").Handler(negroni.New(append(config["NamingSystemUpdate"], negroni.HandlerFunc(NamingSystemUpdateHandler))...))
	namingsystem.Methods("DELETE").Handler(negroni.New(append(config["NamingSystemDelete"], negroni.HandlerFunc(NamingSystemDeleteHandler))...))

	scheduleBase := router.Path("/Schedule").Subrouter()
	scheduleBase.Methods("GET").Handler(negroni.New(append(config["ScheduleIndex"], negroni.HandlerFunc(ScheduleIndexHandler))...))
	scheduleBase.Methods("POST").Handler(negroni.New(append(config["ScheduleCreate"], negroni.HandlerFunc(ScheduleCreateHandler))...))

	schedule := router.Path("/Schedule/{id}").Subrouter()
	schedule.Methods("GET").Handler(negroni.New(append(config["ScheduleShow"], negroni.HandlerFunc(ScheduleShowHandler))...))
	schedule.Methods("PUT").Handler(negroni.New(append(config["ScheduleUpdate"], negroni.HandlerFunc(ScheduleUpdateHandler))...))
	schedule.Methods("DELETE").Handler(negroni.New(append(config["ScheduleDelete"], negroni.HandlerFunc(ScheduleDeleteHandler))...))

	supplydeliveryBase := router.Path("/SupplyDelivery").Subrouter()
	supplydeliveryBase.Methods("GET").Handler(negroni.New(append(config["SupplyDeliveryIndex"], negroni.HandlerFunc(SupplyDeliveryIndexHandler))...))
	supplydeliveryBase.Methods("POST").Handler(negroni.New(append(config["SupplyDeliveryCreate"], negroni.HandlerFunc(SupplyDeliveryCreateHandler))...))

	supplydelivery := router.Path("/SupplyDelivery/{id}").Subrouter()
	supplydelivery.Methods("GET").Handler(negroni.New(append(config["SupplyDeliveryShow"], negroni.HandlerFunc(SupplyDeliveryShowHandler))...))
	supplydelivery.Methods("PUT").Handler(negroni.New(append(config["SupplyDeliveryUpdate"], negroni.HandlerFunc(SupplyDeliveryUpdateHandler))...))
	supplydelivery.Methods("DELETE").Handler(negroni.New(append(config["SupplyDeliveryDelete"], negroni.HandlerFunc(SupplyDeliveryDeleteHandler))...))

	clinicalimpressionBase := router.Path("/ClinicalImpression").Subrouter()
	clinicalimpressionBase.Methods("GET").Handler(negroni.New(append(config["ClinicalImpressionIndex"], negroni.HandlerFunc(ClinicalImpressionIndexHandler))...))
	clinicalimpressionBase.Methods("POST").Handler(negroni.New(append(config["ClinicalImpressionCreate"], negroni.HandlerFunc(ClinicalImpressionCreateHandler))...))

	clinicalimpression := router.Path("/ClinicalImpression/{id}").Subrouter()
	clinicalimpression.Methods("GET").Handler(negroni.New(append(config["ClinicalImpressionShow"], negroni.HandlerFunc(ClinicalImpressionShowHandler))...))
	clinicalimpression.Methods("PUT").Handler(negroni.New(append(config["ClinicalImpressionUpdate"], negroni.HandlerFunc(ClinicalImpressionUpdateHandler))...))
	clinicalimpression.Methods("DELETE").Handler(negroni.New(append(config["ClinicalImpressionDelete"], negroni.HandlerFunc(ClinicalImpressionDeleteHandler))...))

	messageheaderBase := router.Path("/MessageHeader").Subrouter()
	messageheaderBase.Methods("GET").Handler(negroni.New(append(config["MessageHeaderIndex"], negroni.HandlerFunc(MessageHeaderIndexHandler))...))
	messageheaderBase.Methods("POST").Handler(negroni.New(append(config["MessageHeaderCreate"], negroni.HandlerFunc(MessageHeaderCreateHandler))...))

	messageheader := router.Path("/MessageHeader/{id}").Subrouter()
	messageheader.Methods("GET").Handler(negroni.New(append(config["MessageHeaderShow"], negroni.HandlerFunc(MessageHeaderShowHandler))...))
	messageheader.Methods("PUT").Handler(negroni.New(append(config["MessageHeaderUpdate"], negroni.HandlerFunc(MessageHeaderUpdateHandler))...))
	messageheader.Methods("DELETE").Handler(negroni.New(append(config["MessageHeaderDelete"], negroni.HandlerFunc(MessageHeaderDeleteHandler))...))

	claimBase := router.Path("/Claim").Subrouter()
	claimBase.Methods("GET").Handler(negroni.New(append(config["ClaimIndex"], negroni.HandlerFunc(ClaimIndexHandler))...))
	claimBase.Methods("POST").Handler(negroni.New(append(config["ClaimCreate"], negroni.HandlerFunc(ClaimCreateHandler))...))

	claim := router.Path("/Claim/{id}").Subrouter()
	claim.Methods("GET").Handler(negroni.New(append(config["ClaimShow"], negroni.HandlerFunc(ClaimShowHandler))...))
	claim.Methods("PUT").Handler(negroni.New(append(config["ClaimUpdate"], negroni.HandlerFunc(ClaimUpdateHandler))...))
	claim.Methods("DELETE").Handler(negroni.New(append(config["ClaimDelete"], negroni.HandlerFunc(ClaimDeleteHandler))...))

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

	bodysiteBase := router.Path("/BodySite").Subrouter()
	bodysiteBase.Methods("GET").Handler(negroni.New(append(config["BodySiteIndex"], negroni.HandlerFunc(BodySiteIndexHandler))...))
	bodysiteBase.Methods("POST").Handler(negroni.New(append(config["BodySiteCreate"], negroni.HandlerFunc(BodySiteCreateHandler))...))

	bodysite := router.Path("/BodySite/{id}").Subrouter()
	bodysite.Methods("GET").Handler(negroni.New(append(config["BodySiteShow"], negroni.HandlerFunc(BodySiteShowHandler))...))
	bodysite.Methods("PUT").Handler(negroni.New(append(config["BodySiteUpdate"], negroni.HandlerFunc(BodySiteUpdateHandler))...))
	bodysite.Methods("DELETE").Handler(negroni.New(append(config["BodySiteDelete"], negroni.HandlerFunc(BodySiteDeleteHandler))...))

}
