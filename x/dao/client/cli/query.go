package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/onomyprotocol/onomy/x/dao/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd(_ string) *cobra.Command {
	// Group dao queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2, // nolint:gomnd
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdShowParams())
	cmd.AddCommand(CmdShowTreasury())
	cmd.AddCommand(CmdShowAccountBalances())

	return cmd
}

// CmdShowParams returns CmdShowParams cobra.Command.
func CmdShowParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-params",
		Short: "shows the parameters of the module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdShowTreasury returns CmdShowTreasury cobra.Command.
func CmdShowTreasury() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-treasury",
		Short: "Shows treasury balance.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Treasury(context.Background(), &types.QueryTreasuryRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdShowTreasury returns CmdShowTreasury cobra.Command.
func CmdShowAccountBalances() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account-balances [denom]",
		Short: "Shows all account balances.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			denom := args[0]

			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AccountBalances(context.Background(), &types.QueryAccountBalancesRequest{
				Denom: denom,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
