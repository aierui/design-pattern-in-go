package proxy

type server interface {
	handleRequest(string, string) (int, string)
}
