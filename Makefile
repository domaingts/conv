build:
	CGO_ENABLED=0 GOOS=linux go build -v -trimpath -ldflags "-X 'github.com/domaingts/conv/cmd.Version=0.0.2' -s -w -buildid=" -o conv ./

pack:
	tar czvf conv.tar.gz conv