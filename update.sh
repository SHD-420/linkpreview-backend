# Abort on errors
set -e

# Remove the client directory if already exists
sudo rm -rf client

# Clone from frontend repo (branch: deploy) into client directory
git clone --depth=1 --branch=deploy https://github.com/SHD-420/linkpreview-frontend.git client

# Eliminate source control from client (directory)
sudo rm -rf ./client/.git

echo "Frontend updated!"