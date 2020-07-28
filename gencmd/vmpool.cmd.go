// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vmpool.proto

package gencmd

import edgeproto "github.com/mobiledgex/edge-cloud/edgeproto"
import google_protobuf "github.com/gogo/protobuf/types"
import "strings"
import "github.com/spf13/cobra"
import "context"
import "io"
import "github.com/mobiledgex/edge-cloud/cli"
import "google.golang.org/grpc/status"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/mobiledgex/edge-cloud/protogen"
import _ "github.com/mobiledgex/edge-cloud/d-match-engine/dme-proto"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT
func VMHideTags(in *edgeproto.VM) {
	if cli.HideTags == "" {
		return
	}
	tags := make(map[string]struct{})
	for _, tag := range strings.Split(cli.HideTags, ",") {
		tags[tag] = struct{}{}
	}
	if _, found := tags["timestamp"]; found {
		in.UpdatedAt = google_protobuf.Timestamp{}
	}
}

func VMPoolHideTags(in *edgeproto.VMPool) {
	if cli.HideTags == "" {
		return
	}
	tags := make(map[string]struct{})
	for _, tag := range strings.Split(cli.HideTags, ",") {
		tags[tag] = struct{}{}
	}
	for i0 := 0; i0 < len(in.Vms); i0++ {
		if _, found := tags["timestamp"]; found {
			in.Vms[i0].UpdatedAt = google_protobuf.Timestamp{}
		}
	}
	if _, found := tags["nocmp"]; found {
		in.State = 0
	}
	if _, found := tags["nocmp"]; found {
		in.Errors = nil
	}
	if _, found := tags["nocmp"]; found {
		in.CrmOverride = 0
	}
}

func VMPoolMemberHideTags(in *edgeproto.VMPoolMember) {
	if cli.HideTags == "" {
		return
	}
	tags := make(map[string]struct{})
	for _, tag := range strings.Split(cli.HideTags, ",") {
		tags[tag] = struct{}{}
	}
	if _, found := tags["timestamp"]; found {
		in.Vm.UpdatedAt = google_protobuf.Timestamp{}
	}
	if _, found := tags["nocmp"]; found {
		in.CrmOverride = 0
	}
}

func VMPoolInfoHideTags(in *edgeproto.VMPoolInfo) {
	if cli.HideTags == "" {
		return
	}
	tags := make(map[string]struct{})
	for _, tag := range strings.Split(cli.HideTags, ",") {
		tags[tag] = struct{}{}
	}
	if _, found := tags["nocmp"]; found {
		in.NotifyId = 0
	}
	for i0 := 0; i0 < len(in.Vms); i0++ {
		if _, found := tags["timestamp"]; found {
			in.Vms[i0].UpdatedAt = google_protobuf.Timestamp{}
		}
	}
	if _, found := tags["nocmp"]; found {
		in.State = 0
	}
	if _, found := tags["nocmp"]; found {
		in.Errors = nil
	}
}

var VMPoolApiCmd edgeproto.VMPoolApiClient

var CreateVMPoolCmd = &cli.Command{
	Use:          "CreateVMPool",
	RequiredArgs: strings.Join(VMPoolRequiredArgs, " "),
	OptionalArgs: strings.Join(VMPoolOptionalArgs, " "),
	AliasArgs:    strings.Join(VMPoolAliasArgs, " "),
	SpecialArgs:  &VMPoolSpecialArgs,
	Comments:     VMPoolComments,
	ReqData:      &edgeproto.VMPool{},
	ReplyData:    &edgeproto.Result{},
	Run:          runCreateVMPool,
}

func runCreateVMPool(c *cli.Command, args []string) error {
	if cli.SilenceUsage {
		c.CobraCmd.SilenceUsage = true
	}
	obj := c.ReqData.(*edgeproto.VMPool)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return CreateVMPool(c, obj)
}

func CreateVMPool(c *cli.Command, in *edgeproto.VMPool) error {
	if VMPoolApiCmd == nil {
		return fmt.Errorf("VMPoolApi client not initialized")
	}
	ctx := context.Background()
	obj, err := VMPoolApiCmd.CreateVMPool(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("CreateVMPool failed: %s", errstr)
	}
	c.WriteOutput(obj, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func CreateVMPools(c *cli.Command, data []edgeproto.VMPool, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("CreateVMPool %v\n", data[ii])
		myerr := CreateVMPool(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var DeleteVMPoolCmd = &cli.Command{
	Use:          "DeleteVMPool",
	RequiredArgs: strings.Join(VMPoolRequiredArgs, " "),
	OptionalArgs: strings.Join(VMPoolOptionalArgs, " "),
	AliasArgs:    strings.Join(VMPoolAliasArgs, " "),
	SpecialArgs:  &VMPoolSpecialArgs,
	Comments:     VMPoolComments,
	ReqData:      &edgeproto.VMPool{},
	ReplyData:    &edgeproto.Result{},
	Run:          runDeleteVMPool,
}

func runDeleteVMPool(c *cli.Command, args []string) error {
	if cli.SilenceUsage {
		c.CobraCmd.SilenceUsage = true
	}
	obj := c.ReqData.(*edgeproto.VMPool)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return DeleteVMPool(c, obj)
}

func DeleteVMPool(c *cli.Command, in *edgeproto.VMPool) error {
	if VMPoolApiCmd == nil {
		return fmt.Errorf("VMPoolApi client not initialized")
	}
	ctx := context.Background()
	obj, err := VMPoolApiCmd.DeleteVMPool(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("DeleteVMPool failed: %s", errstr)
	}
	c.WriteOutput(obj, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func DeleteVMPools(c *cli.Command, data []edgeproto.VMPool, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("DeleteVMPool %v\n", data[ii])
		myerr := DeleteVMPool(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var UpdateVMPoolCmd = &cli.Command{
	Use:          "UpdateVMPool",
	RequiredArgs: strings.Join(VMPoolRequiredArgs, " "),
	OptionalArgs: strings.Join(VMPoolOptionalArgs, " "),
	AliasArgs:    strings.Join(VMPoolAliasArgs, " "),
	SpecialArgs:  &VMPoolSpecialArgs,
	Comments:     VMPoolComments,
	ReqData:      &edgeproto.VMPool{},
	ReplyData:    &edgeproto.Result{},
	Run:          runUpdateVMPool,
}

func runUpdateVMPool(c *cli.Command, args []string) error {
	if cli.SilenceUsage {
		c.CobraCmd.SilenceUsage = true
	}
	obj := c.ReqData.(*edgeproto.VMPool)
	jsonMap, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	obj.Fields = cli.GetSpecifiedFields(jsonMap, c.ReqData, cli.JsonNamespace)
	return UpdateVMPool(c, obj)
}

func UpdateVMPool(c *cli.Command, in *edgeproto.VMPool) error {
	if VMPoolApiCmd == nil {
		return fmt.Errorf("VMPoolApi client not initialized")
	}
	ctx := context.Background()
	obj, err := VMPoolApiCmd.UpdateVMPool(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("UpdateVMPool failed: %s", errstr)
	}
	c.WriteOutput(obj, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func UpdateVMPools(c *cli.Command, data []edgeproto.VMPool, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("UpdateVMPool %v\n", data[ii])
		myerr := UpdateVMPool(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var ShowVMPoolCmd = &cli.Command{
	Use:          "ShowVMPool",
	OptionalArgs: strings.Join(append(VMPoolRequiredArgs, VMPoolOptionalArgs...), " "),
	AliasArgs:    strings.Join(VMPoolAliasArgs, " "),
	SpecialArgs:  &VMPoolSpecialArgs,
	Comments:     VMPoolComments,
	ReqData:      &edgeproto.VMPool{},
	ReplyData:    &edgeproto.VMPool{},
	Run:          runShowVMPool,
}

func runShowVMPool(c *cli.Command, args []string) error {
	if cli.SilenceUsage {
		c.CobraCmd.SilenceUsage = true
	}
	obj := c.ReqData.(*edgeproto.VMPool)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return ShowVMPool(c, obj)
}

func ShowVMPool(c *cli.Command, in *edgeproto.VMPool) error {
	if VMPoolApiCmd == nil {
		return fmt.Errorf("VMPoolApi client not initialized")
	}
	ctx := context.Background()
	stream, err := VMPoolApiCmd.ShowVMPool(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("ShowVMPool failed: %s", errstr)
	}

	objs := make([]*edgeproto.VMPool, 0)
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			errstr := err.Error()
			st, ok := status.FromError(err)
			if ok {
				errstr = st.Message()
			}
			return fmt.Errorf("ShowVMPool recv failed: %s", errstr)
		}
		VMPoolHideTags(obj)
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		return nil
	}
	c.WriteOutput(objs, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func ShowVMPools(c *cli.Command, data []edgeproto.VMPool, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("ShowVMPool %v\n", data[ii])
		myerr := ShowVMPool(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var AddVMPoolMemberCmd = &cli.Command{
	Use:          "AddVMPoolMember",
	RequiredArgs: strings.Join(AddVMPoolMemberRequiredArgs, " "),
	OptionalArgs: strings.Join(AddVMPoolMemberOptionalArgs, " "),
	AliasArgs:    strings.Join(VMPoolMemberAliasArgs, " "),
	SpecialArgs:  &VMPoolMemberSpecialArgs,
	Comments:     VMPoolMemberComments,
	ReqData:      &edgeproto.VMPoolMember{},
	ReplyData:    &edgeproto.Result{},
	Run:          runAddVMPoolMember,
}

func runAddVMPoolMember(c *cli.Command, args []string) error {
	if cli.SilenceUsage {
		c.CobraCmd.SilenceUsage = true
	}
	obj := c.ReqData.(*edgeproto.VMPoolMember)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return AddVMPoolMember(c, obj)
}

func AddVMPoolMember(c *cli.Command, in *edgeproto.VMPoolMember) error {
	if VMPoolApiCmd == nil {
		return fmt.Errorf("VMPoolApi client not initialized")
	}
	ctx := context.Background()
	obj, err := VMPoolApiCmd.AddVMPoolMember(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("AddVMPoolMember failed: %s", errstr)
	}
	c.WriteOutput(obj, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func AddVMPoolMembers(c *cli.Command, data []edgeproto.VMPoolMember, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("AddVMPoolMember %v\n", data[ii])
		myerr := AddVMPoolMember(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var RemoveVMPoolMemberCmd = &cli.Command{
	Use:          "RemoveVMPoolMember",
	RequiredArgs: strings.Join(RemoveVMPoolMemberRequiredArgs, " "),
	OptionalArgs: strings.Join(RemoveVMPoolMemberOptionalArgs, " "),
	AliasArgs:    strings.Join(VMPoolMemberAliasArgs, " "),
	SpecialArgs:  &VMPoolMemberSpecialArgs,
	Comments:     VMPoolMemberComments,
	ReqData:      &edgeproto.VMPoolMember{},
	ReplyData:    &edgeproto.Result{},
	Run:          runRemoveVMPoolMember,
}

func runRemoveVMPoolMember(c *cli.Command, args []string) error {
	if cli.SilenceUsage {
		c.CobraCmd.SilenceUsage = true
	}
	obj := c.ReqData.(*edgeproto.VMPoolMember)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return RemoveVMPoolMember(c, obj)
}

func RemoveVMPoolMember(c *cli.Command, in *edgeproto.VMPoolMember) error {
	if VMPoolApiCmd == nil {
		return fmt.Errorf("VMPoolApi client not initialized")
	}
	ctx := context.Background()
	obj, err := VMPoolApiCmd.RemoveVMPoolMember(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("RemoveVMPoolMember failed: %s", errstr)
	}
	c.WriteOutput(obj, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func RemoveVMPoolMembers(c *cli.Command, data []edgeproto.VMPoolMember, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("RemoveVMPoolMember %v\n", data[ii])
		myerr := RemoveVMPoolMember(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var VMPoolApiCmds = []*cobra.Command{
	CreateVMPoolCmd.GenCmd(),
	DeleteVMPoolCmd.GenCmd(),
	UpdateVMPoolCmd.GenCmd(),
	ShowVMPoolCmd.GenCmd(),
	AddVMPoolMemberCmd.GenCmd(),
	RemoveVMPoolMemberCmd.GenCmd(),
}

var VMNetInfoRequiredArgs = []string{}
var VMNetInfoOptionalArgs = []string{
	"externalip",
	"internalip",
}
var VMNetInfoAliasArgs = []string{}
var VMNetInfoComments = map[string]string{
	"externalip": "External IP",
	"internalip": "Internal IP",
}
var VMNetInfoSpecialArgs = map[string]string{}
var VMRequiredArgs = []string{}
var VMOptionalArgs = []string{
	"name",
	"netinfo.externalip",
	"netinfo.internalip",
	"groupname",
	"state",
	"updatedat.seconds",
	"updatedat.nanos",
	"internalname",
}
var VMAliasArgs = []string{}
var VMComments = map[string]string{
	"name":               "VM Name",
	"netinfo.externalip": "External IP",
	"netinfo.internalip": "Internal IP",
	"groupname":          "VM Group Name",
	"state":              "VM State, one of VmFree, VmInProgress, VmInUse, VmAdd, VmRemove, VmUpdate",
	"updatedat.seconds":  "Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.",
	"updatedat.nanos":    "Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive.",
	"internalname":       "VM Internal Name",
}
var VMSpecialArgs = map[string]string{}
var VMPoolKeyRequiredArgs = []string{}
var VMPoolKeyOptionalArgs = []string{
	"organization",
	"name",
}
var VMPoolKeyAliasArgs = []string{}
var VMPoolKeyComments = map[string]string{
	"organization": "Organization of the vmpool",
	"name":         "Name of the vmpool",
}
var VMPoolKeySpecialArgs = map[string]string{}
var VMPoolRequiredArgs = []string{
	"vmpool-org",
	"vmpool",
}
var VMPoolOptionalArgs = []string{
	"vms:#.name",
	"vms:#.netinfo.externalip",
	"vms:#.netinfo.internalip",
	"crmoverride",
}
var VMPoolAliasArgs = []string{
	"vmpool-org=key.organization",
	"vmpool=key.name",
}
var VMPoolComments = map[string]string{
	"fields":                   "Fields are used for the Update API to specify which fields to apply",
	"vmpool-org":               "Organization of the vmpool",
	"vmpool":                   "Name of the vmpool",
	"vms:#.name":               "VM Name",
	"vms:#.netinfo.externalip": "External IP",
	"vms:#.netinfo.internalip": "Internal IP",
	"vms:#.groupname":          "VM Group Name",
	"vms:#.state":              "VM State, one of VmFree, VmInProgress, VmInUse, VmAdd, VmRemove, VmUpdate",
	"vms:#.updatedat.seconds":  "Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.",
	"vms:#.updatedat.nanos":    "Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive.",
	"vms:#.internalname":       "VM Internal Name",
	"state":                    "Current state of the VM pool, one of TrackedStateUnknown, NotPresent, CreateRequested, Creating, CreateError, Ready, UpdateRequested, Updating, UpdateError, DeleteRequested, Deleting, DeleteError, DeletePrepare, CrmInitok, CreatingDependencies",
	"errors":                   "Any errors trying to add/remove VM to/from VM Pool",
	"crmoverride":              "Override actions to CRM, one of NoOverride, IgnoreCrmErrors, IgnoreCrm, IgnoreTransientState, IgnoreCrmAndTransientState",
}
var VMPoolSpecialArgs = map[string]string{
	"errors": "StringArray",
	"fields": "StringArray",
}
var VMPoolMemberRequiredArgs = []string{
	"vmpool-org",
	"vmpool",
}
var VMPoolMemberOptionalArgs = []string{
	"vm.name",
	"vm.netinfo.externalip",
	"vm.netinfo.internalip",
	"vm.updatedat.seconds",
	"vm.updatedat.nanos",
	"crmoverride",
}
var VMPoolMemberAliasArgs = []string{
	"vmpool-org=key.organization",
	"vmpool=key.name",
}
var VMPoolMemberComments = map[string]string{
	"vmpool-org":            "Organization of the vmpool",
	"vmpool":                "Name of the vmpool",
	"vm.name":               "VM Name",
	"vm.netinfo.externalip": "External IP",
	"vm.netinfo.internalip": "Internal IP",
	"vm.groupname":          "VM Group Name",
	"vm.state":              "VM State, one of VmFree, VmInProgress, VmInUse, VmAdd, VmRemove, VmUpdate",
	"vm.updatedat.seconds":  "Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.",
	"vm.updatedat.nanos":    "Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive.",
	"vm.internalname":       "VM Internal Name",
	"crmoverride":           "Override actions to CRM, one of NoOverride, IgnoreCrmErrors, IgnoreCrm, IgnoreTransientState, IgnoreCrmAndTransientState",
}
var VMPoolMemberSpecialArgs = map[string]string{}
var VMSpecRequiredArgs = []string{}
var VMSpecOptionalArgs = []string{
	"internalname",
	"externalnetwork",
	"internalnetwork",
}
var VMSpecAliasArgs = []string{}
var VMSpecComments = map[string]string{
	"internalname":    "VM internal name",
	"externalnetwork": "VM has external network defined or not",
	"internalnetwork": "VM has internal network defined or not",
}
var VMSpecSpecialArgs = map[string]string{}
var VMPoolInfoRequiredArgs = []string{
	"vmpool-org",
	"vmpool",
}
var VMPoolInfoOptionalArgs = []string{
	"notifyid",
	"vms:#.name",
	"vms:#.netinfo.externalip",
	"vms:#.netinfo.internalip",
	"vms:#.groupname",
	"vms:#.state",
	"vms:#.updatedat.seconds",
	"vms:#.updatedat.nanos",
	"vms:#.internalname",
	"state",
	"errors",
	"status.tasknumber",
	"status.maxtasks",
	"status.taskname",
	"status.stepname",
}
var VMPoolInfoAliasArgs = []string{
	"vmpool-org=key.organization",
	"vmpool=key.name",
}
var VMPoolInfoComments = map[string]string{
	"fields":                   "Fields are used for the Update API to specify which fields to apply",
	"vmpool-org":               "Organization of the vmpool",
	"vmpool":                   "Name of the vmpool",
	"notifyid":                 "Id of client assigned by server (internal use only)",
	"vms:#.name":               "VM Name",
	"vms:#.netinfo.externalip": "External IP",
	"vms:#.netinfo.internalip": "Internal IP",
	"vms:#.groupname":          "VM Group Name",
	"vms:#.state":              "VM State, one of VmFree, VmInProgress, VmInUse, VmAdd, VmRemove, VmUpdate",
	"vms:#.updatedat.seconds":  "Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.",
	"vms:#.updatedat.nanos":    "Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive.",
	"vms:#.internalname":       "VM Internal Name",
	"state":                    "Current state of the VM pool on the Cloudlet, one of TrackedStateUnknown, NotPresent, CreateRequested, Creating, CreateError, Ready, UpdateRequested, Updating, UpdateError, DeleteRequested, Deleting, DeleteError, DeletePrepare, CrmInitok, CreatingDependencies",
	"errors":                   "Any errors trying to add/remove VM to/from VM Pool",
}
var VMPoolInfoSpecialArgs = map[string]string{
	"errors": "StringArray",
	"fields": "StringArray",
}
var AddVMPoolMemberRequiredArgs = []string{
	"vmpool-org",
	"vmpool",
	"vm.name",
	"vm.netinfo.internalip",
}
var AddVMPoolMemberOptionalArgs = []string{
	"vm.netinfo.externalip",
	"vm.updatedat.seconds",
	"vm.updatedat.nanos",
	"crmoverride",
}
var RemoveVMPoolMemberRequiredArgs = []string{
	"vmpool-org",
	"vmpool",
	"vm.name",
}
var RemoveVMPoolMemberOptionalArgs = []string{
	"vm.netinfo.externalip",
	"vm.netinfo.internalip",
	"vm.updatedat.seconds",
	"vm.updatedat.nanos",
	"crmoverride",
}
