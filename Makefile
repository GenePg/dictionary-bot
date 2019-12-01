ZONE = asia-northeast1-b
PROJECT_ID = dictionary-chatbot
DOMAIN = asia.gcr.io

IMAGE_NAME = dictionary-bot
IMAGE_TAG = prod-v1
PROD_IMAGE = $(DOMAIN)/$(PROJECT_ID)/$(IMAGE_NAME):$(IMAGE_TAG)

env:
	gcloud config set project $(PROJECT_ID)
	gcloud config set compute/zone $(ZONE)

build-image:
	docker build -f Dockerfile -t $(PROD_IMAGE) .

push-image:
	docker push $(PROD_IMAGE)

cloud-build:
	gcloud builds submit --tag $(PROD_IMAGE)

deploy:
	gcloud run deploy --image $(PROD_IMAGE) --platform managed
