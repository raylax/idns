package api

import (
	"../protocol"
	"../storage"
	"fmt"
	"github.com/kataras/iris"
)

const (
	Host = "0.0.0.0"
	Port = 8080
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() {
	app := iris.Default()
	app.Get("/ping", func(ctx iris.Context) {
		writeResult(ctx, nil, "pong")
	})

	zone := app.Party("/zone")
	{
		zone.Get("s", func(ctx iris.Context) {
			writeResult(ctx, nil, storage.GetZones())
		})
		zone.Post("", func(ctx iris.Context) {
			z := struct {
				Name string `json:"name"`
			}{}
			ctx.ReadJSON(&z)
			err := storage.CreateZone(z.Name)
			writeResult(ctx, err, nil)
		})
		zone.Post("/{zone:string}", func(ctx iris.Context) {
			z := ctx.Params().GetString("zone")
			err := storage.RemoveZone(z)
			writeResult(ctx, err, nil)
		})
	}

	domain := zone.Party("/{zone:string}/domain")
	domain.Get("s", func(ctx iris.Context) {
		z := ctx.Params().GetString("zone")
		domains, err := storage.GetDomains(z)
		writeResult(ctx, err, domains)
	})

	domain.Post("", func(ctx iris.Context) {
		z := ctx.Params().GetString("zone")
		d := struct {
			Domain string `json:"domain"`
			Type   int    `json:"type"`
			Data   string `json:"data"`
			TTL    int    `json:"ttl"`
		}{}
		ctx.ReadJSON(&d)
		dtype := protocol.ParseType(d.Type)
		err := storage.CreateDomain(z, d.Domain, dtype, d.Data, d.TTL)
		writeResult(ctx, err, nil)
	})

	domain.Post("/{did:int}", func(ctx iris.Context) {
		z := ctx.Params().GetString("zone")
		did, _ := ctx.Params().GetInt("did")
		d := struct {
			Domain string `json:"domain"`
			Type   int    `json:"type"`
			Data   string `json:"data"`
			TTL    int    `json:"ttl"`
		}{}
		ctx.ReadJSON(&d)
		dtype := protocol.ParseType(d.Type)
		err := storage.UpdateDomain(z, did, d.Domain, dtype, d.Data, d.TTL)
		writeResult(ctx, err, nil)
	})

	domain.Delete("/{did:int}", func(ctx iris.Context) {
		z := ctx.Params().GetString("zone")
		did, _ := ctx.Params().GetInt("did")
		err := storage.RemoveDomain(z, did)
		writeResult(ctx, err, nil)
	})

	err := app.Run(iris.Addr(fmt.Sprintf("%v:%v", Host, Port)))
	if err != nil {
		panic(err)
	}
}

type result struct {
	Code int         `json:"code"`
	Data interface{} `json:"message,omitempty"`
}

func writeResult(ctx iris.Context, err error, data interface{}) {
	if err != nil {
		ctx.JSON(result{Code: 500, Data: err.Error()})
	} else {
		ctx.JSON(result{Code: 0, Data: data})
	}
}
