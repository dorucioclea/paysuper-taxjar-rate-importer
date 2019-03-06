# TaxJar VAT rate importer

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache2.0-green.svg)](https://opensource.org/licenses/Apache2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/paysuper/paysuper-taxjar-rate-importer)](https://goreportcard.com/report/github.com/paysuper/paysuper-taxjar-rate-importer)

# Motivation
We use local tax rate table in PaySuper and this application used to per day sync data from TaxJar API to
local storage. This service uses the Simplemaps database of USA USPS zip codes (5 digits ) to fetch data 
from TaxJar API. This application sync only combined rates.

# Usage

Application designed to be launched with Kubernetes and handle all configuration from env variables:

| Variable      | Description                                                                                            |
|---------------|--------------------------------------------------------------------------------------------------------|
| TAX_JAR_TOKEN | TaxJar uses API keys to allow access to the API. This token could be generate in Account > API Access. |
| ZIP_CODE_FILE | The path to Simplemaps CSV postal code file.                                                           |
| CACHE_PATH    | The path to folder local LevelDB cache for rates. Default is `./cache`                                 |
| MAX_RPS       | The max allowed RPS for TaxJar API. Default is `250`                                                   |



 