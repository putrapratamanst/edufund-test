package loan

import "edufund-test/service/loan"

type Controller struct {
	ln *loan.Service
}

func ControllerHandler(ln *loan.Service) *Controller {
	handler := &Controller{
		ln: ln,
	}
	return handler
}
