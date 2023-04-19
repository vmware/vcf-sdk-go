# VMware Cloud Foundation SDK for Golang

[![License](https://img.shields.io/github/license/vmware/vcf-sdk-go.svg?style=for-the-badge)](LICENSE)

## Overview

A Go SDK for the VMware Cloud Foundation.

## Requirements

Required if building the SDK.

* [Go][golang-install] v1.19
* [go-swagger][go-swagger] v0.30.3
## Building the SDK

The SDK is auto-generated from the VMware Cloud Foundation swagger files.

### Check the `go-swagger` Version

Run the following command to check the version of [go-swagger][go-swagger]:

```console
make check-swagger
```

### Generate the SDK

Run the following command to generate the SDK:

```console
make swagger
```

>**Note**
>
> Generating the SDK will validate the swagger and run a script to normalize
the naming.

### Build the SDK

Run the following command to build the SDK:

```console
make build
```

## Contributing

The project team welcomes contributions from the community. Before you start
working with the project, please read our [Developer Certificate of Origin][vmware-cla-dco].

All contributions to this repository must be signed as described on that page.
Your signature certifies that you wrote the patch or have the right to pass it
on as an open-source patch.

Please open a GitHub [issue][issues] for the maintainers to update
the SDK instead of submitting a pull request.

For more detailed information, refer to the [contribution guidelines][contributing]
to get started.

## Support

We welcome you to use the GitHub [issues][issues] tracker to report bugs or
suggest enhancements.

When filing an issue, please check existing open, or recently closed, issues
to make sure someone else hasn't already reported the issue.

Please try to include as much information as you can. Details like these are
incredibly useful:

- A reproducible test case or series of steps.
- Any modifications you've made relevant to the bug.
- Anything unusual about your environment or deployment.

## License

Copyright 2023 VMware, Inc.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice,
this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation
and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
POSSIBILITY OF SUCH DAMAGE.

[//]: Links

[contributing]: CONTRIBUTING.md
[issues]: https://github.com/vmware/vcf-sdk-go/issues
[golang-install]: https://golang.org/doc/install
[go-swagger]: https://github.com/go-swagger/go-swagger
[vmware-cla-dco]: https://cla.vmware.com/dco
