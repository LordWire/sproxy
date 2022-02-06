package main

// this is going to be our global object that everybody's gonna be operating on
// this simple solution
// For atomic operations (e.g. if we want to perform host uptime checks) we need to add
// a mutex on ServiceBackends and lock on that.
//
// On a first thought, I would prefer to remove the non-working backends from the
// ServiceBackends.Backend[] slice (instead of e.g. holding a status boolean flag
// in the BackendService itself), for a number of reasons:
// - the rest of the get<some>Backend() functions could still operate without checking if the backend is alive (reduces operations)
// - the operation of removing a backend from the slice is exponentially cheaper if we don't care about the slice order

import (
	"log"
	"math/rand"
	"strconv"
)

type BackendService struct {
	Address string
	Port    int
}

type ServiceBackends struct {
	Backend []BackendService
	len     int
}

func (sb *ServiceBackends) loadBackends(config ServicesConfig) {
	for i := 0; i < len(config.Hosts); i++ {
		sb.Backend = append(sb.Backend, BackendService(config.Hosts[i]))
		log.Printf("appended backend %s, port %s", config.Hosts[i].Address, strconv.Itoa(config.Hosts[i].Port))
	}
	sb.len = len(config.Hosts)
}

// operate on an array of backends and get a random one back.
// no atomicity needed here yet
// For an actual production service, it would be really nice to avoid
// Itoa conversions on every request.
func (sb *ServiceBackends) getRandomBackend() string {
	index := rand.Intn(sb.len)
	return sb.Backend[index].Address + ":" + strconv.Itoa(sb.Backend[index].Port)
}

// loop through the backends and do a liveness check
// this should be invoked from a goroutine somewhere else,
// _after_ the ServiceBackends struct becomes thread safe.
// func (sb *ServiceBackends) liveCheck(config ServicesConfig) {
// 	// do stuff
// }

// RoundRobin
// func (sb *ServiceBackends) getRoundRobinBackend() string {
// 	//do stuff
// 	return "the next proxy in line"
// }
