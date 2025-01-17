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

// V1AssemblyDatasetDescriptorsFilterAssemblySource the model 'V1AssemblyDatasetDescriptorsFilterAssemblySource'
type V1AssemblyDatasetDescriptorsFilterAssemblySource string

// List of v1AssemblyDatasetDescriptorsFilterAssemblySource
const (
	V1ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYSOURCE_ALL V1AssemblyDatasetDescriptorsFilterAssemblySource = "all"
	V1ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYSOURCE_REFSEQ V1AssemblyDatasetDescriptorsFilterAssemblySource = "refseq"
	V1ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYSOURCE_GENBANK V1AssemblyDatasetDescriptorsFilterAssemblySource = "genbank"
)

// All allowed values of V1AssemblyDatasetDescriptorsFilterAssemblySource enum
var AllowedV1AssemblyDatasetDescriptorsFilterAssemblySourceEnumValues = []V1AssemblyDatasetDescriptorsFilterAssemblySource{
	"all",
	"refseq",
	"genbank",
}

func (v *V1AssemblyDatasetDescriptorsFilterAssemblySource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1AssemblyDatasetDescriptorsFilterAssemblySource(value)
	for _, existing := range AllowedV1AssemblyDatasetDescriptorsFilterAssemblySourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1AssemblyDatasetDescriptorsFilterAssemblySource", value)
}

// NewV1AssemblyDatasetDescriptorsFilterAssemblySourceFromValue returns a pointer to a valid V1AssemblyDatasetDescriptorsFilterAssemblySource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1AssemblyDatasetDescriptorsFilterAssemblySourceFromValue(v string) (*V1AssemblyDatasetDescriptorsFilterAssemblySource, error) {
	ev := V1AssemblyDatasetDescriptorsFilterAssemblySource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1AssemblyDatasetDescriptorsFilterAssemblySource: valid values are %v", v, AllowedV1AssemblyDatasetDescriptorsFilterAssemblySourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1AssemblyDatasetDescriptorsFilterAssemblySource) IsValid() bool {
	for _, existing := range AllowedV1AssemblyDatasetDescriptorsFilterAssemblySourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1AssemblyDatasetDescriptorsFilterAssemblySource value
func (v V1AssemblyDatasetDescriptorsFilterAssemblySource) Ptr() *V1AssemblyDatasetDescriptorsFilterAssemblySource {
	return &v
}

type NullableV1AssemblyDatasetDescriptorsFilterAssemblySource struct {
	value *V1AssemblyDatasetDescriptorsFilterAssemblySource
	isSet bool
}

func (v NullableV1AssemblyDatasetDescriptorsFilterAssemblySource) Get() *V1AssemblyDatasetDescriptorsFilterAssemblySource {
	return v.value
}

func (v *NullableV1AssemblyDatasetDescriptorsFilterAssemblySource) Set(val *V1AssemblyDatasetDescriptorsFilterAssemblySource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AssemblyDatasetDescriptorsFilterAssemblySource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AssemblyDatasetDescriptorsFilterAssemblySource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AssemblyDatasetDescriptorsFilterAssemblySource(val *V1AssemblyDatasetDescriptorsFilterAssemblySource) *NullableV1AssemblyDatasetDescriptorsFilterAssemblySource {
	return &NullableV1AssemblyDatasetDescriptorsFilterAssemblySource{value: val, isSet: true}
}

func (v NullableV1AssemblyDatasetDescriptorsFilterAssemblySource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AssemblyDatasetDescriptorsFilterAssemblySource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

