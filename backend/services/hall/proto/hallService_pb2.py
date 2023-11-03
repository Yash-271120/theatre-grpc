# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: hallService.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import movieService_pb2 as movieService__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x11hallService.proto\x12\x05proto\x1a\x12movieService.proto\"T\n\x04Hall\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x1b\n\x05movie\x18\x02 \x01(\x0b\x32\x0c.proto.Movie\x12\x10\n\x08\x63\x61pacity\x18\x03 \x01(\x05\x12\x11\n\tavailable\x18\x04 \x01(\x05\"U\n\x11\x43reateHallRequest\x12\x1b\n\x05movie\x18\x01 \x01(\x0b\x32\x0c.proto.Movie\x12\x10\n\x08\x63\x61pacity\x18\x02 \x01(\x05\x12\x11\n\tavailable\x18\x03 \x01(\x05\"/\n\x12\x43reateHallResponse\x12\x19\n\x04hall\x18\x01 \x01(\x0b\x32\x0b.proto.Hall\"\x1c\n\x0eGetHallRequest\x12\n\n\x02id\x18\x01 \x01(\x05\",\n\x0fGetHallResponse\x12\x19\n\x04hall\x18\x01 \x01(\x0b\x32\x0b.proto.Hall\";\n\x1aUpdateHallAvailableRequest\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x11\n\tavailable\x18\x02 \x01(\x05\"8\n\x1bUpdateHallAvailableResponse\x12\x19\n\x04hall\x18\x01 \x01(\x0b\x32\x0b.proto.Hall2\xe8\x01\n\x0bHallService\x12\x41\n\nCreateHall\x12\x18.proto.CreateHallRequest\x1a\x19.proto.CreateHallResponse\x12\x38\n\x07GetHall\x12\x15.proto.GetHallRequest\x1a\x16.proto.GetHallResponse\x12\\\n\x13UpdateHallAvailable\x12!.proto.UpdateHallAvailableRequest\x1a\".proto.UpdateHallAvailableResponseB\tZ\x07./protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'hallService_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\007./proto'
  _globals['_HALL']._serialized_start=48
  _globals['_HALL']._serialized_end=132
  _globals['_CREATEHALLREQUEST']._serialized_start=134
  _globals['_CREATEHALLREQUEST']._serialized_end=219
  _globals['_CREATEHALLRESPONSE']._serialized_start=221
  _globals['_CREATEHALLRESPONSE']._serialized_end=268
  _globals['_GETHALLREQUEST']._serialized_start=270
  _globals['_GETHALLREQUEST']._serialized_end=298
  _globals['_GETHALLRESPONSE']._serialized_start=300
  _globals['_GETHALLRESPONSE']._serialized_end=344
  _globals['_UPDATEHALLAVAILABLEREQUEST']._serialized_start=346
  _globals['_UPDATEHALLAVAILABLEREQUEST']._serialized_end=405
  _globals['_UPDATEHALLAVAILABLERESPONSE']._serialized_start=407
  _globals['_UPDATEHALLAVAILABLERESPONSE']._serialized_end=463
  _globals['_HALLSERVICE']._serialized_start=466
  _globals['_HALLSERVICE']._serialized_end=698
# @@protoc_insertion_point(module_scope)
