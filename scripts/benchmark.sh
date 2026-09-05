FILENAME=$(git rev-parse HEAD)

mkdir -p benchmarks
TEST_PROGRAM=../programs/test_all go test -bench=. -count=10 ./... >"benchmarks/$FILENAME"
