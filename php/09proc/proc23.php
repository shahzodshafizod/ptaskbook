<?php
function Quarter($x, $y){
	if($x > 0)
		if($y > 0)
			$nat = 1;
		else
			$nat = 4;
	else
		if($y > 0)
			$nat = 2;
		else
			$nat = 3;
	return $nat;
}
$hisob = 1;
while($hisob <= 3){
	do{
		$x = mt_rand(-100, 100);
		$y = mt_rand(-100, 100);
	}
	while(($x == 0) || ($y == 0));
	$chor = Quarter($x, $y);
	echo "Координатаҳои нуқтаи ".$hisob;
	echo "<br>x = ".$x;
	echo "<br>y = ".$y;
	echo "<br>Чоряк: ".$chor."<br><br>";
	$hisob++;
}
?>