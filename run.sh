

docker build \
--no-cache \
-t go_runner \
.


docker run \
--rm \
-ti \
go_runner

