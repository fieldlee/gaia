package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types"

var (
	ErrInvalidCollection = sdkerrors.ErrUnknownRequest( "invalid NFT collection")
	ErrUnknownCollection = sdkerrors.ErrUnknownRequest( "unknown NFT collection")
	ErrInvalidNFT        = sdkerrors.ErrUnknownRequest( "invalid NFT")
	ErrUnknownNFT        = sdkerrors.ErrUnknownRequest( "unknown NFT")
	ErrNFTAlreadyExists  = sdkerrors.ErrUnknownRequest( "NFT already exists")
	ErrEmptyMetadata     = sdkerrors.ErrUnknownRequest( "NFT metadata can't be empty")
)
