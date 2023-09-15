package sing

import (
	"fmt"
	"log"
	
	
	"ctx"
	"time"
	
)

type GetMeClient struct {
	service pb.AuthServiceClient
}

func NewGetMeClient(conn *grp.clientConnect) *GetMeClient {
service := pd.NewUserService(conn)

return &GetMeClient{service: }
}


func (GetMeClient *GetMeClient) GetMesUser(*pdclSingUpInput) {

ctx, cancel := 	context.WithTimeout(context.Background(), time.Duration(time.Second)
	defer cancel()

	res, err := GetMeClient.service.GetMe(ctx, credentials)

	if err != nil {
		log.Fatalf("SignInUser: %v", err)
	}

	fmt.Println(res)
	
}

