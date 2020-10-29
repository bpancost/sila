package sila

func (client ClientImpl) LinkBusinessMember(adminUserHandle string, businessHandle string) LinkBusinessMember {
	return &LinkBusinessMemberMsg{
		Header: client.generateHeader().setUserHandle(adminUserHandle).setBusinessHandle(businessHandle),
	}
}

type LinkBusinessMemberMsg struct {
	Header         *Header `json:"header"`
	MemberHandle   string  `json:"member_handle"`
	Role           string  `json:"role"`
	OwnershipStake float64 `json:"ownership_stake,omitempty"`
	Description    string  `json:"description,omitempty"`
}

func (msg *LinkBusinessMemberMsg) SetRef(ref string) LinkBusinessMember {
	msg.Header.setRef(ref)
	return msg
}

func (msg *LinkBusinessMemberMsg) SetAdminMember(newMemberHandle string) LinkBusinessMember {
	msg.MemberHandle = newMemberHandle
	msg.Role = "administrator"
	return msg
}

func (msg *LinkBusinessMemberMsg) SetControllingOfficerMember(newMemberHandle string) LinkBusinessMember {
	msg.MemberHandle = newMemberHandle
	msg.Role = "controlling_officer"
	return msg
}

func (msg *LinkBusinessMemberMsg) SetBeneficialOwnerMember(newMemberHandle string, ownershipStake float64) LinkBusinessMember {
	msg.MemberHandle = newMemberHandle
	msg.Role = "beneficial_owner"
	msg.OwnershipStake = ownershipStake
	return msg
}

func (msg *LinkBusinessMemberMsg) SetMemberDescription(description string) LinkBusinessMember {
	msg.Description = description
	return msg
}

type LinkBusinessMemberResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	Role              string                 `json:"role"`
	Details           string                 `json:"details"`
	VerificationUuid  string                 `json:"verification_uuid"`
}

func (msg *LinkBusinessMemberMsg) Do(userWalletPrivateKey string, businessWalletPrivateKey string) (LinkBusinessMemberResponse, error) {
	var responseBody LinkBusinessMemberResponse
	err := instance.performCallWithUserAndBusinessAuth("/link_business_member", msg, &responseBody, userWalletPrivateKey, businessWalletPrivateKey)
	return responseBody, err
}
