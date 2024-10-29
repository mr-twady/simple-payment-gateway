package service

import "dunsin-olubobokun/simple-payment-gateway/internal/models"

func (s *TransactionService) GetUserBalance(email string) (float64, error) {
	var user models.User
	if err := s.UserRepo.FindUserByEmail(email, &user); err != nil {
		return 0, err
	}
	return user.Balance, nil
}
