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
*   BIP86 - Derivation scheme for Pay-to-Taproot (P2TR) based accounts
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
$ go install github.com/modood/btckeygen@latest
```

Usage
-----

```
Usage of btckeygen:
  -bip39     bool   mnemonic code for generating deterministic keys (default false)
  -pass      string protect bip39 mnemonic with a passphrase (default "")
  -n         int    set number of keys to generate (default 20)
  -mnemonic  string optional list of words to re-generate a root key
```

Example
-------

**generate compressed public key:**

```txt
$ btckeygen

Bitcoin Address                    WIF(Wallet Import Format)                            SegWit(bech32)                             SegWit(nested)                     Taproot(bech32m)
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
1BfzdS1VvhXHKjgaip4yc883BrALD9F3ko L1r9CqkKjLMHZyKt2ZcDuZdeRE7ByhXxx5n3D12tyULotfL4wPF8 bc1qw5gjkllsm0wt0ktxmszxkmxwgld7zh5k2k07dn 3MmWN4z2z9eQ1bUs4AznqURM8j87o5zDrC bc1pqczjkxu0866lgllm7mzvh02a6c3uj5jd83jm5w7j0nghlq6ggpmq8sgf59
1DtSnQBtZ3WcT6P6FXbnAGrXyUmQvZepTZ L55zXW3fjz4wS1As1LFFysKxky7d3bEgiooWTt7vriGSMCo8XUP7 bc1q34wramw4m8rdrjnyv4afmj9dmyh4ucd9kh75ja 3HGBQT7nQ5rpxfTKiHjGnMQdiGxYwQdQnM bc1phusg4j78ln2xvsquqmsyq4lykvhmy9jxjcfqs5e9sf9p7dlnk86sjqlg72
1Aa3LktrJtY6i6geAkM2C5p5T3AjzvtxVF KxPFpvRVuEQtjc6wiR7dAbhqXzyeCVcn2XUvhqHPdF8ugJ32Fnvf bc1qdrudqz3uh2gq63vlmg48elnfajy35l2kt7zt68 3Fxayd1WVcNQjAVTmopQshFdzSH939eLxe bc1pvn7ywglyvfrukl03t7u0kkwfwkq5wyh0q4udg60gw398j2zylshqwusaat
1P5BVoSHCwVvNZFmEzkuNEMTeWwqHvJSHo L5iuKs2HVw8PhjpFLtFhACxb5517jcQknDeCFEVMbSti2j1EXEhC bc1q7gwnzapjnqefstq5mekxhm9myfe7pue6jrnqxx 3L1CDyUnH7aBTpTgc6poqYdskWLwqDAVnm bc1pq9p7rejtaax8fqukge9deg83a32lzd4znqsp3k2ks3yfpcjmm5nsvwlkh3
1PRCLwwR1SKEV9sjjK42ZbKkLezMt7qeWp L5jPHzyp59TSZXfe6CqqV1cmnvBDBRz9VcM8C1E29saji2otfYtC bc1q7hnrssdp93cuw8tpf3cceax7uj3jv064sweslk 3JYK7Anostwh8QQiLP2sLPbdNN3zh16be6 bc1p4rk7vdztzzkkjle0lvryzgfv6dprl064hnrqlevlkr43c6urwl4snjav7l
142gJ6os6XKNDJ5NCaUaMqhFXqnwebPSGV L5bYP893gREoebEfRoogCj4UMZ5Mnhpzh4NgujqUu9qAGZp9bfKP bc1qyyues585e4kfrz2prcye2kttcd7wywaf06cfsq 33ywn3t2WKzHsEdhidb8HZfpc5v7sWv5pj bc1p8r3skhedk9xeemnrv9ezjnzd8rw686mc2h8yml38hm75t3wjnlnqez3z46
13ifkDRnACfzwuHrRz4fY6456fQXg78Qda L1QPU2PVqZ1wKH2agv3iuq9ihZa8r53imbkWugeofCW6ostQN8NP bc1qrhg6v9g6985uz5qd0yw7yal889yp3vxajmmlmp 3P74HYFdaBNMPVegmsgUPJeffNJTvM7wix bc1phchzcmvkcytkr0r58k7zh5fhq35n4e029t3phhrsr7wf64p043sqqptrpx
1HXodg1CjQ9x1WL8NqyRV7nRYnLNaQgcEH L3oWpekUjGAXE3qTgzYKUFBsVqaX27vuRbSDnPkdhNWa2XNNytjE bc1qk42kkpjxvqzz0gnwmd86lfysmvv0g5q64kj83k 3EUGNwTZThUdHpSZXyBkBtBt6jh7GKC774 bc1pkj5s906ene39kz7c8p63ra8mw7kyrc2vhnx8v2p0hv8qalc8dzuq3hf842
1CM5ShXqqyCNW9t8SW3obR4jMXQrLeAWRZ Kx2zYwXtF2HmtiJuRgabhmJraz5RomKPAdE1KnUjRgy7i8Ew7gDr bc1q036kj6xae7rn8lkzf6k2uph0tzaqzfzqtmdktt 36ujhq41LhoTqpDh1Jp6pAgCNigNRpqtcV bc1p4ccxz68q4h3ad7s2554n44t6f7272uwdlw9mw7w8mnrf3r925nqq03nlq4
1Gmj7P3W37jj6Y5nytuNsD1orqHVCGKXZD KyujfST4b9AWPiyP1RLjb1rD6At9BbQp4Te3AGyy1wKYGJ6gmEhd bc1q4nl4wwjkapt3aa9dtgu0uu5976wxld4s4z0rqx 3CUhcCkzfX1bVgafHXVQKAQKGgmsQd2WfM bc1prumcfp2fm0jtpez2592rdlglceres4dw58fa6dudunkhnhy4hcfsfnee39

```

**generate bip39 mnemonic with passphrase:**

```txt
$ btckeygen -bip39 -pass=123456 -n 10

BIP39 Mnemonic:    ordinary volcano company hedgehog usage success awkward filter state energy wool point
BIP39 Passphrase:  123456
BIP39 Seed:        a30ee132965b779387ed58c09a6ecf74d2f1776bcee37ee5193f28b6747c94f31edf3514faf3b3c63fc788a0f33d97bc2f2fa8564f183d5168829d4af545b2fa
BIP32 Root Key:    xprv9s21ZrQH143K3hyFUrgynTfADUoUk8ERhWnVCpRX9KZSWJxfR87dwfuaYFMUC2N3HTQsPXA7hbEbGZSUzNok7siN5EVFvzStzoJmXNvKJJK

Path(BIP44)        Bitcoin Address                    WIF(Wallet Import Format)                           
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

Path(BIP49)        SegWit(nested)                     WIF(Wallet Import Format)
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

Path(BIP84)        SegWit(bech32)                             WIF(Wallet Import Format)
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


Path(BIP86)        Taproot(bech32m)                                               WIF(Wallet Import Format)
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
m/86'/0'/0'/0/9    bc1paj39mjp28znlg5mttzss93e0lxx8tnf56lfc0834p0zq3ujaw2hq7zcmkr Kx45eoQxLmjFsCRuYcAe5crTcdiNgcWiqDTTFaiWUmrqgZ8UH64

```

**re-generate a root key with existing mnemonic:**

```txt
$ btckeygen -bip39 -pass=123456 -mnemonic "voyage blind unit shoulder yellow attitude mule all hire above obvious swap"
```

License
-------

This repo is released under the [WTFPL](http://www.wtfpl.net/) – Do What the Fuck You Want to Public License.
