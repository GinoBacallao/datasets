/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
	"fmt"
)

// V1GeneDatasetRequestSortField the model 'V1GeneDatasetRequestSortField'
type V1GeneDatasetRequestSortField string

// List of v1GeneDatasetRequestSortField
const (
	V1GENEDATASETREQUESTSORTFIELD_UNSPECIFIED V1GeneDatasetRequestSortField = "SORT_FIELD_UNSPECIFIED"
	V1GENEDATASETREQUESTSORTFIELD_GENE_ID V1GeneDatasetRequestSortField = "SORT_FIELD_GENE_ID"
	V1GENEDATASETREQUESTSORTFIELD_GENE_TYPE V1GeneDatasetRequestSortField = "SORT_FIELD_GENE_TYPE"
	V1GENEDATASETREQUESTSORTFIELD_GENE_SYMBOL V1GeneDatasetRequestSortField = "SORT_FIELD_GENE_SYMBOL"
)

// All allowed values of V1GeneDatasetRequestSortField enum
var AllowedV1GeneDatasetRequestSortFieldEnumValues = []V1GeneDatasetRequestSortField{
	"SORT_FIELD_UNSPECIFIED",
	"SORT_FIELD_GENE_ID",
	"SORT_FIELD_GENE_TYPE",
	"SORT_FIELD_GENE_SYMBOL",
}

func (v *V1GeneDatasetRequestSortField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1GeneDatasetRequestSortField(value)
	for _, existing := range AllowedV1GeneDatasetRequestSortFieldEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1GeneDatasetRequestSortField", value)
}

// NewV1GeneDatasetRequestSortFieldFromValue returns a pointer to a valid V1GeneDatasetRequestSortField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1GeneDatasetRequestSortFieldFromValue(v string) (*V1GeneDatasetRequestSortField, error) {
	ev := V1GeneDatasetRequestSortField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1GeneDatasetRequestSortField: valid values are %v", v, AllowedV1GeneDatasetRequestSortFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1GeneDatasetRequestSortField) IsValid() bool {
	for _, existing := range AllowedV1GeneDatasetRequestSortFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1GeneDatasetRequestSortField value
func (v V1GeneDatasetRequestSortField) Ptr() *V1GeneDatasetRequestSortField {
	return &v
}

type NullableV1GeneDatasetRequestSortField struct {
	value *V1GeneDatasetRequestSortField
	isSet bool
}

func (v NullableV1GeneDatasetRequestSortField) Get() *V1GeneDatasetRequestSortField {
	return v.value
}

func (v *NullableV1GeneDatasetRequestSortField) Set(val *V1GeneDatasetRequestSortField) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GeneDatasetRequestSortField) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GeneDatasetRequestSortField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GeneDatasetRequestSortField(val *V1GeneDatasetRequestSortField) *NullableV1GeneDatasetRequestSortField {
	return &NullableV1GeneDatasetRequestSortField{value: val, isSet: true}
}

func (v NullableV1GeneDatasetRequestSortField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GeneDatasetRequestSortField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

