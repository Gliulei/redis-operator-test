package controller

import (
	"github.com/Gliulei/redis-operator-test/pkg/controller/redisservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, redisservice.Add)
}
