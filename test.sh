#!/bin/bash

set -x
trap "kill 0" EXIT

STEAMPIPE_PLUGIN_LOCATION=$HOME/.steampipe/plugins/local/theinternet/theinternet.plugin
BINARY_NAME=steampipe-plugin-the-internet
TEST_QUERY="select * from theinternet where url = 'https://www.hughrawlinson.me'"

go build
mv $BINARY_NAME $STEAMPIPE_PLUGIN_LOCATION

tail -f "$HOME/.steampipe/logs/plugin-$(date +"%F").log" &

STEAMPIPE_LOG=INFO steampipe query "$TEST_QUERY"
