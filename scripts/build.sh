echo "Begin build...\n";

go build -o ../build/archie ./main.go;

echo "Begin docker build...\n";

docker build -t archie .