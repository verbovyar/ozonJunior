package interfaces

type Repository interface {
	Create(key string, value string) bool
	Read(key string) string
	Update(key string, value string) bool
	Delete(key string) bool
}
