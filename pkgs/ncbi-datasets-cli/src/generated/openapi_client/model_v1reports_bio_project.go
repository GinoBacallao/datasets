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

// V1reportsBioProject struct for V1reportsBioProject
type V1reportsBioProject struct {
	Accession *string `json:"accession,omitempty"`
	Title *string `json:"title,omitempty"`
	ParentAccession *string `json:"parent_accession,omitempty"`
	ParentAccessions *[]string `json:"parent_accessions,omitempty"`
}

// NewV1reportsBioProject instantiates a new V1reportsBioProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsBioProject() *V1reportsBioProject {
	this := V1reportsBioProject{}
	return &this
}

// NewV1reportsBioProjectWithDefaults instantiates a new V1reportsBioProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsBioProjectWithDefaults() *V1reportsBioProject {
	this := V1reportsBioProject{}
	return &this
}

// GetAccession returns the Accession field value if set, zero value otherwise.
func (o *V1reportsBioProject) GetAccession() string {
	if o == nil || o.Accession == nil {
		var ret string
		return ret
	}
	return *o.Accession
}

// GetAccessionOk returns a tuple with the Accession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioProject) GetAccessionOk() (*string, bool) {
	if o == nil || o.Accession == nil {
		return nil, false
	}
	return o.Accession, true
}

// HasAccession returns a boolean if a field has been set.
func (o *V1reportsBioProject) HasAccession() bool {
	if o != nil && o.Accession != nil {
		return true
	}

	return false
}

// SetAccession gets a reference to the given string and assigns it to the Accession field.
func (o *V1reportsBioProject) SetAccession(v string) {
	o.Accession = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *V1reportsBioProject) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioProject) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *V1reportsBioProject) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *V1reportsBioProject) SetTitle(v string) {
	o.Title = &v
}

// GetParentAccession returns the ParentAccession field value if set, zero value otherwise.
func (o *V1reportsBioProject) GetParentAccession() string {
	if o == nil || o.ParentAccession == nil {
		var ret string
		return ret
	}
	return *o.ParentAccession
}

// GetParentAccessionOk returns a tuple with the ParentAccession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioProject) GetParentAccessionOk() (*string, bool) {
	if o == nil || o.ParentAccession == nil {
		return nil, false
	}
	return o.ParentAccession, true
}

// HasParentAccession returns a boolean if a field has been set.
func (o *V1reportsBioProject) HasParentAccession() bool {
	if o != nil && o.ParentAccession != nil {
		return true
	}

	return false
}

// SetParentAccession gets a reference to the given string and assigns it to the ParentAccession field.
func (o *V1reportsBioProject) SetParentAccession(v string) {
	o.ParentAccession = &v
}

// GetParentAccessions returns the ParentAccessions field value if set, zero value otherwise.
func (o *V1reportsBioProject) GetParentAccessions() []string {
	if o == nil || o.ParentAccessions == nil {
		var ret []string
		return ret
	}
	return *o.ParentAccessions
}

// GetParentAccessionsOk returns a tuple with the ParentAccessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioProject) GetParentAccessionsOk() (*[]string, bool) {
	if o == nil || o.ParentAccessions == nil {
		return nil, false
	}
	return o.ParentAccessions, true
}

// HasParentAccessions returns a boolean if a field has been set.
func (o *V1reportsBioProject) HasParentAccessions() bool {
	if o != nil && o.ParentAccessions != nil {
		return true
	}

	return false
}

// SetParentAccessions gets a reference to the given []string and assigns it to the ParentAccessions field.
func (o *V1reportsBioProject) SetParentAccessions(v []string) {
	o.ParentAccessions = &v
}

func (o V1reportsBioProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accession != nil  {
		toSerialize["accession"] = o.Accession
	}
	if o.Title != nil  {
		toSerialize["title"] = o.Title
	}
	if o.ParentAccession != nil  {
		toSerialize["parent_accession"] = o.ParentAccession
	}
	if o.ParentAccessions != nil && len(o.GetParentAccessions()) > 0  {
		toSerialize["parent_accessions"] = o.ParentAccessions
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsBioProject struct {
	value *V1reportsBioProject
	isSet bool
}

func (v NullableV1reportsBioProject) Get() *V1reportsBioProject {
	return v.value
}

func (v *NullableV1reportsBioProject) Set(val *V1reportsBioProject) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsBioProject) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsBioProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsBioProject(val *V1reportsBioProject) *NullableV1reportsBioProject {
	return &NullableV1reportsBioProject{value: val, isSet: true}
}

func (v NullableV1reportsBioProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsBioProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


