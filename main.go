package main

import (
	"cargo-hm1/structure"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/google/uuid"
)

//If we want to update the addresses of customers update anything about customer we have to use pointers because
// the customer update also has to change in the orders..
var customers = map[int]*structure.Customer{}

//We might be use the pointer for key values..
var orders = map[uuid.UUID]*structure.Order{}

func main() {
	//Dummy Datas
	//Customers
	customer1 := structure.NewCustomer(12345678900, "Adem", "Polat", 5554443322, "Erzurum")
	customer2 := structure.NewCustomer(12345678902, "Sevket", "Yılmaz", 5556667788, "Aydın")
	customers[customer1.IdNumber] = &customer1
	customers[customer2.IdNumber] = &customer2
	//Orders
	order1 := structure.NewOrder(customers[12345678900], customers[12345678902])
	order2 := structure.NewOrder(customers[12345678902], customers[12345678900])
	orders[order1.Id] = &order1
	orders[order2.Id] = &order2

	menu()
}

func menu() {
	clearScreen()
	fmt.Println("\t\t\t..::Welcome Cargo App::..")
	fmt.Println("[1] - See Orders \n[2] - See Customers \n[3] - Create New Order")
	fmt.Println("[4] - Change Order Status \n[5] - Change Customer Address\n[6] - Order Check")
	fmt.Println("[7] - Exit")
	fmt.Print("Your Choise: ")
	var choise int
	fmt.Scan(&choise)

	switch choise {
	case 1: //See Orders
		clearScreen()
		seeOrders()
		fmt.Println("\n---------------------------------------------------------------------")

		fmt.Println("\nTo See details Please Press 'D'\nReturn Menu Please Press 'Y'")
		var choise string
		fmt.Scan(&choise)
		if choise == "D" || choise == "d" {
			checkOrder()
		} else if choise == "Y" || choise == "y" {
			clearScreen()
			menu()
		}
		returnMenu()
	case 2: //See Customers
		clearScreen()
		seeCustomers()
		returnMenu()
	case 3: //Create New Order
		clearScreen()
		createNewOrder()
		returnMenu()
	case 4: //Change Order Status
		clearScreen()
		orderStatusUpdate()
		returnMenu()
	case 5: // Change Customer Address
		clearScreen()
		customerAddressUpdate()
		returnMenu()
	case 6: //Order Check
		clearScreen()
		checkOrder()
		returnMenu()
	case 7: //Exit
		return
	default:
		clearScreen()
		fmt.Println("Wrong Choise")
		returnMenu()
	}
}
func returnMenu() {
	fmt.Println("\n---------------------------------------------------------------------")

	fmt.Println("\nTo Return Menu Please Press 'Y'")
	var choise string
	fmt.Scan(&choise)
	if choise == "Y" || choise == "y" {
		clearScreen()
		menu()
	}
}
func clearScreen() {
	//To clear console in windows..
	// cmd := exec.Command("cmd", "/c", "cls")
	// cmd.Stdout = os.Stdout
	// cmd.Run()

	// To clear console in Mac or Linux
	// fmt.Println("\033[2J")
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func seeOrders() {
	fmt.Println("\nOrders: ")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("Order ID \t\t\t\t   Order Status \t\t Order Receiver \t\t\t\t Order Sender")
	fmt.Print("-------------------------------------- \t ---------------")
	fmt.Println("\t----------------------------------\t\t---------------------------")
	for k, v := range orders {
		fmt.Print(k)
		fmt.Print("\t  ", v.Status, "  \t")
		fmt.Print("\t", v.Receiver.IdNumber, " - ", v.Receiver.Name, " ", v.Receiver.LastName, "\t")
		fmt.Println("\t\t", v.Sender.IdNumber, " - ", v.Sender.Name, " ", v.Sender.LastName)
	}
}

func seeCustomers() {
	fmt.Println("\nCustomers: ")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("Customer ID \t\t Name \t\t\t LastName \t\t Phone Number \t\t Address")
	fmt.Println("--------------- \t ---------------\t-----------------\t------------------\t--------------")

	for k, v := range customers {
		fmt.Print(k, "\t")
		fmt.Println("\t", v.Name, "  \t\t", v.LastName, "  \t\t", v.PhoneNumber, "\t\t ", v.Address)
	}
}

func createNewOrder() {
	var newOrder structure.Order
	var newSender structure.Customer
	var newReceiver structure.Customer

	fmt.Println("Sender Values:")
	fmt.Println("---------------------------------------------------------------------")

	fmt.Print("Sender Id Number: \t")
	fmt.Scan(&newSender.IdNumber)
	if v, found := customers[newSender.IdNumber]; found {
		fmt.Println("We found a customer: ", v)
		fmt.Println("We'll continue with this customer..")
		newSender = *customers[newSender.IdNumber]
	} else {
		fmt.Print("Sender Name: \t")
		fmt.Scan(&newSender.Name)
		fmt.Print("Sender LastName: \t")
		fmt.Scan(&newSender.LastName)
		fmt.Print("Sender Phone: \t")
		fmt.Scan(&newSender.PhoneNumber)
		fmt.Print("Sender Address: \t")
		fmt.Scan(&newSender.Address)

		customers[newSender.IdNumber] = &newSender
	}

	fmt.Println("\nReceiver Values:")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Print("Receiver Id Number: \t")
	fmt.Scan(&newReceiver.IdNumber)
	if v, found := customers[newReceiver.IdNumber]; found {
		fmt.Println("We found a customer: ", v)
		fmt.Println("We'll continue with this customer..")
		newReceiver = *customers[newReceiver.IdNumber]
	} else {
		fmt.Print("Receiver Name: \t")
		fmt.Scan(&newReceiver.Name)
		fmt.Print("Receiver LastName: \t")
		fmt.Scan(&newReceiver.LastName)
		fmt.Print("Receiver Phone: \t")
		fmt.Scan(&newReceiver.PhoneNumber)
		fmt.Print("Receiver Address: \t")
		fmt.Scan(&newReceiver.Address)

		customers[newReceiver.IdNumber] = &newReceiver
	}
	newOrder = structure.NewOrder(&newReceiver, &newSender)
	orders[newOrder.Id] = &newOrder
	fmt.Println("\nOrder Created Successfully..")
}

func checkOrder() {
	fmt.Println("\nOrder Statuses: ")
	fmt.Println("---------------------------------------------------------------------")
	for k, v := range orders {
		fmt.Print("Order ID: ", k)
		fmt.Println("\tOrder Status: ", v.Status)
	}

	// var inquiry string
	// fmt.Print("\nOrder Status Filter: ")
	// fmt.Scan(&inquiry)

	// fmt.Println("Order Status Filter ")
	// fmt.Println("---------------------------------------------------------------------")

	// for k, v := range orders {
	// 	if v.Status == inquiry {
	// 		fmt.Println(k)
	// 	}
	// }

	var idInquiry string
	fmt.Print("\nPlease Enter Order Id You Want To Control: ")
	fmt.Scan(&idInquiry)

	fmt.Println("\nOrder Status")
	fmt.Println("---------------------------------------------------------------------")

	for k, v := range orders {
		if k == uuid.MustParse(idInquiry) {
			fmt.Println("\nOrder ID: ", k)
			fmt.Println("Order Status: ", v.Status)
			fmt.Println("--------------------")
			fmt.Println("\nOrder Sender: ")
			fmt.Println("Sender Id: ", v.Sender.IdNumber)
			fmt.Println("Sender Name LastName: ", v.Sender.Name, " ", v.Sender.LastName)
			fmt.Println("Sender Phone Number: ", v.Sender.PhoneNumber)
			fmt.Println("Sender Address: ", v.Sender.Address)
			fmt.Println("--------------------")
			fmt.Println("\nOrder Receiver: ")
			fmt.Println("Receiver Id: ", v.Receiver.IdNumber)
			fmt.Println("Receiver Name LastName: ", v.Receiver.Name, " ", v.Receiver.LastName)
			fmt.Println("Receiver Phone Number: ", v.Receiver.PhoneNumber)
			fmt.Println("Receiver Address: ", v.Receiver.Address)
		}
	}
}

func customerAddressUpdate() {
	fmt.Println("__Customer Address Update__")
	seeCustomers()

	fmt.Println("\nCustomer Address Update")
	fmt.Println("---------------------------------------------------------------------")

	var customerId int
	var newAddress string
	fmt.Print("\nPlease Enter the customers id: ")
	fmt.Scan(&customerId)
	if _, found := customers[customerId]; found {
		fmt.Print("Please Enter the new Address: \t")
		fmt.Scan(&newAddress)

		newCustomer := customers[customerId]
		// customers[customerId].SetCustomerAddress(newAddress)
		// customers[customerId].Address = newAddress
		newCustomer.SetCustomerAddress(newAddress)
		delete(customers, customerId)
		customers[newCustomer.IdNumber] = newCustomer
		// fmt.Println("Customers new List")
		// fmt.Println("---------------------------------------------------------------------")
		// for k, v := range customers {
		// 	fmt.Println(k, ": \t", v)
		// }
		fmt.Print("\nCustomer Address of ", customerId, " is cahnged to '", newAddress, "' successfully..")
	} else {
		fmt.Println("User not exist.. Try Again..")
		countDown(3)
		clearScreen()
		customerAddressUpdate()
	}

}

func orderStatusUpdate() {
	fmt.Println("__OrderStatus Update__")
	seeOrders()

	fmt.Println("\nOrder Status Update")
	fmt.Println("---------------------------------------------------------------------")

	var orderId string
	var statusChoise int
	fmt.Print("\nPlease Enter the Order Id: ")
	fmt.Scan(&orderId)

	//Validation
	_, err := uuid.Parse(orderId)
	if err == nil {
		if _, found := orders[uuid.MustParse(orderId)]; found {
			updateOrder := orders[uuid.MustParse(orderId)]

			fmt.Println("\t..Statuses..")
			fmt.Println("[1] - Order Picked up\n[2] - Order Delivered\n[3] - Order Cancelled")
			fmt.Print("Please Choose New Status: ")
			fmt.Scan(&statusChoise)
			if statusChoise == 1 {
				updateOrder.PickedUp()
			} else if statusChoise == 2 {
				updateOrder.Delivered()
			} else if statusChoise == 3 {
				updateOrder.Cancelled()
			} else {
				fmt.Println("Wrong Choise.. Status of order won't change..")
			}

			// customers[customerId].SetCustomerAddress(newAddress)
			// customers[customerId].Address = newAddress
			delete(orders, uuid.MustParse(orderId))
			orders[updateOrder.Id] = updateOrder

			fmt.Print("\nOrder Status of ", orderId, " is changed to '", updateOrder.Status, "' successfully..")
		} else {
			fmt.Println("Order not exist.. Try Again..")
			countDown(3)
			clearScreen()
			orderStatusUpdate()
		}

	} else {
		fmt.Println("Invalid Order Id.. Try Again..")
		countDown(2)
		clearScreen()
		orderStatusUpdate()
	}
}

func countDown(second int) {
	for second > 0 {
		if rand.Intn(100) == 1 {
			break
		}
		fmt.Println("Refreshing in ", second, "s")
		time.Sleep(time.Second)
		second--
	}
	return
}
