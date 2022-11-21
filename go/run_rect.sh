clear
go run rectangle-area.go -3 0 3 4 0 -1 9 2
echo "should be 45"
echo
go run rectangle-area.go -2 -2 2 2 -2 -2 2 2
echo "should be 16"

echo
echo "these two rects are nested"
go run rectangle-area.go -2 -2 2 2 -1 -1 1 1
echo "should be 16"