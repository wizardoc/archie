DIST_NAME="archie";
DIST_DIR="../build";

echo "Begin build...\n";

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $DIST_DIR/$DIST_NAME ./main.go;

echo "Begin docker build...\n";

docker build -t archie .