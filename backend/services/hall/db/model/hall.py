from sqlalchemy.orm import relationships
from typing import List
from typing import Optional
from sqlalchemy import ForeignKey
from sqlalchemy import String,Integer,select, DateTime
from sqlalchemy.orm import DeclarativeBase
from sqlalchemy.orm import Mapped
from sqlalchemy.orm import mapped_column
from datetime import datetime

class Base(DeclarativeBase):
    __abstract__ = True

class Hall(Base):
    __tablename__ = 'halls'
    id = mapped_column('id', Integer, primary_key=True,autoincrement=True)
    movie_id = mapped_column('movie_id', Integer,nullable=False)
    capacity = mapped_column('capacity', Integer,nullable=False)
    available = mapped_column('available', Integer,nullable=False)
    movie_name = mapped_column('movie_name', String,nullable=False)
    created_at = mapped_column('created_at', DateTime,nullable=False,default=datetime.now())
    updated_at = mapped_column('updated_at', DateTime,nullable=False,default=datetime.now(),onupdate=datetime.now())

    def __repr__(self) -> str:
        return f"Hall(id={self.id},movie_id={self.movie_id},capacity={self.capacity},available={self.available},movie_name={self.movie_name})"
    
    def create(self,session):
        session.add(self)
        session.commit()
        return self
    
    def update(self,session):
        session.commit()
        return self
    
    def delete(self,session):
        session.delete(self)
        session.commit()
        return self
    def find_by_id(self,session,id):
        stmt = select(Hall).where(Hall.id==id)
        hall = session.scalar(stmt)
        return hall
