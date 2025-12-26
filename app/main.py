from fastapi import FastAPI, Depends, HTTPException, status
from sqlalchemy.orm import Session
from fastapi.security import OAuth2PasswordBearer, OAuth2PasswordRequestForm

from core import models
from core.database import SessionLocal, engine
from core.schemas import SiteDataCreate, SiteDataResponse, UserCreate, UserResponse, Token, InfoDataResponse
from core import auth
from fastapi.middleware.cors import CORSMiddleware

models.Base.metadata.create_all(bind=engine)

app = FastAPI(title="ISAProg Site API",  docs_url=None, redoc_url=None, openapi_url=None)

origins = [
    "https://isaprog.com",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

oauth2_scheme = OAuth2PasswordBearer(tokenUrl="token")

# --- Зависимость для сессии БД ---
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

# --- Регистрация пользователя ---
# @app.post("/register/", response_model=UserResponse)
def register(user: UserCreate, db: Session = Depends(get_db)):
    db_user = db.query(models.User).filter(models.User.username == user.username).first()
    if db_user:
        raise HTTPException(status_code=400, detail="Username already registered")
    hashed_password = auth.get_password_hash(user.password)
    new_user = models.User(username=user.username, hashed_password=hashed_password)
    db.add(new_user)
    db.commit()
    db.refresh(new_user)
    return new_user

# --- Логин и получение JWT ---
@app.post("/token", response_model=Token)
def login(form_data: OAuth2PasswordRequestForm = Depends(), db: Session = Depends(get_db)):
    user = db.query(models.User).filter(models.User.username == form_data.username).first()
    if not user or not auth.verify_password(form_data.password, user.hashed_password):
        raise HTTPException(status_code=401, detail="Incorrect username or password")
    access_token = auth.create_access_token(data={"sub": user.username})
    return {"access_token": access_token, "token_type": "bearer"}

# --- Получение текущего пользователя ---
def get_current_user(token: str = Depends(oauth2_scheme), db: Session = Depends(get_db)):
    username = auth.decode_access_token(token)
    if not username:
        raise HTTPException(status_code=status.HTTP_401_UNAUTHORIZED, detail="Invalid token")
    user = db.query(models.User).filter(models.User.username == username).first()
    if not user:
        raise HTTPException(status_code=status.HTTP_401_UNAUTHORIZED, detail="User not found")
    return user

# --- CRUD для sitedata ---
@app.post("/sitedata/", response_model=SiteDataResponse)
def create_item(item: SiteDataCreate, db: Session = Depends(get_db), current_user: models.User = Depends(get_current_user)):
    db_item = models.SiteData(**item.dict())
    db.add(db_item)
    db.commit()
    db.refresh(db_item)
    return db_item

@app.get("/sitedata/", response_model=list[SiteDataResponse])
def get_items(db: Session = Depends(get_db)):
    return db.query(models.SiteData).all()

@app.get("/sitedata/{item_id}", response_model=SiteDataResponse)
def get_item(item_id: int, db: Session = Depends(get_db)):
    item = db.query(models.SiteData).filter(models.SiteData.id == item_id).first()
    if not item:
        raise HTTPException(status_code=404, detail="Item not found")
    return item

@app.get("/infopagedata/", response_model=list[InfoDataResponse])
def get_items(db: Session = Depends(get_db)):
    return db.query(models.InfoPageData).all()

@app.put("/sitedata/{item_id}", response_model=SiteDataResponse)
def update_item(item_id: int, data: SiteDataCreate, db: Session = Depends(get_db), current_user: models.User = Depends(get_current_user)):
    item = db.query(models.SiteData).filter(models.SiteData.id == item_id).first()
    if not item:
        raise HTTPException(status_code=404, detail="Item not found")
    for key, value in data.dict().items():
        setattr(item, key, value)
    db.commit()
    db.refresh(item)
    return item

@app.delete("/sitedata/{item_id}")
def delete_item(item_id: int, db: Session = Depends(get_db), current_user: models.User = Depends(get_current_user)):
    item = db.query(models.SiteData).filter(models.SiteData.id == item_id).first()
    if not item:
        raise HTTPException(status_code=404, detail="Item not found")
    db.delete(item)
    db.commit()
    return {"status": "deleted"}
