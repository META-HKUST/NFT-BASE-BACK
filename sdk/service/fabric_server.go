package service

import (
	"NFT-BASE-BACK/sdk/pb"
	"context"
)

// Copy code from fabricsdk_service_grpc.pb.go
// Code to populate protobuf

type FabricServer struct {
	pb.UnimplementedFabricSDKServer
}

func NewFabricServer() *FabricServer {
	return &FabricServer{}
}

func (server *FabricServer) BalanceOf(ctx context.Context, req *pb.BalanceOfResuest) (*pb.BalanceOfResponse, error) {
	result, err := Evaluate(req.Username, "BalanceOf", req.Owner)
	if err != nil {
		return nil, err
	}
	return &pb.BalanceOfResponse{Num: result}, err
	// return nil, status.Errorf(codes.Unimplemented, "method BalanceOf not implemented")
}

func (server *FabricServer) OwnerOf(ctx context.Context, req *pb.OwnerOfRequest) (*pb.OwnerOfResponse, error) {
	result, err := Evaluate(req.Username, "OwnerOf", req.TokenId)
	if err != nil {
		return nil, err
	}
	return &pb.OwnerOfResponse{Owner: result}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method BalanceOf not implemented")
}

func (server *FabricServer) Approve(ctx context.Context, req *pb.ApproveRequest) (*pb.ApproveResponse, error) {
	result, err := Submit(req.Username, "Approve", req.Operator, req.TokenId)
	if err != nil {
		return nil, err
	}
	return &pb.ApproveResponse{Ok: result}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method Approve not implemented")
}

func (server *FabricServer) SetApprovalForAll(ctx context.Context, req *pb.SetApprovalForAllRequest) (*pb.SetApprovalForAllResponse, error) {
	result, err := Submit(req.Username, "SetApprovalForAll", req.Operator, req.Operator2)
	if err != nil {
		return nil, err
	}
	return &pb.SetApprovalForAllResponse{Ok: result}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method SetApprovalForAll not implemented")
}
func (server *FabricServer) IsApprovedForAll(ctx context.Context, req *pb.IsApprovedForAllRequest) (*pb.IsApprovedForAllResponse, error) {
	result, err := Submit(req.Username, "IsApprovedForAll", req.Owner, req.Operator)
	if err != nil {
		return nil, err
	}
	return &pb.IsApprovedForAllResponse{Ok: result}, nil

	// return nil, status.Errorf(codes.Unimplemented, "method IsApprovedForAll not implemented")
}
func (server *FabricServer) GetApproved(ctx context.Context, req *pb.GetApprovedRequest) (*pb.GetApprovedResponse, error) {
	_, err := Submit(req.Username, "GetApproved", req.TokenId)
	if err != nil {
		return nil, err
	}
	return &pb.GetApprovedResponse{}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method GetApproved not implemented")
}
func (server *FabricServer) TransferFrom(ctx context.Context, req *pb.TransferFromRequest) (*pb.TransferFromResponse, error) {
	result, err := Submit(req.Username, "TransferFrom", req.From, req.To, req.TokenId)
	if err != nil {
		return nil, err
	}
	return &pb.TransferFromResponse{Ok: result}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method TransferFrom not implemented")
}
func (server *FabricServer) ClientAccountBalance(ctx context.Context, req *pb.ClientAccountBalanceRequest) (*pb.ClientAccountBalanceResponse, error) {
	result, err := Evaluate(req.Username, "ClientAccountBalance")
	if err != nil {
		return nil, err
	}
	return &pb.ClientAccountBalanceResponse{Num: result}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method ClientAccountBalance not implemented")
}
func (server *FabricServer) ClientAccountID(ctx context.Context, req *pb.ClientAccountIDRequest) (*pb.ClientAccountIDResponse, error) {
	result, err := Evaluate(req.Username, "ClientAccountID")
	if err != nil {
		return nil, err
	}
	return &pb.ClientAccountIDResponse{X509: result}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method ClientAccountID not implemented")
}
func (server *FabricServer) TokenURI(ctx context.Context, req *pb.TokenURIRequest) (*pb.TokenURIReqponse, error) {
	result, err := Evaluate(req.Username, "TokenURI", req.TokenId)
	if err != nil {
		return nil, err
	}
	return &pb.TokenURIReqponse{URI: result}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method TokenURI not implemented")
}
func (server *FabricServer) SetChainCodeOwner(ctx context.Context, req *pb.SetChainCodeOwnerRequest) (*pb.SetChainCodeOwnerResponse, error) {
	result, err := Submit(req.Username, "SetChainCodeOwner", req.CodeOwner)
	if err != nil {
		return nil, err
	}
	return &pb.SetChainCodeOwnerResponse{Owner: result}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method SetChainCodeOwner not implemented")
}
func (server *FabricServer) GetChainCodeOwner(ctx context.Context, req *pb.GetChainCodeOwnerRequest) (*pb.GetChainCodeOwnerResponse, error) {
	result, err := Submit(req.Username, "GetChainCodeOwner")
	if err != nil {
		return nil, err
	}
	return &pb.GetChainCodeOwnerResponse{Owner: result}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method GetChainCodeOwner not implemented")
}
func (server *FabricServer) SetBaseURI(ctx context.Context, req *pb.SetBaseURIRequest) (*pb.SetBaseURIResponse, error) {
	result, err := Submit(req.Username, "SetBaseURI")
	if err != nil {
		return nil, err
	}
	return &pb.SetBaseURIResponse{Result: result}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method SetBaseURI not implemented")
}
func (server *FabricServer) PublicMint(ctx context.Context, req *pb.PublicMintRequest) (*pb.PublicMintResponse, error) {
	result, err := Submit(req.Username, "PublicMint")
	if err != nil {
		return nil, err
	}
	return &pb.PublicMintResponse{TokenId: result}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method PublicMint not implemented")
}
func (server *FabricServer) PrivateBurn(ctx context.Context, req *pb.PrivateBurnRequest) (*pb.PrivateBurnResponse, error) {
	result, err := Submit(req.Username, "PrivateBurn", req.TokenId)
	if err != nil {
		return nil, err
	}
	return &pb.PrivateBurnResponse{Ok: result}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method PrivateBurn not implemented")
}
func (server *FabricServer) TotalSupply(ctx context.Context, req *pb.TokenURIRequest) (*pb.TotalSupplyResponse, error) {
	result, err := Submit(req.Username, "TotalSupply", req.TokenId)
	if err != nil {
		return nil, err
	}
	return &pb.TotalSupplyResponse{Num: result}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method TotalSupply not implemented")
}
