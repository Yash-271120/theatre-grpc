import movieService_pb2 as _movieService_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Hall(_message.Message):
    __slots__ = ["id", "movie", "capacity", "available"]
    ID_FIELD_NUMBER: _ClassVar[int]
    MOVIE_FIELD_NUMBER: _ClassVar[int]
    CAPACITY_FIELD_NUMBER: _ClassVar[int]
    AVAILABLE_FIELD_NUMBER: _ClassVar[int]
    id: int
    movie: _movieService_pb2.Movie
    capacity: int
    available: int
    def __init__(self, id: _Optional[int] = ..., movie: _Optional[_Union[_movieService_pb2.Movie, _Mapping]] = ..., capacity: _Optional[int] = ..., available: _Optional[int] = ...) -> None: ...

class CreateHallRequest(_message.Message):
    __slots__ = ["movie", "capacity", "available"]
    MOVIE_FIELD_NUMBER: _ClassVar[int]
    CAPACITY_FIELD_NUMBER: _ClassVar[int]
    AVAILABLE_FIELD_NUMBER: _ClassVar[int]
    movie: _movieService_pb2.Movie
    capacity: int
    available: int
    def __init__(self, movie: _Optional[_Union[_movieService_pb2.Movie, _Mapping]] = ..., capacity: _Optional[int] = ..., available: _Optional[int] = ...) -> None: ...

class CreateHallResponse(_message.Message):
    __slots__ = ["hall"]
    HALL_FIELD_NUMBER: _ClassVar[int]
    hall: Hall
    def __init__(self, hall: _Optional[_Union[Hall, _Mapping]] = ...) -> None: ...

class GetHallRequest(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class GetHallResponse(_message.Message):
    __slots__ = ["hall"]
    HALL_FIELD_NUMBER: _ClassVar[int]
    hall: Hall
    def __init__(self, hall: _Optional[_Union[Hall, _Mapping]] = ...) -> None: ...

class UpdateHallAvailableRequest(_message.Message):
    __slots__ = ["id", "available"]
    ID_FIELD_NUMBER: _ClassVar[int]
    AVAILABLE_FIELD_NUMBER: _ClassVar[int]
    id: int
    available: int
    def __init__(self, id: _Optional[int] = ..., available: _Optional[int] = ...) -> None: ...

class UpdateHallAvailableResponse(_message.Message):
    __slots__ = ["hall"]
    HALL_FIELD_NUMBER: _ClassVar[int]
    hall: Hall
    def __init__(self, hall: _Optional[_Union[Hall, _Mapping]] = ...) -> None: ...
