from google.protobuf import Timestamp_pb2 as _Timestamp_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Meta(_message.Message):
    __slots__ = ["id", "create_at", "update_at", "update_by"]
    ID_FIELD_NUMBER: _ClassVar[int]
    CREATE_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATE_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATE_BY_FIELD_NUMBER: _ClassVar[int]
    id: str
    create_at: _Timestamp_pb2.Timestamp
    update_at: _Timestamp_pb2.Timestamp
    update_by: str
    def __init__(self, id: _Optional[str] = ..., create_at: _Optional[_Union[_Timestamp_pb2.Timestamp, _Mapping]] = ..., update_at: _Optional[_Union[_Timestamp_pb2.Timestamp, _Mapping]] = ..., update_by: _Optional[str] = ...) -> None: ...

class Scope(_message.Message):
    __slots__ = ["domain", "app_id"]
    DOMAIN_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    domain: str
    app_id: str
    def __init__(self, domain: _Optional[str] = ..., app_id: _Optional[str] = ...) -> None: ...
