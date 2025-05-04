
TAG=$(git rev-parse HEAD)

docker build -t learn .

aws ecr get-login-password --region "eu-west-2" | docker login --username AWS --password-stdin 864899866496.dkr.ecr.eu-west-2.amazonaws.com/learn

docker tag learn 864899866496.dkr.ecr.eu-west-2.amazonaws.com/learn:$TAG

docker push 864899866496.dkr.ecr.eu-west-2.amazonaws.com/learn:$TAG