OUTPUT_DIR=${OUTPUT_DIR:-"_output"}

./${OUTPUT_DIR}/requestgenerator \
 --v=3 \
 --host="http://simple-server" \
 --qps="1" > "/tmp/requestgenerator.log" 2>&1 &
