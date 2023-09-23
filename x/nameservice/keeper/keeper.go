
func (k Keeper) GetProduct(ctx sdk.Context, key string) types.Product {
	store := ctx.KVStore(k.storeKey)

	if !k.IsProductPresent(ctx, key) {
		return types.NewProduct()
	}

	bz := store.Get([]byte(key))  // lấy value theo key

	var product types.Product

	k.cdc.MustUnmarshalBinaryBare(bz, &product) // ép kiểu từ []bytes về  struct Product

	return product
}

func (k Keeper) SetProduct(ctx sdk.Context, key string, product types.Product) {
	if product.Owner.Empty() {
		return
	}

	store := ctx.KVStore(k.storeKey)

	store.Set([]byte(key), k.cdc.MustMarshalBinaryBare(product)) // set vào store theo key-value
}

func (k Keeper) DeleteProduct(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(key)) # delete value theo key
}

func (k Keeper) IsProductPresent(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(key)) # check xem key có tồn tại hay không
}

func (k Keeper) GetProductsIterator(ctx sdk.Context) sdk.Iterator { // lấy mảng key
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
