// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//nolint
package main

import (
	"fmt"
	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/kpmctypes"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"google.golang.org/protobuf/proto"
)

type servicemodel string

const smname = "e2sm_kpm_v2"
const smversion = "v2"
const modulename = "e2sm_kpm_v2.so.2.0"

func (sm servicemodel) ServiceModelData() (string, string, string) {
	return smname, smversion, modulename
}

func (sm servicemodel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := kpmv2ctypes.PerDecodeE2SmKpmIndicationHeader(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmKpmIndicationHeader to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmKpmIndicationHeader %s", err)
	}

	return protoBytes, nil
}

func (sm servicemodel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_kpm_v2.E2SmKpmIndicationHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmIndicationHeader %s", err)
	}

	perBytes, err := kpmv2ctypes.PerEncodeE2SmKpmIndicationHeader(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmIndicationHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm servicemodel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

//func (sm servicemodel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
//	perBytes, err := kpmv2ctypes.PerDecodeE2SmKpmIndicationMessage(asn1Bytes)
//	if err != nil {
//		return nil, fmt.Errorf("error decoding E2SmKpmIndicationMessage to PER %s", err)
//	}
//
//	protoBytes, err := proto.Marshal(perBytes)
//	if err != nil {
//		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmKpmIndicationMessage %s", err)
//	}
//
//	return protoBytes, nil
//}

func (sm servicemodel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

//func (sm servicemodel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
//	protoObj := new(e2sm_kpm_v2.E2SmKpmIndicationMessage)
//	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
//		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmIndicationMessage %s", err)
//	}
//
//	perBytes, err := kpmv2ctypes.PerEncodeE2SmKpmIndicationMessage(protoObj)
//	if err != nil {
//		return nil, fmt.Errorf("error encoding E2SmKpmIndicationMessage to PER %s", err)
//	}
//
//	return perBytes, nil
//}

func (sm servicemodel) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := kpmv2ctypes.PerDecodeE2SmKpmRanfunctionDescription(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmKpmRanfunctionDescription to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmKpmRanfunctionDescription %s", err)
	}

	return protoBytes, nil
}

func (sm servicemodel) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_kpm_v2.E2SmKpmRanfunctionDescription)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmRanfunctionDescription %s", err)
	}

	perBytes, err := kpmv2ctypes.PerEncodeE2SmKpmRanfunctionDescription(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmRanfunctionDescription to PER %s", err)
	}

	return perBytes, nil
}

func (sm servicemodel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := kpmv2ctypes.PerDecodeE2SmKpmEventTriggerDefinition(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmKpmEventTriggerDefinition to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmKpmEventTriggerDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm servicemodel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_kpm_v2.E2SmKpmEventTriggerDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmEventTriggerDefinition %s", err)
	}

	perBytes, err := kpmv2ctypes.PerEncodeE2SmKpmEventTriggerDefinition(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmEventTriggerDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm servicemodel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

//func (sm servicemodel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
//	perBytes, err := kpmv2ctypes.PerDecodeE2SmKpmActionDefinition(asn1Bytes)
//	if err != nil {
//		return nil, fmt.Errorf("error decoding E2SmKpmActionDefinitio to PER %s", err)
//	}
//
//	protoBytes, err := proto.Marshal(perBytes)
//	if err != nil {
//		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmKpmActionDefinition %s", err)
//	}
//
//	return protoBytes, nil
//}

func (sm servicemodel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

//func (sm servicemodel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
//	protoObj := new(e2sm_kpm_v2.E2SmKpmActionDefinition)
//	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
//		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmActionDefinition %s", err)
//	}
//
//	perBytes, err := kpmv2ctypes.PerEncodeE2SmKpmActionDefinition(protoObj)
//	if err != nil {
//		return nil, fmt.Errorf("error encoding E2SmKpmActionDefinition to PER %s", err)
//	}
//
//	return perBytes, nil
//}

func (sm servicemodel) DecodeRanFunctionDescription(asn1bytes []byte) (*types.RanfunctionNameDef, *types.RicEventTriggerList, *types.RicReportList, error){
    return nil, nil, nil, fmt.Errorf("not implemented")
}

// It is redundant so far - could be reused for future, if you need to extract something specific from RanFunctionDescription message
//func (sm servicemodel) DecodeRanFunctionDescription(asn1bytes []byte) (*types.RanfunctionNameDef, *types.RicEventTriggerList, *types.RicReportList, error){
//	e2SmKpmPdu, err := kpmv2ctypes.PerDecodeE2SmKpmRanfunctionDescription(asn1bytes)
//	if err != nil {
//		return nil, nil, nil, err
//	}
//	return pdudecoder.DecodeE2SmKpmRanfunctionDescription(e2SmKpmPdu)
//}

func (sm servicemodel) ControlHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm servicemodel) ControlHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm servicemodel) ControlMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm servicemodel) ControlMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm servicemodel) ControlOutcomeASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm servicemodel) ControlOutcomeProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

//ServiceModel is the exported symbol that gives an entry point to this shared module
var ServiceModel servicemodel