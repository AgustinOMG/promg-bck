package routes

import (
	"promg/controllers"

	"github.com/gin-gonic/gin"
)

func QuotesRoutes(router *gin.Engine) {

}

// #********** New registry
// api.add_resource(newRecord, '/Admin/NewRecord')                                  #       *? New Record registry POST
// api.add_resource(test, '/test')
// #********** User end points
// api.add_resource(newUser, '/Admin/NewUser')                                      #       *? Creates a new user in the staff of a certain Company GET
// api.add_resource(getUser, '/Admin/GetUser')                                      #       *? Returns de requested user GET
// api.add_resource(getStaff, '/Admin/GetStaff')                                     #      *? Returns all Users from a certain company GET
// api.add_resource(updateUser, '/Admin/UpdateUser')                                #       *? Updates the data of a certain user POST
// api.add_resource(deleteUser, '/Admin/DeleteUser')                                #       *? Complete delete the data of a certain Staff User POST

// #********** Company end points
// api.add_resource(validateCompany, '/Admin/ValidateCompany')                      #       *? Checks if a company is already on record to avoid errors on creation of new records GET
// api.add_resource(getCompany, '/Admin/GetCompany')                                #       *? Gets all the company data GET
// api.add_resource(updateCompany, '/Admin/UpdateCompany')                          #       *? Update the data of the company , POST
// api.add_resource(configDoc, '/Admin/ConfigDoc')                                  #       *? COnfiguration of the COmpany documents POTS, GET
// #
// #
// #  ****+++++++++++++++++++++++++++++++++++++++++++++++++++++                          ** COMPANY **

// api.add_resource(newCS, '/Company/NewCS')                                         #       *? Add a new client and / or supplier POST
// api.add_resource(getCS, '/Company/GetCS')                                         #       *? Get the information from Clients and Suppliers GET
// api.add_resource(deleteCS, '/Company/DeleteCS')                                   #       *? Delete client and / or supplier POST
// api.add_resource(updateCS, '/Company/UpdateCS')                                   #       *? Update a Client or Supplier Information POST
// api.add_resource(searchClient, '/Company/SearchClient')
// api.add_resource(searchSupplier, '/Company/SearchSupplier')

// api.add_resource(newI,'/Company/NewI')
// api.add_resource(getI,'/Company/GetI')
// api.add_resource(deleteI, '/Company/DeleteI')
// api.add_resource(updateI, '/Company/UpdateI')
// api.add_resource(searchItem, '/Company/SearchItem')

// api.add_resource(getP,'/Company/GetP')
// api.add_resource(newP,'/Company/NewP')
// api.add_resource(deleteP,'/Company/DeleteP')
// api.add_resource(updateP, '/Company/UpdateP')
// api.add_resource(getCurrencyRates, '/Company/GetCurrencyRates')
// #
// #

// #  ****+++++++++++++++++++++++++++++++++++++++++++++++++++++                             ** SALES **
// api.add_resource(newQuote, '/Sales/NewQuote')
// api.add_resource(updateQuote, '/Sales/UpdateQuote')
// api.add_resource(getQuote, '/Sales/GetQuote')                                      #       *? Get all the quote Information GET
// api.add_resource(deleteQuote, '/Sales/DeleteQuote')
// api.add_resource(searchQuote, '/Sales/SearchQuote')                                #       *? SEarch for specific quote by folio, nombre o titulo POST
// api.add_resource(searchQuoteClnt, '/Sales/SearchQuoteClnt')                        #       *? Search for the Registered Clientes for the quotes POST
// api.add_resource(numberToLetter, '/Sales/NumberToLetter')
// #
// #
// #                                                                                           ** COMPRAS **
// api.add_resource(newPO, '/Purchases/NewPO')
// api.add_resource(updatePO, '/Purchases/UpdatePO')
// api.add_resource(getPO, '/Purchases/GetPO')
// api.add_resource(deletePO, '/Purchases/DeletePO')
// api.add_resource(searchPO, '/Purchases/SearchPO')
// api.add_resource(searchPOSplr, '/Purchases/SearchPOSplr')

func UserRoutes(router *gin.Engine) {
	// llamada para la infromacion del usuario
	router.POST("/v1/administration/company/users/newuser", controllers.NewUser())
	router.GET("/v1/administration/company/users/getuser", controllers.GetAUser())
	router.PUT("/v1/administration/company/users/getstaff", controllers.GetAllStaff())
	router.DELETE("/v1/administration/company/users/updateuser", controllers.UpdateUser())
	router.GET("/v1/administration/company/users/deleteuser", controllers.DeleteAUser())
	router.GET("/v1/administration/access/checkuserstatus", controllers.CheckUserAccount()) //add this
}

func ItemRoutes(router *gin.Engine) {
	// llamada para la infromacion del usuario
	router.POST("/v1/administration/company/items/newitem", controllers.Newitem())
	router.GET("/v1/administration/company/items/getitems", controllers.Getitems())
	router.DELETE("/v1/administration/company/items/updateitem", controllers.Updateitem())
	router.GET("/v1/administration/company/items/deleteitem", controllers.DeleteAitem()) //add this
}
