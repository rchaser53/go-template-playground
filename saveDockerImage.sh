cache='cache_'
images=$(docker images -q)
i=1

for image in $images; do
  docker save $image -o ./cache/$cache$i
  i=`expr $i + 1`
done