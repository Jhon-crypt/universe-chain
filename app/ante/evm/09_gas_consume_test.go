// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)
package evm_test

import (
	"fmt"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	evmante "github.com/evmos/evmos/v16/app/ante/evm"
	commonfactory "github.com/evmos/evmos/v16/testutil/integration/common/factory"
	evmosfactory "github.com/evmos/evmos/v16/testutil/integration/evmos/factory"
	"github.com/evmos/evmos/v16/testutil/integration/evmos/grpc"
	testkeyring "github.com/evmos/evmos/v16/testutil/integration/evmos/keyring"
	"github.com/evmos/evmos/v16/testutil/integration/evmos/network"
)

func (suite *EvmAnteTestSuite) TestUpdateComulativeGasWanted() {
	keyring := testkeyring.New(1)
	unitNetwork := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)

	testCases := []struct {
		name                string
		msgGasWanted        uint64
		maxTxGasWanted      uint64
		cumulativeGasWanted uint64
		getCtx              func() sdktypes.Context
		expectedResponse    uint64
	}{
		{
			name:                "when is NOT checkTx and cumulativeGasWanted is 0, returns msgGasWanted",
			msgGasWanted:        100,
			maxTxGasWanted:      150,
			cumulativeGasWanted: 0,
			getCtx: func() sdktypes.Context {
				return unitNetwork.GetContext().WithIsCheckTx(false)
			},
			expectedResponse: 100,
		},
		{
			name:                "when is NOT checkTx and cumulativeGasWanted has value, returns cumulativeGasWanted + msgGasWanted",
			msgGasWanted:        100,
			maxTxGasWanted:      150,
			cumulativeGasWanted: 50,
			getCtx: func() sdktypes.Context {
				return unitNetwork.GetContext().WithIsCheckTx(false)
			},
			expectedResponse: 150,
		},
		{
			name:                "when is checkTx, maxTxGasWanted is not 0 and msgGasWanted > maxTxGasWanted, returns cumulativeGasWanted + maxTxGasWanted",
			msgGasWanted:        200,
			maxTxGasWanted:      100,
			cumulativeGasWanted: 50,
			getCtx: func() sdktypes.Context {
				return unitNetwork.GetContext().WithIsCheckTx(true)
			},
			expectedResponse: 150,
		},
		{
			name:                "when is checkTx, maxTxGasWanted is not 0 and msgGasWanted < maxTxGasWanted, returns cumulativeGasWanted + msgGasWanted",
			msgGasWanted:        50,
			maxTxGasWanted:      100,
			cumulativeGasWanted: 50,
			getCtx: func() sdktypes.Context {
				return unitNetwork.GetContext().WithIsCheckTx(true)
			},
			expectedResponse: 100,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			// Function under test
			gasWanted := evmante.UpdateComulativeGasWanted(
				tc.getCtx(),
				tc.msgGasWanted,
				tc.maxTxGasWanted,
				tc.cumulativeGasWanted,
			)

			suite.Require().Equal(tc.expectedResponse, gasWanted)
		})
	}
}

func (suite *EvmAnteTestSuite) TestConsumeGasAndEmitEvent() {
	keyring := testkeyring.New(1)
	unitNetwork := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)
	grpcHandler := grpc.NewIntegrationHandler(unitNetwork)
	txFactory := evmosfactory.New(unitNetwork, grpcHandler)

	testCases := []struct {
		name          string
		expectedError error
		fees          sdktypes.Coins
		getSender     func() sdktypes.AccAddress
	}{
		// {
		// 	name:          "success: fees are zero and event emitted",
		// 	expectedError: nil,
		// 	fees:          sdktypes.Coins{},
		// 	getSender: func() sdktypes.AccAddress {
		// 		// Return prefunded sender
		// 		return keyring.GetKey(0).AccAddr
		// 	},
		// },
		// {
		// 	name:          "success: there are non zero fees, user has sufficient bank balances and event emitted",
		// 	expectedError: nil,
		// 	fees: sdktypes.Coins{
		// 		sdktypes.NewCoin(unitNetwork.GetDenom(), sdktypes.NewInt(1000)),
		// 	},
		// 	getSender: func() sdktypes.AccAddress {
		// 		// Return prefunded sender
		// 		return keyring.GetKey(0).AccAddr
		// 	},
		// },
		// {
		// 	name:          "fail: insufficient user balance",
		// 	expectedError: sdkerrors.ErrInsufficientFee,
		// 	fees: sdktypes.Coins{
		// 		sdktypes.NewCoin(unitNetwork.GetDenom(), sdktypes.NewInt(1000)),
		// 	},
		// 	getSender: func() sdktypes.AccAddress {
		// 		// Return unfunded account
		// 		index := keyring.AddKey()
		// 		return keyring.GetKey(index).AccAddr
		// 	},
		// },
		{
			name:          "success: fees are non zero, bank balance is insufficient but user has sufficient staking balance and event is emitted",
			expectedError: sdkerrors.ErrInsufficientFee,
			fees: sdktypes.Coins{
				sdktypes.NewCoin(unitNetwork.GetDenom(), sdktypes.NewInt(1000)),
			},
			getSender: func() sdktypes.AccAddress {
				// Return unfunded account
				index := keyring.AddKey()
				senderKey := keyring.GetKey(index)

				// bank balance can't be 0
				bankBalance := sdktypes.NewCoin(unitNetwork.GetDenom(), sdktypes.NewInt(10))
				stakingBalance := sdktypes.NewCoin(unitNetwork.GetDenom(), sdktypes.NewInt(1000000000000000000))

				err := unitNetwork.FundAccount(senderKey.AccAddr, sdktypes.NewCoins(bankBalance))
				suite.Require().NoError(err)

				validators := unitNetwork.GetValidators()
				bechVal, err := sdktypes.ValAddressFromBech32(validators[0].OperatorAddress)
				suite.Require().NoError(err)

				delegateMsg := stakingtypes.NewMsgDelegate(
					senderKey.AccAddr,
					bechVal,
					stakingBalance,
				)

				txArgs := commonfactory.CosmosTxArgs{Msgs: []sdktypes.Msg{delegateMsg}}
				resp, err := txFactory.ExecuteCosmosTx(senderKey.Priv, txArgs)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(0), resp.Code)

				return keyring.GetKey(index).AccAddr
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			keepers := &evmante.ConsumeGasKeepers{
				Bank:         unitNetwork.App.BankKeeper,
				Distribution: unitNetwork.App.DistrKeeper,
				Evm:          unitNetwork.App.EvmKeeper,
				Staking:      unitNetwork.App.StakingKeeper,
			}
			sender := tc.getSender()
			prevBalance, err := grpcHandler.GetAllBalances(
				sender,
			)
			suite.Require().NoError(err)
			fmt.Println("prevBalance", prevBalance)

			// Function under test
			err = evmante.ConsumeGasAndEmitEvent(
				unitNetwork.GetContext(),
				keepers,
				tc.fees,
				sender,
			)

			if tc.expectedError != nil {
				suite.Require().Error(err)
				suite.Contains(err.Error(), tc.expectedError.Error())

				// Check events are not present
				events := unitNetwork.GetContext().EventManager().Events()
				suite.Require().Zero(len(events))
			} else {
				suite.Require().NoError(err)

				// Check fees are deducted
				afterBalance, err := grpcHandler.GetAllBalances(
					sender,
				)
				suite.Require().NoError(err)
				expectedBalance := prevBalance.Balances.Sub(tc.fees...)
				suite.Require().True(
					expectedBalance.IsEqual(afterBalance.Balances),
				)

				// Event to be emitted
				expectedEvent := sdktypes.NewEvent(
					sdktypes.EventTypeTx,
					sdktypes.NewAttribute(sdktypes.AttributeKeyFee, tc.fees.String()),
				)
				// Check events are present
				events := unitNetwork.GetContext().EventManager().Events()
				suite.Require().NotZero(len(events))
				suite.Require().Contains(
					events,
					expectedEvent,
				)
			}

			// Reset the context
			err = unitNetwork.NextBlock()
			suite.Require().NoError(err)
		})
	}
}
