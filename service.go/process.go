package service

import (
	"math"

	"github.com/tshubham7/adventure-island/models"
)

// returns the index
// we will use binary searching to find the user's current position
func binarySearch(arr []*models.User, min, max int, points int32, userID string) int {
	if min > max {
		return -1
	}
	if len(arr) == 1 {
		return 0
	}

	mid := int(math.Floor(float64((min + max) / 2)))

	if arr[mid].Points == points {
		// there may be possibilities for multiple users to have same points
		// we need to find the right user

		for i := mid; i <= max && arr[i].Points == points; i++ {
			if arr[i].ID == userID {
				return i
			}
		}
		for j := mid; j >= min && arr[j].Points == points; j-- {
			if arr[j].ID == userID {
				return j
			}
		}
		return -1
	}

	if arr[mid].Points < points {
		return binarySearch(arr, min, mid-1, points, userID)
	}

	return binarySearch(arr, mid, max+1, points, userID)
}

// shift user to the new rank position
func shift(arr []*models.User, index int, points int32) []*models.User {
	// go to left until you get the value greater than or equal to that

	for i := index; i > 0; i-- {

		if arr[i-1].Points >= points {
			break
		}
		arr[i], arr[i-1] = arr[i-1], arr[i]
	}

	return arr
}
