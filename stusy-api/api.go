package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        int
	Username  string
	Password  []byte
	Token     string
	ExpiresAt int64
}

type UserData struct {
	ID         int
	UserID     int
	Email      string
	FirstName  string
	MiddleName string
	LastName   string
}

type Group struct {
	ID   int    `json:"gid"`
	Name string `json:"name"`
	Year int    `json:"year"`
}

type GroupData struct {
	ID  int
	UID int
	GID int
}

type Server struct {
	Config
	Router *mux.Router
	DB     *gorm.DB
}

type Config struct {
	Port   string
	DSN    string
	Secret string
}

const (
	ErrorInternal       = "Internal server error. Try again later."
	ErrorBadJSON        = "You have supplied an invalid JSON body."
	ErrorBadCredentials = "Invalid credentials. Invalid email or password."
	ErrorUserExist      = "This username is taken."
	ErrorGroupExist     = "Group with that name already exists"
	ErrorEmailFormat    = "Invalid email format."
	ErrorShortPass      = "Password is too short (should be >= 8)."
	ErrorBadToken       = "Token is invalid or expired."
	ErrorAccess         = "You don't have permissions to make this request."
	ErrorNameFormat     = "Name should only consist of alphabetical characters."
	ErrorNotFound       = "Resource not found."
)

func (a *Server) Init() {
	var err error
	log.Println("Initializing config struct..")
	err = a.Config.Init()
	if err != nil {
		log.Fatal(err)
	}
	for {
		a.DB, err = gorm.Open(mysql.Open(a.Config.DSN), &gorm.Config{})
		if err != nil {
			log.Println("Failed to establish database connection, retrying...")
			time.Sleep(time.Second * 3)
		} else {
			log.Println("Database connection established.")
			break
		}
	}
	err = a.DB.AutoMigrate(&User{}, &UserData{}, &Group{}, &GroupData{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Initializing routes..")
	a.Router = mux.NewRouter()
	a.initRoutes()
}

func (c *Config) Init() error {
	c.DSN = fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	c.Port = os.Getenv("PORT")
	c.Secret = os.Getenv("SECRET")
	if len(c.Port) == 0 || len(c.Secret) == 0 {
		return fmt.Errorf("some of the env vars are blank.")
	}
	return nil
}

func (a *Server) initRoutes() {
	a.Router.HandleFunc("/{.+}", a.preflightCORS).Methods("OPTIONS")
	a.Router.HandleFunc("/{.+}/{.+}", a.preflightCORS).Methods("OPTIONS")
	a.Router.HandleFunc("/", a.listRoutes).Methods("GET")
	a.Router.HandleFunc("/user", a.createUser).Methods("POST")
	a.Router.HandleFunc("/user/batch", a.batchUserCreation).Methods("POST")
	a.Router.HandleFunc("/user/auth", a.loginUser).Methods("POST")
	a.Router.HandleFunc("/user/{id:[0-9]+}", a.authCheck(a.permsCheck(a.getUser))).Methods("GET")
	a.Router.HandleFunc("/user/{id:[0-9]+}", a.authCheck(a.permsCheck(a.updateUser))).Methods("PUT")
	a.Router.HandleFunc("/groups", a.fetchGroups).Methods("GET")
	a.Router.HandleFunc("/group", a.createGroup).Methods("POST")
	a.Router.HandleFunc("/group/{id:[0-9]+}", a.getGroup).Methods("GET")
}

func (a *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func main() {
	var a Server
	a.Init()
	log.Println("Starting local server..")
	a.Run(":" + a.Config.Port)
}

func RespondWithJSON(w *http.ResponseWriter, status int, pl interface{}) {
	res, _ := json.MarshalIndent(pl, "", "	")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).WriteHeader(status)
	(*w).Write(res)
}

func RespondWithError(w *http.ResponseWriter, status int, err string) {
	RespondWithJSON(w, status, map[string]string{"error": err})
}

func (a *Server) authCheck(n http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		header := strings.Split(req.Header.Get("Authorization"), "Bearer ")
		if len(header) != 2 || !a.validateToken(header[1]) {
			RespondWithError(&w, http.StatusBadRequest, ErrorBadToken)
			return
		}
		n.ServeHTTP(w, req)
	}
}

func (a *Server) generateToken(iss string) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 12).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    iss,
		ExpiresAt: exp,
	})
	tok, err := claims.SignedString([]byte(a.Config.Secret))
	if err != nil {
		return "", 0, err
	}
	return tok, exp, nil
}

func (a *Server) validateToken(s string) bool {
	tok, err := a.parseToken(s)
	if err != nil {
		return false
	}
	if _, ok := tok.Claims.(jwt.MapClaims); ok && tok.Valid {
		return true
	}
	return false
}

func (a *Server) parseToken(s string) (*jwt.Token, error) {
	tok, err := jwt.Parse(s, func(tok *jwt.Token) (interface{}, error) {
		if _, ok := tok.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(a.Config.Secret), nil
	})
	return tok, err
}

func (a *Server) preflightCORS(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}

func (a *Server) listRoutes(w http.ResponseWriter, _ *http.Request) {
	var (
		path    string
		methods []string
		err     error
	)
	list := make(map[string]string)
	err = a.Router.Walk(func(r *mux.Route, _ *mux.Router, _ []*mux.Route) error {
		path, err = r.GetPathTemplate()
		if err != nil {
			return err
		}
		methods, err = r.GetMethods()
		if err != nil {
			return err
		}
		list[path] = strings.Join(methods, ",")
		return nil
	})
	if err != nil {
		RespondWithError(&w, http.StatusInternalServerError, ErrorInternal)
		log.Println(err)
		return
	}
	RespondWithJSON(&w, http.StatusOK, list)
}

func (a *Server) createUser(w http.ResponseWriter, req *http.Request) {
	var (
		pl struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		u User
	)
	if err := json.NewDecoder(req.Body).Decode(&pl); err != nil {
		RespondWithError(&w, http.StatusBadRequest, ErrorBadJSON)
		return
	}
	defer req.Body.Close()
	a.DB.Where("username = ?", pl.Username).First(&u)
	if u.ID != 0 {
		RespondWithError(&w, http.StatusForbidden, ErrorUserExist)
		return
	}
	if !validateRegData(pl.Password, pl.Username) {
		RespondWithError(&w, http.StatusBadRequest, ErrorBadJSON)
		return
	}
	pass, _ := bcrypt.GenerateFromPassword([]byte(pl.Password), 12)
	u.Username = strings.ToLower(pl.Username)
	u.Password = pass
	a.DB.Create(&u)
	RespondWithJSON(&w, http.StatusCreated, map[string]int{"uid": u.ID})
}

func validateRegData(pass, name string) bool {
	// TODO: Make sure name is valid (i.e. doesn't contain spaces, symbols)
	if len(pass) < 8 {
		return false
	}
	return true
}

func (a *Server) loginUser(w http.ResponseWriter, req *http.Request) {
	var (
		pl struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		u User
	)
	dec := json.NewDecoder(req.Body)
	if err := dec.Decode(&pl); err != nil {
		RespondWithError(&w, http.StatusBadRequest, ErrorBadJSON)
		return
	}
	defer req.Body.Close()
	a.DB.Where("username = ?", strings.ToLower(pl.Username)).First(&u)
	if u.ID == 0 || !validatePassHash(u.Password, []byte(pl.Password)) {
		RespondWithError(&w, http.StatusForbidden, ErrorBadCredentials)
		return
	}
	if !a.validateToken(u.Token) {
		var err error
		u.Token, u.ExpiresAt, err = a.generateToken(strconv.Itoa(u.ID))
		if err != nil {
			log.Println(err)
			RespondWithError(&w, http.StatusInternalServerError, ErrorInternal)
			return
		}
		a.DB.Save(&u)
	}
	RespondWithJSON(&w, http.StatusOK, map[string]interface{}{
		"uid": u.ID, "access_token": u.Token})
}

func validatePassHash(a []byte, b []byte) bool {
	err := bcrypt.CompareHashAndPassword(a, b)
	if err != nil {
		return false
	}
	return true
}

func (a *Server) fetchUsers(w http.ResponseWriter, req *http.Request) {
}

func (a *Server) getUser(w http.ResponseWriter, req *http.Request) {
	var pl UserData
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	a.DB.Where("user_id = ?", id).First(&pl)
	if pl.UserID == 0 {
		RespondWithError(&w, http.StatusNotFound, ErrorNotFound)
		return
	}
	RespondWithJSON(&w, http.StatusOK, &pl)
}

func (a *Server) updateUser(w http.ResponseWriter, req *http.Request) {
	var pl, tmp UserData
	if err := json.NewDecoder(req.Body).Decode(&pl); err != nil {
		RespondWithError(&w, http.StatusBadRequest, ErrorBadJSON)
		return
	}
	defer req.Body.Close()
	re := regexp.MustCompile(`^\p{L}+$`).MatchString
	if !re(pl.FirstName) || !re(pl.LastName) || (len(pl.MiddleName) > 0 && !re(pl.MiddleName)) {
		RespondWithError(&w, http.StatusBadRequest, ErrorNameFormat)
		return
	}
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	a.DB.Where("user_id = ?", id).First(&tmp)
	if tmp.ID != 0 {
		RespondWithError(&w, http.StatusForbidden, "You can't change your name")
		return
	}
	pl.UserID = id
	a.DB.Create(&pl)
	RespondWithJSON(&w, http.StatusCreated, &pl)
}

func (a *Server) permsCheck(n http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		tok := strings.Split(req.Header.Get("Authorization"), "Bearer ")
		id, _ := strconv.Atoi(mux.Vars(req)["id"])
		var u User
		a.DB.Where("token = ?", tok[1]).First(&u)
		if u.ID != id {
			RespondWithError(&w, http.StatusForbidden, ErrorAccess)
			return
		}
		n.ServeHTTP(w, req)
	}
}

func (a *Server) fetchGroups(w http.ResponseWriter, req *http.Request) {
}

func (a *Server) getGroup(w http.ResponseWriter, req *http.Request) {
	var pl Group
	if err := json.NewDecoder(req.Body).Decode(&pl); err != nil {
		RespondWithError(&w, http.StatusBadRequest, ErrorBadJSON)
		return
	}
	a.DB.Where(&Group{ID: pl.ID}).First(&pl)
	if pl.ID == 0 {
		RespondWithError(&w, http.StatusBadRequest, "Group with corresponding ID doesn't exist")
		return
	}
	RespondWithJSON(&w, http.StatusOK, pl)
}

func (a *Server) createGroup(w http.ResponseWriter, req *http.Request) {
	var pl Group
	if err := json.NewDecoder(req.Body).Decode(&pl); err != nil {
		RespondWithError(&w, http.StatusBadRequest, ErrorBadJSON)
		return
	}
	defer req.Body.Close()
	a.DB.Where("name = ?", pl.Name).First(&pl)
	if pl.ID != 0 {
		RespondWithError(&w, http.StatusForbidden, ErrorGroupExist)
		return
	}
	// TODO: MAYBE validate a group name and year?
	a.DB.Create(&pl)
	RespondWithJSON(&w, http.StatusCreated, map[string]int{"gid": pl.ID})
}

func (a *Server) batchUserCreation(w http.ResponseWriter, req *http.Request) {
	var (
		pl struct {
			GID   int `json:"gid"`
			Count int `json:"count"`
		}
		g  Group
		i  int
	)
	if err := json.NewDecoder(req.Body).Decode(&pl); err != nil {
		RespondWithError(&w, http.StatusBadRequest, ErrorBadJSON)
		return
	}
	defer req.Body.Close()
	a.DB.Where(&Group{ID: pl.GID}).First(&g)
	if g.ID == 0 {
		RespondWithError(&w, http.StatusBadRequest, "Group with corresponding ID doesn't exist")
		return
	}
	users := make([]struct {
		Username string
		Password string
	}, pl.Count)
	for i = 0; i < pl.Count; i++ {
		users[i].Username = g.Name + "_" + strconv.Itoa(i)
		users[i].Password = "testtest" // TODO: add pass generator
		pass, _ := bcrypt.GenerateFromPassword([]byte(users[i].Password), 12)
		u := User{
			Username: users[i].Username,
			Password: pass,
		}
		a.DB.Create(&u)
		gd := GroupData{
			GID: g.ID,
			UID: u.ID,
		}
		a.DB.Create(&gd)
	}
	RespondWithJSON(&w, http.StatusCreated, users)
}
