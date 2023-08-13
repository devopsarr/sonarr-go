# Go API client for sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

[comment]: # (x-release-please-start-version)
- Package version: 0.6.0

[comment]: # (x-release-please-end)
- API version: 3.0.0

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
*ApiInfoApi* | [**GetApi**](sonarr/docs/ApiInfoApi.md#getapi) | **Get** /api | 
*AuthenticationApi* | [**CreateLogin**](sonarr/docs/AuthenticationApi.md#createlogin) | **Post** /login | 
*AuthenticationApi* | [**GetLogout**](sonarr/docs/AuthenticationApi.md#getlogout) | **Get** /logout | 
*AutoTaggingApi* | [**CreateAutoTagging**](sonarr/docs/AutoTaggingApi.md#createautotagging) | **Post** /api/v3/autotagging | 
*AutoTaggingApi* | [**DeleteAutoTagging**](sonarr/docs/AutoTaggingApi.md#deleteautotagging) | **Delete** /api/v3/autotagging/{id} | 
*AutoTaggingApi* | [**GetAutoTaggingById**](sonarr/docs/AutoTaggingApi.md#getautotaggingbyid) | **Get** /api/v3/autotagging/{id} | 
*AutoTaggingApi* | [**GetAutoTaggingSchema**](sonarr/docs/AutoTaggingApi.md#getautotaggingschema) | **Get** /api/v3/autotagging/schema | 
*AutoTaggingApi* | [**ListAutoTagging**](sonarr/docs/AutoTaggingApi.md#listautotagging) | **Get** /api/v3/autotagging | 
*AutoTaggingApi* | [**UpdateAutoTagging**](sonarr/docs/AutoTaggingApi.md#updateautotagging) | **Put** /api/v3/autotagging/{id} | 
*BackupApi* | [**CreateSystemBackupRestoreById**](sonarr/docs/BackupApi.md#createsystembackuprestorebyid) | **Post** /api/v3/system/backup/restore/{id} | 
*BackupApi* | [**CreateSystemBackupRestoreUpload**](sonarr/docs/BackupApi.md#createsystembackuprestoreupload) | **Post** /api/v3/system/backup/restore/upload | 
*BackupApi* | [**DeleteSystemBackup**](sonarr/docs/BackupApi.md#deletesystembackup) | **Delete** /api/v3/system/backup/{id} | 
*BackupApi* | [**ListSystemBackup**](sonarr/docs/BackupApi.md#listsystembackup) | **Get** /api/v3/system/backup | 
*BlocklistApi* | [**DeleteBlocklist**](sonarr/docs/BlocklistApi.md#deleteblocklist) | **Delete** /api/v3/blocklist/{id} | 
*BlocklistApi* | [**DeleteBlocklistBulk**](sonarr/docs/BlocklistApi.md#deleteblocklistbulk) | **Delete** /api/v3/blocklist/bulk | 
*BlocklistApi* | [**GetBlocklist**](sonarr/docs/BlocklistApi.md#getblocklist) | **Get** /api/v3/blocklist | 
*CalendarApi* | [**GetCalendarById**](sonarr/docs/CalendarApi.md#getcalendarbyid) | **Get** /api/v3/calendar/{id} | 
*CalendarApi* | [**ListCalendar**](sonarr/docs/CalendarApi.md#listcalendar) | **Get** /api/v3/calendar | 
*CalendarFeedApi* | [**GetFeedV3CalendarSonarrIcs**](sonarr/docs/CalendarFeedApi.md#getfeedv3calendarsonarrics) | **Get** /feed/v3/calendar/sonarr.ics | 
*CommandApi* | [**CreateCommand**](sonarr/docs/CommandApi.md#createcommand) | **Post** /api/v3/command | 
*CommandApi* | [**DeleteCommand**](sonarr/docs/CommandApi.md#deletecommand) | **Delete** /api/v3/command/{id} | 
*CommandApi* | [**GetCommandById**](sonarr/docs/CommandApi.md#getcommandbyid) | **Get** /api/v3/command/{id} | 
*CommandApi* | [**ListCommand**](sonarr/docs/CommandApi.md#listcommand) | **Get** /api/v3/command | 
*CustomFilterApi* | [**CreateCustomFilter**](sonarr/docs/CustomFilterApi.md#createcustomfilter) | **Post** /api/v3/customfilter | 
*CustomFilterApi* | [**DeleteCustomFilter**](sonarr/docs/CustomFilterApi.md#deletecustomfilter) | **Delete** /api/v3/customfilter/{id} | 
*CustomFilterApi* | [**GetCustomFilterById**](sonarr/docs/CustomFilterApi.md#getcustomfilterbyid) | **Get** /api/v3/customfilter/{id} | 
*CustomFilterApi* | [**ListCustomFilter**](sonarr/docs/CustomFilterApi.md#listcustomfilter) | **Get** /api/v3/customfilter | 
*CustomFilterApi* | [**UpdateCustomFilter**](sonarr/docs/CustomFilterApi.md#updatecustomfilter) | **Put** /api/v3/customfilter/{id} | 
*CustomFormatApi* | [**CreateCustomFormat**](sonarr/docs/CustomFormatApi.md#createcustomformat) | **Post** /api/v3/customformat | 
*CustomFormatApi* | [**DeleteCustomFormat**](sonarr/docs/CustomFormatApi.md#deletecustomformat) | **Delete** /api/v3/customformat/{id} | 
*CustomFormatApi* | [**GetCustomFormatById**](sonarr/docs/CustomFormatApi.md#getcustomformatbyid) | **Get** /api/v3/customformat/{id} | 
*CustomFormatApi* | [**ListCustomFormat**](sonarr/docs/CustomFormatApi.md#listcustomformat) | **Get** /api/v3/customformat | 
*CustomFormatApi* | [**ListCustomFormatSchema**](sonarr/docs/CustomFormatApi.md#listcustomformatschema) | **Get** /api/v3/customformat/schema | 
*CustomFormatApi* | [**UpdateCustomFormat**](sonarr/docs/CustomFormatApi.md#updatecustomformat) | **Put** /api/v3/customformat/{id} | 
*CutoffApi* | [**GetWantedCutoff**](sonarr/docs/CutoffApi.md#getwantedcutoff) | **Get** /api/v3/wanted/cutoff | 
*CutoffApi* | [**GetWantedCutoffById**](sonarr/docs/CutoffApi.md#getwantedcutoffbyid) | **Get** /api/v3/wanted/cutoff/{id} | 
*DelayProfileApi* | [**CreateDelayProfile**](sonarr/docs/DelayProfileApi.md#createdelayprofile) | **Post** /api/v3/delayprofile | 
*DelayProfileApi* | [**DeleteDelayProfile**](sonarr/docs/DelayProfileApi.md#deletedelayprofile) | **Delete** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**GetDelayProfileById**](sonarr/docs/DelayProfileApi.md#getdelayprofilebyid) | **Get** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**ListDelayProfile**](sonarr/docs/DelayProfileApi.md#listdelayprofile) | **Get** /api/v3/delayprofile | 
*DelayProfileApi* | [**UpdateDelayProfile**](sonarr/docs/DelayProfileApi.md#updatedelayprofile) | **Put** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**UpdateDelayProfileReorder**](sonarr/docs/DelayProfileApi.md#updatedelayprofilereorder) | **Put** /api/v3/delayprofile/reorder/{id} | 
*DiskSpaceApi* | [**ListDiskSpace**](sonarr/docs/DiskSpaceApi.md#listdiskspace) | **Get** /api/v3/diskspace | 
*DownloadClientApi* | [**CreateDownloadClient**](sonarr/docs/DownloadClientApi.md#createdownloadclient) | **Post** /api/v3/downloadclient | 
*DownloadClientApi* | [**CreateDownloadClientActionByName**](sonarr/docs/DownloadClientApi.md#createdownloadclientactionbyname) | **Post** /api/v3/downloadclient/action/{name} | 
*DownloadClientApi* | [**DeleteDownloadClient**](sonarr/docs/DownloadClientApi.md#deletedownloadclient) | **Delete** /api/v3/downloadclient/{id} | 
*DownloadClientApi* | [**DeleteDownloadClientBulk**](sonarr/docs/DownloadClientApi.md#deletedownloadclientbulk) | **Delete** /api/v3/downloadclient/bulk | 
*DownloadClientApi* | [**GetDownloadClientById**](sonarr/docs/DownloadClientApi.md#getdownloadclientbyid) | **Get** /api/v3/downloadclient/{id} | 
*DownloadClientApi* | [**ListDownloadClient**](sonarr/docs/DownloadClientApi.md#listdownloadclient) | **Get** /api/v3/downloadclient | 
*DownloadClientApi* | [**ListDownloadClientSchema**](sonarr/docs/DownloadClientApi.md#listdownloadclientschema) | **Get** /api/v3/downloadclient/schema | 
*DownloadClientApi* | [**PutDownloadClientBulk**](sonarr/docs/DownloadClientApi.md#putdownloadclientbulk) | **Put** /api/v3/downloadclient/bulk | 
*DownloadClientApi* | [**TestDownloadClient**](sonarr/docs/DownloadClientApi.md#testdownloadclient) | **Post** /api/v3/downloadclient/test | 
*DownloadClientApi* | [**TestallDownloadClient**](sonarr/docs/DownloadClientApi.md#testalldownloadclient) | **Post** /api/v3/downloadclient/testall | 
*DownloadClientApi* | [**UpdateDownloadClient**](sonarr/docs/DownloadClientApi.md#updatedownloadclient) | **Put** /api/v3/downloadclient/{id} | 
*DownloadClientConfigApi* | [**GetDownloadClientConfig**](sonarr/docs/DownloadClientConfigApi.md#getdownloadclientconfig) | **Get** /api/v3/config/downloadclient | 
*DownloadClientConfigApi* | [**GetDownloadClientConfigById**](sonarr/docs/DownloadClientConfigApi.md#getdownloadclientconfigbyid) | **Get** /api/v3/config/downloadclient/{id} | 
*DownloadClientConfigApi* | [**UpdateDownloadClientConfig**](sonarr/docs/DownloadClientConfigApi.md#updatedownloadclientconfig) | **Put** /api/v3/config/downloadclient/{id} | 
*EpisodeApi* | [**GetEpisodeById**](sonarr/docs/EpisodeApi.md#getepisodebyid) | **Get** /api/v3/episode/{id} | 
*EpisodeApi* | [**ListEpisode**](sonarr/docs/EpisodeApi.md#listepisode) | **Get** /api/v3/episode | 
*EpisodeApi* | [**PutEpisodeMonitor**](sonarr/docs/EpisodeApi.md#putepisodemonitor) | **Put** /api/v3/episode/monitor | 
*EpisodeApi* | [**UpdateEpisode**](sonarr/docs/EpisodeApi.md#updateepisode) | **Put** /api/v3/episode/{id} | 
*EpisodeFileApi* | [**DeleteEpisodeFile**](sonarr/docs/EpisodeFileApi.md#deleteepisodefile) | **Delete** /api/v3/episodefile/{id} | 
*EpisodeFileApi* | [**DeleteEpisodeFileBulk**](sonarr/docs/EpisodeFileApi.md#deleteepisodefilebulk) | **Delete** /api/v3/episodefile/bulk | 
*EpisodeFileApi* | [**GetEpisodeFileById**](sonarr/docs/EpisodeFileApi.md#getepisodefilebyid) | **Get** /api/v3/episodefile/{id} | 
*EpisodeFileApi* | [**ListEpisodeFile**](sonarr/docs/EpisodeFileApi.md#listepisodefile) | **Get** /api/v3/episodefile | 
*EpisodeFileApi* | [**PutEpisodeFileBulk**](sonarr/docs/EpisodeFileApi.md#putepisodefilebulk) | **Put** /api/v3/episodefile/bulk | 
*EpisodeFileApi* | [**PutEpisodeFileEditor**](sonarr/docs/EpisodeFileApi.md#putepisodefileeditor) | **Put** /api/v3/episodefile/editor | 
*EpisodeFileApi* | [**UpdateEpisodeFile**](sonarr/docs/EpisodeFileApi.md#updateepisodefile) | **Put** /api/v3/episodefile/{id} | 
*FileSystemApi* | [**GetFileSystem**](sonarr/docs/FileSystemApi.md#getfilesystem) | **Get** /api/v3/filesystem | 
*FileSystemApi* | [**GetFileSystemMediafiles**](sonarr/docs/FileSystemApi.md#getfilesystemmediafiles) | **Get** /api/v3/filesystem/mediafiles | 
*FileSystemApi* | [**GetFileSystemType**](sonarr/docs/FileSystemApi.md#getfilesystemtype) | **Get** /api/v3/filesystem/type | 
*HealthApi* | [**ListHealth**](sonarr/docs/HealthApi.md#listhealth) | **Get** /api/v3/health | 
*HistoryApi* | [**CreateHistoryFailedById**](sonarr/docs/HistoryApi.md#createhistoryfailedbyid) | **Post** /api/v3/history/failed/{id} | 
*HistoryApi* | [**GetHistory**](sonarr/docs/HistoryApi.md#gethistory) | **Get** /api/v3/history | 
*HistoryApi* | [**ListHistorySeries**](sonarr/docs/HistoryApi.md#listhistoryseries) | **Get** /api/v3/history/series | 
*HistoryApi* | [**ListHistorySince**](sonarr/docs/HistoryApi.md#listhistorysince) | **Get** /api/v3/history/since | 
*HostConfigApi* | [**GetHostConfig**](sonarr/docs/HostConfigApi.md#gethostconfig) | **Get** /api/v3/config/host | 
*HostConfigApi* | [**GetHostConfigById**](sonarr/docs/HostConfigApi.md#gethostconfigbyid) | **Get** /api/v3/config/host/{id} | 
*HostConfigApi* | [**UpdateHostConfig**](sonarr/docs/HostConfigApi.md#updatehostconfig) | **Put** /api/v3/config/host/{id} | 
*ImportListApi* | [**CreateImportList**](sonarr/docs/ImportListApi.md#createimportlist) | **Post** /api/v3/importlist | 
*ImportListApi* | [**CreateImportListActionByName**](sonarr/docs/ImportListApi.md#createimportlistactionbyname) | **Post** /api/v3/importlist/action/{name} | 
*ImportListApi* | [**DeleteImportList**](sonarr/docs/ImportListApi.md#deleteimportlist) | **Delete** /api/v3/importlist/{id} | 
*ImportListApi* | [**DeleteImportListBulk**](sonarr/docs/ImportListApi.md#deleteimportlistbulk) | **Delete** /api/v3/importlist/bulk | 
*ImportListApi* | [**GetImportListById**](sonarr/docs/ImportListApi.md#getimportlistbyid) | **Get** /api/v3/importlist/{id} | 
*ImportListApi* | [**ListImportList**](sonarr/docs/ImportListApi.md#listimportlist) | **Get** /api/v3/importlist | 
*ImportListApi* | [**ListImportListSchema**](sonarr/docs/ImportListApi.md#listimportlistschema) | **Get** /api/v3/importlist/schema | 
*ImportListApi* | [**PutImportListBulk**](sonarr/docs/ImportListApi.md#putimportlistbulk) | **Put** /api/v3/importlist/bulk | 
*ImportListApi* | [**TestImportList**](sonarr/docs/ImportListApi.md#testimportlist) | **Post** /api/v3/importlist/test | 
*ImportListApi* | [**TestallImportList**](sonarr/docs/ImportListApi.md#testallimportlist) | **Post** /api/v3/importlist/testall | 
*ImportListApi* | [**UpdateImportList**](sonarr/docs/ImportListApi.md#updateimportlist) | **Put** /api/v3/importlist/{id} | 
*ImportListExclusionApi* | [**CreateImportListExclusion**](sonarr/docs/ImportListExclusionApi.md#createimportlistexclusion) | **Post** /api/v3/importlistexclusion | 
*ImportListExclusionApi* | [**DeleteImportListExclusion**](sonarr/docs/ImportListExclusionApi.md#deleteimportlistexclusion) | **Delete** /api/v3/importlistexclusion/{id} | 
*ImportListExclusionApi* | [**GetImportListExclusionById**](sonarr/docs/ImportListExclusionApi.md#getimportlistexclusionbyid) | **Get** /api/v3/importlistexclusion/{id} | 
*ImportListExclusionApi* | [**ListImportListExclusion**](sonarr/docs/ImportListExclusionApi.md#listimportlistexclusion) | **Get** /api/v3/importlistexclusion | 
*ImportListExclusionApi* | [**UpdateImportListExclusion**](sonarr/docs/ImportListExclusionApi.md#updateimportlistexclusion) | **Put** /api/v3/importlistexclusion/{id} | 
*IndexerApi* | [**CreateIndexer**](sonarr/docs/IndexerApi.md#createindexer) | **Post** /api/v3/indexer | 
*IndexerApi* | [**CreateIndexerActionByName**](sonarr/docs/IndexerApi.md#createindexeractionbyname) | **Post** /api/v3/indexer/action/{name} | 
*IndexerApi* | [**DeleteIndexer**](sonarr/docs/IndexerApi.md#deleteindexer) | **Delete** /api/v3/indexer/{id} | 
*IndexerApi* | [**DeleteIndexerBulk**](sonarr/docs/IndexerApi.md#deleteindexerbulk) | **Delete** /api/v3/indexer/bulk | 
*IndexerApi* | [**GetIndexerById**](sonarr/docs/IndexerApi.md#getindexerbyid) | **Get** /api/v3/indexer/{id} | 
*IndexerApi* | [**ListIndexer**](sonarr/docs/IndexerApi.md#listindexer) | **Get** /api/v3/indexer | 
*IndexerApi* | [**ListIndexerSchema**](sonarr/docs/IndexerApi.md#listindexerschema) | **Get** /api/v3/indexer/schema | 
*IndexerApi* | [**PutIndexerBulk**](sonarr/docs/IndexerApi.md#putindexerbulk) | **Put** /api/v3/indexer/bulk | 
*IndexerApi* | [**TestIndexer**](sonarr/docs/IndexerApi.md#testindexer) | **Post** /api/v3/indexer/test | 
*IndexerApi* | [**TestallIndexer**](sonarr/docs/IndexerApi.md#testallindexer) | **Post** /api/v3/indexer/testall | 
*IndexerApi* | [**UpdateIndexer**](sonarr/docs/IndexerApi.md#updateindexer) | **Put** /api/v3/indexer/{id} | 
*IndexerConfigApi* | [**GetIndexerConfig**](sonarr/docs/IndexerConfigApi.md#getindexerconfig) | **Get** /api/v3/config/indexer | 
*IndexerConfigApi* | [**GetIndexerConfigById**](sonarr/docs/IndexerConfigApi.md#getindexerconfigbyid) | **Get** /api/v3/config/indexer/{id} | 
*IndexerConfigApi* | [**UpdateIndexerConfig**](sonarr/docs/IndexerConfigApi.md#updateindexerconfig) | **Put** /api/v3/config/indexer/{id} | 
*LanguageApi* | [**GetLanguageById**](sonarr/docs/LanguageApi.md#getlanguagebyid) | **Get** /api/v3/language/{id} | 
*LanguageApi* | [**ListLanguage**](sonarr/docs/LanguageApi.md#listlanguage) | **Get** /api/v3/language | 
*LanguageProfileApi* | [**CreateLanguageProfile**](sonarr/docs/LanguageProfileApi.md#createlanguageprofile) | **Post** /api/v3/languageprofile | 
*LanguageProfileApi* | [**DeleteLanguageProfile**](sonarr/docs/LanguageProfileApi.md#deletelanguageprofile) | **Delete** /api/v3/languageprofile/{id} | 
*LanguageProfileApi* | [**GetLanguageProfileById**](sonarr/docs/LanguageProfileApi.md#getlanguageprofilebyid) | **Get** /api/v3/languageprofile/{id} | 
*LanguageProfileApi* | [**ListLanguageProfile**](sonarr/docs/LanguageProfileApi.md#listlanguageprofile) | **Get** /api/v3/languageprofile | 
*LanguageProfileApi* | [**UpdateLanguageProfile**](sonarr/docs/LanguageProfileApi.md#updatelanguageprofile) | **Put** /api/v3/languageprofile/{id} | 
*LanguageProfileSchemaApi* | [**GetLanguageprofileSchema**](sonarr/docs/LanguageProfileSchemaApi.md#getlanguageprofileschema) | **Get** /api/v3/languageprofile/schema | 
*LocalizationApi* | [**GetLocalization**](sonarr/docs/LocalizationApi.md#getlocalization) | **Get** /api/v3/localization | 
*LocalizationApi* | [**GetLocalizationById**](sonarr/docs/LocalizationApi.md#getlocalizationbyid) | **Get** /api/v3/localization/{id} | 
*LocalizationApi* | [**GetLocalizationLanguage**](sonarr/docs/LocalizationApi.md#getlocalizationlanguage) | **Get** /api/v3/localization/language | 
*LogApi* | [**GetLog**](sonarr/docs/LogApi.md#getlog) | **Get** /api/v3/log | 
*LogFileApi* | [**GetLogFileByFilename**](sonarr/docs/LogFileApi.md#getlogfilebyfilename) | **Get** /api/v3/log/file/{filename} | 
*LogFileApi* | [**ListLogFile**](sonarr/docs/LogFileApi.md#listlogfile) | **Get** /api/v3/log/file | 
*ManualImportApi* | [**CreateManualImport**](sonarr/docs/ManualImportApi.md#createmanualimport) | **Post** /api/v3/manualimport | 
*ManualImportApi* | [**ListManualImport**](sonarr/docs/ManualImportApi.md#listmanualimport) | **Get** /api/v3/manualimport | 
*MediaCoverApi* | [**GetMediaCoverseriesIdByFilename**](sonarr/docs/MediaCoverApi.md#getmediacoverseriesidbyfilename) | **Get** /api/v3/mediacover/{seriesId}/{filename} | 
*MediaManagementConfigApi* | [**GetMediaManagementConfig**](sonarr/docs/MediaManagementConfigApi.md#getmediamanagementconfig) | **Get** /api/v3/config/mediamanagement | 
*MediaManagementConfigApi* | [**GetMediaManagementConfigById**](sonarr/docs/MediaManagementConfigApi.md#getmediamanagementconfigbyid) | **Get** /api/v3/config/mediamanagement/{id} | 
*MediaManagementConfigApi* | [**UpdateMediaManagementConfig**](sonarr/docs/MediaManagementConfigApi.md#updatemediamanagementconfig) | **Put** /api/v3/config/mediamanagement/{id} | 
*MetadataApi* | [**CreateMetadata**](sonarr/docs/MetadataApi.md#createmetadata) | **Post** /api/v3/metadata | 
*MetadataApi* | [**CreateMetadataActionByName**](sonarr/docs/MetadataApi.md#createmetadataactionbyname) | **Post** /api/v3/metadata/action/{name} | 
*MetadataApi* | [**DeleteMetadata**](sonarr/docs/MetadataApi.md#deletemetadata) | **Delete** /api/v3/metadata/{id} | 
*MetadataApi* | [**GetMetadataById**](sonarr/docs/MetadataApi.md#getmetadatabyid) | **Get** /api/v3/metadata/{id} | 
*MetadataApi* | [**ListMetadata**](sonarr/docs/MetadataApi.md#listmetadata) | **Get** /api/v3/metadata | 
*MetadataApi* | [**ListMetadataSchema**](sonarr/docs/MetadataApi.md#listmetadataschema) | **Get** /api/v3/metadata/schema | 
*MetadataApi* | [**TestMetadata**](sonarr/docs/MetadataApi.md#testmetadata) | **Post** /api/v3/metadata/test | 
*MetadataApi* | [**TestallMetadata**](sonarr/docs/MetadataApi.md#testallmetadata) | **Post** /api/v3/metadata/testall | 
*MetadataApi* | [**UpdateMetadata**](sonarr/docs/MetadataApi.md#updatemetadata) | **Put** /api/v3/metadata/{id} | 
*MissingApi* | [**GetWantedMissing**](sonarr/docs/MissingApi.md#getwantedmissing) | **Get** /api/v3/wanted/missing | 
*MissingApi* | [**GetWantedMissingById**](sonarr/docs/MissingApi.md#getwantedmissingbyid) | **Get** /api/v3/wanted/missing/{id} | 
*NamingConfigApi* | [**GetNamingConfig**](sonarr/docs/NamingConfigApi.md#getnamingconfig) | **Get** /api/v3/config/naming | 
*NamingConfigApi* | [**GetNamingConfigById**](sonarr/docs/NamingConfigApi.md#getnamingconfigbyid) | **Get** /api/v3/config/naming/{id} | 
*NamingConfigApi* | [**GetNamingConfigExamples**](sonarr/docs/NamingConfigApi.md#getnamingconfigexamples) | **Get** /api/v3/config/naming/examples | 
*NamingConfigApi* | [**UpdateNamingConfig**](sonarr/docs/NamingConfigApi.md#updatenamingconfig) | **Put** /api/v3/config/naming/{id} | 
*NotificationApi* | [**CreateNotification**](sonarr/docs/NotificationApi.md#createnotification) | **Post** /api/v3/notification | 
*NotificationApi* | [**CreateNotificationActionByName**](sonarr/docs/NotificationApi.md#createnotificationactionbyname) | **Post** /api/v3/notification/action/{name} | 
*NotificationApi* | [**DeleteNotification**](sonarr/docs/NotificationApi.md#deletenotification) | **Delete** /api/v3/notification/{id} | 
*NotificationApi* | [**GetNotificationById**](sonarr/docs/NotificationApi.md#getnotificationbyid) | **Get** /api/v3/notification/{id} | 
*NotificationApi* | [**ListNotification**](sonarr/docs/NotificationApi.md#listnotification) | **Get** /api/v3/notification | 
*NotificationApi* | [**ListNotificationSchema**](sonarr/docs/NotificationApi.md#listnotificationschema) | **Get** /api/v3/notification/schema | 
*NotificationApi* | [**TestNotification**](sonarr/docs/NotificationApi.md#testnotification) | **Post** /api/v3/notification/test | 
*NotificationApi* | [**TestallNotification**](sonarr/docs/NotificationApi.md#testallnotification) | **Post** /api/v3/notification/testall | 
*NotificationApi* | [**UpdateNotification**](sonarr/docs/NotificationApi.md#updatenotification) | **Put** /api/v3/notification/{id} | 
*ParseApi* | [**GetParse**](sonarr/docs/ParseApi.md#getparse) | **Get** /api/v3/parse | 
*PingApi* | [**GetPing**](sonarr/docs/PingApi.md#getping) | **Get** /ping | 
*QualityDefinitionApi* | [**GetQualityDefinitionById**](sonarr/docs/QualityDefinitionApi.md#getqualitydefinitionbyid) | **Get** /api/v3/qualitydefinition/{id} | 
*QualityDefinitionApi* | [**ListQualityDefinition**](sonarr/docs/QualityDefinitionApi.md#listqualitydefinition) | **Get** /api/v3/qualitydefinition | 
*QualityDefinitionApi* | [**PutQualityDefinitionUpdate**](sonarr/docs/QualityDefinitionApi.md#putqualitydefinitionupdate) | **Put** /api/v3/qualitydefinition/update | 
*QualityDefinitionApi* | [**UpdateQualityDefinition**](sonarr/docs/QualityDefinitionApi.md#updatequalitydefinition) | **Put** /api/v3/qualitydefinition/{id} | 
*QualityProfileApi* | [**CreateQualityProfile**](sonarr/docs/QualityProfileApi.md#createqualityprofile) | **Post** /api/v3/qualityprofile | 
*QualityProfileApi* | [**DeleteQualityProfile**](sonarr/docs/QualityProfileApi.md#deletequalityprofile) | **Delete** /api/v3/qualityprofile/{id} | 
*QualityProfileApi* | [**GetQualityProfileById**](sonarr/docs/QualityProfileApi.md#getqualityprofilebyid) | **Get** /api/v3/qualityprofile/{id} | 
*QualityProfileApi* | [**ListQualityProfile**](sonarr/docs/QualityProfileApi.md#listqualityprofile) | **Get** /api/v3/qualityprofile | 
*QualityProfileApi* | [**UpdateQualityProfile**](sonarr/docs/QualityProfileApi.md#updatequalityprofile) | **Put** /api/v3/qualityprofile/{id} | 
*QualityProfileSchemaApi* | [**GetQualityprofileSchema**](sonarr/docs/QualityProfileSchemaApi.md#getqualityprofileschema) | **Get** /api/v3/qualityprofile/schema | 
*QueueApi* | [**DeleteQueue**](sonarr/docs/QueueApi.md#deletequeue) | **Delete** /api/v3/queue/{id} | 
*QueueApi* | [**DeleteQueueBulk**](sonarr/docs/QueueApi.md#deletequeuebulk) | **Delete** /api/v3/queue/bulk | 
*QueueApi* | [**GetQueue**](sonarr/docs/QueueApi.md#getqueue) | **Get** /api/v3/queue | 
*QueueActionApi* | [**CreateQueueGrabBulk**](sonarr/docs/QueueActionApi.md#createqueuegrabbulk) | **Post** /api/v3/queue/grab/bulk | 
*QueueActionApi* | [**CreateQueueGrabById**](sonarr/docs/QueueActionApi.md#createqueuegrabbyid) | **Post** /api/v3/queue/grab/{id} | 
*QueueDetailsApi* | [**ListQueueDetails**](sonarr/docs/QueueDetailsApi.md#listqueuedetails) | **Get** /api/v3/queue/details | 
*QueueStatusApi* | [**GetQueueStatus**](sonarr/docs/QueueStatusApi.md#getqueuestatus) | **Get** /api/v3/queue/status | 
*ReleaseApi* | [**CreateRelease**](sonarr/docs/ReleaseApi.md#createrelease) | **Post** /api/v3/release | 
*ReleaseApi* | [**ListRelease**](sonarr/docs/ReleaseApi.md#listrelease) | **Get** /api/v3/release | 
*ReleaseProfileApi* | [**CreateReleaseProfile**](sonarr/docs/ReleaseProfileApi.md#createreleaseprofile) | **Post** /api/v3/releaseprofile | 
*ReleaseProfileApi* | [**DeleteReleaseProfile**](sonarr/docs/ReleaseProfileApi.md#deletereleaseprofile) | **Delete** /api/v3/releaseprofile/{id} | 
*ReleaseProfileApi* | [**GetReleaseProfileById**](sonarr/docs/ReleaseProfileApi.md#getreleaseprofilebyid) | **Get** /api/v3/releaseprofile/{id} | 
*ReleaseProfileApi* | [**ListReleaseProfile**](sonarr/docs/ReleaseProfileApi.md#listreleaseprofile) | **Get** /api/v3/releaseprofile | 
*ReleaseProfileApi* | [**UpdateReleaseProfile**](sonarr/docs/ReleaseProfileApi.md#updatereleaseprofile) | **Put** /api/v3/releaseprofile/{id} | 
*ReleasePushApi* | [**CreateReleasePush**](sonarr/docs/ReleasePushApi.md#createreleasepush) | **Post** /api/v3/release/push | 
*RemotePathMappingApi* | [**CreateRemotePathMapping**](sonarr/docs/RemotePathMappingApi.md#createremotepathmapping) | **Post** /api/v3/remotepathmapping | 
*RemotePathMappingApi* | [**DeleteRemotePathMapping**](sonarr/docs/RemotePathMappingApi.md#deleteremotepathmapping) | **Delete** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingApi* | [**GetRemotePathMappingById**](sonarr/docs/RemotePathMappingApi.md#getremotepathmappingbyid) | **Get** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingApi* | [**ListRemotePathMapping**](sonarr/docs/RemotePathMappingApi.md#listremotepathmapping) | **Get** /api/v3/remotepathmapping | 
*RemotePathMappingApi* | [**UpdateRemotePathMapping**](sonarr/docs/RemotePathMappingApi.md#updateremotepathmapping) | **Put** /api/v3/remotepathmapping/{id} | 
*RenameEpisodeApi* | [**ListRename**](sonarr/docs/RenameEpisodeApi.md#listrename) | **Get** /api/v3/rename | 
*RootFolderApi* | [**CreateRootFolder**](sonarr/docs/RootFolderApi.md#createrootfolder) | **Post** /api/v3/rootfolder | 
*RootFolderApi* | [**DeleteRootFolder**](sonarr/docs/RootFolderApi.md#deleterootfolder) | **Delete** /api/v3/rootfolder/{id} | 
*RootFolderApi* | [**GetRootFolderById**](sonarr/docs/RootFolderApi.md#getrootfolderbyid) | **Get** /api/v3/rootfolder/{id} | 
*RootFolderApi* | [**ListRootFolder**](sonarr/docs/RootFolderApi.md#listrootfolder) | **Get** /api/v3/rootfolder | 
*SeasonPassApi* | [**CreateSeasonPass**](sonarr/docs/SeasonPassApi.md#createseasonpass) | **Post** /api/v3/seasonpass | 
*SeriesApi* | [**CreateSeries**](sonarr/docs/SeriesApi.md#createseries) | **Post** /api/v3/series | 
*SeriesApi* | [**DeleteSeries**](sonarr/docs/SeriesApi.md#deleteseries) | **Delete** /api/v3/series/{id} | 
*SeriesApi* | [**GetSeriesById**](sonarr/docs/SeriesApi.md#getseriesbyid) | **Get** /api/v3/series/{id} | 
*SeriesApi* | [**ListSeries**](sonarr/docs/SeriesApi.md#listseries) | **Get** /api/v3/series | 
*SeriesApi* | [**UpdateSeries**](sonarr/docs/SeriesApi.md#updateseries) | **Put** /api/v3/series/{id} | 
*SeriesEditorApi* | [**DeleteSeriesEditor**](sonarr/docs/SeriesEditorApi.md#deleteserieseditor) | **Delete** /api/v3/series/editor | 
*SeriesEditorApi* | [**PutSeriesEditor**](sonarr/docs/SeriesEditorApi.md#putserieseditor) | **Put** /api/v3/series/editor | 
*SeriesImportApi* | [**CreateSeriesImport**](sonarr/docs/SeriesImportApi.md#createseriesimport) | **Post** /api/v3/series/import | 
*SeriesLookupApi* | [**ListSeriesLookup**](sonarr/docs/SeriesLookupApi.md#listserieslookup) | **Get** /api/v3/series/lookup | 
*StaticResourceApi* | [**Get**](sonarr/docs/StaticResourceApi.md#get) | **Get** / | 
*StaticResourceApi* | [**GetByPath**](sonarr/docs/StaticResourceApi.md#getbypath) | **Get** /{path} | 
*StaticResourceApi* | [**GetContentByPath**](sonarr/docs/StaticResourceApi.md#getcontentbypath) | **Get** /content/{path} | 
*StaticResourceApi* | [**GetLogin**](sonarr/docs/StaticResourceApi.md#getlogin) | **Get** /login | 
*SystemApi* | [**CreateSystemRestart**](sonarr/docs/SystemApi.md#createsystemrestart) | **Post** /api/v3/system/restart | 
*SystemApi* | [**CreateSystemShutdown**](sonarr/docs/SystemApi.md#createsystemshutdown) | **Post** /api/v3/system/shutdown | 
*SystemApi* | [**GetSystemRoutes**](sonarr/docs/SystemApi.md#getsystemroutes) | **Get** /api/v3/system/routes | 
*SystemApi* | [**GetSystemRoutesDuplicate**](sonarr/docs/SystemApi.md#getsystemroutesduplicate) | **Get** /api/v3/system/routes/duplicate | 
*SystemApi* | [**GetSystemStatus**](sonarr/docs/SystemApi.md#getsystemstatus) | **Get** /api/v3/system/status | 
*TagApi* | [**CreateTag**](sonarr/docs/TagApi.md#createtag) | **Post** /api/v3/tag | 
*TagApi* | [**DeleteTag**](sonarr/docs/TagApi.md#deletetag) | **Delete** /api/v3/tag/{id} | 
*TagApi* | [**GetTagById**](sonarr/docs/TagApi.md#gettagbyid) | **Get** /api/v3/tag/{id} | 
*TagApi* | [**ListTag**](sonarr/docs/TagApi.md#listtag) | **Get** /api/v3/tag | 
*TagApi* | [**UpdateTag**](sonarr/docs/TagApi.md#updatetag) | **Put** /api/v3/tag/{id} | 
*TagDetailsApi* | [**GetTagDetailById**](sonarr/docs/TagDetailsApi.md#gettagdetailbyid) | **Get** /api/v3/tag/detail/{id} | 
*TagDetailsApi* | [**ListTagDetail**](sonarr/docs/TagDetailsApi.md#listtagdetail) | **Get** /api/v3/tag/detail | 
*TaskApi* | [**GetSystemTaskById**](sonarr/docs/TaskApi.md#getsystemtaskbyid) | **Get** /api/v3/system/task/{id} | 
*TaskApi* | [**ListSystemTask**](sonarr/docs/TaskApi.md#listsystemtask) | **Get** /api/v3/system/task | 
*UiConfigApi* | [**GetUiConfig**](sonarr/docs/UiConfigApi.md#getuiconfig) | **Get** /api/v3/config/ui | 
*UiConfigApi* | [**GetUiConfigById**](sonarr/docs/UiConfigApi.md#getuiconfigbyid) | **Get** /api/v3/config/ui/{id} | 
*UiConfigApi* | [**UpdateUiConfig**](sonarr/docs/UiConfigApi.md#updateuiconfig) | **Put** /api/v3/config/ui/{id} | 
*UpdateApi* | [**ListUpdate**](sonarr/docs/UpdateApi.md#listupdate) | **Get** /api/v3/update | 
*UpdateLogFileApi* | [**GetLogFileUpdateByFilename**](sonarr/docs/UpdateLogFileApi.md#getlogfileupdatebyfilename) | **Get** /api/v3/log/file/update/{filename} | 
*UpdateLogFileApi* | [**ListLogFileUpdate**](sonarr/docs/UpdateLogFileApi.md#listlogfileupdate) | **Get** /api/v3/log/file/update | 


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
 - [ImportListExclusionResource](docs/ImportListExclusionResource.md)
 - [ImportListResource](docs/ImportListResource.md)
 - [ImportListType](docs/ImportListType.md)
 - [IndexerBulkResource](docs/IndexerBulkResource.md)
 - [IndexerConfigResource](docs/IndexerConfigResource.md)
 - [IndexerResource](docs/IndexerResource.md)
 - [Language](docs/Language.md)
 - [LanguageProfileItemResource](docs/LanguageProfileItemResource.md)
 - [LanguageProfileResource](docs/LanguageProfileResource.md)
 - [LanguageResource](docs/LanguageResource.md)
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
 - [ReleaseEpisodeResource](docs/ReleaseEpisodeResource.md)
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
 - [TrackedDownloadState](docs/TrackedDownloadState.md)
 - [TrackedDownloadStatus](docs/TrackedDownloadStatus.md)
 - [TrackedDownloadStatusMessage](docs/TrackedDownloadStatusMessage.md)
 - [UiConfigResource](docs/UiConfigResource.md)
 - [UnmappedFolder](docs/UnmappedFolder.md)
 - [UpdateChanges](docs/UpdateChanges.md)
 - [UpdateMechanism](docs/UpdateMechanism.md)
 - [UpdateResource](docs/UpdateResource.md)


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



