package targets

import (
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/targeter/targets/la_greenvn"
	"fmt"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/constants"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils"
	logger "github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/lib"
)

// GetLARandomAPITargeter provides a target of type LA_RANDOM_SEQUENCE.
func GetLARandomAPITargeter(seed int64, act models.Act, urls []string) (vegeta.Targeter, error) {
	return func(target *vegeta.Target) error {
		if target == nil {
			return vegeta.ErrNilTarget
		}
		var api string
		item, err := utils.GetRandomItemFromList(constants.LaGreenvnApis[:])
		if err != nil {
			logger.Error(err)
			return err
		}
		api = fmt.Sprint(item)
		switch api {
		case constants.LaGreenvnCreateAllocation:
			logger.Infof("%s API was used as target.", constants.LaGreenvnCreateAllocation)
			err = la_greenvn.GetLACreateGreenVNAllocationAPITarget(act, target)
			if err != nil {
				logger.Error(err)
				return err
			}
		case constants.LaGreenvnGetAllocationDetails:
			logger.Infof("%s API was used as target.", constants.LaGreenvnGetAllocationDetails)
			la_greenvn.GetLAGetGreenVNAllocationDetailsAPITarget(act, urls[0], target)

		case constants.LaGreenvnDeleteAllocation:
			logger.Infof("%s API was used as target.", constants.LaGreenvnDeleteAllocation)
			err = la_greenvn.GetLADeleteGreenVNAllocationAPITarget(act, urls, target)
			if err != nil {
				logger.Error(err)
				return err
			}
		}
		return nil
	}, nil
}
