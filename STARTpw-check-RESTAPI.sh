echo "BE AWARE SUDO IS REQUIRED"
echo "This code will run docker script"
echo "a error will happen if don't have docker"
echo "every container will be deleted"
sudo docker rmi pwcheck-image
sudo docker rm $(docker ps -aq)
sudo docker stop pwcheck-container
sudo docker rm pwcheck-container
sudo docker build . -t pwcheck-image
sudo docker run -d -p 8080:8080 --name pwcheck-container pwcheck-image
echo "Please tell new folder for tests Results.html"
testpath=
mkdir $testpath
docker cp pwcheck-container:/app/cover.html $testpath
docker cp pwcheck-container:/app/cover.txt $testpath
docker cp pwcheck-container:/app/checkEveryTexts.txt $testpath
cat $testpath/checkEveryTexts.txt
cat $testpath/cover.txt


