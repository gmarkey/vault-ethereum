// Copyright (C) Immutability, LLC - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Jeff Ploughman <jeff@immutability.io>, August 2019

package main

import (
        "bytes"
        "context"
        "fmt"
        "math/big"

        "github.com/ethereum/go-ethereum/accounts/abi/bind"
        "github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/common/hexutil"
        "github.com/ethereum/go-ethereum/ethclient"
        "github.com/ethereum/go-ethereum/common/math"
        "github.com/hashicorp/vault/sdk/framework"
        "github.com/hashicorp/vault/sdk/logical"
        "github.com/gmarkey/vault-ethereum/contracts/starkwarePerpetuals"
        "github.com/gmarkey/vault-ethereum/util"
)

const dydx string = "dydx"

// contract ERC20Interface {
//     string public constant name = "";
//     string public constant symbol = "";
//     uint8 public constant decimals = 0;

//     function totalSupply() public view returns (uint);
//     function balanceOf(address tokenOwner) public view returns (uint balance);
//     function allowance(address tokenOwner, address spender) public view returns (uint remaining);
//     function transfer(address to, uint tokens) public returns (bool success);
//     function approve(address spender, uint tokens) public returns (bool success);
//     function transferFrom(address from, address to, uint tokens) public returns (bool success);

//     event Transfer(address indexed from, address indexed to, uint tokens);
//     event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
// }

func dydxPaths(b *PluginBackend) []*framework.Path {
        return []*framework.Path{
                {
                        Pattern:      ContractPath(dydx, "deposit"),
                        HelpSynopsis: "Deposit collateral into dydx account",
                        HelpDescription: `

Deposit collateral into dydx account.

`,
                        Fields: map[string]*framework.FieldSchema{
                                "name": {Type: framework.TypeString},
                                "stark_key": {
                                        Type:        framework.TypeString,
                                        Description: "The starkware API key.",
                                },
                                "vault_id": {
                                        Type:        framework.TypeString,
                                        Description: "Account into which we are depositing.",
                                },
                                "tokens": {
                                        Type:        framework.TypeString,
                                        Default:     "0",
                                        Description: "The number of tokens to transfer.",
                                },
                                "asset_type": {
                                        Type:        framework.TypeString,
                                        Default:     "0x8707A5bf4C2842d46B31A405Ba41b858C0F876c4",
                                        Description: "Address of the deposit assed token contract.",
                                },
                                "contract": {
                                        Type:        framework.TypeString,
                                        Default:     "0x014F738EAd8Ec6C50BCD456a971F8B84Cd693BBe",
                                        Description: "Address of the destination contract.",
                                },
                        },
                        ExistenceCheck: pathExistenceCheck,
                        Callbacks: map[logical.Operation]framework.OperationFunc{
                                logical.CreateOperation: b.pathDydxDeposit,
                                logical.UpdateOperation: b.pathDydxDeposit,
                        },
                },
        }
}

func (b *PluginBackend) pathDydxDeposit(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
        var tokens *big.Int
        var starkKey *big.Int
        var vaultId *big.Int
        var assetType *big.Int
        //assetType := common.HexToAddress("0x8707a5bf4c2842d46b31a405ba41b858c0f876c4") // ropsten
        //assetType := util.ValidNumber("0x8707a5bf4c2842d46b31a405ba41b858c0f876c4") // ropsten

        config, err := b.configured(ctx, req)
        if err != nil {
                return nil, err
        }

        name := data.Get("name").(string)
        // contractAddress := "0x014F738EAd8Ec6C50BCD456a971F8B84Cd693BBe"
        contractAddress := data.Get("contract").(string)

        accountJSON, err := readAccount(ctx, req, name)
        if err != nil {
                return nil, err
        }
        wallet, account, err := getWalletAndAccount(*accountJSON)
        if err != nil {
                return nil, err
        }

        _, ok := data.GetOk("vault_id")
        if ok {
                vaultId = util.ValidNumber(data.Get("vault_id").(string))
                if vaultId == nil {
                        return nil, fmt.Errorf("missing dydx account id")
                }
        } else {
                vaultId = util.ValidNumber("0")
        }

        _, ok = data.GetOk("stark_key")
        if ok {
                starkKeyRaw := data.Get("stark_key").(string)
                starkKey = math.MustParseBig256("0x" + starkKeyRaw)
                if starkKey == nil {
                        return nil, fmt.Errorf("missing stark key")
                }
        } else {
                starkKey = big.NewInt(0)
        }

        _, ok = data.GetOk("asset_type")
        if ok {
                assetType = util.ValidNumber(data.Get("asset_type").(string))
                if assetType == nil {
                        return nil, fmt.Errorf("missing asset type")
                }
        } else {
                assetType = util.ValidNumber("0")
        }

        chainID := util.ValidNumber(config.ChainID)
        if chainID == nil {
                return nil, fmt.Errorf("invalid chain ID")
        }

        client, err := ethclient.Dial(config.getRPCURL())
        if err != nil {
                return nil, err
        }

        instance, err := starkwarePerpetuals.NewStarkwarePerpetuals(common.HexToAddress(contractAddress), client)
        if err != nil {
                return nil, err
        }
        callOpts := &bind.CallOpts{}

//      starkCallerSession := &starkwarePerpetuals.StarkwarePerpetualsSession{
//              Contract: instance, // Generic contract caller binding to set the session for
//              CallOpts: *callOpts,             // Call options to use throughout this session
//      }

        transactionParams, err := b.getBaseData(client, account.Address, data, "contract")
        if err != nil {
                return nil, err
        }

        _, ok = data.GetOk("tokens")
        if ok {
                tokens = util.ValidNumber(data.Get("tokens").(string))
                if tokens == nil {
                        return nil, fmt.Errorf("number of tokens are required")
                }
        } else {
                tokens = util.ValidNumber("0")
        }

        err = config.ValidAddress(transactionParams.Address)
        if err != nil {
                return nil, err
        }
        err = accountJSON.ValidAddress(transactionParams.Address)
        if err != nil {
                return nil, err
        }
        //tokenAmount := util.TokenAmount(tokens.Int64(), decimals)
        transactOpts, err := b.NewWalletTransactor(chainID, wallet, account)
        if err != nil {
                return nil, err
        }

        transactOpts.GasPrice = transactionParams.GasPrice
        transactOpts.GasLimit = transactionParams.GasLimit
        transactOpts.GasTipCap = transactionParams.GasTipCap
        transactOpts.GasFeeCap = transactionParams.GasFeeCap

        b.Logger().Info(fmt.Sprintf("\nGAS LIMIT: %d\n", transactOpts.GasLimit))
        b.Logger().Info(fmt.Sprintf("\nGAS PRICE: %d\n", transactOpts.GasPrice))
        b.Logger().Info(fmt.Sprintf("\nGAS TIPCAP: %d\n", transactOpts.GasTipCap))
        b.Logger().Info(fmt.Sprintf("\nGAS FEECAP: %d\n", transactOpts.GasFeeCap))

        //transactOpts needs gas etc.
        contractSession := &starkwarePerpetuals.StarkwarePerpetualsSession{
                Contract:     instance,  // Generic contract caller binding to set the session for
                CallOpts:     *callOpts, // Call options to use throughout this session
                TransactOpts: *transactOpts,
        }

        tx, err := contractSession.Deposit0(starkKey, assetType, vaultId, tokens)
        if err != nil {
                return nil, err
        }

        var signedTxBuff bytes.Buffer
        tx.EncodeRLP(&signedTxBuff)
        return &logical.Response{
                Data: map[string]interface{}{
                        //"contract":           tokenAddress.Hex(),
                        //"symbol":             symbol,
                        //"name":               tokenName,
                        "transaction_hash":   tx.Hash().Hex(),
                        "signed_transaction": hexutil.Encode(signedTxBuff.Bytes()),
                        "from":               account.Address.Hex(),
                        "to":                 transactionParams.Address.String(),
                        "amount":             tokens.String(),
                        "nonce":              tx.Nonce(),
                        "gas_price":          tx.GasPrice(),
                        "gas_limit":          tx.Gas(),
                },
        }, nil

}
