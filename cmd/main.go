package cmd

//
// import (
// 	"fmt"
//
// 	"assignment3/domain"
// 	"assignment3/mapstore"
// )
//
// func Run() {
// 	controller := CustomerController{
// 		store: mapstore.NewMapStore(),
// 	}
// 	c1 := domain.Customer{
// 		ID:    "cust101",
// 		Name:  "John",
// 		Email: "john.doe@marketfeed.com",
// 	}
// 	controller.Add(c1)
// 	c2 := domain.Customer{
// 		ID:    "cust102",
// 		Name:  "Jack",
// 		Email: "jack@marketfeed.com",
// 	}
// 	controller.Add(c2)
// 	c3 := domain.Customer{
// 		ID:    "cust103",
// 		Name:  "Jerry",
// 		Email: "jerry@marketfeed.com",
// 	}
// 	controller.Add(c3)
// 	fmt.Println()
//
// 	controller.Update("cust101", domain.Customer{
// 		ID:    c1.ID,
// 		Name:  "Smarak",
// 		Email: "smarak@marketfeed.com",
// 	})
//
// 	controller.Update("cust102", domain.Customer{
// 		ID:    c2.ID,
// 		Name:  "Rob",
// 		Email: "rob@marketfeed.com",
// 	})
// 	fmt.Println()
//
// 	controller.Delete("cust103")
// 	fmt.Println()
// 	controller.GetByID("cust101")
// 	controller.GetByID("cust103")
// 	fmt.Println()
// 	controller.GetAll()
// }
