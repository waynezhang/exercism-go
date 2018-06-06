package erratum

// Use doc here
func Use(o ResourceOpener, input string) (err error) {
	var res Resource
	for {
		var r, err = o()
		if err == nil {
			res = r
			break
		}
		if _, ok := err.(TransientError); ok {
			continue
		}
		return err
	}
	defer res.Close()

	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case FrobError:
				res.Defrob(r.(FrobError).defrobTag)
				err = r.(FrobError).inner
			default:
				err = r.(error)
			}
		}
	}()

	res.Frob("hello")
	return nil
}
