# Changelog

## [3.4.1](https://github.com/devopsarr/terraform-provider-sonarr/compare/v3.4.0...v3.4.1) (2025-01-24)


### Bug Fixes

* **deps:** update hotio/sonarr docker tag to release-4.0.12.2823 ([bc736db](https://github.com/devopsarr/terraform-provider-sonarr/commit/bc736dbd2e2c45db023f5741ddd3149f5e8ccf29))
* **deps:** update module github.com/devopsarr/sonarr-go to v1.1.0 ([d2ad11d](https://github.com/devopsarr/terraform-provider-sonarr/commit/d2ad11d204c303cbf3816883a2fed2f57c73de09))

## [3.4.0](https://github.com/devopsarr/terraform-provider-sonarr/compare/v3.3.0...v3.4.0) (2024-12-22)


### Features

* **#341:** add quality definition preferred size field ([954efec](https://github.com/devopsarr/terraform-provider-sonarr/commit/954efecba5e6a209a29993185810a8b8ffdac8fc))
* add min upgrade format score field to quality profiles ([eb55711](https://github.com/devopsarr/terraform-provider-sonarr/commit/eb557110df91b684e1b5afa4fc19bbaca1ba5939))


### Bug Fixes

* **deps:** update hotio/sonarr docker tag to release-4.0.11.2680 ([115932d](https://github.com/devopsarr/terraform-provider-sonarr/commit/115932d361abc272bd45cd5941fe53e7351df342))
* **deps:** update module github.com/devopsarr/sonarr-go to v1.0.2 ([b07f43b](https://github.com/devopsarr/terraform-provider-sonarr/commit/b07f43b8ea00608623fa966e9aa5665c998a4fd9))
* **deps:** update module github.com/stretchr/testify to v1.10.0 ([fed1951](https://github.com/devopsarr/terraform-provider-sonarr/commit/fed19515c0c4287a5ec7078f046ba1659b575fa3))

## [3.3.0](https://github.com/devopsarr/terraform-provider-sonarr/compare/v3.2.0...v3.3.0) (2024-10-16)


### Features

* add on_import_complete notification support ([3c72a5f](https://github.com/devopsarr/terraform-provider-sonarr/commit/3c72a5fdc80f5a495a92a4c72d61ea8a5ce633ee))


### Bug Fixes

* add support for log size limit ([1ed9c19](https://github.com/devopsarr/terraform-provider-sonarr/commit/1ed9c19651194c4ffb6d11ac1d06e856e7b108f8))
* bump golangci version ([37cd12d](https://github.com/devopsarr/terraform-provider-sonarr/commit/37cd12d04d53d0a74559500d550735d020172e0e))
* **deps:** update module github.com/devopsarr/sonarr-go to v1.0.1 ([cd8d2d0](https://github.com/devopsarr/terraform-provider-sonarr/commit/cd8d2d06f377e69dccf3169fe3132d45b71331d9))
* **deps:** update terraform-framework ([374e965](https://github.com/devopsarr/terraform-provider-sonarr/commit/374e965ac1e53775f73478dcde57cfcb973cbc10))
* goreleaser syntax ([1292adf](https://github.com/devopsarr/terraform-provider-sonarr/commit/1292adf9c6818dee423222cb21f20ca3a7975147))

## [3.2.0](https://github.com/devopsarr/terraform-provider-sonarr/compare/v3.1.1...v3.2.0) (2024-03-09)


### Features

* add fields to use client extra headers ([2056b42](https://github.com/devopsarr/terraform-provider-sonarr/commit/2056b42d65bbc1f2bafa49246722bfdc243c5a80))
* align with new sdk version ([4961b78](https://github.com/devopsarr/terraform-provider-sonarr/commit/4961b78dd0f388c2a9887f6f620784c127216302))
* move to context based authentication ([e3183de](https://github.com/devopsarr/terraform-provider-sonarr/commit/e3183def5462bb7e6c2f6316996843fcd8f5606b))
* remove deprecated boxcar notification ([c1b8500](https://github.com/devopsarr/terraform-provider-sonarr/commit/c1b850018ecfcaf80e4edbf20203416e8dfd58ae))
* update go to 1.21 ([d633029](https://github.com/devopsarr/terraform-provider-sonarr/commit/d633029bc6a6f50522b3b4c17c8f47be676da152))


### Bug Fixes

* email encryption field ([e4b1976](https://github.com/devopsarr/terraform-provider-sonarr/commit/e4b1976574cc8e0095d8058cd5632b4790605163))

## [3.1.1](https://github.com/devopsarr/terraform-provider-sonarr/compare/v3.1.0...v3.1.1) (2023-10-25)


### Bug Fixes

* correct and simplify sensitive field management ([c68245f](https://github.com/devopsarr/terraform-provider-sonarr/commit/c68245f0bfaf348e602990198624f9d2fa2915fe))

## [3.1.0](https://github.com/devopsarr/terraform-provider-sonarr/compare/v3.0.0...v3.1.0) (2023-09-26)


### Features

* **devopsarr/terraform-provider-radarr#203:** add host data source ([de7d2cd](https://github.com/devopsarr/terraform-provider-sonarr/commit/de7d2cd32a755be0aa6e0540803128da2e984616))
* **devopsarr/terraform-provider-radarr#203:** add host resource ([fe89050](https://github.com/devopsarr/terraform-provider-sonarr/commit/fe89050e8da66d8d2a56173763dc3d45efea6957))
* improve diagnostics part 1 ([bfa7b62](https://github.com/devopsarr/terraform-provider-sonarr/commit/bfa7b623b30a77ac73b1094029a3c29c9b10a1fa))
* improve diagnostics part 2 ([b0037e1](https://github.com/devopsarr/terraform-provider-sonarr/commit/b0037e12f3208f3657f9e0acb7052eb0b1a0ad33))
* remove closed rarbg ([f3c0228](https://github.com/devopsarr/terraform-provider-sonarr/commit/f3c0228ef5243e6238255931ff8b1745e8526f16))
* use only ID for delete ([1f16f66](https://github.com/devopsarr/terraform-provider-sonarr/commit/1f16f6628a5d3334b148947a868c984e09e192ac))


### Bug Fixes

* delete error message ([f1ba3fd](https://github.com/devopsarr/terraform-provider-sonarr/commit/f1ba3fd3815a3ec11490a477e03d5c019fefad04))
* quality profile use all formats and ordered quality groups ([206bd72](https://github.com/devopsarr/terraform-provider-sonarr/commit/206bd728320564b85e4318740304c63a22371d7b))

## [3.0.0](https://github.com/devopsarr/terraform-provider-sonarr/compare/v3.0.0...v2.7.0) (2023-05-21)


### ⚠ BREAKING CHANGES

* remove unsupported indexer omgwtfnzbs
* remove deprecated language profile

### Features

* add auto tag condition data source ([acd22cb](https://github.com/devopsarr/terraform-provider-sonarr/commit/acd22cbf13a80260baa7bfe31fbaf5698dd1b446))
* add auto tag condition genres data source ([279cedf](https://github.com/devopsarr/terraform-provider-sonarr/commit/279cedf4cc0eaa8c803808dbbb70d95c8d5eff1b))
* add auto tag condition root folder data source ([f818eb1](https://github.com/devopsarr/terraform-provider-sonarr/commit/f818eb1c75df16a6e69ad6e433b6f50875397020))
* add auto tag condition series type data source ([6b295dc](https://github.com/devopsarr/terraform-provider-sonarr/commit/6b295dcd3043bc68884335d77d1941928d9ca5c1))
* add auto tag data source ([c49f93a](https://github.com/devopsarr/terraform-provider-sonarr/commit/c49f93aa1ca39ed2805aba2287e44f952d2b563e))
* add auto tag resource ([d9b3932](https://github.com/devopsarr/terraform-provider-sonarr/commit/d9b3932ddacf8344e972f54966ec655a69fd5ca6))
* add auto tags data source ([7d685bc](https://github.com/devopsarr/terraform-provider-sonarr/commit/7d685bce962d024896de5b5e2e3eaa9260febdbd))
* add custom format condition data source ([b747056](https://github.com/devopsarr/terraform-provider-sonarr/commit/b747056e9ab8aabdfdc8b99bd53290ca4f809ce9))
* add custom format condition language data source ([45c5fd0](https://github.com/devopsarr/terraform-provider-sonarr/commit/45c5fd0cb43a46ca99019c662dfa7d1ff58c6aeb))
* add custom format condition release group data source ([e7a44af](https://github.com/devopsarr/terraform-provider-sonarr/commit/e7a44afdc9ddf73a73f0bc745f19c95c49d0de02))
* add custom format condition release title data source ([c080e87](https://github.com/devopsarr/terraform-provider-sonarr/commit/c080e87371856eddcc1065a5384cdbcf49c96a3e))
* add custom format condition resolution data source ([afac201](https://github.com/devopsarr/terraform-provider-sonarr/commit/afac201caf7e2299547dd0e04d508f1b98c4b692))
* add custom format condition size data source ([1a96419](https://github.com/devopsarr/terraform-provider-sonarr/commit/1a964196fe6339849158988823da1fc54830a7cd))
* add custom format condition source data source ([b583eb1](https://github.com/devopsarr/terraform-provider-sonarr/commit/b583eb15adcc29880fa5a6a937720cc31928810d))
* add custom format datasource ([62d01ab](https://github.com/devopsarr/terraform-provider-sonarr/commit/62d01ab3e7edc844e5c98f5e6fde0dd76c0c80ca))
* add custom format resource ([ec9cb73](https://github.com/devopsarr/terraform-provider-sonarr/commit/ec9cb73e3b868a4dca64b58193638be67a5c8762))
* add custom formats datasource ([a015810](https://github.com/devopsarr/terraform-provider-sonarr/commit/a015810749c47bee3d2160b51020e30eceae67ac))
* add import list custom resource ([a34ab91](https://github.com/devopsarr/terraform-provider-sonarr/commit/a34ab91fa534c8dc271fa12a6b24c2a39ef7c7d6))
* add import list exclusion datasource ([dd56755](https://github.com/devopsarr/terraform-provider-sonarr/commit/dd56755d5488b28d745cfc120ea278742e1ed8d6))
* add import list exclusion resource ([cae99bb](https://github.com/devopsarr/terraform-provider-sonarr/commit/cae99bbe8e26373ab6601cefa86400411bcd2278))
* add import list exclusions datasource ([1195ef0](https://github.com/devopsarr/terraform-provider-sonarr/commit/1195ef0aca007705280974dede1ae59297802dbd))
* add import list imdb resource ([d918ed8](https://github.com/devopsarr/terraform-provider-sonarr/commit/d918ed8dc445305aac82356565533730cde6b3cb))
* add import list plex rss resource ([ac270fd](https://github.com/devopsarr/terraform-provider-sonarr/commit/ac270fd0a28598351b38dd9ae64f68b6a19736d2))
* add language data source ([da546a2](https://github.com/devopsarr/terraform-provider-sonarr/commit/da546a2b53b7e4f9a48fe24cd6b93041099cc383))
* add languages data source ([68609a1](https://github.com/devopsarr/terraform-provider-sonarr/commit/68609a1ad600eaaf43086efe6a560d92f9116cb0))
* add metadata consumers data source ([6a738b8](https://github.com/devopsarr/terraform-provider-sonarr/commit/6a738b893dbd76a8ab6a4085aafe9622caf78ca1))
* add metadata data source ([7ac7994](https://github.com/devopsarr/terraform-provider-sonarr/commit/7ac79946e6e893bc5d213bf7c675c8f66a7f87ba))
* add metadata kodi resource ([96c1bb0](https://github.com/devopsarr/terraform-provider-sonarr/commit/96c1bb038387f8242cf5b730d49a763f848880a4))
* add metadata resource ([32c046e](https://github.com/devopsarr/terraform-provider-sonarr/commit/32c046eaf29c4fee71be14513d933a38d2495780))
* add metadata roksbox resource ([96c2749](https://github.com/devopsarr/terraform-provider-sonarr/commit/96c27498fb25732164275e143a2bf662515af944))
* add metadata wdtv resource ([ef59c0d](https://github.com/devopsarr/terraform-provider-sonarr/commit/ef59c0d70398e48b31133588418161d27ffe1ed0))
* add naming colon replacement format ([237f21a](https://github.com/devopsarr/terraform-provider-sonarr/commit/237f21ae0e772ab63709209f0c7ff9b6a6e6123b))
* add new notification flags ([ca7e03c](https://github.com/devopsarr/terraform-provider-sonarr/commit/ca7e03c927d26fc406b6d3b847e07f25927ca3c5))
* add notificaiton apprise resource ([3523b11](https://github.com/devopsarr/terraform-provider-sonarr/commit/3523b115ea211e25645054f9f2df102c2d55f9e4))
* add notification ntfy resource ([43446ef](https://github.com/devopsarr/terraform-provider-sonarr/commit/43446ef536f79719d4f6ee8b2af96d5dbd71e417))
* add notification signal resource ([ead9e09](https://github.com/devopsarr/terraform-provider-sonarr/commit/ead9e090e77dbc76e44c7b0112b6f15d0a174548))
* add notification simplepush resource ([10b38fe](https://github.com/devopsarr/terraform-provider-sonarr/commit/10b38fe37a4e442757a7279f73b62ba18e0bfc29))
* add quality data source ([628ff46](https://github.com/devopsarr/terraform-provider-sonarr/commit/628ff468f49226f323825b6928a4a1f6f5769e2e))
* remove deprecated language profile ([5a76731](https://github.com/devopsarr/terraform-provider-sonarr/commit/5a76731b65a6057f61b53fbf6c1128d58fb694d0))
* remove obsolete languages helper ([a9cf584](https://github.com/devopsarr/terraform-provider-sonarr/commit/a9cf584df20e2d2ad866cd83c7ae88937b623f66))
* remove unsupported indexer omgwtfnzbs ([8d515ce](https://github.com/devopsarr/terraform-provider-sonarr/commit/8d515ce0ea2175810920964c76f3b8909295134f))
* set indexer flags to optional ([7a291e4](https://github.com/devopsarr/terraform-provider-sonarr/commit/7a291e426dd7b9c04aed00b3dd4eef287db8232a))
* update delay profile for custom format ([d262ace](https://github.com/devopsarr/terraform-provider-sonarr/commit/d262acec46142b53462c2a85adeba14433bd4b8e))
* update quality profile with custom formats ([f656ea3](https://github.com/devopsarr/terraform-provider-sonarr/commit/f656ea39b5dfe62067e3030d17999c7cf95830c0))


### Bug Fixes

* add back import list language profile ids ([74714ed](https://github.com/devopsarr/terraform-provider-sonarr/commit/74714edcf023482077feeae723943d28ea5372a3))
* quality profile failing with one item group ([81b59fd](https://github.com/devopsarr/terraform-provider-sonarr/commit/81b59fd517526f4ebf76d5927022301c1eb3e991))
* read data source from request ([f8e65ab](https://github.com/devopsarr/terraform-provider-sonarr/commit/f8e65abcea47be0dc569120601720a3bff848094))
* some download clients has wrong parameters ([61d5451](https://github.com/devopsarr/terraform-provider-sonarr/commit/61d5451b9670e022d05cc757f08bc22959c51c3a))
* update sdk method naming ([0414d62](https://github.com/devopsarr/terraform-provider-sonarr/commit/0414d621c5893fa140cf9aa67131c460b099f4f1))
* update sensitive data management ([642ffb7](https://github.com/devopsarr/terraform-provider-sonarr/commit/642ffb72bc2b44d41892718590df7bf8edc75d7c))
* use get function for sdk fields ([c07b9e6](https://github.com/devopsarr/terraform-provider-sonarr/commit/c07b9e6950b010f545dc75de2300efddc7a665df))

## [2.7.0](https://github.com/devopsarr/terraform-provider-sonarr/compare/v2.6.0...v2.7.0) (2022-12-13)


### Features

* add import list datasource ([3f0a3e2](https://github.com/devopsarr/terraform-provider-sonarr/commit/3f0a3e2f1f0d0e8c8fbd7a4a39fd75a66487d5f3))
* add import list plex resource ([0c434e0](https://github.com/devopsarr/terraform-provider-sonarr/commit/0c434e0481f6480e89ce549bd5eab07adabb9b84))
* add import list resource ([abc6b1c](https://github.com/devopsarr/terraform-provider-sonarr/commit/abc6b1cc12f45e3668923fd598c1a551aca2a033))
* add import list sonarr resource ([1b24ae1](https://github.com/devopsarr/terraform-provider-sonarr/commit/1b24ae1ce53cb4bacd107c8629064cabcc0a40b8))
* add import list trakt list resource ([00d4a01](https://github.com/devopsarr/terraform-provider-sonarr/commit/00d4a01558421696ebf0da21fd2b8026eca3cf9b))
* add import list trakt popular resource ([c2f37ad](https://github.com/devopsarr/terraform-provider-sonarr/commit/c2f37ad8c821e2f0bb782bc1ab8bbcd378bae890))
* add import list trakt user resource ([bcd309b](https://github.com/devopsarr/terraform-provider-sonarr/commit/bcd309bb9e24965100bfaff5fead8c7c7cb68ffe))
* add import lists datasource ([2411b4e](https://github.com/devopsarr/terraform-provider-sonarr/commit/2411b4ecf156dbc965cc12c57d6d1f85bcb2c45d))
* add search series datasource ([0f2f5cb](https://github.com/devopsarr/terraform-provider-sonarr/commit/0f2f5cb0258c2ebcf58625f09b040983c959360b))

## [2.6.0](https://github.com/devopsarr/terraform-provider-sonarr/compare/v2.5.0...v2.6.0) (2022-12-06)


### Features

* add quality definition datasource ([b2a06db](https://github.com/devopsarr/terraform-provider-sonarr/commit/b2a06dbba80cbac772f59ecd583d6448c60b73de))
* add quality definition resource ([efebfb4](https://github.com/devopsarr/terraform-provider-sonarr/commit/efebfb49c57abd46ce56ae453c7a7636a9854078))
* add quality definitions datasource ([566d657](https://github.com/devopsarr/terraform-provider-sonarr/commit/566d657bbce6737f5560783d516b5c87cdbff879))

## [2.5.0](https://github.com/devopsarr/terraform-provider-sonarr/compare/v2.4.0...v2.5.0) (2022-12-06)


### Features

* add notification boxcar resource ([5366ec2](https://github.com/devopsarr/terraform-provider-sonarr/commit/5366ec29fd5298186c4e96b2c89ac0b4f1116446))
* add notification discord resource ([b1e0303](https://github.com/devopsarr/terraform-provider-sonarr/commit/b1e030370ec1391b045a12697966b474c57f593c))
* add notification email resource ([8344c92](https://github.com/devopsarr/terraform-provider-sonarr/commit/8344c92be1e1bd47a5972d23fe69d6fceec59558))
* add notification emby resource ([be53d09](https://github.com/devopsarr/terraform-provider-sonarr/commit/be53d090e78f556166c316fae2287af89a987a6b))
* add notification gotify resource ([cb17bb9](https://github.com/devopsarr/terraform-provider-sonarr/commit/cb17bb99a668d83bdbae8b928bc551fdb1717d1c))
* add notification join resource ([446374e](https://github.com/devopsarr/terraform-provider-sonarr/commit/446374e20566c14a4f7d88e5aa5ec4e97b1e11f9))
* add notification kodi resource ([6b75f52](https://github.com/devopsarr/terraform-provider-sonarr/commit/6b75f5249fe85a33ef045a98a2b94ff9cd7fbfcd))
* add notification mailgun resource ([35c60a8](https://github.com/devopsarr/terraform-provider-sonarr/commit/35c60a89ba695e975342eb74de1d2897c406d226))
* add notification plex resource ([c6b0948](https://github.com/devopsarr/terraform-provider-sonarr/commit/c6b09480eebdf0c90487245cce93a4c11d524edc))
* add notification prowl resource ([9ed0337](https://github.com/devopsarr/terraform-provider-sonarr/commit/9ed03376aa63ce08bddee77073117b87266056ed))
* add notification pushbullet resource ([3735a00](https://github.com/devopsarr/terraform-provider-sonarr/commit/3735a00dee5c9064c989364a6bf9070e0433d1ec))
* add notification pushover resource ([37c8912](https://github.com/devopsarr/terraform-provider-sonarr/commit/37c8912332bc000290ae71822202f0716636b548))
* add notification sendgrid resource ([0e6e2a7](https://github.com/devopsarr/terraform-provider-sonarr/commit/0e6e2a7ed06bd73063dd1d694af6d17b3db0293d))
* add notification slack resource ([7afbf19](https://github.com/devopsarr/terraform-provider-sonarr/commit/7afbf194790da4500f4162b919912e74d0e292a3))
* add notification synology indexer resource ([8a2973e](https://github.com/devopsarr/terraform-provider-sonarr/commit/8a2973e451547873f99d5e9b5b4ccdb8b2a21066))
* add notification telegram resource ([ac3960b](https://github.com/devopsarr/terraform-provider-sonarr/commit/ac3960b912c7fe1cbf9e8d6e84d9ad5bf6030c1b))
* add notification twitter resource ([b82990d](https://github.com/devopsarr/terraform-provider-sonarr/commit/b82990dc9acd667a25f210710377d39497e9c70a))


### Bug Fixes

* correct few notification types ([203d1ae](https://github.com/devopsarr/terraform-provider-sonarr/commit/203d1aecd6c56b44089941bae07e7f67ba8934cc))
* notification email fields type ([53b5dcd](https://github.com/devopsarr/terraform-provider-sonarr/commit/53b5dcdabe4dd7ddc015327f625bef87826cf5c7))
* removed unused on_rename ([51bafc9](https://github.com/devopsarr/terraform-provider-sonarr/commit/51bafc96435e5ae60b853ca0c0d3cf9a81170110))
* tags behaviour with null and empty array ([8936dc7](https://github.com/devopsarr/terraform-provider-sonarr/commit/8936dc7d02e305d30f93f870b507d58e222aadb4))
* update grabfields and importfields type ([ed4472a](https://github.com/devopsarr/terraform-provider-sonarr/commit/ed4472a42e95c48f67ca156a3290284749810cd5))

## [2.4.0](https://github.com/devopsarr/terraform-provider-sonarr/compare/v2.3.0...v2.4.0) (2022-11-25)


### Features

* add download client aria2 resource ([bed4a61](https://github.com/devopsarr/terraform-provider-sonarr/commit/bed4a61938af886d0cd8b875729cff7501b7cecb))
* add download client deluge resource ([158c17b](https://github.com/devopsarr/terraform-provider-sonarr/commit/158c17b2bb4a6dd6d6bb800039ae17f477469a51))
* add download client flood resource ([61871a2](https://github.com/devopsarr/terraform-provider-sonarr/commit/61871a2e350467eb5d80fac86500e73ee3d47e64))
* add download client hadouken resource ([b1a4c00](https://github.com/devopsarr/terraform-provider-sonarr/commit/b1a4c004b577d27458775a014c0151dd237fcac7))
* add download client nzbget resource ([0c11272](https://github.com/devopsarr/terraform-provider-sonarr/commit/0c1127230e1aabec2d81fb241f0a0e96e0724d4e))
* add download client nzbvortex resource ([d673fd9](https://github.com/devopsarr/terraform-provider-sonarr/commit/d673fd9c910f1a2cf09d05b2305923ba9efaf57d))
* add download client pneumatic resource ([f05cdfd](https://github.com/devopsarr/terraform-provider-sonarr/commit/f05cdfddd96acc0868e55c81c9a309091c48f4f8))
* add download client qbittorrent resource ([0f3f24f](https://github.com/devopsarr/terraform-provider-sonarr/commit/0f3f24fdb15e82778487415b99aa6c9ad022a39c))
* add download client rtorrent resource ([175cc6c](https://github.com/devopsarr/terraform-provider-sonarr/commit/175cc6cba98b017fa204b1022e0f096c80c83847))
* add download client sabnzbd resource ([bd9c190](https://github.com/devopsarr/terraform-provider-sonarr/commit/bd9c19054683fec3b96acc6273c45c62016e45c3))
* add download client torrent blackhole resource ([ad5c855](https://github.com/devopsarr/terraform-provider-sonarr/commit/ad5c85579af3fa79b52abff7fca85c99ad0b0546))
* add download client torrent download station ([cf66358](https://github.com/devopsarr/terraform-provider-sonarr/commit/cf66358002c6a9c5286fb380c673b9ada015d860))
* add download client usenet blackhole resource ([dcaa5fe](https://github.com/devopsarr/terraform-provider-sonarr/commit/dcaa5fe578b108035420e4aff0d2b61b89004c64))
* add download client usenet download station resource ([3c7aee2](https://github.com/devopsarr/terraform-provider-sonarr/commit/3c7aee226ddfbab69002c7139dd9a6ba85a8641b))
* add download client utorrent resource ([7c56c38](https://github.com/devopsarr/terraform-provider-sonarr/commit/7c56c3803cbc1028662a287b5415d76c5cee58de))
* add download client vuze resource ([7733bb3](https://github.com/devopsarr/terraform-provider-sonarr/commit/7733bb35080caa2828c3d37fb17e373646f2f59c))
* add fields exceptions ([ca7e4f5](https://github.com/devopsarr/terraform-provider-sonarr/commit/ca7e4f5d827af85e3259c8cc9a70119c00aa957b))
* add indexer broadcasthenet resource ([cbe27f4](https://github.com/devopsarr/terraform-provider-sonarr/commit/cbe27f4c1d9b5dccdb6b40d6fbd3fd6561a672d5))
* add indexer fanzub resource ([187af1c](https://github.com/devopsarr/terraform-provider-sonarr/commit/187af1c9c4acdcb4891a0a7500c71b64665e803d))
* add indexer filelist resource ([8489b2d](https://github.com/devopsarr/terraform-provider-sonarr/commit/8489b2dad65ab1336a5cec3ed0d6964ef49106de))
* add indexer hdbits resource ([f2c9cd2](https://github.com/devopsarr/terraform-provider-sonarr/commit/f2c9cd2412adb09a403a4f7cdb59e9efefef26e3))
* add indexer iptorrents resource ([e43c573](https://github.com/devopsarr/terraform-provider-sonarr/commit/e43c573bd4c9354b53ad587c136b87d10839784d))
* add indexer nyaa resource ([ef6df58](https://github.com/devopsarr/terraform-provider-sonarr/commit/ef6df5817af79cb5b648961b65559a7e1fa06a59))
* add indexer omgwtfnzbs resource ([bf74074](https://github.com/devopsarr/terraform-provider-sonarr/commit/bf74074fa221133f671a70dd62614a2019eac0aa))
* add indexer orrentleech resource ([c5d1bf5](https://github.com/devopsarr/terraform-provider-sonarr/commit/c5d1bf53c92bbfeb6c0df8e1081bc35f47a5e2ba))
* add indexer torrent rss resource ([9d4098d](https://github.com/devopsarr/terraform-provider-sonarr/commit/9d4098d4e898f4e7bba8f61b0d6aa76a524a81ab))
* add indexer torznab resource ([5849a65](https://github.com/devopsarr/terraform-provider-sonarr/commit/5849a6549fbe2037b664518389de88516e4784ea))


### Bug Fixes

* download client ds watch folder type ([f4afe44](https://github.com/devopsarr/terraform-provider-sonarr/commit/f4afe446651555975c908b76cbfb4c7621767ab1))
* download client watch_folder field to string ([c240caa](https://github.com/devopsarr/terraform-provider-sonarr/commit/c240caad8885e1d2ef2bf9430101c9b6caab55ec))
* seedcriteria fields ([69466c6](https://github.com/devopsarr/terraform-provider-sonarr/commit/69466c61c3be6f227ca1395e9a2ab4489a55fdcf))
* sensitive fields ([b24abed](https://github.com/devopsarr/terraform-provider-sonarr/commit/b24abed139fb53dfa3241b291e0b7b9070eb53bf))
* **tool:** WriteStringField work with other types ([a40d15f](https://github.com/devopsarr/terraform-provider-sonarr/commit/a40d15f2bf5e8e48936017d90b148222ae39014c))
* torrent blackhole required parameters ([52ccc52](https://github.com/devopsarr/terraform-provider-sonarr/commit/52ccc5237b507148f676ac42bd43d08ee942c9e4))

## [2.3.0](https://github.com/devopsarr/terraform-provider-sonarr/compare/v2.2.0...v2.3.0) (2022-11-03)


### Features

* add notification webhook resource ([114a34e](https://github.com/devopsarr/terraform-provider-sonarr/commit/114a34ecd3c3c47e5eba5dbf91b75d099a8929e3))


### Bug Fixes

* delayprofile order on resource creation ([a64a48e](https://github.com/devopsarr/terraform-provider-sonarr/commit/a64a48e9676eeda13fd6ffab2933fdcfb656eae1))

## [2.2.0](https://github.com/devopsarr/terraform-provider-sonarr/compare/v2.1.0...v2.2.0) (2022-11-02)


### Features

* add download client transmission ([07c14f3](https://github.com/devopsarr/terraform-provider-sonarr/commit/07c14f313f62103a5825dad824014913f7b1ea68))
* add field helpers ([7502cb1](https://github.com/devopsarr/terraform-provider-sonarr/commit/7502cb167717d116a944ee2e0c3d42870a369b66))
* add notification custom script ([3325ff5](https://github.com/devopsarr/terraform-provider-sonarr/commit/3325ff5cbd39b21c679fe3f7891a3235818bc8dd))
* add rarbg indexer resource ([e69d9be](https://github.com/devopsarr/terraform-provider-sonarr/commit/e69d9be2c0694d0fbac530c221b6e2c3f116039a))
* add release profile datasource ([891dc5c](https://github.com/devopsarr/terraform-provider-sonarr/commit/891dc5cf005743c86ef955b690b627fd97912cba))
* add release profile resource ([388db30](https://github.com/devopsarr/terraform-provider-sonarr/commit/388db302572a7c2866b412ffa0fa80ac0f0e19e5))
* add release profiles datasource ([bbdc254](https://github.com/devopsarr/terraform-provider-sonarr/commit/bbdc2544f093389df19e59704371077be5323795))
* add zewznab indexer ([760cef7](https://github.com/devopsarr/terraform-provider-sonarr/commit/760cef744745984a31458836c418996ccd711d23))


### Bug Fixes

* remove useless init calls ([71fa6de](https://github.com/devopsarr/terraform-provider-sonarr/commit/71fa6defed5bc5c731be2d4b6e5ef107465a0b80))

## [2.1.0](https://github.com/devopsarr/terraform-provider-sonarr/compare/v2.0.0...v2.1.0) (2022-09-22)


### Features

* add download client ([25a1a85](https://github.com/devopsarr/terraform-provider-sonarr/commit/25a1a85cca35159d14e791564587a1c033c74f0b))
* add download client config datasource ([eda31c2](https://github.com/devopsarr/terraform-provider-sonarr/commit/eda31c2123e26d4abe329b7882d456768ad8870c))
* add download client config resource ([2791a4a](https://github.com/devopsarr/terraform-provider-sonarr/commit/2791a4a637bdbfa1a5c74526311ade6d44a61ad2))
* add download client datasource ([e06e3a2](https://github.com/devopsarr/terraform-provider-sonarr/commit/e06e3a250e66b7d26ad378f652b677f7a6de212d))
* add download clients datasource ([bb8c794](https://github.com/devopsarr/terraform-provider-sonarr/commit/bb8c7944ba079d3e17d5781a52ba7cfdf6b3b351))
* add int match helper ([7dc4fd4](https://github.com/devopsarr/terraform-provider-sonarr/commit/7dc4fd4b2f2bfbd478d3ac5307367836d6ee033d))
* add notification datasource ([65a91e5](https://github.com/devopsarr/terraform-provider-sonarr/commit/65a91e5c132b157428f3feed9f8513b6772051a1))
* add notification resource ([ea406f1](https://github.com/devopsarr/terraform-provider-sonarr/commit/ea406f1c5f9d87f72d96ce9e1ee7544e900f1e01))
* add notifications datasource ([bae788f](https://github.com/devopsarr/terraform-provider-sonarr/commit/bae788f16f699f6d428be9f61b5315de8f1eb077))
* add remote path mapping data source ([18b347a](https://github.com/devopsarr/terraform-provider-sonarr/commit/18b347a3d330aaac18e7c1994aad40052529de12))
* add remote path mapping resource ([7d1a6a6](https://github.com/devopsarr/terraform-provider-sonarr/commit/7d1a6a602e45ce8fc57bb6ae4f2d623a2964e995))
* add remote path mappings data source ([a1d93a9](https://github.com/devopsarr/terraform-provider-sonarr/commit/a1d93a9f04e4ac0d67555c92ee7ae07d34e81d52))
* add system status datasource ([6613dff](https://github.com/devopsarr/terraform-provider-sonarr/commit/6613dffa5e5eb797408b64c63624f4a1f4bafc80))

## [2.0.0](https://github.com/devopsarr/terraform-provider-sonarr/compare/v1.2.2...v2.0.0) (2022-09-01)


### ⚠ BREAKING CHANGES

* move series to all_series

### Features

* add delay profile data source ([7118dd1](https://github.com/devopsarr/terraform-provider-sonarr/commit/7118dd14e18d124bdb5c58570254a503591f065e))
* add indexer config data source ([049cfe5](https://github.com/devopsarr/terraform-provider-sonarr/commit/049cfe5affd6f54b29094393b98bd289407e141a))
* add indexer data source ([af79355](https://github.com/devopsarr/terraform-provider-sonarr/commit/af79355f18f49ea80bb286c31af6f24ad6944b57))
* add language profile data source ([c980f9d](https://github.com/devopsarr/terraform-provider-sonarr/commit/c980f9d743d8f6c2ee03428282cacbbd1579145d))
* add media management data source ([6ee6c80](https://github.com/devopsarr/terraform-provider-sonarr/commit/6ee6c800357c2a3412c6d48c5706c8a8ee28f66a))
* add naming data source ([c4f5c8e](https://github.com/devopsarr/terraform-provider-sonarr/commit/c4f5c8e9e30866df18629662c9bdf65a62846690))
* add quality profile data soruce ([3f46e2b](https://github.com/devopsarr/terraform-provider-sonarr/commit/3f46e2b9ade3d839b8e91299a47c17f5902c2538))
* add root folder data source ([1958c1f](https://github.com/devopsarr/terraform-provider-sonarr/commit/1958c1f00b9c210103ca6bc406c1013f10a90e02))
* add series data source ([2a7f3ca](https://github.com/devopsarr/terraform-provider-sonarr/commit/2a7f3caf93c31073cc55ebe17cbfc7a0d5e98737))
* add tag data source ([4c260e0](https://github.com/devopsarr/terraform-provider-sonarr/commit/4c260e053b256a45e9b23b64a7a7f2f43593ccfd))
* move series to all_series ([b98d12b](https://github.com/devopsarr/terraform-provider-sonarr/commit/b98d12b38a38ed3002c41b838af1742a0814297e))


### Bug Fixes

* data source read from config ([ec64614](https://github.com/devopsarr/terraform-provider-sonarr/commit/ec646149bf6fcc876aac91c69d9e541d5df6c77d))
* list data source requirements ([c3aa214](https://github.com/devopsarr/terraform-provider-sonarr/commit/c3aa2140131e12a2892f0de5e5414a41f8aafb95))

## [1.2.2](https://github.com/devopsarr/terraform-provider-sonarr/compare/v1.2.1...v1.2.2) (2022-08-30)


### Bug Fixes

* indexer fields type assertion ([d4a746c](https://github.com/devopsarr/terraform-provider-sonarr/commit/d4a746ceaeb88483f21128d8399dfa1adc18ad9c))

## [1.2.1](https://github.com/devopsarr/terraform-provider-sonarr/compare/v1.2.0...v1.2.1) (2022-08-28)


### Bug Fixes

* release gpg dependencies ([b2e7065](https://github.com/devopsarr/terraform-provider-sonarr/commit/b2e7065aa4869219981759cf2af5aa06497d1db8))

## [1.2.1](https://github.com/devopsarr/terraform-provider-sonarr/compare/v1.2.0...v1.2.1) (2022-08-28)


### Bug Fixes

* release gpg dependencies ([0c5a8da](https://github.com/devopsarr/terraform-provider-sonarr/commit/0c5a8da58d37e9443daad242105ca47dbaf82c4e))

## [1.2.0](https://github.com/devopsarr/terraform-provider-sonarr/compare/v1.1.0...v1.2.0) (2022-08-28)


### Features

* add indexer config resource ([375172e](https://github.com/devopsarr/terraform-provider-sonarr/commit/375172e0a572efa3f7293db043bcba7aaa1cacec))
* add indexer resource ([b0602c5](https://github.com/devopsarr/terraform-provider-sonarr/commit/b0602c5b77e7939d6e81cf5ed7212154b0bb932a))
* add indexers data source ([1b2cc84](https://github.com/devopsarr/terraform-provider-sonarr/commit/1b2cc8479b7da70727edc2ba4fcabc289c272934))


### Bug Fixes

* remove set parameter for framework 0.9.0 ([370a452](https://github.com/devopsarr/terraform-provider-sonarr/commit/370a45213b53f87e46516f1404e4310c9ba384ae))
* series and qp datasource set ([3bdac5a](https://github.com/devopsarr/terraform-provider-sonarr/commit/3bdac5a62b55dd04b60cdefe19ec2ae79fa23308))

## [1.1.0](https://github.com/devopsarr/terraform-provider-sonarr/compare/v1.0.0...v1.1.0) (2022-06-06)


### Features

* add media_management resource ([00bfbc5](https://github.com/devopsarr/terraform-provider-sonarr/commit/00bfbc5a9d6d41e6d81b213d3ab228a29602b3dc))
* add naming resource ([79fc714](https://github.com/devopsarr/terraform-provider-sonarr/commit/79fc7142235d79f5cb23d1f5c2566ed8315b4d19))
* add profile validators ([53ce110](https://github.com/devopsarr/terraform-provider-sonarr/commit/53ce1105f5a8ff561039f7e66beb19bd57e425ad))
* add root_folder resources ([6a2d0bd](https://github.com/devopsarr/terraform-provider-sonarr/commit/6a2d0bd4942e9478b58f3b87d2f0f6104e57f184))
* add series resource ([3ec587a](https://github.com/devopsarr/terraform-provider-sonarr/commit/3ec587a31e3c73336ef3ff289ded1644149a23e6))

## 1.0.0 (2022-03-15)


### Features

* add profiles ([3f37682](https://github.com/devopsarr/terraform-provider-sonarr/commit/3f376822edaecbab1129ad202c093eef5d22ff4f))
* add tag resource ([98bb542](https://github.com/devopsarr/terraform-provider-sonarr/commit/98bb5426c912466173e61d1855064110b55ba384))
* align tag to sdk ([987ac70](https://github.com/devopsarr/terraform-provider-sonarr/commit/987ac702ec0f4bf3dd9e0b6a9b3ac119a036f7db))
* first configuration ([b6e0b48](https://github.com/devopsarr/terraform-provider-sonarr/commit/b6e0b486af8a547a2326c3fe8989c0d1abf77a6b))
