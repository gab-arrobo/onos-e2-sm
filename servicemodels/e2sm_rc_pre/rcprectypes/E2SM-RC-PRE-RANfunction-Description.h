/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RC-PRE-IEs"
 * 	found in "e2sm-rc-pre-v1.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_E2SM_RC_PRE_RANfunction_Description_H_
#define	_E2SM_RC_PRE_RANfunction_Description_H_


#include "asn_application.h"

/* Including external dependencies */
#include "RANfunction-Name.h"
#include "asn_SEQUENCE_OF.h"
#include "constr_SEQUENCE_OF.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct RIC_EventTriggerStyle_List;
struct RIC_ReportStyle_List;

/* E2SM-RC-PRE-RANfunction-Description */
typedef struct E2SM_RC_PRE_RANfunction_Description {
	RANfunction_Name_t	 ranFunction_Name;
	struct E2SM_RC_PRE_RANfunction_Description__e2SM_RC_PRE_RANfunction_Item {
		struct E2SM_RC_PRE_RANfunction_Description__e2SM_RC_PRE_RANfunction_Item__ric_EventTriggerStyle_List {
			A_SEQUENCE_OF(struct RIC_EventTriggerStyle_List) list;
			
			/* Context for parsing across buffer boundaries */
			asn_struct_ctx_t _asn_ctx;
		} *ric_EventTriggerStyle_List;
		struct E2SM_RC_PRE_RANfunction_Description__e2SM_RC_PRE_RANfunction_Item__ric_ReportStyle_List {
			A_SEQUENCE_OF(struct RIC_ReportStyle_List) list;
			
			/* Context for parsing across buffer boundaries */
			asn_struct_ctx_t _asn_ctx;
		} *ric_ReportStyle_List;
		/*
		 * This type is extensible,
		 * possible extensions are below.
		 */
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} e2SM_RC_PRE_RANfunction_Item;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} E2SM_RC_PRE_RANfunction_Description_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_E2SM_RC_PRE_RANfunction_Description;

#ifdef __cplusplus
}
#endif

#endif	/* _E2SM_RC_PRE_RANfunction_Description_H_ */
#include "asn_internal.h"