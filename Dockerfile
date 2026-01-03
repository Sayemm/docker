FROM ubuntu:24.04

RUN echo "Hello World" > /hello.txt

CMD ["cat", "/hello.txt"]

# Dockerfile is NOT an image. It is a recipe / blueprint that tells Docker how to build an image.

# docker build -t custom_ubuntu .    -> Build an image from a Dockerfile
# docker run -it custom_ubuntu bash  -> create container