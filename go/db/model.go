package db

type User struct {
    Id   int64
    Name string  `xorm:"varchar(25) not null unique 'usr_name' comment('NickName')"`
}

type Conversion interface {
    FromDB([]byte) error
    ToDB() ([]byte, error)
}


