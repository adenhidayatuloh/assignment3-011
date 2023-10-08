# assignment3-011

Nama : Aden Hidayatuloh
Kode : GLNG-08-011

table database : 

create table waterwind (
id int primary key,
water int,
wind int
);

insert dengan data id 1 terlebih dahulu : 
insert into waterwind values (1,0,0);

data koneksi : 
const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "assigment2"
)

