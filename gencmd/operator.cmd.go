// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: operator.proto

package gencmd

import edgeproto "github.com/mobiledgex/edge-cloud/edgeproto"
import "strings"
import "time"
import "github.com/spf13/cobra"
import "context"
import "os"
import "io"
import "text/tabwriter"
import "github.com/spf13/pflag"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/mobiledgex/edge-cloud/protogen"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT
var OperatorApiCmd edgeproto.OperatorApiClient
var OperatorIn edgeproto.Operator
var OperatorFlagSet = pflag.NewFlagSet("Operator", pflag.ExitOnError)

func OperatorCodeSlicer(in *edgeproto.OperatorCode) []string {
	s := make([]string, 0, 2)
	s = append(s, in.MNC)
	s = append(s, in.MCC)
	return s
}

func OperatorCodeHeaderSlicer() []string {
	s := make([]string, 0, 2)
	s = append(s, "MNC")
	s = append(s, "MCC")
	return s
}

func OperatorKeySlicer(in *edgeproto.OperatorKey) []string {
	s := make([]string, 0, 1)
	s = append(s, in.Name)
	return s
}

func OperatorKeyHeaderSlicer() []string {
	s := make([]string, 0, 1)
	s = append(s, "Name")
	return s
}

func OperatorSlicer(in *edgeproto.Operator) []string {
	s := make([]string, 0, 2)
	if in.Fields == nil {
		in.Fields = make([]string, 1)
	}
	s = append(s, in.Fields[0])
	s = append(s, in.Key.Name)
	return s
}

func OperatorHeaderSlicer() []string {
	s := make([]string, 0, 2)
	s = append(s, "Fields")
	s = append(s, "Key-Name")
	return s
}

var CreateOperatorCmd = &cobra.Command{
	Use: "CreateOperator",
	Run: func(cmd *cobra.Command, args []string) {
		if OperatorApiCmd == nil {
			fmt.Println("OperatorApi client not initialized")
			return
		}
		var err error
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		out, err := OperatorApiCmd.CreateOperator(ctx, &OperatorIn)
		cancel()
		if err != nil {
			fmt.Println("CreateOperator failed: ", err)
		} else {
			headers := ResultHeaderSlicer()
			data := ResultSlicer(out)
			for ii := 0; ii < len(headers) && ii < len(data); ii++ {
				fmt.Println(headers[ii] + ": " + data[ii])
			}
		}
	},
}

var DeleteOperatorCmd = &cobra.Command{
	Use: "DeleteOperator",
	Run: func(cmd *cobra.Command, args []string) {
		if OperatorApiCmd == nil {
			fmt.Println("OperatorApi client not initialized")
			return
		}
		var err error
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		out, err := OperatorApiCmd.DeleteOperator(ctx, &OperatorIn)
		cancel()
		if err != nil {
			fmt.Println("DeleteOperator failed: ", err)
		} else {
			headers := ResultHeaderSlicer()
			data := ResultSlicer(out)
			for ii := 0; ii < len(headers) && ii < len(data); ii++ {
				fmt.Println(headers[ii] + ": " + data[ii])
			}
		}
	},
}

var UpdateOperatorCmd = &cobra.Command{
	Use: "UpdateOperator",
	Run: func(cmd *cobra.Command, args []string) {
		if OperatorApiCmd == nil {
			fmt.Println("OperatorApi client not initialized")
			return
		}
		var err error
		OperatorSetFields()
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		out, err := OperatorApiCmd.UpdateOperator(ctx, &OperatorIn)
		cancel()
		if err != nil {
			fmt.Println("UpdateOperator failed: ", err)
		} else {
			headers := ResultHeaderSlicer()
			data := ResultSlicer(out)
			for ii := 0; ii < len(headers) && ii < len(data); ii++ {
				fmt.Println(headers[ii] + ": " + data[ii])
			}
		}
	},
}

var ShowOperatorCmd = &cobra.Command{
	Use: "ShowOperator",
	Run: func(cmd *cobra.Command, args []string) {
		if OperatorApiCmd == nil {
			fmt.Println("OperatorApi client not initialized")
			return
		}
		var err error
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		count := 0
		fmt.Fprintln(output, strings.Join(OperatorHeaderSlicer(), "\t"))
		defer cancel()
		stream, err := OperatorApiCmd.ShowOperator(ctx, &OperatorIn)
		if err != nil {
			fmt.Println("ShowOperator failed: ", err)
			return
		}
		for {
			obj, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("ShowOperator recv failed: ", err)
				break
			}
			fmt.Fprintln(output, strings.Join(OperatorSlicer(obj), "\t"))
			count++
		}
		if count > 0 {
			output.Flush()
		}
	},
}

func init() {
	OperatorFlagSet.StringVar(&OperatorIn.Key.Name, "key-name", "", "Key.Name")
	CreateOperatorCmd.Flags().AddFlagSet(OperatorFlagSet)
	DeleteOperatorCmd.Flags().AddFlagSet(OperatorFlagSet)
	UpdateOperatorCmd.Flags().AddFlagSet(OperatorFlagSet)
	ShowOperatorCmd.Flags().AddFlagSet(OperatorFlagSet)
}

func OperatorSetFields() {
	OperatorIn.Fields = make([]string, 0)
	if OperatorFlagSet.Lookup("key-name").Changed {
		OperatorIn.Fields = append(OperatorIn.Fields, "2.1")
	}
}
