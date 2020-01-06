// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cloudlet.proto

package gencmd

import edgeproto "github.com/mobiledgex/edge-cloud/edgeproto"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT
func CloudletHideTags(in *edgeproto.Cloudlet) {
	if cli.HideTags == "" {
		return
	}
	tags := make(map[string]struct{})
	for _, tag := range strings.Split(cli.HideTags, ",") {
		tags[tag] = struct{}{}
	}
	if _, found := tags["nocmp"]; found {
		in.TimeLimits = edgeproto.OperationTimeLimits{}
	}
	if _, found := tags["nocmp"]; found {
		in.Errors = nil
	}
	if _, found := tags["nocmp"]; found {
		in.Status = edgeproto.StatusInfo{}
	}
	if _, found := tags["nocmp"]; found {
		in.State = 0
	}
	if _, found := tags["nocmp"]; found {
		in.CrmOverride = 0
	}
	if _, found := tags["nocmp"]; found {
		in.DeploymentLocal = false
	}
	if _, found := tags["nocmp"]; found {
		in.NotifySrvAddr = ""
	}
	if _, found := tags["nocmp"]; found {
		in.Config = edgeproto.PlatformConfig{}
	}
}

var CloudletApiCmd edgeproto.CloudletApiClient

var CreateCloudletCmd = &cli.Command{
	Use:          "CreateCloudlet",
	RequiredArgs: strings.Join(CreateCloudletRequiredArgs, " "),
	OptionalArgs: strings.Join(CreateCloudletOptionalArgs, " "),
	AliasArgs:    strings.Join(CloudletAliasArgs, " "),
	SpecialArgs:  &CloudletSpecialArgs,
	Comments:     CloudletComments,
	ReqData:      &edgeproto.Cloudlet{},
	ReplyData:    &edgeproto.Result{},
	Run:          runCreateCloudlet,
}

func runCreateCloudlet(c *cli.Command, args []string) error {
	obj := c.ReqData.(*edgeproto.Cloudlet)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return CreateCloudlet(c, obj)
}

func CreateCloudlet(c *cli.Command, in *edgeproto.Cloudlet) error {
	if CloudletApiCmd == nil {
		return fmt.Errorf("CloudletApi client not initialized")
	}
	ctx := context.Background()
	stream, err := CloudletApiCmd.CreateCloudlet(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("CreateCloudlet failed: %s", errstr)
	}
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
			return fmt.Errorf("CreateCloudlet recv failed: %s", errstr)
		}
		c.WriteOutput(obj, cli.OutputFormat)
	}
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func CreateCloudlets(c *cli.Command, data []edgeproto.Cloudlet, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("CreateCloudlet %v\n", data[ii])
		myerr := CreateCloudlet(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var DeleteCloudletCmd = &cli.Command{
	Use:          "DeleteCloudlet",
	RequiredArgs: strings.Join(CloudletRequiredArgs, " "),
	OptionalArgs: strings.Join(CloudletOptionalArgs, " "),
	AliasArgs:    strings.Join(CloudletAliasArgs, " "),
	SpecialArgs:  &CloudletSpecialArgs,
	Comments:     CloudletComments,
	ReqData:      &edgeproto.Cloudlet{},
	ReplyData:    &edgeproto.Result{},
	Run:          runDeleteCloudlet,
}

func runDeleteCloudlet(c *cli.Command, args []string) error {
	obj := c.ReqData.(*edgeproto.Cloudlet)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return DeleteCloudlet(c, obj)
}

func DeleteCloudlet(c *cli.Command, in *edgeproto.Cloudlet) error {
	if CloudletApiCmd == nil {
		return fmt.Errorf("CloudletApi client not initialized")
	}
	ctx := context.Background()
	stream, err := CloudletApiCmd.DeleteCloudlet(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("DeleteCloudlet failed: %s", errstr)
	}
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
			return fmt.Errorf("DeleteCloudlet recv failed: %s", errstr)
		}
		c.WriteOutput(obj, cli.OutputFormat)
	}
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func DeleteCloudlets(c *cli.Command, data []edgeproto.Cloudlet, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("DeleteCloudlet %v\n", data[ii])
		myerr := DeleteCloudlet(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var UpdateCloudletCmd = &cli.Command{
	Use:          "UpdateCloudlet",
	RequiredArgs: strings.Join(CloudletRequiredArgs, " "),
	OptionalArgs: strings.Join(CloudletOptionalArgs, " "),
	AliasArgs:    strings.Join(CloudletAliasArgs, " "),
	SpecialArgs:  &CloudletSpecialArgs,
	Comments:     CloudletComments,
	ReqData:      &edgeproto.Cloudlet{},
	ReplyData:    &edgeproto.Result{},
	Run:          runUpdateCloudlet,
}

func runUpdateCloudlet(c *cli.Command, args []string) error {
	obj := c.ReqData.(*edgeproto.Cloudlet)
	jsonMap, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	obj.Fields = cli.GetSpecifiedFields(jsonMap, c.ReqData, cli.JsonNamespace)
	return UpdateCloudlet(c, obj)
}

func UpdateCloudlet(c *cli.Command, in *edgeproto.Cloudlet) error {
	if CloudletApiCmd == nil {
		return fmt.Errorf("CloudletApi client not initialized")
	}
	ctx := context.Background()
	stream, err := CloudletApiCmd.UpdateCloudlet(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("UpdateCloudlet failed: %s", errstr)
	}
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
			return fmt.Errorf("UpdateCloudlet recv failed: %s", errstr)
		}
		c.WriteOutput(obj, cli.OutputFormat)
	}
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func UpdateCloudlets(c *cli.Command, data []edgeproto.Cloudlet, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("UpdateCloudlet %v\n", data[ii])
		myerr := UpdateCloudlet(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var ShowCloudletCmd = &cli.Command{
	Use:          "ShowCloudlet",
	OptionalArgs: strings.Join(append(CloudletRequiredArgs, CloudletOptionalArgs...), " "),
	AliasArgs:    strings.Join(CloudletAliasArgs, " "),
	SpecialArgs:  &CloudletSpecialArgs,
	Comments:     CloudletComments,
	ReqData:      &edgeproto.Cloudlet{},
	ReplyData:    &edgeproto.Cloudlet{},
	Run:          runShowCloudlet,
}

func runShowCloudlet(c *cli.Command, args []string) error {
	obj := c.ReqData.(*edgeproto.Cloudlet)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return ShowCloudlet(c, obj)
}

func ShowCloudlet(c *cli.Command, in *edgeproto.Cloudlet) error {
	if CloudletApiCmd == nil {
		return fmt.Errorf("CloudletApi client not initialized")
	}
	ctx := context.Background()
	stream, err := CloudletApiCmd.ShowCloudlet(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("ShowCloudlet failed: %s", errstr)
	}
	objs := make([]*edgeproto.Cloudlet, 0)
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
			return fmt.Errorf("ShowCloudlet recv failed: %s", errstr)
		}
		CloudletHideTags(obj)
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		return nil
	}
	c.WriteOutput(objs, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func ShowCloudlets(c *cli.Command, data []edgeproto.Cloudlet, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("ShowCloudlet %v\n", data[ii])
		myerr := ShowCloudlet(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var AddCloudletResMappingCmd = &cli.Command{
	Use:          "AddCloudletResMapping",
	RequiredArgs: strings.Join(CloudletResMapRequiredArgs, " "),
	OptionalArgs: strings.Join(CloudletResMapOptionalArgs, " "),
	AliasArgs:    strings.Join(CloudletResMapAliasArgs, " "),
	SpecialArgs:  &CloudletResMapSpecialArgs,
	Comments:     CloudletResMapComments,
	ReqData:      &edgeproto.CloudletResMap{},
	ReplyData:    &edgeproto.Result{},
	Run:          runAddCloudletResMapping,
}

func runAddCloudletResMapping(c *cli.Command, args []string) error {
	obj := c.ReqData.(*edgeproto.CloudletResMap)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return AddCloudletResMapping(c, obj)
}

func AddCloudletResMapping(c *cli.Command, in *edgeproto.CloudletResMap) error {
	if CloudletApiCmd == nil {
		return fmt.Errorf("CloudletApi client not initialized")
	}
	ctx := context.Background()
	obj, err := CloudletApiCmd.AddCloudletResMapping(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("AddCloudletResMapping failed: %s", errstr)
	}
	c.WriteOutput(obj, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func AddCloudletResMappings(c *cli.Command, data []edgeproto.CloudletResMap, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("AddCloudletResMapping %v\n", data[ii])
		myerr := AddCloudletResMapping(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var RemoveCloudletResMappingCmd = &cli.Command{
	Use:          "RemoveCloudletResMapping",
	RequiredArgs: strings.Join(CloudletResMapRequiredArgs, " "),
	OptionalArgs: strings.Join(CloudletResMapOptionalArgs, " "),
	AliasArgs:    strings.Join(CloudletResMapAliasArgs, " "),
	SpecialArgs:  &CloudletResMapSpecialArgs,
	Comments:     CloudletResMapComments,
	ReqData:      &edgeproto.CloudletResMap{},
	ReplyData:    &edgeproto.Result{},
	Run:          runRemoveCloudletResMapping,
}

func runRemoveCloudletResMapping(c *cli.Command, args []string) error {
	obj := c.ReqData.(*edgeproto.CloudletResMap)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return RemoveCloudletResMapping(c, obj)
}

func RemoveCloudletResMapping(c *cli.Command, in *edgeproto.CloudletResMap) error {
	if CloudletApiCmd == nil {
		return fmt.Errorf("CloudletApi client not initialized")
	}
	ctx := context.Background()
	obj, err := CloudletApiCmd.RemoveCloudletResMapping(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("RemoveCloudletResMapping failed: %s", errstr)
	}
	c.WriteOutput(obj, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func RemoveCloudletResMappings(c *cli.Command, data []edgeproto.CloudletResMap, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("RemoveCloudletResMapping %v\n", data[ii])
		myerr := RemoveCloudletResMapping(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var FindFlavorMatchCmd = &cli.Command{
	Use:          "FindFlavorMatch",
	RequiredArgs: strings.Join(FlavorMatchRequiredArgs, " "),
	OptionalArgs: strings.Join(FlavorMatchOptionalArgs, " "),
	AliasArgs:    strings.Join(FlavorMatchAliasArgs, " "),
	SpecialArgs:  &FlavorMatchSpecialArgs,
	Comments:     FlavorMatchComments,
	ReqData:      &edgeproto.FlavorMatch{},
	ReplyData:    &edgeproto.FlavorMatch{},
	Run:          runFindFlavorMatch,
}

func runFindFlavorMatch(c *cli.Command, args []string) error {
	obj := c.ReqData.(*edgeproto.FlavorMatch)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return FindFlavorMatch(c, obj)
}

func FindFlavorMatch(c *cli.Command, in *edgeproto.FlavorMatch) error {
	if CloudletApiCmd == nil {
		return fmt.Errorf("CloudletApi client not initialized")
	}
	ctx := context.Background()
	obj, err := CloudletApiCmd.FindFlavorMatch(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("FindFlavorMatch failed: %s", errstr)
	}
	c.WriteOutput(obj, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func FindFlavorMatchs(c *cli.Command, data []edgeproto.FlavorMatch, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("FindFlavorMatch %v\n", data[ii])
		myerr := FindFlavorMatch(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var CloudletApiCmds = []*cobra.Command{
	CreateCloudletCmd.GenCmd(),
	DeleteCloudletCmd.GenCmd(),
	UpdateCloudletCmd.GenCmd(),
	ShowCloudletCmd.GenCmd(),
	AddCloudletResMappingCmd.GenCmd(),
	RemoveCloudletResMappingCmd.GenCmd(),
	FindFlavorMatchCmd.GenCmd(),
}

var CloudletInfoApiCmd edgeproto.CloudletInfoApiClient

var ShowCloudletInfoCmd = &cli.Command{
	Use:          "ShowCloudletInfo",
	OptionalArgs: strings.Join(append(CloudletInfoRequiredArgs, CloudletInfoOptionalArgs...), " "),
	AliasArgs:    strings.Join(CloudletInfoAliasArgs, " "),
	SpecialArgs:  &CloudletInfoSpecialArgs,
	Comments:     CloudletInfoComments,
	ReqData:      &edgeproto.CloudletInfo{},
	ReplyData:    &edgeproto.CloudletInfo{},
	Run:          runShowCloudletInfo,
}

func runShowCloudletInfo(c *cli.Command, args []string) error {
	obj := c.ReqData.(*edgeproto.CloudletInfo)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return ShowCloudletInfo(c, obj)
}

func ShowCloudletInfo(c *cli.Command, in *edgeproto.CloudletInfo) error {
	if CloudletInfoApiCmd == nil {
		return fmt.Errorf("CloudletInfoApi client not initialized")
	}
	ctx := context.Background()
	stream, err := CloudletInfoApiCmd.ShowCloudletInfo(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("ShowCloudletInfo failed: %s", errstr)
	}
	objs := make([]*edgeproto.CloudletInfo, 0)
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
			return fmt.Errorf("ShowCloudletInfo recv failed: %s", errstr)
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		return nil
	}
	c.WriteOutput(objs, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func ShowCloudletInfos(c *cli.Command, data []edgeproto.CloudletInfo, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("ShowCloudletInfo %v\n", data[ii])
		myerr := ShowCloudletInfo(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var InjectCloudletInfoCmd = &cli.Command{
	Use:          "InjectCloudletInfo",
	RequiredArgs: strings.Join(CloudletInfoRequiredArgs, " "),
	OptionalArgs: strings.Join(CloudletInfoOptionalArgs, " "),
	AliasArgs:    strings.Join(CloudletInfoAliasArgs, " "),
	SpecialArgs:  &CloudletInfoSpecialArgs,
	Comments:     CloudletInfoComments,
	ReqData:      &edgeproto.CloudletInfo{},
	ReplyData:    &edgeproto.Result{},
	Run:          runInjectCloudletInfo,
}

func runInjectCloudletInfo(c *cli.Command, args []string) error {
	obj := c.ReqData.(*edgeproto.CloudletInfo)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return InjectCloudletInfo(c, obj)
}

func InjectCloudletInfo(c *cli.Command, in *edgeproto.CloudletInfo) error {
	if CloudletInfoApiCmd == nil {
		return fmt.Errorf("CloudletInfoApi client not initialized")
	}
	ctx := context.Background()
	obj, err := CloudletInfoApiCmd.InjectCloudletInfo(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("InjectCloudletInfo failed: %s", errstr)
	}
	c.WriteOutput(obj, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func InjectCloudletInfos(c *cli.Command, data []edgeproto.CloudletInfo, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("InjectCloudletInfo %v\n", data[ii])
		myerr := InjectCloudletInfo(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var EvictCloudletInfoCmd = &cli.Command{
	Use:          "EvictCloudletInfo",
	RequiredArgs: strings.Join(CloudletInfoRequiredArgs, " "),
	OptionalArgs: strings.Join(CloudletInfoOptionalArgs, " "),
	AliasArgs:    strings.Join(CloudletInfoAliasArgs, " "),
	SpecialArgs:  &CloudletInfoSpecialArgs,
	Comments:     CloudletInfoComments,
	ReqData:      &edgeproto.CloudletInfo{},
	ReplyData:    &edgeproto.Result{},
	Run:          runEvictCloudletInfo,
}

func runEvictCloudletInfo(c *cli.Command, args []string) error {
	obj := c.ReqData.(*edgeproto.CloudletInfo)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return EvictCloudletInfo(c, obj)
}

func EvictCloudletInfo(c *cli.Command, in *edgeproto.CloudletInfo) error {
	if CloudletInfoApiCmd == nil {
		return fmt.Errorf("CloudletInfoApi client not initialized")
	}
	ctx := context.Background()
	obj, err := CloudletInfoApiCmd.EvictCloudletInfo(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("EvictCloudletInfo failed: %s", errstr)
	}
	c.WriteOutput(obj, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func EvictCloudletInfos(c *cli.Command, data []edgeproto.CloudletInfo, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("EvictCloudletInfo %v\n", data[ii])
		myerr := EvictCloudletInfo(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var CloudletInfoApiCmds = []*cobra.Command{
	ShowCloudletInfoCmd.GenCmd(),
	InjectCloudletInfoCmd.GenCmd(),
	EvictCloudletInfoCmd.GenCmd(),
}

var CloudletMetricsApiCmd edgeproto.CloudletMetricsApiClient

var ShowCloudletMetricsCmd = &cli.Command{
	Use:          "ShowCloudletMetrics",
	OptionalArgs: strings.Join(append(CloudletMetricsRequiredArgs, CloudletMetricsOptionalArgs...), " "),
	AliasArgs:    strings.Join(CloudletMetricsAliasArgs, " "),
	SpecialArgs:  &CloudletMetricsSpecialArgs,
	Comments:     CloudletMetricsComments,
	ReqData:      &edgeproto.CloudletMetrics{},
	ReplyData:    &edgeproto.CloudletMetrics{},
	Run:          runShowCloudletMetrics,
}

func runShowCloudletMetrics(c *cli.Command, args []string) error {
	obj := c.ReqData.(*edgeproto.CloudletMetrics)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return ShowCloudletMetrics(c, obj)
}

func ShowCloudletMetrics(c *cli.Command, in *edgeproto.CloudletMetrics) error {
	if CloudletMetricsApiCmd == nil {
		return fmt.Errorf("CloudletMetricsApi client not initialized")
	}
	ctx := context.Background()
	stream, err := CloudletMetricsApiCmd.ShowCloudletMetrics(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("ShowCloudletMetrics failed: %s", errstr)
	}
	objs := make([]*edgeproto.CloudletMetrics, 0)
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
			return fmt.Errorf("ShowCloudletMetrics recv failed: %s", errstr)
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		return nil
	}
	c.WriteOutput(objs, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func ShowCloudletMetricss(c *cli.Command, data []edgeproto.CloudletMetrics, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("ShowCloudletMetrics %v\n", data[ii])
		myerr := ShowCloudletMetrics(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var CloudletMetricsApiCmds = []*cobra.Command{
	ShowCloudletMetricsCmd.GenCmd(),
}

var CloudletKeyRequiredArgs = []string{}
var CloudletKeyOptionalArgs = []string{
	"operator",
	"name",
}
var CloudletKeyAliasArgs = []string{
	"operator=operatorkey.name",
}
var CloudletKeyComments = map[string]string{
	"operator": "Company or Organization name of the operator",
	"name":     "Name of the cloudlet",
}
var CloudletKeySpecialArgs = map[string]string{}
var OperationTimeLimitsRequiredArgs = []string{}
var OperationTimeLimitsOptionalArgs = []string{
	"createclusterinsttimeout",
	"updateclusterinsttimeout",
	"deleteclusterinsttimeout",
	"createappinsttimeout",
	"updateappinsttimeout",
	"deleteappinsttimeout",
}
var OperationTimeLimitsAliasArgs = []string{}
var OperationTimeLimitsComments = map[string]string{
	"createclusterinsttimeout": "max time to create a cluster instance",
	"updateclusterinsttimeout": "max time to update a cluster instance",
	"deleteclusterinsttimeout": "max time to delete a cluster instance",
	"createappinsttimeout":     "max time to create an app instance",
	"updateappinsttimeout":     "max time to update an app instance",
	"deleteappinsttimeout":     "max time to delete an app instance",
}
var OperationTimeLimitsSpecialArgs = map[string]string{}
var CloudletInfraCommonRequiredArgs = []string{}
var CloudletInfraCommonOptionalArgs = []string{
	"dockerregistry",
	"dnszone",
	"registryfileserver",
	"cfkey",
	"cfuser",
	"dockerregpass",
	"networkscheme",
	"dockerregistrysecret",
}
var CloudletInfraCommonAliasArgs = []string{}
var CloudletInfraCommonComments = map[string]string{
	"dockerregistry":       "the mex docker registry, e.g.  registry.mobiledgex.net:5000.",
	"dnszone":              "DNS Zone",
	"registryfileserver":   "registry file server contains files which get pulled on instantiation such as certs and images",
	"cfkey":                "Cloudflare key",
	"cfuser":               "Cloudflare key",
	"dockerregpass":        "Docker registry password",
	"networkscheme":        "network scheme",
	"dockerregistrysecret": "the name of the docker registry secret, e.g. mexgitlabsecret",
}
var CloudletInfraCommonSpecialArgs = map[string]string{}
var AzurePropertiesRequiredArgs = []string{}
var AzurePropertiesOptionalArgs = []string{
	"location",
	"resourcegroup",
	"username",
	"password",
}
var AzurePropertiesAliasArgs = []string{}
var AzurePropertiesComments = map[string]string{
	"location":      "azure region e.g. uswest2",
	"resourcegroup": "azure resource group",
	"username":      "azure username",
	"password":      "azure password",
}
var AzurePropertiesSpecialArgs = map[string]string{}
var GcpPropertiesRequiredArgs = []string{}
var GcpPropertiesOptionalArgs = []string{
	"project",
	"zone",
	"serviceaccount",
	"gcpauthkeyurl",
}
var GcpPropertiesAliasArgs = []string{}
var GcpPropertiesComments = map[string]string{
	"project":        "gcp project for billing",
	"zone":           "availability zone",
	"serviceaccount": "service account to login with",
	"gcpauthkeyurl":  "vault credentials link",
}
var GcpPropertiesSpecialArgs = map[string]string{}
var OpenStackPropertiesRequiredArgs = []string{}
var OpenStackPropertiesOptionalArgs = []string{
	"osexternalnetworkname",
	"osimagename",
	"osexternalroutername",
	"osmexnetwork",
	"openrcvars",
}
var OpenStackPropertiesAliasArgs = []string{}
var OpenStackPropertiesComments = map[string]string{
	"osexternalnetworkname": "name of the external network, e.g. external-network-shared",
	"osimagename":           "openstack image , e.g. mobiledgex",
	"osexternalroutername":  "openstack router",
	"osmexnetwork":          "openstack internal network",
	"openrcvars":            "openrc env vars",
}
var OpenStackPropertiesSpecialArgs = map[string]string{
	"openrcvars": "StringToString",
}
var CloudletInfraPropertiesRequiredArgs = []string{}
var CloudletInfraPropertiesOptionalArgs = []string{
	"cloudletkind",
	"mexoscontainerimagename",
	"openstackproperties.osexternalnetworkname",
	"openstackproperties.osimagename",
	"openstackproperties.osexternalroutername",
	"openstackproperties.osmexnetwork",
	"openstackproperties.openrcvars",
	"azureproperties.location",
	"azureproperties.resourcegroup",
	"azureproperties.username",
	"azureproperties.password",
	"gcpproperties.project",
	"gcpproperties.zone",
	"gcpproperties.serviceaccount",
	"gcpproperties.gcpauthkeyurl",
}
var CloudletInfraPropertiesAliasArgs = []string{}
var CloudletInfraPropertiesComments = map[string]string{
	"cloudletkind":                              "what kind of infrastructure: Azure, GCP, Openstack",
	"mexoscontainerimagename":                   "name and version of the docker image container image that mexos runs in",
	"openstackproperties.osexternalnetworkname": "name of the external network, e.g. external-network-shared",
	"openstackproperties.osimagename":           "openstack image , e.g. mobiledgex",
	"openstackproperties.osexternalroutername":  "openstack router",
	"openstackproperties.osmexnetwork":          "openstack internal network",
	"openstackproperties.openrcvars":            "openrc env vars",
	"azureproperties.location":                  "azure region e.g. uswest2",
	"azureproperties.resourcegroup":             "azure resource group",
	"azureproperties.username":                  "azure username",
	"azureproperties.password":                  "azure password",
	"gcpproperties.project":                     "gcp project for billing",
	"gcpproperties.zone":                        "availability zone",
	"gcpproperties.serviceaccount":              "service account to login with",
	"gcpproperties.gcpauthkeyurl":               "vault credentials link",
}
var CloudletInfraPropertiesSpecialArgs = map[string]string{
	"openstackproperties.openrcvars": "StringToString",
}
var PlatformConfigRequiredArgs = []string{}
var PlatformConfigOptionalArgs = []string{
	"registrypath",
	"imagepath",
	"notifyctrladdrs",
	"vaultaddr",
	"tlscertfile",
	"envvar",
	"platformtag",
	"testmode",
	"span",
	"cleanupmode",
	"region",
}
var PlatformConfigAliasArgs = []string{}
var PlatformConfigComments = map[string]string{
	"registrypath":    "Path to Docker registry holding edge-cloud image",
	"imagepath":       "Path to platform base image",
	"notifyctrladdrs": "Address of controller notify port (can be multiple of these)",
	"vaultaddr":       "Vault address",
	"tlscertfile":     "TLS cert file",
	"envvar":          "Environment variables",
	"platformtag":     "Tag of edge-cloud image",
	"testmode":        "Internal Test flag",
	"span":            "Span string",
	"cleanupmode":     "Internal cleanup flag",
	"region":          "Region",
}
var PlatformConfigSpecialArgs = map[string]string{
	"envvar": "StringToString",
}
var CloudletResMapRequiredArgs = []string{
	"operator",
	"name",
	"mapping",
}
var CloudletResMapOptionalArgs = []string{}
var CloudletResMapAliasArgs = []string{
	"operator=key.operatorkey.name",
	"name=key.name",
}
var CloudletResMapComments = map[string]string{
	"operator": "Company or Organization name of the operator",
	"name":     "Name of the cloudlet",
	"mapping":  "Resource mapping info",
}
var CloudletResMapSpecialArgs = map[string]string{
	"mapping": "StringToString",
}
var CloudletRequiredArgs = []string{
	"operator",
	"name",
}
var CloudletOptionalArgs = []string{
	"location.latitude",
	"location.longitude",
	"location.altitude",
	"location.timestamp.seconds",
	"location.timestamp.nanos",
	"ipsupport",
	"staticips",
	"numdynamicips",
	"errors",
	"state",
	"crmoverride",
	"deploymentlocal",
	"platformtype",
	"flavor.name",
	"physicalname",
	"envvar",
	"version",
	"restagmap.key",
	"restagmap.value.name",
	"restagmap.value.operatorkey.name",
	"accessvars",
}
var CloudletAliasArgs = []string{
	"operator=key.operatorkey.name",
	"name=key.name",
}
var CloudletComments = map[string]string{
	"operator":                            "Company or Organization name of the operator",
	"name":                                "Name of the cloudlet",
	"location.latitude":                   "latitude in WGS 84 coordinates",
	"location.longitude":                  "longitude in WGS 84 coordinates",
	"location.horizontalaccuracy":         "horizontal accuracy (radius in meters)",
	"location.verticalaccuracy":           "vertical accuracy (meters)",
	"location.altitude":                   "On android only lat and long are guaranteed to be supplied altitude in meters",
	"location.course":                     "course (IOS) / bearing (Android) (degrees east relative to true north)",
	"location.speed":                      "speed (IOS) / velocity (Android) (meters/sec)",
	"ipsupport":                           "Type of IP support provided by Cloudlet (see IpSupport), one of IpSupportUnknown, IpSupportStatic, IpSupportDynamic",
	"staticips":                           "List of static IPs for static IP support",
	"numdynamicips":                       "Number of dynamic IPs available for dynamic IP support",
	"timelimits.createclusterinsttimeout": "max time to create a cluster instance",
	"timelimits.updateclusterinsttimeout": "max time to update a cluster instance",
	"timelimits.deleteclusterinsttimeout": "max time to delete a cluster instance",
	"timelimits.createappinsttimeout":     "max time to create an app instance",
	"timelimits.updateappinsttimeout":     "max time to update an app instance",
	"timelimits.deleteappinsttimeout":     "max time to delete an app instance",
	"errors":                              "Any errors trying to create, update, or delete the Cloudlet.",
	"state":                               "Current state of the cloudlet, one of TrackedStateUnknown, NotPresent, CreateRequested, Creating, CreateError, Ready, UpdateRequested, Updating, UpdateError, DeleteRequested, Deleting, DeleteError, DeletePrepare, CrmInitok, CreatingDependencies",
	"crmoverride":                         "Override actions to CRM, one of NoOverride, IgnoreCrmErrors, IgnoreCrm, IgnoreTransientState, IgnoreCrmAndTransientState",
	"deploymentlocal":                     "Deploy cloudlet services locally",
	"platformtype":                        "Platform type, one of PlatformTypeFake, PlatformTypeDind, PlatformTypeOpenstack, PlatformTypeAzure, PlatformTypeGcp, PlatformTypeEdgebox, PlatformTypeFakeinfra",
	"notifysrvaddr":                       "Address for the CRM notify listener to run on",
	"flavor.name":                         "Flavor name",
	"physicalname":                        "Physical infrastructure cloudlet name",
	"envvar":                              "Single Key-Value pair of env var to be passed to CRM",
	"version":                             "Cloudlet version",
	"config.registrypath":                 "Path to Docker registry holding edge-cloud image",
	"config.imagepath":                    "Path to platform base image",
	"config.notifyctrladdrs":              "Address of controller notify port (can be multiple of these)",
	"config.vaultaddr":                    "Vault address",
	"config.tlscertfile":                  "TLS cert file",
	"config.envvar":                       "Environment variables",
	"config.platformtag":                  "Tag of edge-cloud image",
	"config.testmode":                     "Internal Test flag",
	"config.span":                         "Span string",
	"config.cleanupmode":                  "Internal cleanup flag",
	"config.region":                       "Region",
	"restagmap.value.name":                "Resource Table Name",
	"restagmap.value.operatorkey.name":    "Company or Organization name of the operator",
	"accessvars":                          "Variables required to access cloudlet",
}
var CloudletSpecialArgs = map[string]string{
	"accessvars":    "StringToString",
	"config.envvar": "StringToString",
	"envvar":        "StringToString",
	"errors":        "StringArray",
}
var FlavorMatchRequiredArgs = []string{
	"operator",
	"cloudlet",
}
var FlavorMatchOptionalArgs = []string{
	"flavor",
	"availabilityzone",
}
var FlavorMatchAliasArgs = []string{
	"operator=key.operatorkey.name",
	"cloudlet=key.name",
	"flavor=flavorname",
}
var FlavorMatchComments = map[string]string{
	"operator": "Company or Organization name of the operator",
	"cloudlet": "Name of the cloudlet",
}
var FlavorMatchSpecialArgs = map[string]string{}
var FlavorInfoRequiredArgs = []string{}
var FlavorInfoOptionalArgs = []string{
	"name",
	"vcpus",
	"ram",
	"disk",
	"properties",
}
var FlavorInfoAliasArgs = []string{}
var FlavorInfoComments = map[string]string{
	"name":       "Name of the flavor on the Cloudlet",
	"vcpus":      "Number of VCPU cores on the Cloudlet",
	"ram":        "Ram in MB on the Cloudlet",
	"disk":       "Amount of disk in GB on the Cloudlet",
	"properties": "OS Flavor Properties, if any",
}
var FlavorInfoSpecialArgs = map[string]string{}
var OSAZoneRequiredArgs = []string{}
var OSAZoneOptionalArgs = []string{
	"name",
	"status",
}
var OSAZoneAliasArgs = []string{}
var OSAZoneComments = map[string]string{}
var OSAZoneSpecialArgs = map[string]string{}
var OSImageRequiredArgs = []string{}
var OSImageOptionalArgs = []string{
	"name",
	"tags",
	"properties",
	"diskformat",
}
var OSImageAliasArgs = []string{}
var OSImageComments = map[string]string{
	"name":       "image name",
	"tags":       "optional tags present on image",
	"properties": "image properties/metadata",
	"diskformat": "format qcow2, img, etc",
}
var OSImageSpecialArgs = map[string]string{}
var CloudletInfoRequiredArgs = []string{
	"operator",
	"name",
}
var CloudletInfoOptionalArgs = []string{
	"state",
	"notifyid",
	"controller",
	"osmaxram",
	"osmaxvcores",
	"osmaxvolgb",
	"errors",
	"flavors.name",
	"flavors.vcpus",
	"flavors.ram",
	"flavors.disk",
	"flavors.properties",
	"status.tasknumber",
	"status.maxtasks",
	"status.taskname",
	"status.stepname",
	"version",
	"availabilityzones.name",
	"availabilityzones.status",
	"osimages.name",
	"osimages.tags",
	"osimages.properties",
	"osimages.diskformat",
}
var CloudletInfoAliasArgs = []string{
	"operator=key.operatorkey.name",
	"name=key.name",
}
var CloudletInfoComments = map[string]string{
	"operator":            "Company or Organization name of the operator",
	"name":                "Name of the cloudlet",
	"state":               "State of cloudlet, one of CloudletStateUnknown, CloudletStateErrors, CloudletStateReady, CloudletStateOffline, CloudletStateNotPresent, CloudletStateInit, CloudletStateUpgrade",
	"notifyid":            "Id of client assigned by server (internal use only)",
	"controller":          "Connected controller unique id",
	"osmaxram":            "Maximum Ram in MB on the Cloudlet",
	"osmaxvcores":         "Maximum number of VCPU cores on the Cloudlet",
	"osmaxvolgb":          "Maximum amount of disk in GB on the Cloudlet",
	"errors":              "Any errors encountered while making changes to the Cloudlet",
	"flavors.name":        "Name of the flavor on the Cloudlet",
	"flavors.vcpus":       "Number of VCPU cores on the Cloudlet",
	"flavors.ram":         "Ram in MB on the Cloudlet",
	"flavors.disk":        "Amount of disk in GB on the Cloudlet",
	"flavors.properties":  "OS Flavor Properties, if any",
	"version":             "Cloudlet version",
	"osimages.name":       "image name",
	"osimages.tags":       "optional tags present on image",
	"osimages.properties": "image properties/metadata",
	"osimages.diskformat": "format qcow2, img, etc",
}
var CloudletInfoSpecialArgs = map[string]string{
	"errors": "StringArray",
}
var CloudletMetricsRequiredArgs = []string{}
var CloudletMetricsOptionalArgs = []string{
	"foo",
}
var CloudletMetricsAliasArgs = []string{}
var CloudletMetricsComments = map[string]string{
	"foo": "what goes here?",
}
var CloudletMetricsSpecialArgs = map[string]string{}
var CreateCloudletRequiredArgs = []string{
	"operator",
	"name",
	"location.latitude",
	"location.longitude",
	"numdynamicips",
}
var CreateCloudletOptionalArgs = []string{
	"location.altitude",
	"location.timestamp.seconds",
	"location.timestamp.nanos",
	"ipsupport",
	"staticips",
	"errors",
	"state",
	"crmoverride",
	"deploymentlocal",
	"platformtype",
	"flavor.name",
	"physicalname",
	"envvar",
	"version",
	"restagmap.key",
	"restagmap.value.name",
	"restagmap.value.operatorkey.name",
	"accessvars",
}
