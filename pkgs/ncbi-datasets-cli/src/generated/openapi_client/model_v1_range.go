/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/rehydrate/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
)

// V1Range struct for V1Range
type V1Range struct {
	Begin *string `json:"begin,omitempty"`
	End *string `json:"end,omitempty"`
	Orientation *V1Orientation `json:"orientation,omitempty"`
	Order *int32 `json:"order,omitempty"`
	RibosomalSlippage *int32 `json:"ribosomal_slippage,omitempty"`
}

// NewV1Range instantiates a new V1Range object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Range() *V1Range {
	this := V1Range{}
	var orientation V1Orientation = V1ORIENTATION_NONE
	this.Orientation = &orientation
	return &this
}

// NewV1RangeWithDefaults instantiates a new V1Range object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1RangeWithDefaults() *V1Range {
	this := V1Range{}
	var orientation V1Orientation = V1ORIENTATION_NONE
	this.Orientation = &orientation
	return &this
}

// GetBegin returns the Begin field value if set, zero value otherwise.
func (o *V1Range) GetBegin() string {
	if o == nil || o.Begin == nil {
		var ret string
		return ret
	}
	return *o.Begin
}

// GetBeginOk returns a tuple with the Begin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Range) GetBeginOk() (*string, bool) {
	if o == nil || o.Begin == nil {
		return nil, false
	}
	return o.Begin, true
}

// HasBegin returns a boolean if a field has been set.
func (o *V1Range) HasBegin() bool {
	if o != nil && o.Begin != nil {
		return true
	}

	return false
}

// SetBegin gets a reference to the given string and assigns it to the Begin field.
func (o *V1Range) SetBegin(v string) {
	o.Begin = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *V1Range) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Range) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *V1Range) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *V1Range) SetEnd(v string) {
	o.End = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *V1Range) GetOrientation() V1Orientation {
	if o == nil || o.Orientation == nil {
		var ret V1Orientation
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Range) GetOrientationOk() (*V1Orientation, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *V1Range) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given V1Orientation and assigns it to the Orientation field.
func (o *V1Range) SetOrientation(v V1Orientation) {
	o.Orientation = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *V1Range) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Range) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *V1Range) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *V1Range) SetOrder(v int32) {
	o.Order = &v
}

// GetRibosomalSlippage returns the RibosomalSlippage field value if set, zero value otherwise.
func (o *V1Range) GetRibosomalSlippage() int32 {
	if o == nil || o.RibosomalSlippage == nil {
		var ret int32
		return ret
	}
	return *o.RibosomalSlippage
}

// GetRibosomalSlippageOk returns a tuple with the RibosomalSlippage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Range) GetRibosomalSlippageOk() (*int32, bool) {
	if o == nil || o.RibosomalSlippage == nil {
		return nil, false
	}
	return o.RibosomalSlippage, true
}

// HasRibosomalSlippage returns a boolean if a field has been set.
func (o *V1Range) HasRibosomalSlippage() bool {
	if o != nil && o.RibosomalSlippage != nil {
		return true
	}

	return false
}

// SetRibosomalSlippage gets a reference to the given int32 and assigns it to the RibosomalSlippage field.
func (o *V1Range) SetRibosomalSlippage(v int32) {
	o.RibosomalSlippage = &v
}

func (o V1Range) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Begin != nil  {
		toSerialize["begin"] = o.Begin
	}
	if o.End != nil  {
		toSerialize["end"] = o.End
	}
	if o.Orientation != nil  {
		toSerialize["orientation"] = o.Orientation
	}
	if o.Order != nil  {
		toSerialize["order"] = o.Order
	}
	if o.RibosomalSlippage != nil  {
		toSerialize["ribosomal_slippage"] = o.RibosomalSlippage
	}
	return json.Marshal(toSerialize)
}

type NullableV1Range struct {
	value *V1Range
	isSet bool
}

func (v NullableV1Range) Get() *V1Range {
	return v.value
}

func (v *NullableV1Range) Set(val *V1Range) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Range) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Range) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Range(val *V1Range) *NullableV1Range {
	return &NullableV1Range{value: val, isSet: true}
}

func (v NullableV1Range) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Range) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


