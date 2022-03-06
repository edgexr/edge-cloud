package crmutil

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
)

//OpenstackImageArgs lists args that are part of openstack image data
type OpenstackImageArgs struct {
	Name            string
	ID              string
	Visibility      string
	Tags            []string
	ContainerFormat string
	DiskFormat      string
	MinDisk         int
	MinRAM          int
	Protected       bool
	Properties      map[string]string
}

//CreateOpenstackImage creates data in the image database in openstack glance
func CreateOpenstackImage(client *gophercloud.ServiceClient,
	args *OpenstackImageArgs) error {
	res := images.Create(client, images.CreateOpts{
		Name:            args.Name,
		ContainerFormat: args.ContainerFormat,
		DiskFormat:      args.DiskFormat,
	})
	//TODO: Tags, Properties, ...

	if res.Err != nil {
		return res.Err
	}

	return nil
}