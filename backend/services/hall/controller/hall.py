import time
import random
from proto import hallService_pb2, hallService_pb2_grpc, movieService_pb2
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
        fake_movie = movieService_pb2.Movie(id=1, title="fake_movie")
        fake_capacity = 100
        fake_available = random.randint(0, 100)
        hall = hallService_pb2.Hall(id=id, movie=fake_movie, capacity=fake_capacity, available=fake_available)
        return hallService_pb2.GetHallResponse(hall=hall)

    def UpdateHallAvailable(self, request, context):
        id = request.id
        available = request.available
        fake_movie = movieService_pb2.Movie(id=1, title="fake_movie")
        fake_capacity = 100
        hall = hallService_pb2.Hall(id=id, movie=fake_movie, capacity=fake_capacity, available=available)
        return hallService_pb2.UpdateHallAvailableResponse(hall=hall)

def getHallService():
    return HallService()