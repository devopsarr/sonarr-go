# SystemResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **NullableString** |  | [optional] 
**InstanceName** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**BuildTime** | Pointer to **time.Time** |  | [optional] 
**IsDebug** | Pointer to **bool** |  | [optional] 
**IsProduction** | Pointer to **bool** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**IsUserInteractive** | Pointer to **bool** |  | [optional] 
**StartupPath** | Pointer to **NullableString** |  | [optional] 
**AppData** | Pointer to **NullableString** |  | [optional] 
**OsName** | Pointer to **NullableString** |  | [optional] 
**OsVersion** | Pointer to **NullableString** |  | [optional] 
**IsNetCore** | Pointer to **bool** |  | [optional] 
**IsLinux** | Pointer to **bool** |  | [optional] 
**IsOsx** | Pointer to **bool** |  | [optional] 
**IsWindows** | Pointer to **bool** |  | [optional] 
**IsDocker** | Pointer to **bool** |  | [optional] 
**Mode** | Pointer to [**RuntimeMode**](RuntimeMode.md) |  | [optional] 
**Branch** | Pointer to **NullableString** |  | [optional] 
**Authentication** | Pointer to [**AuthenticationType**](AuthenticationType.md) |  | [optional] 
**SqliteVersion** | Pointer to **string** |  | [optional] 
**MigrationVersion** | Pointer to **int32** |  | [optional] 
**UrlBase** | Pointer to **NullableString** |  | [optional] 
**RuntimeVersion** | Pointer to **string** |  | [optional] 
**RuntimeName** | Pointer to **NullableString** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**PackageVersion** | Pointer to **NullableString** |  | [optional] 
**PackageAuthor** | Pointer to **NullableString** |  | [optional] 
**PackageUpdateMechanism** | Pointer to [**UpdateMechanism**](UpdateMechanism.md) |  | [optional] 
**PackageUpdateMechanismMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSystemResource

`func NewSystemResource() *SystemResource`

NewSystemResource instantiates a new SystemResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemResourceWithDefaults

`func NewSystemResourceWithDefaults() *SystemResource`

NewSystemResourceWithDefaults instantiates a new SystemResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *SystemResource) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *SystemResource) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *SystemResource) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *SystemResource) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### SetAppNameNil

`func (o *SystemResource) SetAppNameNil(b bool)`

 SetAppNameNil sets the value for AppName to be an explicit nil

### UnsetAppName
`func (o *SystemResource) UnsetAppName()`

UnsetAppName ensures that no value is present for AppName, not even an explicit nil
### GetInstanceName

`func (o *SystemResource) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *SystemResource) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *SystemResource) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *SystemResource) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### SetInstanceNameNil

`func (o *SystemResource) SetInstanceNameNil(b bool)`

 SetInstanceNameNil sets the value for InstanceName to be an explicit nil

### UnsetInstanceName
`func (o *SystemResource) UnsetInstanceName()`

UnsetInstanceName ensures that no value is present for InstanceName, not even an explicit nil
### GetVersion

`func (o *SystemResource) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SystemResource) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SystemResource) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SystemResource) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *SystemResource) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *SystemResource) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetBuildTime

`func (o *SystemResource) GetBuildTime() time.Time`

GetBuildTime returns the BuildTime field if non-nil, zero value otherwise.

### GetBuildTimeOk

`func (o *SystemResource) GetBuildTimeOk() (*time.Time, bool)`

GetBuildTimeOk returns a tuple with the BuildTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTime

`func (o *SystemResource) SetBuildTime(v time.Time)`

SetBuildTime sets BuildTime field to given value.

### HasBuildTime

`func (o *SystemResource) HasBuildTime() bool`

HasBuildTime returns a boolean if a field has been set.

### GetIsDebug

`func (o *SystemResource) GetIsDebug() bool`

GetIsDebug returns the IsDebug field if non-nil, zero value otherwise.

### GetIsDebugOk

`func (o *SystemResource) GetIsDebugOk() (*bool, bool)`

GetIsDebugOk returns a tuple with the IsDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDebug

`func (o *SystemResource) SetIsDebug(v bool)`

SetIsDebug sets IsDebug field to given value.

### HasIsDebug

`func (o *SystemResource) HasIsDebug() bool`

HasIsDebug returns a boolean if a field has been set.

### GetIsProduction

`func (o *SystemResource) GetIsProduction() bool`

GetIsProduction returns the IsProduction field if non-nil, zero value otherwise.

### GetIsProductionOk

`func (o *SystemResource) GetIsProductionOk() (*bool, bool)`

GetIsProductionOk returns a tuple with the IsProduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProduction

`func (o *SystemResource) SetIsProduction(v bool)`

SetIsProduction sets IsProduction field to given value.

### HasIsProduction

`func (o *SystemResource) HasIsProduction() bool`

HasIsProduction returns a boolean if a field has been set.

### GetIsAdmin

`func (o *SystemResource) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *SystemResource) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *SystemResource) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *SystemResource) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetIsUserInteractive

`func (o *SystemResource) GetIsUserInteractive() bool`

GetIsUserInteractive returns the IsUserInteractive field if non-nil, zero value otherwise.

### GetIsUserInteractiveOk

`func (o *SystemResource) GetIsUserInteractiveOk() (*bool, bool)`

GetIsUserInteractiveOk returns a tuple with the IsUserInteractive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserInteractive

`func (o *SystemResource) SetIsUserInteractive(v bool)`

SetIsUserInteractive sets IsUserInteractive field to given value.

### HasIsUserInteractive

`func (o *SystemResource) HasIsUserInteractive() bool`

HasIsUserInteractive returns a boolean if a field has been set.

### GetStartupPath

`func (o *SystemResource) GetStartupPath() string`

GetStartupPath returns the StartupPath field if non-nil, zero value otherwise.

### GetStartupPathOk

`func (o *SystemResource) GetStartupPathOk() (*string, bool)`

GetStartupPathOk returns a tuple with the StartupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupPath

`func (o *SystemResource) SetStartupPath(v string)`

SetStartupPath sets StartupPath field to given value.

### HasStartupPath

`func (o *SystemResource) HasStartupPath() bool`

HasStartupPath returns a boolean if a field has been set.

### SetStartupPathNil

`func (o *SystemResource) SetStartupPathNil(b bool)`

 SetStartupPathNil sets the value for StartupPath to be an explicit nil

### UnsetStartupPath
`func (o *SystemResource) UnsetStartupPath()`

UnsetStartupPath ensures that no value is present for StartupPath, not even an explicit nil
### GetAppData

`func (o *SystemResource) GetAppData() string`

GetAppData returns the AppData field if non-nil, zero value otherwise.

### GetAppDataOk

`func (o *SystemResource) GetAppDataOk() (*string, bool)`

GetAppDataOk returns a tuple with the AppData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppData

`func (o *SystemResource) SetAppData(v string)`

SetAppData sets AppData field to given value.

### HasAppData

`func (o *SystemResource) HasAppData() bool`

HasAppData returns a boolean if a field has been set.

### SetAppDataNil

`func (o *SystemResource) SetAppDataNil(b bool)`

 SetAppDataNil sets the value for AppData to be an explicit nil

### UnsetAppData
`func (o *SystemResource) UnsetAppData()`

UnsetAppData ensures that no value is present for AppData, not even an explicit nil
### GetOsName

`func (o *SystemResource) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *SystemResource) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *SystemResource) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *SystemResource) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### SetOsNameNil

`func (o *SystemResource) SetOsNameNil(b bool)`

 SetOsNameNil sets the value for OsName to be an explicit nil

### UnsetOsName
`func (o *SystemResource) UnsetOsName()`

UnsetOsName ensures that no value is present for OsName, not even an explicit nil
### GetOsVersion

`func (o *SystemResource) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *SystemResource) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *SystemResource) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *SystemResource) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersionNil

`func (o *SystemResource) SetOsVersionNil(b bool)`

 SetOsVersionNil sets the value for OsVersion to be an explicit nil

### UnsetOsVersion
`func (o *SystemResource) UnsetOsVersion()`

UnsetOsVersion ensures that no value is present for OsVersion, not even an explicit nil
### GetIsNetCore

`func (o *SystemResource) GetIsNetCore() bool`

GetIsNetCore returns the IsNetCore field if non-nil, zero value otherwise.

### GetIsNetCoreOk

`func (o *SystemResource) GetIsNetCoreOk() (*bool, bool)`

GetIsNetCoreOk returns a tuple with the IsNetCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNetCore

`func (o *SystemResource) SetIsNetCore(v bool)`

SetIsNetCore sets IsNetCore field to given value.

### HasIsNetCore

`func (o *SystemResource) HasIsNetCore() bool`

HasIsNetCore returns a boolean if a field has been set.

### GetIsLinux

`func (o *SystemResource) GetIsLinux() bool`

GetIsLinux returns the IsLinux field if non-nil, zero value otherwise.

### GetIsLinuxOk

`func (o *SystemResource) GetIsLinuxOk() (*bool, bool)`

GetIsLinuxOk returns a tuple with the IsLinux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinux

`func (o *SystemResource) SetIsLinux(v bool)`

SetIsLinux sets IsLinux field to given value.

### HasIsLinux

`func (o *SystemResource) HasIsLinux() bool`

HasIsLinux returns a boolean if a field has been set.

### GetIsOsx

`func (o *SystemResource) GetIsOsx() bool`

GetIsOsx returns the IsOsx field if non-nil, zero value otherwise.

### GetIsOsxOk

`func (o *SystemResource) GetIsOsxOk() (*bool, bool)`

GetIsOsxOk returns a tuple with the IsOsx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOsx

`func (o *SystemResource) SetIsOsx(v bool)`

SetIsOsx sets IsOsx field to given value.

### HasIsOsx

`func (o *SystemResource) HasIsOsx() bool`

HasIsOsx returns a boolean if a field has been set.

### GetIsWindows

`func (o *SystemResource) GetIsWindows() bool`

GetIsWindows returns the IsWindows field if non-nil, zero value otherwise.

### GetIsWindowsOk

`func (o *SystemResource) GetIsWindowsOk() (*bool, bool)`

GetIsWindowsOk returns a tuple with the IsWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindows

`func (o *SystemResource) SetIsWindows(v bool)`

SetIsWindows sets IsWindows field to given value.

### HasIsWindows

`func (o *SystemResource) HasIsWindows() bool`

HasIsWindows returns a boolean if a field has been set.

### GetIsDocker

`func (o *SystemResource) GetIsDocker() bool`

GetIsDocker returns the IsDocker field if non-nil, zero value otherwise.

### GetIsDockerOk

`func (o *SystemResource) GetIsDockerOk() (*bool, bool)`

GetIsDockerOk returns a tuple with the IsDocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDocker

`func (o *SystemResource) SetIsDocker(v bool)`

SetIsDocker sets IsDocker field to given value.

### HasIsDocker

`func (o *SystemResource) HasIsDocker() bool`

HasIsDocker returns a boolean if a field has been set.

### GetMode

`func (o *SystemResource) GetMode() RuntimeMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SystemResource) GetModeOk() (*RuntimeMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SystemResource) SetMode(v RuntimeMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SystemResource) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetBranch

`func (o *SystemResource) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *SystemResource) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *SystemResource) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *SystemResource) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### SetBranchNil

`func (o *SystemResource) SetBranchNil(b bool)`

 SetBranchNil sets the value for Branch to be an explicit nil

### UnsetBranch
`func (o *SystemResource) UnsetBranch()`

UnsetBranch ensures that no value is present for Branch, not even an explicit nil
### GetAuthentication

`func (o *SystemResource) GetAuthentication() AuthenticationType`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *SystemResource) GetAuthenticationOk() (*AuthenticationType, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *SystemResource) SetAuthentication(v AuthenticationType)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *SystemResource) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetSqliteVersion

`func (o *SystemResource) GetSqliteVersion() string`

GetSqliteVersion returns the SqliteVersion field if non-nil, zero value otherwise.

### GetSqliteVersionOk

`func (o *SystemResource) GetSqliteVersionOk() (*string, bool)`

GetSqliteVersionOk returns a tuple with the SqliteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqliteVersion

`func (o *SystemResource) SetSqliteVersion(v string)`

SetSqliteVersion sets SqliteVersion field to given value.

### HasSqliteVersion

`func (o *SystemResource) HasSqliteVersion() bool`

HasSqliteVersion returns a boolean if a field has been set.

### GetMigrationVersion

`func (o *SystemResource) GetMigrationVersion() int32`

GetMigrationVersion returns the MigrationVersion field if non-nil, zero value otherwise.

### GetMigrationVersionOk

`func (o *SystemResource) GetMigrationVersionOk() (*int32, bool)`

GetMigrationVersionOk returns a tuple with the MigrationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationVersion

`func (o *SystemResource) SetMigrationVersion(v int32)`

SetMigrationVersion sets MigrationVersion field to given value.

### HasMigrationVersion

`func (o *SystemResource) HasMigrationVersion() bool`

HasMigrationVersion returns a boolean if a field has been set.

### GetUrlBase

`func (o *SystemResource) GetUrlBase() string`

GetUrlBase returns the UrlBase field if non-nil, zero value otherwise.

### GetUrlBaseOk

`func (o *SystemResource) GetUrlBaseOk() (*string, bool)`

GetUrlBaseOk returns a tuple with the UrlBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlBase

`func (o *SystemResource) SetUrlBase(v string)`

SetUrlBase sets UrlBase field to given value.

### HasUrlBase

`func (o *SystemResource) HasUrlBase() bool`

HasUrlBase returns a boolean if a field has been set.

### SetUrlBaseNil

`func (o *SystemResource) SetUrlBaseNil(b bool)`

 SetUrlBaseNil sets the value for UrlBase to be an explicit nil

### UnsetUrlBase
`func (o *SystemResource) UnsetUrlBase()`

UnsetUrlBase ensures that no value is present for UrlBase, not even an explicit nil
### GetRuntimeVersion

`func (o *SystemResource) GetRuntimeVersion() string`

GetRuntimeVersion returns the RuntimeVersion field if non-nil, zero value otherwise.

### GetRuntimeVersionOk

`func (o *SystemResource) GetRuntimeVersionOk() (*string, bool)`

GetRuntimeVersionOk returns a tuple with the RuntimeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeVersion

`func (o *SystemResource) SetRuntimeVersion(v string)`

SetRuntimeVersion sets RuntimeVersion field to given value.

### HasRuntimeVersion

`func (o *SystemResource) HasRuntimeVersion() bool`

HasRuntimeVersion returns a boolean if a field has been set.

### GetRuntimeName

`func (o *SystemResource) GetRuntimeName() string`

GetRuntimeName returns the RuntimeName field if non-nil, zero value otherwise.

### GetRuntimeNameOk

`func (o *SystemResource) GetRuntimeNameOk() (*string, bool)`

GetRuntimeNameOk returns a tuple with the RuntimeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeName

`func (o *SystemResource) SetRuntimeName(v string)`

SetRuntimeName sets RuntimeName field to given value.

### HasRuntimeName

`func (o *SystemResource) HasRuntimeName() bool`

HasRuntimeName returns a boolean if a field has been set.

### SetRuntimeNameNil

`func (o *SystemResource) SetRuntimeNameNil(b bool)`

 SetRuntimeNameNil sets the value for RuntimeName to be an explicit nil

### UnsetRuntimeName
`func (o *SystemResource) UnsetRuntimeName()`

UnsetRuntimeName ensures that no value is present for RuntimeName, not even an explicit nil
### GetStartTime

`func (o *SystemResource) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SystemResource) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SystemResource) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SystemResource) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetPackageVersion

`func (o *SystemResource) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *SystemResource) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *SystemResource) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *SystemResource) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### SetPackageVersionNil

`func (o *SystemResource) SetPackageVersionNil(b bool)`

 SetPackageVersionNil sets the value for PackageVersion to be an explicit nil

### UnsetPackageVersion
`func (o *SystemResource) UnsetPackageVersion()`

UnsetPackageVersion ensures that no value is present for PackageVersion, not even an explicit nil
### GetPackageAuthor

`func (o *SystemResource) GetPackageAuthor() string`

GetPackageAuthor returns the PackageAuthor field if non-nil, zero value otherwise.

### GetPackageAuthorOk

`func (o *SystemResource) GetPackageAuthorOk() (*string, bool)`

GetPackageAuthorOk returns a tuple with the PackageAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageAuthor

`func (o *SystemResource) SetPackageAuthor(v string)`

SetPackageAuthor sets PackageAuthor field to given value.

### HasPackageAuthor

`func (o *SystemResource) HasPackageAuthor() bool`

HasPackageAuthor returns a boolean if a field has been set.

### SetPackageAuthorNil

`func (o *SystemResource) SetPackageAuthorNil(b bool)`

 SetPackageAuthorNil sets the value for PackageAuthor to be an explicit nil

### UnsetPackageAuthor
`func (o *SystemResource) UnsetPackageAuthor()`

UnsetPackageAuthor ensures that no value is present for PackageAuthor, not even an explicit nil
### GetPackageUpdateMechanism

`func (o *SystemResource) GetPackageUpdateMechanism() UpdateMechanism`

GetPackageUpdateMechanism returns the PackageUpdateMechanism field if non-nil, zero value otherwise.

### GetPackageUpdateMechanismOk

`func (o *SystemResource) GetPackageUpdateMechanismOk() (*UpdateMechanism, bool)`

GetPackageUpdateMechanismOk returns a tuple with the PackageUpdateMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUpdateMechanism

`func (o *SystemResource) SetPackageUpdateMechanism(v UpdateMechanism)`

SetPackageUpdateMechanism sets PackageUpdateMechanism field to given value.

### HasPackageUpdateMechanism

`func (o *SystemResource) HasPackageUpdateMechanism() bool`

HasPackageUpdateMechanism returns a boolean if a field has been set.

### GetPackageUpdateMechanismMessage

`func (o *SystemResource) GetPackageUpdateMechanismMessage() string`

GetPackageUpdateMechanismMessage returns the PackageUpdateMechanismMessage field if non-nil, zero value otherwise.

### GetPackageUpdateMechanismMessageOk

`func (o *SystemResource) GetPackageUpdateMechanismMessageOk() (*string, bool)`

GetPackageUpdateMechanismMessageOk returns a tuple with the PackageUpdateMechanismMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUpdateMechanismMessage

`func (o *SystemResource) SetPackageUpdateMechanismMessage(v string)`

SetPackageUpdateMechanismMessage sets PackageUpdateMechanismMessage field to given value.

### HasPackageUpdateMechanismMessage

`func (o *SystemResource) HasPackageUpdateMechanismMessage() bool`

HasPackageUpdateMechanismMessage returns a boolean if a field has been set.

### SetPackageUpdateMechanismMessageNil

`func (o *SystemResource) SetPackageUpdateMechanismMessageNil(b bool)`

 SetPackageUpdateMechanismMessageNil sets the value for PackageUpdateMechanismMessage to be an explicit nil

### UnsetPackageUpdateMechanismMessage
`func (o *SystemResource) UnsetPackageUpdateMechanismMessage()`

UnsetPackageUpdateMechanismMessage ensures that no value is present for PackageUpdateMechanismMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


