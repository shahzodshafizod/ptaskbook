<?php
do{
	$A = mt_rand(1, 100);
	$B = mt_rand(1, 100);
}
while($A <= $B);
$A_baq_B = $A % $B;
echo "Дарозии порчаи якӯм: ".$A;
echo "<br>Дарозии порчаи дуюм: ".$B;
echo "<br>Қисми бандНАбудаи порчаи A: ".$A_baq_B;
?>