# Go API client for sonarr

Sonarr API docs

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.0
- Package version: 0.0.1
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sonarr "github.com/devopsarr/sonarr-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sonarr.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sonarr.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), sonarr.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sonarr.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8989*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApiInfoApi* | [**GetApi**](docs/ApiInfoApi.md#getapi) | **Get** /api | 
*AuthenticationApi* | [**CreateLogin**](docs/AuthenticationApi.md#createlogin) | **Post** /login | 
*AuthenticationApi* | [**GetLogout**](docs/AuthenticationApi.md#getlogout) | **Get** /logout | 
*BackupApi* | [**CreateSystemBackupRestoreById**](docs/BackupApi.md#createsystembackuprestorebyid) | **Post** /api/v3/system/backup/restore/{id} | 
*BackupApi* | [**CreateSystemBackupRestoreUpload**](docs/BackupApi.md#createsystembackuprestoreupload) | **Post** /api/v3/system/backup/restore/upload | 
*BackupApi* | [**DeleteSystemBackup**](docs/BackupApi.md#deletesystembackup) | **Delete** /api/v3/system/backup/{id} | 
*BackupApi* | [**ListSystemBackup**](docs/BackupApi.md#listsystembackup) | **Get** /api/v3/system/backup | 
*BlocklistApi* | [**DeleteBlocklist**](docs/BlocklistApi.md#deleteblocklist) | **Delete** /api/v3/blocklist/{id} | 
*BlocklistApi* | [**DeleteBlocklistBulk**](docs/BlocklistApi.md#deleteblocklistbulk) | **Delete** /api/v3/blocklist/bulk | 
*BlocklistApi* | [**GetBlocklist**](docs/BlocklistApi.md#getblocklist) | **Get** /api/v3/blocklist | 
*CalendarApi* | [**GetCalendarById**](docs/CalendarApi.md#getcalendarbyid) | **Get** /api/v3/calendar/{id} | 
*CalendarApi* | [**ListCalendar**](docs/CalendarApi.md#listcalendar) | **Get** /api/v3/calendar | 
*CalendarFeedApi* | [**GetFeedV3CalendarSonarrIcs**](docs/CalendarFeedApi.md#getfeedv3calendarsonarrics) | **Get** /feed/v3/calendar/sonarr.ics | 
*CommandApi* | [**CreateCommand**](docs/CommandApi.md#createcommand) | **Post** /api/v3/command | 
*CommandApi* | [**DeleteCommand**](docs/CommandApi.md#deletecommand) | **Delete** /api/v3/command/{id} | 
*CommandApi* | [**GetCommandById**](docs/CommandApi.md#getcommandbyid) | **Get** /api/v3/command/{id} | 
*CommandApi* | [**ListCommand**](docs/CommandApi.md#listcommand) | **Get** /api/v3/command | 
*CustomFilterApi* | [**CreateCustomfilter**](docs/CustomFilterApi.md#createcustomfilter) | **Post** /api/v3/customfilter | 
*CustomFilterApi* | [**DeleteCustomfilter**](docs/CustomFilterApi.md#deletecustomfilter) | **Delete** /api/v3/customfilter/{id} | 
*CustomFilterApi* | [**GetCustomfilterById**](docs/CustomFilterApi.md#getcustomfilterbyid) | **Get** /api/v3/customfilter/{id} | 
*CustomFilterApi* | [**ListCustomfilter**](docs/CustomFilterApi.md#listcustomfilter) | **Get** /api/v3/customfilter | 
*CustomFilterApi* | [**UpdateCustomfilter**](docs/CustomFilterApi.md#updatecustomfilter) | **Put** /api/v3/customfilter/{id} | 
*CustomFormatApi* | [**CreateCustomformat**](docs/CustomFormatApi.md#createcustomformat) | **Post** /api/v3/customformat | 
*CustomFormatApi* | [**DeleteCustomformat**](docs/CustomFormatApi.md#deletecustomformat) | **Delete** /api/v3/customformat/{id} | 
*CustomFormatApi* | [**GetCustomformatById**](docs/CustomFormatApi.md#getcustomformatbyid) | **Get** /api/v3/customformat/{id} | 
*CustomFormatApi* | [**GetCustomformatSchema**](docs/CustomFormatApi.md#getcustomformatschema) | **Get** /api/v3/customformat/schema | 
*CustomFormatApi* | [**ListCustomformat**](docs/CustomFormatApi.md#listcustomformat) | **Get** /api/v3/customformat | 
*CustomFormatApi* | [**UpdateCustomformat**](docs/CustomFormatApi.md#updatecustomformat) | **Put** /api/v3/customformat/{id} | 
*CutoffApi* | [**GetWantedCutoff**](docs/CutoffApi.md#getwantedcutoff) | **Get** /api/v3/wanted/cutoff | 
*CutoffApi* | [**GetWantedCutoffById**](docs/CutoffApi.md#getwantedcutoffbyid) | **Get** /api/v3/wanted/cutoff/{id} | 
*DelayProfileApi* | [**CreateDelayprofile**](docs/DelayProfileApi.md#createdelayprofile) | **Post** /api/v3/delayprofile | 
*DelayProfileApi* | [**DeleteDelayprofile**](docs/DelayProfileApi.md#deletedelayprofile) | **Delete** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**GetDelayprofileById**](docs/DelayProfileApi.md#getdelayprofilebyid) | **Get** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**ListDelayprofile**](docs/DelayProfileApi.md#listdelayprofile) | **Get** /api/v3/delayprofile | 
*DelayProfileApi* | [**UpdateDelayprofile**](docs/DelayProfileApi.md#updatedelayprofile) | **Put** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**UpdateDelayprofileReorder**](docs/DelayProfileApi.md#updatedelayprofilereorder) | **Put** /api/v3/delayprofile/reorder/{id} | 
*DiskSpaceApi* | [**ListDiskspace**](docs/DiskSpaceApi.md#listdiskspace) | **Get** /api/v3/diskspace | 
*DownloadClientApi* | [**CreateDownloadclient**](docs/DownloadClientApi.md#createdownloadclient) | **Post** /api/v3/downloadclient | 
*DownloadClientApi* | [**CreateDownloadclientActionByName**](docs/DownloadClientApi.md#createdownloadclientactionbyname) | **Post** /api/v3/downloadclient/action/{name} | 
*DownloadClientApi* | [**DeleteDownloadclient**](docs/DownloadClientApi.md#deletedownloadclient) | **Delete** /api/v3/downloadclient/{id} | 
*DownloadClientApi* | [**GetDownloadclientById**](docs/DownloadClientApi.md#getdownloadclientbyid) | **Get** /api/v3/downloadclient/{id} | 
*DownloadClientApi* | [**ListDownloadclient**](docs/DownloadClientApi.md#listdownloadclient) | **Get** /api/v3/downloadclient | 
*DownloadClientApi* | [**ListDownloadclientSchema**](docs/DownloadClientApi.md#listdownloadclientschema) | **Get** /api/v3/downloadclient/schema | 
*DownloadClientApi* | [**TestDownloadclient**](docs/DownloadClientApi.md#testdownloadclient) | **Post** /api/v3/downloadclient/test | 
*DownloadClientApi* | [**TestallDownloadclient**](docs/DownloadClientApi.md#testalldownloadclient) | **Post** /api/v3/downloadclient/testall | 
*DownloadClientApi* | [**UpdateDownloadclient**](docs/DownloadClientApi.md#updatedownloadclient) | **Put** /api/v3/downloadclient/{id} | 
*DownloadClientConfigApi* | [**GetConfigDownloadclient**](docs/DownloadClientConfigApi.md#getconfigdownloadclient) | **Get** /api/v3/config/downloadclient | 
*DownloadClientConfigApi* | [**GetConfigDownloadclientById**](docs/DownloadClientConfigApi.md#getconfigdownloadclientbyid) | **Get** /api/v3/config/downloadclient/{id} | 
*DownloadClientConfigApi* | [**UpdateConfigDownloadclient**](docs/DownloadClientConfigApi.md#updateconfigdownloadclient) | **Put** /api/v3/config/downloadclient/{id} | 
*EpisodeApi* | [**GetEpisodeById**](docs/EpisodeApi.md#getepisodebyid) | **Get** /api/v3/episode/{id} | 
*EpisodeApi* | [**ListEpisode**](docs/EpisodeApi.md#listepisode) | **Get** /api/v3/episode | 
*EpisodeApi* | [**PutEpisodeMonitor**](docs/EpisodeApi.md#putepisodemonitor) | **Put** /api/v3/episode/monitor | 
*EpisodeApi* | [**UpdateEpisode**](docs/EpisodeApi.md#updateepisode) | **Put** /api/v3/episode/{id} | 
*EpisodeFileApi* | [**DeleteEpisodefile**](docs/EpisodeFileApi.md#deleteepisodefile) | **Delete** /api/v3/episodefile/{id} | 
*EpisodeFileApi* | [**DeleteEpisodefileBulk**](docs/EpisodeFileApi.md#deleteepisodefilebulk) | **Delete** /api/v3/episodefile/bulk | 
*EpisodeFileApi* | [**GetEpisodefileById**](docs/EpisodeFileApi.md#getepisodefilebyid) | **Get** /api/v3/episodefile/{id} | 
*EpisodeFileApi* | [**ListEpisodefile**](docs/EpisodeFileApi.md#listepisodefile) | **Get** /api/v3/episodefile | 
*EpisodeFileApi* | [**PutEpisodefileBulk**](docs/EpisodeFileApi.md#putepisodefilebulk) | **Put** /api/v3/episodefile/bulk | 
*EpisodeFileApi* | [**PutEpisodefileEditor**](docs/EpisodeFileApi.md#putepisodefileeditor) | **Put** /api/v3/episodefile/editor | 
*EpisodeFileApi* | [**UpdateEpisodefile**](docs/EpisodeFileApi.md#updateepisodefile) | **Put** /api/v3/episodefile/{id} | 
*FileSystemApi* | [**GetFilesystem**](docs/FileSystemApi.md#getfilesystem) | **Get** /api/v3/filesystem | 
*FileSystemApi* | [**GetFilesystemMediafiles**](docs/FileSystemApi.md#getfilesystemmediafiles) | **Get** /api/v3/filesystem/mediafiles | 
*FileSystemApi* | [**GetFilesystemType**](docs/FileSystemApi.md#getfilesystemtype) | **Get** /api/v3/filesystem/type | 
*HealthApi* | [**GetHealthById**](docs/HealthApi.md#gethealthbyid) | **Get** /api/v3/health/{id} | 
*HealthApi* | [**ListHealth**](docs/HealthApi.md#listhealth) | **Get** /api/v3/health | 
*HistoryApi* | [**CreateHistoryFailedById**](docs/HistoryApi.md#createhistoryfailedbyid) | **Post** /api/v3/history/failed/{id} | 
*HistoryApi* | [**GetHistory**](docs/HistoryApi.md#gethistory) | **Get** /api/v3/history | 
*HistoryApi* | [**ListHistorySeries**](docs/HistoryApi.md#listhistoryseries) | **Get** /api/v3/history/series | 
*HistoryApi* | [**ListHistorySince**](docs/HistoryApi.md#listhistorysince) | **Get** /api/v3/history/since | 
*HostConfigApi* | [**GetConfigHost**](docs/HostConfigApi.md#getconfighost) | **Get** /api/v3/config/host | 
*HostConfigApi* | [**GetConfigHostById**](docs/HostConfigApi.md#getconfighostbyid) | **Get** /api/v3/config/host/{id} | 
*HostConfigApi* | [**UpdateConfigHost**](docs/HostConfigApi.md#updateconfighost) | **Put** /api/v3/config/host/{id} | 
*ImportListApi* | [**CreateImportlist**](docs/ImportListApi.md#createimportlist) | **Post** /api/v3/importlist | 
*ImportListApi* | [**CreateImportlistActionByName**](docs/ImportListApi.md#createimportlistactionbyname) | **Post** /api/v3/importlist/action/{name} | 
*ImportListApi* | [**DeleteImportlist**](docs/ImportListApi.md#deleteimportlist) | **Delete** /api/v3/importlist/{id} | 
*ImportListApi* | [**GetImportlistById**](docs/ImportListApi.md#getimportlistbyid) | **Get** /api/v3/importlist/{id} | 
*ImportListApi* | [**ListImportlist**](docs/ImportListApi.md#listimportlist) | **Get** /api/v3/importlist | 
*ImportListApi* | [**ListImportlistSchema**](docs/ImportListApi.md#listimportlistschema) | **Get** /api/v3/importlist/schema | 
*ImportListApi* | [**TestImportlist**](docs/ImportListApi.md#testimportlist) | **Post** /api/v3/importlist/test | 
*ImportListApi* | [**TestallImportlist**](docs/ImportListApi.md#testallimportlist) | **Post** /api/v3/importlist/testall | 
*ImportListApi* | [**UpdateImportlist**](docs/ImportListApi.md#updateimportlist) | **Put** /api/v3/importlist/{id} | 
*ImportListExclusionApi* | [**CreateImportlistexclusion**](docs/ImportListExclusionApi.md#createimportlistexclusion) | **Post** /api/v3/importlistexclusion | 
*ImportListExclusionApi* | [**DeleteImportlistexclusion**](docs/ImportListExclusionApi.md#deleteimportlistexclusion) | **Delete** /api/v3/importlistexclusion/{id} | 
*ImportListExclusionApi* | [**GetImportlistexclusionById**](docs/ImportListExclusionApi.md#getimportlistexclusionbyid) | **Get** /api/v3/importlistexclusion/{id} | 
*ImportListExclusionApi* | [**ListImportlistexclusion**](docs/ImportListExclusionApi.md#listimportlistexclusion) | **Get** /api/v3/importlistexclusion | 
*ImportListExclusionApi* | [**UpdateImportlistexclusion**](docs/ImportListExclusionApi.md#updateimportlistexclusion) | **Put** /api/v3/importlistexclusion/{id} | 
*IndexerApi* | [**CreateIndexer**](docs/IndexerApi.md#createindexer) | **Post** /api/v3/indexer | 
*IndexerApi* | [**CreateIndexerActionByName**](docs/IndexerApi.md#createindexeractionbyname) | **Post** /api/v3/indexer/action/{name} | 
*IndexerApi* | [**DeleteIndexer**](docs/IndexerApi.md#deleteindexer) | **Delete** /api/v3/indexer/{id} | 
*IndexerApi* | [**GetIndexerById**](docs/IndexerApi.md#getindexerbyid) | **Get** /api/v3/indexer/{id} | 
*IndexerApi* | [**ListIndexer**](docs/IndexerApi.md#listindexer) | **Get** /api/v3/indexer | 
*IndexerApi* | [**ListIndexerSchema**](docs/IndexerApi.md#listindexerschema) | **Get** /api/v3/indexer/schema | 
*IndexerApi* | [**TestIndexer**](docs/IndexerApi.md#testindexer) | **Post** /api/v3/indexer/test | 
*IndexerApi* | [**TestallIndexer**](docs/IndexerApi.md#testallindexer) | **Post** /api/v3/indexer/testall | 
*IndexerApi* | [**UpdateIndexer**](docs/IndexerApi.md#updateindexer) | **Put** /api/v3/indexer/{id} | 
*IndexerConfigApi* | [**GetConfigIndexer**](docs/IndexerConfigApi.md#getconfigindexer) | **Get** /api/v3/config/indexer | 
*IndexerConfigApi* | [**GetConfigIndexerById**](docs/IndexerConfigApi.md#getconfigindexerbyid) | **Get** /api/v3/config/indexer/{id} | 
*IndexerConfigApi* | [**UpdateConfigIndexer**](docs/IndexerConfigApi.md#updateconfigindexer) | **Put** /api/v3/config/indexer/{id} | 
*InitializeJsApi* | [**GetInitializeJs**](docs/InitializeJsApi.md#getinitializejs) | **Get** /initialize.js | 
*LanguageApi* | [**GetLanguageById**](docs/LanguageApi.md#getlanguagebyid) | **Get** /api/v3/language/{id} | 
*LanguageApi* | [**ListLanguage**](docs/LanguageApi.md#listlanguage) | **Get** /api/v3/language | 
*LanguageProfileApi* | [**CreateLanguageprofile**](docs/LanguageProfileApi.md#createlanguageprofile) | **Post** /api/v3/languageprofile | 
*LanguageProfileApi* | [**DeleteLanguageprofile**](docs/LanguageProfileApi.md#deletelanguageprofile) | **Delete** /api/v3/languageprofile/{id} | 
*LanguageProfileApi* | [**GetLanguageprofileById**](docs/LanguageProfileApi.md#getlanguageprofilebyid) | **Get** /api/v3/languageprofile/{id} | 
*LanguageProfileApi* | [**ListLanguageprofile**](docs/LanguageProfileApi.md#listlanguageprofile) | **Get** /api/v3/languageprofile | 
*LanguageProfileApi* | [**UpdateLanguageprofile**](docs/LanguageProfileApi.md#updatelanguageprofile) | **Put** /api/v3/languageprofile/{id} | 
*LanguageProfileSchemaApi* | [**GetLanguageprofileSchema**](docs/LanguageProfileSchemaApi.md#getlanguageprofileschema) | **Get** /api/v3/languageprofile/schema | 
*LanguageProfileSchemaApi* | [**GetLanguageprofileSchemaById**](docs/LanguageProfileSchemaApi.md#getlanguageprofileschemabyid) | **Get** /api/v3/languageprofile/schema/{id} | 
*LocalizationApi* | [**GetLocalization**](docs/LocalizationApi.md#getlocalization) | **Get** /api/v3/localization | 
*LocalizationApi* | [**GetLocalizationById**](docs/LocalizationApi.md#getlocalizationbyid) | **Get** /api/v3/localization/{id} | 
*LogApi* | [**GetLog**](docs/LogApi.md#getlog) | **Get** /api/v3/log | 
*LogFileApi* | [**GetLogFileByFilename**](docs/LogFileApi.md#getlogfilebyfilename) | **Get** /api/v3/log/file/{filename} | 
*LogFileApi* | [**ListLogFile**](docs/LogFileApi.md#listlogfile) | **Get** /api/v3/log/file | 
*ManualImportApi* | [**CreateManualimport**](docs/ManualImportApi.md#createmanualimport) | **Post** /api/v3/manualimport | 
*ManualImportApi* | [**ListManualimport**](docs/ManualImportApi.md#listmanualimport) | **Get** /api/v3/manualimport | 
*MediaCoverApi* | [**GetMediacoverseriesidByFilename**](docs/MediaCoverApi.md#getmediacoverseriesidbyfilename) | **Get** /api/v3/mediacover/{seriesId}/{filename} | 
*MediaManagementConfigApi* | [**GetConfigMediamanagement**](docs/MediaManagementConfigApi.md#getconfigmediamanagement) | **Get** /api/v3/config/mediamanagement | 
*MediaManagementConfigApi* | [**GetConfigMediamanagementById**](docs/MediaManagementConfigApi.md#getconfigmediamanagementbyid) | **Get** /api/v3/config/mediamanagement/{id} | 
*MediaManagementConfigApi* | [**UpdateConfigMediamanagement**](docs/MediaManagementConfigApi.md#updateconfigmediamanagement) | **Put** /api/v3/config/mediamanagement/{id} | 
*MetadataApi* | [**CreateMetadata**](docs/MetadataApi.md#createmetadata) | **Post** /api/v3/metadata | 
*MetadataApi* | [**CreateMetadataActionByName**](docs/MetadataApi.md#createmetadataactionbyname) | **Post** /api/v3/metadata/action/{name} | 
*MetadataApi* | [**DeleteMetadata**](docs/MetadataApi.md#deletemetadata) | **Delete** /api/v3/metadata/{id} | 
*MetadataApi* | [**GetMetadataById**](docs/MetadataApi.md#getmetadatabyid) | **Get** /api/v3/metadata/{id} | 
*MetadataApi* | [**ListMetadata**](docs/MetadataApi.md#listmetadata) | **Get** /api/v3/metadata | 
*MetadataApi* | [**ListMetadataSchema**](docs/MetadataApi.md#listmetadataschema) | **Get** /api/v3/metadata/schema | 
*MetadataApi* | [**TestMetadata**](docs/MetadataApi.md#testmetadata) | **Post** /api/v3/metadata/test | 
*MetadataApi* | [**TestallMetadata**](docs/MetadataApi.md#testallmetadata) | **Post** /api/v3/metadata/testall | 
*MetadataApi* | [**UpdateMetadata**](docs/MetadataApi.md#updatemetadata) | **Put** /api/v3/metadata/{id} | 
*MissingApi* | [**GetWantedMissing**](docs/MissingApi.md#getwantedmissing) | **Get** /api/v3/wanted/missing | 
*MissingApi* | [**GetWantedMissingById**](docs/MissingApi.md#getwantedmissingbyid) | **Get** /api/v3/wanted/missing/{id} | 
*NamingConfigApi* | [**GetConfigNaming**](docs/NamingConfigApi.md#getconfignaming) | **Get** /api/v3/config/naming | 
*NamingConfigApi* | [**GetConfigNamingById**](docs/NamingConfigApi.md#getconfignamingbyid) | **Get** /api/v3/config/naming/{id} | 
*NamingConfigApi* | [**GetConfigNamingExamples**](docs/NamingConfigApi.md#getconfignamingexamples) | **Get** /api/v3/config/naming/examples | 
*NamingConfigApi* | [**UpdateConfigNaming**](docs/NamingConfigApi.md#updateconfignaming) | **Put** /api/v3/config/naming/{id} | 
*NotificationApi* | [**CreateNotification**](docs/NotificationApi.md#createnotification) | **Post** /api/v3/notification | 
*NotificationApi* | [**CreateNotificationActionByName**](docs/NotificationApi.md#createnotificationactionbyname) | **Post** /api/v3/notification/action/{name} | 
*NotificationApi* | [**DeleteNotification**](docs/NotificationApi.md#deletenotification) | **Delete** /api/v3/notification/{id} | 
*NotificationApi* | [**GetNotificationById**](docs/NotificationApi.md#getnotificationbyid) | **Get** /api/v3/notification/{id} | 
*NotificationApi* | [**ListNotification**](docs/NotificationApi.md#listnotification) | **Get** /api/v3/notification | 
*NotificationApi* | [**ListNotificationSchema**](docs/NotificationApi.md#listnotificationschema) | **Get** /api/v3/notification/schema | 
*NotificationApi* | [**TestNotification**](docs/NotificationApi.md#testnotification) | **Post** /api/v3/notification/test | 
*NotificationApi* | [**TestallNotification**](docs/NotificationApi.md#testallnotification) | **Post** /api/v3/notification/testall | 
*NotificationApi* | [**UpdateNotification**](docs/NotificationApi.md#updatenotification) | **Put** /api/v3/notification/{id} | 
*ParseApi* | [**GetParse**](docs/ParseApi.md#getparse) | **Get** /api/v3/parse | 
*PingApi* | [**GetPing**](docs/PingApi.md#getping) | **Get** /ping | 
*QualityDefinitionApi* | [**GetQualitydefinitionById**](docs/QualityDefinitionApi.md#getqualitydefinitionbyid) | **Get** /api/v3/qualitydefinition/{id} | 
*QualityDefinitionApi* | [**ListQualitydefinition**](docs/QualityDefinitionApi.md#listqualitydefinition) | **Get** /api/v3/qualitydefinition | 
*QualityDefinitionApi* | [**PutQualitydefinitionUpdate**](docs/QualityDefinitionApi.md#putqualitydefinitionupdate) | **Put** /api/v3/qualitydefinition/update | 
*QualityDefinitionApi* | [**UpdateQualitydefinition**](docs/QualityDefinitionApi.md#updatequalitydefinition) | **Put** /api/v3/qualitydefinition/{id} | 
*QualityProfileApi* | [**CreateQualityprofile**](docs/QualityProfileApi.md#createqualityprofile) | **Post** /api/v3/qualityprofile | 
*QualityProfileApi* | [**DeleteQualityprofile**](docs/QualityProfileApi.md#deletequalityprofile) | **Delete** /api/v3/qualityprofile/{id} | 
*QualityProfileApi* | [**GetQualityprofileById**](docs/QualityProfileApi.md#getqualityprofilebyid) | **Get** /api/v3/qualityprofile/{id} | 
*QualityProfileApi* | [**ListQualityprofile**](docs/QualityProfileApi.md#listqualityprofile) | **Get** /api/v3/qualityprofile | 
*QualityProfileApi* | [**UpdateQualityprofile**](docs/QualityProfileApi.md#updatequalityprofile) | **Put** /api/v3/qualityprofile/{id} | 
*QualityProfileSchemaApi* | [**GetQualityprofileSchema**](docs/QualityProfileSchemaApi.md#getqualityprofileschema) | **Get** /api/v3/qualityprofile/schema | 
*QueueApi* | [**DeleteQueue**](docs/QueueApi.md#deletequeue) | **Delete** /api/v3/queue/{id} | 
*QueueApi* | [**DeleteQueueBulk**](docs/QueueApi.md#deletequeuebulk) | **Delete** /api/v3/queue/bulk | 
*QueueApi* | [**GetQueue**](docs/QueueApi.md#getqueue) | **Get** /api/v3/queue | 
*QueueApi* | [**GetQueueById**](docs/QueueApi.md#getqueuebyid) | **Get** /api/v3/queue/{id} | 
*QueueActionApi* | [**CreateQueueGrabBulk**](docs/QueueActionApi.md#createqueuegrabbulk) | **Post** /api/v3/queue/grab/bulk | 
*QueueActionApi* | [**CreateQueueGrabById**](docs/QueueActionApi.md#createqueuegrabbyid) | **Post** /api/v3/queue/grab/{id} | 
*QueueDetailsApi* | [**GetQueueDetailsById**](docs/QueueDetailsApi.md#getqueuedetailsbyid) | **Get** /api/v3/queue/details/{id} | 
*QueueDetailsApi* | [**ListQueueDetails**](docs/QueueDetailsApi.md#listqueuedetails) | **Get** /api/v3/queue/details | 
*QueueStatusApi* | [**GetQueueStatus**](docs/QueueStatusApi.md#getqueuestatus) | **Get** /api/v3/queue/status | 
*QueueStatusApi* | [**GetQueueStatusById**](docs/QueueStatusApi.md#getqueuestatusbyid) | **Get** /api/v3/queue/status/{id} | 
*ReleaseApi* | [**CreateRelease**](docs/ReleaseApi.md#createrelease) | **Post** /api/v3/release | 
*ReleaseApi* | [**GetReleaseById**](docs/ReleaseApi.md#getreleasebyid) | **Get** /api/v3/release/{id} | 
*ReleaseApi* | [**ListRelease**](docs/ReleaseApi.md#listrelease) | **Get** /api/v3/release | 
*ReleaseProfileApi* | [**CreateReleaseprofile**](docs/ReleaseProfileApi.md#createreleaseprofile) | **Post** /api/v3/releaseprofile | 
*ReleaseProfileApi* | [**DeleteReleaseprofile**](docs/ReleaseProfileApi.md#deletereleaseprofile) | **Delete** /api/v3/releaseprofile/{id} | 
*ReleaseProfileApi* | [**GetReleaseprofileById**](docs/ReleaseProfileApi.md#getreleaseprofilebyid) | **Get** /api/v3/releaseprofile/{id} | 
*ReleaseProfileApi* | [**ListReleaseprofile**](docs/ReleaseProfileApi.md#listreleaseprofile) | **Get** /api/v3/releaseprofile | 
*ReleaseProfileApi* | [**UpdateReleaseprofile**](docs/ReleaseProfileApi.md#updatereleaseprofile) | **Put** /api/v3/releaseprofile/{id} | 
*ReleasePushApi* | [**CreateReleasePush**](docs/ReleasePushApi.md#createreleasepush) | **Post** /api/v3/release/push | 
*ReleasePushApi* | [**GetReleasePushById**](docs/ReleasePushApi.md#getreleasepushbyid) | **Get** /api/v3/release/push/{id} | 
*RemotePathMappingApi* | [**CreateRemotepathmapping**](docs/RemotePathMappingApi.md#createremotepathmapping) | **Post** /api/v3/remotepathmapping | 
*RemotePathMappingApi* | [**DeleteRemotepathmapping**](docs/RemotePathMappingApi.md#deleteremotepathmapping) | **Delete** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingApi* | [**GetRemotepathmappingById**](docs/RemotePathMappingApi.md#getremotepathmappingbyid) | **Get** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingApi* | [**ListRemotepathmapping**](docs/RemotePathMappingApi.md#listremotepathmapping) | **Get** /api/v3/remotepathmapping | 
*RemotePathMappingApi* | [**UpdateRemotepathmapping**](docs/RemotePathMappingApi.md#updateremotepathmapping) | **Put** /api/v3/remotepathmapping/{id} | 
*RenameEpisodeApi* | [**ListRename**](docs/RenameEpisodeApi.md#listrename) | **Get** /api/v3/rename | 
*RootFolderApi* | [**CreateRootfolder**](docs/RootFolderApi.md#createrootfolder) | **Post** /api/v3/rootfolder | 
*RootFolderApi* | [**DeleteRootfolder**](docs/RootFolderApi.md#deleterootfolder) | **Delete** /api/v3/rootfolder/{id} | 
*RootFolderApi* | [**GetRootfolderById**](docs/RootFolderApi.md#getrootfolderbyid) | **Get** /api/v3/rootfolder/{id} | 
*RootFolderApi* | [**ListRootfolder**](docs/RootFolderApi.md#listrootfolder) | **Get** /api/v3/rootfolder | 
*SeasonPassApi* | [**CreateSeasonpass**](docs/SeasonPassApi.md#createseasonpass) | **Post** /api/v3/seasonpass | 
*SeriesApi* | [**CreateSeries**](docs/SeriesApi.md#createseries) | **Post** /api/v3/series | 
*SeriesApi* | [**DeleteSeries**](docs/SeriesApi.md#deleteseries) | **Delete** /api/v3/series/{id} | 
*SeriesApi* | [**GetSeriesById**](docs/SeriesApi.md#getseriesbyid) | **Get** /api/v3/series/{id} | 
*SeriesApi* | [**ListSeries**](docs/SeriesApi.md#listseries) | **Get** /api/v3/series | 
*SeriesApi* | [**UpdateSeries**](docs/SeriesApi.md#updateseries) | **Put** /api/v3/series/{id} | 
*SeriesEditorApi* | [**DeleteSeriesEditor**](docs/SeriesEditorApi.md#deleteserieseditor) | **Delete** /api/v3/series/editor | 
*SeriesEditorApi* | [**PutSeriesEditor**](docs/SeriesEditorApi.md#putserieseditor) | **Put** /api/v3/series/editor | 
*SeriesImportApi* | [**CreateSeriesImport**](docs/SeriesImportApi.md#createseriesimport) | **Post** /api/v3/series/import | 
*SeriesLookupApi* | [**GetSeriesLookup**](docs/SeriesLookupApi.md#getserieslookup) | **Get** /api/v3/series/lookup | 
*StaticResourceApi* | [**Get**](docs/StaticResourceApi.md#get) | **Get** / | 
*StaticResourceApi* | [**GetByPath**](docs/StaticResourceApi.md#getbypath) | **Get** /{path} | 
*StaticResourceApi* | [**GetContentByPath**](docs/StaticResourceApi.md#getcontentbypath) | **Get** /content/{path} | 
*StaticResourceApi* | [**GetLogin**](docs/StaticResourceApi.md#getlogin) | **Get** /login | 
*SystemApi* | [**CreateSystemRestart**](docs/SystemApi.md#createsystemrestart) | **Post** /api/v3/system/restart | 
*SystemApi* | [**CreateSystemShutdown**](docs/SystemApi.md#createsystemshutdown) | **Post** /api/v3/system/shutdown | 
*SystemApi* | [**GetSystemRoutes**](docs/SystemApi.md#getsystemroutes) | **Get** /api/v3/system/routes | 
*SystemApi* | [**GetSystemRoutesDuplicate**](docs/SystemApi.md#getsystemroutesduplicate) | **Get** /api/v3/system/routes/duplicate | 
*SystemApi* | [**GetSystemStatus**](docs/SystemApi.md#getsystemstatus) | **Get** /api/v3/system/status | 
*TagApi* | [**CreateTag**](docs/TagApi.md#createtag) | **Post** /api/v3/tag | 
*TagApi* | [**DeleteTag**](docs/TagApi.md#deletetag) | **Delete** /api/v3/tag/{id} | 
*TagApi* | [**GetTagById**](docs/TagApi.md#gettagbyid) | **Get** /api/v3/tag/{id} | 
*TagApi* | [**ListTag**](docs/TagApi.md#listtag) | **Get** /api/v3/tag | 
*TagApi* | [**UpdateTag**](docs/TagApi.md#updatetag) | **Put** /api/v3/tag/{id} | 
*TagDetailsApi* | [**GetTagDetailById**](docs/TagDetailsApi.md#gettagdetailbyid) | **Get** /api/v3/tag/detail/{id} | 
*TagDetailsApi* | [**ListTagDetail**](docs/TagDetailsApi.md#listtagdetail) | **Get** /api/v3/tag/detail | 
*TaskApi* | [**GetSystemTaskById**](docs/TaskApi.md#getsystemtaskbyid) | **Get** /api/v3/system/task/{id} | 
*TaskApi* | [**ListSystemTask**](docs/TaskApi.md#listsystemtask) | **Get** /api/v3/system/task | 
*UiConfigApi* | [**GetConfigUi**](docs/UiConfigApi.md#getconfigui) | **Get** /api/v3/config/ui | 
*UiConfigApi* | [**GetConfigUiById**](docs/UiConfigApi.md#getconfiguibyid) | **Get** /api/v3/config/ui/{id} | 
*UiConfigApi* | [**UpdateConfigUi**](docs/UiConfigApi.md#updateconfigui) | **Put** /api/v3/config/ui/{id} | 
*UpdateApi* | [**ListUpdate**](docs/UpdateApi.md#listupdate) | **Get** /api/v3/update | 
*UpdateLogFileApi* | [**GetLogFileUpdateByFilename**](docs/UpdateLogFileApi.md#getlogfileupdatebyfilename) | **Get** /api/v3/log/file/update/{filename} | 
*UpdateLogFileApi* | [**ListLogFileUpdate**](docs/UpdateLogFileApi.md#listlogfileupdate) | **Get** /api/v3/log/file/update | 


## Documentation For Models

 - [AddSeriesOptions](docs/AddSeriesOptions.md)
 - [AlternateTitleResource](docs/AlternateTitleResource.md)
 - [ApplyTags](docs/ApplyTags.md)
 - [AuthenticationRequiredType](docs/AuthenticationRequiredType.md)
 - [AuthenticationType](docs/AuthenticationType.md)
 - [BackupResource](docs/BackupResource.md)
 - [BackupType](docs/BackupType.md)
 - [BlocklistBulkResource](docs/BlocklistBulkResource.md)
 - [BlocklistResource](docs/BlocklistResource.md)
 - [BlocklistResourcePagingResource](docs/BlocklistResourcePagingResource.md)
 - [CertificateValidationType](docs/CertificateValidationType.md)
 - [Command](docs/Command.md)
 - [CommandPriority](docs/CommandPriority.md)
 - [CommandResource](docs/CommandResource.md)
 - [CommandStatus](docs/CommandStatus.md)
 - [CommandTrigger](docs/CommandTrigger.md)
 - [CustomFilterResource](docs/CustomFilterResource.md)
 - [CustomFormatResource](docs/CustomFormatResource.md)
 - [CustomFormatSpecificationSchema](docs/CustomFormatSpecificationSchema.md)
 - [DelayProfileResource](docs/DelayProfileResource.md)
 - [DiskSpaceResource](docs/DiskSpaceResource.md)
 - [DownloadClientConfigResource](docs/DownloadClientConfigResource.md)
 - [DownloadClientResource](docs/DownloadClientResource.md)
 - [DownloadProtocol](docs/DownloadProtocol.md)
 - [EpisodeFileListResource](docs/EpisodeFileListResource.md)
 - [EpisodeFileResource](docs/EpisodeFileResource.md)
 - [EpisodeHistoryEventType](docs/EpisodeHistoryEventType.md)
 - [EpisodeResource](docs/EpisodeResource.md)
 - [EpisodeResourcePagingResource](docs/EpisodeResourcePagingResource.md)
 - [EpisodeTitleRequiredType](docs/EpisodeTitleRequiredType.md)
 - [EpisodesMonitoredResource](docs/EpisodesMonitoredResource.md)
 - [Field](docs/Field.md)
 - [FileDateType](docs/FileDateType.md)
 - [HealthCheckResult](docs/HealthCheckResult.md)
 - [HealthResource](docs/HealthResource.md)
 - [HistoryResource](docs/HistoryResource.md)
 - [HistoryResourcePagingResource](docs/HistoryResourcePagingResource.md)
 - [HostConfigResource](docs/HostConfigResource.md)
 - [HttpUri](docs/HttpUri.md)
 - [ImportListExclusionResource](docs/ImportListExclusionResource.md)
 - [ImportListResource](docs/ImportListResource.md)
 - [ImportListType](docs/ImportListType.md)
 - [IndexerConfigResource](docs/IndexerConfigResource.md)
 - [IndexerResource](docs/IndexerResource.md)
 - [Language](docs/Language.md)
 - [LanguageProfileItemResource](docs/LanguageProfileItemResource.md)
 - [LanguageProfileResource](docs/LanguageProfileResource.md)
 - [LanguageResource](docs/LanguageResource.md)
 - [LocalizationResource](docs/LocalizationResource.md)
 - [LogFileResource](docs/LogFileResource.md)
 - [LogResource](docs/LogResource.md)
 - [LogResourcePagingResource](docs/LogResourcePagingResource.md)
 - [ManualImportReprocessResource](docs/ManualImportReprocessResource.md)
 - [ManualImportResource](docs/ManualImportResource.md)
 - [MediaCover](docs/MediaCover.md)
 - [MediaCoverTypes](docs/MediaCoverTypes.md)
 - [MediaInfoResource](docs/MediaInfoResource.md)
 - [MediaManagementConfigResource](docs/MediaManagementConfigResource.md)
 - [MetadataResource](docs/MetadataResource.md)
 - [MonitorTypes](docs/MonitorTypes.md)
 - [MonitoringOptions](docs/MonitoringOptions.md)
 - [NamingConfigResource](docs/NamingConfigResource.md)
 - [NotificationResource](docs/NotificationResource.md)
 - [PagingResourceFilter](docs/PagingResourceFilter.md)
 - [ParseResource](docs/ParseResource.md)
 - [ParsedEpisodeInfo](docs/ParsedEpisodeInfo.md)
 - [PingResource](docs/PingResource.md)
 - [PrivacyLevel](docs/PrivacyLevel.md)
 - [ProfileFormatItemResource](docs/ProfileFormatItemResource.md)
 - [ProperDownloadTypes](docs/ProperDownloadTypes.md)
 - [ProviderMessage](docs/ProviderMessage.md)
 - [ProviderMessageType](docs/ProviderMessageType.md)
 - [ProxyType](docs/ProxyType.md)
 - [Quality](docs/Quality.md)
 - [QualityDefinitionResource](docs/QualityDefinitionResource.md)
 - [QualityModel](docs/QualityModel.md)
 - [QualityProfileQualityItemResource](docs/QualityProfileQualityItemResource.md)
 - [QualityProfileResource](docs/QualityProfileResource.md)
 - [QualitySource](docs/QualitySource.md)
 - [QueueBulkResource](docs/QueueBulkResource.md)
 - [QueueResource](docs/QueueResource.md)
 - [QueueResourcePagingResource](docs/QueueResourcePagingResource.md)
 - [QueueStatusResource](docs/QueueStatusResource.md)
 - [Ratings](docs/Ratings.md)
 - [Rejection](docs/Rejection.md)
 - [RejectionType](docs/RejectionType.md)
 - [ReleaseProfileResource](docs/ReleaseProfileResource.md)
 - [ReleaseResource](docs/ReleaseResource.md)
 - [RemotePathMappingResource](docs/RemotePathMappingResource.md)
 - [RenameEpisodeResource](docs/RenameEpisodeResource.md)
 - [RescanAfterRefreshType](docs/RescanAfterRefreshType.md)
 - [Revision](docs/Revision.md)
 - [RootFolderResource](docs/RootFolderResource.md)
 - [RuntimeMode](docs/RuntimeMode.md)
 - [SeasonPassResource](docs/SeasonPassResource.md)
 - [SeasonPassSeriesResource](docs/SeasonPassSeriesResource.md)
 - [SeasonResource](docs/SeasonResource.md)
 - [SeasonStatisticsResource](docs/SeasonStatisticsResource.md)
 - [SelectOption](docs/SelectOption.md)
 - [SeriesEditorResource](docs/SeriesEditorResource.md)
 - [SeriesResource](docs/SeriesResource.md)
 - [SeriesStatisticsResource](docs/SeriesStatisticsResource.md)
 - [SeriesStatusType](docs/SeriesStatusType.md)
 - [SeriesTitleInfo](docs/SeriesTitleInfo.md)
 - [SeriesTypes](docs/SeriesTypes.md)
 - [SortDirection](docs/SortDirection.md)
 - [SystemResource](docs/SystemResource.md)
 - [TagDetailsResource](docs/TagDetailsResource.md)
 - [TagResource](docs/TagResource.md)
 - [TaskResource](docs/TaskResource.md)
 - [TimeSpan](docs/TimeSpan.md)
 - [TrackedDownloadState](docs/TrackedDownloadState.md)
 - [TrackedDownloadStatus](docs/TrackedDownloadStatus.md)
 - [TrackedDownloadStatusMessage](docs/TrackedDownloadStatusMessage.md)
 - [UiConfigResource](docs/UiConfigResource.md)
 - [UnmappedFolder](docs/UnmappedFolder.md)
 - [UpdateChanges](docs/UpdateChanges.md)
 - [UpdateMechanism](docs/UpdateMechanism.md)
 - [UpdateResource](docs/UpdateResource.md)
 - [Version](docs/Version.md)


## Documentation For Authorization



### X-Api-Key

- **Type**: API key
- **API key parameter name**: X-Api-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Api-Key and passed in as the auth context for each request.


### apikey

- **Type**: API key
- **API key parameter name**: apikey
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: apikey and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



