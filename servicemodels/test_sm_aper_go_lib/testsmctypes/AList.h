/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "TEST-SM-IEs"
 * 	found in "../v1/test_sm.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_AList_H_
#define	_AList_H_


#include "asn_application.h"

/* Including external dependencies */
#include "asn_SEQUENCE_OF.h"
#include "constr_SEQUENCE_OF.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct TEST_NestedExtension;

/* AList */
typedef struct AList {
	struct AList__item {
		A_SEQUENCE_OF(struct TEST_NestedExtension) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} item;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} AList_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_AList;
extern asn_SEQUENCE_specifics_t asn_SPC_AList_specs_1;
extern asn_TYPE_member_t asn_MBR_AList_1[1];

#ifdef __cplusplus
}
#endif

#endif	/* _AList_H_ */
#include "asn_internal.h"
