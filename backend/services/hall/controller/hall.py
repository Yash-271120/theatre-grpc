import time
import random
from proto import hallService_pb2, hallService_pb2_grpc, movieService_pb2
import grpc
from db.model import hall
from db import db


class HallService(hallService_pb2_grpc.HallServiceServicer):
    def CreateHall(self, request, context):
        capacity = request.capacity
        available = request.available
        movie = request.movie
        new_hall = hall.Hall(movie_id=movie.id, capacity=capacity, available=available, movie_name=movie.title) 
        db_session = db.get_session()
        new_hall = new_hall.create(db_session)
        hall_pb = hallService_pb2.Hall(id=new_hall.id, movie=movie, capacity=capacity, available=available)
        return hallService_pb2.CreateHallResponse(hall=hall_pb)

    def GetHall(self, request, context):
        id = request.id
        db_session = db.get_session()
        hallObj = hall.Hall().find_by_id(db_session, id)
        if(hallObj == None):
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("hall not found")
            return hallService_pb2.GetHallResponse()
        res_movie = movieService_pb2.Movie(id=hallObj.movie_id, title=hallObj.movie_name)
        res_capacity = hallObj.capacity
        res_available = hallObj.available
        hall_pb = hallService_pb2.Hall(id=id, movie=res_movie, capacity=res_capacity, available=res_available)
        return hallService_pb2.GetHallResponse(hall=hall_pb)

    def UpdateHallAvailable(self, request, context):
        id = request.id
        db_session = db.get_session()
        hallObj = hall.Hall().find_by_id(db_session, id)
        if(hallObj == None):
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("hall not found")
            return hallService_pb2.GetHallResponse()
        res_available = request.available
        if(hallObj.available-res_available < 0):
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details("available seats not enough")
            return hallService_pb2.UpdateHallAvailableResponse()
        hallObj.available = hallObj.available - res_available
        hallObj.update(session=db_session)
        res_movie = movieService_pb2.Movie(id=hallObj.movie_id, title=hallObj.movie_name)
        res_capacity = hallObj.capacity
        hall_pb = hallService_pb2.Hall(id=id, movie=res_movie, capacity=res_capacity, available=res_available)
        return hallService_pb2.UpdateHallAvailableResponse(hall=hall_pb)

def getHallService():
    return HallService()