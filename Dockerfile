# https://hub.docker.com/_/golang
FROM golang:1.13

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

ENV PORT=8000
ENV CHANNEL_SECRET=aa97d859a0052efb70c819f91ef050a0
ENV CHANNEL_ACCESS_TOKEN=VjHGPM68AQW1OjorD96mQKMQhsZleXzBOI1gucJYZEv0JcookrjQy9gX42lIhvDcb2KkaPCukxDnUN5eM4uvXYcUgnH2S92hTNIeOeZN1kHaiCfl3CVeU71rBoa5gyIRlCA3p39tIODk/m0KzilnzAdB04t89/1O/w1cDnyilFU=
ENV APP_ID=735bad59
ENV APP_KEY=46c857cc9f9f797440f59e425ce82ead

RUN go build -o server main.go
CMD [ "./server" ]
