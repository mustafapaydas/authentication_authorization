package repo

import (
	"authenticaiton-authorization/entity"
	"authenticaiton-authorization/utils"
	"context"
)

type UserRepoPsql struct {
	*AbstractRepo
}

var repo = AbstractRepo{DriverName: "postgres"}

func (psql *UserRepoPsql) Create(user entity.User) (*int, error) {

	insertQuery := `
		 INSERT INTO tbl_user
			(username, first_name, last_name, email, full_name, "password", phone_number)
			VALUES($1,  $2,  $3,  $4,  $5, $6, $7) returning id;`
	ctx := context.Background()
	var id = new(int)

	err := repo.Create(ctx, insertQuery, id, user.UserName, user.FirstName, user.LastName, user.Email, user.FullName, user.Password, user.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (psql *UserRepoPsql) FindById(id int) (*entity.User, error) {

	var user = new(entity.User)
	err := repo.FindById(`select u.*, json_agg(r.*) as roles  from tbl_user as u
		left join tbl_user_role_relation as rel on rel.user_id  = u.id
		left join tbl_role as r on r.id = rel.role_id
		where u.id = $1
		group by u.id`, user, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (psql *UserRepoPsql) GetUserByUsernameOrEmail(user entity.User) (*entity.User, error) {
	err := repo.RunQuery(`select u.*, json_agg(r.*) as roles  from tbl_user as u
		left join tbl_user_role_relation as rel on rel.user_id  = u.id
		left join tbl_role as r on r.id = rel.role_id
		where u.username = $1 or u.email = $2
		group by u.id`, &user, user.UserName, user.Email)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, &utils.BusinessException{
				Message: "Not Matched User Name or Email",
			}
		}
		return nil, err
	}
	return &user, nil
}

func (psql *UserRepoPsql) ExistUser(user entity.User) error {
	query := `select username, email, phone_number  from tbl_user tu where username = $1 or email = $2 or phone_number = $3`
	var existUser = new(entity.User)
	err := repo.RunQuery(query, existUser, user.UserName, user.Email, user.PhoneNumber)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil
		}
		return err
	}
	if existUser != nil {
		err = user.CompareUniqueColumns(existUser)
		return err
	}

	return nil
}
