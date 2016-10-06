package main

import (
	"strconv"
	"strings"

	rr "github.com/larskluge/babl/babl-rpc/roundrobin"
)

func ParseEndpoints(endpoints string) []rr.Endpoint {
	res := []rr.Endpoint{}
	es := strings.Split(endpoints, ",")
	for _, e := range es {
		e = strings.TrimSpace(e)
		if e != "" {
			slice := strings.SplitN(e, ";q=", 2)
			endpoint := slice[0]
			weight := 1
			if len(slice) > 1 {
				w, err := strconv.ParseInt(slice[1], 10, 0)
				check(err)
				weight = int(w)
			}
			res = append(res, rr.Endpoint{endpoint, weight})
		}
	}
	return res
}
