// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package starkwarePerpetuals

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StarkwarePerpetualsMetaData contains all meta data concerning the StarkwarePerpetuals contract.
var StarkwarePerpetualsMetaData = &bind.MetaData{
	ABI: "[{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"UPGRADE_DELAY_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"configurationHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globalConfigurationHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"acceptedGovernor\",\"type\":\"address\"}],\"name\":\"LogNewGovernorAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nominatedGovernor\",\"type\":\"address\"}],\"name\":\"LogNominatedGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogNominationCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entryId\",\"type\":\"string\"}],\"name\":\"LogRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entryId\",\"type\":\"string\"}],\"name\":\"LogRemovalIntent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entryId\",\"type\":\"string\"}],\"name\":\"LogRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"removedGovernor\",\"type\":\"address\"}],\"name\":\"LogRemovedGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogUnFrozen\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEPOSIT_CANCEL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FREEZE_GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAIN_GOVERNANCE_INFO_TAG\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_FORCED_ACTIONS_REQS_PER_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_VERIFIER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"UNFREEZE_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERIFIER_REMOVAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"announceAvailabilityVerifierRemovalIntent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"announceVerifierRemovalIntent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRegisteredAvailabilityVerifiers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_verifers\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRegisteredVerifiers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_verifers\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"name\":\"isAvailabilityVerifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"name\":\"isVerifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"mainAcceptGovernance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"mainCancelNomination\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"testGovernor\",\"type\":\"address\"}],\"name\":\"mainIsGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"mainNominateNewGovernor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"governorForRemoval\",\"type\":\"address\"}],\"name\":\"mainRemoveGovernor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"name\":\"registerAvailabilityVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"name\":\"registerVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"removeAvailabilityVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"removeVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unFreeze\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositorEthKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogDepositCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDepositCancelReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogDepositNftCancelReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogMintWithdrawalPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogMintableWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositorEthKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogNftDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogNftWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"LogNftWithdrawalPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"LogSystemAssetType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAdmin\",\"type\":\"address\"}],\"name\":\"LogTokenAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAdmin\",\"type\":\"address\"}],\"name\":\"LogTokenAdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"}],\"name\":\"LogTokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAdmin\",\"type\":\"address\"}],\"name\":\"LogUserAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAdmin\",\"type\":\"address\"}],\"name\":\"LogUserAdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"LogUserRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"LogWithdrawalPerformed\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositCancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositNft\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositNftReclaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositReclaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"getAssetInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getCancellationRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"request\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getDepositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"}],\"name\":\"getEthKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getFullWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getQuantizedDepositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"presumedAssetType\",\"type\":\"uint256\"}],\"name\":\"getQuantum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quantum\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSystemAssetType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"testedAdmin\",\"type\":\"address\"}],\"name\":\"isTokenAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"testedAdmin\",\"type\":\"address\"}],\"name\":\"isUserAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"}],\"name\":\"registerSystemAssetType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"registerToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registerToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"registerTokenAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"registerUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"registerUserAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"}],\"name\":\"unregisterTokenAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"}],\"name\":\"unregisterUserAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"mintingBlob\",\"type\":\"bytes\"}],\"name\":\"withdrawAndMint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawNft\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawNftTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configHash\",\"type\":\"bytes32\"}],\"name\":\"LogGlobalConfiguration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"configItem\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configHash\",\"type\":\"bytes32\"}],\"name\":\"LogItemConfiguration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"LogOperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"LogOperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateTransitionFact\",\"type\":\"bytes32\"}],\"name\":\"LogStateTransitionFact\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"name\":\"LogUpdateState\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"escape\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKeyA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"starkKeyB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultIdA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultIdB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"syntheticAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountSynthetic\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"aIsBuyingSynthetic\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getForcedTradeRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"getForcedWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastBatchId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOrderRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOrderTreeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVaultRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVaultTreeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"testedOperator\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"registerOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"configHash\",\"type\":\"bytes32\"}],\"name\":\"setAssetConfiguration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"configHash\",\"type\":\"bytes32\"}],\"name\":\"setGlobalConfiguration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"removedOperator\",\"type\":\"address\"}],\"name\":\"unregisterOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"programOutput\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"applicationData\",\"type\":\"uint256[]\"}],\"name\":\"updateState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKeyA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKeyB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultIdA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultIdB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"syntheticAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountCollateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSynthetic\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"aIsBuyingSynthetic\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"LogForcedTradeRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogForcedWithdrawalRequest\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKeyA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"starkKeyB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultIdA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultIdB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"syntheticAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountSynthetic\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"aIsBuyingSynthetic\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submissionExpirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"premiumCost\",\"type\":\"bool\"}],\"name\":\"forcedTradeRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"premiumCost\",\"type\":\"bool\"}],\"name\":\"forcedWithdrawalRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKeyA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"starkKeyB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultIdA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultIdB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"syntheticAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountSynthetic\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"aIsBuyingSynthetic\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"freezeRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"freezeRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StarkwarePerpetualsABI is the input ABI used to generate the binding from.
// Deprecated: Use StarkwarePerpetualsMetaData.ABI instead.
var StarkwarePerpetualsABI = StarkwarePerpetualsMetaData.ABI

// StarkwarePerpetuals is an auto generated Go binding around an Ethereum contract.
type StarkwarePerpetuals struct {
	StarkwarePerpetualsCaller     // Read-only binding to the contract
	StarkwarePerpetualsTransactor // Write-only binding to the contract
	StarkwarePerpetualsFilterer   // Log filterer for contract events
}

// StarkwarePerpetualsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StarkwarePerpetualsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StarkwarePerpetualsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StarkwarePerpetualsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StarkwarePerpetualsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StarkwarePerpetualsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StarkwarePerpetualsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StarkwarePerpetualsSession struct {
	Contract     *StarkwarePerpetuals // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StarkwarePerpetualsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StarkwarePerpetualsCallerSession struct {
	Contract *StarkwarePerpetualsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// StarkwarePerpetualsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StarkwarePerpetualsTransactorSession struct {
	Contract     *StarkwarePerpetualsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// StarkwarePerpetualsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StarkwarePerpetualsRaw struct {
	Contract *StarkwarePerpetuals // Generic contract binding to access the raw methods on
}

// StarkwarePerpetualsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StarkwarePerpetualsCallerRaw struct {
	Contract *StarkwarePerpetualsCaller // Generic read-only contract binding to access the raw methods on
}

// StarkwarePerpetualsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StarkwarePerpetualsTransactorRaw struct {
	Contract *StarkwarePerpetualsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStarkwarePerpetuals creates a new instance of StarkwarePerpetuals, bound to a specific deployed contract.
func NewStarkwarePerpetuals(address common.Address, backend bind.ContractBackend) (*StarkwarePerpetuals, error) {
	contract, err := bindStarkwarePerpetuals(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetuals{StarkwarePerpetualsCaller: StarkwarePerpetualsCaller{contract: contract}, StarkwarePerpetualsTransactor: StarkwarePerpetualsTransactor{contract: contract}, StarkwarePerpetualsFilterer: StarkwarePerpetualsFilterer{contract: contract}}, nil
}

// NewStarkwarePerpetualsCaller creates a new read-only instance of StarkwarePerpetuals, bound to a specific deployed contract.
func NewStarkwarePerpetualsCaller(address common.Address, caller bind.ContractCaller) (*StarkwarePerpetualsCaller, error) {
	contract, err := bindStarkwarePerpetuals(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsCaller{contract: contract}, nil
}

// NewStarkwarePerpetualsTransactor creates a new write-only instance of StarkwarePerpetuals, bound to a specific deployed contract.
func NewStarkwarePerpetualsTransactor(address common.Address, transactor bind.ContractTransactor) (*StarkwarePerpetualsTransactor, error) {
	contract, err := bindStarkwarePerpetuals(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsTransactor{contract: contract}, nil
}

// NewStarkwarePerpetualsFilterer creates a new log filterer instance of StarkwarePerpetuals, bound to a specific deployed contract.
func NewStarkwarePerpetualsFilterer(address common.Address, filterer bind.ContractFilterer) (*StarkwarePerpetualsFilterer, error) {
	contract, err := bindStarkwarePerpetuals(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsFilterer{contract: contract}, nil
}

// bindStarkwarePerpetuals binds a generic wrapper to an already deployed contract.
func bindStarkwarePerpetuals(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StarkwarePerpetualsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StarkwarePerpetuals *StarkwarePerpetualsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StarkwarePerpetuals.Contract.StarkwarePerpetualsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StarkwarePerpetuals *StarkwarePerpetualsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.StarkwarePerpetualsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StarkwarePerpetuals *StarkwarePerpetualsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.StarkwarePerpetualsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StarkwarePerpetuals.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) DEPOSITCANCELDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "DEPOSIT_CANCEL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.DEPOSITCANCELDELAY(&_StarkwarePerpetuals.CallOpts)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.DEPOSITCANCELDELAY(&_StarkwarePerpetuals.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) FREEZEGRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "FREEZE_GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.FREEZEGRACEPERIOD(&_StarkwarePerpetuals.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.FREEZEGRACEPERIOD(&_StarkwarePerpetuals.CallOpts)
}

// MAINGOVERNANCEINFOTAG is a free data retrieval call binding the contract method 0xc23b60ef.
//
// Solidity: function MAIN_GOVERNANCE_INFO_TAG() view returns(string)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) MAINGOVERNANCEINFOTAG(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "MAIN_GOVERNANCE_INFO_TAG")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MAINGOVERNANCEINFOTAG is a free data retrieval call binding the contract method 0xc23b60ef.
//
// Solidity: function MAIN_GOVERNANCE_INFO_TAG() view returns(string)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) MAINGOVERNANCEINFOTAG() (string, error) {
	return _StarkwarePerpetuals.Contract.MAINGOVERNANCEINFOTAG(&_StarkwarePerpetuals.CallOpts)
}

// MAINGOVERNANCEINFOTAG is a free data retrieval call binding the contract method 0xc23b60ef.
//
// Solidity: function MAIN_GOVERNANCE_INFO_TAG() view returns(string)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) MAINGOVERNANCEINFOTAG() (string, error) {
	return _StarkwarePerpetuals.Contract.MAINGOVERNANCEINFOTAG(&_StarkwarePerpetuals.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) MAXFORCEDACTIONSREQSPERBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "MAX_FORCED_ACTIONS_REQS_PER_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_StarkwarePerpetuals.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_StarkwarePerpetuals.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) MAXVERIFIERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "MAX_VERIFIER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.MAXVERIFIERCOUNT(&_StarkwarePerpetuals.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.MAXVERIFIERCOUNT(&_StarkwarePerpetuals.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) UNFREEZEDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "UNFREEZE_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) UNFREEZEDELAY() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.UNFREEZEDELAY(&_StarkwarePerpetuals.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) UNFREEZEDELAY() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.UNFREEZEDELAY(&_StarkwarePerpetuals.CallOpts)
}

// UPGRADEDELAYSLOT is a free data retrieval call binding the contract method 0x20cea94d.
//
// Solidity: function UPGRADE_DELAY_SLOT() view returns(bytes32)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) UPGRADEDELAYSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "UPGRADE_DELAY_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADEDELAYSLOT is a free data retrieval call binding the contract method 0x20cea94d.
//
// Solidity: function UPGRADE_DELAY_SLOT() view returns(bytes32)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) UPGRADEDELAYSLOT() ([32]byte, error) {
	return _StarkwarePerpetuals.Contract.UPGRADEDELAYSLOT(&_StarkwarePerpetuals.CallOpts)
}

// UPGRADEDELAYSLOT is a free data retrieval call binding the contract method 0x20cea94d.
//
// Solidity: function UPGRADE_DELAY_SLOT() view returns(bytes32)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) UPGRADEDELAYSLOT() ([32]byte, error) {
	return _StarkwarePerpetuals.Contract.UPGRADEDELAYSLOT(&_StarkwarePerpetuals.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) VERIFIERREMOVALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "VERIFIER_REMOVAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.VERIFIERREMOVALDELAY(&_StarkwarePerpetuals.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.VERIFIERREMOVALDELAY(&_StarkwarePerpetuals.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) VERSION() (string, error) {
	return _StarkwarePerpetuals.Contract.VERSION(&_StarkwarePerpetuals.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) VERSION() (string, error) {
	return _StarkwarePerpetuals.Contract.VERSION(&_StarkwarePerpetuals.CallOpts)
}

// ConfigurationHash is a free data retrieval call binding the contract method 0xf2011f66.
//
// Solidity: function configurationHash(uint256 ) view returns(bytes32)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) ConfigurationHash(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "configurationHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConfigurationHash is a free data retrieval call binding the contract method 0xf2011f66.
//
// Solidity: function configurationHash(uint256 ) view returns(bytes32)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) ConfigurationHash(arg0 *big.Int) ([32]byte, error) {
	return _StarkwarePerpetuals.Contract.ConfigurationHash(&_StarkwarePerpetuals.CallOpts, arg0)
}

// ConfigurationHash is a free data retrieval call binding the contract method 0xf2011f66.
//
// Solidity: function configurationHash(uint256 ) view returns(bytes32)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) ConfigurationHash(arg0 *big.Int) ([32]byte, error) {
	return _StarkwarePerpetuals.Contract.ConfigurationHash(&_StarkwarePerpetuals.CallOpts, arg0)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetAssetInfo(opts *bind.CallOpts, assetType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getAssetInfo", assetType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _StarkwarePerpetuals.Contract.GetAssetInfo(&_StarkwarePerpetuals.CallOpts, assetType)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _StarkwarePerpetuals.Contract.GetAssetInfo(&_StarkwarePerpetuals.CallOpts, assetType)
}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetCancellationRequest(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getCancellationRequest", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetCancellationRequest(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetCancellationRequest(&_StarkwarePerpetuals.CallOpts, starkKey, assetId, vaultId)
}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetCancellationRequest(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetCancellationRequest(&_StarkwarePerpetuals.CallOpts, starkKey, assetId, vaultId)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetDepositBalance(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getDepositBalance", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetDepositBalance(&_StarkwarePerpetuals.CallOpts, starkKey, assetId, vaultId)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetDepositBalance(&_StarkwarePerpetuals.CallOpts, starkKey, assetId, vaultId)
}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 starkKey) view returns(address ethKey)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetEthKey(opts *bind.CallOpts, starkKey *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getEthKey", starkKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 starkKey) view returns(address ethKey)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetEthKey(starkKey *big.Int) (common.Address, error) {
	return _StarkwarePerpetuals.Contract.GetEthKey(&_StarkwarePerpetuals.CallOpts, starkKey)
}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 starkKey) view returns(address ethKey)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetEthKey(starkKey *big.Int) (common.Address, error) {
	return _StarkwarePerpetuals.Contract.GetEthKey(&_StarkwarePerpetuals.CallOpts, starkKey)
}

// GetForcedTradeRequest is a free data retrieval call binding the contract method 0xf00cec4b.
//
// Solidity: function getForcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce) view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetForcedTradeRequest(opts *bind.CallOpts, starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, nonce *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getForcedTradeRequest", starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, nonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedTradeRequest is a free data retrieval call binding the contract method 0xf00cec4b.
//
// Solidity: function getForcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce) view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetForcedTradeRequest(starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, nonce *big.Int) (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetForcedTradeRequest(&_StarkwarePerpetuals.CallOpts, starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, nonce)
}

// GetForcedTradeRequest is a free data retrieval call binding the contract method 0xf00cec4b.
//
// Solidity: function getForcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce) view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetForcedTradeRequest(starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, nonce *big.Int) (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetForcedTradeRequest(&_StarkwarePerpetuals.CallOpts, starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, nonce)
}

// GetForcedWithdrawalRequest is a free data retrieval call binding the contract method 0x260e96dd.
//
// Solidity: function getForcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetForcedWithdrawalRequest(opts *bind.CallOpts, starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getForcedWithdrawalRequest", starkKey, vaultId, quantizedAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedWithdrawalRequest is a free data retrieval call binding the contract method 0x260e96dd.
//
// Solidity: function getForcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetForcedWithdrawalRequest(starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetForcedWithdrawalRequest(&_StarkwarePerpetuals.CallOpts, starkKey, vaultId, quantizedAmount)
}

// GetForcedWithdrawalRequest is a free data retrieval call binding the contract method 0x260e96dd.
//
// Solidity: function getForcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetForcedWithdrawalRequest(starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetForcedWithdrawalRequest(&_StarkwarePerpetuals.CallOpts, starkKey, vaultId, quantizedAmount)
}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 starkKey, uint256 vaultId) view returns(uint256 res)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetFullWithdrawalRequest(opts *bind.CallOpts, starkKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getFullWithdrawalRequest", starkKey, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 starkKey, uint256 vaultId) view returns(uint256 res)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetFullWithdrawalRequest(starkKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetFullWithdrawalRequest(&_StarkwarePerpetuals.CallOpts, starkKey, vaultId)
}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 starkKey, uint256 vaultId) view returns(uint256 res)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetFullWithdrawalRequest(starkKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetFullWithdrawalRequest(&_StarkwarePerpetuals.CallOpts, starkKey, vaultId)
}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256 batchId)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetLastBatchId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getLastBatchId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256 batchId)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetLastBatchId() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetLastBatchId(&_StarkwarePerpetuals.CallOpts)
}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256 batchId)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetLastBatchId() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetLastBatchId(&_StarkwarePerpetuals.CallOpts)
}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256 root)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetOrderRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getOrderRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256 root)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetOrderRoot() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetOrderRoot(&_StarkwarePerpetuals.CallOpts)
}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256 root)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetOrderRoot() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetOrderRoot(&_StarkwarePerpetuals.CallOpts)
}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256 height)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetOrderTreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getOrderTreeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256 height)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetOrderTreeHeight() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetOrderTreeHeight(&_StarkwarePerpetuals.CallOpts)
}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256 height)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetOrderTreeHeight() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetOrderTreeHeight(&_StarkwarePerpetuals.CallOpts)
}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetQuantizedDepositBalance(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getQuantizedDepositBalance", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetQuantizedDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetQuantizedDepositBalance(&_StarkwarePerpetuals.CallOpts, starkKey, assetId, vaultId)
}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetQuantizedDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetQuantizedDepositBalance(&_StarkwarePerpetuals.CallOpts, starkKey, assetId, vaultId)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetQuantum(opts *bind.CallOpts, presumedAssetType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getQuantum", presumedAssetType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetQuantum(&_StarkwarePerpetuals.CallOpts, presumedAssetType)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetQuantum(&_StarkwarePerpetuals.CallOpts, presumedAssetType)
}

// GetRegisteredAvailabilityVerifiers is a free data retrieval call binding the contract method 0x1ac347f2.
//
// Solidity: function getRegisteredAvailabilityVerifiers() view returns(address[] _verifers)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetRegisteredAvailabilityVerifiers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getRegisteredAvailabilityVerifiers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredAvailabilityVerifiers is a free data retrieval call binding the contract method 0x1ac347f2.
//
// Solidity: function getRegisteredAvailabilityVerifiers() view returns(address[] _verifers)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetRegisteredAvailabilityVerifiers() ([]common.Address, error) {
	return _StarkwarePerpetuals.Contract.GetRegisteredAvailabilityVerifiers(&_StarkwarePerpetuals.CallOpts)
}

// GetRegisteredAvailabilityVerifiers is a free data retrieval call binding the contract method 0x1ac347f2.
//
// Solidity: function getRegisteredAvailabilityVerifiers() view returns(address[] _verifers)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetRegisteredAvailabilityVerifiers() ([]common.Address, error) {
	return _StarkwarePerpetuals.Contract.GetRegisteredAvailabilityVerifiers(&_StarkwarePerpetuals.CallOpts)
}

// GetRegisteredVerifiers is a free data retrieval call binding the contract method 0x4eab9ed3.
//
// Solidity: function getRegisteredVerifiers() view returns(address[] _verifers)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetRegisteredVerifiers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getRegisteredVerifiers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredVerifiers is a free data retrieval call binding the contract method 0x4eab9ed3.
//
// Solidity: function getRegisteredVerifiers() view returns(address[] _verifers)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetRegisteredVerifiers() ([]common.Address, error) {
	return _StarkwarePerpetuals.Contract.GetRegisteredVerifiers(&_StarkwarePerpetuals.CallOpts)
}

// GetRegisteredVerifiers is a free data retrieval call binding the contract method 0x4eab9ed3.
//
// Solidity: function getRegisteredVerifiers() view returns(address[] _verifers)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetRegisteredVerifiers() ([]common.Address, error) {
	return _StarkwarePerpetuals.Contract.GetRegisteredVerifiers(&_StarkwarePerpetuals.CallOpts)
}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256 seq)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetSequenceNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getSequenceNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256 seq)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetSequenceNumber() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetSequenceNumber(&_StarkwarePerpetuals.CallOpts)
}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256 seq)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetSequenceNumber() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetSequenceNumber(&_StarkwarePerpetuals.CallOpts)
}

// GetSystemAssetType is a free data retrieval call binding the contract method 0x27b66a4d.
//
// Solidity: function getSystemAssetType() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetSystemAssetType(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getSystemAssetType")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSystemAssetType is a free data retrieval call binding the contract method 0x27b66a4d.
//
// Solidity: function getSystemAssetType() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetSystemAssetType() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetSystemAssetType(&_StarkwarePerpetuals.CallOpts)
}

// GetSystemAssetType is a free data retrieval call binding the contract method 0x27b66a4d.
//
// Solidity: function getSystemAssetType() view returns(uint256)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetSystemAssetType() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetSystemAssetType(&_StarkwarePerpetuals.CallOpts)
}

// GetVaultRoot is a free data retrieval call binding the contract method 0x64da5dfe.
//
// Solidity: function getVaultRoot() view returns(uint256 root)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetVaultRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getVaultRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultRoot is a free data retrieval call binding the contract method 0x64da5dfe.
//
// Solidity: function getVaultRoot() view returns(uint256 root)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetVaultRoot() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetVaultRoot(&_StarkwarePerpetuals.CallOpts)
}

// GetVaultRoot is a free data retrieval call binding the contract method 0x64da5dfe.
//
// Solidity: function getVaultRoot() view returns(uint256 root)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetVaultRoot() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetVaultRoot(&_StarkwarePerpetuals.CallOpts)
}

// GetVaultTreeHeight is a free data retrieval call binding the contract method 0xf288a3ff.
//
// Solidity: function getVaultTreeHeight() view returns(uint256 height)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetVaultTreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getVaultTreeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultTreeHeight is a free data retrieval call binding the contract method 0xf288a3ff.
//
// Solidity: function getVaultTreeHeight() view returns(uint256 height)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetVaultTreeHeight() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetVaultTreeHeight(&_StarkwarePerpetuals.CallOpts)
}

// GetVaultTreeHeight is a free data retrieval call binding the contract method 0xf288a3ff.
//
// Solidity: function getVaultTreeHeight() view returns(uint256 height)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetVaultTreeHeight() (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetVaultTreeHeight(&_StarkwarePerpetuals.CallOpts)
}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 starkKey, uint256 assetId) view returns(uint256 balance)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GetWithdrawalBalance(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "getWithdrawalBalance", starkKey, assetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 starkKey, uint256 assetId) view returns(uint256 balance)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GetWithdrawalBalance(starkKey *big.Int, assetId *big.Int) (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetWithdrawalBalance(&_StarkwarePerpetuals.CallOpts, starkKey, assetId)
}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 starkKey, uint256 assetId) view returns(uint256 balance)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GetWithdrawalBalance(starkKey *big.Int, assetId *big.Int) (*big.Int, error) {
	return _StarkwarePerpetuals.Contract.GetWithdrawalBalance(&_StarkwarePerpetuals.CallOpts, starkKey, assetId)
}

// GlobalConfigurationHash is a free data retrieval call binding the contract method 0xadac3e15.
//
// Solidity: function globalConfigurationHash() view returns(bytes32)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) GlobalConfigurationHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "globalConfigurationHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GlobalConfigurationHash is a free data retrieval call binding the contract method 0xadac3e15.
//
// Solidity: function globalConfigurationHash() view returns(bytes32)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) GlobalConfigurationHash() ([32]byte, error) {
	return _StarkwarePerpetuals.Contract.GlobalConfigurationHash(&_StarkwarePerpetuals.CallOpts)
}

// GlobalConfigurationHash is a free data retrieval call binding the contract method 0xadac3e15.
//
// Solidity: function globalConfigurationHash() view returns(bytes32)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) GlobalConfigurationHash() ([32]byte, error) {
	return _StarkwarePerpetuals.Contract.GlobalConfigurationHash(&_StarkwarePerpetuals.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address _implementation)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address _implementation)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) Implementation() (common.Address, error) {
	return _StarkwarePerpetuals.Contract.Implementation(&_StarkwarePerpetuals.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address _implementation)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) Implementation() (common.Address, error) {
	return _StarkwarePerpetuals.Contract.Implementation(&_StarkwarePerpetuals.CallOpts)
}

// IsAvailabilityVerifier is a free data retrieval call binding the contract method 0xbd1279ae.
//
// Solidity: function isAvailabilityVerifier(address verifierAddress) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) IsAvailabilityVerifier(opts *bind.CallOpts, verifierAddress common.Address) (bool, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "isAvailabilityVerifier", verifierAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAvailabilityVerifier is a free data retrieval call binding the contract method 0xbd1279ae.
//
// Solidity: function isAvailabilityVerifier(address verifierAddress) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) IsAvailabilityVerifier(verifierAddress common.Address) (bool, error) {
	return _StarkwarePerpetuals.Contract.IsAvailabilityVerifier(&_StarkwarePerpetuals.CallOpts, verifierAddress)
}

// IsAvailabilityVerifier is a free data retrieval call binding the contract method 0xbd1279ae.
//
// Solidity: function isAvailabilityVerifier(address verifierAddress) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) IsAvailabilityVerifier(verifierAddress common.Address) (bool, error) {
	return _StarkwarePerpetuals.Contract.IsAvailabilityVerifier(&_StarkwarePerpetuals.CallOpts, verifierAddress)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool frozen)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) IsFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "isFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool frozen)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) IsFrozen() (bool, error) {
	return _StarkwarePerpetuals.Contract.IsFrozen(&_StarkwarePerpetuals.CallOpts)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool frozen)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) IsFrozen() (bool, error) {
	return _StarkwarePerpetuals.Contract.IsFrozen(&_StarkwarePerpetuals.CallOpts)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address testedOperator) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) IsOperator(opts *bind.CallOpts, testedOperator common.Address) (bool, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "isOperator", testedOperator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address testedOperator) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) IsOperator(testedOperator common.Address) (bool, error) {
	return _StarkwarePerpetuals.Contract.IsOperator(&_StarkwarePerpetuals.CallOpts, testedOperator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address testedOperator) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) IsOperator(testedOperator common.Address) (bool, error) {
	return _StarkwarePerpetuals.Contract.IsOperator(&_StarkwarePerpetuals.CallOpts, testedOperator)
}

// IsTokenAdmin is a free data retrieval call binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address testedAdmin) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) IsTokenAdmin(opts *bind.CallOpts, testedAdmin common.Address) (bool, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "isTokenAdmin", testedAdmin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenAdmin is a free data retrieval call binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address testedAdmin) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) IsTokenAdmin(testedAdmin common.Address) (bool, error) {
	return _StarkwarePerpetuals.Contract.IsTokenAdmin(&_StarkwarePerpetuals.CallOpts, testedAdmin)
}

// IsTokenAdmin is a free data retrieval call binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address testedAdmin) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) IsTokenAdmin(testedAdmin common.Address) (bool, error) {
	return _StarkwarePerpetuals.Contract.IsTokenAdmin(&_StarkwarePerpetuals.CallOpts, testedAdmin)
}

// IsUserAdmin is a free data retrieval call binding the contract method 0x74d523a8.
//
// Solidity: function isUserAdmin(address testedAdmin) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) IsUserAdmin(opts *bind.CallOpts, testedAdmin common.Address) (bool, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "isUserAdmin", testedAdmin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserAdmin is a free data retrieval call binding the contract method 0x74d523a8.
//
// Solidity: function isUserAdmin(address testedAdmin) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) IsUserAdmin(testedAdmin common.Address) (bool, error) {
	return _StarkwarePerpetuals.Contract.IsUserAdmin(&_StarkwarePerpetuals.CallOpts, testedAdmin)
}

// IsUserAdmin is a free data retrieval call binding the contract method 0x74d523a8.
//
// Solidity: function isUserAdmin(address testedAdmin) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) IsUserAdmin(testedAdmin common.Address) (bool, error) {
	return _StarkwarePerpetuals.Contract.IsUserAdmin(&_StarkwarePerpetuals.CallOpts, testedAdmin)
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address verifierAddress) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) IsVerifier(opts *bind.CallOpts, verifierAddress common.Address) (bool, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "isVerifier", verifierAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address verifierAddress) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) IsVerifier(verifierAddress common.Address) (bool, error) {
	return _StarkwarePerpetuals.Contract.IsVerifier(&_StarkwarePerpetuals.CallOpts, verifierAddress)
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address verifierAddress) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) IsVerifier(verifierAddress common.Address) (bool, error) {
	return _StarkwarePerpetuals.Contract.IsVerifier(&_StarkwarePerpetuals.CallOpts, verifierAddress)
}

// MainIsGovernor is a free data retrieval call binding the contract method 0x45f5cd97.
//
// Solidity: function mainIsGovernor(address testGovernor) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsCaller) MainIsGovernor(opts *bind.CallOpts, testGovernor common.Address) (bool, error) {
	var out []interface{}
	err := _StarkwarePerpetuals.contract.Call(opts, &out, "mainIsGovernor", testGovernor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MainIsGovernor is a free data retrieval call binding the contract method 0x45f5cd97.
//
// Solidity: function mainIsGovernor(address testGovernor) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) MainIsGovernor(testGovernor common.Address) (bool, error) {
	return _StarkwarePerpetuals.Contract.MainIsGovernor(&_StarkwarePerpetuals.CallOpts, testGovernor)
}

// MainIsGovernor is a free data retrieval call binding the contract method 0x45f5cd97.
//
// Solidity: function mainIsGovernor(address testGovernor) view returns(bool)
func (_StarkwarePerpetuals *StarkwarePerpetualsCallerSession) MainIsGovernor(testGovernor common.Address) (bool, error) {
	return _StarkwarePerpetuals.Contract.MainIsGovernor(&_StarkwarePerpetuals.CallOpts, testGovernor)
}

// AnnounceAvailabilityVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x1d078bbb.
//
// Solidity: function announceAvailabilityVerifierRemovalIntent(address verifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) AnnounceAvailabilityVerifierRemovalIntent(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "announceAvailabilityVerifierRemovalIntent", verifier)
}

// AnnounceAvailabilityVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x1d078bbb.
//
// Solidity: function announceAvailabilityVerifierRemovalIntent(address verifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) AnnounceAvailabilityVerifierRemovalIntent(verifier common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.AnnounceAvailabilityVerifierRemovalIntent(&_StarkwarePerpetuals.TransactOpts, verifier)
}

// AnnounceAvailabilityVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x1d078bbb.
//
// Solidity: function announceAvailabilityVerifierRemovalIntent(address verifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) AnnounceAvailabilityVerifierRemovalIntent(verifier common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.AnnounceAvailabilityVerifierRemovalIntent(&_StarkwarePerpetuals.TransactOpts, verifier)
}

// AnnounceVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x418573b7.
//
// Solidity: function announceVerifierRemovalIntent(address verifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) AnnounceVerifierRemovalIntent(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "announceVerifierRemovalIntent", verifier)
}

// AnnounceVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x418573b7.
//
// Solidity: function announceVerifierRemovalIntent(address verifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) AnnounceVerifierRemovalIntent(verifier common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.AnnounceVerifierRemovalIntent(&_StarkwarePerpetuals.TransactOpts, verifier)
}

// AnnounceVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x418573b7.
//
// Solidity: function announceVerifierRemovalIntent(address verifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) AnnounceVerifierRemovalIntent(verifier common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.AnnounceVerifierRemovalIntent(&_StarkwarePerpetuals.TransactOpts, verifier)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) Deposit(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "deposit", starkKey, assetType, vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) Deposit(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.Deposit(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType, vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) Deposit(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.Deposit(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType, vaultId)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) Deposit0(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "deposit0", starkKey, assetType, vaultId, quantizedAmount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) Deposit0(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.Deposit0(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) Deposit0(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.Deposit0(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) DepositCancel(opts *bind.TransactOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "depositCancel", starkKey, assetId, vaultId)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) DepositCancel(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.DepositCancel(&_StarkwarePerpetuals.TransactOpts, starkKey, assetId, vaultId)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) DepositCancel(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.DepositCancel(&_StarkwarePerpetuals.TransactOpts, starkKey, assetId, vaultId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) DepositNft(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "depositNft", starkKey, assetType, vaultId, tokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) DepositNft(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.DepositNft(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) DepositNft(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.DepositNft(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) DepositNftReclaim(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "depositNftReclaim", starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) DepositNftReclaim(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.DepositNftReclaim(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) DepositNftReclaim(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.DepositNftReclaim(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) DepositReclaim(opts *bind.TransactOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "depositReclaim", starkKey, assetId, vaultId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) DepositReclaim(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.DepositReclaim(&_StarkwarePerpetuals.TransactOpts, starkKey, assetId, vaultId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) DepositReclaim(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.DepositReclaim(&_StarkwarePerpetuals.TransactOpts, starkKey, assetId, vaultId)
}

// Escape is a paid mutator transaction binding the contract method 0x366a2670.
//
// Solidity: function escape(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) Escape(opts *bind.TransactOpts, starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "escape", starkKey, vaultId, quantizedAmount)
}

// Escape is a paid mutator transaction binding the contract method 0x366a2670.
//
// Solidity: function escape(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) Escape(starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.Escape(&_StarkwarePerpetuals.TransactOpts, starkKey, vaultId, quantizedAmount)
}

// Escape is a paid mutator transaction binding the contract method 0x366a2670.
//
// Solidity: function escape(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) Escape(starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.Escape(&_StarkwarePerpetuals.TransactOpts, starkKey, vaultId, quantizedAmount)
}

// ForcedTradeRequest is a paid mutator transaction binding the contract method 0x2ecb8162.
//
// Solidity: function forcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 submissionExpirationTime, uint256 nonce, bytes signature, bool premiumCost) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) ForcedTradeRequest(opts *bind.TransactOpts, starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, submissionExpirationTime *big.Int, nonce *big.Int, signature []byte, premiumCost bool) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "forcedTradeRequest", starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, submissionExpirationTime, nonce, signature, premiumCost)
}

// ForcedTradeRequest is a paid mutator transaction binding the contract method 0x2ecb8162.
//
// Solidity: function forcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 submissionExpirationTime, uint256 nonce, bytes signature, bool premiumCost) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) ForcedTradeRequest(starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, submissionExpirationTime *big.Int, nonce *big.Int, signature []byte, premiumCost bool) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.ForcedTradeRequest(&_StarkwarePerpetuals.TransactOpts, starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, submissionExpirationTime, nonce, signature, premiumCost)
}

// ForcedTradeRequest is a paid mutator transaction binding the contract method 0x2ecb8162.
//
// Solidity: function forcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 submissionExpirationTime, uint256 nonce, bytes signature, bool premiumCost) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) ForcedTradeRequest(starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, submissionExpirationTime *big.Int, nonce *big.Int, signature []byte, premiumCost bool) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.ForcedTradeRequest(&_StarkwarePerpetuals.TransactOpts, starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, submissionExpirationTime, nonce, signature, premiumCost)
}

// ForcedWithdrawalRequest is a paid mutator transaction binding the contract method 0xaf1437a3.
//
// Solidity: function forcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount, bool premiumCost) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) ForcedWithdrawalRequest(opts *bind.TransactOpts, starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int, premiumCost bool) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "forcedWithdrawalRequest", starkKey, vaultId, quantizedAmount, premiumCost)
}

// ForcedWithdrawalRequest is a paid mutator transaction binding the contract method 0xaf1437a3.
//
// Solidity: function forcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount, bool premiumCost) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) ForcedWithdrawalRequest(starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int, premiumCost bool) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.ForcedWithdrawalRequest(&_StarkwarePerpetuals.TransactOpts, starkKey, vaultId, quantizedAmount, premiumCost)
}

// ForcedWithdrawalRequest is a paid mutator transaction binding the contract method 0xaf1437a3.
//
// Solidity: function forcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount, bool premiumCost) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) ForcedWithdrawalRequest(starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int, premiumCost bool) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.ForcedWithdrawalRequest(&_StarkwarePerpetuals.TransactOpts, starkKey, vaultId, quantizedAmount, premiumCost)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x3153a386.
//
// Solidity: function freezeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) FreezeRequest(opts *bind.TransactOpts, starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, nonce *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "freezeRequest", starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, nonce)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x3153a386.
//
// Solidity: function freezeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) FreezeRequest(starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, nonce *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.FreezeRequest(&_StarkwarePerpetuals.TransactOpts, starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, nonce)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x3153a386.
//
// Solidity: function freezeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) FreezeRequest(starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, nonce *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.FreezeRequest(&_StarkwarePerpetuals.TransactOpts, starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, nonce)
}

// FreezeRequest0 is a paid mutator transaction binding the contract method 0x75d4eefd.
//
// Solidity: function freezeRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) FreezeRequest0(opts *bind.TransactOpts, starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "freezeRequest0", starkKey, vaultId, quantizedAmount)
}

// FreezeRequest0 is a paid mutator transaction binding the contract method 0x75d4eefd.
//
// Solidity: function freezeRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) FreezeRequest0(starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.FreezeRequest0(&_StarkwarePerpetuals.TransactOpts, starkKey, vaultId, quantizedAmount)
}

// FreezeRequest0 is a paid mutator transaction binding the contract method 0x75d4eefd.
//
// Solidity: function freezeRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) FreezeRequest0(starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.FreezeRequest0(&_StarkwarePerpetuals.TransactOpts, starkKey, vaultId, quantizedAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) Initialize(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "initialize", data)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) Initialize(data []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.Initialize(&_StarkwarePerpetuals.TransactOpts, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) Initialize(data []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.Initialize(&_StarkwarePerpetuals.TransactOpts, data)
}

// MainAcceptGovernance is a paid mutator transaction binding the contract method 0x28700a15.
//
// Solidity: function mainAcceptGovernance() returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) MainAcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "mainAcceptGovernance")
}

// MainAcceptGovernance is a paid mutator transaction binding the contract method 0x28700a15.
//
// Solidity: function mainAcceptGovernance() returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) MainAcceptGovernance() (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.MainAcceptGovernance(&_StarkwarePerpetuals.TransactOpts)
}

// MainAcceptGovernance is a paid mutator transaction binding the contract method 0x28700a15.
//
// Solidity: function mainAcceptGovernance() returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) MainAcceptGovernance() (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.MainAcceptGovernance(&_StarkwarePerpetuals.TransactOpts)
}

// MainCancelNomination is a paid mutator transaction binding the contract method 0x72eb3688.
//
// Solidity: function mainCancelNomination() returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) MainCancelNomination(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "mainCancelNomination")
}

// MainCancelNomination is a paid mutator transaction binding the contract method 0x72eb3688.
//
// Solidity: function mainCancelNomination() returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) MainCancelNomination() (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.MainCancelNomination(&_StarkwarePerpetuals.TransactOpts)
}

// MainCancelNomination is a paid mutator transaction binding the contract method 0x72eb3688.
//
// Solidity: function mainCancelNomination() returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) MainCancelNomination() (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.MainCancelNomination(&_StarkwarePerpetuals.TransactOpts)
}

// MainNominateNewGovernor is a paid mutator transaction binding the contract method 0x8c4bce1c.
//
// Solidity: function mainNominateNewGovernor(address newGovernor) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) MainNominateNewGovernor(opts *bind.TransactOpts, newGovernor common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "mainNominateNewGovernor", newGovernor)
}

// MainNominateNewGovernor is a paid mutator transaction binding the contract method 0x8c4bce1c.
//
// Solidity: function mainNominateNewGovernor(address newGovernor) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) MainNominateNewGovernor(newGovernor common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.MainNominateNewGovernor(&_StarkwarePerpetuals.TransactOpts, newGovernor)
}

// MainNominateNewGovernor is a paid mutator transaction binding the contract method 0x8c4bce1c.
//
// Solidity: function mainNominateNewGovernor(address newGovernor) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) MainNominateNewGovernor(newGovernor common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.MainNominateNewGovernor(&_StarkwarePerpetuals.TransactOpts, newGovernor)
}

// MainRemoveGovernor is a paid mutator transaction binding the contract method 0xa1cc921e.
//
// Solidity: function mainRemoveGovernor(address governorForRemoval) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) MainRemoveGovernor(opts *bind.TransactOpts, governorForRemoval common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "mainRemoveGovernor", governorForRemoval)
}

// MainRemoveGovernor is a paid mutator transaction binding the contract method 0xa1cc921e.
//
// Solidity: function mainRemoveGovernor(address governorForRemoval) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) MainRemoveGovernor(governorForRemoval common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.MainRemoveGovernor(&_StarkwarePerpetuals.TransactOpts, governorForRemoval)
}

// MainRemoveGovernor is a paid mutator transaction binding the contract method 0xa1cc921e.
//
// Solidity: function mainRemoveGovernor(address governorForRemoval) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) MainRemoveGovernor(governorForRemoval common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.MainRemoveGovernor(&_StarkwarePerpetuals.TransactOpts, governorForRemoval)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.OnERC721Received(&_StarkwarePerpetuals.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.OnERC721Received(&_StarkwarePerpetuals.TransactOpts, arg0, arg1, arg2, arg3)
}

// RegisterAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbdb75785.
//
// Solidity: function registerAvailabilityVerifier(address verifier, string identifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) RegisterAvailabilityVerifier(opts *bind.TransactOpts, verifier common.Address, identifier string) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "registerAvailabilityVerifier", verifier, identifier)
}

// RegisterAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbdb75785.
//
// Solidity: function registerAvailabilityVerifier(address verifier, string identifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) RegisterAvailabilityVerifier(verifier common.Address, identifier string) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterAvailabilityVerifier(&_StarkwarePerpetuals.TransactOpts, verifier, identifier)
}

// RegisterAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbdb75785.
//
// Solidity: function registerAvailabilityVerifier(address verifier, string identifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) RegisterAvailabilityVerifier(verifier common.Address, identifier string) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterAvailabilityVerifier(&_StarkwarePerpetuals.TransactOpts, verifier, identifier)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) RegisterOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "registerOperator", newOperator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) RegisterOperator(newOperator common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterOperator(&_StarkwarePerpetuals.TransactOpts, newOperator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) RegisterOperator(newOperator common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterOperator(&_StarkwarePerpetuals.TransactOpts, newOperator)
}

// RegisterSystemAssetType is a paid mutator transaction binding the contract method 0x7fbf9ba9.
//
// Solidity: function registerSystemAssetType(uint256 assetType, bytes assetInfo) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) RegisterSystemAssetType(opts *bind.TransactOpts, assetType *big.Int, assetInfo []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "registerSystemAssetType", assetType, assetInfo)
}

// RegisterSystemAssetType is a paid mutator transaction binding the contract method 0x7fbf9ba9.
//
// Solidity: function registerSystemAssetType(uint256 assetType, bytes assetInfo) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) RegisterSystemAssetType(assetType *big.Int, assetInfo []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterSystemAssetType(&_StarkwarePerpetuals.TransactOpts, assetType, assetInfo)
}

// RegisterSystemAssetType is a paid mutator transaction binding the contract method 0x7fbf9ba9.
//
// Solidity: function registerSystemAssetType(uint256 assetType, bytes assetInfo) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) RegisterSystemAssetType(assetType *big.Int, assetInfo []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterSystemAssetType(&_StarkwarePerpetuals.TransactOpts, assetType, assetInfo)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 , bytes ) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) RegisterToken(opts *bind.TransactOpts, arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "registerToken", arg0, arg1)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 , bytes ) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) RegisterToken(arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterToken(&_StarkwarePerpetuals.TransactOpts, arg0, arg1)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 , bytes ) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) RegisterToken(arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterToken(&_StarkwarePerpetuals.TransactOpts, arg0, arg1)
}

// RegisterToken0 is a paid mutator transaction binding the contract method 0xd88d8b38.
//
// Solidity: function registerToken(uint256 , bytes , uint256 ) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) RegisterToken0(opts *bind.TransactOpts, arg0 *big.Int, arg1 []byte, arg2 *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "registerToken0", arg0, arg1, arg2)
}

// RegisterToken0 is a paid mutator transaction binding the contract method 0xd88d8b38.
//
// Solidity: function registerToken(uint256 , bytes , uint256 ) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) RegisterToken0(arg0 *big.Int, arg1 []byte, arg2 *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterToken0(&_StarkwarePerpetuals.TransactOpts, arg0, arg1, arg2)
}

// RegisterToken0 is a paid mutator transaction binding the contract method 0xd88d8b38.
//
// Solidity: function registerToken(uint256 , bytes , uint256 ) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) RegisterToken0(arg0 *big.Int, arg1 []byte, arg2 *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterToken0(&_StarkwarePerpetuals.TransactOpts, arg0, arg1, arg2)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address newAdmin) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) RegisterTokenAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "registerTokenAdmin", newAdmin)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address newAdmin) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) RegisterTokenAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterTokenAdmin(&_StarkwarePerpetuals.TransactOpts, newAdmin)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address newAdmin) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) RegisterTokenAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterTokenAdmin(&_StarkwarePerpetuals.TransactOpts, newAdmin)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xdd2414d4.
//
// Solidity: function registerUser(address ethKey, uint256 starkKey, bytes signature) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) RegisterUser(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, signature []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "registerUser", ethKey, starkKey, signature)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xdd2414d4.
//
// Solidity: function registerUser(address ethKey, uint256 starkKey, bytes signature) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) RegisterUser(ethKey common.Address, starkKey *big.Int, signature []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterUser(&_StarkwarePerpetuals.TransactOpts, ethKey, starkKey, signature)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xdd2414d4.
//
// Solidity: function registerUser(address ethKey, uint256 starkKey, bytes signature) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) RegisterUser(ethKey common.Address, starkKey *big.Int, signature []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterUser(&_StarkwarePerpetuals.TransactOpts, ethKey, starkKey, signature)
}

// RegisterUserAdmin is a paid mutator transaction binding the contract method 0xf3e0c3fb.
//
// Solidity: function registerUserAdmin(address newAdmin) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) RegisterUserAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "registerUserAdmin", newAdmin)
}

// RegisterUserAdmin is a paid mutator transaction binding the contract method 0xf3e0c3fb.
//
// Solidity: function registerUserAdmin(address newAdmin) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) RegisterUserAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterUserAdmin(&_StarkwarePerpetuals.TransactOpts, newAdmin)
}

// RegisterUserAdmin is a paid mutator transaction binding the contract method 0xf3e0c3fb.
//
// Solidity: function registerUserAdmin(address newAdmin) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) RegisterUserAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterUserAdmin(&_StarkwarePerpetuals.TransactOpts, newAdmin)
}

// RegisterVerifier is a paid mutator transaction binding the contract method 0x3776fe2a.
//
// Solidity: function registerVerifier(address verifier, string identifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) RegisterVerifier(opts *bind.TransactOpts, verifier common.Address, identifier string) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "registerVerifier", verifier, identifier)
}

// RegisterVerifier is a paid mutator transaction binding the contract method 0x3776fe2a.
//
// Solidity: function registerVerifier(address verifier, string identifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) RegisterVerifier(verifier common.Address, identifier string) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterVerifier(&_StarkwarePerpetuals.TransactOpts, verifier, identifier)
}

// RegisterVerifier is a paid mutator transaction binding the contract method 0x3776fe2a.
//
// Solidity: function registerVerifier(address verifier, string identifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) RegisterVerifier(verifier common.Address, identifier string) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RegisterVerifier(&_StarkwarePerpetuals.TransactOpts, verifier, identifier)
}

// RemoveAvailabilityVerifier is a paid mutator transaction binding the contract method 0xb1e640bf.
//
// Solidity: function removeAvailabilityVerifier(address verifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) RemoveAvailabilityVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "removeAvailabilityVerifier", verifier)
}

// RemoveAvailabilityVerifier is a paid mutator transaction binding the contract method 0xb1e640bf.
//
// Solidity: function removeAvailabilityVerifier(address verifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) RemoveAvailabilityVerifier(verifier common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RemoveAvailabilityVerifier(&_StarkwarePerpetuals.TransactOpts, verifier)
}

// RemoveAvailabilityVerifier is a paid mutator transaction binding the contract method 0xb1e640bf.
//
// Solidity: function removeAvailabilityVerifier(address verifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) RemoveAvailabilityVerifier(verifier common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RemoveAvailabilityVerifier(&_StarkwarePerpetuals.TransactOpts, verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) RemoveVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "removeVerifier", verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) RemoveVerifier(verifier common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RemoveVerifier(&_StarkwarePerpetuals.TransactOpts, verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) RemoveVerifier(verifier common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.RemoveVerifier(&_StarkwarePerpetuals.TransactOpts, verifier)
}

// SetAssetConfiguration is a paid mutator transaction binding the contract method 0x3e8e8839.
//
// Solidity: function setAssetConfiguration(uint256 assetId, bytes32 configHash) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) SetAssetConfiguration(opts *bind.TransactOpts, assetId *big.Int, configHash [32]byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "setAssetConfiguration", assetId, configHash)
}

// SetAssetConfiguration is a paid mutator transaction binding the contract method 0x3e8e8839.
//
// Solidity: function setAssetConfiguration(uint256 assetId, bytes32 configHash) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) SetAssetConfiguration(assetId *big.Int, configHash [32]byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.SetAssetConfiguration(&_StarkwarePerpetuals.TransactOpts, assetId, configHash)
}

// SetAssetConfiguration is a paid mutator transaction binding the contract method 0x3e8e8839.
//
// Solidity: function setAssetConfiguration(uint256 assetId, bytes32 configHash) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) SetAssetConfiguration(assetId *big.Int, configHash [32]byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.SetAssetConfiguration(&_StarkwarePerpetuals.TransactOpts, assetId, configHash)
}

// SetGlobalConfiguration is a paid mutator transaction binding the contract method 0x474f249a.
//
// Solidity: function setGlobalConfiguration(bytes32 configHash) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) SetGlobalConfiguration(opts *bind.TransactOpts, configHash [32]byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "setGlobalConfiguration", configHash)
}

// SetGlobalConfiguration is a paid mutator transaction binding the contract method 0x474f249a.
//
// Solidity: function setGlobalConfiguration(bytes32 configHash) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) SetGlobalConfiguration(configHash [32]byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.SetGlobalConfiguration(&_StarkwarePerpetuals.TransactOpts, configHash)
}

// SetGlobalConfiguration is a paid mutator transaction binding the contract method 0x474f249a.
//
// Solidity: function setGlobalConfiguration(bytes32 configHash) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) SetGlobalConfiguration(configHash [32]byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.SetGlobalConfiguration(&_StarkwarePerpetuals.TransactOpts, configHash)
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) UnFreeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "unFreeze")
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) UnFreeze() (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.UnFreeze(&_StarkwarePerpetuals.TransactOpts)
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) UnFreeze() (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.UnFreeze(&_StarkwarePerpetuals.TransactOpts)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) UnregisterOperator(opts *bind.TransactOpts, removedOperator common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "unregisterOperator", removedOperator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) UnregisterOperator(removedOperator common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.UnregisterOperator(&_StarkwarePerpetuals.TransactOpts, removedOperator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) UnregisterOperator(removedOperator common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.UnregisterOperator(&_StarkwarePerpetuals.TransactOpts, removedOperator)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address oldAdmin) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) UnregisterTokenAdmin(opts *bind.TransactOpts, oldAdmin common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "unregisterTokenAdmin", oldAdmin)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address oldAdmin) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) UnregisterTokenAdmin(oldAdmin common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.UnregisterTokenAdmin(&_StarkwarePerpetuals.TransactOpts, oldAdmin)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address oldAdmin) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) UnregisterTokenAdmin(oldAdmin common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.UnregisterTokenAdmin(&_StarkwarePerpetuals.TransactOpts, oldAdmin)
}

// UnregisterUserAdmin is a paid mutator transaction binding the contract method 0xb04b6179.
//
// Solidity: function unregisterUserAdmin(address oldAdmin) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) UnregisterUserAdmin(opts *bind.TransactOpts, oldAdmin common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "unregisterUserAdmin", oldAdmin)
}

// UnregisterUserAdmin is a paid mutator transaction binding the contract method 0xb04b6179.
//
// Solidity: function unregisterUserAdmin(address oldAdmin) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) UnregisterUserAdmin(oldAdmin common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.UnregisterUserAdmin(&_StarkwarePerpetuals.TransactOpts, oldAdmin)
}

// UnregisterUserAdmin is a paid mutator transaction binding the contract method 0xb04b6179.
//
// Solidity: function unregisterUserAdmin(address oldAdmin) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) UnregisterUserAdmin(oldAdmin common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.UnregisterUserAdmin(&_StarkwarePerpetuals.TransactOpts, oldAdmin)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] programOutput, uint256[] applicationData) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) UpdateState(opts *bind.TransactOpts, programOutput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "updateState", programOutput, applicationData)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] programOutput, uint256[] applicationData) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) UpdateState(programOutput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.UpdateState(&_StarkwarePerpetuals.TransactOpts, programOutput, applicationData)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] programOutput, uint256[] applicationData) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) UpdateState(programOutput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.UpdateState(&_StarkwarePerpetuals.TransactOpts, programOutput, applicationData)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 starkKey, uint256 assetType) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) Withdraw(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "withdraw", starkKey, assetType)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 starkKey, uint256 assetType) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) Withdraw(starkKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.Withdraw(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 starkKey, uint256 assetType) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) Withdraw(starkKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.Withdraw(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 starkKey, uint256 assetType, bytes mintingBlob) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) WithdrawAndMint(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "withdrawAndMint", starkKey, assetType, mintingBlob)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 starkKey, uint256 assetType, bytes mintingBlob) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) WithdrawAndMint(starkKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.WithdrawAndMint(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType, mintingBlob)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 starkKey, uint256 assetType, bytes mintingBlob) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) WithdrawAndMint(starkKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.WithdrawAndMint(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType, mintingBlob)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 starkKey, uint256 assetType, uint256 tokenId) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) WithdrawNft(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "withdrawNft", starkKey, assetType, tokenId)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 starkKey, uint256 assetType, uint256 tokenId) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) WithdrawNft(starkKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.WithdrawNft(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType, tokenId)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 starkKey, uint256 assetType, uint256 tokenId) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) WithdrawNft(starkKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.WithdrawNft(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType, tokenId)
}

// WithdrawNftTo is a paid mutator transaction binding the contract method 0xebef0fd0.
//
// Solidity: function withdrawNftTo(uint256 starkKey, uint256 assetType, uint256 tokenId, address recipient) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) WithdrawNftTo(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, tokenId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "withdrawNftTo", starkKey, assetType, tokenId, recipient)
}

// WithdrawNftTo is a paid mutator transaction binding the contract method 0xebef0fd0.
//
// Solidity: function withdrawNftTo(uint256 starkKey, uint256 assetType, uint256 tokenId, address recipient) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) WithdrawNftTo(starkKey *big.Int, assetType *big.Int, tokenId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.WithdrawNftTo(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType, tokenId, recipient)
}

// WithdrawNftTo is a paid mutator transaction binding the contract method 0xebef0fd0.
//
// Solidity: function withdrawNftTo(uint256 starkKey, uint256 assetType, uint256 tokenId, address recipient) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) WithdrawNftTo(starkKey *big.Int, assetType *big.Int, tokenId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.WithdrawNftTo(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType, tokenId, recipient)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 starkKey, uint256 assetType, address recipient) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) WithdrawTo(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.Transact(opts, "withdrawTo", starkKey, assetType, recipient)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 starkKey, uint256 assetType, address recipient) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) WithdrawTo(starkKey *big.Int, assetType *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.WithdrawTo(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType, recipient)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 starkKey, uint256 assetType, address recipient) returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) WithdrawTo(starkKey *big.Int, assetType *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.WithdrawTo(&_StarkwarePerpetuals.TransactOpts, starkKey, assetType, recipient)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.Fallback(&_StarkwarePerpetuals.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StarkwarePerpetuals *StarkwarePerpetualsTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StarkwarePerpetuals.Contract.Fallback(&_StarkwarePerpetuals.TransactOpts, calldata)
}

// StarkwarePerpetualsLogDepositIterator is returned from FilterLogDeposit and is used to iterate over the raw logs and unpacked data for LogDeposit events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogDepositIterator struct {
	Event *StarkwarePerpetualsLogDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogDeposit represents a LogDeposit event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogDeposit struct {
	DepositorEthKey    common.Address
	StarkKey           *big.Int
	VaultId            *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDeposit is a free log retrieval operation binding the contract event 0x06724742ccc8c330a39a641ef02a0b419bd09248360680bb38159b0a8c2635d6.
//
// Solidity: event LogDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogDeposit(opts *bind.FilterOpts) (*StarkwarePerpetualsLogDepositIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogDeposit")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogDepositIterator{contract: _StarkwarePerpetuals.contract, event: "LogDeposit", logs: logs, sub: sub}, nil
}

// WatchLogDeposit is a free log subscription operation binding the contract event 0x06724742ccc8c330a39a641ef02a0b419bd09248360680bb38159b0a8c2635d6.
//
// Solidity: event LogDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogDeposit(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogDeposit) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogDeposit)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogDeposit is a log parse operation binding the contract event 0x06724742ccc8c330a39a641ef02a0b419bd09248360680bb38159b0a8c2635d6.
//
// Solidity: event LogDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogDeposit(log types.Log) (*StarkwarePerpetualsLogDeposit, error) {
	event := new(StarkwarePerpetualsLogDeposit)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogDepositCancelIterator is returned from FilterLogDepositCancel and is used to iterate over the raw logs and unpacked data for LogDepositCancel events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogDepositCancelIterator struct {
	Event *StarkwarePerpetualsLogDepositCancel // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogDepositCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogDepositCancel)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogDepositCancel)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogDepositCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogDepositCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogDepositCancel represents a LogDepositCancel event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogDepositCancel struct {
	StarkKey *big.Int
	VaultId  *big.Int
	AssetId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogDepositCancel is a free log retrieval operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogDepositCancel(opts *bind.FilterOpts) (*StarkwarePerpetualsLogDepositCancelIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogDepositCancel")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogDepositCancelIterator{contract: _StarkwarePerpetuals.contract, event: "LogDepositCancel", logs: logs, sub: sub}, nil
}

// WatchLogDepositCancel is a free log subscription operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogDepositCancel(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogDepositCancel) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogDepositCancel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogDepositCancel)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogDepositCancel", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogDepositCancel is a log parse operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogDepositCancel(log types.Log) (*StarkwarePerpetualsLogDepositCancel, error) {
	event := new(StarkwarePerpetualsLogDepositCancel)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogDepositCancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogDepositCancelReclaimedIterator is returned from FilterLogDepositCancelReclaimed and is used to iterate over the raw logs and unpacked data for LogDepositCancelReclaimed events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogDepositCancelReclaimedIterator struct {
	Event *StarkwarePerpetualsLogDepositCancelReclaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogDepositCancelReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogDepositCancelReclaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogDepositCancelReclaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogDepositCancelReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogDepositCancelReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogDepositCancelReclaimed represents a LogDepositCancelReclaimed event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogDepositCancelReclaimed struct {
	StarkKey           *big.Int
	VaultId            *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDepositCancelReclaimed is a free log retrieval operation binding the contract event 0xe3e46ecf1138180bf93cba62a0b7e661d976a8ab3d40243f7b082667d8f500af.
//
// Solidity: event LogDepositCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogDepositCancelReclaimed(opts *bind.FilterOpts) (*StarkwarePerpetualsLogDepositCancelReclaimedIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogDepositCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogDepositCancelReclaimedIterator{contract: _StarkwarePerpetuals.contract, event: "LogDepositCancelReclaimed", logs: logs, sub: sub}, nil
}

// WatchLogDepositCancelReclaimed is a free log subscription operation binding the contract event 0xe3e46ecf1138180bf93cba62a0b7e661d976a8ab3d40243f7b082667d8f500af.
//
// Solidity: event LogDepositCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogDepositCancelReclaimed(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogDepositCancelReclaimed) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogDepositCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogDepositCancelReclaimed)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogDepositCancelReclaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogDepositCancelReclaimed is a log parse operation binding the contract event 0xe3e46ecf1138180bf93cba62a0b7e661d976a8ab3d40243f7b082667d8f500af.
//
// Solidity: event LogDepositCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogDepositCancelReclaimed(log types.Log) (*StarkwarePerpetualsLogDepositCancelReclaimed, error) {
	event := new(StarkwarePerpetualsLogDepositCancelReclaimed)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogDepositCancelReclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogDepositNftCancelReclaimedIterator is returned from FilterLogDepositNftCancelReclaimed and is used to iterate over the raw logs and unpacked data for LogDepositNftCancelReclaimed events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogDepositNftCancelReclaimedIterator struct {
	Event *StarkwarePerpetualsLogDepositNftCancelReclaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogDepositNftCancelReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogDepositNftCancelReclaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogDepositNftCancelReclaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogDepositNftCancelReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogDepositNftCancelReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogDepositNftCancelReclaimed represents a LogDepositNftCancelReclaimed event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogDepositNftCancelReclaimed struct {
	StarkKey  *big.Int
	VaultId   *big.Int
	AssetType *big.Int
	TokenId   *big.Int
	AssetId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogDepositNftCancelReclaimed is a free log retrieval operation binding the contract event 0xf00c0c1a754f6df7545d96a7e12aad552728b94ca6aa94f81e297bdbcf1dab9c.
//
// Solidity: event LogDepositNftCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogDepositNftCancelReclaimed(opts *bind.FilterOpts) (*StarkwarePerpetualsLogDepositNftCancelReclaimedIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogDepositNftCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogDepositNftCancelReclaimedIterator{contract: _StarkwarePerpetuals.contract, event: "LogDepositNftCancelReclaimed", logs: logs, sub: sub}, nil
}

// WatchLogDepositNftCancelReclaimed is a free log subscription operation binding the contract event 0xf00c0c1a754f6df7545d96a7e12aad552728b94ca6aa94f81e297bdbcf1dab9c.
//
// Solidity: event LogDepositNftCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogDepositNftCancelReclaimed(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogDepositNftCancelReclaimed) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogDepositNftCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogDepositNftCancelReclaimed)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogDepositNftCancelReclaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogDepositNftCancelReclaimed is a log parse operation binding the contract event 0xf00c0c1a754f6df7545d96a7e12aad552728b94ca6aa94f81e297bdbcf1dab9c.
//
// Solidity: event LogDepositNftCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogDepositNftCancelReclaimed(log types.Log) (*StarkwarePerpetualsLogDepositNftCancelReclaimed, error) {
	event := new(StarkwarePerpetualsLogDepositNftCancelReclaimed)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogDepositNftCancelReclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogForcedTradeRequestIterator is returned from FilterLogForcedTradeRequest and is used to iterate over the raw logs and unpacked data for LogForcedTradeRequest events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogForcedTradeRequestIterator struct {
	Event *StarkwarePerpetualsLogForcedTradeRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogForcedTradeRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogForcedTradeRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogForcedTradeRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogForcedTradeRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogForcedTradeRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogForcedTradeRequest represents a LogForcedTradeRequest event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogForcedTradeRequest struct {
	StarkKeyA          *big.Int
	StarkKeyB          *big.Int
	VaultIdA           *big.Int
	VaultIdB           *big.Int
	CollateralAssetId  *big.Int
	SyntheticAssetId   *big.Int
	AmountCollateral   *big.Int
	AmountSynthetic    *big.Int
	AIsBuyingSynthetic bool
	Nonce              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogForcedTradeRequest is a free log retrieval operation binding the contract event 0x79acb9227c98b68d7628d2c99a7719926eff77e8c2275f6ffe7cf255b32772be.
//
// Solidity: event LogForcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogForcedTradeRequest(opts *bind.FilterOpts) (*StarkwarePerpetualsLogForcedTradeRequestIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogForcedTradeRequest")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogForcedTradeRequestIterator{contract: _StarkwarePerpetuals.contract, event: "LogForcedTradeRequest", logs: logs, sub: sub}, nil
}

// WatchLogForcedTradeRequest is a free log subscription operation binding the contract event 0x79acb9227c98b68d7628d2c99a7719926eff77e8c2275f6ffe7cf255b32772be.
//
// Solidity: event LogForcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogForcedTradeRequest(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogForcedTradeRequest) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogForcedTradeRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogForcedTradeRequest)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogForcedTradeRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogForcedTradeRequest is a log parse operation binding the contract event 0x79acb9227c98b68d7628d2c99a7719926eff77e8c2275f6ffe7cf255b32772be.
//
// Solidity: event LogForcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogForcedTradeRequest(log types.Log) (*StarkwarePerpetualsLogForcedTradeRequest, error) {
	event := new(StarkwarePerpetualsLogForcedTradeRequest)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogForcedTradeRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogForcedWithdrawalRequestIterator is returned from FilterLogForcedWithdrawalRequest and is used to iterate over the raw logs and unpacked data for LogForcedWithdrawalRequest events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogForcedWithdrawalRequestIterator struct {
	Event *StarkwarePerpetualsLogForcedWithdrawalRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogForcedWithdrawalRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogForcedWithdrawalRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogForcedWithdrawalRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogForcedWithdrawalRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogForcedWithdrawalRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogForcedWithdrawalRequest represents a LogForcedWithdrawalRequest event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogForcedWithdrawalRequest struct {
	StarkKey        *big.Int
	VaultId         *big.Int
	QuantizedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogForcedWithdrawalRequest is a free log retrieval operation binding the contract event 0x5f3a3d7d7c2b8cc20fa315d7504580e81d5f74ea46e1b3dc2985281df08e204d.
//
// Solidity: event LogForcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogForcedWithdrawalRequest(opts *bind.FilterOpts) (*StarkwarePerpetualsLogForcedWithdrawalRequestIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogForcedWithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogForcedWithdrawalRequestIterator{contract: _StarkwarePerpetuals.contract, event: "LogForcedWithdrawalRequest", logs: logs, sub: sub}, nil
}

// WatchLogForcedWithdrawalRequest is a free log subscription operation binding the contract event 0x5f3a3d7d7c2b8cc20fa315d7504580e81d5f74ea46e1b3dc2985281df08e204d.
//
// Solidity: event LogForcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogForcedWithdrawalRequest(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogForcedWithdrawalRequest) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogForcedWithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogForcedWithdrawalRequest)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogForcedWithdrawalRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogForcedWithdrawalRequest is a log parse operation binding the contract event 0x5f3a3d7d7c2b8cc20fa315d7504580e81d5f74ea46e1b3dc2985281df08e204d.
//
// Solidity: event LogForcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogForcedWithdrawalRequest(log types.Log) (*StarkwarePerpetualsLogForcedWithdrawalRequest, error) {
	event := new(StarkwarePerpetualsLogForcedWithdrawalRequest)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogForcedWithdrawalRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogFrozenIterator is returned from FilterLogFrozen and is used to iterate over the raw logs and unpacked data for LogFrozen events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogFrozenIterator struct {
	Event *StarkwarePerpetualsLogFrozen // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogFrozen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogFrozen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogFrozen represents a LogFrozen event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogFrozen struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogFrozen is a free log retrieval operation binding the contract event 0xf5b8e6419478ab140eb98026ab5bd607038cb0ac4d4dad5b1fc0848dfd203d1f.
//
// Solidity: event LogFrozen()
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogFrozen(opts *bind.FilterOpts) (*StarkwarePerpetualsLogFrozenIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogFrozen")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogFrozenIterator{contract: _StarkwarePerpetuals.contract, event: "LogFrozen", logs: logs, sub: sub}, nil
}

// WatchLogFrozen is a free log subscription operation binding the contract event 0xf5b8e6419478ab140eb98026ab5bd607038cb0ac4d4dad5b1fc0848dfd203d1f.
//
// Solidity: event LogFrozen()
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogFrozen(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogFrozen) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogFrozen)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogFrozen", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogFrozen is a log parse operation binding the contract event 0xf5b8e6419478ab140eb98026ab5bd607038cb0ac4d4dad5b1fc0848dfd203d1f.
//
// Solidity: event LogFrozen()
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogFrozen(log types.Log) (*StarkwarePerpetualsLogFrozen, error) {
	event := new(StarkwarePerpetualsLogFrozen)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogGlobalConfigurationIterator is returned from FilterLogGlobalConfiguration and is used to iterate over the raw logs and unpacked data for LogGlobalConfiguration events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogGlobalConfigurationIterator struct {
	Event *StarkwarePerpetualsLogGlobalConfiguration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogGlobalConfigurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogGlobalConfiguration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogGlobalConfiguration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogGlobalConfigurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogGlobalConfigurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogGlobalConfiguration represents a LogGlobalConfiguration event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogGlobalConfiguration struct {
	ConfigHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogGlobalConfiguration is a free log retrieval operation binding the contract event 0xf90677ef35e0a4cde2563a22b999bac16426cded89b170434fe6295f2191afb2.
//
// Solidity: event LogGlobalConfiguration(bytes32 configHash)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogGlobalConfiguration(opts *bind.FilterOpts) (*StarkwarePerpetualsLogGlobalConfigurationIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogGlobalConfiguration")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogGlobalConfigurationIterator{contract: _StarkwarePerpetuals.contract, event: "LogGlobalConfiguration", logs: logs, sub: sub}, nil
}

// WatchLogGlobalConfiguration is a free log subscription operation binding the contract event 0xf90677ef35e0a4cde2563a22b999bac16426cded89b170434fe6295f2191afb2.
//
// Solidity: event LogGlobalConfiguration(bytes32 configHash)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogGlobalConfiguration(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogGlobalConfiguration) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogGlobalConfiguration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogGlobalConfiguration)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogGlobalConfiguration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogGlobalConfiguration is a log parse operation binding the contract event 0xf90677ef35e0a4cde2563a22b999bac16426cded89b170434fe6295f2191afb2.
//
// Solidity: event LogGlobalConfiguration(bytes32 configHash)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogGlobalConfiguration(log types.Log) (*StarkwarePerpetualsLogGlobalConfiguration, error) {
	event := new(StarkwarePerpetualsLogGlobalConfiguration)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogGlobalConfiguration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogItemConfigurationIterator is returned from FilterLogItemConfiguration and is used to iterate over the raw logs and unpacked data for LogItemConfiguration events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogItemConfigurationIterator struct {
	Event *StarkwarePerpetualsLogItemConfiguration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogItemConfigurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogItemConfiguration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogItemConfiguration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogItemConfigurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogItemConfigurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogItemConfiguration represents a LogItemConfiguration event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogItemConfiguration struct {
	ConfigItem *big.Int
	ConfigHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogItemConfiguration is a free log retrieval operation binding the contract event 0x2bd1a82da0f5450a3233065b7000191dabedfdb2d44279950af7fa32c31e7132.
//
// Solidity: event LogItemConfiguration(uint256 configItem, bytes32 configHash)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogItemConfiguration(opts *bind.FilterOpts) (*StarkwarePerpetualsLogItemConfigurationIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogItemConfiguration")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogItemConfigurationIterator{contract: _StarkwarePerpetuals.contract, event: "LogItemConfiguration", logs: logs, sub: sub}, nil
}

// WatchLogItemConfiguration is a free log subscription operation binding the contract event 0x2bd1a82da0f5450a3233065b7000191dabedfdb2d44279950af7fa32c31e7132.
//
// Solidity: event LogItemConfiguration(uint256 configItem, bytes32 configHash)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogItemConfiguration(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogItemConfiguration) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogItemConfiguration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogItemConfiguration)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogItemConfiguration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogItemConfiguration is a log parse operation binding the contract event 0x2bd1a82da0f5450a3233065b7000191dabedfdb2d44279950af7fa32c31e7132.
//
// Solidity: event LogItemConfiguration(uint256 configItem, bytes32 configHash)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogItemConfiguration(log types.Log) (*StarkwarePerpetualsLogItemConfiguration, error) {
	event := new(StarkwarePerpetualsLogItemConfiguration)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogItemConfiguration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogMintWithdrawalPerformedIterator is returned from FilterLogMintWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogMintWithdrawalPerformed events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogMintWithdrawalPerformedIterator struct {
	Event *StarkwarePerpetualsLogMintWithdrawalPerformed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogMintWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogMintWithdrawalPerformed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogMintWithdrawalPerformed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogMintWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogMintWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogMintWithdrawalPerformed represents a LogMintWithdrawalPerformed event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogMintWithdrawalPerformed struct {
	StarkKey           *big.Int
	TokenId            *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	AssetId            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogMintWithdrawalPerformed is a free log retrieval operation binding the contract event 0x7e6e15df814c1a309a57686de672b2bedd128eacde35c5370c36d6840d4e9a92.
//
// Solidity: event LogMintWithdrawalPerformed(uint256 starkKey, uint256 tokenId, uint256 nonQuantizedAmount, uint256 quantizedAmount, uint256 assetId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogMintWithdrawalPerformed(opts *bind.FilterOpts) (*StarkwarePerpetualsLogMintWithdrawalPerformedIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogMintWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogMintWithdrawalPerformedIterator{contract: _StarkwarePerpetuals.contract, event: "LogMintWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogMintWithdrawalPerformed is a free log subscription operation binding the contract event 0x7e6e15df814c1a309a57686de672b2bedd128eacde35c5370c36d6840d4e9a92.
//
// Solidity: event LogMintWithdrawalPerformed(uint256 starkKey, uint256 tokenId, uint256 nonQuantizedAmount, uint256 quantizedAmount, uint256 assetId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogMintWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogMintWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogMintWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogMintWithdrawalPerformed)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogMintWithdrawalPerformed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogMintWithdrawalPerformed is a log parse operation binding the contract event 0x7e6e15df814c1a309a57686de672b2bedd128eacde35c5370c36d6840d4e9a92.
//
// Solidity: event LogMintWithdrawalPerformed(uint256 starkKey, uint256 tokenId, uint256 nonQuantizedAmount, uint256 quantizedAmount, uint256 assetId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogMintWithdrawalPerformed(log types.Log) (*StarkwarePerpetualsLogMintWithdrawalPerformed, error) {
	event := new(StarkwarePerpetualsLogMintWithdrawalPerformed)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogMintWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogMintableWithdrawalAllowedIterator is returned from FilterLogMintableWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogMintableWithdrawalAllowed events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogMintableWithdrawalAllowedIterator struct {
	Event *StarkwarePerpetualsLogMintableWithdrawalAllowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogMintableWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogMintableWithdrawalAllowed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogMintableWithdrawalAllowed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogMintableWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogMintableWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogMintableWithdrawalAllowed represents a LogMintableWithdrawalAllowed event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogMintableWithdrawalAllowed struct {
	StarkKey        *big.Int
	AssetId         *big.Int
	QuantizedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogMintableWithdrawalAllowed is a free log retrieval operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 starkKey, uint256 assetId, uint256 quantizedAmount)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogMintableWithdrawalAllowed(opts *bind.FilterOpts) (*StarkwarePerpetualsLogMintableWithdrawalAllowedIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogMintableWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogMintableWithdrawalAllowedIterator{contract: _StarkwarePerpetuals.contract, event: "LogMintableWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogMintableWithdrawalAllowed is a free log subscription operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 starkKey, uint256 assetId, uint256 quantizedAmount)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogMintableWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogMintableWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogMintableWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogMintableWithdrawalAllowed)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogMintableWithdrawalAllowed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogMintableWithdrawalAllowed is a log parse operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 starkKey, uint256 assetId, uint256 quantizedAmount)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogMintableWithdrawalAllowed(log types.Log) (*StarkwarePerpetualsLogMintableWithdrawalAllowed, error) {
	event := new(StarkwarePerpetualsLogMintableWithdrawalAllowed)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogMintableWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogNewGovernorAcceptedIterator is returned from FilterLogNewGovernorAccepted and is used to iterate over the raw logs and unpacked data for LogNewGovernorAccepted events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogNewGovernorAcceptedIterator struct {
	Event *StarkwarePerpetualsLogNewGovernorAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogNewGovernorAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogNewGovernorAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogNewGovernorAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogNewGovernorAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogNewGovernorAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogNewGovernorAccepted represents a LogNewGovernorAccepted event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogNewGovernorAccepted struct {
	AcceptedGovernor common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogNewGovernorAccepted is a free log retrieval operation binding the contract event 0xcfb473e6c03f9a29ddaf990e736fa3de5188a0bd85d684f5b6e164ebfbfff5d2.
//
// Solidity: event LogNewGovernorAccepted(address acceptedGovernor)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogNewGovernorAccepted(opts *bind.FilterOpts) (*StarkwarePerpetualsLogNewGovernorAcceptedIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogNewGovernorAccepted")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogNewGovernorAcceptedIterator{contract: _StarkwarePerpetuals.contract, event: "LogNewGovernorAccepted", logs: logs, sub: sub}, nil
}

// WatchLogNewGovernorAccepted is a free log subscription operation binding the contract event 0xcfb473e6c03f9a29ddaf990e736fa3de5188a0bd85d684f5b6e164ebfbfff5d2.
//
// Solidity: event LogNewGovernorAccepted(address acceptedGovernor)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogNewGovernorAccepted(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogNewGovernorAccepted) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogNewGovernorAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogNewGovernorAccepted)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogNewGovernorAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNewGovernorAccepted is a log parse operation binding the contract event 0xcfb473e6c03f9a29ddaf990e736fa3de5188a0bd85d684f5b6e164ebfbfff5d2.
//
// Solidity: event LogNewGovernorAccepted(address acceptedGovernor)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogNewGovernorAccepted(log types.Log) (*StarkwarePerpetualsLogNewGovernorAccepted, error) {
	event := new(StarkwarePerpetualsLogNewGovernorAccepted)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogNewGovernorAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogNftDepositIterator is returned from FilterLogNftDeposit and is used to iterate over the raw logs and unpacked data for LogNftDeposit events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogNftDepositIterator struct {
	Event *StarkwarePerpetualsLogNftDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogNftDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogNftDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogNftDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogNftDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogNftDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogNftDeposit represents a LogNftDeposit event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogNftDeposit struct {
	DepositorEthKey common.Address
	StarkKey        *big.Int
	VaultId         *big.Int
	AssetType       *big.Int
	TokenId         *big.Int
	AssetId         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogNftDeposit is a free log retrieval operation binding the contract event 0x0fcf2162832b2d6033d4d34d2f45a28d9cfee523f1899945bbdd32529cfda67b.
//
// Solidity: event LogNftDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogNftDeposit(opts *bind.FilterOpts) (*StarkwarePerpetualsLogNftDepositIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogNftDeposit")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogNftDepositIterator{contract: _StarkwarePerpetuals.contract, event: "LogNftDeposit", logs: logs, sub: sub}, nil
}

// WatchLogNftDeposit is a free log subscription operation binding the contract event 0x0fcf2162832b2d6033d4d34d2f45a28d9cfee523f1899945bbdd32529cfda67b.
//
// Solidity: event LogNftDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogNftDeposit(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogNftDeposit) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogNftDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogNftDeposit)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogNftDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNftDeposit is a log parse operation binding the contract event 0x0fcf2162832b2d6033d4d34d2f45a28d9cfee523f1899945bbdd32529cfda67b.
//
// Solidity: event LogNftDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogNftDeposit(log types.Log) (*StarkwarePerpetualsLogNftDeposit, error) {
	event := new(StarkwarePerpetualsLogNftDeposit)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogNftDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogNftWithdrawalAllowedIterator is returned from FilterLogNftWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogNftWithdrawalAllowed events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogNftWithdrawalAllowedIterator struct {
	Event *StarkwarePerpetualsLogNftWithdrawalAllowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogNftWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogNftWithdrawalAllowed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogNftWithdrawalAllowed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogNftWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogNftWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogNftWithdrawalAllowed represents a LogNftWithdrawalAllowed event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogNftWithdrawalAllowed struct {
	StarkKey *big.Int
	AssetId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNftWithdrawalAllowed is a free log retrieval operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 starkKey, uint256 assetId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogNftWithdrawalAllowed(opts *bind.FilterOpts) (*StarkwarePerpetualsLogNftWithdrawalAllowedIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogNftWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogNftWithdrawalAllowedIterator{contract: _StarkwarePerpetuals.contract, event: "LogNftWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogNftWithdrawalAllowed is a free log subscription operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 starkKey, uint256 assetId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogNftWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogNftWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogNftWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogNftWithdrawalAllowed)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogNftWithdrawalAllowed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNftWithdrawalAllowed is a log parse operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 starkKey, uint256 assetId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogNftWithdrawalAllowed(log types.Log) (*StarkwarePerpetualsLogNftWithdrawalAllowed, error) {
	event := new(StarkwarePerpetualsLogNftWithdrawalAllowed)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogNftWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogNftWithdrawalPerformedIterator is returned from FilterLogNftWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogNftWithdrawalPerformed events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogNftWithdrawalPerformedIterator struct {
	Event *StarkwarePerpetualsLogNftWithdrawalPerformed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogNftWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogNftWithdrawalPerformed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogNftWithdrawalPerformed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogNftWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogNftWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogNftWithdrawalPerformed represents a LogNftWithdrawalPerformed event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogNftWithdrawalPerformed struct {
	StarkKey  *big.Int
	AssetType *big.Int
	TokenId   *big.Int
	AssetId   *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogNftWithdrawalPerformed is a free log retrieval operation binding the contract event 0xa5cfa8e2199ec5b8ca319288bcab72734207d30569756ee594a74b4df7abbf41.
//
// Solidity: event LogNftWithdrawalPerformed(uint256 starkKey, uint256 assetType, uint256 tokenId, uint256 assetId, address recipient)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogNftWithdrawalPerformed(opts *bind.FilterOpts) (*StarkwarePerpetualsLogNftWithdrawalPerformedIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogNftWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogNftWithdrawalPerformedIterator{contract: _StarkwarePerpetuals.contract, event: "LogNftWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogNftWithdrawalPerformed is a free log subscription operation binding the contract event 0xa5cfa8e2199ec5b8ca319288bcab72734207d30569756ee594a74b4df7abbf41.
//
// Solidity: event LogNftWithdrawalPerformed(uint256 starkKey, uint256 assetType, uint256 tokenId, uint256 assetId, address recipient)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogNftWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogNftWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogNftWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogNftWithdrawalPerformed)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogNftWithdrawalPerformed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNftWithdrawalPerformed is a log parse operation binding the contract event 0xa5cfa8e2199ec5b8ca319288bcab72734207d30569756ee594a74b4df7abbf41.
//
// Solidity: event LogNftWithdrawalPerformed(uint256 starkKey, uint256 assetType, uint256 tokenId, uint256 assetId, address recipient)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogNftWithdrawalPerformed(log types.Log) (*StarkwarePerpetualsLogNftWithdrawalPerformed, error) {
	event := new(StarkwarePerpetualsLogNftWithdrawalPerformed)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogNftWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogNominatedGovernorIterator is returned from FilterLogNominatedGovernor and is used to iterate over the raw logs and unpacked data for LogNominatedGovernor events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogNominatedGovernorIterator struct {
	Event *StarkwarePerpetualsLogNominatedGovernor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogNominatedGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogNominatedGovernor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogNominatedGovernor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogNominatedGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogNominatedGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogNominatedGovernor represents a LogNominatedGovernor event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogNominatedGovernor struct {
	NominatedGovernor common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogNominatedGovernor is a free log retrieval operation binding the contract event 0x6166272c8d3f5f579082f2827532732f97195007983bb5b83ac12c56700b01a6.
//
// Solidity: event LogNominatedGovernor(address nominatedGovernor)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogNominatedGovernor(opts *bind.FilterOpts) (*StarkwarePerpetualsLogNominatedGovernorIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogNominatedGovernor")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogNominatedGovernorIterator{contract: _StarkwarePerpetuals.contract, event: "LogNominatedGovernor", logs: logs, sub: sub}, nil
}

// WatchLogNominatedGovernor is a free log subscription operation binding the contract event 0x6166272c8d3f5f579082f2827532732f97195007983bb5b83ac12c56700b01a6.
//
// Solidity: event LogNominatedGovernor(address nominatedGovernor)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogNominatedGovernor(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogNominatedGovernor) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogNominatedGovernor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogNominatedGovernor)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogNominatedGovernor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNominatedGovernor is a log parse operation binding the contract event 0x6166272c8d3f5f579082f2827532732f97195007983bb5b83ac12c56700b01a6.
//
// Solidity: event LogNominatedGovernor(address nominatedGovernor)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogNominatedGovernor(log types.Log) (*StarkwarePerpetualsLogNominatedGovernor, error) {
	event := new(StarkwarePerpetualsLogNominatedGovernor)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogNominatedGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogNominationCancelledIterator is returned from FilterLogNominationCancelled and is used to iterate over the raw logs and unpacked data for LogNominationCancelled events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogNominationCancelledIterator struct {
	Event *StarkwarePerpetualsLogNominationCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogNominationCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogNominationCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogNominationCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogNominationCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogNominationCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogNominationCancelled represents a LogNominationCancelled event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogNominationCancelled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNominationCancelled is a free log retrieval operation binding the contract event 0x7a8dc7dd7fffb43c4807438fa62729225156941e641fd877938f4edade3429f5.
//
// Solidity: event LogNominationCancelled()
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogNominationCancelled(opts *bind.FilterOpts) (*StarkwarePerpetualsLogNominationCancelledIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogNominationCancelled")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogNominationCancelledIterator{contract: _StarkwarePerpetuals.contract, event: "LogNominationCancelled", logs: logs, sub: sub}, nil
}

// WatchLogNominationCancelled is a free log subscription operation binding the contract event 0x7a8dc7dd7fffb43c4807438fa62729225156941e641fd877938f4edade3429f5.
//
// Solidity: event LogNominationCancelled()
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogNominationCancelled(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogNominationCancelled) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogNominationCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogNominationCancelled)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogNominationCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNominationCancelled is a log parse operation binding the contract event 0x7a8dc7dd7fffb43c4807438fa62729225156941e641fd877938f4edade3429f5.
//
// Solidity: event LogNominationCancelled()
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogNominationCancelled(log types.Log) (*StarkwarePerpetualsLogNominationCancelled, error) {
	event := new(StarkwarePerpetualsLogNominationCancelled)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogNominationCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogOperatorAddedIterator is returned from FilterLogOperatorAdded and is used to iterate over the raw logs and unpacked data for LogOperatorAdded events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogOperatorAddedIterator struct {
	Event *StarkwarePerpetualsLogOperatorAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogOperatorAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogOperatorAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogOperatorAdded represents a LogOperatorAdded event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogOperatorAdded struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogOperatorAdded is a free log retrieval operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogOperatorAdded(opts *bind.FilterOpts) (*StarkwarePerpetualsLogOperatorAddedIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogOperatorAdded")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogOperatorAddedIterator{contract: _StarkwarePerpetuals.contract, event: "LogOperatorAdded", logs: logs, sub: sub}, nil
}

// WatchLogOperatorAdded is a free log subscription operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogOperatorAdded(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogOperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogOperatorAdded)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogOperatorAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogOperatorAdded is a log parse operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogOperatorAdded(log types.Log) (*StarkwarePerpetualsLogOperatorAdded, error) {
	event := new(StarkwarePerpetualsLogOperatorAdded)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogOperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogOperatorRemovedIterator is returned from FilterLogOperatorRemoved and is used to iterate over the raw logs and unpacked data for LogOperatorRemoved events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogOperatorRemovedIterator struct {
	Event *StarkwarePerpetualsLogOperatorRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogOperatorRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogOperatorRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogOperatorRemoved represents a LogOperatorRemoved event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogOperatorRemoved struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogOperatorRemoved is a free log retrieval operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogOperatorRemoved(opts *bind.FilterOpts) (*StarkwarePerpetualsLogOperatorRemovedIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogOperatorRemoved")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogOperatorRemovedIterator{contract: _StarkwarePerpetuals.contract, event: "LogOperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchLogOperatorRemoved is a free log subscription operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogOperatorRemoved(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogOperatorRemoved) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogOperatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogOperatorRemoved)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogOperatorRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogOperatorRemoved is a log parse operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogOperatorRemoved(log types.Log) (*StarkwarePerpetualsLogOperatorRemoved, error) {
	event := new(StarkwarePerpetualsLogOperatorRemoved)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogOperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogRegisteredIterator is returned from FilterLogRegistered and is used to iterate over the raw logs and unpacked data for LogRegistered events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogRegisteredIterator struct {
	Event *StarkwarePerpetualsLogRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogRegistered represents a LogRegistered event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogRegistered struct {
	Entry   common.Address
	EntryId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogRegistered is a free log retrieval operation binding the contract event 0xaa7f29c97c6763919ef56006f07fbf5c1ac734b0414665df43cecdbce9010c9b.
//
// Solidity: event LogRegistered(address entry, string entryId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogRegistered(opts *bind.FilterOpts) (*StarkwarePerpetualsLogRegisteredIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogRegistered")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogRegisteredIterator{contract: _StarkwarePerpetuals.contract, event: "LogRegistered", logs: logs, sub: sub}, nil
}

// WatchLogRegistered is a free log subscription operation binding the contract event 0xaa7f29c97c6763919ef56006f07fbf5c1ac734b0414665df43cecdbce9010c9b.
//
// Solidity: event LogRegistered(address entry, string entryId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogRegistered(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogRegistered) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogRegistered)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogRegistered is a log parse operation binding the contract event 0xaa7f29c97c6763919ef56006f07fbf5c1ac734b0414665df43cecdbce9010c9b.
//
// Solidity: event LogRegistered(address entry, string entryId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogRegistered(log types.Log) (*StarkwarePerpetualsLogRegistered, error) {
	event := new(StarkwarePerpetualsLogRegistered)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogRemovalIntentIterator is returned from FilterLogRemovalIntent and is used to iterate over the raw logs and unpacked data for LogRemovalIntent events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogRemovalIntentIterator struct {
	Event *StarkwarePerpetualsLogRemovalIntent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogRemovalIntentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogRemovalIntent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogRemovalIntent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogRemovalIntentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogRemovalIntentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogRemovalIntent represents a LogRemovalIntent event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogRemovalIntent struct {
	Entry   common.Address
	EntryId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogRemovalIntent is a free log retrieval operation binding the contract event 0xa98ac1f696177f16ca482653307aa4a02b68cf03b14409a546de5adf946484fc.
//
// Solidity: event LogRemovalIntent(address entry, string entryId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogRemovalIntent(opts *bind.FilterOpts) (*StarkwarePerpetualsLogRemovalIntentIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogRemovalIntent")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogRemovalIntentIterator{contract: _StarkwarePerpetuals.contract, event: "LogRemovalIntent", logs: logs, sub: sub}, nil
}

// WatchLogRemovalIntent is a free log subscription operation binding the contract event 0xa98ac1f696177f16ca482653307aa4a02b68cf03b14409a546de5adf946484fc.
//
// Solidity: event LogRemovalIntent(address entry, string entryId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogRemovalIntent(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogRemovalIntent) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogRemovalIntent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogRemovalIntent)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogRemovalIntent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogRemovalIntent is a log parse operation binding the contract event 0xa98ac1f696177f16ca482653307aa4a02b68cf03b14409a546de5adf946484fc.
//
// Solidity: event LogRemovalIntent(address entry, string entryId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogRemovalIntent(log types.Log) (*StarkwarePerpetualsLogRemovalIntent, error) {
	event := new(StarkwarePerpetualsLogRemovalIntent)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogRemovalIntent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogRemovedIterator is returned from FilterLogRemoved and is used to iterate over the raw logs and unpacked data for LogRemoved events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogRemovedIterator struct {
	Event *StarkwarePerpetualsLogRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogRemoved represents a LogRemoved event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogRemoved struct {
	Entry   common.Address
	EntryId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogRemoved is a free log retrieval operation binding the contract event 0x35b176cf4f09df896aa55e10eec90bb4c4689ea1d076135a8de3138d829d0142.
//
// Solidity: event LogRemoved(address entry, string entryId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogRemoved(opts *bind.FilterOpts) (*StarkwarePerpetualsLogRemovedIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogRemoved")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogRemovedIterator{contract: _StarkwarePerpetuals.contract, event: "LogRemoved", logs: logs, sub: sub}, nil
}

// WatchLogRemoved is a free log subscription operation binding the contract event 0x35b176cf4f09df896aa55e10eec90bb4c4689ea1d076135a8de3138d829d0142.
//
// Solidity: event LogRemoved(address entry, string entryId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogRemoved(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogRemoved) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogRemoved)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogRemoved is a log parse operation binding the contract event 0x35b176cf4f09df896aa55e10eec90bb4c4689ea1d076135a8de3138d829d0142.
//
// Solidity: event LogRemoved(address entry, string entryId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogRemoved(log types.Log) (*StarkwarePerpetualsLogRemoved, error) {
	event := new(StarkwarePerpetualsLogRemoved)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogRemovedGovernorIterator is returned from FilterLogRemovedGovernor and is used to iterate over the raw logs and unpacked data for LogRemovedGovernor events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogRemovedGovernorIterator struct {
	Event *StarkwarePerpetualsLogRemovedGovernor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogRemovedGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogRemovedGovernor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogRemovedGovernor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogRemovedGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogRemovedGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogRemovedGovernor represents a LogRemovedGovernor event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogRemovedGovernor struct {
	RemovedGovernor common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogRemovedGovernor is a free log retrieval operation binding the contract event 0xd75f94825e770b8b512be8e74759e252ad00e102e38f50cce2f7c6f868a29599.
//
// Solidity: event LogRemovedGovernor(address removedGovernor)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogRemovedGovernor(opts *bind.FilterOpts) (*StarkwarePerpetualsLogRemovedGovernorIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogRemovedGovernor")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogRemovedGovernorIterator{contract: _StarkwarePerpetuals.contract, event: "LogRemovedGovernor", logs: logs, sub: sub}, nil
}

// WatchLogRemovedGovernor is a free log subscription operation binding the contract event 0xd75f94825e770b8b512be8e74759e252ad00e102e38f50cce2f7c6f868a29599.
//
// Solidity: event LogRemovedGovernor(address removedGovernor)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogRemovedGovernor(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogRemovedGovernor) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogRemovedGovernor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogRemovedGovernor)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogRemovedGovernor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogRemovedGovernor is a log parse operation binding the contract event 0xd75f94825e770b8b512be8e74759e252ad00e102e38f50cce2f7c6f868a29599.
//
// Solidity: event LogRemovedGovernor(address removedGovernor)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogRemovedGovernor(log types.Log) (*StarkwarePerpetualsLogRemovedGovernor, error) {
	event := new(StarkwarePerpetualsLogRemovedGovernor)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogRemovedGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogStateTransitionFactIterator is returned from FilterLogStateTransitionFact and is used to iterate over the raw logs and unpacked data for LogStateTransitionFact events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogStateTransitionFactIterator struct {
	Event *StarkwarePerpetualsLogStateTransitionFact // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogStateTransitionFactIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogStateTransitionFact)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogStateTransitionFact)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogStateTransitionFactIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogStateTransitionFactIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogStateTransitionFact represents a LogStateTransitionFact event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogStateTransitionFact struct {
	StateTransitionFact [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLogStateTransitionFact is a free log retrieval operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogStateTransitionFact(opts *bind.FilterOpts) (*StarkwarePerpetualsLogStateTransitionFactIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogStateTransitionFact")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogStateTransitionFactIterator{contract: _StarkwarePerpetuals.contract, event: "LogStateTransitionFact", logs: logs, sub: sub}, nil
}

// WatchLogStateTransitionFact is a free log subscription operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogStateTransitionFact(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogStateTransitionFact) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogStateTransitionFact")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogStateTransitionFact)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogStateTransitionFact", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogStateTransitionFact is a log parse operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogStateTransitionFact(log types.Log) (*StarkwarePerpetualsLogStateTransitionFact, error) {
	event := new(StarkwarePerpetualsLogStateTransitionFact)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogStateTransitionFact", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogSystemAssetTypeIterator is returned from FilterLogSystemAssetType and is used to iterate over the raw logs and unpacked data for LogSystemAssetType events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogSystemAssetTypeIterator struct {
	Event *StarkwarePerpetualsLogSystemAssetType // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogSystemAssetTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogSystemAssetType)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogSystemAssetType)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogSystemAssetTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogSystemAssetTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogSystemAssetType represents a LogSystemAssetType event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogSystemAssetType struct {
	AssetType *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSystemAssetType is a free log retrieval operation binding the contract event 0x51f48293a5ef1940e2b4eb2580372cf384aaa5bc739639e4624ce8d18c9644ab.
//
// Solidity: event LogSystemAssetType(uint256 assetType)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogSystemAssetType(opts *bind.FilterOpts) (*StarkwarePerpetualsLogSystemAssetTypeIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogSystemAssetType")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogSystemAssetTypeIterator{contract: _StarkwarePerpetuals.contract, event: "LogSystemAssetType", logs: logs, sub: sub}, nil
}

// WatchLogSystemAssetType is a free log subscription operation binding the contract event 0x51f48293a5ef1940e2b4eb2580372cf384aaa5bc739639e4624ce8d18c9644ab.
//
// Solidity: event LogSystemAssetType(uint256 assetType)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogSystemAssetType(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogSystemAssetType) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogSystemAssetType")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogSystemAssetType)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogSystemAssetType", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogSystemAssetType is a log parse operation binding the contract event 0x51f48293a5ef1940e2b4eb2580372cf384aaa5bc739639e4624ce8d18c9644ab.
//
// Solidity: event LogSystemAssetType(uint256 assetType)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogSystemAssetType(log types.Log) (*StarkwarePerpetualsLogSystemAssetType, error) {
	event := new(StarkwarePerpetualsLogSystemAssetType)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogSystemAssetType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogTokenAdminAddedIterator is returned from FilterLogTokenAdminAdded and is used to iterate over the raw logs and unpacked data for LogTokenAdminAdded events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogTokenAdminAddedIterator struct {
	Event *StarkwarePerpetualsLogTokenAdminAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogTokenAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogTokenAdminAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogTokenAdminAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogTokenAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogTokenAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogTokenAdminAdded represents a LogTokenAdminAdded event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogTokenAdminAdded struct {
	TokenAdmin common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogTokenAdminAdded is a free log retrieval operation binding the contract event 0x9085a9044aeb6daeeb5b4bf84af42b1a1613d4056f503c4e992b6396c16bd52f.
//
// Solidity: event LogTokenAdminAdded(address tokenAdmin)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogTokenAdminAdded(opts *bind.FilterOpts) (*StarkwarePerpetualsLogTokenAdminAddedIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogTokenAdminAdded")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogTokenAdminAddedIterator{contract: _StarkwarePerpetuals.contract, event: "LogTokenAdminAdded", logs: logs, sub: sub}, nil
}

// WatchLogTokenAdminAdded is a free log subscription operation binding the contract event 0x9085a9044aeb6daeeb5b4bf84af42b1a1613d4056f503c4e992b6396c16bd52f.
//
// Solidity: event LogTokenAdminAdded(address tokenAdmin)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogTokenAdminAdded(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogTokenAdminAdded) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogTokenAdminAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogTokenAdminAdded)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogTokenAdminAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogTokenAdminAdded is a log parse operation binding the contract event 0x9085a9044aeb6daeeb5b4bf84af42b1a1613d4056f503c4e992b6396c16bd52f.
//
// Solidity: event LogTokenAdminAdded(address tokenAdmin)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogTokenAdminAdded(log types.Log) (*StarkwarePerpetualsLogTokenAdminAdded, error) {
	event := new(StarkwarePerpetualsLogTokenAdminAdded)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogTokenAdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogTokenAdminRemovedIterator is returned from FilterLogTokenAdminRemoved and is used to iterate over the raw logs and unpacked data for LogTokenAdminRemoved events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogTokenAdminRemovedIterator struct {
	Event *StarkwarePerpetualsLogTokenAdminRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogTokenAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogTokenAdminRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogTokenAdminRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogTokenAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogTokenAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogTokenAdminRemoved represents a LogTokenAdminRemoved event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogTokenAdminRemoved struct {
	TokenAdmin common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogTokenAdminRemoved is a free log retrieval operation binding the contract event 0xfa49aecb996ea8d99950bb051552dfcc0b5460a0bb209867a1ed8067c32c2177.
//
// Solidity: event LogTokenAdminRemoved(address tokenAdmin)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogTokenAdminRemoved(opts *bind.FilterOpts) (*StarkwarePerpetualsLogTokenAdminRemovedIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogTokenAdminRemoved")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogTokenAdminRemovedIterator{contract: _StarkwarePerpetuals.contract, event: "LogTokenAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchLogTokenAdminRemoved is a free log subscription operation binding the contract event 0xfa49aecb996ea8d99950bb051552dfcc0b5460a0bb209867a1ed8067c32c2177.
//
// Solidity: event LogTokenAdminRemoved(address tokenAdmin)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogTokenAdminRemoved(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogTokenAdminRemoved) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogTokenAdminRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogTokenAdminRemoved)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogTokenAdminRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogTokenAdminRemoved is a log parse operation binding the contract event 0xfa49aecb996ea8d99950bb051552dfcc0b5460a0bb209867a1ed8067c32c2177.
//
// Solidity: event LogTokenAdminRemoved(address tokenAdmin)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogTokenAdminRemoved(log types.Log) (*StarkwarePerpetualsLogTokenAdminRemoved, error) {
	event := new(StarkwarePerpetualsLogTokenAdminRemoved)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogTokenAdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogTokenRegisteredIterator is returned from FilterLogTokenRegistered and is used to iterate over the raw logs and unpacked data for LogTokenRegistered events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogTokenRegisteredIterator struct {
	Event *StarkwarePerpetualsLogTokenRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogTokenRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogTokenRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogTokenRegistered represents a LogTokenRegistered event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogTokenRegistered struct {
	AssetType *big.Int
	AssetInfo []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogTokenRegistered is a free log retrieval operation binding the contract event 0x4d2c7bfd8df1ba4f331f1abd2562bf3088e8b378c7dd1308113a82c64e518dbf.
//
// Solidity: event LogTokenRegistered(uint256 assetType, bytes assetInfo)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogTokenRegistered(opts *bind.FilterOpts) (*StarkwarePerpetualsLogTokenRegisteredIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogTokenRegistered")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogTokenRegisteredIterator{contract: _StarkwarePerpetuals.contract, event: "LogTokenRegistered", logs: logs, sub: sub}, nil
}

// WatchLogTokenRegistered is a free log subscription operation binding the contract event 0x4d2c7bfd8df1ba4f331f1abd2562bf3088e8b378c7dd1308113a82c64e518dbf.
//
// Solidity: event LogTokenRegistered(uint256 assetType, bytes assetInfo)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogTokenRegistered(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogTokenRegistered) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogTokenRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogTokenRegistered)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogTokenRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogTokenRegistered is a log parse operation binding the contract event 0x4d2c7bfd8df1ba4f331f1abd2562bf3088e8b378c7dd1308113a82c64e518dbf.
//
// Solidity: event LogTokenRegistered(uint256 assetType, bytes assetInfo)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogTokenRegistered(log types.Log) (*StarkwarePerpetualsLogTokenRegistered, error) {
	event := new(StarkwarePerpetualsLogTokenRegistered)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogTokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogUnFrozenIterator is returned from FilterLogUnFrozen and is used to iterate over the raw logs and unpacked data for LogUnFrozen events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogUnFrozenIterator struct {
	Event *StarkwarePerpetualsLogUnFrozen // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogUnFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogUnFrozen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogUnFrozen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogUnFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogUnFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogUnFrozen represents a LogUnFrozen event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogUnFrozen struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogUnFrozen is a free log retrieval operation binding the contract event 0x07017fe9180629cfffba412f65a9affcf9a121de02294179f5c058f881dcc9f8.
//
// Solidity: event LogUnFrozen()
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogUnFrozen(opts *bind.FilterOpts) (*StarkwarePerpetualsLogUnFrozenIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogUnFrozen")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogUnFrozenIterator{contract: _StarkwarePerpetuals.contract, event: "LogUnFrozen", logs: logs, sub: sub}, nil
}

// WatchLogUnFrozen is a free log subscription operation binding the contract event 0x07017fe9180629cfffba412f65a9affcf9a121de02294179f5c058f881dcc9f8.
//
// Solidity: event LogUnFrozen()
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogUnFrozen(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogUnFrozen) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogUnFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogUnFrozen)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogUnFrozen", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogUnFrozen is a log parse operation binding the contract event 0x07017fe9180629cfffba412f65a9affcf9a121de02294179f5c058f881dcc9f8.
//
// Solidity: event LogUnFrozen()
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogUnFrozen(log types.Log) (*StarkwarePerpetualsLogUnFrozen, error) {
	event := new(StarkwarePerpetualsLogUnFrozen)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogUnFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogUpdateStateIterator is returned from FilterLogUpdateState and is used to iterate over the raw logs and unpacked data for LogUpdateState events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogUpdateStateIterator struct {
	Event *StarkwarePerpetualsLogUpdateState // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogUpdateStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogUpdateState)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogUpdateState)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogUpdateStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogUpdateStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogUpdateState represents a LogUpdateState event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogUpdateState struct {
	SequenceNumber *big.Int
	BatchId        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogUpdateState is a free log retrieval operation binding the contract event 0x2672b53d25204094519f7b0fba8d2b5cd0cc1e426f49554c89461cdb9dcec08f.
//
// Solidity: event LogUpdateState(uint256 sequenceNumber, uint256 batchId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogUpdateState(opts *bind.FilterOpts) (*StarkwarePerpetualsLogUpdateStateIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogUpdateState")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogUpdateStateIterator{contract: _StarkwarePerpetuals.contract, event: "LogUpdateState", logs: logs, sub: sub}, nil
}

// WatchLogUpdateState is a free log subscription operation binding the contract event 0x2672b53d25204094519f7b0fba8d2b5cd0cc1e426f49554c89461cdb9dcec08f.
//
// Solidity: event LogUpdateState(uint256 sequenceNumber, uint256 batchId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogUpdateState(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogUpdateState) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogUpdateState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogUpdateState)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogUpdateState", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogUpdateState is a log parse operation binding the contract event 0x2672b53d25204094519f7b0fba8d2b5cd0cc1e426f49554c89461cdb9dcec08f.
//
// Solidity: event LogUpdateState(uint256 sequenceNumber, uint256 batchId)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogUpdateState(log types.Log) (*StarkwarePerpetualsLogUpdateState, error) {
	event := new(StarkwarePerpetualsLogUpdateState)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogUpdateState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogUserAdminAddedIterator is returned from FilterLogUserAdminAdded and is used to iterate over the raw logs and unpacked data for LogUserAdminAdded events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogUserAdminAddedIterator struct {
	Event *StarkwarePerpetualsLogUserAdminAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogUserAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogUserAdminAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogUserAdminAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogUserAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogUserAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogUserAdminAdded represents a LogUserAdminAdded event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogUserAdminAdded struct {
	UserAdmin common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogUserAdminAdded is a free log retrieval operation binding the contract event 0x7284e8b42a1333a4f23e858e513b3b28d2667a3531b7c1872cce3f7720a25046.
//
// Solidity: event LogUserAdminAdded(address userAdmin)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogUserAdminAdded(opts *bind.FilterOpts) (*StarkwarePerpetualsLogUserAdminAddedIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogUserAdminAdded")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogUserAdminAddedIterator{contract: _StarkwarePerpetuals.contract, event: "LogUserAdminAdded", logs: logs, sub: sub}, nil
}

// WatchLogUserAdminAdded is a free log subscription operation binding the contract event 0x7284e8b42a1333a4f23e858e513b3b28d2667a3531b7c1872cce3f7720a25046.
//
// Solidity: event LogUserAdminAdded(address userAdmin)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogUserAdminAdded(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogUserAdminAdded) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogUserAdminAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogUserAdminAdded)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogUserAdminAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogUserAdminAdded is a log parse operation binding the contract event 0x7284e8b42a1333a4f23e858e513b3b28d2667a3531b7c1872cce3f7720a25046.
//
// Solidity: event LogUserAdminAdded(address userAdmin)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogUserAdminAdded(log types.Log) (*StarkwarePerpetualsLogUserAdminAdded, error) {
	event := new(StarkwarePerpetualsLogUserAdminAdded)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogUserAdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogUserAdminRemovedIterator is returned from FilterLogUserAdminRemoved and is used to iterate over the raw logs and unpacked data for LogUserAdminRemoved events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogUserAdminRemovedIterator struct {
	Event *StarkwarePerpetualsLogUserAdminRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogUserAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogUserAdminRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogUserAdminRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogUserAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogUserAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogUserAdminRemoved represents a LogUserAdminRemoved event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogUserAdminRemoved struct {
	UserAdmin common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogUserAdminRemoved is a free log retrieval operation binding the contract event 0xb32f8aed6bedf93605e95bc99e0e229b8bbfcd0fe2e76a6748450d3e9522db46.
//
// Solidity: event LogUserAdminRemoved(address userAdmin)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogUserAdminRemoved(opts *bind.FilterOpts) (*StarkwarePerpetualsLogUserAdminRemovedIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogUserAdminRemoved")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogUserAdminRemovedIterator{contract: _StarkwarePerpetuals.contract, event: "LogUserAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchLogUserAdminRemoved is a free log subscription operation binding the contract event 0xb32f8aed6bedf93605e95bc99e0e229b8bbfcd0fe2e76a6748450d3e9522db46.
//
// Solidity: event LogUserAdminRemoved(address userAdmin)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogUserAdminRemoved(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogUserAdminRemoved) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogUserAdminRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogUserAdminRemoved)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogUserAdminRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogUserAdminRemoved is a log parse operation binding the contract event 0xb32f8aed6bedf93605e95bc99e0e229b8bbfcd0fe2e76a6748450d3e9522db46.
//
// Solidity: event LogUserAdminRemoved(address userAdmin)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogUserAdminRemoved(log types.Log) (*StarkwarePerpetualsLogUserAdminRemoved, error) {
	event := new(StarkwarePerpetualsLogUserAdminRemoved)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogUserAdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogUserRegisteredIterator is returned from FilterLogUserRegistered and is used to iterate over the raw logs and unpacked data for LogUserRegistered events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogUserRegisteredIterator struct {
	Event *StarkwarePerpetualsLogUserRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogUserRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogUserRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogUserRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogUserRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogUserRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogUserRegistered represents a LogUserRegistered event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogUserRegistered struct {
	EthKey   common.Address
	StarkKey *big.Int
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogUserRegistered is a free log retrieval operation binding the contract event 0xcab1cf17c190e4e2195a7b8f7b362023246fa774390432b4704ab6b29d56b07b.
//
// Solidity: event LogUserRegistered(address ethKey, uint256 starkKey, address sender)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogUserRegistered(opts *bind.FilterOpts) (*StarkwarePerpetualsLogUserRegisteredIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogUserRegistered")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogUserRegisteredIterator{contract: _StarkwarePerpetuals.contract, event: "LogUserRegistered", logs: logs, sub: sub}, nil
}

// WatchLogUserRegistered is a free log subscription operation binding the contract event 0xcab1cf17c190e4e2195a7b8f7b362023246fa774390432b4704ab6b29d56b07b.
//
// Solidity: event LogUserRegistered(address ethKey, uint256 starkKey, address sender)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogUserRegistered(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogUserRegistered) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogUserRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogUserRegistered)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogUserRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogUserRegistered is a log parse operation binding the contract event 0xcab1cf17c190e4e2195a7b8f7b362023246fa774390432b4704ab6b29d56b07b.
//
// Solidity: event LogUserRegistered(address ethKey, uint256 starkKey, address sender)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogUserRegistered(log types.Log) (*StarkwarePerpetualsLogUserRegistered, error) {
	event := new(StarkwarePerpetualsLogUserRegistered)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogUserRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogWithdrawalAllowedIterator is returned from FilterLogWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogWithdrawalAllowed events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogWithdrawalAllowedIterator struct {
	Event *StarkwarePerpetualsLogWithdrawalAllowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogWithdrawalAllowed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogWithdrawalAllowed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogWithdrawalAllowed represents a LogWithdrawalAllowed event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogWithdrawalAllowed struct {
	StarkKey           *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawalAllowed is a free log retrieval operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 starkKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogWithdrawalAllowed(opts *bind.FilterOpts) (*StarkwarePerpetualsLogWithdrawalAllowedIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogWithdrawalAllowedIterator{contract: _StarkwarePerpetuals.contract, event: "LogWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawalAllowed is a free log subscription operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 starkKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogWithdrawalAllowed)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogWithdrawalAllowed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogWithdrawalAllowed is a log parse operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 starkKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogWithdrawalAllowed(log types.Log) (*StarkwarePerpetualsLogWithdrawalAllowed, error) {
	event := new(StarkwarePerpetualsLogWithdrawalAllowed)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StarkwarePerpetualsLogWithdrawalPerformedIterator is returned from FilterLogWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogWithdrawalPerformed events raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogWithdrawalPerformedIterator struct {
	Event *StarkwarePerpetualsLogWithdrawalPerformed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StarkwarePerpetualsLogWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarkwarePerpetualsLogWithdrawalPerformed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StarkwarePerpetualsLogWithdrawalPerformed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StarkwarePerpetualsLogWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarkwarePerpetualsLogWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarkwarePerpetualsLogWithdrawalPerformed represents a LogWithdrawalPerformed event raised by the StarkwarePerpetuals contract.
type StarkwarePerpetualsLogWithdrawalPerformed struct {
	StarkKey           *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Recipient          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawalPerformed is a free log retrieval operation binding the contract event 0xb7477a7b93b2addc5272bbd7ad0986ef1c0d0bd265f26c3dc4bbe42727c2ac0c.
//
// Solidity: event LogWithdrawalPerformed(uint256 starkKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) FilterLogWithdrawalPerformed(opts *bind.FilterOpts) (*StarkwarePerpetualsLogWithdrawalPerformedIterator, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.FilterLogs(opts, "LogWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &StarkwarePerpetualsLogWithdrawalPerformedIterator{contract: _StarkwarePerpetuals.contract, event: "LogWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawalPerformed is a free log subscription operation binding the contract event 0xb7477a7b93b2addc5272bbd7ad0986ef1c0d0bd265f26c3dc4bbe42727c2ac0c.
//
// Solidity: event LogWithdrawalPerformed(uint256 starkKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) WatchLogWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *StarkwarePerpetualsLogWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _StarkwarePerpetuals.contract.WatchLogs(opts, "LogWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarkwarePerpetualsLogWithdrawalPerformed)
				if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogWithdrawalPerformed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogWithdrawalPerformed is a log parse operation binding the contract event 0xb7477a7b93b2addc5272bbd7ad0986ef1c0d0bd265f26c3dc4bbe42727c2ac0c.
//
// Solidity: event LogWithdrawalPerformed(uint256 starkKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_StarkwarePerpetuals *StarkwarePerpetualsFilterer) ParseLogWithdrawalPerformed(log types.Log) (*StarkwarePerpetualsLogWithdrawalPerformed, error) {
	event := new(StarkwarePerpetualsLogWithdrawalPerformed)
	if err := _StarkwarePerpetuals.contract.UnpackLog(event, "LogWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
