package storage

import (
	"../protocol"
	"sync"
	"time"
)

var (
	zoneMap  = make(map[string][]ZoneDomain)
	mutex    sync.RWMutex
	did      int
	didMutex sync.Mutex
)

func CreateZone(zone string) error {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := zoneMap[zone]; ok {
		return &protocol.Error{Message: "Zone already exist"}
	}
	zoneMap[zone] = make([]ZoneDomain, 0)
	return nil
}

func RemoveZone(zone string) error {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := zoneMap[zone]; !ok {
		return &protocol.Error{Message: "Zone not exist"}
	}
	delete(zoneMap, zone)
	return nil
}

func GetZones() []string {
	mutex.RLock()
	defer mutex.RUnlock()
	zones := make([]string, len(zoneMap))
	i := 0
	for k := range zoneMap {
		zones[i] = k
		i++
	}
	return zones
}

func CreateDomain(zone string, domain string, dtype protocol.Type, data string, ttl int) error {
	mutex.Lock()
	defer mutex.Unlock()
	if domains, ok := zoneMap[zone]; !ok {
		return &protocol.Error{Message: "Zone not exist"}
	} else {
		var domain = ZoneDomain{
			DomainId: getDid(),
			Domain:   domain,
			Zone:     zone,
			Type:     dtype,
			Data:     data,
			TTL:      ttl,
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
		}
		domains = append(domains, domain)
		zoneMap[zone] = domains
	}
	return nil
}

func UpdateDomain(zone string, domainId int, domain string, dtype protocol.Type, data string, ttl int) error {
	mutex.Lock()
	defer mutex.Unlock()
	if domains, ok := zoneMap[zone]; !ok {
		return &protocol.Error{Message: "Zone not exist"}
	} else {
		for i, d := range domains {
			if d.DomainId == domainId {
				d.Domain = domain
				d.Type = dtype
				d.Data = data
				d.TTL = ttl
				d.UpdateAt = time.Now()
				domains[i] = d
				zoneMap[zone] = domains
				break
			}
		}
	}
	return nil
}

func RemoveDomain(zone string, domainId int) error {
	mutex.Lock()
	defer mutex.Unlock()
	if domains, ok := zoneMap[zone]; !ok {
		return &protocol.Error{Message: "Zone not exist"}
	} else {
		for i, d := range domains {
			if d.DomainId == domainId {
				zoneMap[zone] = append(domains[:i], domains[i+1:]...)
				break
			}
		}
	}
	return nil
}

func GetDomains(zone string) ([]ZoneDomain, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	if domains, ok := zoneMap[zone]; ok {
		return domains, nil
	}
	return nil, &protocol.Error{Message: "Zone not exist"}
}

func getDid() int {
	didMutex.Lock()
	defer didMutex.Unlock()
	did++
	return did
}
