telepresence-install:
	@echo "installing telepresence for $(OS)"
ifeq ($(OS),MAC)
	brew cask install osxfuse
	brew install datawire/blackbird/telepresence
else ifeq ($(OS),UBUNTU)
	curl -s https://packagecloud.io/install/repositories/datawireio/telepresence/script.deb.sh | sudo bash
	sudo apt install --no-install-recommends telepresence
endif

telepresence-local-dev:
	telepresence --namespace $(NAMESPACE) --new-deployment $(DEPLOYMENT) --expose $(PORT)
