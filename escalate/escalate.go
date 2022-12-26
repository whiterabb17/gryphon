package escalate

func Escalate(path string) string {
	err := escalate(path)
	return err.Error()
}
