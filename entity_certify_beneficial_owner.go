package sila

import "github.com/bpancost/sila/domain"

func (client ClientImpl) CertifyBeneficialOwner(adminUserHandle string, businessHandle string) CertifyBeneficialOwner {
	return &CertifyBeneficialOwnerMsg{
		Header: client.generateHeader().setUserHandle(adminUserHandle).setBusinessHandle(businessHandle),
	}
}

type CertifyBeneficialOwnerMsg struct {
	Header *Header `json:"header"`
}

func (msg *CertifyBeneficialOwnerMsg) SetRef(ref string) CertifyBeneficialOwner {
	msg.Header.setRef(ref)
	return msg
}

func (msg *CertifyBeneficialOwnerMsg) Do(userWalletPrivateKey string, businessWalletPrivateKey string) (domain.SuccessResponse, error) {
	var responseBody domain.SuccessResponse
	err := instance.performCallWithUserAndBusinessAuth("/certify_beneficial_owner", msg, &responseBody, userWalletPrivateKey, businessWalletPrivateKey)
	return responseBody, err
}
