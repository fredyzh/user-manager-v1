import annotations_pb2 as _annotations_pb2
import empty_pb2 as _empty_pb2
import user_pb2 as _user_pb2
import meta_pb2 as _meta_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetUserReq(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetUserReply(_message.Message):
    __slots__ = ["id", "user"]
    ID_FIELD_NUMBER: _ClassVar[int]
    USER_FIELD_NUMBER: _ClassVar[int]
    id: int
    user: _user_pb2.User
    def __init__(self, id: _Optional[int] = ..., user: _Optional[_Union[_user_pb2.User, _Mapping]] = ...) -> None: ...

class GetUserByUsernameReq(_message.Message):
    __slots__ = ["username"]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    username: str
    def __init__(self, username: _Optional[str] = ...) -> None: ...

class GetUserByUsernameReply(_message.Message):
    __slots__ = ["id", "user"]
    ID_FIELD_NUMBER: _ClassVar[int]
    USER_FIELD_NUMBER: _ClassVar[int]
    id: int
    user: _user_pb2.User
    def __init__(self, id: _Optional[int] = ..., user: _Optional[_Union[_user_pb2.User, _Mapping]] = ...) -> None: ...

class CreateUserReq(_message.Message):
    __slots__ = ["username", "password", "scope"]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    SCOPE_FIELD_NUMBER: _ClassVar[int]
    username: str
    password: str
    scope: _meta_pb2.Scope
    def __init__(self, username: _Optional[str] = ..., password: _Optional[str] = ..., scope: _Optional[_Union[_meta_pb2.Scope, _Mapping]] = ...) -> None: ...

class CreateUserReply(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class ListUserReq(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class ListUserReply(_message.Message):
    __slots__ = ["id", "results"]
    ID_FIELD_NUMBER: _ClassVar[int]
    RESULTS_FIELD_NUMBER: _ClassVar[int]
    id: int
    results: _containers.RepeatedCompositeFieldContainer[_user_pb2.User]
    def __init__(self, id: _Optional[int] = ..., results: _Optional[_Iterable[_Union[_user_pb2.User, _Mapping]]] = ...) -> None: ...

class VerifyPasswordReq(_message.Message):
    __slots__ = ["username", "password", "scope"]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    SCOPE_FIELD_NUMBER: _ClassVar[int]
    username: str
    password: str
    scope: _meta_pb2.Scope
    def __init__(self, username: _Optional[str] = ..., password: _Optional[str] = ..., scope: _Optional[_Union[_meta_pb2.Scope, _Mapping]] = ...) -> None: ...

class VerifyPasswordReply(_message.Message):
    __slots__ = ["id", "ok"]
    ID_FIELD_NUMBER: _ClassVar[int]
    OK_FIELD_NUMBER: _ClassVar[int]
    id: int
    ok: bool
    def __init__(self, id: _Optional[int] = ..., ok: bool = ...) -> None: ...

class CreateProfileReq(_message.Message):
    __slots__ = ["id", "profile"]
    ID_FIELD_NUMBER: _ClassVar[int]
    PROFILE_FIELD_NUMBER: _ClassVar[int]
    id: str
    profile: _user_pb2.Profile
    def __init__(self, id: _Optional[str] = ..., profile: _Optional[_Union[_user_pb2.Profile, _Mapping]] = ...) -> None: ...

class CreateProfileReply(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class GetProfileReq(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetProfileReply(_message.Message):
    __slots__ = ["id", "profile"]
    ID_FIELD_NUMBER: _ClassVar[int]
    PROFILE_FIELD_NUMBER: _ClassVar[int]
    id: int
    profile: _user_pb2.Profile
    def __init__(self, id: _Optional[int] = ..., profile: _Optional[_Union[_user_pb2.Profile, _Mapping]] = ...) -> None: ...

class CreateAddressReq(_message.Message):
    __slots__ = ["id", "address"]
    ID_FIELD_NUMBER: _ClassVar[int]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    id: str
    address: _user_pb2.Address
    def __init__(self, id: _Optional[str] = ..., address: _Optional[_Union[_user_pb2.Address, _Mapping]] = ...) -> None: ...

class CreateAddressReply(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class GetAddressReq(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetAddressReply(_message.Message):
    __slots__ = ["id", "address"]
    ID_FIELD_NUMBER: _ClassVar[int]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    id: int
    address: _user_pb2.Address
    def __init__(self, id: _Optional[int] = ..., address: _Optional[_Union[_user_pb2.Address, _Mapping]] = ...) -> None: ...

class ResetUserPasswordReq(_message.Message):
    __slots__ = ["id", "new_password", "old_password"]
    ID_FIELD_NUMBER: _ClassVar[int]
    NEW_PASSWORD_FIELD_NUMBER: _ClassVar[int]
    OLD_PASSWORD_FIELD_NUMBER: _ClassVar[int]
    id: str
    new_password: str
    old_password: str
    def __init__(self, id: _Optional[str] = ..., new_password: _Optional[str] = ..., old_password: _Optional[str] = ...) -> None: ...

class ResetUserPasswordReply(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class HealthReply(_message.Message):
    __slots__ = ["ok"]
    OK_FIELD_NUMBER: _ClassVar[int]
    ok: bool
    def __init__(self, ok: bool = ...) -> None: ...
