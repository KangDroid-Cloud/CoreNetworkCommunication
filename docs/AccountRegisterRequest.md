# AccountRegisterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email Address to register. | [optional] 
**Password** | Pointer to **string** | The Password | [optional] 
**NickName** | Pointer to **string** | Nickname to use in our service. | [optional] 

## Methods

### NewAccountRegisterRequest

`func NewAccountRegisterRequest() *AccountRegisterRequest`

NewAccountRegisterRequest instantiates a new AccountRegisterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountRegisterRequestWithDefaults

`func NewAccountRegisterRequestWithDefaults() *AccountRegisterRequest`

NewAccountRegisterRequestWithDefaults instantiates a new AccountRegisterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AccountRegisterRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountRegisterRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountRegisterRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountRegisterRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *AccountRegisterRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AccountRegisterRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AccountRegisterRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AccountRegisterRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetNickName

`func (o *AccountRegisterRequest) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *AccountRegisterRequest) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *AccountRegisterRequest) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *AccountRegisterRequest) HasNickName() bool`

HasNickName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


