package sila

import "github.com/bpancost/sila/domain"

func (client ClientImpl) CertifyBusiness(adminUserHandle string, businessHandle string) CertifyBusiness {
	return &CertifyBusinessMsg{
		Header: client.generateHeader().setUserHandle(adminUserHandle).setBusinessHandle(businessHandle),
	}
}

type CertifyBusinessMsg struct {
	Header             *Header `json:"header"`
	MemberHandle       string  `json:"member_handle"`
	CertificationToken string  `json:"certification_token"`
}

func (msg *CertifyBusinessMsg) SetRef(ref string) CertifyBusiness {
	msg.Header.setRef(ref)
	return msg
}

func (msg *CertifyBusinessMsg) Do(userWalletPrivateKey string, businessWalletPrivateKey string) (domain.SuccessResponse, error) {
	var responseBody domain.SuccessResponse
	err := instance.performCallWithUserAndBusinessAuth("/certify_business", msg, &responseBody, userWalletPrivateKey, businessWalletPrivateKey)
	return responseBody, err
}
