package roundrobin

import (
	"fmt"
	"sync"
)

type Endpoint struct {
	Endpoint string
	Weight   int
}

type RoundRobin struct {
	mutex *sync.Mutex
	// Current index (starts from -1)
	index         int
	endpoints     []Endpoint
	currentWeight int
}

func New(endpoints []Endpoint) (*RoundRobin, error) {
	if len(endpoints) == 0 {
		return nil, fmt.Errorf("no endpoints in the pool")
	}
	rr := &RoundRobin{
		index:     -1,
		mutex:     &sync.Mutex{},
		endpoints: endpoints,
	}
	return rr, nil
}

func (r *RoundRobin) NextEndpoint() string {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// The algo below may look messy, but is actually very simple
	// it calculates the GCD  and subtracts it on every iteration, what interleaves endpoints
	// and allows us not to build an iterator every time we readjust weights

	// GCD across all enabled endpoints
	gcd := r.weightGcd()
	// Maximum weight across all enabled endpoints
	max := r.maxWeight()

	for {
		r.index = (r.index + 1) % len(r.endpoints)
		if r.index == 0 {
			r.currentWeight = r.currentWeight - gcd
			if r.currentWeight <= 0 {
				r.currentWeight = max
				if r.currentWeight == 0 {
					panic("all endpoints have 0 weight")
				}
			}
		}
		ep := r.endpoints[r.index]
		if ep.Weight >= r.currentWeight {
			return ep.Endpoint
		}
	}
}

func (rr *RoundRobin) maxWeight() int {
	max := -1
	for _, e := range rr.endpoints {
		if e.Weight > max {
			max = e.Weight
		}
	}
	return max
}

func (rr *RoundRobin) weightGcd() int {
	divisor := -1
	for _, e := range rr.endpoints {
		if divisor == -1 {
			divisor = e.Weight
		} else {
			divisor = gcd(divisor, e.Weight)
		}
	}
	return divisor
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
