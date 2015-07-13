package models

import "encoding/json"

/* When FHIR JSON is unmarshalled, types that are interface{} just get unmarshaled to map[string]interface{}.
   This function converts that unmarshaled map to a specific resource type. */
func MapToResource(resourceMap interface{}) interface{} {
	b, _ := json.Marshal(&resourceMap)
	m := resourceMap.(map[string]interface{})
	t := m["resourceType"].(string)
	switch t {
	case "AllergyIntolerance":
		x := AllergyIntolerance{}
		json.Unmarshal(b, &x)
		return x
	case "Appointment":
		x := Appointment{}
		json.Unmarshal(b, &x)
		return x
	case "AppointmentResponse":
		x := AppointmentResponse{}
		json.Unmarshal(b, &x)
		return x
	case "AuditEvent":
		x := AuditEvent{}
		json.Unmarshal(b, &x)
		return x
	case "Basic":
		x := Basic{}
		json.Unmarshal(b, &x)
		return x
	case "Binary":
		x := Binary{}
		json.Unmarshal(b, &x)
		return x
	case "BodySite":
		x := BodySite{}
		json.Unmarshal(b, &x)
		return x
	case "Bundle":
		x := Bundle{}
		json.Unmarshal(b, &x)
		return x
	case "CarePlan":
		x := CarePlan{}
		json.Unmarshal(b, &x)
		return x
	case "Claim":
		x := Claim{}
		json.Unmarshal(b, &x)
		return x
	case "ClaimResponse":
		x := ClaimResponse{}
		json.Unmarshal(b, &x)
		return x
	case "ClinicalImpression":
		x := ClinicalImpression{}
		json.Unmarshal(b, &x)
		return x
	case "Communication":
		x := Communication{}
		json.Unmarshal(b, &x)
		return x
	case "CommunicationRequest":
		x := CommunicationRequest{}
		json.Unmarshal(b, &x)
		return x
	case "Composition":
		x := Composition{}
		json.Unmarshal(b, &x)
		return x
	case "ConceptMap":
		x := ConceptMap{}
		json.Unmarshal(b, &x)
		return x
	case "Condition":
		x := Condition{}
		json.Unmarshal(b, &x)
		return x
	case "Conformance":
		x := Conformance{}
		json.Unmarshal(b, &x)
		return x
	case "Contract":
		x := Contract{}
		json.Unmarshal(b, &x)
		return x
	case "Contraindication":
		x := Contraindication{}
		json.Unmarshal(b, &x)
		return x
	case "Coverage":
		x := Coverage{}
		json.Unmarshal(b, &x)
		return x
	case "DataElement":
		x := DataElement{}
		json.Unmarshal(b, &x)
		return x
	case "Device":
		x := Device{}
		json.Unmarshal(b, &x)
		return x
	case "DeviceComponent":
		x := DeviceComponent{}
		json.Unmarshal(b, &x)
		return x
	case "DeviceMetric":
		x := DeviceMetric{}
		json.Unmarshal(b, &x)
		return x
	case "DeviceUseRequest":
		x := DeviceUseRequest{}
		json.Unmarshal(b, &x)
		return x
	case "DeviceUseStatement":
		x := DeviceUseStatement{}
		json.Unmarshal(b, &x)
		return x
	case "DiagnosticOrder":
		x := DiagnosticOrder{}
		json.Unmarshal(b, &x)
		return x
	case "DiagnosticReport":
		x := DiagnosticReport{}
		json.Unmarshal(b, &x)
		return x
	case "DocumentManifest":
		x := DocumentManifest{}
		json.Unmarshal(b, &x)
		return x
	case "DocumentReference":
		x := DocumentReference{}
		json.Unmarshal(b, &x)
		return x
	case "EligibilityRequest":
		x := EligibilityRequest{}
		json.Unmarshal(b, &x)
		return x
	case "EligibilityResponse":
		x := EligibilityResponse{}
		json.Unmarshal(b, &x)
		return x
	case "Encounter":
		x := Encounter{}
		json.Unmarshal(b, &x)
		return x
	case "EnrollmentRequest":
		x := EnrollmentRequest{}
		json.Unmarshal(b, &x)
		return x
	case "EnrollmentResponse":
		x := EnrollmentResponse{}
		json.Unmarshal(b, &x)
		return x
	case "EpisodeOfCare":
		x := EpisodeOfCare{}
		json.Unmarshal(b, &x)
		return x
	case "ExplanationOfBenefit":
		x := ExplanationOfBenefit{}
		json.Unmarshal(b, &x)
		return x
	case "FamilyMemberHistory":
		x := FamilyMemberHistory{}
		json.Unmarshal(b, &x)
		return x
	case "Flag":
		x := Flag{}
		json.Unmarshal(b, &x)
		return x
	case "Goal":
		x := Goal{}
		json.Unmarshal(b, &x)
		return x
	case "Group":
		x := Group{}
		json.Unmarshal(b, &x)
		return x
	case "HealthcareService":
		x := HealthcareService{}
		json.Unmarshal(b, &x)
		return x
	case "ImagingObjectSelection":
		x := ImagingObjectSelection{}
		json.Unmarshal(b, &x)
		return x
	case "ImagingStudy":
		x := ImagingStudy{}
		json.Unmarshal(b, &x)
		return x
	case "Immunization":
		x := Immunization{}
		json.Unmarshal(b, &x)
		return x
	case "ImmunizationRecommendation":
		x := ImmunizationRecommendation{}
		json.Unmarshal(b, &x)
		return x
	case "List":
		x := List{}
		json.Unmarshal(b, &x)
		return x
	case "Location":
		x := Location{}
		json.Unmarshal(b, &x)
		return x
	case "Media":
		x := Media{}
		json.Unmarshal(b, &x)
		return x
	case "Medication":
		x := Medication{}
		json.Unmarshal(b, &x)
		return x
	case "MedicationAdministration":
		x := MedicationAdministration{}
		json.Unmarshal(b, &x)
		return x
	case "MedicationDispense":
		x := MedicationDispense{}
		json.Unmarshal(b, &x)
		return x
	case "MedicationPrescription":
		x := MedicationPrescription{}
		json.Unmarshal(b, &x)
		return x
	case "MedicationStatement":
		x := MedicationStatement{}
		json.Unmarshal(b, &x)
		return x
	case "MessageHeader":
		x := MessageHeader{}
		json.Unmarshal(b, &x)
		return x
	case "NamingSystem":
		x := NamingSystem{}
		json.Unmarshal(b, &x)
		return x
	case "NutritionOrder":
		x := NutritionOrder{}
		json.Unmarshal(b, &x)
		return x
	case "Observation":
		x := Observation{}
		json.Unmarshal(b, &x)
		return x
	case "OperationDefinition":
		x := OperationDefinition{}
		json.Unmarshal(b, &x)
		return x
	case "OperationOutcome":
		x := OperationOutcome{}
		json.Unmarshal(b, &x)
		return x
	case "Order":
		x := Order{}
		json.Unmarshal(b, &x)
		return x
	case "OrderResponse":
		x := OrderResponse{}
		json.Unmarshal(b, &x)
		return x
	case "Organization":
		x := Organization{}
		json.Unmarshal(b, &x)
		return x
	case "Patient":
		x := Patient{}
		json.Unmarshal(b, &x)
		return x
	case "PaymentNotice":
		x := PaymentNotice{}
		json.Unmarshal(b, &x)
		return x
	case "PaymentReconciliation":
		x := PaymentReconciliation{}
		json.Unmarshal(b, &x)
		return x
	case "Person":
		x := Person{}
		json.Unmarshal(b, &x)
		return x
	case "Practitioner":
		x := Practitioner{}
		json.Unmarshal(b, &x)
		return x
	case "Procedure":
		x := Procedure{}
		json.Unmarshal(b, &x)
		return x
	case "ProcedureRequest":
		x := ProcedureRequest{}
		json.Unmarshal(b, &x)
		return x
	case "ProcessRequest":
		x := ProcessRequest{}
		json.Unmarshal(b, &x)
		return x
	case "ProcessResponse":
		x := ProcessResponse{}
		json.Unmarshal(b, &x)
		return x
	case "Provenance":
		x := Provenance{}
		json.Unmarshal(b, &x)
		return x
	case "Questionnaire":
		x := Questionnaire{}
		json.Unmarshal(b, &x)
		return x
	case "QuestionnaireAnswers":
		x := QuestionnaireAnswers{}
		json.Unmarshal(b, &x)
		return x
	case "ReferralRequest":
		x := ReferralRequest{}
		json.Unmarshal(b, &x)
		return x
	case "RelatedPerson":
		x := RelatedPerson{}
		json.Unmarshal(b, &x)
		return x
	case "RiskAssessment":
		x := RiskAssessment{}
		json.Unmarshal(b, &x)
		return x
	case "Schedule":
		x := Schedule{}
		json.Unmarshal(b, &x)
		return x
	case "SearchParameter":
		x := SearchParameter{}
		json.Unmarshal(b, &x)
		return x
	case "Slot":
		x := Slot{}
		json.Unmarshal(b, &x)
		return x
	case "Specimen":
		x := Specimen{}
		json.Unmarshal(b, &x)
		return x
	case "StructureDefinition":
		x := StructureDefinition{}
		json.Unmarshal(b, &x)
		return x
	case "Subscription":
		x := Subscription{}
		json.Unmarshal(b, &x)
		return x
	case "Substance":
		x := Substance{}
		json.Unmarshal(b, &x)
		return x
	case "Supply":
		x := Supply{}
		json.Unmarshal(b, &x)
		return x
	case "SupplyDelivery":
		x := SupplyDelivery{}
		json.Unmarshal(b, &x)
		return x
	case "SupplyRequest":
		x := SupplyRequest{}
		json.Unmarshal(b, &x)
		return x
	case "TestScript":
		x := TestScript{}
		json.Unmarshal(b, &x)
		return x
	case "ValueSet":
		x := ValueSet{}
		json.Unmarshal(b, &x)
		return x
	case "VisionPrescription":
		x := VisionPrescription{}
		json.Unmarshal(b, &x)
		return x

	}
	return nil
}
