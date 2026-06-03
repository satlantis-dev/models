package models

type RegistrationAnswerType string

const (
	RegistrationAnswerTypeText         RegistrationAnswerType = "text"
	RegistrationAnswerTypeLongText     RegistrationAnswerType = "long-text"
	RegistrationAnswerTypeEmail        RegistrationAnswerType = "email"
	RegistrationAnswerTypePhone        RegistrationAnswerType = "phone"
	RegistrationAnswerTypeURL          RegistrationAnswerType = "url"
	RegistrationAnswerTypeNpub         RegistrationAnswerType = "npub"
	RegistrationAnswerTypeAgreeCheck   RegistrationAnswerType = "agree-check"
	RegistrationAnswerTypeFacebook     RegistrationAnswerType = "facebook"
	RegistrationAnswerTypeX            RegistrationAnswerType = "x"
	RegistrationAnswerTypeInsta        RegistrationAnswerType = "insta"
	RegistrationAnswerTypeLinkedIn     RegistrationAnswerType = "linkedin"
	RegistrationAnswerTypeSingleSelect RegistrationAnswerType = "single-select"
	RegistrationAnswerTypeMultiSelect  RegistrationAnswerType = "multi-select"
)

type RegistrationQuestion struct {
	Label      string                 `json:"label"`
	Required   bool                   `json:"required"`
	AnswerType RegistrationAnswerType `json:"answerType"`
	Options    []string               `json:"options,omitempty"`
}

type RegistrationQuestionsPayload struct {
	Questions []RegistrationQuestion `json:"questions"`
}

type RegistrationAnswer struct {
	Label      string                 `json:"label"`
	Answer     any                    `json:"answer"`
	AnswerType RegistrationAnswerType `json:"answerType"`
}

type RegistrationAnswersPayload struct {
	Answers []RegistrationAnswer `json:"answers"`
}
