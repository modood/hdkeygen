btckeygen <img src="https://www.buybitcoinworldwide.com/img/segwit.png" width="100">
=========

[![license](https://img.shields.io/badge/license-WTFPL%20--%20Do%20What%20the%20Fuck%20You%20Want%20to%20Public%20License-green.svg)](https://github.com/modood/btckeygen/blob/master/LICENSE)

A simple and easy to use bitcoin key generator

*   BIP32 - Hierarchical Deterministic Wallets
*   BIP39 - Mnemonic code for generating deterministic keys
*   BIP43 - Purpose Field for Deterministic Wallets
*   BIP44 - Multi-Account Hierarchy for Deterministic Wallets
*   BIP49 - Derivation scheme for P2WPKH-nested-in-P2SH based accounts
*   BIP84 - Derivation scheme for P2WPKH based accounts
*   BIP173 - Base32 address format for native v0-16 witness outputs
*   SLIP44 - Registered coin types for BIP-0044


Can I trust this code?
----------------------

> Don't Trust. Verify.

> We recommend every user of this library audit and verify any underlying code for its validity and suitability.
>
> You can do so by using this tool: https://iancoleman.io/bip39/

Installation
------------

```
$ go get -u github.com/modood/btckeygen
```

Usage
-----

```
Usage of btckeygen:
  -bip39     bool   mnemonic code for generating deterministic keys (default false)
  -pass      string protect bip39 mnemonic with a passphrase (default "")
  -n         int    set number of keys to generate (default 20)
```

Example
-------

**generate compressed public key:**

```txt
$ btckeygen

Bitcoin Address                    WIF(Wallet Import Format)                            SegWit(bech32)                             SegWit(nested)
---------------------------------------------------------------------------------------------------------------------------------------------------------------------
18nndBdh9VpfBsw8DJKg6KzYrC2o8v2KPo Kz1nmzu2DcLnHQy9kLbr794RAspLTf3ocx2jD2UnZRHN2EGuregS bc1q24c67qhugfla6zl3jj7x5uc48y6z5q7eecua4x 38fupM9MgfLj7fu2DWyXqnk9N9Um757G7D
1zdJnv2PZjP8EkfNJo27sotRntNyZCjEw  L3KR1wo35ngz4ATuAcciEXRrYAyosSB6s6XUP43H834hXpNqNNqC bc1qptnqgetfx7y220hwy3tl565mc73s75nq05u4mu 3HkC5ajZungAeFg2KjLUpV9Xa8Ahfw3Sqz
1Fw7MmDLMU5vGiLLqMKDdwRTPPMUaZi5vZ L1yEQi2D3YUyxcyNDmkUS2aHk8Tw3q7sF2DLEiamMr8gF79xhzXr bc1q50xc0sp2jnljv93el3at0v62gtv3dtwc2neh4q 338J9irj1kUN3CLVCo92oMizvPaYohq4BV
1Q4jALMrXo7iV5aGhgvFKoFgPxPFyRCMBU L2epjvSWacyZ1yUxLs6XUaJHCNPRzkTRZ89nxuMDM7eAQSKG1m2A bc1qlnl4t3l0sn0sllyu749qnj94edsstspd2w36ft 3GM6dMTEUrXi2oUhffqWzbwdUBb5wcFFRH
1MkQLmfcEcrjHDWVu3jZDQjCatZKzNZqWv KwHTiPwM5YpspUaHhvT7Cd5cePcXD1NtiygKpkMebo42jYPidhSU bc1quwt795fwxemp7vq5328xptvjwwpeh2036etjk7 3KRYMYGyNg8tZh9z43MDHowZ9SLqVyxvvx
17TxrHkbZftp2bxDehpikZWmEZvMvnzSXp L1RzZbyGa8EFfodabAB9jCcekR8FkzDDkmXvqTKkPvfhwN178h68 bc1qgm4rp4ha8fxuhqz40klqu65t528uz2q2gwlcvp 3FD42UuPBKwAfoXbEd5oCcUvrebpuhwDtW
1FmHRVvhtvL6GCsJEtpbHx7PqHMH13iZR9 Kx8akR9U1tdUSt24pW3WgAapDCiyf5t915YfHmQCSQ2zETHVN4fX bc1q58cutg45zz9d20jr6nsah0rf2vtx7falvl0jgc 3PHHzdZhM2QD4fDQHS7bnYx3ArCzYeQ3rv
1LHgqpG2zbW9PnhgNN5pNMJR94zZQrrjho L3AeREULfYV5Jq5quM7v5iYLPQ9aME6tEXnPQrDiCaxfGdTqMKFu bc1q6wfym27yfmvk4j3kvetjm0900wnzum4ys8a82q 34UKCs4q2NoG8MqZ17KNnNh1hfS173pMxp
1BC6VCkmTqsqRPMyThAjBtMW7QmpeUP6pa KzP7DYWscsvh5RdRnNh1nabjp7H431GxkEJLdxHJ9qZiHKr8kKQA bc1qdl9xnsvlj0ktqyy27v7lnkryq5t83yfgucuks8 3Ax9TfbihnUecfQFyfFMcN3wCAqcbhY5nX
13SY8L4wpmP45LtWDaX187WKDc864H1Qnf L5MprHUJ3HohKCi2MYLW7BR7ah321kqcQtecao9PyXsZ2xA4wb5P bc1qrtz2yvkg2dxrtqwdl08eaf0tkmd4ly9wct42tf 3JJGAp2cXkTpZGHCVWugLgEryyJ2nMC2po

```

**generate bip39 mnemonic with passphrase:**

```txt
$ btckeygen -bip39 -pass=123456 -n 10

BIP39 Mnemonic:    voyage blind unit shoulder yellow attitude mule all hire above obvious swap
BIP39 Passphrase:  123456
BIP39 Seed:        f4b8043e3b3b4d0b9e3c7cda81d6868c331aaecc80555dc7b2d0edce6b73ea50a91d67586f7461cd46caccee6e240a598a9aaa3063cdd9bec65a3d24d3aa551b
BIP32 Root Key:    xprv9s21ZrQH143K45MBGTeN7zrQxBgh7v3XNtAMrQvYBfm6xdtaVkjCFNyFHZ262PpMoiaA8JEFGUDPVV6qzB459nGgR1mjuigdTaG2NsKr5BG

Path(BIP44)        Bitcoin Address                    WIF(Wallet Import Format)
----------------------------------------------------------------------------------------------------------
m/44'/0'/0'/0/0    1LaDeWmc3oTSMjY1nZ2kiY9dFp9g1ySEem L5aA9x2EQar8Fe3yiwjLko39uGU8Unba4ZQ1kE8pejNAqJqpdbpY
m/44'/0'/0'/0/1    16xFS1oZW6ReE2CPhP59J875wE3PJd6GDG KyjJ9A1k646PoVWu67NcvhjBC6Y7ReofBUjmnBc4cgNXeGp8KWbJ
m/44'/0'/0'/0/2    1HkiKju2QSSWN1vvq6AU8jZ8MdExsZmpjw L2R4CE1jN7gE9Lhmm1oM2ELFjdn4etcxUuQm7YbFEFMxakdPu9rX
m/44'/0'/0'/0/3    1JAHQXAr5UfYpxVTaQnXuLv5xSuNY9CBpx L4Xf2kUs4vB9ayke59FVzCcZWVg8zA3MedR2QZwQHea6ubZi9Rko
m/44'/0'/0'/0/4    1L3aX9Zb4Be6goHdD2ruGaevsfiZ1Z2KeF Kx9YmkB39dGuo7dZXctc1jsjDRxrKcsiz6pWRJAAcNYydM3Z3rqE
m/44'/0'/0'/0/5    1BJ8zZMavS7MXU9hjV6atdHkumLbdvUhS1 KyjR4qQrEQ5siwPW1BoYtzLMhaFve1AqmHF2XTEjpZuH9D4ghFSf
m/44'/0'/0'/0/6    18V5CgntMUSqaDoBTDJg8UFv6vHhrYwK8o L5ZV2cycbbfetQR1HM6fcqkghUKyaKo5YE2PeDHBYg3qfwGMbJAb
m/44'/0'/0'/0/7    19euAnbCCWkfBZ88HrTG9FYY8KtcDBNWXk L4M78czByXYWsuwjcaFsAskfA86S6YgFSMtoRg62YYdQ1SNGPofX
m/44'/0'/0'/0/8    1MzQpwKs2gmAgNQmdJ946Rjj6ooxg5q3oE Kzd7bpaKaTsshfU9Hfn5WGGnEDph3XRfbgPFkkaVgsno2x6ng1ii
m/44'/0'/0'/0/9    1B6myJyMjNJyR3xCYnBeVjB4ENarhhRnYc L3vqyjSkzSiRumoN5D69hhYLCeJ6v1zwfbTGAEEgBnr5gntwURPj

Path(BIP49)        SegWit(nested)                     WIF(Wallet Import Format)
----------------------------------------------------------------------------------------------------------
m/49'/0'/0'/0/0    3GPKjBFRrXnmKLHJtqbiBgXQx9N4UQQ1m3 KzfvJ9kCBswJFEnEqtVwLguq7cWfx84o6qWxiT8HAWASuT1ScHMJ
m/49'/0'/0'/0/1    3MxC93Ny5Y6WQSbPwT3nVz1uU1vohASehb L2moJqSFws3vyeSFnTSeWpGGqXwCMDt3vaH4sFpHCgq7PDnU5upT
m/49'/0'/0'/0/2    38R8oVGe9RdFwjWqyiespp1ZCLQECnULkT KzfHcbSz19fhoUEVDuk1x4FzQwaPHdtoSEhap2CDzQL835x1G7CG
m/49'/0'/0'/0/3    35afr7YCmVLEQGUxCNEit9NnXFtJCAEnPV KyD6FFguKDss5QTMXqFD3tjrVTTKB1vcsPhwSepcRRts7ZkzSWFS
m/49'/0'/0'/0/4    3CAwxE45dhEvYC5gV8zhkETRQEK4SZV5uZ KzRKJPFqqMkV324sLP3XHJa95Zc1wrW1yF5m6pvmx2Dy2ve26qfe
m/49'/0'/0'/0/5    35mcCZ6dLsbcxZsv664ucN3pcQWbvob1c4 L23T8kYXzEffy32R6u8CdbDLph2xqXUkS4tNd4WuxQwGZAefrRAY
m/49'/0'/0'/0/6    3E1DWJt7eR4wVMmQHQur7mNbx2hUCpXNhp L5PyJjHH5sp1ZjBtupZrYHjVAWp6mxa2t4AsXdYr56CT1KtDAuYv
m/49'/0'/0'/0/7    347d6kiUQmyoR1sZ9DdgE43rBpHqGHKGsJ L4Z4yLTuhUWWtFZVEzL2Y8dMjuFGt8nJcDNzBg6m2qhQpMntJCg3
m/49'/0'/0'/0/8    3NkvBbq6wGKLMvjD1jw5R6662SGKA9J2bU L1YLja4mc8wxbqg5zMqPm18JT9RNT6vbkKVs17VAnXKSR6QNRVcp
m/49'/0'/0'/0/9    3GSbNfZZ2pbtzUQ1cdnJFXniB6j9Ds1FLC KzBVtGieeVCav3rvEYD9Rs7QtjN6HAZYQvQTD1rw3pLcZTWkayVZ

Path(BIP84)        SegWit(bech32)                             WIF(Wallet Import Format)
------------------------------------------------------------------------------------------------------------------
m/84'/0'/0'/0/0    bc1q84kvjqueplk6h80x89j4mjujvq3svfwcz8pavd KzaU5N8JLHEv8fwMmfjvxahD556a26QDDsGM8ETKrTUVSnTocS19
m/84'/0'/0'/0/1    bc1qw77nhsl2dsquhnpej0d7dua9wdptm8q8r6xv6e L2mTAeqLhJmmf9hK1MZ43t73cqwc7t1pg5yJHbhdYuiLPajmxLJU
m/84'/0'/0'/0/2    bc1qg85tc9ugn76qa6exgwu3rgxcmmjjy0jrn7d6hw L48Zk1B99S9FyKKcoD7KP8pEJ5WxsTW7d8wvRxVNypwdiZ1wckYf
m/84'/0'/0'/0/3    bc1qnzwg3cnadxqtvkaj420aj2kz506cd8pxh30e0x KwMpidUmtdp536ZA5sH4Tqn9nhmKMs93DUkN4WtQjksVyehi9qP7
m/84'/0'/0'/0/4    bc1qrn99yrfes85zq76rx877kgaw2kc7jeqja4m0f3 L2JSF9vH9oV8mqZPJjmbsswjqxdevhuMKcNyHZB3hfcpe2r79ena
m/84'/0'/0'/0/5    bc1qqqmmm033m2pzpawjy3m96nw7fu76a8l08wxykt KwUnSCNCQCLBymQPLwx62odRfEn4xwSKo9DRa3TgfXjG8Egw4cXq
m/84'/0'/0'/0/6    bc1q00qctf0uvvhd4kr7qsr5s90mwpc9466z6rxpre L4stA9VYKnRC8cEXk61HAFzqPHxSVcaUkPbynFc3wci8BxPf6wkP
m/84'/0'/0'/0/7    bc1qmntu8hq7cg799x42lw4apn63s6w0pjyjy60gep L35w9XduSfYhBDgXDzTKpf4y7f1VWT4uozTTE9xGonWdKa8XiZnN
m/84'/0'/0'/0/8    bc1q8ydc6runthe97n6k7slt5xqeh9dl8enl6kvk0j L1p29HYEgB7f1HGh1vEnneuzne8D7LaEPt5AgnArNnJ2LtiemwUK
m/84'/0'/0'/0/9    bc1qw4jlevwft2yw08g9auga8vtk0zsvzxfmm4jmlc L2p5DvXgJ8xLq6hjVuaWeyZe6pyPRP8hCbW7me6UVHfquNVoK8QP

```

License
-------

This repo is released under the [WTFPL](http://www.wtfpl.net/) – Do What the Fuck You Want to Public License.
