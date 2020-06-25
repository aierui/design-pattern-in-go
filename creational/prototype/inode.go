package prototype

type inode interface {
	print(string)
	clone() inode
}
