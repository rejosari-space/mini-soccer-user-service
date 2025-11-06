package error

func ErrMapping(err error) bool {
	allErrors := make([]error, 0) //make empty slice
	allErrors = append(append(GeneralErrors[:], userErrors[:]...))

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}

	return false

}

// versi readable

// func IsKnownError(err error) bool {
//     if err == nil {
//         return false
//     }

//     knownErrors := append(GeneralErrors, userErrors...) //tampung semua error ke dalam variable untuk di looping

//     for _, e := range knownErrors {
//         if errors.Is(err, e) {
//             return true
//         }
//     }
//     return false
// }
