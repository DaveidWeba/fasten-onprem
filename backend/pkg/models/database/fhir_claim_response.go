// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-onprem/blob/main/backend/pkg/models/database/generate.go
// PLEASE DO NOT EDIT BY HAND

package database

import (
	"encoding/json"
	"fmt"
	goja "github.com/dop251/goja"
	models "github.com/fastenhealth/fasten-onprem/backend/pkg/models"
	datatypes "gorm.io/datatypes"
	"time"
)

type FhirClaimResponse struct {
	models.ResourceBase
	// The creation date
	// https://hl7.org/fhir/r4/search.html#date
	Created *time.Time `gorm:"column:created;type:datetime" json:"created,omitempty"`
	// The contents of the disposition message
	// https://hl7.org/fhir/r4/search.html#string
	Disposition datatypes.JSON `gorm:"column:disposition;type:text;serializer:json" json:"disposition,omitempty"`
	// The identity of the ClaimResponse
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// The organization which generated this resource
	// https://hl7.org/fhir/r4/search.html#reference
	Insurer datatypes.JSON `gorm:"column:insurer;type:text;serializer:json" json:"insurer,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	MetaLastUpdated *time.Time `gorm:"column:metaLastUpdated;type:datetime" json:"metaLastUpdated,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	MetaProfile datatypes.JSON `gorm:"column:metaProfile;type:text;serializer:json" json:"metaProfile,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	MetaTag datatypes.JSON `gorm:"column:metaTag;type:text;serializer:json" json:"metaTag,omitempty"`
	// Tags applied to this resource
	// This is a primitive string literal (`keyword` type). It is not a recognized SearchParameter type from https://hl7.org/fhir/r4/search.html, it's Fasten Health-specific
	MetaVersionId string `gorm:"column:metaVersionId;type:text" json:"metaVersionId,omitempty"`
	// Notes/comments
	// https://hl7.org/fhir/r4/search.html#string
	Note datatypes.JSON `gorm:"column:note;type:text;serializer:json" json:"note,omitempty"`
	// The processing outcome
	// https://hl7.org/fhir/r4/search.html#token
	Outcome datatypes.JSON `gorm:"column:outcome;type:text;serializer:json" json:"outcome,omitempty"`
	// The expected payment date
	// https://hl7.org/fhir/r4/search.html#date
	PaymentDate *time.Time `gorm:"column:paymentDate;type:datetime" json:"paymentDate,omitempty"`
	// The claim reference
	// https://hl7.org/fhir/r4/search.html#reference
	Request datatypes.JSON `gorm:"column:request;type:text;serializer:json" json:"request,omitempty"`
	// The Provider of the claim
	// https://hl7.org/fhir/r4/search.html#reference
	Requestor datatypes.JSON `gorm:"column:requestor;type:text;serializer:json" json:"requestor,omitempty"`
	// The status of the ClaimResponse
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text datatypes.JSON `gorm:"column:text;type:text;serializer:json" json:"text,omitempty"`
	// The type of claim
	// https://hl7.org/fhir/r4/search.html#token
	Use datatypes.JSON `gorm:"column:use;type:text;serializer:json" json:"use,omitempty"`
}

func (s *FhirClaimResponse) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"created":              "date",
		"disposition":          "string",
		"id":                   "keyword",
		"identifier":           "token",
		"insurer":              "reference",
		"language":             "token",
		"metaLastUpdated":      "date",
		"metaProfile":          "reference",
		"metaTag":              "token",
		"metaVersionId":        "keyword",
		"note":                 "string",
		"outcome":              "token",
		"paymentDate":          "date",
		"request":              "reference",
		"requestor":            "reference",
		"sort_date":            "date",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"status":               "token",
		"text":                 "string",
		"use":                  "token",
	}
	return searchParameters
}
func (s *FhirClaimResponse) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
	s.ResourceRaw = datatypes.JSON(resourceRaw)
	// unmarshal the raw resource (bytes) into a map
	var resourceRawMap map[string]interface{}
	err := json.Unmarshal(resourceRaw, &resourceRawMap)
	if err != nil {
		return err
	}
	if len(fhirPathJs) == 0 {
		return fmt.Errorf("fhirPathJs script is empty")
	}
	vm := goja.New()
	// setup the global window object
	vm.Set("window", vm.NewObject())
	// set the global FHIR Resource object
	vm.Set("fhirResource", resourceRawMap)
	// compile the fhirpath library
	fhirPathJsProgram, err := goja.Compile("fhirpath.min.js", fhirPathJs, true)
	if err != nil {
		return err
	}
	// compile the searchParametersExtractor library
	searchParametersExtractorJsProgram, err := goja.Compile("searchParameterExtractor.js", searchParameterExtractorJs, true)
	if err != nil {
		return err
	}
	// add the fhirpath library in the goja vm
	_, err = vm.RunProgram(fhirPathJsProgram)
	if err != nil {
		return err
	}
	// add the searchParametersExtractor library in the goja vm
	_, err = vm.RunProgram(searchParametersExtractorJsProgram)
	if err != nil {
		return err
	}
	// execute the fhirpath expression for each search parameter
	// extracting Created
	createdResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'ClaimResponse.created')")
	if err == nil && createdResult.String() != "undefined" {
		if t, err := time.Parse(time.RFC3339, createdResult.String()); err == nil {
			s.Created = &t
		} else if t, err = time.Parse("2006-01-02", createdResult.String()); err == nil {
			s.Created = &t
		} else if t, err = time.Parse("2006-01", createdResult.String()); err == nil {
			s.Created = &t
		} else if t, err = time.Parse("2006", createdResult.String()); err == nil {
			s.Created = &t
		}
	}
	// extracting Disposition
	dispositionResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'ClaimResponse.disposition')")
	if err == nil && dispositionResult.String() != "undefined" {
		s.Disposition = []byte(dispositionResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'ClaimResponse.identifier')")
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Insurer
	insurerResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'ClaimResponse.insurer')")
	if err == nil && insurerResult.String() != "undefined" {
		s.Insurer = []byte(insurerResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'language')")
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting MetaLastUpdated
	metaLastUpdatedResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'meta.lastUpdated')")
	if err == nil && metaLastUpdatedResult.String() != "undefined" {
		if t, err := time.Parse(time.RFC3339, metaLastUpdatedResult.String()); err == nil {
			s.MetaLastUpdated = &t
		} else if t, err = time.Parse("2006-01-02", metaLastUpdatedResult.String()); err == nil {
			s.MetaLastUpdated = &t
		} else if t, err = time.Parse("2006-01", metaLastUpdatedResult.String()); err == nil {
			s.MetaLastUpdated = &t
		} else if t, err = time.Parse("2006", metaLastUpdatedResult.String()); err == nil {
			s.MetaLastUpdated = &t
		}
	}
	// extracting MetaProfile
	metaProfileResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'meta.profile')")
	if err == nil && metaProfileResult.String() != "undefined" {
		s.MetaProfile = []byte(metaProfileResult.String())
	}
	// extracting MetaTag
	metaTagResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'meta.tag')")
	if err == nil && metaTagResult.String() != "undefined" {
		s.MetaTag = []byte(metaTagResult.String())
	}
	// extracting MetaVersionId
	metaVersionIdResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'meta.versionId')")
	if err == nil && metaVersionIdResult.String() != "undefined" {
		s.MetaVersionId = metaVersionIdResult.String()
	}
	// extracting Note
	noteResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'note')")
	if err == nil && noteResult.String() != "undefined" {
		s.Note = []byte(noteResult.String())
	}
	// extracting Outcome
	outcomeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'ClaimResponse.outcome')")
	if err == nil && outcomeResult.String() != "undefined" {
		s.Outcome = []byte(outcomeResult.String())
	}
	// extracting PaymentDate
	paymentDateResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'ClaimResponse.payment.date')")
	if err == nil && paymentDateResult.String() != "undefined" {
		if t, err := time.Parse(time.RFC3339, paymentDateResult.String()); err == nil {
			s.PaymentDate = &t
		} else if t, err = time.Parse("2006-01-02", paymentDateResult.String()); err == nil {
			s.PaymentDate = &t
		} else if t, err = time.Parse("2006-01", paymentDateResult.String()); err == nil {
			s.PaymentDate = &t
		} else if t, err = time.Parse("2006", paymentDateResult.String()); err == nil {
			s.PaymentDate = &t
		}
	}
	// extracting Request
	requestResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'ClaimResponse.request')")
	if err == nil && requestResult.String() != "undefined" {
		s.Request = []byte(requestResult.String())
	}
	// extracting Requestor
	requestorResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'ClaimResponse.requestor')")
	if err == nil && requestorResult.String() != "undefined" {
		s.Requestor = []byte(requestorResult.String())
	}
	// extracting Status
	statusResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'ClaimResponse.status')")
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Text
	textResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'text')")
	if err == nil && textResult.String() != "undefined" {
		s.Text = []byte(textResult.String())
	}
	// extracting Use
	useResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'ClaimResponse.use')")
	if err == nil && useResult.String() != "undefined" {
		s.Use = []byte(useResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirClaimResponse) TableName() string {
	return "fhir_claim_response"
}
