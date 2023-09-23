func GetCmdCreateProduct(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "create-product [productID] [description] [price]",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			coins, err := sdk.ParseCoins(args[2]) // ép về kiểu sdk.Coins trong cosmos sdk
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateProduct(args[0], args[1], coins, cliCtx.GetFromAddress()) // cliCtx.GetFromAddress() hàm này dùng để lấy address của user đã thực hiện lệnh create-product
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdUpdateProduct(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "update-product [productID] [description] [price]",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			coins, err := sdk.ParseCoins(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateProduct(args[0], args[1], coins, cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteProduct(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "delete-product [productID]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteProduct(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdBuyProduct(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "buy-product [productID]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgBuyProduct(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

nameserviceTxCmd.AddCommand(flags.PostCommands(
		GetCmdBuyName(cdc),
		GetCmdSetName(cdc),
		GetCmdDeleteName(cdc),

		GetCmdCreateProduct(cdc),
		GetCmdUpdateProduct(cdc),
		GetCmdDeleteProduct(cdc),
		GetCmdBuyProduct(cdc),
	)...)
