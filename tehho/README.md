Din data placeras i data foldern

docker build -t tehho:latest .
docker run -it -v $(pwd)/data:/app/data tehho:latest