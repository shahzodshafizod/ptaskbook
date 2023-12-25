<?php
do{
	$A = mt_rand(-50, 50);
	$B = mt_rand(-50, 50);
	$C = mt_rand(-50, 50);
}
while(($A != $B) && ($A != $C) && ($B != $C));
echo "Адади якӯм: ".$A;
echo "<br>Адади дуюм: ".$B;
echo "<br>Адади сеюм: ".$C;
if(($A == $B) && ($B == $C))
	$index = 0; //Иҷро нашудани шартҳо
else if($A == $B) //Баробар будани A ва B
	$index = 3;
else if($A == $C) //Баробар будани A ва C
	$index = 2;
else if($B == $C) //Баробар будани C ва B
	$index = 1;
if($index == 0)
	echo "<br>Шарт иҷро нашуд!";
else
	echo "<br>Рақами тартибии адади фарқкунанда: ".$index;
?>