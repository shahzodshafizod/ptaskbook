<?php
$k = 1;
do{
	$A = mt_rand(10, 50);
	$B = mt_rand(10, 50);
}
while($A >= $B);
echo "Aдади якӯм: ".$A;
echo "<br>Aдади дуюм: ".$B."<br>";
for($i = $A; $i <= $B; $i++){
	for($j = 1; $j <= $k; $j++)
		echo $i." ";
	echo "<br>";
	$k++;
}
?>