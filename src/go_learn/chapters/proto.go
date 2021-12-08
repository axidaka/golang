package chapters

import (
	"fmt"
	proto "github.com/golang/protobuf/proto"
	protocol "golang/src/go_learn/protocolbuf/go"
)

func AddressBook_test() {

	p := protocol.Person{
		Id: 1234,
		Name:"zzzz",
		Email: "xxx@qq.com",
		Phones: []*protocol.Person_PhoneNumber{
			{Number: "111111", Type: protocol.Person_HOME},
			{Number: "222222", Type: protocol.Person_WORK},
		}}

	book := &protocol.AddressBook{
		Person: []*protocol.Person{&p,
		},
	}

	out, err := proto.Marshal(book)
	if err != nil {
		fmt.Println("Failed to encode address book:", err)
	} else {
		fmt.Println("out:", out)
	}
}
