# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: cluster.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import annotations_pb2 as annotations__pb2
import protogen_pb2 as protogen__pb2
import protocmd_pb2 as protocmd__pb2
import result_pb2 as result__pb2
import developer_pb2 as developer__pb2
import clusterflavor_pb2 as clusterflavor__pb2
import gogo_pb2 as gogo__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='cluster.proto',
  package='edgeproto',
  syntax='proto3',
  serialized_pb=_b('\n\rcluster.proto\x12\tedgeproto\x1a\x11\x61nnotations.proto\x1a\x0eprotogen.proto\x1a\x0eprotocmd.proto\x1a\x0cresult.proto\x1a\x0f\x64\x65veloper.proto\x1a\x13\x63lusterflavor.proto\x1a\ngogo.proto\"(\n\nClusterKey\x12\x0c\n\x04name\x18\x01 \x01(\t:\x0c\xe8\xf3\x18\x01\xf8\xf3\x18\x01\xb0\xa0\x1f\x01\"\xac\x01\n\x07\x43luster\x12\x0e\n\x06\x66ields\x18\x01 \x03(\t\x12(\n\x03key\x18\x02 \x01(\x0b\x32\x15.edgeproto.ClusterKeyB\x04\xc8\xde\x1f\x00\x12\x39\n\x0e\x64\x65\x66\x61ult_flavor\x18\x03 \x01(\x0b\x32\x1b.edgeproto.ClusterFlavorKeyB\x04\xc8\xde\x1f\x00\x12\x12\n\x04\x61uto\x18\x05 \x01(\x08\x42\x04\xa8\xf4\x18\x01:\x18\xe8\xf3\x18\x01\xf0\xf3\x18\x01\x98\xf4\x18\x01\x80\xf4\x18\x01\x8a\xb2\x19\x04\x41uto2\xdb\x02\n\nClusterApi\x12R\n\rCreateCluster\x12\x12.edgeproto.Cluster\x1a\x11.edgeproto.Result\"\x1a\x82\xd3\xe4\x93\x02\x14\"\x0f/create/cluster:\x01*\x12R\n\rDeleteCluster\x12\x12.edgeproto.Cluster\x1a\x11.edgeproto.Result\"\x1a\x82\xd3\xe4\x93\x02\x14\"\x0f/delete/cluster:\x01*\x12R\n\rUpdateCluster\x12\x12.edgeproto.Cluster\x1a\x11.edgeproto.Result\"\x1a\x82\xd3\xe4\x93\x02\x14\"\x0f/update/cluster:\x01*\x12Q\n\x0bShowCluster\x12\x12.edgeproto.Cluster\x1a\x12.edgeproto.Cluster\"\x18\x82\xd3\xe4\x93\x02\x12\"\r/show/cluster:\x01*0\x01\x62\x06proto3')
  ,
  dependencies=[annotations__pb2.DESCRIPTOR,protogen__pb2.DESCRIPTOR,protocmd__pb2.DESCRIPTOR,result__pb2.DESCRIPTOR,developer__pb2.DESCRIPTOR,clusterflavor__pb2.DESCRIPTOR,gogo__pb2.DESCRIPTOR,])




_CLUSTERKEY = _descriptor.Descriptor(
  name='ClusterKey',
  full_name='edgeproto.ClusterKey',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='edgeproto.ClusterKey.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('\350\363\030\001\370\363\030\001\260\240\037\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=143,
  serialized_end=183,
)


_CLUSTER = _descriptor.Descriptor(
  name='Cluster',
  full_name='edgeproto.Cluster',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='fields', full_name='edgeproto.Cluster.fields', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='key', full_name='edgeproto.Cluster.key', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\310\336\037\000')), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default_flavor', full_name='edgeproto.Cluster.default_flavor', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\310\336\037\000')), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='auto', full_name='edgeproto.Cluster.auto', index=3,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\250\364\030\001')), file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('\350\363\030\001\360\363\030\001\230\364\030\001\200\364\030\001\212\262\031\004Auto')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=186,
  serialized_end=358,
)

_CLUSTER.fields_by_name['key'].message_type = _CLUSTERKEY
_CLUSTER.fields_by_name['default_flavor'].message_type = clusterflavor__pb2._CLUSTERFLAVORKEY
DESCRIPTOR.message_types_by_name['ClusterKey'] = _CLUSTERKEY
DESCRIPTOR.message_types_by_name['Cluster'] = _CLUSTER
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ClusterKey = _reflection.GeneratedProtocolMessageType('ClusterKey', (_message.Message,), dict(
  DESCRIPTOR = _CLUSTERKEY,
  __module__ = 'cluster_pb2'
  # @@protoc_insertion_point(class_scope:edgeproto.ClusterKey)
  ))
_sym_db.RegisterMessage(ClusterKey)

Cluster = _reflection.GeneratedProtocolMessageType('Cluster', (_message.Message,), dict(
  DESCRIPTOR = _CLUSTER,
  __module__ = 'cluster_pb2'
  # @@protoc_insertion_point(class_scope:edgeproto.Cluster)
  ))
_sym_db.RegisterMessage(Cluster)


_CLUSTERKEY.has_options = True
_CLUSTERKEY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('\350\363\030\001\370\363\030\001\260\240\037\001'))
_CLUSTER.fields_by_name['key'].has_options = True
_CLUSTER.fields_by_name['key']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\310\336\037\000'))
_CLUSTER.fields_by_name['default_flavor'].has_options = True
_CLUSTER.fields_by_name['default_flavor']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\310\336\037\000'))
_CLUSTER.fields_by_name['auto'].has_options = True
_CLUSTER.fields_by_name['auto']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\250\364\030\001'))
_CLUSTER.has_options = True
_CLUSTER._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('\350\363\030\001\360\363\030\001\230\364\030\001\200\364\030\001\212\262\031\004Auto'))

_CLUSTERAPI = _descriptor.ServiceDescriptor(
  name='ClusterApi',
  full_name='edgeproto.ClusterApi',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=361,
  serialized_end=708,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateCluster',
    full_name='edgeproto.ClusterApi.CreateCluster',
    index=0,
    containing_service=None,
    input_type=_CLUSTER,
    output_type=result__pb2._RESULT,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\024\"\017/create/cluster:\001*')),
  ),
  _descriptor.MethodDescriptor(
    name='DeleteCluster',
    full_name='edgeproto.ClusterApi.DeleteCluster',
    index=1,
    containing_service=None,
    input_type=_CLUSTER,
    output_type=result__pb2._RESULT,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\024\"\017/delete/cluster:\001*')),
  ),
  _descriptor.MethodDescriptor(
    name='UpdateCluster',
    full_name='edgeproto.ClusterApi.UpdateCluster',
    index=2,
    containing_service=None,
    input_type=_CLUSTER,
    output_type=result__pb2._RESULT,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\024\"\017/update/cluster:\001*')),
  ),
  _descriptor.MethodDescriptor(
    name='ShowCluster',
    full_name='edgeproto.ClusterApi.ShowCluster',
    index=3,
    containing_service=None,
    input_type=_CLUSTER,
    output_type=_CLUSTER,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\022\"\r/show/cluster:\001*')),
  ),
])
_sym_db.RegisterServiceDescriptor(_CLUSTERAPI)

DESCRIPTOR.services_by_name['ClusterApi'] = _CLUSTERAPI

# @@protoc_insertion_point(module_scope)
