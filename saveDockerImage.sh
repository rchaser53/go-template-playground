cache='cache_'
images=$(docker images -q)
i=1

for image in $images; do
  docker save $image -o ./cache/$cache$i &
  # echo docker image_$i saved 
  i=`expr $i + 1`
done

# i think it shouldn't use 'wait' for parallel
# wait