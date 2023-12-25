<?php
$dar = 1;
$N = mt_rand(1, 50);
$A = mt_rand(1, 50);
echo "Ҳудуди давр: ".$N;
echo "<br>Адади додашуда: ".$A."<br>";
for($i = 1; $i <= $N; $i++){
	$dar *= $A;
	echo $A."^".$i." = ".$dar."<br>";
}
?>