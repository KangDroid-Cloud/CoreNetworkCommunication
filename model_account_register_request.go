/*
KangDroid Cloud OpenAPI Description

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AccountRegisterRequest struct for AccountRegisterRequest
type AccountRegisterRequest struct {
	// Email Address to register.
	Email *string `json:"email,omitempty"`
	// The Password
	Password *string `json:"password,omitempty"`
	// Nickname to use in our service.
	NickName *string `json:"nickName,omitempty"`
}

// NewAccountRegisterRequest instantiates a new AccountRegisterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountRegisterRequest() *AccountRegisterRequest {
	this := AccountRegisterRequest{}
	return &this
}

// NewAccountRegisterRequestWithDefaults instantiates a new AccountRegisterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountRegisterRequestWithDefaults() *AccountRegisterRequest {
	this := AccountRegisterRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AccountRegisterRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRegisterRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AccountRegisterRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AccountRegisterRequest) SetEmail(v string) {
	o.Email = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AccountRegisterRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRegisterRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AccountRegisterRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AccountRegisterRequest) SetPassword(v string) {
	o.Password = &v
}

// GetNickName returns the NickName field value if set, zero value otherwise.
func (o *AccountRegisterRequest) GetNickName() string {
	if o == nil || o.NickName == nil {
		var ret string
		return ret
	}
	return *o.NickName
}

// GetNickNameOk returns a tuple with the NickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRegisterRequest) GetNickNameOk() (*string, bool) {
	if o == nil || o.NickName == nil {
		return nil, false
	}
	return o.NickName, true
}

// HasNickName returns a boolean if a field has been set.
func (o *AccountRegisterRequest) HasNickName() bool {
	if o != nil && o.NickName != nil {
		return true
	}

	return false
}

// SetNickName gets a reference to the given string and assigns it to the NickName field.
func (o *AccountRegisterRequest) SetNickName(v string) {
	o.NickName = &v
}

func (o AccountRegisterRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.NickName != nil {
		toSerialize["nickName"] = o.NickName
	}
	return json.Marshal(toSerialize)
}

type NullableAccountRegisterRequest struct {
	value *AccountRegisterRequest
	isSet bool
}

func (v NullableAccountRegisterRequest) Get() *AccountRegisterRequest {
	return v.value
}

func (v *NullableAccountRegisterRequest) Set(val *AccountRegisterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountRegisterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountRegisterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountRegisterRequest(val *AccountRegisterRequest) *NullableAccountRegisterRequest {
	return &NullableAccountRegisterRequest{value: val, isSet: true}
}

func (v NullableAccountRegisterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountRegisterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

