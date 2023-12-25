<?php
do{
	$x = mt_rand(-50, 50);
	$y = mt_rand(-50, 50);
}
while(($x == 0) || ($y == 0));
echo "x = ".$x;
echo "<br>y = ".$y."<br>Чоряк: ";
if($x > 0)
	if($y > 0)
		echo "I";
	else
		echo "IV";
else
	if($y > 0)
		echo "II";
	else
		echo "III";
?>