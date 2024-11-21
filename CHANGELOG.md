# CHANGELOG

## [v0.5.0](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.5.0)

> Release Date: 2024-11-25

- Added support for VCF 5.2.1 APIs.

## [v0.4.1](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.4.1)

> Release Date: 2024-11-11

- Added a custom template for response models.

## [v0.4.0](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.4.0)

> Release Date: 2024-10-14

- Added support for VCF 5.2.0 APIs.
- Replaced `go-swagger` with `oapi-codegen`.

## [v0.3.3](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.3.3)

> Release Date: 2024-07-24

- Fixed vSAN ESA config serialization when value is `false`.

## [v0.3.2](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.3.2)

> Release Date: 2024-06-07

- Accept 0 as a valid VLAN ID for network pools and edge node management
  portgroups.

## [v0.3.1](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.3.1)

> Release Date: 2024-06-07

> [!IMPORTANT]
> Do **not** use this release. Use v0.3.2 instead.

## [v0.3.0](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.3.0)

> Release Date: 2024-05-23

- Added support for VCF 5.1.1 APIs.
- Updated `go-swagger` to v0.30.5.

## [v0.2.4](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.2.4)

> Release Date: 2024-03-11

- Patched personalities API.

## [v0.2.3](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.2.3)

> Release Date: 2024-02-27

- Updated `GeneveVlanID` in `NsxTClusterSpec` to а required field.

## [v0.2.2](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.2.2)

> Release Date: 2023-12-11

- Updated `AutoRotateCredentialPolicyInputSpec.EnableAutoRotatePolicy` from
  being omitted when set to `false`

## [v0.2.1](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.2.1)

> Release Date: 2023-11-22

- Added `202` as response code to `/v1/domains/{id}/resource-certificates` API
  endpoint.

## [v0.2.0](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.2.0)

> Release Date: 2023-10-18

- Added support for VCF 4.5.2 APIs.

> [!NOTE]
> `PUT /v1/personalities/files` API is omitted as it made the swagger
> file invalid.

## [v0.1.3](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.1.3)

> Release Date: 2023-10-02

- Updated `MemorySharesValue` in SDDC's `ResourcePoolSpec` а required field.

## [v0.1.2](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.1.2)

> Release Date: 2023-10-02

- Updated `TransportVlanID` in `SDDCNSXTSpec` to а required field.

## [v0.1.1](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.1.1)

> Release Date: 2023-05-31

- Removed the `GET` and `POST` abbreviations from the additional initialism
  list, for better dev experience when using the SDK.

## [v0.1.0](https://github.com/vmware/vcf-sdk-go/releases/tag/v0.1.0)

> Release Date: 2023-04-21

- Initial release with support for VCF 4.5.0.
