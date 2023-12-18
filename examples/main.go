package main

import (
	"fmt"
	"github.com/yahya077/parasut"
	"os"
)

func main() {
	_ = os.Setenv(parasut.UsernameEnv, "yahyahindioglu@gmail.com")
	_ = os.Setenv(parasut.PasswordEnv, "aGFc95!M!KpjX8k")
	_ = os.Setenv(parasut.CompanyIDEnv, "668246")

	room := parasut.NewRoom()

	response := room.ContactIndex(parasut.ContactIndexParams{
		HasInclude: parasut.HasInclude{
			Include: "contact_people",
		},
	})

	fmt.Println("RequestUri", response.RequestUri())
	fmt.Println("RequestHeader", response.RequestHeader())
	fmt.Println("StatusCode", response.StatusCode())
	fmt.Println("Header", response.Header())
	fmt.Println("Status", response.Status())
	fmt.Println("Ok", response.Ok())
	fmt.Println("Dto", response.Dto().(*parasut.ContactResponse).Data[0].Relationships.ContactPeople.Data)
	fmt.Println("RequestMethod", response.RequestMethod())

	fmt.Println("TimeElapsed", room.GetSegment().GetElapsedTime())
}
