/* Basic if else */
if condition {
	//do something
}

if condition {
	do something
} else {
	//do something else
}

//Example
if userName != "" {
	//process user name
} else {
	//display error message
}

/*Nesting if statements */
if age >= 21 {
	//permit entry
} else if parentalApproval {
	//permit entry
} else {
	//cannot enter
}
//Can have multiple else if statements

/*if_else with Logical Operators*/
if hasTicket && ticketValid {
	//permit entry
}

if funds > cost && accountActive {
	//allow transaction
}

if usingDebitCard || usingCreditCard {
	//display card # entry screen
}

if !quizComplete {
	//display incomplete submission
	//warning
}


/*Usage With Functions*/
if termperature("freezer") > 0 {
	//display alert
}

if temp("freezer") > preferredTemp("freezer") {
	//display alert
}


/* Statement Initializtion*/
//Can create a variable and do a comparison at the same time
if i := 5; i < 10 {
	//do something with i
} else {
	//do something else with i
} 

if rank := getUserRank(); rank == "admin" {
	//do admin stuff here
} else if rank == "manager" {
	//do manager stuff here
} else {
	//display error message
}


/*Early Return*/
/*
Use "early return" to stop processing within a function as soon as possible
	Maximizes performance
	All data needed for the function is ready after error checking
		Leads to clean code*/

token, err := getSession("alice")
if err != nil {
	return
}
//if there is an error then the code returns early

cart, err := getCart(token)
if err != nil {
	return
}
//Cases like this is good because it prevents proceeding forward when the error is already found