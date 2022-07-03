package main

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func theInternet() *plugin.Table {
	return &plugin.Table{
		Name: "theinternet",
		Description: "Make HTTP Requests to the internet by querying this table",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"url"}),
			Hydrate: listHttpResource,
		},
		// HydrateDependencies: []plugin.HydrateDependencies{
		// 	{
		// 		Func: hydrate2,
		// 		Depends: []plugin.HydrateFunc{hydrate1},
		// 	},
		// },
		// Get: &plugin.GetConfig{
		// 	// KeyColumns: plugin.AllColumns([]string{"url", "request_headers","request_timestamp", "method"}),
		// 	// KeyColumns: plugin.AllColumns([]string{"url"}),
		// 	KeyColumns: plugin.SingleColumn("url"),
		// 	Hydrate: getHttpResource,
		// },
		Columns: []*plugin.Column{
		  {Name: "body", Type: proto.ColumnType_STRING, Transform: transform.FromField("body").NullIfZero()},
			{Name: "status_code", Type: proto.ColumnType_INT, Transform: transform.FromField("statusCode").NullIfZero()},
			// {Name: "response_headers", Type: proto.ColumnType_JSON},
			// {Name: "request_headers", Type: proto.ColumnType_JSON},
			// {Name: "request_timestamp", Type: proto.ColumnType_TIMESTAMP},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromQual("url")},
			// {Name: "method", Type: proto.ColumnType_STRING},
		},
	}
}

// func getUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
//     conn, err := connect(ctx)
//     if err != nil {
//         return nil, err
//     }
//     quals := d.KeyColumnQuals
//     plugin.Logger(ctx).Warn("getUser", "quals", quals)
//     id := quals["id"].GetInt64Value()
//     plugin.Logger(ctx).Warn("getUser", "id", id)
//     result, err := conn.GetUser(ctx, id)
//     if err != nil {
//         return nil, err
//     }
//     return result, nil
// }

type TheInternetHttpUrl struct {
	url string
}

type TheInternetHttpResponse struct{
	url string
	statusCode int
	body string
}

func getUrlList(ctx context.Context, quals plugin.KeyColumnEqualsQualMap) ([]string) {
	urlArgs := quals["url"].GetListValue()

	values := urlArgs.GetValues()

	if (values == nil) {
		plugin.Logger(ctx).Info("error getting values\n")
		return []string{quals["url"].GetStringValue()}
	}
	
	urls := []string{}
	
	for _, urlValue := range values {
		url := urlValue.GetStringValue()
		urls = append(urls, url)
	}
	return urls
}

func listHttpResource(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals

	plugin.Logger(ctx).Info("Here I am, rock me like a hurricane\n")

	values := getUrlList(ctx, quals)

	plugin.Logger(ctx).Info("values: %v\n", values)

	for _, urlValue := range values {
		plugin.Logger(ctx).Info("url: %s\n", urlValue)
		// res, err := http.Get(url)
		// if err != nil {
		// 	plugin.Logger(ctx).Info("error making http request: %s\n", err)
		// }

		// plugin.Logger(ctx).Info("client: got response!\n")
		// plugin.Logger(ctx).Info("client: status code: %d\n", res.StatusCode)
		// resBody, err := ioutil.ReadAll(res.Body)
		// // return TheInternetHttpResponse{
		// // 	statusCode: res.StatusCode,
		// // 	body: string(resBody),
		// // 	url: url,
		// // }, nil

		// d.StreamListItem(ctx, TheInternetHttpResponse{url, res.StatusCode, string(resBody)})
	}
	
	return nil, nil
}

func getHttpResource(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	requestURL := quals["url"].GetStringValue()

	res, err := http.Get(requestURL)
	if err != nil {
		plugin.Logger(ctx).Info("error making http request: %s\n", err)
	}

	plugin.Logger(ctx).Info("client: got response!\n")
	plugin.Logger(ctx).Info("client: status code: %d\n", res.StatusCode)
	resBody, err := ioutil.ReadAll(res.Body)
	return TheInternetHttpResponse{
		statusCode: res.StatusCode,
		body: string(resBody),
		url: requestURL,
	}, nil
}