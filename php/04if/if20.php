<?php
$A = mt_rand(-100, 100);
$B = mt_rand(-100, 100);
$C = mt_rand(-100, 100);
echo "A = ".$A;
echo "<br>B = ".$B;
echo "<br>C = ".$C;
$AB = abs($A - $B);
$AC = abs($A - $C);
if($AB < $AC){
	echo "<br>Нуқтаи B Ба нуқтаи A наздик ҷойгир аст.";
	echo "<br>Масофаи AB ба ".$AB." баробар аст.";
}
else if($AB > $AC){
	echo "<br>Нуқтаи C Ба нуқтаи A наздик ҷойгир аст.";
	echo "<br>Масофаи AC ба ".$AC." баробар аст.";
}
else
	echo "<br>Ҳар дуи нуқтаҳо ба нуқтаи A дар масофаи якхела ҷойгиранд.";
?>