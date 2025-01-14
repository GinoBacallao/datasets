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

// V1ProkaryoteGeneRequest struct for V1ProkaryoteGeneRequest
type V1ProkaryoteGeneRequest struct {
	Accessions *[]string `json:"accessions,omitempty"`
	IncludeAnnotationType *[]V1Fasta `json:"include_annotation_type,omitempty"`
	GeneFlankConfig *V1ProkaryoteGeneRequestGeneFlankConfig `json:"gene_flank_config,omitempty"`
	Taxon *string `json:"taxon,omitempty"`
}

// NewV1ProkaryoteGeneRequest instantiates a new V1ProkaryoteGeneRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ProkaryoteGeneRequest() *V1ProkaryoteGeneRequest {
	this := V1ProkaryoteGeneRequest{}
	return &this
}

// NewV1ProkaryoteGeneRequestWithDefaults instantiates a new V1ProkaryoteGeneRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ProkaryoteGeneRequestWithDefaults() *V1ProkaryoteGeneRequest {
	this := V1ProkaryoteGeneRequest{}
	return &this
}

// GetAccessions returns the Accessions field value if set, zero value otherwise.
func (o *V1ProkaryoteGeneRequest) GetAccessions() []string {
	if o == nil || o.Accessions == nil {
		var ret []string
		return ret
	}
	return *o.Accessions
}

// GetAccessionsOk returns a tuple with the Accessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ProkaryoteGeneRequest) GetAccessionsOk() (*[]string, bool) {
	if o == nil || o.Accessions == nil {
		return nil, false
	}
	return o.Accessions, true
}

// HasAccessions returns a boolean if a field has been set.
func (o *V1ProkaryoteGeneRequest) HasAccessions() bool {
	if o != nil && o.Accessions != nil {
		return true
	}

	return false
}

// SetAccessions gets a reference to the given []string and assigns it to the Accessions field.
func (o *V1ProkaryoteGeneRequest) SetAccessions(v []string) {
	o.Accessions = &v
}

// GetIncludeAnnotationType returns the IncludeAnnotationType field value if set, zero value otherwise.
func (o *V1ProkaryoteGeneRequest) GetIncludeAnnotationType() []V1Fasta {
	if o == nil || o.IncludeAnnotationType == nil {
		var ret []V1Fasta
		return ret
	}
	return *o.IncludeAnnotationType
}

// GetIncludeAnnotationTypeOk returns a tuple with the IncludeAnnotationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ProkaryoteGeneRequest) GetIncludeAnnotationTypeOk() (*[]V1Fasta, bool) {
	if o == nil || o.IncludeAnnotationType == nil {
		return nil, false
	}
	return o.IncludeAnnotationType, true
}

// HasIncludeAnnotationType returns a boolean if a field has been set.
func (o *V1ProkaryoteGeneRequest) HasIncludeAnnotationType() bool {
	if o != nil && o.IncludeAnnotationType != nil {
		return true
	}

	return false
}

// SetIncludeAnnotationType gets a reference to the given []V1Fasta and assigns it to the IncludeAnnotationType field.
func (o *V1ProkaryoteGeneRequest) SetIncludeAnnotationType(v []V1Fasta) {
	o.IncludeAnnotationType = &v
}

// GetGeneFlankConfig returns the GeneFlankConfig field value if set, zero value otherwise.
func (o *V1ProkaryoteGeneRequest) GetGeneFlankConfig() V1ProkaryoteGeneRequestGeneFlankConfig {
	if o == nil || o.GeneFlankConfig == nil {
		var ret V1ProkaryoteGeneRequestGeneFlankConfig
		return ret
	}
	return *o.GeneFlankConfig
}

// GetGeneFlankConfigOk returns a tuple with the GeneFlankConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ProkaryoteGeneRequest) GetGeneFlankConfigOk() (*V1ProkaryoteGeneRequestGeneFlankConfig, bool) {
	if o == nil || o.GeneFlankConfig == nil {
		return nil, false
	}
	return o.GeneFlankConfig, true
}

// HasGeneFlankConfig returns a boolean if a field has been set.
func (o *V1ProkaryoteGeneRequest) HasGeneFlankConfig() bool {
	if o != nil && o.GeneFlankConfig != nil {
		return true
	}

	return false
}

// SetGeneFlankConfig gets a reference to the given V1ProkaryoteGeneRequestGeneFlankConfig and assigns it to the GeneFlankConfig field.
func (o *V1ProkaryoteGeneRequest) SetGeneFlankConfig(v V1ProkaryoteGeneRequestGeneFlankConfig) {
	o.GeneFlankConfig = &v
}

// GetTaxon returns the Taxon field value if set, zero value otherwise.
func (o *V1ProkaryoteGeneRequest) GetTaxon() string {
	if o == nil || o.Taxon == nil {
		var ret string
		return ret
	}
	return *o.Taxon
}

// GetTaxonOk returns a tuple with the Taxon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ProkaryoteGeneRequest) GetTaxonOk() (*string, bool) {
	if o == nil || o.Taxon == nil {
		return nil, false
	}
	return o.Taxon, true
}

// HasTaxon returns a boolean if a field has been set.
func (o *V1ProkaryoteGeneRequest) HasTaxon() bool {
	if o != nil && o.Taxon != nil {
		return true
	}

	return false
}

// SetTaxon gets a reference to the given string and assigns it to the Taxon field.
func (o *V1ProkaryoteGeneRequest) SetTaxon(v string) {
	o.Taxon = &v
}

func (o V1ProkaryoteGeneRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accessions != nil && len(o.GetAccessions()) > 0  {
		toSerialize["accessions"] = o.Accessions
	}
	if o.IncludeAnnotationType != nil && len(o.GetIncludeAnnotationType()) > 0  {
		toSerialize["include_annotation_type"] = o.IncludeAnnotationType
	}
	if o.GeneFlankConfig != nil  {
		toSerialize["gene_flank_config"] = o.GeneFlankConfig
	}
	if o.Taxon != nil  {
		toSerialize["taxon"] = o.Taxon
	}
	return json.Marshal(toSerialize)
}

type NullableV1ProkaryoteGeneRequest struct {
	value *V1ProkaryoteGeneRequest
	isSet bool
}

func (v NullableV1ProkaryoteGeneRequest) Get() *V1ProkaryoteGeneRequest {
	return v.value
}

func (v *NullableV1ProkaryoteGeneRequest) Set(val *V1ProkaryoteGeneRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ProkaryoteGeneRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ProkaryoteGeneRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ProkaryoteGeneRequest(val *V1ProkaryoteGeneRequest) *NullableV1ProkaryoteGeneRequest {
	return &NullableV1ProkaryoteGeneRequest{value: val, isSet: true}
}

func (v NullableV1ProkaryoteGeneRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ProkaryoteGeneRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


