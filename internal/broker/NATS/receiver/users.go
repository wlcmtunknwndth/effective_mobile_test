package receiver

//GetUserByPassport(ctx context.Context, passportNumber string) (*models.User, error)
//
//CreateUser(ctx context.Context, user *models.User) error
//GetUser(ctx context.Context, id uint64) (*models.User, error)
//DeleteUser(ctx context.Context, id uint64) error
//UpdateUser(ctx context.Context, user *models.User) error
//IsAdmin(ctx context.Context, id uint64) (bool, error)

//func (b *Receiver) createUser(ctx context.Context, user *models.User) error {
//	const op = scope + "CreateUser"
//	sub, err := b.conn.Subscribe(NATS.UserCreate, func(msg *nats.Msg) {
//		//msg.Data
//
//		var usrInf models.CreateUserAPI
//		if err := json.Unmarshal(msg.Data, &usrInf); err != nil {
//			b.log.Error("couldn't unmarshal data", sl.Op(op), sl.Err(err))
//			return
//		}
//
//		usr, usrInfo, err := models.CreateUserToUsersDB(&usrInf)
//		if err != nil{
//			b.log.Error("couldn't parse input data to ")
//		}
//		b.users.CreateUser(ctx, &usr, &usrInfo)
//
//	})
//}
