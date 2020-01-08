keygen
======

Bitcoin Key Generator

A Golang implementation of the BIP32/BIP39/BIP44 for mnemonic seeds and Hierarchical Deterministic (HD) addresses

Can I trust this code?
----------------------

> Don't Trust. Verify.

> We recommend every user of this library audit and verify any underlying code for its validity and suitability.
>
> You can do so by using this tool. https://iancoleman.io/bip39/

Installation
------------

```
$ go get github.com/modood/keygen
```

Usage
-----

```
Usage of keygen:
  -bip39     bool   mnemonic code for generating deterministic keys (default false)
  -pass      string protect bip39 mnemonic with a passphrase (default "")
  -n         int    set number of keys to generate (default 20)
  -compress  bool   generate a compressed public key (default true)
```

Example
-------

**generate compressed public key:**

```txt
$ keygen

Bitcoin address                    WIF(Wallet Import Format)
--------------------------------------------------------------------------------------
139bkuCvFLKMH1KPnBkX1yqpSnChTUbd35 L21qKM4Br31RFZKkEVUUtHyMN6erugcNuRPRdJHcM9Df6jjVkMLm
12BdLfiYhRKUev68ipx8RvmSMtLYN1cdTX L5UTSfMEo2AbQTDCKSiMfCXgap62Db6TrEx6BBpnm1HtykTLpALa
1L66AUcEYKCi4QFgED76dXhvaSP8C6WF5A L5n3yp2owxMS9PnYocgGBFHPXLMHEE13zXEutXnmyxn8uyWxV9aP
1NgsN6ZY7wbvzhw4FNYJ6ZGbvcSmesD7Zb KwvvX9VEqKdzqcwCcynUWDbGViy4Wt1Edo9WQHfCUhQkS9Ew3TyM
1JHzc4nVNxtVk9EaXWhqNvkWQto3rswKYv L53JSKLLpoEBeoTCYNk7x7BptjjAQfahPmZWZkLPMYcKyNKWCaJR
1GJpo9m4JpzPzXGQ7rAfseLPZnRuydAcc3 KyAUfUcS3vfFN6hVos8iKVTvkQB77UYa9cPiqTYLphvJpE3z6d6r
1F6xEVNTjoViq23LgWTZbmrSx4wL7PzsfP KxW5oio92af2yS5v75oAaLddAG3K2f3cxWv9M2HpNMys8E8qSKmj
1A1VMCos1G2UCeH7gYcAvzar8dgQrMRKCH KzECAAJwHuzRRFpCoSNfbqnyAJ5trAn6UaSAuvz6yZmMEZVfPRzA
1PRZKXwsA3M7CAJDNTzykhBELiaTYYP2g2 L4UdwgQKBedz7PoEPHoEYGQzF63nAkZSZH1jGezYpHa9C1sHZW2Q
18u9cCB8uW7MPxL3CFHr4GXTH6HXXCVDiJ KxRZqT4s33AFNhLxC8uWvWWsuYEqer1hDV1dAnYUFiAhpHqWJQ1P
186oB9RizJvsiRMPzr4c68wAPWj1U8bScc L35fskpp8X42JwRnxpaVCFYEy9hgQhCYmNmg5hvtuQhwmAffJPPt
15E7f36R1dHrbcRqnQDAjARSndHCSLFzPh KyV72CibfF1oYJWKJzejp3LdQNNbpMtGD4few93a2yry9WiX495u
16MW9gxtmwbAKjHcycYeXUZaQ3DkDbPbZw L2psGNvjHm5kPDw5uq1DrtPgFULQUWYuKwrFqykJxrtGS33RMG5S
1N2Mdg41aA6vQX1pg82rgT7id9JiBftoqq L4PYeDBKkcisEjCsh5FwW2HEXmiRoXY6nDhnNYYDCQSMjTHQAvQU
1JiRHz4onYwhZPZBWDHfJD4T5T3XByLLbZ L2ZJALPds1YLLQVf6rbfEBHAqwQcQdrHZzkcKzmEjSmQqMWrxqP2
1RsemWFv5ThYq6VW4CStLs2GXBqZpLHyN  L1RCnDZwEUidGMC2thUvejC3ChGGx3JRPmvHivHrWqUQXyQ8bsgx
13FQsUVKQQyzLtYiiR2J1U9upyqivk4Fuz KxBGyZxRCrkCmCB5YFLRhqGs6eFDRpDR1w3MYaEyqQ3vbXKihqm8
1MJQDqB1dHdPFxcHeRuPbn4CE8WY3XNGww L1TZNRd3S5pEmKThUABGkmURjVkXVHdxUmjRaYzCXvhDJ6vay8KK
1LWTkwaYgrpDgwfNT9zafJBtgiXKkuotc8 KwxcE9z1ieCE37b8UZCKph924jvYvcrFJNt4QBQe3WSjh3kdfH7h
1FykDsmzqogch2dvXKXkrFSuT6iHapCYb2 L4QD2RzjYM1snVW54F1SnbHUormHQvNtD3GDHB59PGkn2Lv6tjmN

```

**generate uncompressed public key:**

```txt
$ keygen -compress=false -n 5

Bitcoin address                    WIF(Wallet Import Format)
--------------------------------------------------------------------------------------
1C2on8GDgZjMHz4hpUUTccdbwiHmDm82qp 5KMxQBPLyfLLpAgnxnDdDgtHF323jEaqCyM7u9TgRxJvkTPSoM1
1JymTfXZQWQqPZV7GENiPb6rAWrnYg9B8T 5Hqwm9jYdLgMZawGxFW3kdbUmu1hWV8Krjnc33vRPj1T8Zz2eVZ
1KZKt4N1QpcR9qnyWnZvAe9ccK9SWzv2AT 5KXZFiE3Rz5s1zYym8voutZs4a4NBkz3pPVQfJhrSCeCVZfgJDU
1C91FBVHU5t3N4oUFe7XDCVwjbcRKvSSSe 5JwKt754DN8SjustfVuuHo9wFzDWSZDirSwGbiXC3hUnopbiw4J
1DN7rfvAYB77acrk91SmQHGWtjLh23uG5K 5K7RMVjdWRaE6PwgvGYW1jE8ydLLjnzYWFx3jBGEWekUD8pcZmC

```

**generate bip39 mnemonic with passphrase:**

```txt
$ keygen -bip39 -pass=123456 -n 20

BIP39 Mnemonic:                     arctic weird hurry attract morning mention island laundry gasp lecture urban season
BIP39 Passphrase:                   123456
BIP39 Seed:                         fe69ed69eca284eaf0146a7aafedf99986cba2f8805fcf8e538bd342683086de830fecda7dde9140c48737c97467fd87df8f5c16a6f90b8af82d51cdb3180e8d
BIP32 Root Key:                     m                xprv9s21ZrQH143K4EUqf9hLagzxnpEkGJGxFD7q7ANbMYrYrzk8gsJDGCDiVwxbiBBAiJHNN8idtUG6YTmefdJaat9WwpZG7xfbqY8VUWAipmQ
BIP44 Account Extended Private Key: m/44'/0'/0'      xprv9yN4thyP1Py9uhiTh4Z4h4q3ed4wPr6boSugrCwyQSd2D1wCc7LFepjUx2c6JZW2WZYwKZA42vVQRQAQy9W2TpuqsEr3nWrVzaWsNYVyait
BIP44 Account Extended Public  Key:                  xpub6CMRJDWGqmXT8Bnvo6654CmnCeuRoJpTAfqHebMaxnA15pGM9eeWCd3xoHbuBncbfxiQ1ueRBU8hHmxg3AUvJ9V4HsKfvS52HZnfi5H8g9N
BIP32 Extended Private Key:         m/44'/0'/0'/0    xprv9zzuDNBQd7f9XAVDx7kZYdKFyQ2KYxLGYyhdFcRqZSqWNA9GqZayGSVtFj4PPGFqMyUyyP5JCYc3neuuk2y9xM87fbKseAcVr3GXz6ooCrf
BIP32 Extended Public  Key:                          xpub6DzFcsiJTVDSjeZh49HZumFzXRroxR47vCdE3zqT7nNVExURP6uDpEpN73jgZxSoqiEJ2t6QwgRouR7UKzzAosBwxzCBkC2AqwqQqf6V42J

Path               Bitcoin address                    WIF(Wallet Import Format)
----------------------------------------------------------------------------------------------------------
m/44'/0'/0'/0/0    1NjXHEGvoCddZqAX4gKkNpRR2vBBLSDNyR L3b8Jogy6Q3347mKNkiF8ccDTvjC6w1GQaZx386sfNQdYPUbQmW5
m/44'/0'/0'/0/1    1AxENeqLKZ7fnprCXFWo2n9F1X8b6qahX4 KyWyJApkpoa1nzkTDKWJ6dPHVBBub7xb6cwH7eeq4keJNzrsokpv
m/44'/0'/0'/0/2    1B5EXNUsKks1kdN2ojRh8ysJhBA2v4ZYrM KywK8FGqvjG7v7Xyo96Bw1s6LA9qUxTEafXmjWpgMKEf23mtC3nx
m/44'/0'/0'/0/3    19xa4Ehkk1SeDRzi5Q2fv2V9x8s5oEVmeD Kx9buvsAkMmdTK3vj85Twe4n7JtxurvyfcawwR9EaYVvFrjMQrzn
m/44'/0'/0'/0/4    1DXTaFQevfgh2TLnPawCMi9qXbek47WjpY L4oXfAWAzMSYWe5fz5hkVq1C9N7K9nE5rey3LR97ze64rJaggWpq
m/44'/0'/0'/0/5    1DLygBKy4msVcNmRtX77FCczjCNiVhmeka L2qSH5yG5fyTEyjpTkQY1SpP2ZH1hA6oFEG6r62adVj62tsoeEps
m/44'/0'/0'/0/6    1DvvtZuiV1pLWy8up4Bek9vo1BmW9NQdZE KyGZhM8VV1ee5J9yxV2LfGtzbEGTRMwEtq1z1FqFAC1E9oToyTjq
m/44'/0'/0'/0/7    1AxDkG3pB4BVvvFdjr4evfPEk7W6pFE9si KzUAheye2X8ZZSnXonYtRhgRYYDmKdVXujH3U66o1XL5GJgzDvFJ
m/44'/0'/0'/0/8    13ioixpQZktK8545LdDx9Mnyhfmnf97EFd KwGLLqCr8PGQTRBmY4eB9WioWwCcwLZTyvVjgYCHuQFuvxN28Xnd
m/44'/0'/0'/0/9    1DAZb5GoCfweEhfKrPvQhpCrw4tRzR6ncb L3muRZdupCEToHEeiJggJxH8KXnqmqsnVEtLHQkeSvfv3WHWFGJW
m/44'/0'/0'/0/10   1CvtHeSKJRpAySKjVpXFjwkjoW9WBRSUys L4GaeuCrroeRtL7aNXMopseT8HVohvnFtzViuddeVffSv2LrZxAL
m/44'/0'/0'/0/11   16D9K61PJEcoQi87BNCkV6MvGCgRbFBmmy Ky8k3vBntawmZdzzdsL94TsCCmpnzVcvMGPg2HmvtL93C3im45RZ
m/44'/0'/0'/0/12   1CVtw63QxLteXUh9HYFu7ZwXy9yfhDD2ud L49yfGJgsVL1UZEmR95JpHzwM7rEwUAbasE5dzNGvoyFStyxEUR8
m/44'/0'/0'/0/13   1CUhG3eLmaiMDhKTJEnQsyi5mosKiN4HJS L1MGutW83uAdxR6mHv5gQzLZpMEkg544pENL3PVRsLcmq3kjZxYJ
m/44'/0'/0'/0/14   1ASfwbiqz8VAhPbT3oyfKPe1Bq9phdstby KyoPSvTCZ9a7DuXpJysM681wQNZRgMLGwm4sTFG9u9uNovaGJw98
m/44'/0'/0'/0/15   1AkJsCgdibnYogKzFDYLbEqYssrMi9iwfG Kyevv4ty2KFu7dCfYXoKKtoJaZngsa3vY1tf1jK1S6c1n93Yp33e
m/44'/0'/0'/0/16   1R53m5qXCNwFLhLdZbDH8NAe7sw878Nqi  L3371CF2p3RiZN4SJJR9YbogvrRn9webuk8uniTbMTHj1URX52A2
m/44'/0'/0'/0/17   1Lj8Y2ruA8EevVWhKzP7cQPgm7EtGefXSH L5UfKHAemFQLeerJpf9oZy3ybChSiYeTH4bJku4FHGPXeV7XKC81
m/44'/0'/0'/0/18   17KLNxs3SoJzxPkhMrEFVBKSMNNJrzLNN7 KxyPLEtUJWkb6njR3bc4HB9EucbgyhPPDTFD3hb3gFRD8Z9UM19q
m/44'/0'/0'/0/19   1FS7ddyAe1gBKJpk11qNjRdhGQB6RsJNUz L4Nu5DGEEh48aVRLfF8wSjnKjERQ3UyLaUXn7mdmDdqqGKBwvJYf
```

License
-------

This repo is released under the [WTFPL](http://www.wtfpl.net/) – Do What the Fuck You Want to Public License.
