<?php
$N = mt_rand(1, 4);
$A = mt_rand(-50, 50);
do{
	$B = mt_rand(-50, 50);
}
while($B == 0);
echo "N = ".$N;
echo "<br>A = ".$A;
echo "<br>B = ".$B;
switch($N){
	case 1:
		echo "<br>A + B = ".($A + $B);
		break;
	case 2:
		echo "<br>A - B = ".($A - $B);
		break;
	case 3:
		echo "<br>A * B = ".($A * $B);
		break;
	case 4:
		echo "<br>A / B = ".($A / $B);
}
?>

/*
<?php
$N = mt_rand(1, 4);
$A = mt_rand(-50, 50);
do{
	$B = mt_rand(-50, 50);
}
while($B == 0);
echo "N = ".$N;
echo "<br>A = ".$A;
echo "<br>B = ".$B;
switch($N){
	case 1:
		echo "<br>A + B = ".($A + $B);
		break;
	case 2:
		echo "<br>A - B = ".($A - $B);
		break;
	case 3:
		echo "<br>A * B = ".($A * $B);
		break;
	case 4:
		echo "<br>A / B = ".($A / $B);
}
?>
*/