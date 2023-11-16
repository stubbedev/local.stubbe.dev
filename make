#!/bin/bash
echo "Go Project Builder"
echo "(1). Run Project"
echo "(2). Build Project"
read  -n 1 -p "Enter desired option number: " GOOPTIONS

echo $GOOPTIONS

if [[ $GOOPTIONS = "1" ]]; then
  go run $PWD/cmd/main/main.go
  echo "Running Application in ./cmd/main/main.go"
fi

if [[ $GOOPTIONS = "2"]]; then
  go build $PWD/cmd/main/main.go -o $PWD/bin/main
  echo "Built Binary Package in ./bin/"
fi

