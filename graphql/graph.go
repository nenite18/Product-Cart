package main

// central file to all graphql

//type Server struct {
//	accountClient *account.Client
//	catalogClient *catalog.Client
//	orderClient   *order.Client
//}

func NewGraphqlServer(accountUrl, catalogUrl, orderUrl) (*server, error) {
	accountClient, err := account.NewClient(accountUrl)

	catalogClient, err := order
}
