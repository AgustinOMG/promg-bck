package routes

import (
	"promg/controllers"

	"github.com/gin-gonic/gin"
)

func QuotesRoutes(router *gin.Engine) {

}

func UserRoutes(router *gin.Engine) {
	// llamada para la infromacion del usuario
	router.POST("/v1/administration/company/users/newuser", controllers.NewUser())
	router.GET("/v1/administration/company/users/getuser", controllers.GetAUser())
	router.GET("/v1/administration/company/users/getstaff", controllers.GetStaff())
	router.POST("/v1/administration/company/users/updateuser", controllers.UpdateUser())
	//router.DELETE("/v1/administration/company/users/deleteuser", controllers.DeleteAUser())
	router.GET("/v1/administration/access/checkuserstatus", controllers.CheckUserAccount())
	router.POST("/v1/administration/access/registeruser", controllers.RegisterUser()) //add this
}

func CompanyRoutes(router *gin.Engine) {
	// llamada para la infromacion del usuario
	router.GET("/v1/administration/company/getcompany", controllers.GetACompany())
	router.POST("/v1/administration/company/updatecompany", controllers.UpdateCompany())
	router.POST("/v1/administration/company/updateconfiguration", controllers.UpdateConfiguration())
}

func ClientRoutes(router *gin.Engine) {
	// llamada para la infromacion del usuario
	router.POST("/v1/clients/newclient", controllers.NewClient())
	router.GET("/v1/clients/getclients", controllers.GetClients())
	router.POST("/v1/clients/updateclient", controllers.UpdateClient())
	router.DELETE("/v1/clients/deleteclient", controllers.DeleteClient())

}

func SupplierRoutes(router *gin.Engine) {
	// llamada para la infromacion del usuario
	router.POST("/v1/suppliers/newsupplier", controllers.NewSupplier())
	router.GET("/v1/suppliers/getsuppliers", controllers.GetSuppliers())
	router.POST("/v1/suppliers/updatesupplier", controllers.UpdateSupplier())
	router.DELETE("/v1/suppliers/deletesupplier", controllers.DeleteSupplier())

}

func ItemRoutes(router *gin.Engine) {
	// llamada para la infromacion del usuario
	router.POST("/v1/items/newitem", controllers.NewItem())
	router.GET("/v1/items/getitems", controllers.GetItem())
	router.POST("/v1/items/updateitem", controllers.UpdateItem())
	router.DELETE("/v1/items/deleteitem", controllers.DeleteItem())

}

func QuoteRoutes(router *gin.Engine) {
	// llamada para la infromacion del usuario
	router.POST("/v1/sales/quotes/newquote", controllers.NewQuote())
	router.POST("/v1/sales/quotes/copyquote", controllers.CopyQuote())
	router.GET("/v1/sales/quotes/getquotes", controllers.GetQuotes())
	// router.POST("/v1/quotes/updatequote", controllers.UpdateQuote())
	router.DELETE("/v1/sales/quotes/deletequote", controllers.DeleteQuote())

}
