# Changelog

## 2.1.0 (2026-01-16)

Full Changelog: [v2.0.0...v2.1.0](https://github.com/imagekit-developer/imagekit-go/compare/v2.0.0...v2.1.0)

### Features

* **api:** add GetImageAttributesOptions and ResponsiveImageAttributes schemas; update resource references in main.yaml; remove dummy endpoint ([41072da](https://github.com/imagekit-developer/imagekit-go/commit/41072da63cd2ba891a911d932af3bc8b70c90588))
* **api:** Add saved extensions API and enhance transformation options ([1f7e772](https://github.com/imagekit-developer/imagekit-go/commit/1f7e7723cb74828044a489d0b2de58b195b05183))
* **api:** fix go sdk breaking changes ([6cbddff](https://github.com/imagekit-developer/imagekit-go/commit/6cbddffab95c89b964fc29ce119ceb70d7ebded5))
* **encoder:** support bracket encoding form-data object members ([cb3e557](https://github.com/imagekit-developer/imagekit-go/commit/cb3e5572b00fe978ee93d595cc4d8775edccbc89))
* **tests:** add transformations for radius, color replace, and distort; enhance overlay tests with layer modes ([382b9ed](https://github.com/imagekit-developer/imagekit-go/commit/382b9ed5e05f5fb1810481c0fd5a2299bbf5cfde))


### Bug Fixes

* add ai-tasks property to response schemas with enum values ([0551693](https://github.com/imagekit-developer/imagekit-go/commit/0551693b9a7eeb35984c198046e80a6a9debdbf8))
* **client:** correctly specify Accept header with */* instead of empty ([21a30a4](https://github.com/imagekit-developer/imagekit-go/commit/21a30a4d60b4b2da84d2314ddf6bcf76759da64d))
* **client:** invalid URL ([aea74b3](https://github.com/imagekit-developer/imagekit-go/commit/aea74b387869ce1fa566bf7a1fda3b2432eff9a9))
* **client:** properly marshal embedded structs ([e55e614](https://github.com/imagekit-developer/imagekit-go/commit/e55e614fab629dfab68d7249e53ef602dd1a36b3))
* **docs:** update go get command to include version path in README.md ([d7d4c82](https://github.com/imagekit-developer/imagekit-go/commit/d7d4c829ebccafd1242d79a03651f1189c9f24d0))
* **mcp:** correct code tool API endpoint ([b32395e](https://github.com/imagekit-developer/imagekit-go/commit/b32395e36a3fdd2f8e37313a303a68135f13400f))
* rename param to avoid collision ([5067fd4](https://github.com/imagekit-developer/imagekit-go/commit/5067fd4adfe3a7108f3799f270c4211ade385882))
* skip usage tests that don't work with Prism ([429ad75](https://github.com/imagekit-developer/imagekit-go/commit/429ad75eb8c267b44a9c8e4f2c542344189563e3))


### Chores

* add float64 to valid types for RegisterFieldValidator ([2dc3cae](https://github.com/imagekit-developer/imagekit-go/commit/2dc3cae63386dc6b8c7743af3fdb0a9c4ed93ef5))
* bump gjson version ([87ad44d](https://github.com/imagekit-developer/imagekit-go/commit/87ad44d7016dcb641158e26cc4f0d08e89770dc5))
* elide duplicate aliases ([2f9eee1](https://github.com/imagekit-developer/imagekit-go/commit/2f9eee11c82bd0b6641dec1f0c20042863ad169f))
* **internal:** codegen related update ([23edba8](https://github.com/imagekit-developer/imagekit-go/commit/23edba82802b3bc322649f9693ff71134db70b5c))
* **internal:** codegen related update ([2fdd961](https://github.com/imagekit-developer/imagekit-go/commit/2fdd961c6a458140239af2e88cf1b1ecf7dc27f2))
* **internal:** codegen related update ([8877b4f](https://github.com/imagekit-developer/imagekit-go/commit/8877b4fbdce39a56785613a1172dda44399c6fe7))
* **internal:** codegen related update ([d83769d](https://github.com/imagekit-developer/imagekit-go/commit/d83769df0486737f479694acc9421415fb11c523))
* **internal:** codegen related update ([63165ac](https://github.com/imagekit-developer/imagekit-go/commit/63165ac51ec10df842ed8f9496c82294b4f9e61e))
* **internal:** grammar fix (it's -&gt; its) ([e35e192](https://github.com/imagekit-developer/imagekit-go/commit/e35e1922f7ad2541ed116db557bcedd8c9c088de))
* remove MCP Server section from README.md ([58749a4](https://github.com/imagekit-developer/imagekit-go/commit/58749a446146fc365b10ced5db1556eb5dbd1de3))


### Documentation

* prominently feature MCP server setup in root SDK readmes ([e2a2d90](https://github.com/imagekit-developer/imagekit-go/commit/e2a2d9047c24a8fd32471aaf96315cad0bd9f935))

## 2.0.0 (2025-10-05)

Full Changelog: [v0.0.1...v2.0.0](https://github.com/imagekit-developer/imagekit-go/compare/v0.0.1...v2.0.0)

### Features

* **api:** add BaseWebhookEvent ([49aee77](https://github.com/imagekit-developer/imagekit-go/commit/49aee77cc145d6f45eb44dd5e6fc080c7fe7a001))
* **api:** add dummy endpoint for shared model generation in go ([c5b8b5f](https://github.com/imagekit-developer/imagekit-go/commit/c5b8b5f8abce0058c19693590dc00b575da039d3))
* **api:** add path policy related non-breaking changes ([385c57b](https://github.com/imagekit-developer/imagekit-go/commit/385c57b53db52f1ac00ad75f99f9015541aaf45c))
* **api:** add selectedFieldsSchema in upload and list API response ([7e3af2b](https://github.com/imagekit-developer/imagekit-go/commit/7e3af2b3672a75cac175489564579df8ed6fc98a))
* **api:** extract UpdateFileDetailsRequest to model ([8b5aea9](https://github.com/imagekit-developer/imagekit-go/commit/8b5aea9dde90cbe13cba1d451fb01ad35dfd2e13))
* **api:** fix upload API request params ([db0d74c](https://github.com/imagekit-developer/imagekit-go/commit/db0d74c7e7be7f041110b967ad2aac6d0b90684e))
* **api:** manual updates ([acff3e5](https://github.com/imagekit-developer/imagekit-go/commit/acff3e528b86aa7423ae204e2910556f0246e554))
* **api:** manual updates ([98f9de6](https://github.com/imagekit-developer/imagekit-go/commit/98f9de6e111944c0d81a129ad511e79d34c974a5))
* **api:** manual updates ([49c8193](https://github.com/imagekit-developer/imagekit-go/commit/49c8193c3192fb49491cbc5ae5e2b4acb212c21e))
* **api:** manual updates ([a6a9d9c](https://github.com/imagekit-developer/imagekit-go/commit/a6a9d9cc5d0f3d40aa536aa14ea3a3c3b181a6d8))
* **api:** manual updates ([67b5049](https://github.com/imagekit-developer/imagekit-go/commit/67b5049b88e5feb8815ab087d3efffc9173596fc))
* **api:** manual updates ([5829397](https://github.com/imagekit-developer/imagekit-go/commit/58293972018671d07621207fabde810e368be16e))
* **api:** manual updates ([cc22fef](https://github.com/imagekit-developer/imagekit-go/commit/cc22fef638f68ff664019f5ba24f4292c7866771))
* **api:** manual updates ([051d622](https://github.com/imagekit-developer/imagekit-go/commit/051d6221e188479c80197daaf448ce70c7d1248a))
* **api:** manual updates ([452fcb4](https://github.com/imagekit-developer/imagekit-go/commit/452fcb4070c0141322f5b49180986a10ee07be22))
* **api:** manual updates ([66bd0b0](https://github.com/imagekit-developer/imagekit-go/commit/66bd0b077f7ef37b59294dd0761648f35c028747))
* **api:** manual updates ([9a4ab31](https://github.com/imagekit-developer/imagekit-go/commit/9a4ab3117285fd91a09f4b0ef3b868ca5d46ae63))
* **api:** manual updates ([67adab7](https://github.com/imagekit-developer/imagekit-go/commit/67adab7a0d64f06f0fc9beed70c67cad3b5f1706))
* **api:** manual updates ([5237fb5](https://github.com/imagekit-developer/imagekit-go/commit/5237fb51c965361f7d96d3dada6726d54cb7349c))
* **api:** manual updates ([b7169f4](https://github.com/imagekit-developer/imagekit-go/commit/b7169f426e0c8bcb6cebbd9413a5674cdaf79b5b))
* **api:** manual updates ([bc53c1a](https://github.com/imagekit-developer/imagekit-go/commit/bc53c1a5493d2fbe02f90523d3578e3cd19b8260))
* **api:** manual updates ([e197563](https://github.com/imagekit-developer/imagekit-go/commit/e1975634e1a07e8276fe79314f91b84d06a4399d))
* **api:** modify multipart marshaling for file upload parameters as per our backend expects ([aba9cfa](https://github.com/imagekit-developer/imagekit-go/commit/aba9cfaec6a5f0c97c60da752f821433cbd718c4))
* **api:** remove Stainless attribution from readme ([e41a4bd](https://github.com/imagekit-developer/imagekit-go/commit/e41a4bd5601df12655cff51873049c0280e439ed))
* **api:** update api docs link ([9c185d4](https://github.com/imagekit-developer/imagekit-go/commit/9c185d44f474acc5d5da001d0da62378bfd276fc))
* **api:** Update env var name ([d7aace7](https://github.com/imagekit-developer/imagekit-go/commit/d7aace7fc7604c66ad73d701595b440ae1578c75))
* **api:** updated docs ([2f0ac45](https://github.com/imagekit-developer/imagekit-go/commit/2f0ac45bea1908568a08cd1b44f76eb46e6a4dbc))
* **examples:** add file upload examples for local, base64, URL, and advanced uploads ([799c172](https://github.com/imagekit-developer/imagekit-go/commit/799c172fb612ea829609e2471f7ec61439a354b3))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([07ab70c](https://github.com/imagekit-developer/imagekit-go/commit/07ab70ce76d2472aaf15c021719e2fed71b5cd32))
* **internal:** unmarshal correctly when there are multiple discriminators ([b87188b](https://github.com/imagekit-developer/imagekit-go/commit/b87188b901ba7422570ed594ee0185fc6e4545b3))
* **README:** update example webhook secrets and improve event logging for clarity ([cbf2a51](https://github.com/imagekit-developer/imagekit-go/commit/cbf2a51cc744d3222c4d029ffcbfbe6c41d9c845))
* **tests:** update gradient transformation test case with new parameters ([02ab219](https://github.com/imagekit-developer/imagekit-go/commit/02ab21975a27393904724f120d6f837a2dc094c0))
* update import paths to use the correct imagekit-developer repository ([09f4ca0](https://github.com/imagekit-developer/imagekit-go/commit/09f4ca00e565bd2ac6f47ac52fcc66aa14610dd9))
* use slices.Concat instead of sometimes modifying r.Options ([7822f5a](https://github.com/imagekit-developer/imagekit-go/commit/7822f5a23a73feacec29a3db1011bc01c3318e10))
* **webhook:** encode webhook secret in base64 before verification ([6f18756](https://github.com/imagekit-developer/imagekit-go/commit/6f187566f76695ae21e56b8f52ddc963be519058))


### Chores

* bump minimum go version to 1.22 ([2cd7a9b](https://github.com/imagekit-developer/imagekit-go/commit/2cd7a9bd62b9cc830bf59d3cce7b11f2b05528e2))
* do not install brew dependencies in ./scripts/bootstrap by default ([ac2daaa](https://github.com/imagekit-developer/imagekit-go/commit/ac2daaa1b8bfa78b1bc814e98a5c220c402f973d))
* **internal:** codegen related update ([a9dae33](https://github.com/imagekit-developer/imagekit-go/commit/a9dae337065c0b204dcf4a3a65da0d9abbe1b376))
* sync repo ([9c87c18](https://github.com/imagekit-developer/imagekit-go/commit/9c87c185f17e8c881e666149f590559512df501f))
* update more docs for 1.22 ([6cfe89f](https://github.com/imagekit-developer/imagekit-go/commit/6cfe89f7c0d5968ca9d4bc10e54f5debf83f2c1b))
* update SDK settings ([369f85d](https://github.com/imagekit-developer/imagekit-go/commit/369f85da041fd5a4b6b379c994e3ac1d9ebed36f))
* update SDK settings ([2d43aee](https://github.com/imagekit-developer/imagekit-go/commit/2d43aee37d06353a55a9894a52ed5b4b8300d91b))


### Documentation

* correct typo in default value description for custom metadata field ([64ff337](https://github.com/imagekit-developer/imagekit-go/commit/64ff3372bf902c9b7db2e2273abd15fa0010d1ab))
* **README:** add section on using Raw transformations for undocumented features ([09d84f7](https://github.com/imagekit-developer/imagekit-go/commit/09d84f71136a819c3638e4e894f0f009890c5ca0))
* **README:** update file upload examples to use local files and improve error handling ([c20f7b7](https://github.com/imagekit-developer/imagekit-go/commit/c20f7b731608bb3e0b60d508da1a365cca11861b))
* **README:** update SDK description and enhance table of contents with detailed sections ([acc22b8](https://github.com/imagekit-developer/imagekit-go/commit/acc22b80ba4c10a13bb576417c8099b931349fbc))


### Refactors

* **helper:** remove legacy Helper struct and streamline methods in HelperService ([534383e](https://github.com/imagekit-developer/imagekit-go/commit/534383e3c9f70514ff66fc8f2932293b2ba49a9d))
* **helper:** replace UUID token generation with random byte generation and update tests for hex format ([4e30f5a](https://github.com/imagekit-developer/imagekit-go/commit/4e30f5a42441dc3436c2058a9db68d10de0f08bb))
* **helper:** streamline transformation string generation with mapping functions ([36ccee1](https://github.com/imagekit-developer/imagekit-go/commit/36ccee1628edcb86929c8446c03060a310d840d6))
* **helper:** update HelperService to return errors instead of panicking and adjust related tests ([0f5bb17](https://github.com/imagekit-developer/imagekit-go/commit/0f5bb17456294880973720f38a5cc717fdc45f95))
* **tests:** improve test case descriptions in helper_overlay_test.go ([eb8c7d9](https://github.com/imagekit-developer/imagekit-go/commit/eb8c7d998be411141493fe9f79ba91df36e1f673))
* **tests:** reorder import statements and enhance test descriptions in helper tests ([eff5fb6](https://github.com/imagekit-developer/imagekit-go/commit/eff5fb60832e50dde37b982e4ff5c12b74e40dad))
* **upload:** update file upload parameter documentation to clarify required io.Reader type ([c5aabc6](https://github.com/imagekit-developer/imagekit-go/commit/c5aabc6de66ab4e43a37e939bdc279421ed0a4fb))
