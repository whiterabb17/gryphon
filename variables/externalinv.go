package variables

func ExternalRunPE(payloadPath string, targetPath string, arguments string) (int, bool) {
	return externalRunPE(payloadPath, targetPath, arguments)
}
