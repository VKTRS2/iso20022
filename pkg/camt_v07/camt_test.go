// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v07

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/iso20022/pkg/common"
	"github.com/stretchr/testify/assert"
)

const (
	testTimeString = "2014-11-12T11:45:26.371Z"
)

func TestDocumentCamt00300107(t *testing.T) {
	sample := DocumentCamt00300107{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt00300107{
		GetAcct: GetAccountV07{
			MsgHdr: MessageHeader9{
				MsgId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"GetAcct":{"MsgHdr":{"MsgId":"Id"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.003.001.07" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><GetAcct xmlns="urn:iso:std:iso:20022:tech:xsd:camt.003.001.07"><MsgHdr><MsgId>Id</MsgId></MsgHdr></GetAcct></Document>`)
}

func TestDocumentCamt00900107(t *testing.T) {
	sample := DocumentCamt00900107{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt00900107{
		GetLmt: GetLimitV07{
			MsgHdr: MessageHeader9{
				MsgId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"GetLmt":{"MsgHdr":{"MsgId":"Id"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.009.001.07" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><GetLmt xmlns="urn:iso:std:iso:20022:tech:xsd:camt.009.001.07"><MsgHdr><MsgId>Id</MsgId></MsgHdr></GetLmt></Document>`)
}

func TestDocumentCamt01100107(t *testing.T) {
	sample := DocumentCamt01100107{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt01100107{
		ModfyLmt: ModifyLimitV07{
			MsgHdr: MessageHeader1{
				MsgId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"ModfyLmt":{"MsgHdr":{"MsgId":"Id"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.011.001.07" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ModfyLmt xmlns="urn:iso:std:iso:20022:tech:xsd:camt.011.001.07"><MsgHdr><MsgId>Id</MsgId></MsgHdr></ModfyLmt></Document>`)
}

func TestDocumentCamt01900107(t *testing.T) {
	sample := DocumentCamt01900107{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt01900107{
		RtrBizDayInf: ReturnBusinessDayInformationV07{
			MsgHdr: MessageHeader7{
				MsgId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"RtrBizDayInf":{"MsgHdr":{"MsgId":"Id"},"RptOrErr":{}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.019.001.07" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><RtrBizDayInf xmlns="urn:iso:std:iso:20022:tech:xsd:camt.019.001.07"><MsgHdr><MsgId>Id</MsgId></MsgHdr><RptOrErr></RptOrErr></RtrBizDayInf></Document>`)
}

func TestDocumentCamt02300107(t *testing.T) {
	sample := DocumentCamt02300107{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt02300107{
		BckpPmt: BackupPaymentV07{
			MsgHdr: MessageHeader1{
				MsgId: "Id",
			},
			TrfdAmt: Amount2Choice{
				AmtWthCcy: ActiveCurrencyAndAmount{Ccy: "ABC"},
			},
			Cdtr: SystemMember3{
				MmbId: MemberIdentification3Choice{
					BICFI:       "0000AA00BBB",
					ClrSysMmbId: ClearingSystemMemberIdentification2{MmbId: "Id"},
					Othr:        GenericFinancialIdentification1{Id: "ID"},
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"BckpPmt":{"MsgHdr":{"MsgId":"Id"},"TrfdAmt":{"AmtWthtCcy":0,"AmtWthCcy":{"Value":0,"Ccy":"ABC"}},"Cdtr":{"MmbId":{"BICFI":"0000AA00BBB","ClrSysMmbId":{"MmbId":"Id"},"Othr":{"Id":"ID"}}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.023.001.07" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><BckpPmt xmlns="urn:iso:std:iso:20022:tech:xsd:camt.023.001.07"><MsgHdr><MsgId>Id</MsgId></MsgHdr><TrfdAmt><AmtWthtCcy>0</AmtWthtCcy><AmtWthCcy Ccy="ABC">0</AmtWthCcy></TrfdAmt><Cdtr><MmbId><BICFI>0000AA00BBB</BICFI><ClrSysMmbId><MmbId>Id</MmbId></ClrSysMmbId><Othr><Id>ID</Id></Othr></MmbId></Cdtr></BckpPmt></Document>`)
}

func TestDocumentCamt01200107(t *testing.T) {
	sample := DocumentCamt01200107{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt01200107{
		DelLmt: DeleteLimitV07{
			MsgHdr: MessageHeader1{
				MsgId: "Id",
			},
			LmtDtls: LimitStructure2Choice{
				AllCurLmts: LimitIdentification6{
					Tp: LimitType1Choice{
						Cd:    "INBI",
						Prtry: "Prtry",
					},
				},
				CurLmtId: LimitIdentification5{
					Tp: LimitType1Choice{
						Cd:    "INBI",
						Prtry: "Prtry",
					},
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"DelLmt":{"MsgHdr":{"MsgId":"Id"},"LmtDtls":{"CurLmtId":{"Tp":{"Cd":"INBI","Prtry":"Prtry"}},"AllCurLmts":{"Tp":{"Cd":"INBI","Prtry":"Prtry"}}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.012.001.07" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><DelLmt xmlns="urn:iso:std:iso:20022:tech:xsd:camt.012.001.07"><MsgHdr><MsgId>Id</MsgId></MsgHdr><LmtDtls><CurLmtId><Tp><Cd>INBI</Cd><Prtry>Prtry</Prtry></Tp></CurLmtId><AllCurLmts><Tp><Cd>INBI</Cd><Prtry>Prtry</Prtry></Tp></AllCurLmts></LmtDtls></DelLmt></Document>`)
}

func TestDocumentCamt08700107(t *testing.T) {
	sample := DocumentCamt08700107{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentCamt08700107{
		ReqToModfyPmt: RequestToModifyPaymentV07{
			Assgnmt: CaseAssignment5{
				Id:      "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			Undrlyg: UnderlyingTransaction6Choice{
				Initn: UnderlyingPaymentInstruction6{
					OrgnlInstdAmt: ActiveOrHistoricCurrencyAndAmount{
						Ccy: "ABC",
					},
				},
				IntrBk: UnderlyingPaymentTransaction5{
					OrgnlIntrBkSttlmAmt: ActiveOrHistoricCurrencyAndAmount{
						Ccy: "ABC",
					},
					OrgnlIntrBkSttlmDt: common.ISODate(testTime),
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"ReqToModfyPmt":{"Assgnmt":{"Id":"Id","Assgnr":{"Pty":{},"Agt":{"FinInstnId":{}}},"Assgne":{"Pty":{},"Agt":{"FinInstnId":{}}},"CreDtTm":"2014-11-12T11:45:26.371"},"Undrlyg":{"Initn":{"OrgnlInstdAmt":{"Value":0,"Ccy":"ABC"}},"IntrBk":{"OrgnlIntrBkSttlmAmt":{"Value":0,"Ccy":"ABC"},"OrgnlIntrBkSttlmDt":"2014-11-12"},"StmtNtry":{}},"Mod":{}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.087.001.07" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ReqToModfyPmt xmlns="urn:iso:std:iso:20022:tech:xsd:camt.087.001.07"><Assgnmt><Id>Id</Id><Assgnr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgnr><Assgne><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgne><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></Assgnmt><Undrlyg><Initn><OrgnlInstdAmt Ccy="ABC">0</OrgnlInstdAmt></Initn><IntrBk><OrgnlIntrBkSttlmAmt Ccy="ABC">0</OrgnlIntrBkSttlmAmt><OrgnlIntrBkSttlmDt>2014-11-12</OrgnlIntrBkSttlmDt></IntrBk><StmtNtry></StmtNtry></Undrlyg><Mod></Mod></ReqToModfyPmt></Document>`)
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AccountCriteria3Choice{}.Validate())
	assert.Nil(t, AccountCriteria7{}.Validate())
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountIdentificationSearchCriteria2Choice{}.Validate())
	assert.Nil(t, AccountQuery3{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.NotNil(t, BalanceType11Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.Nil(t, CashAccountReturnCriteria5{}.Validate())
	assert.Nil(t, CashAccountSearchCriteria7{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.Nil(t, CashBalance12{}.Validate())
	assert.Nil(t, CashBalanceReturnCriteria2{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.Nil(t, DateAndDateTimeSearch4Choice{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.Nil(t, DatePeriod2{}.Validate())
	assert.Nil(t, DatePeriodSearch1Choice{}.Validate())
	assert.Nil(t, DateTimePeriod1{}.Validate())
	assert.Nil(t, DateTimeSearch2Choice{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, GetAccountV07{}.Validate())
	assert.NotNil(t, MessageHeader9{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, RequestType4Choice{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, ActiveAmountRange3Choice{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmountRange3{}.Validate())
	assert.Nil(t, AmountRangeBoundary1{}.Validate())
	assert.Nil(t, DateAndPeriod2Choice{}.Validate())
	assert.Nil(t, FromToAmountRange1{}.Validate())
	assert.Nil(t, FromToPercentageRange1{}.Validate())
	assert.NotNil(t, GetLimitV07{}.Validate())
	assert.Nil(t, ImpliedCurrencyAmountRange1Choice{}.Validate())
	assert.Nil(t, ImpliedCurrencyAndAmountRange1{}.Validate())
	assert.Nil(t, LimitCriteria6{}.Validate())
	assert.NotNil(t, LimitCriteria6Choice{}.Validate())
	assert.Nil(t, LimitQuery4{}.Validate())
	assert.Nil(t, LimitReturnCriteria2{}.Validate())
	assert.Nil(t, LimitSearchCriteria6{}.Validate())
	assert.NotNil(t, LimitType1Choice{}.Validate())
	assert.NotNil(t, MarketInfrastructureIdentification1Choice{}.Validate())
	assert.Nil(t, PercentageRange1Choice{}.Validate())
	assert.Nil(t, PercentageRangeBoundary1{}.Validate())
	assert.Nil(t, Period2{}.Validate())
	assert.NotNil(t, SystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmount{}.Validate())
	assert.NotNil(t, Amount2Choice{}.Validate())
	assert.Nil(t, DateAndDateTime2Choice{}.Validate())
	assert.NotNil(t, Limit8{}.Validate())
	assert.NotNil(t, LimitIdentification2Choice{}.Validate())
	assert.NotNil(t, LimitIdentification5{}.Validate())
	assert.NotNil(t, LimitIdentification6{}.Validate())
	assert.NotNil(t, LimitStructure3{}.Validate())
	assert.NotNil(t, MessageHeader1{}.Validate())
	assert.NotNil(t, ModifyLimitV07{}.Validate())
	assert.NotNil(t, DeleteLimitV07{}.Validate())
	assert.NotNil(t, LimitStructure2Choice{}.Validate())
	assert.Nil(t, BusinessDay8{}.Validate())
	assert.Nil(t, BusinessDay9{}.Validate())
	assert.Nil(t, BusinessDayReportOrError10Choice{}.Validate())
	assert.Nil(t, BusinessDayReportOrError9Choice{}.Validate())
	assert.NotNil(t, ClosureReason2Choice{}.Validate())
	assert.Nil(t, DateTimePeriod1Choice{}.Validate())
	assert.NotNil(t, ErrorHandling3Choice{}.Validate())
	assert.NotNil(t, ErrorHandling5{}.Validate())
	assert.NotNil(t, MessageHeader7{}.Validate())
	assert.NotNil(t, OriginalBusinessQuery1{}.Validate())
	assert.NotNil(t, ReturnBusinessDayInformationV07{}.Validate())
	assert.Nil(t, SystemAvailabilityAndEvents3{}.Validate())
	assert.NotNil(t, SystemClosure2{}.Validate())
	assert.NotNil(t, SystemEvent3{}.Validate())
	assert.NotNil(t, SystemEventType4Choice{}.Validate())
	assert.NotNil(t, SystemStatus2Choice{}.Validate())
	assert.NotNil(t, SystemStatus3{}.Validate())
	assert.Nil(t, TimePeriod1{}.Validate())
	assert.NotNil(t, BackupPaymentV07{}.Validate())
	assert.NotNil(t, MemberIdentification3Choice{}.Validate())
	assert.Nil(t, PaymentInstruction13{}.Validate())
	assert.NotNil(t, PaymentType4Choice{}.Validate())
	assert.NotNil(t, SystemMember3{}.Validate())
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmount{}.Validate())
	assert.Nil(t, AmendmentInformationDetails13{}.Validate())
	assert.NotNil(t, AmountType4Choice{}.Validate())
	assert.NotNil(t, Case5{}.Validate())
	assert.NotNil(t, CaseAssignment5{}.Validate())
	assert.NotNil(t, CashAccount38{}.Validate())
	assert.NotNil(t, CategoryPurpose1Choice{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification3Choice{}.Validate())
	assert.Nil(t, CreditTransferMandateData1{}.Validate())
	assert.Nil(t, CreditorReferenceInformation2{}.Validate())
	assert.NotNil(t, CreditorReferenceType1Choice{}.Validate())
	assert.NotNil(t, CreditorReferenceType2{}.Validate())
	assert.NotNil(t, DiscountAmountAndType1{}.Validate())
	assert.NotNil(t, DiscountAmountType1Choice{}.Validate())
	assert.NotNil(t, DocumentAdjustment1{}.Validate())
	assert.Nil(t, DocumentLineIdentification1{}.Validate())
	assert.Nil(t, DocumentLineInformation1{}.Validate())
	assert.NotNil(t, DocumentLineType1{}.Validate())
	assert.NotNil(t, DocumentLineType1Choice{}.Validate())
	assert.NotNil(t, EquivalentAmount2{}.Validate())
	assert.NotNil(t, Frequency36Choice{}.Validate())
	assert.NotNil(t, FrequencyAndMoment1{}.Validate())
	assert.NotNil(t, FrequencyPeriod1{}.Validate())
	assert.NotNil(t, Garnishment3{}.Validate())
	assert.NotNil(t, GarnishmentType1{}.Validate())
	assert.NotNil(t, GarnishmentType1Choice{}.Validate())
	assert.Nil(t, InstructionForAssignee1{}.Validate())
	assert.Nil(t, InstructionForCreditorAgent3{}.Validate())
	assert.Nil(t, InstructionForNextAgent1{}.Validate())
	assert.NotNil(t, LocalInstrument2Choice{}.Validate())
	assert.NotNil(t, MandateClassification1Choice{}.Validate())
	assert.Nil(t, MandateRelatedData1Choice{}.Validate())
	assert.Nil(t, MandateRelatedInformation14{}.Validate())
	assert.NotNil(t, MandateSetupReason1Choice{}.Validate())
	assert.Nil(t, MandateTypeInformation2{}.Validate())
	assert.NotNil(t, OriginalGroupInformation29{}.Validate())
	assert.Nil(t, OriginalTransactionReference31{}.Validate())
	assert.Nil(t, Party40Choice{}.Validate())
	assert.Nil(t, PaymentTypeInformation27{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.NotNil(t, ProxyAccountType1Choice{}.Validate())
	assert.NotNil(t, Purpose2Choice{}.Validate())
	assert.Nil(t, ReferredDocumentInformation7{}.Validate())
	assert.NotNil(t, ReferredDocumentType3Choice{}.Validate())
	assert.NotNil(t, ReferredDocumentType4{}.Validate())
	assert.Nil(t, RemittanceAmount2{}.Validate())
	assert.Nil(t, RemittanceAmount3{}.Validate())
	assert.Nil(t, RemittanceInformation16{}.Validate())
	assert.NotNil(t, RequestToModifyPaymentV07{}.Validate())
	assert.Nil(t, RequestedModification9{}.Validate())
	assert.NotNil(t, ServiceLevel8Choice{}.Validate())
	assert.Nil(t, SettlementDateTimeIndication1{}.Validate())
	assert.Nil(t, SettlementInstruction6{}.Validate())
	assert.NotNil(t, SettlementInstruction7{}.Validate())
	assert.Nil(t, StructuredRemittanceInformation16{}.Validate())
	assert.Nil(t, TaxAmount2{}.Validate())
	assert.NotNil(t, TaxAmountAndType1{}.Validate())
	assert.NotNil(t, TaxAmountType1Choice{}.Validate())
	assert.Nil(t, TaxAuthorisation1{}.Validate())
	assert.Nil(t, TaxInformation7{}.Validate())
	assert.Nil(t, TaxParty1{}.Validate())
	assert.Nil(t, TaxParty2{}.Validate())
	assert.Nil(t, TaxPeriod2{}.Validate())
	assert.Nil(t, TaxRecord2{}.Validate())
	assert.NotNil(t, TaxRecordDetails2{}.Validate())
	assert.NotNil(t, UnderlyingGroupInformation1{}.Validate())
	assert.NotNil(t, UnderlyingPaymentInstruction6{}.Validate())
	assert.NotNil(t, UnderlyingPaymentTransaction5{}.Validate())
	assert.Nil(t, UnderlyingStatementEntry3{}.Validate())
	assert.NotNil(t, UnderlyingTransaction6Choice{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 ExternalAccountIdentification1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalCashAccountType1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.Nil(t, type2.Validate())

	var type3 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalEnquiryRequestType1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type6 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type7 ExternalPaymentControlRequestType1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 ExternalPersonIdentification1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type9 ExternalSystemBalanceType1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.Nil(t, type9.Validate())

	var type10 ExternalMarketInfrastructure1Code
	assert.NotNil(t, type10.Validate())
	type10 = "tes"
	assert.Nil(t, type10.Validate())

	var type11 ExternalSystemErrorHandling1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.Nil(t, type11.Validate())

	var type12 ExternalSystemEventType1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())

	var type13 ExternalAgentInstruction1Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.Nil(t, type13.Validate())

	var type14 ExternalCashClearingSystem1Code
	assert.NotNil(t, type14.Validate())
	type14 = "tes"
	assert.Nil(t, type14.Validate())

	var type15 ExternalCategoryPurpose1Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.Nil(t, type15.Validate())

	var type16 ExternalCreditorAgentInstruction1Code
	assert.NotNil(t, type16.Validate())
	type16 = "test"
	assert.Nil(t, type16.Validate())

	var type17 ExternalDiscountAmountType1Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.Nil(t, type17.Validate())

	var type18 ExternalDocumentLineType1Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.Nil(t, type18.Validate())

	var type19 ExternalGarnishmentType1Code
	assert.NotNil(t, type19.Validate())
	type19 = "test"
	assert.Nil(t, type19.Validate())

	var type20 ExternalLocalInstrument1Code
	assert.NotNil(t, type20.Validate())
	type20 = "test"
	assert.Nil(t, type20.Validate())

	var type21 ExternalMandateSetupReason1Code
	assert.NotNil(t, type21.Validate())
	type21 = "test"
	assert.Nil(t, type21.Validate())

	var type22 ExternalProxyAccountType1Code
	assert.NotNil(t, type22.Validate())
	type22 = "test"
	assert.Nil(t, type22.Validate())

	var type23 ExternalPurpose1Code
	assert.NotNil(t, type23.Validate())
	type23 = "test"
	assert.Nil(t, type23.Validate())

	var type24 ExternalServiceLevel1Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.Nil(t, type24.Validate())

	var type25 ExternalTaxAmountType1Code
	assert.NotNil(t, type25.Validate())
	type25 = "test"
	assert.Nil(t, type25.Validate())

	var type26 PreferredContactMethod1Code
	assert.NotNil(t, type26.Validate())
	type26 = "test"
	assert.NotNil(t, type26.Validate())
	type26 = "CELL"
	assert.Nil(t, type26.Validate())

	var type27 QueryType2Code
	assert.NotNil(t, type27.Validate())
	type27 = "test"
	assert.NotNil(t, type27.Validate())
	type27 = "DELD"
	assert.Nil(t, type27.Validate())

	var type28 BalanceCounterparty1Code
	assert.NotNil(t, type28.Validate())
	type28 = "test"
	assert.NotNil(t, type28.Validate())
	type28 = "MULT"
	assert.Nil(t, type28.Validate())

	var type29 LimitType3Code
	assert.NotNil(t, type29.Validate())
	type29 = "test"
	assert.NotNil(t, type29.Validate())
	type29 = "MULT"
	assert.Nil(t, type29.Validate())

	var type30 SystemClosureReason1Code
	assert.NotNil(t, type30.Validate())
	type30 = "test"
	assert.NotNil(t, type30.Validate())
	type30 = "ADTW"
	assert.Nil(t, type30.Validate())

	var type31 SystemStatus2Code
	assert.NotNil(t, type31.Validate())
	type31 = "test"
	assert.NotNil(t, type31.Validate())
	type31 = "CLSG"
	assert.Nil(t, type31.Validate())

	var type32 PaymentType3Code
	assert.NotNil(t, type32.Validate())
	type32 = "test"
	assert.NotNil(t, type32.Validate())
	type32 = "CBS"
	assert.Nil(t, type32.Validate())

	var type33 ChargeBearerType1Code
	assert.NotNil(t, type33.Validate())
	type33 = "test"
	assert.NotNil(t, type33.Validate())
	type33 = "DEBT"
	assert.Nil(t, type33.Validate())

	var type34 ClearingChannel2Code
	assert.NotNil(t, type34.Validate())
	type34 = "test"
	assert.NotNil(t, type34.Validate())
	type34 = "BOOK"
	assert.Nil(t, type34.Validate())

	var type35 DocumentType3Code
	assert.NotNil(t, type35.Validate())
	type35 = "test"
	assert.NotNil(t, type35.Validate())
	type35 = "SCOR"
	assert.Nil(t, type35.Validate())

	var type36 DocumentType6Code
	assert.NotNil(t, type36.Validate())
	type36 = "test"
	assert.NotNil(t, type36.Validate())
	type36 = "MSIN"
	assert.Nil(t, type36.Validate())

	var type37 Frequency6Code
	assert.NotNil(t, type37.Validate())
	type37 = "test"
	assert.NotNil(t, type37.Validate())
	type37 = "YEAR"
	assert.Nil(t, type37.Validate())

	var type38 Instruction4Code
	assert.NotNil(t, type38.Validate())
	type38 = "test"
	assert.NotNil(t, type38.Validate())
	type38 = "TELA"
	assert.Nil(t, type38.Validate())

	var type39 PaymentMethod4Code
	assert.NotNil(t, type39.Validate())
	type39 = "test"
	assert.NotNil(t, type39.Validate())
	type39 = "DD"
	assert.Nil(t, type39.Validate())

	var type40 Priority2Code
	assert.NotNil(t, type40.Validate())
	type40 = "test"
	assert.NotNil(t, type40.Validate())
	type40 = "NORM"
	assert.Nil(t, type40.Validate())

	var type41 SequenceType3Code
	assert.NotNil(t, type41.Validate())
	type41 = "test"
	assert.NotNil(t, type41.Validate())
	type41 = "RPRE"
	assert.Nil(t, type41.Validate())

	var type42 SettlementMethod1Code
	assert.NotNil(t, type42.Validate())
	type42 = "test"
	assert.NotNil(t, type42.Validate())
	type42 = "CLRG"
	assert.Nil(t, type42.Validate())

	var type43 TaxRecordPeriod1Code
	assert.NotNil(t, type43.Validate())
	type43 = "test"
	assert.NotNil(t, type43.Validate())
	type43 = "MM01"
	assert.Nil(t, type43.Validate())
}
