package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"

	"github.com/webProj/Admin/repository"
	"github.com/webProj/Admin/service"
	"github.com/webProj/delivery/http/handler"
)

var tmpl = template.Must(template.ParseGlob("../ui/template/*.html"))

func main() {

	dbconn, err := gorm.Open("postgres", "postgres://postgres:P@$$w0rDd@localhost/hospital?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	//Login handlers
	// loginRepo := repository.NewLoginGormRepo(dbconn)
	// loginSrv := service.NewLoginService(loginRepo)
	// allLoginHand := handler.NewAllLoginHandler(loginSrv)
	// router := httprouter.New()
	// fmt.Println("To call the /login method")
	// router.POST("/login", allLoginHand.UserLogin)
	// fmt.Println("After calling the /login method")

	manageDocRepo := repository.NewManageDoctorsRepository(dbconn)
	manageDocServ := service.NewManageDoctorsService(manageDocRepo)
	managDocHand := handler.NewManageDoctorsHandler(manageDocServ)
	router := httprouter.New()
	router.GET("/admin/doctors", managDocHand.GetDoctors)
	router.GET("/admin/doctor/?id", managDocHand.GetSingleDoctor)
	router.PUT("/admin/doctor/?id", managDocHand.UpdateDoctor)
	router.POST("/admin/doctors", managDocHand.AddDoctor)
	router.DELETE("/admin/doctors/?id", managDocHand.DeleteDoctor)

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("../ui/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Println("Before registerring the / method")
	mux.HandleFunc("/", index)
	fmt.Println("After registering the / method")

	http.ListenAndServe(":8012", router)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Here in index")
	tmpl.ExecuteTemplate(w, "login.html", nil)
	fmt.Println("Here in index, after display")
}

// type LoginGormRepo struct {
// 	conn *gorm.DB
// }

// // NewCategoryGormRepo will create a new object of CategoryGormRepo
// func NewLoginGormRepo(db *gorm.DB) LoginGormRepo {
// 	return LoginGormRepo{conn: db}
// }

// func login(w http.ResponseWriter, r *http.Request) {
// 	username := r.FormValue("username")
// 	password := r.FormValue("password")
// 	fmt.Println(username)
// 	fmt.Println(password)

// }
// func UserChecker(db gorm.DB, uname string, pass string) (*entity.Profile, []error) {
// 	ctg := entity.Profile{}
// 	errs := db.First(&ctg, uname, pass).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return &ctg, errs
// }
