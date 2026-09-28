package simple

type SimpleRepository struct {
}

type SimpleService struct {
	*SimpleRepository
}

// provider (constructor)

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

func NewSimpleService(repository *SimpleRepository) *SimpleService {
	return &SimpleService{
		SimpleRepository: repository,
	}
}
