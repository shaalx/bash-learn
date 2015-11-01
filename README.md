# bash-learn

## output redirect

	sh -x run.sh 1>>std.out 2>>err.out
	
	./main.exe  1>>std.out 2>>err.out
	
## for

for i in `seq 1 2 10`;do

echo $i

done



for i in `ls -a`;do

file=`pwd`"/$i"

echo $file

if [ `echo $i|grep .log`];then

echo delete $file ...

rm -rf $file

fi

done



