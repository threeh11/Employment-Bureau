package controllers

import "threeh.com/Employment_Bureau/internal/rep"

type DealingController struct {
	dealingRep *rep.DealingRepository
}

func NewDealingController(dealingRep *rep.DealingRepository) *DealingController {
	return &DealingController{dealingRep: dealingRep}
}
