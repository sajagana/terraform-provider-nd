// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_groups

// Generated conversion remains authoritative for API values. These adjustments
// retain configuration intent for Optional-only fields: an omitted field is
// unmanaged, and an explicitly empty route-server list remains empty rather
// than becoming null after the generated converter sees a zero-length slice.
// Import intentionally skips this adjustment so it can expose API values.
func (v *FabricGroupsModel) normalizeOptionalState(previous FabricGroupsModel) {
	if v.FabricType.ValueString() != "vxlan" && v.FabricType.ValueString() != "vxlanAci" {
		return
	}
	if previous.RouteServerCollection.IsNull() {
		v.RouteServerCollection = previous.RouteServerCollection
	} else if !previous.RouteServerCollection.IsUnknown() && len(previous.RouteServerCollection.Elements()) == 0 && v.RouteServerCollection.IsNull() {
		v.RouteServerCollection = previous.RouteServerCollection
	}
	if previous.ScheduledBackup.IsNull() {
		v.ScheduledBackup = previous.ScheduledBackup
	}
	if previous.ScheduledBackupTime.IsNull() {
		v.ScheduledBackupTime = previous.ScheduledBackupTime
	}
	if previous.MultisiteInterConnectBgpKey.IsNull() {
		v.MultisiteInterConnectBgpKey = previous.MultisiteInterConnectBgpKey
	}
	if previous.CloudSecKey.IsNull() {
		v.CloudSecKey = previous.CloudSecKey
	}
}
