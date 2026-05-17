# Changelog

## 2.6.0 (2026-05-17)

Full Changelog: [v2.5.0...v2.6.0](https://github.com/imagekit-developer/imagekit-go/compare/v2.5.0...v2.6.0)

### Features

* **client:** optimize json encoder for internal types ([33c708e](https://github.com/imagekit-developer/imagekit-go/commit/33c708ed32562ccf0e16a736e77573d4c1b9b204))

## 2.5.0 (2026-05-13)

Full Changelog: [v2.4.0...v2.5.0](https://github.com/imagekit-developer/imagekit-go/compare/v2.4.0...v2.5.0)

### Features

* **api:** add no-enlarge crop modes and colorize transformation ([ad25990](https://github.com/imagekit-developer/imagekit-go/commit/ad259909e99498bfda6beada5eaff2e87ffb2e54))
* **api:** manual updates ([791dd8c](https://github.com/imagekit-developer/imagekit-go/commit/791dd8c9fa7b8e9ee247918f6aceffc77ad7f8cf))
* **client:** add compatibility aliases for old type names ([b238a42](https://github.com/imagekit-developer/imagekit-go/commit/b238a423928c3f00584a4894bc1759e21bafb33d))
* **go:** add default http client with timeout ([f4d5e6b](https://github.com/imagekit-developer/imagekit-go/commit/f4d5e6bc9096d5e51368b2aa8810c22352900643))
* **helper:** add colorize transformation support and update tests ([53b441a](https://github.com/imagekit-developer/imagekit-go/commit/53b441a512d73384714df827ad202fb678d73a5e))
* support setting headers via env ([3872a5e](https://github.com/imagekit-developer/imagekit-go/commit/3872a5eb58f1c091b8be08183dd4c972e8a68f75))
* **tests:** comment update ([4bef04e](https://github.com/imagekit-developer/imagekit-go/commit/4bef04e89c89449faff1e6bf642ccc7158070790))


### Bug Fixes

* **go:** avoid panic when http.DefaultTransport is wrapped ([c762867](https://github.com/imagekit-developer/imagekit-go/commit/c7628672829b9037cdae1851a93bb8f1304a9c86))


### Chores

* avoid embedding reflect.Type for dead code elimination ([a8ce4c3](https://github.com/imagekit-developer/imagekit-go/commit/a8ce4c343ff55b3f698e96438167f340cddfc4c2))
* configure new SDK language ([0951290](https://github.com/imagekit-developer/imagekit-go/commit/0951290e27f1568894dd90e717fc49612e24bedd))
* **internal:** codegen related update ([1622f80](https://github.com/imagekit-developer/imagekit-go/commit/1622f80e1bae58177eae2b260e269e5b1b32de9a))
* **internal:** codegen related update ([56516d9](https://github.com/imagekit-developer/imagekit-go/commit/56516d9f70279d35c597800d8c580ed18650e6d7))
* **internal:** codegen related update ([14d3582](https://github.com/imagekit-developer/imagekit-go/commit/14d3582a90625f39b54ffb31000db0e35a849027))
* **internal:** codegen related update ([904d9f3](https://github.com/imagekit-developer/imagekit-go/commit/904d9f3ebf4291f4142dcd99db48b832761fba7f))
* **internal:** more robust bootstrap script ([5888120](https://github.com/imagekit-developer/imagekit-go/commit/5888120dc9032b1826cc0b7018f53f34db89666e))
* **internal:** simplify release-please config ([48c9497](https://github.com/imagekit-developer/imagekit-go/commit/48c949735a94871559549cd63d9dd9bad83f4241))
* redact api-key headers in debug logs ([add677a](https://github.com/imagekit-developer/imagekit-go/commit/add677afd39e072a7f0ef313f2a0f7d6ceff1f19))


### Documentation

* remove example code for verifying webhook signatures ([799e47c](https://github.com/imagekit-developer/imagekit-go/commit/799e47c306c5f681891804c6e3ca738b0b2c053c))

## 2.4.0 (2026-04-13)

Full Changelog: [v2.3.0...v2.4.0](https://github.com/imagekit-developer/imagekit-go/compare/v2.3.0...v2.4.0)

### Features

* **api:** dam related webhook events ([b364b42](https://github.com/imagekit-developer/imagekit-go/commit/b364b4250885a0c3d172d26e0121979970c55990))
* **api:** fix spec indentation ([f254242](https://github.com/imagekit-developer/imagekit-go/commit/f2542428409b7963a25d629c667291292d5b0126))
* **api:** indentation fix ([d9bbacf](https://github.com/imagekit-developer/imagekit-go/commit/d9bbacf78dd49e892ea97ead76000115c863495f))
* **api:** merge with main to bring back missing parameters ([4ab1555](https://github.com/imagekit-developer/imagekit-go/commit/4ab1555f5ecf9f5aeca0de356f67b7e7ce831a7c))
* **api:** update webhook event names and remove DAM prefix ([5054493](https://github.com/imagekit-developer/imagekit-go/commit/5054493498d777340e64c267da0621cc1fa3f804))


### Bug Fixes

* **api:** extract shared schemas to prevent Go webhook union breaking changes ([57ff5e2](https://github.com/imagekit-developer/imagekit-go/commit/57ff5e25ba5dfb11ebed03cdfb145001bf717d40))
* **api:** fix references of schema ([991d381](https://github.com/imagekit-developer/imagekit-go/commit/991d3815942efcb68655192ad7b81ab3fb37b791))
* **api:** remove optional password parameter from client initialization ([e322edc](https://github.com/imagekit-developer/imagekit-go/commit/e322edc056ae3145fcc1e10af75b6a1208692cd0))
* **api:** rename DamFile events to File for consistency ([6d142f5](https://github.com/imagekit-developer/imagekit-go/commit/6d142f5d4d48d2368a7b59d6508062dd63a46c72))


### Documentation

* improve examples ([3953c51](https://github.com/imagekit-developer/imagekit-go/commit/3953c51e4f8582b9ac3e317f846e8fcc17725c05))


### Refactors

* add backwards-compatible type aliases and import removed json package ([e12a0f6](https://github.com/imagekit-developer/imagekit-go/commit/e12a0f663b5fb2f89306aa02f2da40217a18a2b3))
* AITags to singular AITag schema with array items pattern ([1fe0465](https://github.com/imagekit-developer/imagekit-go/commit/1fe04652d68fccf6ce2465a8148b4aa892f9e3e0))

## 2.3.0 (2026-04-06)

Full Changelog: [v2.2.0...v2.3.0](https://github.com/imagekit-developer/imagekit-go/compare/v2.2.0...v2.3.0)

### Features

* **api:** add support for xCenter, yCenter, and anchorPoint in overlay URL generation ([50ef82e](https://github.com/imagekit-developer/imagekit-go/commit/50ef82eb5eb19302a7c9933166e2643c4b60b6ec))
* **api:** dpr type update ([7b5ed76](https://github.com/imagekit-developer/imagekit-go/commit/7b5ed7649a30f369e05741ad26e87866ab369341))
* **api:** Introduce lxc, lyc, lap parameters in overlays. ([1d33db8](https://github.com/imagekit-developer/imagekit-go/commit/1d33db8dbb8860ca132781e558ce3dab880860b2))
* **api:** revert dpr breaking change ([4f76b02](https://github.com/imagekit-developer/imagekit-go/commit/4f76b02da01cce562522509abc4aeb9f43632227))
* **internal:** support comma format in multipart form encoding ([e0ecf81](https://github.com/imagekit-developer/imagekit-go/commit/e0ecf8124b1c7a1f02cef55883fb139e0bbf2363))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([c772a06](https://github.com/imagekit-developer/imagekit-go/commit/c772a06aa381125614083d572ce99b73b4eb3ce5))
* **encoder:** correctly serialize NullStruct ([c8a223a](https://github.com/imagekit-developer/imagekit-go/commit/c8a223a358428b56cd5f13843928252c58acff90))
* fix issue with unmarshaling in some cases ([c05145e](https://github.com/imagekit-developer/imagekit-go/commit/c05145eb2eced30cb2398810796c0b6a2fa81e00))
* fix request delays for retrying to be more respectful of high requested delays ([065df63](https://github.com/imagekit-developer/imagekit-go/commit/065df633342605e2d5f69d707170c8b92e5f59fb))
* prevent duplicate ? in query params ([1c53d0d](https://github.com/imagekit-developer/imagekit-go/commit/1c53d0da2b0a2eaa26b29b3fa6b69dc84d78a8be))
* **types:** generate shared enum types that are not referenced by other schemas ([420472d](https://github.com/imagekit-developer/imagekit-go/commit/420472d6e58c1626a7a15fa0c81cfbe02539c2d0))


### Chores

* **ci:** add build step ([950faa7](https://github.com/imagekit-developer/imagekit-go/commit/950faa70ba325a3af7e2149618d1690d681888a1))
* **ci:** skip lint on metadata-only changes ([cb58aa5](https://github.com/imagekit-developer/imagekit-go/commit/cb58aa516481e088ef0924c1d2e8979df33800aa))
* **ci:** skip uploading artifacts on stainless-internal branches ([1fedd9f](https://github.com/imagekit-developer/imagekit-go/commit/1fedd9f85df8d3c231fb7eb7e342ef98a2fe3acc))
* **ci:** support opting out of skipping builds on metadata-only commits ([9781c64](https://github.com/imagekit-developer/imagekit-go/commit/9781c64025356760a3b067d0f13c897282747997))
* **client:** fix multipart serialisation of Default() fields ([facb97a](https://github.com/imagekit-developer/imagekit-go/commit/facb97a34c051ca7f0b5ce1c18fca7eddf94b68e))
* **internal:** codegen related update ([a4bb374](https://github.com/imagekit-developer/imagekit-go/commit/a4bb374e09dc89d66cc143dc3b6bb1f315e0e4cb))
* **internal:** codegen related update ([255a031](https://github.com/imagekit-developer/imagekit-go/commit/255a0313de39110cdedc4002d451bdc39e3bc308))
* **internal:** codegen related update ([e1e4755](https://github.com/imagekit-developer/imagekit-go/commit/e1e47559733779920f899cb665b244e431702668))
* **internal:** codegen related update ([5071176](https://github.com/imagekit-developer/imagekit-go/commit/50711763c0f3bee01b36ab00500db61aa371415b))
* **internal:** minor cleanup ([bba2f4a](https://github.com/imagekit-developer/imagekit-go/commit/bba2f4a012b76ed18936d5659db28505d88cf694))
* **internal:** move custom custom `json` tags to `api` ([4609e1f](https://github.com/imagekit-developer/imagekit-go/commit/4609e1fd02735cc92d442ab6c6360429b5a8318d))
* **internal:** remove mock server code ([ab007f7](https://github.com/imagekit-developer/imagekit-go/commit/ab007f7fa428ff639b7c1177976d76b4c66eb5f5))
* **internal:** support default value struct tag ([86a3f7c](https://github.com/imagekit-developer/imagekit-go/commit/86a3f7cc8bce7c25b59720b35dfe590b974933eb))
* **internal:** tweak CI branches ([28a816d](https://github.com/imagekit-developer/imagekit-go/commit/28a816d34c0bb30be0a550ade3fc7e3a848c98bf))
* **internal:** update gitignore ([bbc7450](https://github.com/imagekit-developer/imagekit-go/commit/bbc745035b8765ad93a60ab91114135d1343e1d0))
* **internal:** use explicit returns ([3d59762](https://github.com/imagekit-developer/imagekit-go/commit/3d59762abc0b08d20e1d3aa2f10e1e46b8e605b0))
* **internal:** use explicit returns in more places ([93dfb11](https://github.com/imagekit-developer/imagekit-go/commit/93dfb117b673e0fbb123d47862bd48f7e685bfdd))
* remove unnecessary error check for url parsing ([7517de5](https://github.com/imagekit-developer/imagekit-go/commit/7517de5c74a9c6cc7bab2ed0fa6c48278558d1b5))
* **tests:** update webhook tests ([f3ef31e](https://github.com/imagekit-developer/imagekit-go/commit/f3ef31e396305fb69fa7905821fd2e94258f3b18))
* update docs for api:"required" ([d5de86f](https://github.com/imagekit-developer/imagekit-go/commit/d5de86f78052cefc4c28da6049f66980d7f79112))
* update mock server docs ([1ba9e87](https://github.com/imagekit-developer/imagekit-go/commit/1ba9e87f6b30e3992999808434dd39298de4d0a0))
* update placeholder string ([a783267](https://github.com/imagekit-developer/imagekit-go/commit/a7832673f2d1d19e5c1d87b2ab3ad3b152288f73))

## 2.2.0 (2026-02-02)

Full Changelog: [v2.1.1...v2.2.0](https://github.com/imagekit-developer/imagekit-go/compare/v2.1.1...v2.2.0)

### Features

* **api:** add customMetadata property to folder schema ([71fcd55](https://github.com/imagekit-developer/imagekit-go/commit/71fcd554cadcbc4981b57078a2ab9c2fbe72be51))
* **client:** add a convenient param.SetJSON helper ([7db0b48](https://github.com/imagekit-developer/imagekit-go/commit/7db0b4829c864151bda74f002c122f8a018118a0))


### Bug Fixes

* **api:** add missing embeddedMetadata and video properties to FileDetails ([b0e6909](https://github.com/imagekit-developer/imagekit-go/commit/b0e69096abac1ceba154150495cc2ae508ba63a0))
* **docs:** fix mcp installation instructions for remote servers ([f3875dd](https://github.com/imagekit-developer/imagekit-go/commit/f3875dddc1017d0db1ac2f3a86bd905bb941353d))
* **tests:** update subtitle references to use plural form ([7d180b7](https://github.com/imagekit-developer/imagekit-go/commit/7d180b758a232951e4b5c918c9dcc47aa11e99ab))

## 2.1.1 (2026-01-20)

Full Changelog: [v2.1.0...v2.1.1](https://github.com/imagekit-developer/imagekit-go/compare/v2.1.0...v2.1.1)

### Bug Fixes

* **docs:** add missing pointer prefix to api.md return types ([f3a6f9f](https://github.com/imagekit-developer/imagekit-go/commit/f3a6f9fa69c8d6c41a018fcdd5534f0dbd00b7a9))
* vocab field is required ([b18d890](https://github.com/imagekit-developer/imagekit-go/commit/b18d89015fc6f7eb213c89071fb43582ee98ca17))


### Chores

* **internal:** update `actions/checkout` version ([f0a9fc0](https://github.com/imagekit-developer/imagekit-go/commit/f0a9fc0d1c3e84934e78815169f763b468c22de6))

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
