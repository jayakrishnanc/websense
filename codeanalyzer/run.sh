for entry in "$1"/*
do
  echo "$entry"
  ./codeanalyzer "$entry"
done


