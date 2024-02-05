package boothstrap

func Init() (err error) {
	err = LoadConfig()
	if err != nil {
		return err
	}
	InitLogger()
	err = InitDatabase()
	if err != nil {
		return err
	}
	return err
}
