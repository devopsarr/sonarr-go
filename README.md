# Go API client for sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v4.0.9.2244
- Package version: 1.0.1 <!--- x-release-please-version -->
- Generator version: 7.9.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import sonarr "github.com/devopsarr/sonarr-go/sonarr"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sonarr.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), sonarr.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sonarr.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), sonarr.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sonarr.ContextOperationServerIndices` and `sonarr.ContextOperationServerVariables` context maps.

```go
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
*ApiInfoAPI* | [**GetApi**](docs/ApiInfoAPI.md#getapi) | **Get** /api | 
*AuthenticationAPI* | [**CreateLogin**](docs/AuthenticationAPI.md#createlogin) | **Post** /login | 
*AuthenticationAPI* | [**GetLogout**](docs/AuthenticationAPI.md#getlogout) | **Get** /logout | 
*AutoTaggingAPI* | [**CreateAutoTagging**](docs/AutoTaggingAPI.md#createautotagging) | **Post** /api/v3/autotagging | 
*AutoTaggingAPI* | [**DeleteAutoTagging**](docs/AutoTaggingAPI.md#deleteautotagging) | **Delete** /api/v3/autotagging/{id} | 
*AutoTaggingAPI* | [**GetAutoTaggingById**](docs/AutoTaggingAPI.md#getautotaggingbyid) | **Get** /api/v3/autotagging/{id} | 
*AutoTaggingAPI* | [**ListAutoTagging**](docs/AutoTaggingAPI.md#listautotagging) | **Get** /api/v3/autotagging | 
*AutoTaggingAPI* | [**ListAutoTaggingSchema**](docs/AutoTaggingAPI.md#listautotaggingschema) | **Get** /api/v3/autotagging/schema | 
*AutoTaggingAPI* | [**UpdateAutoTagging**](docs/AutoTaggingAPI.md#updateautotagging) | **Put** /api/v3/autotagging/{id} | 
*BackupAPI* | [**CreateSystemBackupRestoreById**](docs/BackupAPI.md#createsystembackuprestorebyid) | **Post** /api/v3/system/backup/restore/{id} | 
*BackupAPI* | [**CreateSystemBackupRestoreUpload**](docs/BackupAPI.md#createsystembackuprestoreupload) | **Post** /api/v3/system/backup/restore/upload | 
*BackupAPI* | [**DeleteSystemBackup**](docs/BackupAPI.md#deletesystembackup) | **Delete** /api/v3/system/backup/{id} | 
*BackupAPI* | [**ListSystemBackup**](docs/BackupAPI.md#listsystembackup) | **Get** /api/v3/system/backup | 
*BlocklistAPI* | [**DeleteBlocklist**](docs/BlocklistAPI.md#deleteblocklist) | **Delete** /api/v3/blocklist/{id} | 
*BlocklistAPI* | [**DeleteBlocklistBulk**](docs/BlocklistAPI.md#deleteblocklistbulk) | **Delete** /api/v3/blocklist/bulk | 
*BlocklistAPI* | [**GetBlocklist**](docs/BlocklistAPI.md#getblocklist) | **Get** /api/v3/blocklist | 
*CalendarAPI* | [**GetCalendarById**](docs/CalendarAPI.md#getcalendarbyid) | **Get** /api/v3/calendar/{id} | 
*CalendarAPI* | [**ListCalendar**](docs/CalendarAPI.md#listcalendar) | **Get** /api/v3/calendar | 
*CalendarFeedAPI* | [**GetFeedV3CalendarSonarrIcs**](docs/CalendarFeedAPI.md#getfeedv3calendarsonarrics) | **Get** /feed/v3/calendar/sonarr.ics | 
*CommandAPI* | [**CreateCommand**](docs/CommandAPI.md#createcommand) | **Post** /api/v3/command | 
*CommandAPI* | [**DeleteCommand**](docs/CommandAPI.md#deletecommand) | **Delete** /api/v3/command/{id} | 
*CommandAPI* | [**GetCommandById**](docs/CommandAPI.md#getcommandbyid) | **Get** /api/v3/command/{id} | 
*CommandAPI* | [**ListCommand**](docs/CommandAPI.md#listcommand) | **Get** /api/v3/command | 
*CustomFilterAPI* | [**CreateCustomFilter**](docs/CustomFilterAPI.md#createcustomfilter) | **Post** /api/v3/customfilter | 
*CustomFilterAPI* | [**DeleteCustomFilter**](docs/CustomFilterAPI.md#deletecustomfilter) | **Delete** /api/v3/customfilter/{id} | 
*CustomFilterAPI* | [**GetCustomFilterById**](docs/CustomFilterAPI.md#getcustomfilterbyid) | **Get** /api/v3/customfilter/{id} | 
*CustomFilterAPI* | [**ListCustomFilter**](docs/CustomFilterAPI.md#listcustomfilter) | **Get** /api/v3/customfilter | 
*CustomFilterAPI* | [**UpdateCustomFilter**](docs/CustomFilterAPI.md#updatecustomfilter) | **Put** /api/v3/customfilter/{id} | 
*CustomFormatAPI* | [**CreateCustomFormat**](docs/CustomFormatAPI.md#createcustomformat) | **Post** /api/v3/customformat | 
*CustomFormatAPI* | [**DeleteCustomFormat**](docs/CustomFormatAPI.md#deletecustomformat) | **Delete** /api/v3/customformat/{id} | 
*CustomFormatAPI* | [**GetCustomFormatById**](docs/CustomFormatAPI.md#getcustomformatbyid) | **Get** /api/v3/customformat/{id} | 
*CustomFormatAPI* | [**ListCustomFormat**](docs/CustomFormatAPI.md#listcustomformat) | **Get** /api/v3/customformat | 
*CustomFormatAPI* | [**ListCustomFormatSchema**](docs/CustomFormatAPI.md#listcustomformatschema) | **Get** /api/v3/customformat/schema | 
*CustomFormatAPI* | [**UpdateCustomFormat**](docs/CustomFormatAPI.md#updatecustomformat) | **Put** /api/v3/customformat/{id} | 
*CutoffAPI* | [**GetWantedCutoff**](docs/CutoffAPI.md#getwantedcutoff) | **Get** /api/v3/wanted/cutoff | 
*CutoffAPI* | [**GetWantedCutoffById**](docs/CutoffAPI.md#getwantedcutoffbyid) | **Get** /api/v3/wanted/cutoff/{id} | 
*DelayProfileAPI* | [**CreateDelayProfile**](docs/DelayProfileAPI.md#createdelayprofile) | **Post** /api/v3/delayprofile | 
*DelayProfileAPI* | [**DeleteDelayProfile**](docs/DelayProfileAPI.md#deletedelayprofile) | **Delete** /api/v3/delayprofile/{id} | 
*DelayProfileAPI* | [**GetDelayProfileById**](docs/DelayProfileAPI.md#getdelayprofilebyid) | **Get** /api/v3/delayprofile/{id} | 
*DelayProfileAPI* | [**ListDelayProfile**](docs/DelayProfileAPI.md#listdelayprofile) | **Get** /api/v3/delayprofile | 
*DelayProfileAPI* | [**UpdateDelayProfile**](docs/DelayProfileAPI.md#updatedelayprofile) | **Put** /api/v3/delayprofile/{id} | 
*DelayProfileAPI* | [**UpdateDelayProfileReorder**](docs/DelayProfileAPI.md#updatedelayprofilereorder) | **Put** /api/v3/delayprofile/reorder/{id} | 
*DiskSpaceAPI* | [**ListDiskSpace**](docs/DiskSpaceAPI.md#listdiskspace) | **Get** /api/v3/diskspace | 
*DownloadClientAPI* | [**CreateDownloadClient**](docs/DownloadClientAPI.md#createdownloadclient) | **Post** /api/v3/downloadclient | 
*DownloadClientAPI* | [**CreateDownloadClientActionByName**](docs/DownloadClientAPI.md#createdownloadclientactionbyname) | **Post** /api/v3/downloadclient/action/{name} | 
*DownloadClientAPI* | [**DeleteDownloadClient**](docs/DownloadClientAPI.md#deletedownloadclient) | **Delete** /api/v3/downloadclient/{id} | 
*DownloadClientAPI* | [**DeleteDownloadClientBulk**](docs/DownloadClientAPI.md#deletedownloadclientbulk) | **Delete** /api/v3/downloadclient/bulk | 
*DownloadClientAPI* | [**GetDownloadClientById**](docs/DownloadClientAPI.md#getdownloadclientbyid) | **Get** /api/v3/downloadclient/{id} | 
*DownloadClientAPI* | [**ListDownloadClient**](docs/DownloadClientAPI.md#listdownloadclient) | **Get** /api/v3/downloadclient | 
*DownloadClientAPI* | [**ListDownloadClientSchema**](docs/DownloadClientAPI.md#listdownloadclientschema) | **Get** /api/v3/downloadclient/schema | 
*DownloadClientAPI* | [**PutDownloadClientBulk**](docs/DownloadClientAPI.md#putdownloadclientbulk) | **Put** /api/v3/downloadclient/bulk | 
*DownloadClientAPI* | [**TestDownloadClient**](docs/DownloadClientAPI.md#testdownloadclient) | **Post** /api/v3/downloadclient/test | 
*DownloadClientAPI* | [**TestallDownloadClient**](docs/DownloadClientAPI.md#testalldownloadclient) | **Post** /api/v3/downloadclient/testall | 
*DownloadClientAPI* | [**UpdateDownloadClient**](docs/DownloadClientAPI.md#updatedownloadclient) | **Put** /api/v3/downloadclient/{id} | 
*DownloadClientConfigAPI* | [**GetDownloadClientConfig**](docs/DownloadClientConfigAPI.md#getdownloadclientconfig) | **Get** /api/v3/config/downloadclient | 
*DownloadClientConfigAPI* | [**GetDownloadClientConfigById**](docs/DownloadClientConfigAPI.md#getdownloadclientconfigbyid) | **Get** /api/v3/config/downloadclient/{id} | 
*DownloadClientConfigAPI* | [**UpdateDownloadClientConfig**](docs/DownloadClientConfigAPI.md#updatedownloadclientconfig) | **Put** /api/v3/config/downloadclient/{id} | 
*EpisodeAPI* | [**GetEpisodeById**](docs/EpisodeAPI.md#getepisodebyid) | **Get** /api/v3/episode/{id} | 
*EpisodeAPI* | [**ListEpisode**](docs/EpisodeAPI.md#listepisode) | **Get** /api/v3/episode | 
*EpisodeAPI* | [**PutEpisodeMonitor**](docs/EpisodeAPI.md#putepisodemonitor) | **Put** /api/v3/episode/monitor | 
*EpisodeAPI* | [**UpdateEpisode**](docs/EpisodeAPI.md#updateepisode) | **Put** /api/v3/episode/{id} | 
*EpisodeFileAPI* | [**DeleteEpisodeFile**](docs/EpisodeFileAPI.md#deleteepisodefile) | **Delete** /api/v3/episodefile/{id} | 
*EpisodeFileAPI* | [**DeleteEpisodeFileBulk**](docs/EpisodeFileAPI.md#deleteepisodefilebulk) | **Delete** /api/v3/episodefile/bulk | 
*EpisodeFileAPI* | [**GetEpisodeFileById**](docs/EpisodeFileAPI.md#getepisodefilebyid) | **Get** /api/v3/episodefile/{id} | 
*EpisodeFileAPI* | [**ListEpisodeFile**](docs/EpisodeFileAPI.md#listepisodefile) | **Get** /api/v3/episodefile | 
*EpisodeFileAPI* | [**PutEpisodeFileBulk**](docs/EpisodeFileAPI.md#putepisodefilebulk) | **Put** /api/v3/episodefile/bulk | 
*EpisodeFileAPI* | [**PutEpisodeFileEditor**](docs/EpisodeFileAPI.md#putepisodefileeditor) | **Put** /api/v3/episodefile/editor | 
*EpisodeFileAPI* | [**UpdateEpisodeFile**](docs/EpisodeFileAPI.md#updateepisodefile) | **Put** /api/v3/episodefile/{id} | 
*FileSystemAPI* | [**GetFileSystem**](docs/FileSystemAPI.md#getfilesystem) | **Get** /api/v3/filesystem | 
*FileSystemAPI* | [**GetFileSystemMediafiles**](docs/FileSystemAPI.md#getfilesystemmediafiles) | **Get** /api/v3/filesystem/mediafiles | 
*FileSystemAPI* | [**GetFileSystemType**](docs/FileSystemAPI.md#getfilesystemtype) | **Get** /api/v3/filesystem/type | 
*HealthAPI* | [**ListHealth**](docs/HealthAPI.md#listhealth) | **Get** /api/v3/health | 
*HistoryAPI* | [**CreateHistoryFailedById**](docs/HistoryAPI.md#createhistoryfailedbyid) | **Post** /api/v3/history/failed/{id} | 
*HistoryAPI* | [**GetHistory**](docs/HistoryAPI.md#gethistory) | **Get** /api/v3/history | 
*HistoryAPI* | [**ListHistorySeries**](docs/HistoryAPI.md#listhistoryseries) | **Get** /api/v3/history/series | 
*HistoryAPI* | [**ListHistorySince**](docs/HistoryAPI.md#listhistorysince) | **Get** /api/v3/history/since | 
*HostConfigAPI* | [**GetHostConfig**](docs/HostConfigAPI.md#gethostconfig) | **Get** /api/v3/config/host | 
*HostConfigAPI* | [**GetHostConfigById**](docs/HostConfigAPI.md#gethostconfigbyid) | **Get** /api/v3/config/host/{id} | 
*HostConfigAPI* | [**UpdateHostConfig**](docs/HostConfigAPI.md#updatehostconfig) | **Put** /api/v3/config/host/{id} | 
*ImportListAPI* | [**CreateImportList**](docs/ImportListAPI.md#createimportlist) | **Post** /api/v3/importlist | 
*ImportListAPI* | [**CreateImportListActionByName**](docs/ImportListAPI.md#createimportlistactionbyname) | **Post** /api/v3/importlist/action/{name} | 
*ImportListAPI* | [**DeleteImportList**](docs/ImportListAPI.md#deleteimportlist) | **Delete** /api/v3/importlist/{id} | 
*ImportListAPI* | [**DeleteImportListBulk**](docs/ImportListAPI.md#deleteimportlistbulk) | **Delete** /api/v3/importlist/bulk | 
*ImportListAPI* | [**GetImportListById**](docs/ImportListAPI.md#getimportlistbyid) | **Get** /api/v3/importlist/{id} | 
*ImportListAPI* | [**ListImportList**](docs/ImportListAPI.md#listimportlist) | **Get** /api/v3/importlist | 
*ImportListAPI* | [**ListImportListSchema**](docs/ImportListAPI.md#listimportlistschema) | **Get** /api/v3/importlist/schema | 
*ImportListAPI* | [**PutImportListBulk**](docs/ImportListAPI.md#putimportlistbulk) | **Put** /api/v3/importlist/bulk | 
*ImportListAPI* | [**TestImportList**](docs/ImportListAPI.md#testimportlist) | **Post** /api/v3/importlist/test | 
*ImportListAPI* | [**TestallImportList**](docs/ImportListAPI.md#testallimportlist) | **Post** /api/v3/importlist/testall | 
*ImportListAPI* | [**UpdateImportList**](docs/ImportListAPI.md#updateimportlist) | **Put** /api/v3/importlist/{id} | 
*ImportListConfigAPI* | [**GetImportListConfig**](docs/ImportListConfigAPI.md#getimportlistconfig) | **Get** /api/v3/config/importlist | 
*ImportListConfigAPI* | [**GetImportListConfigById**](docs/ImportListConfigAPI.md#getimportlistconfigbyid) | **Get** /api/v3/config/importlist/{id} | 
*ImportListConfigAPI* | [**UpdateImportListConfig**](docs/ImportListConfigAPI.md#updateimportlistconfig) | **Put** /api/v3/config/importlist/{id} | 
*ImportListExclusionAPI* | [**CreateImportListExclusion**](docs/ImportListExclusionAPI.md#createimportlistexclusion) | **Post** /api/v3/importlistexclusion | 
*ImportListExclusionAPI* | [**DeleteImportListExclusion**](docs/ImportListExclusionAPI.md#deleteimportlistexclusion) | **Delete** /api/v3/importlistexclusion/{id} | 
*ImportListExclusionAPI* | [**DeleteImportListExclusionBulk**](docs/ImportListExclusionAPI.md#deleteimportlistexclusionbulk) | **Delete** /api/v3/importlistexclusion/bulk | 
*ImportListExclusionAPI* | [**GetImportListExclusionById**](docs/ImportListExclusionAPI.md#getimportlistexclusionbyid) | **Get** /api/v3/importlistexclusion/{id} | 
*ImportListExclusionAPI* | [**GetImportListExclusionPaged**](docs/ImportListExclusionAPI.md#getimportlistexclusionpaged) | **Get** /api/v3/importlistexclusion/paged | 
*ImportListExclusionAPI* | [**ListImportListExclusion**](docs/ImportListExclusionAPI.md#listimportlistexclusion) | **Get** /api/v3/importlistexclusion | 
*ImportListExclusionAPI* | [**UpdateImportListExclusion**](docs/ImportListExclusionAPI.md#updateimportlistexclusion) | **Put** /api/v3/importlistexclusion/{id} | 
*IndexerAPI* | [**CreateIndexer**](docs/IndexerAPI.md#createindexer) | **Post** /api/v3/indexer | 
*IndexerAPI* | [**CreateIndexerActionByName**](docs/IndexerAPI.md#createindexeractionbyname) | **Post** /api/v3/indexer/action/{name} | 
*IndexerAPI* | [**DeleteIndexer**](docs/IndexerAPI.md#deleteindexer) | **Delete** /api/v3/indexer/{id} | 
*IndexerAPI* | [**DeleteIndexerBulk**](docs/IndexerAPI.md#deleteindexerbulk) | **Delete** /api/v3/indexer/bulk | 
*IndexerAPI* | [**GetIndexerById**](docs/IndexerAPI.md#getindexerbyid) | **Get** /api/v3/indexer/{id} | 
*IndexerAPI* | [**ListIndexer**](docs/IndexerAPI.md#listindexer) | **Get** /api/v3/indexer | 
*IndexerAPI* | [**ListIndexerSchema**](docs/IndexerAPI.md#listindexerschema) | **Get** /api/v3/indexer/schema | 
*IndexerAPI* | [**PutIndexerBulk**](docs/IndexerAPI.md#putindexerbulk) | **Put** /api/v3/indexer/bulk | 
*IndexerAPI* | [**TestIndexer**](docs/IndexerAPI.md#testindexer) | **Post** /api/v3/indexer/test | 
*IndexerAPI* | [**TestallIndexer**](docs/IndexerAPI.md#testallindexer) | **Post** /api/v3/indexer/testall | 
*IndexerAPI* | [**UpdateIndexer**](docs/IndexerAPI.md#updateindexer) | **Put** /api/v3/indexer/{id} | 
*IndexerConfigAPI* | [**GetIndexerConfig**](docs/IndexerConfigAPI.md#getindexerconfig) | **Get** /api/v3/config/indexer | 
*IndexerConfigAPI* | [**GetIndexerConfigById**](docs/IndexerConfigAPI.md#getindexerconfigbyid) | **Get** /api/v3/config/indexer/{id} | 
*IndexerConfigAPI* | [**UpdateIndexerConfig**](docs/IndexerConfigAPI.md#updateindexerconfig) | **Put** /api/v3/config/indexer/{id} | 
*IndexerFlagAPI* | [**ListIndexerFlag**](docs/IndexerFlagAPI.md#listindexerflag) | **Get** /api/v3/indexerflag | 
*LanguageAPI* | [**GetLanguageById**](docs/LanguageAPI.md#getlanguagebyid) | **Get** /api/v3/language/{id} | 
*LanguageAPI* | [**ListLanguage**](docs/LanguageAPI.md#listlanguage) | **Get** /api/v3/language | 
*LanguageProfileAPI* | [**CreateLanguageProfile**](docs/LanguageProfileAPI.md#createlanguageprofile) | **Post** /api/v3/languageprofile | 
*LanguageProfileAPI* | [**DeleteLanguageProfile**](docs/LanguageProfileAPI.md#deletelanguageprofile) | **Delete** /api/v3/languageprofile/{id} | 
*LanguageProfileAPI* | [**GetLanguageProfileById**](docs/LanguageProfileAPI.md#getlanguageprofilebyid) | **Get** /api/v3/languageprofile/{id} | 
*LanguageProfileAPI* | [**ListLanguageProfile**](docs/LanguageProfileAPI.md#listlanguageprofile) | **Get** /api/v3/languageprofile | 
*LanguageProfileAPI* | [**UpdateLanguageProfile**](docs/LanguageProfileAPI.md#updatelanguageprofile) | **Put** /api/v3/languageprofile/{id} | 
*LanguageProfileSchemaAPI* | [**GetLanguageprofileSchema**](docs/LanguageProfileSchemaAPI.md#getlanguageprofileschema) | **Get** /api/v3/languageprofile/schema | 
*LocalizationAPI* | [**GetLocalization**](docs/LocalizationAPI.md#getlocalization) | **Get** /api/v3/localization | 
*LocalizationAPI* | [**GetLocalizationById**](docs/LocalizationAPI.md#getlocalizationbyid) | **Get** /api/v3/localization/{id} | 
*LocalizationAPI* | [**GetLocalizationLanguage**](docs/LocalizationAPI.md#getlocalizationlanguage) | **Get** /api/v3/localization/language | 
*LogAPI* | [**GetLog**](docs/LogAPI.md#getlog) | **Get** /api/v3/log | 
*LogFileAPI* | [**GetLogFileByFilename**](docs/LogFileAPI.md#getlogfilebyfilename) | **Get** /api/v3/log/file/{filename} | 
*LogFileAPI* | [**ListLogFile**](docs/LogFileAPI.md#listlogfile) | **Get** /api/v3/log/file | 
*ManualImportAPI* | [**CreateManualImport**](docs/ManualImportAPI.md#createmanualimport) | **Post** /api/v3/manualimport | 
*ManualImportAPI* | [**ListManualImport**](docs/ManualImportAPI.md#listmanualimport) | **Get** /api/v3/manualimport | 
*MediaCoverAPI* | [**GetMediaCoverByFilename**](docs/MediaCoverAPI.md#getmediacoverbyfilename) | **Get** /api/v3/mediacover/{seriesId}/{filename} | 
*MediaManagementConfigAPI* | [**GetMediaManagementConfig**](docs/MediaManagementConfigAPI.md#getmediamanagementconfig) | **Get** /api/v3/config/mediamanagement | 
*MediaManagementConfigAPI* | [**GetMediaManagementConfigById**](docs/MediaManagementConfigAPI.md#getmediamanagementconfigbyid) | **Get** /api/v3/config/mediamanagement/{id} | 
*MediaManagementConfigAPI* | [**UpdateMediaManagementConfig**](docs/MediaManagementConfigAPI.md#updatemediamanagementconfig) | **Put** /api/v3/config/mediamanagement/{id} | 
*MetadataAPI* | [**CreateMetadata**](docs/MetadataAPI.md#createmetadata) | **Post** /api/v3/metadata | 
*MetadataAPI* | [**CreateMetadataActionByName**](docs/MetadataAPI.md#createmetadataactionbyname) | **Post** /api/v3/metadata/action/{name} | 
*MetadataAPI* | [**DeleteMetadata**](docs/MetadataAPI.md#deletemetadata) | **Delete** /api/v3/metadata/{id} | 
*MetadataAPI* | [**GetMetadataById**](docs/MetadataAPI.md#getmetadatabyid) | **Get** /api/v3/metadata/{id} | 
*MetadataAPI* | [**ListMetadata**](docs/MetadataAPI.md#listmetadata) | **Get** /api/v3/metadata | 
*MetadataAPI* | [**ListMetadataSchema**](docs/MetadataAPI.md#listmetadataschema) | **Get** /api/v3/metadata/schema | 
*MetadataAPI* | [**TestMetadata**](docs/MetadataAPI.md#testmetadata) | **Post** /api/v3/metadata/test | 
*MetadataAPI* | [**TestallMetadata**](docs/MetadataAPI.md#testallmetadata) | **Post** /api/v3/metadata/testall | 
*MetadataAPI* | [**UpdateMetadata**](docs/MetadataAPI.md#updatemetadata) | **Put** /api/v3/metadata/{id} | 
*MissingAPI* | [**GetWantedMissing**](docs/MissingAPI.md#getwantedmissing) | **Get** /api/v3/wanted/missing | 
*MissingAPI* | [**GetWantedMissingById**](docs/MissingAPI.md#getwantedmissingbyid) | **Get** /api/v3/wanted/missing/{id} | 
*NamingConfigAPI* | [**GetNamingConfig**](docs/NamingConfigAPI.md#getnamingconfig) | **Get** /api/v3/config/naming | 
*NamingConfigAPI* | [**GetNamingConfigById**](docs/NamingConfigAPI.md#getnamingconfigbyid) | **Get** /api/v3/config/naming/{id} | 
*NamingConfigAPI* | [**GetNamingConfigExamples**](docs/NamingConfigAPI.md#getnamingconfigexamples) | **Get** /api/v3/config/naming/examples | 
*NamingConfigAPI* | [**UpdateNamingConfig**](docs/NamingConfigAPI.md#updatenamingconfig) | **Put** /api/v3/config/naming/{id} | 
*NotificationAPI* | [**CreateNotification**](docs/NotificationAPI.md#createnotification) | **Post** /api/v3/notification | 
*NotificationAPI* | [**CreateNotificationActionByName**](docs/NotificationAPI.md#createnotificationactionbyname) | **Post** /api/v3/notification/action/{name} | 
*NotificationAPI* | [**DeleteNotification**](docs/NotificationAPI.md#deletenotification) | **Delete** /api/v3/notification/{id} | 
*NotificationAPI* | [**GetNotificationById**](docs/NotificationAPI.md#getnotificationbyid) | **Get** /api/v3/notification/{id} | 
*NotificationAPI* | [**ListNotification**](docs/NotificationAPI.md#listnotification) | **Get** /api/v3/notification | 
*NotificationAPI* | [**ListNotificationSchema**](docs/NotificationAPI.md#listnotificationschema) | **Get** /api/v3/notification/schema | 
*NotificationAPI* | [**TestNotification**](docs/NotificationAPI.md#testnotification) | **Post** /api/v3/notification/test | 
*NotificationAPI* | [**TestallNotification**](docs/NotificationAPI.md#testallnotification) | **Post** /api/v3/notification/testall | 
*NotificationAPI* | [**UpdateNotification**](docs/NotificationAPI.md#updatenotification) | **Put** /api/v3/notification/{id} | 
*ParseAPI* | [**GetParse**](docs/ParseAPI.md#getparse) | **Get** /api/v3/parse | 
*PingAPI* | [**GetPing**](docs/PingAPI.md#getping) | **Get** /ping | 
*PingAPI* | [**HeadPing**](docs/PingAPI.md#headping) | **Head** /ping | 
*QualityDefinitionAPI* | [**GetQualityDefinitionById**](docs/QualityDefinitionAPI.md#getqualitydefinitionbyid) | **Get** /api/v3/qualitydefinition/{id} | 
*QualityDefinitionAPI* | [**ListQualityDefinition**](docs/QualityDefinitionAPI.md#listqualitydefinition) | **Get** /api/v3/qualitydefinition | 
*QualityDefinitionAPI* | [**PutQualityDefinitionUpdate**](docs/QualityDefinitionAPI.md#putqualitydefinitionupdate) | **Put** /api/v3/qualitydefinition/update | 
*QualityDefinitionAPI* | [**UpdateQualityDefinition**](docs/QualityDefinitionAPI.md#updatequalitydefinition) | **Put** /api/v3/qualitydefinition/{id} | 
*QualityProfileAPI* | [**CreateQualityProfile**](docs/QualityProfileAPI.md#createqualityprofile) | **Post** /api/v3/qualityprofile | 
*QualityProfileAPI* | [**DeleteQualityProfile**](docs/QualityProfileAPI.md#deletequalityprofile) | **Delete** /api/v3/qualityprofile/{id} | 
*QualityProfileAPI* | [**GetQualityProfileById**](docs/QualityProfileAPI.md#getqualityprofilebyid) | **Get** /api/v3/qualityprofile/{id} | 
*QualityProfileAPI* | [**ListQualityProfile**](docs/QualityProfileAPI.md#listqualityprofile) | **Get** /api/v3/qualityprofile | 
*QualityProfileAPI* | [**UpdateQualityProfile**](docs/QualityProfileAPI.md#updatequalityprofile) | **Put** /api/v3/qualityprofile/{id} | 
*QualityProfileSchemaAPI* | [**GetQualityprofileSchema**](docs/QualityProfileSchemaAPI.md#getqualityprofileschema) | **Get** /api/v3/qualityprofile/schema | 
*QueueAPI* | [**DeleteQueue**](docs/QueueAPI.md#deletequeue) | **Delete** /api/v3/queue/{id} | 
*QueueAPI* | [**DeleteQueueBulk**](docs/QueueAPI.md#deletequeuebulk) | **Delete** /api/v3/queue/bulk | 
*QueueAPI* | [**GetQueue**](docs/QueueAPI.md#getqueue) | **Get** /api/v3/queue | 
*QueueActionAPI* | [**CreateQueueGrabBulk**](docs/QueueActionAPI.md#createqueuegrabbulk) | **Post** /api/v3/queue/grab/bulk | 
*QueueActionAPI* | [**CreateQueueGrabById**](docs/QueueActionAPI.md#createqueuegrabbyid) | **Post** /api/v3/queue/grab/{id} | 
*QueueDetailsAPI* | [**ListQueueDetails**](docs/QueueDetailsAPI.md#listqueuedetails) | **Get** /api/v3/queue/details | 
*QueueStatusAPI* | [**GetQueueStatus**](docs/QueueStatusAPI.md#getqueuestatus) | **Get** /api/v3/queue/status | 
*ReleaseAPI* | [**CreateRelease**](docs/ReleaseAPI.md#createrelease) | **Post** /api/v3/release | 
*ReleaseAPI* | [**ListRelease**](docs/ReleaseAPI.md#listrelease) | **Get** /api/v3/release | 
*ReleaseProfileAPI* | [**CreateReleaseProfile**](docs/ReleaseProfileAPI.md#createreleaseprofile) | **Post** /api/v3/releaseprofile | 
*ReleaseProfileAPI* | [**DeleteReleaseProfile**](docs/ReleaseProfileAPI.md#deletereleaseprofile) | **Delete** /api/v3/releaseprofile/{id} | 
*ReleaseProfileAPI* | [**GetReleaseProfileById**](docs/ReleaseProfileAPI.md#getreleaseprofilebyid) | **Get** /api/v3/releaseprofile/{id} | 
*ReleaseProfileAPI* | [**ListReleaseProfile**](docs/ReleaseProfileAPI.md#listreleaseprofile) | **Get** /api/v3/releaseprofile | 
*ReleaseProfileAPI* | [**UpdateReleaseProfile**](docs/ReleaseProfileAPI.md#updatereleaseprofile) | **Put** /api/v3/releaseprofile/{id} | 
*ReleasePushAPI* | [**CreateReleasePush**](docs/ReleasePushAPI.md#createreleasepush) | **Post** /api/v3/release/push | 
*RemotePathMappingAPI* | [**CreateRemotePathMapping**](docs/RemotePathMappingAPI.md#createremotepathmapping) | **Post** /api/v3/remotepathmapping | 
*RemotePathMappingAPI* | [**DeleteRemotePathMapping**](docs/RemotePathMappingAPI.md#deleteremotepathmapping) | **Delete** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingAPI* | [**GetRemotePathMappingById**](docs/RemotePathMappingAPI.md#getremotepathmappingbyid) | **Get** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingAPI* | [**ListRemotePathMapping**](docs/RemotePathMappingAPI.md#listremotepathmapping) | **Get** /api/v3/remotepathmapping | 
*RemotePathMappingAPI* | [**UpdateRemotePathMapping**](docs/RemotePathMappingAPI.md#updateremotepathmapping) | **Put** /api/v3/remotepathmapping/{id} | 
*RenameEpisodeAPI* | [**ListRename**](docs/RenameEpisodeAPI.md#listrename) | **Get** /api/v3/rename | 
*RootFolderAPI* | [**CreateRootFolder**](docs/RootFolderAPI.md#createrootfolder) | **Post** /api/v3/rootfolder | 
*RootFolderAPI* | [**DeleteRootFolder**](docs/RootFolderAPI.md#deleterootfolder) | **Delete** /api/v3/rootfolder/{id} | 
*RootFolderAPI* | [**GetRootFolderById**](docs/RootFolderAPI.md#getrootfolderbyid) | **Get** /api/v3/rootfolder/{id} | 
*RootFolderAPI* | [**ListRootFolder**](docs/RootFolderAPI.md#listrootfolder) | **Get** /api/v3/rootfolder | 
*SeasonPassAPI* | [**CreateSeasonPass**](docs/SeasonPassAPI.md#createseasonpass) | **Post** /api/v3/seasonpass | 
*SeriesAPI* | [**CreateSeries**](docs/SeriesAPI.md#createseries) | **Post** /api/v3/series | 
*SeriesAPI* | [**DeleteSeries**](docs/SeriesAPI.md#deleteseries) | **Delete** /api/v3/series/{id} | 
*SeriesAPI* | [**GetSeriesById**](docs/SeriesAPI.md#getseriesbyid) | **Get** /api/v3/series/{id} | 
*SeriesAPI* | [**ListSeries**](docs/SeriesAPI.md#listseries) | **Get** /api/v3/series | 
*SeriesAPI* | [**UpdateSeries**](docs/SeriesAPI.md#updateseries) | **Put** /api/v3/series/{id} | 
*SeriesEditorAPI* | [**DeleteSeriesEditor**](docs/SeriesEditorAPI.md#deleteserieseditor) | **Delete** /api/v3/series/editor | 
*SeriesEditorAPI* | [**PutSeriesEditor**](docs/SeriesEditorAPI.md#putserieseditor) | **Put** /api/v3/series/editor | 
*SeriesImportAPI* | [**CreateSeriesImport**](docs/SeriesImportAPI.md#createseriesimport) | **Post** /api/v3/series/import | 
*SeriesLookupAPI* | [**ListSeriesLookup**](docs/SeriesLookupAPI.md#listserieslookup) | **Get** /api/v3/series/lookup | 
*StaticResourceAPI* | [**Get**](docs/StaticResourceAPI.md#get) | **Get** / | 
*StaticResourceAPI* | [**GetByPath**](docs/StaticResourceAPI.md#getbypath) | **Get** /{path} | 
*StaticResourceAPI* | [**GetContentByPath**](docs/StaticResourceAPI.md#getcontentbypath) | **Get** /content/{path} | 
*StaticResourceAPI* | [**GetLogin**](docs/StaticResourceAPI.md#getlogin) | **Get** /login | 
*SystemAPI* | [**CreateSystemRestart**](docs/SystemAPI.md#createsystemrestart) | **Post** /api/v3/system/restart | 
*SystemAPI* | [**CreateSystemShutdown**](docs/SystemAPI.md#createsystemshutdown) | **Post** /api/v3/system/shutdown | 
*SystemAPI* | [**GetSystemRoutes**](docs/SystemAPI.md#getsystemroutes) | **Get** /api/v3/system/routes | 
*SystemAPI* | [**GetSystemRoutesDuplicate**](docs/SystemAPI.md#getsystemroutesduplicate) | **Get** /api/v3/system/routes/duplicate | 
*SystemAPI* | [**GetSystemStatus**](docs/SystemAPI.md#getsystemstatus) | **Get** /api/v3/system/status | 
*TagAPI* | [**CreateTag**](docs/TagAPI.md#createtag) | **Post** /api/v3/tag | 
*TagAPI* | [**DeleteTag**](docs/TagAPI.md#deletetag) | **Delete** /api/v3/tag/{id} | 
*TagAPI* | [**GetTagById**](docs/TagAPI.md#gettagbyid) | **Get** /api/v3/tag/{id} | 
*TagAPI* | [**ListTag**](docs/TagAPI.md#listtag) | **Get** /api/v3/tag | 
*TagAPI* | [**UpdateTag**](docs/TagAPI.md#updatetag) | **Put** /api/v3/tag/{id} | 
*TagDetailsAPI* | [**GetTagDetailById**](docs/TagDetailsAPI.md#gettagdetailbyid) | **Get** /api/v3/tag/detail/{id} | 
*TagDetailsAPI* | [**ListTagDetail**](docs/TagDetailsAPI.md#listtagdetail) | **Get** /api/v3/tag/detail | 
*TaskAPI* | [**GetSystemTaskById**](docs/TaskAPI.md#getsystemtaskbyid) | **Get** /api/v3/system/task/{id} | 
*TaskAPI* | [**ListSystemTask**](docs/TaskAPI.md#listsystemtask) | **Get** /api/v3/system/task | 
*UiConfigAPI* | [**GetUiConfig**](docs/UiConfigAPI.md#getuiconfig) | **Get** /api/v3/config/ui | 
*UiConfigAPI* | [**GetUiConfigById**](docs/UiConfigAPI.md#getuiconfigbyid) | **Get** /api/v3/config/ui/{id} | 
*UiConfigAPI* | [**UpdateUiConfig**](docs/UiConfigAPI.md#updateuiconfig) | **Put** /api/v3/config/ui/{id} | 
*UpdateAPI* | [**ListUpdate**](docs/UpdateAPI.md#listupdate) | **Get** /api/v3/update | 
*UpdateLogFileAPI* | [**GetLogFileUpdateByFilename**](docs/UpdateLogFileAPI.md#getlogfileupdatebyfilename) | **Get** /api/v3/log/file/update/{filename} | 
*UpdateLogFileAPI* | [**ListLogFileUpdate**](docs/UpdateLogFileAPI.md#listlogfileupdate) | **Get** /api/v3/log/file/update | 


## Documentation For Models

 - [AddSeriesOptions](docs/AddSeriesOptions.md)
 - [AlternateTitleResource](docs/AlternateTitleResource.md)
 - [ApplyTags](docs/ApplyTags.md)
 - [AuthenticationRequiredType](docs/AuthenticationRequiredType.md)
 - [AuthenticationType](docs/AuthenticationType.md)
 - [AutoTaggingResource](docs/AutoTaggingResource.md)
 - [AutoTaggingSpecificationSchema](docs/AutoTaggingSpecificationSchema.md)
 - [BackupResource](docs/BackupResource.md)
 - [BackupType](docs/BackupType.md)
 - [BlocklistBulkResource](docs/BlocklistBulkResource.md)
 - [BlocklistResource](docs/BlocklistResource.md)
 - [BlocklistResourcePagingResource](docs/BlocklistResourcePagingResource.md)
 - [CertificateValidationType](docs/CertificateValidationType.md)
 - [Command](docs/Command.md)
 - [CommandPriority](docs/CommandPriority.md)
 - [CommandResource](docs/CommandResource.md)
 - [CommandResult](docs/CommandResult.md)
 - [CommandStatus](docs/CommandStatus.md)
 - [CommandTrigger](docs/CommandTrigger.md)
 - [CustomFilterResource](docs/CustomFilterResource.md)
 - [CustomFormatResource](docs/CustomFormatResource.md)
 - [CustomFormatSpecificationSchema](docs/CustomFormatSpecificationSchema.md)
 - [DatabaseType](docs/DatabaseType.md)
 - [DelayProfileResource](docs/DelayProfileResource.md)
 - [DiskSpaceResource](docs/DiskSpaceResource.md)
 - [DownloadClientBulkResource](docs/DownloadClientBulkResource.md)
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
 - [ImportListBulkResource](docs/ImportListBulkResource.md)
 - [ImportListConfigResource](docs/ImportListConfigResource.md)
 - [ImportListExclusionBulkResource](docs/ImportListExclusionBulkResource.md)
 - [ImportListExclusionResource](docs/ImportListExclusionResource.md)
 - [ImportListExclusionResourcePagingResource](docs/ImportListExclusionResourcePagingResource.md)
 - [ImportListResource](docs/ImportListResource.md)
 - [ImportListType](docs/ImportListType.md)
 - [IndexerBulkResource](docs/IndexerBulkResource.md)
 - [IndexerConfigResource](docs/IndexerConfigResource.md)
 - [IndexerFlagResource](docs/IndexerFlagResource.md)
 - [IndexerResource](docs/IndexerResource.md)
 - [Language](docs/Language.md)
 - [LanguageProfileItemResource](docs/LanguageProfileItemResource.md)
 - [LanguageProfileResource](docs/LanguageProfileResource.md)
 - [LanguageResource](docs/LanguageResource.md)
 - [ListSyncLevelType](docs/ListSyncLevelType.md)
 - [LocalizationLanguageResource](docs/LocalizationLanguageResource.md)
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
 - [NewItemMonitorTypes](docs/NewItemMonitorTypes.md)
 - [NotificationResource](docs/NotificationResource.md)
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
 - [ReleaseEpisodeResource](docs/ReleaseEpisodeResource.md)
 - [ReleaseProfileResource](docs/ReleaseProfileResource.md)
 - [ReleaseResource](docs/ReleaseResource.md)
 - [ReleaseType](docs/ReleaseType.md)
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
 - [TrackedDownloadState](docs/TrackedDownloadState.md)
 - [TrackedDownloadStatus](docs/TrackedDownloadStatus.md)
 - [TrackedDownloadStatusMessage](docs/TrackedDownloadStatusMessage.md)
 - [UiConfigResource](docs/UiConfigResource.md)
 - [UnmappedFolder](docs/UnmappedFolder.md)
 - [UpdateChanges](docs/UpdateChanges.md)
 - [UpdateMechanism](docs/UpdateMechanism.md)
 - [UpdateResource](docs/UpdateResource.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### X-Api-Key

- **Type**: API key
- **API key parameter name**: X-Api-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Api-Key and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		sonarr.ContextAPIKeys,
		map[string]sonarr.APIKey{
			"X-Api-Key": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### apikey

- **Type**: API key
- **API key parameter name**: apikey
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: apikey and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		sonarr.ContextAPIKeys,
		map[string]sonarr.APIKey{
			"apikey": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


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



