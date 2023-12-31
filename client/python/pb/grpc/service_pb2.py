# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: service.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import annotations_pb2 as annotations__pb2
import empty_pb2 as empty__pb2
import user_pb2 as user__pb2
import meta_pb2 as meta__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rservice.proto\x12\x0fuser.v1.service\x1a\x11\x61nnotations.proto\x1a\x0b\x65mpty.proto\x1a\nuser.proto\x1a\nmeta.proto\"\x18\n\nGetUserReq\x12\n\n\x02id\x18\x01 \x01(\t\":\n\x0cGetUserReply\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x1e\n\x04user\x18\x02 \x01(\x0b\x32\x10.user.v1.pb.User\"(\n\x14GetUserByUsernameReq\x12\x10\n\x08username\x18\x01 \x01(\t\"D\n\x16GetUserByUsernameReply\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x1e\n\x04user\x18\x02 \x01(\x0b\x32\x10.user.v1.pb.User\"U\n\rCreateUserReq\x12\x10\n\x08username\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\x12 \n\x05scope\x18\x03 \x01(\x0b\x32\x11.meta.v1.pb.Scope\"\x1d\n\x0f\x43reateUserReply\x12\n\n\x02id\x18\x01 \x01(\x03\"\r\n\x0bListUserReq\">\n\rListUserReply\x12\n\n\x02id\x18\x01 \x01(\x03\x12!\n\x07results\x18\x02 \x03(\x0b\x32\x10.user.v1.pb.User\"Y\n\x11VerifyPasswordReq\x12\x10\n\x08username\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\x12 \n\x05scope\x18\x03 \x01(\x0b\x32\x11.meta.v1.pb.Scope\"-\n\x13VerifyPasswordReply\x12\n\n\x02id\x18\x01 \x01(\x03\x12\n\n\x02ok\x18\x02 \x01(\x08\"D\n\x10\x43reateProfileReq\x12\n\n\x02id\x18\x01 \x01(\t\x12$\n\x07profile\x18\x02 \x01(\x0b\x32\x13.user.v1.pb.Profile\" \n\x12\x43reateProfileReply\x12\n\n\x02id\x18\x01 \x01(\x03\"\x1b\n\rGetProfileReq\x12\n\n\x02id\x18\x01 \x01(\t\"C\n\x0fGetProfileReply\x12\n\n\x02id\x18\x01 \x01(\x03\x12$\n\x07profile\x18\x02 \x01(\x0b\x32\x13.user.v1.pb.Profile\"D\n\x10\x43reateAddressReq\x12\n\n\x02id\x18\x01 \x01(\t\x12$\n\x07\x61\x64\x64ress\x18\x02 \x01(\x0b\x32\x13.user.v1.pb.Address\" \n\x12\x43reateAddressReply\x12\n\n\x02id\x18\x01 \x01(\x03\"\x1b\n\rGetAddressReq\x12\n\n\x02id\x18\x01 \x01(\t\"C\n\x0fGetAddressReply\x12\n\n\x02id\x18\x01 \x01(\x03\x12$\n\x07\x61\x64\x64ress\x18\x02 \x01(\x0b\x32\x13.user.v1.pb.Address\"N\n\x14ResetUserPasswordReq\x12\n\n\x02id\x18\x01 \x01(\t\x12\x14\n\x0cnew_password\x18\x02 \x01(\t\x12\x14\n\x0cold_password\x18\x03 \x01(\t\"$\n\x16ResetUserPasswordReply\x12\n\n\x02id\x18\x01 \x01(\x03\"\x19\n\x0bHealthReply\x12\n\n\x02ok\x18\x01 \x01(\x08\x32\x81\x0b\n\x04User\x12_\n\x06Health\x12\x16.google.protobuf.Empty\x1a\x1c.user.v1.service.HealthReply\"\x1f\x82\xd3\xe4\x93\x02\x19\x12\x17/user/v1/service/health\x12l\n\x07GetUser\x12\x1b.user.v1.service.GetUserReq\x1a\x1d.user.v1.service.GetUserReply\"%\x82\xd3\xe4\x93\x02\x1f\x12\x1d/user/v1/service/getuser/{id}\x12\x96\x01\n\x11GetUserByUsername\x12%.user.v1.service.GetUserByUsernameReq\x1a\'.user.v1.service.GetUserByUsernameReply\"1\x82\xd3\xe4\x93\x02+\x12)/user/v1/service/getuserbyname/{username}\x12\x92\x01\n\x11ResetUserPassword\x12%.user.v1.service.ResetUserPasswordReq\x1a\'.user.v1.service.ResetUserPasswordReply\"-\x82\xd3\xe4\x93\x02\'\"\"/user/v1/service/resetuserpassword:\x01*\x12v\n\nCreateUser\x12\x1e.user.v1.service.CreateUserReq\x1a .user.v1.service.CreateUserReply\"&\x82\xd3\xe4\x93\x02 \"\x1b/user/v1/service/createuser:\x01*\x12k\n\x08ListUser\x12\x1c.user.v1.service.ListUserReq\x1a\x1e.user.v1.service.ListUserReply\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/user/v1/service/listuser\x12\x86\x01\n\x0eVerifyPassword\x12\".user.v1.service.VerifyPasswordReq\x1a$.user.v1.service.VerifyPasswordReply\"*\x82\xd3\xe4\x93\x02$\"\x1f/user/v1/service/verifypassword:\x01*\x12\x82\x01\n\rCreateAddress\x12!.user.v1.service.CreateAddressReq\x1a#.user.v1.service.CreateAddressReply\")\x82\xd3\xe4\x93\x02#\"\x1e/user/v1/service/createaddress:\x01*\x12\x80\x01\n\x0eGetUserAddress\x12\x1e.user.v1.service.GetAddressReq\x1a .user.v1.service.GetAddressReply\",\x82\xd3\xe4\x93\x02&\x12$/user/v1/service/getuseraddress/{id}\x12\x82\x01\n\rCreateProfile\x12!.user.v1.service.CreateProfileReq\x1a#.user.v1.service.CreateProfileReply\")\x82\xd3\xe4\x93\x02#\"\x1e/user/v1/service/createprofile:\x01*\x12\x80\x01\n\x0eGetUserProfile\x12\x1e.user.v1.service.GetProfileReq\x1a .user.v1.service.GetProfileReply\",\x82\xd3\xe4\x93\x02&\x12$/user/v1/service/getuserprofile/{id}B5Z3github.com/fredyzh/user_manager/api/user/v1/serviceb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z3github.com/fredyzh/user_manager/api/user/v1/service'
  _USER.methods_by_name['Health']._options = None
  _USER.methods_by_name['Health']._serialized_options = b'\202\323\344\223\002\031\022\027/user/v1/service/health'
  _USER.methods_by_name['GetUser']._options = None
  _USER.methods_by_name['GetUser']._serialized_options = b'\202\323\344\223\002\037\022\035/user/v1/service/getuser/{id}'
  _USER.methods_by_name['GetUserByUsername']._options = None
  _USER.methods_by_name['GetUserByUsername']._serialized_options = b'\202\323\344\223\002+\022)/user/v1/service/getuserbyname/{username}'
  _USER.methods_by_name['ResetUserPassword']._options = None
  _USER.methods_by_name['ResetUserPassword']._serialized_options = b'\202\323\344\223\002\'\"\"/user/v1/service/resetuserpassword:\001*'
  _USER.methods_by_name['CreateUser']._options = None
  _USER.methods_by_name['CreateUser']._serialized_options = b'\202\323\344\223\002 \"\033/user/v1/service/createuser:\001*'
  _USER.methods_by_name['ListUser']._options = None
  _USER.methods_by_name['ListUser']._serialized_options = b'\202\323\344\223\002\033\022\031/user/v1/service/listuser'
  _USER.methods_by_name['VerifyPassword']._options = None
  _USER.methods_by_name['VerifyPassword']._serialized_options = b'\202\323\344\223\002$\"\037/user/v1/service/verifypassword:\001*'
  _USER.methods_by_name['CreateAddress']._options = None
  _USER.methods_by_name['CreateAddress']._serialized_options = b'\202\323\344\223\002#\"\036/user/v1/service/createaddress:\001*'
  _USER.methods_by_name['GetUserAddress']._options = None
  _USER.methods_by_name['GetUserAddress']._serialized_options = b'\202\323\344\223\002&\022$/user/v1/service/getuseraddress/{id}'
  _USER.methods_by_name['CreateProfile']._options = None
  _USER.methods_by_name['CreateProfile']._serialized_options = b'\202\323\344\223\002#\"\036/user/v1/service/createprofile:\001*'
  _USER.methods_by_name['GetUserProfile']._options = None
  _USER.methods_by_name['GetUserProfile']._serialized_options = b'\202\323\344\223\002&\022$/user/v1/service/getuserprofile/{id}'
  _globals['_GETUSERREQ']._serialized_start=90
  _globals['_GETUSERREQ']._serialized_end=114
  _globals['_GETUSERREPLY']._serialized_start=116
  _globals['_GETUSERREPLY']._serialized_end=174
  _globals['_GETUSERBYUSERNAMEREQ']._serialized_start=176
  _globals['_GETUSERBYUSERNAMEREQ']._serialized_end=216
  _globals['_GETUSERBYUSERNAMEREPLY']._serialized_start=218
  _globals['_GETUSERBYUSERNAMEREPLY']._serialized_end=286
  _globals['_CREATEUSERREQ']._serialized_start=288
  _globals['_CREATEUSERREQ']._serialized_end=373
  _globals['_CREATEUSERREPLY']._serialized_start=375
  _globals['_CREATEUSERREPLY']._serialized_end=404
  _globals['_LISTUSERREQ']._serialized_start=406
  _globals['_LISTUSERREQ']._serialized_end=419
  _globals['_LISTUSERREPLY']._serialized_start=421
  _globals['_LISTUSERREPLY']._serialized_end=483
  _globals['_VERIFYPASSWORDREQ']._serialized_start=485
  _globals['_VERIFYPASSWORDREQ']._serialized_end=574
  _globals['_VERIFYPASSWORDREPLY']._serialized_start=576
  _globals['_VERIFYPASSWORDREPLY']._serialized_end=621
  _globals['_CREATEPROFILEREQ']._serialized_start=623
  _globals['_CREATEPROFILEREQ']._serialized_end=691
  _globals['_CREATEPROFILEREPLY']._serialized_start=693
  _globals['_CREATEPROFILEREPLY']._serialized_end=725
  _globals['_GETPROFILEREQ']._serialized_start=727
  _globals['_GETPROFILEREQ']._serialized_end=754
  _globals['_GETPROFILEREPLY']._serialized_start=756
  _globals['_GETPROFILEREPLY']._serialized_end=823
  _globals['_CREATEADDRESSREQ']._serialized_start=825
  _globals['_CREATEADDRESSREQ']._serialized_end=893
  _globals['_CREATEADDRESSREPLY']._serialized_start=895
  _globals['_CREATEADDRESSREPLY']._serialized_end=927
  _globals['_GETADDRESSREQ']._serialized_start=929
  _globals['_GETADDRESSREQ']._serialized_end=956
  _globals['_GETADDRESSREPLY']._serialized_start=958
  _globals['_GETADDRESSREPLY']._serialized_end=1025
  _globals['_RESETUSERPASSWORDREQ']._serialized_start=1027
  _globals['_RESETUSERPASSWORDREQ']._serialized_end=1105
  _globals['_RESETUSERPASSWORDREPLY']._serialized_start=1107
  _globals['_RESETUSERPASSWORDREPLY']._serialized_end=1143
  _globals['_HEALTHREPLY']._serialized_start=1145
  _globals['_HEALTHREPLY']._serialized_end=1170
  _globals['_USER']._serialized_start=1173
  _globals['_USER']._serialized_end=2582
# @@protoc_insertion_point(module_scope)
