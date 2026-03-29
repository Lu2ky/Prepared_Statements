package main
import(
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"fmt"
	"crypto/sha256"
)

var db *sql.DB
type User struct {
	ID int
	nombre string
	ciudad_id int
	fecha_creacion string
}
type UserUC struct {
	nombre string
	ciudad string
}
func main()  {
	cfg := mysql.Config{
		User: "root",
		Passwd: "2611",
		Net: "tcp",
		Addr: "localhost:3306",
		DBName: "prepared_statements",
		AllowNativePasswords: true,
		MultiStatements:      false,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	result, err1 := getUserByUsername("' OR id=4 OR '1' = '2'; DROP DATABASE prepared_statements; -- ")
	if err1 != nil {
		fmt.Println("Error:", err1)
		return
	}
	if result != nil {
		fmt.Printf("ID: %d, Nombre: %s, Ciudad ID: %d, Fecha Creacion: %s\n", result.ID, result.nombre, result.ciudad_id, result.fecha_creacion)
	}

	result1, err2 := getUserByUsernameStmt("' OR id='3")
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	if result1 != nil {
		fmt.Printf("ID: %d, Nombre: %s, Ciudad ID: %d, Fecha Creacion: %s\n", result1.ID, result1.nombre, result1.ciudad_id, result1.fecha_creacion)
	}

	var log_Login = login("Ana", "clave123")
	fmt.Printf("Login exitoso: %t\n", log_Login)

	userCity, err3 := getUserByUsernameAndCity("Maria", 1)
	if err3 != nil {
		fmt.Println("Error:", err3)
	}
	if userCity != nil {
		fmt.Printf("Nombre: %s, Ciudad: %s\n", userCity.nombre, userCity.ciudad)
	}

	defer db.Close()
}


func getUserByUsername(username string) (*User, error) {
	fmt.Println("-----------------------------------------SQL Injection-----------------------------------------")
	var user User
	var query = "SELECT id, nombre, ciudad_id, fecha_creacion FROM usuario WHERE nombre = '" + username + "'";
	fmt.Println("Query ejecutada:", query)
	err := db.QueryRow(query).Scan(&user.ID, &user.nombre, &user.ciudad_id, &user.fecha_creacion);

	if err != nil {
		return nil, err
	}
	return &user, nil
}
func getUserByUsernameStmt(username string) (*User, error) {
	fmt.Println( "-----------------------------------------Query preparada-----------------------------------------" ) 
	var user User
	var query = "SELECT id, nombre, ciudad_id, fecha_creacion FROM usuario WHERE nombre = ?"
	err := db.QueryRow(query, username).Scan(&user.ID, &user.nombre, &user.ciudad_id, &user.fecha_creacion);

	if err != nil {
		return nil, err
	}
	return &user, nil
}	

func hashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	encode := h.Sum(nil)
	return fmt.Sprintf("%x", encode)
	
}

func login(username string, password string) bool{
	fmt.Println("-----------------------------------------Login-----------------------------------------")
	passwdHash := hashPassword(password)
	var passStored string
	var query = "SELECT id FROM usuario WHERE nombre = ? AND clave = ?"
	err := db.QueryRow(query, username, passwdHash).Scan(&passStored);
	if err != nil {
		fmt.Println("Error al obtener la contraseña:", err)
		return false
	}else{
		return true;
	}
}

func getUserByUsernameAndCity(username string, cityID int) ([]*UserUC, error) {
	fmt.Println("-----------------------------------------getUserByUsernameAndCity-----------------------------------------")
	var query = "SELECT u.nombre, c.nombre FROM usuario u JOIN ciudad c ON u.ciudad_id = c.id WHERE u.nombre = ? AND c.id = ?"
	rows, err := db.Query(query, username, cityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*UserUC
	for rows.Next() {
		var user UserUC
		err := rows.Scan(&user.nombre, &user.ciudad);
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
