package usecase

import (
	"go_inventory_ctrl/entity"
	"go_inventory_ctrl/repository"
)

type TrxInStUsecase interface {
	EnrollInsertTrx(EnrollTrxInSt *entity.TrxInST) string
}

type trxInStUsecase struct {
	trxInStRepo repository.TrxInStRepo
}

func (tx *trxInStUsecase) EnrollInsertTrx(EnrollTrxInSt *entity.TrxInST) string {
	return tx.trxInStRepo.EnrollInsertTrxInSt(EnrollTrxInSt)
}

func NewTrxInStUsecase(trxInSt repository.TrxInStRepo) TrxInStUsecase {
	return &trxInStUsecase{
		trxInStRepo: trxInSt,
	}
}
