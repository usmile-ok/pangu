package main_test

import (
	`github.com/pangum/pangu`
)

const stringBanner = `example`

func bannerWithString() {
	panic(pangu.New(
		pangu.Name("example"),
		pangu.Banner(stringBanner, pangu.BannerTypeString),
	).Run(newBootstrap))
}
