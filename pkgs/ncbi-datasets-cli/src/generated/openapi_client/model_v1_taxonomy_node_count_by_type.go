/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
)

// V1TaxonomyNodeCountByType struct for V1TaxonomyNodeCountByType
type V1TaxonomyNodeCountByType struct {
	Type *V1CountType `json:"type,omitempty"`
	Count *int32 `json:"count,omitempty"`
}

// NewV1TaxonomyNodeCountByType instantiates a new V1TaxonomyNodeCountByType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1TaxonomyNodeCountByType() *V1TaxonomyNodeCountByType {
	this := V1TaxonomyNodeCountByType{}
	var type_ V1CountType = V1COUNTTYPE_UNSPECIFIED
	this.Type = &type_
	return &this
}

// NewV1TaxonomyNodeCountByTypeWithDefaults instantiates a new V1TaxonomyNodeCountByType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1TaxonomyNodeCountByTypeWithDefaults() *V1TaxonomyNodeCountByType {
	this := V1TaxonomyNodeCountByType{}
	var type_ V1CountType = V1COUNTTYPE_UNSPECIFIED
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1TaxonomyNodeCountByType) GetType() V1CountType {
	if o == nil || o.Type == nil {
		var ret V1CountType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyNodeCountByType) GetTypeOk() (*V1CountType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1TaxonomyNodeCountByType) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given V1CountType and assigns it to the Type field.
func (o *V1TaxonomyNodeCountByType) SetType(v V1CountType) {
	o.Type = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V1TaxonomyNodeCountByType) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyNodeCountByType) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V1TaxonomyNodeCountByType) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *V1TaxonomyNodeCountByType) SetCount(v int32) {
	o.Count = &v
}

func (o V1TaxonomyNodeCountByType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil  {
		toSerialize["type"] = o.Type
	}
	if o.Count != nil  {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableV1TaxonomyNodeCountByType struct {
	value *V1TaxonomyNodeCountByType
	isSet bool
}

func (v NullableV1TaxonomyNodeCountByType) Get() *V1TaxonomyNodeCountByType {
	return v.value
}

func (v *NullableV1TaxonomyNodeCountByType) Set(val *V1TaxonomyNodeCountByType) {
	v.value = val
	v.isSet = true
}

func (v NullableV1TaxonomyNodeCountByType) IsSet() bool {
	return v.isSet
}

func (v *NullableV1TaxonomyNodeCountByType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1TaxonomyNodeCountByType(val *V1TaxonomyNodeCountByType) *NullableV1TaxonomyNodeCountByType {
	return &NullableV1TaxonomyNodeCountByType{value: val, isSet: true}
}

func (v NullableV1TaxonomyNodeCountByType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1TaxonomyNodeCountByType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


