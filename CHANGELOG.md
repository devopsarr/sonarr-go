# Changelog

## [1.0.1](https://github.com/devopsarr/sonarr-go/compare/v1.0.0...v1.0.1) (2024-10-08)


### Bug Fixes

* **deps:** update openapitools/openapi-generator-cli docker tag to v7.9.0 ([a5deb71](https://github.com/devopsarr/sonarr-go/commit/a5deb71993a11284c9acd12dc9ed43931bf70cfc))

## [1.0.0](https://github.com/devopsarr/sonarr-go/compare/v0.6.0...v1.0.0) (2024-02-15)


### ⚠ BREAKING CHANGES

* **devopsarr/prowlarr-py#39:** update sdk generator to use pydantic 2.x

### Features

* add sonarr auto tagging schema return type ([9fcbdf3](https://github.com/devopsarr/sonarr-go/commit/9fcbdf3c2de6eb696324560f72cb63c02135a46d))
* align to go sdk ([43f9b85](https://github.com/devopsarr/sonarr-go/commit/43f9b85acee5718d5800c86b6fe3e921d42672b1))
* **devopsarr/prowlarr-py#39:** update sdk generator to use pydantic 2.x ([cdc62a1](https://github.com/devopsarr/sonarr-go/commit/cdc62a12319882e13e65cdcc6a2b7ce823454de4))
* inject api version into readme ([c76458e](https://github.com/devopsarr/sonarr-go/commit/c76458ebf173d104a5687ecc6f65fae4046457db))


### Bug Fixes

* release please commented lines ([859267d](https://github.com/devopsarr/sonarr-go/commit/859267db39fcb77533a88df38d43d8e2e32a4327))
* remove middle elements from method name ([084f14d](https://github.com/devopsarr/sonarr-go/commit/084f14d21e4f9903209387e0e733b9c38c129a9b))

## [0.6.0](https://github.com/devopsarr/sonarr-go/compare/v0.5.1...v0.6.0) (2023-07-21)


### Features

* **sonarr:** add response type to custom format ([8864f7a](https://github.com/devopsarr/sonarr-go/commit/8864f7a090ba2716666599b32314aa5f6313a774))


### Bug Fixes

* set pydantic version ([85a91e0](https://github.com/devopsarr/sonarr-go/commit/85a91e0416f16e8dcf3f7942450f9e1f6d9f68de))

## [0.5.1](https://github.com/devopsarr/sonarr-go/compare/v0.5.0...v0.5.1) (2023-06-21)


### Bug Fixes

* sonarr release profile required and ignored list ([391a45b](https://github.com/devopsarr/sonarr-go/commit/391a45bd3787971a7624224b1f922cd8ce8b2f03))

## [0.5.0](https://github.com/devopsarr/sonarr-go/compare/v0.4.1...v0.5.0) (2023-05-15)


### Features

* pin openapi version and add version management ([087c1d5](https://github.com/devopsarr/sonarr-go/commit/087c1d5c0742d94a18d7302faf22e8427d023dcf))

## [0.4.1](https://github.com/devopsarr/sonarr-go/compare/v0.4.0...v0.4.1) (2023-03-27)


### Bug Fixes

* **devopsarr/sonarr-py#6:** remove timespan validator ([c1a2fb6](https://github.com/devopsarr/sonarr-go/commit/c1a2fb60d92b43b1555124c543792430058044f2))

## [0.4.0](https://github.com/devopsarr/sonarr-go/compare/v0.3.0...v0.4.0) (2023-03-23)


### Features

* sdk generation alignment ([989ad66](https://github.com/devopsarr/sonarr-go/commit/989ad665dcd6311b1e93dcd29e2ed23c7737b99a))

## [0.3.0](https://github.com/devopsarr/sonarr-go/compare/v0.2.2...v0.3.0) (2023-03-23)


### Features

* sdk generation alignment ([cf10f9d](https://github.com/devopsarr/sonarr-go/commit/cf10f9dccda61fb7ec46a0623b0e4ab625fdbe49))

## [0.2.2](https://github.com/devopsarr/sonarr-go/compare/v0.2.1...v0.2.2) (2023-02-27)


### Bug Fixes

* use time span type as string ([76571f8](https://github.com/devopsarr/sonarr-go/commit/76571f82ab462137ed903b777c0036572a76f6db))

## [0.2.1](https://github.com/devopsarr/sonarr-go/compare/v0.2.0...v0.2.1) (2023-02-27)


### Bug Fixes

* remove wrongly used time span ([85bf196](https://github.com/devopsarr/sonarr-go/commit/85bf196f9506dadf9fa1e2b6eaa915c0815f54b2))

## [0.2.0](https://github.com/devopsarr/sonarr-go/compare/v0.1.0...v0.2.0) (2023-02-21)


### Features

* add autotagging ([08c0c9a](https://github.com/devopsarr/sonarr-go/commit/08c0c9ae1235e7716304e5838905b7ddbce2a333))

## 0.1.0 (2023-01-24)


### Bug Fixes

* add missing language profile id to ImportListModel ([dce6273](https://github.com/devopsarr/sonarr-go/commit/dce627372a10c28e3908ed57a351b151883bc69b))
* move array to pointers ([5f5d01f](https://github.com/devopsarr/sonarr-go/commit/5f5d01f17750b9453a67f5552df4dad7dd59f570))
* package name ([7f7ca7b](https://github.com/devopsarr/sonarr-go/commit/7f7ca7b964d45b51801dad0585cf73a387545d27))
* remove language profile ID for v4 ([718768c](https://github.com/devopsarr/sonarr-go/commit/718768cb8b5298a0c54e607557b52e8ea54b76a5))
* rename methods CamelCase ([1a55b9b](https://github.com/devopsarr/sonarr-go/commit/1a55b9b11b0935a99cb8d4fd87fa0876d98003fb))
* SeriesLookup return type ([45a1301](https://github.com/devopsarr/sonarr-go/commit/45a1301d19e56e65ec6113ff4ffd7cbf75d9b5bf))
* version to string ([87f52f6](https://github.com/devopsarr/sonarr-go/commit/87f52f67793d95d687acb2a7d556e9952d0bfbc4))


### Miscellaneous Chores

* release 0.1.0 ([aa0814d](https://github.com/devopsarr/sonarr-go/commit/aa0814ddc60ca50226845de7c054dfce7526f2c8))
