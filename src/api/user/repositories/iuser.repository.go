package user

import (
	userModel "app/src/api/user/models"
	baseRepository "app/src/shared/repository"
)

type IUserRepository interface {
	baseRepository.IBaseRepository[userModel.User]
}
