.DEFAULT_GOAL := all
BACKEND = trivia-api
FRONTEND = trivia-ui
BACKEND_DIR = backend
FRONTEND_DIR = frontend
BACKEND_VERSION = 1.0.0
FRONTEND_VERSION = 1.0.0
HUB_HOST = hub.docker.com
HUB_UNAME = rezaramadhan

build: build_backend build_frontend

build_backend:
	docker build $(BACKEND_DIR)/ -t $(BACKEND):$(BACKEND_VERSION)

build_frontend:
	docker build $(FRONTEND_DIR)/ -t $(FRONTEND):$(FRONTEND_VERSION)

push_backend:
	docker push $(HUB_HOST)/$HUB_UNAME/$(BACKEND):$(BACKEND_VERSION)

push_frontend:
	docker push $(HUB_HOST)/$HUB_UNAME/$(FRONTEND):$(FRONTEND_VERSION)

pull_backend:
	docker pull $(HUB_HOST)/$HUB_UNAME/$(BACKEND):$(BACKEND_VERSION)

pull_frontend:
	docker pull $(HUB_HOST)/$HUB_UNAME/$(FRONTEND):$(FRONTEND_VERSION)

all: build push

build: build_frontend build_backend

push: push_backend push_frontend

pull: pull_backend pull_frontendg
