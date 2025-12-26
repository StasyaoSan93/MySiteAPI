from pydantic import BaseModel

# DATA BASE

class SiteDataBase(BaseModel):
    headertext: str
    bodytext: str
    imageurl: str
    siteurl: str

class SiteDataCreate(SiteDataBase):
    pass

class SiteDataResponse(SiteDataBase):
    id: int

    class Config:
        from_attributes = True

# INFO PAGE DATA BASE

class InfoDataBase(BaseModel):
    headertext: str
    bodytext: str

class InfoDataResponse(InfoDataBase):
    id: int

    class Config:
        from_attributes = True

# USER BASE

class UserCreate(BaseModel):
    username: str
    password: str

class UserResponse(BaseModel):
    id: int
    username: str

    class Config:
        from_attributes = True

class Token(BaseModel):
    access_token: str
    token_type: str