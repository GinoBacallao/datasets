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

// V1AssemblyMatch struct for V1AssemblyMatch
type V1AssemblyMatch struct {
	Messages *[]V1Message `json:"messages,omitempty"`
	Assembly *V1AssemblyDatasetDescriptor `json:"assembly,omitempty"`
}

// NewV1AssemblyMatch instantiates a new V1AssemblyMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AssemblyMatch() *V1AssemblyMatch {
	this := V1AssemblyMatch{}
	return &this
}

// NewV1AssemblyMatchWithDefaults instantiates a new V1AssemblyMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AssemblyMatchWithDefaults() *V1AssemblyMatch {
	this := V1AssemblyMatch{}
	return &this
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *V1AssemblyMatch) GetMessages() []V1Message {
	if o == nil || o.Messages == nil {
		var ret []V1Message
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyMatch) GetMessagesOk() (*[]V1Message, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *V1AssemblyMatch) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []V1Message and assigns it to the Messages field.
func (o *V1AssemblyMatch) SetMessages(v []V1Message) {
	o.Messages = &v
}

// GetAssembly returns the Assembly field value if set, zero value otherwise.
func (o *V1AssemblyMatch) GetAssembly() V1AssemblyDatasetDescriptor {
	if o == nil || o.Assembly == nil {
		var ret V1AssemblyDatasetDescriptor
		return ret
	}
	return *o.Assembly
}

// GetAssemblyOk returns a tuple with the Assembly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyMatch) GetAssemblyOk() (*V1AssemblyDatasetDescriptor, bool) {
	if o == nil || o.Assembly == nil {
		return nil, false
	}
	return o.Assembly, true
}

// HasAssembly returns a boolean if a field has been set.
func (o *V1AssemblyMatch) HasAssembly() bool {
	if o != nil && o.Assembly != nil {
		return true
	}

	return false
}

// SetAssembly gets a reference to the given V1AssemblyDatasetDescriptor and assigns it to the Assembly field.
func (o *V1AssemblyMatch) SetAssembly(v V1AssemblyDatasetDescriptor) {
	o.Assembly = &v
}

func (o V1AssemblyMatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Messages != nil && len(o.GetMessages()) > 0  {
		toSerialize["messages"] = o.Messages
	}
	if o.Assembly != nil  {
		toSerialize["assembly"] = o.Assembly
	}
	return json.Marshal(toSerialize)
}

type NullableV1AssemblyMatch struct {
	value *V1AssemblyMatch
	isSet bool
}

func (v NullableV1AssemblyMatch) Get() *V1AssemblyMatch {
	return v.value
}

func (v *NullableV1AssemblyMatch) Set(val *V1AssemblyMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AssemblyMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AssemblyMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AssemblyMatch(val *V1AssemblyMatch) *NullableV1AssemblyMatch {
	return &NullableV1AssemblyMatch{value: val, isSet: true}
}

func (v NullableV1AssemblyMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AssemblyMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


