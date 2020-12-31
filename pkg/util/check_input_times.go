package util

import (
	"time"
	mErrors "wallet/errors"
	"wallet/types"
)

type CheckBase struct {
	Is        bool
	CreatedAt types.Second
}

func AnalyseRecord(record []CheckBase, size int, gap time.Duration, banTime time.Duration) (int, error) {
	var (
		rightNum    int
		continueNum int
		isRight     bool
		isFind      bool
	)

	if len(record) == 0 {
		return 1, nil
	}
	if len(record) == size {
		for i := len(record) - 1; i >= 0; i-- {
			if record[i].Is {
				rightNum = i + 2
				isRight = true
			}
			if time.Unix(record[i].CreatedAt.ToInt64(), 0).
				Add(gap).After(time.Now()) && !isFind {
				continueNum = i + 2
				isFind = true
			}
		}
		if !isRight {
			timeOver := time.Unix(record[size-1].CreatedAt.ToInt64(), 0).Add(banTime)
			if !timeOver.After(time.Now()) {
				return continueNum, nil
			}
			return 0, mErrors.OverLimitTimes.
				Format("The input times more than 5")
		}
	} else {
		if time.Now().After(time.Unix(record[0].CreatedAt.ToInt64(), 0).
			Add(gap)) {
			return 1, nil
		}

		for i := len(record) - 1; i >= 0; i-- {
			if record[i].Is {
				rightNum = i + 2
				isRight = true
			}
			if time.Unix(record[i].CreatedAt.ToInt64(), 0).
				Add(gap).After(time.Now()) && !isFind {
				continueNum = i + 2
				isFind = true
			}
		}
	}

	if rightNum != continueNum {
		if rightNum > continueNum || rightNum == 0 {
			return continueNum, nil
		}
		return rightNum - 1, nil
	}

	if rightNum == 0 {
		return 1, nil
	}

	return rightNum - 1, nil
}
