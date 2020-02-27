package main

import (
	"sync"

	dmecommon "github.com/mobiledgex/edge-cloud/d-match-engine/dme-common"
	"github.com/mobiledgex/edge-cloud/edgeproto"
	"golang.org/x/net/context"
)

type ClientsMap struct {
	sync.RWMutex
	clients map[edgeproto.AppInstKey][]edgeproto.AppInstClient
}

var clientsMap *ClientsMap

func InitAppInstClients() {
	clientsMap = new(ClientsMap)
	clientsMap.clients = make(map[edgeproto.AppInstKey][]edgeproto.AppInstClient)
}

// Add a new client to the list of clients
func UpdateClientsBuffer(ctx context.Context, msg *edgeproto.AppInstClient) {
	clientsMap.Lock()
	defer clientsMap.Unlock()
	list, found := clientsMap.clients[msg.ClientKey.Key]
	if !found {
		clientsMap.clients[msg.ClientKey.Key] = []edgeproto.AppInstClient{*msg}
	} else {
		// We need to either update, or add the client to the list
		for ii, c := range clientsMap.clients[msg.ClientKey.Key] {
			// Found the same client from before
			if c.ClientKey.Uuid == msg.ClientKey.Uuid {
				if len(clientsMap.clients[msg.ClientKey.Key]) > ii+1 {
					// remove this client the and append it at the end, since it's new
					clientsMap.clients[msg.ClientKey.Key] =
						append(clientsMap.clients[msg.ClientKey.Key][:ii],
							clientsMap.clients[msg.ClientKey.Key][ii+1:]...)
				} else {
					// if this is already the last element
					clientsMap.clients[msg.ClientKey.Key] =
						clientsMap.clients[msg.ClientKey.Key][:ii]

				}
				break
			}
		}
		//  We reached the limit of clients - remove the first one
		if len(list) == int(dmecommon.Settings.MaxTrackedDmeClients) {
			list = list[1:]
		}
		clientsMap.clients[msg.ClientKey.Key] = append(list, *msg)
	}
	// If there is an outstanding request for this appInst, send it out
	if appInstClientKeyCache.HasKey(msg.ClientKey.GetKey()) {
		ClientSender.Update(ctx, msg)
	}
}

// If an AppInst is deleted, clean up all the clients from it
func PurgeAppInstClients(ctx context.Context, msg *edgeproto.AppInstKey) {
	clientsMap.Lock()
	defer clientsMap.Unlock()
	_, found := clientsMap.clients[*msg]
	if found {
		delete(clientsMap.clients, *msg)
	}
}

func SendCachedClients(ctx context.Context, old *edgeproto.AppInstClientKey, new *edgeproto.AppInstClientKey) {
	clientsMap.Lock()
	defer clientsMap.Unlock()
	list, found := clientsMap.clients[new.Key]
	if !found {
		return
	}
	for ii, _ := range list {
		ClientSender.Update(ctx, &list[ii])
	}
}

// TODO - function to periodically timeout the clients