#!/usr/bin/env sh
#Copyright 2023 VMware, Inc.
#SPDX-License-Identifier: BSD-2-Clause

# These operations are required due to way initialisms are handled in the go-swagger tool

rename() {
    find client -type d -name "*$1*" | sed "s/\(.*\)$1\(.*\)/mv & \1$2\2/" | sh
    find client -type f -name "*$1*" | sed "s/\(.*\)$1\(.*\)/mv & \1$2\2/" | sh
}

echo "Renaming directories under /client"
rename a_v_ns avns
rename n_s_x_t_clusters nsxt_clusters
rename nsx_t_edge_clusters nsxt_edge_clusters
rename p_s_cs pscs
rename v_centers vcenters
rename v_r_l_i vrli
rename v_r_o_ps vrops
rename v_r_s_l_c_m vrslcm
rename v_r_a vra
rename v_san_health_check vsan_health_check

echo "Replace package names and occurrences of incorrect abbreviations with correct ones"
find client -type f -name '*.go' -exec sed -i "" 's/a_v_ns/avns/g' {} +
find client -type f -name '*.go' -exec sed -i "" 's/n_s_x_t/nsxt/g' {} +
find client -type f -name '*.go' -exec sed -i "" 's/nsx_t/nsxt/g' {} +
find client -type f -name '*.go' -exec sed -i "" 's/p_s_cs/pscs/g' {} +
find client -type f -name '*.go' -exec sed -i "" 's/v_r_l_i/vrli/g' {} +
find client -type f -name '*.go' -exec sed -i "" 's/v_centers/vcenters/g' {} +
find client -type f -name '*.go' -exec sed -i "" 's/v_r_o_ps/vrops/g' {} +
find client -type f -name '*.go' -exec sed -i "" 's/v_r_s_l_c_m/vrslcm/g' {} +
find client -type f -name '*.go' -exec sed -i "" 's/v_r_a/vra/g' {} +
find client -type f -name '*.go' -exec sed -i "" 's/v_san_health_check/vsan_health_check/g' {} +

echo "Changing main client filename"
mv client/vcf_client_client.go client/vcf_client.go
echo "Changing avns client filename"
mv client/avns/av_ns_client.go client/avns/avns_client.go
echo "Changing c_s_rs filenames"
mv client/certificates/generates_c_s_rs_parameters.go client/certificates/generates_csrs_parameters.go
mv client/certificates/generates_c_s_rs_responses.go client/certificates/generates_csrs_responses.go
mv client/certificates/get_c_s_rs_parameters.go client/certificates/get_csrs_parameters.go
mv client/certificates/get_c_s_rs_responses.go client/certificates/get_csrs_responses.go
echo "Changing p_s_cs client filename"
mv client/pscs/ps_cs_client.go client/pscs/pscs_client.go
echo "Changing vro_ps client filename"
mv client/vrops/vro_ps_client.go client/vrops/vrops_client.go
echo "Changing g_e_t filenames"
mv client/bundles/get_bundles_for_skip_upgrade_using_g_e_t_parameters.go client/bundles/get_bundles_for_skip_upgrade_using_get_parameters.go
mv client/bundles/get_bundles_for_skip_upgrade_using_g_e_t_responses.go client/bundles/get_bundles_for_skip_upgrade_using_get_responses.go
mv client/nsxt_clusters/get_validation_result_using_g_e_t_parameters.go client/nsxt_clusters/get_validation_result_using_get_parameters.go
mv client/nsxt_clusters/get_validation_result_using_g_e_t_responses.go client/nsxt_clusters/get_validation_result_using_get_responses.go
mv client/nsxt_clusters/validate_ip_pool_using_p_o_s_t_parameters.go client/nsxt_clusters/validate_ip_pool_using_post_parameters.go
mv client/nsxt_clusters/validate_ip_pool_using_p_o_s_t_responses.go client/nsxt_clusters/validate_ip_pool_using_post_responses.go
mv client/upgrades/get_precheck_using_g_e_t_parameters.go client/upgrades/get_precheck_using_get_parameters.go
mv client/upgrades/get_precheck_using_g_e_t_responses.go client/upgrades/get_precheck_using_get_responses.go
mv client/upgrades/perform_prechecks_using_p_o_s_t_parameters.go client/upgrades/perform_prechecks_using_post_parameters.go
mv client/upgrades/perform_prechecks_using_p_o_s_t_responses.go client/upgrades/perform_prechecks_using_post_responses.go
mv client/users/get_all_ui_users_using_g_e_t_parameters.go client/users/get_all_ui_users_using_get_parameters.go
mv client/users/get_all_ui_users_using_g_e_t_responses.go client/users/get_all_ui_users_using_get_responses.go