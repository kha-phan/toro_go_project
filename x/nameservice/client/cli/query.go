func GetCmdProduct(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "product [productID]",
		Short: "Query whois info of name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			productID := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/product/%s", queryRoute, productID), nil)
			if err != nil {
				fmt.Printf("could not query product - %s \n", productID)
				return nil
			}

			var out types.Product
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdAllProducts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "all-products",
		Short: "all-products",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/allProducts", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get all products \n")
				return nil
			}

			var out types.QueryResAllProducts
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

nameserviceQueryCmd.AddCommand(flags.GetCommands(
		GetCmdResolveName(storeKey, cdc),
		GetCmdWhois(storeKey, cdc),
		GetCmdNames(storeKey, cdc),

		GetCmdProduct(storeKey, cdc),
		GetCmdAllProducts(storeKey, cdc),
	)...)
