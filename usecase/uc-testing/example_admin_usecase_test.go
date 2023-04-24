package uc_testing

import (
	"testing"

	"go_inventory_ctrl/entity"

	"go_inventory_ctrl/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyExampleAdmin = []entity.ExampleAdmin{
	{ID: "ea1", Name: "admin1", Email: "admin1@gmail.com", Password: "12345", Phone: "010101", Photo: "admin1.png"},
	{ID: "ea2", Name: "admin2", Email: "admin2@gmail.com", Password: "12345", Phone: "010101", Photo: "admin2.png"},
	{ID: "ea3", Name: "admin3", Email: "admin3@gmail.com", Password: "12345", Phone: "010101", Photo: "admin3.png"},
}

type repoMockExampleAdmin struct {
	mock.Mock
}

func (r *repoMockExampleAdmin) GetAllExampleAdmin() any {
	args := r.Called()
	if args.Get(0) == nil {
		return "no data"
	}
	return args.Get(0).([]entity.ExampleAdmin)
}

func (r *repoMockExampleAdmin) GetByIdExampleAdmin(id string) any {
	return nil
}

func (r *repoMockExampleAdmin) CreateExampleAdmin(newExampleAdmin *entity.ExampleAdmin) string {
	args := r.Called(newExampleAdmin)
	if args[0] != nil {
		return "failed to create example admin"
	}
	return "example admin created successfully"
}

func (r *repoMockExampleAdmin) UpdateExampleAdmin(exampleAdmin *entity.ExampleAdmin) string {
	return "test"
}

func (r *repoMockExampleAdmin) DeleteExampleAdmin(id string) string {
	return "test"
}

type ExampleAdminUsecaseTesSuite struct {
	repoMockExampleAdmin *repoMockExampleAdmin
	suite.Suite
}

// // ======================================== FindAll Test ====================================================================
func (suite *ExampleAdminUsecaseTesSuite) TestFindAllExampleAdmin_Success() {
	exampleAdminUc := usecase.NewExampleAdminUsecase(suite.repoMockExampleAdmin)
	suite.repoMockExampleAdmin.On("GetAllExampleAdmin").Return(dummyExampleAdmin)
	exampleAdmin := exampleAdminUc.FindAllExampleAdmin()
	assert.Equal(suite.T(), exampleAdmin, dummyExampleAdmin)
}

func (suite *ExampleAdminUsecaseTesSuite) TestFindAllExampleAdmin_Failed() {
	exampleAdminUc := usecase.NewExampleAdminUsecase(suite.repoMockExampleAdmin)
	suite.repoMockExampleAdmin.On("GetAllExampleAdmin").Return(nil)
	exampleAdmin := exampleAdminUc.FindAllExampleAdmin()
	assert.Equal(suite.T(), exampleAdmin, "no data")
}

// // ======================================== FindById Test ====================================================================
// func (suite *ExampleAdminUsecaseTesSuite) TestFindByIdlExampleAdmin_Success() {

// 	exampleAdminUc := usecase.NewExampleAdminUsecase(suite.repoMockExampleAdmin)
// 	suite.repoMockExampleAdmin.On("GetByIdExampleAdmin", ).Return(dummyExampleAdmin)
// 	exampleAdmin := exampleAdminUc.FindByIdExampleAdmin()
// 	assert.Equal(suite.T(), exampleAdmin, dummyExampleAdmin)
// }

// func (suite *ExampleAdminUsecaseTesSuite) TestFindByIdExampleAdmin_Failed() {
// 	exampleAdminUc := usecase.NewExampleAdminUsecase(suite.repoMockExampleAdmin)
// 	suite.repoMockExampleAdmin.On("GetByIdExampleAdmin").Return("")
// 	exampleAdmin := exampleAdminUc.FindByIdExampleAdmin()
// 	assert.Equal(suite.T(), exampleAdmin, "no data")
// }

// ======================================== Register Test ====================================================================
func (suite *ExampleAdminUsecaseTesSuite) TestRegisterExampleAdmin_Success() {
	exampleAdminUc := usecase.NewExampleAdminUsecase(suite.repoMockExampleAdmin)
	suite.repoMockExampleAdmin.On("CreateExampleAdmin", &dummyExampleAdmin[0]).Return(nil)
	value := exampleAdminUc.RegisterExampleAdmin(&dummyExampleAdmin[0])
	assert.Equal(suite.T(), value, "example admin created successfully")
}

func (suite *ExampleAdminUsecaseTesSuite) TestRegisterExampleAdmin_Failed() {
	exampleAdminUc := usecase.NewExampleAdminUsecase(suite.repoMockExampleAdmin)
	suite.repoMockExampleAdmin.On("CreateExampleAdmin", &dummyExampleAdmin[0]).Return("failed to create example admin")
	value := exampleAdminUc.RegisterExampleAdmin(&dummyExampleAdmin[0])
	assert.Equal(suite.T(), value, "failed to create example admin")
}

// //==============================================Set Up Test	=======================================================================

func (suite *ExampleAdminUsecaseTesSuite) SetupTest() {
	suite.repoMockExampleAdmin = new(repoMockExampleAdmin)
}

func TestExampleAdminUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleAdminUsecaseTesSuite))
}
