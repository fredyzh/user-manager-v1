# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import meta_pb2 as meta__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nuser.proto\x12\nuser.v1.pb\x1a\nmeta.proto\"\x99\x01\n\x04User\x12\x1e\n\x04meta\x18\x01 \x01(\x0b\x32\x10.meta.v1.pb.Meta\x12)\n\nbasic_auth\x18\x02 \x01(\x0b\x32\x15.user.v1.pb.BasicAuth\x12$\n\x07profile\x18\x03 \x01(\x0b\x32\x13.user.v1.pb.Profile\x12 \n\x05scope\x18\x04 \x01(\x0b\x32\x11.meta.v1.pb.Scope\"/\n\tBasicAuth\x12\x10\n\x08username\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\"\xa6\x01\n\x07Profile\x12\x12\n\nfirst_name\x18\x01 \x01(\t\x12\x11\n\tlast_name\x18\x02 \x01(\t\x12\r\n\x05\x65mail\x18\x03 \x01(\t\x12\r\n\x05phone\x18\x04 \x01(\t\x12\x0c\n\x04\x63\x65ll\x18\x05 \x01(\t\x12\"\n\x06gender\x18\x06 \x01(\x0e\x32\x12.user.v1.pb.Gender\x12$\n\x07\x61\x64\x64ress\x18\x07 \x01(\x0b\x32\x13.user.v1.pb.Address\"C\n\x07\x41\x64\x64ress\x12\x0e\n\x06street\x18\x01 \x01(\t\x12\x0c\n\x04\x63ity\x18\x02 \x01(\t\x12\r\n\x05state\x18\x03 \x01(\t\x12\x0b\n\x03zip\x18\x04 \x01(\t*+\n\x06Gender\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x08\n\x04MALE\x10\x01\x12\n\n\x06\x46\x45MALE\x10\x02\x42\x30Z.github.com/fredyzh/user_manager/api/user/v1/pbb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'user_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z.github.com/fredyzh/user_manager/api/user/v1/pb'
  _globals['_GENDER']._serialized_start=481
  _globals['_GENDER']._serialized_end=524
  _globals['_USER']._serialized_start=39
  _globals['_USER']._serialized_end=192
  _globals['_BASICAUTH']._serialized_start=194
  _globals['_BASICAUTH']._serialized_end=241
  _globals['_PROFILE']._serialized_start=244
  _globals['_PROFILE']._serialized_end=410
  _globals['_ADDRESS']._serialized_start=412
  _globals['_ADDRESS']._serialized_end=479
# @@protoc_insertion_point(module_scope)