echo "BE AWARE SUDO IS REQUIRED"
echo "This code will run docker script"
echo "a error will happen if don't have docker"

sudo docker build . -t pwcheck-image
sudo docker run  -p 8080:8080 --name pwcheck-container pwcheck-image
echo "Please tell path for tests Results html"
read testpath
mkdir $testpath
docker cp pwcheck-container:/app/cover.html $testpath


