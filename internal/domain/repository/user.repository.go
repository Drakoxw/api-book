package repository

// import (
// 	"api-book/internal/domain/models"
// 	"database/sql"
// 	"log"
// )

// type UserRepository interface {
// 	FindAllUsers() ([]models.User, error)
// 	Save(user *models.User) error
// 	Update(user *models.User) error
// 	Delete(user *models.User) error
// }

// type UserRepositoryImpl struct {
// 	db *sql.DB
// }

// func NewUserRepository(db *sql.DB) UserRepository {
// 	return &UserRepositoryImpl{
// 		db: db,
// 	}
// }

// func (ur *UserRepositoryImpl) FindAllUsers() ([]models.User, error) {
// 	query := "SELECT id, username, password, email, created_at, updated_at FROM users"

// 	rows, err := ur.db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users []models.User

// 	for rows.Next() {
// 		var user models.User
// 		err := rows.Scan(&user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
// 		if err != nil {
// 			log.Printf("Error al escanear los resultados: %v", err)
// 			continue
// 		}

// 		users = append(users, user)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }

// func (repo *UserRepositoryImpl) Save(user *models.User) error {
// 	// Implementación para guardar un usuario en la capa de persistencia
// 	return nil
// }

// func (repo *UserRepositoryImpl) Update(user *models.User) error {
// 	// Implementación para actualizar un usuario en la capa de persistencia
// 	return nil
// }

// func (repo *UserRepositoryImpl) Delete(user *models.User) error {
// 	// Implementación para eliminar un usuario en la capa de persistencia
// 	return nil
// }
