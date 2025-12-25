from sqlalchemy import Column, Integer, Text
from .database import Base

class SiteData(Base):
    __tablename__ = "sitedata"

    id = Column(Integer, primary_key=True, index=True)
    headertext = Column(Text)
    bodytext = Column(Text)
    imageurl = Column(Text)
    siteurl = Column(Text)

class User(Base):
    __tablename__ = "users"

    id = Column(Integer, primary_key=True, index=True)
    username = Column(Text, unique=True, index=True, nullable=False)
    hashed_password = Column(Text, nullable=False)