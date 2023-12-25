<?php
$N = mt_rand(1, 12);
echo "N = ".$N."<br>";
switch($N){
	case 1:
	case 2:
	case 12:
		echo "Зима";
		break;
	case 3:
	case 4:
	case 5:
		echo "Весна";
		break;
	case 6:
	case 7:
	case 8:
		echo "Лето";
		break;
	case 9:
	case 10:
	case 11:
		echo "Осень";
}
?>

/*
<?php
$N = mt_rand(1, 12);
echo "N = ".$N."<br>";
switch($N){
	case 1:
	case 2:
	case 12:
		echo "Зимистон";
		break;
	case 3:
	case 4:
	case 5:
		echo "Баҳор";
		break;
	case 6:
	case 7:
	case 8:
		echo "Тобистон";
		break;
	case 9:
	case 10:
	case 11:
		echo "Тирамоҳ";
}
?>
*/