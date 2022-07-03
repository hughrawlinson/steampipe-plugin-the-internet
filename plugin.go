package main

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(cts context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-the-internet",
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"theinternet": theInternet(),
		},
	}
	return p
}