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

// V1GeneMatch struct for V1GeneMatch
type V1GeneMatch struct {
	Messages *[]V1Message `json:"messages,omitempty"`
	Warnings *[]V1Warning `json:"warnings,omitempty"`
	Query *[]string `json:"query,omitempty"`
	Gene *V1GeneDescriptor `json:"gene,omitempty"`
	Errors *[]V1Error `json:"errors,omitempty"`
}

// NewV1GeneMatch instantiates a new V1GeneMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GeneMatch() *V1GeneMatch {
	this := V1GeneMatch{}
	return &this
}

// NewV1GeneMatchWithDefaults instantiates a new V1GeneMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GeneMatchWithDefaults() *V1GeneMatch {
	this := V1GeneMatch{}
	return &this
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *V1GeneMatch) GetMessages() []V1Message {
	if o == nil || o.Messages == nil {
		var ret []V1Message
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneMatch) GetMessagesOk() (*[]V1Message, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *V1GeneMatch) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []V1Message and assigns it to the Messages field.
func (o *V1GeneMatch) SetMessages(v []V1Message) {
	o.Messages = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V1GeneMatch) GetWarnings() []V1Warning {
	if o == nil || o.Warnings == nil {
		var ret []V1Warning
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneMatch) GetWarningsOk() (*[]V1Warning, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V1GeneMatch) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V1Warning and assigns it to the Warnings field.
func (o *V1GeneMatch) SetWarnings(v []V1Warning) {
	o.Warnings = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *V1GeneMatch) GetQuery() []string {
	if o == nil || o.Query == nil {
		var ret []string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneMatch) GetQueryOk() (*[]string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *V1GeneMatch) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given []string and assigns it to the Query field.
func (o *V1GeneMatch) SetQuery(v []string) {
	o.Query = &v
}

// GetGene returns the Gene field value if set, zero value otherwise.
func (o *V1GeneMatch) GetGene() V1GeneDescriptor {
	if o == nil || o.Gene == nil {
		var ret V1GeneDescriptor
		return ret
	}
	return *o.Gene
}

// GetGeneOk returns a tuple with the Gene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneMatch) GetGeneOk() (*V1GeneDescriptor, bool) {
	if o == nil || o.Gene == nil {
		return nil, false
	}
	return o.Gene, true
}

// HasGene returns a boolean if a field has been set.
func (o *V1GeneMatch) HasGene() bool {
	if o != nil && o.Gene != nil {
		return true
	}

	return false
}

// SetGene gets a reference to the given V1GeneDescriptor and assigns it to the Gene field.
func (o *V1GeneMatch) SetGene(v V1GeneDescriptor) {
	o.Gene = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V1GeneMatch) GetErrors() []V1Error {
	if o == nil || o.Errors == nil {
		var ret []V1Error
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneMatch) GetErrorsOk() (*[]V1Error, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V1GeneMatch) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V1Error and assigns it to the Errors field.
func (o *V1GeneMatch) SetErrors(v []V1Error) {
	o.Errors = &v
}

func (o V1GeneMatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Messages != nil && len(o.GetMessages()) > 0  {
		toSerialize["messages"] = o.Messages
	}
	if o.Warnings != nil && len(o.GetWarnings()) > 0  {
		toSerialize["warnings"] = o.Warnings
	}
	if o.Query != nil && len(o.GetQuery()) > 0  {
		toSerialize["query"] = o.Query
	}
	if o.Gene != nil  {
		toSerialize["gene"] = o.Gene
	}
	if o.Errors != nil && len(o.GetErrors()) > 0  {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableV1GeneMatch struct {
	value *V1GeneMatch
	isSet bool
}

func (v NullableV1GeneMatch) Get() *V1GeneMatch {
	return v.value
}

func (v *NullableV1GeneMatch) Set(val *V1GeneMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GeneMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GeneMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GeneMatch(val *V1GeneMatch) *NullableV1GeneMatch {
	return &NullableV1GeneMatch{value: val, isSet: true}
}

func (v NullableV1GeneMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GeneMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


