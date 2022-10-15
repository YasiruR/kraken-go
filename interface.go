package main

type PriceFeeder interface {
	AssetInfo(pair string)
	Subscribe(pair string)
	Feed(pair string)
}
