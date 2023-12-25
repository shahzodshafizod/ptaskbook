<?php
$X_kg = mt_rand(1, 100) / 4;
$A_somoni = mt_rand(1, 100) / 4;
$Y_kg = mt_rand(1, 100) / 4;
$kg_somoni = $A_somoni / $X_kg;
$Y_somoni = $A_somoni * $Y_kg / $X_kg;
$X_kg = number_format($X_kg, 2);
$A_somoni = number_format($A_somoni, 2);
$Y_kg = number_format($Y_kg, 2);
$kg_somoni = number_format($kg_somoni, 2);
$Y_somoni = number_format($Y_somoni, 2);
echo "Данный вес конфета: ".$X_kg;
echo "<br>Данная стоимость конфета: ".$A_somoni;
echo "<br>Вес для определения стоимости: ".$Y_kg;
echo "<br>".$X_kg." кг = ".$A_somoni." сомони";
echo "<br>1 кг = ".$kg_somoni." сомони";
echo "<br>".$Y_kg." кг = ".$Y_somoni." сомони";
?>

/*
<?php
$X_kg = mt_rand(1, 100) / 4;
$A_somoni = mt_rand(1, 100) / 4;
$Y_kg = mt_rand(1, 100) / 4;
$kg_somoni = $A_somoni / $X_kg;
$Y_somoni = $A_somoni * $Y_kg / $X_kg;
$X_kg = number_format($X_kg, 2);
$A_somoni = number_format($A_somoni, 2);
$Y_kg = number_format($Y_kg, 2);
$kg_somoni = number_format($kg_somoni, 2);
$Y_somoni = number_format($Y_somoni, 2);
echo "Вазни маълумбуда: ".$X_kg;
echo "<br>Вазни маълумбуда: ".$A_somoni;
echo "<br>Вазн барои маълумкунии нарх: ".$Y_kg;
echo "<br>".$X_kg." кг = ".$A_somoni." сомонӣ";
echo "<br>1 кг = ".$kg_somoni." сомонӣ";
echo "<br>".$Y_kg." кг = ".$Y_somoni." сомонӣ";
?>
*/