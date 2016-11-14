package router
import(
	"github.com/gorilla/mux"
	"metrics/api"
)

type systemMetricsRouter struct {}

func NewSystemMetricsRouter() *systemMetricsRouter{
	return &systemMetricsRouter{}
}

func (_ *systemMetricsRouter) registerRoutes(r *mux.Router){
	r.HandleFunc("/cpu",api.CpuMetricsHandler).Methods("GET")
	r.HandleFunc("/mem",api.MemMetricsHandler).Methods("GET")
	r.HandleFunc("/diskio",api.DiskIOMetricsHandler).Methods("GET")
	r.HandleFunc("/nwio",api.NwIOMetricsHandler).Methods("GET")

}

func (sMR *systemMetricsRouter) RegisterRoutes(){
	//sMR.registerRoutes(baseRouter)
	sMR.registerRoutes(RootRouter())
}