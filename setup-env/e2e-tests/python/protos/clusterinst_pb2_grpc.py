# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import clusterinst_pb2 as clusterinst__pb2
import result_pb2 as result__pb2


class ClusterInstApiStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.CreateClusterInst = channel.unary_stream(
        '/edgeproto.ClusterInstApi/CreateClusterInst',
        request_serializer=clusterinst__pb2.ClusterInst.SerializeToString,
        response_deserializer=result__pb2.Result.FromString,
        )
    self.DeleteClusterInst = channel.unary_stream(
        '/edgeproto.ClusterInstApi/DeleteClusterInst',
        request_serializer=clusterinst__pb2.ClusterInst.SerializeToString,
        response_deserializer=result__pb2.Result.FromString,
        )
    self.UpdateClusterInst = channel.unary_stream(
        '/edgeproto.ClusterInstApi/UpdateClusterInst',
        request_serializer=clusterinst__pb2.ClusterInst.SerializeToString,
        response_deserializer=result__pb2.Result.FromString,
        )
    self.ShowClusterInst = channel.unary_stream(
        '/edgeproto.ClusterInstApi/ShowClusterInst',
        request_serializer=clusterinst__pb2.ClusterInst.SerializeToString,
        response_deserializer=clusterinst__pb2.ClusterInst.FromString,
        )


class ClusterInstApiServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def CreateClusterInst(self, request, context):
    """Create a Cluster instance
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteClusterInst(self, request, context):
    """Delete a Cluster instance
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UpdateClusterInst(self, request, context):
    """Update a Cluster instance
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ShowClusterInst(self, request, context):
    """Show Cluster instances
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_ClusterInstApiServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'CreateClusterInst': grpc.unary_stream_rpc_method_handler(
          servicer.CreateClusterInst,
          request_deserializer=clusterinst__pb2.ClusterInst.FromString,
          response_serializer=result__pb2.Result.SerializeToString,
      ),
      'DeleteClusterInst': grpc.unary_stream_rpc_method_handler(
          servicer.DeleteClusterInst,
          request_deserializer=clusterinst__pb2.ClusterInst.FromString,
          response_serializer=result__pb2.Result.SerializeToString,
      ),
      'UpdateClusterInst': grpc.unary_stream_rpc_method_handler(
          servicer.UpdateClusterInst,
          request_deserializer=clusterinst__pb2.ClusterInst.FromString,
          response_serializer=result__pb2.Result.SerializeToString,
      ),
      'ShowClusterInst': grpc.unary_stream_rpc_method_handler(
          servicer.ShowClusterInst,
          request_deserializer=clusterinst__pb2.ClusterInst.FromString,
          response_serializer=clusterinst__pb2.ClusterInst.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'edgeproto.ClusterInstApi', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))


class ClusterInstInfoApiStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.ShowClusterInstInfo = channel.unary_stream(
        '/edgeproto.ClusterInstInfoApi/ShowClusterInstInfo',
        request_serializer=clusterinst__pb2.ClusterInstInfo.SerializeToString,
        response_deserializer=clusterinst__pb2.ClusterInstInfo.FromString,
        )


class ClusterInstInfoApiServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def ShowClusterInstInfo(self, request, context):
    """Show Cluster instances state.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_ClusterInstInfoApiServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'ShowClusterInstInfo': grpc.unary_stream_rpc_method_handler(
          servicer.ShowClusterInstInfo,
          request_deserializer=clusterinst__pb2.ClusterInstInfo.FromString,
          response_serializer=clusterinst__pb2.ClusterInstInfo.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'edgeproto.ClusterInstInfoApi', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
