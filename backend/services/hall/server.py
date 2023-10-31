import logging
import grpc
import os

from concurrent import futures

from sqlalchemy import text
from proto import hallService_pb2
from proto import hallService_pb2_grpc
from proto import movieService_pb2
from db import db
from dotenv import load_dotenv
from controller import hall



def serve():
    load_dotenv()
    host = 'localhost'
    port = 7777
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

    hallService_pb2_grpc.add_HallServiceServicer_to_server(hall.getHallService(), server)
    server.add_insecure_port('{}:{}'.format(host, port))
    server.start()
    print("Server started at...{}".format(port))
    connection = db.get_connection()
    print("Connection established to db: {}".format(connection))
    server.wait_for_termination()


if __name__=="__main__":
    logging.basicConfig()
    serve()