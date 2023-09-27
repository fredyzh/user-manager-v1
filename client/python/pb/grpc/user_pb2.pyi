import meta_pb2 as _meta_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Gender(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    UNKNOWN: _ClassVar[Gender]
    MALE: _ClassVar[Gender]
    FEMALE: _ClassVar[Gender]
UNKNOWN: Gender
MALE: Gender
FEMALE: Gender

class User(_message.Message):
    __slots__ = ["meta", "basic_auth", "profile", "scope"]
    META_FIELD_NUMBER: _ClassVar[int]
    BASIC_AUTH_FIELD_NUMBER: _ClassVar[int]
    PROFILE_FIELD_NUMBER: _ClassVar[int]
    SCOPE_FIELD_NUMBER: _ClassVar[int]
    meta: _meta_pb2.Meta
    basic_auth: BasicAuth
    profile: Profile
    scope: _meta_pb2.Scope
    def __init__(self, meta: _Optional[_Union[_meta_pb2.Meta, _Mapping]] = ..., basic_auth: _Optional[_Union[BasicAuth, _Mapping]] = ..., profile: _Optional[_Union[Profile, _Mapping]] = ..., scope: _Optional[_Union[_meta_pb2.Scope, _Mapping]] = ...) -> None: ...

class BasicAuth(_message.Message):
    __slots__ = ["username", "password"]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    username: str
    password: str
    def __init__(self, username: _Optional[str] = ..., password: _Optional[str] = ...) -> None: ...

class Profile(_message.Message):
    __slots__ = ["first_name", "last_name", "email", "phone", "cell", "gender", "address"]
    FIRST_NAME_FIELD_NUMBER: _ClassVar[int]
    LAST_NAME_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    PHONE_FIELD_NUMBER: _ClassVar[int]
    CELL_FIELD_NUMBER: _ClassVar[int]
    GENDER_FIELD_NUMBER: _ClassVar[int]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    first_name: str
    last_name: str
    email: str
    phone: str
    cell: str
    gender: Gender
    address: Address
    def __init__(self, first_name: _Optional[str] = ..., last_name: _Optional[str] = ..., email: _Optional[str] = ..., phone: _Optional[str] = ..., cell: _Optional[str] = ..., gender: _Optional[_Union[Gender, str]] = ..., address: _Optional[_Union[Address, _Mapping]] = ...) -> None: ...

class Address(_message.Message):
    __slots__ = ["street", "city", "state", "zip"]
    STREET_FIELD_NUMBER: _ClassVar[int]
    CITY_FIELD_NUMBER: _ClassVar[int]
    STATE_FIELD_NUMBER: _ClassVar[int]
    ZIP_FIELD_NUMBER: _ClassVar[int]
    street: str
    city: str
    state: str
    zip: str
    def __init__(self, street: _Optional[str] = ..., city: _Optional[str] = ..., state: _Optional[str] = ..., zip: _Optional[str] = ...) -> None: ...
