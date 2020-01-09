package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"

	"github.com/webProj/Login/lrepository"
	"github.com/webProj/Login/lservice"
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
	loginRepo := lrepository.NewLoginGormRepo(dbconn)
	loginSrv := lservice.NewLoginService(loginRepo)
	allLoginHand := handler.NewAllLoginHandler(loginSrv)
	router := httprouter.New()
	fmt.Println("To call the /login method")
	router.POST("/login", allLoginHand.UserLogin)
	fmt.Println("After calling the /login method")

	//ManageDoctorHandler
	manageDocRepo := repository.NewManageDoctorsRepository(dbconn)
	manageDocServ := service.NewManageDoctorsService(manageDocRepo)
	managDocHand := handler.NewManageDoctorsHandler(manageDocServ)
	router := httprouter.New()
	router.GET("/admin/doctors", managDocHand.GetDoctors)
	router.GET("/admin/doctor/?id", managDocHand.GetSingleDoctor)
	router.PUT("/admin/doctor/?id", managDocHand.UpdateDoctor)
	router.POST("/admin/doctor", managDocHand.AddDoctor)
	router.DELETE("/admin/doctor/?id", managDocHand.DeleteDoctor)

	//ManagePatientHandler
	managePatRepo := repository.NewManagePatientsRepository(dbconn)
	managepatServ := service.NewManagePatientsService(managePatRepo)
	managpatHand := handler.NewManagePatientHandler(managepatServ)
	router.GET("/admin/patients", managpatHand.GetPatients)
	router.GET("/admin/patient/?id", managpatHand.GetSinglePatient)
	router.PUT("/admin/patient/?id", managpatHand.UpdatePatient)
	router.POST("/admin/patient", managpatHand.AddPatient)
	router.DELETE("/admin/patient/?id", managpatHand.DeletePatient)

	//ManageLaboratoristHander
	manageLabRepo := repository.NewManageLaboratoristsRepository(dbconn)
	manageLabServ := service.NewManageLaboratoristsService(manageLabRepo)
	managLabHand := handler.NewManageLaboratoristHandler(manageLabServ)
	router.GET("/admin/laboratorists", managLabHand.GetLaboratorists)
	router.GET("/admin/laboratorist/?id", managLabHand.GetSingleLaboratorist)
	router.PUT("/admin/laboratorist/?id", managLabHand.UpdateLaboratorist)
	router.POST("/admin/laboratorist", managLabHand.AddLaboratorist)
	router.DELETE("/admin/laboratorist/?id", managLabHand.DeleteLaboratorist)

	//ManageAppointmentHandler
	manageAppRepo := repository.NewManageAppointmetRepository(dbconn)
	manageAppServ := service.NewManageAppointmetService(manageAppRepo)
	managAppHand := handler.NewManageAppointmentHandler(manageAppServ)
	router.GET("/admin/appointments", managAppHand.GetAppointments)
	router.GET("/admin/appointment/?id", managAppHand.GetSingleAppointment)
	router.PUT("/admin/appointment/?id", managAppHand.UpdateAppointment)
	router.POST("/admin/appointment", managAppHand.AddAppointment)
	router.DELETE("/admin/appointment/?id", managAppHand.DeleteAppointment)

	//ManageProfileHandler
	manageProRepo := repository.NewManageProfileRepository(dbconn)
	manageProServ := service.NewManageProfileService(manageProRepo)
	managProHand := handler.NewManageProfileHandler(manageProServ)
	router.GET("/admin/profiles", managProHand.GetProfiles)
	router.GET("/admin/profile/?id", managProHand.GetSingleProfile)
	router.PUT("/admin/profile/?id", managProHand.UpdateProfile)
	router.POST("/admin/profile", managProHand.AddProfile)
	router.DELETE("/admin/profile/?id", managProHand.DeleteProfile)

	//ManagePharmasistHandler
	managePhaRepo := repository.NewManagePharmasistsRepository(dbconn)
	managePhaServ := service.NewManagePharmasistsService(managePhaRepo)
	managPhaHand := handler.NewManagePharmasistHandler(managePhaServ)
	router.GET("/admin/doctors", managPhaHand.GetPharmasists)
	router.GET("/admin/doctor/?id", managPhaHand.GetSinglePharmasist)
	router.PUT("/admin/doctor/?id", managPhaHand.UpdatePharmasist)
	router.POST("/admin/doctors", managPhaHand.AddPharmasist)
	router.DELETE("/admin/doctors/?id", managPhaHand.DeletePharmasist)

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
