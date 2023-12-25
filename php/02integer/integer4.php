<?php
do{
	$A = mt_rand(1, 100);
	$B = mt_rand(1, 100);
}
while($A <= $B);
$B_dar_A = intval($A / $B);
echo "Дарозии порчаи якӯм: ".$A;
echo "<br>Дарозии порчаи дуюм: ".$B;
echo "<br>Миқдори порчаҳои B, ки дар порчаи A ҷойгиранд: ".$B_dar_A;
?>