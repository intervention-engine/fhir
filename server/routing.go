package server

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, config map[string][]negroni.Handler) {

	// Batch Support

	batchBase := router.Path("/").Subrouter()
	batchBase.Methods("POST").Handler(negroni.New(append(config["Batch"], negroni.HandlerFunc(BatchHandler))...))

	// Resources

	appointmentController := ResourceController{"Appointment"}
	appointmentBase := router.Path("/Appointment").Subrouter()
	appointmentBase.Methods("GET").Handler(negroni.New(append(config["AppointmentIndex"], negroni.HandlerFunc(appointmentController.IndexHandler))...))
	appointmentBase.Methods("POST").Handler(negroni.New(append(config["AppointmentCreate"], negroni.HandlerFunc(appointmentController.CreateHandler))...))

	appointment := router.Path("/Appointment/{id}").Subrouter()
	appointment.Methods("GET").Handler(negroni.New(append(config["AppointmentShow"], negroni.HandlerFunc(appointmentController.ShowHandler))...))
	appointment.Methods("PUT").Handler(negroni.New(append(config["AppointmentUpdate"], negroni.HandlerFunc(appointmentController.UpdateHandler))...))
	appointment.Methods("DELETE").Handler(negroni.New(append(config["AppointmentDelete"], negroni.HandlerFunc(appointmentController.DeleteHandler))...))

	referralrequestController := ResourceController{"ReferralRequest"}
	referralrequestBase := router.Path("/ReferralRequest").Subrouter()
	referralrequestBase.Methods("GET").Handler(negroni.New(append(config["ReferralRequestIndex"], negroni.HandlerFunc(referralrequestController.IndexHandler))...))
	referralrequestBase.Methods("POST").Handler(negroni.New(append(config["ReferralRequestCreate"], negroni.HandlerFunc(referralrequestController.CreateHandler))...))

	referralrequest := router.Path("/ReferralRequest/{id}").Subrouter()
	referralrequest.Methods("GET").Handler(negroni.New(append(config["ReferralRequestShow"], negroni.HandlerFunc(referralrequestController.ShowHandler))...))
	referralrequest.Methods("PUT").Handler(negroni.New(append(config["ReferralRequestUpdate"], negroni.HandlerFunc(referralrequestController.UpdateHandler))...))
	referralrequest.Methods("DELETE").Handler(negroni.New(append(config["ReferralRequestDelete"], negroni.HandlerFunc(referralrequestController.DeleteHandler))...))

	accountController := ResourceController{"Account"}
	accountBase := router.Path("/Account").Subrouter()
	accountBase.Methods("GET").Handler(negroni.New(append(config["AccountIndex"], negroni.HandlerFunc(accountController.IndexHandler))...))
	accountBase.Methods("POST").Handler(negroni.New(append(config["AccountCreate"], negroni.HandlerFunc(accountController.CreateHandler))...))

	account := router.Path("/Account/{id}").Subrouter()
	account.Methods("GET").Handler(negroni.New(append(config["AccountShow"], negroni.HandlerFunc(accountController.ShowHandler))...))
	account.Methods("PUT").Handler(negroni.New(append(config["AccountUpdate"], negroni.HandlerFunc(accountController.UpdateHandler))...))
	account.Methods("DELETE").Handler(negroni.New(append(config["AccountDelete"], negroni.HandlerFunc(accountController.DeleteHandler))...))

	provenanceController := ResourceController{"Provenance"}
	provenanceBase := router.Path("/Provenance").Subrouter()
	provenanceBase.Methods("GET").Handler(negroni.New(append(config["ProvenanceIndex"], negroni.HandlerFunc(provenanceController.IndexHandler))...))
	provenanceBase.Methods("POST").Handler(negroni.New(append(config["ProvenanceCreate"], negroni.HandlerFunc(provenanceController.CreateHandler))...))

	provenance := router.Path("/Provenance/{id}").Subrouter()
	provenance.Methods("GET").Handler(negroni.New(append(config["ProvenanceShow"], negroni.HandlerFunc(provenanceController.ShowHandler))...))
	provenance.Methods("PUT").Handler(negroni.New(append(config["ProvenanceUpdate"], negroni.HandlerFunc(provenanceController.UpdateHandler))...))
	provenance.Methods("DELETE").Handler(negroni.New(append(config["ProvenanceDelete"], negroni.HandlerFunc(provenanceController.DeleteHandler))...))

	questionnaireController := ResourceController{"Questionnaire"}
	questionnaireBase := router.Path("/Questionnaire").Subrouter()
	questionnaireBase.Methods("GET").Handler(negroni.New(append(config["QuestionnaireIndex"], negroni.HandlerFunc(questionnaireController.IndexHandler))...))
	questionnaireBase.Methods("POST").Handler(negroni.New(append(config["QuestionnaireCreate"], negroni.HandlerFunc(questionnaireController.CreateHandler))...))

	questionnaire := router.Path("/Questionnaire/{id}").Subrouter()
	questionnaire.Methods("GET").Handler(negroni.New(append(config["QuestionnaireShow"], negroni.HandlerFunc(questionnaireController.ShowHandler))...))
	questionnaire.Methods("PUT").Handler(negroni.New(append(config["QuestionnaireUpdate"], negroni.HandlerFunc(questionnaireController.UpdateHandler))...))
	questionnaire.Methods("DELETE").Handler(negroni.New(append(config["QuestionnaireDelete"], negroni.HandlerFunc(questionnaireController.DeleteHandler))...))

	explanationofbenefitController := ResourceController{"ExplanationOfBenefit"}
	explanationofbenefitBase := router.Path("/ExplanationOfBenefit").Subrouter()
	explanationofbenefitBase.Methods("GET").Handler(negroni.New(append(config["ExplanationOfBenefitIndex"], negroni.HandlerFunc(explanationofbenefitController.IndexHandler))...))
	explanationofbenefitBase.Methods("POST").Handler(negroni.New(append(config["ExplanationOfBenefitCreate"], negroni.HandlerFunc(explanationofbenefitController.CreateHandler))...))

	explanationofbenefit := router.Path("/ExplanationOfBenefit/{id}").Subrouter()
	explanationofbenefit.Methods("GET").Handler(negroni.New(append(config["ExplanationOfBenefitShow"], negroni.HandlerFunc(explanationofbenefitController.ShowHandler))...))
	explanationofbenefit.Methods("PUT").Handler(negroni.New(append(config["ExplanationOfBenefitUpdate"], negroni.HandlerFunc(explanationofbenefitController.UpdateHandler))...))
	explanationofbenefit.Methods("DELETE").Handler(negroni.New(append(config["ExplanationOfBenefitDelete"], negroni.HandlerFunc(explanationofbenefitController.DeleteHandler))...))

	documentmanifestController := ResourceController{"DocumentManifest"}
	documentmanifestBase := router.Path("/DocumentManifest").Subrouter()
	documentmanifestBase.Methods("GET").Handler(negroni.New(append(config["DocumentManifestIndex"], negroni.HandlerFunc(documentmanifestController.IndexHandler))...))
	documentmanifestBase.Methods("POST").Handler(negroni.New(append(config["DocumentManifestCreate"], negroni.HandlerFunc(documentmanifestController.CreateHandler))...))

	documentmanifest := router.Path("/DocumentManifest/{id}").Subrouter()
	documentmanifest.Methods("GET").Handler(negroni.New(append(config["DocumentManifestShow"], negroni.HandlerFunc(documentmanifestController.ShowHandler))...))
	documentmanifest.Methods("PUT").Handler(negroni.New(append(config["DocumentManifestUpdate"], negroni.HandlerFunc(documentmanifestController.UpdateHandler))...))
	documentmanifest.Methods("DELETE").Handler(negroni.New(append(config["DocumentManifestDelete"], negroni.HandlerFunc(documentmanifestController.DeleteHandler))...))

	specimenController := ResourceController{"Specimen"}
	specimenBase := router.Path("/Specimen").Subrouter()
	specimenBase.Methods("GET").Handler(negroni.New(append(config["SpecimenIndex"], negroni.HandlerFunc(specimenController.IndexHandler))...))
	specimenBase.Methods("POST").Handler(negroni.New(append(config["SpecimenCreate"], negroni.HandlerFunc(specimenController.CreateHandler))...))

	specimen := router.Path("/Specimen/{id}").Subrouter()
	specimen.Methods("GET").Handler(negroni.New(append(config["SpecimenShow"], negroni.HandlerFunc(specimenController.ShowHandler))...))
	specimen.Methods("PUT").Handler(negroni.New(append(config["SpecimenUpdate"], negroni.HandlerFunc(specimenController.UpdateHandler))...))
	specimen.Methods("DELETE").Handler(negroni.New(append(config["SpecimenDelete"], negroni.HandlerFunc(specimenController.DeleteHandler))...))

	allergyintoleranceController := ResourceController{"AllergyIntolerance"}
	allergyintoleranceBase := router.Path("/AllergyIntolerance").Subrouter()
	allergyintoleranceBase.Methods("GET").Handler(negroni.New(append(config["AllergyIntoleranceIndex"], negroni.HandlerFunc(allergyintoleranceController.IndexHandler))...))
	allergyintoleranceBase.Methods("POST").Handler(negroni.New(append(config["AllergyIntoleranceCreate"], negroni.HandlerFunc(allergyintoleranceController.CreateHandler))...))

	allergyintolerance := router.Path("/AllergyIntolerance/{id}").Subrouter()
	allergyintolerance.Methods("GET").Handler(negroni.New(append(config["AllergyIntoleranceShow"], negroni.HandlerFunc(allergyintoleranceController.ShowHandler))...))
	allergyintolerance.Methods("PUT").Handler(negroni.New(append(config["AllergyIntoleranceUpdate"], negroni.HandlerFunc(allergyintoleranceController.UpdateHandler))...))
	allergyintolerance.Methods("DELETE").Handler(negroni.New(append(config["AllergyIntoleranceDelete"], negroni.HandlerFunc(allergyintoleranceController.DeleteHandler))...))

	careplanController := ResourceController{"CarePlan"}
	careplanBase := router.Path("/CarePlan").Subrouter()
	careplanBase.Methods("GET").Handler(negroni.New(append(config["CarePlanIndex"], negroni.HandlerFunc(careplanController.IndexHandler))...))
	careplanBase.Methods("POST").Handler(negroni.New(append(config["CarePlanCreate"], negroni.HandlerFunc(careplanController.CreateHandler))...))

	careplan := router.Path("/CarePlan/{id}").Subrouter()
	careplan.Methods("GET").Handler(negroni.New(append(config["CarePlanShow"], negroni.HandlerFunc(careplanController.ShowHandler))...))
	careplan.Methods("PUT").Handler(negroni.New(append(config["CarePlanUpdate"], negroni.HandlerFunc(careplanController.UpdateHandler))...))
	careplan.Methods("DELETE").Handler(negroni.New(append(config["CarePlanDelete"], negroni.HandlerFunc(careplanController.DeleteHandler))...))

	goalController := ResourceController{"Goal"}
	goalBase := router.Path("/Goal").Subrouter()
	goalBase.Methods("GET").Handler(negroni.New(append(config["GoalIndex"], negroni.HandlerFunc(goalController.IndexHandler))...))
	goalBase.Methods("POST").Handler(negroni.New(append(config["GoalCreate"], negroni.HandlerFunc(goalController.CreateHandler))...))

	goal := router.Path("/Goal/{id}").Subrouter()
	goal.Methods("GET").Handler(negroni.New(append(config["GoalShow"], negroni.HandlerFunc(goalController.ShowHandler))...))
	goal.Methods("PUT").Handler(negroni.New(append(config["GoalUpdate"], negroni.HandlerFunc(goalController.UpdateHandler))...))
	goal.Methods("DELETE").Handler(negroni.New(append(config["GoalDelete"], negroni.HandlerFunc(goalController.DeleteHandler))...))

	structuredefinitionController := ResourceController{"StructureDefinition"}
	structuredefinitionBase := router.Path("/StructureDefinition").Subrouter()
	structuredefinitionBase.Methods("GET").Handler(negroni.New(append(config["StructureDefinitionIndex"], negroni.HandlerFunc(structuredefinitionController.IndexHandler))...))
	structuredefinitionBase.Methods("POST").Handler(negroni.New(append(config["StructureDefinitionCreate"], negroni.HandlerFunc(structuredefinitionController.CreateHandler))...))

	structuredefinition := router.Path("/StructureDefinition/{id}").Subrouter()
	structuredefinition.Methods("GET").Handler(negroni.New(append(config["StructureDefinitionShow"], negroni.HandlerFunc(structuredefinitionController.ShowHandler))...))
	structuredefinition.Methods("PUT").Handler(negroni.New(append(config["StructureDefinitionUpdate"], negroni.HandlerFunc(structuredefinitionController.UpdateHandler))...))
	structuredefinition.Methods("DELETE").Handler(negroni.New(append(config["StructureDefinitionDelete"], negroni.HandlerFunc(structuredefinitionController.DeleteHandler))...))

	enrollmentrequestController := ResourceController{"EnrollmentRequest"}
	enrollmentrequestBase := router.Path("/EnrollmentRequest").Subrouter()
	enrollmentrequestBase.Methods("GET").Handler(negroni.New(append(config["EnrollmentRequestIndex"], negroni.HandlerFunc(enrollmentrequestController.IndexHandler))...))
	enrollmentrequestBase.Methods("POST").Handler(negroni.New(append(config["EnrollmentRequestCreate"], negroni.HandlerFunc(enrollmentrequestController.CreateHandler))...))

	enrollmentrequest := router.Path("/EnrollmentRequest/{id}").Subrouter()
	enrollmentrequest.Methods("GET").Handler(negroni.New(append(config["EnrollmentRequestShow"], negroni.HandlerFunc(enrollmentrequestController.ShowHandler))...))
	enrollmentrequest.Methods("PUT").Handler(negroni.New(append(config["EnrollmentRequestUpdate"], negroni.HandlerFunc(enrollmentrequestController.UpdateHandler))...))
	enrollmentrequest.Methods("DELETE").Handler(negroni.New(append(config["EnrollmentRequestDelete"], negroni.HandlerFunc(enrollmentrequestController.DeleteHandler))...))

	episodeofcareController := ResourceController{"EpisodeOfCare"}
	episodeofcareBase := router.Path("/EpisodeOfCare").Subrouter()
	episodeofcareBase.Methods("GET").Handler(negroni.New(append(config["EpisodeOfCareIndex"], negroni.HandlerFunc(episodeofcareController.IndexHandler))...))
	episodeofcareBase.Methods("POST").Handler(negroni.New(append(config["EpisodeOfCareCreate"], negroni.HandlerFunc(episodeofcareController.CreateHandler))...))

	episodeofcare := router.Path("/EpisodeOfCare/{id}").Subrouter()
	episodeofcare.Methods("GET").Handler(negroni.New(append(config["EpisodeOfCareShow"], negroni.HandlerFunc(episodeofcareController.ShowHandler))...))
	episodeofcare.Methods("PUT").Handler(negroni.New(append(config["EpisodeOfCareUpdate"], negroni.HandlerFunc(episodeofcareController.UpdateHandler))...))
	episodeofcare.Methods("DELETE").Handler(negroni.New(append(config["EpisodeOfCareDelete"], negroni.HandlerFunc(episodeofcareController.DeleteHandler))...))

	operationoutcomeController := ResourceController{"OperationOutcome"}
	operationoutcomeBase := router.Path("/OperationOutcome").Subrouter()
	operationoutcomeBase.Methods("GET").Handler(negroni.New(append(config["OperationOutcomeIndex"], negroni.HandlerFunc(operationoutcomeController.IndexHandler))...))
	operationoutcomeBase.Methods("POST").Handler(negroni.New(append(config["OperationOutcomeCreate"], negroni.HandlerFunc(operationoutcomeController.CreateHandler))...))

	operationoutcome := router.Path("/OperationOutcome/{id}").Subrouter()
	operationoutcome.Methods("GET").Handler(negroni.New(append(config["OperationOutcomeShow"], negroni.HandlerFunc(operationoutcomeController.ShowHandler))...))
	operationoutcome.Methods("PUT").Handler(negroni.New(append(config["OperationOutcomeUpdate"], negroni.HandlerFunc(operationoutcomeController.UpdateHandler))...))
	operationoutcome.Methods("DELETE").Handler(negroni.New(append(config["OperationOutcomeDelete"], negroni.HandlerFunc(operationoutcomeController.DeleteHandler))...))

	medicationController := ResourceController{"Medication"}
	medicationBase := router.Path("/Medication").Subrouter()
	medicationBase.Methods("GET").Handler(negroni.New(append(config["MedicationIndex"], negroni.HandlerFunc(medicationController.IndexHandler))...))
	medicationBase.Methods("POST").Handler(negroni.New(append(config["MedicationCreate"], negroni.HandlerFunc(medicationController.CreateHandler))...))

	medication := router.Path("/Medication/{id}").Subrouter()
	medication.Methods("GET").Handler(negroni.New(append(config["MedicationShow"], negroni.HandlerFunc(medicationController.ShowHandler))...))
	medication.Methods("PUT").Handler(negroni.New(append(config["MedicationUpdate"], negroni.HandlerFunc(medicationController.UpdateHandler))...))
	medication.Methods("DELETE").Handler(negroni.New(append(config["MedicationDelete"], negroni.HandlerFunc(medicationController.DeleteHandler))...))

	procedureController := ResourceController{"Procedure"}
	procedureBase := router.Path("/Procedure").Subrouter()
	procedureBase.Methods("GET").Handler(negroni.New(append(config["ProcedureIndex"], negroni.HandlerFunc(procedureController.IndexHandler))...))
	procedureBase.Methods("POST").Handler(negroni.New(append(config["ProcedureCreate"], negroni.HandlerFunc(procedureController.CreateHandler))...))

	procedure := router.Path("/Procedure/{id}").Subrouter()
	procedure.Methods("GET").Handler(negroni.New(append(config["ProcedureShow"], negroni.HandlerFunc(procedureController.ShowHandler))...))
	procedure.Methods("PUT").Handler(negroni.New(append(config["ProcedureUpdate"], negroni.HandlerFunc(procedureController.UpdateHandler))...))
	procedure.Methods("DELETE").Handler(negroni.New(append(config["ProcedureDelete"], negroni.HandlerFunc(procedureController.DeleteHandler))...))

	listController := ResourceController{"List"}
	listBase := router.Path("/List").Subrouter()
	listBase.Methods("GET").Handler(negroni.New(append(config["ListIndex"], negroni.HandlerFunc(listController.IndexHandler))...))
	listBase.Methods("POST").Handler(negroni.New(append(config["ListCreate"], negroni.HandlerFunc(listController.CreateHandler))...))

	list := router.Path("/List/{id}").Subrouter()
	list.Methods("GET").Handler(negroni.New(append(config["ListShow"], negroni.HandlerFunc(listController.ShowHandler))...))
	list.Methods("PUT").Handler(negroni.New(append(config["ListUpdate"], negroni.HandlerFunc(listController.UpdateHandler))...))
	list.Methods("DELETE").Handler(negroni.New(append(config["ListDelete"], negroni.HandlerFunc(listController.DeleteHandler))...))

	conceptmapController := ResourceController{"ConceptMap"}
	conceptmapBase := router.Path("/ConceptMap").Subrouter()
	conceptmapBase.Methods("GET").Handler(negroni.New(append(config["ConceptMapIndex"], negroni.HandlerFunc(conceptmapController.IndexHandler))...))
	conceptmapBase.Methods("POST").Handler(negroni.New(append(config["ConceptMapCreate"], negroni.HandlerFunc(conceptmapController.CreateHandler))...))

	conceptmap := router.Path("/ConceptMap/{id}").Subrouter()
	conceptmap.Methods("GET").Handler(negroni.New(append(config["ConceptMapShow"], negroni.HandlerFunc(conceptmapController.ShowHandler))...))
	conceptmap.Methods("PUT").Handler(negroni.New(append(config["ConceptMapUpdate"], negroni.HandlerFunc(conceptmapController.UpdateHandler))...))
	conceptmap.Methods("DELETE").Handler(negroni.New(append(config["ConceptMapDelete"], negroni.HandlerFunc(conceptmapController.DeleteHandler))...))

	subscriptionController := ResourceController{"Subscription"}
	subscriptionBase := router.Path("/Subscription").Subrouter()
	subscriptionBase.Methods("GET").Handler(negroni.New(append(config["SubscriptionIndex"], negroni.HandlerFunc(subscriptionController.IndexHandler))...))
	subscriptionBase.Methods("POST").Handler(negroni.New(append(config["SubscriptionCreate"], negroni.HandlerFunc(subscriptionController.CreateHandler))...))

	subscription := router.Path("/Subscription/{id}").Subrouter()
	subscription.Methods("GET").Handler(negroni.New(append(config["SubscriptionShow"], negroni.HandlerFunc(subscriptionController.ShowHandler))...))
	subscription.Methods("PUT").Handler(negroni.New(append(config["SubscriptionUpdate"], negroni.HandlerFunc(subscriptionController.UpdateHandler))...))
	subscription.Methods("DELETE").Handler(negroni.New(append(config["SubscriptionDelete"], negroni.HandlerFunc(subscriptionController.DeleteHandler))...))

	valuesetController := ResourceController{"ValueSet"}
	valuesetBase := router.Path("/ValueSet").Subrouter()
	valuesetBase.Methods("GET").Handler(negroni.New(append(config["ValueSetIndex"], negroni.HandlerFunc(valuesetController.IndexHandler))...))
	valuesetBase.Methods("POST").Handler(negroni.New(append(config["ValueSetCreate"], negroni.HandlerFunc(valuesetController.CreateHandler))...))

	valueset := router.Path("/ValueSet/{id}").Subrouter()
	valueset.Methods("GET").Handler(negroni.New(append(config["ValueSetShow"], negroni.HandlerFunc(valuesetController.ShowHandler))...))
	valueset.Methods("PUT").Handler(negroni.New(append(config["ValueSetUpdate"], negroni.HandlerFunc(valuesetController.UpdateHandler))...))
	valueset.Methods("DELETE").Handler(negroni.New(append(config["ValueSetDelete"], negroni.HandlerFunc(valuesetController.DeleteHandler))...))

	operationdefinitionController := ResourceController{"OperationDefinition"}
	operationdefinitionBase := router.Path("/OperationDefinition").Subrouter()
	operationdefinitionBase.Methods("GET").Handler(negroni.New(append(config["OperationDefinitionIndex"], negroni.HandlerFunc(operationdefinitionController.IndexHandler))...))
	operationdefinitionBase.Methods("POST").Handler(negroni.New(append(config["OperationDefinitionCreate"], negroni.HandlerFunc(operationdefinitionController.CreateHandler))...))

	operationdefinition := router.Path("/OperationDefinition/{id}").Subrouter()
	operationdefinition.Methods("GET").Handler(negroni.New(append(config["OperationDefinitionShow"], negroni.HandlerFunc(operationdefinitionController.ShowHandler))...))
	operationdefinition.Methods("PUT").Handler(negroni.New(append(config["OperationDefinitionUpdate"], negroni.HandlerFunc(operationdefinitionController.UpdateHandler))...))
	operationdefinition.Methods("DELETE").Handler(negroni.New(append(config["OperationDefinitionDelete"], negroni.HandlerFunc(operationdefinitionController.DeleteHandler))...))

	documentreferenceController := ResourceController{"DocumentReference"}
	documentreferenceBase := router.Path("/DocumentReference").Subrouter()
	documentreferenceBase.Methods("GET").Handler(negroni.New(append(config["DocumentReferenceIndex"], negroni.HandlerFunc(documentreferenceController.IndexHandler))...))
	documentreferenceBase.Methods("POST").Handler(negroni.New(append(config["DocumentReferenceCreate"], negroni.HandlerFunc(documentreferenceController.CreateHandler))...))

	documentreference := router.Path("/DocumentReference/{id}").Subrouter()
	documentreference.Methods("GET").Handler(negroni.New(append(config["DocumentReferenceShow"], negroni.HandlerFunc(documentreferenceController.ShowHandler))...))
	documentreference.Methods("PUT").Handler(negroni.New(append(config["DocumentReferenceUpdate"], negroni.HandlerFunc(documentreferenceController.UpdateHandler))...))
	documentreference.Methods("DELETE").Handler(negroni.New(append(config["DocumentReferenceDelete"], negroni.HandlerFunc(documentreferenceController.DeleteHandler))...))

	orderController := ResourceController{"Order"}
	orderBase := router.Path("/Order").Subrouter()
	orderBase.Methods("GET").Handler(negroni.New(append(config["OrderIndex"], negroni.HandlerFunc(orderController.IndexHandler))...))
	orderBase.Methods("POST").Handler(negroni.New(append(config["OrderCreate"], negroni.HandlerFunc(orderController.CreateHandler))...))

	order := router.Path("/Order/{id}").Subrouter()
	order.Methods("GET").Handler(negroni.New(append(config["OrderShow"], negroni.HandlerFunc(orderController.ShowHandler))...))
	order.Methods("PUT").Handler(negroni.New(append(config["OrderUpdate"], negroni.HandlerFunc(orderController.UpdateHandler))...))
	order.Methods("DELETE").Handler(negroni.New(append(config["OrderDelete"], negroni.HandlerFunc(orderController.DeleteHandler))...))

	immunizationController := ResourceController{"Immunization"}
	immunizationBase := router.Path("/Immunization").Subrouter()
	immunizationBase.Methods("GET").Handler(negroni.New(append(config["ImmunizationIndex"], negroni.HandlerFunc(immunizationController.IndexHandler))...))
	immunizationBase.Methods("POST").Handler(negroni.New(append(config["ImmunizationCreate"], negroni.HandlerFunc(immunizationController.CreateHandler))...))

	immunization := router.Path("/Immunization/{id}").Subrouter()
	immunization.Methods("GET").Handler(negroni.New(append(config["ImmunizationShow"], negroni.HandlerFunc(immunizationController.ShowHandler))...))
	immunization.Methods("PUT").Handler(negroni.New(append(config["ImmunizationUpdate"], negroni.HandlerFunc(immunizationController.UpdateHandler))...))
	immunization.Methods("DELETE").Handler(negroni.New(append(config["ImmunizationDelete"], negroni.HandlerFunc(immunizationController.DeleteHandler))...))

	deviceController := ResourceController{"Device"}
	deviceBase := router.Path("/Device").Subrouter()
	deviceBase.Methods("GET").Handler(negroni.New(append(config["DeviceIndex"], negroni.HandlerFunc(deviceController.IndexHandler))...))
	deviceBase.Methods("POST").Handler(negroni.New(append(config["DeviceCreate"], negroni.HandlerFunc(deviceController.CreateHandler))...))

	device := router.Path("/Device/{id}").Subrouter()
	device.Methods("GET").Handler(negroni.New(append(config["DeviceShow"], negroni.HandlerFunc(deviceController.ShowHandler))...))
	device.Methods("PUT").Handler(negroni.New(append(config["DeviceUpdate"], negroni.HandlerFunc(deviceController.UpdateHandler))...))
	device.Methods("DELETE").Handler(negroni.New(append(config["DeviceDelete"], negroni.HandlerFunc(deviceController.DeleteHandler))...))

	visionprescriptionController := ResourceController{"VisionPrescription"}
	visionprescriptionBase := router.Path("/VisionPrescription").Subrouter()
	visionprescriptionBase.Methods("GET").Handler(negroni.New(append(config["VisionPrescriptionIndex"], negroni.HandlerFunc(visionprescriptionController.IndexHandler))...))
	visionprescriptionBase.Methods("POST").Handler(negroni.New(append(config["VisionPrescriptionCreate"], negroni.HandlerFunc(visionprescriptionController.CreateHandler))...))

	visionprescription := router.Path("/VisionPrescription/{id}").Subrouter()
	visionprescription.Methods("GET").Handler(negroni.New(append(config["VisionPrescriptionShow"], negroni.HandlerFunc(visionprescriptionController.ShowHandler))...))
	visionprescription.Methods("PUT").Handler(negroni.New(append(config["VisionPrescriptionUpdate"], negroni.HandlerFunc(visionprescriptionController.UpdateHandler))...))
	visionprescription.Methods("DELETE").Handler(negroni.New(append(config["VisionPrescriptionDelete"], negroni.HandlerFunc(visionprescriptionController.DeleteHandler))...))

	mediaController := ResourceController{"Media"}
	mediaBase := router.Path("/Media").Subrouter()
	mediaBase.Methods("GET").Handler(negroni.New(append(config["MediaIndex"], negroni.HandlerFunc(mediaController.IndexHandler))...))
	mediaBase.Methods("POST").Handler(negroni.New(append(config["MediaCreate"], negroni.HandlerFunc(mediaController.CreateHandler))...))

	media := router.Path("/Media/{id}").Subrouter()
	media.Methods("GET").Handler(negroni.New(append(config["MediaShow"], negroni.HandlerFunc(mediaController.ShowHandler))...))
	media.Methods("PUT").Handler(negroni.New(append(config["MediaUpdate"], negroni.HandlerFunc(mediaController.UpdateHandler))...))
	media.Methods("DELETE").Handler(negroni.New(append(config["MediaDelete"], negroni.HandlerFunc(mediaController.DeleteHandler))...))

	conformanceController := ResourceController{"Conformance"}
	conformanceBase := router.Path("/Conformance").Subrouter()
	conformanceBase.Methods("GET").Handler(negroni.New(append(config["ConformanceIndex"], negroni.HandlerFunc(conformanceController.IndexHandler))...))
	conformanceBase.Methods("POST").Handler(negroni.New(append(config["ConformanceCreate"], negroni.HandlerFunc(conformanceController.CreateHandler))...))

	conformance := router.Path("/Conformance/{id}").Subrouter()
	conformance.Methods("GET").Handler(negroni.New(append(config["ConformanceShow"], negroni.HandlerFunc(conformanceController.ShowHandler))...))
	conformance.Methods("PUT").Handler(negroni.New(append(config["ConformanceUpdate"], negroni.HandlerFunc(conformanceController.UpdateHandler))...))
	conformance.Methods("DELETE").Handler(negroni.New(append(config["ConformanceDelete"], negroni.HandlerFunc(conformanceController.DeleteHandler))...))

	procedurerequestController := ResourceController{"ProcedureRequest"}
	procedurerequestBase := router.Path("/ProcedureRequest").Subrouter()
	procedurerequestBase.Methods("GET").Handler(negroni.New(append(config["ProcedureRequestIndex"], negroni.HandlerFunc(procedurerequestController.IndexHandler))...))
	procedurerequestBase.Methods("POST").Handler(negroni.New(append(config["ProcedureRequestCreate"], negroni.HandlerFunc(procedurerequestController.CreateHandler))...))

	procedurerequest := router.Path("/ProcedureRequest/{id}").Subrouter()
	procedurerequest.Methods("GET").Handler(negroni.New(append(config["ProcedureRequestShow"], negroni.HandlerFunc(procedurerequestController.ShowHandler))...))
	procedurerequest.Methods("PUT").Handler(negroni.New(append(config["ProcedureRequestUpdate"], negroni.HandlerFunc(procedurerequestController.UpdateHandler))...))
	procedurerequest.Methods("DELETE").Handler(negroni.New(append(config["ProcedureRequestDelete"], negroni.HandlerFunc(procedurerequestController.DeleteHandler))...))

	eligibilityresponseController := ResourceController{"EligibilityResponse"}
	eligibilityresponseBase := router.Path("/EligibilityResponse").Subrouter()
	eligibilityresponseBase.Methods("GET").Handler(negroni.New(append(config["EligibilityResponseIndex"], negroni.HandlerFunc(eligibilityresponseController.IndexHandler))...))
	eligibilityresponseBase.Methods("POST").Handler(negroni.New(append(config["EligibilityResponseCreate"], negroni.HandlerFunc(eligibilityresponseController.CreateHandler))...))

	eligibilityresponse := router.Path("/EligibilityResponse/{id}").Subrouter()
	eligibilityresponse.Methods("GET").Handler(negroni.New(append(config["EligibilityResponseShow"], negroni.HandlerFunc(eligibilityresponseController.ShowHandler))...))
	eligibilityresponse.Methods("PUT").Handler(negroni.New(append(config["EligibilityResponseUpdate"], negroni.HandlerFunc(eligibilityresponseController.UpdateHandler))...))
	eligibilityresponse.Methods("DELETE").Handler(negroni.New(append(config["EligibilityResponseDelete"], negroni.HandlerFunc(eligibilityresponseController.DeleteHandler))...))

	deviceuserequestController := ResourceController{"DeviceUseRequest"}
	deviceuserequestBase := router.Path("/DeviceUseRequest").Subrouter()
	deviceuserequestBase.Methods("GET").Handler(negroni.New(append(config["DeviceUseRequestIndex"], negroni.HandlerFunc(deviceuserequestController.IndexHandler))...))
	deviceuserequestBase.Methods("POST").Handler(negroni.New(append(config["DeviceUseRequestCreate"], negroni.HandlerFunc(deviceuserequestController.CreateHandler))...))

	deviceuserequest := router.Path("/DeviceUseRequest/{id}").Subrouter()
	deviceuserequest.Methods("GET").Handler(negroni.New(append(config["DeviceUseRequestShow"], negroni.HandlerFunc(deviceuserequestController.ShowHandler))...))
	deviceuserequest.Methods("PUT").Handler(negroni.New(append(config["DeviceUseRequestUpdate"], negroni.HandlerFunc(deviceuserequestController.UpdateHandler))...))
	deviceuserequest.Methods("DELETE").Handler(negroni.New(append(config["DeviceUseRequestDelete"], negroni.HandlerFunc(deviceuserequestController.DeleteHandler))...))

	devicemetricController := ResourceController{"DeviceMetric"}
	devicemetricBase := router.Path("/DeviceMetric").Subrouter()
	devicemetricBase.Methods("GET").Handler(negroni.New(append(config["DeviceMetricIndex"], negroni.HandlerFunc(devicemetricController.IndexHandler))...))
	devicemetricBase.Methods("POST").Handler(negroni.New(append(config["DeviceMetricCreate"], negroni.HandlerFunc(devicemetricController.CreateHandler))...))

	devicemetric := router.Path("/DeviceMetric/{id}").Subrouter()
	devicemetric.Methods("GET").Handler(negroni.New(append(config["DeviceMetricShow"], negroni.HandlerFunc(devicemetricController.ShowHandler))...))
	devicemetric.Methods("PUT").Handler(negroni.New(append(config["DeviceMetricUpdate"], negroni.HandlerFunc(devicemetricController.UpdateHandler))...))
	devicemetric.Methods("DELETE").Handler(negroni.New(append(config["DeviceMetricDelete"], negroni.HandlerFunc(devicemetricController.DeleteHandler))...))

	flagController := ResourceController{"Flag"}
	flagBase := router.Path("/Flag").Subrouter()
	flagBase.Methods("GET").Handler(negroni.New(append(config["FlagIndex"], negroni.HandlerFunc(flagController.IndexHandler))...))
	flagBase.Methods("POST").Handler(negroni.New(append(config["FlagCreate"], negroni.HandlerFunc(flagController.CreateHandler))...))

	flag := router.Path("/Flag/{id}").Subrouter()
	flag.Methods("GET").Handler(negroni.New(append(config["FlagShow"], negroni.HandlerFunc(flagController.ShowHandler))...))
	flag.Methods("PUT").Handler(negroni.New(append(config["FlagUpdate"], negroni.HandlerFunc(flagController.UpdateHandler))...))
	flag.Methods("DELETE").Handler(negroni.New(append(config["FlagDelete"], negroni.HandlerFunc(flagController.DeleteHandler))...))

	relatedpersonController := ResourceController{"RelatedPerson"}
	relatedpersonBase := router.Path("/RelatedPerson").Subrouter()
	relatedpersonBase.Methods("GET").Handler(negroni.New(append(config["RelatedPersonIndex"], negroni.HandlerFunc(relatedpersonController.IndexHandler))...))
	relatedpersonBase.Methods("POST").Handler(negroni.New(append(config["RelatedPersonCreate"], negroni.HandlerFunc(relatedpersonController.CreateHandler))...))

	relatedperson := router.Path("/RelatedPerson/{id}").Subrouter()
	relatedperson.Methods("GET").Handler(negroni.New(append(config["RelatedPersonShow"], negroni.HandlerFunc(relatedpersonController.ShowHandler))...))
	relatedperson.Methods("PUT").Handler(negroni.New(append(config["RelatedPersonUpdate"], negroni.HandlerFunc(relatedpersonController.UpdateHandler))...))
	relatedperson.Methods("DELETE").Handler(negroni.New(append(config["RelatedPersonDelete"], negroni.HandlerFunc(relatedpersonController.DeleteHandler))...))

	supplyrequestController := ResourceController{"SupplyRequest"}
	supplyrequestBase := router.Path("/SupplyRequest").Subrouter()
	supplyrequestBase.Methods("GET").Handler(negroni.New(append(config["SupplyRequestIndex"], negroni.HandlerFunc(supplyrequestController.IndexHandler))...))
	supplyrequestBase.Methods("POST").Handler(negroni.New(append(config["SupplyRequestCreate"], negroni.HandlerFunc(supplyrequestController.CreateHandler))...))

	supplyrequest := router.Path("/SupplyRequest/{id}").Subrouter()
	supplyrequest.Methods("GET").Handler(negroni.New(append(config["SupplyRequestShow"], negroni.HandlerFunc(supplyrequestController.ShowHandler))...))
	supplyrequest.Methods("PUT").Handler(negroni.New(append(config["SupplyRequestUpdate"], negroni.HandlerFunc(supplyrequestController.UpdateHandler))...))
	supplyrequest.Methods("DELETE").Handler(negroni.New(append(config["SupplyRequestDelete"], negroni.HandlerFunc(supplyrequestController.DeleteHandler))...))

	practitionerController := ResourceController{"Practitioner"}
	practitionerBase := router.Path("/Practitioner").Subrouter()
	practitionerBase.Methods("GET").Handler(negroni.New(append(config["PractitionerIndex"], negroni.HandlerFunc(practitionerController.IndexHandler))...))
	practitionerBase.Methods("POST").Handler(negroni.New(append(config["PractitionerCreate"], negroni.HandlerFunc(practitionerController.CreateHandler))...))

	practitioner := router.Path("/Practitioner/{id}").Subrouter()
	practitioner.Methods("GET").Handler(negroni.New(append(config["PractitionerShow"], negroni.HandlerFunc(practitionerController.ShowHandler))...))
	practitioner.Methods("PUT").Handler(negroni.New(append(config["PractitionerUpdate"], negroni.HandlerFunc(practitionerController.UpdateHandler))...))
	practitioner.Methods("DELETE").Handler(negroni.New(append(config["PractitionerDelete"], negroni.HandlerFunc(practitionerController.DeleteHandler))...))

	appointmentresponseController := ResourceController{"AppointmentResponse"}
	appointmentresponseBase := router.Path("/AppointmentResponse").Subrouter()
	appointmentresponseBase.Methods("GET").Handler(negroni.New(append(config["AppointmentResponseIndex"], negroni.HandlerFunc(appointmentresponseController.IndexHandler))...))
	appointmentresponseBase.Methods("POST").Handler(negroni.New(append(config["AppointmentResponseCreate"], negroni.HandlerFunc(appointmentresponseController.CreateHandler))...))

	appointmentresponse := router.Path("/AppointmentResponse/{id}").Subrouter()
	appointmentresponse.Methods("GET").Handler(negroni.New(append(config["AppointmentResponseShow"], negroni.HandlerFunc(appointmentresponseController.ShowHandler))...))
	appointmentresponse.Methods("PUT").Handler(negroni.New(append(config["AppointmentResponseUpdate"], negroni.HandlerFunc(appointmentresponseController.UpdateHandler))...))
	appointmentresponse.Methods("DELETE").Handler(negroni.New(append(config["AppointmentResponseDelete"], negroni.HandlerFunc(appointmentresponseController.DeleteHandler))...))

	observationController := ResourceController{"Observation"}
	observationBase := router.Path("/Observation").Subrouter()
	observationBase.Methods("GET").Handler(negroni.New(append(config["ObservationIndex"], negroni.HandlerFunc(observationController.IndexHandler))...))
	observationBase.Methods("POST").Handler(negroni.New(append(config["ObservationCreate"], negroni.HandlerFunc(observationController.CreateHandler))...))

	observation := router.Path("/Observation/{id}").Subrouter()
	observation.Methods("GET").Handler(negroni.New(append(config["ObservationShow"], negroni.HandlerFunc(observationController.ShowHandler))...))
	observation.Methods("PUT").Handler(negroni.New(append(config["ObservationUpdate"], negroni.HandlerFunc(observationController.UpdateHandler))...))
	observation.Methods("DELETE").Handler(negroni.New(append(config["ObservationDelete"], negroni.HandlerFunc(observationController.DeleteHandler))...))

	medicationadministrationController := ResourceController{"MedicationAdministration"}
	medicationadministrationBase := router.Path("/MedicationAdministration").Subrouter()
	medicationadministrationBase.Methods("GET").Handler(negroni.New(append(config["MedicationAdministrationIndex"], negroni.HandlerFunc(medicationadministrationController.IndexHandler))...))
	medicationadministrationBase.Methods("POST").Handler(negroni.New(append(config["MedicationAdministrationCreate"], negroni.HandlerFunc(medicationadministrationController.CreateHandler))...))

	medicationadministration := router.Path("/MedicationAdministration/{id}").Subrouter()
	medicationadministration.Methods("GET").Handler(negroni.New(append(config["MedicationAdministrationShow"], negroni.HandlerFunc(medicationadministrationController.ShowHandler))...))
	medicationadministration.Methods("PUT").Handler(negroni.New(append(config["MedicationAdministrationUpdate"], negroni.HandlerFunc(medicationadministrationController.UpdateHandler))...))
	medicationadministration.Methods("DELETE").Handler(negroni.New(append(config["MedicationAdministrationDelete"], negroni.HandlerFunc(medicationadministrationController.DeleteHandler))...))

	slotController := ResourceController{"Slot"}
	slotBase := router.Path("/Slot").Subrouter()
	slotBase.Methods("GET").Handler(negroni.New(append(config["SlotIndex"], negroni.HandlerFunc(slotController.IndexHandler))...))
	slotBase.Methods("POST").Handler(negroni.New(append(config["SlotCreate"], negroni.HandlerFunc(slotController.CreateHandler))...))

	slot := router.Path("/Slot/{id}").Subrouter()
	slot.Methods("GET").Handler(negroni.New(append(config["SlotShow"], negroni.HandlerFunc(slotController.ShowHandler))...))
	slot.Methods("PUT").Handler(negroni.New(append(config["SlotUpdate"], negroni.HandlerFunc(slotController.UpdateHandler))...))
	slot.Methods("DELETE").Handler(negroni.New(append(config["SlotDelete"], negroni.HandlerFunc(slotController.DeleteHandler))...))

	enrollmentresponseController := ResourceController{"EnrollmentResponse"}
	enrollmentresponseBase := router.Path("/EnrollmentResponse").Subrouter()
	enrollmentresponseBase.Methods("GET").Handler(negroni.New(append(config["EnrollmentResponseIndex"], negroni.HandlerFunc(enrollmentresponseController.IndexHandler))...))
	enrollmentresponseBase.Methods("POST").Handler(negroni.New(append(config["EnrollmentResponseCreate"], negroni.HandlerFunc(enrollmentresponseController.CreateHandler))...))

	enrollmentresponse := router.Path("/EnrollmentResponse/{id}").Subrouter()
	enrollmentresponse.Methods("GET").Handler(negroni.New(append(config["EnrollmentResponseShow"], negroni.HandlerFunc(enrollmentresponseController.ShowHandler))...))
	enrollmentresponse.Methods("PUT").Handler(negroni.New(append(config["EnrollmentResponseUpdate"], negroni.HandlerFunc(enrollmentresponseController.UpdateHandler))...))
	enrollmentresponse.Methods("DELETE").Handler(negroni.New(append(config["EnrollmentResponseDelete"], negroni.HandlerFunc(enrollmentresponseController.DeleteHandler))...))

	binaryController := ResourceController{"Binary"}
	binaryBase := router.Path("/Binary").Subrouter()
	binaryBase.Methods("GET").Handler(negroni.New(append(config["BinaryIndex"], negroni.HandlerFunc(binaryController.IndexHandler))...))
	binaryBase.Methods("POST").Handler(negroni.New(append(config["BinaryCreate"], negroni.HandlerFunc(binaryController.CreateHandler))...))

	binary := router.Path("/Binary/{id}").Subrouter()
	binary.Methods("GET").Handler(negroni.New(append(config["BinaryShow"], negroni.HandlerFunc(binaryController.ShowHandler))...))
	binary.Methods("PUT").Handler(negroni.New(append(config["BinaryUpdate"], negroni.HandlerFunc(binaryController.UpdateHandler))...))
	binary.Methods("DELETE").Handler(negroni.New(append(config["BinaryDelete"], negroni.HandlerFunc(binaryController.DeleteHandler))...))

	medicationstatementController := ResourceController{"MedicationStatement"}
	medicationstatementBase := router.Path("/MedicationStatement").Subrouter()
	medicationstatementBase.Methods("GET").Handler(negroni.New(append(config["MedicationStatementIndex"], negroni.HandlerFunc(medicationstatementController.IndexHandler))...))
	medicationstatementBase.Methods("POST").Handler(negroni.New(append(config["MedicationStatementCreate"], negroni.HandlerFunc(medicationstatementController.CreateHandler))...))

	medicationstatement := router.Path("/MedicationStatement/{id}").Subrouter()
	medicationstatement.Methods("GET").Handler(negroni.New(append(config["MedicationStatementShow"], negroni.HandlerFunc(medicationstatementController.ShowHandler))...))
	medicationstatement.Methods("PUT").Handler(negroni.New(append(config["MedicationStatementUpdate"], negroni.HandlerFunc(medicationstatementController.UpdateHandler))...))
	medicationstatement.Methods("DELETE").Handler(negroni.New(append(config["MedicationStatementDelete"], negroni.HandlerFunc(medicationstatementController.DeleteHandler))...))

	personController := ResourceController{"Person"}
	personBase := router.Path("/Person").Subrouter()
	personBase.Methods("GET").Handler(negroni.New(append(config["PersonIndex"], negroni.HandlerFunc(personController.IndexHandler))...))
	personBase.Methods("POST").Handler(negroni.New(append(config["PersonCreate"], negroni.HandlerFunc(personController.CreateHandler))...))

	person := router.Path("/Person/{id}").Subrouter()
	person.Methods("GET").Handler(negroni.New(append(config["PersonShow"], negroni.HandlerFunc(personController.ShowHandler))...))
	person.Methods("PUT").Handler(negroni.New(append(config["PersonUpdate"], negroni.HandlerFunc(personController.UpdateHandler))...))
	person.Methods("DELETE").Handler(negroni.New(append(config["PersonDelete"], negroni.HandlerFunc(personController.DeleteHandler))...))

	contractController := ResourceController{"Contract"}
	contractBase := router.Path("/Contract").Subrouter()
	contractBase.Methods("GET").Handler(negroni.New(append(config["ContractIndex"], negroni.HandlerFunc(contractController.IndexHandler))...))
	contractBase.Methods("POST").Handler(negroni.New(append(config["ContractCreate"], negroni.HandlerFunc(contractController.CreateHandler))...))

	contract := router.Path("/Contract/{id}").Subrouter()
	contract.Methods("GET").Handler(negroni.New(append(config["ContractShow"], negroni.HandlerFunc(contractController.ShowHandler))...))
	contract.Methods("PUT").Handler(negroni.New(append(config["ContractUpdate"], negroni.HandlerFunc(contractController.UpdateHandler))...))
	contract.Methods("DELETE").Handler(negroni.New(append(config["ContractDelete"], negroni.HandlerFunc(contractController.DeleteHandler))...))

	communicationrequestController := ResourceController{"CommunicationRequest"}
	communicationrequestBase := router.Path("/CommunicationRequest").Subrouter()
	communicationrequestBase.Methods("GET").Handler(negroni.New(append(config["CommunicationRequestIndex"], negroni.HandlerFunc(communicationrequestController.IndexHandler))...))
	communicationrequestBase.Methods("POST").Handler(negroni.New(append(config["CommunicationRequestCreate"], negroni.HandlerFunc(communicationrequestController.CreateHandler))...))

	communicationrequest := router.Path("/CommunicationRequest/{id}").Subrouter()
	communicationrequest.Methods("GET").Handler(negroni.New(append(config["CommunicationRequestShow"], negroni.HandlerFunc(communicationrequestController.ShowHandler))...))
	communicationrequest.Methods("PUT").Handler(negroni.New(append(config["CommunicationRequestUpdate"], negroni.HandlerFunc(communicationrequestController.UpdateHandler))...))
	communicationrequest.Methods("DELETE").Handler(negroni.New(append(config["CommunicationRequestDelete"], negroni.HandlerFunc(communicationrequestController.DeleteHandler))...))

	riskassessmentController := ResourceController{"RiskAssessment"}
	riskassessmentBase := router.Path("/RiskAssessment").Subrouter()
	riskassessmentBase.Methods("GET").Handler(negroni.New(append(config["RiskAssessmentIndex"], negroni.HandlerFunc(riskassessmentController.IndexHandler))...))
	riskassessmentBase.Methods("POST").Handler(negroni.New(append(config["RiskAssessmentCreate"], negroni.HandlerFunc(riskassessmentController.CreateHandler))...))

	riskassessment := router.Path("/RiskAssessment/{id}").Subrouter()
	riskassessment.Methods("GET").Handler(negroni.New(append(config["RiskAssessmentShow"], negroni.HandlerFunc(riskassessmentController.ShowHandler))...))
	riskassessment.Methods("PUT").Handler(negroni.New(append(config["RiskAssessmentUpdate"], negroni.HandlerFunc(riskassessmentController.UpdateHandler))...))
	riskassessment.Methods("DELETE").Handler(negroni.New(append(config["RiskAssessmentDelete"], negroni.HandlerFunc(riskassessmentController.DeleteHandler))...))

	testscriptController := ResourceController{"TestScript"}
	testscriptBase := router.Path("/TestScript").Subrouter()
	testscriptBase.Methods("GET").Handler(negroni.New(append(config["TestScriptIndex"], negroni.HandlerFunc(testscriptController.IndexHandler))...))
	testscriptBase.Methods("POST").Handler(negroni.New(append(config["TestScriptCreate"], negroni.HandlerFunc(testscriptController.CreateHandler))...))

	testscript := router.Path("/TestScript/{id}").Subrouter()
	testscript.Methods("GET").Handler(negroni.New(append(config["TestScriptShow"], negroni.HandlerFunc(testscriptController.ShowHandler))...))
	testscript.Methods("PUT").Handler(negroni.New(append(config["TestScriptUpdate"], negroni.HandlerFunc(testscriptController.UpdateHandler))...))
	testscript.Methods("DELETE").Handler(negroni.New(append(config["TestScriptDelete"], negroni.HandlerFunc(testscriptController.DeleteHandler))...))

	basicController := ResourceController{"Basic"}
	basicBase := router.Path("/Basic").Subrouter()
	basicBase.Methods("GET").Handler(negroni.New(append(config["BasicIndex"], negroni.HandlerFunc(basicController.IndexHandler))...))
	basicBase.Methods("POST").Handler(negroni.New(append(config["BasicCreate"], negroni.HandlerFunc(basicController.CreateHandler))...))

	basic := router.Path("/Basic/{id}").Subrouter()
	basic.Methods("GET").Handler(negroni.New(append(config["BasicShow"], negroni.HandlerFunc(basicController.ShowHandler))...))
	basic.Methods("PUT").Handler(negroni.New(append(config["BasicUpdate"], negroni.HandlerFunc(basicController.UpdateHandler))...))
	basic.Methods("DELETE").Handler(negroni.New(append(config["BasicDelete"], negroni.HandlerFunc(basicController.DeleteHandler))...))

	groupController := ResourceController{"Group"}
	groupBase := router.Path("/Group").Subrouter()
	groupBase.Methods("GET").Handler(negroni.New(append(config["GroupIndex"], negroni.HandlerFunc(groupController.IndexHandler))...))
	groupBase.Methods("POST").Handler(negroni.New(append(config["GroupCreate"], negroni.HandlerFunc(groupController.CreateHandler))...))

	group := router.Path("/Group/{id}").Subrouter()
	group.Methods("GET").Handler(negroni.New(append(config["GroupShow"], negroni.HandlerFunc(groupController.ShowHandler))...))
	group.Methods("PUT").Handler(negroni.New(append(config["GroupUpdate"], negroni.HandlerFunc(groupController.UpdateHandler))...))
	group.Methods("DELETE").Handler(negroni.New(append(config["GroupDelete"], negroni.HandlerFunc(groupController.DeleteHandler))...))

	paymentnoticeController := ResourceController{"PaymentNotice"}
	paymentnoticeBase := router.Path("/PaymentNotice").Subrouter()
	paymentnoticeBase.Methods("GET").Handler(negroni.New(append(config["PaymentNoticeIndex"], negroni.HandlerFunc(paymentnoticeController.IndexHandler))...))
	paymentnoticeBase.Methods("POST").Handler(negroni.New(append(config["PaymentNoticeCreate"], negroni.HandlerFunc(paymentnoticeController.CreateHandler))...))

	paymentnotice := router.Path("/PaymentNotice/{id}").Subrouter()
	paymentnotice.Methods("GET").Handler(negroni.New(append(config["PaymentNoticeShow"], negroni.HandlerFunc(paymentnoticeController.ShowHandler))...))
	paymentnotice.Methods("PUT").Handler(negroni.New(append(config["PaymentNoticeUpdate"], negroni.HandlerFunc(paymentnoticeController.UpdateHandler))...))
	paymentnotice.Methods("DELETE").Handler(negroni.New(append(config["PaymentNoticeDelete"], negroni.HandlerFunc(paymentnoticeController.DeleteHandler))...))

	organizationController := ResourceController{"Organization"}
	organizationBase := router.Path("/Organization").Subrouter()
	organizationBase.Methods("GET").Handler(negroni.New(append(config["OrganizationIndex"], negroni.HandlerFunc(organizationController.IndexHandler))...))
	organizationBase.Methods("POST").Handler(negroni.New(append(config["OrganizationCreate"], negroni.HandlerFunc(organizationController.CreateHandler))...))

	organization := router.Path("/Organization/{id}").Subrouter()
	organization.Methods("GET").Handler(negroni.New(append(config["OrganizationShow"], negroni.HandlerFunc(organizationController.ShowHandler))...))
	organization.Methods("PUT").Handler(negroni.New(append(config["OrganizationUpdate"], negroni.HandlerFunc(organizationController.UpdateHandler))...))
	organization.Methods("DELETE").Handler(negroni.New(append(config["OrganizationDelete"], negroni.HandlerFunc(organizationController.DeleteHandler))...))

	implementationguideController := ResourceController{"ImplementationGuide"}
	implementationguideBase := router.Path("/ImplementationGuide").Subrouter()
	implementationguideBase.Methods("GET").Handler(negroni.New(append(config["ImplementationGuideIndex"], negroni.HandlerFunc(implementationguideController.IndexHandler))...))
	implementationguideBase.Methods("POST").Handler(negroni.New(append(config["ImplementationGuideCreate"], negroni.HandlerFunc(implementationguideController.CreateHandler))...))

	implementationguide := router.Path("/ImplementationGuide/{id}").Subrouter()
	implementationguide.Methods("GET").Handler(negroni.New(append(config["ImplementationGuideShow"], negroni.HandlerFunc(implementationguideController.ShowHandler))...))
	implementationguide.Methods("PUT").Handler(negroni.New(append(config["ImplementationGuideUpdate"], negroni.HandlerFunc(implementationguideController.UpdateHandler))...))
	implementationguide.Methods("DELETE").Handler(negroni.New(append(config["ImplementationGuideDelete"], negroni.HandlerFunc(implementationguideController.DeleteHandler))...))

	claimresponseController := ResourceController{"ClaimResponse"}
	claimresponseBase := router.Path("/ClaimResponse").Subrouter()
	claimresponseBase.Methods("GET").Handler(negroni.New(append(config["ClaimResponseIndex"], negroni.HandlerFunc(claimresponseController.IndexHandler))...))
	claimresponseBase.Methods("POST").Handler(negroni.New(append(config["ClaimResponseCreate"], negroni.HandlerFunc(claimresponseController.CreateHandler))...))

	claimresponse := router.Path("/ClaimResponse/{id}").Subrouter()
	claimresponse.Methods("GET").Handler(negroni.New(append(config["ClaimResponseShow"], negroni.HandlerFunc(claimresponseController.ShowHandler))...))
	claimresponse.Methods("PUT").Handler(negroni.New(append(config["ClaimResponseUpdate"], negroni.HandlerFunc(claimresponseController.UpdateHandler))...))
	claimresponse.Methods("DELETE").Handler(negroni.New(append(config["ClaimResponseDelete"], negroni.HandlerFunc(claimresponseController.DeleteHandler))...))

	eligibilityrequestController := ResourceController{"EligibilityRequest"}
	eligibilityrequestBase := router.Path("/EligibilityRequest").Subrouter()
	eligibilityrequestBase.Methods("GET").Handler(negroni.New(append(config["EligibilityRequestIndex"], negroni.HandlerFunc(eligibilityrequestController.IndexHandler))...))
	eligibilityrequestBase.Methods("POST").Handler(negroni.New(append(config["EligibilityRequestCreate"], negroni.HandlerFunc(eligibilityrequestController.CreateHandler))...))

	eligibilityrequest := router.Path("/EligibilityRequest/{id}").Subrouter()
	eligibilityrequest.Methods("GET").Handler(negroni.New(append(config["EligibilityRequestShow"], negroni.HandlerFunc(eligibilityrequestController.ShowHandler))...))
	eligibilityrequest.Methods("PUT").Handler(negroni.New(append(config["EligibilityRequestUpdate"], negroni.HandlerFunc(eligibilityrequestController.UpdateHandler))...))
	eligibilityrequest.Methods("DELETE").Handler(negroni.New(append(config["EligibilityRequestDelete"], negroni.HandlerFunc(eligibilityrequestController.DeleteHandler))...))

	processrequestController := ResourceController{"ProcessRequest"}
	processrequestBase := router.Path("/ProcessRequest").Subrouter()
	processrequestBase.Methods("GET").Handler(negroni.New(append(config["ProcessRequestIndex"], negroni.HandlerFunc(processrequestController.IndexHandler))...))
	processrequestBase.Methods("POST").Handler(negroni.New(append(config["ProcessRequestCreate"], negroni.HandlerFunc(processrequestController.CreateHandler))...))

	processrequest := router.Path("/ProcessRequest/{id}").Subrouter()
	processrequest.Methods("GET").Handler(negroni.New(append(config["ProcessRequestShow"], negroni.HandlerFunc(processrequestController.ShowHandler))...))
	processrequest.Methods("PUT").Handler(negroni.New(append(config["ProcessRequestUpdate"], negroni.HandlerFunc(processrequestController.UpdateHandler))...))
	processrequest.Methods("DELETE").Handler(negroni.New(append(config["ProcessRequestDelete"], negroni.HandlerFunc(processrequestController.DeleteHandler))...))

	medicationdispenseController := ResourceController{"MedicationDispense"}
	medicationdispenseBase := router.Path("/MedicationDispense").Subrouter()
	medicationdispenseBase.Methods("GET").Handler(negroni.New(append(config["MedicationDispenseIndex"], negroni.HandlerFunc(medicationdispenseController.IndexHandler))...))
	medicationdispenseBase.Methods("POST").Handler(negroni.New(append(config["MedicationDispenseCreate"], negroni.HandlerFunc(medicationdispenseController.CreateHandler))...))

	medicationdispense := router.Path("/MedicationDispense/{id}").Subrouter()
	medicationdispense.Methods("GET").Handler(negroni.New(append(config["MedicationDispenseShow"], negroni.HandlerFunc(medicationdispenseController.ShowHandler))...))
	medicationdispense.Methods("PUT").Handler(negroni.New(append(config["MedicationDispenseUpdate"], negroni.HandlerFunc(medicationdispenseController.UpdateHandler))...))
	medicationdispense.Methods("DELETE").Handler(negroni.New(append(config["MedicationDispenseDelete"], negroni.HandlerFunc(medicationdispenseController.DeleteHandler))...))

	diagnosticreportController := ResourceController{"DiagnosticReport"}
	diagnosticreportBase := router.Path("/DiagnosticReport").Subrouter()
	diagnosticreportBase.Methods("GET").Handler(negroni.New(append(config["DiagnosticReportIndex"], negroni.HandlerFunc(diagnosticreportController.IndexHandler))...))
	diagnosticreportBase.Methods("POST").Handler(negroni.New(append(config["DiagnosticReportCreate"], negroni.HandlerFunc(diagnosticreportController.CreateHandler))...))

	diagnosticreport := router.Path("/DiagnosticReport/{id}").Subrouter()
	diagnosticreport.Methods("GET").Handler(negroni.New(append(config["DiagnosticReportShow"], negroni.HandlerFunc(diagnosticreportController.ShowHandler))...))
	diagnosticreport.Methods("PUT").Handler(negroni.New(append(config["DiagnosticReportUpdate"], negroni.HandlerFunc(diagnosticreportController.UpdateHandler))...))
	diagnosticreport.Methods("DELETE").Handler(negroni.New(append(config["DiagnosticReportDelete"], negroni.HandlerFunc(diagnosticreportController.DeleteHandler))...))

	imagingstudyController := ResourceController{"ImagingStudy"}
	imagingstudyBase := router.Path("/ImagingStudy").Subrouter()
	imagingstudyBase.Methods("GET").Handler(negroni.New(append(config["ImagingStudyIndex"], negroni.HandlerFunc(imagingstudyController.IndexHandler))...))
	imagingstudyBase.Methods("POST").Handler(negroni.New(append(config["ImagingStudyCreate"], negroni.HandlerFunc(imagingstudyController.CreateHandler))...))

	imagingstudy := router.Path("/ImagingStudy/{id}").Subrouter()
	imagingstudy.Methods("GET").Handler(negroni.New(append(config["ImagingStudyShow"], negroni.HandlerFunc(imagingstudyController.ShowHandler))...))
	imagingstudy.Methods("PUT").Handler(negroni.New(append(config["ImagingStudyUpdate"], negroni.HandlerFunc(imagingstudyController.UpdateHandler))...))
	imagingstudy.Methods("DELETE").Handler(negroni.New(append(config["ImagingStudyDelete"], negroni.HandlerFunc(imagingstudyController.DeleteHandler))...))

	imagingobjectselectionController := ResourceController{"ImagingObjectSelection"}
	imagingobjectselectionBase := router.Path("/ImagingObjectSelection").Subrouter()
	imagingobjectselectionBase.Methods("GET").Handler(negroni.New(append(config["ImagingObjectSelectionIndex"], negroni.HandlerFunc(imagingobjectselectionController.IndexHandler))...))
	imagingobjectselectionBase.Methods("POST").Handler(negroni.New(append(config["ImagingObjectSelectionCreate"], negroni.HandlerFunc(imagingobjectselectionController.CreateHandler))...))

	imagingobjectselection := router.Path("/ImagingObjectSelection/{id}").Subrouter()
	imagingobjectselection.Methods("GET").Handler(negroni.New(append(config["ImagingObjectSelectionShow"], negroni.HandlerFunc(imagingobjectselectionController.ShowHandler))...))
	imagingobjectselection.Methods("PUT").Handler(negroni.New(append(config["ImagingObjectSelectionUpdate"], negroni.HandlerFunc(imagingobjectselectionController.UpdateHandler))...))
	imagingobjectselection.Methods("DELETE").Handler(negroni.New(append(config["ImagingObjectSelectionDelete"], negroni.HandlerFunc(imagingobjectselectionController.DeleteHandler))...))

	healthcareserviceController := ResourceController{"HealthcareService"}
	healthcareserviceBase := router.Path("/HealthcareService").Subrouter()
	healthcareserviceBase.Methods("GET").Handler(negroni.New(append(config["HealthcareServiceIndex"], negroni.HandlerFunc(healthcareserviceController.IndexHandler))...))
	healthcareserviceBase.Methods("POST").Handler(negroni.New(append(config["HealthcareServiceCreate"], negroni.HandlerFunc(healthcareserviceController.CreateHandler))...))

	healthcareservice := router.Path("/HealthcareService/{id}").Subrouter()
	healthcareservice.Methods("GET").Handler(negroni.New(append(config["HealthcareServiceShow"], negroni.HandlerFunc(healthcareserviceController.ShowHandler))...))
	healthcareservice.Methods("PUT").Handler(negroni.New(append(config["HealthcareServiceUpdate"], negroni.HandlerFunc(healthcareserviceController.UpdateHandler))...))
	healthcareservice.Methods("DELETE").Handler(negroni.New(append(config["HealthcareServiceDelete"], negroni.HandlerFunc(healthcareserviceController.DeleteHandler))...))

	dataelementController := ResourceController{"DataElement"}
	dataelementBase := router.Path("/DataElement").Subrouter()
	dataelementBase.Methods("GET").Handler(negroni.New(append(config["DataElementIndex"], negroni.HandlerFunc(dataelementController.IndexHandler))...))
	dataelementBase.Methods("POST").Handler(negroni.New(append(config["DataElementCreate"], negroni.HandlerFunc(dataelementController.CreateHandler))...))

	dataelement := router.Path("/DataElement/{id}").Subrouter()
	dataelement.Methods("GET").Handler(negroni.New(append(config["DataElementShow"], negroni.HandlerFunc(dataelementController.ShowHandler))...))
	dataelement.Methods("PUT").Handler(negroni.New(append(config["DataElementUpdate"], negroni.HandlerFunc(dataelementController.UpdateHandler))...))
	dataelement.Methods("DELETE").Handler(negroni.New(append(config["DataElementDelete"], negroni.HandlerFunc(dataelementController.DeleteHandler))...))

	devicecomponentController := ResourceController{"DeviceComponent"}
	devicecomponentBase := router.Path("/DeviceComponent").Subrouter()
	devicecomponentBase.Methods("GET").Handler(negroni.New(append(config["DeviceComponentIndex"], negroni.HandlerFunc(devicecomponentController.IndexHandler))...))
	devicecomponentBase.Methods("POST").Handler(negroni.New(append(config["DeviceComponentCreate"], negroni.HandlerFunc(devicecomponentController.CreateHandler))...))

	devicecomponent := router.Path("/DeviceComponent/{id}").Subrouter()
	devicecomponent.Methods("GET").Handler(negroni.New(append(config["DeviceComponentShow"], negroni.HandlerFunc(devicecomponentController.ShowHandler))...))
	devicecomponent.Methods("PUT").Handler(negroni.New(append(config["DeviceComponentUpdate"], negroni.HandlerFunc(devicecomponentController.UpdateHandler))...))
	devicecomponent.Methods("DELETE").Handler(negroni.New(append(config["DeviceComponentDelete"], negroni.HandlerFunc(devicecomponentController.DeleteHandler))...))

	familymemberhistoryController := ResourceController{"FamilyMemberHistory"}
	familymemberhistoryBase := router.Path("/FamilyMemberHistory").Subrouter()
	familymemberhistoryBase.Methods("GET").Handler(negroni.New(append(config["FamilyMemberHistoryIndex"], negroni.HandlerFunc(familymemberhistoryController.IndexHandler))...))
	familymemberhistoryBase.Methods("POST").Handler(negroni.New(append(config["FamilyMemberHistoryCreate"], negroni.HandlerFunc(familymemberhistoryController.CreateHandler))...))

	familymemberhistory := router.Path("/FamilyMemberHistory/{id}").Subrouter()
	familymemberhistory.Methods("GET").Handler(negroni.New(append(config["FamilyMemberHistoryShow"], negroni.HandlerFunc(familymemberhistoryController.ShowHandler))...))
	familymemberhistory.Methods("PUT").Handler(negroni.New(append(config["FamilyMemberHistoryUpdate"], negroni.HandlerFunc(familymemberhistoryController.UpdateHandler))...))
	familymemberhistory.Methods("DELETE").Handler(negroni.New(append(config["FamilyMemberHistoryDelete"], negroni.HandlerFunc(familymemberhistoryController.DeleteHandler))...))

	nutritionorderController := ResourceController{"NutritionOrder"}
	nutritionorderBase := router.Path("/NutritionOrder").Subrouter()
	nutritionorderBase.Methods("GET").Handler(negroni.New(append(config["NutritionOrderIndex"], negroni.HandlerFunc(nutritionorderController.IndexHandler))...))
	nutritionorderBase.Methods("POST").Handler(negroni.New(append(config["NutritionOrderCreate"], negroni.HandlerFunc(nutritionorderController.CreateHandler))...))

	nutritionorder := router.Path("/NutritionOrder/{id}").Subrouter()
	nutritionorder.Methods("GET").Handler(negroni.New(append(config["NutritionOrderShow"], negroni.HandlerFunc(nutritionorderController.ShowHandler))...))
	nutritionorder.Methods("PUT").Handler(negroni.New(append(config["NutritionOrderUpdate"], negroni.HandlerFunc(nutritionorderController.UpdateHandler))...))
	nutritionorder.Methods("DELETE").Handler(negroni.New(append(config["NutritionOrderDelete"], negroni.HandlerFunc(nutritionorderController.DeleteHandler))...))

	encounterController := ResourceController{"Encounter"}
	encounterBase := router.Path("/Encounter").Subrouter()
	encounterBase.Methods("GET").Handler(negroni.New(append(config["EncounterIndex"], negroni.HandlerFunc(encounterController.IndexHandler))...))
	encounterBase.Methods("POST").Handler(negroni.New(append(config["EncounterCreate"], negroni.HandlerFunc(encounterController.CreateHandler))...))

	encounter := router.Path("/Encounter/{id}").Subrouter()
	encounter.Methods("GET").Handler(negroni.New(append(config["EncounterShow"], negroni.HandlerFunc(encounterController.ShowHandler))...))
	encounter.Methods("PUT").Handler(negroni.New(append(config["EncounterUpdate"], negroni.HandlerFunc(encounterController.UpdateHandler))...))
	encounter.Methods("DELETE").Handler(negroni.New(append(config["EncounterDelete"], negroni.HandlerFunc(encounterController.DeleteHandler))...))

	substanceController := ResourceController{"Substance"}
	substanceBase := router.Path("/Substance").Subrouter()
	substanceBase.Methods("GET").Handler(negroni.New(append(config["SubstanceIndex"], negroni.HandlerFunc(substanceController.IndexHandler))...))
	substanceBase.Methods("POST").Handler(negroni.New(append(config["SubstanceCreate"], negroni.HandlerFunc(substanceController.CreateHandler))...))

	substance := router.Path("/Substance/{id}").Subrouter()
	substance.Methods("GET").Handler(negroni.New(append(config["SubstanceShow"], negroni.HandlerFunc(substanceController.ShowHandler))...))
	substance.Methods("PUT").Handler(negroni.New(append(config["SubstanceUpdate"], negroni.HandlerFunc(substanceController.UpdateHandler))...))
	substance.Methods("DELETE").Handler(negroni.New(append(config["SubstanceDelete"], negroni.HandlerFunc(substanceController.DeleteHandler))...))

	auditeventController := ResourceController{"AuditEvent"}
	auditeventBase := router.Path("/AuditEvent").Subrouter()
	auditeventBase.Methods("GET").Handler(negroni.New(append(config["AuditEventIndex"], negroni.HandlerFunc(auditeventController.IndexHandler))...))
	auditeventBase.Methods("POST").Handler(negroni.New(append(config["AuditEventCreate"], negroni.HandlerFunc(auditeventController.CreateHandler))...))

	auditevent := router.Path("/AuditEvent/{id}").Subrouter()
	auditevent.Methods("GET").Handler(negroni.New(append(config["AuditEventShow"], negroni.HandlerFunc(auditeventController.ShowHandler))...))
	auditevent.Methods("PUT").Handler(negroni.New(append(config["AuditEventUpdate"], negroni.HandlerFunc(auditeventController.UpdateHandler))...))
	auditevent.Methods("DELETE").Handler(negroni.New(append(config["AuditEventDelete"], negroni.HandlerFunc(auditeventController.DeleteHandler))...))

	medicationorderController := ResourceController{"MedicationOrder"}
	medicationorderBase := router.Path("/MedicationOrder").Subrouter()
	medicationorderBase.Methods("GET").Handler(negroni.New(append(config["MedicationOrderIndex"], negroni.HandlerFunc(medicationorderController.IndexHandler))...))
	medicationorderBase.Methods("POST").Handler(negroni.New(append(config["MedicationOrderCreate"], negroni.HandlerFunc(medicationorderController.CreateHandler))...))

	medicationorder := router.Path("/MedicationOrder/{id}").Subrouter()
	medicationorder.Methods("GET").Handler(negroni.New(append(config["MedicationOrderShow"], negroni.HandlerFunc(medicationorderController.ShowHandler))...))
	medicationorder.Methods("PUT").Handler(negroni.New(append(config["MedicationOrderUpdate"], negroni.HandlerFunc(medicationorderController.UpdateHandler))...))
	medicationorder.Methods("DELETE").Handler(negroni.New(append(config["MedicationOrderDelete"], negroni.HandlerFunc(medicationorderController.DeleteHandler))...))

	searchparameterController := ResourceController{"SearchParameter"}
	searchparameterBase := router.Path("/SearchParameter").Subrouter()
	searchparameterBase.Methods("GET").Handler(negroni.New(append(config["SearchParameterIndex"], negroni.HandlerFunc(searchparameterController.IndexHandler))...))
	searchparameterBase.Methods("POST").Handler(negroni.New(append(config["SearchParameterCreate"], negroni.HandlerFunc(searchparameterController.CreateHandler))...))

	searchparameter := router.Path("/SearchParameter/{id}").Subrouter()
	searchparameter.Methods("GET").Handler(negroni.New(append(config["SearchParameterShow"], negroni.HandlerFunc(searchparameterController.ShowHandler))...))
	searchparameter.Methods("PUT").Handler(negroni.New(append(config["SearchParameterUpdate"], negroni.HandlerFunc(searchparameterController.UpdateHandler))...))
	searchparameter.Methods("DELETE").Handler(negroni.New(append(config["SearchParameterDelete"], negroni.HandlerFunc(searchparameterController.DeleteHandler))...))

	paymentreconciliationController := ResourceController{"PaymentReconciliation"}
	paymentreconciliationBase := router.Path("/PaymentReconciliation").Subrouter()
	paymentreconciliationBase.Methods("GET").Handler(negroni.New(append(config["PaymentReconciliationIndex"], negroni.HandlerFunc(paymentreconciliationController.IndexHandler))...))
	paymentreconciliationBase.Methods("POST").Handler(negroni.New(append(config["PaymentReconciliationCreate"], negroni.HandlerFunc(paymentreconciliationController.CreateHandler))...))

	paymentreconciliation := router.Path("/PaymentReconciliation/{id}").Subrouter()
	paymentreconciliation.Methods("GET").Handler(negroni.New(append(config["PaymentReconciliationShow"], negroni.HandlerFunc(paymentreconciliationController.ShowHandler))...))
	paymentreconciliation.Methods("PUT").Handler(negroni.New(append(config["PaymentReconciliationUpdate"], negroni.HandlerFunc(paymentreconciliationController.UpdateHandler))...))
	paymentreconciliation.Methods("DELETE").Handler(negroni.New(append(config["PaymentReconciliationDelete"], negroni.HandlerFunc(paymentreconciliationController.DeleteHandler))...))

	communicationController := ResourceController{"Communication"}
	communicationBase := router.Path("/Communication").Subrouter()
	communicationBase.Methods("GET").Handler(negroni.New(append(config["CommunicationIndex"], negroni.HandlerFunc(communicationController.IndexHandler))...))
	communicationBase.Methods("POST").Handler(negroni.New(append(config["CommunicationCreate"], negroni.HandlerFunc(communicationController.CreateHandler))...))

	communication := router.Path("/Communication/{id}").Subrouter()
	communication.Methods("GET").Handler(negroni.New(append(config["CommunicationShow"], negroni.HandlerFunc(communicationController.ShowHandler))...))
	communication.Methods("PUT").Handler(negroni.New(append(config["CommunicationUpdate"], negroni.HandlerFunc(communicationController.UpdateHandler))...))
	communication.Methods("DELETE").Handler(negroni.New(append(config["CommunicationDelete"], negroni.HandlerFunc(communicationController.DeleteHandler))...))

	conditionController := ResourceController{"Condition"}
	conditionBase := router.Path("/Condition").Subrouter()
	conditionBase.Methods("GET").Handler(negroni.New(append(config["ConditionIndex"], negroni.HandlerFunc(conditionController.IndexHandler))...))
	conditionBase.Methods("POST").Handler(negroni.New(append(config["ConditionCreate"], negroni.HandlerFunc(conditionController.CreateHandler))...))

	condition := router.Path("/Condition/{id}").Subrouter()
	condition.Methods("GET").Handler(negroni.New(append(config["ConditionShow"], negroni.HandlerFunc(conditionController.ShowHandler))...))
	condition.Methods("PUT").Handler(negroni.New(append(config["ConditionUpdate"], negroni.HandlerFunc(conditionController.UpdateHandler))...))
	condition.Methods("DELETE").Handler(negroni.New(append(config["ConditionDelete"], negroni.HandlerFunc(conditionController.DeleteHandler))...))

	compositionController := ResourceController{"Composition"}
	compositionBase := router.Path("/Composition").Subrouter()
	compositionBase.Methods("GET").Handler(negroni.New(append(config["CompositionIndex"], negroni.HandlerFunc(compositionController.IndexHandler))...))
	compositionBase.Methods("POST").Handler(negroni.New(append(config["CompositionCreate"], negroni.HandlerFunc(compositionController.CreateHandler))...))

	composition := router.Path("/Composition/{id}").Subrouter()
	composition.Methods("GET").Handler(negroni.New(append(config["CompositionShow"], negroni.HandlerFunc(compositionController.ShowHandler))...))
	composition.Methods("PUT").Handler(negroni.New(append(config["CompositionUpdate"], negroni.HandlerFunc(compositionController.UpdateHandler))...))
	composition.Methods("DELETE").Handler(negroni.New(append(config["CompositionDelete"], negroni.HandlerFunc(compositionController.DeleteHandler))...))

	detectedissueController := ResourceController{"DetectedIssue"}
	detectedissueBase := router.Path("/DetectedIssue").Subrouter()
	detectedissueBase.Methods("GET").Handler(negroni.New(append(config["DetectedIssueIndex"], negroni.HandlerFunc(detectedissueController.IndexHandler))...))
	detectedissueBase.Methods("POST").Handler(negroni.New(append(config["DetectedIssueCreate"], negroni.HandlerFunc(detectedissueController.CreateHandler))...))

	detectedissue := router.Path("/DetectedIssue/{id}").Subrouter()
	detectedissue.Methods("GET").Handler(negroni.New(append(config["DetectedIssueShow"], negroni.HandlerFunc(detectedissueController.ShowHandler))...))
	detectedissue.Methods("PUT").Handler(negroni.New(append(config["DetectedIssueUpdate"], negroni.HandlerFunc(detectedissueController.UpdateHandler))...))
	detectedissue.Methods("DELETE").Handler(negroni.New(append(config["DetectedIssueDelete"], negroni.HandlerFunc(detectedissueController.DeleteHandler))...))

	bundleController := ResourceController{"Bundle"}
	bundleBase := router.Path("/Bundle").Subrouter()
	bundleBase.Methods("GET").Handler(negroni.New(append(config["BundleIndex"], negroni.HandlerFunc(bundleController.IndexHandler))...))
	bundleBase.Methods("POST").Handler(negroni.New(append(config["BundleCreate"], negroni.HandlerFunc(bundleController.CreateHandler))...))

	bundle := router.Path("/Bundle/{id}").Subrouter()
	bundle.Methods("GET").Handler(negroni.New(append(config["BundleShow"], negroni.HandlerFunc(bundleController.ShowHandler))...))
	bundle.Methods("PUT").Handler(negroni.New(append(config["BundleUpdate"], negroni.HandlerFunc(bundleController.UpdateHandler))...))
	bundle.Methods("DELETE").Handler(negroni.New(append(config["BundleDelete"], negroni.HandlerFunc(bundleController.DeleteHandler))...))

	diagnosticorderController := ResourceController{"DiagnosticOrder"}
	diagnosticorderBase := router.Path("/DiagnosticOrder").Subrouter()
	diagnosticorderBase.Methods("GET").Handler(negroni.New(append(config["DiagnosticOrderIndex"], negroni.HandlerFunc(diagnosticorderController.IndexHandler))...))
	diagnosticorderBase.Methods("POST").Handler(negroni.New(append(config["DiagnosticOrderCreate"], negroni.HandlerFunc(diagnosticorderController.CreateHandler))...))

	diagnosticorder := router.Path("/DiagnosticOrder/{id}").Subrouter()
	diagnosticorder.Methods("GET").Handler(negroni.New(append(config["DiagnosticOrderShow"], negroni.HandlerFunc(diagnosticorderController.ShowHandler))...))
	diagnosticorder.Methods("PUT").Handler(negroni.New(append(config["DiagnosticOrderUpdate"], negroni.HandlerFunc(diagnosticorderController.UpdateHandler))...))
	diagnosticorder.Methods("DELETE").Handler(negroni.New(append(config["DiagnosticOrderDelete"], negroni.HandlerFunc(diagnosticorderController.DeleteHandler))...))

	patientController := ResourceController{"Patient"}
	patientBase := router.Path("/Patient").Subrouter()
	patientBase.Methods("GET").Handler(negroni.New(append(config["PatientIndex"], negroni.HandlerFunc(patientController.IndexHandler))...))
	patientBase.Methods("POST").Handler(negroni.New(append(config["PatientCreate"], negroni.HandlerFunc(patientController.CreateHandler))...))

	patient := router.Path("/Patient/{id}").Subrouter()
	patient.Methods("GET").Handler(negroni.New(append(config["PatientShow"], negroni.HandlerFunc(patientController.ShowHandler))...))
	patient.Methods("PUT").Handler(negroni.New(append(config["PatientUpdate"], negroni.HandlerFunc(patientController.UpdateHandler))...))
	patient.Methods("DELETE").Handler(negroni.New(append(config["PatientDelete"], negroni.HandlerFunc(patientController.DeleteHandler))...))

	orderresponseController := ResourceController{"OrderResponse"}
	orderresponseBase := router.Path("/OrderResponse").Subrouter()
	orderresponseBase.Methods("GET").Handler(negroni.New(append(config["OrderResponseIndex"], negroni.HandlerFunc(orderresponseController.IndexHandler))...))
	orderresponseBase.Methods("POST").Handler(negroni.New(append(config["OrderResponseCreate"], negroni.HandlerFunc(orderresponseController.CreateHandler))...))

	orderresponse := router.Path("/OrderResponse/{id}").Subrouter()
	orderresponse.Methods("GET").Handler(negroni.New(append(config["OrderResponseShow"], negroni.HandlerFunc(orderresponseController.ShowHandler))...))
	orderresponse.Methods("PUT").Handler(negroni.New(append(config["OrderResponseUpdate"], negroni.HandlerFunc(orderresponseController.UpdateHandler))...))
	orderresponse.Methods("DELETE").Handler(negroni.New(append(config["OrderResponseDelete"], negroni.HandlerFunc(orderresponseController.DeleteHandler))...))

	coverageController := ResourceController{"Coverage"}
	coverageBase := router.Path("/Coverage").Subrouter()
	coverageBase.Methods("GET").Handler(negroni.New(append(config["CoverageIndex"], negroni.HandlerFunc(coverageController.IndexHandler))...))
	coverageBase.Methods("POST").Handler(negroni.New(append(config["CoverageCreate"], negroni.HandlerFunc(coverageController.CreateHandler))...))

	coverage := router.Path("/Coverage/{id}").Subrouter()
	coverage.Methods("GET").Handler(negroni.New(append(config["CoverageShow"], negroni.HandlerFunc(coverageController.ShowHandler))...))
	coverage.Methods("PUT").Handler(negroni.New(append(config["CoverageUpdate"], negroni.HandlerFunc(coverageController.UpdateHandler))...))
	coverage.Methods("DELETE").Handler(negroni.New(append(config["CoverageDelete"], negroni.HandlerFunc(coverageController.DeleteHandler))...))

	questionnaireresponseController := ResourceController{"QuestionnaireResponse"}
	questionnaireresponseBase := router.Path("/QuestionnaireResponse").Subrouter()
	questionnaireresponseBase.Methods("GET").Handler(negroni.New(append(config["QuestionnaireResponseIndex"], negroni.HandlerFunc(questionnaireresponseController.IndexHandler))...))
	questionnaireresponseBase.Methods("POST").Handler(negroni.New(append(config["QuestionnaireResponseCreate"], negroni.HandlerFunc(questionnaireresponseController.CreateHandler))...))

	questionnaireresponse := router.Path("/QuestionnaireResponse/{id}").Subrouter()
	questionnaireresponse.Methods("GET").Handler(negroni.New(append(config["QuestionnaireResponseShow"], negroni.HandlerFunc(questionnaireresponseController.ShowHandler))...))
	questionnaireresponse.Methods("PUT").Handler(negroni.New(append(config["QuestionnaireResponseUpdate"], negroni.HandlerFunc(questionnaireresponseController.UpdateHandler))...))
	questionnaireresponse.Methods("DELETE").Handler(negroni.New(append(config["QuestionnaireResponseDelete"], negroni.HandlerFunc(questionnaireresponseController.DeleteHandler))...))

	deviceusestatementController := ResourceController{"DeviceUseStatement"}
	deviceusestatementBase := router.Path("/DeviceUseStatement").Subrouter()
	deviceusestatementBase.Methods("GET").Handler(negroni.New(append(config["DeviceUseStatementIndex"], negroni.HandlerFunc(deviceusestatementController.IndexHandler))...))
	deviceusestatementBase.Methods("POST").Handler(negroni.New(append(config["DeviceUseStatementCreate"], negroni.HandlerFunc(deviceusestatementController.CreateHandler))...))

	deviceusestatement := router.Path("/DeviceUseStatement/{id}").Subrouter()
	deviceusestatement.Methods("GET").Handler(negroni.New(append(config["DeviceUseStatementShow"], negroni.HandlerFunc(deviceusestatementController.ShowHandler))...))
	deviceusestatement.Methods("PUT").Handler(negroni.New(append(config["DeviceUseStatementUpdate"], negroni.HandlerFunc(deviceusestatementController.UpdateHandler))...))
	deviceusestatement.Methods("DELETE").Handler(negroni.New(append(config["DeviceUseStatementDelete"], negroni.HandlerFunc(deviceusestatementController.DeleteHandler))...))

	processresponseController := ResourceController{"ProcessResponse"}
	processresponseBase := router.Path("/ProcessResponse").Subrouter()
	processresponseBase.Methods("GET").Handler(negroni.New(append(config["ProcessResponseIndex"], negroni.HandlerFunc(processresponseController.IndexHandler))...))
	processresponseBase.Methods("POST").Handler(negroni.New(append(config["ProcessResponseCreate"], negroni.HandlerFunc(processresponseController.CreateHandler))...))

	processresponse := router.Path("/ProcessResponse/{id}").Subrouter()
	processresponse.Methods("GET").Handler(negroni.New(append(config["ProcessResponseShow"], negroni.HandlerFunc(processresponseController.ShowHandler))...))
	processresponse.Methods("PUT").Handler(negroni.New(append(config["ProcessResponseUpdate"], negroni.HandlerFunc(processresponseController.UpdateHandler))...))
	processresponse.Methods("DELETE").Handler(negroni.New(append(config["ProcessResponseDelete"], negroni.HandlerFunc(processresponseController.DeleteHandler))...))

	namingsystemController := ResourceController{"NamingSystem"}
	namingsystemBase := router.Path("/NamingSystem").Subrouter()
	namingsystemBase.Methods("GET").Handler(negroni.New(append(config["NamingSystemIndex"], negroni.HandlerFunc(namingsystemController.IndexHandler))...))
	namingsystemBase.Methods("POST").Handler(negroni.New(append(config["NamingSystemCreate"], negroni.HandlerFunc(namingsystemController.CreateHandler))...))

	namingsystem := router.Path("/NamingSystem/{id}").Subrouter()
	namingsystem.Methods("GET").Handler(negroni.New(append(config["NamingSystemShow"], negroni.HandlerFunc(namingsystemController.ShowHandler))...))
	namingsystem.Methods("PUT").Handler(negroni.New(append(config["NamingSystemUpdate"], negroni.HandlerFunc(namingsystemController.UpdateHandler))...))
	namingsystem.Methods("DELETE").Handler(negroni.New(append(config["NamingSystemDelete"], negroni.HandlerFunc(namingsystemController.DeleteHandler))...))

	scheduleController := ResourceController{"Schedule"}
	scheduleBase := router.Path("/Schedule").Subrouter()
	scheduleBase.Methods("GET").Handler(negroni.New(append(config["ScheduleIndex"], negroni.HandlerFunc(scheduleController.IndexHandler))...))
	scheduleBase.Methods("POST").Handler(negroni.New(append(config["ScheduleCreate"], negroni.HandlerFunc(scheduleController.CreateHandler))...))

	schedule := router.Path("/Schedule/{id}").Subrouter()
	schedule.Methods("GET").Handler(negroni.New(append(config["ScheduleShow"], negroni.HandlerFunc(scheduleController.ShowHandler))...))
	schedule.Methods("PUT").Handler(negroni.New(append(config["ScheduleUpdate"], negroni.HandlerFunc(scheduleController.UpdateHandler))...))
	schedule.Methods("DELETE").Handler(negroni.New(append(config["ScheduleDelete"], negroni.HandlerFunc(scheduleController.DeleteHandler))...))

	supplydeliveryController := ResourceController{"SupplyDelivery"}
	supplydeliveryBase := router.Path("/SupplyDelivery").Subrouter()
	supplydeliveryBase.Methods("GET").Handler(negroni.New(append(config["SupplyDeliveryIndex"], negroni.HandlerFunc(supplydeliveryController.IndexHandler))...))
	supplydeliveryBase.Methods("POST").Handler(negroni.New(append(config["SupplyDeliveryCreate"], negroni.HandlerFunc(supplydeliveryController.CreateHandler))...))

	supplydelivery := router.Path("/SupplyDelivery/{id}").Subrouter()
	supplydelivery.Methods("GET").Handler(negroni.New(append(config["SupplyDeliveryShow"], negroni.HandlerFunc(supplydeliveryController.ShowHandler))...))
	supplydelivery.Methods("PUT").Handler(negroni.New(append(config["SupplyDeliveryUpdate"], negroni.HandlerFunc(supplydeliveryController.UpdateHandler))...))
	supplydelivery.Methods("DELETE").Handler(negroni.New(append(config["SupplyDeliveryDelete"], negroni.HandlerFunc(supplydeliveryController.DeleteHandler))...))

	clinicalimpressionController := ResourceController{"ClinicalImpression"}
	clinicalimpressionBase := router.Path("/ClinicalImpression").Subrouter()
	clinicalimpressionBase.Methods("GET").Handler(negroni.New(append(config["ClinicalImpressionIndex"], negroni.HandlerFunc(clinicalimpressionController.IndexHandler))...))
	clinicalimpressionBase.Methods("POST").Handler(negroni.New(append(config["ClinicalImpressionCreate"], negroni.HandlerFunc(clinicalimpressionController.CreateHandler))...))

	clinicalimpression := router.Path("/ClinicalImpression/{id}").Subrouter()
	clinicalimpression.Methods("GET").Handler(negroni.New(append(config["ClinicalImpressionShow"], negroni.HandlerFunc(clinicalimpressionController.ShowHandler))...))
	clinicalimpression.Methods("PUT").Handler(negroni.New(append(config["ClinicalImpressionUpdate"], negroni.HandlerFunc(clinicalimpressionController.UpdateHandler))...))
	clinicalimpression.Methods("DELETE").Handler(negroni.New(append(config["ClinicalImpressionDelete"], negroni.HandlerFunc(clinicalimpressionController.DeleteHandler))...))

	messageheaderController := ResourceController{"MessageHeader"}
	messageheaderBase := router.Path("/MessageHeader").Subrouter()
	messageheaderBase.Methods("GET").Handler(negroni.New(append(config["MessageHeaderIndex"], negroni.HandlerFunc(messageheaderController.IndexHandler))...))
	messageheaderBase.Methods("POST").Handler(negroni.New(append(config["MessageHeaderCreate"], negroni.HandlerFunc(messageheaderController.CreateHandler))...))

	messageheader := router.Path("/MessageHeader/{id}").Subrouter()
	messageheader.Methods("GET").Handler(negroni.New(append(config["MessageHeaderShow"], negroni.HandlerFunc(messageheaderController.ShowHandler))...))
	messageheader.Methods("PUT").Handler(negroni.New(append(config["MessageHeaderUpdate"], negroni.HandlerFunc(messageheaderController.UpdateHandler))...))
	messageheader.Methods("DELETE").Handler(negroni.New(append(config["MessageHeaderDelete"], negroni.HandlerFunc(messageheaderController.DeleteHandler))...))

	claimController := ResourceController{"Claim"}
	claimBase := router.Path("/Claim").Subrouter()
	claimBase.Methods("GET").Handler(negroni.New(append(config["ClaimIndex"], negroni.HandlerFunc(claimController.IndexHandler))...))
	claimBase.Methods("POST").Handler(negroni.New(append(config["ClaimCreate"], negroni.HandlerFunc(claimController.CreateHandler))...))

	claim := router.Path("/Claim/{id}").Subrouter()
	claim.Methods("GET").Handler(negroni.New(append(config["ClaimShow"], negroni.HandlerFunc(claimController.ShowHandler))...))
	claim.Methods("PUT").Handler(negroni.New(append(config["ClaimUpdate"], negroni.HandlerFunc(claimController.UpdateHandler))...))
	claim.Methods("DELETE").Handler(negroni.New(append(config["ClaimDelete"], negroni.HandlerFunc(claimController.DeleteHandler))...))

	immunizationrecommendationController := ResourceController{"ImmunizationRecommendation"}
	immunizationrecommendationBase := router.Path("/ImmunizationRecommendation").Subrouter()
	immunizationrecommendationBase.Methods("GET").Handler(negroni.New(append(config["ImmunizationRecommendationIndex"], negroni.HandlerFunc(immunizationrecommendationController.IndexHandler))...))
	immunizationrecommendationBase.Methods("POST").Handler(negroni.New(append(config["ImmunizationRecommendationCreate"], negroni.HandlerFunc(immunizationrecommendationController.CreateHandler))...))

	immunizationrecommendation := router.Path("/ImmunizationRecommendation/{id}").Subrouter()
	immunizationrecommendation.Methods("GET").Handler(negroni.New(append(config["ImmunizationRecommendationShow"], negroni.HandlerFunc(immunizationrecommendationController.ShowHandler))...))
	immunizationrecommendation.Methods("PUT").Handler(negroni.New(append(config["ImmunizationRecommendationUpdate"], negroni.HandlerFunc(immunizationrecommendationController.UpdateHandler))...))
	immunizationrecommendation.Methods("DELETE").Handler(negroni.New(append(config["ImmunizationRecommendationDelete"], negroni.HandlerFunc(immunizationrecommendationController.DeleteHandler))...))

	locationController := ResourceController{"Location"}
	locationBase := router.Path("/Location").Subrouter()
	locationBase.Methods("GET").Handler(negroni.New(append(config["LocationIndex"], negroni.HandlerFunc(locationController.IndexHandler))...))
	locationBase.Methods("POST").Handler(negroni.New(append(config["LocationCreate"], negroni.HandlerFunc(locationController.CreateHandler))...))

	location := router.Path("/Location/{id}").Subrouter()
	location.Methods("GET").Handler(negroni.New(append(config["LocationShow"], negroni.HandlerFunc(locationController.ShowHandler))...))
	location.Methods("PUT").Handler(negroni.New(append(config["LocationUpdate"], negroni.HandlerFunc(locationController.UpdateHandler))...))
	location.Methods("DELETE").Handler(negroni.New(append(config["LocationDelete"], negroni.HandlerFunc(locationController.DeleteHandler))...))

	bodysiteController := ResourceController{"BodySite"}
	bodysiteBase := router.Path("/BodySite").Subrouter()
	bodysiteBase.Methods("GET").Handler(negroni.New(append(config["BodySiteIndex"], negroni.HandlerFunc(bodysiteController.IndexHandler))...))
	bodysiteBase.Methods("POST").Handler(negroni.New(append(config["BodySiteCreate"], negroni.HandlerFunc(bodysiteController.CreateHandler))...))

	bodysite := router.Path("/BodySite/{id}").Subrouter()
	bodysite.Methods("GET").Handler(negroni.New(append(config["BodySiteShow"], negroni.HandlerFunc(bodysiteController.ShowHandler))...))
	bodysite.Methods("PUT").Handler(negroni.New(append(config["BodySiteUpdate"], negroni.HandlerFunc(bodysiteController.UpdateHandler))...))
	bodysite.Methods("DELETE").Handler(negroni.New(append(config["BodySiteDelete"], negroni.HandlerFunc(bodysiteController.DeleteHandler))...))

}
