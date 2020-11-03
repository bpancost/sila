package sila

import "github.com/bpancost/sila/domain"

type CheckHandle interface {
	SetRef(ref string) CheckHandle
	Do() (domain.SuccessResponse, error)
}

type Register interface {
	SetRef(ref string) Register
	SetAddress(address domain.RegistrationAddress) Register
	SetIdentity(identityType domain.IdentityType, identityValue string) Register
	SetContact(contactAlias string, phone string, email string) Register
	SetCrypto(nickname string, address string) Register
	SetIndividualEntity(firstName string, lastName string, birthDate string) Register
	SetBusinessEntity(entityName string, businessType string, naicsCode int) Register
	SetBusinessWebsite(businessWebsite string) Register
	SetDoingBusinessAs(dba string) Register
	Do() (domain.SuccessResponse, error)
}

type RequestKyc interface {
	SetRef(ref string) RequestKyc
	SetKycLevel(kycLevel string) RequestKyc
	Do(userWalletPrivateKey string) (domain.RequestKycResponse, error)
}

type CheckKyc interface {
	SetRef(ref string) CheckKyc
	SetKycLevel(kycLevel string) CheckKyc
	Do(userWalletPrivateKey string) (domain.CheckKycResponse, error)
}

type GetEntity interface {
	Do(userWalletPrivateKey string) (domain.GetEntityResponse, error)
}

type GetEntities interface {
	SetEntityType(entityType string) GetEntities
	SetPage(page int32) GetEntities
	SetPerPage(perPage int32) GetEntities
	Do() (domain.GetEntitiesResponse, error)
}

type LinkBusinessMember interface {
	SetAdminMemberAsAdmin(newMemberHandle string) LinkBusinessMember
	SetAdminMember() LinkBusinessMember
	SetControllingOfficerMemberAsAdmin(newMemberHandle string) LinkBusinessMember
	SetControllingOfficerMember() LinkBusinessMember
	SetBeneficialOwnerMemberAsAdmin(newMemberHandle string, ownershipStake float64) LinkBusinessMember
	SetBeneficialOwnerMember(ownershipStake float64) LinkBusinessMember
	SetMemberDescription(description string) LinkBusinessMember
	Do(userWalletPrivateKey string, businessWalletPrivateKey string) (domain.LinkBusinessMemberResponse, error)
}

type UnlinkBusinessMember interface {
	SetAdminRole() UnlinkBusinessMember
	SetBeneficialOwnerRole() UnlinkBusinessMember
	SetControllingOfficerRole() UnlinkBusinessMember
	Do(userWalletPrivateKey string, businessWalletPrivateKey string) (domain.UnlinkBusinessMemberResponse, error)
}

type CertifyBusiness interface {
	Do(userWalletPrivateKey string, businessWalletPrivateKey string) (domain.SuccessResponse, error)
}

type CertifyBeneficialOwner interface {
	SetCertificationToken(userHandleToCertify string, certificationToken string) CertifyBeneficialOwner
	Do(userWalletPrivateKey string, businessWalletPrivateKey string) (domain.SuccessResponse, error)
}

type AddRegistrationData interface {
	SetEmail(email string) AddRegistrationData
	SetPhone(phone string) AddRegistrationData
	SetIdentity(identityAlias string, identityValue string) AddRegistrationData
	SetAddress(address domain.RegistrationAddress) AddRegistrationData
	Do(userWalletPrivateKey string) (domain.ModifyRegistrationDataResponse, error)
}

type UpdateRegistrationData interface {
	SetEmail(emailUuid string, email string) UpdateRegistrationData
	SetPhone(phoneUuid string, phone string) UpdateRegistrationData
	SetIdentity(identityUuid string, identityAlias string, identityValue string) UpdateRegistrationData
	SetAddress(addressUuid string, address domain.RegistrationAddress) UpdateRegistrationData
	SetIndividualEntity(firstName string, lastName string, fullName string, birthDate string) UpdateRegistrationData
	SetBusinessEntity(businessName string, startDate string, businessType string, naicsCode int, doingBusinessAs string, businessWebsite string) UpdateRegistrationData
	Do(userWalletPrivateKey string) (domain.ModifyRegistrationDataResponse, error)
}

type DeleteRegistrationData interface {
	SetEmail(emailUuid string) DeleteRegistrationData
	SetPhone(phoneUuid string) DeleteRegistrationData
	SetIdentity(identityUuid string) DeleteRegistrationData
	SetAddress(addressUuid string) DeleteRegistrationData
	Do(userWalletPrivateKey string) (domain.SuccessResponse, error)
}
