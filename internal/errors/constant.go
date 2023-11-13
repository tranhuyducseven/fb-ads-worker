package errors

import "errors"

var (
	ErrUserIdRequired      = errors.New("USER_ID_REQUIRED")
	ErrOrderIdRequired     = errors.New("ORDER_ID_REQUIRED")
	ErrOrderNotFound       = errors.New("ORDER_NOT_FOUND")
	ErrOrderNotCharged     = errors.New("ORDER_NOT_CHARGED")
	ErrOrderAlreadyProcess = errors.New("ORDER_ALREADY_PROCESS")
	ErrVolumeRequired      = errors.New("VOLUME_REQUIRED")
	ErrLockFailed          = errors.New("LOCK_FAILED")
	ErrUnlockFailed        = errors.New("UNLOCK_FAILED")

	ErrGetFeeWalletFailed        = errors.New("GET_ADMIN_FEE_FAILED")
	ErrCommissionWalletNotFound  = errors.New("COMMISSION_WALLET_NOT_FOUND")
	ErrCommissionWalletNotEnough = errors.New("COMMISSION_WALLET_NOT_ENOUGH")

	ErrUserWalletNotEnough = errors.New("USER_WALLET_NOT_ENOUGH")
	ErrUserWalletNotFound  = errors.New("USER_WALLET_NOT_FOUND")

	ErrStructInvalid = errors.New("STRUCT_INVALID")

	ErrUpdateTokenBalanceFailed = errors.New("UPDATE_TOKEN_BALANCE_FAILED")
	ErrCreateTokenBalanceFailed = errors.New("CREATE_TOKEN_BALANCE_FAILED")

	ErrCreateOrderFailed = errors.New("CREATE_ORDER_FAILED")
	ErrUpdateOrderFailed = errors.New("UPDATE_ORDER_FAILED")

	ErrCommitTransactionFailed = errors.New("COMMIT_TRANSACTION_FAILED")

	ErrRewardWalletNotFound  = errors.New("REWARD_WALLET_NOT_FOUND")
	ErrRewardWalletNotEnough = errors.New("REWARD_WALLET_NOT_ENOUGH")

	ErrStakeWalletNotFound  = errors.New("STAKE_WALLET_NOT_FOUND")
	ErrStakeWalletNotEnough = errors.New("STAKE_WALLET_NOT_ENOUGH")

	ErrCurrencyConvertFailed = errors.New("CURRENCY_CONVERT_FAILED")
)
