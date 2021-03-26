/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-KPM-IEs"
 * 	found in "../v2/e2sm_kpm_v2.0.2-rm.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_E2SM_KPMv2_IndicationMessage_Format1_H_
#define	_E2SM_KPMv2_IndicationMessage_Format1_H_


#include "asn_application.h"

/* Including external dependencies */
#include "SubscriptionID.h"
#include "CellObjectID.h"
#include "GranularityPeriod.h"
#include "MeasurementData.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct MeasurementInfoList;

/* E2SM-KPMv2-IndicationMessage-Format1 */
typedef struct E2SM_KPMv2_IndicationMessage_Format1 {
	SubscriptionID_t	 subscriptID;
	CellObjectID_t	*cellObjID;	/* OPTIONAL */
	GranularityPeriod_t	*granulPeriod;	/* OPTIONAL */
	struct MeasurementInfoList	*measInfoList;	/* OPTIONAL */
	MeasurementData_t	 measData;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} E2SM_KPMv2_IndicationMessage_Format1_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_E2SM_KPMv2_IndicationMessage_Format1;
extern asn_SEQUENCE_specifics_t asn_SPC_E2SM_KPMv2_IndicationMessage_Format1_specs_1;
extern asn_TYPE_member_t asn_MBR_E2SM_KPMv2_IndicationMessage_Format1_1[5];

#ifdef __cplusplus
}
#endif

#endif	/* _E2SM_KPMv2_IndicationMessage_Format1_H_ */
#include "asn_internal.h"