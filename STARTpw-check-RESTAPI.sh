echo "BE AWARE SUDO IS REQUIRED"
echo "This code will run docker script"
echo "a error will happen if don't have docker"

sudo docker build . -t pwcheck-image
sudo docker run -d -p 8080:8080 --name pwcheck-container pwcheck-imag
mkdir testResults
docker cp pwCheck-container:/app/PasswordCheckJsonAPI/cover.html /testResults


