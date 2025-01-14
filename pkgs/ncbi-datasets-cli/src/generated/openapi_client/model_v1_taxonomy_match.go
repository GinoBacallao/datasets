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

// V1TaxonomyMatch struct for V1TaxonomyMatch
type V1TaxonomyMatch struct {
	Warnings *[]V1Warning `json:"warnings,omitempty"`
	Errors *[]V1Error `json:"errors,omitempty"`
	Query *[]string `json:"query,omitempty"`
	Taxonomy *V1TaxonomyNode `json:"taxonomy,omitempty"`
}

// NewV1TaxonomyMatch instantiates a new V1TaxonomyMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1TaxonomyMatch() *V1TaxonomyMatch {
	this := V1TaxonomyMatch{}
	return &this
}

// NewV1TaxonomyMatchWithDefaults instantiates a new V1TaxonomyMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1TaxonomyMatchWithDefaults() *V1TaxonomyMatch {
	this := V1TaxonomyMatch{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V1TaxonomyMatch) GetWarnings() []V1Warning {
	if o == nil || o.Warnings == nil {
		var ret []V1Warning
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyMatch) GetWarningsOk() (*[]V1Warning, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V1TaxonomyMatch) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V1Warning and assigns it to the Warnings field.
func (o *V1TaxonomyMatch) SetWarnings(v []V1Warning) {
	o.Warnings = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V1TaxonomyMatch) GetErrors() []V1Error {
	if o == nil || o.Errors == nil {
		var ret []V1Error
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyMatch) GetErrorsOk() (*[]V1Error, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V1TaxonomyMatch) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V1Error and assigns it to the Errors field.
func (o *V1TaxonomyMatch) SetErrors(v []V1Error) {
	o.Errors = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *V1TaxonomyMatch) GetQuery() []string {
	if o == nil || o.Query == nil {
		var ret []string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyMatch) GetQueryOk() (*[]string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *V1TaxonomyMatch) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given []string and assigns it to the Query field.
func (o *V1TaxonomyMatch) SetQuery(v []string) {
	o.Query = &v
}

// GetTaxonomy returns the Taxonomy field value if set, zero value otherwise.
func (o *V1TaxonomyMatch) GetTaxonomy() V1TaxonomyNode {
	if o == nil || o.Taxonomy == nil {
		var ret V1TaxonomyNode
		return ret
	}
	return *o.Taxonomy
}

// GetTaxonomyOk returns a tuple with the Taxonomy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyMatch) GetTaxonomyOk() (*V1TaxonomyNode, bool) {
	if o == nil || o.Taxonomy == nil {
		return nil, false
	}
	return o.Taxonomy, true
}

// HasTaxonomy returns a boolean if a field has been set.
func (o *V1TaxonomyMatch) HasTaxonomy() bool {
	if o != nil && o.Taxonomy != nil {
		return true
	}

	return false
}

// SetTaxonomy gets a reference to the given V1TaxonomyNode and assigns it to the Taxonomy field.
func (o *V1TaxonomyMatch) SetTaxonomy(v V1TaxonomyNode) {
	o.Taxonomy = &v
}

func (o V1TaxonomyMatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Warnings != nil && len(o.GetWarnings()) > 0  {
		toSerialize["warnings"] = o.Warnings
	}
	if o.Errors != nil && len(o.GetErrors()) > 0  {
		toSerialize["errors"] = o.Errors
	}
	if o.Query != nil && len(o.GetQuery()) > 0  {
		toSerialize["query"] = o.Query
	}
	if o.Taxonomy != nil  {
		toSerialize["taxonomy"] = o.Taxonomy
	}
	return json.Marshal(toSerialize)
}

type NullableV1TaxonomyMatch struct {
	value *V1TaxonomyMatch
	isSet bool
}

func (v NullableV1TaxonomyMatch) Get() *V1TaxonomyMatch {
	return v.value
}

func (v *NullableV1TaxonomyMatch) Set(val *V1TaxonomyMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableV1TaxonomyMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableV1TaxonomyMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1TaxonomyMatch(val *V1TaxonomyMatch) *NullableV1TaxonomyMatch {
	return &NullableV1TaxonomyMatch{value: val, isSet: true}
}

func (v NullableV1TaxonomyMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1TaxonomyMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


