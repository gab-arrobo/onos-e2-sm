// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createMeasurementCondUEIDItem() (*e2sm_kpm_v2.MeasurementCondUeidItem, error) {

	muei := &e2sm_kpm_v2.MatchingUeidItem{
		UeId: &e2sm_kpm_v2.UeIdentity{
			Value: "SomeUE",
		},
	}
	muel := &e2sm_kpm_v2.MatchingUeidList{
		Value: make([]*e2sm_kpm_v2.MatchingUeidItem, 0),
	}
	muel.Value = append(muel.Value, muei)

	mci1 := &e2sm_kpm_v2.MatchingCondItem{
		MatchingCondItem: &e2sm_kpm_v2.MatchingCondItem_MeasLabel{
			MeasLabel: &e2sm_kpm_v2.MeasurementLabel{
				PlmnId: &e2sm_kpm_v2.PlmnIdentity{
					Value: []byte{0x1, 0x2, 0x3},
				},
				SliceId: &e2sm_kpm_v2.Snssai{
					SD:  []byte{0x01, 0x02, 0x03},
					SSt: []byte{0x01},
				},
				FiveQi: &e2sm_kpm_v2.FiveQi{
					Value: 23,
				},
				QFi: &e2sm_kpm_v2.Qfi{
					Value: 52,
				},
				QCi: &e2sm_kpm_v2.Qci{
					Value: 24,
				},
				QCimax: &e2sm_kpm_v2.Qci{
					Value: 30,
				},
				QCimin: &e2sm_kpm_v2.Qci{
					Value: 1,
				},
				ARpmax: &e2sm_kpm_v2.Arp{
					Value: 15,
				},
				ARpmin: &e2sm_kpm_v2.Arp{
					Value: 1,
				},
				BitrateRange:     25,
				LayerMuMimo:      1,
				SUm:              e2sm_kpm_v2.SUM_SUM_TRUE,
				DistBinX:         123,
				DistBinY:         456,
				DistBinZ:         789,
				PreLabelOverride: e2sm_kpm_v2.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE,
				StartEndInd:      e2sm_kpm_v2.StartEndInd_START_END_IND_END,
			},
		},
	}

	if err := mci1.Validate(); err != nil {
		return nil, err
	}

	mci2 := &e2sm_kpm_v2.MatchingCondItem{
		MatchingCondItem: &e2sm_kpm_v2.MatchingCondItem_TestCondInfo{
			TestCondInfo: &e2sm_kpm_v2.TestCondInfo{
				TestType: &e2sm_kpm_v2.TestCondType{
					TestCondType: &e2sm_kpm_v2.TestCondType_AMbr{
						AMbr: e2sm_kpm_v2.AMBR_AMBR_TRUE,
					},
				},
				TestExpr: e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_GREATERTHAN,
				TestValue: &e2sm_kpm_v2.TestCondValue{
					TestCondValue: &e2sm_kpm_v2.TestCondValue_ValueInt{
						ValueInt: 21,
					},
				},
			},
		},
	}

	if err := mci2.Validate(); err != nil {
		return nil, err
	}

	mcl := &e2sm_kpm_v2.MatchingCondList{
		Value: make([]*e2sm_kpm_v2.MatchingCondItem, 0),
	}
	mcl.Value = append(mcl.Value, mci1)
	mcl.Value = append(mcl.Value, mci2)

	measType := &e2sm_kpm_v2.MeasurementType{
		MeasurementType: &e2sm_kpm_v2.MeasurementType_MeasName{
			MeasName: &e2sm_kpm_v2.MeasurementTypeName{
				Value: "onf",
			},
		},
	}

	res := &e2sm_kpm_v2.MeasurementCondUeidItem{
		MatchingCond: mcl,
		MeasType:     measType,
		//MatchingUeidList: muel,
	}
	if err := res.Validate(); err != nil {
		return nil, err
	}
	return res, nil
}

func Test_xerEncodeMeasurementCondUEIDItem(t *testing.T) {

	mcueIDi, err := createMeasurementCondUEIDItem()
	assert.NilError(t, err)

	xer, err := xerEncodeMeasurementCondUeIDItem(mcueIDi)
	assert.NilError(t, err)
	//assert.Equal(t, 1441, len(xer))
	//assert.Equal(t, 1296, len(xer)) // without MatchingUeidList
	t.Logf("MeasurementCondUEidItem XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementCondUEIDItem(t *testing.T) {

	mcueIDi, err := createMeasurementCondUEIDItem()
	assert.NilError(t, err)

	xer, err := xerEncodeMeasurementCondUeIDItem(mcueIDi)
	assert.NilError(t, err)
	//assert.Equal(t, 1441, len(xer))
	//assert.Equal(t, 1296, len(xer)) // without MatchingUeidList
	t.Logf("MeasurementCondUEidItem XER\n%s", string(xer))

	result, err := xerDecodeMeasurementCondUeIDItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementCondUEidItem XER - decoded\n%s", result)
}

func Test_perEncodeMeasurementCondUEIDItem(t *testing.T) {

	mcueIDi, err := createMeasurementCondUEIDItem()
	assert.NilError(t, err)

	per, err := perEncodeMeasurementCondUeIDItem(mcueIDi)
	assert.NilError(t, err)
	//assert.Equal(t, 57, len(per))
	//assert.Equal(t, 47, len(per)) // without MatchingUeidList
	t.Logf("MeasurementCondUEidItem PER\n%s", string(per))
}

func Test_perDecodeMeasurementCondUEIDItem(t *testing.T) {

	mcueIDi, err := createMeasurementCondUEIDItem()
	assert.NilError(t, err)

	per, err := perEncodeMeasurementCondUeIDItem(mcueIDi)
	assert.NilError(t, err)
	//assert.Equal(t, 57, len(per))
	//assert.Equal(t, 47, len(per)) // without MatchingUeidList
	t.Logf("MeasurementCondUEidItem PER\n%s", string(per))

	result, err := perDecodeMeasurementCondUeIDItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementCondUEidItem PER - decoded\n%v", result)
}