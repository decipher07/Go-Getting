result=$(git status | grep ".*\.go$" )

git add . && git commit -m "$result" > /dev/null 2>&1

result=$(git status | grep ".*\.go$" )

git add . && git commit -m "$result" > /dev/null 2>&1
