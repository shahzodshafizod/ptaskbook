<?php
do{
	$A = mt_rand(10, 46);
	$B = mt_rand(10, 46);
}
while($A >= $B);
echo "Aдади якӯм: ".$A;
echo "<br>Aдади дуюм: ".$B."<br>";
for($i = $A; $i <= $B; $i++){
	for($j = 1; $j <= $i; $j++)
			echo $i." ";
	echo "<br>";
}
?>