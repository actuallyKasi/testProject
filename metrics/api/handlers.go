package api

import (
	"net/http"
	"metrics/base/httputils"
	"metrics/base/utils"
	//"metrics/base/errors"
)

func CpuMetricsHandler(rw http.ResponseWriter, request *http.Request){

	CpuMetricsResponseObject := GetCpuMetrics()
	ReturnJson, err := utils.ObjectToJson(CpuMetricsResponseObject)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteResponse(&rw, "json", ReturnJson, http.StatusOK)
}

func MemMetricsHandler(rw http.ResponseWriter, request *http.Request){
	MemMetricsResponseObject := GetMemMetrics()
	ReturnJson, err := utils.ObjectToJson(MemMetricsResponseObject)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteResponse(&rw, "json", ReturnJson, http.StatusOK)
}

func DiskIOMetricsHandler(rw http.ResponseWriter, request *http.Request){
	DiskIOMetricsResponseObject := GetDiskIOMetrics()
	ReturnJson, err := utils.ObjectToJson(DiskIOMetricsResponseObject)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteResponse(&rw, "json", ReturnJson, http.StatusOK)
}

func NwIOMetricsHandler(rw http.ResponseWriter, request *http.Request){
	NwIOMetricsResponseObject := GetNwIOMetrics()
	ReturnJson, err := utils.ObjectToJson(NwIOMetricsResponseObject)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteResponse(&rw, "json", ReturnJson, http.StatusOK)
}


