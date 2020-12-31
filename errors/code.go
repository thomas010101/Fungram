package errors

import (
	"fmt"
	"wallet/types"
)

type code int64

const (
	success code = iota
	unknown
	authorizationFailed      code = 1001
	newEquipmentLogin        code = 1002
	errAuthorizationNotFound code = 1003
	dataNotFound             code = 10000
	_                             = iota + dataNotFound
	wrongParameter
	accountAlreadyExist
	invitationCodeNotFound
	accountNameAndPwdFailed
	accountCannotLogin
	captchaTimeout
	mobileAlreadyBind
	emailAlreadyBind
	mobileAlreadyExist
	emailAlreadyExist
	googleAuthAlreadyBind
	googleAuthNodeBind
	googleAuthSecretEmpty
	googleAuthValidate
	overLimit
	accountNameAlreadyExist
	accountInvitationAlreadyExist
	realNameAlreadyApprove
	realNameAlreadyPass
	realNameCardAlreadyExist
	currencyNotFound
	sendEmailError
	registerError
	addressNotFound
	googleSecretHasExist
	reqParamsIsWrong
	createWalletErr
	updateNickLenErr
	overLimitTimes
	sendCaptchaOverLimitTimes
	comment
	getImeiErr
	captchaErr
	updatePasswordErr
	samePasswordErr
	accountIsNotExist
	addressNotAllowToLogin
	addressNotBindWithUser
	changeAddressPowerOfLoginErr
	balanceTooLowError
	balanceTooLargeError
	amountLessZero
	coinNotExist
	getHDSupportChainErr
	walletTransferErr
	changeAmountTooLittleErr
	getTransferListErr
	getTransferDetailErr
	getAddressErr
	delAddressErr
	addAddressErr
	referralCodeErr
	payPassWordErr
	bindPayPassWordErr
	verifyGoogleCodeErr
	verifyTimeStampErr
	walletINTransferErr
	walletWithdrawErr
	walletWithDrawMinErr
	operationErr
	excessMaxDebtRatioErr
	excessMaxLoanAmountErr
	settledErr
	generateSecret
	getInformationErr
	emailTheSameErr
	amountLessParameterErr
	rateErr
	getWalletBalanceErr
	logoutErr
	accountNameOrPwdErr
	customErr
	getWalletTransferGasErr
	getAuditingListErr
	getBackTransferListErr
	getAssetListErr
	getWalletAddressListErr
	updateAuditStatusErr
	getTransactionListErr
	userMasterAssetsErr
	isAuditedErr
	auditingDetailErr
	withdrawConfigErr
	editWithdrawConfigErr
	transactionInfoErr
	cannotCheckoutErr
	userIsBlockErr
	repaymentErr
	isFreezeErr
	googleSecretTimeoutErr
	applyGoogleSecretErr
	todayWithdrawErr
	addBannerErr
	updateBannerErr
	deleteBannerErr
	addRecommendErr
	updateRecommendErr
	deleteRecommendErr
	previewErr
	sameDataErr
	applicationErr
	reachedLimitErr
)

var messages = map[code]map[types.Language]string{
	success: {
		types.ZHCNLanguage: "成功",
		types.ENUSLanguage: "success",
	},
	unknown: {
		types.ZHCNLanguage: "未知错误",
		types.ENUSLanguage: "unknown error",
	},
	dataNotFound: {
		types.ZHCNLanguage: "数据不存在",
		types.ENUSLanguage: "data not found",
	},
	wrongParameter: {
		types.ZHCNLanguage: "参数错误",
		types.ENUSLanguage: "wrong parameter",
	},
	accountAlreadyExist: {
		types.ZHCNLanguage: "账号已经存在",
		types.ENUSLanguage: "the account already exist",
	},
	invitationCodeNotFound: {
		types.ZHCNLanguage: "邀请码不存在",
		types.ENUSLanguage: "the invitation code not found",
	},
	accountNameAndPwdFailed: {
		types.ZHCNLanguage: "用户名或密码错误, 如连续输错5次,则24小时后才可操作",
		types.ENUSLanguage: "account name or password failed,If you have entered five consecutive times is wrong," +
			" you can operate after 24 hours",
	},
	accountCannotLogin: {
		types.ZHCNLanguage: "该账户已被限制",
		types.ENUSLanguage: "account name or password failed",
	},
	captchaTimeout: {
		types.ZHCNLanguage: "验证码错误或过期",
		types.ENUSLanguage: "captcha err or timeout,If you have entered five consecutive times is wrong," +
			" you can operate after 24 hours",
	},
	newEquipmentLogin: {
		types.ZHCNLanguage: "新设备登录",
		types.ENUSLanguage: "new equipment login",
	},
	authorizationFailed: {
		types.ZHCNLanguage: "未登录或登录超时",
		types.ENUSLanguage: "account not login or login timeout",
	},
	mobileAlreadyBind: {
		types.ZHCNLanguage: "手机号已经绑定, 请不要重复绑定",
		types.ENUSLanguage: "The mobile already bind, Please do not repeat the binding",
	},
	emailAlreadyBind: {
		types.ZHCNLanguage: "邮箱号已经绑定, 请不要重复绑定",
		types.ENUSLanguage: "The email already bind, Please do not repeat the binding",
	},
	mobileAlreadyExist: {
		types.ZHCNLanguage: "手机号已存在",
		types.ENUSLanguage: "The mobile already exist",
	},
	emailAlreadyExist: {
		types.ZHCNLanguage: "邮箱已存在",
		types.ENUSLanguage: "The email already exist",
	},
	googleAuthAlreadyBind: {
		types.ZHCNLanguage: "您已经绑定谷歌验证",
		types.ENUSLanguage: "You have bind Google validation",
	},
	googleAuthSecretEmpty: {
		types.ZHCNLanguage: "您还没有申请谷歌验证私钥, 请先申请",
		types.ENUSLanguage: "You not have google secret, Please apple",
	},
	googleAuthNodeBind: {
		types.ZHCNLanguage: "您未绑定谷歌验证, 请先绑定",
		types.ENUSLanguage: "You not bind google auth, Please bind",
	},
	googleAuthValidate: {
		types.ZHCNLanguage: "谷歌验证失败",
		types.ENUSLanguage: "Failed google validate",
	},
	overLimit: {
		types.ZHCNLanguage: "超过限制, 不能操作",
		types.ENUSLanguage: "Over the limit, can't operation",
	},
	accountNameAlreadyExist: {
		types.ZHCNLanguage: "用户名已经存在",
		types.ENUSLanguage: "The account name already exist",
	},
	accountInvitationAlreadyExist: {
		types.ZHCNLanguage: "用户名邀请人已经存在",
		types.ENUSLanguage: "The account invitation already exist",
	},
	realNameAlreadyApprove: {
		types.ZHCNLanguage: "实名认证已经申请中, 请不要重复申请",
		types.ENUSLanguage: "In the approve, Please can't repeat",
	},
	realNameAlreadyPass: {
		types.ZHCNLanguage: "实名认证已经通过",
		types.ENUSLanguage: "The real-name authentication has been passed",
	},
	realNameCardAlreadyExist: {
		types.ZHCNLanguage: "该证件号已经认证",
		types.ENUSLanguage: "The card has been authentication",
	},
	currencyNotFound: {
		types.ZHCNLanguage: "暂不支持该币种: %s",
		types.ENUSLanguage: "currency not found: %s",
	},
	sendEmailError: {
		types.ZHCNLanguage: "发送邮件错误",
		types.ENUSLanguage: "send email error",
	},
	addressNotFound: {
		types.ZHCNLanguage: "地址不存在",
		types.ENUSLanguage: "the address not found",
	},
	registerError: {
		types.ZHCNLanguage: "用户注册失败",
		types.ENUSLanguage: "fail to register",
	},
	googleSecretHasExist: {
		types.ZHCNLanguage: "谷歌秘钥已存在, 请不要重复申请",
		types.ENUSLanguage: "google secret has exist, not repeat apply",
	},
	reqParamsIsWrong: {
		types.ZHCNLanguage: "您的输入参数有误，请重新输入",
		types.ENUSLanguage: "your request's param was wrong",
	},
	createWalletErr: {
		types.ZHCNLanguage: "创建钱包失败，请您重新创建",
		types.ENUSLanguage: "fail to create a new wallet,please try again",
	},
	errAuthorizationNotFound: {
		types.ZHCNLanguage: "请你重新登录",
		types.ENUSLanguage: "please login this website",
	},
	updateNickLenErr: {
		types.ZHCNLanguage: "昵称长度不能少于6个字符",
		types.ENUSLanguage: "nicknames cannot be less than 6 characters long",
	},
	overLimitTimes: {
		types.ZHCNLanguage: "输入错误超过5次，请于24小时之后重新输入",
		types.ENUSLanguage: "Input error more than 5 times, please re-enter at %s",
	},
	sendCaptchaOverLimitTimes: {
		types.ZHCNLanguage: "发送太频繁，请稍后再试",
		types.ENUSLanguage: "Request verification code more than 5 times, please re request after %s",
	},
	comment: {
		types.ZHCNLanguage: "提交失败",
		types.ENUSLanguage: "fail to submit",
	},
	getImeiErr: {
		types.ZHCNLanguage: "获取设备信息失败",
		types.ENUSLanguage: "Failed to get device information",
	},
	captchaErr: {
		types.ZHCNLanguage: "验证码错误",
		types.ENUSLanguage: "Verification code error",
	},
	updatePasswordErr: {
		types.ZHCNLanguage: "修改密码失败，请重新尝试",
		types.ENUSLanguage: "Password change failed，please try again",
	},
	samePasswordErr: {
		types.ZHCNLanguage: "修改前后密码一致，请重新输入密码",
		types.ENUSLanguage: "The password before and after modification is the same, please re-enter the password",
	},
	accountIsNotExist: {
		types.ZHCNLanguage: "用户不存在",
		types.ENUSLanguage: "user does not exist",
	},
	addressNotAllowToLogin: {
		types.ZHCNLanguage: "该地址不允许用来登录",
		types.ENUSLanguage: "This address is not allowed to log in",
	},
	addressNotBindWithUser: {
		types.ZHCNLanguage: "该地址未绑定用户",
		types.ENUSLanguage: "The address is not bound to a user",
	},
	changeAddressPowerOfLoginErr: {
		types.ZHCNLanguage: "更改用户是否启用地址登录失败",
		types.ENUSLanguage: "Failed to change whether to enable address login for user",
	},
	balanceTooLowError: {
		types.ZHCNLanguage: "余额不足",
		types.ENUSLanguage: "balance too low",
	},
	balanceTooLargeError: {
		types.ZHCNLanguage: "余额过大",
		types.ENUSLanguage: "balance too large",
	},
	amountLessZero: {
		types.ZHCNLanguage: "金额必须大于0",
		types.ENUSLanguage: "amount less than zero",
	},
	coinNotExist: {
		types.ZHCNLanguage: "当前币种不存在",
		types.ENUSLanguage: "coin dose not exist",
	},
	getHDSupportChainErr: {
		types.ZHCNLanguage: "获取钱包支持链列表失败",
		types.ENUSLanguage: "Failed to get wallet support chain list",
	},
	walletTransferErr: {
		types.ZHCNLanguage: "转账失败",
		types.ENUSLanguage: "Failed to transfer",
	},
	changeAmountTooLittleErr: {
		types.ZHCNLanguage: "转账金额过少,必须大于%d",
		types.ENUSLanguage: "Too little transfer amount,Must be greater than %d ",
	},
	getTransferListErr: {
		types.ZHCNLanguage: "获取交易清单失败",
		types.ENUSLanguage: "Failed to obtain transaction list of master account",
	},
	getTransferDetailErr: {
		types.ZHCNLanguage: "获取交易详情失败",
		types.ENUSLanguage: "Failed to obtain transaction detail of master account",
	},
	getAddressErr: {
		types.ZHCNLanguage: "获取地址簿失败",
		types.ENUSLanguage: "Failed to get address book",
	},
	delAddressErr: {
		types.ZHCNLanguage: "删除地址簿失败",
		types.ENUSLanguage: "Failed to del address book",
	},
	addAddressErr: {
		types.ZHCNLanguage: "增加地址簿失败",
		types.ENUSLanguage: "Failed to add address",
	},
	referralCodeErr: {
		types.ZHCNLanguage: "邀请码错误",
		types.ENUSLanguage: "referralCode is wrong",
	},
	payPassWordErr: {
		types.ZHCNLanguage: "支付密码错误,连续输错5次,则24小时后才可操作",
		types.ENUSLanguage: "payPassWord is wrong,If you have entered five consecutive times is wrong," +
			" you can operate after 24 hours",
	},
	bindPayPassWordErr: {
		types.ZHCNLanguage: "请先绑定邮箱或者谷歌验证码",
		types.ENUSLanguage: "please bind email or google code",
	},
	verifyGoogleCodeErr: {
		types.ZHCNLanguage: "谷歌验证码错误或过期",
		types.ENUSLanguage: "google code error or expiration",
	},
	verifyTimeStampErr: {
		types.ZHCNLanguage: "时间戳错误",
		types.ENUSLanguage: "timestamp is wrong",
	},
	walletINTransferErr: {
		types.ZHCNLanguage: "转账失败",
		types.ENUSLanguage: "Failed to transfer",
	},
	walletWithdrawErr: {
		types.ZHCNLanguage: "提现失败",
		types.ENUSLanguage: "Failed to withdraw",
	},
	walletWithDrawMinErr: {
		types.ZHCNLanguage: "提现失败,转账数额必须大于%s",
		types.ENUSLanguage: "Withdrawal failed, transfer amount must be greater than %s",
	},
	operationErr: {
		types.ZHCNLanguage: "操作失败",
		types.ENUSLanguage: "the operation failure",
	},
	excessMaxDebtRatioErr: {
		types.ZHCNLanguage: "当前您的负债率已超过150%",
		types.ENUSLanguage: "excess max debt ratio 150%",
	},
	excessMaxLoanAmountErr: {
		types.ZHCNLanguage: "您剩余最大可借贷:%s",
		types.ENUSLanguage: "you excess max loan:%s",
	},
	settledErr: {
		types.ZHCNLanguage: "已结算",
		types.ENUSLanguage: "has been settled",
	},
	generateSecret: {
		types.ZHCNLanguage: "生成谷歌秘钥出错",
		types.ENUSLanguage: "generate secret error",
	},
	getInformationErr: {
		types.ZHCNLanguage: "获取信息失败，请刷新页面",
		types.ENUSLanguage: "fail to get information,please try again",
	},
	emailTheSameErr: {
		types.ZHCNLanguage: "当前邮箱一样",
		types.ENUSLanguage: "email is the same",
	},
	amountLessParameterErr: {
		types.ZHCNLanguage: "金额必须大于等于:%s",
		types.ENUSLanguage: "amount great than or equal to:%s",
	},
	rateErr: {
		types.ZHCNLanguage: "汇率出错",
		types.ENUSLanguage: "rate error",
	},
	getWalletBalanceErr: {
		types.ZHCNLanguage: "获取云钱包余额出错",
		types.ENUSLanguage: "Error in obtaining cloud wallet balance",
	},
	logoutErr: {
		types.ZHCNLanguage: "退出失败",
		types.ENUSLanguage: "logout error",
	},
	accountNameOrPwdErr: {
		types.ZHCNLanguage: "账户或密码错误",
		types.ENUSLanguage: "account or passWord error",
	},
	customErr: {
		types.ZHCNLanguage: "%s",
		types.ENUSLanguage: "%s",
	},
	getWalletTransferGasErr: {
		types.ZHCNLanguage: "获取手续费失败",
		types.ENUSLanguage: "failed to obtain service charge",
	},
	getAuditingListErr: {
		types.ZHCNLanguage: "查询审核列表错误",
		types.ENUSLanguage: "get auditing list error",
	},
	getBackTransferListErr: {
		types.ZHCNLanguage: "获取转账列表错误",
		types.ENUSLanguage: "get transfer list error",
	},
	getAssetListErr: {
		types.ZHCNLanguage: "获取用户资产列表错误",
		types.ENUSLanguage: "get assets list error",
	},
	getWalletAddressListErr: {
		types.ZHCNLanguage: "获取去中心化钱包地址列表错误",
		types.ENUSLanguage: "get wallet address list error",
	},
	updateAuditStatusErr: {
		types.ZHCNLanguage: "更新审核状态错误",
		types.ENUSLanguage: "update audit status error",
	},
	getTransactionListErr: {
		types.ZHCNLanguage: "获取交易列表异常",
		types.ENUSLanguage: "get transaction list error",
	},
	userMasterAssetsErr: {
		types.ZHCNLanguage: "获取用户主账户资产异常",
		types.ENUSLanguage: "get user master asset error",
	},
	isAuditedErr: {
		types.ZHCNLanguage: "审核已处理",
		types.ENUSLanguage: "auditing has handled",
	},
	auditingDetailErr: {
		types.ZHCNLanguage: "获取审核详情异常",
		types.ENUSLanguage: "get auditing info error",
	},
	withdrawConfigErr: {
		types.ZHCNLanguage: "查询提现最小值设定异常",
		types.ENUSLanguage: "get withdraw config error",
	},
	editWithdrawConfigErr: {
		types.ZHCNLanguage: "修改提现最小值异常",
		types.ENUSLanguage: "edit withdraw min error",
	},
	transactionInfoErr: {
		types.ZHCNLanguage: "查询交易详情异常",
		types.ENUSLanguage: "get transaction info error",
	},
	cannotCheckoutErr: {
		types.ZHCNLanguage: "没有权限查看接口",
		types.ENUSLanguage: "cannot checkout interface error",
	},
	userIsBlockErr: {
		types.ZHCNLanguage: "用户被禁止使用",
		types.ENUSLanguage: "user ",
	},
	repaymentErr: {
		types.ZHCNLanguage: "24小时之后才能还款",
		types.ENUSLanguage: "after 24 hours pay it back",
	},
	isFreezeErr: {
		types.ZHCNLanguage: "资金冻结",
		types.ENUSLanguage: "money is frozen",
	},
	googleSecretTimeoutErr: {
		types.ZHCNLanguage: "谷歌秘钥过期或错误",
		types.ENUSLanguage: "google secret key expired or error",
	},
	applyGoogleSecretErr: {
		types.ZHCNLanguage: "申请谷歌秘钥失败",
		types.ENUSLanguage: "failed to apply for Google secret key",
	},
	todayWithdrawErr: {
		types.ZHCNLanguage: "查询今日提现异常",
		types.ENUSLanguage: "getTodayWithdraw error",
	},
	addBannerErr: {
		types.ZHCNLanguage: "新增广告异常",
		types.ENUSLanguage: "add banner error",
	},
	updateBannerErr: {
		types.ZHCNLanguage: "编辑广告异常",
		types.ENUSLanguage: "update banner error",
	},
	deleteBannerErr: {
		types.ZHCNLanguage: "删除广告异常",
		types.ENUSLanguage: "delete banner error",
	},
	addRecommendErr: {
		types.ZHCNLanguage: "新增推荐异常",
		types.ENUSLanguage: "update recommend error",
	},
	updateRecommendErr: {
		types.ZHCNLanguage: "编辑推荐异常",
		types.ENUSLanguage: "update recommend error",
	},
	deleteRecommendErr: {
		types.ZHCNLanguage: "删除推荐异常",
		types.ENUSLanguage: "delete recommend error",
	},
	previewErr: {
		types.ZHCNLanguage: "查看预览异常",
		types.ENUSLanguage: "get preview error",
	},
	sameDataErr: {
		types.ZHCNLanguage: "已有相同数据",
		types.ENUSLanguage: "has same data error",
	},
	applicationErr: {
		types.ZHCNLanguage: "应用异常",
		types.ENUSLanguage: "application error",
	},
	reachedLimitErr: {
		types.ZHCNLanguage: "已达上限",
		types.ENUSLanguage: "reached limit error",
	},
}

func (c code) msg(language types.Language, params []interface{}) string {
	s := c.origin(language)
	if len(params) == 0 {
		return s
	}
	return fmt.Sprintf(s, params...)
}

func (c code) origin(language types.Language) string {
	if m, ok := messages[c]; ok {
		return m[language]
	}
	return messages[unknown][language]
}

func (c code) isNeedData() bool {
	switch c {
	case newEquipmentLogin, success:
		return true
	}
	return false
}
