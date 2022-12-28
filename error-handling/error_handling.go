package erratum

func Use(opener ResourceOpener, input string) error {
	var res Resource
	var err error

	for {
		res, err = opener()
		if err != nil {
			if _, ok := err.(TransientError); ok {
				continue
			}
			return err
		}
		break
	}

	func() {
		defer func() {
			if rec := recover(); rec != nil {
				switch v := rec.(type) {
				case FrobError:
					res.Defrob(v.defrobTag)
					err = v
				case error:
					err = v
				}
			}
		}()

		res.Frob(input)
	}()

	res.Close()
	return err
}
