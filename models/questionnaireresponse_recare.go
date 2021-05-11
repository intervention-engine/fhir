package models

type Meta_Recare struct {
	LastUpdated string `json:"lastUpdated,omitempty"`
}

type Link []struct {
	Relation string `json:"relation,omitempty"`
	Url      string `json:"url,omitempty"`
}

type QuestionnaireResponses struct {
	Entry        QuestionnaireEntry `json:"entry,omitempty"`
	ID           *string            `json:"id,omitempty"`
	Meta         *Meta_Recare       `json:"meta,omitempty"`
	Type         *string            `json:"type,omitempty"`
	Total        *int               `json:"total,omitempty"`
	ResourceType *string            `json:"resourceType,omitempty"`
	Link         Link               `json:"link,omitempty"`
}

type Identifier_Recare struct {
	System *string `json:"system,omitempty"`
	Value  *string `json:"value,omitempty"`
}

type Resource struct {
	ResourceType  *string            `json:"resourceType,omitempty"`
	ID            *string            `json:"id,omitempty"`
	Identifier    *Identifier_Recare `json:"identifier,omitempty"`
	Status        *string            `json:"status,omitempty"`
	Authored      *string            `json:"authored,omitempty"`
	Questionnaire string             `json:"questionnaire,omitempty"`
	Item          QuestionnaireItem  `json:"item,omitempty"`
}

type QuestionnaireEntry []struct {
	FullUrl  *string   `json:"fullUrl,omitempty"`
	Resource *Resource `json:"resource,omitempty"`
}

type QuestionnaireItem []struct {
	LinkId string                `json:"linkId,omitempty"`
	Text   *string               `json:"text,omitempty"`
	Item   QuestionnaireItem     `json:"item,omitempty"`
	Answer []QuestionnaireAnswer `json:"answer,omitempty"`
}

type QuestionnaireAnswer struct {
	ValueBoolean *bool    `json:"valueBoolean,omitempty"`
	ValueDecimal *float64 `json:"valueDecimal,omitempty"`
	ValueInteger *int32   `json:"valueInteger,omitempty"`
	ValueString  *string  `json:"valueString,omitempty"`
	ValueUri     *string  `json:"valueUri,omitempty"`
}
