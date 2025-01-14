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

// V1reportsOrganism struct for V1reportsOrganism
type V1reportsOrganism struct {
	TaxId *int32 `json:"tax_id,omitempty"`
	SciName *string `json:"sci_name,omitempty"`
	OrganismName *string `json:"organism_name,omitempty"`
	CommonName *string `json:"common_name,omitempty"`
	Lineage *[]V1reportsLineageOrganism `json:"lineage,omitempty"`
	Strain *string `json:"strain,omitempty"`
	PangolinClassification *string `json:"pangolin_classification,omitempty"`
}

// NewV1reportsOrganism instantiates a new V1reportsOrganism object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsOrganism() *V1reportsOrganism {
	this := V1reportsOrganism{}
	return &this
}

// NewV1reportsOrganismWithDefaults instantiates a new V1reportsOrganism object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsOrganismWithDefaults() *V1reportsOrganism {
	this := V1reportsOrganism{}
	return &this
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *V1reportsOrganism) GetTaxId() int32 {
	if o == nil || o.TaxId == nil {
		var ret int32
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsOrganism) GetTaxIdOk() (*int32, bool) {
	if o == nil || o.TaxId == nil {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *V1reportsOrganism) HasTaxId() bool {
	if o != nil && o.TaxId != nil {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given int32 and assigns it to the TaxId field.
func (o *V1reportsOrganism) SetTaxId(v int32) {
	o.TaxId = &v
}

// GetSciName returns the SciName field value if set, zero value otherwise.
func (o *V1reportsOrganism) GetSciName() string {
	if o == nil || o.SciName == nil {
		var ret string
		return ret
	}
	return *o.SciName
}

// GetSciNameOk returns a tuple with the SciName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsOrganism) GetSciNameOk() (*string, bool) {
	if o == nil || o.SciName == nil {
		return nil, false
	}
	return o.SciName, true
}

// HasSciName returns a boolean if a field has been set.
func (o *V1reportsOrganism) HasSciName() bool {
	if o != nil && o.SciName != nil {
		return true
	}

	return false
}

// SetSciName gets a reference to the given string and assigns it to the SciName field.
func (o *V1reportsOrganism) SetSciName(v string) {
	o.SciName = &v
}

// GetOrganismName returns the OrganismName field value if set, zero value otherwise.
func (o *V1reportsOrganism) GetOrganismName() string {
	if o == nil || o.OrganismName == nil {
		var ret string
		return ret
	}
	return *o.OrganismName
}

// GetOrganismNameOk returns a tuple with the OrganismName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsOrganism) GetOrganismNameOk() (*string, bool) {
	if o == nil || o.OrganismName == nil {
		return nil, false
	}
	return o.OrganismName, true
}

// HasOrganismName returns a boolean if a field has been set.
func (o *V1reportsOrganism) HasOrganismName() bool {
	if o != nil && o.OrganismName != nil {
		return true
	}

	return false
}

// SetOrganismName gets a reference to the given string and assigns it to the OrganismName field.
func (o *V1reportsOrganism) SetOrganismName(v string) {
	o.OrganismName = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *V1reportsOrganism) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsOrganism) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *V1reportsOrganism) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *V1reportsOrganism) SetCommonName(v string) {
	o.CommonName = &v
}

// GetLineage returns the Lineage field value if set, zero value otherwise.
func (o *V1reportsOrganism) GetLineage() []V1reportsLineageOrganism {
	if o == nil || o.Lineage == nil {
		var ret []V1reportsLineageOrganism
		return ret
	}
	return *o.Lineage
}

// GetLineageOk returns a tuple with the Lineage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsOrganism) GetLineageOk() (*[]V1reportsLineageOrganism, bool) {
	if o == nil || o.Lineage == nil {
		return nil, false
	}
	return o.Lineage, true
}

// HasLineage returns a boolean if a field has been set.
func (o *V1reportsOrganism) HasLineage() bool {
	if o != nil && o.Lineage != nil {
		return true
	}

	return false
}

// SetLineage gets a reference to the given []V1reportsLineageOrganism and assigns it to the Lineage field.
func (o *V1reportsOrganism) SetLineage(v []V1reportsLineageOrganism) {
	o.Lineage = &v
}

// GetStrain returns the Strain field value if set, zero value otherwise.
func (o *V1reportsOrganism) GetStrain() string {
	if o == nil || o.Strain == nil {
		var ret string
		return ret
	}
	return *o.Strain
}

// GetStrainOk returns a tuple with the Strain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsOrganism) GetStrainOk() (*string, bool) {
	if o == nil || o.Strain == nil {
		return nil, false
	}
	return o.Strain, true
}

// HasStrain returns a boolean if a field has been set.
func (o *V1reportsOrganism) HasStrain() bool {
	if o != nil && o.Strain != nil {
		return true
	}

	return false
}

// SetStrain gets a reference to the given string and assigns it to the Strain field.
func (o *V1reportsOrganism) SetStrain(v string) {
	o.Strain = &v
}

// GetPangolinClassification returns the PangolinClassification field value if set, zero value otherwise.
func (o *V1reportsOrganism) GetPangolinClassification() string {
	if o == nil || o.PangolinClassification == nil {
		var ret string
		return ret
	}
	return *o.PangolinClassification
}

// GetPangolinClassificationOk returns a tuple with the PangolinClassification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsOrganism) GetPangolinClassificationOk() (*string, bool) {
	if o == nil || o.PangolinClassification == nil {
		return nil, false
	}
	return o.PangolinClassification, true
}

// HasPangolinClassification returns a boolean if a field has been set.
func (o *V1reportsOrganism) HasPangolinClassification() bool {
	if o != nil && o.PangolinClassification != nil {
		return true
	}

	return false
}

// SetPangolinClassification gets a reference to the given string and assigns it to the PangolinClassification field.
func (o *V1reportsOrganism) SetPangolinClassification(v string) {
	o.PangolinClassification = &v
}

func (o V1reportsOrganism) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaxId != nil  {
		toSerialize["tax_id"] = o.TaxId
	}
	if o.SciName != nil  {
		toSerialize["sci_name"] = o.SciName
	}
	if o.OrganismName != nil  {
		toSerialize["organism_name"] = o.OrganismName
	}
	if o.CommonName != nil  {
		toSerialize["common_name"] = o.CommonName
	}
	if o.Lineage != nil && len(o.GetLineage()) > 0  {
		toSerialize["lineage"] = o.Lineage
	}
	if o.Strain != nil  {
		toSerialize["strain"] = o.Strain
	}
	if o.PangolinClassification != nil  {
		toSerialize["pangolin_classification"] = o.PangolinClassification
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsOrganism struct {
	value *V1reportsOrganism
	isSet bool
}

func (v NullableV1reportsOrganism) Get() *V1reportsOrganism {
	return v.value
}

func (v *NullableV1reportsOrganism) Set(val *V1reportsOrganism) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsOrganism) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsOrganism) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsOrganism(val *V1reportsOrganism) *NullableV1reportsOrganism {
	return &NullableV1reportsOrganism{value: val, isSet: true}
}

func (v NullableV1reportsOrganism) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsOrganism) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


