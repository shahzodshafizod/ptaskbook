<?php
do{
	$a = mt_rand(-50, 50);
	$b = mt_rand(-50, 50);
	$c = mt_rand(-50, 50);
	$d = mt_rand(-50, 50);
}
while((($a!=$b)||($b!=$c))&&(($a!=$b)||($b!=$d))&&(($b!=$c)||($c!=$d))&&(($a!=$c)||($c!=$d)));
echo "Адади якӯм: ".$a;
echo "<br>Адади дуюм: ".$b;
echo "<br>Адади сеюм: ".$c;
echo "<br>Адади чорӯм: ".$d;
if(($a==$b)&&($b==$c)&&($c==$d))
	$index = 0;
else if(($a == $b) && ($b == $c))
	$index = 4;
else if(($a == $b) && ($b == $d))
	$index = 3;
else if(($a == $c) && ($c == $d))
	$index = 2;
else
	$index = 1;
if($index == 0)
	echo "<br>Шарт иҷро нашуд!";
else
	echo "<br>Рақами тартибии адади фарқкунанда: ".$index;
?>