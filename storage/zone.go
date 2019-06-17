package storage

import (
	"../protocol"
	"time"
)

type ZoneDomain struct {
	DomainId int
	Domain   string
	Zone     string
	Type     protocol.Type // supported A/CNAME only
	Data     string
	TTL      int
	CreateAt time.Time
	UpdateAt time.Time
}
