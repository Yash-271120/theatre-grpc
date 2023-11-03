from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Movie(_message.Message):
    __slots__ = ["id", "title"]
    ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    id: int
    title: str
    def __init__(self, id: _Optional[int] = ..., title: _Optional[str] = ...) -> None: ...

class GetMoviesRequest(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class GetMoviesResponse(_message.Message):
    __slots__ = ["movies"]
    MOVIES_FIELD_NUMBER: _ClassVar[int]
    movies: _containers.RepeatedCompositeFieldContainer[Movie]
    def __init__(self, movies: _Optional[_Iterable[_Union[Movie, _Mapping]]] = ...) -> None: ...

class CreateMovieRequest(_message.Message):
    __slots__ = ["title"]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    title: str
    def __init__(self, title: _Optional[str] = ...) -> None: ...

class CreateMovieResponse(_message.Message):
    __slots__ = ["movie"]
    MOVIE_FIELD_NUMBER: _ClassVar[int]
    movie: Movie
    def __init__(self, movie: _Optional[_Union[Movie, _Mapping]] = ...) -> None: ...
