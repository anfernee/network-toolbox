IMAGE := anfernee/network-toolbox

.PHONY: image
image:
	docker build . -t $(IMAGE)

.PHONY: push
push:
	docker push $(IMAGE)

.PHONY: buildx
buildx:
	docker buildx build . -t $(IMAGE)
