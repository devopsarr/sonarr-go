# HostConfigResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**BindAddress** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**SslPort** | Pointer to **int32** |  | [optional] 
**EnableSsl** | Pointer to **bool** |  | [optional] 
**LaunchBrowser** | Pointer to **bool** |  | [optional] 
**AuthenticationMethod** | Pointer to [**AuthenticationType**](AuthenticationType.md) |  | [optional] 
**AuthenticationRequired** | Pointer to [**AuthenticationRequiredType**](AuthenticationRequiredType.md) |  | [optional] 
**AnalyticsEnabled** | Pointer to **bool** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**PasswordConfirmation** | Pointer to **NullableString** |  | [optional] 
**LogLevel** | Pointer to **NullableString** |  | [optional] 
**ConsoleLogLevel** | Pointer to **NullableString** |  | [optional] 
**Branch** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | Pointer to **NullableString** |  | [optional] 
**SslCertPath** | Pointer to **NullableString** |  | [optional] 
**SslCertPassword** | Pointer to **NullableString** |  | [optional] 
**UrlBase** | Pointer to **NullableString** |  | [optional] 
**InstanceName** | Pointer to **NullableString** |  | [optional] 
**ApplicationUrl** | Pointer to **NullableString** |  | [optional] 
**UpdateAutomatically** | Pointer to **bool** |  | [optional] 
**UpdateMechanism** | Pointer to [**UpdateMechanism**](UpdateMechanism.md) |  | [optional] 
**UpdateScriptPath** | Pointer to **NullableString** |  | [optional] 
**ProxyEnabled** | Pointer to **bool** |  | [optional] 
**ProxyType** | Pointer to [**ProxyType**](ProxyType.md) |  | [optional] 
**ProxyHostname** | Pointer to **NullableString** |  | [optional] 
**ProxyPort** | Pointer to **int32** |  | [optional] 
**ProxyUsername** | Pointer to **NullableString** |  | [optional] 
**ProxyPassword** | Pointer to **NullableString** |  | [optional] 
**ProxyBypassFilter** | Pointer to **NullableString** |  | [optional] 
**ProxyBypassLocalAddresses** | Pointer to **bool** |  | [optional] 
**CertificateValidation** | Pointer to [**CertificateValidationType**](CertificateValidationType.md) |  | [optional] 
**BackupFolder** | Pointer to **NullableString** |  | [optional] 
**BackupInterval** | Pointer to **int32** |  | [optional] 
**BackupRetention** | Pointer to **int32** |  | [optional] 

## Methods

### NewHostConfigResource

`func NewHostConfigResource() *HostConfigResource`

NewHostConfigResource instantiates a new HostConfigResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostConfigResourceWithDefaults

`func NewHostConfigResourceWithDefaults() *HostConfigResource`

NewHostConfigResourceWithDefaults instantiates a new HostConfigResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HostConfigResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostConfigResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostConfigResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HostConfigResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBindAddress

`func (o *HostConfigResource) GetBindAddress() string`

GetBindAddress returns the BindAddress field if non-nil, zero value otherwise.

### GetBindAddressOk

`func (o *HostConfigResource) GetBindAddressOk() (*string, bool)`

GetBindAddressOk returns a tuple with the BindAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindAddress

`func (o *HostConfigResource) SetBindAddress(v string)`

SetBindAddress sets BindAddress field to given value.

### HasBindAddress

`func (o *HostConfigResource) HasBindAddress() bool`

HasBindAddress returns a boolean if a field has been set.

### SetBindAddressNil

`func (o *HostConfigResource) SetBindAddressNil(b bool)`

 SetBindAddressNil sets the value for BindAddress to be an explicit nil

### UnsetBindAddress
`func (o *HostConfigResource) UnsetBindAddress()`

UnsetBindAddress ensures that no value is present for BindAddress, not even an explicit nil
### GetPort

`func (o *HostConfigResource) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HostConfigResource) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HostConfigResource) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HostConfigResource) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSslPort

`func (o *HostConfigResource) GetSslPort() int32`

GetSslPort returns the SslPort field if non-nil, zero value otherwise.

### GetSslPortOk

`func (o *HostConfigResource) GetSslPortOk() (*int32, bool)`

GetSslPortOk returns a tuple with the SslPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslPort

`func (o *HostConfigResource) SetSslPort(v int32)`

SetSslPort sets SslPort field to given value.

### HasSslPort

`func (o *HostConfigResource) HasSslPort() bool`

HasSslPort returns a boolean if a field has been set.

### GetEnableSsl

`func (o *HostConfigResource) GetEnableSsl() bool`

GetEnableSsl returns the EnableSsl field if non-nil, zero value otherwise.

### GetEnableSslOk

`func (o *HostConfigResource) GetEnableSslOk() (*bool, bool)`

GetEnableSslOk returns a tuple with the EnableSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSsl

`func (o *HostConfigResource) SetEnableSsl(v bool)`

SetEnableSsl sets EnableSsl field to given value.

### HasEnableSsl

`func (o *HostConfigResource) HasEnableSsl() bool`

HasEnableSsl returns a boolean if a field has been set.

### GetLaunchBrowser

`func (o *HostConfigResource) GetLaunchBrowser() bool`

GetLaunchBrowser returns the LaunchBrowser field if non-nil, zero value otherwise.

### GetLaunchBrowserOk

`func (o *HostConfigResource) GetLaunchBrowserOk() (*bool, bool)`

GetLaunchBrowserOk returns a tuple with the LaunchBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchBrowser

`func (o *HostConfigResource) SetLaunchBrowser(v bool)`

SetLaunchBrowser sets LaunchBrowser field to given value.

### HasLaunchBrowser

`func (o *HostConfigResource) HasLaunchBrowser() bool`

HasLaunchBrowser returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *HostConfigResource) GetAuthenticationMethod() AuthenticationType`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *HostConfigResource) GetAuthenticationMethodOk() (*AuthenticationType, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *HostConfigResource) SetAuthenticationMethod(v AuthenticationType)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *HostConfigResource) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetAuthenticationRequired

`func (o *HostConfigResource) GetAuthenticationRequired() AuthenticationRequiredType`

GetAuthenticationRequired returns the AuthenticationRequired field if non-nil, zero value otherwise.

### GetAuthenticationRequiredOk

`func (o *HostConfigResource) GetAuthenticationRequiredOk() (*AuthenticationRequiredType, bool)`

GetAuthenticationRequiredOk returns a tuple with the AuthenticationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationRequired

`func (o *HostConfigResource) SetAuthenticationRequired(v AuthenticationRequiredType)`

SetAuthenticationRequired sets AuthenticationRequired field to given value.

### HasAuthenticationRequired

`func (o *HostConfigResource) HasAuthenticationRequired() bool`

HasAuthenticationRequired returns a boolean if a field has been set.

### GetAnalyticsEnabled

`func (o *HostConfigResource) GetAnalyticsEnabled() bool`

GetAnalyticsEnabled returns the AnalyticsEnabled field if non-nil, zero value otherwise.

### GetAnalyticsEnabledOk

`func (o *HostConfigResource) GetAnalyticsEnabledOk() (*bool, bool)`

GetAnalyticsEnabledOk returns a tuple with the AnalyticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsEnabled

`func (o *HostConfigResource) SetAnalyticsEnabled(v bool)`

SetAnalyticsEnabled sets AnalyticsEnabled field to given value.

### HasAnalyticsEnabled

`func (o *HostConfigResource) HasAnalyticsEnabled() bool`

HasAnalyticsEnabled returns a boolean if a field has been set.

### GetUsername

`func (o *HostConfigResource) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HostConfigResource) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HostConfigResource) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HostConfigResource) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *HostConfigResource) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *HostConfigResource) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *HostConfigResource) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HostConfigResource) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HostConfigResource) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HostConfigResource) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *HostConfigResource) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *HostConfigResource) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordConfirmation

`func (o *HostConfigResource) GetPasswordConfirmation() string`

GetPasswordConfirmation returns the PasswordConfirmation field if non-nil, zero value otherwise.

### GetPasswordConfirmationOk

`func (o *HostConfigResource) GetPasswordConfirmationOk() (*string, bool)`

GetPasswordConfirmationOk returns a tuple with the PasswordConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordConfirmation

`func (o *HostConfigResource) SetPasswordConfirmation(v string)`

SetPasswordConfirmation sets PasswordConfirmation field to given value.

### HasPasswordConfirmation

`func (o *HostConfigResource) HasPasswordConfirmation() bool`

HasPasswordConfirmation returns a boolean if a field has been set.

### SetPasswordConfirmationNil

`func (o *HostConfigResource) SetPasswordConfirmationNil(b bool)`

 SetPasswordConfirmationNil sets the value for PasswordConfirmation to be an explicit nil

### UnsetPasswordConfirmation
`func (o *HostConfigResource) UnsetPasswordConfirmation()`

UnsetPasswordConfirmation ensures that no value is present for PasswordConfirmation, not even an explicit nil
### GetLogLevel

`func (o *HostConfigResource) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *HostConfigResource) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *HostConfigResource) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *HostConfigResource) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### SetLogLevelNil

`func (o *HostConfigResource) SetLogLevelNil(b bool)`

 SetLogLevelNil sets the value for LogLevel to be an explicit nil

### UnsetLogLevel
`func (o *HostConfigResource) UnsetLogLevel()`

UnsetLogLevel ensures that no value is present for LogLevel, not even an explicit nil
### GetConsoleLogLevel

`func (o *HostConfigResource) GetConsoleLogLevel() string`

GetConsoleLogLevel returns the ConsoleLogLevel field if non-nil, zero value otherwise.

### GetConsoleLogLevelOk

`func (o *HostConfigResource) GetConsoleLogLevelOk() (*string, bool)`

GetConsoleLogLevelOk returns a tuple with the ConsoleLogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleLogLevel

`func (o *HostConfigResource) SetConsoleLogLevel(v string)`

SetConsoleLogLevel sets ConsoleLogLevel field to given value.

### HasConsoleLogLevel

`func (o *HostConfigResource) HasConsoleLogLevel() bool`

HasConsoleLogLevel returns a boolean if a field has been set.

### SetConsoleLogLevelNil

`func (o *HostConfigResource) SetConsoleLogLevelNil(b bool)`

 SetConsoleLogLevelNil sets the value for ConsoleLogLevel to be an explicit nil

### UnsetConsoleLogLevel
`func (o *HostConfigResource) UnsetConsoleLogLevel()`

UnsetConsoleLogLevel ensures that no value is present for ConsoleLogLevel, not even an explicit nil
### GetBranch

`func (o *HostConfigResource) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *HostConfigResource) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *HostConfigResource) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *HostConfigResource) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### SetBranchNil

`func (o *HostConfigResource) SetBranchNil(b bool)`

 SetBranchNil sets the value for Branch to be an explicit nil

### UnsetBranch
`func (o *HostConfigResource) UnsetBranch()`

UnsetBranch ensures that no value is present for Branch, not even an explicit nil
### GetApiKey

`func (o *HostConfigResource) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *HostConfigResource) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *HostConfigResource) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *HostConfigResource) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *HostConfigResource) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *HostConfigResource) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetSslCertPath

`func (o *HostConfigResource) GetSslCertPath() string`

GetSslCertPath returns the SslCertPath field if non-nil, zero value otherwise.

### GetSslCertPathOk

`func (o *HostConfigResource) GetSslCertPathOk() (*string, bool)`

GetSslCertPathOk returns a tuple with the SslCertPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertPath

`func (o *HostConfigResource) SetSslCertPath(v string)`

SetSslCertPath sets SslCertPath field to given value.

### HasSslCertPath

`func (o *HostConfigResource) HasSslCertPath() bool`

HasSslCertPath returns a boolean if a field has been set.

### SetSslCertPathNil

`func (o *HostConfigResource) SetSslCertPathNil(b bool)`

 SetSslCertPathNil sets the value for SslCertPath to be an explicit nil

### UnsetSslCertPath
`func (o *HostConfigResource) UnsetSslCertPath()`

UnsetSslCertPath ensures that no value is present for SslCertPath, not even an explicit nil
### GetSslCertPassword

`func (o *HostConfigResource) GetSslCertPassword() string`

GetSslCertPassword returns the SslCertPassword field if non-nil, zero value otherwise.

### GetSslCertPasswordOk

`func (o *HostConfigResource) GetSslCertPasswordOk() (*string, bool)`

GetSslCertPasswordOk returns a tuple with the SslCertPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertPassword

`func (o *HostConfigResource) SetSslCertPassword(v string)`

SetSslCertPassword sets SslCertPassword field to given value.

### HasSslCertPassword

`func (o *HostConfigResource) HasSslCertPassword() bool`

HasSslCertPassword returns a boolean if a field has been set.

### SetSslCertPasswordNil

`func (o *HostConfigResource) SetSslCertPasswordNil(b bool)`

 SetSslCertPasswordNil sets the value for SslCertPassword to be an explicit nil

### UnsetSslCertPassword
`func (o *HostConfigResource) UnsetSslCertPassword()`

UnsetSslCertPassword ensures that no value is present for SslCertPassword, not even an explicit nil
### GetUrlBase

`func (o *HostConfigResource) GetUrlBase() string`

GetUrlBase returns the UrlBase field if non-nil, zero value otherwise.

### GetUrlBaseOk

`func (o *HostConfigResource) GetUrlBaseOk() (*string, bool)`

GetUrlBaseOk returns a tuple with the UrlBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlBase

`func (o *HostConfigResource) SetUrlBase(v string)`

SetUrlBase sets UrlBase field to given value.

### HasUrlBase

`func (o *HostConfigResource) HasUrlBase() bool`

HasUrlBase returns a boolean if a field has been set.

### SetUrlBaseNil

`func (o *HostConfigResource) SetUrlBaseNil(b bool)`

 SetUrlBaseNil sets the value for UrlBase to be an explicit nil

### UnsetUrlBase
`func (o *HostConfigResource) UnsetUrlBase()`

UnsetUrlBase ensures that no value is present for UrlBase, not even an explicit nil
### GetInstanceName

`func (o *HostConfigResource) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *HostConfigResource) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *HostConfigResource) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *HostConfigResource) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### SetInstanceNameNil

`func (o *HostConfigResource) SetInstanceNameNil(b bool)`

 SetInstanceNameNil sets the value for InstanceName to be an explicit nil

### UnsetInstanceName
`func (o *HostConfigResource) UnsetInstanceName()`

UnsetInstanceName ensures that no value is present for InstanceName, not even an explicit nil
### GetApplicationUrl

`func (o *HostConfigResource) GetApplicationUrl() string`

GetApplicationUrl returns the ApplicationUrl field if non-nil, zero value otherwise.

### GetApplicationUrlOk

`func (o *HostConfigResource) GetApplicationUrlOk() (*string, bool)`

GetApplicationUrlOk returns a tuple with the ApplicationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUrl

`func (o *HostConfigResource) SetApplicationUrl(v string)`

SetApplicationUrl sets ApplicationUrl field to given value.

### HasApplicationUrl

`func (o *HostConfigResource) HasApplicationUrl() bool`

HasApplicationUrl returns a boolean if a field has been set.

### SetApplicationUrlNil

`func (o *HostConfigResource) SetApplicationUrlNil(b bool)`

 SetApplicationUrlNil sets the value for ApplicationUrl to be an explicit nil

### UnsetApplicationUrl
`func (o *HostConfigResource) UnsetApplicationUrl()`

UnsetApplicationUrl ensures that no value is present for ApplicationUrl, not even an explicit nil
### GetUpdateAutomatically

`func (o *HostConfigResource) GetUpdateAutomatically() bool`

GetUpdateAutomatically returns the UpdateAutomatically field if non-nil, zero value otherwise.

### GetUpdateAutomaticallyOk

`func (o *HostConfigResource) GetUpdateAutomaticallyOk() (*bool, bool)`

GetUpdateAutomaticallyOk returns a tuple with the UpdateAutomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAutomatically

`func (o *HostConfigResource) SetUpdateAutomatically(v bool)`

SetUpdateAutomatically sets UpdateAutomatically field to given value.

### HasUpdateAutomatically

`func (o *HostConfigResource) HasUpdateAutomatically() bool`

HasUpdateAutomatically returns a boolean if a field has been set.

### GetUpdateMechanism

`func (o *HostConfigResource) GetUpdateMechanism() UpdateMechanism`

GetUpdateMechanism returns the UpdateMechanism field if non-nil, zero value otherwise.

### GetUpdateMechanismOk

`func (o *HostConfigResource) GetUpdateMechanismOk() (*UpdateMechanism, bool)`

GetUpdateMechanismOk returns a tuple with the UpdateMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMechanism

`func (o *HostConfigResource) SetUpdateMechanism(v UpdateMechanism)`

SetUpdateMechanism sets UpdateMechanism field to given value.

### HasUpdateMechanism

`func (o *HostConfigResource) HasUpdateMechanism() bool`

HasUpdateMechanism returns a boolean if a field has been set.

### GetUpdateScriptPath

`func (o *HostConfigResource) GetUpdateScriptPath() string`

GetUpdateScriptPath returns the UpdateScriptPath field if non-nil, zero value otherwise.

### GetUpdateScriptPathOk

`func (o *HostConfigResource) GetUpdateScriptPathOk() (*string, bool)`

GetUpdateScriptPathOk returns a tuple with the UpdateScriptPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateScriptPath

`func (o *HostConfigResource) SetUpdateScriptPath(v string)`

SetUpdateScriptPath sets UpdateScriptPath field to given value.

### HasUpdateScriptPath

`func (o *HostConfigResource) HasUpdateScriptPath() bool`

HasUpdateScriptPath returns a boolean if a field has been set.

### SetUpdateScriptPathNil

`func (o *HostConfigResource) SetUpdateScriptPathNil(b bool)`

 SetUpdateScriptPathNil sets the value for UpdateScriptPath to be an explicit nil

### UnsetUpdateScriptPath
`func (o *HostConfigResource) UnsetUpdateScriptPath()`

UnsetUpdateScriptPath ensures that no value is present for UpdateScriptPath, not even an explicit nil
### GetProxyEnabled

`func (o *HostConfigResource) GetProxyEnabled() bool`

GetProxyEnabled returns the ProxyEnabled field if non-nil, zero value otherwise.

### GetProxyEnabledOk

`func (o *HostConfigResource) GetProxyEnabledOk() (*bool, bool)`

GetProxyEnabledOk returns a tuple with the ProxyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyEnabled

`func (o *HostConfigResource) SetProxyEnabled(v bool)`

SetProxyEnabled sets ProxyEnabled field to given value.

### HasProxyEnabled

`func (o *HostConfigResource) HasProxyEnabled() bool`

HasProxyEnabled returns a boolean if a field has been set.

### GetProxyType

`func (o *HostConfigResource) GetProxyType() ProxyType`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *HostConfigResource) GetProxyTypeOk() (*ProxyType, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *HostConfigResource) SetProxyType(v ProxyType)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *HostConfigResource) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetProxyHostname

`func (o *HostConfigResource) GetProxyHostname() string`

GetProxyHostname returns the ProxyHostname field if non-nil, zero value otherwise.

### GetProxyHostnameOk

`func (o *HostConfigResource) GetProxyHostnameOk() (*string, bool)`

GetProxyHostnameOk returns a tuple with the ProxyHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHostname

`func (o *HostConfigResource) SetProxyHostname(v string)`

SetProxyHostname sets ProxyHostname field to given value.

### HasProxyHostname

`func (o *HostConfigResource) HasProxyHostname() bool`

HasProxyHostname returns a boolean if a field has been set.

### SetProxyHostnameNil

`func (o *HostConfigResource) SetProxyHostnameNil(b bool)`

 SetProxyHostnameNil sets the value for ProxyHostname to be an explicit nil

### UnsetProxyHostname
`func (o *HostConfigResource) UnsetProxyHostname()`

UnsetProxyHostname ensures that no value is present for ProxyHostname, not even an explicit nil
### GetProxyPort

`func (o *HostConfigResource) GetProxyPort() int32`

GetProxyPort returns the ProxyPort field if non-nil, zero value otherwise.

### GetProxyPortOk

`func (o *HostConfigResource) GetProxyPortOk() (*int32, bool)`

GetProxyPortOk returns a tuple with the ProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPort

`func (o *HostConfigResource) SetProxyPort(v int32)`

SetProxyPort sets ProxyPort field to given value.

### HasProxyPort

`func (o *HostConfigResource) HasProxyPort() bool`

HasProxyPort returns a boolean if a field has been set.

### GetProxyUsername

`func (o *HostConfigResource) GetProxyUsername() string`

GetProxyUsername returns the ProxyUsername field if non-nil, zero value otherwise.

### GetProxyUsernameOk

`func (o *HostConfigResource) GetProxyUsernameOk() (*string, bool)`

GetProxyUsernameOk returns a tuple with the ProxyUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUsername

`func (o *HostConfigResource) SetProxyUsername(v string)`

SetProxyUsername sets ProxyUsername field to given value.

### HasProxyUsername

`func (o *HostConfigResource) HasProxyUsername() bool`

HasProxyUsername returns a boolean if a field has been set.

### SetProxyUsernameNil

`func (o *HostConfigResource) SetProxyUsernameNil(b bool)`

 SetProxyUsernameNil sets the value for ProxyUsername to be an explicit nil

### UnsetProxyUsername
`func (o *HostConfigResource) UnsetProxyUsername()`

UnsetProxyUsername ensures that no value is present for ProxyUsername, not even an explicit nil
### GetProxyPassword

`func (o *HostConfigResource) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *HostConfigResource) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *HostConfigResource) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *HostConfigResource) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### SetProxyPasswordNil

`func (o *HostConfigResource) SetProxyPasswordNil(b bool)`

 SetProxyPasswordNil sets the value for ProxyPassword to be an explicit nil

### UnsetProxyPassword
`func (o *HostConfigResource) UnsetProxyPassword()`

UnsetProxyPassword ensures that no value is present for ProxyPassword, not even an explicit nil
### GetProxyBypassFilter

`func (o *HostConfigResource) GetProxyBypassFilter() string`

GetProxyBypassFilter returns the ProxyBypassFilter field if non-nil, zero value otherwise.

### GetProxyBypassFilterOk

`func (o *HostConfigResource) GetProxyBypassFilterOk() (*string, bool)`

GetProxyBypassFilterOk returns a tuple with the ProxyBypassFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyBypassFilter

`func (o *HostConfigResource) SetProxyBypassFilter(v string)`

SetProxyBypassFilter sets ProxyBypassFilter field to given value.

### HasProxyBypassFilter

`func (o *HostConfigResource) HasProxyBypassFilter() bool`

HasProxyBypassFilter returns a boolean if a field has been set.

### SetProxyBypassFilterNil

`func (o *HostConfigResource) SetProxyBypassFilterNil(b bool)`

 SetProxyBypassFilterNil sets the value for ProxyBypassFilter to be an explicit nil

### UnsetProxyBypassFilter
`func (o *HostConfigResource) UnsetProxyBypassFilter()`

UnsetProxyBypassFilter ensures that no value is present for ProxyBypassFilter, not even an explicit nil
### GetProxyBypassLocalAddresses

`func (o *HostConfigResource) GetProxyBypassLocalAddresses() bool`

GetProxyBypassLocalAddresses returns the ProxyBypassLocalAddresses field if non-nil, zero value otherwise.

### GetProxyBypassLocalAddressesOk

`func (o *HostConfigResource) GetProxyBypassLocalAddressesOk() (*bool, bool)`

GetProxyBypassLocalAddressesOk returns a tuple with the ProxyBypassLocalAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyBypassLocalAddresses

`func (o *HostConfigResource) SetProxyBypassLocalAddresses(v bool)`

SetProxyBypassLocalAddresses sets ProxyBypassLocalAddresses field to given value.

### HasProxyBypassLocalAddresses

`func (o *HostConfigResource) HasProxyBypassLocalAddresses() bool`

HasProxyBypassLocalAddresses returns a boolean if a field has been set.

### GetCertificateValidation

`func (o *HostConfigResource) GetCertificateValidation() CertificateValidationType`

GetCertificateValidation returns the CertificateValidation field if non-nil, zero value otherwise.

### GetCertificateValidationOk

`func (o *HostConfigResource) GetCertificateValidationOk() (*CertificateValidationType, bool)`

GetCertificateValidationOk returns a tuple with the CertificateValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateValidation

`func (o *HostConfigResource) SetCertificateValidation(v CertificateValidationType)`

SetCertificateValidation sets CertificateValidation field to given value.

### HasCertificateValidation

`func (o *HostConfigResource) HasCertificateValidation() bool`

HasCertificateValidation returns a boolean if a field has been set.

### GetBackupFolder

`func (o *HostConfigResource) GetBackupFolder() string`

GetBackupFolder returns the BackupFolder field if non-nil, zero value otherwise.

### GetBackupFolderOk

`func (o *HostConfigResource) GetBackupFolderOk() (*string, bool)`

GetBackupFolderOk returns a tuple with the BackupFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFolder

`func (o *HostConfigResource) SetBackupFolder(v string)`

SetBackupFolder sets BackupFolder field to given value.

### HasBackupFolder

`func (o *HostConfigResource) HasBackupFolder() bool`

HasBackupFolder returns a boolean if a field has been set.

### SetBackupFolderNil

`func (o *HostConfigResource) SetBackupFolderNil(b bool)`

 SetBackupFolderNil sets the value for BackupFolder to be an explicit nil

### UnsetBackupFolder
`func (o *HostConfigResource) UnsetBackupFolder()`

UnsetBackupFolder ensures that no value is present for BackupFolder, not even an explicit nil
### GetBackupInterval

`func (o *HostConfigResource) GetBackupInterval() int32`

GetBackupInterval returns the BackupInterval field if non-nil, zero value otherwise.

### GetBackupIntervalOk

`func (o *HostConfigResource) GetBackupIntervalOk() (*int32, bool)`

GetBackupIntervalOk returns a tuple with the BackupInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupInterval

`func (o *HostConfigResource) SetBackupInterval(v int32)`

SetBackupInterval sets BackupInterval field to given value.

### HasBackupInterval

`func (o *HostConfigResource) HasBackupInterval() bool`

HasBackupInterval returns a boolean if a field has been set.

### GetBackupRetention

`func (o *HostConfigResource) GetBackupRetention() int32`

GetBackupRetention returns the BackupRetention field if non-nil, zero value otherwise.

### GetBackupRetentionOk

`func (o *HostConfigResource) GetBackupRetentionOk() (*int32, bool)`

GetBackupRetentionOk returns a tuple with the BackupRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetention

`func (o *HostConfigResource) SetBackupRetention(v int32)`

SetBackupRetention sets BackupRetention field to given value.

### HasBackupRetention

`func (o *HostConfigResource) HasBackupRetention() bool`

HasBackupRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


