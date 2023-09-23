// Handle a message to create product
func handleMsgCreateProduct(ctx sdk.Context, keeper Keeper, msg MsgCreateProduct) (*sdk.Result, error) {
   key := "Product-" + msg.ProductID

   if keeper.IsProductPresent(ctx, key) { // nếu key đã tồn tai thì trả về lỗi
   	return nil, sdkerrors.Wrap(types.ErrProductAlreadyExists, msg.ProductID)
   }

   var product = Product{ // khởi tạo một product mới
   	ProductID:   msg.ProductID,
   	Description: msg.Description,
   	Price:       msg.Price,
   	Owner:       msg.Signer,
   }

   keeper.SetProduct(ctx, key, product)
   return &sdk.Result{}, nil
}


// Handle a message to update product
func handleMsgUpdateProduct(ctx sdk.Context, keeper Keeper, msg MsgUpdateProduct) (*sdk.Result, error) {
  key := "Product-" + msg.ProductID

  if !keeper.IsProductPresent(ctx, key) { // kiểm tra xem product muốn cập nhật có tồn tại không
  	return nil, sdkerrors.Wrap(types.ErrProductDoesNotExist, msg.ProductID)
  }

  product := keeper.GetProduct(ctx, key)

  if !msg.Signer.Equals(product.Owner) { // kiểm tả xem người cập nhật product có phải là chủ hiện tại không
  	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // nếu không thì trả về lỗi
  }

  product.Description = msg.Description // Cập nhật các trường
  product.Price = msg.Price

  keeper.SetProduct(ctx, key, product) //
  return &sdk.Result{}, nil            //
}

// Handle a message to delete product
func handleMsgDeleteProduct(ctx sdk.Context, keeper Keeper, msg MsgDeleteProduct) (*sdk.Result, error) {
   key := "Product-" + msg.ProductID

   if !keeper.IsProductPresent(ctx, key) {
   	return nil, sdkerrors.Wrap(types.ErrNameDoesNotExist, msg.ProductID)
   }
   if !msg.Signer.Equals(keeper.GetProduct(ctx, key).Owner) { // kiểm tra xem người muốn xóa product có phải là chủ hiện tại của product không
   	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
   }

   keeper.DeleteProduct(ctx, key)
   return &sdk.Result{}, nil
}

// Handle a message to buy product
func handleMsgBuyProduct(ctx sdk.Context, keeper Keeper, msg MsgBuyProduct) (*sdk.Result, error) {
   key := "Product-" + msg.ProductID

   if !keeper.IsProductPresent(ctx, key) {
   	return nil, sdkerrors.Wrap(types.ErrNameDoesNotExist, msg.ProductID)
   }

   product := keeper.GetProduct(ctx, key)

   if msg.Signer.Equals(product.Owner) {
   	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You are product owner")
   }

   err := keeper.CoinKeeper.SendCoins(ctx, msg.Signer, product.Owner, product.Price) // chuyển một lượng token bằng đúng với price của product từ người mua sang cho người bán
   if err != nil {
   	return nil, err
   }

   product.Owner = msg.Signer // chuyển quyền sở hữu cho người mua

   keeper.SetProduct(ctx, key, product) // set lại product
   return &sdk.Result{}, nil
}

// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case MsgSetName:
			return handleMsgSetName(ctx, keeper, msg)
		case MsgBuyName:
			return handleMsgBuyName(ctx, keeper, msg)
		case MsgDeleteName:
			return handleMsgDeleteName(ctx, keeper, msg)
		case MsgCreateProduct:
			return handleMsgCreateProduct(ctx, keeper, msg)
		case MsgUpdateProduct:
			return handleMsgUpdateProduct(ctx, keeper, msg)
		case MsgDeleteProduct:
			return handleMsgDeleteProduct(ctx, keeper, msg)
		case MsgBuyProduct:
			return handleMsgBuyProduct(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type()))
		}
	}
}
