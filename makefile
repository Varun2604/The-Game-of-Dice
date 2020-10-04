install:
	echo "Installing go!"
	echo "PS: make command in progress, stay tuned for updates!"

build:
	go build

dist: 
	GOOS=darwin GOARCH=amd64 go build -v -o build/darwin_amd64/game_of_dice
	GOOS=linux GOARCH=amd64 go build -v -o build/linux_amd64/game_of_dice
	GOOS=windows GOARCH=amd64 go build -v -o build/windows_amd64/game_of_dice