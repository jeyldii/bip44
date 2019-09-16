You can read more about BIP44 here:
https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki

```
type Constants []constant

// Create is starting point. You get constants.
Create() Constants

// If arguments are not passed, all constants will be returned, you can also pass coin-symbol and recieve slice constant as a result.
func(c Constants) Get(coinSymbol ...string) []constant
```

```
Format: constant, coinSymbol, coinName 
ie. [{"2147483708" "ETH" "Ether"}] 

Usage:

consts := bip44.Create() 
eth := c.Get("ETH") // [{"2147483708" "ETH" "Ether"}]
ethBtc := c.Get("ETH", "BTC") // [{"2147483648" "BTC" "Bitcoin"} {"2147483708" "ETH" "Ether"}]

Also you can get all constants:
all := c.Get() 
// [{"2147483648" "BTC" "Bitcoin"} {"2147483649"  Testnet (all coins)} {"2147483650" "LTC" "Litecoin"} {"2147483651" "DOGE" "Dogecoin"} ...]
```

LICENSE MIT