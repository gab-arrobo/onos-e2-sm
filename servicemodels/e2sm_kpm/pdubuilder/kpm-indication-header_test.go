// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package sandbox

import (
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmIndicationHeader(t *testing.T) {
	var plmnID = "ONF"
	var gNbCuUpId int64 = 0
	var gNbDuId int64 = 0
	var plmnIDnrcgi = "onf"
	var sst = "1"
	var sd = "SD1"
	var fiveQi int32 = 0
	var qCi int32 = 0

	newE2SmKpmPdu, err := CreateE2SmKpmIndicationHeader(plmnID, gNbCuUpId, gNbDuId, plmnIDnrcgi, sst, sd, fiveQi, qCi)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}