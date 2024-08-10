CREATE_DMG ?= $(shell which create-dmg)

dev:
	wails dev

build: clean
	wails build

clean:
	rm -rf ./build/bin/*
	rm -rf ./mydesktopapp.dmg

dmg: build
	@@$(call check,create-dmg,brew install create-dmg)
	$(CREATE_DMG) \
		--volname "mydesktopapp" \
		--background "assets/background.png" \
		--window-size 558 367 \
		--icon-size 100 \
		--icon "mydesktopapp.app" 200 190 \
		--hide-extension "mydesktopapp.app" \
		--app-drop-link 370 190 \
		"mydesktopapp.dmg" \
		"build/bin/"