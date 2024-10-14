# CHANGELOG

## [v0.4.0](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.4.0)

> Release Date: 2024-10-14

Support for VCF 5.2.0 APIs\
Replace `go-swagger` with `oapi-codegen`

## [v0.3.3](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.3.3)

> Release Date: 2024-07-24

Fix vSAN ESA config serialization when value is false

## [v0.3.2](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.3.2)

> Release Date: 2024-06-07

Accept 0 as a valid VLAN ID for network pools and edge node management portgroups

## [v0.3.1](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.3.1)

> Release Date: 2024-06-07

Do not use this release. Use v0.3.2 instead.

## [v0.3.0](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.3.0)

> Release Date: 2024-05-23

Support for VCF 5.1.1 APIs\
Move to go-swagger 0.30.5

## [v0.2.4](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.2.4)

> Release Date: 2024-03-11

Patch personalities API

## [v0.2.3](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.2.3)

> Release Date: 2024-02-27

Made GeneveVlanID in NsxTClusterSpec а required field.

## [v0.2.2](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.2.2)

> Release Date: 2023-12-11

Prevent `AutoRotateCredentialPolicyInputSpec.EnableAutoRotatePolicy` being omitted when set to `false` 

## [v0.2.1](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.2.1)

> Release Date: 2023-11-22

Include 202 as response code to "/v1/domains/{id}/resource-certificates" API endpoint

## [v0.2.0](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.2.0)

> Release Date: 2023-10-18
 
Support APIs from VCF 4.5.2\
Note: PUT “/v1/personalities/files” API omitted as it made the swagger file invalid.

## [v0.1.3](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.1.3)

> Release Date: 2023-10-02

Made MemorySharesValue in SDDC's ResourcePoolSpec а required field.

## [v0.1.2](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.1.2)

> Release Date: 2023-10-02

Made TransportVlanID in SDDCNSXTSpec а required field.

## [v0.1.1](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.1.1)

> Release Date: 2023-05-31

Removed the GET and POST abbreviations from the additional initialism list, for better dev experience when using the SDK.

## [v0.1.0](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.1.0)

> Release Date: 2023-04-21

Initial release, bringing support for VCF 4.5.0
