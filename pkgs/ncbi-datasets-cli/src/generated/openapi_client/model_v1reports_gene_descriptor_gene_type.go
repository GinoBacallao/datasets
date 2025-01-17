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

// V1reportsGeneDescriptorGeneType the model 'V1reportsGeneDescriptorGeneType'
type V1reportsGeneDescriptorGeneType string

// List of v1reportsGeneDescriptorGeneType
const (
	V1REPORTSGENEDESCRIPTORGENETYPE_UNKNOWN V1reportsGeneDescriptorGeneType = "UNKNOWN"
	V1REPORTSGENEDESCRIPTORGENETYPE_T_RNA V1reportsGeneDescriptorGeneType = "tRNA"
	V1REPORTSGENEDESCRIPTORGENETYPE_R_RNA V1reportsGeneDescriptorGeneType = "rRNA"
	V1REPORTSGENEDESCRIPTORGENETYPE_SN_RNA V1reportsGeneDescriptorGeneType = "snRNA"
	V1REPORTSGENEDESCRIPTORGENETYPE_SC_RNA V1reportsGeneDescriptorGeneType = "scRNA"
	V1REPORTSGENEDESCRIPTORGENETYPE_SNO_RNA V1reportsGeneDescriptorGeneType = "snoRNA"
	V1REPORTSGENEDESCRIPTORGENETYPE_PROTEIN_CODING V1reportsGeneDescriptorGeneType = "PROTEIN_CODING"
	V1REPORTSGENEDESCRIPTORGENETYPE_PSEUDO V1reportsGeneDescriptorGeneType = "PSEUDO"
	V1REPORTSGENEDESCRIPTORGENETYPE_TRANSPOSON V1reportsGeneDescriptorGeneType = "TRANSPOSON"
	V1REPORTSGENEDESCRIPTORGENETYPE_MISC_RNA V1reportsGeneDescriptorGeneType = "miscRNA"
	V1REPORTSGENEDESCRIPTORGENETYPE_NC_RNA V1reportsGeneDescriptorGeneType = "ncRNA"
	V1REPORTSGENEDESCRIPTORGENETYPE_BIOLOGICAL_REGION V1reportsGeneDescriptorGeneType = "BIOLOGICAL_REGION"
	V1REPORTSGENEDESCRIPTORGENETYPE_OTHER V1reportsGeneDescriptorGeneType = "OTHER"
)

// All allowed values of V1reportsGeneDescriptorGeneType enum
var AllowedV1reportsGeneDescriptorGeneTypeEnumValues = []V1reportsGeneDescriptorGeneType{
	"UNKNOWN",
	"tRNA",
	"rRNA",
	"snRNA",
	"scRNA",
	"snoRNA",
	"PROTEIN_CODING",
	"PSEUDO",
	"TRANSPOSON",
	"miscRNA",
	"ncRNA",
	"BIOLOGICAL_REGION",
	"OTHER",
}

func (v *V1reportsGeneDescriptorGeneType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1reportsGeneDescriptorGeneType(value)
	for _, existing := range AllowedV1reportsGeneDescriptorGeneTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1reportsGeneDescriptorGeneType", value)
}

// NewV1reportsGeneDescriptorGeneTypeFromValue returns a pointer to a valid V1reportsGeneDescriptorGeneType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1reportsGeneDescriptorGeneTypeFromValue(v string) (*V1reportsGeneDescriptorGeneType, error) {
	ev := V1reportsGeneDescriptorGeneType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1reportsGeneDescriptorGeneType: valid values are %v", v, AllowedV1reportsGeneDescriptorGeneTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1reportsGeneDescriptorGeneType) IsValid() bool {
	for _, existing := range AllowedV1reportsGeneDescriptorGeneTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1reportsGeneDescriptorGeneType value
func (v V1reportsGeneDescriptorGeneType) Ptr() *V1reportsGeneDescriptorGeneType {
	return &v
}

type NullableV1reportsGeneDescriptorGeneType struct {
	value *V1reportsGeneDescriptorGeneType
	isSet bool
}

func (v NullableV1reportsGeneDescriptorGeneType) Get() *V1reportsGeneDescriptorGeneType {
	return v.value
}

func (v *NullableV1reportsGeneDescriptorGeneType) Set(val *V1reportsGeneDescriptorGeneType) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsGeneDescriptorGeneType) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsGeneDescriptorGeneType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsGeneDescriptorGeneType(val *V1reportsGeneDescriptorGeneType) *NullableV1reportsGeneDescriptorGeneType {
	return &NullableV1reportsGeneDescriptorGeneType{value: val, isSet: true}
}

func (v NullableV1reportsGeneDescriptorGeneType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsGeneDescriptorGeneType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

