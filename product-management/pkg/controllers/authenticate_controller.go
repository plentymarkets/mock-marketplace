package controllers

type Person struct {
	ID string `uri:"id" binding:"required,uuid"`
}

//type AuthenticateController struct {
//	authenticateRepository repositories.UserRepositoryContract
//}
//
//func NewAuthenticateController(authenticateRepository repositories.AuthenticateRepositoryContract) AuthenticateController {
//	return AuthenticateController{
//		authenticateRepository: authenticateRepository,
//	}
//}
//
//func (controller *AuthenticateController) GetAll() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Done()
//	}
//}
