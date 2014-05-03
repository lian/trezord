package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"./go.hid"
)

func main() {
	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
	}
	handler.SetRoutes(
		&rest.Route{"GET", "/devices", GetDevices},
		&rest.Route{"GET", "/devices/:id", PostDevice},
	)
	http.ListenAndServe(":8080", &handler)
}

const vendor_id  = 0x534c
const product_id = 0x0001

func GetDevices(w rest.ResponseWriter, r *rest.Request) {
	list, err := hid.Enumerate(vendor_id, product_id)
	devices := make(map[string]string, 0)
	if err == nil {
		for _, v := range list {
			if v.InterfaceNumber == 0 || v.InterfaceNumber == -1 {
				dev, err := hid.OpenPath(v.Path)
				if err == nil {
					sn, _ := dev.SerialNumberString()
					devices[v.Path] = sn
				}
				dev.Close()
			}
		}
	}
	w.WriteJson(&devices)
}

func PostDevice(w rest.ResponseWriter, r *rest.Request) {
	code := r.PathParam("id")
	// r.DecodeJsonPayload
	result := make(map[string]string, 0)
	result["code"] = code
	w.WriteJson(&result)
}
