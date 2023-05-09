package B

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Use() {
	s := `password123`

	// fmt.Println("start", time.Now())

	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost) //generate hash of password

	// fmt.Println("end", time.Now())

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs)) //print hashed password

	fmt.Println(bcrypt.Cost(bs)) //get cost require to password

	// s = s + "34"
	isCorrect := bcrypt.CompareHashAndPassword(bs, []byte(s)) //verify the password

	//check status of password nil=>success and error=>not match(fail)
	if isCorrect != nil {
		fmt.Println("password incorrect", isCorrect)
	} else {
		fmt.Println("Password match")
	}

}
