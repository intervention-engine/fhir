package main

import (
	"fmt"
	"gitlab.mitre.org/fhir/server"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func main() {
	var err error

	// Setup the database
	if server.MongoSession, err = mgo.Dial("localhost"); err != nil {
		panic(err)
	}
	log.Println("Connected to mongodb")
	defer server.MongoSession.Close()

	server.Database = server.MongoSession.DB("fhir")
	server.Router.StrictSlash(true)
	server.Router.HandleFunc("/", HomeHandler)

	conditionBase := server.Router.Path("/Condition").Subrouter()
	conditionBase.Methods("GET").HandlerFunc(server.ConditionIndexHandler)
	conditionBase.Methods("POST").HandlerFunc(server.ConditionCreateHandler)

	condition := server.Router.PathPrefix("/Condition/{id}").Subrouter()
	condition.Methods("GET").HandlerFunc(server.ConditionShowHandler)
	condition.Methods("PUT").HandlerFunc(server.ConditionUpdateHandler)
	condition.Methods("DELETE").HandlerFunc(server.ConditionDeleteHandler)

	supplyBase := server.Router.Path("/Supply").Subrouter()
	supplyBase.Methods("GET").HandlerFunc(server.SupplyIndexHandler)
	supplyBase.Methods("POST").HandlerFunc(server.SupplyCreateHandler)

	supply := server.Router.PathPrefix("/Supply/{id}").Subrouter()
	supply.Methods("GET").HandlerFunc(server.SupplyShowHandler)
	supply.Methods("PUT").HandlerFunc(server.SupplyUpdateHandler)
	supply.Methods("DELETE").HandlerFunc(server.SupplyDeleteHandler)

	organizationBase := server.Router.Path("/Organization").Subrouter()
	organizationBase.Methods("GET").HandlerFunc(server.OrganizationIndexHandler)
	organizationBase.Methods("POST").HandlerFunc(server.OrganizationCreateHandler)

	organization := server.Router.PathPrefix("/Organization/{id}").Subrouter()
	organization.Methods("GET").HandlerFunc(server.OrganizationShowHandler)
	organization.Methods("PUT").HandlerFunc(server.OrganizationUpdateHandler)
	organization.Methods("DELETE").HandlerFunc(server.OrganizationDeleteHandler)

	groupBase := server.Router.Path("/Group").Subrouter()
	groupBase.Methods("GET").HandlerFunc(server.GroupIndexHandler)
	groupBase.Methods("POST").HandlerFunc(server.GroupCreateHandler)

	group := server.Router.PathPrefix("/Group/{id}").Subrouter()
	group.Methods("GET").HandlerFunc(server.GroupShowHandler)
	group.Methods("PUT").HandlerFunc(server.GroupUpdateHandler)
	group.Methods("DELETE").HandlerFunc(server.GroupDeleteHandler)

	valuesetBase := server.Router.Path("/ValueSet").Subrouter()
	valuesetBase.Methods("GET").HandlerFunc(server.ValueSetIndexHandler)
	valuesetBase.Methods("POST").HandlerFunc(server.ValueSetCreateHandler)

	valueset := server.Router.PathPrefix("/ValueSet/{id}").Subrouter()
	valueset.Methods("GET").HandlerFunc(server.ValueSetShowHandler)
	valueset.Methods("PUT").HandlerFunc(server.ValueSetUpdateHandler)
	valueset.Methods("DELETE").HandlerFunc(server.ValueSetDeleteHandler)

	immunizationrecommendationBase := server.Router.Path("/ImmunizationRecommendation").Subrouter()
	immunizationrecommendationBase.Methods("GET").HandlerFunc(server.ImmunizationRecommendationIndexHandler)
	immunizationrecommendationBase.Methods("POST").HandlerFunc(server.ImmunizationRecommendationCreateHandler)

	immunizationrecommendation := server.Router.PathPrefix("/ImmunizationRecommendation/{id}").Subrouter()
	immunizationrecommendation.Methods("GET").HandlerFunc(server.ImmunizationRecommendationShowHandler)
	immunizationrecommendation.Methods("PUT").HandlerFunc(server.ImmunizationRecommendationUpdateHandler)
	immunizationrecommendation.Methods("DELETE").HandlerFunc(server.ImmunizationRecommendationDeleteHandler)

	appointmentBase := server.Router.Path("/Appointment").Subrouter()
	appointmentBase.Methods("GET").HandlerFunc(server.AppointmentIndexHandler)
	appointmentBase.Methods("POST").HandlerFunc(server.AppointmentCreateHandler)

	appointment := server.Router.PathPrefix("/Appointment/{id}").Subrouter()
	appointment.Methods("GET").HandlerFunc(server.AppointmentShowHandler)
	appointment.Methods("PUT").HandlerFunc(server.AppointmentUpdateHandler)
	appointment.Methods("DELETE").HandlerFunc(server.AppointmentDeleteHandler)

	medicationdispenseBase := server.Router.Path("/MedicationDispense").Subrouter()
	medicationdispenseBase.Methods("GET").HandlerFunc(server.MedicationDispenseIndexHandler)
	medicationdispenseBase.Methods("POST").HandlerFunc(server.MedicationDispenseCreateHandler)

	medicationdispense := server.Router.PathPrefix("/MedicationDispense/{id}").Subrouter()
	medicationdispense.Methods("GET").HandlerFunc(server.MedicationDispenseShowHandler)
	medicationdispense.Methods("PUT").HandlerFunc(server.MedicationDispenseUpdateHandler)
	medicationdispense.Methods("DELETE").HandlerFunc(server.MedicationDispenseDeleteHandler)

	medicationprescriptionBase := server.Router.Path("/MedicationPrescription").Subrouter()
	medicationprescriptionBase.Methods("GET").HandlerFunc(server.MedicationPrescriptionIndexHandler)
	medicationprescriptionBase.Methods("POST").HandlerFunc(server.MedicationPrescriptionCreateHandler)

	medicationprescription := server.Router.PathPrefix("/MedicationPrescription/{id}").Subrouter()
	medicationprescription.Methods("GET").HandlerFunc(server.MedicationPrescriptionShowHandler)
	medicationprescription.Methods("PUT").HandlerFunc(server.MedicationPrescriptionUpdateHandler)
	medicationprescription.Methods("DELETE").HandlerFunc(server.MedicationPrescriptionDeleteHandler)

	slotBase := server.Router.Path("/Slot").Subrouter()
	slotBase.Methods("GET").HandlerFunc(server.SlotIndexHandler)
	slotBase.Methods("POST").HandlerFunc(server.SlotCreateHandler)

	slot := server.Router.PathPrefix("/Slot/{id}").Subrouter()
	slot.Methods("GET").HandlerFunc(server.SlotShowHandler)
	slot.Methods("PUT").HandlerFunc(server.SlotUpdateHandler)
	slot.Methods("DELETE").HandlerFunc(server.SlotDeleteHandler)

	contraindicationBase := server.Router.Path("/Contraindication").Subrouter()
	contraindicationBase.Methods("GET").HandlerFunc(server.ContraindicationIndexHandler)
	contraindicationBase.Methods("POST").HandlerFunc(server.ContraindicationCreateHandler)

	contraindication := server.Router.PathPrefix("/Contraindication/{id}").Subrouter()
	contraindication.Methods("GET").HandlerFunc(server.ContraindicationShowHandler)
	contraindication.Methods("PUT").HandlerFunc(server.ContraindicationUpdateHandler)
	contraindication.Methods("DELETE").HandlerFunc(server.ContraindicationDeleteHandler)

	appointmentresponseBase := server.Router.Path("/AppointmentResponse").Subrouter()
	appointmentresponseBase.Methods("GET").HandlerFunc(server.AppointmentResponseIndexHandler)
	appointmentresponseBase.Methods("POST").HandlerFunc(server.AppointmentResponseCreateHandler)

	appointmentresponse := server.Router.PathPrefix("/AppointmentResponse/{id}").Subrouter()
	appointmentresponse.Methods("GET").HandlerFunc(server.AppointmentResponseShowHandler)
	appointmentresponse.Methods("PUT").HandlerFunc(server.AppointmentResponseUpdateHandler)
	appointmentresponse.Methods("DELETE").HandlerFunc(server.AppointmentResponseDeleteHandler)

	medicationstatementBase := server.Router.Path("/MedicationStatement").Subrouter()
	medicationstatementBase.Methods("GET").HandlerFunc(server.MedicationStatementIndexHandler)
	medicationstatementBase.Methods("POST").HandlerFunc(server.MedicationStatementCreateHandler)

	medicationstatement := server.Router.PathPrefix("/MedicationStatement/{id}").Subrouter()
	medicationstatement.Methods("GET").HandlerFunc(server.MedicationStatementShowHandler)
	medicationstatement.Methods("PUT").HandlerFunc(server.MedicationStatementUpdateHandler)
	medicationstatement.Methods("DELETE").HandlerFunc(server.MedicationStatementDeleteHandler)

	compositionBase := server.Router.Path("/Composition").Subrouter()
	compositionBase.Methods("GET").HandlerFunc(server.CompositionIndexHandler)
	compositionBase.Methods("POST").HandlerFunc(server.CompositionCreateHandler)

	composition := server.Router.PathPrefix("/Composition/{id}").Subrouter()
	composition.Methods("GET").HandlerFunc(server.CompositionShowHandler)
	composition.Methods("PUT").HandlerFunc(server.CompositionUpdateHandler)
	composition.Methods("DELETE").HandlerFunc(server.CompositionDeleteHandler)

	questionnaireBase := server.Router.Path("/Questionnaire").Subrouter()
	questionnaireBase.Methods("GET").HandlerFunc(server.QuestionnaireIndexHandler)
	questionnaireBase.Methods("POST").HandlerFunc(server.QuestionnaireCreateHandler)

	questionnaire := server.Router.PathPrefix("/Questionnaire/{id}").Subrouter()
	questionnaire.Methods("GET").HandlerFunc(server.QuestionnaireShowHandler)
	questionnaire.Methods("PUT").HandlerFunc(server.QuestionnaireUpdateHandler)
	questionnaire.Methods("DELETE").HandlerFunc(server.QuestionnaireDeleteHandler)

	operationoutcomeBase := server.Router.Path("/OperationOutcome").Subrouter()
	operationoutcomeBase.Methods("GET").HandlerFunc(server.OperationOutcomeIndexHandler)
	operationoutcomeBase.Methods("POST").HandlerFunc(server.OperationOutcomeCreateHandler)

	operationoutcome := server.Router.PathPrefix("/OperationOutcome/{id}").Subrouter()
	operationoutcome.Methods("GET").HandlerFunc(server.OperationOutcomeShowHandler)
	operationoutcome.Methods("PUT").HandlerFunc(server.OperationOutcomeUpdateHandler)
	operationoutcome.Methods("DELETE").HandlerFunc(server.OperationOutcomeDeleteHandler)

	conformanceBase := server.Router.Path("/Conformance").Subrouter()
	conformanceBase.Methods("GET").HandlerFunc(server.ConformanceIndexHandler)
	conformanceBase.Methods("POST").HandlerFunc(server.ConformanceCreateHandler)

	conformance := server.Router.PathPrefix("/Conformance/{id}").Subrouter()
	conformance.Methods("GET").HandlerFunc(server.ConformanceShowHandler)
	conformance.Methods("PUT").HandlerFunc(server.ConformanceUpdateHandler)
	conformance.Methods("DELETE").HandlerFunc(server.ConformanceDeleteHandler)

	mediaBase := server.Router.Path("/Media").Subrouter()
	mediaBase.Methods("GET").HandlerFunc(server.MediaIndexHandler)
	mediaBase.Methods("POST").HandlerFunc(server.MediaCreateHandler)

	media := server.Router.PathPrefix("/Media/{id}").Subrouter()
	media.Methods("GET").HandlerFunc(server.MediaShowHandler)
	media.Methods("PUT").HandlerFunc(server.MediaUpdateHandler)
	media.Methods("DELETE").HandlerFunc(server.MediaDeleteHandler)

	otherBase := server.Router.Path("/Other").Subrouter()
	otherBase.Methods("GET").HandlerFunc(server.OtherIndexHandler)
	otherBase.Methods("POST").HandlerFunc(server.OtherCreateHandler)

	other := server.Router.PathPrefix("/Other/{id}").Subrouter()
	other.Methods("GET").HandlerFunc(server.OtherShowHandler)
	other.Methods("PUT").HandlerFunc(server.OtherUpdateHandler)
	other.Methods("DELETE").HandlerFunc(server.OtherDeleteHandler)

	profileBase := server.Router.Path("/Profile").Subrouter()
	profileBase.Methods("GET").HandlerFunc(server.ProfileIndexHandler)
	profileBase.Methods("POST").HandlerFunc(server.ProfileCreateHandler)

	profile := server.Router.PathPrefix("/Profile/{id}").Subrouter()
	profile.Methods("GET").HandlerFunc(server.ProfileShowHandler)
	profile.Methods("PUT").HandlerFunc(server.ProfileUpdateHandler)
	profile.Methods("DELETE").HandlerFunc(server.ProfileDeleteHandler)

	documentreferenceBase := server.Router.Path("/DocumentReference").Subrouter()
	documentreferenceBase.Methods("GET").HandlerFunc(server.DocumentReferenceIndexHandler)
	documentreferenceBase.Methods("POST").HandlerFunc(server.DocumentReferenceCreateHandler)

	documentreference := server.Router.PathPrefix("/DocumentReference/{id}").Subrouter()
	documentreference.Methods("GET").HandlerFunc(server.DocumentReferenceShowHandler)
	documentreference.Methods("PUT").HandlerFunc(server.DocumentReferenceUpdateHandler)
	documentreference.Methods("DELETE").HandlerFunc(server.DocumentReferenceDeleteHandler)

	immunizationBase := server.Router.Path("/Immunization").Subrouter()
	immunizationBase.Methods("GET").HandlerFunc(server.ImmunizationIndexHandler)
	immunizationBase.Methods("POST").HandlerFunc(server.ImmunizationCreateHandler)

	immunization := server.Router.PathPrefix("/Immunization/{id}").Subrouter()
	immunization.Methods("GET").HandlerFunc(server.ImmunizationShowHandler)
	immunization.Methods("PUT").HandlerFunc(server.ImmunizationUpdateHandler)
	immunization.Methods("DELETE").HandlerFunc(server.ImmunizationDeleteHandler)

	subscriptionBase := server.Router.Path("/Subscription").Subrouter()
	subscriptionBase.Methods("GET").HandlerFunc(server.SubscriptionIndexHandler)
	subscriptionBase.Methods("POST").HandlerFunc(server.SubscriptionCreateHandler)

	subscription := server.Router.PathPrefix("/Subscription/{id}").Subrouter()
	subscription.Methods("GET").HandlerFunc(server.SubscriptionShowHandler)
	subscription.Methods("PUT").HandlerFunc(server.SubscriptionUpdateHandler)
	subscription.Methods("DELETE").HandlerFunc(server.SubscriptionDeleteHandler)

	orderresponseBase := server.Router.Path("/OrderResponse").Subrouter()
	orderresponseBase.Methods("GET").HandlerFunc(server.OrderResponseIndexHandler)
	orderresponseBase.Methods("POST").HandlerFunc(server.OrderResponseCreateHandler)

	orderresponse := server.Router.PathPrefix("/OrderResponse/{id}").Subrouter()
	orderresponse.Methods("GET").HandlerFunc(server.OrderResponseShowHandler)
	orderresponse.Methods("PUT").HandlerFunc(server.OrderResponseUpdateHandler)
	orderresponse.Methods("DELETE").HandlerFunc(server.OrderResponseDeleteHandler)

	conceptmapBase := server.Router.Path("/ConceptMap").Subrouter()
	conceptmapBase.Methods("GET").HandlerFunc(server.ConceptMapIndexHandler)
	conceptmapBase.Methods("POST").HandlerFunc(server.ConceptMapCreateHandler)

	conceptmap := server.Router.PathPrefix("/ConceptMap/{id}").Subrouter()
	conceptmap.Methods("GET").HandlerFunc(server.ConceptMapShowHandler)
	conceptmap.Methods("PUT").HandlerFunc(server.ConceptMapUpdateHandler)
	conceptmap.Methods("DELETE").HandlerFunc(server.ConceptMapDeleteHandler)

	imagingstudyBase := server.Router.Path("/ImagingStudy").Subrouter()
	imagingstudyBase.Methods("GET").HandlerFunc(server.ImagingStudyIndexHandler)
	imagingstudyBase.Methods("POST").HandlerFunc(server.ImagingStudyCreateHandler)

	imagingstudy := server.Router.PathPrefix("/ImagingStudy/{id}").Subrouter()
	imagingstudy.Methods("GET").HandlerFunc(server.ImagingStudyShowHandler)
	imagingstudy.Methods("PUT").HandlerFunc(server.ImagingStudyUpdateHandler)
	imagingstudy.Methods("DELETE").HandlerFunc(server.ImagingStudyDeleteHandler)

	practitionerBase := server.Router.Path("/Practitioner").Subrouter()
	practitionerBase.Methods("GET").HandlerFunc(server.PractitionerIndexHandler)
	practitionerBase.Methods("POST").HandlerFunc(server.PractitionerCreateHandler)

	practitioner := server.Router.PathPrefix("/Practitioner/{id}").Subrouter()
	practitioner.Methods("GET").HandlerFunc(server.PractitionerShowHandler)
	practitioner.Methods("PUT").HandlerFunc(server.PractitionerUpdateHandler)
	practitioner.Methods("DELETE").HandlerFunc(server.PractitionerDeleteHandler)

	careplanBase := server.Router.Path("/CarePlan").Subrouter()
	careplanBase.Methods("GET").HandlerFunc(server.CarePlanIndexHandler)
	careplanBase.Methods("POST").HandlerFunc(server.CarePlanCreateHandler)

	careplan := server.Router.PathPrefix("/CarePlan/{id}").Subrouter()
	careplan.Methods("GET").HandlerFunc(server.CarePlanShowHandler)
	careplan.Methods("PUT").HandlerFunc(server.CarePlanUpdateHandler)
	careplan.Methods("DELETE").HandlerFunc(server.CarePlanDeleteHandler)

	provenanceBase := server.Router.Path("/Provenance").Subrouter()
	provenanceBase.Methods("GET").HandlerFunc(server.ProvenanceIndexHandler)
	provenanceBase.Methods("POST").HandlerFunc(server.ProvenanceCreateHandler)

	provenance := server.Router.PathPrefix("/Provenance/{id}").Subrouter()
	provenance.Methods("GET").HandlerFunc(server.ProvenanceShowHandler)
	provenance.Methods("PUT").HandlerFunc(server.ProvenanceUpdateHandler)
	provenance.Methods("DELETE").HandlerFunc(server.ProvenanceDeleteHandler)

	deviceBase := server.Router.Path("/Device").Subrouter()
	deviceBase.Methods("GET").HandlerFunc(server.DeviceIndexHandler)
	deviceBase.Methods("POST").HandlerFunc(server.DeviceCreateHandler)

	device := server.Router.PathPrefix("/Device/{id}").Subrouter()
	device.Methods("GET").HandlerFunc(server.DeviceShowHandler)
	device.Methods("PUT").HandlerFunc(server.DeviceUpdateHandler)
	device.Methods("DELETE").HandlerFunc(server.DeviceDeleteHandler)

	queryBase := server.Router.Path("/Query").Subrouter()
	queryBase.Methods("GET").HandlerFunc(server.QueryIndexHandler)
	queryBase.Methods("POST").HandlerFunc(server.QueryCreateHandler)

	query := server.Router.PathPrefix("/Query/{id}").Subrouter()
	query.Methods("GET").HandlerFunc(server.QueryShowHandler)
	query.Methods("PUT").HandlerFunc(server.QueryUpdateHandler)
	query.Methods("DELETE").HandlerFunc(server.QueryDeleteHandler)

	orderBase := server.Router.Path("/Order").Subrouter()
	orderBase.Methods("GET").HandlerFunc(server.OrderIndexHandler)
	orderBase.Methods("POST").HandlerFunc(server.OrderCreateHandler)

	order := server.Router.PathPrefix("/Order/{id}").Subrouter()
	order.Methods("GET").HandlerFunc(server.OrderShowHandler)
	order.Methods("PUT").HandlerFunc(server.OrderUpdateHandler)
	order.Methods("DELETE").HandlerFunc(server.OrderDeleteHandler)

	procedureBase := server.Router.Path("/Procedure").Subrouter()
	procedureBase.Methods("GET").HandlerFunc(server.ProcedureIndexHandler)
	procedureBase.Methods("POST").HandlerFunc(server.ProcedureCreateHandler)

	procedure := server.Router.PathPrefix("/Procedure/{id}").Subrouter()
	procedure.Methods("GET").HandlerFunc(server.ProcedureShowHandler)
	procedure.Methods("PUT").HandlerFunc(server.ProcedureUpdateHandler)
	procedure.Methods("DELETE").HandlerFunc(server.ProcedureDeleteHandler)

	substanceBase := server.Router.Path("/Substance").Subrouter()
	substanceBase.Methods("GET").HandlerFunc(server.SubstanceIndexHandler)
	substanceBase.Methods("POST").HandlerFunc(server.SubstanceCreateHandler)

	substance := server.Router.PathPrefix("/Substance/{id}").Subrouter()
	substance.Methods("GET").HandlerFunc(server.SubstanceShowHandler)
	substance.Methods("PUT").HandlerFunc(server.SubstanceUpdateHandler)
	substance.Methods("DELETE").HandlerFunc(server.SubstanceDeleteHandler)

	diagnosticreportBase := server.Router.Path("/DiagnosticReport").Subrouter()
	diagnosticreportBase.Methods("GET").HandlerFunc(server.DiagnosticReportIndexHandler)
	diagnosticreportBase.Methods("POST").HandlerFunc(server.DiagnosticReportCreateHandler)

	diagnosticreport := server.Router.PathPrefix("/DiagnosticReport/{id}").Subrouter()
	diagnosticreport.Methods("GET").HandlerFunc(server.DiagnosticReportShowHandler)
	diagnosticreport.Methods("PUT").HandlerFunc(server.DiagnosticReportUpdateHandler)
	diagnosticreport.Methods("DELETE").HandlerFunc(server.DiagnosticReportDeleteHandler)

	medicationBase := server.Router.Path("/Medication").Subrouter()
	medicationBase.Methods("GET").HandlerFunc(server.MedicationIndexHandler)
	medicationBase.Methods("POST").HandlerFunc(server.MedicationCreateHandler)

	medication := server.Router.PathPrefix("/Medication/{id}").Subrouter()
	medication.Methods("GET").HandlerFunc(server.MedicationShowHandler)
	medication.Methods("PUT").HandlerFunc(server.MedicationUpdateHandler)
	medication.Methods("DELETE").HandlerFunc(server.MedicationDeleteHandler)

	messageheaderBase := server.Router.Path("/MessageHeader").Subrouter()
	messageheaderBase.Methods("GET").HandlerFunc(server.MessageHeaderIndexHandler)
	messageheaderBase.Methods("POST").HandlerFunc(server.MessageHeaderCreateHandler)

	messageheader := server.Router.PathPrefix("/MessageHeader/{id}").Subrouter()
	messageheader.Methods("GET").HandlerFunc(server.MessageHeaderShowHandler)
	messageheader.Methods("PUT").HandlerFunc(server.MessageHeaderUpdateHandler)
	messageheader.Methods("DELETE").HandlerFunc(server.MessageHeaderDeleteHandler)

	documentmanifestBase := server.Router.Path("/DocumentManifest").Subrouter()
	documentmanifestBase.Methods("GET").HandlerFunc(server.DocumentManifestIndexHandler)
	documentmanifestBase.Methods("POST").HandlerFunc(server.DocumentManifestCreateHandler)

	documentmanifest := server.Router.PathPrefix("/DocumentManifest/{id}").Subrouter()
	documentmanifest.Methods("GET").HandlerFunc(server.DocumentManifestShowHandler)
	documentmanifest.Methods("PUT").HandlerFunc(server.DocumentManifestUpdateHandler)
	documentmanifest.Methods("DELETE").HandlerFunc(server.DocumentManifestDeleteHandler)

	dataelementBase := server.Router.Path("/DataElement").Subrouter()
	dataelementBase.Methods("GET").HandlerFunc(server.DataElementIndexHandler)
	dataelementBase.Methods("POST").HandlerFunc(server.DataElementCreateHandler)

	dataelement := server.Router.PathPrefix("/DataElement/{id}").Subrouter()
	dataelement.Methods("GET").HandlerFunc(server.DataElementShowHandler)
	dataelement.Methods("PUT").HandlerFunc(server.DataElementUpdateHandler)
	dataelement.Methods("DELETE").HandlerFunc(server.DataElementDeleteHandler)

	availabilityBase := server.Router.Path("/Availability").Subrouter()
	availabilityBase.Methods("GET").HandlerFunc(server.AvailabilityIndexHandler)
	availabilityBase.Methods("POST").HandlerFunc(server.AvailabilityCreateHandler)

	availability := server.Router.PathPrefix("/Availability/{id}").Subrouter()
	availability.Methods("GET").HandlerFunc(server.AvailabilityShowHandler)
	availability.Methods("PUT").HandlerFunc(server.AvailabilityUpdateHandler)
	availability.Methods("DELETE").HandlerFunc(server.AvailabilityDeleteHandler)

	medicationadministrationBase := server.Router.Path("/MedicationAdministration").Subrouter()
	medicationadministrationBase.Methods("GET").HandlerFunc(server.MedicationAdministrationIndexHandler)
	medicationadministrationBase.Methods("POST").HandlerFunc(server.MedicationAdministrationCreateHandler)

	medicationadministration := server.Router.PathPrefix("/MedicationAdministration/{id}").Subrouter()
	medicationadministration.Methods("GET").HandlerFunc(server.MedicationAdministrationShowHandler)
	medicationadministration.Methods("PUT").HandlerFunc(server.MedicationAdministrationUpdateHandler)
	medicationadministration.Methods("DELETE").HandlerFunc(server.MedicationAdministrationDeleteHandler)

	questionnaireanswersBase := server.Router.Path("/QuestionnaireAnswers").Subrouter()
	questionnaireanswersBase.Methods("GET").HandlerFunc(server.QuestionnaireAnswersIndexHandler)
	questionnaireanswersBase.Methods("POST").HandlerFunc(server.QuestionnaireAnswersCreateHandler)

	questionnaireanswers := server.Router.PathPrefix("/QuestionnaireAnswers/{id}").Subrouter()
	questionnaireanswers.Methods("GET").HandlerFunc(server.QuestionnaireAnswersShowHandler)
	questionnaireanswers.Methods("PUT").HandlerFunc(server.QuestionnaireAnswersUpdateHandler)
	questionnaireanswers.Methods("DELETE").HandlerFunc(server.QuestionnaireAnswersDeleteHandler)

	encounterBase := server.Router.Path("/Encounter").Subrouter()
	encounterBase.Methods("GET").HandlerFunc(server.EncounterIndexHandler)
	encounterBase.Methods("POST").HandlerFunc(server.EncounterCreateHandler)

	encounter := server.Router.PathPrefix("/Encounter/{id}").Subrouter()
	encounter.Methods("GET").HandlerFunc(server.EncounterShowHandler)
	encounter.Methods("PUT").HandlerFunc(server.EncounterUpdateHandler)
	encounter.Methods("DELETE").HandlerFunc(server.EncounterDeleteHandler)

	securityeventBase := server.Router.Path("/SecurityEvent").Subrouter()
	securityeventBase.Methods("GET").HandlerFunc(server.SecurityEventIndexHandler)
	securityeventBase.Methods("POST").HandlerFunc(server.SecurityEventCreateHandler)

	securityevent := server.Router.PathPrefix("/SecurityEvent/{id}").Subrouter()
	securityevent.Methods("GET").HandlerFunc(server.SecurityEventShowHandler)
	securityevent.Methods("PUT").HandlerFunc(server.SecurityEventUpdateHandler)
	securityevent.Methods("DELETE").HandlerFunc(server.SecurityEventDeleteHandler)

	listBase := server.Router.Path("/List").Subrouter()
	listBase.Methods("GET").HandlerFunc(server.ListIndexHandler)
	listBase.Methods("POST").HandlerFunc(server.ListCreateHandler)

	list := server.Router.PathPrefix("/List/{id}").Subrouter()
	list.Methods("GET").HandlerFunc(server.ListShowHandler)
	list.Methods("PUT").HandlerFunc(server.ListUpdateHandler)
	list.Methods("DELETE").HandlerFunc(server.ListDeleteHandler)

	operationdefinitionBase := server.Router.Path("/OperationDefinition").Subrouter()
	operationdefinitionBase.Methods("GET").HandlerFunc(server.OperationDefinitionIndexHandler)
	operationdefinitionBase.Methods("POST").HandlerFunc(server.OperationDefinitionCreateHandler)

	operationdefinition := server.Router.PathPrefix("/OperationDefinition/{id}").Subrouter()
	operationdefinition.Methods("GET").HandlerFunc(server.OperationDefinitionShowHandler)
	operationdefinition.Methods("PUT").HandlerFunc(server.OperationDefinitionUpdateHandler)
	operationdefinition.Methods("DELETE").HandlerFunc(server.OperationDefinitionDeleteHandler)

	deviceobservationreportBase := server.Router.Path("/DeviceObservationReport").Subrouter()
	deviceobservationreportBase.Methods("GET").HandlerFunc(server.DeviceObservationReportIndexHandler)
	deviceobservationreportBase.Methods("POST").HandlerFunc(server.DeviceObservationReportCreateHandler)

	deviceobservationreport := server.Router.PathPrefix("/DeviceObservationReport/{id}").Subrouter()
	deviceobservationreport.Methods("GET").HandlerFunc(server.DeviceObservationReportShowHandler)
	deviceobservationreport.Methods("PUT").HandlerFunc(server.DeviceObservationReportUpdateHandler)
	deviceobservationreport.Methods("DELETE").HandlerFunc(server.DeviceObservationReportDeleteHandler)

	nutritionorderBase := server.Router.Path("/NutritionOrder").Subrouter()
	nutritionorderBase.Methods("GET").HandlerFunc(server.NutritionOrderIndexHandler)
	nutritionorderBase.Methods("POST").HandlerFunc(server.NutritionOrderCreateHandler)

	nutritionorder := server.Router.PathPrefix("/NutritionOrder/{id}").Subrouter()
	nutritionorder.Methods("GET").HandlerFunc(server.NutritionOrderShowHandler)
	nutritionorder.Methods("PUT").HandlerFunc(server.NutritionOrderUpdateHandler)
	nutritionorder.Methods("DELETE").HandlerFunc(server.NutritionOrderDeleteHandler)

	referralrequestBase := server.Router.Path("/ReferralRequest").Subrouter()
	referralrequestBase.Methods("GET").HandlerFunc(server.ReferralRequestIndexHandler)
	referralrequestBase.Methods("POST").HandlerFunc(server.ReferralRequestCreateHandler)

	referralrequest := server.Router.PathPrefix("/ReferralRequest/{id}").Subrouter()
	referralrequest.Methods("GET").HandlerFunc(server.ReferralRequestShowHandler)
	referralrequest.Methods("PUT").HandlerFunc(server.ReferralRequestUpdateHandler)
	referralrequest.Methods("DELETE").HandlerFunc(server.ReferralRequestDeleteHandler)

	riskassessmentBase := server.Router.Path("/RiskAssessment").Subrouter()
	riskassessmentBase.Methods("GET").HandlerFunc(server.RiskAssessmentIndexHandler)
	riskassessmentBase.Methods("POST").HandlerFunc(server.RiskAssessmentCreateHandler)

	riskassessment := server.Router.PathPrefix("/RiskAssessment/{id}").Subrouter()
	riskassessment.Methods("GET").HandlerFunc(server.RiskAssessmentShowHandler)
	riskassessment.Methods("PUT").HandlerFunc(server.RiskAssessmentUpdateHandler)
	riskassessment.Methods("DELETE").HandlerFunc(server.RiskAssessmentDeleteHandler)

	familyhistoryBase := server.Router.Path("/FamilyHistory").Subrouter()
	familyhistoryBase.Methods("GET").HandlerFunc(server.FamilyHistoryIndexHandler)
	familyhistoryBase.Methods("POST").HandlerFunc(server.FamilyHistoryCreateHandler)

	familyhistory := server.Router.PathPrefix("/FamilyHistory/{id}").Subrouter()
	familyhistory.Methods("GET").HandlerFunc(server.FamilyHistoryShowHandler)
	familyhistory.Methods("PUT").HandlerFunc(server.FamilyHistoryUpdateHandler)
	familyhistory.Methods("DELETE").HandlerFunc(server.FamilyHistoryDeleteHandler)

	locationBase := server.Router.Path("/Location").Subrouter()
	locationBase.Methods("GET").HandlerFunc(server.LocationIndexHandler)
	locationBase.Methods("POST").HandlerFunc(server.LocationCreateHandler)

	location := server.Router.PathPrefix("/Location/{id}").Subrouter()
	location.Methods("GET").HandlerFunc(server.LocationShowHandler)
	location.Methods("PUT").HandlerFunc(server.LocationUpdateHandler)
	location.Methods("DELETE").HandlerFunc(server.LocationDeleteHandler)

	allergyintoleranceBase := server.Router.Path("/AllergyIntolerance").Subrouter()
	allergyintoleranceBase.Methods("GET").HandlerFunc(server.AllergyIntoleranceIndexHandler)
	allergyintoleranceBase.Methods("POST").HandlerFunc(server.AllergyIntoleranceCreateHandler)

	allergyintolerance := server.Router.PathPrefix("/AllergyIntolerance/{id}").Subrouter()
	allergyintolerance.Methods("GET").HandlerFunc(server.AllergyIntoleranceShowHandler)
	allergyintolerance.Methods("PUT").HandlerFunc(server.AllergyIntoleranceUpdateHandler)
	allergyintolerance.Methods("DELETE").HandlerFunc(server.AllergyIntoleranceDeleteHandler)

	observationBase := server.Router.Path("/Observation").Subrouter()
	observationBase.Methods("GET").HandlerFunc(server.ObservationIndexHandler)
	observationBase.Methods("POST").HandlerFunc(server.ObservationCreateHandler)

	observation := server.Router.PathPrefix("/Observation/{id}").Subrouter()
	observation.Methods("GET").HandlerFunc(server.ObservationShowHandler)
	observation.Methods("PUT").HandlerFunc(server.ObservationUpdateHandler)
	observation.Methods("DELETE").HandlerFunc(server.ObservationDeleteHandler)

	relatedpersonBase := server.Router.Path("/RelatedPerson").Subrouter()
	relatedpersonBase.Methods("GET").HandlerFunc(server.RelatedPersonIndexHandler)
	relatedpersonBase.Methods("POST").HandlerFunc(server.RelatedPersonCreateHandler)

	relatedperson := server.Router.PathPrefix("/RelatedPerson/{id}").Subrouter()
	relatedperson.Methods("GET").HandlerFunc(server.RelatedPersonShowHandler)
	relatedperson.Methods("PUT").HandlerFunc(server.RelatedPersonUpdateHandler)
	relatedperson.Methods("DELETE").HandlerFunc(server.RelatedPersonDeleteHandler)

	specimenBase := server.Router.Path("/Specimen").Subrouter()
	specimenBase.Methods("GET").HandlerFunc(server.SpecimenIndexHandler)
	specimenBase.Methods("POST").HandlerFunc(server.SpecimenCreateHandler)

	specimen := server.Router.PathPrefix("/Specimen/{id}").Subrouter()
	specimen.Methods("GET").HandlerFunc(server.SpecimenShowHandler)
	specimen.Methods("PUT").HandlerFunc(server.SpecimenUpdateHandler)
	specimen.Methods("DELETE").HandlerFunc(server.SpecimenDeleteHandler)

	alertBase := server.Router.Path("/Alert").Subrouter()
	alertBase.Methods("GET").HandlerFunc(server.AlertIndexHandler)
	alertBase.Methods("POST").HandlerFunc(server.AlertCreateHandler)

	alert := server.Router.PathPrefix("/Alert/{id}").Subrouter()
	alert.Methods("GET").HandlerFunc(server.AlertShowHandler)
	alert.Methods("PUT").HandlerFunc(server.AlertUpdateHandler)
	alert.Methods("DELETE").HandlerFunc(server.AlertDeleteHandler)

	namespaceBase := server.Router.Path("/Namespace").Subrouter()
	namespaceBase.Methods("GET").HandlerFunc(server.NamespaceIndexHandler)
	namespaceBase.Methods("POST").HandlerFunc(server.NamespaceCreateHandler)

	namespace := server.Router.PathPrefix("/Namespace/{id}").Subrouter()
	namespace.Methods("GET").HandlerFunc(server.NamespaceShowHandler)
	namespace.Methods("PUT").HandlerFunc(server.NamespaceUpdateHandler)
	namespace.Methods("DELETE").HandlerFunc(server.NamespaceDeleteHandler)

	patientBase := server.Router.Path("/Patient").Subrouter()
	patientBase.Methods("GET").HandlerFunc(server.PatientIndexHandler)
	patientBase.Methods("POST").HandlerFunc(server.PatientCreateHandler)

	patient := server.Router.PathPrefix("/Patient/{id}").Subrouter()
	patient.Methods("GET").HandlerFunc(server.PatientShowHandler)
	patient.Methods("PUT").HandlerFunc(server.PatientUpdateHandler)
	patient.Methods("DELETE").HandlerFunc(server.PatientDeleteHandler)

	adversereactionBase := server.Router.Path("/AdverseReaction").Subrouter()
	adversereactionBase.Methods("GET").HandlerFunc(server.AdverseReactionIndexHandler)
	adversereactionBase.Methods("POST").HandlerFunc(server.AdverseReactionCreateHandler)

	adversereaction := server.Router.PathPrefix("/AdverseReaction/{id}").Subrouter()
	adversereaction.Methods("GET").HandlerFunc(server.AdverseReactionShowHandler)
	adversereaction.Methods("PUT").HandlerFunc(server.AdverseReactionUpdateHandler)
	adversereaction.Methods("DELETE").HandlerFunc(server.AdverseReactionDeleteHandler)

	diagnosticorderBase := server.Router.Path("/DiagnosticOrder").Subrouter()
	diagnosticorderBase.Methods("GET").HandlerFunc(server.DiagnosticOrderIndexHandler)
	diagnosticorderBase.Methods("POST").HandlerFunc(server.DiagnosticOrderCreateHandler)

	diagnosticorder := server.Router.PathPrefix("/DiagnosticOrder/{id}").Subrouter()
	diagnosticorder.Methods("GET").HandlerFunc(server.DiagnosticOrderShowHandler)
	diagnosticorder.Methods("PUT").HandlerFunc(server.DiagnosticOrderUpdateHandler)
	diagnosticorder.Methods("DELETE").HandlerFunc(server.DiagnosticOrderDeleteHandler)

	http.ListenAndServe(":8080", server.Router)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "FHIR Server Yay! \\o/")
}
