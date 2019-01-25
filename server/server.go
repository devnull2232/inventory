package server;

import (
	pb "github.com/devnull2232/inventory/catalog/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	log "log"
)

type catalogServer struct{}

// Gets all items
func (c *catalogServer) GetAllItems(e *empty.Empty, stream pb.Catalog_GetAllItemsServer) error{
	log.Println("Starting stream...")
	go func(){
		for {
			// TODO: get msgs from DB (msg := something)
			// msg is GetAllItemsResponse
			stream.Send(msg)
		}
	}()
}