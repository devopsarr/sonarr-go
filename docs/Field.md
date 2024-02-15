# Field

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**Unit** | Pointer to **NullableString** |  | [optional] 
**HelpText** | Pointer to **NullableString** |  | [optional] 
**HelpTextWarning** | Pointer to **NullableString** |  | [optional] 
**HelpLink** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Advanced** | Pointer to **bool** |  | [optional] 
**SelectOptions** | Pointer to [**[]SelectOption**](SelectOption.md) |  | [optional] 
**SelectOptionsProviderAction** | Pointer to **NullableString** |  | [optional] 
**Section** | Pointer to **NullableString** |  | [optional] 
**Hidden** | Pointer to **NullableString** |  | [optional] 
**Privacy** | Pointer to [**PrivacyLevel**](PrivacyLevel.md) |  | [optional] 
**Placeholder** | Pointer to **NullableString** |  | [optional] 
**IsFloat** | Pointer to **bool** |  | [optional] 

## Methods

### NewField

`func NewField() *Field`

NewField instantiates a new Field object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldWithDefaults

`func NewFieldWithDefaults() *Field`

NewFieldWithDefaults instantiates a new Field object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *Field) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Field) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Field) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Field) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetName

`func (o *Field) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Field) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Field) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Field) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Field) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Field) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLabel

`func (o *Field) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Field) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Field) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Field) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *Field) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *Field) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetUnit

`func (o *Field) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Field) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Field) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Field) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *Field) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *Field) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetHelpText

`func (o *Field) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *Field) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *Field) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *Field) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### SetHelpTextNil

`func (o *Field) SetHelpTextNil(b bool)`

 SetHelpTextNil sets the value for HelpText to be an explicit nil

### UnsetHelpText
`func (o *Field) UnsetHelpText()`

UnsetHelpText ensures that no value is present for HelpText, not even an explicit nil
### GetHelpTextWarning

`func (o *Field) GetHelpTextWarning() string`

GetHelpTextWarning returns the HelpTextWarning field if non-nil, zero value otherwise.

### GetHelpTextWarningOk

`func (o *Field) GetHelpTextWarningOk() (*string, bool)`

GetHelpTextWarningOk returns a tuple with the HelpTextWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpTextWarning

`func (o *Field) SetHelpTextWarning(v string)`

SetHelpTextWarning sets HelpTextWarning field to given value.

### HasHelpTextWarning

`func (o *Field) HasHelpTextWarning() bool`

HasHelpTextWarning returns a boolean if a field has been set.

### SetHelpTextWarningNil

`func (o *Field) SetHelpTextWarningNil(b bool)`

 SetHelpTextWarningNil sets the value for HelpTextWarning to be an explicit nil

### UnsetHelpTextWarning
`func (o *Field) UnsetHelpTextWarning()`

UnsetHelpTextWarning ensures that no value is present for HelpTextWarning, not even an explicit nil
### GetHelpLink

`func (o *Field) GetHelpLink() string`

GetHelpLink returns the HelpLink field if non-nil, zero value otherwise.

### GetHelpLinkOk

`func (o *Field) GetHelpLinkOk() (*string, bool)`

GetHelpLinkOk returns a tuple with the HelpLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpLink

`func (o *Field) SetHelpLink(v string)`

SetHelpLink sets HelpLink field to given value.

### HasHelpLink

`func (o *Field) HasHelpLink() bool`

HasHelpLink returns a boolean if a field has been set.

### SetHelpLinkNil

`func (o *Field) SetHelpLinkNil(b bool)`

 SetHelpLinkNil sets the value for HelpLink to be an explicit nil

### UnsetHelpLink
`func (o *Field) UnsetHelpLink()`

UnsetHelpLink ensures that no value is present for HelpLink, not even an explicit nil
### GetValue

`func (o *Field) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Field) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Field) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *Field) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Field) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Field) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetType

`func (o *Field) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Field) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Field) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Field) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Field) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Field) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAdvanced

`func (o *Field) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *Field) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *Field) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *Field) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.

### GetSelectOptions

`func (o *Field) GetSelectOptions() []SelectOption`

GetSelectOptions returns the SelectOptions field if non-nil, zero value otherwise.

### GetSelectOptionsOk

`func (o *Field) GetSelectOptionsOk() (*[]SelectOption, bool)`

GetSelectOptionsOk returns a tuple with the SelectOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectOptions

`func (o *Field) SetSelectOptions(v []SelectOption)`

SetSelectOptions sets SelectOptions field to given value.

### HasSelectOptions

`func (o *Field) HasSelectOptions() bool`

HasSelectOptions returns a boolean if a field has been set.

### SetSelectOptionsNil

`func (o *Field) SetSelectOptionsNil(b bool)`

 SetSelectOptionsNil sets the value for SelectOptions to be an explicit nil

### UnsetSelectOptions
`func (o *Field) UnsetSelectOptions()`

UnsetSelectOptions ensures that no value is present for SelectOptions, not even an explicit nil
### GetSelectOptionsProviderAction

`func (o *Field) GetSelectOptionsProviderAction() string`

GetSelectOptionsProviderAction returns the SelectOptionsProviderAction field if non-nil, zero value otherwise.

### GetSelectOptionsProviderActionOk

`func (o *Field) GetSelectOptionsProviderActionOk() (*string, bool)`

GetSelectOptionsProviderActionOk returns a tuple with the SelectOptionsProviderAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectOptionsProviderAction

`func (o *Field) SetSelectOptionsProviderAction(v string)`

SetSelectOptionsProviderAction sets SelectOptionsProviderAction field to given value.

### HasSelectOptionsProviderAction

`func (o *Field) HasSelectOptionsProviderAction() bool`

HasSelectOptionsProviderAction returns a boolean if a field has been set.

### SetSelectOptionsProviderActionNil

`func (o *Field) SetSelectOptionsProviderActionNil(b bool)`

 SetSelectOptionsProviderActionNil sets the value for SelectOptionsProviderAction to be an explicit nil

### UnsetSelectOptionsProviderAction
`func (o *Field) UnsetSelectOptionsProviderAction()`

UnsetSelectOptionsProviderAction ensures that no value is present for SelectOptionsProviderAction, not even an explicit nil
### GetSection

`func (o *Field) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *Field) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *Field) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *Field) HasSection() bool`

HasSection returns a boolean if a field has been set.

### SetSectionNil

`func (o *Field) SetSectionNil(b bool)`

 SetSectionNil sets the value for Section to be an explicit nil

### UnsetSection
`func (o *Field) UnsetSection()`

UnsetSection ensures that no value is present for Section, not even an explicit nil
### GetHidden

`func (o *Field) GetHidden() string`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Field) GetHiddenOk() (*string, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Field) SetHidden(v string)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Field) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHiddenNil

`func (o *Field) SetHiddenNil(b bool)`

 SetHiddenNil sets the value for Hidden to be an explicit nil

### UnsetHidden
`func (o *Field) UnsetHidden()`

UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
### GetPrivacy

`func (o *Field) GetPrivacy() PrivacyLevel`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *Field) GetPrivacyOk() (*PrivacyLevel, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *Field) SetPrivacy(v PrivacyLevel)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *Field) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetPlaceholder

`func (o *Field) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *Field) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *Field) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *Field) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### SetPlaceholderNil

`func (o *Field) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *Field) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
### GetIsFloat

`func (o *Field) GetIsFloat() bool`

GetIsFloat returns the IsFloat field if non-nil, zero value otherwise.

### GetIsFloatOk

`func (o *Field) GetIsFloatOk() (*bool, bool)`

GetIsFloatOk returns a tuple with the IsFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFloat

`func (o *Field) SetIsFloat(v bool)`

SetIsFloat sets IsFloat field to given value.

### HasIsFloat

`func (o *Field) HasIsFloat() bool`

HasIsFloat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


