// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/client/avns"
	"github.com/vmware/vcf-sdk-go/client/backup_restore"
	"github.com/vmware/vcf-sdk-go/client/bundles"
	"github.com/vmware/vcf-sdk-go/client/ceip"
	"github.com/vmware/vcf-sdk-go/client/certificates"
	"github.com/vmware/vcf-sdk-go/client/clusters"
	"github.com/vmware/vcf-sdk-go/client/credentials"
	"github.com/vmware/vcf-sdk-go/client/depot_settings"
	"github.com/vmware/vcf-sdk-go/client/domains"
	"github.com/vmware/vcf-sdk-go/client/fips_mode_details"
	"github.com/vmware/vcf-sdk-go/client/hosts"
	"github.com/vmware/vcf-sdk-go/client/identity_providers"
	"github.com/vmware/vcf-sdk-go/client/license_keys"
	"github.com/vmware/vcf-sdk-go/client/manifests"
	"github.com/vmware/vcf-sdk-go/client/nsxt_clusters"
	"github.com/vmware/vcf-sdk-go/client/network_pools"
	"github.com/vmware/vcf-sdk-go/client/nsxt_edge_clusters"
	"github.com/vmware/vcf-sdk-go/client/pscs"
	"github.com/vmware/vcf-sdk-go/client/personalities"
	"github.com/vmware/vcf-sdk-go/client/proxy_configuration"
	"github.com/vmware/vcf-sdk-go/client/releases"
	"github.com/vmware/vcf-sdk-go/client/resource_functionalities"
	"github.com/vmware/vcf-sdk-go/client/resource_warnings"
	"github.com/vmware/vcf-sdk-go/client/sddc"
	"github.com/vmware/vcf-sdk-go/client/sddc_managers"
	"github.com/vmware/vcf-sdk-go/client/sos"
	"github.com/vmware/vcf-sdk-go/client/system"
	"github.com/vmware/vcf-sdk-go/client/system_prechecks"
	"github.com/vmware/vcf-sdk-go/client/tasks"
	"github.com/vmware/vcf-sdk-go/client/tokens"
	"github.com/vmware/vcf-sdk-go/client/trusted_certificates"
	"github.com/vmware/vcf-sdk-go/client/upgradables"
	"github.com/vmware/vcf-sdk-go/client/upgrades"
	"github.com/vmware/vcf-sdk-go/client/users"
	"github.com/vmware/vcf-sdk-go/client/vcenters"
	"github.com/vmware/vcf-sdk-go/client/vrli"
	"github.com/vmware/vcf-sdk-go/client/vrops"
	"github.com/vmware/vcf-sdk-go/client/vrslcm"
	"github.com/vmware/vcf-sdk-go/client/vasa_providers"
	"github.com/vmware/vcf-sdk-go/client/vcf_services"
	"github.com/vmware/vcf-sdk-go/client/version_aliases_for_bundle_component_type"
	"github.com/vmware/vcf-sdk-go/client/vra"
	"github.com/vmware/vcf-sdk-go/client/vsan_health_check"
	"github.com/vmware/vcf-sdk-go/client/wsa"
)

// Default vcf client HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "sfo-vcf01.rainpole.io"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new vcf client HTTP client.
func NewHTTPClient(formats strfmt.Registry) *VcfClient {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new vcf client HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *VcfClient {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new vcf client client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *VcfClient {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(VcfClient)
	cli.Transport = transport
	cli.AvNs = avns.New(transport, formats)
	cli.BackupRestore = backup_restore.New(transport, formats)
	cli.Bundles = bundles.New(transport, formats)
	cli.CEIP = ceip.New(transport, formats)
	cli.Certificates = certificates.New(transport, formats)
	cli.Clusters = clusters.New(transport, formats)
	cli.Credentials = credentials.New(transport, formats)
	cli.DepotSettings = depot_settings.New(transport, formats)
	cli.Domains = domains.New(transport, formats)
	cli.FIPSModeDetails = fips_mode_details.New(transport, formats)
	cli.Hosts = hosts.New(transport, formats)
	cli.IdentityProviders = identity_providers.New(transport, formats)
	cli.LicenseKeys = license_keys.New(transport, formats)
	cli.Manifests = manifests.New(transport, formats)
	cli.NSXTClusters = nsxt_clusters.New(transport, formats)
	cli.NetworkPools = network_pools.New(transport, formats)
	cli.NSXTEdgeClusters = nsxt_edge_clusters.New(transport, formats)
	cli.PsCs = pscs.New(transport, formats)
	cli.Personalities = personalities.New(transport, formats)
	cli.ProxyConfiguration = proxy_configuration.New(transport, formats)
	cli.Releases = releases.New(transport, formats)
	cli.ResourceFunctionalities = resource_functionalities.New(transport, formats)
	cli.ResourceWarnings = resource_warnings.New(transport, formats)
	cli.SDDC = sddc.New(transport, formats)
	cli.SDDCManagers = sddc_managers.New(transport, formats)
	cli.SOS = sos.New(transport, formats)
	cli.System = system.New(transport, formats)
	cli.SystemPrechecks = system_prechecks.New(transport, formats)
	cli.Tasks = tasks.New(transport, formats)
	cli.Tokens = tokens.New(transport, formats)
	cli.TrustedCertificates = trusted_certificates.New(transport, formats)
	cli.Upgradables = upgradables.New(transport, formats)
	cli.Upgrades = upgrades.New(transport, formats)
	cli.Users = users.New(transport, formats)
	cli.VCenters = vcenters.New(transport, formats)
	cli.Vrli = vrli.New(transport, formats)
	cli.VroPs = vrops.New(transport, formats)
	cli.VRSLCM = vrslcm.New(transport, formats)
	cli.VasaProviders = vasa_providers.New(transport, formats)
	cli.VcfServices = vcf_services.New(transport, formats)
	cli.VersionAliasesForBundleComponentType = version_aliases_for_bundle_component_type.New(transport, formats)
	cli.VRA = vra.New(transport, formats)
	cli.VSANHealthCheck = vsan_health_check.New(transport, formats)
	cli.WSA = wsa.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// VcfClient is a client for vcf client
type VcfClient struct {
	AvNs avns.ClientService

	BackupRestore backup_restore.ClientService

	Bundles bundles.ClientService

	CEIP ceip.ClientService

	Certificates certificates.ClientService

	Clusters clusters.ClientService

	Credentials credentials.ClientService

	DepotSettings depot_settings.ClientService

	Domains domains.ClientService

	FIPSModeDetails fips_mode_details.ClientService

	Hosts hosts.ClientService

	IdentityProviders identity_providers.ClientService

	LicenseKeys license_keys.ClientService

	Manifests manifests.ClientService

	NSXTClusters nsxt_clusters.ClientService

	NetworkPools network_pools.ClientService

	NSXTEdgeClusters nsxt_edge_clusters.ClientService

	PsCs pscs.ClientService

	Personalities personalities.ClientService

	ProxyConfiguration proxy_configuration.ClientService

	Releases releases.ClientService

	ResourceFunctionalities resource_functionalities.ClientService

	ResourceWarnings resource_warnings.ClientService

	SDDC sddc.ClientService

	SDDCManagers sddc_managers.ClientService

	SOS sos.ClientService

	System system.ClientService

	SystemPrechecks system_prechecks.ClientService

	Tasks tasks.ClientService

	Tokens tokens.ClientService

	TrustedCertificates trusted_certificates.ClientService

	Upgradables upgradables.ClientService

	Upgrades upgrades.ClientService

	Users users.ClientService

	VCenters vcenters.ClientService

	Vrli vrli.ClientService

	VroPs vrops.ClientService

	VRSLCM vrslcm.ClientService

	VasaProviders vasa_providers.ClientService

	VcfServices vcf_services.ClientService

	VersionAliasesForBundleComponentType version_aliases_for_bundle_component_type.ClientService

	VRA vra.ClientService

	VSANHealthCheck vsan_health_check.ClientService

	WSA wsa.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *VcfClient) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AvNs.SetTransport(transport)
	c.BackupRestore.SetTransport(transport)
	c.Bundles.SetTransport(transport)
	c.CEIP.SetTransport(transport)
	c.Certificates.SetTransport(transport)
	c.Clusters.SetTransport(transport)
	c.Credentials.SetTransport(transport)
	c.DepotSettings.SetTransport(transport)
	c.Domains.SetTransport(transport)
	c.FIPSModeDetails.SetTransport(transport)
	c.Hosts.SetTransport(transport)
	c.IdentityProviders.SetTransport(transport)
	c.LicenseKeys.SetTransport(transport)
	c.Manifests.SetTransport(transport)
	c.NSXTClusters.SetTransport(transport)
	c.NetworkPools.SetTransport(transport)
	c.NSXTEdgeClusters.SetTransport(transport)
	c.PsCs.SetTransport(transport)
	c.Personalities.SetTransport(transport)
	c.ProxyConfiguration.SetTransport(transport)
	c.Releases.SetTransport(transport)
	c.ResourceFunctionalities.SetTransport(transport)
	c.ResourceWarnings.SetTransport(transport)
	c.SDDC.SetTransport(transport)
	c.SDDCManagers.SetTransport(transport)
	c.SOS.SetTransport(transport)
	c.System.SetTransport(transport)
	c.SystemPrechecks.SetTransport(transport)
	c.Tasks.SetTransport(transport)
	c.Tokens.SetTransport(transport)
	c.TrustedCertificates.SetTransport(transport)
	c.Upgradables.SetTransport(transport)
	c.Upgrades.SetTransport(transport)
	c.Users.SetTransport(transport)
	c.VCenters.SetTransport(transport)
	c.Vrli.SetTransport(transport)
	c.VroPs.SetTransport(transport)
	c.VRSLCM.SetTransport(transport)
	c.VasaProviders.SetTransport(transport)
	c.VcfServices.SetTransport(transport)
	c.VersionAliasesForBundleComponentType.SetTransport(transport)
	c.VRA.SetTransport(transport)
	c.VSANHealthCheck.SetTransport(transport)
	c.WSA.SetTransport(transport)
}
