/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/rehydrate/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
	"time"
)

// V1AssemblyDatasetDescriptorsFilter struct for V1AssemblyDatasetDescriptorsFilter
type V1AssemblyDatasetDescriptorsFilter struct {
	ReferenceOnly *bool `json:"reference_only,omitempty"`
	RefseqOnly *bool `json:"refseq_only,omitempty"`
	AssemblySource *V1AssemblyDatasetDescriptorsFilterAssemblySource `json:"assembly_source,omitempty"`
	HasAnnotation *bool `json:"has_annotation,omitempty"`
	AssemblyLevel *[]V1AssemblyDatasetDescriptorsFilterAssemblyLevel `json:"assembly_level,omitempty"`
	FirstReleaseDate *time.Time `json:"first_release_date,omitempty"`
	LastReleaseDate *time.Time `json:"last_release_date,omitempty"`
	SearchText *[]string `json:"search_text,omitempty"`
}

// NewV1AssemblyDatasetDescriptorsFilter instantiates a new V1AssemblyDatasetDescriptorsFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AssemblyDatasetDescriptorsFilter() *V1AssemblyDatasetDescriptorsFilter {
	this := V1AssemblyDatasetDescriptorsFilter{}
	var assemblySource V1AssemblyDatasetDescriptorsFilterAssemblySource = V1ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYSOURCE_ALL
	this.AssemblySource = &assemblySource
	return &this
}

// NewV1AssemblyDatasetDescriptorsFilterWithDefaults instantiates a new V1AssemblyDatasetDescriptorsFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AssemblyDatasetDescriptorsFilterWithDefaults() *V1AssemblyDatasetDescriptorsFilter {
	this := V1AssemblyDatasetDescriptorsFilter{}
	var assemblySource V1AssemblyDatasetDescriptorsFilterAssemblySource = V1ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYSOURCE_ALL
	this.AssemblySource = &assemblySource
	return &this
}

// GetReferenceOnly returns the ReferenceOnly field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsFilter) GetReferenceOnly() bool {
	if o == nil || o.ReferenceOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReferenceOnly
}

// GetReferenceOnlyOk returns a tuple with the ReferenceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsFilter) GetReferenceOnlyOk() (*bool, bool) {
	if o == nil || o.ReferenceOnly == nil {
		return nil, false
	}
	return o.ReferenceOnly, true
}

// HasReferenceOnly returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsFilter) HasReferenceOnly() bool {
	if o != nil && o.ReferenceOnly != nil {
		return true
	}

	return false
}

// SetReferenceOnly gets a reference to the given bool and assigns it to the ReferenceOnly field.
func (o *V1AssemblyDatasetDescriptorsFilter) SetReferenceOnly(v bool) {
	o.ReferenceOnly = &v
}

// GetRefseqOnly returns the RefseqOnly field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsFilter) GetRefseqOnly() bool {
	if o == nil || o.RefseqOnly == nil {
		var ret bool
		return ret
	}
	return *o.RefseqOnly
}

// GetRefseqOnlyOk returns a tuple with the RefseqOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsFilter) GetRefseqOnlyOk() (*bool, bool) {
	if o == nil || o.RefseqOnly == nil {
		return nil, false
	}
	return o.RefseqOnly, true
}

// HasRefseqOnly returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsFilter) HasRefseqOnly() bool {
	if o != nil && o.RefseqOnly != nil {
		return true
	}

	return false
}

// SetRefseqOnly gets a reference to the given bool and assigns it to the RefseqOnly field.
func (o *V1AssemblyDatasetDescriptorsFilter) SetRefseqOnly(v bool) {
	o.RefseqOnly = &v
}

// GetAssemblySource returns the AssemblySource field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsFilter) GetAssemblySource() V1AssemblyDatasetDescriptorsFilterAssemblySource {
	if o == nil || o.AssemblySource == nil {
		var ret V1AssemblyDatasetDescriptorsFilterAssemblySource
		return ret
	}
	return *o.AssemblySource
}

// GetAssemblySourceOk returns a tuple with the AssemblySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsFilter) GetAssemblySourceOk() (*V1AssemblyDatasetDescriptorsFilterAssemblySource, bool) {
	if o == nil || o.AssemblySource == nil {
		return nil, false
	}
	return o.AssemblySource, true
}

// HasAssemblySource returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsFilter) HasAssemblySource() bool {
	if o != nil && o.AssemblySource != nil {
		return true
	}

	return false
}

// SetAssemblySource gets a reference to the given V1AssemblyDatasetDescriptorsFilterAssemblySource and assigns it to the AssemblySource field.
func (o *V1AssemblyDatasetDescriptorsFilter) SetAssemblySource(v V1AssemblyDatasetDescriptorsFilterAssemblySource) {
	o.AssemblySource = &v
}

// GetHasAnnotation returns the HasAnnotation field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsFilter) GetHasAnnotation() bool {
	if o == nil || o.HasAnnotation == nil {
		var ret bool
		return ret
	}
	return *o.HasAnnotation
}

// GetHasAnnotationOk returns a tuple with the HasAnnotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsFilter) GetHasAnnotationOk() (*bool, bool) {
	if o == nil || o.HasAnnotation == nil {
		return nil, false
	}
	return o.HasAnnotation, true
}

// HasHasAnnotation returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsFilter) HasHasAnnotation() bool {
	if o != nil && o.HasAnnotation != nil {
		return true
	}

	return false
}

// SetHasAnnotation gets a reference to the given bool and assigns it to the HasAnnotation field.
func (o *V1AssemblyDatasetDescriptorsFilter) SetHasAnnotation(v bool) {
	o.HasAnnotation = &v
}

// GetAssemblyLevel returns the AssemblyLevel field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsFilter) GetAssemblyLevel() []V1AssemblyDatasetDescriptorsFilterAssemblyLevel {
	if o == nil || o.AssemblyLevel == nil {
		var ret []V1AssemblyDatasetDescriptorsFilterAssemblyLevel
		return ret
	}
	return *o.AssemblyLevel
}

// GetAssemblyLevelOk returns a tuple with the AssemblyLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsFilter) GetAssemblyLevelOk() (*[]V1AssemblyDatasetDescriptorsFilterAssemblyLevel, bool) {
	if o == nil || o.AssemblyLevel == nil {
		return nil, false
	}
	return o.AssemblyLevel, true
}

// HasAssemblyLevel returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsFilter) HasAssemblyLevel() bool {
	if o != nil && o.AssemblyLevel != nil {
		return true
	}

	return false
}

// SetAssemblyLevel gets a reference to the given []V1AssemblyDatasetDescriptorsFilterAssemblyLevel and assigns it to the AssemblyLevel field.
func (o *V1AssemblyDatasetDescriptorsFilter) SetAssemblyLevel(v []V1AssemblyDatasetDescriptorsFilterAssemblyLevel) {
	o.AssemblyLevel = &v
}

// GetFirstReleaseDate returns the FirstReleaseDate field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsFilter) GetFirstReleaseDate() time.Time {
	if o == nil || o.FirstReleaseDate == nil {
		var ret time.Time
		return ret
	}
	return *o.FirstReleaseDate
}

// GetFirstReleaseDateOk returns a tuple with the FirstReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsFilter) GetFirstReleaseDateOk() (*time.Time, bool) {
	if o == nil || o.FirstReleaseDate == nil {
		return nil, false
	}
	return o.FirstReleaseDate, true
}

// HasFirstReleaseDate returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsFilter) HasFirstReleaseDate() bool {
	if o != nil && o.FirstReleaseDate != nil {
		return true
	}

	return false
}

// SetFirstReleaseDate gets a reference to the given time.Time and assigns it to the FirstReleaseDate field.
func (o *V1AssemblyDatasetDescriptorsFilter) SetFirstReleaseDate(v time.Time) {
	o.FirstReleaseDate = &v
}

// GetLastReleaseDate returns the LastReleaseDate field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsFilter) GetLastReleaseDate() time.Time {
	if o == nil || o.LastReleaseDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastReleaseDate
}

// GetLastReleaseDateOk returns a tuple with the LastReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsFilter) GetLastReleaseDateOk() (*time.Time, bool) {
	if o == nil || o.LastReleaseDate == nil {
		return nil, false
	}
	return o.LastReleaseDate, true
}

// HasLastReleaseDate returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsFilter) HasLastReleaseDate() bool {
	if o != nil && o.LastReleaseDate != nil {
		return true
	}

	return false
}

// SetLastReleaseDate gets a reference to the given time.Time and assigns it to the LastReleaseDate field.
func (o *V1AssemblyDatasetDescriptorsFilter) SetLastReleaseDate(v time.Time) {
	o.LastReleaseDate = &v
}

// GetSearchText returns the SearchText field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsFilter) GetSearchText() []string {
	if o == nil || o.SearchText == nil {
		var ret []string
		return ret
	}
	return *o.SearchText
}

// GetSearchTextOk returns a tuple with the SearchText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsFilter) GetSearchTextOk() (*[]string, bool) {
	if o == nil || o.SearchText == nil {
		return nil, false
	}
	return o.SearchText, true
}

// HasSearchText returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsFilter) HasSearchText() bool {
	if o != nil && o.SearchText != nil {
		return true
	}

	return false
}

// SetSearchText gets a reference to the given []string and assigns it to the SearchText field.
func (o *V1AssemblyDatasetDescriptorsFilter) SetSearchText(v []string) {
	o.SearchText = &v
}

func (o V1AssemblyDatasetDescriptorsFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReferenceOnly != nil  {
		toSerialize["reference_only"] = o.ReferenceOnly
	}
	if o.RefseqOnly != nil  {
		toSerialize["refseq_only"] = o.RefseqOnly
	}
	if o.AssemblySource != nil  {
		toSerialize["assembly_source"] = o.AssemblySource
	}
	if o.HasAnnotation != nil  {
		toSerialize["has_annotation"] = o.HasAnnotation
	}
	if o.AssemblyLevel != nil && len(o.GetAssemblyLevel()) > 0  {
		toSerialize["assembly_level"] = o.AssemblyLevel
	}
	if o.FirstReleaseDate != nil  {
		toSerialize["first_release_date"] = o.FirstReleaseDate
	}
	if o.LastReleaseDate != nil  {
		toSerialize["last_release_date"] = o.LastReleaseDate
	}
	if o.SearchText != nil && len(o.GetSearchText()) > 0  {
		toSerialize["search_text"] = o.SearchText
	}
	return json.Marshal(toSerialize)
}

type NullableV1AssemblyDatasetDescriptorsFilter struct {
	value *V1AssemblyDatasetDescriptorsFilter
	isSet bool
}

func (v NullableV1AssemblyDatasetDescriptorsFilter) Get() *V1AssemblyDatasetDescriptorsFilter {
	return v.value
}

func (v *NullableV1AssemblyDatasetDescriptorsFilter) Set(val *V1AssemblyDatasetDescriptorsFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AssemblyDatasetDescriptorsFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AssemblyDatasetDescriptorsFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AssemblyDatasetDescriptorsFilter(val *V1AssemblyDatasetDescriptorsFilter) *NullableV1AssemblyDatasetDescriptorsFilter {
	return &NullableV1AssemblyDatasetDescriptorsFilter{value: val, isSet: true}
}

func (v NullableV1AssemblyDatasetDescriptorsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AssemblyDatasetDescriptorsFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


