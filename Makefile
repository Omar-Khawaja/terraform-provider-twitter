build:
	@echo "building plugin"
	@cd scripts && ./build.sh

install:
	@echo "building plugin for provider based on latest release and installing in ~/.terraform.d/plugins/"
	@cd scripts/ && ./install.sh
