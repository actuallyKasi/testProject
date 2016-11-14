#Docker installation
sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D
echo 'deb https://apt.dockerproject.org/repo ubuntu-trusty main' > /etc/apt/sources.list.d/docker.list
sudo apt-get update
sudo apt-get install docker-engine

#Pull docker image from docker hub
docker pull razorpay/devops:latest

#Create Dockerfile
echo 'FROM razorpay/devops:latest
RUN apk update
RUN apk add bash
ENV MYSQL_HOST="104.199.167.114"
ENV MYSQL_USER="razor"
ENV MYSQL_PASS="razor"
ENV MYSQL_PORT="razor"
EXPOSE 8080' > Dockerfile

#Build new image testkasi using Dockerfile
docker build -t testkasi .

#Start container using just created image
docker run -it -d -p 8080:8080 testkasi /rzp-interview