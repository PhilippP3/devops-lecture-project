# devops-lecture-project-2025

### Project for the lecture "Einführung in DevOps, Continuous Delivery Tools und Mindset"

## Webshop
This Webshop provides different modern Smartphones. For example The Samsung Galaxy A35.


## Rollout
The application can be used via a docker image on linux/amd64 and linux/arm64 pulled using
docker pull.

The docker image is available at following link: https://hub.docker.com/repository/docker/fl012/webshop.

Tip for windows: docker run -d -p 8080:8080 webshop