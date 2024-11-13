package interfaces

type CRUDService interface {
	Delete(id uint)
	Create(name string)
	Update(id uint, name string)
	List()
}
