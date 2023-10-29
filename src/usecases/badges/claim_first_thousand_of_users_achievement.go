package badge_usecases

import (
	"github.com/robertokbr/bero-events/src/domain/enums"
	"github.com/robertokbr/bero-events/src/infra/database/repositories"
	"github.com/robertokbr/bero-events/src/logger"
)

type ClaimFirstThousandOfUsersAchievement struct {
	repository *repositories.MySqlRepository
}

func NewClaimFirstThousandOfUsersAchievement(repository *repositories.MySqlRepository) *ClaimFirstThousandOfUsersAchievement {
	return &ClaimFirstThousandOfUsersAchievement{
		repository: repository,
	}
}

func (self *ClaimFirstThousandOfUsersAchievement) Execute(userID int64) error {
	amountOfFirstHundredAchievementClaims, err := self.repository.GetAmountOfAchievementClaimsByAchievementID(int64(enums.FIRST_HUNDRED_OF_USERS))

	if err != nil {
		logger.Errorf("Error while getting amount of achievement claims for achievement %d: %s", enums.FIRST_HUNDRED_OF_USERS, err.Error())
		return err
	}

	if amountOfFirstHundredAchievementClaims < 101 {
		return nil
	}

	amountOfFirstThousandAchievementClaims, err := self.repository.GetAmountOfAchievementClaimsByAchievementID(int64(enums.FIRST_THOUSAND_OF_USERS))

	if err != nil {
		logger.Errorf("Error while getting amount of achievement claims for achievement %d: %s", enums.FIRST_THOUSAND_OF_USERS, err.Error())
		return err
	}

	if amountOfFirstThousandAchievementClaims > 1001 {
		return nil
	}

	userAchievements, err := self.repository.GetUserAchievementsByUserAndAchievementID(
		userID,
		int64(enums.FIRST_THOUSAND_OF_USERS),
	)

	if err != nil {
		logger.Errorf("Error while getting user %d achievement %d: %s", userID, enums.FIRST_THOUSAND_OF_USERS, err.Error())
		return err
	}

	if len(userAchievements) > 0 {
		logger.Debugf("User %d has already collected achievement %d", userID, enums.FIRST_THOUSAND_OF_USERS)
		return nil
	}

	err = self.repository.CreateUserAchievement(userID, int64(enums.FIRST_THOUSAND_OF_USERS))

	if err != nil {
		logger.Errorf("Error while creating user %d achievement %d: %s", userID, enums.FIRST_THOUSAND_OF_USERS, err.Error())
		return err
	}

	logger.Debugf("Achievement %d unlocked for user %d", enums.FIRST_THOUSAND_OF_USERS, userID)

	return nil
}