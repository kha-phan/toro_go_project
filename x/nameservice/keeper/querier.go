const (
	QueryResolve = "resolve"
	QueryWhois   = "whois"
	QueryNames   = "names"

	QueryProduct     = "product"
	QueryAllProducts = "allProducts"
)


func queryProduct(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) { // query một sản phẩm theo productID

	key := "Product-" + path[0] // Tại sao mình lại thêm tiền tố "Product-" ở đây, đọc xuống dưới các bạn sẽ hiểu

	product := keeper.GetProduct(ctx, key)

	res, err := codec.MarshalJSONIndent(keeper.cdc, product)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryAllProducts(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) { // lấy tất cả products

	var productsList types.QueryResProducts

	iterator := keeper.GetProductsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		if "Product-" <= key && key <= "Product-zzzzzzzz" {
			product := keeper.GetProduct(ctx, key)
			productsList = append(productsList, product)

		}
	}
	res, err := codec.MarshalJSONIndent(keeper.cdc, productsList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryResolve:
			return queryResolve(ctx, path[1:], req, keeper)
		case QueryWhois:
			return queryWhois(ctx, path[1:], req, keeper)
		case QueryNames:
			return queryNames(ctx, req, keeper)
		case QueryProduct:
			return queryProduct(ctx, path[1:], req, keeper)
		case QueryAllProducts:
			return queryAllProducts(ctx, req, keeper)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown nameservice query endpoint")
		}
	}
}
