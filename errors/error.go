package errors

import (
	"errors"
	"fmt"
	"wallet/types"
)

var (
	Success                          = newError(success)
	UnknownErr                       = newError(unknown)
	NewEquipmentLoginErr             = newError(newEquipmentLogin)
	DataNotFoundErr                  = newError(dataNotFound)
	WrongParameterErr                = newError(wrongParameter)
	AccountAlreadyExistErr           = newError(accountAlreadyExist)
	InvitationCodeNotFoundErr        = newError(invitationCodeNotFound)
	AccountNameAndPwdErr             = newError(accountNameAndPwdFailed)
	AccountCannotLoginErr            = newError(accountCannotLogin)
	CaptchaTimeoutErr                = newError(captchaTimeout)
	AuthorizationFailedErr           = newError(authorizationFailed)
	MobileAlreadyBindErr             = newError(mobileAlreadyBind)
	EmailAlreadyBindErr              = newError(emailAlreadyBind)
	MobileAlreadyExistErr            = newError(mobileAlreadyExist)
	EmailAlreadyExistErr             = newError(emailAlreadyExist)
	GoogleAuthAlreadyBindErr         = newError(googleAuthAlreadyBind)
	GoogleAuthSecretEmptyErr         = newError(googleAuthSecretEmpty)
	GoogleAuthNodeBindErr            = newError(googleAuthNodeBind)
	GoogleAuthValidateErr            = newError(googleAuthValidate)
	OverLimitErr                     = newError(overLimit)
	AccountNameAlreadyExistErr       = newError(accountNameAlreadyExist)
	AccountInvitationAlreadyExistErr = newError(accountInvitationAlreadyExist)
	RealNameAlreadyApproveErr        = newError(realNameAlreadyApprove)
	RealNameAlreadyPassErr           = newError(realNameAlreadyPass)
	RealNameCardAlreadyExistErr      = newError(realNameCardAlreadyExist)
	CurrencyNotFoundErr              = newError(currencyNotFound)
	SendEmailError                   = newError(sendEmailError)
	RegisterError                    = newError(registerError)
	AddressNotFound                  = newError(addressNotFound)
	GoogleSecretHasExist             = newError(googleSecretHasExist)
	ReqParamsIsWrong                 = newError(reqParamsIsWrong)
	CreateWalletErr                  = newError(createWalletErr)

	ErrAuthorizationNotFound     = newError(errAuthorizationNotFound)
	UpdateNickLen                = newError(updateNickLenErr)
	OverLimitTimes               = newError(overLimitTimes)
	SendCaptchaOverLimitTimes    = newError(sendCaptchaOverLimitTimes)
	CaptchaErr                   = newError(captchaErr)
	CommentError                 = newError(comment)
	GetImeiErr                   = newError(getImeiErr)
	UpdatePasswordErr            = newError(updatePasswordErr)
	SamePasswordErr              = newError(samePasswordErr)
	AccountIsNotExist            = newError(accountIsNotExist)
	AddressNotAllowToLogin       = newError(addressNotAllowToLogin)
	AddressNotBindWithUser       = newError(addressNotBindWithUser)
	ChangeAddressPowerOfLoginErr = newError(changeAddressPowerOfLoginErr)
	BalanceTooLow                = newError(balanceTooLowError)
	BalanceTooLarge              = newError(balanceTooLargeError)
	AmountLessZero               = newError(amountLessZero)
	CoinNotExist                 = newError(coinNotExist)
	GetHDSupportChainErr         = newError(getHDSupportChainErr)
	WalletTransferErr            = newError(walletTransferErr)
	ChangeAmountTooLittleErr     = newError(changeAmountTooLittleErr)
	GetTransferListErr           = newError(getTransferListErr)
	GetTransferDetailErr         = newError(getTransferDetailErr)
	GetAddressErr                = newError(getAddressErr)
	DelAddressErr                = newError(delAddressErr)
	AddAddressErr                = newError(addAddressErr)
	ReferralCodeErr              = newError(referralCodeErr)
	PayPassWordErr               = newError(payPassWordErr)
	BindPayPassWordErr           = newError(bindPayPassWordErr)
	VerifyGoogleCodeErr          = newError(verifyGoogleCodeErr)
	WalletINTransferErr          = newError(walletINTransferErr)
	VerifyTimeStampErr           = newError(verifyTimeStampErr)
	WalletWithdrawErr            = newError(walletWithdrawErr)
	WalletWithDrawMinErr         = newError(walletWithDrawMinErr)
	OperationErr                 = newError(operationErr)
	ExcessMaxDebtRatioErr        = newError(excessMaxDebtRatioErr)
	ExcessMaxLoanAmountErr       = newError(excessMaxLoanAmountErr)
	SettledErr                   = newError(settledErr)
	GenerateSecret               = newError(generateSecret)
	GetInformationErr            = newError(getInformationErr)
	EmailTheSameErr              = newError(emailTheSameErr)
	AmountLessParameterErr       = newError(amountLessParameterErr)
	RateErr                      = newError(rateErr)
	GetWalletBalanceErr          = newError(getWalletBalanceErr)
	LogoutErr                    = newError(logoutErr)
	AccountNameOrPwdErr          = newError(accountNameOrPwdErr)
	CustomErr                    = newError(customErr)
	GetWalletTransferGasErr      = newError(getWalletTransferGasErr)
	GetAuditingListErr           = newError(getAuditingListErr)
	GetBackTransferListErr       = newError(getBackTransferListErr)
	GetAssetListErr              = newError(getAssetListErr)
	GetWalletAddressListErr      = newError(getWalletAddressListErr)
	UpdateAuditStatusErr         = newError(updateAuditStatusErr)
	GetTransactionListErr        = newError(getTransactionListErr)
	UserMasterAssetsErr          = newError(userMasterAssetsErr)
	IsAuditedErr                 = newError(isAuditedErr)
	AuditingDetailErr            = newError(auditingDetailErr)
	WithdrawConfigErr            = newError(withdrawConfigErr)
	EditWithdrawConfigErr        = newError(editWithdrawConfigErr)
	TransactionInfoErr           = newError(transactionInfoErr)
	CannotCheckoutErr            = newError(cannotCheckoutErr)
	UserIsBlock                  = newError(userIsBlockErr)
	RepaymentErr                 = newError(repaymentErr)
	IsFreezeErr                  = newError(isFreezeErr)
	GoogleSecretTimeoutErr       = newError(googleSecretTimeoutErr)
	ApplyGoogleSecretErr       = newError(applyGoogleSecretErr)
	TodayWithdrawErr             = newError(todayWithdrawErr)
	AddBannerErr                 = newError(addBannerErr)
	UpdateBannerErr              = newError(updateBannerErr)
	DeleteBannerErr              = newError(deleteBannerErr)
	AddRecommendErr              = newError(addRecommendErr)
	UpdateRecommendErr           = newError(updateRecommendErr)
	DeleteRecommendErr           = newError(deleteRecommendErr)
	PreviewErr                   = newError(previewErr)
	SameDataErr                  = newError(sameDataErr)
	ApplicationErr               = newError(applicationErr)
	ReachedLimitErr              = newError(reachedLimitErr)
)

func newError(c code) *mError {
	return &mError{
		code: c,
	}
}

type mError struct {
	code  code
	paras []interface{}
	err   error
}

func (e *mError) Error() string {
	if e.err == nil {
		return e.code.msg(types.DefaultLanguage, e.paras)
	}
	return e.err.Error()
}

func (e *mError) Code() int64 {
	return int64(e.code)
}

func (e *mError) Data(data interface{}) interface{} {
	if e.code.isNeedData() && data != nil {
		return data
	} else {
		return nil
	}
}

func (e *mError) Msg(language types.Language) string {
	return e.code.msg(language, e.paras)
}

func (e *mError) With(param ...interface{}) *mError {
	e.paras = make([]interface{}, 0, len(param))
	for _, v := range param {
		e.paras = append(e.paras, v)
	}
	return e
}

func (e *mError) Format(format string, a ...interface{}) *mError {
	e.err = fmt.Errorf(format, a...)
	return e
}

func As(err error) *mError {
	if err == nil {
		return Success
	}
	var e *mError
	if errors.As(err, &e) {
		return e
	}
	return UnknownErr
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}
