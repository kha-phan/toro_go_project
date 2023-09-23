// MsgCreateProduct defines a CreateProduct message
type MsgCreateProduct struct {
	ProductID   string         `json:"productID"`
	Description string         `json:"description"`
	Price       sdk.Coins      `json:"price"`
	Signer      sdk.AccAddress `json:"signer"`
}

// NewMsgCreateProduct is a constructor function for MsgCreateProduct
func NewMsgCreateProduct(productID string, description string, price sdk.Coins, signer sdk.AccAddress) MsgCreateProduct {
	return MsgCreateProduct{
		ProductID:   productID,
		Description: description,
		Price:       price,
		Signer:      signer,
	}
}

// Route should return the name of the module
func (msg MsgCreateProduct) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateProduct) Type() string { return "create_product" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateProduct) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ProductID) == 0 || len(msg.Description) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID and/or Description cannot be empty")
	}
	if !msg.Price.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateProduct) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateProduct) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgUpdateProduct defines a UpdateProduct message
type MsgUpdateProduct struct {
	ProductID   string         `json:"productID"`
	Description string         `json:"description"`
	Price       sdk.Coins      `json:"price"`
	Signer      sdk.AccAddress `json:"signer"`
}

// NewMsgUpdateProduct is a constructor function for MsgUpdateProduct
func NewMsgUpdateProduct(productID string, description string, price sdk.Coins, signer sdk.AccAddress) MsgUpdateProduct {
	return MsgUpdateProduct{
		ProductID:   productID,
		Description: description,
		Price:       price,
		Signer:      signer,
	}
}

// Route should return the name of the module
func (msg MsgUpdateProduct) Route() string { return RouterKey }

// Type should return the action
func (msg MsgUpdateProduct) Type() string { return "update_product" }

// ValidateBasic runs stateless checks on the message
func (msg MsgUpdateProduct) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ProductID) == 0 || len(msg.Description) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID and/or Description cannot be empty")
	}
	if !msg.Price.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUpdateProduct) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUpdateProduct) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgDeleteProduct defines a DeleteProduct message
type MsgDeleteProduct struct {
	ProductID   string         `json:"productID"`
	Signer      sdk.AccAddress `json:"signer"`
}

// NewMsgDeleteProduct is a constructor function for MsgUpdateProduct
func NewMsgDeleteProduct(productID string, signer sdk.AccAddress) MsgDeleteProduct {
	return MsgDeleteProduct{
		ProductID:   productID,
		Signer:      signer,
	}
}

// Route should return the name of the module
func (msg MsgUpdateProduct) Route() string { return RouterKey }

// Type should return the action
func (msg MsgUpdateProduct) Type() string { return "update_product" }

// ValidateBasic runs stateless checks on the message
func (msg MsgUpdateProduct) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ProductID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID and/or Description cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteProduct) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteProduct) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgBuyProduct defines a DeleteName message
type MsgBuyProduct struct {
	ProductID string         `json:"productID"`
	Signer    sdk.AccAddress `json:"signer"`
}

// NewMsgBuyProduct is a constructor function for MsgBuyProduct
func NewMsgBuyProduct(productID string, signer sdk.AccAddress) MsgBuyProduct {
	return MsgBuyProduct{
		ProductID: productID,
		Signer:    signer,
	}
}

// Route should return the name of the module
func (msg MsgBuyProduct) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBuyProduct) Type() string { return "buy_product" }

// ValidateBasic runs stateless checks on the message
func (msg MsgBuyProduct) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ProductID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBuyProduct) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBuyProduct) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}
