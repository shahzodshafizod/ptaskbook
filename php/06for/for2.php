<?php
$hisob = 0;
do{
	$A = mt_rand(-100, 100);
	$B = mt_rand(-100, 100);
}
while($A >= $B);
echo "A = ".$A;
echo "<br>B = ".$B."<br><br>";
for($i = $A; $i <= $B; $i++){
	echo $i."<br>";
	$hisob++;
}
echo "Миқдор: ".$hisob;
?>