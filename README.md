# TaxJar VAT rate importer

[![License: GPL 3.0](https://img.shields.io/badge/License-GPL3.0-green.svg)](https://opensource.org/licenses/Gpl3.0)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/paysuper/paysuper-taxjar-rate-importer/issues)
[![Build Status](https://travis-ci.org/paysuper/paysuper-tax-service.svg?branch=develop)](https://travis-ci.org/paysuper/paysuper-tax-service)
[![Go Report Card](https://goreportcard.com/badge/github.com/paysuper/paysuper-taxjar-rate-importer)](https://goreportcard.com/report/github.com/paysuper/paysuper-taxjar-rate-importer)

PaySuper TaxJar VAT everyday syncs data from the TaxJar API to the PaySuper local storage for the taxes rates (only combined rates are synchronised). This service uses the Simplemaps database of the USA USPS zip codes (5 digits) to fetch data from the TaxJar API.

***

## Table of Contents

- [Usage](#usage)
- [Contributing](#contributing-feature-requests-and-support)
- [License](#license)

# Usage

Application can be launched with Kubernetes and handles all configuration from the environment variables:

| Variable      | Description                                                                                            |
|---------------|--------------------------------------------------------------------------------------------------------|
| TAX_JAR_TOKEN | TaxJar uses API keys to allow access to the API. This token could be generated in Account > API Access. |
| ZIP_CODE_FILE | The path to the Simplemaps CSV postal code file.                                                           |
| CACHE_PATH    | The path to the folder of the local LevelDB cache for rates. Default is `./cache`.                                 |
| MAX_RPS       | The max allowed RPS for the TaxJar API. Default is `250`.                                                   |

## Contributing, Feature Requests and Support

If you like this project then you can put a ⭐️ on it. It means a lot to us.

If you have an idea of how to improve PaySuper (or any of the product parts) or have general feedback, you're welcome to submit a [feature request](../../issues/new?assignees=&labels=&template=feature_request.md&title=).

Chances are, you like what we have already but you may require a custom integration, a special license or something else big and specific to your needs. We're generally open to such conversations.

If you have a question and can't find the answer yourself, you can [raise an issue](../../issues/new?assignees=&labels=&template=support-request.md&title=I+have+a+question+about+%3Cthis+and+that%3E+%5BSupport%5D) and describe what exactly you're trying to do. We'll do our best to reply in a meaningful time.

We feel that a welcoming community is important and we ask that you follow PaySuper's [Open Source Code of Conduct](https://github.com/paysuper/code-of-conduct/blob/master/README.md) in all interactions with the community.

PaySuper welcomes contributions from anyone and everyone. Please refer to [our contribution guide to learn more](CONTRIBUTING.md).

## License

The project is available as open source under the terms of the [GPL v3 License](https://www.gnu.org/licenses/gpl-3.0).
