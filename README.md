hdkeygen <img src="https://www.buybitcoinworldwide.com/img/segwit.png" width="100">
========

[![license](https://img.shields.io/badge/license-WTFPL%20--%20Do%20What%20the%20Fuck%20You%20Want%20to%20Public%20License-green.svg)](https://github.com/modood/hdkeygen/blob/master/LICENSE)

A very simple and easy to use bitcoin(btc)/ethereum(eth) key/wallet generator.

*   BIP32 - Hierarchical Deterministic Wallets
*   BIP39 - Mnemonic code for generating deterministic keys
*   BIP43 - Purpose Field for Deterministic Wallets
*   BIP44 - Multi-Account Hierarchy for Deterministic Wallets
*   BIP49 - Derivation scheme for P2WPKH-nested-in-P2SH based accounts
*   BIP84 - Derivation scheme for P2WPKH based accounts
*   BIP86 - Derivation scheme for Pay-to-Taproot (P2TR) based accounts
*   BIP173 - Base32 address format for native v0-16 witness outputs
*   SLIP44 - Registered coin types for BIP-0044
*   EIP55 - Ethereum Mixed-case checksum address encoding


Can I trust this code?
----------------------

> Don't Trust. Verify.

> We recommend every user of this library audit and verify any underlying code for its validity and suitability.

Installation
------------

```
$ go install github.com/modood/hdkeygen@latest
```

Running in Docker
-----------------
For security reasons it's recommended to start a Docker container with disabled networking:
```
$ docker run --network none --rm -it modood/hdkeygen
```

Usage
-----

```
Usage of hdkeygen:

  -mnemonic string  optional list of words to re-generate a root key
  -pass     string  protect bip39 mnemonic with a passphrase
  -n        int     set number of keys to generate (default 10)

  -wif      string  decode the private key from wif, then generate the bitcoin address.
```

Example
-------

**generate with existing wif:**

```txt
$ hdkeygen -wif KyBmvuMWYT9aPRf8tT3oBJsmjwXpgWbFLwbUDhezAx9EQX1PpJXN

 Wallet Import Format:
 *   WIF(compressed):         KyBmvuMWYT9aPRf8tT3oBJsmjwXpgWbFLwbUDhezAx9EQX1PpJXN
 *   WIF(uncompressed):       5JG8TiKEwhpFpDqWt5YSsNGahYPPasf7sjjFKjk3XoUhCEUvBj4

 Public Addresses:
 *   Legacy(compresed):       138dtSCfryzZQ2zkEDLC43KjtmBMk9Diw1
 *   Legacy(uncompressed):    1CbBvFZbdfCzfPHeYGBneGE8jr6FigocnY
 *   SegWit(nested):          3KiGDUjkstR1zxcnfPcSNoGHoLFFKBCmpe
 *   SegWit(bech32):          bc1qzaslgufxmwrrj9f224h3rztyjn5jusmy0hkk8j
 *   Taproot(bech32m):        bc1pzf49m706xk5xnmkqj368tvmf48aav8jqg3mhr5ujdp79u8jcvx4s7scxzc
```

**generate bip39 mnemonic with passphrase:**

```txt
$ hdkeygen -pass=123456 -n 10

BIP39 Mnemonic:    ordinary volcano company hedgehog usage success awkward filter state energy wool point
BIP39 Passphrase:  123456
BIP39 Seed:        a30ee132965b779387ed58c09a6ecf74d2f1776bcee37ee5193f28b6747c94f31edf3514faf3b3c63fc788a0f33d97bc2f2fa8564f183d5168829d4af545b2fa
BIP32 Root Key:    xprv9s21ZrQH143K3hyFUrgynTfADUoUk8ERhWnVCpRX9KZSWJxfR87dwfuaYFMUC2N3HTQsPXA7hbEbGZSUzNok7siN5EVFvzStzoJmXNvKJJK

Path(BIP44)        Legacy(P2PKH, compresed)           WIF(Wallet Import Format)
----------------------------------------------------------------------------------------------------------
m/44'/0'/0'/0/0    138dtSCfryzZQ2zkEDLC43KjtmBMk9Diw1 KyBmvuMWYT9aPRf8tT3oBJsmjwXpgWbFLwbUDhezAx9EQX1PpJXN
m/44'/0'/0'/0/1    1LQs9YUK5f1tZbHd4BgZhwBrAkjiEijRBk L462k718zfnSErmr5N2z6nD5HiczeN5VRnpsbz7e3bKWCTCRAXrX
m/44'/0'/0'/0/2    1HAGmsNM9YHL9HaEf52C2ZnMxV3hpGhXcK KxdPiFrN8ugsTKvZmWsmEupmfTtZgHV7R7SAhNNiiEjdM8JYDvh3
m/44'/0'/0'/0/3    1LQmutrWhvsSanZobZbHorTuuqYUtmDLEc L49LJ2gnhrkJUxwJ8tzh6ZqL18Jf5vbU6Q9WwhcQL1wdGRuaiYvR
m/44'/0'/0'/0/4    156XqUkKvQtZ9bsKzosubmTbHKmXw5gWx4 L3p2C3rsakMP5TcUS5BMpxXdLu8fopCieKc6sdchundgg6VvRXan
m/44'/0'/0'/0/5    1Jk2nLWL4jFHgwEd8DN6JwVd2qCkLYNwJH L4bKArmoRuQVsogjpQMtS1fZQ1CJ3VH5GPMTJncYyXksukN3wm7A
m/44'/0'/0'/0/6    1KAT5dBLjR4YX5AZGDcUHpPY5L2cQfXbkY L23n4GcnSbNpxKy8Jq9iE3ec7xZgJ6o9RyRxivEwyE7ZwjX7Dfod
m/44'/0'/0'/0/7    1Hy4VqNsNcCjcn8nF5br9bD2kNBnHKszEh L2mMTYVZrAivCsGMuSxQAfdauXvbhfJLdFyC44MF5mTuat6mh5pn
m/44'/0'/0'/0/8    1J67UZKPGU429x85vLtN3bHfkTzk7U5n4T L3Fxn161sBpLrDJk6XUmmhYdJFdjwezZG4HKmWfqf5CMiwBCgMex
m/44'/0'/0'/0/9    1Er5UU9Dcs6HRoNXF2koHGG2E1vzKGtNCy Kz3fxwbdKNAStDv1XHQ5oHsZpY1BV7zGoMERM4yVfXd3aKsEWFHq

Path(BIP49)        SegWit(P2WPKH-nested-in-P2SH)      WIF(Wallet Import Format)
----------------------------------------------------------------------------------------------------------
m/49'/0'/0'/0/0    3CPT584CJX8FhwpznbbkNSxpp9MeLGmzP9 L1d2TNsQp7X1FfPuYVLVcrJC5YRZeyPVF9zdy8rk674mq3Ak5Cki
m/49'/0'/0'/0/1    3EAd2eiSUEDW9Edg8VWLzVCn3XGJRiRzJg KzQjLHtrkS6j6CUTucEY4NhkFwykC6mQCfM4BNmxajg6vx4SyAje
m/49'/0'/0'/0/2    3FtSgvTDEGV56krUDeaZUmEcJ6FmECg11A L2k72EpVgYH1NjvFUUWSa5F9ckvQXPpVwkeCqn7Q95Pg9jemDG14
m/49'/0'/0'/0/3    32D62fEJ6ML7dWm2LV3LW84wBqbxyWJRtt Kyy3HgT5BeEtNtLoaUXDSxE1sBfX5nP5usmxNnPs2TpMG45mWhaf
m/49'/0'/0'/0/4    3QZTnR3fVQuZCpPCntQ8ezxAAq7aRGupwM L2Qgp3jHLaVx9xeznG45trwxWkGRSd3KdRKEg9JzLwCjJ2eUing8
m/49'/0'/0'/0/5    3HrPi8oDU19RyWJxcWUft2inSWxPDo1yq5 KxhKzV5fdsYjWzsdbuCWCaRsaXcNC4xzVbuu3q8tqyvWZ1o3d4xX
m/49'/0'/0'/0/6    3ETz3XcLnj6dUBrckSy9cgPet5Qprvjsxv L4miaKHCC6P2v8T7mSXnePdem8WVZzCM2mpbcpry66sjNnRb7HTJ
m/49'/0'/0'/0/7    36sX8P8oM8cQWYBgyBZRkYNrGhKcxPrmB6 L22djdrLE9TTc8yg2qUqEc2tYt1LFEFubDXGMehq6nqyDzgzh9T6
m/49'/0'/0'/0/8    39tf3FX3vWLkY1xnPtJajpnwNGu7jVgSzp L3t9TzAgaD4kUg9efRm1a467Yuim9H4AD1FbN9zf7UyRhy9yZD9e
m/49'/0'/0'/0/9    3PUGGsnRsHABw3HUDy3z4vztMmhQ4WeeNG L2i4svg2qsoGwaQR3HZNBnbB7pQwZGsLKXYHhCgxroLx6YTKefHg

Path(BIP84)        SegWit(P2WPKH, bech32)                     WIF(Wallet Import Format)
------------------------------------------------------------------------------------------------------------------
m/84'/0'/0'/0/0    bc1qadjnz3r4cfun9vz40gn2npzc65mm74q8z2w2y2 L1WM3pNwupPP1pJ66jEN9fWZLgRwr9v2woSxxLPg18GKmAJnSDQ5
m/84'/0'/0'/0/1    bc1ql6dp9ck9gukygda7yuhkhdkqd7p4t2jjy84njv KwMQo6GqbWq1yCxgMbENkBu7p4P3prqueSZydrSRbFsA63mcxQUJ
m/84'/0'/0'/0/2    bc1q388gwptytt92nhawkcg5gyrg8eswt7r78jvcqy L2d3XxZz8dbeGNcHJG3VpzhtvV5s1CxkGKGcL1o85hcjkvMpiuUT
m/84'/0'/0'/0/3    bc1q45z20nk9v6662rgew93j2ay94hxtlm977xu0md L43W9Q5P9A3fGZdV3vzA7jjuGYB6auh68vLp3BzEzARtR9sxsKx1
m/84'/0'/0'/0/4    bc1q48f3vxu20kap94xfaawf06c4y8fmk6gkqngwxk KybUQTHkK8civ3bfKSSJwamzJTmnazJJfZPBK8fGD2mhCZ6zYCf1
m/84'/0'/0'/0/5    bc1qkfpc8p93lxnqremv4cwe5tg5zqhff2c60qsjdm L2g65k55yiDD35SdXhaYdE8inZgvRkr4MLMebWqusfB6gTS89FCX
m/84'/0'/0'/0/6    bc1qcjnn60h63ps5rwm3x2qeymz6nqcxg7qx5ppl5z KyCMX5tQDHCsibDhQ2j9NLaqcoD2Dsn6ozx6BCjsypMpVVZoxMDB
m/84'/0'/0'/0/7    bc1qjjejva2su9ez7r52qe69pz2fhy9dgc9vmq0we6 L42dCmEDMo7c1G5VAAfM1uJD51c43AgL6bn52npet51nRPhW8vwB
m/84'/0'/0'/0/8    bc1q9llvncwl2mwh74nx0gtecw204lwlvr56arv570 Kxt2sBoe1Y7sbvaehBoDDpBcDAMEQ9xcQNPpXGqJ4SydQG47QSZc
m/84'/0'/0'/0/9    bc1qqzy44jc24akv20euhvedwlw674r2rwe5ylrpnx L4qUmctrk6z6vbwH9CKPkRqu3SqRdeSySw9DchbADaHo3PcYuCHL

Path(BIP86)        Taproot(P2TR, bech32m)                                         WIF(Wallet Import Format)
--------------------------------------------------------------------------------------------------------------------------------------
m/86'/0'/0'/0/0    bc1pwp0ypr3cn09sfn4en7vmcep4857dawftz9t7ays326v3yrwnk06sfwew84 Kwtz3TpxuuWhJ7dqm4QBLNnY1qxcD9DodXZsyAoL4FCFUGRfdg6R
m/86'/0'/0'/0/1    bc1psttavnhu382yq24cdac5j9qruf2cvkxhgx8rxkvf2jgpt3pduuzspa2n2k Kzhyp2W2Nd13wkJbNRkd7FEEJ5rDALkq1HPfLgSfz53mRmZsBgNX
m/86'/0'/0'/0/2    bc1pltlwnzste88yjmrzsx7y6ls5ggggharflll8uku3786glvz73hnq0ps336 KzYxEXhja6R1QkLx43Zir5tfpdK1sLzvMguZELapLb5cUoqpEYGL
m/86'/0'/0'/0/3    bc1pwwjzkvdg5hshxtje0gumet6rn3pxlc6unxenlshumqhmdede0mfqk7gryh L57sF9zkbAKbztpkLYxKBZREqrdU6zStzatANekR4EpXpJbYXipV
m/86'/0'/0'/0/4    bc1p0kv69cmelkcdsgz9mpg7nzrghupp6lm9da8uarfdtlradgsufpxsf7vwdl L3DuXn97eQt7CfRxNc1NHEkieBkMbdbWhnKHEGJHDRqrPkehH8WM
m/86'/0'/0'/0/5    bc1pcqjufjs3l29gpx8lkmzrlfw3vzjxpavnpra2z5k3sh027d8azhuqhxakks L4wm3MBugbGphX8WX5Ta71uNzvsQprLGg6DNZ7werGVKhnGtVg65
m/86'/0'/0'/0/6    bc1pzd0umy9ynhx0dvl93udvx7u3kwad04az9zgmf8hn26ey2aj7zngqhds47l L54RbQi8C5eQdLWL5VBB8KqV5tCXNXJjrzvgUxdZBFgW4TeDLKxQ
m/86'/0'/0'/0/7    bc1p52dykhfv7tdkmnset7hkzr5uc2g72vntv0tcgvggf25ddlkweu6s5e7299 L1LG9HoB6Z1a23nqdYsWLk4oxzh5U49StqN3STXf4cjX3tWvCeGu
m/86'/0'/0'/0/8    bc1pr34ruah5pq6s5l0w6ntueazesaqrqeugkrqwexe9qd325kyma2hsgw3knw Kxf3oR9dquE7PNUaVbbn4yA7E1mrmFzY8A9Md6pNrcdfS8UjW52K
m/86'/0'/0'/0/9    bc1paj39mjp28znlg5mttzss93e0lxx8tnf56lfc0834p0zq3ujaw2hq7zcmkr Kx45eoQxLmjFsCRuYcAe5crTcdiNgcWiqDTTFaiWUmrqgZ8UH64a

Path(BIP44)        Ethereum(EIP55)                            Private Key(hex)
------------------------------------------------------------------------------------------------------------------------------
m/44'/60'/0'/0/0   0x84176087d47dFfc816dC22AEd0B73c1a0858A9a9 e9390301c5fc4d18569d6e55176a0759ca2c4bde0037b2d4edb46d85bc819efa
m/44'/60'/0'/0/1   0x40F5Ea577985c169BE62F30E21448873782801Ce 96638b915ac0d867d557465165369c2908e7c5b45ec64ef23d075564ddddd38c
m/44'/60'/0'/0/2   0xdCE2C1a45F1b9b39f4b2B18276B6FDCb1252Cd14 c1c276e77478b38327459cc64c749130c50ba4bd971c0bde4a2ed83564858650
m/44'/60'/0'/0/3   0xe1314B362eec5A1E9A0AdFc7e6085A9eE5dcB478 4e802e15994d0e76073c94ae4b692e203a799017e3a71d6d7a514ac243945b87
m/44'/60'/0'/0/4   0x5d56F4081D1c3d44b80A98F5C1E48a3964B1F065 7af9e145e789969d9231c3f0134feb0573ada06a78a98a054cb590ef752a2ef9
m/44'/60'/0'/0/5   0xA1F5aeb76Dd50d8ed41F8Cfa8FA89A2F75691b72 87c11816f5a4348eef022f382fd2febdd148dc06cd766dffd96914363b6fc6a1
m/44'/60'/0'/0/6   0x2Cf9477B0AAE3C1CE85c034616B7cB3702576E45 d0f687259719372d0d6d6eaf7e5d1f839279000ac7de7bddaaf37fcf959e656f
m/44'/60'/0'/0/7   0xab1e882A3E07192D45595fEDe116F5ce7AE73210 ee6a8125095536b00b219007a562898f49d0708c8d4cb43926cb613773e90f2a
m/44'/60'/0'/0/8   0x4cB8B7984818b5deA7c3dD638D9931145af26e92 23f8d138ec530626bfe553775c4db9dfbfa47cd27f9ccef051f399e7760c05c9
m/44'/60'/0'/0/9   0xd997f7076FE26510Be4966F5cb42bcf5c1aD83Df e41c4d7ea125fa28be536687b70ea58d4d85338fd773e46edc08f90e8fdcd2a5
```

**re-generate a root key with existing mnemonic:**

```txt
$ hdkeygen -mnemonic "ordinary volcano company hedgehog usage success awkward filter state energy wool point" -pass=123456
```

License
-------

This repo is released under the [WTFPL](http://www.wtfpl.net/) – Do What the Fuck You Want to Public License.
