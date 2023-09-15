import (
	"fmt"
	"log"
	
	
	"ctx"
	"time"
	
)

type SingUpClient struct {
	service pb.AuthServiceClient
}

func NewSingUpUserClient(conn *grp.clientConnect) *SingUpClient {
service := pd.AutoServiceClient(conn)

return &SingUpClient{service: }
}


func (singUpUserClient *SingInUserClient) SingUpsUser(*pdclSingUpInput) {

ctx, cancel := 	context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := signUpUserClient.service.SignInUser(ctx, credentials)

	if err != nil {
		log.Fatalf("SignInUser: %v", err)
	}

	fmt.Println(res)
	
}

