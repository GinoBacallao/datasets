/*
 * NCBI Datasets API
 *
 * ### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/rehydrate/), and retrieve the individual data files at a later time. 
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
)

// V1reportsProkaryoteGeneProteinNameEvidence struct for V1reportsProkaryoteGeneProteinNameEvidence
type V1reportsProkaryoteGeneProteinNameEvidence struct {
	Accession *string `json:"accession,omitempty"`
	Category *string `json:"category,omitempty"`
	Source *string `json:"source,omitempty"`
}

// NewV1reportsProkaryoteGeneProteinNameEvidence instantiates a new V1reportsProkaryoteGeneProteinNameEvidence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsProkaryoteGeneProteinNameEvidence() *V1reportsProkaryoteGeneProteinNameEvidence {
	this := V1reportsProkaryoteGeneProteinNameEvidence{}
	return &this
}

// NewV1reportsProkaryoteGeneProteinNameEvidenceWithDefaults instantiates a new V1reportsProkaryoteGeneProteinNameEvidence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsProkaryoteGeneProteinNameEvidenceWithDefaults() *V1reportsProkaryoteGeneProteinNameEvidence {
	this := V1reportsProkaryoteGeneProteinNameEvidence{}
	return &this
}

// GetAccession returns the Accession field value if set, zero value otherwise.
func (o *V1reportsProkaryoteGeneProteinNameEvidence) GetAccession() string {
	if o == nil || o.Accession == nil {
		var ret string
		return ret
	}
	return *o.Accession
}

// GetAccessionOk returns a tuple with the Accession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsProkaryoteGeneProteinNameEvidence) GetAccessionOk() (*string, bool) {
	if o == nil || o.Accession == nil {
		return nil, false
	}
	return o.Accession, true
}

// HasAccession returns a boolean if a field has been set.
func (o *V1reportsProkaryoteGeneProteinNameEvidence) HasAccession() bool {
	if o != nil && o.Accession != nil {
		return true
	}

	return false
}

// SetAccession gets a reference to the given string and assigns it to the Accession field.
func (o *V1reportsProkaryoteGeneProteinNameEvidence) SetAccession(v string) {
	o.Accession = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *V1reportsProkaryoteGeneProteinNameEvidence) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsProkaryoteGeneProteinNameEvidence) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *V1reportsProkaryoteGeneProteinNameEvidence) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *V1reportsProkaryoteGeneProteinNameEvidence) SetCategory(v string) {
	o.Category = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *V1reportsProkaryoteGeneProteinNameEvidence) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsProkaryoteGeneProteinNameEvidence) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *V1reportsProkaryoteGeneProteinNameEvidence) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *V1reportsProkaryoteGeneProteinNameEvidence) SetSource(v string) {
	o.Source = &v
}

func (o V1reportsProkaryoteGeneProteinNameEvidence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accession != nil  {
		toSerialize["accession"] = o.Accession
	}
	if o.Category != nil  {
		toSerialize["category"] = o.Category
	}
	if o.Source != nil  {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsProkaryoteGeneProteinNameEvidence struct {
	value *V1reportsProkaryoteGeneProteinNameEvidence
	isSet bool
}

func (v NullableV1reportsProkaryoteGeneProteinNameEvidence) Get() *V1reportsProkaryoteGeneProteinNameEvidence {
	return v.value
}

func (v *NullableV1reportsProkaryoteGeneProteinNameEvidence) Set(val *V1reportsProkaryoteGeneProteinNameEvidence) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsProkaryoteGeneProteinNameEvidence) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsProkaryoteGeneProteinNameEvidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsProkaryoteGeneProteinNameEvidence(val *V1reportsProkaryoteGeneProteinNameEvidence) *NullableV1reportsProkaryoteGeneProteinNameEvidence {
	return &NullableV1reportsProkaryoteGeneProteinNameEvidence{value: val, isSet: true}
}

func (v NullableV1reportsProkaryoteGeneProteinNameEvidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsProkaryoteGeneProteinNameEvidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

