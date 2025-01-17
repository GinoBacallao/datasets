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

// V1GeneMetadata struct for V1GeneMetadata
type V1GeneMetadata struct {
	Messages *[]V1Message `json:"messages,omitempty"`
	Genes *[]V1GeneMatch `json:"genes,omitempty"`
	TotalCount *int32 `json:"total_count,omitempty"`
}

// NewV1GeneMetadata instantiates a new V1GeneMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GeneMetadata() *V1GeneMetadata {
	this := V1GeneMetadata{}
	return &this
}

// NewV1GeneMetadataWithDefaults instantiates a new V1GeneMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GeneMetadataWithDefaults() *V1GeneMetadata {
	this := V1GeneMetadata{}
	return &this
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *V1GeneMetadata) GetMessages() []V1Message {
	if o == nil || o.Messages == nil {
		var ret []V1Message
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneMetadata) GetMessagesOk() (*[]V1Message, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *V1GeneMetadata) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []V1Message and assigns it to the Messages field.
func (o *V1GeneMetadata) SetMessages(v []V1Message) {
	o.Messages = &v
}

// GetGenes returns the Genes field value if set, zero value otherwise.
func (o *V1GeneMetadata) GetGenes() []V1GeneMatch {
	if o == nil || o.Genes == nil {
		var ret []V1GeneMatch
		return ret
	}
	return *o.Genes
}

// GetGenesOk returns a tuple with the Genes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneMetadata) GetGenesOk() (*[]V1GeneMatch, bool) {
	if o == nil || o.Genes == nil {
		return nil, false
	}
	return o.Genes, true
}

// HasGenes returns a boolean if a field has been set.
func (o *V1GeneMetadata) HasGenes() bool {
	if o != nil && o.Genes != nil {
		return true
	}

	return false
}

// SetGenes gets a reference to the given []V1GeneMatch and assigns it to the Genes field.
func (o *V1GeneMetadata) SetGenes(v []V1GeneMatch) {
	o.Genes = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *V1GeneMetadata) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneMetadata) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *V1GeneMetadata) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *V1GeneMetadata) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o V1GeneMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Messages != nil && len(o.GetMessages()) > 0  {
		toSerialize["messages"] = o.Messages
	}
	if o.Genes != nil && len(o.GetGenes()) > 0  {
		toSerialize["genes"] = o.Genes
	}
	if o.TotalCount != nil  {
		toSerialize["total_count"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableV1GeneMetadata struct {
	value *V1GeneMetadata
	isSet bool
}

func (v NullableV1GeneMetadata) Get() *V1GeneMetadata {
	return v.value
}

func (v *NullableV1GeneMetadata) Set(val *V1GeneMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GeneMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GeneMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GeneMetadata(val *V1GeneMetadata) *NullableV1GeneMetadata {
	return &NullableV1GeneMetadata{value: val, isSet: true}
}

func (v NullableV1GeneMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GeneMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


