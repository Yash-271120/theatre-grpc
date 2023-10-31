from sqlalchemy import create_engine, text
from sqlalchemy.orm import Session
from .model import hall

from dotenv import load_dotenv
import os

load_dotenv()
DBNAME = os.getenv("DBNAME")
DBUSER = os.getenv("DBUSER")
DBPASS = os.getenv("DBPASS")
DBHOST = os.getenv("DBHOST")
DBPORT = os.getenv("DBPORT")

connection_string = ('postgresql://{}:{}@{}:{}/{}'.format(
                            DBUSER,
                            DBPASS,
                            DBHOST,
                            DBPORT,
                            DBNAME))

engine = create_engine(connection_string,echo=True)
hall.Base.metadata.create_all(engine)
connection = engine.connect()

def get_engine():
    return engine

def get_connection():
    return connection

def get_session():
    return Session(engine)

def close_connection():
    connection.close()
    engine.dispose()